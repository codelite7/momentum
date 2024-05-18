locals {
  # build a default spec using implmented variables
  default_application_spec = {
    "destination" = {
      "namespace" = "argo-cd"
      "server"    = "https://kubernetes.default.svc"
    }
    "project" = var.project
    "source" = {
      "repoURL" = var.source_repo_url
      "path"   = var.source_path
      "helm" = {
        "values"  = var.helm_values
      }
      "targetRevision" = var.source_target_revision
    }
    "syncPolicy" = var.sync_policy
  }

  # allow any spec field to be specified if it isn't implemented in variables
  application_spec = merge(local.default_application_spec, var.spec_override)

  application = {
    "apiVersion" = "argoproj.io/v1alpha1"
    "kind"       = "Application"
    "metadata" = {
      "name"      = var.name
      "namespace" = var.namespace
    }
    "spec" = local.application_spec
  }

  application_yaml = yamlencode(local.application)
}

resource "kubectl_manifest" "application" {
  yaml_body        = local.application_yaml
  sensitive_fields = ["spec.source.helm.values"]
}
