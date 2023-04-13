terraform {
  required_providers {
    cudo = {
      source  = "cudo.org/v1/cudo"
    }
  }
}


provider "cudo" {
  api_key      = "545fb7f3fa7841b72861cc0e1eeaf2a200d927a19dd663a8c6512081c81f4f9f"
#  api_key = "bb2844fdf6c70535f3e882bbb7241e77e6d5fd69f63cc2208488388b4a9110fe"
  endpoint = "127.0.0.1:9000"
  disable_tls = true
}
#
# data "cudo_regions" "reg1" {
#
# }
#
# output "regions" {
#   value = data.cudo_regions.reg1
# }
#
# data "cudo_images" "img1" {
#
# }
#
# output "images" {
#   value = data.cudo_images.img1.images
# }

data "cudo_compute_configs" "cfgs" {

}

output "configs" {
value = data.cudo_compute_configs.cfgs
}