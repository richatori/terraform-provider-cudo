package provider

import (
	"context"
	"cudo.org/v1/terraform-provider-cudo/internal/client/projects"
	"cudo.org/v1/terraform-provider-cudo/internal/client/virtual_machines"
	"cudo.org/v1/terraform-provider-cudo/internal/models"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
var _ resource.Resource = &VMResource{}
var _ resource.ResourceWithImportState = &VMResource{}

func NewVMResource() resource.Resource {
	return &VMResource{}
}

// VMResource defines the resource implementation.
type VMResource struct {
	client *CudoClientData
}

// VMResourceModel describes the resource data model.
type VMResourceModel struct {
	BootDiskClass types.String `tfsdk:"boot_disk_class"`
	BootDiskSize  types.Int64  `tfsdk:"boot_disk_size_gib"`
	ConfigId      types.String `tfsdk:"config_id"`
	GPUs          types.Int64  `tfsdk:"gpu_quantity"`
	ImageID       types.String `tfsdk:"image_id"`
	Memory        types.Int64  `tfsdk:"memory_gib"`
	Password      types.String `tfsdk:"password"`
	VCPUs         types.Int64  `tfsdk:"vcpu_quantity"`
	VMId          types.String `tfsdk:"vm_id"`
	// Response
	CPUClass     types.String `tfsdk:"cpu_class"`
	CPUModel     types.String `tfsdk:"cpu_model"`
	CreateBy     types.String `tfsdk:"create_by"`
	DatacenterID types.String `tfsdk:"datacenter_id"`
	GpuMem       types.Int64  `tfsdk:"gpu_mem"`
	GpuModel     types.String `tfsdk:"gpu_model"`
	ImageDesc    types.String `tfsdk:"image_desc"`

	ImageName       types.String  `tfsdk:"image_name"`
	LcmState        types.String  `tfsdk:"lcm_state"`
	LocalIPAddress  types.String  `tfsdk:"local_ip_address"`
	OneState        types.String  `tfsdk:"one_state"`
	PriceHr         types.Float64 `tfsdk:"price_hr"`
	PublicIPAddress types.String  `tfsdk:"public_ip_address"`
	RegionID        types.String  `tfsdk:"region_id"`
	RegionName      types.String  `tfsdk:"region_name"`
	RenewableEnergy types.Bool    `tfsdk:"renewable_energy"`
}

func (r *VMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_vm"
}

func (r *VMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM resource",
		Attributes: map[string]schema.Attribute{
			"boot_disk_class": schema.StringAttribute{
				MarkdownDescription: "Storage class for boot disk, either 'local' or 'network'",
				Required:            true,
				Validators:          []validator.String{stringvalidator.OneOf("local", "network")},
			},
			"boot_disk_size_gib": schema.Int64Attribute{
				MarkdownDescription: "Size of the boot disk in GiB",
				Required:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(10)},
			},
			"config_id": schema.StringAttribute{
				MarkdownDescription: "VM config id, from vm config data source",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"gpu_quantity": schema.Int64Attribute{
				MarkdownDescription: "Number of GPUs",
				Optional:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(0), int64validator.AtMost(10)},
			},
			"image_id": schema.StringAttribute{
				MarkdownDescription: "OS image ID on boot disk",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"memory_gib": schema.Int64Attribute{
				MarkdownDescription: "Amount of VM memory in GiB",
				Required:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(1), int64validator.AtMost(1000)},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "VM password",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.LengthBetween(6, 64)},
			},
			"vcpu_quantity": schema.Int64Attribute{
				MarkdownDescription: "Number of VCPUs",
				Required:            true,
				Validators:          []validator.Int64{int64validator.AtLeast(1), int64validator.AtMost(100)},
			},
			"vm_id": schema.StringAttribute{
				MarkdownDescription: "Your chosen VM identifier",
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id e.g. my-vm")},
			},

			"cpu_class": schema.StringAttribute{
				MarkdownDescription: "The class of the CPU.",
				Computed:            true,
			},
			"cpu_model": schema.StringAttribute{
				MarkdownDescription: "The model of the CPU.",
				Computed:            true,
			},
			"create_by": schema.StringAttribute{
				MarkdownDescription: "The name of the user who created the VM instance.",
				Computed:            true,
			},

			"datacenter_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the VM instance is located.",
				Computed:            true,
			},
			"gpu_mem": schema.Int64Attribute{
				MarkdownDescription: "The amount of memory on the GPU.",
				Computed:            true,
			},
			"gpu_model": schema.StringAttribute{
				MarkdownDescription: "The model of the GPU.",
				Computed:            true,
			},
			"image_desc": schema.StringAttribute{
				MarkdownDescription: "The description of the image used to create the VM instance.",
				Computed:            true,
			},
			"image_name": schema.StringAttribute{
				MarkdownDescription: "The name of the image used to create the VM instance.",
				Computed:            true,
			},
			"lcm_state": schema.StringAttribute{
				MarkdownDescription: "The state of the VM instance in the LCM.",
				Computed:            true,
			},
			"local_ip_address": schema.StringAttribute{
				MarkdownDescription: "The local IP address of the VM instance.",
				Computed:            true,
			},
			"one_state": schema.StringAttribute{
				MarkdownDescription: "The state of the VM instance in OpenNebula.",
				Computed:            true,
			},
			"price_hr": schema.Float64Attribute{
				MarkdownDescription: "The price per hour for the VM instance.",
				Computed:            true,
			},
			"public_ip_address": schema.StringAttribute{
				MarkdownDescription: "The public IP address of the VM instance.",
				Computed:            true,
			},
			"region_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the region where the VM instance is located.",
				Computed:            true,
			},
			"region_name": schema.StringAttribute{
				MarkdownDescription: "The name of the region where the VM instance is located.",
				Computed:            true,
			},
			"renewable_energy": schema.BoolAttribute{
				MarkdownDescription: "Whether the VM instance is powered by renewable energy",
				Computed:            true,
			},
		},
	}
}

