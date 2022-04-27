// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosPIMSSMRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPIMSSMRangePrerequisitesConfig + testAccNxosPIMSSMRangeConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "vrf_name", "default"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "group_list_1", "232.0.0.0/8"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "group_list_2", "233.0.0.0/8"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "group_list_3", "0.0.0.0"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "group_list_4", "0.0.0.0"),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "prefix_list", ""),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "route_map", ""),
					resource.TestCheckResourceAttr("nxos_pim_ssm_range.test", "ssm_none", "false"),
				),
			},
			{
				ResourceName:  "nxos_pim_ssm_range.test",
				ImportState:   true,
				ImportStateId: "sys/pim/inst/dom-[default]/ssm/range",
			},
		},
	})
}

const testAccNxosPIMSSMRangePrerequisitesConfig = `
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

resource "nxos_rest" "PreReq4" {
  dn = "sys/pim/inst/dom-[default]/ssm"
  class_name = "pimSSMPatP"
  content = {
  }
  depends_on = [nxos_rest.PreReq3, ]
}

`

func testAccNxosPIMSSMRangeConfig_minimum() string {
	return `
	resource "nxos_pim_ssm_range" "test" {
		vrf_name = "default"
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}

func testAccNxosPIMSSMRangeConfig_all() string {
	return `
	resource "nxos_pim_ssm_range" "test" {
		vrf_name = "default"
		group_list_1 = "232.0.0.0/8"
		group_list_2 = "233.0.0.0/8"
		group_list_3 = "0.0.0.0"
		group_list_4 = "0.0.0.0"
		prefix_list = ""
		route_map = ""
		ssm_none = false
  		depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, nxos_rest.PreReq3, nxos_rest.PreReq4, ]
	}
	`
}
