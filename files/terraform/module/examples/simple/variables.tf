variable "tags" {
  default = {
    testTagKey : "testTagValue"
  }
}

variable "ecr_proxy_username" {
  default = "kvendingoldo"
}

variable "ecr_proxy_access_token" {
  default = "ghp_xxx"
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
      "reImportThreshold" : 100,
      "issueType" : "default",
      "storeCertInSM" : true
    }
  ]
}
