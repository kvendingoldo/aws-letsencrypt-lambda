module "iam_assumable_role" {
  source = "terraform-aws-modules/iam/aws//modules/iam-assumable-role"

  trusted_role_arns = [
    "arn:aws:iam::004867756392:root"
  ]
  trusted_role_services = [
    "lambda.amazonaws.com"
  ]
  create_role = true

  role_name         = "example-certs-use2"
  role_requires_mfa = false

  custom_role_policy_arns = [
    "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess",
    "arn:aws:iam::aws:policy/AmazonRoute53FullAccess",
    "arn:aws:iam::aws:policy/AWSCertificateManagerFullAccess"
  ]
  number_of_custom_role_policy_arns = 3
}

module "certs_renew_lambda" {
  source = "../../"

  blank_name = "example-certs-use2"
  tags       = {}

  cron_schedule = var.certs_renew_lambda_cron_schedule
  events        = var.certs_renew_lambda_events

  iam_role_arn = module.iam_assumable_role.iam_role_arn

  ecr_proxy_enabled               = true
  ecr_proxy_upstream_registry_url = "registry-1.docker.io"
  ecr_proxy_username              = "yourDockerHubUsername"
  ecr_proxy_access_token          = "yourDockerHubAccessToken"
}
