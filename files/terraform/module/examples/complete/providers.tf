provider "aws" {
  region = "us-east-2"
}

provider "aws" {
  alias  = "use1"
  region = "us-east-1"
}
