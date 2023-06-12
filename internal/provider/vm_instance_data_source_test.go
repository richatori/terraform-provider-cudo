package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccVMInstanceDataSource(t *testing.T) {

	name, err := helper.NewNanoID(12)

	if err != nil {
		return
	}

	resourcesConfig := fmt.Sprintf(`
resource "cudo_vm" "my-vm" {
   machine_type       = "standard"
   data_center_id     = "black-mesa"
   vcpus              = 1
   boot_disk_size_gib = 1
   image_id           = "alpine-linux-317"
   memory_gib         = 2
   vm_id              = "%s"
   boot_disk_class    = "network"
 }`, name)

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			time.Sleep(time.Second * 10)

			cl := getClient()

			getParams := virtual_machines.NewGetVMParams()
			getParams.ID = name
			getParams.ProjectID = project_id

			ins, err := cl.VirtualMachines.GetVM(getParams)

			if err == nil {
				return fmt.Errorf("vm resource not destroyed %s , %s,  %s", ins.Payload.VM.ID, ins.Payload.VM.LcmState, ins.Payload.VM.OneState)
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
