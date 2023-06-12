package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cudo": providerserver.NewProtocol6WithError(New("test", endpoint)()),
}

var billing_account_id = "billing-account"
var data_center_id = "black-mesa"
var project_id = "terraform-testing"
var endpoint = "rest.staging.compute.cudo.org"
var apiKey string

func getProviderConfig() string {
	apiKey = os.Getenv("TF_TEST_CUDO_API_KEY")

	return fmt.Sprintf(`
provider "cudo" {
  api_key            = "%s"
  remote_addr        = "%s"
  project_id         = "%s"
  data_center_id     = "%s"
  billing_account_id = "%s"
}
`, apiKey, endpoint, project_id, data_center_id, billing_account_id)
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

}

func getClient() *client.CudoComputeService {
	tx := httptransport.New(endpoint, client.DefaultBasePath, client.DefaultSchemes)
	tx.DefaultAuthentication = httptransport.BearerToken(apiKey)
	clientx := client.New(tx, strfmt.Default)
	return clientx
}
