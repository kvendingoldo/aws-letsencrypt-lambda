package cloud

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

func getZoneIdByName(ctx context.Context, client Client, domainName string) (string, error) {
	params := route53.ListHostedZonesByNameInput{
		MaxItems: aws.Int32(100),
	}
	out, err := client.Route53Client.ListHostedZonesByName(ctx, &params)
	if err != nil {
		return "", err
	}

	for _, zone := range out.HostedZones {
		if domainName == *zone.Name || fmt.Sprintf("%s.", domainName) == *zone.Name {
			return *zone.Id, nil
		}
	}

	return "", errors.New("Not found")
}

func ChangeRecord(ctx context.Context, client Client, changeType types.ChangeAction, domainName, recordName, recordValue string) error {
	zoneId, err := getZoneIdByName(ctx, client, domainName)

	if err != nil {
		return err
	}

	params := route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(zoneId),
		ChangeBatch: &types.ChangeBatch{
			Changes: []types.Change{
				{
					Action: changeType,
					ResourceRecordSet: &types.ResourceRecordSet{
						Name: aws.String(fmt.Sprintf("%v.%v", recordName, domainName)),
						Type: types.RRTypeTxt,
						TTL:  aws.Int64(60),
						ResourceRecords: []types.ResourceRecord{
							{
								Value: aws.String(recordValue),
							},
						},
					},
				},
			},
		},
	}

	_, err = client.Route53Client.ChangeResourceRecordSets(ctx, &params)
	if err != nil {
		return err
	}

	return nil
}
