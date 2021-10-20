resource "aws_iam_role" "main" {
  name               = var.blank_name
  description        = "TODO"
  tags               = var.tags
  assume_role_policy = <<-POLICY
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
  POLICY
}

resource "aws_iam_policy" "logging" {
  name        = "lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}
resource "aws_iam_role_policy_attachment" "logs" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.logging.arn
}

resource "aws_iam_policy" "acm" {
  name        = "lambda_acm"
  path        = "/"
  description = "IAM policy for work with ACM from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "acm:AddTagsToCertificate",
        "acm:DescribeCertificate",
        "acm:GetCertificate",
        "acm:ImportCertificate",
        "acm:ListCertificates",
        "acm:ListTagsForCertificate"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
EOF
}
resource "aws_iam_role_policy_attachment" "acm" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.acm.arn
}

#
#resource "aws_iam_role_policy" "lambda_policy" {
#  role = aws_iam_role.iam_for_lambda.id
#  name = "policy"
#
#  policy = var.lambda_role_policy
#}
#
#resource "aws_iam_role_policy_attachment" "vpc_permissions" {
#  role       = aws_iam_role.iam_for_lambda.name
#  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
#
#  count = length(var.subnet_ids) != 0 ? 1 : 0
#}
