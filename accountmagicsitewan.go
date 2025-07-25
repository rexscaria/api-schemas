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

// AccountMagicSiteWanService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicSiteWanService] method instead.
type AccountMagicSiteWanService struct {
	Options []option.RequestOption
}

// NewAccountMagicSiteWanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMagicSiteWanService(opts ...option.RequestOption) (r *AccountMagicSiteWanService) {
	r = &AccountMagicSiteWanService{}
	r.Options = opts
	return
}

// Creates a new Site WAN.
func (r *AccountMagicSiteWanService) New(ctx context.Context, accountID string, siteID string, body AccountMagicSiteWanNewParams, opts ...option.RequestOption) (res *MagicWansCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a specific Site WAN.
func (r *AccountMagicSiteWanService) Get(ctx context.Context, accountID string, siteID string, wanID string, opts ...option.RequestOption) (res *AccountMagicSiteWanGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if wanID == "" {
		err = errors.New("missing required wan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a specific Site WAN.
func (r *AccountMagicSiteWanService) Update(ctx context.Context, accountID string, siteID string, wanID string, body AccountMagicSiteWanUpdateParams, opts ...option.RequestOption) (res *MagicWanModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if wanID == "" {
		err = errors.New("missing required wan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists Site WANs associated with an account.
func (r *AccountMagicSiteWanService) List(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *MagicWansCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove a specific Site WAN.
func (r *AccountMagicSiteWanService) Delete(ctx context.Context, accountID string, siteID string, wanID string, opts ...option.RequestOption) (res *AccountMagicSiteWanDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if wanID == "" {
		err = errors.New("missing required wan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patch a specific Site WAN.
func (r *AccountMagicSiteWanService) Patch(ctx context.Context, accountID string, siteID string, wanID string, body AccountMagicSiteWanPatchParams, opts ...option.RequestOption) (res *MagicWanModifiedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if wanID == "" {
		err = errors.New("missing required wan_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/wans/%s", accountID, siteID, wanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MagicWan struct {
	// Identifier
	ID string `json:"id"`
	// Magic WAN health check rate for tunnels created on this link. The default value
	// is `mid`.
	HealthCheckRate MagicWanHealthCheckRate `json:"health_check_rate"`
	Name            string                  `json:"name"`
	Physport        int64                   `json:"physport"`
	// Priority of WAN for traffic loadbalancing.
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string `json:"site_id"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing MagicWanStaticAddressing `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag int64        `json:"vlan_tag"`
	JSON    magicWanJSON `json:"-"`
}

// magicWanJSON contains the JSON metadata for the struct [MagicWan]
type magicWanJSON struct {
	ID               apijson.Field
	HealthCheckRate  apijson.Field
	Name             apijson.Field
	Physport         apijson.Field
	Priority         apijson.Field
	SiteID           apijson.Field
	StaticAddressing apijson.Field
	VlanTag          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicWan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicWanJSON) RawJSON() string {
	return r.raw
}

// Magic WAN health check rate for tunnels created on this link. The default value
// is `mid`.
type MagicWanHealthCheckRate string

const (
	MagicWanHealthCheckRateLow  MagicWanHealthCheckRate = "low"
	MagicWanHealthCheckRateMid  MagicWanHealthCheckRate = "mid"
	MagicWanHealthCheckRateHigh MagicWanHealthCheckRate = "high"
)

func (r MagicWanHealthCheckRate) IsKnown() bool {
	switch r {
	case MagicWanHealthCheckRateLow, MagicWanHealthCheckRateMid, MagicWanHealthCheckRateHigh:
		return true
	}
	return false
}

type MagicWanModifiedResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicWan           `json:"result,required"`
	// Whether the API call was successful
	Success MagicWanModifiedResponseSuccess `json:"success,required"`
	JSON    magicWanModifiedResponseJSON    `json:"-"`
}

// magicWanModifiedResponseJSON contains the JSON metadata for the struct
// [MagicWanModifiedResponse]
type magicWanModifiedResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicWanModifiedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicWanModifiedResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicWanModifiedResponseSuccess bool

const (
	MagicWanModifiedResponseSuccessTrue MagicWanModifiedResponseSuccess = true
)

func (r MagicWanModifiedResponseSuccess) IsKnown() bool {
	switch r {
	case MagicWanModifiedResponseSuccessTrue:
		return true
	}
	return false
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type MagicWanStaticAddressing struct {
	// A valid CIDR notation representing an IP range.
	Address string `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress string `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress string                       `json:"secondary_address"`
	JSON             magicWanStaticAddressingJSON `json:"-"`
}

// magicWanStaticAddressingJSON contains the JSON metadata for the struct
// [MagicWanStaticAddressing]
type magicWanStaticAddressingJSON struct {
	Address          apijson.Field
	GatewayAddress   apijson.Field
	SecondaryAddress apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicWanStaticAddressing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicWanStaticAddressingJSON) RawJSON() string {
	return r.raw
}

// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
// availability mode.
type MagicWanStaticAddressingParam struct {
	// A valid CIDR notation representing an IP range.
	Address param.Field[string] `json:"address,required"`
	// A valid IPv4 address.
	GatewayAddress param.Field[string] `json:"gateway_address,required"`
	// A valid CIDR notation representing an IP range.
	SecondaryAddress param.Field[string] `json:"secondary_address"`
}

func (r MagicWanStaticAddressingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicWanUpdateRequestParam struct {
	Name     param.Field[string] `json:"name"`
	Physport param.Field[int64]  `json:"physport"`
	Priority param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[MagicWanStaticAddressingParam] `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r MagicWanUpdateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MagicWansCollectionResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   []MagicWan         `json:"result,required"`
	// Whether the API call was successful
	Success MagicWansCollectionResponseSuccess `json:"success,required"`
	JSON    magicWansCollectionResponseJSON    `json:"-"`
}

// magicWansCollectionResponseJSON contains the JSON metadata for the struct
// [MagicWansCollectionResponse]
type magicWansCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicWansCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicWansCollectionResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicWansCollectionResponseSuccess bool

const (
	MagicWansCollectionResponseSuccessTrue MagicWansCollectionResponseSuccess = true
)

func (r MagicWansCollectionResponseSuccess) IsKnown() bool {
	switch r {
	case MagicWansCollectionResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteWanGetResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicWan           `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteWanGetResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteWanGetResponseJSON    `json:"-"`
}

// accountMagicSiteWanGetResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteWanGetResponse]
type accountMagicSiteWanGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteWanGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteWanGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteWanGetResponseSuccess bool

const (
	AccountMagicSiteWanGetResponseSuccessTrue AccountMagicSiteWanGetResponseSuccess = true
)

func (r AccountMagicSiteWanGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteWanGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteWanDeleteResponse struct {
	Errors   []MagicMessageItem `json:"errors,required"`
	Messages []MagicMessageItem `json:"messages,required"`
	Result   MagicWan           `json:"result,required"`
	// Whether the API call was successful
	Success AccountMagicSiteWanDeleteResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteWanDeleteResponseJSON    `json:"-"`
}

// accountMagicSiteWanDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMagicSiteWanDeleteResponse]
type accountMagicSiteWanDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteWanDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteWanDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteWanDeleteResponseSuccess bool

const (
	AccountMagicSiteWanDeleteResponseSuccessTrue AccountMagicSiteWanDeleteResponseSuccess = true
)

func (r AccountMagicSiteWanDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteWanDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteWanNewParams struct {
	Physport param.Field[int64]  `json:"physport,required"`
	Name     param.Field[string] `json:"name"`
	Priority param.Field[int64]  `json:"priority"`
	// (optional) if omitted, use DHCP. Submit secondary_address when site is in high
	// availability mode.
	StaticAddressing param.Field[MagicWanStaticAddressingParam] `json:"static_addressing"`
	// VLAN ID. Use zero for untagged.
	VlanTag param.Field[int64] `json:"vlan_tag"`
}

func (r AccountMagicSiteWanNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMagicSiteWanUpdateParams struct {
	MagicWanUpdateRequest MagicWanUpdateRequestParam `json:"magic_wan_update_request,required"`
}

func (r AccountMagicSiteWanUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicWanUpdateRequest)
}

type AccountMagicSiteWanPatchParams struct {
	MagicWanUpdateRequest MagicWanUpdateRequestParam `json:"magic_wan_update_request,required"`
}

func (r AccountMagicSiteWanPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MagicWanUpdateRequest)
}
