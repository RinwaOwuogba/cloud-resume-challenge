resource "google_storage_bucket" "static_frontend" {
  name          = "${var.name}-static-bucket"
  location      = var.region
  force_destroy = true

  uniform_bucket_level_access = true
}

resource "google_storage_bucket" "archive_bucket" {
  name          = "${var.name}-archive-bucket"
  location      = var.region
  force_destroy = true

  uniform_bucket_level_access = true
}

resource "google_storage_bucket_object" "backend_archive" {
  name   = "backend.zip"
  bucket = google_storage_bucket.archive_bucket.name
  source = var.backend_code_path
}