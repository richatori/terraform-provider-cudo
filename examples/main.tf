terraform {
  required_providers {
    cudo = {
      source = "CudoVentures/cudo"
    }
  }
}

provider "cudo" {
  api_key    = "api-key"
  project_id = "my-project"
}

data "cudo_data_centers" "data_centers" {
}

output "data_centers" {
  value = data.data_centers
}

data "cudo_images" "images" {
}

output "images" {
  value = data.cudo_images.images
}

data "cudo_vm" "vm" {
  id = "test-vm"
}

output "vm_external_ip_address" {
  value = data.cudo_vm.ins.external_ip_address
}

resource "cudo_vm" "my-vm" {
  id             = "terra-vm-1"
  data_center_id = "cudo-ca-montreal-2"
  machine_type   = ""
  vcpus          = 1
  memory_gib     = 2
  boot_disk = {
    image_id = "debian-11"
    size_gib = 25
  }
}

output "new-instance" {
  value = resource.cudo_vm.my-vm
}
