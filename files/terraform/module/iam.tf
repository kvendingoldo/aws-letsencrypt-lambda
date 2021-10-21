resource "aws_iam_role" "main" {
  name               = var.blank_name
  description        = "IAM role for for Lambda ${var.blank_name}"
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

#
# VPC permissions
#
resource "aws_iam_role_policy_attachment" "vpc_permissions" {
  count      = length(var.subnet_ids) != 0 ? 1 : 0
  role       = aws_iam_role.main.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

#
# Logging policy
#
resource "aws_iam_policy" "logging" {
  name        = format("%s-%s", var.blank_name, "logging")
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<-POLICY
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
POLICY
}
resource "aws_iam_role_policy_attachment" "logging" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.logging.arn
}

#
# ACM policy
#
resource "aws_iam_policy" "acm" {
  name        = format("%s-%s", var.blank_name, "acm")
  path        = "/"
  description = "IAM policy for working with ACM from a lambda"

  policy = <<-POLICY
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
POLICY
}
resource "aws_iam_role_policy_attachment" "acm" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.acm.arn
}

#
# Route53 policy
#
resource "aws_iam_policy" "route53" {
  name        = format("%s-%s", var.blank_name, "route53")
  path        = "/"
  description = "IAM policy for working with Route53 from a lambda"

  policy = <<-POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "route53:ListHostedZonesByName",
        "route53:ListResourceRecordSets",
        "route53:GetChange",
        "route53:ChangeResourceRecordSets"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
POLICY
}
resource "aws_iam_role_policy_attachment" "route53" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.route53.arn
}


