// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type QueuingQOSPolicySystemOut struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
}

func (data QueuingQOSPolicySystemOut) getDn() string {
	return "sys/ipqos/queuing/policy/out/sys"
}

func (data QueuingQOSPolicySystemOut) getClassName() string {
	return "ipqosSystem"
}

func (data QueuingQOSPolicySystemOut) toBody() nxos.Body {
	return nxos.Body{Str: `{"` + data.getClassName() + `":{"attributes":{}}}`}
}

func (data *QueuingQOSPolicySystemOut) fromBody(res gjson.Result) {
}

func (data *QueuingQOSPolicySystemOut) fromPlan(plan QueuingQOSPolicySystemOut) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
}
