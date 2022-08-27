#
# Naming variables
#
variable "blank_name" {
  type        = string
  description = "Blank name for AWS resources"
}
variable "tags" {
  type        = map(string)
  description = "Any tags that should be present on AWS resources"
  default     = {}
}

#
# Lambda variables
#
variable "subnet_ids" {
  type        = list(string)
  description = "The VPC subnets in which the Lambda runs"
  default     = []
}
variable "security_group_ids" {
  type        = list(string)
  description = "The VPC security groups assigned to the Lambda"
  default     = []
}

variable "description" {
  type        = string
  description = "Lambda description"
  default     = ""
}
variable "image_uri" {
  type        = string
  description = "ECR image URI containing the function's deployment package"
}
variable "timeout" {
  type        = string
  description = "The maximum time in seconds that the Lambda can run for"
  default     = 900
}
variable "memory_size" {
  type        = string
  description = "The memory in Mb that the function can use"
  default     = 128
}
variable "environ" {
  description = "Environment variables passed to the Lambda function"
  type        = map(string)
  default     = {}
}

#
# IAM configuration
#
variable "create_iam_role" {
  description = "Create IAM role with a defined name that permits Lambda to work with Route53 & ACM"
  type        = bool
  default     = false
}
variable "iam_role_arn" {
  description = "The ARN for the IAM role that permits Lambda to work with Route53 & ACM. Must be specified if monitoring_interval is non-zero"
  type        = string
  default     = null
}

#
# Lambda events
#
variable "events" {
  type        = any
  description = "List of events for Lambda function (each event contains info about one certificate)"
  default     = []
}

#
# Cron variables
#
variable "cron_enabled" {
  type    = bool
  default = true
}
variable "cron_schedule" {
  type        = string
  description = "The schedule expression for how often the Lambda function runs"
  default     = "rate(24 hours)"
}

#
# Logging
#
variable "retention" {
  type        = number
  description = "Number of days to retain log events in the specified log group"
  default     = 7
}
