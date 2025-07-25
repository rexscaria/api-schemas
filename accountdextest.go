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

// AccountDexTestService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexTestService] method instead.
type AccountDexTestService struct {
	Options []option.RequestOption
}

// NewAccountDexTestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDexTestService(opts ...option.RequestOption) (r *AccountDexTestService) {
	r = &AccountDexTestService{}
	r.Options = opts
	return
}

// Returns unique count of devices that have run synthetic application monitoring
// tests in the past 7 days.
func (r *AccountDexTestService) CountUniqueDevices(ctx context.Context, accountID string, query AccountDexTestCountUniqueDevicesParams, opts ...option.RequestOption) (res *AccountDexTestCountUniqueDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/tests/unique-devices", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// List DEX tests with overview metrics
func (r *AccountDexTestService) ListOverview(ctx context.Context, accountID string, query AccountDexTestListOverviewParams, opts ...option.RequestOption) (res *AccountDexTestListOverviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/tests/overview", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AggregateTimePeriod struct {
	Units AggregateTimePeriodUnits `json:"units,required"`
	Value int64                    `json:"value,required"`
	JSON  aggregateTimePeriodJSON  `json:"-"`
}

// aggregateTimePeriodJSON contains the JSON metadata for the struct
// [AggregateTimePeriod]
type aggregateTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AggregateTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregateTimePeriodJSON) RawJSON() string {
	return r.raw
}

type AggregateTimePeriodUnits string

const (
	AggregateTimePeriodUnitsHours    AggregateTimePeriodUnits = "hours"
	AggregateTimePeriodUnitsDays     AggregateTimePeriodUnits = "days"
	AggregateTimePeriodUnitsTestRuns AggregateTimePeriodUnits = "testRuns"
)

func (r AggregateTimePeriodUnits) IsKnown() bool {
	switch r {
	case AggregateTimePeriodUnitsHours, AggregateTimePeriodUnitsDays, AggregateTimePeriodUnitsTestRuns:
		return true
	}
	return false
}

type TimingAggregates struct {
	History  []TimingAggregatesHistory `json:"history,required"`
	AvgMs    int64                     `json:"avgMs,nullable"`
	OverTime TimingAggregatesOverTime  `json:"overTime,nullable"`
	JSON     timingAggregatesJSON      `json:"-"`
}

// timingAggregatesJSON contains the JSON metadata for the struct
// [TimingAggregates]
type timingAggregatesJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimingAggregates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timingAggregatesJSON) RawJSON() string {
	return r.raw
}

type TimingAggregatesHistory struct {
	TimePeriod AggregateTimePeriod         `json:"timePeriod,required"`
	AvgMs      int64                       `json:"avgMs,nullable"`
	DeltaPct   float64                     `json:"deltaPct,nullable"`
	JSON       timingAggregatesHistoryJSON `json:"-"`
}

// timingAggregatesHistoryJSON contains the JSON metadata for the struct
// [TimingAggregatesHistory]
type timingAggregatesHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimingAggregatesHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timingAggregatesHistoryJSON) RawJSON() string {
	return r.raw
}

type TimingAggregatesOverTime struct {
	TimePeriod AggregateTimePeriod             `json:"timePeriod,required"`
	Values     []TimingAggregatesOverTimeValue `json:"values,required"`
	JSON       timingAggregatesOverTimeJSON    `json:"-"`
}

// timingAggregatesOverTimeJSON contains the JSON metadata for the struct
// [TimingAggregatesOverTime]
type timingAggregatesOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimingAggregatesOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timingAggregatesOverTimeJSON) RawJSON() string {
	return r.raw
}

type TimingAggregatesOverTimeValue struct {
	AvgMs     int64                             `json:"avgMs,required"`
	Timestamp string                            `json:"timestamp,required"`
	JSON      timingAggregatesOverTimeValueJSON `json:"-"`
}

