output "load_balancer_ip" {
  value = google_compute_global_address.default.address
}

output "static_frontend_bucket_url" {
  value = google_storage_bucket.static_frontend.url
}

