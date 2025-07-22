// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/tidwall/gjson"
)

// ZoneSpectrumAnalyticsEventService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpectrumAnalyticsEventService] method instead.
type ZoneSpectrumAnalyticsEventService struct {
	Options []option.RequestOption
}

// NewZoneSpectrumAnalyticsEventService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSpectrumAnalyticsEventService(opts ...option.RequestOption) (r *ZoneSpectrumAnalyticsEventService) {
	r = &ZoneSpectrumAnalyticsEventService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
func (r *ZoneSpectrumAnalyticsEventService) GetByTime(ctx context.Context, zoneID string, query ZoneSpectrumAnalyticsEventGetByTimeParams, opts ...option.RequestOption) (res *QueryResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/bytime", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
func (r *ZoneSpectrumAnalyticsEventService) GetSummary(ctx context.Context, zoneID string, query ZoneSpectrumAnalyticsEventGetSummaryParams, opts ...option.RequestOption) (res *QueryResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/summary", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DimensionItem string

const (
	DimensionItemEvent     DimensionItem = "event"
	DimensionItemAppID     DimensionItem = "appID"
	DimensionItemColoName  DimensionItem = "coloName"
	DimensionItemIPVersion DimensionItem = "ipVersion"
)

func (r DimensionItem) IsKnown() bool {
	switch r {
	case DimensionItemEvent, DimensionItemAppID, DimensionItemColoName, DimensionItemIPVersion:
		return true
	}
	return false
}

type MetricItem string

const (
	MetricItemCount          MetricItem = "count"
	MetricItemBytesIngress   MetricItem = "bytesIngress"
	MetricItemBytesEgress    MetricItem = "bytesEgress"
	MetricItemDurationAvg    MetricItem = "durationAvg"
	MetricItemDurationMedian MetricItem = "durationMedian"
	MetricItemDuration90th   MetricItem = "duration90th"
	MetricItemDuration99th   MetricItem = "duration99th"
)

func (r MetricItem) IsKnown() bool {
	switch r {
	case MetricItemCount, MetricItemBytesIngress, MetricItemBytesEgress, MetricItemDurationAvg, MetricItemDurationMedian, MetricItemDuration90th, MetricItemDuration99th:
		return true
	}
	return false
}

type QueryResponseSingle struct {
	Result QueryResponseSingleResult `json:"result"`
	JSON   queryResponseSingleJSON   `json:"-"`
	APIResponseSingleSpectrumAnalytics
}

// queryResponseSingleJSON contains the JSON metadata for the struct
// [QueryResponseSingle]
type queryResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResponseSingleJSON) RawJSON() string {
	return r.raw
}

type QueryResponseSingleResult struct {
	// List of columns returned by the analytics query.
	Data []QueryResponseSingleResultData `json:"data,required"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum result for each selected metrics across all data.
	Max map[string]float64 `json:"max,required"`
	// Minimum result for each selected metrics across all data.
	Min   map[string]float64             `json:"min,required"`
	Query QueryResponseSingleResultQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total result for each selected metrics across all data.
	Totals map[string]float64 `json:"totals,required"`
	// List of time interval buckets: [start, end]
	TimeIntervals [][]time.Time                 `json:"time_intervals" format:"date-time"`
	JSON          queryResponseSingleResultJSON `json:"-"`
}

// queryResponseSingleResultJSON contains the JSON metadata for the struct
// [QueryResponseSingleResult]
type queryResponseSingleResultJSON struct {
	Data          apijson.Field
	DataLag       apijson.Field
	Max           apijson.Field
	Min           apijson.Field
	Query         apijson.Field
	Rows          apijson.Field
	Totals        apijson.Field
	TimeIntervals apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *QueryResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

type QueryResponseSingleResultData struct {
	Dimensions []string                                  `json:"dimensions"`
	Metrics    QueryResponseSingleResultDataMetricsUnion `json:"metrics"`
	JSON       queryResponseSingleResultDataJSON         `json:"-"`
}

// queryResponseSingleResultDataJSON contains the JSON metadata for the struct
// [QueryResponseSingleResultData]
type queryResponseSingleResultDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResponseSingleResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResponseSingleResultDataJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [QueryResponseSingleResultDataMetricsArray] or
// [QueryResponseSingleResultDataMetricsArray].
type QueryResponseSingleResultDataMetricsUnion interface {
	implementsQueryResponseSingleResultDataMetricsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*QueryResponseSingleResultDataMetricsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QueryResponseSingleResultDataMetricsArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(QueryResponseSingleResultDataMetricsArray{}),
		},
	)
}

type QueryResponseSingleResultDataMetricsArray []float64

func (r QueryResponseSingleResultDataMetricsArray) implementsQueryResponseSingleResultDataMetricsUnion() {
}

type QueryResponseSingleResultQuery struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions []DimensionItem `json:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | \>       | Greater Than             | %3E         |
	// | \<       | Less Than                | %3C         |
	// | \>=      | Greater than or equal to | %3E%3D      |
	// | \<=      | Less than or equal to    | %3C%3D      |
	Filters string `json:"filters"`
	// Limit number of returned metrics.
	Limit float64 `json:"limit"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics []MetricItem `json:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since Since `json:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort []string `json:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until Until                              `json:"until" format:"date-time"`
	JSON  queryResponseSingleResultQueryJSON `json:"-"`
}

// queryResponseSingleResultQueryJSON contains the JSON metadata for the struct
// [QueryResponseSingleResultQuery]
type queryResponseSingleResultQueryJSON struct {
	Dimensions  apijson.Field
	Filters     apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Sort        apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResponseSingleResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResponseSingleResultQueryJSON) RawJSON() string {
	return r.raw
}

type SinceParam = time.Time

type UntilParam = time.Time

type ZoneSpectrumAnalyticsEventGetByTimeParams struct {
	// Used to select time series resolution.
	TimeDelta param.Field[ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta] `query:"time_delta,required"`
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]DimensionItem] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | \>       | Greater Than             | %3E         |
	// | \<       | Less Than                | %3C         |
	// | \>=      | Greater than or equal to | %3E%3D      |
	// | \<=      | Less than or equal to    | %3C%3D      |
	Filters param.Field[string] `query:"filters"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics param.Field[[]MetricItem] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[SinceParam] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]string] `query:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[UntilParam] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneSpectrumAnalyticsEventGetByTimeParams]'s query
