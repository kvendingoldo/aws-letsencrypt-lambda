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
  default = ""
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
  description = "The memory in Gb that the function can use"
  default     = 128
}
variable "environ" {
  description = "Environment parameters passed to the Lambda function"
  type        = map(string)
  default     = {}
}

#
# Cron variables
#
variable "cron_enabled" {
  type    = bool
  default = false
}
variable "cron_schedule" {
  type        = string
  description = "The schedule expression for how often the Lambda function runs"
  default     = "rate(24 hours)"
}