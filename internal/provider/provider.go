package provider

import (
	"context"
	"github.com/CudoVentures/cudo-terraform-provider-pf/internal/client"
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
				MarkdownDescription: "Example provider attribute",
				Optional:            true,
			},
		},
	}
}

func (p *CudoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data CudoProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	// if data.Endpoint.IsNull() { /* ... */ }
	transport := httptransport.New("rest.staging.compute.cudo.org", "", []string{"https"})
	transport.DefaultAuthentication = httptransport.BearerToken("545fb7f3fa7841b72861cc0e1eeaf2a200d927a19dd663a8c6512081c81f4f9f")
	transport.Debug = true

	tx := httptransport.New("rest.staging.compute.cudo.org", client.DefaultBasePath, client.DefaultSchemes)
	tx.DefaultAuthentication = httptransport.BearerToken(os.Getenv("API_ACCESS_TOKEN"))
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
		NewExampleDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CudoProvider{
			version: version,
		}
	}
}
