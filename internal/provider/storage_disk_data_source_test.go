package provider

import (
	"context"
	"fmt"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/disks"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAcc_StorageDiskDataSource(t *testing.T) {
	var cancel context.CancelFunc
	ctx := context.Background()
	deadline, ok := t.Deadline()
	if ok {
		ctx, cancel = context.WithDeadline(ctx, deadline)
		defer cancel()
	}
	name := "tf-ds-test-" + testRunID

	resourcesConfig := fmt.Sprintf(`
resource "cudo_storage_disk" "disk" {
	data_center_id = "black-mesa"
	id = "%s"
	size_gib = 15
}`, name)

	testAccVMInstanceDataSourceConfig := fmt.Sprintf(`
data "cudo_storage_disk" "test" {
	id = "%s"
}`, name)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		CheckDestroy: func(state *terraform.State) error {
			cl := getClient()

			getParams := disks.NewGetDiskParamsWithContext(ctx)
			getParams.ID = name
			getParams.ProjectID = projectID

			ins, err := cl.Disks.GetDisk(getParams)

			if err == nil && ins.Payload.Disk.DiskState != "dele" {
				terminateParams := disks.NewDeleteStorageDiskParamsWithContext(ctx)
				terminateParams.ID = name
				terminateParams.ProjectID = projectID
				res, err := cl.Disks.DeleteStorageDisk(terminateParams)
				t.Logf("(%s) %#v: %v", ins.Payload.Disk.DiskState, res, err)

				return fmt.Errorf("disk resource not destroyed %s , %s", *ins.Payload.Disk.ID, ins.Payload.Disk.DiskState)
			}
			return nil
		},

		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + resourcesConfig,
			},
			{
				Config: getProviderConfig() + resourcesConfig + testAccVMInstanceDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_storage_disk.test", "id", name),
					resource.TestCheckResourceAttr("data.cudo_storage_disk.test", "size_gib", "15"),
				),
			},
		},
	})
}
