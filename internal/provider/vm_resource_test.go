package provider

import (
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVMResource(t *testing.T) {

	name, err := helper.NewNanoID(12)

	if err != nil {
		return
	}

	vmConfig := fmt.Sprintf(`
resource "cudo_vm" "test-vm" {
   machine_type       = "standard"
   datacenter_id      = "black-mesa"
   vcpu_quantity      = 1
   boot_disk_size_gib = 50
   image_id           = "ubuntu-minimal-2004"
   memory_gib         = 4
   vm_id              = "%s"
   boot_disk_class    = "network"
 }`, name)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: getProviderConfig() + vmConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "boot_disk_class", "network"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "boot_disk_size_gib", "50"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "machine_type", "standard"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "datacenter_id", "black-mesa"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "gpu_model", ""),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "image_id", "ubuntu-minimal-2004"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "memory_gib", "4"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "one_state", "INIT"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "price_hr"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "public_ip_address"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "external_ip_address"),
					resource.TestCheckResourceAttrSet("cudo_vm.test-vm", "internal_address"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "region_id", "gb-bournemouth"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "region_name", "Bournemouth, United Kingdom"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "renewable_energy", "true"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "vcpu_quantity", "1"),
					resource.TestCheckResourceAttr("cudo_vm.test-vm", "vm_id", name),
				),
			},
		},
	})
}
