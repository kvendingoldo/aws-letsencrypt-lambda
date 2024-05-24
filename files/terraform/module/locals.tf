locals {
  events = {for event in var.events : event["domainName"] => event if var.cron_enabled}

  ecr_domain   = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com"
  image_prefix = var.enable_ecr_proxy ? "${local.ecr_domain}/${var.ecr_repository_prefix}/" : ""
  lambda_image = var.enable_ecr_proxy ? "${local.image_prefix}${var.dockerhub_image}" : var.ecr_image_uri
}
