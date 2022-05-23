// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
	"github.com/tidwall/gjson"
)

type PIMSSMPolicy struct {
	Device   types.String `tfsdk:"device"`
	Dn       types.String `tfsdk:"id"`
	Vrf_name types.String `tfsdk:"vrf_name"`
	Name     types.String `tfsdk:"name"`
}

func (data PIMSSMPolicy) getDn() string {
	return fmt.Sprintf("sys/pim/inst/dom-[%s]/ssm", data.Vrf_name.Value)
}

func (data PIMSSMPolicy) getClassName() string {
	return "pimSSMPatP"
}

func (data PIMSSMPolicy) toBody() nxos.Body {
	attrs := nxos.Body{}.
		Set("name", data.Name.Value)
	return nxos.Body{}.SetRaw(data.getClassName()+".attributes", attrs.Str)
}

func (data *PIMSSMPolicy) fromBody(res gjson.Result) {
	data.Name.Value = res.Get("*.attributes.name").String()
}

func (data *PIMSSMPolicy) fromPlan(plan PIMSSMPolicy) {
	data.Device = plan.Device
	data.Dn.Value = plan.Dn.Value
	data.Vrf_name.Value = plan.Vrf_name.Value
}
