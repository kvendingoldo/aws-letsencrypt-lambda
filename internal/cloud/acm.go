package cloud
//
//import (
//	"context"
//	"fmt"
//	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/service/acm/types"
//	"github.com/aws/aws-sdk-go/service/acm"
//)
//
//func ListCertificateSummaries(ctx context.Context, api ACMListCertificatesAPI) ([]acmTypes.CertificateSummary, error) {
//	in := acm.ListCertificatesInput{}
//	out, err := api.ListCertificates(ctx, &in)
//	if err != nil {
//		return nil, err
//	}
//
//	return out.CertificateSummaryList, nil
//}
//
//func ListCertificates(ctx context.Context, api ACMAPI) ([]Certificate, error) {
//	summary, err := ListCertificateSummaries(ctx, api)
//	if err != nil {
//		return nil, err
//	}
//
//	var cList []Certificate
//	for _, s := range summary {
//		c, err := GetCertificate(ctx, api, aws.ToString(s.CertificateArn))
//		if err != nil {
//			fmt.Println(err.Error())
//			continue
//		}
//		cList = append(cList, c)
//	}
//
//	return cList, nil
//}
//
//func GetCertificate(ctx context.Context, api ACMDescribeCertificateAPI, arn string) (Certificate, error) {
//	in := acm.DescribeCertificateInput{
//		CertificateArn: aws.String(arn),
//	}
//	out, err := api.DescribeCertificate(ctx, &in)
//	if err != nil {
//		return Certificate{}, err
//	}
//
//	vMethod := ""
//	recordSet := RecordSet{}
//	if out.Certificate.DomainValidationOptions != nil {
//		vMethod = string(out.Certificate.DomainValidationOptions[0].ValidationMethod)
//		if vMethod == string(types.ValidationMethodDns) {
//			recordSet.HostedDomainName = aws.ToString(out.Certificate.DomainValidationOptions[0].ValidationDomain)
//			recordSet.Name = aws.ToString(out.Certificate.DomainValidationOptions[0].ResourceRecord.Name)
//			recordSet.Value = aws.ToString(out.Certificate.DomainValidationOptions[0].ResourceRecord.Value)
//			recordSet.Type = string(out.Certificate.DomainValidationOptions[0].ResourceRecord.Type)
//		}
//	}
//
//	return Certificate{
//		Arn:                 arn,
//		DomainName:          aws.ToString(out.Certificate.DomainName),
//		Status:              string(out.Certificate.Status),
//		Type:                string(out.Certificate.Type),
//		FailureReason:       string(out.Certificate.FailureReason),
//		ValidationMethod:    vMethod,
//		ValidationRecordSet: recordSet,
//	}, nil
//}
