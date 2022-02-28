# Enable Compute Engine API
resource "google_project_service" "compute" {
  project = var.project
  service = "compute.googleapis.com"

  disable_dependent_services = true
}

# Enable Cloud DNS API
resource "google_project_service" "dns" {
  project = var.project
  service = "dns.googleapis.com"

  disable_dependent_services = true
}

# Enable Cloud Functions API
resource "google_project_service" "functions" {
  project = var.project
  service = "cloudfunctions.googleapis.com"

  disable_dependent_services = true
}

# Enable Cloud Build API
resource "google_project_service" "build" {
  project = var.project
  service = "cloudbuild.googleapis.com"

  disable_dependent_services = true
}