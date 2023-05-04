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
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "id", "placeholder"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.cpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.data_center_id", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.gpu_count", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.gpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.memory_gib", "2"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.order_by", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.page_number", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.page_size", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.region_id", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.storage_gib", "50"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "search_params.vcpu", "1"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.#", "3"),

					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.cpu_model", "AMD EPYC 7262 8-Core Processor"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.data_center_id", "black-mesa"),
					//resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.gpu_memory_gib", "16"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.gpu_model", "GA102GL [RTX A4000]"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.gpu_price_hr", "0.170000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.id", "tbwf26agg51g"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.memory_gib_price_hr", "0.005600"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.storage_gib_price_hr", "0.0001"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.total_gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.total_memory_price_hr", "0.011200"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.total_price_hr", "0.037200"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.total_storage_price_hr", "0.0050"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.total_vcpu_price_hr", "0.021000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.0.vcpu_price_hr", "0.021000"),

					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.cpu_model", "Intel(R) Xeon(R) CPU E5-2690 v3 @ 2.60GHz"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.gpu_memory_gib", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.gpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.id", "oaml6hca4fb0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.memory_gib_price_hr", "0.011000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.storage_gib_price_hr", "0.0001"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.total_gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.total_memory_price_hr", "0.022000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.total_price_hr", "0.054000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.total_storage_price_hr", "0.0050"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.total_vcpu_price_hr", "0.027000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.1.vcpu_price_hr", "0.027000"),

					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.cpu_model", "Intel(R) Xeon(R) CPU E5-2690 v3 @ 2.60GHz"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.data_center_id", "black-mesa"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.gpu_memory_gib", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.gpu_model", ""),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.id", "p72jorepwn-l"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.memory_gib_price_hr", "0.570000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.storage_gib_price_hr", "0.0001"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.total_gpu_price_hr", "0"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.total_memory_price_hr", "1.140000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.total_price_hr", "2.645000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.total_storage_price_hr", "0.0050"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.total_vcpu_price_hr", "1.500000"),
					resource.TestCheckResourceAttr("data.cudo_vm_configs.test", "vm_configs.2.vcpu_price_hr", "1.500000")),
			},
		},
	})
}

const testAccVMConfigDataSourceConfig = `
data "cudo_vm_configs" "test" {
    search_params =  {
      memory_gib=2
      vcpu=1
      storage_gib=50
     }
}`
