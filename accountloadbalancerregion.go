// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountLoadBalancerRegionService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLoadBalancerRegionService] method instead.
type AccountLoadBalancerRegionService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerRegionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerRegionService(opts ...option.RequestOption) (r *AccountLoadBalancerRegionService) {
	r = &AccountLoadBalancerRegionService{}
	r.Options = opts
	return
}

// Get a single region mapping.
func (r *AccountLoadBalancerRegionService) Get(ctx context.Context, accountID string, regionID AccountLoadBalancerRegionGetParamsRegionID, opts ...option.RequestOption) (res *AccountLoadBalancerRegionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", accountID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all region mappings.
func (r *AccountLoadBalancerRegionService) List(ctx context.Context, accountID string, query AccountLoadBalancerRegionListParams, opts ...option.RequestOption) (res *AccountLoadBalancerRegionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountLoadBalancerRegionGetResponse struct {
	// A list of countries and subdivisions mapped to a region.
	Result interface{}                              `json:"result"`
	JSON   accountLoadBalancerRegionGetResponseJSON `json:"-"`
	SingleResponseMonitor
}

// accountLoadBalancerRegionGetResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerRegionGetResponse]
type accountLoadBalancerRegionGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLoadBalancerRegionGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLoadBalancerRegionListResponse struct {
	Result interface{}                               `json:"result"`
	JSON   accountLoadBalancerRegionListResponseJSON `json:"-"`
	SingleResponseMonitor
}

// accountLoadBalancerRegionListResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerRegionListResponse]
type accountLoadBalancerRegionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerRegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLoadBalancerRegionListResponseJSON) RawJSON() string {
	return r.raw
}

// A list of Cloudflare regions. WNAM: Western North America, ENAM: Eastern North
// America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America,
// SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa,
// SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East
// Asia).
type AccountLoadBalancerRegionGetParamsRegionID string

const (
	AccountLoadBalancerRegionGetParamsRegionIDWnam AccountLoadBalancerRegionGetParamsRegionID = "WNAM"
	AccountLoadBalancerRegionGetParamsRegionIDEnam AccountLoadBalancerRegionGetParamsRegionID = "ENAM"
	AccountLoadBalancerRegionGetParamsRegionIDWeu  AccountLoadBalancerRegionGetParamsRegionID = "WEU"
	AccountLoadBalancerRegionGetParamsRegionIDEeu  AccountLoadBalancerRegionGetParamsRegionID = "EEU"
	AccountLoadBalancerRegionGetParamsRegionIDNsam AccountLoadBalancerRegionGetParamsRegionID = "NSAM"
	AccountLoadBalancerRegionGetParamsRegionIDSsam AccountLoadBalancerRegionGetParamsRegionID = "SSAM"
	AccountLoadBalancerRegionGetParamsRegionIDOc   AccountLoadBalancerRegionGetParamsRegionID = "OC"
	AccountLoadBalancerRegionGetParamsRegionIDMe   AccountLoadBalancerRegionGetParamsRegionID = "ME"
	AccountLoadBalancerRegionGetParamsRegionIDNaf  AccountLoadBalancerRegionGetParamsRegionID = "NAF"
	AccountLoadBalancerRegionGetParamsRegionIDSaf  AccountLoadBalancerRegionGetParamsRegionID = "SAF"
	AccountLoadBalancerRegionGetParamsRegionIDSas  AccountLoadBalancerRegionGetParamsRegionID = "SAS"
	AccountLoadBalancerRegionGetParamsRegionIDSeas AccountLoadBalancerRegionGetParamsRegionID = "SEAS"
	AccountLoadBalancerRegionGetParamsRegionIDNeas AccountLoadBalancerRegionGetParamsRegionID = "NEAS"
)

func (r AccountLoadBalancerRegionGetParamsRegionID) IsKnown() bool {
	switch r {
	case AccountLoadBalancerRegionGetParamsRegionIDWnam, AccountLoadBalancerRegionGetParamsRegionIDEnam, AccountLoadBalancerRegionGetParamsRegionIDWeu, AccountLoadBalancerRegionGetParamsRegionIDEeu, AccountLoadBalancerRegionGetParamsRegionIDNsam, AccountLoadBalancerRegionGetParamsRegionIDSsam, AccountLoadBalancerRegionGetParamsRegionIDOc, AccountLoadBalancerRegionGetParamsRegionIDMe, AccountLoadBalancerRegionGetParamsRegionIDNaf, AccountLoadBalancerRegionGetParamsRegionIDSaf, AccountLoadBalancerRegionGetParamsRegionIDSas, AccountLoadBalancerRegionGetParamsRegionIDSeas, AccountLoadBalancerRegionGetParamsRegionIDNeas:
		return true
	}
	return false
}

type AccountLoadBalancerRegionListParams struct {
	// Two-letter alpha-2 country code followed in ISO 3166-1.
	CountryCodeA2 param.Field[string] `query:"country_code_a2"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCode param.Field[string] `query:"subdivision_code"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCodeA2 param.Field[string] `query:"subdivision_code_a2"`
}

// URLQuery serializes [AccountLoadBalancerRegionListParams]'s query parameters as
// `url.Values`.
func (r AccountLoadBalancerRegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
