// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosOSPFArea(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosOSPFAreaPrerequisitesConfig + testAccDataSourceNxosOSPFAreaConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "area_id", "0.0.0.10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "authentication_type", "none"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "cost", "10"),
					resource.TestCheckResourceAttr("data.nxos_ospf_area.test", "type", "stub"),
				),
			},
		},
	})
}

const testAccDataSourceNxosOSPFAreaPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/ospf"
  class_name = "ospfEntity"
  content = {
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/ospf/inst-[OSPF1]"
  class_name = "ospfInst"
  content = {
      name = "OSPF1"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/ospf/inst-[OSPF1]/dom-[VRF1]"
  class_name = "ospfDom"
  content = {
      name = "VRF1"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosOSPFAreaConfig = `

resource "nxos_ospf_area" "test" {
  instance_name = "OSPF1"
  vrf_name = "VRF1"
  area_id = "0.0.0.10"
  authentication_type = "none"
  cost = 10
  type = "stub"
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_ospf_area" "test" {
  instance_name = "OSPF1"
  vrf_name = "VRF1"
  area_id = "0.0.0.10"
  depends_on = [nxos_ospf_area.test]
}
`
