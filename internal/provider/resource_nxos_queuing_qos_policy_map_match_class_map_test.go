// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNxosQueuingQOSPolicyMapMatchClassMap(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNxosQueuingQOSPolicyMapConfig_all(),
			},
			{
				Config: testAccNxosQueuingQOSPolicyMapConfig_all() + testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "policy_map_name", "PM1"),
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "name", "c-out-q1"),
				),
			},
			{
				Config: testAccNxosQueuingQOSPolicyMapConfig_all() + testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "policy_map_name", "PM1"),
					resource.TestCheckResourceAttr("nxos_queuing_qos_policy_map_match_class_map.test", "name", "c-out-q1"),
				),
			},
			{
				ResourceName:  "nxos_queuing_qos_policy_map_match_class_map.test",
				ImportState:   true,
				ImportStateId: "sys/ipqos/queuing/p/name-[PM1]/cmap-[c-out-q1]",
			},
		},
	})
}

func testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_minimum() string {
	return `
	resource "nxos_queuing_qos_policy_map_match_class_map" "test" {
		policy_map_name = "PM1"
		name = "c-out-q1"
	}
	`
}

func testAccNxosQueuingQOSPolicyMapMatchClassMapConfig_all() string {
	return `
	resource "nxos_queuing_qos_policy_map_match_class_map" "test" {
		policy_map_name = "PM1"
		name = "c-out-q1"
	}
	`
}
