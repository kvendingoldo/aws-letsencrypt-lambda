## How to use it within AWS

1. Lambda image should be pulled from docker hub and pushed into your personal ECR repository; AWS Lambda is not able to
   work with any other docker registry except ECR.
2. Apply TF module into your infrastructure

### How to trigger lambda manually via UI

1. Go to Lambda function that has been created via Terraform -> Tests
2. Fill "Test Event" and click "Test"

```
{
  "domainName": "<YOUR_VALID_DOMAIN>",
  "acmeUrl": "stage",
  "acmeEmail": "<ANY_VALID_EMAIL>",
  "reImportThreshold": 10,
  "issueType": "<default | force>"
}
```

#### Example #1:

```json
{
   "domainName": "mypersonaldomain.com",
   "acmeUrl": "stage",
   "acmeEmail": "mypersonal@gmail.com",
   "reImportThreshold": 10,
   "issueType": "default"
}
```