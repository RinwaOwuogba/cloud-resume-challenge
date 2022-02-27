# resource "google_storage_bucket_object" "backend_archive" {
#   name   = "backend.zip"
#   bucket = google_storage_bucket.archive_bucket.name
#   source = var.backend_code_path
# }

# resource "google_cloudfunctions_function" "function" {
#   name        = "${var.name}-api-function"
#   description = "API for cloud resume project backend"
#   runtime     = "go116"

#   available_memory_mb   = 256
#   source_archive_bucket = google_storage_bucket.archive_bucket.name
#   source_archive_object = google_storage_bucket_object.backend_archive.name
#   trigger_http          = true
#   entry_point           = "ServerEntry"
# }

