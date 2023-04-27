package provider

import (
	"context"
	"cudo.org/v1/terraform-provider-cudo/internal/client/ssh_keys"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SshKeysDataSource{}

func NewSshKeysDataSource() datasource.DataSource {
	return &SshKeysDataSource{}
}

// SshKeysDataSource defines the data source implementation.
type SshKeysDataSource struct {
	client *CudoClientData
}

type SshKeysModel struct {
	Id          types.String `tfsdk:"id"`
	PublicKey   types.String `tfsdk:"public_key"`
	Fingerprint types.String `tfsdk:"fingerprint"`
	Comment     types.String `tfsdk:"comment"`
	Type        types.String `tfsdk:"type"`
}

// SshKeysDataSourceModel describes the data source data model.
type SshKeysDataSourceModel struct {
	SshKeys []SshKeysModel `tfsdk:"ssh_keys"`
	//ID      types.String   `tfsdk:"id"`
}

func (d *SshKeysDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "cudo_ssh_keys"
}

func (d *SshKeysDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "SshKeys data source",
		Description:         "Fetches the list of SSH keys",
		Attributes: map[string]schema.Attribute{
			//"id": schema.StringAttribute{
			//	Description: "Placeholder identifier attribute.",
			//	Computed:    true,
			//},
			"ssh_keys": schema.ListNestedAttribute{
				Description: "List of SSH keys",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							MarkdownDescription: "SSH key identifier",
							Computed:            true,
						},
						"public_key": schema.StringAttribute{
							MarkdownDescription: "SSH key public key",
							Computed:            true,
						},
						"fingerprint": schema.StringAttribute{
							MarkdownDescription: "SSH key finger print",
							Computed:            true,
						},
						"comment": schema.StringAttribute{
							MarkdownDescription: "SSH key comment",
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: "SSH key type",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *SshKeysDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*CudoClientData)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *CudoClientData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *SshKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state SshKeysDataSourceModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)

	res, err := d.client.Client.SSHKeys.ListSSHKeys(ssh_keys.NewListSSHKeysParams())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to read sh keys",
			err.Error(),
		)
		return
	}

	for _, key := range res.Payload.SSHKeys {
		SshKeystate := SshKeysModel{
			//Id:          types.StringValue(key.ID),
			PublicKey:   types.StringValue(key.PublicKey),
			Fingerprint: types.StringValue(key.Fingerprint),
			Comment:     types.StringValue(key.Comment),
			Type:        types.StringValue(key.Type),
		}

		state.SshKeys = append(state.SshKeys, SshKeystate)
	}

	//state.ID = types.StringValue("placeholder")

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}
