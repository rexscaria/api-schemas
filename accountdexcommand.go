// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountDexCommandService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexCommandService] method instead.
type AccountDexCommandService struct {
	Options []option.RequestOption
}

// NewAccountDexCommandService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDexCommandService(opts ...option.RequestOption) (r *AccountDexCommandService) {
	r = &AccountDexCommandService{}
	r.Options = opts
	return
}

// Initiate commands for up to 10 devices per account
func (r *AccountDexCommandService) New(ctx context.Context, accountID string, body AccountDexCommandNewParams, opts ...option.RequestOption) (res *AccountDexCommandNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a paginated list of commands issued to devices under the specified
// account, optionally filtered by time range, device, or other parameters
func (r *AccountDexCommandService) List(ctx context.Context, accountID string, query AccountDexCommandListParams, opts ...option.RequestOption) (res *AccountDexCommandListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Downloads artifacts for an executed command. Bulk downloads are not supported
func (r *AccountDexCommandService) DownloadOutput(ctx context.Context, accountID string, commandID string, filename string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if commandID == "" {
		err = errors.New("missing required command_id parameter")
		return
	}
	if filename == "" {
		err = errors.New("missing required filename parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/%s/downloads/%s", accountID, commandID, filename)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves the current quota usage and limits for device commands within a
// specific account, including the time when the quota will reset
func (r *AccountDexCommandService) GetQuota(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountDexCommandGetQuotaResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/quota", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List devices with WARP client support for remote captures which have been
// connected in the last 1 hour.
func (r *AccountDexCommandService) ListEligibleDevices(ctx context.Context, accountID string, query AccountDexCommandListEligibleDevicesParams, opts ...option.RequestOption) (res *AccountDexCommandListEligibleDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexCommandNewResponse struct {
	Result AccountDexCommandNewResponseResult `json:"result"`
	JSON   accountDexCommandNewResponseJSON   `json:"-"`
	APIResponseCollectionDigitalExperienceMonitoring
}

// accountDexCommandNewResponseJSON contains the JSON metadata for the struct
// [AccountDexCommandNewResponse]
type accountDexCommandNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandNewResponseResult struct {
	// List of created commands
	Commands []AccountDexCommandNewResponseResultCommand `json:"commands"`
	JSON     accountDexCommandNewResponseResultJSON      `json:"-"`
}

// accountDexCommandNewResponseResultJSON contains the JSON metadata for the struct
// [AccountDexCommandNewResponseResult]
type accountDexCommandNewResponseResultJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandNewResponseResultCommand struct {
	// Unique identifier for the command
	ID string `json:"id"`
	// Command arguments
	Args map[string]string `json:"args"`
	// Identifier for the device associated with the command
	DeviceID string `json:"device_id"`
	// Current status of the command
	Status AccountDexCommandNewResponseResultCommandsStatus `json:"status"`
	// Type of the command (e.g., "pcap" or "warp-diag")
	Type string                                        `json:"type"`
	JSON accountDexCommandNewResponseResultCommandJSON `json:"-"`
}

// accountDexCommandNewResponseResultCommandJSON contains the JSON metadata for the
// struct [AccountDexCommandNewResponseResultCommand]
type accountDexCommandNewResponseResultCommandJSON struct {
	ID          apijson.Field
	Args        apijson.Field
	DeviceID    apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandNewResponseResultCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandNewResponseResultCommandJSON) RawJSON() string {
	return r.raw
}

// Current status of the command
type AccountDexCommandNewResponseResultCommandsStatus string

const (
	AccountDexCommandNewResponseResultCommandsStatusPendingExec   AccountDexCommandNewResponseResultCommandsStatus = "PENDING_EXEC"
	AccountDexCommandNewResponseResultCommandsStatusPendingUpload AccountDexCommandNewResponseResultCommandsStatus = "PENDING_UPLOAD"
	AccountDexCommandNewResponseResultCommandsStatusSuccess       AccountDexCommandNewResponseResultCommandsStatus = "SUCCESS"
	AccountDexCommandNewResponseResultCommandsStatusFailed        AccountDexCommandNewResponseResultCommandsStatus = "FAILED"
)

func (r AccountDexCommandNewResponseResultCommandsStatus) IsKnown() bool {
	switch r {
	case AccountDexCommandNewResponseResultCommandsStatusPendingExec, AccountDexCommandNewResponseResultCommandsStatusPendingUpload, AccountDexCommandNewResponseResultCommandsStatusSuccess, AccountDexCommandNewResponseResultCommandsStatusFailed:
		return true
	}
	return false
}

type AccountDexCommandListResponse struct {
	Result AccountDexCommandListResponseResult `json:"result"`
	JSON   accountDexCommandListResponseJSON   `json:"-"`
	APIResponseCollectionDigitalExperienceMonitoring
}

// accountDexCommandListResponseJSON contains the JSON metadata for the struct
// [AccountDexCommandListResponse]
type accountDexCommandListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandListResponseResult struct {
	Commands []AccountDexCommandListResponseResultCommand `json:"commands"`
	JSON     accountDexCommandListResponseResultJSON      `json:"-"`
}

// accountDexCommandListResponseResultJSON contains the JSON metadata for the
// struct [AccountDexCommandListResponseResult]
type accountDexCommandListResponseResultJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandListResponseResultCommand struct {
	ID            string                                         `json:"id"`
	CompletedDate time.Time                                      `json:"completed_date,nullable" format:"date-time"`
	CreatedDate   time.Time                                      `json:"created_date" format:"date-time"`
	DeviceID      string                                         `json:"device_id"`
	Filename      string                                         `json:"filename,nullable"`
	Status        string                                         `json:"status"`
	Type          string                                         `json:"type"`
	UserEmail     string                                         `json:"user_email"`
	JSON          accountDexCommandListResponseResultCommandJSON `json:"-"`
}

// accountDexCommandListResponseResultCommandJSON contains the JSON metadata for
// the struct [AccountDexCommandListResponseResultCommand]
type accountDexCommandListResponseResultCommandJSON struct {
	ID            apijson.Field
	CompletedDate apijson.Field
	CreatedDate   apijson.Field
	DeviceID      apijson.Field
	Filename      apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	UserEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexCommandListResponseResultCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListResponseResultCommandJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandGetQuotaResponse struct {
	Result AccountDexCommandGetQuotaResponseResult `json:"result"`
	JSON   accountDexCommandGetQuotaResponseJSON   `json:"-"`
	APIResponseCollectionDigitalExperienceMonitoring
}

// accountDexCommandGetQuotaResponseJSON contains the JSON metadata for the struct
// [AccountDexCommandGetQuotaResponse]
type accountDexCommandGetQuotaResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandGetQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandGetQuotaResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandGetQuotaResponseResult struct {
	// The remaining number of commands that can be initiated for an account
	Quota float64 `json:"quota,required"`
	// The number of commands that have been initiated for an account
	QuotaUsage float64 `json:"quota_usage,required"`
	// The time when the quota resets
	ResetTime time.Time                                   `json:"reset_time,required" format:"date-time"`
	JSON      accountDexCommandGetQuotaResponseResultJSON `json:"-"`
}

// accountDexCommandGetQuotaResponseResultJSON contains the JSON metadata for the
// struct [AccountDexCommandGetQuotaResponseResult]
type accountDexCommandGetQuotaResponseResultJSON struct {
	Quota       apijson.Field
	QuotaUsage  apijson.Field
	ResetTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandGetQuotaResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandGetQuotaResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandListEligibleDevicesResponse struct {
	Result AccountDexCommandListEligibleDevicesResponseResult `json:"result"`
	JSON   accountDexCommandListEligibleDevicesResponseJSON   `json:"-"`
	APIResponseCollectionDigitalExperienceMonitoring
}

// accountDexCommandListEligibleDevicesResponseJSON contains the JSON metadata for
// the struct [AccountDexCommandListEligibleDevicesResponse]
type accountDexCommandListEligibleDevicesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandListEligibleDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListEligibleDevicesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandListEligibleDevicesResponseResult struct {
	// List of eligible devices
	Devices []AccountDexCommandListEligibleDevicesResponseResultDevice `json:"devices"`
	JSON    accountDexCommandListEligibleDevicesResponseResultJSON     `json:"-"`
}

// accountDexCommandListEligibleDevicesResponseResultJSON contains the JSON
// metadata for the struct [AccountDexCommandListEligibleDevicesResponseResult]
type accountDexCommandListEligibleDevicesResponseResultJSON struct {
	Devices     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexCommandListEligibleDevicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListEligibleDevicesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandListEligibleDevicesResponseResultDevice struct {
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId"`
	// Device identifier (human readable)
	DeviceName string `json:"deviceName"`
	// Whether the device is eligible for remote captures
	Eligible bool `json:"eligible"`
	// If the device is not eligible, the reason why.
	IneligibleReason string `json:"ineligibleReason"`
	// User contact email address
	PersonEmail string `json:"personEmail"`
	// Operating system
	Platform string `json:"platform"`
	// Network status
	Status string `json:"status"`
	// Timestamp in ISO format
	Timestamp string `json:"timestamp"`
	// WARP client version
	Version string                                                       `json:"version"`
	JSON    accountDexCommandListEligibleDevicesResponseResultDeviceJSON `json:"-"`
}

// accountDexCommandListEligibleDevicesResponseResultDeviceJSON contains the JSON
// metadata for the struct
// [AccountDexCommandListEligibleDevicesResponseResultDevice]
type accountDexCommandListEligibleDevicesResponseResultDeviceJSON struct {
	DeviceID         apijson.Field
	DeviceName       apijson.Field
	Eligible         apijson.Field
	IneligibleReason apijson.Field
	PersonEmail      apijson.Field
	Platform         apijson.Field
	Status           apijson.Field
	Timestamp        apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountDexCommandListEligibleDevicesResponseResultDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexCommandListEligibleDevicesResponseResultDeviceJSON) RawJSON() string {
	return r.raw
}

type AccountDexCommandNewParams struct {
	// List of device-level commands to execute
	Commands param.Field[[]AccountDexCommandNewParamsCommand] `json:"commands,required"`
}

func (r AccountDexCommandNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDexCommandNewParamsCommand struct {
	// Type of command to execute on the device
	CommandType param.Field[AccountDexCommandNewParamsCommandsCommandType] `json:"command_type,required"`
	// Unique identifier for the device
	DeviceID param.Field[string] `json:"device_id,required"`
	// Email tied to the device
	UserEmail   param.Field[string]                                        `json:"user_email,required"`
	CommandArgs param.Field[AccountDexCommandNewParamsCommandsCommandArgs] `json:"command_args"`
}

func (r AccountDexCommandNewParamsCommand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of command to execute on the device
type AccountDexCommandNewParamsCommandsCommandType string

const (
	AccountDexCommandNewParamsCommandsCommandTypePcap     AccountDexCommandNewParamsCommandsCommandType = "pcap"
	AccountDexCommandNewParamsCommandsCommandTypeWarpDiag AccountDexCommandNewParamsCommandsCommandType = "warp-diag"
)

func (r AccountDexCommandNewParamsCommandsCommandType) IsKnown() bool {
	switch r {
	case AccountDexCommandNewParamsCommandsCommandTypePcap, AccountDexCommandNewParamsCommandsCommandTypeWarpDiag:
		return true
	}
	return false
}

type AccountDexCommandNewParamsCommandsCommandArgs struct {
	// List of interfaces to capture packets on
	Interfaces param.Field[[]AccountDexCommandNewParamsCommandsCommandArgsInterface] `json:"interfaces"`
	// Maximum file size (in MB) for the capture file. Specifies the maximum file size
	// of the warp-diag zip artifact that can be uploaded. If the zip artifact exceeds
	// the specified max file size, it will NOT be uploaded
	MaxFileSizeMB param.Field[float64] `json:"max-file-size-mb"`
	// Maximum number of bytes to save for each packet
	PacketSizeBytes param.Field[float64] `json:"packet-size-bytes"`
	// Test an IP address from all included or excluded ranges. Tests an IP address
	// from all included or excluded ranges. Essentially the same as running 'route get
	// <ip>” and collecting the results. This option may increase the time taken to
	// collect the warp-diag
	TestAllRoutes param.Field[bool] `json:"test-all-routes"`
	// Limit on capture duration (in minutes)
	TimeLimitMin param.Field[float64] `json:"time-limit-min"`
}

func (r AccountDexCommandNewParamsCommandsCommandArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDexCommandNewParamsCommandsCommandArgsInterface string

const (
	AccountDexCommandNewParamsCommandsCommandArgsInterfaceDefault AccountDexCommandNewParamsCommandsCommandArgsInterface = "default"
	AccountDexCommandNewParamsCommandsCommandArgsInterfaceTunnel  AccountDexCommandNewParamsCommandsCommandArgsInterface = "tunnel"
)

func (r AccountDexCommandNewParamsCommandsCommandArgsInterface) IsKnown() bool {
	switch r {
	case AccountDexCommandNewParamsCommandsCommandArgsInterfaceDefault, AccountDexCommandNewParamsCommandsCommandArgsInterfaceTunnel:
		return true
	}
	return false
}

type AccountDexCommandListParams struct {
	// Page number for pagination
	Page param.Field[float64] `query:"page,required"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Optionally filter executed commands by command type
	CommandType param.Field[string] `query:"command_type"`
	// Unique identifier for a device
	DeviceID param.Field[string] `query:"device_id"`
	// Start time for the query in ISO (RFC3339 - ISO 8601) format
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Optionally filter executed commands by status
	Status param.Field[AccountDexCommandListParamsStatus] `query:"status"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format
	To param.Field[time.Time] `query:"to" format:"date-time"`
	// Email tied to the device
	UserEmail param.Field[string] `query:"user_email"`
}

// URLQuery serializes [AccountDexCommandListParams]'s query parameters as
// `url.Values`.
func (r AccountDexCommandListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optionally filter executed commands by status
type AccountDexCommandListParamsStatus string

const (
	AccountDexCommandListParamsStatusPendingExec   AccountDexCommandListParamsStatus = "PENDING_EXEC"
	AccountDexCommandListParamsStatusPendingUpload AccountDexCommandListParamsStatus = "PENDING_UPLOAD"
	AccountDexCommandListParamsStatusSuccess       AccountDexCommandListParamsStatus = "SUCCESS"
	AccountDexCommandListParamsStatusFailed        AccountDexCommandListParamsStatus = "FAILED"
)

func (r AccountDexCommandListParamsStatus) IsKnown() bool {
	switch r {
	case AccountDexCommandListParamsStatusPendingExec, AccountDexCommandListParamsStatusPendingUpload, AccountDexCommandListParamsStatusSuccess, AccountDexCommandListParamsStatusFailed:
		return true
	}
	return false
}

type AccountDexCommandListEligibleDevicesParams struct {
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Filter devices by name or email
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountDexCommandListEligibleDevicesParams]'s query
// parameters as `url.Values`.
func (r AccountDexCommandListEligibleDevicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
