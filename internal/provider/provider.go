// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/frankgreco/terraform-helpers/validators"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netascode/go-nxos"
)

// provider satisfies the tfsdk.Provider interface and usually is included
// with all Resource and DataSource implementations.
type provider struct {
	client  *nxos.Client
	devices map[string]string

	// configured is set to true at the end of the Configure method.
	// This can be used in Resource and DataSource implementations to verify
	// that the provider was previously configured.
	configured bool

	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// providerData can be used to store data from the Terraform configuration.
type providerData struct {
	Username types.String         `tfsdk:"username"`
	Password types.String         `tfsdk:"password"`
	URL      types.String         `tfsdk:"url"`
	Insecure types.Bool           `tfsdk:"insecure"`
	Retries  types.Int64          `tfsdk:"retries"`
	Devices  []providerDataDevice `tfsdk:"devices"`
}

type providerDataDevice struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

func (p *provider) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"username": {
				MarkdownDescription: "Username for the NXOS device account. This can also be set as the NXOS_USERNAME environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"password": {
				MarkdownDescription: "Password for the NXOS device account. This can also be set as the NXOS_PASSWORD environment variable.",
				Type:                types.StringType,
				Optional:            true,
				Sensitive:           true,
			},
			"url": {
				MarkdownDescription: "URL of the Cisco NXOS device. This can also be set as the NXOS_URL environment variable.",
				Type:                types.StringType,
				Optional:            true,
			},
			"insecure": {
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the NXOS_INSECURE environment variable. Defaults to `true`.",
				Type:                types.BoolType,
				Optional:            true,
			},
			"retries": {
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the NXOS_RETRIES environment variable. Defaults to `3`.",
				Type:                types.Int64Type,
				Optional:            true,
				Validators: []tfsdk.AttributeValidator{
					validators.Range(0, 9),
				},
			},
			"devices": {
				MarkdownDescription: "This can be used to manage a list of devices from a single provider. All devices must use the same credentials. Each resource and data source has an optional attribute named `device`, which can then select a device by its name from this list.",
				Optional:            true,
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						MarkdownDescription: "Device name.",
						Type:                types.StringType,
						Required:            true,
					},
					"url": {
						MarkdownDescription: "URL of the Cisco NXOS device.",
						Type:                types.StringType,
						Required:            true,
					},
				}, tfsdk.ListNestedAttributesOptions{}),
			},
		},
	}, nil
}