// timingAggregatesOverTimeValueJSON contains the JSON metadata for the struct
// [TimingAggregatesOverTimeValue]
type timingAggregatesOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TimingAggregatesOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r timingAggregatesOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestCountUniqueDevicesResponse struct {
	Errors   []Item `json:"errors,required"`
	Messages []Item `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDexTestCountUniqueDevicesResponseSuccess `json:"success,required"`
	Result  AccountDexTestCountUniqueDevicesResponseResult  `json:"result"`
	JSON    accountDexTestCountUniqueDevicesResponseJSON    `json:"-"`
}

// accountDexTestCountUniqueDevicesResponseJSON contains the JSON metadata for the
// struct [AccountDexTestCountUniqueDevicesResponse]
type accountDexTestCountUniqueDevicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestCountUniqueDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestCountUniqueDevicesResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDexTestCountUniqueDevicesResponseSuccess bool

const (
	AccountDexTestCountUniqueDevicesResponseSuccessTrue AccountDexTestCountUniqueDevicesResponseSuccess = true
)

func (r AccountDexTestCountUniqueDevicesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDexTestCountUniqueDevicesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDexTestCountUniqueDevicesResponseResult struct {
	// total number of unique devices
	UniqueDevicesTotal int64                                              `json:"uniqueDevicesTotal,required"`
	JSON               accountDexTestCountUniqueDevicesResponseResultJSON `json:"-"`
}

// accountDexTestCountUniqueDevicesResponseResultJSON contains the JSON metadata
// for the struct [AccountDexTestCountUniqueDevicesResponseResult]
type accountDexTestCountUniqueDevicesResponseResultJSON struct {
	UniqueDevicesTotal apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountDexTestCountUniqueDevicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestCountUniqueDevicesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponse struct {
	Errors   []Item `json:"errors,required"`
	Messages []Item `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDexTestListOverviewResponseSuccess `json:"success,required"`
	Result  AccountDexTestListOverviewResponseResult  `json:"result"`
	JSON    accountDexTestListOverviewResponseJSON    `json:"-"`
}

// accountDexTestListOverviewResponseJSON contains the JSON metadata for the struct
// [AccountDexTestListOverviewResponse]
type accountDexTestListOverviewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDexTestListOverviewResponseSuccess bool

const (
	AccountDexTestListOverviewResponseSuccessTrue AccountDexTestListOverviewResponseSuccess = true
)

func (r AccountDexTestListOverviewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDexTestListOverviewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDexTestListOverviewResponseResult struct {
	OverviewMetrics AccountDexTestListOverviewResponseResultOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []AccountDexTestListOverviewResponseResultTest `json:"tests,required"`
	JSON  accountDexTestListOverviewResponseResultJSON   `json:"-"`
}

// accountDexTestListOverviewResponseResultJSON contains the JSON metadata for the
// struct [AccountDexTestListOverviewResponseResult]
type accountDexTestListOverviewResponseResultJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all HTTP test results in response
	AvgHTTPAvailabilityPct float64 `json:"avgHttpAvailabilityPct,nullable"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                                                     `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         accountDexTestListOverviewResponseResultOverviewMetricsJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultOverviewMetricsJSON contains the JSON
// metadata for the struct
// [AccountDexTestListOverviewResponseResultOverviewMetrics]
type accountDexTestListOverviewResponseResultOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgHTTPAvailabilityPct       apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultOverviewMetricsJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultTest struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// date the test was created.
	Created string `json:"created,required"`
	// the test description defined during configuration
	Description string `json:"description,required"`
	// if true, then the test will run on targeted devices. Else, the test will not
	// run.
	Enabled bool   `json:"enabled,required"`
	Host    string `json:"host,required"`
	// The interval at which the synthetic application test is set to run.
	Interval string `json:"interval,required"`
	// test type, http or traceroute
	Kind AccountDexTestListOverviewResponseResultTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                                                           `json:"name,required"`
	Updated           string                                                           `json:"updated,required"`
	HTTPResults       AccountDexTestListOverviewResponseResultTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []AccountDexTestListOverviewResponseResultTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                                                                 `json:"method"`
	TargetPolicies          []AccountDexTestListOverviewResponseResultTestsTargetPolicy            `json:"target_policies,nullable"`
	Targeted                bool                                                                   `json:"targeted"`
	TracerouteResults       AccountDexTestListOverviewResponseResultTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []AccountDexTestListOverviewResponseResultTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    accountDexTestListOverviewResponseResultTestJSON                       `json:"-"`
}

