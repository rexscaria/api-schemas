// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDNSFirewallDNSAnalyticReportService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallDNSAnalyticReportService] method instead.
type AccountDNSFirewallDNSAnalyticReportService struct {
	Options []option.RequestOption
}

// NewAccountDNSFirewallDNSAnalyticReportService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountDNSFirewallDNSAnalyticReportService(opts ...option.RequestOption) (r *AccountDNSFirewallDNSAnalyticReportService) {
	r = &AccountDNSFirewallDNSAnalyticReportService{}
	r.Options = opts
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AccountDNSFirewallDNSAnalyticReportService) Get(ctx context.Context, accountID string, dnsFirewallID string, query AccountDNSFirewallDNSAnalyticReportGetParams, opts ...option.RequestOption) (res *AccountDNSFirewallDNSAnalyticReportGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AccountDNSFirewallDNSAnalyticReportService) ListByTime(ctx context.Context, accountID string, dnsFirewallID string, query AccountDNSFirewallDNSAnalyticReportListByTimeParams, opts ...option.RequestOption) (res *AccountDNSFirewallDNSAnalyticReportListByTimeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type APIResponseSingleDNSAnalytics struct {
	Errors   []MessagesDNSAnalyticsItem `json:"errors,required"`
	Messages []MessagesDNSAnalyticsItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleDNSAnalyticsSuccess `json:"success,required"`
	JSON    apiResponseSingleDNSAnalyticsJSON    `json:"-"`
}

// apiResponseSingleDNSAnalyticsJSON contains the JSON metadata for the struct
// [APIResponseSingleDNSAnalytics]
type apiResponseSingleDNSAnalyticsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleDNSAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleDNSAnalyticsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleDNSAnalyticsSuccess bool

const (
	APIResponseSingleDNSAnalyticsSuccessTrue APIResponseSingleDNSAnalyticsSuccess = true
)

func (r APIResponseSingleDNSAnalyticsSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleDNSAnalyticsSuccessTrue:
		return true
	}
	return false
}

type DataReport struct {
	Data []DataReportData `json:"data,required"`
	JSON dataReportJSON   `json:"-"`
	Result
}

// dataReportJSON contains the JSON metadata for the struct [DataReport]
type dataReportJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataReportJSON) RawJSON() string {
	return r.raw
}

type DataReportData struct {
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64          `json:"metrics,required"`
	JSON    dataReportDataJSON `json:"-"`
}

// dataReportDataJSON contains the JSON metadata for the struct [DataReportData]
type dataReportDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DataReportData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dataReportDataJSON) RawJSON() string {
	return r.raw
}

type MessagesDNSAnalyticsItem struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    messagesDNSAnalyticsItemJSON `json:"-"`
}

// messagesDNSAnalyticsItemJSON contains the JSON metadata for the struct
// [MessagesDNSAnalyticsItem]
type messagesDNSAnalyticsItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDNSAnalyticsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDNSAnalyticsItemJSON) RawJSON() string {
	return r.raw
}

