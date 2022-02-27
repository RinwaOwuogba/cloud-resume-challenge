# Create DNS A record for load balancer
resource "google_dns_record_set" "a" {
  name = "${var.domain}."
  type = "A"
  ttl  = 300

  managed_zone = var.dns_zone_name

  rrdatas = [google_compute_global_address.default.address]
}

# todo: add IPv6 address for lb
# # Create DNS AAAA record for load balancer
# resource "google_dns_record_set" "aaaa" {
#   name = "${google_dns_managed_zone.default.dns_name}"
#   type = "AAAA"
#   ttl  = 300

#   managed_zone = google_dns_managed_zone.default.name

#   rrdatas = [google_compute_global_address.default.address]
# }

