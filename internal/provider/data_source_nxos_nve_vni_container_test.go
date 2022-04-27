// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosNVEVNIContainer(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNxosNVEVNIContainerPrerequisitesConfig + testAccDataSourceNxosNVEVNIContainerConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

const testAccDataSourceNxosNVEVNIContainerPrerequisitesConfig = `
resource "nxos_rest" "PreReq0" {
  dn = "sys/fm/nvo"
  class_name = "fmNvo"
  delete = false
  content = {
      adminSt = "enabled"
  }
}

resource "nxos_rest" "PreReq1" {
  dn = "sys/fm/evpn"
  class_name = "fmEvpn"
  delete = false
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq0, ]
}

resource "nxos_rest" "PreReq2" {
  dn = "sys/eps/epId-[1]"
  class_name = "nvoEp"
  content = {
      adminSt = "enabled"
  }
  depends_on = [nxos_rest.PreReq1, ]
}

`

const testAccDataSourceNxosNVEVNIContainerConfig = `

resource "nxos_nve_vni_container" "test" {
  depends_on = [nxos_rest.PreReq0, nxos_rest.PreReq1, nxos_rest.PreReq2, ]
}

data "nxos_nve_vni_container" "test" {
  depends_on = [nxos_nve_vni_container.test]
}
`
