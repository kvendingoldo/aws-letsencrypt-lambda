## How to use it locally

1. Look at [environments variables](environment_variables.md) and set at least required variables

### example:
```shell
export AWS_REGION="us-east-2"
export MODE=local
export DOMAIN_NAME="t22est.com"
export ACME_URL="stage"
export ACME_EMAIL="mytestemail@gmail.com"
export REIMPORT_THRESHOLD=10
export ISSUE_TYPE="default"
```

2. Run lambda locally

```sh
go run main.go
```

