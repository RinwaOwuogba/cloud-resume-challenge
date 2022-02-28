# Create default app engine application to activate firestore database
resource "google_app_engine_application" "app" {
  project       = var.project
  location_id   = var.region
  database_type = "CLOUD_FIRESTORE"
  depends_on = [
    google_project_service.app_engine
  ]
}