// parameters as `url.Values`.
func (r ZoneSpectrumAnalyticsEventGetByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Used to select time series resolution.
type ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta string

const (
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaYear       ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "year"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaQuarter    ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "quarter"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaMonth      ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "month"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaWeek       ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "week"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaDay        ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "day"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaHour       ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "hour"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaDekaminute ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "dekaminute"
	ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaMinute     ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta = "minute"
)

func (r ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDelta) IsKnown() bool {
	switch r {
	case ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaYear, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaQuarter, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaMonth, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaWeek, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaDay, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaHour, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaDekaminute, ZoneSpectrumAnalyticsEventGetByTimeParamsTimeDeltaMinute:
		return true
	}
	return false
}

type ZoneSpectrumAnalyticsEventGetSummaryParams struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]DimensionItem] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | \>       | Greater Than             | %3E         |
	// | \<       | Less Than                | %3C         |
	// | \>=      | Greater than or equal to | %3E%3D      |
	// | \<=      | Less than or equal to    | %3C%3D      |
	Filters param.Field[string] `query:"filters"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics param.Field[[]MetricItem] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[SinceParam] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]string] `query:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[UntilParam] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneSpectrumAnalyticsEventGetSummaryParams]'s query
// parameters as `url.Values`.
func (r ZoneSpectrumAnalyticsEventGetSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
