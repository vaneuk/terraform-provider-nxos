// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosPIMInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosPIMInterfacePrerequisitesConfig + testAccDataSourceNxosPIMInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "bfd", "enabled"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "dr_priority", "10"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "passive", "false"),
					resource.TestCheckResourceAttr("data.nxos_pim_interface.test", "sparse_mode", "true"),
				),
			},
		},
	})
}

const testAccDataSourceNxosPIMInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/pim"
  class_name = "fmPim"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/pim"
  class_name = "pimEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/pim/inst"
  class_name = "pimInst"
  content = {
  }
  depends_on = [nxos_rest.PreReq1, ]
}

resource "nxos_rest" "PreReq3" {
  dn = "sys/pim/inst/dom-[default]"
  class_name = "pimDom"
  content = {
      name = "default"
  }
  depends_on = [nxos_rest.PreReq2, ]
}

`

const testAccDataSourceNxosPIMInterfaceConfig = `

resource "nxos_pim_interface" "test" {
  vrf_name = "default"
  interface_id = "eth1/10"
  admin_state = "enabled"
  bfd = "enabled"
  dr_priority = 10
  passive = false
  sparse_mode = true
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, ]
}

data "nxos_pim_interface" "test" {
  vrf_name = "default"
  interface_id = "eth1/10"
  depends_on = [nxos_pim_interface.test]
}
`
