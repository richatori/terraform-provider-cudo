package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/CudoVentures/terraform-provider-cudo/internal/client/networks"
	"github.com/CudoVentures/terraform-provider-cudo/internal/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SecurityGroupResource{}
var _ resource.ResourceWithImportState = &SecurityGroupResource{}

func NewSecurityGroupResource() resource.Resource {
	return &SecurityGroupResource{}
}

// SecurityGroupResource defines the resource implementation.
type SecurityGroupResource struct {
	client *CudoClientData
}

type RuleModel struct {
	Id          types.String `tfsdk:"id"`
	IcmpType    types.String `tfsdk:"icmp_type"`
	IpRangeCidr types.String `tfsdk:"ip_range"`
	Ports       types.String `tfsdk:"ports"`
	Protocol    types.String `tfsdk:"protocol"`
	RuleType    types.String `tfsdk:"rule_type"`
}

// SecurityGroupResourceModel describes the resource data model.
type SecurityGroupResourceModel struct {
	Id           types.String `tfsdk:"id"`
	DataCenterID types.String `tfsdk:"data_center_id"`
	Description  types.String `tfsdk:"description"`
	Rules        []RuleModel  `tfsdk:"rules"`
}

func (r *SecurityGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "cudo_security_group"
}

func (r *SecurityGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Security group resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Security Group ID",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"data_center_id": schema.StringAttribute{
				MarkdownDescription: "The unique identifier of the datacenter where the network is located.",
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{stringvalidator.RegexMatches(
					regexp.MustCompile("^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$"), "must be a valid resource id")},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description of the security group",
				Computed:            true,
				Optional:            true,
				Default:             stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.RegexMatches(securityGroupDescriptionRegex,
					"must be a valid security group description up to 255 characters, commas, periods, & spaces")},
			},
			"rules": schema.ListNestedAttribute{
				Description: "List of security group rules",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "The unique identifier of the rule",
							Computed:    true,
						},
						"icmp_type": schema.StringAttribute{
							Description: "Specific ICMP type of the rule. If a type has multiple codes, it includes all the codes within. This can only be used with ICMP",
							Optional:    true,
							Validators: []validator.String{stringvalidator.RegexMatches(
								icmpTypesRegex, "must be a valid icmp type i.e. 0,3,4,5,8,9,10,11,12,13,14,17,18")},
						},
						"ip_range": schema.StringAttribute{
							Description: "A single IP address or CIDR format range to apply rule to",
							Optional:    true,
						},
						"ports": schema.StringAttribute{
							Description: "A comma separated list of ports (80,443,8080) or a single port range (1024:2048)",
							Optional:    true,
							Validators:  []validator.String{portListValidator{}},
						},
						"protocol": schema.StringAttribute{
							Description: "Protocol for rule, use one of: all, tcp, udp, icmp, icmpv6, ipsec",
							Computed:    true,
							Optional:    true,
							Validators:  []validator.String{stringvalidator.OneOf("all", "tcp", "udp", "icmp", "icmpv6", "ipsec")},
						},
						"rule_type": schema.StringAttribute{
							Description: "Type for rule either 'inbound' or 'outbound'",
							Required:    true,
							Validators:  []validator.String{stringvalidator.OneOf("inbound", "outbound")},
						},
					},
				},
			},
		},
	}
}

