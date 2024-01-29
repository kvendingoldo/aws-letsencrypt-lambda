# aws-letsencrypt-lambda

[![Github release](https://img.shields.io/github/v/release/kvendingoldo/aws-letsencrypt-lambda?style=for-the-badge)](https://github.com/kvendingoldo/aws-letsencrypt-lambda/releases) [![Contributors](https://img.shields.io/github/contributors/kvendingoldo/aws-letsencrypt-lambda?style=for-the-badge)](https://github.com/kvendingoldo/aws-letsencrypt-lambda/graphs/contributors) ![maintenance status](https://img.shields.io/maintenance/yes/2024.svg?style=for-the-badge)
[![support opentofu](https://img.shields.io/badge/supports-opentofu-green?style=for-the-badge)](https://opentofu.org/)

## Overview
It's common for people to desire having valid TLS certificates without wanting to pay for them.
This is where the [Let's Encrypt](https://letsencrypt.org) project can help. Although it offers an excellent service, issuing and renewing certifications is not made simple. It is particularly visible in cloudy environments.

This repository represents a straightforward Lambda function for AWS that uses CRON (cloud watch events) and can simply issue and renew certificates without any manual operation on the part of the operator. In addition, this repository offers a Terraform module that speeds up Lambda's onboarding.

## Documentation
You can review the following documents on the Lambda to learn more:
* [How to use the Lambda inside of AWS](docs/how_to_use_aws.md)
* [How to use the Lambda locally](docs/how_to_use_locally.md)
* [How to use Terraform automation](docs/how_to_use_terraform.md)
* [Lambda's environment variables](docs/environment_variables.md)
