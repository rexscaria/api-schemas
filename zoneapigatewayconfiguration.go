// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneAPIGatewayConfigurationService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAPIGatewayConfigurationService] method instead.
type ZoneAPIGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewZoneAPIGatewayConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAPIGatewayConfigurationService(opts ...option.RequestOption) (r *ZoneAPIGatewayConfigurationService) {
	r = &ZoneAPIGatewayConfigurationService{}
	r.Options = opts
	return
}

// Retrieve information about specific configuration properties
func (r *ZoneAPIGatewayConfigurationService) Get(ctx context.Context, zoneID string, query ZoneAPIGatewayConfigurationGetParams, opts ...option.RequestOption) (res *ZoneAPIGatewayConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Set configuration properties
func (r *ZoneAPIGatewayConfigurationService) Update(ctx context.Context, zoneID string, body ZoneAPIGatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *APIResponseAPIShield, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/configuration", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type APIResponseAPIShield struct {
	Errors   []MessagesAPIShieldItem `json:"errors,required"`
	Messages []MessagesAPIShieldItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseAPIShieldSuccess `json:"success,required"`
	JSON    apiResponseAPIShieldJSON    `json:"-"`
}

// apiResponseAPIShieldJSON contains the JSON metadata for the struct
// [APIResponseAPIShield]
type apiResponseAPIShieldJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAPIShield) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAPIShieldJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseAPIShieldSuccess bool

const (
	APIResponseAPIShieldSuccessTrue APIResponseAPIShieldSuccess = true
)

func (r APIResponseAPIShieldSuccess) IsKnown() bool {
	switch r {
	case APIResponseAPIShieldSuccessTrue:
		return true
	}
	return false
}

type Configuration struct {
	AuthIDCharacteristics []ConfigurationAuthIDCharacteristic `json:"auth_id_characteristics,required"`
	JSON                  configurationJSON                   `json:"-"`
}

// configurationJSON contains the JSON metadata for the struct [Configuration]
type configurationJSON struct {
	AuthIDCharacteristics apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationJSON) RawJSON() string {
	return r.raw
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type  ConfigurationAuthIDCharacteristicsType `json:"type,required"`
	JSON  configurationAuthIDCharacteristicJSON  `json:"-"`
	union ConfigurationAuthIDCharacteristicsUnion
}

// configurationAuthIDCharacteristicJSON contains the JSON metadata for the struct
// [ConfigurationAuthIDCharacteristic]
type configurationAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r configurationAuthIDCharacteristicJSON) RawJSON() string {
	return r.raw
}

func (r *ConfigurationAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	*r = ConfigurationAuthIDCharacteristic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConfigurationAuthIDCharacteristicsUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic],
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim].
func (r ConfigurationAuthIDCharacteristic) AsUnion() ConfigurationAuthIDCharacteristicsUnion {
	return r.union
}

// Auth ID Characteristic
//
// Union satisfied by
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic] or
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim].
type ConfigurationAuthIDCharacteristicsUnion interface {
	implementsConfigurationAuthIDCharacteristic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConfigurationAuthIDCharacteristicsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim{}),
		},
	)
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType `json:"type,required"`
	JSON configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON `json:"-"`
}

// configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON contains the
// JSON metadata for the struct
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic]
type configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristic) implementsConfigurationAuthIDCharacteristic() {
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType string

const (
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeHeader ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType = "header"
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeCookie ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType = "cookie"
)

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeHeader, ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicTypeCookie:
		return true
	}
	return false
}

// Auth ID Characteristic extracted from JWT Token Claims
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim struct {
	// Claim location expressed as `$(token_config_id):$(json_path)`, where
	// `token_config_id` is the ID of the token configuration used in validating the
	// JWT, and `json_path` is a RFC 9535 JSONPath
	// (https://goessner.net/articles/JsonPath/,
	// https://www.rfc-editor.org/rfc/rfc9535.html). The JSONPath expression may be in
	// dot or bracket notation, may only specify literal keys or array indexes, and
	// must return a singleton value, which will be interpreted as a string.
	Name string `json:"name,required"`
	// The type of characteristic.
	Type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType `json:"type,required"`
	JSON configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON `json:"-"`
}

// configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON
// contains the JSON metadata for the struct
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim]
type configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimJSON) RawJSON() string {
	return r.raw
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaim) implementsConfigurationAuthIDCharacteristic() {
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType string

const (
	ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimTypeJwt ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType = "jwt"
)

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimTypeJwt:
		return true
	}
	return false
}

