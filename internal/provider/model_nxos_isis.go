// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type ISIS struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data ISIS) getDn() string {
	return "sys/isis"
}

func (data ISIS) getClassName() string {
	return "isisEntity"
}

func (data ISIS) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *ISIS) fromBody(res gjson.Result) {
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *ISIS) fromPlan(plan ISIS) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
}
