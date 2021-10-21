locals {
  events = {for event in var.events : event["DomainName"] => event if var.cron_enabled}
}