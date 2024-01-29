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
variable "image" {
  type        = string
  description = "Docker image containing the function's deployment package"
  default     = "kvendingoldo/aws-letsencrypt-lambda:rc-0.9.0"
}
variable "description" {
  type        = string
  description = "Lambda description"
  default     = ""
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
# ECR proxy
#
variable "ecr_proxy_enabled" {
  default = false
}
variable "ecr_proxy_repository_prefix" {
  description = "The repository name prefix to use when caching images from the source registry."
  type        = string
  default     = "proxy-cache"
}
variable "ecr_proxy_upstream_registry_url" {
  description = "The registry URL of the upstream public registry to use as the source."
  type        = string
  default     = "registry-1.docker.io"

  validation {
    condition     = can(regex("^((public\\.ecr\\.aws)|(registry-1\\.docker\\.io)|(registry\\.k8s\\.io)|(quay\\.io)|(ghcr\\.io)|(\\w+\\.azurecr\\.io))$", var.container_registry_url))
    error_message = "Invalid container registry URL. It must be one of: public.ecr.aws, registry-1.docker.io, registry.k8s.io, quay.io, ghcr.io, {custom}.azurecr.io"
  }
}
variable "ecr_proxy_username" {
  description = "The username to access to public registry"
  type        = string
  default     = null
}
variable "ecr_proxy_access_token" {
  description = "The username to access to public registry"
  type        = string
  default     = null
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
  type        = bool
  description = "If true, CRON schedule rules will be enabled"
  default     = true
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
