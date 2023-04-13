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
var _ datasource.DataSource = &ComputeConfigsDataSource{}

func NewComputeConfigsDataSource() datasource.DataSource {
	return &ComputeConfigsDataSource{}
}

// ComputeConfigsDataSource defines the data source implementation.
type ComputeConfigsDataSource struct {
	client *client.CudoComputeService
}

type ComputeConfigsModel struct {
	Id                  types.String `tfsdk:"id"`
	CpuModel            types.String `tfsdk:"cpu_model"`
	DataCenterId        types.String `tfsdk:"data_center_id"`
	GpuMemoryGib        types.Int64  `tfsdk:"gpu_memory_gib"`
	GpuModel            types.String `tfsdk:"gpu_model"`
	GpuPriceHr          types.String `tfsdk:"gpu_price_hr"`
	MemoryGibPriceHr    types.String `tfsdk:"memory_gib_price_hr"`
	StorageGibPriceHr   types.String `tfsdk:"storage_gib_price_hr"`
	TotalGpuPriceHr     types.String `tfsdk:"total_gpu_price_hr"`
	TotalMemoryPriceHr  types.String `tfsdk:"total_memory_price_hr"`
	TotalPriceHr        types.String `tfsdk:"total_price_hr"`
	TotalStoragePriceHr types.String `tfsdk:"total_storage_price_hr"`
	TotalVcpuPriceHr    types.String `tfsdk:"total_vcpu_price_hr"`
	VcpuPriceHr         types.String `tfsdk:"vcpu_price_hr"`
	CountVmAvailable    types.Int64  `tfsdk:"count_vm_available"`
}

type SearchParamsModel struct {
	CpuModel     types.String `tfsdk:"cpu_model"`
	DataCenterID types.String `tfsdk:"data_center_id"`
	GpuCount     types.Int64  `tfsdk:"gpu_count"`
	GpuModel     types.String `tfsdk:"gpu_model"`
	MemoryGiB    types.Int64  `tfsdk:"memory_gib"`
	OrderBy      types.String `tfsdk:"order_by"`
	PageNumber   types.Int64  `tfsdk:"page_number"`
	PageSize     types.Int64  `tfsdk:"page_size"`
	RegionID     types.String `tfsdk:"region_id"`
	StorageGiB   types.Int64  `tfsdk:"storage_gib"`
	VCPU         types.Int64  `tfsdk:"vcpu"`
}

// ComputeConfigsDataSourceModel describes the data source data model.
type ComputeConfigsDataSourceModel struct {
	ComputeConfigs []ComputeConfigsModel `tfsdk:"compute_configs"`
	ID             types.String          `tfsdk:"id"`
	SearchParams   *SearchParamsModel    `tfsdk:"search_params"`
}

func (d *ComputeConfigsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_compute_configs"
}

