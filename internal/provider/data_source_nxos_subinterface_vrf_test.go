// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceNxosSubinterfaceVRF(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosPhysicalInterfaceConfig_all(),
			},
			{
				Config: testAccNxosSubinterfaceConfig_all(),
			},
			{
				Config: testAccNxosPhysicalInterfaceConfig_all()+testAccNxosSubinterfaceConfig_all()+testAccNxosSubinterfaceVRFConfig_all(),
			},
			{
				Config: testAccDataSourceNxosSubinterfaceVRFConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.nxos_subinterface_vrf.test", "vrf_dn", "sys/inst-VRF123"),
				),
			},
		},
	})
}

const testAccDataSourceNxosSubinterfaceVRFConfig = `
data "nxos_subinterface_vrf" "test" {
  interface_id = "eth1/10.124"
}
`
