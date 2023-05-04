package provider

import (
	"fmt"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"testing"
	"time"
)

func TestAccVMInstanceDataSource(t *testing.T) {

	name, err := helper.NewNanoID(12)

	if err != nil {
		return
	}

	resourcesConfig := fmt.Sprintf(`
resource "cudo_vm" "my-vm" {
   config_id          = "oaml6hca4fb0"
   vcpu_quantity      = 1
   boot_disk_size_gib = 50
   image_id           = "ubuntu-minimal-2004"
   memory_gib         = 4
   vm_id              = "%s"
   boot_disk_class    = "network"
 }`, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			time.Sleep(time.Second * 10)

			cl := getClient()

			getParams := virtual_machines.NewGetInstanceParams()
			getParams.InstanceID = name
			getParams.ProjectID = project_id

			ins, err := cl.VirtualMachines.GetInstance(getParams)

			if err == nil {
				return fmt.Errorf("vm resource not destroyed %s , %s,  %s, %s", ins.Payload.Compute.Lease.Status, ins.Payload.Compute.Instance.ID, ins.Payload.Compute.Instance.OneState, ins.Payload.Compute.Instance.LcmState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + resourcesConfig,
			},
			{
				PreConfig: func() {
					time.Sleep(time.Second * 10)
				},
				Config: getProviderConfig() + testAccVMInstanceDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_vm_instances.test", "instances.#", "1"),
					resource.TestCheckResourceAttr("data.cudo_vm_instances.test", "instances.0.id", name)),
			},
		},
	})
}

const testAccVMInstanceDataSourceConfig = `
data "cudo_vm_instances" "test" {
}`
