terraform {
  required_version = "~> 1.3"
  backend "s3" {
    bucket = "manifold-terraform"
    key    = "deployments/demo/terraform.tfstate"
    region = "us-west-2"
  }
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