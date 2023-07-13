resource "cudo_vm_image" "my-image" {
  id = "my-image"
  source = {
    vm_id = "my-vm-id"
  }
}