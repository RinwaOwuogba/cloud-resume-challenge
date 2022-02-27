# IAM entry for all users to view static files bucket
resource "google_storage_bucket_iam_member" "default" {
  bucket = google_storage_bucket.static_frontend.name
  role = "roles/storage.legacyObjectReader"
  member = "allUsers"
}

# # IAM entry for all users to invoke the function
# resource "google_cloudfunctions_function_iam_member" "invoker" {
#   project        = google_cloudfunctions_function.function.project
#   region         = google_cloudfunctions_function.function.region
#   cloud_function = google_cloudfunctions_function.function.name

#   role   = "roles/cloudfunctions.invoker"
#   member = "allUsers"
# }