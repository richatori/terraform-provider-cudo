package provider

import (
	"context"
	"cudo.org/v1/terraform-provider-cudo/internal/client"
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
}

func (p *CudoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "scaffolding"
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
		},
	}
}

func (p *CudoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config CudoProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.APIKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Cudo API Key",
			"The provider cannot create the client without an API KEY please pass it or set the CUDO_API_KEY environment variable.",
		)
	}

	api_key := os.Getenv("CUDO_API_KEY")

	if !config.APIKey.IsNull() {
		api_key = config.APIKey.ValueString()
	}

	if api_key == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Missing Cudo API Key",
			"The provider cannot create the client without an API KEY please pass it or set the CUDO_API_KEY environment variable.",
		)
	}

	// TODO check endpoint
	// set default for production

	//transport := httptransport.New("rest.staging.compute.cudo.org", "", []string{"https"})
	//transport.DefaultAuthentication = httptransport.BearerToken(api_key)
	//transport.Debug = true

	tx := httptransport.New("rest.staging.compute.cudo.org", client.DefaultBasePath, client.DefaultSchemes)
	tx.DefaultAuthentication = httptransport.BearerToken(api_key)
	clientx := client.New(tx, strfmt.Default)

	resp.DataSourceData = clientx
	resp.ResourceData = clientx
}

func (p *CudoProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *CudoProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewImagesDataSource,
		NewRegionsDataSource,
		NewComputeConfigsDataSource,
		NewSshKeysDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CudoProvider{
			version: version,
		}
	}
}
