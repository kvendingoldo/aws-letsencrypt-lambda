## How to use it within AWS

1. Lambda image should be pulled from docker hub and pushed into your personal ECR repository; AWS Lambda is not able to
   work with any other docker registry except ECR.
2. Apply TF module into your infrastructure

### How to trigger lambda manually via UI

1. Go to Lambda function that has been created via Terraform -> Tests
2. Fill "Test Event" and click "Test"

```
{
  "domain_name": "<TEST_DOMAIN_3>",
  "acme_url": "stage",
  "acme_email": "<EMAIL_2>",
  "reimport_threshold": 10,
  "issue_type": "default"
}
```

#### Example #1:

```json
{
   "domain_name": "t22est.com",
   "acme_url": "stage",
   "acme_email": "mytestemail@gmail.com",
   "reimport_threshold": 10,
   "issue_type": "default"
}
```