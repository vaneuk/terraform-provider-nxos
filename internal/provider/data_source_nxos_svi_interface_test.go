// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosSVIInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosSVIInterfacePrerequisitesConfig + testAccDataSourceNxosSVIInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "interface_id", "vlan293"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "admin_state", "down"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "bandwidth", "1000"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "delay", "10"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "description", "My Description"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "medium", "bcast"),
					resource.TestCheckResourceAttr("data.nxos_svi_interface.test", "mtu", "9216"),
				),
			},
		},
	})
}

const testAccDataSourceNxosSVIInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/ifvlan"
  class_name = "fmInterfaceVlan"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

`

const testAccDataSourceNxosSVIInterfaceConfig = `

resource "nxos_svi_interface" "test" {
  interface_id = "vlan293"
  admin_state = "down"
  bandwidth = 1000
  delay = 10
  description = "My Description"
  medium = "bcast"
  mtu = 9216
  depends_on = [nxos_rest.PreReq0, ]
}

data "nxos_svi_interface" "test" {
  interface_id = "vlan293"
  depends_on = [nxos_svi_interface.test]
}
`
