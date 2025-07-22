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
	JSON magicAppConfigJSON `json:"-"`
}

// magicAppConfigJSON contains the JSON metadata for the struct [MagicAppConfig]
type magicAppConfigJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigJSON) RawJSON() string {
	return r.raw
}

type MagicAppConfigSingleResponse struct {
	// Traffic decision configuration for an app.
	Result MagicAppConfig                   `json:"result"`
	JSON   magicAppConfigSingleResponseJSON `json:"-"`
	MagicAppsResponseObject
}

// magicAppConfigSingleResponseJSON contains the JSON metadata for the struct
// [MagicAppConfigSingleResponse]
type magicAppConfigSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicAppConfigSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicAppConfigSingleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigListResponse struct {
	Result []MagicAppConfig                          `json:"result"`
	JSON   accountMagicSiteAppConfigListResponseJSON `json:"-"`
	MagicAppsResponseObject
}

// accountMagicSiteAppConfigListResponseJSON contains the JSON metadata for the
// struct [AccountMagicSiteAppConfigListResponse]
type accountMagicSiteAppConfigListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMagicSiteAppConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMagicSiteAppConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMagicSiteAppConfigNewParams struct {
	Body AccountMagicSiteAppConfigNewParamsBody `json:"body,required"`
}

func (r AccountMagicSiteAppConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMagicSiteAppConfigNewParamsBody struct {
}

func (r AccountMagicSiteAppConfigNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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
