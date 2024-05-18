# manage secrets for this environment via aws secrets manager for secret
# versioning and aws access control over secrets
data "aws_secretsmanager_secret_version" "secret" {
  secret_id = "infra"
}

locals {
  secrets = jsondecode(
    data.aws_secretsmanager_secret_version.secret.secret_string
  )
}

module "eks" {
  source       = "../modules/eks"
  cluster_name = "manifold"
  eks_managed_node_groups = {
    one = {
      capacity_type  = "SPOT"
      name           = "t3-large"
      instance_types = ["t3.large"]
      min_size       = 1
      max_size       = 3
      desired_size   = 1
    }
  }
  region = "us-west-2"
}

resource "helm_release" "argo-cd" {
  depends_on = [module.eks]
  name       = "argo-cd"
  repository = "https://argoproj.github.io/argo-helm"
  chart      = "argo-cd"
  version = "6.9.2"
  namespace = "argo-cd"
  create_namespace = true
  values = [
    templatefile("argocd-values.yaml", {
      pat: local.secrets.pat
    })
  ]
}

resource "kubernetes_manifest" "app-cluster" {
  manifest = yamldecode(file("app-cluster.yaml"))
}