// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAddressingAddressMapService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingAddressMapService] method instead.
type AccountAddressingAddressMapService struct {
	Options  []option.RequestOption
	Accounts *AccountAddressingAddressMapAccountService
	IPs      *AccountAddressingAddressMapIPService
	Zones    *AccountAddressingAddressMapZoneService
}

// NewAccountAddressingAddressMapService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingAddressMapService(opts ...option.RequestOption) (r *AccountAddressingAddressMapService) {
	r = &AccountAddressingAddressMapService{}
	r.Options = opts
	r.Accounts = NewAccountAddressingAddressMapAccountService(opts...)
	r.IPs = NewAccountAddressingAddressMapIPService(opts...)
	r.Zones = NewAccountAddressingAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AccountAddressingAddressMapService) New(ctx context.Context, accountID string, body AccountAddressingAddressMapNewParams, opts ...option.RequestOption) (res *FullResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Show a particular address map owned by the account.
func (r *AccountAddressingAddressMapService) Get(ctx context.Context, accountID string, addressMapID string, opts ...option.RequestOption) (res *FullResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify properties of an address map owned by the account.
func (r *AccountAddressingAddressMapService) Update(ctx context.Context, accountID string, addressMapID string, body AccountAddressingAddressMapUpdateParams, opts ...option.RequestOption) (res *AccountAddressingAddressMapUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all address maps owned by the account.
func (r *AccountAddressingAddressMapService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAddressingAddressMapListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a particular address map owned by the account. An Address Map must be
// disabled before it can be deleted.
func (r *AccountAddressingAddressMapService) Delete(ctx context.Context, accountID string, addressMapID string, body AccountAddressingAddressMapDeleteParams, opts ...option.RequestOption) (res *APIResponseCollectionAddressing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", accountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AddressMap struct {
	// Identifier of an Address Map.
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled    bool           `json:"enabled,nullable"`
	ModifiedAt time.Time      `json:"modified_at" format:"date-time"`
	JSON       addressMapJSON `json:"-"`
}

// addressMapJSON contains the JSON metadata for the struct [AddressMap]
type addressMapJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSni   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMap) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAddressing struct {
	ResultInfo APIResponseCollectionAddressingResultInfo `json:"result_info"`
	JSON       apiResponseCollectionAddressingJSON       `json:"-"`
	APIResponseAddressing
}

// apiResponseCollectionAddressingJSON contains the JSON metadata for the struct
// [APIResponseCollectionAddressing]
type apiResponseCollectionAddressingJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAddressingJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAddressingResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       apiResponseCollectionAddressingResultInfoJSON `json:"-"`
}

// apiResponseCollectionAddressingResultInfoJSON contains the JSON metadata for the
// struct [APIResponseCollectionAddressingResultInfo]
type apiResponseCollectionAddressingResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAddressingResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAddressingResultInfoJSON) RawJSON() string {
	return r.raw
}

type FullResponse struct {
	Result FullResponseResult `json:"result"`
	JSON   fullResponseJSON   `json:"-"`
	AddressingAPIResponseSingle
}

// fullResponseJSON contains the JSON metadata for the struct [FullResponse]
type fullResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FullResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fullResponseJSON) RawJSON() string {
	return r.raw
}

type FullResponseResult struct {
	// The set of IPs on the Address Map.
	IPs []FullResponseResultIP `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []Membership           `json:"memberships"`
	JSON        fullResponseResultJSON `json:"-"`
	AddressMap
}

// fullResponseResultJSON contains the JSON metadata for the struct
// [FullResponseResult]
type fullResponseResultJSON struct {
	IPs         apijson.Field
	Memberships apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FullResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fullResponseResultJSON) RawJSON() string {
	return r.raw
}

type FullResponseResultIP struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string                   `json:"ip"`
	JSON fullResponseResultIPJSON `json:"-"`
}

// fullResponseResultIPJSON contains the JSON metadata for the struct
// [FullResponseResultIP]
type fullResponseResultIPJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FullResponseResultIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fullResponseResultIPJSON) RawJSON() string {
	return r.raw
}

type Membership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The identifier for the membership (eg. a zone or account tag).
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind MembershipKind `json:"kind"`
	JSON membershipJSON `json:"-"`
}

// membershipJSON contains the JSON metadata for the struct [Membership]
type membershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Membership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type MembershipKind string

const (
	MembershipKindZone    MembershipKind = "zone"
	MembershipKindAccount MembershipKind = "account"
)

func (r MembershipKind) IsKnown() bool {
	switch r {
	case MembershipKindZone, MembershipKindAccount:
		return true
	}
	return false
}

type MembershipParam struct {
	CreatedAt param.Field[time.Time] `json:"created_at" format:"date-time"`
	// The identifier for the membership (eg. a zone or account tag).
	Identifier param.Field[string] `json:"identifier"`
	// The type of the membership.
	Kind param.Field[MembershipKind] `json:"kind"`
}

func (r MembershipParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingAddressMapUpdateResponse struct {
	Result AddressMap                                    `json:"result"`
	JSON   accountAddressingAddressMapUpdateResponseJSON `json:"-"`
	AddressingAPIResponseSingle
}

// accountAddressingAddressMapUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAddressingAddressMapUpdateResponse]
type accountAddressingAddressMapUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingAddressMapUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingAddressMapUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingAddressMapListResponse struct {
	Result []AddressMap                                `json:"result"`
	JSON   accountAddressingAddressMapListResponseJSON `json:"-"`
	APIResponseCollectionAddressing
}

// accountAddressingAddressMapListResponseJSON contains the JSON metadata for the
// struct [AccountAddressingAddressMapListResponse]
type accountAddressingAddressMapListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingAddressMapListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingAddressMapListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingAddressMapNewParams struct {
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool]     `json:"enabled"`
	IPs     param.Field[[]string] `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships param.Field[[]MembershipParam] `json:"memberships"`
}

func (r AccountAddressingAddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingAddressMapUpdateParams struct {
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSni param.Field[string] `json:"default_sni"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountAddressingAddressMapUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingAddressMapDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountAddressingAddressMapDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
