# Enable Compute Engine API
resource "google_project_service" "compute" {
  project = var.project
  service = "compute.googleapis.com"

  disable_dependent_services = true
}

# Enable Firestore API
resource "google_project_service" "firestore" {
  project = var.project
  service = "firestore.googleapis.com"

  disable_dependent_services = true
}

# Enable Cloud DNS API
resource "google_project_service" "dns" {
  project = var.project
  service = "dns.googleapis.com"

  disable_dependent_services = true
}