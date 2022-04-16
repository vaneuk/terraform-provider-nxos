// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosDHCPRelayAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosDHCPRelayAddressPrerequisitesConfig + testAccDataSourceNxosDHCPRelayAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_address.test", "vrf", "VRF1"),
					resource.TestCheckResourceAttr("data.nxos_dhcp_relay_address.test", "address", "1.1.1.1"),
				),
			},
		},
	})
}

const testAccDataSourceNxosDHCPRelayAddressPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/dhcp"
  class_name = "fmDhcp"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/intf/phys-[eth1/10]"
  class_name = "l1PhysIf"
  content = {
      layer = "Layer3"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/dhcp/inst/relayif-[eth1/10]"
  class_name = "dhcpRelayIf"
  content = {
      id = "eth1/10"
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosDHCPRelayAddressConfig = `

resource "nxos_dhcp_relay_address" "test" {
  interface_id = "eth1/10"
  vrf = "VRF1"
  address = "1.1.1.1"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_dhcp_relay_address" "test" {
  interface_id = "eth1/10"
  vrf = "VRF1"
  address = "1.1.1.1"
  depends_on = [nxos_dhcp_relay_address.test]
}
`
