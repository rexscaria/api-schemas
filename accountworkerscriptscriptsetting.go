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

// AccountWorkerScriptScriptSettingService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWorkerScriptScriptSettingService] method instead.
type AccountWorkerScriptScriptSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptScriptSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptScriptSettingService(opts ...option.RequestOption) (r *AccountWorkerScriptScriptSettingService) {
	r = &AccountWorkerScriptScriptSettingService{}
	r.Options = opts
	return
}

// Get script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Includes Logpush and Tail Consumers.
func (r *AccountWorkerScriptScriptSettingService) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *SettingsResponseScriptSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patch script-level settings when using
// [Worker Versions](https://developers.cloudflare.com/api/operations/worker-versions-list-versions).
// Including but not limited to Logpush and Tail Consumers.
func (r *AccountWorkerScriptScriptSettingService) Patch(ctx context.Context, accountID string, scriptName string, body AccountWorkerScriptScriptSettingPatchParams, opts ...option.RequestOption) (res *SettingsResponseScriptSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if scriptName == "" {
		err = errors.New("missing required script_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/script-settings", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Observability settings for the Worker.
type Observability struct {
	// Whether observability is enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate float64 `json:"head_sampling_rate,nullable"`
	// Log settings for the Worker.
	Logs ObservabilityLogs `json:"logs,nullable"`
	JSON observabilityJSON `json:"-"`
}

// observabilityJSON contains the JSON metadata for the struct [Observability]
type observabilityJSON struct {
	Enabled          apijson.Field
	HeadSamplingRate apijson.Field
	Logs             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Observability) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityJSON) RawJSON() string {
	return r.raw
}

// Log settings for the Worker.
type ObservabilityLogs struct {
	// Whether logs are enabled for the Worker.
	Enabled bool `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs bool `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate float64               `json:"head_sampling_rate,nullable"`
	JSON             observabilityLogsJSON `json:"-"`
}

// observabilityLogsJSON contains the JSON metadata for the struct
// [ObservabilityLogs]
type observabilityLogsJSON struct {
	Enabled          apijson.Field
	InvocationLogs   apijson.Field
	HeadSamplingRate apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservabilityLogs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityLogsJSON) RawJSON() string {
	return r.raw
}

// Observability settings for the Worker.
type ObservabilityParam struct {
	// Whether observability is enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// The sampling rate for incoming requests. From 0 to 1 (1 = 100%, 0.1 = 10%).
	// Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
	// Log settings for the Worker.
	Logs param.Field[ObservabilityLogsParam] `json:"logs"`
}

func (r ObservabilityParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Log settings for the Worker.
type ObservabilityLogsParam struct {
	// Whether logs are enabled for the Worker.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Whether
	// [invocation logs](https://developers.cloudflare.com/workers/observability/logs/workers-logs/#invocation-logs)
	// are enabled for the Worker.
	InvocationLogs param.Field[bool] `json:"invocation_logs,required"`
	// The sampling rate for logs. From 0 to 1 (1 = 100%, 0.1 = 10%). Default is 1.
	HeadSamplingRate param.Field[float64] `json:"head_sampling_rate"`
}

func (r ObservabilityLogsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingsItem struct {
	// Whether Logpush is turned on for the Worker.
	Logpush bool `json:"logpush"`
	// Observability settings for the Worker.
	Observability Observability `json:"observability,nullable"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers []TailConsumersScript `json:"tail_consumers,nullable"`
	JSON          settingsItemJSON      `json:"-"`
}

// settingsItemJSON contains the JSON metadata for the struct [SettingsItem]
type settingsItemJSON struct {
	Logpush       apijson.Field
	Observability apijson.Field
	TailConsumers apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SettingsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsItemJSON) RawJSON() string {
	return r.raw
}

type SettingsItemParam struct {
	// Whether Logpush is turned on for the Worker.
	Logpush param.Field[bool] `json:"logpush"`
	// Observability settings for the Worker.
	Observability param.Field[ObservabilityParam] `json:"observability"`
	// List of Workers that will consume logs from the attached Worker.
	TailConsumers param.Field[[]TailConsumersScriptParam] `json:"tail_consumers"`
}

func (r SettingsItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingsResponseScriptSettings struct {
	Errors   []WorkersMessages `json:"errors,required"`
	Messages []WorkersMessages `json:"messages,required"`
	Result   SettingsItem      `json:"result,required"`
	// Whether the API call was successful.
	Success SettingsResponseScriptSettingsSuccess `json:"success,required"`
	JSON    settingsResponseScriptSettingsJSON    `json:"-"`
}

// settingsResponseScriptSettingsJSON contains the JSON metadata for the struct
// [SettingsResponseScriptSettings]
type settingsResponseScriptSettingsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsResponseScriptSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsResponseScriptSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingsResponseScriptSettingsSuccess bool

const (
	SettingsResponseScriptSettingsSuccessTrue SettingsResponseScriptSettingsSuccess = true
)

func (r SettingsResponseScriptSettingsSuccess) IsKnown() bool {
	switch r {
	case SettingsResponseScriptSettingsSuccessTrue:
		return true
	}
	return false
}

type SettingsResponseScriptSettingsParam struct {
	Errors   param.Field[[]WorkersMessagesParam] `json:"errors,required"`
	Messages param.Field[[]WorkersMessagesParam] `json:"messages,required"`
	Result   param.Field[SettingsItemParam]      `json:"result,required"`
	// Whether the API call was successful.
	Success param.Field[SettingsResponseScriptSettingsSuccess] `json:"success,required"`
}

func (r SettingsResponseScriptSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A reference to a script that will consume logs from the attached Worker.
type TailConsumersScript struct {
	// Name of Worker that is to be the consumer.
	Service string `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment string `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace string                  `json:"namespace"`
	JSON      tailConsumersScriptJSON `json:"-"`
}

// tailConsumersScriptJSON contains the JSON metadata for the struct
// [TailConsumersScript]
type tailConsumersScriptJSON struct {
	Service     apijson.Field
	Environment apijson.Field
	Namespace   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TailConsumersScript) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tailConsumersScriptJSON) RawJSON() string {
	return r.raw
}

// A reference to a script that will consume logs from the attached Worker.
type TailConsumersScriptParam struct {
	// Name of Worker that is to be the consumer.
	Service param.Field[string] `json:"service,required"`
	// Optional environment if the Worker utilizes one.
	Environment param.Field[string] `json:"environment"`
	// Optional dispatch namespace the script belongs to.
	Namespace param.Field[string] `json:"namespace"`
}

func (r TailConsumersScriptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerScriptScriptSettingPatchParams struct {
	SettingsItem SettingsItemParam `json:"settings_item,required"`
}

func (r AccountWorkerScriptScriptSettingPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SettingsItem)
}
