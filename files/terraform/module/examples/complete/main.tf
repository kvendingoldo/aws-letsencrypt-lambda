module "certs_renew_lambda_use2" {
  source = "git::https://github.com/kvendingoldo/aws-letsencrypt-lambda.git//files/terraform/module?ref=0.11.0"

  blank_name = "example-certs-use2"
  tags       = {}

  cron_schedule = var.certs_renew_lambda_cron_schedule
  events        = var.certs_renew_lambda_events_use2

  create_iam_role = true

  ecr_proxy_enabled               = true
  ecr_proxy_upstream_registry_url = "registry-1.docker.io"
  ecr_proxy_username              = "yourDockerHubUsername"
  ecr_proxy_access_token          = "yourDockerHubAccessToken"
}

module "certs_renew_lambda_use1" {
  providers = {
    aws = aws.use1
  }

  source = "git::https://github.com/kvendingoldo/aws-letsencrypt-lambda.git//files/terraform/module?ref=0.11.0"

  blank_name = "example-certs-use1"
  tags       = {}

  cron_schedule = var.certs_renew_lambda_cron_schedule
  image         = "004867756392.dkr.ecr.us-east-1.amazonaws.com/aws_letsencrypt_lambda:0.11.0"
  events        = var.certs_renew_lambda_events_use1

  create_iam_role = true
}
