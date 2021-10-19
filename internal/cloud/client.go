package cloud

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/route53"
)

type Client struct {
	ACMClient     *acm.Client
	Route53Client *route53.Client
	Region        string
}

func New(ctx context.Context, region string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	return &Client{
		ACMClient:     acm.NewFromConfig(cfg),
		Route53Client: route53.NewFromConfig(cfg),
		Region:        region,
	}, nil
}
