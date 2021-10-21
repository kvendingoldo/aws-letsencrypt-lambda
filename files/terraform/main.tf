#
# Lambda
#
resource "aws_lambda_function" "main" {
  function_name = var.blank_name
  description   = var.description
  role          = aws_iam_role.main.arn
  image_uri     = var.image_uri
  package_type  = "Image"
  timeout       = var.timeout
  memory_size   = var.memory_size
  tags          = var.tags

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.security_group_ids
  }

  environment {
    variables = var.environ
  }
}

#
# Cloudwatch group
#
resource "aws_cloudwatch_log_group" "main" {
  name              = "/aws/lambda/${var.blank_name}"
  retention_in_days = 7
  depends_on        = [aws_lambda_function.main]
}

#
# Cron schedule
#
resource "aws_cloudwatch_event_rule" "cron_schedule" {
  count = var.cron_enabled ? 1 : 0

  name                = format("%s-cron_schedule", var.blank_name)
  description         = "This event will run according to a schedule for Lambda ${var.blank_name}"
  schedule_expression = var.cron_schedule
  is_enabled          = true
}

resource "aws_lambda_permission" "allow_cloudwatch" {
  count = var.cron_enabled ? 1 : 0

  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.main.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.cron_schedule[0].arn
}

resource "aws_cloudwatch_event_target" "event_target" {
  count = var.cron_enabled ? 1 : 0

  rule = aws_cloudwatch_event_rule.cron_schedule[0].name
  arn  = aws_lambda_function.main.arn
}
