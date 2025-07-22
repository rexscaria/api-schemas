// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDexFleetStatusService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexFleetStatusService] method instead.
type AccountDexFleetStatusService struct {
	Options []option.RequestOption
}

// NewAccountDexFleetStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDexFleetStatusService(opts ...option.RequestOption) (r *AccountDexFleetStatusService) {
	r = &AccountDexFleetStatusService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *AccountDexFleetStatusService) ListDevices(ctx context.Context, accountID string, query AccountDexFleetStatusListDevicesParams, opts ...option.RequestOption) (res *AccountDexFleetStatusListDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List details for live (up to 60 minutes) devices using WARP
func (r *AccountDexFleetStatusService) ListLiveStatus(ctx context.Context, accountID string, query AccountDexFleetStatusListLiveStatusParams, opts ...option.RequestOption) (res *AccountDexFleetStatusListLiveStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/live", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List details for devices using WARP, up to 7 days
func (r *AccountDexFleetStatusService) ListOverTime(ctx context.Context, accountID string, query AccountDexFleetStatusListOverTimeParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/over-time", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, nil, opts...)
	return
}

type LiveStat struct {
	// Number of unique devices
	UniqueDevicesTotal float64      `json:"uniqueDevicesTotal"`
	Value              string       `json:"value"`
	JSON               liveStatJSON `json:"-"`
}

// liveStatJSON contains the JSON metadata for the struct [LiveStat]
type liveStatJSON struct {
	UniqueDevicesTotal apijson.Field
	Value              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *LiveStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r liveStatJSON) RawJSON() string {
	return r.raw
}

type AccountDexFleetStatusListDevicesResponse struct {
	Result []Device                                     `json:"result"`
	JSON   accountDexFleetStatusListDevicesResponseJSON `json:"-"`
	APIResponseCollectionDigitalExperienceMonitoring
}

// accountDexFleetStatusListDevicesResponseJSON contains the JSON metadata for the
// struct [AccountDexFleetStatusListDevicesResponse]
type accountDexFleetStatusListDevicesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusListDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexFleetStatusListDevicesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexFleetStatusListLiveStatusResponse struct {
	Result AccountDexFleetStatusListLiveStatusResponseResult `json:"result"`
	JSON   accountDexFleetStatusListLiveStatusResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexFleetStatusListLiveStatusResponseJSON contains the JSON metadata for
// the struct [AccountDexFleetStatusListLiveStatusResponse]
type accountDexFleetStatusListLiveStatusResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusListLiveStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexFleetStatusListLiveStatusResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexFleetStatusListLiveStatusResponseResult struct {
	DeviceStats AccountDexFleetStatusListLiveStatusResponseResultDeviceStats `json:"deviceStats"`
	JSON        accountDexFleetStatusListLiveStatusResponseResultJSON        `json:"-"`
}

// accountDexFleetStatusListLiveStatusResponseResultJSON contains the JSON metadata
// for the struct [AccountDexFleetStatusListLiveStatusResponseResult]
type accountDexFleetStatusListLiveStatusResponseResultJSON struct {
	DeviceStats apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexFleetStatusListLiveStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexFleetStatusListLiveStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexFleetStatusListLiveStatusResponseResultDeviceStats struct {
	ByColo     []LiveStat `json:"byColo,nullable"`
	ByMode     []LiveStat `json:"byMode,nullable"`
	ByPlatform []LiveStat `json:"byPlatform,nullable"`
	ByStatus   []LiveStat `json:"byStatus,nullable"`
	ByVersion  []LiveStat `json:"byVersion,nullable"`
	// Number of unique devices
	UniqueDevicesTotal float64                                                          `json:"uniqueDevicesTotal"`
	JSON               accountDexFleetStatusListLiveStatusResponseResultDeviceStatsJSON `json:"-"`
}

// accountDexFleetStatusListLiveStatusResponseResultDeviceStatsJSON contains the
// JSON metadata for the struct
// [AccountDexFleetStatusListLiveStatusResponseResultDeviceStats]
type accountDexFleetStatusListLiveStatusResponseResultDeviceStatsJSON struct {
	ByColo             apijson.Field
	ByMode             apijson.Field
	ByPlatform         apijson.Field
	ByStatus           apijson.Field
	ByVersion          apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexFleetStatusListLiveStatusResponseResultDeviceStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexFleetStatusListLiveStatusResponseResultDeviceStatsJSON) RawJSON() string {
	return r.raw
}

type AccountDexFleetStatusListDevicesParams struct {
	// Time range beginning in ISO format
	From param.Field[string] `query:"from,required"`
	// Page number
	Page param.Field[float64] `query:"page,required"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Time range end in ISO format
	To param.Field[string] `query:"to,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
	// The mode under which the WARP client is run
	Mode param.Field[string] `query:"mode"`
	// Operating system
	Platform param.Field[string] `query:"platform"`
	// Dimension to sort results by
	SortBy param.Field[AccountDexFleetStatusListDevicesParamsSortBy] `query:"sort_by"`
	// Source:
	//
	// - `hourly` - device details aggregated hourly, up to 7 days prior
	// - `last_seen` - device details, up to 24 hours prior
	// - `raw` - device details, up to 7 days prior
	Source param.Field[AccountDexFleetStatusListDevicesParamsSource] `query:"source"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [AccountDexFleetStatusListDevicesParams]'s query parameters
// as `url.Values`.
func (r AccountDexFleetStatusListDevicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dimension to sort results by
type AccountDexFleetStatusListDevicesParamsSortBy string

const (
	AccountDexFleetStatusListDevicesParamsSortByColo      AccountDexFleetStatusListDevicesParamsSortBy = "colo"
	AccountDexFleetStatusListDevicesParamsSortByDeviceID  AccountDexFleetStatusListDevicesParamsSortBy = "device_id"
	AccountDexFleetStatusListDevicesParamsSortByMode      AccountDexFleetStatusListDevicesParamsSortBy = "mode"
	AccountDexFleetStatusListDevicesParamsSortByPlatform  AccountDexFleetStatusListDevicesParamsSortBy = "platform"
	AccountDexFleetStatusListDevicesParamsSortByStatus    AccountDexFleetStatusListDevicesParamsSortBy = "status"
	AccountDexFleetStatusListDevicesParamsSortByTimestamp AccountDexFleetStatusListDevicesParamsSortBy = "timestamp"
	AccountDexFleetStatusListDevicesParamsSortByVersion   AccountDexFleetStatusListDevicesParamsSortBy = "version"
)

func (r AccountDexFleetStatusListDevicesParamsSortBy) IsKnown() bool {
	switch r {
	case AccountDexFleetStatusListDevicesParamsSortByColo, AccountDexFleetStatusListDevicesParamsSortByDeviceID, AccountDexFleetStatusListDevicesParamsSortByMode, AccountDexFleetStatusListDevicesParamsSortByPlatform, AccountDexFleetStatusListDevicesParamsSortByStatus, AccountDexFleetStatusListDevicesParamsSortByTimestamp, AccountDexFleetStatusListDevicesParamsSortByVersion:
		return true
	}
	return false
}

// Source:
//
// - `hourly` - device details aggregated hourly, up to 7 days prior
// - `last_seen` - device details, up to 24 hours prior
// - `raw` - device details, up to 7 days prior
type AccountDexFleetStatusListDevicesParamsSource string

const (
	AccountDexFleetStatusListDevicesParamsSourceLastSeen AccountDexFleetStatusListDevicesParamsSource = "last_seen"
	AccountDexFleetStatusListDevicesParamsSourceHourly   AccountDexFleetStatusListDevicesParamsSource = "hourly"
	AccountDexFleetStatusListDevicesParamsSourceRaw      AccountDexFleetStatusListDevicesParamsSource = "raw"
)

func (r AccountDexFleetStatusListDevicesParamsSource) IsKnown() bool {
	switch r {
	case AccountDexFleetStatusListDevicesParamsSourceLastSeen, AccountDexFleetStatusListDevicesParamsSourceHourly, AccountDexFleetStatusListDevicesParamsSourceRaw:
		return true
	}
	return false
}

type AccountDexFleetStatusListLiveStatusParams struct {
	// Number of minutes before current time
	SinceMinutes param.Field[float64] `query:"since_minutes,required"`
}

// URLQuery serializes [AccountDexFleetStatusListLiveStatusParams]'s query
// parameters as `url.Values`.
func (r AccountDexFleetStatusListLiveStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDexFleetStatusListOverTimeParams struct {
	// Time range beginning in ISO format
	From param.Field[string] `query:"from,required"`
	// Time range end in ISO format
	To param.Field[string] `query:"to,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
}

// URLQuery serializes [AccountDexFleetStatusListOverTimeParams]'s query parameters
// as `url.Values`.
func (r AccountDexFleetStatusListOverTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