// The type of characteristic.
type ConfigurationAuthIDCharacteristicsType string

const (
	ConfigurationAuthIDCharacteristicsTypeHeader ConfigurationAuthIDCharacteristicsType = "header"
	ConfigurationAuthIDCharacteristicsTypeCookie ConfigurationAuthIDCharacteristicsType = "cookie"
	ConfigurationAuthIDCharacteristicsTypeJwt    ConfigurationAuthIDCharacteristicsType = "jwt"
)

func (r ConfigurationAuthIDCharacteristicsType) IsKnown() bool {
	switch r {
	case ConfigurationAuthIDCharacteristicsTypeHeader, ConfigurationAuthIDCharacteristicsTypeCookie, ConfigurationAuthIDCharacteristicsTypeJwt:
		return true
	}
	return false
}

type ConfigurationParam struct {
	AuthIDCharacteristics param.Field[[]ConfigurationAuthIDCharacteristicsUnionParam] `json:"auth_id_characteristics,required"`
}

func (r ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicParam struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicParam) implementsConfigurationAuthIDCharacteristicsUnionParam() {
}

// Auth ID Characteristic
//
// Satisfied by
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam],
// [ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam],
// [ConfigurationAuthIDCharacteristicParam].
type ConfigurationAuthIDCharacteristicsUnionParam interface {
	implementsConfigurationAuthIDCharacteristicsUnionParam()
}

// Auth ID Characteristic
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam struct {
	// The name of the characteristic field, i.e., the header or cookie name.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicParam) implementsConfigurationAuthIDCharacteristicsUnionParam() {
}

// Auth ID Characteristic extracted from JWT Token Claims
type ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam struct {
	// Claim location expressed as `$(token_config_id):$(json_path)`, where
	// `token_config_id` is the ID of the token configuration used in validating the
	// JWT, and `json_path` is a RFC 9535 JSONPath
	// (https://goessner.net/articles/JsonPath/,
	// https://www.rfc-editor.org/rfc/rfc9535.html). The JSONPath expression may be in
	// dot or bracket notation, may only specify literal keys or array indexes, and
	// must return a singleton value, which will be interpreted as a string.
	Name param.Field[string] `json:"name,required"`
	// The type of characteristic.
	Type param.Field[ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimType] `json:"type,required"`
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurationAuthIDCharacteristicsAPIShieldAuthIDCharacteristicJwtClaimParam) implementsConfigurationAuthIDCharacteristicsUnionParam() {
}

type MessagesAPIShieldItem struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    messagesAPIShieldItemJSON `json:"-"`
}

// messagesAPIShieldItemJSON contains the JSON metadata for the struct
// [MessagesAPIShieldItem]
type messagesAPIShieldItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesAPIShieldItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesAPIShieldItemJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayConfigurationGetResponse struct {
	Result Configuration                              `json:"result,required"`
	JSON   zoneAPIGatewayConfigurationGetResponseJSON `json:"-"`
	APIResponseAPIShield
}

// zoneAPIGatewayConfigurationGetResponseJSON contains the JSON metadata for the
// struct [ZoneAPIGatewayConfigurationGetResponse]
type zoneAPIGatewayConfigurationGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAPIGatewayConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAPIGatewayConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAPIGatewayConfigurationGetParams struct {
	// Requests information about certain properties.
	Properties param.Field[[]ZoneAPIGatewayConfigurationGetParamsProperty] `query:"properties"`
}

// URLQuery serializes [ZoneAPIGatewayConfigurationGetParams]'s query parameters as
// `url.Values`.
func (r ZoneAPIGatewayConfigurationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneAPIGatewayConfigurationGetParamsProperty string

const (
	ZoneAPIGatewayConfigurationGetParamsPropertyAuthIDCharacteristics ZoneAPIGatewayConfigurationGetParamsProperty = "auth_id_characteristics"
)

func (r ZoneAPIGatewayConfigurationGetParamsProperty) IsKnown() bool {
	switch r {
	case ZoneAPIGatewayConfigurationGetParamsPropertyAuthIDCharacteristics:
		return true
	}
	return false
}

type ZoneAPIGatewayConfigurationUpdateParams struct {
	Configuration ConfigurationParam `json:"configuration,required"`
}

func (r ZoneAPIGatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Configuration)
}
