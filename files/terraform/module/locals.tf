locals {
  events = { for event in var.events : event["domainName"] => event if var.cron_enabled }

  ecr_domain   = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com"
  image_prefix = var.ecr_proxy_enabled ? "${local.ecr_domain}/${var.ecr_proxy_repository_prefix}/" : ""
  lambda_image = var.ecr_proxy_enabled ? "${local.image_prefix}${var.public_image}" : var.ecr_image_uri
}
