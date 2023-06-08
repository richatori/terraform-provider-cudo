package provider

import (
	"context"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/path"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"os"

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
	version string
}

// CudoProviderModel describes the provider data model.
type CudoProviderModel struct {
	Endpoint   types.String `tfsdk:"endpoint"`
	APIKey     types.String `tfsdk:"api_key"`
	DisableTLS types.Bool   `tfsdk:"disable_tls"`
	ProjectID  types.String `tfsdk:"project_id"`
}

type CudoClientData struct {
	Client           *client.CudoComputeService
	DefaultProjectID string
}

func (p *CudoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cudo"
	resp.Version = p.version
}

func (p *CudoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "API endpoint",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "Your API key",
				Optional:            true,
			},
			"disable_tls": schema.BoolAttribute{
				MarkdownDescription: "Whether to connect using TLS",
				Optional:            true,
			},
			"project_id": schema.StringAttribute{
				MarkdownDescription: "Which project id to use",
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

	// API Key checks
	api_key := os.Getenv("CUDO_API_KEY")

	if config.APIKey.ValueString() != "" {
		api_key = config.APIKey.ValueString()
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Cudo API Key",
			"The provider cannot create the client without an API KEY please pass it or set the CUDO_API_KEY environment variable.",
		)
	}

	// Endpoint checks

	endpoint := os.Getenv("CUDO_ENDPOINT")

	if config.Endpoint.ValueString() != "" {
		endpoint = config.Endpoint.ValueString()
	}

	if endpoint == "" {
		endpoint = "rest.compute.cudo.org"
	}

	disable_tls := config.DisableTLS.ValueBool()

	// Project

	project_id := os.Getenv("CUDO_PROJECT_ID")

	if config.ProjectID.ValueString() != "" {
		project_id = config.ProjectID.ValueString()
	}

	if project_id == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("project_id"),
			"Missing Cudo project ID",
			"The provider cannot create the client without a project_id please pass it or set the CUDO_PROJECT_ID environment variable.",
		)
	}

	var scheme []string

	if disable_tls {
		scheme = []string{"https"}
	} else {
		scheme = client.DefaultSchemes
	}

	tx := httptransport.New(endpoint, client.DefaultBasePath, scheme)
	tx.DefaultAuthentication = httptransport.BearerToken(api_key)
	clientx := client.New(tx, strfmt.Default)

	ccd := &CudoClientData{
		Client:           clientx,
		DefaultProjectID: project_id,
	}
	resp.DataSourceData = ccd
	resp.ResourceData = ccd
}

func (p *CudoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewVMResource,
		NewSecurityGroupResource,
		NewNetworkResource,
	}
}

func (p *CudoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewImagesDataSource,
		NewRegionsDataSource,
		NewMachineTypeDataSource,
		NewSshKeysDataSource,
		NewVMInstanceDataSource,
		NewSecurityGroupsDataSource,
		NewNetworksDataSource,
		NewNetworkSearchDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CudoProvider{
			version: version,
		}
	}
}
