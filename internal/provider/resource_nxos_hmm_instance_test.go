// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosHMMInstance(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosHMMInstancePrerequisitesConfig + testAccNxosHMMInstanceConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_hmm_instance.test", "admin_state", "enabled"),
					resource.TestCheckResourceAttr("nxos_hmm_instance.test", "anycast_mac", "20:20:00:00:10:10"),
				),
			},
			{
				ResourceName:  "nxos_hmm_instance.test",
				ImportState:   true,
				ImportStateId: "sys/hmm/fwdinst",
			},
		},
	})
}

const testAccNxosHMMInstancePrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/hmm"
  class_name = "fmHmm"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/hmm"
  class_name = "hmmEntity"
  content = {
  }
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, ]
}

`

func testAccNxosHMMInstanceConfig_minimum() string {
	return `
	resource "nxos_hmm_instance" "test" {
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}

func testAccNxosHMMInstanceConfig_all() string {
	return `
	resource "nxos_hmm_instance" "test" {
		admin_state = "enabled"
		anycast_mac = "20:20:00:00:10:10"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
	}
	`
}
