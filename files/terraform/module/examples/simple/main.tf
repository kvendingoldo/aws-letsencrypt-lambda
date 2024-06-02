module "letsencrypt_lambda" {
  source = "../../"

  blank_name = "test-letsencrypt-lambda"
  tags       = var.tags

  cron_schedule = var.letsencrypt_lambda_cron_schedule
  events        = var.letsencrypt_lambda_events

  ecr_proxy_username     = "myusername"
  ecr_proxy_access_token = "ghp_XXX"
}
