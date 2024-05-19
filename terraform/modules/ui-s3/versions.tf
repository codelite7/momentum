terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.47.0"
    }
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "4.33.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.6.1"
    }
  }
  required_version = "~> 1.3"
}