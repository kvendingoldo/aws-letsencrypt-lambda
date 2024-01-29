variable "certs_renew_lambda_cron_schedule" {
  default = "rate(48 hours)"
}

variable "certs_renew_lambda_events" {
  description = "List of events for Lambda function (each event contains info about one certificate)"
  default = [
    {
      domainName : "dev.referrs.me",
      acmeUrl : "prod",
      acmeEmail : "alex.sharov@referrs.me",
      reImportThreshold : 10,
      issueType : "default"
    }
  ]
}
