package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// 包过滤
var _ resource.Resource = &PfPolicyResource{}
var _ resource.ResourceWithImportState = &PfPolicyResource{}

func NewPfPolicyResource() resource.Resource {
	return &PfPolicyResource{}
}

type PfPolicyResource struct {
	client *Client
}

type PfPolicyResourceModel struct {
	AddPfPolicyParameter AddPfPolicyParameter `tfsdk:"securitypolicylist"`
}

type AddPfPolicyRequest struct {
	AddPfPolicyRequestModel AddPfPolicyRequestModel `json:"securitypolicylist"`
}

type UpdatePfPolicyRequest struct {
	UpdatePfPolicyRequestModel UpdatePfPolicyRequestModel `json:"securitypolicylist"`
}

// 调用接口参数
type AddPfPolicyRequestModel struct {
	Name                    string `json:"name"`
	Enabled                 string `json:"enabled"`
	Action                  string `json:"action"`
	DelAllEnable            string `json:"delAllEnable"`
	IpVersion               string `json:"ipVersion"`
	VsysName                string `json:"vsysName"`
	GroupName               string `json:"groupName"`
	TargetName              string `json:"targetName"`
	Position                string `json:"position"`
	SourceSecurityZone      string `json:"sourceSecurityZone"`
	DestinationSecurityZone string `json:"destinationSecurityZone"`
	SourceIpObjects         string `json:"sourceIpObjects"`
	SourceIpGroups          string `json:"sourceIpGroups"`
	SourceDomains           string `json:"sourceDomains"`
	SourceMacObjects        string `json:"sourceMacObjects"`
	SourceMacGroups         string `json:"sourceMacGroups"`
	DestinationIpObjects    string `json:"destinationIpObjects"`
	DestinationIpGroups     string `json:"destinationIpGroups"`
	DestinationDomains      string `json:"destinationDomains"`
	DestinationMacObjects   string `json:"destinationMacObjects"`
	DestinationMacGroups    string `json:"destinationMacGroups"`
	ServicePreObjects       string `json:"servicePreObjects"`
	ServiceUsrObjects       string `json:"serviceUsrObjects"`
	ServiceGroups           string `json:"serviceGroups"`
	UserObjects             string `json:"userObjects"`
	UserGroups              string `json:"userGroups"`
	Description             string `json:"describe"`
	EffectName              string `json:"effectName"`
	Matchlog                string `json:"matchlog"`
	Sessionlog              string `json:"sessionlog"`
	Longsession             string `json:"longsession"`
	Agingtime               string `json:"agingtime"`
	Fragdrop                string `json:"fragdrop"`
	Dscp                    string `json:"dscp"`
	Cos                     string `json:"cos"`
	RltGroup                string `json:"rltGroup"`
	RltUser                 string `json:"rltUser"`
	Acctl                   string `json:"acctl"`
	UrlClass                string `json:"urlClass"`
	UrlSenior               string `json:"urlSenior"`
	Cam                     string `json:"cam"`
	Ips                     string `json:"ips"`
	Av                      string `json:"av"`
}

