variable "argo-cd-set" {
  type = any
  default = {}
}

variable "argo-cd-set-list" {
  type = any
  default = {}
}

variable "argo-cd-version" {
  type = any
}

variable "argo-cd-values" {
  type = list(string)
  default = []
}