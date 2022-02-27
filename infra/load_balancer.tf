locals {
  health_check = {
    check_interval_sec  = null
    timeout_sec         = null
    healthy_threshold   = null
    unhealthy_threshold = null
    request_path        = "/"
    port                = 80
    host                = null
    logging             = null
  }
}

module "gce-lb-https" {
  source  = "GoogleCloudPlatform/lb-http/google"
  version = "~> 5.1"
  name    = var.network_name
  project = var.project
  firewall_networks = [google_compute_network.default.self_link]
  url_map           = google_compute_url_map.lb-url-map.self_link
  create_url_map    = false
  ssl               = true
  private_key       = tls_private_key.example.private_key_pem
  certificate       = tls_self_signed_cert.example.cert_pem

  backends = {
    default = {
      description                     = null
      protocol                        = "HTTP"
      port                            = 80
      port_name                       = "http"
      timeout_sec                     = 10
      connection_draining_timeout_sec = null
      enable_cdn                      = false
      security_policy                 = null
      session_affinity                = null
      affinity_cookie_ttl_sec         = null
      custom_request_headers          = null
      custom_response_headers         = null

      health_check = local.health_check
      log_config = {
        enable      = true
        sample_rate = 1.0
      }
      iap_config = {
        enable               = false
        oauth2_client_id     = ""
        oauth2_client_secret = ""
      }
    }
  }
}


resource "google_compute_url_map" "lb-url-map" {
  name            = var.lb_name
  default_service = module.gce-lb-https.backend_services["default"].self_link

  host_rule {
    hosts        = ["*"]
    path_matcher = "allpaths"
  }

  path_matcher {
    name            = "allpaths"
    default_service = module.gce-lb-https.backend_services["default"].self_link

    path_rule {
      ath_rule {
        paths = [
          "/resume",
          "/resume/*"
        ]
        service = google_compute_backend_bucket.static_frontend.self_link
      }
    }
  }
}

resource "google_compute_backend_bucket" "static_frontend" {
  name        = "${var.static_bucket_prefix}_backend_bucket"
  description = "Contains static frontend resources"
  bucket_name = google_storage_bucket.static_frontend.name
  enable_cdn  = true
}