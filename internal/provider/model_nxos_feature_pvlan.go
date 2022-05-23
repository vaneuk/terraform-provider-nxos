// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type FeaturePVLAN struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	AdminSt types.String `tfsdk:"admin_state"`
}

func (data FeaturePVLAN) getDn() string {
	return "sys/fm/pvlan"
}

func (data FeaturePVLAN) getClassName() string {
	return "fmPvlan"
}

func (data FeaturePVLAN) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("adminSt", data.AdminSt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *FeaturePVLAN) fromBody(res gjson.Result) {
	data.AdminSt.Value = res.Get("*.attributes.adminSt").String()
}

func (data *FeaturePVLAN) fromPlan(plan FeaturePVLAN) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
}
