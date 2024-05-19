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

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}

provider "kubernetes" {
  config_path = "~/.kube/config"
}

provider "cloudflare" {
  api_token = local.secrets.cloudflareApiToken
}