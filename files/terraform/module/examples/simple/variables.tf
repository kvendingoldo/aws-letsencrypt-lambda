variable "tags" {
  default = {
    testTagKey : "testTagValue"
  }
}
variable "letsencrypt_lambda_cron_schedule" {
  default = "rate(168 hours)"
}
variable "letsencrypt_lambda_events" {
  default = [
    {
      "acmRegion" : "us-east-1",
      "route53Region" : "us-east-1",
      "domainName" : "hackernoon.referrs.me",
      "acmeUrl" : "stage",
      "acmeEmail" : "alex.sharov@referrs.me",
      "reImportThreshold" : 10,
      "issueType" : "default"
    }
  ]
}
