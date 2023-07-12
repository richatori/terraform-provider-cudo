resource "cudo_network" "network" {
  data_center_id = "gb-london-1"
  id             = "my-network"
  ip_range       = "192.168.0.0/24"
}