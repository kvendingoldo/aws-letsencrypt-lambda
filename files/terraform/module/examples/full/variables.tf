variable "tags" {
  default = {
    testTagKey : "testTagValue"
  }
}
variable "letsencrypt_lambda_cron_schedule" {
  default = "rate(168 hours)"
}
variable "letsencrypt_lambda_image_uri" {
  default = "<YOUR_ACCOUNT_ID>.dkr.ecr.us-east-2.amazonaws.com/aws_letsencrypt_lambda:<VERSION>"
}
variable "letsencrypt_lambda_events" {
  default = [
    {
      "domainName" : "<TEST_DOMAIN_1>",
      "acmeUrl" : "stage",
      "acmeEmail" : "<EMAIL_1>",
      "reImportThreshold" : 10,
      "issueType" : "force",
      "storeCertInSM" : false
    },
    {
      "acmRegion" : "us-east-2",
      "route53Region" : "us-east-1",
      "domainName" : "<TEST_DOMAIN_2>",
      "acmeUrl" : "prod",
      "acmeEmail" : "<EMAIL_2>",
      "reImportThreshold" : 30,
      "issueType" : "default",
      "storeCertInSM" : "true"
    }
  ]
}
