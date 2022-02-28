# Create default app engine application to activate firestore database
resource "google_app_engine_application" "app" {
  project       = google_project.my_project.project_id
  location_id   = var.region
  database_type = "CLOUD_FIRESTORE"
}