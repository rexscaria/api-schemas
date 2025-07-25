// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apiform"
	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// ZoneAPIGatewayUserSchemaService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayUserSchemaService] method instead.
type ZoneAPIGatewayUserSchemaService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayUserSchemaService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayUserSchemaService(opts ...option.RequestOption) (r *ZoneAPIGatewayUserSchemaService) {
	r = &ZoneAPIGatewayUserSchemaService{}
	r.Options = opts
	return
}

// Retrieve information about a specific schema on a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) Get(ctx context.Context, zoneID string, schemaID string, query ZoneAPIGatewayUserSchemaGetParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve information about all schemas on a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) List(ctx context.Context, zoneID string, query ZoneAPIGatewayUserSchemaListParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a schema
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) Delete(ctx context.Context, zoneID string, schemaID string, opts ...option.RequestOption) (res *APIResponseAPIShield, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Enable validation for a schema
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) EnableValidation(ctx context.Context, zoneID string, schemaID string, body ZoneAPIGatewayUserSchemaEnableValidationParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaEnableValidationResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Retrieve schema hosts in a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) GetHosts(ctx context.Context, zoneID string, query ZoneAPIGatewayUserSchemaGetHostsParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaGetHostsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/hosts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) GetOperations(ctx context.Context, zoneID string, schemaID string, query ZoneAPIGatewayUserSchemaGetOperationsParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaGetOperationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s/operations", zoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Upload a schema to a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayUserSchemaService) Upload(ctx context.Context, zoneID string, body ZoneAPIGatewayUserSchemaUploadParams, opts ...option.RequestOption) (res *ZoneAPIGatewayUserSchemaUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Kind of schema
type APIShieldKind string

const (
	APIShieldKindOpenAPIV3 APIShieldKind = "openapi_v3"
)

func (r APIShieldKind) IsKnown() bool {
	switch r {
	case APIShieldKindOpenAPIV3:
		return true
	}
	return false
}

// Upper and lower bound for percentile estimate
type ConfidenceIntervalsBounds struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                       `json:"upper"`
	JSON  confidenceIntervalsBoundsJSON `json:"-"`
}

// confidenceIntervalsBoundsJSON contains the JSON metadata for the struct
// [ConfidenceIntervalsBounds]
type confidenceIntervalsBoundsJSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfidenceIntervalsBounds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r confidenceIntervalsBoundsJSON) RawJSON() string {
	return r.raw
}

type Operation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string           `json:"host,required" format:"hostname"`
	LastUpdated SchemasTimestamp `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationMethod `json:"method,required"`
	// UUID.
	OperationID SchemasUuid       `json:"operation_id,required"`
	Features    OperationFeatures `json:"features"`
	JSON        operationJSON     `json:"-"`
}

// operationJSON contains the JSON metadata for the struct [Operation]
type operationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Operation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationJSON) RawJSON() string {
	return r.raw
}

func (r Operation) implementsZoneAPIGatewayUserSchemaGetOperationsResponseResult() {}

// The HTTP method used to access the endpoint.
type OperationMethod string

const (
	OperationMethodGet     OperationMethod = "GET"
	OperationMethodPost    OperationMethod = "POST"
	OperationMethodHead    OperationMethod = "HEAD"
	OperationMethodOptions OperationMethod = "OPTIONS"
	OperationMethodPut     OperationMethod = "PUT"
	OperationMethodDelete  OperationMethod = "DELETE"
	OperationMethodConnect OperationMethod = "CONNECT"
	OperationMethodPatch   OperationMethod = "PATCH"
	OperationMethodTrace   OperationMethod = "TRACE"
)

func (r OperationMethod) IsKnown() bool {
	switch r {
	case OperationMethodGet, OperationMethodPost, OperationMethodHead, OperationMethodOptions, OperationMethodPut, OperationMethodDelete, OperationMethodConnect, OperationMethodPatch, OperationMethodTrace:
		return true
	}
	return false
}

type OperationFeatures struct {
	// This field can have the runtime type of
	// [OperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [OperationFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}           `json:"thresholds"`
	JSON       operationFeaturesJSON `json:"-"`
	union      OperationFeaturesUnion
}

// operationFeaturesJSON contains the JSON metadata for the struct
// [OperationFeatures]
type operationFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationFeaturesUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [OperationFeaturesAPIShieldOperationFeatureThresholds],
// [OperationFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [OperationFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationFeatures) AsUnion() OperationFeaturesUnion {
	return r.union
}

// Union satisfied by [OperationFeaturesAPIShieldOperationFeatureThresholds],
// [OperationFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals] or
// [OperationFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationFeaturesUnion interface {
	implementsOperationFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureThresholdsJSON contains the JSON
// metadata for the struct [OperationFeaturesAPIShieldOperationFeatureThresholds]
type operationFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationFeaturesAPIShieldOperationFeatureThresholds) implementsOperationFeatures() {}

type OperationFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                              `json:"suggested_threshold"`
	JSON               operationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON contains the
// JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureParameterSchemasJSON contains the JSON
// metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationFeaturesAPIShieldOperationFeatureParameterSchemas) implementsOperationFeatures() {}

type OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                    `json:"responses,nullable"`
	JSON      operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains the JSON
// metadata for the struct [OperationFeaturesAPIShieldOperationFeatureAPIRouting]
type operationFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationFeaturesAPIShieldOperationFeatureAPIRouting) implementsOperationFeatures() {}

// API Routing settings on endpoint.
type OperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                             `json:"route"`
	JSON  operationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON contains the
// JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON contains the
// JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsOperationFeatures() {
}

type OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                          `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                `json:"mean"`
	JSON operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 ConfidenceIntervalsBounds `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 ConfidenceIntervalsBounds `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  ConfidenceIntervalsBounds                                                                                                 `json:"p99"`
	JSON operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains the JSON
