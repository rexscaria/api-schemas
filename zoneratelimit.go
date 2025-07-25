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

// ZoneRateLimitService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRateLimitService] method instead.
type ZoneRateLimitService struct {
	Options []option.RequestOption
}

// NewZoneRateLimitService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRateLimitService(opts ...option.RequestOption) (r *ZoneRateLimitService) {
	r = &ZoneRateLimitService{}
	r.Options = opts
	return
}

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
//
// Deprecated: deprecated
func (r *ZoneRateLimitService) New(ctx context.Context, zoneID string, body ZoneRateLimitNewParams, opts ...option.RequestOption) (res *FirewallRatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a rate limit.
//
// Deprecated: deprecated
func (r *ZoneRateLimitService) Get(ctx context.Context, zoneID string, rateLimitID string, opts ...option.RequestOption) (res *FirewallRatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing rate limit.
//
// Deprecated: deprecated
func (r *ZoneRateLimitService) Update(ctx context.Context, zoneID string, rateLimitID string, body ZoneRateLimitUpdateParams, opts ...option.RequestOption) (res *FirewallRatelimitSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the rate limits for a zone.
//
// Deprecated: deprecated
func (r *ZoneRateLimitService) List(ctx context.Context, zoneID string, query ZoneRateLimitListParams, opts ...option.RequestOption) (res *ZoneRateLimitListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing rate limit.
//
// Deprecated: deprecated
func (r *ZoneRateLimitService) Delete(ctx context.Context, zoneID string, rateLimitID string, opts ...option.RequestOption) (res *ZoneRateLimitDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", zoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Determines which traffic the rate limit counts towards the threshold.
type FirewallMatch struct {
	Headers  []FirewallMatchHeader `json:"headers"`
	Request  FirewallMatchRequest  `json:"request"`
	Response FirewallMatchResponse `json:"response"`
	JSON     firewallMatchJSON     `json:"-"`
}

// firewallMatchJSON contains the JSON metadata for the struct [FirewallMatch]
type firewallMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMatchJSON) RawJSON() string {
	return r.raw
}

type FirewallMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op FirewallMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                  `json:"value"`
	JSON  firewallMatchHeaderJSON `json:"-"`
}

// firewallMatchHeaderJSON contains the JSON metadata for the struct
// [FirewallMatchHeader]
type firewallMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMatchHeaderJSON) RawJSON() string {
	return r.raw
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type FirewallMatchHeadersOp string

const (
	FirewallMatchHeadersOpEq FirewallMatchHeadersOp = "eq"
	FirewallMatchHeadersOpNe FirewallMatchHeadersOp = "ne"
)

func (r FirewallMatchHeadersOp) IsKnown() bool {
	switch r {
	case FirewallMatchHeadersOpEq, FirewallMatchHeadersOpNe:
		return true
	}
	return false
}

type FirewallMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []FirewallMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                   `json:"url"`
	JSON firewallMatchRequestJSON `json:"-"`
}

// firewallMatchRequestJSON contains the JSON metadata for the struct
// [FirewallMatchRequest]
type firewallMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMatchRequestJSON) RawJSON() string {
	return r.raw
}

// An HTTP method or `_ALL_` to indicate all methods.
type FirewallMatchRequestMethod string

const (
	FirewallMatchRequestMethodGet    FirewallMatchRequestMethod = "GET"
	FirewallMatchRequestMethodPost   FirewallMatchRequestMethod = "POST"
	FirewallMatchRequestMethodPut    FirewallMatchRequestMethod = "PUT"
	FirewallMatchRequestMethodDelete FirewallMatchRequestMethod = "DELETE"
	FirewallMatchRequestMethodPatch  FirewallMatchRequestMethod = "PATCH"
	FirewallMatchRequestMethodHead   FirewallMatchRequestMethod = "HEAD"
	FirewallMatchRequestMethod_All   FirewallMatchRequestMethod = "_ALL_"
)

func (r FirewallMatchRequestMethod) IsKnown() bool {
	switch r {
	case FirewallMatchRequestMethodGet, FirewallMatchRequestMethodPost, FirewallMatchRequestMethodPut, FirewallMatchRequestMethodDelete, FirewallMatchRequestMethodPatch, FirewallMatchRequestMethodHead, FirewallMatchRequestMethod_All:
		return true
	}
	return false
}

type FirewallMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                      `json:"origin_traffic"`
	JSON          firewallMatchResponseJSON `json:"-"`
}

// firewallMatchResponseJSON contains the JSON metadata for the struct
// [FirewallMatchResponse]
type firewallMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMatchResponseJSON) RawJSON() string {
	return r.raw
}

// Determines which traffic the rate limit counts towards the threshold.
type FirewallMatchParam struct {
	Headers  param.Field[[]FirewallMatchHeaderParam] `json:"headers"`
	Request  param.Field[FirewallMatchRequestParam]  `json:"request"`
	Response param.Field[FirewallMatchResponseParam] `json:"response"`
}