func (d *ComputeConfigsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "ComputeConfigs data source",
		Description:         "Fetches the list of compute_configs",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Placeholder identifier attribute.",
				Computed:    true,
			},
			"search_params": schema.SingleNestedAttribute{
				Description: "Search parameters for limiting compute types",
				Computed:    false,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"cpu_model": schema.StringAttribute{
						MarkdownDescription: "CPU model name",
						Computed:            false,
						Optional:            true,
					},
					"data_center_id": schema.StringAttribute{
						MarkdownDescription: "ID of the data center where the VM is located",
						Computed:            false,
						Optional:            true,
					},
					"gpu_count": schema.Int64Attribute{
						MarkdownDescription: "Number of GPUs",
						Computed:            false,
						Optional:            true,
					},
					"gpu_model": schema.StringAttribute{
						MarkdownDescription: "GPU model name",
						Computed:            false,
						Optional:            true,
					},
					"memory_gib": schema.Int64Attribute{
						MarkdownDescription: "Amount of memory in GiB",
						Computed:            false,
						Optional:            true,
					},
					"order_by": schema.StringAttribute{
						MarkdownDescription: "Field to order results by",
						Computed:            false,
						Optional:            true,
					},
					"page_number": schema.Int64Attribute{
						MarkdownDescription: "Page number of the results to return",
						Computed:            false,
						Optional:            true,
					},
					"page_size": schema.Int64Attribute{
						MarkdownDescription: "Number of results per page",
						Computed:            false,
						Optional:            true,
					},
					"region_id": schema.StringAttribute{
						MarkdownDescription: "ID of the region to search in",
						Computed:            false,
						Optional:            true,
					},
					"storage_gib": schema.Int64Attribute{
						MarkdownDescription: "Amount of storage in GiB",
						Computed:            false,
						Optional:            true,
					},
					"vcpu": schema.Int64Attribute{
						MarkdownDescription: "Number of vCPUs",
						Computed:            false,
						Optional:            true,
					},
				},
			},
			"compute_configs": schema.ListNestedAttribute{
				Description: "List of available vm configurations",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cpu_model": schema.StringAttribute{
							MarkdownDescription: "CPU model name",
							Computed:            true,
						},
						//"cpu_model_category": schema.StringAttribute{
						//	MarkdownDescription: "CPU model category",
						//	Computed:            true,
						//},
						"data_center_id": schema.StringAttribute{
							MarkdownDescription: "ID of the data center where the VM is located",
							Computed:            true,
						},
						"gpu_memory_gib": schema.Int64Attribute{
							MarkdownDescription: "Amount of GPU memory in GiB",
							Computed:            true,
						},
						"gpu_model": schema.StringAttribute{
							MarkdownDescription: "GPU model name",
							Computed:            true,
						},
						//"gpu_model_category": schema.StringAttribute{
						//	MarkdownDescription: "GPU model category",
						//	Computed:            true,
						//},
						"gpu_price_hr": schema.StringAttribute{
							MarkdownDescription: "Price per GPU per hour",
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: "Compute config identifier",
							Computed:            true,
						},
						"memory_gib_price_hr": schema.StringAttribute{
							MarkdownDescription: "Price per GiB of memory per hour",
							Computed:            true,
						},
						"storage_gib_price_hr": schema.StringAttribute{
							MarkdownDescription: "Price per GiB of storage per hour",
							Computed:            true,
						},
						"total_gpu_price_hr": schema.StringAttribute{
							MarkdownDescription: "Total price for all GPUs per hour",
							Computed:            true,
						},
						"total_memory_price_hr": schema.StringAttribute{
							MarkdownDescription: "Total price for all memory per hour",
							Computed:            true,
						},
						"total_price_hr": schema.StringAttribute{
							MarkdownDescription: "Total price for the VM per hour",
							Computed:            true,
						},
						"total_storage_price_hr": schema.StringAttribute{
							MarkdownDescription: "Total price for all storage per hour",
							Computed:            true,
						},
						"total_vcpu_price_hr": schema.StringAttribute{
							MarkdownDescription: "Total price for all vCPUs per hour",
							Computed:            true,
						},
						"vcpu_price_hr": schema.StringAttribute{
							MarkdownDescription: "Price per vCPU per hour",
							Computed:            true,
						},
						"count_vm_available": schema.Int64Attribute{
							MarkdownDescription: "Number of available VMs of this configuration",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *ComputeConfigsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ComputeConfigsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ComputeConfigsDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	params := search.NewSearchComputeParams()

	var v int32
	v = 2

	var m int32
	m = 4

	var gp int32
	gp = 0
	var st int32
	st = 10

	params.MemoryGib = &m
	params.Vcpu = &v
	params.Gpu = &gp
	params.StorageGib = &st

	res, err := d.client.Search.SearchCompute(params)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read compute_configs",
			err.Error(),
		)
		return
	}

	tflog.Debug(ctx, "Host config size: "+res.String()+fmt.Sprintf("%d", len(res.Payload.HostConfigs)))
	for _, cfg := range res.Payload.HostConfigs {
		tflog.Debug(ctx, "Host config: "+cfg.ID)
		computeConfigState := ComputeConfigsModel{
			Id:                  types.StringValue(cfg.ID),
			CpuModel:            types.StringValue(cfg.CPUModel),
			DataCenterId:        types.StringValue(cfg.DataCenterID),
			GpuMemoryGib:        types.Int64Value(int64(cfg.GpuMemoryGib)),
			GpuModel:            types.StringValue(cfg.GpuModel),
			GpuPriceHr:          types.StringValue(cfg.GpuPriceHr.Value),
			MemoryGibPriceHr:    types.StringValue(cfg.MemoryGibPriceHr.Value),
			StorageGibPriceHr:   types.StringValue(cfg.StorageGibPriceHr.Value),
			TotalGpuPriceHr:     types.StringValue(cfg.TotalGpuPriceHr.Value),
			TotalMemoryPriceHr:  types.StringValue(cfg.TotalMemoryPriceHr.Value),
			TotalPriceHr:        types.StringValue(cfg.TotalPriceHr.Value),
			TotalStoragePriceHr: types.StringValue(cfg.TotalStoragePriceHr.Value),
			TotalVcpuPriceHr:    types.StringValue(cfg.TotalVcpuPriceHr.Value),
			VcpuPriceHr:         types.StringValue(cfg.VcpuPriceHr.Value),
			CountVmAvailable:    types.Int64Value(int64(cfg.CountVMAvailable)),
		}

		state.ComputeConfigs = append(state.ComputeConfigs, computeConfigState)
	}

	state.SearchParams = &SearchParamsModel{
		CpuModel:     types.StringValue(""),
		DataCenterID: types.StringValue(""),
		GpuCount:     types.Int64Value(0),
		GpuModel:     types.StringValue(""),
		MemoryGiB:    types.Int64Value(0),
		OrderBy:      types.StringValue(""),
		PageNumber:   types.Int64Value(0),
		PageSize:     types.Int64Value(0),
		RegionID:     types.StringValue(""),
		StorageGiB:   types.Int64Value(0),
		VCPU:         types.Int64Value(0),
	}

	state.ID = types.StringValue("placeholder")

	tflog.Trace(ctx, "read a data source")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
