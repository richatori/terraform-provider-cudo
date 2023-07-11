package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client"
	"github.com/CudoVentures/terraform-provider-cudo/internal/helper"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

var apiKey = os.Getenv("TF_TEST_CUDO_API_KEY")
var projectID = os.Getenv("TF_TEST_CUDO_PROJECT_ID")
var remoteAddr = "rest.staging.compute.cudo.org"
var billingAccountID = os.Getenv("TF_TEST_CUDO_BILLING_ACCOUNT_ID")

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"cudo": providerserver.NewProtocol6WithError(New("test", remoteAddr)()),
}

func getProviderConfig() string {

	return fmt.Sprintf(`
provider "cudo" {
  api_key            = "%s"
  remote_addr        = "%s"
  project_id         = "%s"
  billing_account_id = "%s"
}
`, apiKey, remoteAddr, projectID, billingAccountID)
}

func testAccPreCheck(t *testing.T) {
	// You can add code here to run prior to any test case execution, for example assertions
	// about the appropriate environment variables being set are common to see in a pre-check
	// function.

}

var testRunID, _ = helper.NewNanoID(6)

func getClient() *client.CudoComputeService {
	tx := httptransport.New(remoteAddr, client.DefaultBasePath, client.DefaultSchemes)
	tx.DefaultAuthentication = httptransport.BearerToken(apiKey)
	// TODO: it would be nice to plug the debug logging into t.Log
	// tx.Debug = true
	clientx := client.New(tx, strfmt.Default)
	return clientx
}
