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

// RadarNetflowService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarNetflowService] method instead.
type RadarNetflowService struct {
	Options []option.RequestOption
	Top     *RadarNetflowTopService
}

// NewRadarNetflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarNetflowService(opts ...option.RequestOption) (r *RadarNetflowService) {
	r = &RadarNetflowService{}
	r.Options = opts
	r.Top = NewRadarNetflowTopService(opts...)
	return
}

// Retrieves the distribution of network traffic (NetFlows) by HTTP vs other
// protocols.
func (r *RadarNetflowService) GetSummary(ctx context.Context, query RadarNetflowGetSummaryParams, opts ...option.RequestOption) (res *RadarNetflowGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves network traffic (NetFlows) over time.
func (r *RadarNetflowService) GetTimeseries(ctx context.Context, query RadarNetflowGetTimeseriesParams, opts ...option.RequestOption) (res *RadarNetflowGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarNetflowGetSummaryResponse struct {
	Result  RadarNetflowGetSummaryResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarNetflowGetSummaryResponseJSON   `json:"-"`
}

// radarNetflowGetSummaryResponseJSON contains the JSON metadata for the struct
// [RadarNetflowGetSummaryResponse]
type radarNetflowGetSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResult struct {
	Meta     RadarNetflowGetSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarNetflowGetSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarNetflowGetSummaryResponseResultJSON     `json:"-"`
}

// radarNetflowGetSummaryResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowGetSummaryResponseResult]
type radarNetflowGetSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultMeta struct {
	DateRange      []RadarNetflowGetSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarNetflowGetSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarNetflowGetSummaryResponseResultMetaJSON           `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarNetflowGetSummaryResponseResultMeta]
type radarNetflowGetSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowGetSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarNetflowGetSummaryResponseResultMetaDateRange]
type radarNetflowGetSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarNetflowGetSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarNetflowGetSummaryResponseResultMetaConfidenceInfo]
type radarNetflowGetSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultSummary0 struct {
	HTTP  string                                           `json:"HTTP,required"`
	Other string                                           `json:"OTHER,required"`
	JSON  radarNetflowGetSummaryResponseResultSummary0JSON `json:"-"`
}

// radarNetflowGetSummaryResponseResultSummary0JSON contains the JSON metadata for
// the struct [RadarNetflowGetSummaryResponseResultSummary0]
type radarNetflowGetSummaryResponseResultSummary0JSON struct {
	HTTP        apijson.Field
	Other       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponse struct {
	Result  RadarNetflowGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarNetflowGetTimeseriesResponseJSON   `json:"-"`
}

// radarNetflowGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarNetflowGetTimeseriesResponse]
type radarNetflowGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResult struct {
	Meta   RadarNetflowGetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarNetflowGetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarNetflowGetTimeseriesResponseResultJSON   `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowGetTimeseriesResponseResult]
type radarNetflowGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResultMeta struct {
	AggInterval    string                                                    `json:"aggInterval,required"`
	DateRange      []RadarNetflowGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                 `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarNetflowGetTimeseriesResponseResultMetaJSON           `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarNetflowGetTimeseriesResponseResultMeta]
type radarNetflowGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowGetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarNetflowGetTimeseriesResponseResultMetaDateRange]
type radarNetflowGetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo]
type radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                       `json:"timestamps,required" format:"date-time"`
	Values     []string                                          `json:"values,required"`
	JSON       radarNetflowGetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarNetflowGetTimeseriesResponseResultSerie0]
type radarNetflowGetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarNetflowGetSummaryParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowGetSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowGetSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarNetflowGetSummaryParamsFormat string

const (
	RadarNetflowGetSummaryParamsFormatJson RadarNetflowGetSummaryParamsFormat = "JSON"
	RadarNetflowGetSummaryParamsFormatCsv  RadarNetflowGetSummaryParamsFormat = "CSV"
)

func (r RadarNetflowGetSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarNetflowGetSummaryParamsFormatJson, RadarNetflowGetSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarNetflowGetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarNetflowGetTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarNetflowGetTimeseriesParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarNetflowGetTimeseriesParamsNormalization] `query:"normalization"`
	// Array of network traffic product types to filters results by.
	Product param.Field[[]RadarNetflowGetTimeseriesParamsProduct] `query:"product"`
}

// URLQuery serializes [RadarNetflowGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarNetflowGetTimeseriesParamsAggInterval string

const (
	RadarNetflowGetTimeseriesParamsAggInterval15m RadarNetflowGetTimeseriesParamsAggInterval = "15m"
	RadarNetflowGetTimeseriesParamsAggInterval1h  RadarNetflowGetTimeseriesParamsAggInterval = "1h"
	RadarNetflowGetTimeseriesParamsAggInterval1d  RadarNetflowGetTimeseriesParamsAggInterval = "1d"
	RadarNetflowGetTimeseriesParamsAggInterval1w  RadarNetflowGetTimeseriesParamsAggInterval = "1w"
)

func (r RadarNetflowGetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesParamsAggInterval15m, RadarNetflowGetTimeseriesParamsAggInterval1h, RadarNetflowGetTimeseriesParamsAggInterval1d, RadarNetflowGetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarNetflowGetTimeseriesParamsFormat string

const (
	RadarNetflowGetTimeseriesParamsFormatJson RadarNetflowGetTimeseriesParamsFormat = "JSON"
	RadarNetflowGetTimeseriesParamsFormatCsv  RadarNetflowGetTimeseriesParamsFormat = "CSV"
)

func (r RadarNetflowGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesParamsFormatJson, RadarNetflowGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowGetTimeseriesParamsNormalization string

const (
	RadarNetflowGetTimeseriesParamsNormalizationPercentageChange RadarNetflowGetTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowGetTimeseriesParamsNormalizationMin0Max          RadarNetflowGetTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarNetflowGetTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesParamsNormalizationPercentageChange, RadarNetflowGetTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarNetflowGetTimeseriesParamsProduct string

const (
	RadarNetflowGetTimeseriesParamsProductHTTP RadarNetflowGetTimeseriesParamsProduct = "HTTP"
	RadarNetflowGetTimeseriesParamsProductAll  RadarNetflowGetTimeseriesParamsProduct = "ALL"
)

func (r RadarNetflowGetTimeseriesParamsProduct) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesParamsProductHTTP, RadarNetflowGetTimeseriesParamsProductAll:
		return true
	}
	return false
}
