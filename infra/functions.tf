resource "google_cloudfunctions_function" "function" {
  name        = "${var.name}-api-function"
  description = "API for cloud resume project backend"
  runtime     = "go116"
  depends_on = [
    google_project_service.functions,
    google_project_service.build
  ]

  available_memory_mb   = 256
  source_archive_bucket = google_storage_bucket.archive_bucket.name
  source_archive_object = google_storage_bucket_object.backend_archive.name
  trigger_http          = true
  entry_point           = "ServerEntry"

  environment_variables = {
    "GCP_PROJECT" = var.project
  }
}
