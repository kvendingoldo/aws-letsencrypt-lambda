locals {
  events = {for event in var.events : event["domainName"] => event if var.cron_enabled}
}