// metadata for the struct [OperationFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationFeaturesAPIShieldOperationFeatureSchemaInfo) implementsOperationFeatures() {}

type OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON contains the
// JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        SchemasUuid `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                         `json:"name"`
	JSON operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type PublicSchema struct {
	CreatedAt SchemasTimestamp `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind APIShieldKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID.
	SchemaID SchemasUuid `json:"schema_id,required"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool             `json:"validation_enabled"`
	JSON              publicSchemaJSON `json:"-"`
}

// publicSchemaJSON contains the JSON metadata for the struct [PublicSchema]
type publicSchemaJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PublicSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publicSchemaJSON) RawJSON() string {
	return r.raw
}

type SchemasUuid = string

type SchemasUuidParam = string

type ZoneAPIGatewayUserSchemaGetResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   PublicSchema            `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayUserSchemaGetResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayUserSchemaGetResponseJSON    `json:"-"`
}

// zoneAPIGatewayUserSchemaGetResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaGetResponse]
type zoneAPIGatewayUserSchemaGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaGetResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaGetResponseSuccessTrue ZoneAPIGatewayUserSchemaGetResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaListResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   []PublicSchema          `json:"result,required"`
	// Whether the API call was successful.
	Success    ZoneAPIGatewayUserSchemaListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneAPIGatewayUserSchemaListResponseResultInfo `json:"result_info"`
	JSON       zoneAPIGatewayUserSchemaListResponseJSON       `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaListResponse]
type zoneAPIGatewayUserSchemaListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaListResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaListResponseSuccessTrue ZoneAPIGatewayUserSchemaListResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                            `json:"total_count"`
	JSON       zoneAPIGatewayUserSchemaListResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaListResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayUserSchemaListResponseResultInfo]
type zoneAPIGatewayUserSchemaListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaEnableValidationResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	Result   PublicSchema            `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayUserSchemaEnableValidationResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayUserSchemaEnableValidationResponseJSON    `json:"-"`
}

// zoneAPIGatewayUserSchemaEnableValidationResponseJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayUserSchemaEnableValidationResponse]
type zoneAPIGatewayUserSchemaEnableValidationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaEnableValidationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaEnableValidationResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaEnableValidationResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaEnableValidationResponseSuccessTrue ZoneAPIGatewayUserSchemaEnableValidationResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaEnableValidationResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaEnableValidationResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaGetHostsResponse struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneAPIGatewayUserSchemaGetHostsResponseSuccess    `json:"success,required"`
	Result     []ZoneAPIGatewayUserSchemaGetHostsResponseResult   `json:"result"`
	ResultInfo ZoneAPIGatewayUserSchemaGetHostsResponseResultInfo `json:"result_info"`
	JSON       zoneAPIGatewayUserSchemaGetHostsResponseJSON       `json:"-"`
}

// zoneAPIGatewayUserSchemaGetHostsResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaGetHostsResponse]
type zoneAPIGatewayUserSchemaGetHostsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetHostsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetHostsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaGetHostsResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaGetHostsResponseSuccessTrue ZoneAPIGatewayUserSchemaGetHostsResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaGetHostsResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetHostsResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaGetHostsResponseResult struct {
	CreatedAt SchemasTimestamp `json:"created_at,required" format:"date-time"`
	// Hosts serving the schema, e.g zone.host.com
	Hosts []string `json:"hosts,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID.
	SchemaID SchemasUuid                                        `json:"schema_id,required"`
	JSON     zoneAPIGatewayUserSchemaGetHostsResponseResultJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetHostsResponseResultJSON contains the JSON metadata
