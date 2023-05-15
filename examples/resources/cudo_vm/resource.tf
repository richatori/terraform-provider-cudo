resource "cudo_vm" "my-vm" {
  machine_type       = "standard"
  datacenter         = "gb-bournemouth-1"
  vcpu_quantity      = 1
  boot_disk_size_gib = 50
  image_id           = "ubuntu-minimal-2004"
  memory_gib         = 2
  vm_id              = "terra-vm-1"
  boot_disk_class    = "network"
}