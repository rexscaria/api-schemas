// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarAIBotTimeseriesGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAIBotTimeseriesGroupService] method instead.
type RadarAIBotTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAIBotTimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAIBotTimeseriesGroupService(opts ...option.RequestOption) (r *RadarAIBotTimeseriesGroupService) {
	r = &RadarAIBotTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of traffic by AI user agent over time.
func (r *RadarAIBotTimeseriesGroupService) GetUserAgent(ctx context.Context, query RadarAIBotTimeseriesGroupGetUserAgentParams, opts ...option.RequestOption) (res *RadarAIBotTimeseriesGroupGetUserAgentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries_groups/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAIBotTimeseriesGroupGetUserAgentResponse struct {
	Result  RadarAIBotTimeseriesGroupGetUserAgentResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAIBotTimeseriesGroupGetUserAgentResponseJSON   `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseJSON contains the JSON metadata for
// the struct [RadarAIBotTimeseriesGroupGetUserAgentResponse]
type radarAIBotTimeseriesGroupGetUserAgentResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotTimeseriesGroupGetUserAgentResponseResult struct {
	// Metadata for the results.
	Meta   RadarAIBotTimeseriesGroupGetUserAgentResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAIBotTimeseriesGroupGetUserAgentResponseResultJSON   `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultJSON contains the JSON
// metadata for the struct [RadarAIBotTimeseriesGroupGetUserAgentResponseResult]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaJSON   `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultMeta]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval string

const (
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalFifteenMinutes RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneHour        RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneDay         RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval = "ONE_DAY"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneWeek        RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneMonth       RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalFifteenMinutes, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneHour, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneDay, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneWeek, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                     `json:"level,required"`
	JSON  radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfo]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotation]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRange]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization string

const (
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationPercentage           RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationMin0Max              RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationMinMax               RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "MIN_MAX"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRawValues            RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationPercentageChange     RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRollingAverage       RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationOverlappedPercentage RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRatio                RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationPercentage, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationMin0Max, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationMinMax, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRawValues, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationPercentageChange, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRollingAverage, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationOverlappedPercentage, RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnit struct {
	Name  string                                                          `json:"name,required"`
	Value string                                                          `json:"value,required"`
	JSON  radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnitJSON `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnit]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0 struct {
	Timestamps  []time.Time                                                   `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                           `json:"-,extras"`
	JSON        radarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0JSON `json:"-"`
}

// radarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0]
type radarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAIBotTimeseriesGroupGetUserAgentParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIBotTimeseriesGroupGetUserAgentParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIBotTimeseriesGroupGetUserAgentParams]'s query
// parameters as `url.Values`.
func (r RadarAIBotTimeseriesGroupGetUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval string

const (
	RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval15m RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval = "15m"
	RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1h  RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval = "1h"
	RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1d  RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval = "1d"
	RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1w  RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval = "1w"
)

func (r RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval15m, RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1h, RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1d, RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAIBotTimeseriesGroupGetUserAgentParamsFormat string

const (
	RadarAIBotTimeseriesGroupGetUserAgentParamsFormatJson RadarAIBotTimeseriesGroupGetUserAgentParamsFormat = "JSON"
	RadarAIBotTimeseriesGroupGetUserAgentParamsFormatCsv  RadarAIBotTimeseriesGroupGetUserAgentParamsFormat = "CSV"
)

func (r RadarAIBotTimeseriesGroupGetUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIBotTimeseriesGroupGetUserAgentParamsFormatJson, RadarAIBotTimeseriesGroupGetUserAgentParamsFormatCsv:
		return true
	}
	return false
}
