terraform {
  required_providers {
    cudo = {
      source  = "cudo.org/v1/cudo"
    }
  }
}


provider "cudo" {
  api_key      = "545fb7f3fa7841b72861cc0e1eeaf2a200d927a19dd663a8c6512081c81f4f9f"
  endpoint = "rest.staging.compute.cudo.org"
  disable_tls = true
  project_id = "long-term-test"
}

data "cudo_regions" "reg1" {

}

output "regions" {
  value = data.cudo_regions.reg1
}

data "cudo_images" "img1" {

}

output "images" {
  value = data.cudo_images.img1.images
}

data "cudo_vm_configs" "cfgs" {

}

output "configs" {
value = data.cudo_vm_configs.cfgs
}

data "cudo_vm_instances" "ins" {

}

output "instances" {
value = data.cudo_vm_instances.ins
}

resource "cudo_vm" "my-vm" {
config_id = "oaml6hca4fb0"
vcpu_quantity = 1
boot_disk_size_gib = 50
image_id = "ubuntu-minimal-2004"
memory_gib = 2
vm_id = "terra-vm-1"
boot_disk_class = "network"
}

output "new-instance" {
value = resource.cudo_vm.my-vm
}