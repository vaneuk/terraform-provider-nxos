// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type VRFRouteTargetAddressFamily struct {
	Device  types.String `tfsdk:"device"`
	Dn      types.String `tfsdk:"id"`
	Vrf     types.String `tfsdk:"vrf"`
	Af_type types.String `tfsdk:"address_family"`
	Type    types.String `tfsdk:"route_target_address_family"`
}

func (data VRFRouteTargetAddressFamily) getDn() string {
	return fmt.Sprintf("sys/inst-[%s]/dom-[%[1]s]/af-[%s]/ctrl-[%s]", data.Vrf.Value, data.Af_type.Value, data.Type.Value)
}

func (data VRFRouteTargetAddressFamily) getClassName() string {
	return "rtctrlAfCtrl"
}

func (data VRFRouteTargetAddressFamily) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("type", data.Type.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *VRFRouteTargetAddressFamily) fromBody(res gjson.Result) {
	data.Type.Value = res.Get("*.attributes.type").String()
}

func (data *VRFRouteTargetAddressFamily) fromPlan(plan VRFRouteTargetAddressFamily) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Vrf.Value = plan.Vrf.Value
	data.Af_type.Value = plan.Af_type.Value
}
