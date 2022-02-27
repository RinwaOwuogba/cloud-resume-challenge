terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "~>4.11.0"
    }
  }
  backend "gcs" {
    bucket  = "cloud-resume-project-342521_bucket_1"
    prefix  = "terraform/state"
  }
}

provider "google" {
  project = var.project
  region  = var.region
  zone    = var.zone
}
