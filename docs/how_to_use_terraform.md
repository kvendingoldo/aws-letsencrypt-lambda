#### Terraform code example

1. Add module execution to your TF code

```terraform
module "letsencrypt_lambda" {
  source = "git@github.com:kvendingoldo/aws-letsencrypt-lambda.git//files/terraform/module?ref=rc/0.9.0"

  blank_name = "test-letsencrypt-lambda"
  tags       = var.tags

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  image_uri     = var.letsencrypt_lambda_image_uri
  events        = var.letsencrypt_lambda_events
}
```

2. Specify variables

```terraform
variable "tags" {
  default = {
    testTagKey : "testTagValue"
  }
}
variable "letsencrypt_lambda_cron_schedule" {
  default = "rate(168 hours)"
}
variable "letsencrypt_lambda_image_uri" {
  default = "<YOUR_ACCOUNT_ID>.dkr.ecr.us-east-2.amazonaws.com/aws_letsencrypt_lambda:<VERSION>"
}
variable "letsencrypt_lambda_events" {
  default = [
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
}
```