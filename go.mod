module github.com/kvendingoldo/aws-letsencrypt-lambda

go 1.17

require (
	github.com/aws/aws-lambda-go v1.27.0
	github.com/aws/aws-sdk-go-v2 v1.9.2
	github.com/aws/aws-sdk-go-v2/config v1.8.3
	github.com/aws/aws-sdk-go-v2/service/acm v1.6.2
	github.com/aws/aws-sdk-go-v2/service/route53 v1.11.2
	github.com/go-acme/lego v2.7.2+incompatible
	github.com/go-acme/lego/v4 v4.5.3
	github.com/guregu/null v4.0.0+incompatible
	github.com/sirupsen/logrus v1.8.1
)

require (
	github.com/aws/aws-sdk-go v1.41.5 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.4.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.6.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.3.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.4.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.2 // indirect
	github.com/aws/smithy-go v1.8.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/miekg/dns v1.1.43 // indirect
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)
