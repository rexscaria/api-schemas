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

// AccountGatewayLoggingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayLoggingService] method instead.
type AccountGatewayLoggingService struct {
	Options []option.RequestOption
}

// NewAccountGatewayLoggingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayLoggingService(opts ...option.RequestOption) (r *AccountGatewayLoggingService) {
	r = &AccountGatewayLoggingService{}
	r.Options = opts
	return
}

// Fetches the current logging settings for Zero Trust account.
func (r *AccountGatewayLoggingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *GatewayAccountLoggingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates logging settings for the current Zero Trust account.
func (r *AccountGatewayLoggingService) Update(ctx context.Context, accountID string, body AccountGatewayLoggingUpdateParams, opts ...option.RequestOption) (res *GatewayAccountLoggingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountLogOptions struct {
	// Log all requests to this service.
	LogAll bool `json:"log_all"`
	// Log only blocking requests to this service.
	LogBlocks bool                  `json:"log_blocks"`
	JSON      accountLogOptionsJSON `json:"-"`
}

// accountLogOptionsJSON contains the JSON metadata for the struct
// [AccountLogOptions]
type accountLogOptionsJSON struct {
	LogAll      apijson.Field
	LogBlocks   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogOptionsJSON) RawJSON() string {
	return r.raw
}

type AccountLogOptionsParam struct {
	// Log all requests to this service.
	LogAll param.Field[bool] `json:"log_all"`
	// Log only blocking requests to this service.
	LogBlocks param.Field[bool] `json:"log_blocks"`
}

func (r AccountLogOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayAccountLoggingSettings struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayAccountLoggingSettingsSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayAccountLoggingSettingsJSON               `json:"-"`
}

// gatewayAccountLoggingSettingsJSON contains the JSON metadata for the struct
// [GatewayAccountLoggingSettings]
type gatewayAccountLoggingSettingsJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountLoggingSettingsJSON) RawJSON() string {
	return r.raw
}

// Logging settings by rule type.
type GatewayAccountLoggingSettingsSettingsByRuleType struct {
	DNS  AccountLogOptions                                   `json:"dns"`
	HTTP AccountLogOptions                                   `json:"http"`
	L4   AccountLogOptions                                   `json:"l4"`
	JSON gatewayAccountLoggingSettingsSettingsByRuleTypeJSON `json:"-"`
}

// gatewayAccountLoggingSettingsSettingsByRuleTypeJSON contains the JSON metadata
// for the struct [GatewayAccountLoggingSettingsSettingsByRuleType]
type gatewayAccountLoggingSettingsSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountLoggingSettingsSettingsByRuleTypeJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountLoggingSettingsParam struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[GatewayAccountLoggingSettingsSettingsByRuleTypeParam] `json:"settings_by_rule_type"`
}

func (r GatewayAccountLoggingSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type GatewayAccountLoggingSettingsSettingsByRuleTypeParam struct {
	DNS  param.Field[AccountLogOptionsParam] `json:"dns"`
	HTTP param.Field[AccountLogOptionsParam] `json:"http"`
	L4   param.Field[AccountLogOptionsParam] `json:"l4"`
}

func (r GatewayAccountLoggingSettingsSettingsByRuleTypeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayAccountLoggingSettingsResponse struct {
	Result GatewayAccountLoggingSettings             `json:"result"`
	JSON   gatewayAccountLoggingSettingsResponseJSON `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// gatewayAccountLoggingSettingsResponseJSON contains the JSON metadata for the
// struct [GatewayAccountLoggingSettingsResponse]
type gatewayAccountLoggingSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountLoggingSettingsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayLoggingUpdateParams struct {
	GatewayAccountLoggingSettings GatewayAccountLoggingSettingsParam `json:"gateway_account_logging_settings,required"`
}

func (r AccountGatewayLoggingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GatewayAccountLoggingSettings)
}
