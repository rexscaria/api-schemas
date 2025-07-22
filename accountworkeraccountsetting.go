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

// AccountWorkerAccountSettingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerAccountSettingService] method instead.
type AccountWorkerAccountSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerAccountSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerAccountSettingService(opts ...option.RequestOption) (r *AccountWorkerAccountSettingService) {
	r = &AccountWorkerAccountSettingService{}
	r.Options = opts
	return
}

// Fetches Worker account settings for an account.
func (r *AccountWorkerAccountSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *SettingsResponseAccountSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Creates Worker account settings for an account.
func (r *AccountWorkerAccountSettingService) Update(ctx context.Context, accountID string, body AccountWorkerAccountSettingUpdateParams, opts ...option.RequestOption) (res *SettingsResponseAccountSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SettingsResponseAccountSettings struct {
	Result WorkersAccountSettings              `json:"result"`
	JSON   settingsResponseAccountSettingsJSON `json:"-"`
	CommonResponseWorkers
}

// settingsResponseAccountSettingsJSON contains the JSON metadata for the struct
// [SettingsResponseAccountSettings]
type settingsResponseAccountSettingsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsResponseAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsResponseAccountSettingsJSON) RawJSON() string {
	return r.raw
}

type WorkersAccountSettings struct {
	DefaultUsageModel string                     `json:"default_usage_model"`
	GreenCompute      bool                       `json:"green_compute"`
	JSON              workersAccountSettingsJSON `json:"-"`
}

// workersAccountSettingsJSON contains the JSON metadata for the struct
// [WorkersAccountSettings]
type workersAccountSettingsJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WorkersAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersAccountSettingsJSON) RawJSON() string {
	return r.raw
}

type WorkersAccountSettingsParam struct {
	DefaultUsageModel param.Field[string] `json:"default_usage_model"`
	GreenCompute      param.Field[bool]   `json:"green_compute"`
}

func (r WorkersAccountSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerAccountSettingUpdateParams struct {
	WorkersAccountSettings WorkersAccountSettingsParam `json:"workers_account_settings,required"`
}

func (r AccountWorkerAccountSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.WorkersAccountSettings)
}