type UpdatePfPolicyRequestModel struct {
	IpVersion               string `json:"ipVersion"`
	VsysName                string `json:"vsysName"`
	GroupName               string `json:"groupName"`
	TargetName              string `json:"targetName"`
	Position                string `json:"position"`
	Name                    string `json:"name"`
	OldName                 string `json:"oldName"`
	SourceSecurityZone      string `json:"sourceSecurityZone"`
	DestinationSecurityZone string `json:"destinationSecurityZone"`
	Enabled                 string `json:"enabled"`
	Action                  string `json:"action"`
	SourceIpObjects         string `json:"sourceIpObjects"`
	SourceIpGroups          string `json:"sourceIpGroups"`
	SourceDomains           string `json:"sourceDomains"`
	SourceMacObjects        string `json:"sourceMacObjects"`
	SourceMacGroups         string `json:"sourceMacGroups"`
	DestinationIpObjects    string `json:"destinationIpObjects"`
	DestinationIpGroups     string `json:"destinationIpGroups"`
	DestinationDomains      string `json:"destinationDomains"`
	DestinationMacObjects   string `json:"destinationMacObjects"`
	DestinationMacGroups    string `json:"destinationMacGroups"`
	ServicePreObjects       string `json:"servicePreObjects"`
	ServiceUsrObjects       string `json:"serviceUsrObjects"`
	ServiceGroups           string `json:"serviceGroups"`
	UserObjects             string `json:"userObjects"`
	UserGroups              string `json:"userGroups"`
	Description             string `json:"description"`
	EffectName              string `json:"effectName"`
	Matchlog                string `json:"matchlog"`
	Sessionlog              string `json:"sessionlog"`
	Longsession             string `json:"longsession"`
	Agingtime               string `json:"agingtime"`
	Fragdrop                string `json:"fragdrop"`
	Dscp                    string `json:"dscp"`
	Cos                     string `json:"cos"`
	RltGroup                string `json:"rltGroup"`
	RltUser                 string `json:"rltUser"`
	Acctl                   string `json:"acctl"`
	UrlClass                string `json:"urlClass"`
	UrlSenior               string `json:"urlSenior"`
	Cam                     string `json:"cam"`
	Ips                     string `json:"ips"`
	Av                      string `json:"av"`
}

// 接收外部参数
type AddPfPolicyParameter struct {
	Name                    types.String `tfsdk:"name"`
	Enabled                 types.String `tfsdk:"enabled"`
	Action                  types.String `tfsdk:"action"`
	DelAllEnable            types.String `tfsdk:"delallenable"`
	IpVersion               types.String `tfsdk:"ipversion"`
	VsysName                types.String `tfsdk:"vsysname"`
	GroupName               types.String `tfsdk:"groupname"`
	TargetName              types.String `tfsdk:"targetname"`
	Position                types.String `tfsdk:"position"`
	SourceSecurityZone      types.String `tfsdk:"sourcesecurityzone"`
	DestinationSecurityZone types.String `tfsdk:"destinationsecurityzone"`
	SourceIpObjects         types.String `tfsdk:"sourceipobjects"`
	SourceIpGroups          types.String `tfsdk:"sourceipgroups"`
	SourceDomains           types.String `tfsdk:"sourcedomains"`
	SourceMacObjects        types.String `tfsdk:"sourcemacobjects"`
	SourceMacGroups         types.String `tfsdk:"sourcemacgroups"`
	DestinationIpObjects    types.String `tfsdk:"destinationipobjects"`
	DestinationIpGroups     types.String `tfsdk:"destinationipgroups"`
	DestinationDomains      types.String `tfsdk:"destinationdomains"`
	DestinationMacObjects   types.String `tfsdk:"destinationmacobjects"`
	DestinationMacGroups    types.String `tfsdk:"destinationmacgroups"`
	ServicePreObjects       types.String `tfsdk:"servicepreobjects"`
	ServiceUsrObjects       types.String `tfsdk:"serviceusrobjects"`
	ServiceGroups           types.String `tfsdk:"servicegroups"`
	UserObjects             types.String `tfsdk:"userobjects"`
	UserGroups              types.String `tfsdk:"usergroups"`
	Description             types.String `tfsdk:"describe"`
	EffectName              types.String `tfsdk:"effectname"`
	Matchlog                types.String `tfsdk:"matchlog"`
	Sessionlog              types.String `tfsdk:"sessionlog"`
	Longsession             types.String `tfsdk:"longsession"`
	Agingtime               types.String `tfsdk:"agingtime"`
	Fragdrop                types.String `tfsdk:"fragdrop"`
	Dscp                    types.String `tfsdk:"dscp"`
	Cos                     types.String `tfsdk:"cos"`
	RltGroup                types.String `tfsdk:"rltgroup"`
	RltUser                 types.String `tfsdk:"rltuser"`
	Acctl                   types.String `tfsdk:"acctl"`
	UrlClass                types.String `tfsdk:"urlclass"`
	UrlSenior               types.String `tfsdk:"urlsenior"`
	Cam                     types.String `tfsdk:"cam"`
	Ips                     types.String `tfsdk:"ips"`
	Av                      types.String `tfsdk:"av"`
}

