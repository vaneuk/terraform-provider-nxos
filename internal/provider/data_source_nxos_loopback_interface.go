// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type dataSourceLoopbackInterfaceType struct{}

func (t dataSourceLoopbackInterfaceType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read a loopback interface.",

		Attributes: map[string]tfsdk.Attribute{
			"id": {
				MarkdownDescription: "The distinguished name of the object.",
				Type:                types.StringType,
				Computed:            true,
			},
			"interface_id": {
				MarkdownDescription: "Must match first field in the output of `show intf brief`. Example: `lo123`.",
				Type:                types.StringType,
				Required:            true,
			},
			"admin_state": {
				MarkdownDescription: "Administrative state.",
				Type:                types.StringType,
				Computed:            true,
			},
			"description": {
				MarkdownDescription: "Interface description.",
				Type:                types.StringType,
				Computed:            true,
			},
		},
	}, nil
}

func (t dataSourceLoopbackInterfaceType) NewDataSource(ctx context.Context, in tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	provider, diags := convertProviderType(in)

	return dataSourceLoopbackInterface{
		provider: provider,
	}, diags
}

type dataSourceLoopbackInterface struct {
	provider provider
}

func (d dataSourceLoopbackInterface) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var config, state LoopbackInterface

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.getDn()))

	res, err := d.provider.client.GetDn(config.getDn())

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	state.fromBody(res)
	state.fromPlan(config)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.getDn()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
