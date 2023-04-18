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

data "cudo_compute_configs" "cfgs" {

}

output "configs" {
value = data.cudo_compute_configs.cfgs
}

data "cudo_vm_instances" "ins" {

}

output "instances" {
value = data.cudo_vm_instances.ins
}