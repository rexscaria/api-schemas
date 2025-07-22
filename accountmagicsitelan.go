// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMagicSiteLanService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicSiteLanService] method instead.
type AccountMagicSiteLanService struct {
	Options []option.RequestOption
}

// NewAccountMagicSiteLanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicSiteLanService(opts ...option.RequestOption) (r *AccountMagicSiteLanService) {
	r = &AccountMagicSiteLanService{}
	r.Options = opts
	return
}

// Creates a new Site LAN. If the site is in high availability mode,
// static_addressing is required along with secondary and virtual address.
func (r *AccountMagicSiteLanService) New(ctx context.Context, accountID string, siteID string, body AccountMagicSiteLanNewParams, opts ...option.RequestOption) (res *MagicLansCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific Site LAN.
func (r *AccountMagicSiteLanService) Get(ctx context.Context, accountID string, siteID string, lanID string, opts ...option.RequestOption) (res *AccountMagicSiteLanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Site LAN.
func (r *AccountMagicSiteLanService) Update(ctx context.Context, accountID string, siteID string, lanID string, body AccountMagicSiteLanUpdateParams, opts ...option.RequestOption) (res *MagicLanModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Site LANs associated with an account.
func (r *AccountMagicSiteLanService) List(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *MagicLansCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a specific Site LAN.
func (r *AccountMagicSiteLanService) Delete(ctx context.Context, accountID string, siteID string, lanID string, body AccountMagicSiteLanDeleteParams, opts ...option.RequestOption) (res *AccountMagicSiteLanDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Patch a specific Site LAN.
func (r *AccountMagicSiteLanService) Patch(ctx context.Context, accountID string, siteID string, lanID string, body AccountMagicSiteLanPatchParams, opts ...option.RequestOption) (res *MagicLanModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if lanID == "" {
		err = errors.New("missing required lan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/lans/%s", accountID, siteID, lanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MagicLan struct {
	// Identifier
	ID string `json:"id"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        bool                `json:"ha_link"`
	Name          string              `json:"name"`
	Nat           MagicNat            `json:"nat"`
	Physport      int64               `json:"physport"`
	RoutedSubnets []MagicRoutedSubnet `json:"routed_subnets"`
	// Identifier
	SiteID string `json:"site_id"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing MagicLanStaticAddressing `json:"static_addressing"`
	// VLAN port number.
	VlanTag int64        `json:"vlan_tag"`
	JSON    magicLanJSON `json:"-"`
}

// magicLanJSON contains the JSON metadata for the struct [MagicLan]
type magicLanJSON struct {
	ID               apijson.Field
	HaLink           apijson.Field
	Name             apijson.Field
	Nat              apijson.Field
	Physport         apijson.Field
	RoutedSubnets    apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicLan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanJSON) RawJSON() string {
	return r.raw
}

type MagicLanModifiedResponse struct {
	Result MagicLan                     `json:"result"`
	JSON   magicLanModifiedResponseJSON `json:"-"`
	MagicAPIResponseSingle
}

// magicLanModifiedResponseJSON contains the JSON metadata for the struct
// [MagicLanModifiedResponse]
type magicLanModifiedResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicLanModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanModifiedResponseJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type MagicLanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address    string                             `json:"address,required"`
	DhcpRelay  MagicLanStaticAddressingDhcpRelay  `json:"dhcp_relay"`
	DhcpServer MagicLanStaticAddressingDhcpServer `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress string                       `json:"virtual_address"`
	JSON           magicLanStaticAddressingJSON `json:"-"`
}

// magicLanStaticAddressingJSON contains the JSON metadata for the struct
// [MagicLanStaticAddressing]
type magicLanStaticAddressingJSON struct {
	Address          apijson.Field
	DhcpRelay        apijson.Field
	DhcpServer       apijson.Field
	SecondaryAddress apijson.Field
	VirtualAddress   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicLanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

type MagicLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	ServerAddresses []string                              `json:"server_addresses"`
	JSON            magicLanStaticAddressingDhcpRelayJSON `json:"-"`
}

// magicLanStaticAddressingDhcpRelayJSON contains the JSON metadata for the struct
// [MagicLanStaticAddressingDhcpRelay]
type magicLanStaticAddressingDhcpRelayJSON struct {
	ServerAddresses apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *MagicLanStaticAddressingDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanStaticAddressingDhcpRelayJSON) RawJSON() string {
	return r.raw
}

type MagicLanStaticAddressingDhcpServer struct {
	// A valid IPv4 address.
	DhcpPoolEnd string `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart string `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer  string   `json:"dns_server"`
	DNSServers []string `json:"dns_servers"`
	// Mapping of MAC addresses to IP addresses
	Reservations map[string]string                      `json:"reservations"`
	JSON         magicLanStaticAddressingDhcpServerJSON `json:"-"`
}

// magicLanStaticAddressingDhcpServerJSON contains the JSON metadata for the struct
// [MagicLanStaticAddressingDhcpServer]
type magicLanStaticAddressingDhcpServerJSON struct {
	DhcpPoolEnd   apijson.Field
	DhcpPoolStart apijson.Field
	DNSServer     apijson.Field
	DNSServers    apijson.Field
	Reservations  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MagicLanStaticAddressingDhcpServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLanStaticAddressingDhcpServerJSON) RawJSON() string {
	return r.raw
}

// If the site is not configured in high availability mode, this configuration is
// optional (if omitted, use DHCP). However, if in high availability mode,
// static_address is required along with secondary and virtual address.
type MagicLanStaticAddressingParam struct {
	// A valid CIDR notation representing an IP range.
	Address    param.Field[string]                                  `json:"address,required"`
	DhcpRelay  param.Field[MagicLanStaticAddressingDhcpRelayParam]  `json:"dhcp_relay"`
	DhcpServer param.Field[MagicLanStaticAddressingDhcpServerParam] `json:"dhcp_server"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
	// A valid CIDR notation representing an IP range.
	VirtualAddress param.Field[string] `json:"virtual_address"`
}

func (r MagicLanStaticAddressingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicLanStaticAddressingDhcpRelayParam struct {
	// List of DHCP server IPs.
	ServerAddresses param.Field[[]string] `json:"server_addresses"`
}

func (r MagicLanStaticAddressingDhcpRelayParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicLanStaticAddressingDhcpServerParam struct {
	// A valid IPv4 address.
	DhcpPoolEnd param.Field[string] `json:"dhcp_pool_end"`
	// A valid IPv4 address.
	DhcpPoolStart param.Field[string] `json:"dhcp_pool_start"`
	// A valid IPv4 address.
	DNSServer  param.Field[string]   `json:"dns_server"`
	DNSServers param.Field[[]string] `json:"dns_servers"`
	// Mapping of MAC addresses to IP addresses
	Reservations param.Field[map[string]string] `json:"reservations"`
}

func (r MagicLanStaticAddressingDhcpServerParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicLanUpdateRequestParam struct {
	Name          param.Field[string]                   `json:"name"`
	Nat           param.Field[MagicNatParam]            `json:"nat"`
	Physport      param.Field[int64]                    `json:"physport"`
	RoutedSubnets param.Field[[]MagicRoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[MagicLanStaticAddressingParam] `json:"static_addressing"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r MagicLanUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicLansCollectionResponse struct {
	Result []MagicLan                      `json:"result"`
	JSON   magicLansCollectionResponseJSON `json:"-"`
	MagicAPIResponseSingle
}

// magicLansCollectionResponseJSON contains the JSON metadata for the struct
// [MagicLansCollectionResponse]
type magicLansCollectionResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicLansCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicLansCollectionResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNat struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix string       `json:"static_prefix"`
	JSON         magicNatJSON `json:"-"`
}

// magicNatJSON contains the JSON metadata for the struct [MagicNat]
type magicNatJSON struct {
	StaticPrefix apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicNat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNatJSON) RawJSON() string {
	return r.raw
}

type MagicNatParam struct {
	// A valid CIDR notation representing an IP range.
	StaticPrefix param.Field[string] `json:"static_prefix"`
}

func (r MagicNatParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicRoutedSubnet struct {
	// A valid IPv4 address.
	NextHop string `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix string                `json:"prefix,required"`
	Nat    MagicNat              `json:"nat"`
	JSON   magicRoutedSubnetJSON `json:"-"`
}

// magicRoutedSubnetJSON contains the JSON metadata for the struct
// [MagicRoutedSubnet]
type magicRoutedSubnetJSON struct {
	NextHop     apijson.Field
	Prefix      apijson.Field
	Nat         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicRoutedSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicRoutedSubnetJSON) RawJSON() string {
	return r.raw
}

type MagicRoutedSubnetParam struct {
	// A valid IPv4 address.
	NextHop param.Field[string] `json:"next_hop,required"`
	// A valid CIDR notation representing an IP range.
	Prefix param.Field[string]        `json:"prefix,required"`
	Nat    param.Field[MagicNatParam] `json:"nat"`
}

func (r MagicRoutedSubnetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteLanGetResponse struct {
	Result MagicLan                           `json:"result"`
	JSON   accountMagicSiteLanGetResponseJSON `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicSiteLanGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteLanGetResponse]
type accountMagicSiteLanGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteLanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteLanGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteLanDeleteResponse struct {
	Result MagicLan                              `json:"result"`
	JSON   accountMagicSiteLanDeleteResponseJSON `json:"-"`
	MagicAPIResponseSingle
}

// accountMagicSiteLanDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteLanDeleteResponse]
type accountMagicSiteLanDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteLanDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteLanDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteLanNewParams struct {
	Physport param.Field[int64] `json:"physport,required"`
	// VLAN port number.
	VlanTag param.Field[int64] `json:"vlan_tag,required"`
	// mark true to use this LAN for HA probing. only works for site with HA turned on.
	// only one LAN can be set as the ha_link.
	HaLink        param.Field[bool]                     `json:"ha_link"`
	Name          param.Field[string]                   `json:"name"`
	Nat           param.Field[MagicNatParam]            `json:"nat"`
	RoutedSubnets param.Field[[]MagicRoutedSubnetParam] `json:"routed_subnets"`
	// If the site is not configured in high availability mode, this configuration is
	// optional (if omitted, use DHCP). However, if in high availability mode,
	// static_address is required along with secondary and virtual address.
	StaticAddressing param.Field[MagicLanStaticAddressingParam] `json:"static_addressing"`
}

func (r AccountMagicSiteLanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteLanUpdateParams struct {
	MagicLanUpdateRequest MagicLanUpdateRequestParam `json:"magic_lan_update_request,required"`
}

func (r AccountMagicSiteLanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicLanUpdateRequest)
}

type AccountMagicSiteLanDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMagicSiteLanDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicSiteLanPatchParams struct {
	MagicLanUpdateRequest MagicLanUpdateRequestParam `json:"magic_lan_update_request,required"`
}

func (r AccountMagicSiteLanPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicLanUpdateRequest)
}
