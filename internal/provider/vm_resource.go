package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/virtual_machines"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
	Id            types.String   `tfsdk:"id"`
	BootDiskClass types.String   `tfsdk:"boot_disk_class"`
	BootDiskSize  types.Int64    `tfsdk:"boot_disk_size_gib"`
	MachineType   types.String   `tfsdk:"machine_type"`
	GPUs          types.Int64    `tfsdk:"gpu_quantity"`
	ImageID       types.String   `tfsdk:"image_id"`
	Memory        types.Int64    `tfsdk:"memory_gib"`
	Password      types.String   `tfsdk:"password"`
	VCPUs         types.Int64    `tfsdk:"vcpu_quantity"`
	ID            types.String   `tfsdk:"vm_id"`
	SSHKeySource  types.String   `tfsdk:"ssh_key_source"`
	SSHKeysCustom []types.String `tfsdk:"ssh_keys_custom"`
	StartScript   types.String   `tfsdk:"start_script"`
	// Response
	CPUModel          types.String  `tfsdk:"cpu_model"`
	CreateBy          types.String  `tfsdk:"create_by"`
	DatacenterID      types.String  `tfsdk:"datacenter_id"`
	GpuModel          types.String  `tfsdk:"gpu_model"`
	LcmState          types.String  `tfsdk:"lcm_state"`
	InternalIPAddress types.String  `tfsdk:"internal_ip_address"`
	ExternalIPAddress types.String  `tfsdk:"external_ip_address"`
	OneState          types.String  `tfsdk:"one_state"`
	PriceHr           types.Float64 `tfsdk:"price_hr"`
	PublicIPAddress   types.String  `tfsdk:"public_ip_address"`
	RegionID          types.String  `tfsdk:"region_id"`
	RegionName        types.String  `tfsdk:"region_name"`
	RenewableEnergy   types.Bool    `tfsdk:"renewable_energy"`
}

func (r *VMResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_vm"
}

