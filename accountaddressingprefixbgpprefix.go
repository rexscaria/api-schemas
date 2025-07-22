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
	Asn           int64                    `json:"asn,nullable"`
	BgpSignalOpts BgpPrefixesBgpSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	Cidr       string              `json:"cidr"`
	CreatedAt  time.Time           `json:"created_at" format:"date-time"`
	ModifiedAt time.Time           `json:"modified_at" format:"date-time"`
	OnDemand   BgpPrefixesOnDemand `json:"on_demand"`
	JSON       bgpPrefixesJSON     `json:"-"`
}

// bgpPrefixesJSON contains the JSON metadata for the struct [BgpPrefixes]
type bgpPrefixesJSON struct {
	ID            apijson.Field
	Asn           apijson.Field
	BgpSignalOpts apijson.Field
	Cidr          apijson.Field
	CreatedAt     apijson.Field
	ModifiedAt    apijson.Field
	OnDemand      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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
	Result BgpPrefixes           `json:"result"`
	JSON   singleResponseBgpJSON `json:"-"`
	AddressingAPIResponseSingle
}

// singleResponseBgpJSON contains the JSON metadata for the struct
// [SingleResponseBgp]
type singleResponseBgpJSON struct {
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

type AccountAddressingPrefixBgpPrefixListResponse struct {
	Result []BgpPrefixes                                    `json:"result"`
	JSON   accountAddressingPrefixBgpPrefixListResponseJSON `json:"-"`
	APIResponseCollectionAddressing
}

// accountAddressingPrefixBgpPrefixListResponseJSON contains the JSON metadata for
// the struct [AccountAddressingPrefixBgpPrefixListResponse]
type accountAddressingPrefixBgpPrefixListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingPrefixBgpPrefixListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAddressingPrefixBgpPrefixListResponseJSON) RawJSON() string {
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
	OnDemand param.Field[AccountAddressingPrefixBgpPrefixUpdateParamsOnDemand] `json:"on_demand"`
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
