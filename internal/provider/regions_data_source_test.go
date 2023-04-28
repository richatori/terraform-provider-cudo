package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRegionsDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: getProviderConfig() + testAccRegionsDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_regions.test", "regions.#", "2"),
					resource.TestCheckResourceAttr("data.cudo_regions.test", "regions.0.id", "no-luster"),
					resource.TestCheckResourceAttr("data.cudo_regions.test", "regions.0.name", "Luster, Norway"),
					resource.TestCheckResourceAttr("data.cudo_regions.test", "regions.1.id", "gb-bournemouth"),
					resource.TestCheckResourceAttr("data.cudo_regions.test", "regions.1.name", "Bournemouth, United Kingdom")),
			},
		},
	})
}

const testAccRegionsDataSourceConfig = `
data "cudo_regions" "test" {
}`
