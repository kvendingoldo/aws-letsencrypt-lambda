resource "aws_lambda_function" "main" {
  #description = "TODO"



  s3_bucket     = var.s3_bucket
  s3_key        = var.s3_key
  function_name = var.function_name
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = var.handler
  runtime       = var.runtime
  timeout       = var.timeout
  memory_size   = var.memory_size
  layers        = [var.layer]
  tags          = var.tags

#  vpc_config {
#    subnet_ids         = var.subnet_ids
#    security_group_ids = var.security_group_ids
#  }

  environment {
    variables = var.lambda_env
  }
}

#resource "aws_cloudwatch_event_rule" "cron_schedule" {
#  name                = replace("${var.function_name}-cron_schedule", "/(.{0,64}).*/", "$1")
#  description         = "This event will run according to a schedule for Lambda ${var.function_name}"
#  schedule_expression = var.lambda_cron_schedule
#  is_enabled          = var.is_enabled
#}
#
#resource "aws_cloudwatch_event_target" "event_target" {
#  rule = aws_cloudwatch_event_rule.cron_schedule.name
#  arn  = aws_lambda_function.lambda_function.arn
#}
#
#resource "aws_cloudwatch_log_group" "loggroup" {
#  name              = "/aws/lambda/${var.function_name}"
#  retention_in_days = 7
#  depends_on        = [aws_lambda_function.lambda_function]
#}
#
#resource "aws_lambda_permission" "allow_cloudwatch" {
#  statement_id  = "AllowExecutionFromCloudWatch"
#  action        = "lambda:InvokeFunction"
#  function_name = aws_lambda_function.lambda_function.function_name
#  principal     = "events.amazonaws.com"
#  source_arn    = aws_cloudwatch_event_rule.cron_schedule.arn
#}
#
#resource "aws_cloudwatch_log_subscription_filter" "kinesis_log_stream" {
#  count           = var.datadog_log_subscription_arn != "" ? 1 : 0
#  name            = "kinesis-log-stream-${var.function_name}"
#  destination_arn = var.datadog_log_subscription_arn
#  log_group_name  = aws_cloudwatch_log_group.lambda_loggroup.name
#  distribution    = "ByLogStream"
#  filter_pattern  = ""
#  depends_on      = [aws_lambda_function.lambda_function]
#}
