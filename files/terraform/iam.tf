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

resource "aws_iam_policy" "logging" {
  name        = "lambda_logging"
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
resource "aws_iam_role_policy_attachment" "logs" {
  role       = aws_iam_role.main.name
  policy_arn = aws_iam_policy.logging.arn
}

resource "aws_iam_policy" "acm" {
  name        = "lambda_acm"
  path        = "/"
  description = "IAM policy for work with ACM from a lambda"

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


resource "aws_iam_role_policy_attachment" "vpc_permissions" {
  count      = length(var.subnet_ids) != 0 ? 1 : 0
  role       = aws_iam_role.main.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaVPCAccessExecutionRole"
}
