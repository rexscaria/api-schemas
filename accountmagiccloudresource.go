// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// AccountMagicCloudResourceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicCloudResourceService] method instead.
type AccountMagicCloudResourceService struct {
	Options []option.RequestOption
}

// NewAccountMagicCloudResourceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicCloudResourceService(opts ...option.RequestOption) (r *AccountMagicCloudResourceService) {
	r = &AccountMagicCloudResourceService{}
	r.Options = opts
	return
}

// Read an resource from the Resource Catalog (Closed Beta)
func (r *AccountMagicCloudResourceService) Get(ctx context.Context, accountID string, resourceID string, query AccountMagicCloudResourceGetParams, opts ...option.RequestOption) (res *AccountMagicCloudResourceGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/%s", accountID, resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List resources in the Resource Catalog (Closed Beta)
func (r *AccountMagicCloudResourceService) List(ctx context.Context, accountID string, query AccountMagicCloudResourceListParams, opts ...option.RequestOption) (res *AccountMagicCloudResourceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Export resources in the Resource Catalog as a JSON file (Closed Beta)
func (r *AccountMagicCloudResourceService) Export(ctx context.Context, accountID string, query AccountMagicCloudResourceExportParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/octet-stream")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/export", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Preview Rego query result against the latest resource catalog (Closed Beta)
func (r *AccountMagicCloudResourceService) PreviewPolicy(ctx context.Context, accountID string, body AccountMagicCloudResourcePreviewPolicyParams, opts ...option.RequestOption) (res *AccountMagicCloudResourcePreviewPolicyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cloud/resources/policy-preview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type McnCloudPlatformClient struct {
	ID         string                           `json:"id,required" format:"uuid"`
	ClientType McnCloudPlatformClientClientType `json:"client_type,required"`
	Name       string                           `json:"name,required"`
	JSON       mcnCloudPlatformClientJSON       `json:"-"`
}

// mcnCloudPlatformClientJSON contains the JSON metadata for the struct
// [McnCloudPlatformClient]
type mcnCloudPlatformClientJSON struct {
	ID          apijson.Field
	ClientType  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnCloudPlatformClient) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnCloudPlatformClientJSON) RawJSON() string {
	return r.raw
}

type McnCloudPlatformClientClientType string

const (
	McnCloudPlatformClientClientTypeMagicWanCloudOnramp McnCloudPlatformClientClientType = "MAGIC_WAN_CLOUD_ONRAMP"
)

func (r McnCloudPlatformClientClientType) IsKnown() bool {
	switch r {
	case McnCloudPlatformClientClientTypeMagicWanCloudOnramp:
		return true
	}
	return false
}

type McnResourceDetails struct {
	ID                  string                                   `json:"id,required" format:"uuid"`
	AccountID           string                                   `json:"account_id,required"`
	CloudType           McnCloudType                             `json:"cloud_type,required"`
	Config              map[string]interface{}                   `json:"config,required"`
	DeploymentProvider  string                                   `json:"deployment_provider,required" format:"uuid"`
	Managed             bool                                     `json:"managed,required"`
	MonthlyCostEstimate McnCost                                  `json:"monthly_cost_estimate,required"`
	Name                string                                   `json:"name,required"`
	NativeID            string                                   `json:"native_id,required"`
	Observations        map[string]McnResourceDetailsObservation `json:"observations,required"`
	ProviderIDs         []string                                 `json:"provider_ids,required" format:"uuid"`
	ProviderNamesByID   map[string]string                        `json:"provider_names_by_id,required"`
	Region              string                                   `json:"region,required"`
	ResourceGroup       string                                   `json:"resource_group,required"`
	ResourceType        McnResourceType                          `json:"resource_type,required"`
	Sections            []McnResourceDetailsSection              `json:"sections,required"`
	State               map[string]interface{}                   `json:"state,required"`
	Tags                map[string]string                        `json:"tags,required"`
	UpdatedAt           string                                   `json:"updated_at,required"`
	URL                 string                                   `json:"url,required"`
	ManagedBy           []McnCloudPlatformClient                 `json:"managed_by"`
	JSON                mcnResourceDetailsJSON                   `json:"-"`
}

// mcnResourceDetailsJSON contains the JSON metadata for the struct
// [McnResourceDetails]
type mcnResourceDetailsJSON struct {
	ID                  apijson.Field
	AccountID           apijson.Field
	CloudType           apijson.Field
	Config              apijson.Field
	DeploymentProvider  apijson.Field
	Managed             apijson.Field
	MonthlyCostEstimate apijson.Field
	Name                apijson.Field
	NativeID            apijson.Field
	Observations        apijson.Field
	ProviderIDs         apijson.Field
	ProviderNamesByID   apijson.Field
	Region              apijson.Field
	ResourceGroup       apijson.Field
	ResourceType        apijson.Field
	Sections            apijson.Field
	State               apijson.Field
	Tags                apijson.Field
	UpdatedAt           apijson.Field
	URL                 apijson.Field
	ManagedBy           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *McnResourceDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsJSON) RawJSON() string {
	return r.raw
}

type McnResourceDetailsObservation struct {
	FirstObservedAt string                            `json:"first_observed_at,required"`
	LastObservedAt  string                            `json:"last_observed_at,required"`
	ProviderID      string                            `json:"provider_id,required" format:"uuid"`
	ResourceID      string                            `json:"resource_id,required" format:"uuid"`
	JSON            mcnResourceDetailsObservationJSON `json:"-"`
}

// mcnResourceDetailsObservationJSON contains the JSON metadata for the struct
// [McnResourceDetailsObservation]
type mcnResourceDetailsObservationJSON struct {
	FirstObservedAt apijson.Field
	LastObservedAt  apijson.Field
	ProviderID      apijson.Field
	ResourceID      apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *McnResourceDetailsObservation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsObservationJSON) RawJSON() string {
	return r.raw
}

type McnResourceDetailsSection struct {
	HiddenItems  []McnResourceDetailsSectionItem `json:"hidden_items,required"`
	Name         string                          `json:"name,required"`
	VisibleItems []McnResourceDetailsSectionItem `json:"visible_items,required"`
	HelpText     string                          `json:"help_text"`
	JSON         mcnResourceDetailsSectionJSON   `json:"-"`
}

// mcnResourceDetailsSectionJSON contains the JSON metadata for the struct
// [McnResourceDetailsSection]
type mcnResourceDetailsSectionJSON struct {
	HiddenItems  apijson.Field
	Name         apijson.Field
	VisibleItems apijson.Field
	HelpText     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *McnResourceDetailsSection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsSectionJSON) RawJSON() string {
	return r.raw
}

type McnResourceDetailsSectionItem struct {
	HelpText string                             `json:"helpText"`
	Name     string                             `json:"name"`
	Value    McnResourceDetailsSectionItemValue `json:"value"`
	JSON     mcnResourceDetailsSectionItemJSON  `json:"-"`
}

// mcnResourceDetailsSectionItemJSON contains the JSON metadata for the struct
// [McnResourceDetailsSectionItem]
type mcnResourceDetailsSectionItemJSON struct {
	HelpText    apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnResourceDetailsSectionItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsSectionItemJSON) RawJSON() string {
	return r.raw
}

type McnResourceDetailsSectionItemValue struct {
	ItemType string `json:"item_type,required"`
	// This field can have the runtime type of
	// [[]McnResourceDetailsSectionItemValueMcnListItemList].
	List            interface{}                            `json:"list"`
	ResourcePreview McnResourcePreview                     `json:"resource_preview"`
	String          string                                 `json:"string"`
	Yaml            string                                 `json:"yaml"`
	YamlDiff        McnYamlDiff                            `json:"yaml_diff"`
	JSON            mcnResourceDetailsSectionItemValueJSON `json:"-"`
	union           McnResourceDetailsSectionItemValueUnion
}

// mcnResourceDetailsSectionItemValueJSON contains the JSON metadata for the struct
// [McnResourceDetailsSectionItemValue]
type mcnResourceDetailsSectionItemValueJSON struct {
	ItemType        apijson.Field
	List            apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	Yaml            apijson.Field
	YamlDiff        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r mcnResourceDetailsSectionItemValueJSON) RawJSON() string {
	return r.raw
}

func (r *McnResourceDetailsSectionItemValue) UnmarshalJSON(data []byte) (err error) {
	*r = McnResourceDetailsSectionItemValue{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [McnResourceDetailsSectionItemValueUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [McnStringItem],
// [McnResourceDetailsSectionItemValueMcnYamlItem],
// [McnResourceDetailsSectionItemValueMcnYamlDiffItem], [McnResourcePreviewItem],
// [McnResourceDetailsSectionItemValueMcnListItem].
func (r McnResourceDetailsSectionItemValue) AsUnion() McnResourceDetailsSectionItemValueUnion {
	return r.union
}

// Union satisfied by [McnStringItem],
// [McnResourceDetailsSectionItemValueMcnYamlItem],
// [McnResourceDetailsSectionItemValueMcnYamlDiffItem], [McnResourcePreviewItem] or
// [McnResourceDetailsSectionItemValueMcnListItem].
type McnResourceDetailsSectionItemValueUnion interface {
	implementsMcnResourceDetailsSectionItemValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McnResourceDetailsSectionItemValueUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnResourceDetailsSectionItemValueMcnYamlItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnResourceDetailsSectionItemValueMcnYamlDiffItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnResourcePreviewItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnResourceDetailsSectionItemValueMcnListItem{}),
		},
	)
}

type McnResourceDetailsSectionItemValueMcnYamlItem struct {
	ItemType string                                            `json:"item_type,required"`
	Yaml     string                                            `json:"yaml,required"`
	JSON     mcnResourceDetailsSectionItemValueMcnYamlItemJSON `json:"-"`
}

// mcnResourceDetailsSectionItemValueMcnYamlItemJSON contains the JSON metadata for
// the struct [McnResourceDetailsSectionItemValueMcnYamlItem]
type mcnResourceDetailsSectionItemValueMcnYamlItemJSON struct {
	ItemType    apijson.Field
	Yaml        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnResourceDetailsSectionItemValueMcnYamlItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsSectionItemValueMcnYamlItemJSON) RawJSON() string {
	return r.raw
}

func (r McnResourceDetailsSectionItemValueMcnYamlItem) implementsMcnResourceDetailsSectionItemValue() {
}

type McnResourceDetailsSectionItemValueMcnYamlDiffItem struct {
	ItemType string                                                `json:"item_type,required"`
	YamlDiff McnYamlDiff                                           `json:"yaml_diff,required"`
	JSON     mcnResourceDetailsSectionItemValueMcnYamlDiffItemJSON `json:"-"`
}

// mcnResourceDetailsSectionItemValueMcnYamlDiffItemJSON contains the JSON metadata
// for the struct [McnResourceDetailsSectionItemValueMcnYamlDiffItem]
type mcnResourceDetailsSectionItemValueMcnYamlDiffItemJSON struct {
	ItemType    apijson.Field
	YamlDiff    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnResourceDetailsSectionItemValueMcnYamlDiffItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsSectionItemValueMcnYamlDiffItemJSON) RawJSON() string {
	return r.raw
}

func (r McnResourceDetailsSectionItemValueMcnYamlDiffItem) implementsMcnResourceDetailsSectionItemValue() {
}

type McnResourceDetailsSectionItemValueMcnListItem struct {
	ItemType string                                              `json:"item_type,required"`
	List     []McnResourceDetailsSectionItemValueMcnListItemList `json:"list,required"`
	JSON     mcnResourceDetailsSectionItemValueMcnListItemJSON   `json:"-"`
}

// mcnResourceDetailsSectionItemValueMcnListItemJSON contains the JSON metadata for
// the struct [McnResourceDetailsSectionItemValueMcnListItem]
type mcnResourceDetailsSectionItemValueMcnListItemJSON struct {
	ItemType    apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnResourceDetailsSectionItemValueMcnListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourceDetailsSectionItemValueMcnListItemJSON) RawJSON() string {
	return r.raw
}

func (r McnResourceDetailsSectionItemValueMcnListItem) implementsMcnResourceDetailsSectionItemValue() {
}

type McnResourceDetailsSectionItemValueMcnListItemList struct {
	ItemType        string                                                `json:"item_type,required"`
	ResourcePreview McnResourcePreview                                    `json:"resource_preview"`
	String          string                                                `json:"string"`
	JSON            mcnResourceDetailsSectionItemValueMcnListItemListJSON `json:"-"`
	union           McnResourceDetailsSectionItemValueMcnListItemListUnion
}

// mcnResourceDetailsSectionItemValueMcnListItemListJSON contains the JSON metadata
// for the struct [McnResourceDetailsSectionItemValueMcnListItemList]
type mcnResourceDetailsSectionItemValueMcnListItemListJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	String          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r mcnResourceDetailsSectionItemValueMcnListItemListJSON) RawJSON() string {
	return r.raw
}

func (r *McnResourceDetailsSectionItemValueMcnListItemList) UnmarshalJSON(data []byte) (err error) {
	*r = McnResourceDetailsSectionItemValueMcnListItemList{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [McnResourceDetailsSectionItemValueMcnListItemListUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [McnStringItem],
// [McnResourcePreviewItem].
func (r McnResourceDetailsSectionItemValueMcnListItemList) AsUnion() McnResourceDetailsSectionItemValueMcnListItemListUnion {
	return r.union
}

// Union satisfied by [McnStringItem] or [McnResourcePreviewItem].
type McnResourceDetailsSectionItemValueMcnListItemListUnion interface {
	implementsMcnResourceDetailsSectionItemValueMcnListItemList()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*McnResourceDetailsSectionItemValueMcnListItemListUnion)(nil)).Elem(),
		"item_type",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnStringItem{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(McnResourcePreviewItem{}),
		},
	)
}

type McnResourcePreviewItem struct {
	ItemType        string                     `json:"item_type,required"`
	ResourcePreview McnResourcePreview         `json:"resource_preview,required"`
	JSON            mcnResourcePreviewItemJSON `json:"-"`
}

// mcnResourcePreviewItemJSON contains the JSON metadata for the struct
// [McnResourcePreviewItem]
type mcnResourcePreviewItemJSON struct {
	ItemType        apijson.Field
	ResourcePreview apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *McnResourcePreviewItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnResourcePreviewItemJSON) RawJSON() string {
	return r.raw
}

func (r McnResourcePreviewItem) implementsMcnResourceDetailsSectionItemValue() {}

func (r McnResourcePreviewItem) implementsMcnResourceDetailsSectionItemValueMcnListItemList() {}

type McnResourceType string

const (
	McnResourceTypeAwsCustomerGateway                                         McnResourceType = "aws_customer_gateway"
	McnResourceTypeAwsEgressOnlyInternetGateway                               McnResourceType = "aws_egress_only_internet_gateway"
	McnResourceTypeAwsInternetGateway                                         McnResourceType = "aws_internet_gateway"
	McnResourceTypeAwsInstance                                                McnResourceType = "aws_instance"
	McnResourceTypeAwsNetworkInterface                                        McnResourceType = "aws_network_interface"
	McnResourceTypeAwsRoute                                                   McnResourceType = "aws_route"
	McnResourceTypeAwsRouteTable                                              McnResourceType = "aws_route_table"
	McnResourceTypeAwsRouteTableAssociation                                   McnResourceType = "aws_route_table_association"
	McnResourceTypeAwsSubnet                                                  McnResourceType = "aws_subnet"
	McnResourceTypeAwsVpc                                                     McnResourceType = "aws_vpc"
	McnResourceTypeAwsVpcIpv4CidrBlockAssociation                             McnResourceType = "aws_vpc_ipv4_cidr_block_association"
	McnResourceTypeAwsVpnConnection                                           McnResourceType = "aws_vpn_connection"
	McnResourceTypeAwsVpnConnectionRoute                                      McnResourceType = "aws_vpn_connection_route"
	McnResourceTypeAwsVpnGateway                                              McnResourceType = "aws_vpn_gateway"
	McnResourceTypeAwsSecurityGroup                                           McnResourceType = "aws_security_group"
	McnResourceTypeAwsVpcSecurityGroupIngressRule                             McnResourceType = "aws_vpc_security_group_ingress_rule"
	McnResourceTypeAwsVpcSecurityGroupEgressRule                              McnResourceType = "aws_vpc_security_group_egress_rule"
	McnResourceTypeAwsEc2ManagedPrefixList                                    McnResourceType = "aws_ec2_managed_prefix_list"
	McnResourceTypeAwsEc2TransitGateway                                       McnResourceType = "aws_ec2_transit_gateway"
	McnResourceTypeAwsEc2TransitGatewayPrefixListReference                    McnResourceType = "aws_ec2_transit_gateway_prefix_list_reference"
	McnResourceTypeAwsEc2TransitGatewayVpcAttachment                          McnResourceType = "aws_ec2_transit_gateway_vpc_attachment"
	McnResourceTypeAzurermApplicationSecurityGroup                            McnResourceType = "azurerm_application_security_group"
	McnResourceTypeAzurermLb                                                  McnResourceType = "azurerm_lb"
	McnResourceTypeAzurermLbBackendAddressPool                                McnResourceType = "azurerm_lb_backend_address_pool"
	McnResourceTypeAzurermLbNatPool                                           McnResourceType = "azurerm_lb_nat_pool"
	McnResourceTypeAzurermLbNatRule                                           McnResourceType = "azurerm_lb_nat_rule"
	McnResourceTypeAzurermLbRule                                              McnResourceType = "azurerm_lb_rule"
	McnResourceTypeAzurermLocalNetworkGateway                                 McnResourceType = "azurerm_local_network_gateway"
	McnResourceTypeAzurermNetworkInterface                                    McnResourceType = "azurerm_network_interface"
	McnResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation McnResourceType = "azurerm_network_interface_application_security_group_association"
	McnResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation       McnResourceType = "azurerm_network_interface_backend_address_pool_association"
	McnResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation            McnResourceType = "azurerm_network_interface_security_group_association"
	McnResourceTypeAzurermNetworkSecurityGroup                                McnResourceType = "azurerm_network_security_group"
	McnResourceTypeAzurermPublicIP                                            McnResourceType = "azurerm_public_ip"
	McnResourceTypeAzurermRoute                                               McnResourceType = "azurerm_route"
	McnResourceTypeAzurermRouteTable                                          McnResourceType = "azurerm_route_table"
	McnResourceTypeAzurermSubnet                                              McnResourceType = "azurerm_subnet"
	McnResourceTypeAzurermSubnetRouteTableAssociation                         McnResourceType = "azurerm_subnet_route_table_association"
	McnResourceTypeAzurermVirtualMachine                                      McnResourceType = "azurerm_virtual_machine"
	McnResourceTypeAzurermVirtualNetworkGatewayConnection                     McnResourceType = "azurerm_virtual_network_gateway_connection"
	McnResourceTypeAzurermVirtualNetwork                                      McnResourceType = "azurerm_virtual_network"
	McnResourceTypeAzurermVirtualNetworkGateway                               McnResourceType = "azurerm_virtual_network_gateway"
	McnResourceTypeGoogleComputeNetwork                                       McnResourceType = "google_compute_network"
	McnResourceTypeGoogleComputeSubnetwork                                    McnResourceType = "google_compute_subnetwork"
	McnResourceTypeGoogleComputeVpnGateway                                    McnResourceType = "google_compute_vpn_gateway"
	McnResourceTypeGoogleComputeVpnTunnel                                     McnResourceType = "google_compute_vpn_tunnel"
	McnResourceTypeGoogleComputeRoute                                         McnResourceType = "google_compute_route"
	McnResourceTypeGoogleComputeAddress                                       McnResourceType = "google_compute_address"
	McnResourceTypeGoogleComputeGlobalAddress                                 McnResourceType = "google_compute_global_address"
	McnResourceTypeGoogleComputeRouter                                        McnResourceType = "google_compute_router"
	McnResourceTypeGoogleComputeInterconnectAttachment                        McnResourceType = "google_compute_interconnect_attachment"
	McnResourceTypeGoogleComputeHaVpnGateway                                  McnResourceType = "google_compute_ha_vpn_gateway"
	McnResourceTypeGoogleComputeForwardingRule                                McnResourceType = "google_compute_forwarding_rule"
	McnResourceTypeGoogleComputeNetworkFirewallPolicy                         McnResourceType = "google_compute_network_firewall_policy"
	McnResourceTypeGoogleComputeNetworkFirewallPolicyRule                     McnResourceType = "google_compute_network_firewall_policy_rule"
	McnResourceTypeCloudflareStaticRoute                                      McnResourceType = "cloudflare_static_route"
	McnResourceTypeCloudflareIpsecTunnel                                      McnResourceType = "cloudflare_ipsec_tunnel"
)

func (r McnResourceType) IsKnown() bool {
	switch r {
	case McnResourceTypeAwsCustomerGateway, McnResourceTypeAwsEgressOnlyInternetGateway, McnResourceTypeAwsInternetGateway, McnResourceTypeAwsInstance, McnResourceTypeAwsNetworkInterface, McnResourceTypeAwsRoute, McnResourceTypeAwsRouteTable, McnResourceTypeAwsRouteTableAssociation, McnResourceTypeAwsSubnet, McnResourceTypeAwsVpc, McnResourceTypeAwsVpcIpv4CidrBlockAssociation, McnResourceTypeAwsVpnConnection, McnResourceTypeAwsVpnConnectionRoute, McnResourceTypeAwsVpnGateway, McnResourceTypeAwsSecurityGroup, McnResourceTypeAwsVpcSecurityGroupIngressRule, McnResourceTypeAwsVpcSecurityGroupEgressRule, McnResourceTypeAwsEc2ManagedPrefixList, McnResourceTypeAwsEc2TransitGateway, McnResourceTypeAwsEc2TransitGatewayPrefixListReference, McnResourceTypeAwsEc2TransitGatewayVpcAttachment, McnResourceTypeAzurermApplicationSecurityGroup, McnResourceTypeAzurermLb, McnResourceTypeAzurermLbBackendAddressPool, McnResourceTypeAzurermLbNatPool, McnResourceTypeAzurermLbNatRule, McnResourceTypeAzurermLbRule, McnResourceTypeAzurermLocalNetworkGateway, McnResourceTypeAzurermNetworkInterface, McnResourceTypeAzurermNetworkInterfaceApplicationSecurityGroupAssociation, McnResourceTypeAzurermNetworkInterfaceBackendAddressPoolAssociation, McnResourceTypeAzurermNetworkInterfaceSecurityGroupAssociation, McnResourceTypeAzurermNetworkSecurityGroup, McnResourceTypeAzurermPublicIP, McnResourceTypeAzurermRoute, McnResourceTypeAzurermRouteTable, McnResourceTypeAzurermSubnet, McnResourceTypeAzurermSubnetRouteTableAssociation, McnResourceTypeAzurermVirtualMachine, McnResourceTypeAzurermVirtualNetworkGatewayConnection, McnResourceTypeAzurermVirtualNetwork, McnResourceTypeAzurermVirtualNetworkGateway, McnResourceTypeGoogleComputeNetwork, McnResourceTypeGoogleComputeSubnetwork, McnResourceTypeGoogleComputeVpnGateway, McnResourceTypeGoogleComputeVpnTunnel, McnResourceTypeGoogleComputeRoute, McnResourceTypeGoogleComputeAddress, McnResourceTypeGoogleComputeGlobalAddress, McnResourceTypeGoogleComputeRouter, McnResourceTypeGoogleComputeInterconnectAttachment, McnResourceTypeGoogleComputeHaVpnGateway, McnResourceTypeGoogleComputeForwardingRule, McnResourceTypeGoogleComputeNetworkFirewallPolicy, McnResourceTypeGoogleComputeNetworkFirewallPolicyRule, McnResourceTypeCloudflareStaticRoute, McnResourceTypeCloudflareIpsecTunnel:
		return true
	}
	return false
}

type McnStringItem struct {
	ItemType string            `json:"item_type,required"`
	String   string            `json:"string,required"`
	JSON     mcnStringItemJSON `json:"-"`
}

// mcnStringItemJSON contains the JSON metadata for the struct [McnStringItem]
type mcnStringItemJSON struct {
	ItemType    apijson.Field
	String      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *McnStringItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mcnStringItemJSON) RawJSON() string {
	return r.raw
}

func (r McnStringItem) implementsMcnResourceDetailsSectionItemValue() {}

func (r McnStringItem) implementsMcnResourceDetailsSectionItemValueMcnListItemList() {}

type AccountMagicCloudResourceGetResponse struct {
	Result McnResourceDetails                       `json:"result"`
	JSON   accountMagicCloudResourceGetResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudResourceGetResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudResourceGetResponse]
type accountMagicCloudResourceGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudResourceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudResourceGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudResourceListResponse struct {
	Errors     []McnError                                      `json:"errors,required"`
	Messages   []McnError                                      `json:"messages,required"`
	Result     []McnResourceDetails                            `json:"result,required"`
	Success    bool                                            `json:"success,required"`
	ResultInfo AccountMagicCloudResourceListResponseResultInfo `json:"result_info"`
	JSON       accountMagicCloudResourceListResponseJSON       `json:"-"`
}

// accountMagicCloudResourceListResponseJSON contains the JSON metadata for the
// struct [AccountMagicCloudResourceListResponse]
type accountMagicCloudResourceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudResourceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudResourceListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudResourceListResponseResultInfo struct {
	// The number of items in the current result set.
	Count int64 `json:"count,required"`
	// The current page (starts from zero).
	Page int64 `json:"page,required"`
	// Max items per page.
	PerPage int64 `json:"per_page,required"`
	// The total number of items in the entire result set.
	TotalCount int64 `json:"total_count,required"`
	// The number of total pages in the entire result set.
	TotalPages int64                                               `json:"total_pages"`
	JSON       accountMagicCloudResourceListResponseResultInfoJSON `json:"-"`
}

// accountMagicCloudResourceListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountMagicCloudResourceListResponseResultInfo]
type accountMagicCloudResourceListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudResourceListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudResourceListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudResourcePreviewPolicyResponse struct {
	Result string                                             `json:"result"`
	JSON   accountMagicCloudResourcePreviewPolicyResponseJSON `json:"-"`
	McnGoodResponse
}

// accountMagicCloudResourcePreviewPolicyResponseJSON contains the JSON metadata
// for the struct [AccountMagicCloudResourcePreviewPolicyResponse]
type accountMagicCloudResourcePreviewPolicyResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicCloudResourcePreviewPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicCloudResourcePreviewPolicyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicCloudResourceGetParams struct {
	V2 param.Field[bool] `query:"v2"`
}

// URLQuery serializes [AccountMagicCloudResourceGetParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudResourceGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudResourceListParams struct {
	Cloudflare param.Field[bool] `query:"cloudflare"`
	Desc       param.Field[bool] `query:"desc"`
	Managed    param.Field[bool] `query:"managed"`
	// one of ["id", "resource_type", "region"]
	OrderBy       param.Field[string]            `query:"order_by"`
	Page          param.Field[int64]             `query:"page"`
	PerPage       param.Field[int64]             `query:"per_page"`
	ProviderID    param.Field[string]            `query:"provider_id"`
	Region        param.Field[string]            `query:"region"`
	ResourceGroup param.Field[string]            `query:"resource_group"`
	ResourceID    param.Field[[]string]          `query:"resource_id" format:"uuid"`
	ResourceType  param.Field[[]McnResourceType] `query:"resource_type"`
	Search        param.Field[[]string]          `query:"search"`
	V2            param.Field[bool]              `query:"v2"`
}

// URLQuery serializes [AccountMagicCloudResourceListParams]'s query parameters as
// `url.Values`.
func (r AccountMagicCloudResourceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudResourceExportParams struct {
	Desc param.Field[bool] `query:"desc"`
	// one of ["id", "resource_type", "region"]
	OrderBy       param.Field[string]            `query:"order_by"`
	ProviderID    param.Field[string]            `query:"provider_id"`
	Region        param.Field[string]            `query:"region"`
	ResourceGroup param.Field[string]            `query:"resource_group"`
	ResourceID    param.Field[[]string]          `query:"resource_id" format:"uuid"`
	ResourceType  param.Field[[]McnResourceType] `query:"resource_type"`
	Search        param.Field[[]string]          `query:"search"`
	V2            param.Field[bool]              `query:"v2"`
}

// URLQuery serializes [AccountMagicCloudResourceExportParams]'s query parameters
// as `url.Values`.
func (r AccountMagicCloudResourceExportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountMagicCloudResourcePreviewPolicyParams struct {
	Policy param.Field[string] `json:"policy,required"`
}

func (r AccountMagicCloudResourcePreviewPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