func (r *SecurityGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func getNullableString(value string) basetypes.StringValue {
	var result = basetypes.StringValue{}
	if value != "" {
		result = types.StringValue(value)
	} else {
		result = types.StringNull()
	}

	return result
}

func getRuleModels(rules []*models.Rule) []RuleModel {
	var ruleModels []RuleModel
	for _, rule := range rules {
		protocol := ""
		switch *rule.Protocol {
		case models.ProtocolPROTOCOLALL:
			protocol = "all"
		case models.ProtocolPROTOCOLICMP:
			protocol = "icmp"
		case models.ProtocolPROTOCOLICMPv6:
			protocol = "icmpv6"
		case models.ProtocolPROTOCOLIPSEC:
			protocol = "ipsec"
		case models.ProtocolPROTOCOLTCP:
			protocol = "tcp"
		case models.ProtocolPROTOCOLUDP:
			protocol = "udp"
		}

		ruleType := ""

		switch *rule.RuleType {
		case models.RuleTypeRULETYPEINBOUND:
			ruleType = "inbound"
		case models.RuleTypeRULETYPEOUTBOUND:
			ruleType = "outbound"
		}

		ruleModel := RuleModel{
			Id:          types.StringValue(rule.ID),
			IcmpType:    getNullableString(rule.IcmpType),
			IpRangeCidr: getNullableString(rule.IPRangeCidr),
			Ports:       getNullableString(rule.Ports),
			Protocol:    getNullableString(protocol),
			RuleType:    getNullableString(ruleType),
		}

		ruleModels = append(ruleModels, ruleModel)
	}

	return ruleModels
}

func getRuleParams(stateRules []RuleModel) []*models.Rule {
	var rules []*models.Rule

	for _, r := range stateRules {

		var protocol *models.Protocol
		switch r.Protocol.ValueString() {
		case "tcp":
			protocol = models.NewProtocol(models.ProtocolPROTOCOLTCP)
		case "udp":
			protocol = models.NewProtocol(models.ProtocolPROTOCOLUDP)
		case "icmp":
			protocol = models.NewProtocol(models.ProtocolPROTOCOLICMP)
		case "icmpv6":
			protocol = models.NewProtocol(models.ProtocolPROTOCOLICMPv6)
		case "ipsec":
			protocol = models.NewProtocol(models.ProtocolPROTOCOLIPSEC)
		default:
			protocol = models.NewProtocol(models.ProtocolPROTOCOLALL)
		}

		var ruleType *models.RuleType
		switch r.RuleType.ValueString() {
		case "inbound":
			ruleType = models.NewRuleType(models.RuleTypeRULETYPEINBOUND)
		case "outbound":
			ruleType = models.NewRuleType(models.RuleTypeRULETYPEOUTBOUND)
		}

		rule := models.Rule{
			IcmpType:    r.IcmpType.ValueString(),
			ID:          r.Id.ValueString(),
			IPRangeCidr: r.IpRangeCidr.ValueString(),
			Ports:       r.Ports.ValueString(),
			Protocol:    protocol,
			RuleType:    ruleType,
		}
		rules = append(rules, &rule)
	}

	return rules
}

func (r *SecurityGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var state *SecurityGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewCreateSecurityGroupParamsWithContext(ctx)
	params.SecurityGroup = networks.CreateSecurityGroupBody{
		DataCenterID: state.DataCenterID.ValueStringPointer(),
		Description:  state.Description.ValueString(), //allows up to 255 characters, commas, periods, & spaces has regex
		ID:           state.Id.ValueStringPointer(),   // security group id
		Rules:        getRuleParams(state.Rules),
	}
	params.SecurityGroupProjectID = r.client.DefaultProjectID

	res, err := r.client.Client.Networks.CreateSecurityGroup(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create security group resource",
			err.Error(),
		)
		return
	}

	state.Rules = getRuleModels(res.Payload.SecurityGroup.Rules)

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SecurityGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state *SecurityGroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewGetSecurityGroupParamsWithContext(ctx)
	params.ID = state.Id.ValueString()
	params.ProjectID = r.client.DefaultProjectID

	res, err := r.client.Client.Networks.GetSecurityGroup(params)

	if err != nil {
		if apiErr, ok := err.(*networks.GetSecurityGroupDefault); ok && apiErr.IsCode(404) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Unable to read security group resource",
			err.Error(),
		)
		return
	}

	state.Id = types.StringValue(*res.Payload.SecurityGroup.ID)
	state.Description = types.StringValue(res.Payload.SecurityGroup.Description)
	state.DataCenterID = types.StringValue(*res.Payload.SecurityGroup.DataCenterID)
	state.Id = types.StringValue(*res.Payload.SecurityGroup.ID)
	state.Rules = getRuleModels(res.Payload.SecurityGroup.Rules)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SecurityGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state *SecurityGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := networks.NewUpdateSecurityGroupParamsWithContext(ctx)
	params.SecurityGroup = networks.UpdateSecurityGroupBody{
		DataCenterID: state.DataCenterID.ValueStringPointer(),
		Description:  state.Description.ValueString(), //allows up to 255 characters, commas, periods, & spaces has regex
		Rules:        getRuleParams(state.Rules),
	}
	params.SecurityGroupID = state.Id.ValueString()
	params.SecurityGroupProjectID = r.client.DefaultProjectID

	res, err := r.client.Client.Networks.UpdateSecurityGroup(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to update security group resource",
			err.Error(),
		)
		return
	}

	state.Rules = getRuleModels(res.Payload.SecurityGroup.Rules)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SecurityGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state *SecurityGroupResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	params := networks.NewDeleteSecurityGroupParamsWithContext(ctx)
	params.ProjectID = r.client.DefaultProjectID
	params.ID = state.Id.ValueString()

	_, err := r.client.Client.Networks.DeleteSecurityGroup(params)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to delete security group resource",
			err.Error(),
		)
		return
	}

	tflog.Trace(ctx, "deleted a security group")
}

func (r *SecurityGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

var securityGroupDescriptionRegex = regexp.MustCompile(`^[A-Za-z0-9,'.\s]{1,255}$`)
var icmpTypesRegex = regexp.MustCompile("^(0|3|4|5|8|9|10|11|12|13|14|17|18)$")

type portListValidator struct {
}

func (p portListValidator) Description(ctx context.Context) string {
	return "must be a comma separated list of ports 80,443,8080 or a single port range 1024:2048"
}

func (p portListValidator) MarkdownDescription(ctx context.Context) string {
	return "must be a comma separated list of ports 80,443,8080 or a single port range 1024:2048"
}

func (p portListValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	portString := req.ConfigValue.ValueString()

	_, ok := validatePorts(portString)

	if !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid port range",
			"String length must be a single range i.e. 1024:2048 or a comma separated list i.e. 80,443,8080",
		)
	}

}

func validatePorts(ports string) ([]uint16, bool) {
	var invalid bool
	portRanges := strings.Split(ports, ",")
	ranges := make([]uint16, 0, len(portRanges))
	for _, portRange := range portRanges {
		portRange = strings.Trim(portRange, " ")
		startStr, endStr, found := strings.Cut(portRange, ":")
		if !found {
			endStr = startStr
		}
		start, err := strconv.ParseUint(startStr, 10, 16)
		if err != nil {
			invalid = true
			continue
		}
		end, err := strconv.ParseUint(endStr, 10, 16)
		if err != nil {
			invalid = true
			continue
		}
		if end < start {
			invalid = true
			continue
		}
		ranges = append(ranges, uint16(start), uint16(end))
	}
	return ranges, !invalid
}
