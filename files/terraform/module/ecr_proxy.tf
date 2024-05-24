resource "aws_ecr_repository" "lambda_proxy" {
  count                = var.enable_ecr_proxy  ? 0 : 1
  name                 = var.ecr_repository_prefix
  image_tag_mutability = "MUTABLE"
  force_delete         = true
  image_scanning_configuration {
    scan_on_push = false
  }
}

resource "aws_ecr_pull_through_cache_rule" "docker_hub" {
  count                 = var.enable_ecr_proxy  ? 0 : 1
  ecr_repository_prefix = var.ecr_repository_prefix
  upstream_registry_url = "registry-1.docker.io"
  credential_arn        = var.dockerhub_proxy_secret_arn
}
