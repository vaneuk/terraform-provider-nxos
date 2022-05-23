// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type EVPNVNIRouteTargetDirection struct {
	Device types.String `tfsdk:"device"`
	Dn     types.String `tfsdk:"id"`
	Encap  types.String `tfsdk:"encap"`
	Type   types.String `tfsdk:"direction"`
}

func (data EVPNVNIRouteTargetDirection) getDn() string {
	return fmt.Sprintf("sys/evpn/bdevi-[%s]/rttp-[%s]", data.Encap.Value, data.Type.Value)
}

func (data EVPNVNIRouteTargetDirection) getClassName() string {
	return "rtctrlRttP"
}

func (data EVPNVNIRouteTargetDirection) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *EVPNVNIRouteTargetDirection) fromBody(res gjson.Result) {
	data.Type.Value = res.Get("*.attributes.type").String()
}

func (data *EVPNVNIRouteTargetDirection) fromPlan(plan EVPNVNIRouteTargetDirection) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Encap.Value = plan.Encap.Value
}
