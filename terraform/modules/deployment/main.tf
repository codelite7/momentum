module "ui-s3" {
  source = "../ui-s3"
  content-path = var.content-path
  cloudflareZoneId = var.cloudflareZoneId
  domain = var.domain
  proxied = var.proxied
  ttl = var.ttl
}