output "lambda_arn" {
  value = aws_lambda_function.main.arn
}

output "lambda_iam_role_name" {
  value = aws_iam_role.main.name
}
