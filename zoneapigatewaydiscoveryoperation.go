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

// ZoneAPIGatewayDiscoveryOperationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayDiscoveryOperationService] method instead.
type ZoneAPIGatewayDiscoveryOperationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayDiscoveryOperationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayDiscoveryOperationService(opts ...option.RequestOption) (r *ZoneAPIGatewayDiscoveryOperationService) {
	r = &ZoneAPIGatewayDiscoveryOperationService{}
	r.Options = opts
	return
}

// Update the `state` on one or more discovered operations
func (r *ZoneAPIGatewayDiscoveryOperationService) Update(ctx context.Context, zoneID string, body ZoneAPIGatewayDiscoveryOperationUpdateParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve the most up to date view of discovered operations
func (r *ZoneAPIGatewayDiscoveryOperationService) List(ctx context.Context, zoneID string, query ZoneAPIGatewayDiscoveryOperationListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update the `state` on a discovered operation
func (r *ZoneAPIGatewayDiscoveryOperationService) UpdateSingle(ctx context.Context, zoneID string, operationID SchemasUuidParam, body ZoneAPIGatewayDiscoveryOperationUpdateSingleParams, opts ...option.RequestOption) (res *ZoneAPIGatewayDiscoveryOperationUpdateSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/discovery/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// State of operation in API Discovery
//
// - `review` - Operation is not saved into API Shield Endpoint Management
// - `saved` - Operation is saved into API Shield Endpoint Management
// - `ignored` - Operation is marked as ignored
type APIDiscoveryState string

const (
	APIDiscoveryStateReview  APIDiscoveryState = "review"
	APIDiscoveryStateSaved   APIDiscoveryState = "saved"
	APIDiscoveryStateIgnored APIDiscoveryState = "ignored"
)

func (r APIDiscoveryState) IsKnown() bool {
	switch r {
	case APIDiscoveryStateReview, APIDiscoveryStateSaved, APIDiscoveryStateIgnored:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationUpdateResponse struct {
	Errors   []MessagesAPIShieldItem                                         `json:"errors,required"`
	Messages []MessagesAPIShieldItem                                         `json:"messages,required"`
	Result   map[string]ZoneAPIGatewayDiscoveryOperationUpdateResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateResponseJSON    `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateResponseJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayDiscoveryOperationUpdateResponse]
type zoneAPIGatewayDiscoveryOperationUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Mappings of discovered operations (keys) to objects describing their state
type ZoneAPIGatewayDiscoveryOperationUpdateResponseResult struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State ZoneAPIGatewayDiscoveryOperationUpdateResponseResultState `json:"state"`
	JSON  zoneAPIGatewayDiscoveryOperationUpdateResponseResultJSON  `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateResponseResultJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationUpdateResponseResult]
type zoneAPIGatewayDiscoveryOperationUpdateResponseResultJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type ZoneAPIGatewayDiscoveryOperationUpdateResponseResultState string

const (
	ZoneAPIGatewayDiscoveryOperationUpdateResponseResultStateReview  ZoneAPIGatewayDiscoveryOperationUpdateResponseResultState = "review"
	ZoneAPIGatewayDiscoveryOperationUpdateResponseResultStateIgnored ZoneAPIGatewayDiscoveryOperationUpdateResponseResultState = "ignored"
)

func (r ZoneAPIGatewayDiscoveryOperationUpdateResponseResultState) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationUpdateResponseResultStateReview, ZoneAPIGatewayDiscoveryOperationUpdateResponseResultStateIgnored:
		return true
	}
	return false
}

// Whether the API call was successful.
type ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccess bool

const (
	ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccessTrue ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccess = true
)

func (r ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationListResponse struct {
	Errors   []MessagesAPIShieldItem                              `json:"errors,required"`
	Messages []MessagesAPIShieldItem                              `json:"messages,required"`
	Result   []ZoneAPIGatewayDiscoveryOperationListResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success    ZoneAPIGatewayDiscoveryOperationListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneAPIGatewayDiscoveryOperationListResponseResultInfo `json:"result_info"`
	JSON       zoneAPIGatewayDiscoveryOperationListResponseJSON       `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayDiscoveryOperationListResponse]
type zoneAPIGatewayDiscoveryOperationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayDiscoveryOperationListResponseResult struct {
	// UUID.
	ID SchemasUuid `json:"id,required"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string           `json:"host,required" format:"hostname"`
	LastUpdated SchemasTimestamp `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method ZoneAPIGatewayDiscoveryOperationListResponseResultMethod `json:"method,required"`
	// API discovery engine(s) that discovered this operation
	Origin []ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin `json:"origin,required"`
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State    APIDiscoveryState                                          `json:"state,required"`
	Features ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures `json:"features"`
	JSON     zoneAPIGatewayDiscoveryOperationListResponseResultJSON     `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationListResponseResult]
type zoneAPIGatewayDiscoveryOperationListResponseResultJSON struct {
	ID          apijson.Field
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	Origin      apijson.Field
	State       apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationListResponseResultJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type ZoneAPIGatewayDiscoveryOperationListResponseResultMethod string

const (
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodGet     ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "GET"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPost    ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "POST"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodHead    ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "HEAD"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodOptions ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "OPTIONS"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPut     ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "PUT"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodDelete  ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "DELETE"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodConnect ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "CONNECT"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPatch   ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "PATCH"
	ZoneAPIGatewayDiscoveryOperationListResponseResultMethodTrace   ZoneAPIGatewayDiscoveryOperationListResponseResultMethod = "TRACE"
)

func (r ZoneAPIGatewayDiscoveryOperationListResponseResultMethod) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListResponseResultMethodGet, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPost, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodHead, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodOptions, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPut, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodDelete, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodConnect, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodPatch, ZoneAPIGatewayDiscoveryOperationListResponseResultMethodTrace:
		return true
	}
	return false
}

//   - `ML` - Discovered operation was sourced using ML API Discovery _
//     `SessionIdentifier` - Discovered operation was sourced using Session
//     Identifier API Discovery _ `LabelDiscovery` - Discovered operation was
//     identified to have a specific label
type ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin string

const (
	ZoneAPIGatewayDiscoveryOperationListResponseResultOriginMl                ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin = "ML"
	ZoneAPIGatewayDiscoveryOperationListResponseResultOriginSessionIdentifier ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin = "SessionIdentifier"
	ZoneAPIGatewayDiscoveryOperationListResponseResultOriginLabelDiscovery    ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin = "LabelDiscovery"
)

func (r ZoneAPIGatewayDiscoveryOperationListResponseResultOrigin) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListResponseResultOriginMl, ZoneAPIGatewayDiscoveryOperationListResponseResultOriginSessionIdentifier, ZoneAPIGatewayDiscoveryOperationListResponseResultOriginLabelDiscovery:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures struct {
	TrafficStats ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats `json:"traffic_stats"`
	JSON         zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON         `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures]
type zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON struct {
	TrafficStats apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultFeatures) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats struct {
	LastUpdated SchemasTimestamp `json:"last_updated,required" format:"date-time"`
	// The period in seconds these statistics were computed over
	PeriodSeconds int64 `json:"period_seconds,required"`
	// The average number of requests seen during this period
	Requests float64                                                                    `json:"requests,required"`
	JSON     zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON
// contains the JSON metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats]
type zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON struct {
	LastUpdated   apijson.Field
	PeriodSeconds apijson.Field
	Requests      apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationListResponseResultFeaturesTrafficStatsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayDiscoveryOperationListResponseSuccess bool

const (
	ZoneAPIGatewayDiscoveryOperationListResponseSuccessTrue ZoneAPIGatewayDiscoveryOperationListResponseSuccess = true
)

func (r ZoneAPIGatewayDiscoveryOperationListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                    `json:"total_count"`
	JSON       zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationListResponseResultInfo]
type zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayDiscoveryOperationUpdateSingleResponse struct {
	Errors   []MessagesAPIShieldItem                                    `json:"errors,required"`
	Messages []MessagesAPIShieldItem                                    `json:"messages,required"`
	Result   ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayDiscoveryOperationUpdateSingleResponseJSON    `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateSingleResponseJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayDiscoveryOperationUpdateSingleResponse]
type zoneAPIGatewayDiscoveryOperationUpdateSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationUpdateSingleResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseResult struct {
	// State of operation in API Discovery
	//
	// - `review` - Operation is not saved into API Shield Endpoint Management
	// - `saved` - Operation is saved into API Shield Endpoint Management
	// - `ignored` - Operation is marked as ignored
	State APIDiscoveryState                                              `json:"state"`
	JSON  zoneAPIGatewayDiscoveryOperationUpdateSingleResponseResultJSON `json:"-"`
}

// zoneAPIGatewayDiscoveryOperationUpdateSingleResponseResultJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseResult]
type zoneAPIGatewayDiscoveryOperationUpdateSingleResponseResultJSON struct {
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayDiscoveryOperationUpdateSingleResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccess bool

const (
	ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccessTrue ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccess = true
)

func (r ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationUpdateSingleResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationUpdateParams struct {
	Body map[string]ZoneAPIGatewayDiscoveryOperationUpdateParamsBody `json:"body,required"`
}

func (r ZoneAPIGatewayDiscoveryOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Mappings of discovered operations (keys) to objects describing their state
type ZoneAPIGatewayDiscoveryOperationUpdateParamsBody struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyState] `json:"state"`
}

func (r ZoneAPIGatewayDiscoveryOperationUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyState string

const (
	ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateReview  ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyState = "review"
	ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateIgnored ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyState = "ignored"
)

func (r ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyState) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateReview, ZoneAPIGatewayDiscoveryOperationUpdateParamsBodyStateIgnored:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationListParams struct {
	// When `true`, only return API Discovery results that are not saved into API
	// Shield Endpoint Management
	Diff param.Field[bool] `query:"diff"`
	// Direction to order results.
	Direction param.Field[ZoneAPIGatewayDiscoveryOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by
	Order param.Field[ZoneAPIGatewayDiscoveryOperationListParamsOrder] `query:"order"`
	// Filter results to only include discovery results sourced from a particular
	// discovery engine
	//
	//   - `ML` - Discovered operations that were sourced using ML API Discovery
	//   - `SessionIdentifier` - Discovered operations that were sourced using Session
	//     Identifier API Discovery
	Origin param.Field[ZoneAPIGatewayDiscoveryOperationListParamsOrigin] `query:"origin"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter results to only include discovery results in a particular state. States
	// are as follows
	//
	//   - `review` - Discovered operations that are not saved into API Shield Endpoint
	//     Management
	//   - `saved` - Discovered operations that are already saved into API Shield
	//     Endpoint Management
	//   - `ignored` - Discovered operations that have been marked as ignored
	State param.Field[APIDiscoveryState] `query:"state"`
}

// URLQuery serializes [ZoneAPIGatewayDiscoveryOperationListParams]'s query
// parameters as `url.Values`.
func (r ZoneAPIGatewayDiscoveryOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type ZoneAPIGatewayDiscoveryOperationListParamsDirection string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsDirectionAsc  ZoneAPIGatewayDiscoveryOperationListParamsDirection = "asc"
	ZoneAPIGatewayDiscoveryOperationListParamsDirectionDesc ZoneAPIGatewayDiscoveryOperationListParamsDirection = "desc"
)

func (r ZoneAPIGatewayDiscoveryOperationListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListParamsDirectionAsc, ZoneAPIGatewayDiscoveryOperationListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order by
type ZoneAPIGatewayDiscoveryOperationListParamsOrder string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsOrderHost                    ZoneAPIGatewayDiscoveryOperationListParamsOrder = "host"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderMethod                  ZoneAPIGatewayDiscoveryOperationListParamsOrder = "method"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderEndpoint                ZoneAPIGatewayDiscoveryOperationListParamsOrder = "endpoint"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsRequests    ZoneAPIGatewayDiscoveryOperationListParamsOrder = "traffic_stats.requests"
	ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsLastUpdated ZoneAPIGatewayDiscoveryOperationListParamsOrder = "traffic_stats.last_updated"
)

func (r ZoneAPIGatewayDiscoveryOperationListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListParamsOrderHost, ZoneAPIGatewayDiscoveryOperationListParamsOrderMethod, ZoneAPIGatewayDiscoveryOperationListParamsOrderEndpoint, ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsRequests, ZoneAPIGatewayDiscoveryOperationListParamsOrderTrafficStatsLastUpdated:
		return true
	}
	return false
}

// Filter results to only include discovery results sourced from a particular
// discovery engine
//
//   - `ML` - Discovered operations that were sourced using ML API Discovery
//   - `SessionIdentifier` - Discovered operations that were sourced using Session
//     Identifier API Discovery
type ZoneAPIGatewayDiscoveryOperationListParamsOrigin string

const (
	ZoneAPIGatewayDiscoveryOperationListParamsOriginMl                ZoneAPIGatewayDiscoveryOperationListParamsOrigin = "ML"
	ZoneAPIGatewayDiscoveryOperationListParamsOriginSessionIdentifier ZoneAPIGatewayDiscoveryOperationListParamsOrigin = "SessionIdentifier"
	ZoneAPIGatewayDiscoveryOperationListParamsOriginLabelDiscovery    ZoneAPIGatewayDiscoveryOperationListParamsOrigin = "LabelDiscovery"
)

func (r ZoneAPIGatewayDiscoveryOperationListParamsOrigin) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationListParamsOriginMl, ZoneAPIGatewayDiscoveryOperationListParamsOriginSessionIdentifier, ZoneAPIGatewayDiscoveryOperationListParamsOriginLabelDiscovery:
		return true
	}
	return false
}

type ZoneAPIGatewayDiscoveryOperationUpdateSingleParams struct {
	// Mark state of operation in API Discovery
	//
	// - `review` - Mark operation as for review
	// - `ignored` - Mark operation as ignored
	State param.Field[ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsState] `json:"state"`
}

func (r ZoneAPIGatewayDiscoveryOperationUpdateSingleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark state of operation in API Discovery
//
// - `review` - Mark operation as for review
// - `ignored` - Mark operation as ignored
type ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsState string

const (
	ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsStateReview  ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsState = "review"
	ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsStateIgnored ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsState = "ignored"
)

func (r ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsState) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsStateReview, ZoneAPIGatewayDiscoveryOperationUpdateSingleParamsStateIgnored:
		return true
	}
	return false
}
