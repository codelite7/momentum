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