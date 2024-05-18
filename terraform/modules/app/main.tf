module "argocd-application" {
  source = "../../modules/argocd-application"
  name   = "app-${var.name}"
  source_path = "app-chart"
  source_repo_url = "https://github.com/codelite7/momentum.git"
  source_target_revision = var.targetRevision
  helm_values = var.values
  sync_policy = var.syncPolicy
}