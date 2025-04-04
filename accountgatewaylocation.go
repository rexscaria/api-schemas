// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountGatewayLocationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayLocationService] method instead.
type AccountGatewayLocationService struct {
	Options []option.RequestOption
}

// NewAccountGatewayLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayLocationService(opts ...option.RequestOption) (r *AccountGatewayLocationService) {
	r = &AccountGatewayLocationService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway location.
func (r *AccountGatewayLocationService) New(ctx context.Context, accountID string, body AccountGatewayLocationNewParams, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Get(ctx context.Context, accountID string, locationID string, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Update(ctx context.Context, accountID string, locationID string, body AccountGatewayLocationUpdateParams, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Zero Trust Gateway locations for an account.
func (r *AccountGatewayLocationService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a configured Zero Trust Gateway location.
func (r *AccountGatewayLocationService) Delete(ctx context.Context, accountID string, locationID string, body AccountGatewayLocationDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if locationID == "" {
		err = errors.New("missing required location_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/locations/%s", accountID, locationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// The destination endpoints configured for this location. When updating a
// location, if this field is absent or set with null, the endpoints configuration
// remains unchanged.
type Endpoints struct {
	Doh  EndpointsDoh  `json:"doh"`
	Dot  EndpointsDot  `json:"dot"`
	Ipv4 EndpointsIpv4 `json:"ipv4"`
	Ipv6 EndpointsIpv6 `json:"ipv6"`
	JSON endpointsJSON `json:"-"`
}

// endpointsJSON contains the JSON metadata for the struct [Endpoints]
type endpointsJSON struct {
	Doh         apijson.Field
	Dot         apijson.Field
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Endpoints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsJSON) RawJSON() string {
	return r.raw
}

type EndpointsDoh struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []IPNetwork `json:"networks"`
	// True if the endpoint requires
	// [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user)
	// authentication.
	RequireToken bool             `json:"require_token"`
	JSON         endpointsDohJSON `json:"-"`
}

// endpointsDohJSON contains the JSON metadata for the struct [EndpointsDoh]
type endpointsDohJSON struct {
	Enabled      apijson.Field
	Networks     apijson.Field
	RequireToken apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EndpointsDoh) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsDohJSON) RawJSON() string {
	return r.raw
}

type EndpointsDot struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []IPNetwork      `json:"networks"`
	JSON     endpointsDotJSON `json:"-"`
}

// endpointsDotJSON contains the JSON metadata for the struct [EndpointsDot]
type endpointsDotJSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointsDot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsDotJSON) RawJSON() string {
	return r.raw
}

type EndpointsIpv4 struct {
	// True if the endpoint is enabled for this location.
	Enabled bool              `json:"enabled"`
	JSON    endpointsIpv4JSON `json:"-"`
}

// endpointsIpv4JSON contains the JSON metadata for the struct [EndpointsIpv4]
type endpointsIpv4JSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointsIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsIpv4JSON) RawJSON() string {
	return r.raw
}

type EndpointsIpv6 struct {
	// True if the endpoint is enabled for this location.
	Enabled bool `json:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks []EndpointsIpv6Network `json:"networks"`
	JSON     endpointsIpv6JSON      `json:"-"`
}

// endpointsIpv6JSON contains the JSON metadata for the struct [EndpointsIpv6]
type endpointsIpv6JSON struct {
	Enabled     apijson.Field
	Networks    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointsIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsIpv6JSON) RawJSON() string {
	return r.raw
}

type EndpointsIpv6Network struct {
	// The IPv6 address or IPv6 CIDR.
	Network string                   `json:"network,required"`
	JSON    endpointsIpv6NetworkJSON `json:"-"`
}

// endpointsIpv6NetworkJSON contains the JSON metadata for the struct
// [EndpointsIpv6Network]
type endpointsIpv6NetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointsIpv6Network) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointsIpv6NetworkJSON) RawJSON() string {
	return r.raw
}

// The destination endpoints configured for this location. When updating a
// location, if this field is absent or set with null, the endpoints configuration
// remains unchanged.
type EndpointsParam struct {
	Doh  param.Field[EndpointsDohParam]  `json:"doh"`
	Dot  param.Field[EndpointsDotParam]  `json:"dot"`
	Ipv4 param.Field[EndpointsIpv4Param] `json:"ipv4"`
	Ipv6 param.Field[EndpointsIpv6Param] `json:"ipv6"`
}

func (r EndpointsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EndpointsDohParam struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]IPNetworkParam] `json:"networks"`
	// True if the endpoint requires
	// [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user)
	// authentication.
	RequireToken param.Field[bool] `json:"require_token"`
}

func (r EndpointsDohParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EndpointsDotParam struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IP network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]IPNetworkParam] `json:"networks"`
}

func (r EndpointsDotParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EndpointsIpv4Param struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r EndpointsIpv4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EndpointsIpv6Param struct {
	// True if the endpoint is enabled for this location.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of allowed source IPv6 network ranges for this endpoint. When empty, all
	// source IPs are allowed. A non-empty list is only effective if the endpoint is
	// enabled for this location.
	Networks param.Field[[]EndpointsIpv6NetworkParam] `json:"networks"`
}

func (r EndpointsIpv6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EndpointsIpv6NetworkParam struct {
	// The IPv6 address or IPv6 CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r EndpointsIpv6NetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPNetwork struct {
	// The IP address or IP CIDR.
	Network string        `json:"network,required"`
	JSON    ipNetworkJSON `json:"-"`
}

// ipNetworkJSON contains the JSON metadata for the struct [IPNetwork]
type ipNetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipNetworkJSON) RawJSON() string {
	return r.raw
}

type IPNetworkParam struct {
	// The IP address or IP CIDR.
	Network param.Field[string] `json:"network,required"`
}

func (r IPNetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Ipv4Network struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network string          `json:"network,required"`
	JSON    ipv4NetworkJSON `json:"-"`
}

// ipv4NetworkJSON contains the JSON metadata for the struct [Ipv4Network]
type ipv4NetworkJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ipv4Network) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv4NetworkJSON) RawJSON() string {
	return r.raw
}

type Ipv4NetworkParam struct {
	// The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.
	Network param.Field[string] `json:"network,required"`
}

func (r Ipv4NetworkParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Location struct {
	ID string `json:"id"`
	// True if the location is the default location.
	ClientDefault bool      `json:"client_default"`
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	// The identifier of the pair of IPv4 addresses assigned to this location.
	DNSDestinationIPsID string `json:"dns_destination_ips_id"`
	// The uuid identifier of the IPv6 block brought to the gateway, so that this
	// location's IPv6 address is allocated from the Bring Your Own Ipv6(BYOIPv6) block
	// and not from the standard CloudFlare IPv6 block.
	DNSDestinationIpv6BlockID string `json:"dns_destination_ipv6_block_id"`
	// The DNS over HTTPS domain to send DNS requests to. This field is auto-generated
	// by Gateway.
	DohSubdomain string `json:"doh_subdomain"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport bool `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints Endpoints `json:"endpoints"`
	// IPV6 destination ip assigned to this location. DNS requests sent to this IP will
	// counted as the request under this location. This field is auto-generated by
	// Gateway.
	IP string `json:"ip"`
	// The primary destination IPv4 address from the pair identified by the
	// dns_destination_ips_id. This field is read-only.
	Ipv4Destination string `json:"ipv4_destination"`
	// The backup destination IPv4 address from the pair identified by the
	// dns_destination_ips_id. This field is read-only.
	Ipv4DestinationBackup string `json:"ipv4_destination_backup"`
	// The name of the location.
	Name string `json:"name"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks  []Ipv4Network `json:"networks"`
	UpdatedAt time.Time     `json:"updated_at" format:"date-time"`
	JSON      locationJSON  `json:"-"`
}

// locationJSON contains the JSON metadata for the struct [Location]
type locationJSON struct {
	ID                        apijson.Field
	ClientDefault             apijson.Field
	CreatedAt                 apijson.Field
	DNSDestinationIPsID       apijson.Field
	DNSDestinationIpv6BlockID apijson.Field
	DohSubdomain              apijson.Field
	EcsSupport                apijson.Field
	Endpoints                 apijson.Field
	IP                        apijson.Field
	Ipv4Destination           apijson.Field
	Ipv4DestinationBackup     apijson.Field
	Name                      apijson.Field
	Networks                  apijson.Field
	UpdatedAt                 apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *Location) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationJSON) RawJSON() string {
	return r.raw
}

type SingleResponseLocation struct {
	Result Location                   `json:"result"`
	JSON   singleResponseLocationJSON `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// singleResponseLocationJSON contains the JSON metadata for the struct
// [SingleResponseLocation]
type singleResponseLocationJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseLocationJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayLocationListResponse struct {
	Result []Location                             `json:"result"`
	JSON   accountGatewayLocationListResponseJSON `json:"-"`
	APIResponseCollectionZeroTrustGateway
}

// accountGatewayLocationListResponseJSON contains the JSON metadata for the struct
// [AccountGatewayLocationListResponse]
type accountGatewayLocationListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayLocationNewParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// The identifier of the pair of IPv4 addresses assigned to this location. When
	// creating a location, if this field is absent or set with null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if the field is absent or set with null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints param.Field[EndpointsParam] `json:"endpoints"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks param.Field[[]Ipv4NetworkParam] `json:"networks"`
}

func (r AccountGatewayLocationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationUpdateParams struct {
	// The name of the location.
	Name param.Field[string] `json:"name,required"`
	// True if the location is the default location.
	ClientDefault param.Field[bool] `json:"client_default"`
	// The identifier of the pair of IPv4 addresses assigned to this location. When
	// creating a location, if this field is absent or set with null, the pair of
	// shared IPv4 addresses (0e4a32c6-6fb8-4858-9296-98f51631e8e6) is auto-assigned.
	// When updating a location, if the field is absent or set with null, the
	// pre-assigned pair remains unchanged.
	DNSDestinationIPsID param.Field[string] `json:"dns_destination_ips_id"`
	// True if the location needs to resolve EDNS queries.
	EcsSupport param.Field[bool] `json:"ecs_support"`
	// The destination endpoints configured for this location. When updating a
	// location, if this field is absent or set with null, the endpoints configuration
	// remains unchanged.
	Endpoints param.Field[EndpointsParam] `json:"endpoints"`
	// A list of network ranges that requests from this location would originate from.
	// A non-empty list is only effective if the ipv4 endpoint is enabled for this
	// location.
	Networks param.Field[[]Ipv4NetworkParam] `json:"networks"`
}

func (r AccountGatewayLocationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayLocationDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGatewayLocationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
