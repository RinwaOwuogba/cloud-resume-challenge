resource "google_storage_bucket" "static_frontend" {
  name          = "${var.static_bucket_prefix}_bucket"
  location      = var.region
  force_destroy = true

  uniform_bucket_level_access = true
  enable_cdn = true
}