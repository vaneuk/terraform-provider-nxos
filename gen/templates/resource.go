//go:build ignore
// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-nxos"
	"github.com/netascode/terraform-provider-nxos/internal/provider/helpers"
)

type resource{{camelCase .Name}}Type struct{}

func (t resource{{camelCase .Name}}Type) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "{{.ResDescription}}",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.UseStateForUnknown(),
				},
			},
			{{- range  .Attributes}}
			"{{.TfName}}": {
				MarkdownDescription: helpers.NewDescription("{{.Description}}")
					{{- if len .EnumValues -}}
					.AddStringEnumDescription({{range .EnumValues}}"{{.}}", {{end}})
					{{- end -}}
					{{- if or (ne .MinInt 0) (ne .MaxInt 0) -}}
					.AddIntegerRangeDescription({{.MinInt}}, {{.MaxInt}})
					{{- end -}}
					{{- if len .DefaultValue -}}
					.AddDefaultValueDescription("{{.DefaultValue}}")
					{{- end -}}
					.String,
				Type:                types.{{.Type}}Type,
				{{- if eq .Id true}}
				Required:            true,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					tfsdk.RequiresReplace(),
				},
				{{- else}}
				Optional:            true,
				Computed:            true,
				{{- if len .EnumValues}}
				Validators: []tfsdk.AttributeValidator{
					helpers.StringEnumValidator({{range .EnumValues}}"{{.}}", {{end}}),
				},
				{{- else if or (ne .MinInt 0) (ne .MaxInt 0)}}
				Validators: []tfsdk.AttributeValidator{
					helpers.IntegerRangeValidator({{.MinInt}}, {{.MaxInt}}),
				},
				{{- end}}
				{{- if len .DefaultValue}}
				PlanModifiers: tfsdk.AttributePlanModifiers{
					{{- if eq .Type "Int64"}}
					helpers.IntegerDefaultModifier({{.DefaultValue}}),
					{{- else}}
					helpers.StringDefaultModifier("{{.DefaultValue}}"),
					{{- end}}
				},
				{{- end}}
				{{- end}}
			},
			{{- end}}
		},
	}, nil
}

func (t resource{{camelCase .Name}}Type) NewResource(ctx context.Context, in tfsdk.Provider) (tfsdk.Resource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return resource{{camelCase .Name}}{
		provider: provider,
	}, diags
}

type resource{{camelCase .Name}} struct {
	provider provider
}

func (r resource{{camelCase .Name}}) Create(ctx context.Context, req tfsdk.CreateResourceRequest, resp *tfsdk.CreateResourceResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.getDn()))

	// Post object
	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to post object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Read(ctx context.Context, req tfsdk.ReadResourceRequest, resp *tfsdk.ReadResourceResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Dn.Value))

	res, err := r.provider.client.GetDn(state.Dn.Value, nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Dn.Value))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Update(ctx context.Context, req tfsdk.UpdateResourceRequest, resp *tfsdk.UpdateResourceResponse) {
	var plan, state {{camelCase .Name}}

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.getDn()))

	body := plan.toBody()
	_, err := r.provider.client.Post(plan.getDn(), body.Str)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
		return
	}

	// Read object
	res, err := r.provider.client.GetDn(plan.getDn(), nxos.Query("rsp-prop-include", "config-only"))
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(plan)

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r resource{{camelCase .Name}}) Delete(ctx context.Context, req tfsdk.DeleteResourceRequest, resp *tfsdk.DeleteResourceResponse) {
	var state {{camelCase .Name}}

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Dn.Value))

	res, err := r.provider.client.DeleteDn(state.Dn.Value)
	if err != nil {
		errCode := res.Get("imdata.0.error.attributes.code").Str
		// Ignore errors of type "Cannot delete object"
		if errCode != "107" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to update object, got error: %s", err))
			return
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Dn.Value))

	resp.State.RemoveResource(ctx)
}

func (r resource{{camelCase .Name}}) ImportState(ctx context.Context, req tfsdk.ImportResourceStateRequest, resp *tfsdk.ImportResourceStateResponse) {
	tfsdk.ResourceImportStatePassthroughID(ctx, tftypes.NewAttributePath().WithAttributeName("id"), req, resp)
}
