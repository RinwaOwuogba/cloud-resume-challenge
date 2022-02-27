resource "google_compute_network" "default" {
  name                    = var.network_name
  auto_create_subnetworks = "false"
}