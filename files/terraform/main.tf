#
# Lambda
#
resource "aws_lambda_function" "main" {
  function_name = var.blank_name
  description   = var.description
  tags          = var.tags

  role         = aws_iam_role.main.arn
  image_uri    = var.image_uri
  package_type = "Image"
  timeout      = var.timeout
  memory_size  = var.memory_size

  vpc_config {
    subnet_ids         = var.subnet_ids
    security_group_ids = var.security_group_ids
  }

  environment {
    variables = merge(var.environ, { "MODE" : "cloud" })
  }
}

#
# Cloudwatch group
#
resource "aws_cloudwatch_log_group" "main" {
  name = "/aws/lambda/${var.blank_name}"
  tags = var.tags

  retention_in_days = var.retention
  depends_on        = [aws_lambda_function.main]
}

#
# Event schedule
#
resource "aws_cloudwatch_event_rule" "schedule" {
  for_each = local.events

  name        = format("%s_%s_schedule", var.blank_name, each.key)
  description = "This event will run according to a schedule for Lambda ${var.blank_name}"
  tags        = var.tags

  schedule_expression = var.cron_schedule
  is_enabled          = true
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

  input = <<-INPUT
{
"domain_name": "${each.value["DomainName"]}",
"acme_url": "${each.value["AcmeUrl"]}",
"acme_email": "${each.value["AcmeEmail"]}",
"reimport_threshold": ${each.value["ReImportThreshold"]},
"issue_type": "${each.value["IssueType"]}"
}
INPUT
}
