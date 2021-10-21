## How to use it within AWS

1. Lambda image should be pulled from docker hub and pushed into your personal ECR repository; AWS Lambda is not able to
   work with any other docker registry except ECR.
2. Apply TF module into your infrastructure

#### Example of input variables for Terraform module

```terraform
blank_name = "aws-letsencrypt-lambda"
image_uri  = "<PATH_TO_IMAGE_IN_ECR>"
events     = [
  {
    "DomainName" : "<TEST_DOMAIN_1>",
    "AcmeUrl" : "stage",
    "AcmeEmail" : "<EMAIL_1>",
    "ReImportThreshold" : 10,
    "IssueType" : "force"
  },
  {
    "DomainName" : "<TEST_DOMAIN_2>",
    "AcmeUrl" : "prod",
    "AcmeEmail" : "<EMAIL_2>",
    "ReImportThreshold" : 30,
    "IssueType" : "default"
  }
]
```

### How to trigger lambda manually via UI

1. Go to Lambda function that has been created via Terraform -> Tests
2. Fill "Test Event" and click "Test"

```json
{
  "domain_name": "<TEST_DOMAIN_3>",
  "acme_url": "stage",
  "acme_email": "<EMAIL_2>",
  "reimport_threshold": 10,
  "issue_type": "default"
}
```

## How to use it locally
1. Set the following environment variables (do not forget to change placeholders)
```shell
export AWS_REGION=<REGION>
export MODE=local
export DOMAIN_NAME=<TEST_DOMAIN_4>
export ACME_URL="stage"
export ACME_EMAIL=<TEST_EMAIL_4>
export REIMPORT_THRESHOLD=10
export ISSUE_TYPE="default"
```
2. Run lambda locally
```sh
go run main.go
```

## Environment variables

* AWS_REGION
    * Description: AWS Region. Inside of Lambda it's setting automatically by Lambda
    * Possible values:

* MODE
    * Description: mode of application running
    * Possible values: cloud | local

* DOMAIN_NAME
    * Description: Name of domain for which certificate will be issued/renewed
    * Possible values: *any valid domain name*

* ACME_URL
    * Description: If prod then *production* LE URL will be used, otherwise *stage* URL will be used
    * Possible values: prod | stage

* ACME_EMAIL
    * Description: Email that will be associated with LE certificate
    * Possible values: *any valid email*

* REIMPORT_THRESHOLD
    * Description: If TTL of cert == REIMPORT_THRESHOLD then cert will be renewed
    * Possible values: *any int > 0*

