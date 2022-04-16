// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosOSPFInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosOSPFInterfacePrerequisitesConfig + testAccNxosOSPFInterfaceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "instance_name", "OSPF1"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "vrf_name", "VRF1"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "advertise_secondaries", "false"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "area", "0.0.0.10"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "bfd", "disabled"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "cost", "1000"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "dead_interval", "60"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "hello_interval", "15"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "network_type", "p2p"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "passive", "enabled"),
					resource.TestCheckResourceAttr("nxos_ospf_interface.test", "priority", "10"),
				),
			},
			{
				ResourceName:  "nxos_ospf_interface.test",
				ImportState:   true,
				ImportStateId: "sys/ospf/inst-[OSPF1]/dom-[VRF1]/if-[eth1/10]",
			},
		},
	})
}

const testAccNxosOSPFInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/bfd"
  class_name = "fmBfd"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]"
  class_name = "ospfDom"
  content = {
      name = "VRF1"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

resource "nxos_rest" "PreReq4" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
  depends_on = [nxos_rest.PreReq3, ]
}

resource "nxos_rest" "PreReq5" {
  dn = "sys/intf/phys-[eth1/10]/rtvrfMbr"
  class_name = "nwRtVrfMbr"
  content = {
      tDn = "sys/inst-VRF1"
  }
  depends_on = [nxos_rest.PreReq4, ]
}

`

func testAccNxosOSPFInterfaceConfig_minimum() string {
	return `
	resource "nxos_ospf_interface" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		interface_id = "eth1/10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}

func testAccNxosOSPFInterfaceConfig_all() string {
	return `
	resource "nxos_ospf_interface" "test" {
		instance_name = "OSPF1"
		vrf_name = "VRF1"
		interface_id = "eth1/10"
		advertise_secondaries = false
		area = "0.0.0.10"
		bfd = "disabled"
		cost = 1000
		dead_interval = 60
		hello_interval = 15
		network_type = "p2p"
		passive = "enabled"
		priority = 10
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, nxos_rest.PreReq5, ]
	}
	`
}