type ReportByTime struct {
	Data  []ReportByTimeData `json:"data,required"`
	Query ReportByTimeQuery  `json:"query,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time    `json:"time_intervals,required" format:"date-time"`
	JSON          reportByTimeJSON `json:"-"`
	Result
}

// reportByTimeJSON contains the JSON metadata for the struct [ReportByTime]
type reportByTimeJSON struct {
	Data          apijson.Field
	Query         apijson.Field
	TimeIntervals apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ReportByTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportByTimeJSON) RawJSON() string {
	return r.raw
}

type ReportByTimeData struct {
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}      `json:"metrics,required"`
	JSON    reportByTimeDataJSON `json:"-"`
}

// reportByTimeDataJSON contains the JSON metadata for the struct
// [ReportByTimeData]
type reportByTimeDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportByTimeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportByTimeDataJSON) RawJSON() string {
	return r.raw
}

type ReportByTimeQuery struct {
	// Unit of time to group data by.
	TimeDelta TimeDelta             `json:"time_delta,required"`
	JSON      reportByTimeQueryJSON `json:"-"`
}

// reportByTimeQueryJSON contains the JSON metadata for the struct
// [ReportByTimeQuery]
type reportByTimeQueryJSON struct {
	TimeDelta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportByTimeQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportByTimeQueryJSON) RawJSON() string {
	return r.raw
}

type Result struct {
	// Array with one row per combination of dimension values.
	Data []ResultData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{} `json:"min,required"`
	Query ResultQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{} `json:"totals,required"`
	JSON   resultJSON  `json:"-"`
}

// resultJSON contains the JSON metadata for the struct [Result]
type resultJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultJSON) RawJSON() string {
	return r.raw
}

type ResultData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string       `json:"dimensions,required"`
	JSON       resultDataJSON `json:"-"`
}

// resultDataJSON contains the JSON metadata for the struct [ResultData]
type resultDataJSON struct {
	Dimensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultDataJSON) RawJSON() string {
	return r.raw
}

type ResultQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string        `json:"sort"`
	JSON resultQueryJSON `json:"-"`
}

// resultQueryJSON contains the JSON metadata for the struct [ResultQuery]
type resultQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r resultQueryJSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type TimeDelta string

const (
	TimeDeltaAll        TimeDelta = "all"
	TimeDeltaAuto       TimeDelta = "auto"
	TimeDeltaYear       TimeDelta = "year"
	TimeDeltaQuarter    TimeDelta = "quarter"
	TimeDeltaMonth      TimeDelta = "month"
	TimeDeltaWeek       TimeDelta = "week"
	TimeDeltaDay        TimeDelta = "day"
	TimeDeltaHour       TimeDelta = "hour"
	TimeDeltaDekaminute TimeDelta = "dekaminute"
	TimeDeltaMinute     TimeDelta = "minute"
)

func (r TimeDelta) IsKnown() bool {
	switch r {
	case TimeDeltaAll, TimeDeltaAuto, TimeDeltaYear, TimeDeltaQuarter, TimeDeltaMonth, TimeDeltaWeek, TimeDeltaDay, TimeDeltaHour, TimeDeltaDekaminute, TimeDeltaMinute:
		return true
	}
	return false
}

type AccountDNSFirewallDNSAnalyticReportGetResponse struct {
	Result DataReport                                         `json:"result"`
	JSON   accountDNSFirewallDNSAnalyticReportGetResponseJSON `json:"-"`
	APIResponseSingleDNSAnalytics
}

// accountDNSFirewallDNSAnalyticReportGetResponseJSON contains the JSON metadata
// for the struct [AccountDNSFirewallDNSAnalyticReportGetResponse]
type accountDNSFirewallDNSAnalyticReportGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallDNSAnalyticReportGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallDNSAnalyticReportListByTimeResponse struct {
	Result ReportByTime                                              `json:"result"`
	JSON   accountDNSFirewallDNSAnalyticReportListByTimeResponseJSON `json:"-"`
	APIResponseSingleDNSAnalytics
}

// accountDNSFirewallDNSAnalyticReportListByTimeResponseJSON contains the JSON
// metadata for the struct [AccountDNSFirewallDNSAnalyticReportListByTimeResponse]
type accountDNSFirewallDNSAnalyticReportListByTimeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListByTimeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallDNSAnalyticReportListByTimeResponseJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallDNSAnalyticReportGetParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountDNSFirewallDNSAnalyticReportGetParams]'s query
// parameters as `url.Values`.
func (r AccountDNSFirewallDNSAnalyticReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountDNSFirewallDNSAnalyticReportListByTimeParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// Unit of time to group data by.
	TimeDelta param.Field[TimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountDNSFirewallDNSAnalyticReportListByTimeParams]'s
// query parameters as `url.Values`.
func (r AccountDNSFirewallDNSAnalyticReportListByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