// 查询结果结构体
type QueryPfPolicyResponseListModel struct {
	Securitypolicylist []QueryPfPolicyResponseModel `json:"securitypolicylist"`
}
type QueryPfPolicyResponseModel struct {
	IpVersion               string `json:"ipVersion"`
	VsysName                string `json:"vsysName"`
	Offset                  string `json:"offset"`
	Count                   string `json:"count"`
	GroupName               string `json:"groupName"`
	TargetName              string `json:"targetName"`
	Position                string `json:"position"`
	Name                    string `json:"name"`
	OldName                 string `json:"oldName"`
	SourceSecurityZone      string `json:"sourceSecurityZone"`
	DestinationSecurityZone string `json:"destinationSecurityZone"`
	Sourceaddress           string `json:"sourceaddress"`
	Destinationaddress      string `json:"destinationaddress"`
	Service                 string `json:"service"`
	Enabled                 string `json:"enabled"`
	Action                  string `json:"action"`
	MatchTime               string `json:"matchTime"`
	DelallEnable            string `json:"delallEnable"`
	Id                      string `json:"id"`
	SourceIpObjects         string `json:"sourceIpObjects"`
	SourceIpGroups          string `json:"sourceIpGroups"`
	SourceDomains           string `json:"sourceDomains"`
	SourceMacObjects        string `json:"sourceMacObjects"`
	SourceMacGroups         string `json:"sourceMacGroups"`
	DestinationIpObjects    string `json:"destinationIpObjects"`
	DestinationIpGroups     string `json:"destinationIpGroups"`
	DestinationDomains      string `json:"destinationDomains"`
	DestinationMacObjects   string `json:"destinationMacObjects"`
	DestinationMacGroups    string `json:"destinationMacGroups"`
	ServicePreObjects       string `json:"servicePreObjects"`
	ServiceUsrObjects       string `json:"serviceUsrObjects"`
	ServiceGroups           string `json:"serviceGroups"`
	UserObjects             string `json:"userObjects"`
	UserGroups              string `json:"userGroups"`
	Description             string `json:"description"`
	EffectName              string `json:"effectName"`
	Matchlog                string `json:"matchlog"`
	Sessionlog              string `json:"sessionlog"`
	Longsession             string `json:"longsession"`
	Agingtime               string `json:"agingtime"`
	Fragdrop                string `json:"fragdrop"`
	Dscp                    string `json:"dscp"`
	Cos                     string `json:"cos"`
	RltGroup                string `json:"rltGroup"`
	RltUser                 string `json:"rltUser"`
	Acctl                   string `json:"acctl"`
	UrlClass                string `json:"urlClass"`
	UrlSenior               string `json:"urlSenior"`
	Cam                     string `json:"cam"`
	Ips                     string `json:"ips"`
	Av                      string `json:"av"`
	Createtime              string `json:"createtime"`
	Modifytime              string `json:"modifytime"`
	Matchnum                string `json:"matchnum"`
	FlowStatisticsBytes     string `json:"flowStatisticsBytes"`
	FlowStatisticsPackets   string `json:"flowStatisticsPackets"`
}

func (r *PfPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "dpfirewall_PfPolicy"
}

