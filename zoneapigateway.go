// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneAPIGatewayService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayService] method instead.
type ZoneAPIGatewayService struct {
	Options            []option.RequestOption
	Configuration      *ZoneAPIGatewayConfigurationService
	Discovery          *ZoneAPIGatewayDiscoveryService
	ExpressionTemplate *ZoneAPIGatewayExpressionTemplateService
	Operations         *ZoneAPIGatewayOperationService
	Settings           *ZoneAPIGatewaySettingService
	UserSchemas        *ZoneAPIGatewayUserSchemaService
}

// NewZoneAPIGatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAPIGatewayService(opts ...option.RequestOption) (r *ZoneAPIGatewayService) {
	r = &ZoneAPIGatewayService{}
	r.Options = opts
	r.Configuration = NewZoneAPIGatewayConfigurationService(opts...)
	r.Discovery = NewZoneAPIGatewayDiscoveryService(opts...)
	r.ExpressionTemplate = NewZoneAPIGatewayExpressionTemplateService(opts...)
	r.Operations = NewZoneAPIGatewayOperationService(opts...)
	r.Settings = NewZoneAPIGatewaySettingService(opts...)
	r.UserSchemas = NewZoneAPIGatewayUserSchemaService(opts...)
	return
}

// Retrieve operations and features as OpenAPI schemas
func (r *ZoneAPIGatewayService) GetSchemas(ctx context.Context, zoneID string, query ZoneAPIGatewayGetSchemasParams, opts ...option.RequestOption) (res *ZoneAPIGatewayGetSchemasResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/schemas", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAPIGatewayGetSchemasResponse struct {
	Result ZoneAPIGatewayGetSchemasResponseResult `json:"result,required"`
	JSON   zoneAPIGatewayGetSchemasResponseJSON   `json:"-"`
	APIResponseAPIShield
}

// zoneAPIGatewayGetSchemasResponseJSON contains the JSON metadata for the struct
// [ZoneAPIGatewayGetSchemasResponse]
type zoneAPIGatewayGetSchemasResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayGetSchemasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayGetSchemasResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayGetSchemasResponseResult struct {
	Schemas   []interface{}                              `json:"schemas"`
	Timestamp string                                     `json:"timestamp"`
	JSON      zoneAPIGatewayGetSchemasResponseResultJSON `json:"-"`
}

// zoneAPIGatewayGetSchemasResponseResultJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayGetSchemasResponseResult]
type zoneAPIGatewayGetSchemasResponseResultJSON struct {
	Schemas     apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayGetSchemasResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayGetSchemasResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayGetSchemasParams struct {
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]ZoneAPIGatewayGetSchemasParamsFeature] `query:"feature"`
	// Receive schema only for the given host(s).
	Host param.Field[[]string] `query:"host"`
}

// URLQuery serializes [ZoneAPIGatewayGetSchemasParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayGetSchemasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayGetSchemasParamsFeature string

const (
	ZoneAPIGatewayGetSchemasParamsFeatureThresholds       ZoneAPIGatewayGetSchemasParamsFeature = "thresholds"
	ZoneAPIGatewayGetSchemasParamsFeatureParameterSchemas ZoneAPIGatewayGetSchemasParamsFeature = "parameter_schemas"
	ZoneAPIGatewayGetSchemasParamsFeatureSchemaInfo       ZoneAPIGatewayGetSchemasParamsFeature = "schema_info"
)

func (r ZoneAPIGatewayGetSchemasParamsFeature) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayGetSchemasParamsFeatureThresholds, ZoneAPIGatewayGetSchemasParamsFeatureParameterSchemas, ZoneAPIGatewayGetSchemasParamsFeatureSchemaInfo:
		return true
	}
	return false
}
