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
	client nxos.Client

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
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	URL      types.String `tfsdk:"url"`
	Insecure types.Bool   `tfsdk:"insecure"`
	Retries  types.Int64  `tfsdk:"retries"`
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
			retries = 4
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

	p.client = c

	p.configured = true
}

func (p *provider) GetResources(ctx context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{
		"nxos_rest":                   resourceRestType{},
		"nxos_physical_interface":     resourcePhysicalInterfaceType{},
		"nxos_ipv4_interface":         resourceIPv4InterfaceType{},
		"nxos_ipv4_interface_address": resourceIPv4InterfaceAddressType{},
		"nxos_vrf":                    resourceVRFType{},
		"nxos_vrf_container":          resourceVRFContainerType{},
		"nxos_bridge_domain":          resourceBridgeDomainType{},
		"nxos_svi_interface":          resourceSVIInterfaceType{},
		"nxos_svi_interface_vrf":      resourceSVIInterfaceVRFType{},
		"nxos_feature_ospf":           resourceFeatureOSPFType{},
		"nxos_feature_interface_vlan": resourceFeatureInterfaceVLANType{},
		"nxos_feature_pim":            resourceFeaturePIMType{},
		"nxos_feature_dhcp":           resourceFeatureDHCPType{},
		"nxos_feature_lldp":           resourceFeatureLLDPType{},
		"nxos_loopback_interface":     resourceLoopbackInterfaceType{},
		"nxos_loopback_interface_vrf": resourceLoopbackInterfaceVRFType{},
	}, nil
}

func (p *provider) GetDataSources(ctx context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"nxos_rest":                   dataSourceRestType{},
		"nxos_physical_interface":     dataSourcePhysicalInterfaceType{},
		"nxos_ipv4_interface":         dataSourceIPv4InterfaceType{},
		"nxos_ipv4_interface_address": dataSourceIPv4InterfaceAddressType{},
		"nxos_vrf":                    dataSourceVRFType{},
		"nxos_vrf_container":          dataSourceVRFContainerType{},
		"nxos_bridge_domain":          dataSourceBridgeDomainType{},
		"nxos_svi_interface":          dataSourceSVIInterfaceType{},
		"nxos_svi_interface_vrf":      dataSourceSVIInterfaceVRFType{},
		"nxos_feature_ospf":           dataSourceFeatureOSPFType{},
		"nxos_feature_interface_vlan": dataSourceFeatureInterfaceVLANType{},
		"nxos_feature_pim":            dataSourceFeaturePIMType{},
		"nxos_feature_dhcp":           dataSourceFeatureDHCPType{},
		"nxos_feature_lldp":           dataSourceFeatureLLDPType{},
		"nxos_loopback_interface":     dataSourceLoopbackInterfaceType{},
		"nxos_loopback_interface_vrf": dataSourceLoopbackInterfaceVRFType{},
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
