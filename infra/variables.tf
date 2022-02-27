variable "credentials" {
  description = "Path to service account credentials"
}

variable "project" {
    description = "GCP project ID"
}

variable "region" {
  default = "europe-west3"
}

variable "zone" {
  default = "europe-west3-a"
}

variable "name" {}

variable "domain" {}

variable "backend_code_path" {}

variable "dns_zone_name" {}