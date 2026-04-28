// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDexHTTPTestService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDexHTTPTestService] method instead.
type AccountDexHTTPTestService struct {
	Options []option.RequestOption
}

// NewAccountDexHTTPTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDexHTTPTestService(opts ...option.RequestOption) (r *AccountDexHTTPTestService) {
	r = &AccountDexHTTPTestService{}
	r.Options = opts
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *AccountDexHTTPTestService) Get(ctx context.Context, accountID string, testID string, query AccountDexHTTPTestGetParams, opts ...option.RequestOption) (res *AccountDexHTTPTestGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *AccountDexHTTPTestService) GetPercentiles(ctx context.Context, accountID string, testID string, query AccountDexHTTPTestGetPercentilesParams, opts ...option.RequestOption) (res *AccountDexHTTPTestGetPercentilesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type Percentiles struct {
	// p50 observed in the time period
	P50 float64 `json:"p50" api:"nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90" api:"nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95" api:"nullable"`
	// p99 observed in the time period
	P99  float64         `json:"p99" api:"nullable"`
	JSON percentilesJSON `json:"-"`
}

// percentilesJSON contains the JSON metadata for the struct [Percentiles]
type percentilesJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Percentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentilesJSON) RawJSON() string {
	return r.raw
}

type TestStatOverTime struct {
	Slots []TestStatOverTimeSlot `json:"slots" api:"required"`
	// average observed in the time period
	Avg int64 `json:"avg" api:"nullable"`
	// highest observed in the time period
	Max int64 `json:"max" api:"nullable"`
	// lowest observed in the time period
	Min  int64                `json:"min" api:"nullable"`
	JSON testStatOverTimeJSON `json:"-"`
}

// testStatOverTimeJSON contains the JSON metadata for the struct
// [TestStatOverTime]
type testStatOverTimeJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestStatOverTimeSlot struct {
	Timestamp string                   `json:"timestamp" api:"required"`
	Value     int64                    `json:"value" api:"required"`
	JSON      testStatOverTimeSlotJSON `json:"-"`
}

// testStatOverTimeSlotJSON contains the JSON metadata for the struct
// [TestStatOverTimeSlot]
type testStatOverTimeSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatOverTimeSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatOverTimeSlotJSON) RawJSON() string {
	return r.raw
}

type TestStatPctOverTime struct {
	Slots []TestStatPctOverTimeSlot `json:"slots" api:"required"`
	// average observed in the time period
	Avg float64 `json:"avg" api:"nullable"`
	// highest observed in the time period
	Max float64 `json:"max" api:"nullable"`
	// lowest observed in the time period
	Min  float64                 `json:"min" api:"nullable"`
	JSON testStatPctOverTimeJSON `json:"-"`
}

// testStatPctOverTimeJSON contains the JSON metadata for the struct
// [TestStatPctOverTime]
type testStatPctOverTimeJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatPctOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatPctOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestStatPctOverTimeSlot struct {
	Timestamp string                      `json:"timestamp" api:"required"`
	Value     float64                     `json:"value" api:"required"`
	JSON      testStatPctOverTimeSlotJSON `json:"-"`
}

