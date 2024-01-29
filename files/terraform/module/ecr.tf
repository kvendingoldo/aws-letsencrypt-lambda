resource "aws_secretsmanager_secret" "image_registry" {
  count = var.ecr_proxy_enabled && var.ecr_proxy_username != null && var.ecr_proxy_access_token != null ? 1 : 0

  name                    = format("ecr-pullthroughcache/%s", var.ecr_proxy_repository_prefix)
  recovery_window_in_days = 0
}
resource "aws_secretsmanager_secret_version" "image_registry" {
  count = var.ecr_proxy_enabled && var.ecr_proxy_username != null && var.ecr_proxy_access_token != null ? 1 : 0

  secret_id     = aws_secretsmanager_secret.image_registry.id
  secret_string = format("{'username':'%s','accessToken':'%s'}", var.ecr_proxy_username, var.ecr_proxy_access_token)
}

resource "aws_ecr_pull_through_cache_rule" "lambda" {
  count = var.ecr_proxy_enabled ? 1 : 0

  ecr_repository_prefix = var.ecr_proxy_repository_prefix
  upstream_registry_url = var.ecr_proxy_upstream_registry_url

  depends_on     = [aws_secretsmanager_secret.image_registry]
  credential_arn = var.ecr_proxy_username != null && var.ecr_proxy_access_token != null ? aws_secretsmanager_secret.image_registry[0].arn : null
}
