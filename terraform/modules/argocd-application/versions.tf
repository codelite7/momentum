terraform {
  required_version = ">= 0.13.1"

  required_providers {
    kubectl = {
      source  = "alekc/kubectl"
      version = "2.0.4"
    }
  }
}
