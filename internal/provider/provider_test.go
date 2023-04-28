package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cudo": providerserver.NewProtocol6WithError(New("test")()),
}

func getProviderConfig() string {
	apiKey := os.Getenv("TF_TEST_CUDO_API_KEY")

	return fmt.Sprintf(`
provider "cudo" {
  api_key     = "%s"
  endpoint    = "rest.staging.compute.cudo.org"
  project_id  = "terraform-testing"
}
`, apiKey)
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

}
