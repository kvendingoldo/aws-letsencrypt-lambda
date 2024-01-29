## How to use it within AWS

1. Apply TF module into your infrastructure via the following commands

```sh
cd files/terraform/module
terraform init
terraform plan -out plan.out
terraform apply -auto-approve plan.out
```

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
