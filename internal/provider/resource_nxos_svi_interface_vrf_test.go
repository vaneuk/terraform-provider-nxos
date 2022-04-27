// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosSVIInterfaceVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosSVIInterfaceVRFPrerequisitesConfig + testAccNxosSVIInterfaceVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_svi_interface_vrf.test", "interface_id", "vlan293"),
					resource.TestCheckResourceAttr("nxos_svi_interface_vrf.test", "vrf_dn", "sys/inst-VRF123"),
				),
			},
			{
				ResourceName:  "nxos_svi_interface_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/intf/svi-[vlan293]/rtvrfMbr",
			},
		},
	})
}

const testAccNxosSVIInterfaceVRFPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ifvlan"
  class_name = "fmInterfaceVlan"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/svi-[vlan293]"
  class_name = "sviIf"
  content = {
      id = "vlan293"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

`

func testAccNxosSVIInterfaceVRFConfig_minimum() string {
	return `
	resource "nxos_svi_interface_vrf" "test" {
		interface_id = "vlan293"
		vrf_dn = "sys/inst-VRF123"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}

func testAccNxosSVIInterfaceVRFConfig_all() string {
	return `
	resource "nxos_svi_interface_vrf" "test" {
		interface_id = "vlan293"
		vrf_dn = "sys/inst-VRF123"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
	}
	`
}
