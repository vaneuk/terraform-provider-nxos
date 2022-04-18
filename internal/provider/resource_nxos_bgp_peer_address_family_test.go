// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosBGPPeerAddressFamily(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosBGPPeerAddressFamilyPrerequisitesConfig + testAccNxosBGPPeerAddressFamilyConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "vrf", "default"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "address", "192.168.0.1"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "address_family", "ipv4-ucast"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "control", "rr-client"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "send_community_extended", "enabled"),
					resource.TestCheckResourceAttr("nxos_bgp_peer_address_family.test", "send_community_standard", "enabled"),
				),
			},
			{
				ResourceName:  "nxos_bgp_peer_address_family.test",
				ImportState:   true,
				ImportStateId: "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]",
			},
		},
	})
}

const testAccNxosBGPPeerAddressFamilyPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/bgp"
  class_name = "bgpEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/bgp/inst"
  class_name = "bgpInst"
  content = {
      adminSt = "enabled"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/bgp/inst/dom-[default]"
  class_name = "bgpDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]"
  class_name = "bgpPeer"
  content = {
      addr = "192.168.0.1"
      asn = "65001"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

func testAccNxosBGPPeerAddressFamilyConfig_minimum() string {
	return `
	resource "nxos_bgp_peer_address_family" "test" {
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}

func testAccNxosBGPPeerAddressFamilyConfig_all() string {
	return `
	resource "nxos_bgp_peer_address_family" "test" {
		vrf = "default"
		address = "192.168.0.1"
		address_family = "ipv4-ucast"
		control = "rr-client"
		send_community_extended = "enabled"
		send_community_standard = "enabled"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
	}
	`
}