// accountDexTestListOverviewResponseResultTestJSON contains the JSON metadata for
// the struct [AccountDexTestListOverviewResponseResultTest]
type accountDexTestListOverviewResponseResultTestJSON struct {
	ID                      apijson.Field
	Created                 apijson.Field
	Description             apijson.Field
	Enabled                 apijson.Field
	Host                    apijson.Field
	Interval                apijson.Field
	Kind                    apijson.Field
	Name                    apijson.Field
	Updated                 apijson.Field
	HTTPResults             apijson.Field
	HTTPResultsByColo       apijson.Field
	Method                  apijson.Field
	TargetPolicies          apijson.Field
	Targeted                apijson.Field
	TracerouteResults       apijson.Field
	TracerouteResultsByColo apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestJSON) RawJSON() string {
	return r.raw
}

// test type, http or traceroute
type AccountDexTestListOverviewResponseResultTestsKind string

const (
	AccountDexTestListOverviewResponseResultTestsKindHTTP       AccountDexTestListOverviewResponseResultTestsKind = "http"
	AccountDexTestListOverviewResponseResultTestsKindTraceroute AccountDexTestListOverviewResponseResultTestsKind = "traceroute"
)

func (r AccountDexTestListOverviewResponseResultTestsKind) IsKnown() bool {
	switch r {
	case AccountDexTestListOverviewResponseResultTestsKindHTTP, AccountDexTestListOverviewResponseResultTestsKindTraceroute:
		return true
	}
	return false
}

type AccountDexTestListOverviewResponseResultTestsHTTPResults struct {
	ResourceFetchTime TimingAggregates                                             `json:"resourceFetchTime,required"`
	JSON              accountDexTestListOverviewResponseResultTestsHTTPResultsJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultTestsHTTPResultsJSON contains the JSON
// metadata for the struct
// [AccountDexTestListOverviewResponseResultTestsHTTPResults]
type accountDexTestListOverviewResponseResultTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestsHTTPResultsJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                                             `json:"colo,required"`
	ResourceFetchTime TimingAggregates                                                   `json:"resourceFetchTime,required"`
	JSON              accountDexTestListOverviewResponseResultTestsHTTPResultsByColoJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultTestsHTTPResultsByColoJSON contains the
// JSON metadata for the struct
// [AccountDexTestListOverviewResponseResultTestsHTTPResultsByColo]
type accountDexTestListOverviewResponseResultTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestsHTTPResultsByColoJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultTestsTargetPolicy struct {
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                                                          `json:"default,required"`
	Name    string                                                        `json:"name,required"`
	JSON    accountDexTestListOverviewResponseResultTestsTargetPolicyJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultTestsTargetPolicyJSON contains the JSON
// metadata for the struct
// [AccountDexTestListOverviewResponseResultTestsTargetPolicy]
type accountDexTestListOverviewResponseResultTestsTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTestsTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestsTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultTestsTracerouteResults struct {
	RoundTripTime TimingAggregates                                                   `json:"roundTripTime,required"`
	JSON          accountDexTestListOverviewResponseResultTestsTracerouteResultsJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultTestsTracerouteResultsJSON contains the
// JSON metadata for the struct
// [AccountDexTestListOverviewResponseResultTestsTracerouteResults]
type accountDexTestListOverviewResponseResultTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestsTracerouteResultsJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestListOverviewResponseResultTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                                                   `json:"colo,required"`
	RoundTripTime TimingAggregates                                                         `json:"roundTripTime,required"`
	JSON          accountDexTestListOverviewResponseResultTestsTracerouteResultsByColoJSON `json:"-"`
}

// accountDexTestListOverviewResponseResultTestsTracerouteResultsByColoJSON
// contains the JSON metadata for the struct
// [AccountDexTestListOverviewResponseResultTestsTracerouteResultsByColo]
type accountDexTestListOverviewResponseResultTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDexTestListOverviewResponseResultTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexTestListOverviewResponseResultTestsTracerouteResultsByColoJSON) RawJSON() string {
	return r.raw
}

type AccountDexTestCountUniqueDevicesParams struct {
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [AccountDexTestCountUniqueDevicesParams]'s query parameters
// as `url.Values`.
func (r AccountDexTestCountUniqueDevicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDexTestListOverviewParams struct {
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [AccountDexTestListOverviewParams]'s query parameters as
// `url.Values`.
func (r AccountDexTestListOverviewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
