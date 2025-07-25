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

// AccountAddressingPrefixBgpPrefixService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixBgpPrefixService] method instead.
type AccountAddressingPrefixBgpPrefixService struct {
	Options []option.RequestOption
}

// NewAccountAddressingPrefixBgpPrefixService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixBgpPrefixService(opts ...option.RequestOption) (r *AccountAddressingPrefixBgpPrefixService) {
	r = &AccountAddressingPrefixBgpPrefixService{}
	r.Options = opts
	return
}

// Create a BGP prefix, controlling the BGP advertisement status of a specific
// subnet. When created, BGP prefixes are initially withdrawn, and can be
// advertised with the Update BGP Prefix API.
func (r *AccountAddressingPrefixBgpPrefixService) New(ctx context.Context, accountID string, prefixID string, body AccountAddressingPrefixBgpPrefixNewParams, opts ...option.RequestOption) (res *SingleResponseBgp, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *AccountAddressingPrefixBgpPrefixService) Get(ctx context.Context, accountID string, prefixID string, bgpPrefixID string, opts ...option.RequestOption) (res *SingleResponseBgp, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bgpPrefixID == "" {
		err = errors.New("missing required bgp_prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the properties of a BGP Prefix, such as the on demand advertisement
// status (advertised or withdrawn).
func (r *AccountAddressingPrefixBgpPrefixService) Update(ctx context.Context, accountID string, prefixID string, bgpPrefixID string, body AccountAddressingPrefixBgpPrefixUpdateParams, opts ...option.RequestOption) (res *SingleResponseBgp, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bgpPrefixID == "" {
		err = errors.New("missing required bgp_prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", accountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *AccountAddressingPrefixBgpPrefixService) List(ctx context.Context, accountID string, prefixID string, opts ...option.RequestOption) (res *AccountAddressingPrefixBgpPrefixListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", accountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BgpPrefixes struct {
	// Identifier of BGP Prefix.
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	Asn int64 `json:"asn,nullable"`
	// Number of times to prepend the Cloudflare ASN to the BGP AS-Path attribute
	AsnPrependCount int64 `json:"asn_prepend_count"`
	// Determines if Cloudflare advertises a BYOIP BGP prefix even when there is no
	// matching BGP prefix in the Magic routing table. When true, Cloudflare will
	// automatically withdraw the BGP prefix when there are no matching BGP routes, and
	// will resume advertising when there is at least one matching BGP route.
	AutoAdvertiseWithdraw bool                     `json:"auto_advertise_withdraw"`
	BgpSignalOpts         BgpPrefixesBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string              `json:"cidr"`
	CreatedAt  time.Time           `json:"created_at" format:"date-time"`
	ModifiedAt time.Time           `json:"modified_at" format:"date-time"`
	OnDemand   BgpPrefixesOnDemand `json:"on_demand"`
	JSON       bgpPrefixesJSON     `json:"-"`
}

// bgpPrefixesJSON contains the JSON metadata for the struct [BgpPrefixes]
type bgpPrefixesJSON struct {
	ID                    apijson.Field
	Asn                   apijson.Field
	AsnPrependCount       apijson.Field
	AutoAdvertiseWithdraw apijson.Field
	BgpSignalOpts         apijson.Field
	Cidr                  apijson.Field
	CreatedAt             apijson.Field
	ModifiedAt            apijson.Field
	OnDemand              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BgpPrefixes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixesJSON) RawJSON() string {
	return r.raw
}

type BgpPrefixesBgpSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                    `json:"modified_at,nullable" format:"date-time"`
	JSON       bgpPrefixesBgpSignalOptsJSON `json:"-"`
}

// bgpPrefixesBgpSignalOptsJSON contains the JSON metadata for the struct
// [BgpPrefixesBgpSignalOpts]
type bgpPrefixesBgpSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpPrefixesBgpSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixesBgpSignalOptsJSON) RawJSON() string {
	return r.raw
}

type BgpPrefixesOnDemand struct {
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                    `json:"on_demand_locked"`
	JSON           bgpPrefixesOnDemandJSON `json:"-"`
}

// bgpPrefixesOnDemandJSON contains the JSON metadata for the struct
// [BgpPrefixesOnDemand]
type bgpPrefixesOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BgpPrefixesOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixesOnDemandJSON) RawJSON() string {
	return r.raw
}

type SingleResponseBgp struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SingleResponseBgpSuccess `json:"success,required"`
	Result  BgpPrefixes              `json:"result"`
	JSON    singleResponseBgpJSON    `json:"-"`
}

// singleResponseBgpJSON contains the JSON metadata for the struct
// [SingleResponseBgp]
type singleResponseBgpJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseBgp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseBgpJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleResponseBgpSuccess bool

const (
	SingleResponseBgpSuccessTrue SingleResponseBgpSuccess = true
)

func (r SingleResponseBgpSuccess) IsKnown() bool {
	switch r {
	case SingleResponseBgpSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixBgpPrefixListResponse struct {
	Errors   []AddressingMessages `json:"errors,required"`
	Messages []AddressingMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAddressingPrefixBgpPrefixListResponseSuccess    `json:"success,required"`
	Result     []BgpPrefixes                                          `json:"result"`
	ResultInfo AccountAddressingPrefixBgpPrefixListResponseResultInfo `json:"result_info"`
	JSON       accountAddressingPrefixBgpPrefixListResponseJSON       `json:"-"`
}

// accountAddressingPrefixBgpPrefixListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixBgpPrefixListResponse]
type accountAddressingPrefixBgpPrefixListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBgpPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBgpPrefixListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAddressingPrefixBgpPrefixListResponseSuccess bool

const (
	AccountAddressingPrefixBgpPrefixListResponseSuccessTrue AccountAddressingPrefixBgpPrefixListResponseSuccess = true
)

func (r AccountAddressingPrefixBgpPrefixListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAddressingPrefixBgpPrefixListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAddressingPrefixBgpPrefixListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                    `json:"total_count"`
	JSON       accountAddressingPrefixBgpPrefixListResponseResultInfoJSON `json:"-"`
}

// accountAddressingPrefixBgpPrefixListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAddressingPrefixBgpPrefixListResponseResultInfo]
type accountAddressingPrefixBgpPrefixListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBgpPrefixListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBgpPrefixListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAddressingPrefixBgpPrefixNewParams struct {
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr param.Field[string] `json:"cidr"`
}

func (r AccountAddressingPrefixBgpPrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingPrefixBgpPrefixUpdateParams struct {
	// Number of times to prepend the Cloudflare ASN to the BGP AS-Path attribute
	AsnPrependCount param.Field[int64] `json:"asn_prepend_count"`
	// Determines if Cloudflare advertises a BYOIP BGP prefix even when there is no
	// matching BGP prefix in the Magic routing table. When true, Cloudflare will
	// automatically withdraw the BGP prefix when there are no matching BGP routes, and
	// will resume advertising when there is at least one matching BGP route.
	AutoAdvertiseWithdraw param.Field[bool]                                                 `json:"auto_advertise_withdraw"`
	OnDemand              param.Field[AccountAddressingPrefixBgpPrefixUpdateParamsOnDemand] `json:"on_demand"`
}

func (r AccountAddressingPrefixBgpPrefixUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAddressingPrefixBgpPrefixUpdateParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r AccountAddressingPrefixBgpPrefixUpdateParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
