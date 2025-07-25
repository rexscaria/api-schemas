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

// RadarAttackLayer7Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7Service] method instead.
type RadarAttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer7SummaryService
	TimeseriesGroups *RadarAttackLayer7TimeseriesGroupService
	Top              *RadarAttackLayer7TopService
}

// NewRadarAttackLayer7Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7Service(opts ...option.RequestOption) (r *RadarAttackLayer7Service) {
	r = &RadarAttackLayer7Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer7TopService(opts...)
	return
}

// Retrieves layer 7 attacks over time.
func (r *RadarAttackLayer7Service) GetTimeseries(ctx context.Context, query RadarAttackLayer7GetTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7GetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7GetTimeseriesResponse struct {
	Result  RadarAttackLayer7GetTimeseriesResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAttackLayer7GetTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7GetTimeseriesResponse]
type radarAttackLayer7GetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7GetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7GetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7GetTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer7GetTimeseriesResponseResult]
type radarAttackLayer7GetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7GetTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7GetTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7GetTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7GetTimeseriesResponseResultMeta]
type radarAttackLayer7GetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7GetTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7GetTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7GetTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7GetTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAttackLayer7GetTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAttackLayer7GetTimeseriesResponseResultMetaUnit]
type radarAttackLayer7GetTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                            `json:"timestamps,required" format:"date-time"`
	Values     []string                                               `json:"values,required"`
	JSON       radarAttackLayer7GetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7GetTimeseriesResponseResultSerie0]
type radarAttackLayer7GetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7GetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7GetTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7GetTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7GetTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7GetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7GetTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7GetTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7GetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7GetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7GetTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7GetTimeseriesParamsAggInterval15m RadarAttackLayer7GetTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1h  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1d  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1w  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7GetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsAggInterval15m, RadarAttackLayer7GetTimeseriesParamsAggInterval1h, RadarAttackLayer7GetTimeseriesParamsAggInterval1d, RadarAttackLayer7GetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7GetTimeseriesParamsFormat string

const (
	RadarAttackLayer7GetTimeseriesParamsFormatJson RadarAttackLayer7GetTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7GetTimeseriesParamsFormatCsv  RadarAttackLayer7GetTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7GetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsFormatJson, RadarAttackLayer7GetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodGet             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPost            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPut             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodHead            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodACL             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodLock            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMove            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodReport          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodJson            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCook            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7GetTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsHTTPMethodGet, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPost, RadarAttackLayer7GetTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPut, RadarAttackLayer7GetTimeseriesParamsHTTPMethodHead, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7GetTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodACL, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7GetTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7GetTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7GetTimeseriesParamsHTTPMethodLock, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMove, RadarAttackLayer7GetTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7GetTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7GetTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodReport, RadarAttackLayer7GetTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7GetTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7GetTimeseriesParamsHTTPMethodJson, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCook, RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7GetTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7GetTimeseriesParamsIPVersionIPv4 RadarAttackLayer7GetTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7GetTimeseriesParamsIPVersionIPv6 RadarAttackLayer7GetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7GetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsIPVersionIPv4, RadarAttackLayer7GetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7GetTimeseriesParamsMitigationProductDdos               RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductWaf                RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7GetTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsMitigationProductDdos, RadarAttackLayer7GetTimeseriesParamsMitigationProductWaf, RadarAttackLayer7GetTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7GetTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7GetTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7GetTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7GetTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7GetTimeseriesParamsNormalization string

const (
	RadarAttackLayer7GetTimeseriesParamsNormalizationPercentageChange RadarAttackLayer7GetTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7GetTimeseriesParamsNormalizationMin0Max          RadarAttackLayer7GetTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7GetTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsNormalizationPercentageChange, RadarAttackLayer7GetTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}
