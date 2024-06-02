module "letsencrypt_lambda" {
  source = "../../"

  blank_name = "test-letsencrypt-lambda"
  tags       = var.tags

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  events        = var.letsencrypt_lambda_events

  ecr_proxy_username     = var.ecr_proxy_username
  ecr_proxy_access_token = var.ecr_proxy_access_token
}