func (r FirewallMatchParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallMatchHeaderParam struct {
	// The name of the response header to match.
	Name param.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op param.Field[FirewallMatchHeadersOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value param.Field[string] `json:"value"`
}

func (r FirewallMatchHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallMatchRequestParam struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods param.Field[[]FirewallMatchRequestMethod] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes param.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL param.Field[string] `json:"url"`
}

func (r FirewallMatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallMatchResponseParam struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic param.Field[bool] `json:"origin_traffic"`
}

func (r FirewallMatchResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallRateLimits struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action FirewallAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []FirewallRateLimitsBypass `json:"bypass"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match FirewallMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                `json:"threshold"`
	JSON      firewallRateLimitsJSON `json:"-"`
}

// firewallRateLimitsJSON contains the JSON metadata for the struct
// [FirewallRateLimits]
type firewallRateLimitsJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRateLimits) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRateLimitsJSON) RawJSON() string {
	return r.raw
}

type FirewallRateLimitsBypass struct {
	Name FirewallRateLimitsBypassName `json:"name"`
	// The URL to bypass.
	Value string                       `json:"value"`
	JSON  firewallRateLimitsBypassJSON `json:"-"`
}

// firewallRateLimitsBypassJSON contains the JSON metadata for the struct
// [FirewallRateLimitsBypass]
type firewallRateLimitsBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRateLimitsBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRateLimitsBypassJSON) RawJSON() string {
	return r.raw
}

type FirewallRateLimitsBypassName string

const (
	FirewallRateLimitsBypassNameURL FirewallRateLimitsBypassName = "url"
)

func (r FirewallRateLimitsBypassName) IsKnown() bool {
	switch r {
	case FirewallRateLimitsBypassNameURL:
		return true
	}
	return false
}

type FirewallRatelimitSingle struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   FirewallRateLimits     `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallRatelimitSingleSuccess `json:"success,required"`
	JSON    firewallRatelimitSingleJSON    `json:"-"`
}

// firewallRatelimitSingleJSON contains the JSON metadata for the struct
// [FirewallRatelimitSingle]
type firewallRatelimitSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRatelimitSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRatelimitSingleJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallRatelimitSingleSuccess bool

const (
	FirewallRatelimitSingleSuccessTrue FirewallRatelimitSingleSuccess = true
)

func (r FirewallRatelimitSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallRatelimitSingleSuccessTrue:
		return true
	}
	return false
}

type ZoneRateLimitListResponse struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   []FirewallRateLimits   `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    ZoneRateLimitListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneRateLimitListResponseResultInfo `json:"result_info"`
	JSON       zoneRateLimitListResponseJSON       `json:"-"`
}

// zoneRateLimitListResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitListResponse]
type zoneRateLimitListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRateLimitListResponseJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneRateLimitListResponseSuccess bool

const (
	ZoneRateLimitListResponseSuccessTrue ZoneRateLimitListResponseSuccess = true
)

func (r ZoneRateLimitListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRateLimitListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRateLimitListResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                 `json:"total_count"`
	JSON       zoneRateLimitListResponseResultInfoJSON `json:"-"`
}

// zoneRateLimitListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneRateLimitListResponseResultInfo]
type zoneRateLimitListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRateLimitListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneRateLimitDeleteResponse struct {
	Errors   []FirewallMessagesItem            `json:"errors,required"`
	Messages []FirewallMessagesItem            `json:"messages,required"`
	Result   ZoneRateLimitDeleteResponseResult `json:"result,required"`
	// Defines whether the API call was successful.
	Success ZoneRateLimitDeleteResponseSuccess `json:"success,required"`
	JSON    zoneRateLimitDeleteResponseJSON    `json:"-"`
}

// zoneRateLimitDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponse]
type zoneRateLimitDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRateLimitDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneRateLimitDeleteResponseResult struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action FirewallAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []ZoneRateLimitDeleteResponseResultBypass `json:"bypass"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match FirewallMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                               `json:"threshold"`
	JSON      zoneRateLimitDeleteResponseResultJSON `json:"-"`
}

// zoneRateLimitDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneRateLimitDeleteResponseResult]
type zoneRateLimitDeleteResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRateLimitDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneRateLimitDeleteResponseResultBypass struct {
	Name ZoneRateLimitDeleteResponseResultBypassName `json:"name"`
	// The URL to bypass.
	Value string                                      `json:"value"`
	JSON  zoneRateLimitDeleteResponseResultBypassJSON `json:"-"`
}

// zoneRateLimitDeleteResponseResultBypassJSON contains the JSON metadata for the
// struct [ZoneRateLimitDeleteResponseResultBypass]
type zoneRateLimitDeleteResponseResultBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRateLimitDeleteResponseResultBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRateLimitDeleteResponseResultBypassJSON) RawJSON() string {
	return r.raw
}

type ZoneRateLimitDeleteResponseResultBypassName string

const (
	ZoneRateLimitDeleteResponseResultBypassNameURL ZoneRateLimitDeleteResponseResultBypassName = "url"
)

func (r ZoneRateLimitDeleteResponseResultBypassName) IsKnown() bool {
	switch r {
	case ZoneRateLimitDeleteResponseResultBypassNameURL:
		return true
	}
	return false
}

// Defines whether the API call was successful.
type ZoneRateLimitDeleteResponseSuccess bool

const (
	ZoneRateLimitDeleteResponseSuccessTrue ZoneRateLimitDeleteResponseSuccess = true
)

func (r ZoneRateLimitDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRateLimitDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRateLimitNewParams struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[FirewallActionParam] `json:"action,required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[FirewallMatchParam] `json:"match,required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period,required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold,required"`
}

func (r ZoneRateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRateLimitUpdateParams struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[FirewallActionParam] `json:"action,required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[FirewallMatchParam] `json:"match,required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period,required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold,required"`
}

func (r ZoneRateLimitUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRateLimitListParams struct {
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the maximum number of results per page. You can only set the value to
	// `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneRateLimitListParams]'s query parameters as
// `url.Values`.
func (r ZoneRateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
