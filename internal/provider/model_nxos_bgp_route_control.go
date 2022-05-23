// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type BGPRouteControl struct {
	Device             types.String `tfsdk:"device"`
	Dn                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"vrf"`
	EnforceFirstAs     types.String `tfsdk:"enforce_first_as"`
	FibAccelerate      types.String `tfsdk:"fib_accelerate"`
	LogNeighborChanges types.String `tfsdk:"log_neighbor_changes"`
	SupprRt            types.String `tfsdk:"suppress_routes"`
}

func (data BGPRouteControl) getDn() string {
	return fmt.Sprintf("sys/bgp/inst/dom-[%s]/rtctrl", data.Name.Value)
}

func (data BGPRouteControl) getClassName() string {
	return "bgpRtCtrl"
}

func (data BGPRouteControl) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("enforceFirstAs", data.EnforceFirstAs.Value).
		Set("fibAccelerate", data.FibAccelerate.Value).
		Set("logNeighborChanges", data.LogNeighborChanges.Value).
		Set("supprRt", data.SupprRt.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *BGPRouteControl) fromBody(res gjson.Result) {
	data.EnforceFirstAs.Value = res.Get("*.attributes.enforceFirstAs").String()
	data.FibAccelerate.Value = res.Get("*.attributes.fibAccelerate").String()
	data.LogNeighborChanges.Value = res.Get("*.attributes.logNeighborChanges").String()
	data.SupprRt.Value = res.Get("*.attributes.supprRt").String()
}

func (data *BGPRouteControl) fromPlan(plan BGPRouteControl) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Name.Value = plan.Name.Value
}
