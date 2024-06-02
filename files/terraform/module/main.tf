#
# Lambda
#
resource "aws_lambda_function" "main" {
  function_name = var.blank_name
  description   = var.description
  tags          = var.tags

  role         = var.create_iam_role ? aws_iam_role.main[0].arn : var.iam_role_arn
  image_uri    = local.lambda_image
  package_type = "Image"
  timeout      = var.timeout
  memory_size  = var.memory_size

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.security_group_ids
  }

  environment {
    variables = merge(var.environ, { "MODE" : "cloud", "FORMATTER_TYPE" : "JSON" })
  }

  depends_on = [
    aws_ecr_pull_through_cache_rule.lambda_proxy
  ]
}

#
# Cloudwatch group
#
resource "aws_cloudwatch_log_group" "main" {
  name = "/aws/lambda/${var.blank_name}"
  tags = var.tags

  retention_in_days = var.cloudwatch_log_group_retention
  depends_on        = [aws_lambda_function.main]
}

#
# Event schedule
#
resource "aws_cloudwatch_event_rule" "schedule" {
  for_each = local.events

  name        = format("%s_%s", var.blank_name, each.key)
  description = "This event will run according to a schedule for Lambda ${var.blank_name}"
  tags        = var.tags

  schedule_expression = var.cron_schedule
  state               = "ENABLED"
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  for_each = local.events

  function_name = aws_lambda_function.main.function_name
  statement_id  = format("%s_%s_allowExecutionFromCloudWatch", var.blank_name, replace(each.key, ".", "-"))
  action        = "lambda:InvokeFunction"
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.schedule[each.key].arn
}

resource "aws_cloudwatch_event_target" "event_target" {
  for_each = local.events

  rule = aws_cloudwatch_event_rule.schedule[each.key].name
  arn  = aws_lambda_function.main.arn

  input = jsonencode(each.value)
}