func (p *provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, resp *tfsdk.ConfigureProviderResponse) {
	// Retrieve provider data from configuration
	var config providerData
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Null {
		username = os.Getenv("NXOS_USERNAME")
	} else {
		username = config.Username.Value
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Null {
		password = os.Getenv("NXOS_PASSWORD")
	} else {
		password = config.Password.Value
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.Null {
		url = os.Getenv("NXOS_URL")
	} else {
		url = config.URL.Value
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.Null {
		insecureStr := os.Getenv("NXOS_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.Value
	}

	var retries int64
	if config.Retries.Unknown {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.Null {
		retriesStr := os.Getenv("NXOS_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.Value
	}

	// Create a new NXOS client and set it to the provider client
	c, err := nxos.NewClient(url, username, password, insecure, nxos.MaxRetries(int(retries)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create nxos client:\n\n"+err.Error(),
		)
		return
	}

	devices := make(map[string]string)
	for _, device := range config.Devices {
		devices[device.Name.Value] = device.URL.Value
	}

	p.client = &c
	p.devices = devices
	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"nxos_rest":                                                 resourceRestType{},
		"nxos_bridge_domain":                                        resourceBridgeDomainType{},
		"nxos_default_qos_class_map":                                resourceDefaultQOSClassMapType{},
		"nxos_default_qos_class_map_dscp":                           resourceDefaultQOSClassMapDSCPType{},
		"nxos_default_qos_policy_interface_in":                      resourceDefaultQOSPolicyInterfaceInType{},
		"nxos_default_qos_policy_interface_in_policy_map":           resourceDefaultQOSPolicyInterfaceInPolicyMapType{},
		"nxos_default_qos_policy_map":                               resourceDefaultQOSPolicyMapType{},
		"nxos_default_qos_policy_map_match_class_map":               resourceDefaultQOSPolicyMapMatchClassMapType{},
		"nxos_default_qos_policy_map_match_class_map_police":        resourceDefaultQOSPolicyMapMatchClassMapPoliceType{},
		"nxos_default_qos_policy_map_match_class_map_set_qos_group": resourceDefaultQOSPolicyMapMatchClassMapSetQOSGroupType{},
		"nxos_dhcp_relay_address":                                   resourceDHCPRelayAddressType{},
		"nxos_dhcp_relay_interface":                                 resourceDHCPRelayInterfaceType{},
		"nxos_ethernet":                                             resourceEthernetType{},
		"nxos_feature_bfd":                                          resourceFeatureBFDType{},
		"nxos_feature_bgp":                                          resourceFeatureBGPType{},
		"nxos_feature_dhcp":                                         resourceFeatureDHCPType{},
		"nxos_feature_evpn":                                         resourceFeatureEVPNType{},
		"nxos_feature_hsrp":                                         resourceFeatureHSRPType{},
		"nxos_feature_interface_vlan":                               resourceFeatureInterfaceVLANType{},
		"nxos_feature_isis":                                         resourceFeatureISISType{},
		"nxos_feature_lacp":                                         resourceFeatureLACPType{},
		"nxos_feature_lldp":                                         resourceFeatureLLDPType{},
		"nxos_feature_macsec":                                       resourceFeatureMACsecType{},
		"nxos_feature_ospf":                                         resourceFeatureOSPFType{},
		"nxos_feature_pim":                                          resourceFeaturePIMType{},
		"nxos_ipv4_interface":                                       resourceIPv4InterfaceType{},
		"nxos_ipv4_interface_address":                               resourceIPv4InterfaceAddressType{},
		"nxos_loopback_interface":                                   resourceLoopbackInterfaceType{},
		"nxos_loopback_interface_vrf":                               resourceLoopbackInterfaceVRFType{},
		"nxos_ospf":                                                 resourceOSPFType{},
		"nxos_ospf_area":                                            resourceOSPFAreaType{},
		"nxos_ospf_authentication":                                  resourceOSPFAuthenticationType{},
		"nxos_ospf_instance":                                        resourceOSPFInstanceType{},
		"nxos_ospf_interface":                                       resourceOSPFInterfaceType{},
		"nxos_ospf_vrf":                                             resourceOSPFVRFType{},
		"nxos_physical_interface":                                   resourcePhysicalInterfaceType{},
		"nxos_physical_interface_vrf":                               resourcePhysicalInterfaceVRFType{},
		"nxos_pim":                                                  resourcePIMType{},
		"nxos_pim_instance":                                         resourcePIMInstanceType{},
		"nxos_pim_interface":                                        resourcePIMInterfaceType{},
		"nxos_pim_ssm_policy":                                       resourcePIMSSMPolicyType{},
		"nxos_pim_ssm_range":                                        resourcePIMSSMRangeType{},
		"nxos_pim_static_rp":                                        resourcePIMStaticRPType{},
		"nxos_pim_static_rp_group_list":                             resourcePIMStaticRPGroupListType{},
		"nxos_pim_static_rp_policy":                                 resourcePIMStaticRPPolicyType{},
		"nxos_pim_vrf":                                              resourcePIMVRFType{},
		"nxos_queuing_qos_policy_map":                               resourceQueuingQOSPolicyMapType{},
		"nxos_queuing_qos_policy_map_match_class_map":               resourceQueuingQOSPolicyMapMatchClassMapType{},
		"nxos_queuing_qos_policy_map_match_class_map_priority":      resourceQueuingQOSPolicyMapMatchClassMapPriorityType{},
		"nxos_queuing_qos_policy_map_match_class_map_remaining_bandwidth": resourceQueuingQOSPolicyMapMatchClassMapRemainingBandwidthType{},
		"nxos_queuing_qos_policy_system_out":                              resourceQueuingQOSPolicySystemOutType{},
		"nxos_queuing_qos_policy_system_out_policy_map":                   resourceQueuingQOSPolicySystemOutPolicyMapType{},
		"nxos_subinterface":      resourceSubinterfaceType{},
		"nxos_subinterface_vrf":  resourceSubinterfaceVRFType{},
		"nxos_svi_interface":     resourceSVIInterfaceType{},
		"nxos_svi_interface_vrf": resourceSVIInterfaceVRFType{},
		"nxos_vrf":               resourceVRFType{},
		"nxos_vrf_container":     resourceVRFContainerType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"nxos_rest":                                                 dataSourceRestType{},
		"nxos_bridge_domain":                                        dataSourceBridgeDomainType{},
		"nxos_default_qos_class_map":                                dataSourceDefaultQOSClassMapType{},
		"nxos_default_qos_class_map_dscp":                           dataSourceDefaultQOSClassMapDSCPType{},
		"nxos_default_qos_policy_interface_in":                      dataSourceDefaultQOSPolicyInterfaceInType{},
		"nxos_default_qos_policy_interface_in_policy_map":           dataSourceDefaultQOSPolicyInterfaceInPolicyMapType{},
		"nxos_default_qos_policy_map":                               dataSourceDefaultQOSPolicyMapType{},
		"nxos_default_qos_policy_map_match_class_map":               dataSourceDefaultQOSPolicyMapMatchClassMapType{},
		"nxos_default_qos_policy_map_match_class_map_police":        dataSourceDefaultQOSPolicyMapMatchClassMapPoliceType{},
		"nxos_default_qos_policy_map_match_class_map_set_qos_group": dataSourceDefaultQOSPolicyMapMatchClassMapSetQOSGroupType{},
		"nxos_dhcp_relay_address":                                   dataSourceDHCPRelayAddressType{},
		"nxos_dhcp_relay_interface":                                 dataSourceDHCPRelayInterfaceType{},
		"nxos_ethernet":                                             dataSourceEthernetType{},
		"nxos_feature_bfd":                                          dataSourceFeatureBFDType{},
		"nxos_feature_bgp":                                          dataSourceFeatureBGPType{},
		"nxos_feature_dhcp":                                         dataSourceFeatureDHCPType{},
		"nxos_feature_evpn":                                         dataSourceFeatureEVPNType{},
		"nxos_feature_hsrp":                                         dataSourceFeatureHSRPType{},
		"nxos_feature_interface_vlan":                               dataSourceFeatureInterfaceVLANType{},
		"nxos_feature_isis":                                         dataSourceFeatureISISType{},
		"nxos_feature_lacp":                                         dataSourceFeatureLACPType{},
		"nxos_feature_lldp":                                         dataSourceFeatureLLDPType{},
		"nxos_feature_macsec":                                       dataSourceFeatureMACsecType{},
		"nxos_feature_ospf":                                         dataSourceFeatureOSPFType{},
		"nxos_feature_pim":                                          dataSourceFeaturePIMType{},
		"nxos_ipv4_interface":                                       dataSourceIPv4InterfaceType{},
		"nxos_ipv4_interface_address":                               dataSourceIPv4InterfaceAddressType{},
		"nxos_loopback_interface":                                   dataSourceLoopbackInterfaceType{},
		"nxos_loopback_interface_vrf":                               dataSourceLoopbackInterfaceVRFType{},
		"nxos_ospf":                                                 dataSourceOSPFType{},
		"nxos_ospf_area":                                            dataSourceOSPFAreaType{},
		"nxos_ospf_authentication":                                  dataSourceOSPFAuthenticationType{},
		"nxos_ospf_instance":                                        dataSourceOSPFInstanceType{},
		"nxos_ospf_interface":                                       dataSourceOSPFInterfaceType{},
		"nxos_ospf_vrf":                                             dataSourceOSPFVRFType{},
		"nxos_physical_interface":                                   dataSourcePhysicalInterfaceType{},
		"nxos_physical_interface_vrf":                               dataSourcePhysicalInterfaceVRFType{},
		"nxos_pim":                                                  dataSourcePIMType{},
		"nxos_pim_instance":                                         dataSourcePIMInstanceType{},
		"nxos_pim_interface":                                        dataSourcePIMInterfaceType{},
		"nxos_pim_ssm_policy":                                       dataSourcePIMSSMPolicyType{},
		"nxos_pim_ssm_range":                                        dataSourcePIMSSMRangeType{},
		"nxos_pim_static_rp":                                        dataSourcePIMStaticRPType{},
		"nxos_pim_static_rp_group_list":                             dataSourcePIMStaticRPGroupListType{},
		"nxos_pim_static_rp_policy":                                 dataSourcePIMStaticRPPolicyType{},
		"nxos_pim_vrf":                                              dataSourcePIMVRFType{},
		"nxos_queuing_qos_policy_map":                               dataSourceQueuingQOSPolicyMapType{},
		"nxos_queuing_qos_policy_map_match_class_map":               dataSourceQueuingQOSPolicyMapMatchClassMapType{},
		"nxos_queuing_qos_policy_map_match_class_map_priority":      dataSourceQueuingQOSPolicyMapMatchClassMapPriorityType{},
		"nxos_queuing_qos_policy_map_match_class_map_remaining_bandwidth": dataSourceQueuingQOSPolicyMapMatchClassMapRemainingBandwidthType{},
		"nxos_queuing_qos_policy_system_out":                              dataSourceQueuingQOSPolicySystemOutType{},
		"nxos_queuing_qos_policy_system_out_policy_map":                   dataSourceQueuingQOSPolicySystemOutPolicyMapType{},
		"nxos_subinterface":      dataSourceSubinterfaceType{},
		"nxos_subinterface_vrf":  dataSourceSubinterfaceVRFType{},
		"nxos_svi_interface":     dataSourceSVIInterfaceType{},
		"nxos_svi_interface_vrf": dataSourceSVIInterfaceVRFType{},
		"nxos_vrf":               dataSourceVRFType{},
		"nxos_vrf_container":     dataSourceVRFContainerType{},
	}, nil
}

func New(version string) func() tfsdk.Provider {
	return func() tfsdk.Provider {
		return &provider{
			version: version,
		}
	}
}

// convertProviderType is a helper function for NewResource and NewDataSource
// implementations to associate the concrete provider type. Alternatively,
// this helper can be skipped and the provider type can be directly type
// asserted (e.g. provider: in.(*provider)), however using this can prevent
// potential panics.
func convertProviderType(in tfsdk.Provider) (provider, diag.Diagnostics) {
	var diags diag.Diagnostics

	p, ok := in.(*provider)

	if !ok {
		diags.AddError(
			"Unexpected Provider Instance Type",
			fmt.Sprintf("While creating the data source or resource, an unexpected provider type (%T) was received. This is always a bug in the provider code and should be reported to the provider developers.", p),
		)
		return provider{}, diags
	}

	if p == nil {
		diags.AddError(
			"Unexpected Provider Instance Type",
			"While creating the data source or resource, an unexpected empty provider instance was received. This is always a bug in the provider code and should be reported to the provider developers.",
		)
		return provider{}, diags
	}

	return *p, diags
}