func (r *PfPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"securitypolicylist": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"enabled": schema.StringAttribute{
						Required: true,
					},
					"action": schema.StringAttribute{
						Required: true,
					},
					"delallenable": schema.StringAttribute{
						Optional: true,
					},
					"ipversion": schema.StringAttribute{
						Optional: true,
					},
					"vsysname": schema.StringAttribute{
						Optional: true,
					},
					"groupname": schema.StringAttribute{
						Optional: true,
					},
					"targetname": schema.StringAttribute{
						Optional: true,
					},
					"position": schema.StringAttribute{
						Optional: true,
					},
					"effectname": schema.StringAttribute{
						Optional: true,
					},
					"matchlog": schema.StringAttribute{
						Optional: true,
					},
					"sessionlog": schema.StringAttribute{
						Optional: true,
					},
					"longsession": schema.StringAttribute{
						Optional: true,
					},
					"agingtime": schema.StringAttribute{
						Optional: true,
					},
					"fragdrop": schema.StringAttribute{
						Optional: true,
					},
					"sourcesecurityzone": schema.StringAttribute{
						Optional: true,
					},
					"destinationsecurityzone": schema.StringAttribute{
						Optional: true,
					},
					"sourceipobjects": schema.StringAttribute{
						Optional: true,
					},
					"sourceipgroups": schema.StringAttribute{
						Optional: true,
					},
					"sourcedomains": schema.StringAttribute{
						Optional: true,
					},
					"sourcemacobjects": schema.StringAttribute{
						Optional: true,
					},
					"sourcemacgroups": schema.StringAttribute{
						Optional: true,
					},
					"destinationipobjects": schema.StringAttribute{
						Optional: true,
					},
					"destinationipgroups": schema.StringAttribute{
						Optional: true,
					},
					"destinationdomains": schema.StringAttribute{
						Optional: true,
					},
					"destinationmacobjects": schema.StringAttribute{
						Optional: true,
					},
					"destinationmacgroups": schema.StringAttribute{
						Optional: true,
					},
					"servicepreobjects": schema.StringAttribute{
						Optional: true,
					},
					"serviceusrobjects": schema.StringAttribute{
						Optional: true,
					},
					"servicegroups": schema.StringAttribute{
						Optional: true,
					},
					"userobjects": schema.StringAttribute{
						Optional: true,
					},
					"usergroups": schema.StringAttribute{
						Optional: true,
					},
					"describe": schema.StringAttribute{
						Optional: true,
					},
					"dscp": schema.StringAttribute{
						Optional: true,
					},
					"cos": schema.StringAttribute{
						Optional: true,
					},
					"rltgroup": schema.StringAttribute{
						Optional: true,
					},
					"rltuser": schema.StringAttribute{
						Optional: true,
					},
					"acctl": schema.StringAttribute{
						Optional: true,
					},
					"urlclass": schema.StringAttribute{
						Optional: true,
					},
					"urlsenior": schema.StringAttribute{
						Optional: true,
					},
					"cam": schema.StringAttribute{
						Optional: true,
					},
					"ips": schema.StringAttribute{
						Optional: true,
					},
					"av": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r *PfPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*Client)

	if req.ProviderData == nil {
		return
	}
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *PfPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PfPolicyResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Trace(ctx, "created a resource **************")
	sendToweb_Request(ctx, "POST", r.client, data.AddPfPolicyParameter)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PfPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PfPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " read Start ***************")
	//sendToweb_Request(ctx, "GET", r.client, data.AddPfPolicyParameter)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PfPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PfPolicyResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " Update Start ************")
	//sendToweb_Request(ctx, "PUT", r.client, data.AddPfPolicyParameter)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PfPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PfPolicyResourceModel
	tflog.Info(ctx, " Delete Start *************")

	//sendToweb_Request(ctx, "DELETE", r.client, data.AddPfPolicyParameter)

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PfPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func sendToweb_Request(ctx context.Context, reqmethod string, c *Client, Rsinfo AddPfPolicyParameter) {

	if reqmethod == "POST" {
		// 先查询是否存在，再执行新增操作
		tflog.Info(ctx, "包过滤--开始执行--查询操作")
		responseBody := sendRequest(ctx, "GET", c, nil, "/func/web_main/api/pf_policy/pf_policy/pf_policy/securitypolicylist", "包过滤")
		var queryResList QueryPfPolicyResponseListModel
		err := json.Unmarshal([]byte(responseBody), &queryResList)
		if err != nil {
			panic("转换查询结果json出现异常")
		}
		for _, queryRes := range queryResList.Securitypolicylist {
			if queryRes.Name == Rsinfo.Name.ValueString() {
				tflog.Info(ctx, "包过滤--存在重复数据，执行--修改操作")
				var sendUpdateData UpdatePfPolicyRequestModel
				sendUpdateData = UpdatePfPolicyRequestModel{
					IpVersion:               Rsinfo.IpVersion.ValueString(),
					VsysName:                Rsinfo.VsysName.ValueString(),
					GroupName:               Rsinfo.GroupName.ValueString(),
					TargetName:              Rsinfo.TargetName.ValueString(),
					Position:                Rsinfo.Position.ValueString(),
					Name:                    Rsinfo.Name.ValueString(),
					OldName:                 Rsinfo.Name.ValueString(),
					SourceSecurityZone:      Rsinfo.SourceSecurityZone.ValueString(),
					DestinationSecurityZone: Rsinfo.DestinationSecurityZone.ValueString(),
					Enabled:                 Rsinfo.Enabled.ValueString(),
					Action:                  Rsinfo.Action.ValueString(),
					SourceIpObjects:         Rsinfo.SourceIpObjects.ValueString(),
					SourceIpGroups:          Rsinfo.SourceIpGroups.ValueString(),
					SourceDomains:           Rsinfo.SourceDomains.ValueString(),
					SourceMacObjects:        Rsinfo.SourceMacObjects.ValueString(),
					SourceMacGroups:         Rsinfo.SourceMacGroups.ValueString(),
					DestinationIpObjects:    Rsinfo.DestinationIpObjects.ValueString(),
					DestinationIpGroups:     Rsinfo.DestinationIpGroups.ValueString(),
					DestinationDomains:      Rsinfo.DestinationDomains.ValueString(),
					DestinationMacObjects:   Rsinfo.DestinationMacObjects.ValueString(),
					DestinationMacGroups:    Rsinfo.DestinationMacGroups.ValueString(),
					ServicePreObjects:       Rsinfo.ServicePreObjects.ValueString(),
					ServiceUsrObjects:       Rsinfo.ServiceUsrObjects.ValueString(),
					ServiceGroups:           Rsinfo.ServiceGroups.ValueString(),
					UserObjects:             Rsinfo.UserObjects.ValueString(),
					UserGroups:              Rsinfo.UserGroups.ValueString(),
					Description:             Rsinfo.Description.ValueString(),
					EffectName:              Rsinfo.EffectName.ValueString(),
					Matchlog:                Rsinfo.Matchlog.ValueString(),
					Sessionlog:              Rsinfo.Sessionlog.ValueString(),
					Longsession:             Rsinfo.Longsession.ValueString(),
					Agingtime:               Rsinfo.Agingtime.ValueString(),
					Fragdrop:                Rsinfo.Fragdrop.ValueString(),
					Dscp:                    Rsinfo.Dscp.ValueString(),
					Cos:                     Rsinfo.Cos.ValueString(),
					RltGroup:                Rsinfo.RltGroup.ValueString(),
					RltUser:                 Rsinfo.RltUser.ValueString(),
					Acctl:                   Rsinfo.Acctl.ValueString(),
					UrlClass:                Rsinfo.UrlClass.ValueString(),
					UrlSenior:               Rsinfo.UrlSenior.ValueString(),
					Cam:                     Rsinfo.Cam.ValueString(),
					Ips:                     Rsinfo.Ips.ValueString(),
					Av:                      Rsinfo.Av.ValueString(),
				}

				requstUpdateData := UpdatePfPolicyRequest{
					UpdatePfPolicyRequestModel: sendUpdateData,
				}
				body, _ := json.Marshal(requstUpdateData)

				sendRequest(ctx, "PUT", c, body, "/func/web_main/api/pf_policy/pf_policy/pf_policy/securitypolicylist", "包过滤")
				return
			}
		}
		// 新增操作
		var sendData AddPfPolicyRequestModel
		sendData = AddPfPolicyRequestModel{
			Name:                    Rsinfo.Name.ValueString(),
			Enabled:                 Rsinfo.Enabled.ValueString(),
			Action:                  Rsinfo.Action.ValueString(),
			IpVersion:               Rsinfo.IpVersion.ValueString(),
			VsysName:                Rsinfo.VsysName.ValueString(),
			GroupName:               Rsinfo.GroupName.ValueString(),
			TargetName:              Rsinfo.TargetName.ValueString(),
			Position:                Rsinfo.Position.ValueString(),
			EffectName:              Rsinfo.EffectName.ValueString(),
			Matchlog:                Rsinfo.Matchlog.ValueString(),
			Sessionlog:              Rsinfo.Sessionlog.ValueString(),
			Longsession:             Rsinfo.Longsession.ValueString(),
			Agingtime:               Rsinfo.Agingtime.ValueString(),
			Fragdrop:                Rsinfo.Fragdrop.ValueString(),
			SourceSecurityZone:      Rsinfo.SourceSecurityZone.ValueString(),
			DestinationSecurityZone: Rsinfo.DestinationSecurityZone.ValueString(),
			SourceIpObjects:         Rsinfo.SourceIpObjects.ValueString(),
			SourceIpGroups:          Rsinfo.SourceIpGroups.ValueString(),
			SourceDomains:           Rsinfo.SourceDomains.ValueString(),
			SourceMacObjects:        Rsinfo.SourceMacObjects.ValueString(),
			SourceMacGroups:         Rsinfo.SourceMacGroups.ValueString(),
			DestinationIpObjects:    Rsinfo.DestinationIpObjects.ValueString(),
			DestinationIpGroups:     Rsinfo.DestinationIpGroups.ValueString(),
			DestinationDomains:      Rsinfo.DestinationDomains.ValueString(),
			DestinationMacObjects:   Rsinfo.DestinationMacObjects.ValueString(),
			DestinationMacGroups:    Rsinfo.DestinationMacGroups.ValueString(),
			ServicePreObjects:       Rsinfo.ServicePreObjects.ValueString(),
			ServiceUsrObjects:       Rsinfo.ServiceUsrObjects.ValueString(),
			ServiceGroups:           Rsinfo.ServiceGroups.ValueString(),
			UserObjects:             Rsinfo.UserObjects.ValueString(),
			UserGroups:              Rsinfo.UserGroups.ValueString(),
			Description:             Rsinfo.Description.ValueString(),
			Dscp:                    Rsinfo.Dscp.ValueString(),
			Cos:                     Rsinfo.Cos.ValueString(),
			RltGroup:                Rsinfo.RltGroup.ValueString(),
			RltUser:                 Rsinfo.RltUser.ValueString(),
			Acctl:                   Rsinfo.Acctl.ValueString(),
			UrlClass:                Rsinfo.UrlClass.ValueString(),
			UrlSenior:               Rsinfo.UrlSenior.ValueString(),
			Cam:                     Rsinfo.Cam.ValueString(),
			Ips:                     Rsinfo.Ips.ValueString(),
			Av:                      Rsinfo.Av.ValueString(),
		}

		requstData := AddPfPolicyRequest{
			AddPfPolicyRequestModel: sendData,
		}
		body, _ := json.Marshal(requstData)
		sendRequest(ctx, reqmethod, c, body, "/func/web_main/api/pf_policy/pf_policy/pf_policy/securitypolicylist", "包过滤")
		return
	} else if reqmethod == "GET" {

	} else if reqmethod == "PUT" {

	} else if reqmethod == "DELETE" {

	}
}
