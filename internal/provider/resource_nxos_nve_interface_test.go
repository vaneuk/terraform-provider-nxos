// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosNVEInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosNVEInterfacePrerequisitesConfig + testAccNxosNVEInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "advertise_virtual_mac", "true"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "hold_down_time", "60"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "host_reachability_protocol", "bgp"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "ingress_replication_protocol_bgp", "true"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "multicast_group_l2", "0.0.0.0"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "multicast_group_l3", "0.0.0.0"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "multisite_source_interface", "unspecified"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "source_interface", "lo0"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "suppress_arp", "true"),
					resource.TestCheckResourceAttr("nxos_nve_interface.test", "suppress_mac_route", "false"),
				),
			},
			{
				ResourceName:  "nxos_nve_interface.test",
				ImportState:   true,
				ImportStateId: "sys/eps/epId-[1]",
			},
		},
	})
}

const testAccNxosNVEInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/nvo"
  class_name = "fmNvo"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  delete = false
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

`

func testAccNxosNVEInterfaceConfig_minimum() string {
	return `
	resource "nxos_nve_interface" "test" {
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosNVEInterfaceConfig_all() string {
	return `
	resource "nxos_nve_interface" "test" {
		admin_state = "enabled"
		advertise_virtual_mac = true
		hold_down_time = 60
		host_reachability_protocol = "bgp"
		ingress_replication_protocol_bgp = true
		multicast_group_l2 = "0.0.0.0"
		multicast_group_l3 = "0.0.0.0"
		multisite_source_interface = "unspecified"
		source_interface = "lo0"
		suppress_arp = true
		suppress_mac_route = false
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
