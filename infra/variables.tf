variable "project" {
    description = "GCP project ID"
}

variable "region" {
  default = "europe-west3"
}

variable "zone" {
  default = "europe-west3-a"
}


variable "static_bucket_prefix" {}

variable "lb_name" {}

variable "network_name" {}