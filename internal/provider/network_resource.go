package provider

import (
	"context"
	"fmt"
	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"regexp"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NetworkResource{}
var _ resource.ResourceWithImportState = &NetworkResource{}

func NewNetworkResource() resource.Resource {
	return &NetworkResource{}
}

// NetworkResource defines the resource implementation.
type NetworkResource struct {
	client *CudoClientData
}

// NetworkResourceModel describes the resource data model.
type NetworkResourceModel struct {
	Id           types.String `tfsdk:"id"`
	DataCenterId types.String `tfsdk:"datacenter_id"`
	CIDRPrefix   types.String `tfsdk:"cidr_prefix"`
	Gateway      types.String `tfsdk:"gateway"`
}

func (r *NetworkResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_network"
}

func (r *NetworkResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Network resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Network ID",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"datacenter_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Required:            true,
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"cidr_prefix": schema.StringAttribute{
				MarkdownDescription: "CIDR prefix i.e. 192.168.0.0/24",
				Required:            true,
			},
			"vrouter_size": schema.StringAttribute{ // temporarily only one size (small) so is computed will be set to required later
				MarkdownDescription: "Size of the vrouter 'small' 'medium' or 'large'",
				Computed:            true,
				Validators:          []validator.String{stringvalidator.OneOf("small")}, //, "medium", "large")},
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "gateway",
				Computed:            true,
			},
		},
	}
}

func (r *NetworkResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *CudoClientData got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *NetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *NetworkResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewCreateNetworkParams()
	params.ProjectID = r.client.DefaultProjectID
	params.Body = networks.CreateNetworkBody{
		CidrPrefix:   state.CIDRPrefix.ValueStringPointer(),
		DataCenterID: state.DataCenterId.ValueStringPointer(),
		NetworkID:    state.Id.ValueStringPointer(),
		VrouterSize:  models.VrouterSizeVROUTERINSTANCESMALL.Pointer(),
	}
	_, err := r.client.Client.Networks.CreateNetwork(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	getParams := networks.NewGetNetworkParams()
	getParams.ID = state.Id.ValueString()
	getParams.ProjectID = r.client.DefaultProjectID

	res, err := r.client.Client.Networks.GetNetwork(getParams)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue(res.Payload.Network.ID)
	state.DataCenterId = types.StringValue(res.Payload.Network.DataCenter)
	state.CIDRPrefix = types.StringValue(res.Payload.Network.CidrPrefix)
	state.Gateway = types.StringValue(res.Payload.Network.Gateway)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *NetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	getParams := networks.NewGetNetworkParams()
	getParams.ID = state.Id.ValueString()
	getParams.ProjectID = r.client.DefaultProjectID

	resget, err := r.client.Client.Networks.GetNetwork(getParams)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read network resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue(resget.Payload.Network.ID)
	state.DataCenterId = types.StringValue(resget.Payload.Network.DataCenter)
	state.CIDRPrefix = types.StringValue(resget.Payload.Network.CidrPrefix)
	state.Gateway = types.StringValue(resget.Payload.Network.Gateway)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *NetworkResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	getParams := networks.NewGetNetworkParams()
	getParams.ID = state.Id.ValueString()
	getParams.ProjectID = r.client.DefaultProjectID

	resget, err := r.client.Client.Networks.GetNetwork(getParams)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create network resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue(resget.Payload.Network.ID)
	state.DataCenterId = types.StringValue(resget.Payload.Network.DataCenter)
	state.CIDRPrefix = types.StringValue(resget.Payload.Network.CidrPrefix)
	state.Gateway = types.StringValue(resget.Payload.Network.Gateway)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *NetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *NetworkResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := networks.NewDeleteNetworkParams()
	params.ProjectID = r.client.DefaultProjectID
	params.NetworkID = state.Id.ValueString()

	_, err := r.client.Client.Networks.DeleteNetwork(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete network resource",
			err.Error(),
		)
		return
	}

	tflog.Trace(ctx, "deleted a network")
}

func (r *NetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
