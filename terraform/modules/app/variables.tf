# required
variable "name" {
  description = "Application name"
  type        = string
}

# optional
variable "values" {
  description = "Helm values as a raw string in YAML format"
  type        = string
  default     = ""
}

variable "targetRevision" {
  description = "App target revision"
  type        = string
}

variable "syncPolicy" {
  description = "ArgoCD application sync policy. Defaults to a frequent automatic sync."
  type = object({
    automated = object({
      prune    = bool
      selfHeal = bool
    })
    retry = object({
      backoff = object({
        duration    = string
        factor      = number
        maxDuration = string
      })
      limit = number
    })
    syncOptions = list(string)
  })
  default = {
    "automated" = {
      "prune"    = true
      "selfHeal" = true
    }
    "retry" = {
      "backoff" = {
        "duration"    = "5s"
        "factor"      = 2
        "maxDuration" = "3m"
      }
      "limit" = 3
    }
    "syncOptions" = [
      "CreateNamespace=true",
      "PrunePropagationPolicy=foreground",
      "PruneLast=true",
    ]
  }
}
