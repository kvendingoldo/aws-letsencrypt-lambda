resource "aws_ecr_pull_through_cache_rule" "lambda" {
  ecr_repository_prefix = var.ecr_repository_prefix
  upstream_registry_url = var.ecr_upstream_registry_url
}
