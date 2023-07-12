resource "cudo_security_group" "my-sg" {
  id             = "my-sg"
  data_center_id = "gb-london-1"
  description    = "security group for a web server"
  rules = [
    {
      ports     = "80"
      rule_type = "outbound"
      protocol  = "tcp"
    },
    {
      ip_range  = "192.168.0.0/24"
      ports     = "22,80,443"
      rule_type = "inbound"
      protocol  = "tcp"
    }
  ]
}
