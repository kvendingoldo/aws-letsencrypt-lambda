locals {
  events = { for event in var.events : event["domainName"] => event if var.cron_enabled }
  image  = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com/${var.ecr_repository_prefix}/${var.image}"
}
