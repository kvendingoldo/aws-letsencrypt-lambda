resource "aws_iam_role" "main" {
  count = var.create_iam_role ? 1 : 0

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
  count      = (length(var.subnet_ids) != 0 && var.create_iam_role) ? 1 : 0
  role       = aws_iam_role.main[0].name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}

#
# Logging policy
#
resource "aws_iam_policy" "logging" {
  count = var.create_iam_role ? 1 : 0

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
  count = var.create_iam_role ? 1 : 0

  role       = aws_iam_role.main[0].name
  policy_arn = aws_iam_policy.logging[0].arn
}

#
# ACM policy
#
resource "aws_iam_policy" "acm" {
  count = var.create_iam_role ? 1 : 0

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
  count = var.create_iam_role ? 1 : 0

  role       = aws_iam_role.main[0].name
  policy_arn = aws_iam_policy.acm[0].arn
}

#
# Route53 policy
#
resource "aws_iam_policy" "route53" {
  count = var.create_iam_role ? 1 : 0

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
  count = var.create_iam_role ? 1 : 0

  role       = aws_iam_role.main[0].name
  policy_arn = aws_iam_policy.route53[0].arn
}

#
# Secrets Manager policy
#
resource "aws_iam_policy" "secretsmanager" {
  count = var.create_iam_role ? 1 : 0

  name        = format("%s-%s", var.blank_name, "secretsmanager")
  path        = "/"
  description = "IAM policy for working with secretsmanager from a lambda"

  policy = <<-POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "secretsmanager:*"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
POLICY
}
resource "aws_iam_role_policy_attachment" "secretsmanager" {
  count = var.create_iam_role ? 1 : 0

  role       = aws_iam_role.main[0].name
  policy_arn = aws_iam_policy.secretsmanager[0].arn
}
