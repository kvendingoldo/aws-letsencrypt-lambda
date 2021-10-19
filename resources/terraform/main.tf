# https://github.com/uridium/terraform-aws-lambda-scheduler/blob/master/main.tf
# cron = "cron(01 06 ? * MON-FRI *)"
resource "aws_lambda_function" "this" {
  function_name = local.name
  layers        = var.layer_enabled ? [aws_lambda_layer_version.this[0].arn] : null

  filename         = data.archive_file.function_zip.output_path
  source_code_hash = data.archive_file.function_zip.output_base64sha256

  dynamic "environment" {
    for_each = var.vars[*]
    content { variables = environment.value }
  }

  tags = var.tags

  description = var.description
  handler     = var.handler
  runtime     = var.runtime
  memory_size = var.memory_size
  timeout     = var.timeout

  tracing_config {
    mode = var.tracing_mode
  }

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.security_group_ids
  }

  role = aws_iam_role.this.arn
}