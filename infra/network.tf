resource "google_compute_network" "default" {
  name                    = "${var.name}-network"
  auto_create_subnetworks = "false"
  depends_on = [google_project_service.compute]
}