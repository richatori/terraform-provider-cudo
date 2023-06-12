terraform {
  required_providers {
    cudo = {
      source = "CudoVentures/cudo"
    }
  }
}

provider "cudo" {
  api_key        = "api-key"
  project_id     = "my-project"
  data_center_id = "gb-london-1"
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
  search_params = {
    memory_gib = 4
  }
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
  config_id          = data.cudo_vm_configs.cfgs.vm_configs[0].id
  vcpus              = 1
  boot_disk_size_gib = 50
  image_id           = "ubuntu-minimal-2004"
  memory_gib         = 2
  vm_id              = "terra-vm-1"
  boot_disk_class    = "network"
}

output "new-instance" {
  value = resource.cudo_vm.my-vm
}
