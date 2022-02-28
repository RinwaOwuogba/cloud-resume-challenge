resource "google_compute_network" "default" {
  name                    = "${var.name}-network"
  auto_create_subnetworks = "false"
  depends_on              = [google_project_service.compute]
}

resource "google_compute_region_network_endpoint_group" "function_neg" {
  provider              = google-beta
  name                  = "${var.name}-neg"
  network_endpoint_type = "SERVERLESS"
  region                = var.region
  cloud_function {
    function = google_cloudfunctions_function.function.name
  }
}

