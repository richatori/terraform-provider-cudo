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
var _ datasource.DataSource = &SecurityGroupsDataSource{}

func NewSecurityGroupsDataSource() datasource.DataSource {
	return &SecurityGroupsDataSource{}
}

// SecurityGroupsDataSource defines the data source implementation.
type SecurityGroupsDataSource struct {
	client *CudoClientData
}

// SecurityGroupsDataSourceModel describes the data source data model.
type SecurityGroupsDataSourceModel struct {
	SecurityGroupModels []SecurityGroupResourceModel `tfsdk:"security_groups"`
	ID                  types.String                 `tfsdk:"id"`
	DataCenterId        types.String                 `tfsdk:"datacenter_id"`
}

func (d *SecurityGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_security_groups"
}

func (d *SecurityGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Security groups data source",
		Description:         "Fetches the list of security groups",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"datacenter_id": schema.StringAttribute{
				MarkdownDescription: "Datacenter ID to request security groups from",
				Required:            true,
			},
			"security_groups": schema.ListNestedAttribute{
				Description: "List of security groups",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "Image identifier",
							Computed:            true,
						},
						"datacenter_id": schema.StringAttribute{
							MarkdownDescription: "Datacenter ID",
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: "Security group description",
							Computed:            true,
						},
						"rules": schema.ListNestedAttribute{
							Description: "List of rules in security group",
							Computed:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"icmp_type": schema.StringAttribute{
										MarkdownDescription: "ICMP type",
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: "Rule ID",
										Computed:            true,
									},
									"ip_range_cidr": schema.StringAttribute{
										MarkdownDescription: "IP range",
										Computed:            true,
									},
									"ports": schema.StringAttribute{
										MarkdownDescription: "Image size in GiB",
										Computed:            true,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: "Image size in GiB",
										Computed:            true,
									},
									"rule_type": schema.StringAttribute{
										MarkdownDescription: "Image size in GiB",
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *SecurityGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *SecurityGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SecurityGroupsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewListSecurityGroupsParams()
	params.ProjectID = d.client.DefaultProjectID
	params.DataCenterID = state.DataCenterId.ValueStringPointer()
	res, err := d.client.Client.Networks.ListSecurityGroups(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read security groups",
			err.Error(),
		)
		return
	}

	for _, sg := range res.Payload.SecurityGroups {
		rules := getRuleModels(sg.Rules)

		sgModel := SecurityGroupResourceModel{
			Id:           types.StringValue(*sg.ID),
			DataCenterId: types.StringValue(*sg.DataCenterID),
			Description:  types.StringValue(sg.Description),
			Rules:        rules,
		}

		state.SecurityGroupModels = append(state.SecurityGroupModels, sgModel)
	}

	state.ID = types.StringValue("placeholder")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
