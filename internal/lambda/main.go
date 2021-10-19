package lambda

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	acmTypes "github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/go-acme/lego/v4/certificate"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/cloud"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/config"
	"github.com/kvendingoldo/aws-letsencrypt-lambda/internal/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func getCertificateByDomainFromSlice(domain string, certificates *acm.ListCertificatesOutput) (acmTypes.CertificateSummary, error) {
	for _, crt := range certificates.CertificateSummaryList {
		if domain == *crt.DomainName {
			return crt, nil
		}
	}

	return acmTypes.CertificateSummary{}, errors.New("Not found")
}

func showCertificateInfo(certificate acmTypes.CertificateDetail) {
	log.Infof("Checking certificate for domain '%v' with arn '%v'", *certificate.DomainName, *certificate.CertificateArn)
	log.Infof("Certificate status is '%v'", certificate.Status)
	log.Infof("Certificate in use by %v", certificate.InUseBy)

	notAfterDate := certificate.NotAfter
	certificateDaysLeft := int(notAfterDate.Sub(time.Now()).Hours() / 24)
	log.Infof("Certificate valid untill %v (%v days left)", notAfterDate, certificateDaysLeft)
}

func importCertificate(ctx context.Context, client *cloud.Client, arn *string, tlsCertificates *certificate.Resource) error {
	fmt.Println("===")
	fmt.Println(tlsCertificates)
	fmt.Println("===")
	fmt.Println(string(tlsCertificates.Certificate))
	fmt.Println("===")
	fmt.Println(string(tlsCertificates.PrivateKey))

	crts, err := x509.ParseCertificates(tlsCertificates.Certificate)
	if err != nil {
		return err
	}

	// TODO: delete this code
	fmt.Println("===")
	fmt.Println(string(tlsCertificates.Certificate))
	fmt.Println("===")
	fmt.Println(string(crts[0].Raw))
	fmt.Println("===")
	fmt.Println(string(tlsCertificates.PrivateKey))

	// NOTE: The first one certificate in tlsCertificates.Certificate is certBody,
	// the whole tlsCertificates.Certificate is certChain
	params := &acm.ImportCertificateInput{
		Certificate:      crts[0].Raw,
		PrivateKey:       tlsCertificates.PrivateKey,
		CertificateChain: tlsCertificates.Certificate,
		Tags: []types.Tag{
			{
				Key:   aws.String("issue-date"),
				Value: aws.String(time.Now().String()),
			},
		},
	}

	if arn != nil {
		params.CertificateArn = arn
	}

	output, err := client.ACMClient.ImportCertificate(ctx, params)
	if err != nil {
		return err
	}

	// todo: useless prints
	fmt.Println(params)
	fmt.Println(output)
	return nil
}

func processCertificate(ctx context.Context, config config.Config, client *cloud.Client, certificate acmTypes.CertificateSummary) error {
	info, _ := client.ACMClient.DescribeCertificate(
		ctx,
		&acm.DescribeCertificateInput{
			CertificateArn: certificate.CertificateArn,
		},
	)

	showCertificateInfo(*info.Certificate)
	certificateDaysLeft := int(info.Certificate.NotAfter.Sub(time.Now()).Hours() / 24)

	if certificateDaysLeft <= config.ReImportThreshold {
		tlsCertificates, err := utils.GetCertificates(config, *info.Certificate.DomainName)
		if err != nil {
			fmt.Println(err)
			return err
		}

		err = importCertificate(ctx, client, certificate.CertificateArn, tlsCertificates)
		if err != nil {
			return err
		}
	} else {
		log.Infof("No re-import needed. It has to be done 10 days before expiration")
	}

	return nil
}

func Execute(config config.Config) {
	client, err := cloud.New(context.TODO(), config.Region)
	if err != nil {
		log.Error(fmt.Sprintf("Could not create AWS client"), "error", err)
		os.Exit(1)
	}

	ctx := context.TODO()

	params := &acm.ListCertificatesInput{
		CertificateStatuses: []acmTypes.CertificateStatus{acmTypes.CertificateStatusIssued},
		MaxItems:            aws.Int32(100),
	}
	certificates, err := client.ACMClient.ListCertificates(ctx, params)
	if err != nil {
		log.Error(fmt.Sprintf("Could not get list of AWS certificates"), "error", err)
		os.Exit(1)
	}

	if config.DomainOnly {
		crt, err := getCertificateByDomainFromSlice(config.DomainName, certificates)
		if err != nil {
			log.Warn("Certificate not found; Trying to create ...")

			tlsCertificates, err := utils.GetCertificates(config, config.DomainName)
			if err != nil {
				log.Error("Failed to issue certificate")
				log.Error("error", err)
				os.Exit(1)
			}

			err = importCertificate(ctx, client, nil, tlsCertificates)
			if err != nil {
				log.Error("Failed to import certificate")
				log.Error("error", err)
				os.Exit(1)
			}

		} else {
			err := processCertificate(ctx, config, client, crt)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to proccess certificate\n"), "error", err)
				os.Exit(1)
			}
		}
	} else {
		for _, crt := range certificates.CertificateSummaryList {
			err := processCertificate(ctx, config, client, crt)
			if err != nil {
				log.Error(fmt.Sprintf("Failed to proccess certificate"), "error", err)
				os.Exit(1)
			}
		}
	}
}
