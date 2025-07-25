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

// ZoneAPIGatewayOperationSchemaValidationService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayOperationSchemaValidationService] method instead.
type ZoneAPIGatewayOperationSchemaValidationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayOperationSchemaValidationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneAPIGatewayOperationSchemaValidationService(opts ...option.RequestOption) (r *ZoneAPIGatewayOperationSchemaValidationService) {
	r = &ZoneAPIGatewayOperationSchemaValidationService{}
	r.Options = opts
	return
}

// Retrieves operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayOperationSchemaValidationService) Get(ctx context.Context, zoneID string, operationID string, opts ...option.RequestOption) (res *SchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayOperationSchemaValidationService) Update(ctx context.Context, zoneID string, operationID string, body ZoneAPIGatewayOperationSchemaValidationUpdateParams, opts ...option.RequestOption) (res *SchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", zoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates multiple operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *ZoneAPIGatewayOperationSchemaValidationService) UpdateMultiple(ctx context.Context, zoneID string, body ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParams, opts ...option.RequestOption) (res *ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type SchemaValidationSettings struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction SchemaValidationSettingsMitigationAction `json:"mitigation_action,nullable"`
	// UUID.
	OperationID SchemasUuid                  `json:"operation_id"`
	JSON        schemaValidationSettingsJSON `json:"-"`
}

// schemaValidationSettingsJSON contains the JSON metadata for the struct
// [SchemaValidationSettings]
type schemaValidationSettingsJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SchemaValidationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaValidationSettingsJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type SchemaValidationSettingsMitigationAction string

const (
	SchemaValidationSettingsMitigationActionLog   SchemaValidationSettingsMitigationAction = "log"
	SchemaValidationSettingsMitigationActionBlock SchemaValidationSettingsMitigationAction = "block"
	SchemaValidationSettingsMitigationActionNone  SchemaValidationSettingsMitigationAction = "none"
)

func (r SchemaValidationSettingsMitigationAction) IsKnown() bool {
	switch r {
	case SchemaValidationSettingsMitigationActionLog, SchemaValidationSettingsMitigationActionBlock, SchemaValidationSettingsMitigationActionNone:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponse struct {
	Errors   []MessagesAPIShieldItem                                                        `json:"errors,required"`
	Messages []MessagesAPIShieldItem                                                        `json:"messages,required"`
	Result   map[string]ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResult `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccess `json:"success,required"`
	JSON    zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseJSON    `json:"-"`
}

// zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseJSON contains the
// JSON metadata for the struct
// [ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponse]
type zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseJSON) RawJSON() string {
	return r.raw
}

// Operation ID to mitigation action mappings
type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResult struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction `json:"mitigation_action,nullable"`
	JSON             zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultJSON             `json:"-"`
}

// zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResult]
type zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultJSON struct {
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction string

const (
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionLog   ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction = "log"
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionBlock ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction = "block"
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionNone  ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction = "none"
)

func (r ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationAction) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionLog, ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionBlock, ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseResultMitigationActionNone:
		return true
	}
	return false
}

// Whether the API call was successful.
type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccess bool

const (
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccessTrue ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccess = true
)

func (r ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationSchemaValidationUpdateMultipleResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationSchemaValidationUpdateParams struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction] `json:"mitigation_action"`
}

func (r ZoneAPIGatewayOperationSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction string

const (
	ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionLog   ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction = "log"
	ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionBlock ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction = "block"
	ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionNone  ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction = "none"
)

func (r ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationAction) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionLog, ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionBlock, ZoneAPIGatewayOperationSchemaValidationUpdateParamsMitigationActionNone:
		return true
	}
	return false
}

type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParams struct {
	Body map[string]ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBody `json:"body,required"`
}

func (r ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Operation ID to mitigation action mappings
type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBody struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction] `json:"mitigation_action"`
}

func (r ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction string

const (
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionLog   ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction = "log"
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionBlock ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction = "block"
	ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionNone  ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction = "none"
)

func (r ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationAction) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionLog, ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionBlock, ZoneAPIGatewayOperationSchemaValidationUpdateMultipleParamsBodyMitigationActionNone:
		return true
	}
	return false
}