func (r *VMResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VM resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "placeholder",
				Computed:            true,
			},
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
			"machine_type": schema.StringAttribute{
				MarkdownDescription: "VM machine type, from machine type data source",
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
				Required:            true,
				Validators:          []validator.String{stringvalidator.RegexMatches(regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"gpu_model": schema.StringAttribute{
				MarkdownDescription: "The model of the GPU.",
				Computed:            true,
			},
			"lcm_state": schema.StringAttribute{
				MarkdownDescription: "The state of the VM instance in the LCM.",
				Computed:            true,
			},
			"internal_ip_address": schema.StringAttribute{
				MarkdownDescription: "The internal IP address of the VM instance.",
				Computed:            true,
			},
			"external_ip_address": schema.StringAttribute{
				MarkdownDescription: "The external IP address of the VM instance.",
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
			"ssh_key_source": schema.StringAttribute{
				MarkdownDescription: "Which SSH keys to add to the VM: user (default), project or custom",
				Optional:            true,
				Validators:          []validator.String{stringvalidator.OneOf("user", "project", "custom")},
			},
			"ssh_keys_custom": schema.ListAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "List of custom SSH keys to add to the VM, ssh_key_source must be set to custom",
				Optional:            true,
			},
			"start_script": schema.StringAttribute{
				MarkdownDescription: "A script to run when VM boots",
				Optional:            true,
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

	sshKeySource := "SSH_KEY_SOURCE_USER"
	switch state.SSHKeySource.ValueString() {
	case "project":
		sshKeySource = "SSH_KEY_SOURCE_PROJECT"
	case "custom":
		sshKeySource = "SSH_KEY_SOURCE_NONE"
	}

	ks := models.SSHKeySource(sshKeySource)

	var customKeys []string
	for _, key := range state.SSHKeysCustom {
		customKeys = append(customKeys, key.ValueString())
	}

	params := virtual_machines.NewCreateVMParams()
	params.ProjectID = r.client.DefaultProjectID
	params.Body = virtual_machines.CreateVMBody{
		BootDisk: &models.Disk{
			SizeGib:      int32(state.BootDiskSize.ValueInt64()),
			StorageClass: bootDiskClass,
		},
		DataCenterID:    state.DatacenterID.ValueString(),
		Gpus:            int32(state.GPUs.ValueInt64()),
		MachineType:     state.MachineType.ValueString(),
		MemoryGib:       int32(state.Memory.ValueInt64()),
		BootDiskImageID: state.ImageID.ValueStringPointer(),
		Password:        state.Password.ValueString(),
		Vcpus:           int32(state.VCPUs.ValueInt64()),
		VMID:            state.ID.ValueStringPointer(),
		SSHKeySource:    &ks,
		CustomSSHKeys:   customKeys,
		StartScript:     state.StartScript.ValueString(),
	}

	_, err := r.client.Client.VirtualMachines.CreateVM(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create VM resource",
			err.Error(),
		)
		return
	}

	paramsGet := virtual_machines.NewGetVMParams()
	paramsGet.ProjectID = r.client.DefaultProjectID
	paramsGet.ID = state.ID.ValueString()

	res, err := r.client.Client.VirtualMachines.GetVM(paramsGet)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create VM resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue("placeholder")
	state.CPUModel = types.StringValue(res.Payload.VM.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.VM.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.VM.DatacenterID)
	state.GpuModel = types.StringValue(res.Payload.VM.GpuModel)
	state.ImageID = types.StringValue(res.Payload.VM.ImageID)
	state.LcmState = types.StringValue(res.Payload.VM.LcmState)
	state.InternalIPAddress = types.StringValue(res.Payload.VM.InternalIPAddress)
	state.ExternalIPAddress = types.StringValue(res.Payload.VM.ExternalIPAddress)
	state.OneState = types.StringValue(res.Payload.VM.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.VM.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.VM.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.VM.RegionID)
	state.RegionName = types.StringValue(res.Payload.VM.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.VM.RenewableEnergy)

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

	params := virtual_machines.NewGetVMParams()
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	res, err := r.client.Client.VirtualMachines.GetVM(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VM resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue("placeholder")
	state.CPUModel = types.StringValue(res.Payload.VM.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.VM.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.VM.DatacenterID)
	state.GpuModel = types.StringValue(res.Payload.VM.GpuModel)
	state.ImageID = types.StringValue(res.Payload.VM.ImageID)
	state.LcmState = types.StringValue(res.Payload.VM.LcmState)
	state.InternalIPAddress = types.StringValue(res.Payload.VM.InternalIPAddress)
	state.ExternalIPAddress = types.StringValue(res.Payload.VM.ExternalIPAddress)
	state.OneState = types.StringValue(res.Payload.VM.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.VM.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.VM.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.VM.RegionID)
	state.RegionName = types.StringValue(res.Payload.VM.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.VM.RenewableEnergy)

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

	params := virtual_machines.NewGetVMParams()

	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	res, err := r.client.Client.VirtualMachines.GetVM(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update VM resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue("placeholder")
	state.CPUModel = types.StringValue(res.Payload.VM.CPUModel)
	state.CreateBy = types.StringValue(res.Payload.VM.CreateBy)
	state.DatacenterID = types.StringValue(res.Payload.VM.DatacenterID)
	state.GpuModel = types.StringValue(res.Payload.VM.GpuModel)
	state.ImageID = types.StringValue(res.Payload.VM.ImageID)
	state.LcmState = types.StringValue(res.Payload.VM.LcmState)
	state.InternalIPAddress = types.StringValue(res.Payload.VM.InternalIPAddress)
	state.ExternalIPAddress = types.StringValue(res.Payload.VM.ExternalIPAddress)
	state.OneState = types.StringValue(res.Payload.VM.OneState)
	state.PriceHr = types.Float64Value(float64(res.Payload.VM.PriceHr))
	state.PublicIPAddress = types.StringValue(res.Payload.VM.PublicIPAddress)
	state.RegionID = types.StringValue(res.Payload.VM.RegionID)
	state.RegionName = types.StringValue(res.Payload.VM.RegionName)
	state.RenewableEnergy = types.BoolValue(res.Payload.VM.RenewableEnergy)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *VMResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *VMResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := virtual_machines.NewTerminateVMParams()
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.ID.ValueString()

	_, err := r.client.Client.VirtualMachines.TerminateVM(params)

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
