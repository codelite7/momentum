module "demo-manifold" {
  source = "../modules/manifold"
  cluster_name = "manifold-demo"
  eks_managed_node_groups = {
    one = {
      capacity_type = "spot"
      name = "t3-large"
      instance_types = ["t3.large"]
      min_size     = 1
      max_size     = 3
      desired_size = 1
    }
  }
  region = "us-west-2"
  argo-cd-version = "6.9.2"
}