// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountMagicSiteAppConfigService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMagicSiteAppConfigService] method instead.
type AccountMagicSiteAppConfigService struct {
	Options []option.RequestOption
}

// NewAccountMagicSiteAppConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountMagicSiteAppConfigService(opts ...option.RequestOption) (r *AccountMagicSiteAppConfigService) {
	r = &AccountMagicSiteAppConfigService{}
	r.Options = opts
	return
}

// Creates a new App Config for a site
func (r *AccountMagicSiteAppConfigService) New(ctx context.Context, accountID string, siteID string, body AccountMagicSiteAppConfigNewParams, opts ...option.RequestOption) (res *MagicAppConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an App Config for a site
func (r *AccountMagicSiteAppConfigService) Update(ctx context.Context, accountID string, siteID string, appConfigID string, body AccountMagicSiteAppConfigUpdateParams, opts ...option.RequestOption) (res *MagicAppConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if appConfigID == "" {
		err = errors.New("missing required app_config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs/%s", accountID, siteID, appConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists App Configs associated with a site.
func (r *AccountMagicSiteAppConfigService) List(ctx context.Context, accountID string, siteID string, opts ...option.RequestOption) (res *AccountMagicSiteAppConfigListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs", accountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes specific App Config associated with a site.
func (r *AccountMagicSiteAppConfigService) Delete(ctx context.Context, accountID string, siteID string, appConfigID string, opts ...option.RequestOption) (res *MagicAppConfigSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if appConfigID == "" {
		err = errors.New("missing required app_config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/app_configs/%s", accountID, siteID, appConfigID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Traffic decision configuration for an app.
type MagicAppConfig struct {
	// Identifier
	ID string `json:"id"`
	// Magic account app ID.
	AccountAppID string `json:"account_app_id"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string             `json:"site_id"`
	JSON   magicAppConfigJSON `json:"-"`
	union  MagicAppConfigUnion
}

// magicAppConfigJSON contains the JSON metadata for the struct [MagicAppConfig]
type magicAppConfigJSON struct {
	ID           apijson.Field
	AccountAppID apijson.Field
	Breakout     apijson.Field
	ManagedAppID apijson.Field
	Priority     apijson.Field
	SiteID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r magicAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r *MagicAppConfig) UnmarshalJSON(data []byte) (err error) {
	*r = MagicAppConfig{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [MagicAppConfigUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [MagicAppConfigAccountApp],
// [MagicAppConfigManagedApp].
func (r MagicAppConfig) AsUnion() MagicAppConfigUnion {
	return r.union
}

// Traffic decision configuration for an app.
//
// Union satisfied by [MagicAppConfigAccountApp] or [MagicAppConfigManagedApp].
type MagicAppConfigUnion interface {
	implementsMagicAppConfig()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*MagicAppConfigUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicAppConfigAccountApp{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(MagicAppConfigManagedApp{}),
		},
	)
}

type MagicAppConfigAccountApp struct {
	// Magic account app ID.
	AccountAppID string `json:"account_app_id,required"`
	// Identifier
	ID string `json:"id"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string                       `json:"site_id"`
	JSON   magicAppConfigAccountAppJSON `json:"-"`
}

// magicAppConfigAccountAppJSON contains the JSON metadata for the struct
// [MagicAppConfigAccountApp]
type magicAppConfigAccountAppJSON struct {
	AccountAppID apijson.Field
	ID           apijson.Field
	Breakout     apijson.Field
	Priority     apijson.Field
	SiteID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicAppConfigAccountApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigAccountAppJSON) RawJSON() string {
	return r.raw
}

func (r MagicAppConfigAccountApp) implementsMagicAppConfig() {}

type MagicAppConfigManagedApp struct {
	// Managed app ID.
	ManagedAppID string `json:"managed_app_id,required"`
	// Identifier
	ID string `json:"id"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout bool `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority int64 `json:"priority"`
	// Identifier
	SiteID string                       `json:"site_id"`
	JSON   magicAppConfigManagedAppJSON `json:"-"`
}

// magicAppConfigManagedAppJSON contains the JSON metadata for the struct
// [MagicAppConfigManagedApp]
type magicAppConfigManagedAppJSON struct {
	ManagedAppID apijson.Field
	ID           apijson.Field
	Breakout     apijson.Field
	Priority     apijson.Field
	SiteID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MagicAppConfigManagedApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigManagedAppJSON) RawJSON() string {
	return r.raw
}

func (r MagicAppConfigManagedApp) implementsMagicAppConfig() {}

type MagicAppConfigSingleResponse struct {
	Errors   []MagicAppConfigSingleResponseError   `json:"errors,required"`
	Messages []MagicAppConfigSingleResponseMessage `json:"messages,required"`
	// Traffic decision configuration for an app.
	Result MagicAppConfig `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicAppConfigSingleResponseSuccess `json:"success,required"`
	JSON    magicAppConfigSingleResponseJSON    `json:"-"`
}

// magicAppConfigSingleResponseJSON contains the JSON metadata for the struct
// [MagicAppConfigSingleResponse]
type magicAppConfigSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseJSON) RawJSON() string {
	return r.raw
}

type MagicAppConfigSingleResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           MagicAppConfigSingleResponseErrorsSource `json:"source"`
	JSON             magicAppConfigSingleResponseErrorJSON    `json:"-"`
}

// magicAppConfigSingleResponseErrorJSON contains the JSON metadata for the struct
// [MagicAppConfigSingleResponseError]
type magicAppConfigSingleResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseErrorJSON) RawJSON() string {
	return r.raw
}

type MagicAppConfigSingleResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    magicAppConfigSingleResponseErrorsSourceJSON `json:"-"`
}

// magicAppConfigSingleResponseErrorsSourceJSON contains the JSON metadata for the
// struct [MagicAppConfigSingleResponseErrorsSource]
type magicAppConfigSingleResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MagicAppConfigSingleResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           MagicAppConfigSingleResponseMessagesSource `json:"source"`
	JSON             magicAppConfigSingleResponseMessageJSON    `json:"-"`
}

// magicAppConfigSingleResponseMessageJSON contains the JSON metadata for the
// struct [MagicAppConfigSingleResponseMessage]
type magicAppConfigSingleResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseMessageJSON) RawJSON() string {
	return r.raw
}

type MagicAppConfigSingleResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    magicAppConfigSingleResponseMessagesSourceJSON `json:"-"`
}

// magicAppConfigSingleResponseMessagesSourceJSON contains the JSON metadata for
// the struct [MagicAppConfigSingleResponseMessagesSource]
type magicAppConfigSingleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicAppConfigSingleResponseSuccess bool

const (
	MagicAppConfigSingleResponseSuccessTrue MagicAppConfigSingleResponseSuccess = true
)

func (r MagicAppConfigSingleResponseSuccess) IsKnown() bool {
	switch r {
	case MagicAppConfigSingleResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteAppConfigListResponse struct {
	Errors   []AccountMagicSiteAppConfigListResponseError   `json:"errors,required"`
	Messages []AccountMagicSiteAppConfigListResponseMessage `json:"messages,required"`
	Result   []MagicAppConfig                               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountMagicSiteAppConfigListResponseSuccess `json:"success,required"`
	JSON    accountMagicSiteAppConfigListResponseJSON    `json:"-"`
}

// accountMagicSiteAppConfigListResponseJSON contains the JSON metadata for the
// struct [AccountMagicSiteAppConfigListResponse]
type accountMagicSiteAppConfigListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigListResponseError struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccountMagicSiteAppConfigListResponseErrorsSource `json:"source"`
	JSON             accountMagicSiteAppConfigListResponseErrorJSON    `json:"-"`
}

// accountMagicSiteAppConfigListResponseErrorJSON contains the JSON metadata for
// the struct [AccountMagicSiteAppConfigListResponseError]
type accountMagicSiteAppConfigListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigListResponseErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accountMagicSiteAppConfigListResponseErrorsSourceJSON `json:"-"`
}

// accountMagicSiteAppConfigListResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountMagicSiteAppConfigListResponseErrorsSource]
type accountMagicSiteAppConfigListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigListResponseMessage struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccountMagicSiteAppConfigListResponseMessagesSource `json:"source"`
	JSON             accountMagicSiteAppConfigListResponseMessageJSON    `json:"-"`
}

// accountMagicSiteAppConfigListResponseMessageJSON contains the JSON metadata for
// the struct [AccountMagicSiteAppConfigListResponseMessage]
type accountMagicSiteAppConfigListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigListResponseMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accountMagicSiteAppConfigListResponseMessagesSourceJSON `json:"-"`
}

// accountMagicSiteAppConfigListResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountMagicSiteAppConfigListResponseMessagesSource]
type accountMagicSiteAppConfigListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountMagicSiteAppConfigListResponseSuccess bool

const (
	AccountMagicSiteAppConfigListResponseSuccessTrue AccountMagicSiteAppConfigListResponseSuccess = true
)

func (r AccountMagicSiteAppConfigListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMagicSiteAppConfigListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMagicSiteAppConfigNewParams struct {
	Body AccountMagicSiteAppConfigNewParamsBodyUnion `json:"body,required"`
}

func (r AccountMagicSiteAppConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicSiteAppConfigNewParamsBody struct {
	// Magic account app ID.
	AccountAppID param.Field[string] `json:"account_app_id"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Managed app ID.
	ManagedAppID param.Field[string] `json:"managed_app_id"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r AccountMagicSiteAppConfigNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMagicSiteAppConfigNewParamsBody) implementsAccountMagicSiteAppConfigNewParamsBodyUnion() {
}

// Satisfied by [AccountMagicSiteAppConfigNewParamsBodyAccountApp],
// [AccountMagicSiteAppConfigNewParamsBodyManagedApp],
// [AccountMagicSiteAppConfigNewParamsBody].
type AccountMagicSiteAppConfigNewParamsBodyUnion interface {
	implementsAccountMagicSiteAppConfigNewParamsBodyUnion()
}

type AccountMagicSiteAppConfigNewParamsBodyAccountApp struct {
	// Magic account app ID.
	AccountAppID param.Field[string] `json:"account_app_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r AccountMagicSiteAppConfigNewParamsBodyAccountApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMagicSiteAppConfigNewParamsBodyAccountApp) implementsAccountMagicSiteAppConfigNewParamsBodyUnion() {
}

type AccountMagicSiteAppConfigNewParamsBodyManagedApp struct {
	// Managed app ID.
	ManagedAppID param.Field[string] `json:"managed_app_id,required"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r AccountMagicSiteAppConfigNewParamsBodyManagedApp) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMagicSiteAppConfigNewParamsBodyManagedApp) implementsAccountMagicSiteAppConfigNewParamsBodyUnion() {
}

type AccountMagicSiteAppConfigUpdateParams struct {
	// Magic account app ID.
	AccountAppID param.Field[string] `json:"account_app_id"`
	// Whether to breakout traffic to the app's endpoints directly. Null preserves
	// default behavior.
	Breakout param.Field[bool] `json:"breakout"`
	// Managed app ID.
	ManagedAppID param.Field[string] `json:"managed_app_id"`
	// Priority of traffic. 0 is default, anything greater is prioritized. (Currently
	// only 0 and 1 are supported)
	Priority param.Field[int64] `json:"priority"`
}

func (r AccountMagicSiteAppConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
