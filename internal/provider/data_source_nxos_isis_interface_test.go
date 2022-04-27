// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosISISInterface(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosISISInterfacePrerequisitesConfig + testAccDataSourceNxosISISInterfaceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "interface_id", "eth1/10"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_check", "false"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_check_l1", "false"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_check_l2", "false"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_type", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_type_l1", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "authentication_type_l2", "unknown"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "circuit_type", "l2"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "vrf", "default"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_interval", "20"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_interval_l1", "20"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_interval_l2", "20"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_multiplier", "4"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_multiplier_l1", "4"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_multiplier_l2", "4"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "hello_padding", "never"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "metric_l1", "1000"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "metric_l2", "1000"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "mtu_check", "true"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "mtu_check_l1", "true"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "mtu_check_l2", "true"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "network_type_p2p", "on"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "passive", "l1"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "priority_l1", "80"),
					resource.TestCheckResourceAttr("data.nxos_isis_interface.test", "priority_l2", "80"),
				),
			},
		},
	})
}

const testAccDataSourceNxosISISInterfacePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/isis"
  class_name = "fmIsis"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/isis"
  class_name = "isisEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, ]
}

`

const testAccDataSourceNxosISISInterfaceConfig = `

resource "nxos_isis_interface" "test" {
  interface_id = "eth1/10"
  authentication_check = false
  authentication_check_l1 = false
  authentication_check_l2 = false
  authentication_key = ""
  authentication_key_l1 = ""
  authentication_key_l2 = ""
  authentication_type = "unknown"
  authentication_type_l1 = "unknown"
  authentication_type_l2 = "unknown"
  circuit_type = "l2"
  vrf = "default"
  hello_interval = 20
  hello_interval_l1 = 20
  hello_interval_l2 = 20
  hello_multiplier = 4
  hello_multiplier_l1 = 4
  hello_multiplier_l2 = 4
  hello_padding = "never"
  metric_l1 = 1000
  metric_l2 = 1000
  mtu_check = true
  mtu_check_l1 = true
  mtu_check_l2 = true
  network_type_p2p = "on"
  passive = "l1"
  priority_l1 = 80
  priority_l2 = 80
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

data "nxos_isis_interface" "test" {
  interface_id = "eth1/10"
  depends_on = [nxos_isis_interface.test]
}
`
