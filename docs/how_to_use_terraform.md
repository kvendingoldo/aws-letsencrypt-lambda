# Configuration

## Docker image configuration

AWS Lambda does not provide an ability to use public docker images easily. To use the Terraform automation you have two options:

1. Use ECR image pull cache (**preferred way**) 
   1. Set Terraform variable `ecr_proxy_enabled=true`
   2. Set Terraform variable `ecr_proxy_upstream_registry_url` to `registry-1.docker.io` or `ghcr.io`.
   3. Create Access Token for GitHub Package registry or Docker Hub.
   4. Set Terraform variable `ecr_proxy_username` to your username
   5. Set Terraform variable `ecr_proxy_access_token` to your access token

2. Use your own AWS ECR
   1. Pull kvendingoldo's image from [Docker Hub](https://hub.docker.com/repository/docker/kvendingoldo/aws-letsencrypt-lambda) / [GitHub registry](https://github.com/kvendingoldo?tab=packages&repo_name=aws-letsencrypt-lambda).
   2. Create your own private AWS ECR repository
   3. Retag pulled image and push it to your private ECR repository. 
   4. Change `var.image` to your image URL. E.g.: `image = "004867756392.dkr.ecr.us-east-1.amazonaws.com/aws_letsencrypt_lambda:0.11.0"`


## Terraform examples

To get more examples, explore [`examples`](https://github.com/kvendingoldo/aws-letsencrypt-lambda/tree/main/files/terraform/module/examples) folder.

### Step-by-step example

1. Add lambda module to your TF code

```terraform
module "letsencrypt_lambda" {
  source = "git@github.com:kvendingoldo/aws-letsencrypt-lambda.git//files/terraform/module?ref=0.11.0"

  blank_name = "test-letsencrypt-lambda"
  tags       = {
    testTagKey : "testTagValue"
  }

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  events        = var.letsencrypt_lambda_events
}
```

2. Specify variables

```terraform
variable "letsencrypt_lambda_cron_schedule" {
  default = "rate(168 hours)"
}
variable "letsencrypt_lambda_events" {
  default = [
    {
      "domainName" : "<TEST_DOMAIN_1>",
      "acmeUrl" : "stage",
      "acmeEmail" : "<EMAIL_1>",
      "reImportThreshold" : 10,
      "issueType" : "force"
    },
    {
      "acmRegion" : "us-east-2",
      "route53Region" : "us-east-1",
      "domainName" : "<TEST_DOMAIN_2>",
      "acmeUrl" : "prod",
      "acmeEmail" : "<EMAIL_2>",
      "reImportThreshold" : 30,
      "issueType" : "default"
    }
  ]
}
```
