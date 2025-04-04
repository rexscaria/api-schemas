// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneAPIGatewaySettingSchemaValidationService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewaySettingSchemaValidationService] method instead.
type ZoneAPIGatewaySettingSchemaValidationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewaySettingSchemaValidationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneAPIGatewaySettingSchemaValidationService(opts ...option.RequestOption) (r *ZoneAPIGatewaySettingSchemaValidationService) {
	r = &ZoneAPIGatewaySettingSchemaValidationService{}
	r.Options = opts
	return
}

// Retrieves zone level schema validation settings currently set on the zone
func (r *ZoneAPIGatewaySettingSchemaValidationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates zone level schema validation settings on the zone
func (r *ZoneAPIGatewaySettingSchemaValidationService) Update(ctx context.Context, zoneID string, body ZoneAPIGatewaySettingSchemaValidationUpdateParams, opts ...option.RequestOption) (res *ZoneSchemaValidationSettings, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/settings/schema_validation", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The default mitigation action used when there is no mitigation action defined on
// the operation
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
//
// A special value of of `none` will skip running schema validation entirely for
// the request when there is no mitigation action defined on the operation
type DefaultMitigationAction string

const (
	DefaultMitigationActionNone  DefaultMitigationAction = "none"
	DefaultMitigationActionLog   DefaultMitigationAction = "log"
	DefaultMitigationActionBlock DefaultMitigationAction = "block"
)

func (r DefaultMitigationAction) IsKnown() bool {
	switch r {
	case DefaultMitigationActionNone, DefaultMitigationActionLog, DefaultMitigationActionBlock:
		return true
	}
	return false
}

type ZoneSchemaValidationSettings struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction DefaultMitigationAction `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	ValidationOverrideMitigationAction ZoneSchemaValidationSettingsValidationOverrideMitigationAction `json:"validation_override_mitigation_action,nullable"`
	JSON                               zoneSchemaValidationSettingsJSON                               `json:"-"`
}

// zoneSchemaValidationSettingsJSON contains the JSON metadata for the struct
// [ZoneSchemaValidationSettings]
type zoneSchemaValidationSettingsJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *ZoneSchemaValidationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSchemaValidationSettingsJSON) RawJSON() string {
	return r.raw
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
type ZoneSchemaValidationSettingsValidationOverrideMitigationAction string

const (
	ZoneSchemaValidationSettingsValidationOverrideMitigationActionNone ZoneSchemaValidationSettingsValidationOverrideMitigationAction = "none"
)

func (r ZoneSchemaValidationSettingsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case ZoneSchemaValidationSettingsValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type ZoneAPIGatewaySettingSchemaValidationUpdateParams struct {
	// The default mitigation action used when there is no mitigation action defined on
	// the operation
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	//
	// A special value of of `none` will skip running schema validation entirely for
	// the request when there is no mitigation action defined on the operation
	ValidationDefaultMitigationAction param.Field[DefaultMitigationAction] `json:"validation_default_mitigation_action,required"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `none` will skip running schema validation entirely for the request
	// - `null` indicates that no override is in place
	//
	// To clear any override, use the special value `disable_override` or `null`
	ValidationOverrideMitigationAction param.Field[ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r ZoneAPIGatewaySettingSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `none` will skip running schema validation entirely for the request
// - `null` indicates that no override is in place
//
// To clear any override, use the special value `disable_override` or `null`
type ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction string

const (
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone            ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "none"
	ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction = "disable_override"
)

func (r ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionNone, ZoneAPIGatewaySettingSchemaValidationUpdateParamsValidationOverrideMitigationActionDisableOverride:
		return true
	}
	return false
}
