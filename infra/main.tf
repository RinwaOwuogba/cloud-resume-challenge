terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "~> 3.53"
    }
    google-beta = {
      source = "hashicorp/google-beta"
      version = "4.11.0"
    }
    tls = {
      source = "hashicorp/tls"
      version = "3.1.0"
    }
  }
  backend "gcs" {
    bucket  = "cloud-resume-project-342521_bucket_1"
    prefix  = "terraform/state"
  }
}

provider "google" {
  credentials = file(var.credentials)
  project = var.project
  region  = var.region
  zone    = var.zone
}

provider "google-beta" {
    project = var.project
}
