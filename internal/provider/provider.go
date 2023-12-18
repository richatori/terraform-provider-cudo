package provider

import (
	"context"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"

	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure CudoProvider satisfies various provider interfaces.
var _ provider.Provider = &CudoProvider{}

// CudoProvider defines the provider implementation.
type CudoProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version           string
	defaultRemoteAddr string
}

// CudoProviderModel describes the provider data model.
type CudoProviderModel struct {
	APIKey           types.String `tfsdk:"api_key"`
	DisableTLS       types.Bool   `tfsdk:"disable_tls"`
	RemoteAddr       types.String `tfsdk:"remote_addr"`
	ProjectID        types.String `tfsdk:"project_id"`
	BillingAccountID types.String `tfsdk:"billing_account_id"`
}

type CudoClientData struct {
	Client                  *client.CudoComputeService
	DefaultBillingAccountID string
	DefaultProjectID        string
}

func (p *CudoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cudo"
	resp.Version = p.version
}

func (p *CudoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Your API key",
				Optional:            true,
			},
			"billing_account_id": schema.StringAttribute{
				MarkdownDescription: "Which billing account id to create resources in",
				Optional:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Which project id to use",
				Optional:            true,
			},
			"remote_addr": schema.StringAttribute{
				MarkdownDescription: "API endpoint",
				Optional:            true,
			},
			"disable_tls": schema.BoolAttribute{
				MarkdownDescription: "Whether to connect the API endpoint using TLS",
				Optional:            true,
			},
		},
	}
}

func (p *CudoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config CudoProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	apiKey := config.APIKey.ValueString()
	if apiKey == "" {
		apiKey = os.Getenv("CUDO_API_KEY")
	}
	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Cudo API Key",
			"The provider cannot create the client without an API KEY please configure it, use the CUDO_API_KEY environment variable or set it in your cudo config file.",
		)
	}

	// Endpoint checks
	remoteAddr := config.RemoteAddr.ValueString()
	if remoteAddr == "" {
		remoteAddr = os.Getenv("CUDO_REMOTE_ADDR")
	}
	if remoteAddr == "" {
		remoteAddr = p.defaultRemoteAddr
	}

	// Project
	projectID := config.ProjectID.ValueString()
	if projectID == "" {
		projectID = os.Getenv("CUDO_PROJECT_ID")
	}
	if projectID == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("project_id"),
			"Missing Cudo project ID",
			"The provider cannot create the client without a project_id please pass it or set the CUDO_PROJECT_ID environment variable or set it in your cudo config file.",
		)
	}

	billingAccountID := config.BillingAccountID.ValueString()
	if billingAccountID == "" {
		billingAccountID = os.Getenv("CUDO_BILLING_ACCOUNT_ID")
	}

	var scheme []string
	if config.DisableTLS.ValueBool() {
		scheme = []string{"https"}
	} else {
		scheme = client.DefaultSchemes
	}

	tx := httptransport.New(remoteAddr, client.DefaultBasePath, scheme)
	tx.DefaultAuthentication = httptransport.BearerToken(apiKey)
	clientx := client.New(tx, strfmt.Default)

	ccd := &CudoClientData{
		Client:                  clientx,
		DefaultProjectID:        projectID,
		DefaultBillingAccountID: billingAccountID,
	}
	resp.DataSourceData = ccd
	resp.ResourceData = ccd
}

func (p *CudoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewStorageDiskResource,
		NewSecurityGroupResource,
		NewNetworkResource,
		NewVMImageResource,
		NewVMResource,
	}
}

func (p *CudoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewStorageDiskDataSource,
		NewVMImagesDataSource,
		NewVMDataCentersDataSource,
		NewVMDataSource,
		NewSecurityGroupDataSource,
		NewNetworkDataSource,
	}
}

func New(version string, defaultRemoteAddr string) func() provider.Provider {
	return func() provider.Provider {
		return &CudoProvider{
			version:           version,
			defaultRemoteAddr: defaultRemoteAddr,
		}
	}
}
