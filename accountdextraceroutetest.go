// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountDexTracerouteTestService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexTracerouteTestService] method instead.
type AccountDexTracerouteTestService struct {
	Options []option.RequestOption
}

// NewAccountDexTracerouteTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexTracerouteTestService(opts ...option.RequestOption) (r *AccountDexTracerouteTestService) {
	r = &AccountDexTracerouteTestService{}
	r.Options = opts
	return
}

// Get test details and aggregate performance metrics for an traceroute test for a
// given time period between 1 hour and 7 days.
func (r *AccountDexTracerouteTestService) Get(ctx context.Context, accountID string, testID string, query AccountDexTracerouteTestGetParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a breakdown of metrics by hop for individual traceroute test runs
func (r *AccountDexTracerouteTestService) GetNetworkPath(ctx context.Context, accountID string, testID string, query AccountDexTracerouteTestGetNetworkPathParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestGetNetworkPathResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/network-path", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get percentiles for a traceroute test for a given time period between 1 hour and
// 7 days.
func (r *AccountDexTracerouteTestService) GetPercentiles(ctx context.Context, accountID string, testID string, query AccountDexTracerouteTestGetPercentilesParams, opts ...option.RequestOption) (res *AccountDexTracerouteTestGetPercentilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDexTracerouteTestGetResponse struct {
	Result AccountDexTracerouteTestGetResponseResult `json:"result"`
	JSON   accountDexTracerouteTestGetResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexTracerouteTestGetResponseJSON contains the JSON metadata for the
// struct [AccountDexTracerouteTestGetResponse]
type accountDexTracerouteTestGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetResponseResult struct {
	// The host of the Traceroute synthetic application test
	Host string `json:"host,required"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval string                                        `json:"interval,required"`
	Kind     AccountDexTracerouteTestGetResponseResultKind `json:"kind,required"`
	// The name of the Traceroute synthetic application test
	Name                  string                                                           `json:"name,required"`
	TargetPolicies        []AccountDexTracerouteTestGetResponseResultTargetPolicy          `json:"target_policies,nullable"`
	Targeted              bool                                                             `json:"targeted"`
	TracerouteStats       AccountDexTracerouteTestGetResponseResultTracerouteStats         `json:"tracerouteStats,nullable"`
	TracerouteStatsByColo []AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo `json:"tracerouteStatsByColo"`
	JSON                  accountDexTracerouteTestGetResponseResultJSON                    `json:"-"`
}

// accountDexTracerouteTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDexTracerouteTestGetResponseResult]
type accountDexTracerouteTestGetResponseResultJSON struct {
	Host                  apijson.Field
	Interval              apijson.Field
	Kind                  apijson.Field
	Name                  apijson.Field
	TargetPolicies        apijson.Field
	Targeted              apijson.Field
	TracerouteStats       apijson.Field
	TracerouteStatsByColo apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetResponseResultKind string

const (
	AccountDexTracerouteTestGetResponseResultKindTraceroute AccountDexTracerouteTestGetResponseResultKind = "traceroute"
)

func (r AccountDexTracerouteTestGetResponseResultKind) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestGetResponseResultKindTraceroute:
		return true
	}
	return false
}

type AccountDexTracerouteTestGetResponseResultTargetPolicy struct {
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                                                      `json:"default,required"`
	Name    string                                                    `json:"name,required"`
	JSON    accountDexTracerouteTestGetResponseResultTargetPolicyJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTargetPolicyJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestGetResponseResultTargetPolicy]
type accountDexTracerouteTestGetResponseResultTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetResponseResultTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetResponseResultTracerouteStats struct {
	AvailabilityPct TestStatPctOverTime `json:"availabilityPct,required"`
	HopsCount       TestStatOverTime    `json:"hopsCount,required"`
	PacketLossPct   TestStatPctOverTime `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime    `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                        `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTracerouteTestGetResponseResultTracerouteStatsJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsJSON contains the JSON
// metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStats]
type accountDexTracerouteTestGetResponseResultTracerouteStatsJSON struct {
	AvailabilityPct    apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetResponseResultTracerouteStatsJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo struct {
	AvailabilityPct TestStatPctOverTime `json:"availabilityPct,required"`
	Colo            string              `json:"colo,required"`
	HopsCount       TestStatOverTime    `json:"hopsCount,required"`
	PacketLossPct   TestStatPctOverTime `json:"packetLossPct,required"`
	RoundTripTimeMs TestStatOverTime    `json:"roundTripTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                              `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON `json:"-"`
}

// accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo]
type accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON struct {
	AvailabilityPct    apijson.Field
	Colo               apijson.Field
	HopsCount          apijson.Field
	PacketLossPct      apijson.Field
	RoundTripTimeMs    apijson.Field
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetResponseResultTracerouteStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetResponseResultTracerouteStatsByColoJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetNetworkPathResponse struct {
	Result AccountDexTracerouteTestGetNetworkPathResponseResult `json:"result"`
	JSON   accountDexTracerouteTestGetNetworkPathResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexTracerouteTestGetNetworkPathResponseJSON contains the JSON metadata
// for the struct [AccountDexTracerouteTestGetNetworkPathResponse]
type accountDexTracerouteTestGetNetworkPathResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetNetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetNetworkPathResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetNetworkPathResponseResult struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                                                          `json:"interval"`
	Kind        AccountDexTracerouteTestGetNetworkPathResponseResultKind        `json:"kind"`
	Name        string                                                          `json:"name"`
	NetworkPath AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPath `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                                                   `json:"url"`
	JSON accountDexTracerouteTestGetNetworkPathResponseResultJSON `json:"-"`
}

// accountDexTracerouteTestGetNetworkPathResponseResultJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestGetNetworkPathResponseResult]
type accountDexTracerouteTestGetNetworkPathResponseResultJSON struct {
	ID          apijson.Field
	DeviceName  apijson.Field
	Interval    apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	NetworkPath apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetNetworkPathResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetNetworkPathResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetNetworkPathResponseResultKind string

const (
	AccountDexTracerouteTestGetNetworkPathResponseResultKindTraceroute AccountDexTracerouteTestGetNetworkPathResponseResultKind = "traceroute"
)

func (r AccountDexTracerouteTestGetNetworkPathResponseResultKind) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestGetNetworkPathResponseResultKindTraceroute:
		return true
	}
	return false
}

type AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPath struct {
	Slots []AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSampling `json:"sampling,nullable"`
	JSON     accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathJSON     `json:"-"`
}

// accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathJSON contains the
// JSON metadata for the struct
// [AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPath]
type accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlot struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// Round trip time in ms of the client to app mile
	ClientToAppRttMs int64 `json:"clientToAppRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare egress mile
	ClientToCfEgressRttMs int64 `json:"clientToCfEgressRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare ingress mile
	ClientToCfIngressRttMs int64  `json:"clientToCfIngressRttMs,required,nullable"`
	Timestamp              string `json:"timestamp,required"`
	// Round trip time in ms of the client to ISP mile
	ClientToIspRttMs int64                                                                   `json:"clientToIspRttMs,nullable"`
	JSON             accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlotJSON `json:"-"`
}

// accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlotJSON contains
// the JSON metadata for the struct
// [AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlot]
type accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRttMs       apijson.Field
	ClientToCfEgressRttMs  apijson.Field
	ClientToCfIngressRttMs apijson.Field
	Timestamp              apijson.Field
	ClientToIspRttMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSlotJSON) RawJSON() string {
	return r.raw
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSampling struct {
	Unit  AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnit `json:"unit,required"`
	Value int64                                                                       `json:"value,required"`
	JSON  accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingJSON `json:"-"`
}

// accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingJSON
// contains the JSON metadata for the struct
// [AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSampling]
type accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnit string

const (
	AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnitHours AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnit = "hours"
)

func (r AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnit) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestGetNetworkPathResponseResultNetworkPathSamplingUnitHours:
		return true
	}
	return false
}

type AccountDexTracerouteTestGetPercentilesResponse struct {
	Result AccountDexTracerouteTestGetPercentilesResponseResult `json:"result"`
	JSON   accountDexTracerouteTestGetPercentilesResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexTracerouteTestGetPercentilesResponseJSON contains the JSON metadata
// for the struct [AccountDexTracerouteTestGetPercentilesResponse]
type accountDexTracerouteTestGetPercentilesResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetPercentilesResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetPercentilesResponseResult struct {
	HopsCount       Percentiles                                              `json:"hopsCount"`
	PacketLossPct   Percentiles                                              `json:"packetLossPct"`
	RoundTripTimeMs Percentiles                                              `json:"roundTripTimeMs"`
	JSON            accountDexTracerouteTestGetPercentilesResponseResultJSON `json:"-"`
}

// accountDexTracerouteTestGetPercentilesResponseResultJSON contains the JSON
// metadata for the struct [AccountDexTracerouteTestGetPercentilesResponseResult]
type accountDexTracerouteTestGetPercentilesResponseResultJSON struct {
	HopsCount       apijson.Field
	PacketLossPct   apijson.Field
	RoundTripTimeMs apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDexTracerouteTestGetPercentilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTracerouteTestGetPercentilesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTracerouteTestGetParams struct {
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexTracerouteTestGetParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [AccountDexTracerouteTestGetParams]'s query parameters as
// `url.Values`.
func (r AccountDexTracerouteTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexTracerouteTestGetParamsInterval string

const (
	AccountDexTracerouteTestGetParamsIntervalMinute AccountDexTracerouteTestGetParamsInterval = "minute"
	AccountDexTracerouteTestGetParamsIntervalHour   AccountDexTracerouteTestGetParamsInterval = "hour"
)

func (r AccountDexTracerouteTestGetParamsInterval) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestGetParamsIntervalMinute, AccountDexTracerouteTestGetParamsIntervalHour:
		return true
	}
	return false
}

type AccountDexTracerouteTestGetNetworkPathParams struct {
	// Device to filter tracroute result runs to
	DeviceID param.Field[string] `query:"deviceId,required"`
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexTracerouteTestGetNetworkPathParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
}

// URLQuery serializes [AccountDexTracerouteTestGetNetworkPathParams]'s query
// parameters as `url.Values`.
func (r AccountDexTracerouteTestGetNetworkPathParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexTracerouteTestGetNetworkPathParamsInterval string

const (
	AccountDexTracerouteTestGetNetworkPathParamsIntervalMinute AccountDexTracerouteTestGetNetworkPathParamsInterval = "minute"
	AccountDexTracerouteTestGetNetworkPathParamsIntervalHour   AccountDexTracerouteTestGetNetworkPathParamsInterval = "hour"
)

func (r AccountDexTracerouteTestGetNetworkPathParamsInterval) IsKnown() bool {
	switch r {
	case AccountDexTracerouteTestGetNetworkPathParamsIntervalMinute, AccountDexTracerouteTestGetNetworkPathParamsIntervalHour:
		return true
	}
	return false
}

type AccountDexTracerouteTestGetPercentilesParams struct {
	// Start time for the query in ISO (RFC3339 - ISO 8601) format
	From param.Field[string] `query:"from,required"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format
	To param.Field[string] `query:"to,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [AccountDexTracerouteTestGetPercentilesParams]'s query
// parameters as `url.Values`.
func (r AccountDexTracerouteTestGetPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
