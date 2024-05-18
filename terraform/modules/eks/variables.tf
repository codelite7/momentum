variable "region" {
  type        = string
}

variable "cluster_name" {
  type = string
}

variable "cluster_version" {
  type = string
  default = "1.29"
}

variable "eks_managed_node_groups" {
  type        = any
  default     = {}
}