// for the struct [ZoneAPIGatewayUserSchemaGetHostsResponseResult]
type zoneAPIGatewayUserSchemaGetHostsResponseResultJSON struct {
	CreatedAt   apijson.Field
	Hosts       apijson.Field
	Name        apijson.Field
	SchemaID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetHostsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetHostsResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaGetHostsResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                `json:"total_count"`
	JSON       zoneAPIGatewayUserSchemaGetHostsResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetHostsResponseResultInfoJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayUserSchemaGetHostsResponseResultInfo]
type zoneAPIGatewayUserSchemaGetHostsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetHostsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetHostsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaGetOperationsResponse struct {
	Errors   []MessagesAPIShieldItem                               `json:"errors,required"`
	Messages []MessagesAPIShieldItem                               `json:"messages,required"`
	Result   []ZoneAPIGatewayUserSchemaGetOperationsResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success    ZoneAPIGatewayUserSchemaGetOperationsResponseSuccess    `json:"success,required"`
	ResultInfo ZoneAPIGatewayUserSchemaGetOperationsResponseResultInfo `json:"result_info"`
	JSON       zoneAPIGatewayUserSchemaGetOperationsResponseJSON       `json:"-"`
}

// zoneAPIGatewayUserSchemaGetOperationsResponseJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaGetOperationsResponse]
type zoneAPIGatewayUserSchemaGetOperationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetOperationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetOperationsResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaGetOperationsResponseResult struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod `json:"method,required"`
	// This field can have the runtime type of [OperationFeatures].
	Features    interface{}      `json:"features"`
	LastUpdated shared.UnionTime `json:"last_updated" format:"date-time"`
	// UUID.
	OperationID shared.UnionString                                      `json:"operation_id"`
	JSON        zoneAPIGatewayUserSchemaGetOperationsResponseResultJSON `json:"-"`
	union       ZoneAPIGatewayUserSchemaGetOperationsResponseResultUnion
}

// zoneAPIGatewayUserSchemaGetOperationsResponseResultJSON contains the JSON
// metadata for the struct [ZoneAPIGatewayUserSchemaGetOperationsResponseResult]
type zoneAPIGatewayUserSchemaGetOperationsResponseResultJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	Features    apijson.Field
	LastUpdated apijson.Field
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r zoneAPIGatewayUserSchemaGetOperationsResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneAPIGatewayUserSchemaGetOperationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneAPIGatewayUserSchemaGetOperationsResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneAPIGatewayUserSchemaGetOperationsResponseResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [Operation], [BasicOperation].
func (r ZoneAPIGatewayUserSchemaGetOperationsResponseResult) AsUnion() ZoneAPIGatewayUserSchemaGetOperationsResponseResultUnion {
	return r.union
}

// Union satisfied by [Operation] or [BasicOperation].
type ZoneAPIGatewayUserSchemaGetOperationsResponseResultUnion interface {
	implementsZoneAPIGatewayUserSchemaGetOperationsResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAPIGatewayUserSchemaGetOperationsResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Operation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BasicOperation{}),
		},
	)
}

// The HTTP method used to access the endpoint.
type ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod string

const (
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodGet     ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "GET"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPost    ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "POST"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodHead    ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "HEAD"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodOptions ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "OPTIONS"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPut     ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "PUT"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodDelete  ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "DELETE"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodConnect ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "CONNECT"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPatch   ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "PATCH"
	ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodTrace   ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod = "TRACE"
)

func (r ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethod) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodGet, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPost, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodHead, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodOptions, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPut, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodDelete, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodConnect, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodPatch, ZoneAPIGatewayUserSchemaGetOperationsResponseResultMethodTrace:
		return true
	}
	return false
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaGetOperationsResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaGetOperationsResponseSuccessTrue ZoneAPIGatewayUserSchemaGetOperationsResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaGetOperationsResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetOperationsResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaGetOperationsResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                     `json:"total_count"`
	JSON       zoneAPIGatewayUserSchemaGetOperationsResponseResultInfoJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaGetOperationsResponseResultInfoJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayUserSchemaGetOperationsResponseResultInfo]
type zoneAPIGatewayUserSchemaGetOperationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaGetOperationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaGetOperationsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaUploadResponse struct {
	Errors   []MessagesAPIShieldItem                      `json:"errors,required"`
	Messages []MessagesAPIShieldItem                      `json:"messages,required"`
	Result   ZoneAPIGatewayUserSchemaUploadResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayUserSchemaUploadResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayUserSchemaUploadResponseJSON    `json:"-"`
}

// zoneAPIGatewayUserSchemaUploadResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayUserSchemaUploadResponse]
type zoneAPIGatewayUserSchemaUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaUploadResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaUploadResponseResult struct {
	Schema        PublicSchema                                              `json:"schema,required"`
	UploadDetails ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetails `json:"upload_details"`
	JSON          zoneAPIGatewayUserSchemaUploadResponseResultJSON          `json:"-"`
}

