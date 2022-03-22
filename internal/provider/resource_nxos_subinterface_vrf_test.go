// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosSubinterfaceVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPhysicalInterfaceConfig_all(),
			},
			{
				Config: testAccNxosSubinterfaceConfig_all(),
			},
			{
				Config:testAccNxosPhysicalInterfaceConfig_all()+testAccNxosSubinterfaceConfig_all()+testAccNxosSubinterfaceVRFConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "interface_id", "eth1/10.124"),
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "vrf_dn", "sys/inst-VRF123"),
				),
			},
			{
				Config:testAccNxosPhysicalInterfaceConfig_all()+testAccNxosSubinterfaceConfig_all()+testAccNxosSubinterfaceVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "interface_id", "eth1/10.124"),
					resource.TestCheckResourceAttr("nxos_subinterface_vrf.test", "vrf_dn", "sys/inst-VRF123"),
				),
			},
			{
				ResourceName:  "nxos_subinterface_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/intf/encrtd-[eth1/10.124]/rtvrfMbr",
			},
		},
	})
}

func testAccNxosSubinterfaceVRFConfig_minimum() string {
	return `
	resource "nxos_subinterface_vrf" "test" {
		interface_id = "eth1/10.124"
		vrf_dn = "sys/inst-VRF123"
	}
	`
}

func testAccNxosSubinterfaceVRFConfig_all() string {
	return `
	resource "nxos_subinterface_vrf" "test" {
		interface_id = "eth1/10.124"
		vrf_dn = "sys/inst-VRF123"
	}
	`
}
