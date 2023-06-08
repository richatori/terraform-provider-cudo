package provider

import (
	"context"
	"fmt"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NetworksDataSource{}

func NewNetworksDataSource() datasource.DataSource {
	return &NetworksDataSource{}
}

// NetworksDataSource defines the data source implementation.
type NetworksDataSource struct {
	client *CudoClientData
}

// NetworksDataSourceModel describes the data source data model.
type NetworksDataSourceModel struct {
	Network []NetworkResourceModel `tfsdk:"networks"`
	ID      types.String           `tfsdk:"id"`
}

func (d *NetworksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_networks"
}

func (d *NetworksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Networks data source",
		Description:         "Fetches the list of networks",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"networks": schema.ListNestedAttribute{
				Description: "List of networks.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Network ID",
							Computed:            true,
						},
						"datacenter_id": schema.StringAttribute{
							MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
							Computed:            true,
						},
						"cidr_prefix": schema.StringAttribute{
							MarkdownDescription: "CIDR prefix i.e. 192.168.0.0/24",
							Computed:            true},
						"vrouter_size": schema.StringAttribute{
							MarkdownDescription: "Size of the vrouter 'small' 'medium' or 'large'",
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: "gateway",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CudoClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *NetworksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state NetworksDataSourceModel

	params := networks.NewListNetworksParams()
	params.ProjectID = d.client.DefaultProjectID

	res, err := d.client.Client.Networks.ListNetworks(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read networks",
			err.Error(),
		)
		return
	}

	for _, net := range res.Payload.Networks {
		networkModel := NetworkResourceModel{
			Id:           types.StringValue(net.ID),
			DataCenterId: types.StringValue(net.DataCenter),
			CIDRPrefix:   types.StringValue(net.CidrPrefix),
			Gateway:      types.StringValue(net.Gateway),
		}

		state.Network = append(state.Network, networkModel)
	}

	state.ID = types.StringValue("placeholder")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
