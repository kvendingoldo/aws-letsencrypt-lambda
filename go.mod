module github.com/kvendingoldo/aws-letsencrypt-lambda

go 1.22

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2 v1.27.0
	github.com/aws/aws-sdk-go-v2/config v1.27.16
	github.com/aws/aws-sdk-go-v2/service/acm v1.26.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.40.8
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.29.1
	github.com/go-acme/lego v2.7.2+incompatible
	github.com/go-acme/lego/v4 v4.17.3
	github.com/guregu/null v4.0.0+incompatible
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/aws/aws-sdk-go v1.53.14 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.16 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.7 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.24.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.10 // indirect
	github.com/aws/smithy-go v1.20.2 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/go-jose/go-jose/v4 v4.0.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/miekg/dns v1.1.59 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
