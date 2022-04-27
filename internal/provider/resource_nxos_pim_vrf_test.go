// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosPIMVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMVRFPrerequisitesConfig + testAccNxosPIMVRFConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_vrf.test", "name", "default"),
					resource.TestCheckResourceAttr("nxos_pim_vrf.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_pim_vrf.test", "bfd", "true"),
				),
			},
			{
				ResourceName:  "nxos_pim_vrf.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst/dom-[default]",
			},
		},
	})
}

const testAccNxosPIMVRFPrerequisitesConfig = `
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

`

func testAccNxosPIMVRFConfig_minimum() string {
	return `
	resource "nxos_pim_vrf" "test" {
		name = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosPIMVRFConfig_all() string {
	return `
	resource "nxos_pim_vrf" "test" {
		name = "default"
		admin_state = "enabled"
		bfd = true
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
