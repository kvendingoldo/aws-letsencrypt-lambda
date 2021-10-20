resource "aws_iam_role" "iam_for_lambda" {
  name_prefix = replace(
  replace(var.function_name, "/(.{0,32}).*/", "$1"),
  "/^-+|-+$/",
  "",
  )

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF

}

resource "aws_iam_role_policy" "lambda_policy" {
  role = aws_iam_role.iam_for_lambda.id
  name = "policy"

  policy = var.lambda_role_policy
}

resource "aws_iam_role_policy_attachment" "vpc_permissions" {
  role       = aws_iam_role.iam_for_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"

  count = length(var.subnet_ids) != 0 ? 1 : 0
}
