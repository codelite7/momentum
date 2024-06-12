# manage secrets for this environment via aws secrets manager for secret
# versioning and aws access control over secrets
data "aws_secretsmanager_secret_version" "secret" {
  secret_id = "demo"
}

locals {
  secrets = jsondecode(
    data.aws_secretsmanager_secret_version.secret.secret_string
  )
}

module "app" {
  source = "../../modules/app"
  name = "demo"
  targetRevision = "main"
  syncPolicy = yamldecode(file("syncPolicy.yaml"))
  values = templatefile("values.yaml", {
    defaultAgents: local.secrets.defaultAgents,
    workosApiKey: local.secrets.workosApiKey,
    workosCookiePassword: local.secrets.workosCookiePassword,
    openaiApiKey: local.secrets.openaiApiKey,
    groqApiKey: local.secrets.groqApiKey,
    anthropicApiKey: local.secrets.anthropicApiKey,
  })
}

module "deployment" {
  source = "../../modules/deployment"
  content-path = "../../../ui/dist/ui/browser"
  domain = "manifold.derp.ninja"
  cloudflareZoneId = "e3e1e8f58492cf7bd325b30a6a022295"
  proxied = true
  ttl = 1
}