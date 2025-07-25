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

// ZoneAPIGatewayOperationService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayOperationService] method instead.
type ZoneAPIGatewayOperationService struct {
	Options          []option.RequestOption
	SchemaValidation *ZoneAPIGatewayOperationSchemaValidationService
}

// NewZoneAPIGatewayOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayOperationService(opts ...option.RequestOption) (r *ZoneAPIGatewayOperationService) {
	r = &ZoneAPIGatewayOperationService{}
	r.Options = opts
	r.SchemaValidation = NewZoneAPIGatewayOperationSchemaValidationService(opts...)
	return
}

// Retrieve information about an operation
func (r *ZoneAPIGatewayOperationService) Get(ctx context.Context, zoneID string, operationID string, query ZoneAPIGatewayOperationGetParams, opts ...option.RequestOption) (res *SingleOperationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve information about all operations on a zone
func (r *ZoneAPIGatewayOperationService) List(ctx context.Context, zoneID string, query ZoneAPIGatewayOperationListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayOperationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add one or more operations to a zone. Endpoints can contain path variables.
// Host, method, endpoint will be normalized to a canoncial form when creating an
// operation and must be unique on the zone. Inserting an operation that matches an
// existing one will return the record of the already existing operation and update
// its last_updated date.
func (r *ZoneAPIGatewayOperationService) AddMultiple(ctx context.Context, zoneID string, body ZoneAPIGatewayOperationAddMultipleParams, opts ...option.RequestOption) (res *ZoneAPIGatewayOperationAddMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add one operation to a zone. Endpoints can contain path variables. Host, method,
// endpoint will be normalized to a canoncial form when creating an operation and
// must be unique on the zone. Inserting an operation that matches an existing one
// will return the record of the already existing operation and update its
// last_updated date.
func (r *ZoneAPIGatewayOperationService) AddSingle(ctx context.Context, zoneID string, body ZoneAPIGatewayOperationAddSingleParams, opts ...option.RequestOption) (res *SingleOperationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/item", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete multiple operations
func (r *ZoneAPIGatewayOperationService) DeleteMultiple(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *APIResponseAPIShield, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Delete an operation
func (r *ZoneAPIGatewayOperationService) DeleteSingle(ctx context.Context, zoneID string, operationID string, opts ...option.RequestOption) (res *APIResponseAPIShield, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type BasicOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method BasicOperationMethod `json:"method,required"`
	JSON   basicOperationJSON   `json:"-"`
}

// basicOperationJSON contains the JSON metadata for the struct [BasicOperation]
type basicOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BasicOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r basicOperationJSON) RawJSON() string {
	return r.raw
}

func (r BasicOperation) implementsZoneAPIGatewayUserSchemaGetOperationsResponseResult() {}

// The HTTP method used to access the endpoint.
type BasicOperationMethod string

const (
	BasicOperationMethodGet     BasicOperationMethod = "GET"
	BasicOperationMethodPost    BasicOperationMethod = "POST"
	BasicOperationMethodHead    BasicOperationMethod = "HEAD"
	BasicOperationMethodOptions BasicOperationMethod = "OPTIONS"
	BasicOperationMethodPut     BasicOperationMethod = "PUT"
	BasicOperationMethodDelete  BasicOperationMethod = "DELETE"
	BasicOperationMethodConnect BasicOperationMethod = "CONNECT"
	BasicOperationMethodPatch   BasicOperationMethod = "PATCH"
	BasicOperationMethodTrace   BasicOperationMethod = "TRACE"
)

func (r BasicOperationMethod) IsKnown() bool {
	switch r {
	case BasicOperationMethodGet, BasicOperationMethodPost, BasicOperationMethodHead, BasicOperationMethodOptions, BasicOperationMethodPut, BasicOperationMethodDelete, BasicOperationMethodConnect, BasicOperationMethodPatch, BasicOperationMethodTrace:
		return true
	}
	return false
}

type BasicOperationParam struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint param.Field[string] `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method param.Field[BasicOperationMethod] `json:"method,required"`
}

func (r BasicOperationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleOperationResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   Operation               `json:"result,required"`
	// Whether the API call was successful.
	Success SingleOperationResponseSuccess `json:"success,required"`
	JSON    singleOperationResponseJSON    `json:"-"`
}

// singleOperationResponseJSON contains the JSON metadata for the struct
// [SingleOperationResponse]
type singleOperationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleOperationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleOperationResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SingleOperationResponseSuccess bool

const (
	SingleOperationResponseSuccessTrue SingleOperationResponseSuccess = true
)

func (r SingleOperationResponseSuccess) IsKnown() bool {
	switch r {
	case SingleOperationResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationListResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   []Operation             `json:"result,required"`
	// Whether the API call was successful.
	Success    ZoneAPIGatewayOperationListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneAPIGatewayOperationListResponseResultInfo `json:"result_info"`
	JSON       zoneAPIGatewayOperationListResponseJSON       `json:"-"`
}

// zoneAPIGatewayOperationListResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayOperationListResponse]
type zoneAPIGatewayOperationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayOperationListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayOperationListResponseSuccess bool

const (
	ZoneAPIGatewayOperationListResponseSuccessTrue ZoneAPIGatewayOperationListResponseSuccess = true
)

func (r ZoneAPIGatewayOperationListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       zoneAPIGatewayOperationListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayOperationListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayOperationListResponseResultInfo]
type zoneAPIGatewayOperationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayOperationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayOperationListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayOperationAddMultipleResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   []Operation             `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayOperationAddMultipleResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayOperationAddMultipleResponseJSON    `json:"-"`
}

// zoneAPIGatewayOperationAddMultipleResponseJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayOperationAddMultipleResponse]
type zoneAPIGatewayOperationAddMultipleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayOperationAddMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayOperationAddMultipleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayOperationAddMultipleResponseSuccess bool

const (
	ZoneAPIGatewayOperationAddMultipleResponseSuccessTrue ZoneAPIGatewayOperationAddMultipleResponseSuccess = true
)

func (r ZoneAPIGatewayOperationAddMultipleResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationAddMultipleResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationGetParams struct {
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewayOperationGetParamsFeature] `query:"feature"`
}

// URLQuery serializes [ZoneAPIGatewayOperationGetParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayOperationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayOperationGetParamsFeature string

const (
	ZoneAPIGatewayOperationGetParamsFeatureThresholds       ZoneAPIGatewayOperationGetParamsFeature = "thresholds"
	ZoneAPIGatewayOperationGetParamsFeatureParameterSchemas ZoneAPIGatewayOperationGetParamsFeature = "parameter_schemas"
	ZoneAPIGatewayOperationGetParamsFeatureSchemaInfo       ZoneAPIGatewayOperationGetParamsFeature = "schema_info"
)

func (r ZoneAPIGatewayOperationGetParamsFeature) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationGetParamsFeatureThresholds, ZoneAPIGatewayOperationGetParamsFeatureParameterSchemas, ZoneAPIGatewayOperationGetParamsFeatureSchemaInfo:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationListParams struct {
	// Direction to order results.
	Direction param.Field[ZoneAPIGatewayOperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewayOperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by. When requesting a feature, the feature keys are available for
	// ordering as well, e.g., `thresholds.suggested_threshold`.
	Order param.Field[ZoneAPIGatewayOperationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ZoneAPIGatewayOperationListParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type ZoneAPIGatewayOperationListParamsDirection string

const (
	ZoneAPIGatewayOperationListParamsDirectionAsc  ZoneAPIGatewayOperationListParamsDirection = "asc"
	ZoneAPIGatewayOperationListParamsDirectionDesc ZoneAPIGatewayOperationListParamsDirection = "desc"
)

func (r ZoneAPIGatewayOperationListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationListParamsDirectionAsc, ZoneAPIGatewayOperationListParamsDirectionDesc:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationListParamsFeature string

const (
	ZoneAPIGatewayOperationListParamsFeatureThresholds       ZoneAPIGatewayOperationListParamsFeature = "thresholds"
	ZoneAPIGatewayOperationListParamsFeatureParameterSchemas ZoneAPIGatewayOperationListParamsFeature = "parameter_schemas"
	ZoneAPIGatewayOperationListParamsFeatureSchemaInfo       ZoneAPIGatewayOperationListParamsFeature = "schema_info"
)

func (r ZoneAPIGatewayOperationListParamsFeature) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationListParamsFeatureThresholds, ZoneAPIGatewayOperationListParamsFeatureParameterSchemas, ZoneAPIGatewayOperationListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Field to order by. When requesting a feature, the feature keys are available for
// ordering as well, e.g., `thresholds.suggested_threshold`.
type ZoneAPIGatewayOperationListParamsOrder string

const (
	ZoneAPIGatewayOperationListParamsOrderMethod        ZoneAPIGatewayOperationListParamsOrder = "method"
	ZoneAPIGatewayOperationListParamsOrderHost          ZoneAPIGatewayOperationListParamsOrder = "host"
	ZoneAPIGatewayOperationListParamsOrderEndpoint      ZoneAPIGatewayOperationListParamsOrder = "endpoint"
	ZoneAPIGatewayOperationListParamsOrderThresholdsKey ZoneAPIGatewayOperationListParamsOrder = "thresholds.$key"
)

func (r ZoneAPIGatewayOperationListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationListParamsOrderMethod, ZoneAPIGatewayOperationListParamsOrderHost, ZoneAPIGatewayOperationListParamsOrderEndpoint, ZoneAPIGatewayOperationListParamsOrderThresholdsKey:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationAddMultipleParams struct {
	Body []BasicOperationParam `json:"body,required"`
}

func (r ZoneAPIGatewayOperationAddMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneAPIGatewayOperationAddSingleParams struct {
	BasicOperation BasicOperationParam `json:"basic_operation,required"`
}

func (r ZoneAPIGatewayOperationAddSingleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BasicOperation)
}
