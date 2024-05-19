variable cloudflareZoneId {
  type = string
}

variable "domain" {
  type = string
}

variable "content-path" {
  type = string
}

variable proxied {
  type = bool
}

variable ttl {
  type = number
  default = 60
}