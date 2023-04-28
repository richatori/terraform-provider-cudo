package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSSHKeysDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: getProviderConfig() + testAccSSHKeysDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.cudo_ssh_keys.test", "ssh_keys.#", "1"),
					resource.TestCheckResourceAttr("data.cudo_ssh_keys.test", "ssh_keys.0.fingerprint", "SHA256:z+44hPDi60D+WGSuOYDrjO1zczxmJ+Ng1KZY1Na6Gpo"),
					resource.TestCheckResourceAttr("data.cudo_ssh_keys.test", "ssh_keys.0.id", "ybwna1z1jrw6"),
					resource.TestCheckResourceAttr("data.cudo_ssh_keys.test", "ssh_keys.0.type", "ssh-rsa")),
			},
		},
	})
}

const testAccSSHKeysDataSourceConfig = `
data "cudo_ssh_keys" "test" {
}`
