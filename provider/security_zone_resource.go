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

// 安全域
var _ resource.Resource = &SecurityZoneResource{}
var _ resource.ResourceWithImportState = &SecurityZoneResource{}

func NewSecurityZoneResource() resource.Resource {
	return &SecurityZoneResource{}
}

type SecurityZoneResource struct {
	client *Client
}

type SecurityZoneResourceModel struct {
	AddSecurityZoneParameter AddSecurityZoneParameter `tfsdk:"securityzonelist"`
}

type AddSecurityZoneRequest struct {
	AddSecurityZoneRequestModel AddSecurityZoneRequestModel `json:"securityzonelist"`
}

type UpdateSecurityZoneRequest struct {
	UpdateSecurityZoneRequestModel UpdateSecurityZoneRequestModel `json:"securityzonelist"`
}

// 调用接口参数
type AddSecurityZoneRequestModel struct {
	Vsysname    string `json:"vsysName"`
	Name        string `json:"name"`
	Priority    string `json:"priority"`
	Interfaces  string `json:"interfaces"`
	Desc        string `json:"desc"`
	Inneraction string `json:"innerAction"`
}

type UpdateSecurityZoneRequestModel struct {
	Vsysname    string `json:"vsysName"`
	Name        string `json:"name"`
	OldName     string `json:"oldName"`
	Priority    string `json:"priority"`
	Interfaces  string `json:"interfaces"`
	Desc        string `json:"desc"`
	Inneraction string `json:"innerAction"`
}

// 接收外部参数
type AddSecurityZoneParameter struct {
	Vsysname    types.String `tfsdk:"vsysname"`
	Name        types.String `tfsdk:"name"`
	Priority    types.String `tfsdk:"priority"`
	Interfaces  types.String `tfsdk:"interfaces"`
	Desc        types.String `tfsdk:"desc"`
	Inneraction types.String `tfsdk:"inneraction"`
}

// 查询结果结构体
type QuerySecurityZoneResponseListModel struct {
	Securityzonelist []QuerySecurityZoneResponseModel `json:"securityzonelist"`
}
type QuerySecurityZoneResponseModel struct {
	Vsysname     string `json:"vsysName"`
	Name         string `json:"name"`
	Priority     string `json:"priority"`
	Interfaces   string `json:"interfaces"`
	Desc         string `json:"desc"`
	Inneraction  string `json:"innerAction"`
	OldName      string `json:"oldName"`
	ReferNum     string `json:"referNum"`
	DelAllEnable string `json:"delAllEnable"`
	Offset       string `json:"offset"`
	Count        string `json:"count"`
}

func (r *SecurityZoneResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "dpfirewall_SecurityZone"
}

func (r *SecurityZoneResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"securityzonelist": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"vsysname": schema.StringAttribute{
						Optional: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"priority": schema.StringAttribute{
						Required: true,
					},
					"interfaces": schema.StringAttribute{
						Optional: true,
					},
					"desc": schema.StringAttribute{
						Optional: true,
					},
					"inneraction": schema.StringAttribute{
						Optional: true,
					},
				},
			},
		},
	}
}

func (r *SecurityZoneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *SecurityZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *SecurityZoneResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Trace(ctx, "created a resource ************** ")
	sendToweb_SecurityZoneRequest(ctx, "POST", r.client, data.AddSecurityZoneParameter)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *SecurityZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " read Start ***************")
	//sendToweb_SecurityZoneRequest(ctx, "GET", r.client, data.AddSecurityZoneParameter)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *SecurityZoneResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, " Update Start ************")
	//sendToweb_SecurityZoneRequest(ctx, "PUT", r.client, data.AddSecurityZoneParameter)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SecurityZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *SecurityZoneResourceModel
	tflog.Info(ctx, " Delete Start *************")

	//sendToweb_SecurityZoneRequest(ctx, "DELETE", r.client, data.AddSecurityZoneParameter)

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *SecurityZoneResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func sendToweb_SecurityZoneRequest(ctx context.Context, reqmethod string, c *Client, Rsinfo AddSecurityZoneParameter) {

	if reqmethod == "POST" {
		// 先查询是否存在，再执行新增操作
		tflog.Info(ctx, "安全域--开始执行--查询操作")
		responseBody := sendRequest(ctx, "GET", c, nil, "/func/web_main/api/security_zone/security_zone/securityzonelist?vsysName=PublicSystem&offset=0&count=1000", "安全域")
		var queryResList QuerySecurityZoneResponseListModel
		err := json.Unmarshal([]byte(responseBody), &queryResList)
		if err != nil {
			panic("转换查询结果json出现异常")
		}
		for _, queryRes := range queryResList.Securityzonelist {
			if queryRes.Name == Rsinfo.Name.ValueString() {
				tflog.Info(ctx, "安全域--存在重复数据，执行--修改操作")
				var sendUpdateData UpdateSecurityZoneRequestModel
				sendUpdateData = UpdateSecurityZoneRequestModel{
					Vsysname:    Rsinfo.Vsysname.ValueString(),
					Name:        Rsinfo.Name.ValueString(),
					OldName:     Rsinfo.Name.ValueString(),
					Priority:    Rsinfo.Priority.ValueString(),
					Interfaces:  Rsinfo.Interfaces.ValueString(),
					Desc:        Rsinfo.Desc.ValueString(),
					Inneraction: Rsinfo.Inneraction.ValueString(),
				}

				requstUpdateData := UpdateSecurityZoneRequest{
					UpdateSecurityZoneRequestModel: sendUpdateData,
				}
				body, _ := json.Marshal(requstUpdateData)

				sendRequest(ctx, "PUT", c, body, "/func/web_main/api/security_zone/security_zone/securityzonelist", "安全域")
				return
			}
		}

		// 新增操作
		var sendData AddSecurityZoneRequestModel
		sendData = AddSecurityZoneRequestModel{
			Vsysname:    Rsinfo.Vsysname.ValueString(),
			Name:        Rsinfo.Name.ValueString(),
			Priority:    Rsinfo.Priority.ValueString(),
			Interfaces:  Rsinfo.Interfaces.ValueString(),
			Desc:        Rsinfo.Desc.ValueString(),
			Inneraction: Rsinfo.Inneraction.ValueString(),
		}

		requstData := AddSecurityZoneRequest{
			AddSecurityZoneRequestModel: sendData,
		}
		body, _ := json.Marshal(requstData)

		sendRequest(ctx, reqmethod, c, body, "/func/web_main/api/security_zone/security_zone/securityzonelist", "安全域")
		return
	} else if reqmethod == "GET" {

	} else if reqmethod == "PUT" {

	} else if reqmethod == "DELETE" {

	}
}
