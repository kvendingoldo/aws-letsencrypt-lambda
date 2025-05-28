<!-- BADGES -->
[![Github release](https://img.shields.io/github/v/release/kvendingoldo/aws-letsencrypt-lambda?style=for-the-badge)](https://github.com/kvendingoldo/aws-letsencrypt-lambda/releases) [![Contributors](https://img.shields.io/github/contributors/kvendingoldo/aws-letsencrypt-lambda?style=for-the-badge)](https://github.com/kvendingoldo/aws-letsencrypt-lambda/graphs/contributors) ![maintenance status](https://img.shields.io/maintenance/yes/2025.svg?style=for-the-badge) [![Go report](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/kvendingoldo/aws-letsencrypt-lambda/) [![OpenTofu support](https://img.shields.io/badge/opentofu-supported-blue.svg?logo=opentofu&style=for-the-badge)](https://opentofu.org/) [![OpenTofu support](https://img.shields.io/badge/terraform-supported-blue.svg?logo=terraform&style=for-the-badge)](https://www.terraform.io/)

# aws-letsencrypt-lambda

## Overview
It's common for people to desire having valid TLS certificates without wanting to pay for them.
This is where the [Let's Encrypt](https://letsencrypt.org) project can help. Although it offers an excellent service, issuing and renewing certifications is not made simple. It is particularly visible in cloudy environments.

This repository represents a straightforward Lambda function for AWS that uses CRON (cloud watch events) and can simply issue and renew certificates without any manual operation on the part of the operator. In addition, this repository offers a Terraform module that speeds up Lambda's onboarding.

An article talks about the solution on [HackerNoon](https://hackernoon.com/aws-letsencrypt-lambda-or-why-i-wrote-a-custom-tls-provider-for-aws-using-opentofu-and-go)

## Documentation
You can review the following documents on the Lambda to learn more:
* [How to use the Lambda inside of AWS](docs/how_to_use_aws.md)
* [How to use the Lambda locally](docs/how_to_use_locally.md)
* [How to use OpenTofu automation](docs/how_to_use_opentofu.md)
* [Labmda's environment variables](docs/environment_variables.md)
