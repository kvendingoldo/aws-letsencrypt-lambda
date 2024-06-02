# aws-letsencrypt-lambda Terraform module

OpenTofu module which creates aws-letsencrypt-lambda resources.

## Examples

Examples codified under
the [`examples`](https://github.com/kvendingoldo/aws-letsencrypt-lambda/tree/main/files/terraform/module/examples) are intended
to give users references for how to use the module(s) as well as testing/validating changes to the source code of the
module. If contributing to the project, please be sure to make any appropriate updates to the relevant examples to allow
maintainers to test your changes and to keep the examples up to date for users. Thank you!

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudwatch_event_rule.schedule](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_rule) | resource |
| [aws_cloudwatch_event_target.event_target](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_event_target) | resource |
| [aws_cloudwatch_log_group.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudwatch_log_group) | resource |
| [aws_ecr_pull_through_cache_rule.lambda_proxy](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecr_pull_through_cache_rule) | resource |
| [aws_iam_policy.acm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_policy.logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_policy.route53](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_policy.secretsmanager](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_policy) | resource |
| [aws_iam_role.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role) | resource |
| [aws_iam_role_policy_attachment.acm](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_role_policy_attachment.logging](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_role_policy_attachment.route53](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_role_policy_attachment.secretsmanager](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_iam_role_policy_attachment.vpc_permissions](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/iam_role_policy_attachment) | resource |
| [aws_lambda_function.main](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function) | resource |
| [aws_lambda_permission.allow_cloudwatch](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission) | resource |
| [aws_secretsmanager_secret.image_registry](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret) | resource |
| [aws_secretsmanager_secret_version.image_registry](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/secretsmanager_secret_version) | resource |
| [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
| [aws_region.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/region) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_blank_name"></a> [blank\_name](#input\_blank\_name) | Blank name for AWS resources | `string` | n/a | yes |
| <a name="input_cloudwatch_log_group_retention"></a> [cloudwatch\_log\_group\_retention](#input\_cloudwatch\_log\_group\_retention) | Number of days to retain log events in the specified cloudwatch log group | `number` | `7` | no |
| <a name="input_create_iam_role"></a> [create\_iam\_role](#input\_create\_iam\_role) | Create IAM role with a defined name that permits Lambda to work with Route53 & ACM | `bool` | `true` | no |
| <a name="input_cron_enabled"></a> [cron\_enabled](#input\_cron\_enabled) | If true, CRON schedule rules will be enabled | `bool` | `true` | no |
| <a name="input_cron_schedule"></a> [cron\_schedule](#input\_cron\_schedule) | The schedule expression for how often the Lambda function runs | `string` | `"rate(24 hours)"` | no |
| <a name="input_description"></a> [description](#input\_description) | Lambda description | `string` | `"The AWS Let's Encrypt Lambda. URL: https://github.com/kvendingoldo/aws-letsencrypt-lambda"` | no |
| <a name="input_ecr_image_uri"></a> [ecr\_image\_uri](#input\_ecr\_image\_uri) | ECR image URI. Required only if enable\_ecr\_proxy is false | `string` | `null` | no |
| <a name="input_ecr_proxy_access_token"></a> [ecr\_proxy\_access\_token](#input\_ecr\_proxy\_access\_token) | The username to access to public registry. | `string` | `null` | no |
| <a name="input_ecr_proxy_enabled"></a> [ecr\_proxy\_enabled](#input\_ecr\_proxy\_enabled) | If true, ECR proxy for ghcr.io will be created | `bool` | `true` | no |
| <a name="input_ecr_proxy_repository_prefix"></a> [ecr\_proxy\_repository\_prefix](#input\_ecr\_proxy\_repository\_prefix) | The repository name prefix to use when caching images from the source registry. | `string` | `"ghcr-io-proxy"` | no |
| <a name="input_ecr_proxy_upstream_registry_url"></a> [ecr\_proxy\_upstream\_registry\_url](#input\_ecr\_proxy\_upstream\_registry\_url) | The registry URL of the upstream public registry to use as the source. | `string` | `"ghcr.io"` | no |
| <a name="input_ecr_proxy_username"></a> [ecr\_proxy\_username](#input\_ecr\_proxy\_username) | The username to access to public registry. | `string` | `null` | no |
| <a name="input_environ"></a> [environ](#input\_environ) | Environment variables passed to the Lambda function | `map(string)` | `{}` | no |
| <a name="input_events"></a> [events](#input\_events) | List of events for Lambda function (each event contains info about one certificate) | `any` | `[]` | no |
| <a name="input_iam_role_arn"></a> [iam\_role\_arn](#input\_iam\_role\_arn) | The ARN for the IAM role that permits Lambda to work with Route53 & ACM. Must be specified if monitoring\_interval is non-zero | `string` | `null` | no |
| <a name="input_memory_size"></a> [memory\_size](#input\_memory\_size) | The memory in Mb that the function can use | `string` | `128` | no |
| <a name="input_public_image"></a> [public\_image](#input\_public\_image) | GHCR image containing the function's deployment package | `string` | `"kvendingoldo/aws-letsencrypt-lambda:rc-0.29.0"` | no |
| <a name="input_security_group_ids"></a> [security\_group\_ids](#input\_security\_group\_ids) | The VPC security groups assigned to the Lambda | `list(string)` | `[]` | no |
| <a name="input_subnet_ids"></a> [subnet\_ids](#input\_subnet\_ids) | The VPC subnets in which the Lambda runs | `list(string)` | `[]` | no |
| <a name="input_tags"></a> [tags](#input\_tags) | Any tags that should be present on AWS resources | `map(string)` | `{}` | no |
| <a name="input_timeout"></a> [timeout](#input\_timeout) | The maximum time in seconds that the Lambda can run for | `string` | `900` | no |

## Outputs

No outputs.
<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## License
See [LICENSE](https://github.com/kvendingoldo/aws-letsencrypt-lambda/blob/main/LICENSE).