// testStatPctOverTimeSlotJSON contains the JSON metadata for the struct
// [TestStatPctOverTimeSlot]
type testStatPctOverTimeSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestStatPctOverTimeSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testStatPctOverTimeSlotJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponse struct {
	Errors   []Item `json:"errors" api:"required"`
	Messages []Item `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccountDexHTTPTestGetResponseSuccess `json:"success" api:"required"`
	Result  AccountDexHTTPTestGetResponseResult  `json:"result"`
	JSON    accountDexHTTPTestGetResponseJSON    `json:"-"`
}

// accountDexHTTPTestGetResponseJSON contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponse]
type accountDexHTTPTestGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDexHTTPTestGetResponseSuccess bool

const (
	AccountDexHTTPTestGetResponseSuccessTrue AccountDexHTTPTestGetResponseSuccess = true
)

func (r AccountDexHTTPTestGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDexHTTPTestGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDexHTTPTestGetResponseResult struct {
	// The url of the HTTP synthetic application test
	Host            string                                               `json:"host"`
	HTTPStats       AccountDexHTTPTestGetResponseResultHTTPStats         `json:"httpStats" api:"nullable"`
	HTTPStatsByColo []AccountDexHTTPTestGetResponseResultHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                                  `json:"interval"`
	Kind     AccountDexHTTPTestGetResponseResultKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name           string                                            `json:"name"`
	TargetPolicies []AccountDexHTTPTestGetResponseResultTargetPolicy `json:"target_policies" api:"nullable"`
	Targeted       bool                                              `json:"targeted"`
	JSON           accountDexHTTPTestGetResponseResultJSON           `json:"-"`
}

// accountDexHTTPTestGetResponseResultJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestGetResponseResult]
type accountDexHTTPTestGetResponseResultJSON struct {
	Host            apijson.Field
	HTTPStats       apijson.Field
	HTTPStatsByColo apijson.Field
	Interval        apijson.Field
	Kind            apijson.Field
	Method          apijson.Field
	Name            apijson.Field
	TargetPolicies  apijson.Field
	Targeted        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponseResultHTTPStats struct {
	AvailabilityPct      TestStatPctOverTime                                          `json:"availabilityPct" api:"required"`
	DNSResponseTimeMs    TestStatOverTime                                             `json:"dnsResponseTimeMs" api:"required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode `json:"httpStatusCode" api:"required"`
	ResourceFetchTimeMs  TestStatOverTime                                             `json:"resourceFetchTimeMs" api:"required"`
	ServerResponseTimeMs TestStatOverTime                                             `json:"serverResponseTimeMs" api:"required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                            `json:"uniqueDevicesTotal" api:"required"`
	JSON               accountDexHTTPTestGetResponseResultHTTPStatsJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsJSON contains the JSON metadata for
// the struct [AccountDexHTTPTestGetResponseResultHTTPStats]
type accountDexHTTPTestGetResponseResultHTTPStatsJSON struct {
	AvailabilityPct      apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultHTTPStatsJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode struct {
	Status200 int64                                                          `json:"status200" api:"required"`
	Status300 int64                                                          `json:"status300" api:"required"`
	Status400 int64                                                          `json:"status400" api:"required"`
	Status500 int64                                                          `json:"status500" api:"required"`
	Timestamp string                                                         `json:"timestamp" api:"required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON contains the JSON
// metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode]
type accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColo struct {
	AvailabilityPct      TestStatPctOverTime                                                `json:"availabilityPct" api:"required"`
	Colo                 string                                                             `json:"colo" api:"required"`
	DNSResponseTimeMs    TestStatOverTime                                                   `json:"dnsResponseTimeMs" api:"required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode `json:"httpStatusCode" api:"required"`
	ResourceFetchTimeMs  TestStatOverTime                                                   `json:"resourceFetchTimeMs" api:"required"`
	ServerResponseTimeMs TestStatOverTime                                                   `json:"serverResponseTimeMs" api:"required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                  `json:"uniqueDevicesTotal" api:"required"`
	JSON               accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON contains the JSON
// metadata for the struct [AccountDexHTTPTestGetResponseResultHTTPStatsByColo]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON struct {
	AvailabilityPct      apijson.Field
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultHTTPStatsByColoJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                                                `json:"status200" api:"required"`
	Status300 int64                                                                `json:"status300" api:"required"`
	Status400 int64                                                                `json:"status400" api:"required"`
	Status500 int64                                                                `json:"status500" api:"required"`
	Timestamp string                                                               `json:"timestamp" api:"required"`
	JSON      accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON contains
// the JSON metadata for the struct
// [AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode]
type accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetResponseResultKind string

const (
	AccountDexHTTPTestGetResponseResultKindHTTP AccountDexHTTPTestGetResponseResultKind = "http"
)

func (r AccountDexHTTPTestGetResponseResultKind) IsKnown() bool {
	switch r {
	case AccountDexHTTPTestGetResponseResultKindHTTP:
		return true
	}
	return false
}

type AccountDexHTTPTestGetResponseResultTargetPolicy struct {
	ID string `json:"id" api:"required"`
	// Whether the policy is the default for the account
	Default bool                                                `json:"default" api:"required"`
	Name    string                                              `json:"name" api:"required"`
	JSON    accountDexHTTPTestGetResponseResultTargetPolicyJSON `json:"-"`
}

// accountDexHTTPTestGetResponseResultTargetPolicyJSON contains the JSON metadata
// for the struct [AccountDexHTTPTestGetResponseResultTargetPolicy]
type accountDexHTTPTestGetResponseResultTargetPolicyJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetResponseResultTargetPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetResponseResultTargetPolicyJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetPercentilesResponse struct {
	Errors   []Item `json:"errors" api:"required"`
	Messages []Item `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccountDexHTTPTestGetPercentilesResponseSuccess `json:"success" api:"required"`
	Result  AccountDexHTTPTestGetPercentilesResponseResult  `json:"result"`
	JSON    accountDexHTTPTestGetPercentilesResponseJSON    `json:"-"`
}

// accountDexHTTPTestGetPercentilesResponseJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestGetPercentilesResponse]
type accountDexHTTPTestGetPercentilesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetPercentilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetPercentilesResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDexHTTPTestGetPercentilesResponseSuccess bool

const (
	AccountDexHTTPTestGetPercentilesResponseSuccessTrue AccountDexHTTPTestGetPercentilesResponseSuccess = true
)

func (r AccountDexHTTPTestGetPercentilesResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDexHTTPTestGetPercentilesResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDexHTTPTestGetPercentilesResponseResult struct {
	DNSResponseTimeMs    Percentiles                                        `json:"dnsResponseTimeMs"`
	ResourceFetchTimeMs  Percentiles                                        `json:"resourceFetchTimeMs"`
	ServerResponseTimeMs Percentiles                                        `json:"serverResponseTimeMs"`
	JSON                 accountDexHTTPTestGetPercentilesResponseResultJSON `json:"-"`
}

// accountDexHTTPTestGetPercentilesResponseResultJSON contains the JSON metadata
// for the struct [AccountDexHTTPTestGetPercentilesResponseResult]
type accountDexHTTPTestGetPercentilesResponseResultJSON struct {
	DNSResponseTimeMs    apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountDexHTTPTestGetPercentilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDexHTTPTestGetPercentilesResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDexHTTPTestGetParams struct {
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from" api:"required"`
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexHTTPTestGetParamsInterval] `query:"interval" api:"required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to" api:"required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [AccountDexHTTPTestGetParams]'s query parameters as
// `url.Values`.
func (r AccountDexHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Time interval for aggregate time slots.
type AccountDexHTTPTestGetParamsInterval string

const (
	AccountDexHTTPTestGetParamsIntervalMinute AccountDexHTTPTestGetParamsInterval = "minute"
	AccountDexHTTPTestGetParamsIntervalHour   AccountDexHTTPTestGetParamsInterval = "hour"
)

func (r AccountDexHTTPTestGetParamsInterval) IsKnown() bool {
	switch r {
	case AccountDexHTTPTestGetParamsIntervalMinute, AccountDexHTTPTestGetParamsIntervalHour:
		return true
	}
	return false
}

type AccountDexHTTPTestGetPercentilesParams struct {
	// Start time for the query in ISO (RFC3339 - ISO 8601) format
	From param.Field[string] `query:"from" api:"required"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format
	To param.Field[string] `query:"to" api:"required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [AccountDexHTTPTestGetPercentilesParams]'s query parameters
// as `url.Values`.
func (r AccountDexHTTPTestGetPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
