resource "cudo_security_group" "my-sg" {
  id            = "my-sg"
  datacenter_id = "gb-london-1"
  description   = "security group for a purpose"
  rules = [
    {
      ports     = "22,80,443"
      rule_type = "outbound"
      protocol  = "tcp"
    },
    {
      ports     = "22,80,443"
      rule_type = "inbound"
      protocol  = "tcp"
    }
  ]
}