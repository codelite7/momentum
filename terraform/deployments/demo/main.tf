module "app" {
  source = "../../modules/app"
  name = "demo"
  targetRevision = "init"
  syncPolicy = yamldecode(file("syncPolicy.yaml"))
  values = templatefile("values.yaml", {})
}