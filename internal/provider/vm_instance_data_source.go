package provider

import (
	"context"
	client2 "cudo.org/v1/terraform-provider-cudo/internal/client"
	"cudo.org/v1/terraform-provider-cudo/internal/client/virtual_machines"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &VMInstanceDataSource{}

func NewVMInstanceDataSource() datasource.DataSource {
	return &VMInstanceDataSource{}
}

// VMInstanceDataSource defines the data source implementation.
type VMInstanceDataSource struct {
	client *client2.CudoComputeService
}

type VMModel struct {
	Id              types.String  `tfsdk:"id"`
	BootDiskSizeGib types.Int64   `tfsdk:"boot_disk_size_gib"`
	CPUClass        types.String  `tfsdk:"cpu_class"`
	CPUModel        types.String  `tfsdk:"cpu_model"`
	CreateBy        types.String  `tfsdk:"create_by"`
	DatacenterID    types.String  `tfsdk:"datacenter_id"`
	GpuMem          types.Int64   `tfsdk:"gpu_mem"`
	GpuModel        types.String  `tfsdk:"gpu_model"`
	GpuQuantity     types.Int64   `tfsdk:"gpu_quantity"`
	ImageDesc       types.String  `tfsdk:"image_desc"`
	ImageID         types.String  `tfsdk:"image_id"`
	ImageName       types.String  `tfsdk:"image_name"`
	LcmState        types.String  `tfsdk:"lcm_state"`
	LocalIPAddress  types.String  `tfsdk:"local_ip_address"`
	Memory          types.Int64   `tfsdk:"memory"`
	OneState        types.String  `tfsdk:"one_state"`
	PriceHr         types.Float64 `tfsdk:"price_hr"`
	PublicIPAddress types.String  `tfsdk:"public_ip_address"`
	RegionID        types.String  `tfsdk:"region_id"`
	RegionName      types.String  `tfsdk:"region_name"`
	RenewableEnergy types.Bool    `tfsdk:"renewable_energy"`
	Vcpus           types.Int64   `tfsdk:"vcpus"`
}

type LeaseModel struct {
	Id                 types.String    `tfsdk:"id"`
	CreateBy           types.String    `tfsdk:"create_by"`
	CreateTime         strfmt.DateTime `tfsdk:"create_time"`
	EndTime            strfmt.DateTime `tfsdk:"end_time"`
	InfrastructureType types.String    `tfsdk:"infrastructure_type"`
	Status             types.String    `tfsdk:"status"`
}

type InstanceModel struct {
	VM    *VMModel    `tfsdk:"vm"`
	Lease *LeaseModel `tfsdk:"lease"`
}

// VMInstanceDataSourceModel describes the data source data model.
type VMInstanceDataSourceModel struct {
	Instances []VMModel    `tfsdk:"instances"`
	ID        types.String `tfsdk:"id"`
}

func (d *VMInstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_vm_instances"
}

func (d *VMInstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "VMInstance data source",
		Description:         "Fetches the list of VMInstance",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"instances": schema.ListNestedAttribute{
				Description: "List of VM Instances.",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "The unique identifier of the VM instance.",
							Computed:            true,
						},
						"boot_disk_size_gib": schema.Int64Attribute{
							MarkdownDescription: "The size of the boot disk in gibibytes (GiB).",
							Computed:            true,
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
						"gpu_quantity": schema.Int64Attribute{
							MarkdownDescription: "The number of GPUs attached to the VM instance.",
							Computed:            true,
						},
						"image_desc": schema.StringAttribute{
							MarkdownDescription: "The description of the image used to create the VM instance.",
							Computed:            true,
						},
						"image_id": schema.StringAttribute{
							MarkdownDescription: "The unique identifier of the image used to create the VM instance.",
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
						"memory": schema.Int64Attribute{
							MarkdownDescription: "The amount of memory allocated to the VM instance.",
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
						"vcpus": schema.Int64Attribute{
							MarkdownDescription: "",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *VMInstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client2.CudoComputeService)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CudoClientData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *VMInstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state VMInstanceDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	res, err := d.client.VirtualMachines.ListInstances(&virtual_machines.ListInstancesParams{
		ProjectID: "", //d.client.DefaultProjectID,
	})

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read VMInstance list instances",
			err.Error(),
		)
		return
	}

	for _, i := range res.Payload.Instances {
		vmState := VMModel{
			Id:              types.StringValue(i.Instance.ID),
			BootDiskSizeGib: types.Int64Value(i.Instance.BootDiskSizeGib),
			CPUClass:        types.StringValue(i.Instance.CPUClass),
			CPUModel:        types.StringValue(i.Instance.CPUModel),
			CreateBy:        types.StringValue(i.Instance.CreateBy),
			DatacenterID:    types.StringValue(i.Instance.DatacenterID),
			GpuMem:          types.Int64Value(i.Instance.GpuMem),
			GpuModel:        types.StringValue(i.Instance.GpuModel),
			GpuQuantity:     types.Int64Value(i.Instance.GpuQuantity),
			ImageDesc:       types.StringValue(i.Instance.ImageDesc),
			ImageID:         types.StringValue(i.Instance.ImageID),
			ImageName:       types.StringValue(i.Instance.ImageName),
			LcmState:        types.StringValue(i.Instance.LcmState),
			LocalIPAddress:  types.StringValue(i.Instance.LocalIPAddress),
			Memory:          types.Int64Value(i.Instance.Memory),
			OneState:        types.StringValue(i.Instance.OneState),
			PriceHr:         types.Float64Value(float64(i.Instance.PriceHr)),
			PublicIPAddress: types.StringValue(i.Instance.PublicIPAddress),
			RegionID:        types.StringValue(i.Instance.RegionID),
			RegionName:      types.StringValue(i.Instance.RegionName),
			RenewableEnergy: types.BoolValue(i.Instance.RenewableEnergy),
			Vcpus:           types.Int64Value(i.Instance.Vcpus),
		}

		state.Instances = append(state.Instances, vmState)
	}

	state.ID = types.StringValue("placeholder")

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
