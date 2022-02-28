resource "google_compute_global_address" "default" {
  name = "${var.name}-address"
  depends_on = [google_project_service.compute]
}

resource "google_compute_managed_ssl_certificate" "default" {
  provider = google-beta

  name = "${var.name}-cert"
  managed {
    domains = ["${var.domain}"]
  }
  depends_on = [google_project_service.compute]
}

# Backend bucket for frontend
resource "google_compute_backend_bucket" "default" {
  name        = "${var.name}-backend-bucket"
  description = "Contains static frontend resources"
  bucket_name = google_storage_bucket.static_frontend.name
  enable_cdn  = true
  depends_on = [google_project_service.compute]
}

# Backend service for severless function
resource "google_compute_backend_service" "default" {
  name      = "${var.name}-backend"

  protocol  = "HTTP"
  port_name = "http"
  timeout_sec = 30

  backend {
    group = google_compute_region_network_endpoint_group.function_neg.id
  }
}

resource "google_compute_url_map" "default" {
  name            = "${var.name}-urlmap"
  depends_on = [google_project_service.compute]

  default_service = google_compute_backend_bucket.default.id

   host_rule {
    hosts        = [var.domain]
    path_matcher = "resume-site"
  }

  path_matcher {
    name            = "resume-site"
    default_service = google_compute_backend_bucket.default.id

    path_rule {
      paths   = ["/"]
      service = google_compute_backend_bucket.default.id
    }

    path_rule {
      paths   = ["/api/*"]
      service = google_compute_backend_service.default.id
    }
  }

  test {
    service = google_compute_backend_bucket.default.id
    host    = var.domain
    path    = "/"
  }
}

resource "google_compute_target_https_proxy" "default" {
  name   = "${var.name}-https-proxy"

  url_map          = google_compute_url_map.default.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.default.id
  ]
  depends_on = [google_project_service.compute]
}

resource "google_compute_global_forwarding_rule" "default" {
  name   = "${var.name}-lb"

  target = google_compute_target_https_proxy.default.id
  port_range = "443"
  ip_address = google_compute_global_address.default.address
  
  depends_on = [google_project_service.compute]
}
