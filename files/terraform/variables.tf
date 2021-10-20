#
# Naming variables
#
variable "blank_name" {
  type        = string
  description = "Blank name for Lambda resources"
}
variable "tags" {
  type        = map(string)
  description = "Any tags that should be present on the AWS resources"
  default     = {}
}


variable "image_uri" {
    type = string
}

variable "timeout" {
  description = "The maximum time in seconds that the Lambda can run for"
  # 20 minutes
  default     = 900
}
variable "memory_size" {
  description = "The memory in Gb that the function can use"
  default     = 128
}

variable "environ" {
  description = "Environment parameters passed to the Lambda function"
  #type        = map(string)
  default     = {}
}


variable "cron_enabled" {
  description = "The sceduling expression for how often the Lambda function runs."
  default = false
}
variable "cron_schedule" {
  description = "The sceduling expression for how often the Lambda function runs."
  default = "rate(24 hours)"

}





#
#
#
#variable "lambda_cron_schedule" {
#  description = "The sceduling expression for how often the Lambda function runs."
#}
#
#variable "subnet_ids" {
#  type        = list(string)
#  description = "The VPC subnets in which the Lambda runs"
#}
#
#variable "security_group_ids" {
#  type        = list(string)
#  description = "The VPC security groups assigned to the Lambda"
#}
#
#// Optional Variables
#variable "datadog_log_subscription_arn" {
#  description = "Log subscription arn for shipping logs to datadog"
#  default     = ""
#}
#
#variable "lambda_role_policy" {
#  description = "The Lambda IAM Role Policy."
#
#  default = <<END
#{
#  "Statement": [
#    {
#      "Effect": "Allow",
#      "Action": [
#        "logs:CreateLogGroup",
#        "logs:CreateLogStream",
#        "logs:PutLogEvents"
#      ],
#      "Resource": "arn:aws:logs:*:*:*"
#    }
#  ]
#}
#END
#
#}
#

#


#
#variable "lambda_iam_policy_name" {
#  description = "[DEPRECATED] The name for the Lambda functions IAM policy."
#  default     = ""
#}
#

#

#
#variable "tags" {
#  description = "A mapping of tags to assign to this lambda function."
#  type        = map(string)
#  default     = {}
#}
