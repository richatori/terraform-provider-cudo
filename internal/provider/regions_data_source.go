package provider

import (
	"context"
	"cudo.org/v1/terraform-provider-cudo/internal/client"
	"cudo.org/v1/terraform-provider-cudo/internal/client/search"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &RegionsDataSource{}

func NewRegionsDataSource() datasource.DataSource {
	return &RegionsDataSource{}
}

// RegionsDataSource defines the data source implementation.
type RegionsDataSource struct {
	client *client.CudoComputeService
}

type RegionsModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// RegionsDataSourceModel describes the data source data model.
type RegionsDataSourceModel struct {
	Regions []RegionsModel `tfsdk:"regions"`
	ID      types.String   `tfsdk:"id"`
}

func (d *RegionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_regions"
}

func (d *RegionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Regions data source",
		Description:         "Fetches the list of regions",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"images": schema.ListNestedAttribute{
				Description: "List of images.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Region identifier",
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: "Region name",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *RegionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.CudoComputeService)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.CudoComputeService, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *RegionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state RegionsDataSourceModel

	res, err := d.client.Search.ListRegions(search.NewListRegionsParams())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read regions",
			err.Error(),
		)
		return
	}

	for _, region := range res.Payload.Regions {
		Regionstate := RegionsModel{
			Id:   types.StringValue(region.ID),
			Name: types.StringValue(region.Name),
		}

		state.Regions = append(state.Regions, Regionstate)
	}

	state.ID = types.StringValue("placeholder")

	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
