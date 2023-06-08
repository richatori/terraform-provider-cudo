package provider

import (
	"context"
	"fmt"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NetworkSearchDataSource{}

func NewNetworkSearchDataSource() datasource.DataSource {
	return &NetworkSearchDataSource{}
}

// NetworkSearchDataSource defines the data source implementation.
type NetworkSearchDataSource struct {
	client *CudoClientData
}

// NetworkSearchDataSourceModel describes the data source data model.
type NetworkSearchDataSourceModel struct {
	ID           types.String           `tfsdk:"id"`
	Network      []NetworkResourceModel `tfsdk:"networks"`
	DataCenterId types.String           `tfsdk:"datacenter_id"`
	MachineType  types.String           `tfsdk:"machine_type"`
}

func (d *NetworkSearchDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_network_search"
}

func (d *NetworkSearchDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Network search data source",
		Description:         "Searches networks available to cluster",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"machine_type": schema.StringAttribute{
				Description: "Filter search by machine type",
				Optional:    true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"datacenter_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Optional:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"networks": schema.ListNestedAttribute{
				Description: "Networks search results",
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

func (d *NetworkSearchDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *NetworkSearchDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state NetworkSearchDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewSearchNetworksParams()
	params.ProjectID = d.client.DefaultProjectID
	params.MachineType = state.MachineType.ValueStringPointer()
	params.DataCenterID = state.DataCenterId.ValueStringPointer()

	res, err := d.client.Client.Networks.SearchNetworks(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read network_search",
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
