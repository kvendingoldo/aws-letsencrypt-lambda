module "letsencrypt_lambda" {
  source = "git@github.com:kvendingoldo/aws-letsencrypt-lambda.git//files/terraform/module?ref=rc/0.9.0"

  blank_name = "test-letsencrypt-lambda"
  tags       = var.tags

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  image_uri     = var.letsencrypt_lambda_image_uri
  events        = var.letsencrypt_lambda_events

  enable_ecr_proxy = false
  ecr_image_uri    = "<YOUR_ACCOUNT_ID>.dkr.ecr.us-east-2.amazonaws.com/aws_letsencrypt_lambda:<VERSION>"
}
