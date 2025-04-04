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
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get percentiles for an http test for a given time period between 1 hour and 7
// days.
func (r *AccountDexHTTPTestService) GetPercentiles(ctx context.Context, accountID string, testID string, query AccountDexHTTPTestGetPercentilesParams, opts ...option.RequestOption) (res *AccountDexHTTPTestGetPercentilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s/percentiles", accountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIResponseSingleDigitalExperience struct {
	Errors   []Item `json:"errors,required"`
	Messages []Item `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleDigitalExperienceSuccess `json:"success,required"`
	JSON    apiResponseSingleDigitalExperienceJSON    `json:"-"`
}

// apiResponseSingleDigitalExperienceJSON contains the JSON metadata for the struct
// [APIResponseSingleDigitalExperience]
type apiResponseSingleDigitalExperienceJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleDigitalExperience) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleDigitalExperienceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleDigitalExperienceSuccess bool

const (
	APIResponseSingleDigitalExperienceSuccessTrue APIResponseSingleDigitalExperienceSuccess = true
)

func (r APIResponseSingleDigitalExperienceSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleDigitalExperienceSuccessTrue:
		return true
	}
	return false
}

type Percentiles struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64         `json:"p99,nullable"`
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
	Slots []TestStatOverTimeSlot `json:"slots,required"`
	// average observed in the time period
	Avg int64 `json:"avg,nullable"`
	// highest observed in the time period
	Max int64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  int64                `json:"min,nullable"`
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
	Timestamp string                   `json:"timestamp,required"`
	Value     int64                    `json:"value,required"`
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
	Slots []TestStatPctOverTimeSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                 `json:"min,nullable"`
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
	Timestamp string                      `json:"timestamp,required"`
	Value     float64                     `json:"value,required"`
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
	Result AccountDexHTTPTestGetResponseResult `json:"result"`
	JSON   accountDexHTTPTestGetResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexHTTPTestGetResponseJSON contains the JSON metadata for the struct
// [AccountDexHTTPTestGetResponse]
type accountDexHTTPTestGetResponseJSON struct {
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

type AccountDexHTTPTestGetResponseResult struct {
	// The url of the HTTP synthetic application test
	Host            string                                               `json:"host"`
	HTTPStats       AccountDexHTTPTestGetResponseResultHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []AccountDexHTTPTestGetResponseResultHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string                                  `json:"interval"`
	Kind     AccountDexHTTPTestGetResponseResultKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name           string                                            `json:"name"`
	TargetPolicies []AccountDexHTTPTestGetResponseResultTargetPolicy `json:"target_policies,nullable"`
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
	AvailabilityPct      TestStatPctOverTime                                          `json:"availabilityPct,required"`
	DNSResponseTimeMs    TestStatOverTime                                             `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsHTTPStatusCode `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  TestStatOverTime                                             `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs TestStatOverTime                                             `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                            `json:"uniqueDevicesTotal,required"`
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
	Status200 int64                                                          `json:"status200,required"`
	Status300 int64                                                          `json:"status300,required"`
	Status400 int64                                                          `json:"status400,required"`
	Status500 int64                                                          `json:"status500,required"`
	Timestamp string                                                         `json:"timestamp,required"`
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
	AvailabilityPct      TestStatPctOverTime                                                `json:"availabilityPct,required"`
	Colo                 string                                                             `json:"colo,required"`
	DNSResponseTimeMs    TestStatOverTime                                                   `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []AccountDexHTTPTestGetResponseResultHTTPStatsByColoHTTPStatusCode `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  TestStatOverTime                                                   `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs TestStatOverTime                                                   `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                                                  `json:"uniqueDevicesTotal,required"`
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
	Status200 int64                                                                `json:"status200,required"`
	Status300 int64                                                                `json:"status300,required"`
	Status400 int64                                                                `json:"status400,required"`
	Status500 int64                                                                `json:"status500,required"`
	Timestamp string                                                               `json:"timestamp,required"`
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
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                                                `json:"default,required"`
	Name    string                                              `json:"name,required"`
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
	Result AccountDexHTTPTestGetPercentilesResponseResult `json:"result"`
	JSON   accountDexHTTPTestGetPercentilesResponseJSON   `json:"-"`
	APIResponseSingleDigitalExperience
}

// accountDexHTTPTestGetPercentilesResponseJSON contains the JSON metadata for the
// struct [AccountDexHTTPTestGetPercentilesResponse]
type accountDexHTTPTestGetPercentilesResponseJSON struct {
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
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[AccountDexHTTPTestGetParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
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

// URLQuery serializes [AccountDexHTTPTestGetPercentilesParams]'s query parameters
// as `url.Values`.
func (r AccountDexHTTPTestGetPercentilesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
