package lambda

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	acmTypes "github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/cloud"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/types"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/utils"
	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	secretsManagerTypes "github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

func getCertificateByDomainFromSlice(domain string, certificates *acm.ListCertificatesOutput) (acmTypes.CertificateSummary, error) {
	for _, crt := range certificates.CertificateSummaryList {
		if domain == *crt.DomainName {
			return crt, nil
		}
	}

	return acmTypes.CertificateSummary{}, errors.New("not found")
}

func showCertificateInfo(certificate acmTypes.CertificateDetail) {
	log.Infof("Checking certificate for domain '%v' with arn '%v'", *certificate.DomainName, *certificate.CertificateArn)
	log.Infof("Certificate status is '%v'", certificate.Status)
	log.Infof("Certificate in use by %v", certificate.InUseBy)

	notAfterDate := certificate.NotAfter
	certificateDaysLeft := int(notAfterDate.Sub(time.Now()).Hours() / 24)
	log.Infof("Certificate valid until %v (%v days left)", notAfterDate, certificateDaysLeft)
}

func importCertificate(ctx context.Context, client *cloud.Client, arn *string, tlsCertificates *certificate.Resource, reimport bool) error {
	// NOTE: The first one certificate in tlsCertificates.Certificate is certBody,
	// the whole tlsCertificates.Certificate is certChain
	crts := utils.SplitStringsByEmptyNewline(string(tlsCertificates.Certificate))

	params := &acm.ImportCertificateInput{
		Certificate:      []byte(crts[0]),
		PrivateKey:       tlsCertificates.PrivateKey,
		CertificateChain: tlsCertificates.Certificate,
	}

	if !reimport {
		params.Tags = []acmTypes.Tag{
			{
				Key:   aws.String(lambdaAwsTag),
				Value: aws.String("true"),
			},
		}
	}

	if arn != nil {
		params.CertificateArn = arn
	}

	output, err := client.ACMClient.ImportCertificate(ctx, params)
	if err != nil {
		return err
	}

	log.Infof("Certificate has been successfully imported. Arn is %v", *output.CertificateArn)

	return nil
}

func uploadToSecretManager(ctx context.Context, client *cloud.Client, domainName string, tlsCertificates *certificate.Resource) error {
	secretName := domainName

	// Convert the certificate resource to JSON
	certData := types.SecretsManagerResource{
		Domain:            domainName,
		PrivateKey:        string(tlsCertificates.PrivateKey),
		Certificate:       string(tlsCertificates.Certificate),
		IssuerCertificate: string(tlsCertificates.IssuerCertificate),
		CSR:               string(tlsCertificates.CSR),
	}
	certJSON, err := json.Marshal(certData)
	if err != nil {
		return fmt.Errorf("failed to marshal certificate resource: %w", err)
	}

	// Prepare the secret value
	secretString := aws.String(string(certJSON))

	// Check if the secret exists
	_, err = client.SecretsManagerClient.DescribeSecret(ctx, &secretsmanager.DescribeSecretInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		// If the secret doesn't exist, create it
		var notFound *secretsManagerTypes.ResourceNotFoundException
		if err != nil && !errors.As(err, &notFound) {
			return fmt.Errorf("failed to describe secret: %w", err)
		}

		secret, err := client.SecretsManagerClient.CreateSecret(ctx, &secretsmanager.CreateSecretInput{
			Name:         aws.String(secretName),
			SecretString: secretString,
		})
		if err != nil {
			return fmt.Errorf("failed to create secret: %w", err)
		}
		log.Infof("Secret created successfully. SecretId: %s", *secret.ARN)
	} else {
		// If the secret exists, update it
		secret, err := client.SecretsManagerClient.UpdateSecret(ctx, &secretsmanager.UpdateSecretInput{
			SecretId:     aws.String(secretName),
			SecretString: secretString,
		})
		if err != nil {
			return fmt.Errorf("failed to update secret: %w", err)
		}
		log.Infof("Secret updated successfully. SecretId: %s", *secret.ARN)
	}

	return nil
}

func processCertificate(ctx context.Context, config config.Config, client *cloud.Client, certificate acmTypes.CertificateSummary) error {
	tags, err := client.ACMClient.ListTagsForCertificate(
		ctx,
		&acm.ListTagsForCertificateInput{
			CertificateArn: certificate.CertificateArn,
		},
	)
	if err != nil {
		return err
	}

	isAutomationEnabled := false
	for _, tag := range tags.Tags {
		if *tag.Key == lambdaAwsTag {
			isAutomationEnabled = true

			break
		}
	}

	if !isAutomationEnabled {
		log.Infof("Certificate '%v' is out of the scope of Lambda automation. Re-create it via Lambda or add '%v' tag to certificate", *certificate.CertificateArn, lambdaAwsTag)
	}

	info, err := client.ACMClient.DescribeCertificate(
		ctx,
		&acm.DescribeCertificateInput{
			CertificateArn: certificate.CertificateArn,
		},
	)
	if err != nil {
		return err
	}

	showCertificateInfo(*info.Certificate)
	certificateDaysLeft := int64(info.Certificate.NotAfter.Sub(time.Now()).Hours() / 24)

	//nolint:gocritic
	if (certificateDaysLeft <= config.ReImportThreshold) || (config.IssueType == "force") || ((*info.Certificate).Status == acmTypes.CertificateStatusExpired) {
		if config.IssueType == "force" {
			log.Info("IssueType == force, certificate will be recreated")
		}

		tlsCertificates, err := utils.GetCertificates(config, *info.Certificate.DomainName)
		if err != nil {
			return err
		}

		err = importCertificate(ctx, client, certificate.CertificateArn, tlsCertificates, true)
		if err != nil {
			return err
		}

		err = uploadToSecretManager(ctx, client, config.DomainName, tlsCertificates)
		if err != nil {
			return err
		}

	} else {
		log.Infof("No re-import needed. It has to be done 10 days before expiration")
	}

	return nil
}

func Execute(ctx context.Context, config config.Config) error {
	client, err := cloud.New(ctx, config.ACMRegion, config.Route53Region, config.SecretsManagerRegion)
	if err != nil {
		//nolint:stylecheck
		return fmt.Errorf("Could not create AWS client. Error: %w", err)
	}

	params := &acm.ListCertificatesInput{
		CertificateStatuses: []acmTypes.CertificateStatus{acmTypes.CertificateStatusIssued},
		MaxItems:            aws.Int32(100),
	}
	certificates, err := client.ACMClient.ListCertificates(ctx, params)
	if err != nil {
		//nolint:stylecheck
		return fmt.Errorf("Could not get list of AWS certificates. Error: %w", err)
	}

	crt, err := getCertificateByDomainFromSlice(config.DomainName, certificates)
	if err != nil {
		log.Warnf("Certificate not found for %v domain; Trying to create ...", config.DomainName)

		tlsCertificates, err := utils.GetCertificates(config, config.DomainName)
		if err != nil {
			//nolint:stylecheck
			return fmt.Errorf("Failed to issue certificate. Error: %w", err)
		}

		err = importCertificate(ctx, client, nil, tlsCertificates, false)
		if err != nil {
			//nolint:stylecheck
			return fmt.Errorf("Failed to import certificate. Error: %w", err)
		}
	} else {
		log.Infof("Certificate found, arn is %v. Trying to renew ...", *crt.CertificateArn)
		err := processCertificate(ctx, config, client, crt)
		if err != nil {
			//nolint:stylecheck
			return fmt.Errorf("Failed to process certificate. Error: %w", err)
		}
	}

	return nil
}
