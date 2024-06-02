package cloud

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Client struct {
	ACMClient            *acm.Client
	Route53Client        *route53.Client
	SecretsManagerClient *secretsmanager.Client
}

func New(ctx context.Context, route53Region, acmRegion, secretsManagerRegion string) (*Client, error) {
	acmCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(acmRegion))
	if err != nil {
		return nil, err
	}

	route53Cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(route53Region))
	if err != nil {
		return nil, err
	}

	secretsManagerCfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(secretsManagerRegion))
	if err != nil {
		return nil, err
	}

	return &Client{
		ACMClient:            acm.NewFromConfig(acmCfg),
		Route53Client:        route53.NewFromConfig(route53Cfg),
		SecretsManagerClient: secretsmanager.NewFromConfig(secretsManagerCfg),
	}, nil
}