// zoneAPIGatewayUserSchemaUploadResponseResultJSON contains the JSON metadata for
// the struct [ZoneAPIGatewayUserSchemaUploadResponseResult]
type zoneAPIGatewayUserSchemaUploadResponseResultJSON struct {
	Schema        apijson.Field
	UploadDetails apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetails struct {
	// Diagnostic warning events that occurred during processing. These events are
	// non-critical errors found within the schema.
	Warnings []ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarning `json:"warnings"`
	JSON     zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsJSON      `json:"-"`
}

// zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsJSON contains the JSON
// metadata for the struct
// [ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetails]
type zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsJSON struct {
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarning struct {
	// Code that identifies the event that occurred.
	Code int64 `json:"code,required"`
	// JSONPath location(s) in the schema where these events were encountered. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string `json:"locations"`
	// Diagnostic message that describes the event.
	Message string                                                               `json:"message"`
	JSON    zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarningJSON `json:"-"`
}

// zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarningJSON contains
// the JSON metadata for the struct
// [ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarning]
type zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarningJSON struct {
	Code        apijson.Field
	Locations   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayUserSchemaUploadResponseResultUploadDetailsWarningJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneAPIGatewayUserSchemaUploadResponseSuccess bool

const (
	ZoneAPIGatewayUserSchemaUploadResponseSuccessTrue ZoneAPIGatewayUserSchemaUploadResponseSuccess = true
)

func (r ZoneAPIGatewayUserSchemaUploadResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaUploadResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaGetParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaGetParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayUserSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaListParams struct {
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaListParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayUserSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaEnableValidationParams struct {
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabled] `json:"validation_enabled"`
}

func (r ZoneAPIGatewayUserSchemaEnableValidationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Flag whether schema is enabled for validation.
type ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabled bool

const (
	ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabledTrue ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabled = true
)

func (r ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabled) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaEnableValidationParamsValidationEnabledTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaGetHostsParams struct {
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaGetHostsParams]'s query parameters
// as `url.Values`.
func (r ZoneAPIGatewayUserSchemaGetHostsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaGetOperationsParams struct {
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewayUserSchemaGetOperationsParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Filter results by whether operations exist in API Shield Endpoint Management or
	// not. `new` will just return operations from the schema that do not exist in API
	// Shield Endpoint Management. `existing` will just return operations from the
	// schema that already exist in API Shield Endpoint Management.
	OperationStatus param.Field[ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatus] `query:"operation_status"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ZoneAPIGatewayUserSchemaGetOperationsParams]'s query
// parameters as `url.Values`.
func (r ZoneAPIGatewayUserSchemaGetOperationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayUserSchemaGetOperationsParamsFeature string

const (
	ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureThresholds       ZoneAPIGatewayUserSchemaGetOperationsParamsFeature = "thresholds"
	ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureParameterSchemas ZoneAPIGatewayUserSchemaGetOperationsParamsFeature = "parameter_schemas"
	ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureSchemaInfo       ZoneAPIGatewayUserSchemaGetOperationsParamsFeature = "schema_info"
)

func (r ZoneAPIGatewayUserSchemaGetOperationsParamsFeature) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureThresholds, ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureParameterSchemas, ZoneAPIGatewayUserSchemaGetOperationsParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Filter results by whether operations exist in API Shield Endpoint Management or
// not. `new` will just return operations from the schema that do not exist in API
// Shield Endpoint Management. `existing` will just return operations from the
// schema that already exist in API Shield Endpoint Management.
type ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatus string

const (
	ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatusNew      ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatus = "new"
	ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatusExisting ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatus = "existing"
)

func (r ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatus) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatusNew, ZoneAPIGatewayUserSchemaGetOperationsParamsOperationStatusExisting:
		return true
	}
	return false
}

type ZoneAPIGatewayUserSchemaUploadParams struct {
	// Schema file bytes
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Kind of schema
	Kind param.Field[APIShieldKind] `json:"kind,required"`
	// Name of the schema
	Name param.Field[string] `json:"name"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[ZoneAPIGatewayUserSchemaUploadParamsValidationEnabled] `json:"validation_enabled"`
}

func (r ZoneAPIGatewayUserSchemaUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Flag whether schema is enabled for validation.
type ZoneAPIGatewayUserSchemaUploadParamsValidationEnabled string

const (
	ZoneAPIGatewayUserSchemaUploadParamsValidationEnabledTrue  ZoneAPIGatewayUserSchemaUploadParamsValidationEnabled = "true"
	ZoneAPIGatewayUserSchemaUploadParamsValidationEnabledFalse ZoneAPIGatewayUserSchemaUploadParamsValidationEnabled = "false"
)

func (r ZoneAPIGatewayUserSchemaUploadParamsValidationEnabled) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayUserSchemaUploadParamsValidationEnabledTrue, ZoneAPIGatewayUserSchemaUploadParamsValidationEnabledFalse:
		return true
	}
	return false
}
