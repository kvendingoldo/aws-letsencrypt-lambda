module "letsencrypt_lambda" {
  source = "../../"

  blank_name = "test-letsencrypt-lambda"
  tags       = var.tags

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  events        = var.letsencrypt_lambda_events

  ecr_proxy_enabled = false
  ecr_image_uri     = "<YOUR_ACCOUNT_ID>.dkr.ecr.us-east-2.amazonaws.com/aws_letsencrypt_lambda:<VERSION>"
}