func (r *VMResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *VMResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *VMResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bootDiskClass := models.NewStorageClass(models.StorageClassSTORAGECLASSUNKNOWN)

	switch state.BootDiskClass.ValueString() {
	case "local":
		bootDiskClass = models.NewStorageClass(models.StorageClassSTORAGECLASSLOCAL)
	case "network":
		bootDiskClass = models.NewStorageClass(models.StorageClassSTORAGECLASSNETWORK)
	}

	params := projects.NewBuyComputeParams()
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ConfigId.ValueString()
	params.Body = projects.BuyComputeBody{
		BootDisk: &models.Disk{
			SizeGib:      int32(state.BootDiskSize.ValueInt64()),
			StorageClass: bootDiskClass,
		},
		GpuQuantity: int32(state.GPUs.ValueInt64()),
		MemoryGib:   int32(state.Memory.ValueInt64()),
		OsID:        state.ImageID.ValueString(),
		Password:    state.Password.ValueString(),
		Quantity:    1,
		Vcpu:        int32(state.VCPUs.ValueInt64()),
		VMID:        state.VMId.ValueString(),
	}

	_, err := r.client.Client.Projects.BuyCompute(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create VM resource",
			err.Error(),
		)
		return
	}

	paramsGet := virtual_machines.NewGetInstanceParams()
	paramsGet.ProjectID = r.client.DefaultProjectID
	paramsGet.InstanceID = state.VMId.ValueString()

	res, err := r.client.Client.VirtualMachines.GetInstance(paramsGet)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create VM resource",
			err.Error(),
		)
		return
	}

	state.CPUClass = types.StringValue(res.Payload.Compute.Instance.CPUClass)
	state.CPUModel = types.StringValue(res.Payload.Compute.Instance.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.Compute.Instance.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.Compute.Instance.DatacenterID)
	state.GpuMem = types.Int64Value(res.Payload.Compute.Instance.GpuMem)
	state.GpuModel = types.StringValue(res.Payload.Compute.Instance.GpuModel)
	state.ImageDesc = types.StringValue(res.Payload.Compute.Instance.ImageDesc)
	state.ImageID = types.StringValue(res.Payload.Compute.Instance.ImageID)
	state.ImageName = types.StringValue(res.Payload.Compute.Instance.ImageName)
	state.LcmState = types.StringValue(res.Payload.Compute.Instance.LcmState)
	state.LocalIPAddress = types.StringValue(res.Payload.Compute.Instance.LocalIPAddress)
	state.OneState = types.StringValue(res.Payload.Compute.Instance.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.Compute.Instance.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.Compute.Instance.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.Compute.Instance.RegionID)
	state.RegionName = types.StringValue(res.Payload.Compute.Instance.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.Compute.Instance.RenewableEnergy)

	tflog.Trace(ctx, "created a vm")

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *VMResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := virtual_machines.NewGetInstanceParams()
	params.ProjectID = r.client.DefaultProjectID
	params.InstanceID = state.VMId.ValueString()

	res, err := r.client.Client.VirtualMachines.GetInstance(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VM resource",
			err.Error(),
		)
		return
	}

	state.CPUClass = types.StringValue(res.Payload.Compute.Instance.CPUClass)
	state.CPUModel = types.StringValue(res.Payload.Compute.Instance.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.Compute.Instance.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.Compute.Instance.DatacenterID)
	state.GpuMem = types.Int64Value(res.Payload.Compute.Instance.GpuMem)
	state.GpuModel = types.StringValue(res.Payload.Compute.Instance.GpuModel)
	state.ImageDesc = types.StringValue(res.Payload.Compute.Instance.ImageDesc)
	state.ImageID = types.StringValue(res.Payload.Compute.Instance.ImageID)
	state.ImageName = types.StringValue(res.Payload.Compute.Instance.ImageName)
	state.LcmState = types.StringValue(res.Payload.Compute.Instance.LcmState)
	state.LocalIPAddress = types.StringValue(res.Payload.Compute.Instance.LocalIPAddress)
	state.OneState = types.StringValue(res.Payload.Compute.Instance.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.Compute.Instance.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.Compute.Instance.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.Compute.Instance.RegionID)
	state.RegionName = types.StringValue(res.Payload.Compute.Instance.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.Compute.Instance.RenewableEnergy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *VMResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := virtual_machines.NewGetInstanceParams()

	params.ProjectID = r.client.DefaultProjectID
	params.InstanceID = state.VMId.ValueString()

	res, err := r.client.Client.VirtualMachines.GetInstance(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update VM resource",
			err.Error(),
		)
		return
	}

	state.CPUClass = types.StringValue(res.Payload.Compute.Instance.CPUClass)
	state.CPUModel = types.StringValue(res.Payload.Compute.Instance.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.Compute.Instance.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.Compute.Instance.DatacenterID)
	state.GpuMem = types.Int64Value(res.Payload.Compute.Instance.GpuMem)
	state.GpuModel = types.StringValue(res.Payload.Compute.Instance.GpuModel)
	state.ImageDesc = types.StringValue(res.Payload.Compute.Instance.ImageDesc)
	state.ImageID = types.StringValue(res.Payload.Compute.Instance.ImageID)
	state.ImageName = types.StringValue(res.Payload.Compute.Instance.ImageName)
	state.LcmState = types.StringValue(res.Payload.Compute.Instance.LcmState)
	state.LocalIPAddress = types.StringValue(res.Payload.Compute.Instance.LocalIPAddress)
	state.OneState = types.StringValue(res.Payload.Compute.Instance.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.Compute.Instance.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.Compute.Instance.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.Compute.Instance.RegionID)
	state.RegionName = types.StringValue(res.Payload.Compute.Instance.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.Compute.Instance.RenewableEnergy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *VMResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := virtual_machines.NewTerminateInstanceParams()
	params.ProjectID = r.client.DefaultProjectID
	params.InstanceID = state.VMId.ValueString()

	_, err := r.client.Client.VirtualMachines.TerminateInstance(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete VM resource",
			err.Error(),
		)
		return
	}

	tflog.Trace(ctx, "deleted a vm")
}

func (r *VMResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
