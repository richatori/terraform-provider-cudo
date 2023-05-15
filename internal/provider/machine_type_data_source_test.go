package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVMConfigDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: getProviderConfig() + testAccVMConfigDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "id", "placeholder"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.cpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.data_center_id", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.gpu_count", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.gpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.memory_gib", "2"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.order_by", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.page_number", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.page_size", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.region_id", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.storage_gib", "50"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "search_params.vcpu", "1"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.#", "2"),

					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.cpu_model", "EPYC-Rome"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.gpu_model", "RTX A4000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.gpu_price_hr", "0.170000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.id", "epyc-rtx-a4000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.memory_gib_price_hr", "0.005600"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.storage_gib_price_hr", "0.0001"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.total_gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.total_memory_price_hr", "0.011200"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.total_price_hr", "0.037200"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.total_storage_price_hr", "0.0050"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.total_vcpu_price_hr", "0.021000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.0.vcpu_price_hr", "0.021000"),

					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.cpu_model", "Haswell-noTSX-IBRS"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.gpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.id", "standard"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.memory_gib_price_hr", "0.011000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.storage_gib_price_hr", "0.0001"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.total_gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.total_memory_price_hr", "0.022000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.total_price_hr", "0.054000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.total_storage_price_hr", "0.0050"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.total_vcpu_price_hr", "0.027000"),
					resource.TestCheckResourceAttr("data.cudo_machine_types.test", "machine_types.1.vcpu_price_hr", "0.027000")),
			},
		},
	})
}

const testAccVMConfigDataSourceConfig = `
data "cudo_machine_types" "test" {
    search_params =  {
      memory_gib=2
      vcpu=1
      storage_gib=50
     }
}`
