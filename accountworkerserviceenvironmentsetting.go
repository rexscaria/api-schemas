// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountWorkerServiceEnvironmentSettingService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerServiceEnvironmentSettingService] method instead.
type AccountWorkerServiceEnvironmentSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerServiceEnvironmentSettingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountWorkerServiceEnvironmentSettingService(opts ...option.RequestOption) (r *AccountWorkerServiceEnvironmentSettingService) {
	r = &AccountWorkerServiceEnvironmentSettingService{}
	r.Options = opts
	return
}

// Get script settings from a worker with an environment.
func (r *AccountWorkerServiceEnvironmentSettingService) Get(ctx context.Context, accountID string, serviceName string, environmentName string, opts ...option.RequestOption) (res *SettingsResponseScriptSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceName == "" {
		err = errors.New("missing required service_name parameter")
		return
	}
	if environmentName == "" {
		err = errors.New("missing required environment_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script metadata, such as bindings.
func (r *AccountWorkerServiceEnvironmentSettingService) Patch(ctx context.Context, accountID string, serviceName string, environmentName string, body AccountWorkerServiceEnvironmentSettingPatchParams, opts ...option.RequestOption) (res *SettingsResponseScriptSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceName == "" {
		err = errors.New("missing required service_name parameter")
		return
	}
	if environmentName == "" {
		err = errors.New("missing required environment_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/services/%s/environments/%s/settings", accountID, serviceName, environmentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountWorkerServiceEnvironmentSettingPatchParams struct {
	SettingsResponseScriptSettings SettingsResponseScriptSettingsParam `json:"settings_response_script_settings,required"`
}

func (r AccountWorkerServiceEnvironmentSettingPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SettingsResponseScriptSettings)
}
