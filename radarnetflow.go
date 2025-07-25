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
	// Metadata for the results.
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

// Metadata for the results.
type RadarNetflowGetSummaryResponseResultMeta struct {
	ConfidenceInfo RadarNetflowGetSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarNetflowGetSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarNetflowGetSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarNetflowGetSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarNetflowGetSummaryResponseResultMetaJSON   `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarNetflowGetSummaryResponseResultMeta]
type radarNetflowGetSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  radarNetflowGetSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowGetSummaryResponseResultMetaNormalization string

const (
	RadarNetflowGetSummaryResponseResultMetaNormalizationPercentage           RadarNetflowGetSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarNetflowGetSummaryResponseResultMetaNormalizationMin0Max              RadarNetflowGetSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarNetflowGetSummaryResponseResultMetaNormalizationMinMax               RadarNetflowGetSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarNetflowGetSummaryResponseResultMetaNormalizationRawValues            RadarNetflowGetSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarNetflowGetSummaryResponseResultMetaNormalizationPercentageChange     RadarNetflowGetSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowGetSummaryResponseResultMetaNormalizationRollingAverage       RadarNetflowGetSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarNetflowGetSummaryResponseResultMetaNormalizationOverlappedPercentage RadarNetflowGetSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarNetflowGetSummaryResponseResultMetaNormalizationRatio                RadarNetflowGetSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarNetflowGetSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarNetflowGetSummaryResponseResultMetaNormalizationPercentage, RadarNetflowGetSummaryResponseResultMetaNormalizationMin0Max, RadarNetflowGetSummaryResponseResultMetaNormalizationMinMax, RadarNetflowGetSummaryResponseResultMetaNormalizationRawValues, RadarNetflowGetSummaryResponseResultMetaNormalizationPercentageChange, RadarNetflowGetSummaryResponseResultMetaNormalizationRollingAverage, RadarNetflowGetSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarNetflowGetSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarNetflowGetSummaryResponseResultMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarNetflowGetSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarNetflowGetSummaryResponseResultMetaUnitJSON contains the JSON metadata for
// the struct [RadarNetflowGetSummaryResponseResultMetaUnit]
type radarNetflowGetSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowGetSummaryResponseResultSummary0 struct {
	// A numeric string.
	HTTP string `json:"HTTP,required"`
	// A numeric string.
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
	// Metadata for the results.
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

// Metadata for the results.
type RadarNetflowGetTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarNetflowGetTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarNetflowGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarNetflowGetTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarNetflowGetTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarNetflowGetTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarNetflowGetTimeseriesResponseResultMeta]
type radarNetflowGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarNetflowGetTimeseriesResponseResultMetaAggInterval string

const (
	RadarNetflowGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarNetflowGetTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneHour        RadarNetflowGetTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneDay         RadarNetflowGetTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneWeek        RadarNetflowGetTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneMonth       RadarNetflowGetTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarNetflowGetTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneHour, RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneDay, RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneWeek, RadarNetflowGetTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowGetTimeseriesResponseResultMetaNormalization string

const (
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationPercentage           RadarNetflowGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationMin0Max              RadarNetflowGetTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationMinMax               RadarNetflowGetTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationRawValues            RadarNetflowGetTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationPercentageChange     RadarNetflowGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationRollingAverage       RadarNetflowGetTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarNetflowGetTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarNetflowGetTimeseriesResponseResultMetaNormalizationRatio                RadarNetflowGetTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarNetflowGetTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarNetflowGetTimeseriesResponseResultMetaNormalizationPercentage, RadarNetflowGetTimeseriesResponseResultMetaNormalizationMin0Max, RadarNetflowGetTimeseriesResponseResultMetaNormalizationMinMax, RadarNetflowGetTimeseriesResponseResultMetaNormalizationRawValues, RadarNetflowGetTimeseriesResponseResultMetaNormalizationPercentageChange, RadarNetflowGetTimeseriesResponseResultMetaNormalizationRollingAverage, RadarNetflowGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarNetflowGetTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarNetflowGetTimeseriesResponseResultMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  radarNetflowGetTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarNetflowGetTimeseriesResponseResultMetaUnitJSON contains the JSON metadata
// for the struct [RadarNetflowGetTimeseriesResponseResultMetaUnit]
type radarNetflowGetTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowGetTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowGetTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
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
	Format param.Field[RadarNetflowGetSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarNetflowGetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarNetflowGetTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarNetflowGetTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by network traffic product types.
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

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
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

// Normalization method applied to the results. Refer to
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
