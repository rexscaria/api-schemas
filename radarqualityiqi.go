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

// RadarQualityIqiService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarQualityIqiService] method instead.
type RadarQualityIqiService struct {
	Options []option.RequestOption
}

// NewRadarQualityIqiService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarQualityIqiService(opts ...option.RequestOption) (r *RadarQualityIqiService) {
	r = &RadarQualityIqiService{}
	r.Options = opts
	return
}

// Retrieves a summary (percentiles) of bandwidth, latency, or DNS response time
// from the Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiService) GetSummary(ctx context.Context, query RadarQualityIqiGetSummaryParams, opts ...option.RequestOption) (res *RadarQualityIqiGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a time series (percentiles) of bandwidth, latency, or DNS response
// time from the Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiService) GetTimeseriesGroups(ctx context.Context, query RadarQualityIqiGetTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarQualityIqiGetTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualityIqiGetSummaryResponse struct {
	Result  RadarQualityIqiGetSummaryResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarQualityIqiGetSummaryResponseJSON   `json:"-"`
}

// radarQualityIqiGetSummaryResponseJSON contains the JSON metadata for the struct
// [RadarQualityIqiGetSummaryResponse]
type radarQualityIqiGetSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarQualityIqiGetSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarQualityIqiGetSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarQualityIqiGetSummaryResponseResultJSON     `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultJSON contains the JSON metadata for the
// struct [RadarQualityIqiGetSummaryResponseResult]
type radarQualityIqiGetSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarQualityIqiGetSummaryResponseResultMeta struct {
	ConfidenceInfo RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarQualityIqiGetSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarQualityIqiGetSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarQualityIqiGetSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarQualityIqiGetSummaryResponseResultMetaJSON   `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarQualityIqiGetSummaryResponseResultMeta]
type radarQualityIqiGetSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                         `json:"level,required"`
	JSON  radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfo]
type radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiGetSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarQualityIqiGetSummaryResponseResultMetaDateRange]
type radarQualityIqiGetSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarQualityIqiGetSummaryResponseResultMetaNormalization string

const (
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationPercentage           RadarQualityIqiGetSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationMin0Max              RadarQualityIqiGetSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationMinMax               RadarQualityIqiGetSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationRawValues            RadarQualityIqiGetSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationPercentageChange     RadarQualityIqiGetSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationRollingAverage       RadarQualityIqiGetSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationOverlappedPercentage RadarQualityIqiGetSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarQualityIqiGetSummaryResponseResultMetaNormalizationRatio                RadarQualityIqiGetSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarQualityIqiGetSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetSummaryResponseResultMetaNormalizationPercentage, RadarQualityIqiGetSummaryResponseResultMetaNormalizationMin0Max, RadarQualityIqiGetSummaryResponseResultMetaNormalizationMinMax, RadarQualityIqiGetSummaryResponseResultMetaNormalizationRawValues, RadarQualityIqiGetSummaryResponseResultMetaNormalizationPercentageChange, RadarQualityIqiGetSummaryResponseResultMetaNormalizationRollingAverage, RadarQualityIqiGetSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarQualityIqiGetSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarQualityIqiGetSummaryResponseResultMetaUnit struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  radarQualityIqiGetSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultMetaUnitJSON contains the JSON metadata
// for the struct [RadarQualityIqiGetSummaryResponseResultMetaUnit]
type radarQualityIqiGetSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetSummaryResponseResultSummary0 struct {
	P25  string                                              `json:"p25,required"`
	P50  string                                              `json:"p50,required"`
	P75  string                                              `json:"p75,required"`
	JSON radarQualityIqiGetSummaryResponseResultSummary0JSON `json:"-"`
}

// radarQualityIqiGetSummaryResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarQualityIqiGetSummaryResponseResultSummary0]
type radarQualityIqiGetSummaryResponseResultSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetTimeseriesGroupsResponse struct {
	Result  RadarQualityIqiGetTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarQualityIqiGetTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseJSON contains the JSON metadata for
// the struct [RadarQualityIqiGetTimeseriesGroupsResponse]
type radarQualityIqiGetTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetTimeseriesGroupsResponseResult struct {
	// Metadata for the results.
	Meta   RadarQualityIqiGetTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarQualityIqiGetTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarQualityIqiGetTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultJSON contains the JSON metadata
// for the struct [RadarQualityIqiGetTimeseriesGroupsResponseResult]
type radarQualityIqiGetTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarQualityIqiGetTimeseriesGroupsResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarQualityIqiGetTimeseriesGroupsResponseResultMetaUnit `json:"units,required"`
	JSON  radarQualityIqiGetTimeseriesGroupsResponseResultMetaJSON   `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarQualityIqiGetTimeseriesGroupsResponseResultMeta]
type radarQualityIqiGetTimeseriesGroupsResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval string

const (
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneHour        RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_HOUR"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneDay         RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_DAY"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneWeek        RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_WEEK"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneMonth       RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneHour, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneDay, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneWeek, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                  `json:"level,required"`
	JSON  radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfo]
type radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                        `json:"startDate,required" format:"date-time"`
	JSON            radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation]
type radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRange]
type radarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization string

const (
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationPercentage           RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationMin0Max              RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "MIN0_MAX"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationMinMax               RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "MIN_MAX"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRawValues            RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "RAW_VALUES"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationPercentageChange     RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRollingAverage       RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRatio                RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization = "RATIO"
)

func (r RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationPercentage, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationMin0Max, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationMinMax, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRawValues, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationPercentageChange, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRollingAverage, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage, RadarQualityIqiGetTimeseriesGroupsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarQualityIqiGetTimeseriesGroupsResponseResultMetaUnit struct {
	Name  string                                                       `json:"name,required"`
	Value string                                                       `json:"value,required"`
	JSON  radarQualityIqiGetTimeseriesGroupsResponseResultMetaUnitJSON `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarQualityIqiGetTimeseriesGroupsResponseResultMetaUnit]
type radarQualityIqiGetTimeseriesGroupsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetTimeseriesGroupsResponseResultSerie0 struct {
	P25        []string                                                   `json:"p25,required"`
	P50        []string                                                   `json:"p50,required"`
	P75        []string                                                   `json:"p75,required"`
	Timestamps []string                                                   `json:"timestamps,required"`
	JSON       radarQualityIqiGetTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarQualityIqiGetTimeseriesGroupsResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarQualityIqiGetTimeseriesGroupsResponseResultSerie0]
type radarQualityIqiGetTimeseriesGroupsResponseResultSerie0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiGetTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualityIqiGetTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarQualityIqiGetSummaryParams struct {
	// Defines which metric to return (bandwidth, latency, or DNS response time).
	Metric param.Field[RadarQualityIqiGetSummaryParamsMetric] `query:"metric,required"`
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
	Format param.Field[RadarQualityIqiGetSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiGetSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIqiGetSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Defines which metric to return (bandwidth, latency, or DNS response time).
type RadarQualityIqiGetSummaryParamsMetric string

const (
	RadarQualityIqiGetSummaryParamsMetricBandwidth RadarQualityIqiGetSummaryParamsMetric = "BANDWIDTH"
	RadarQualityIqiGetSummaryParamsMetricDNS       RadarQualityIqiGetSummaryParamsMetric = "DNS"
	RadarQualityIqiGetSummaryParamsMetricLatency   RadarQualityIqiGetSummaryParamsMetric = "LATENCY"
)

func (r RadarQualityIqiGetSummaryParamsMetric) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetSummaryParamsMetricBandwidth, RadarQualityIqiGetSummaryParamsMetricDNS, RadarQualityIqiGetSummaryParamsMetricLatency:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarQualityIqiGetSummaryParamsFormat string

const (
	RadarQualityIqiGetSummaryParamsFormatJson RadarQualityIqiGetSummaryParamsFormat = "JSON"
	RadarQualityIqiGetSummaryParamsFormatCsv  RadarQualityIqiGetSummaryParamsFormat = "CSV"
)

func (r RadarQualityIqiGetSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetSummaryParamsFormatJson, RadarQualityIqiGetSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarQualityIqiGetTimeseriesGroupsParams struct {
	// Defines which metric to return (bandwidth, latency, or DNS response time).
	Metric param.Field[RadarQualityIqiGetTimeseriesGroupsParamsMetric] `query:"metric,required"`
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarQualityIqiGetTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarQualityIqiGetTimeseriesGroupsParamsFormat] `query:"format"`
	// Enables interpolation for all series (using the average).
	Interpolation param.Field[bool] `query:"interpolation"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiGetTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RadarQualityIqiGetTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Defines which metric to return (bandwidth, latency, or DNS response time).
type RadarQualityIqiGetTimeseriesGroupsParamsMetric string

const (
	RadarQualityIqiGetTimeseriesGroupsParamsMetricBandwidth RadarQualityIqiGetTimeseriesGroupsParamsMetric = "BANDWIDTH"
	RadarQualityIqiGetTimeseriesGroupsParamsMetricDNS       RadarQualityIqiGetTimeseriesGroupsParamsMetric = "DNS"
	RadarQualityIqiGetTimeseriesGroupsParamsMetricLatency   RadarQualityIqiGetTimeseriesGroupsParamsMetric = "LATENCY"
)

func (r RadarQualityIqiGetTimeseriesGroupsParamsMetric) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetTimeseriesGroupsParamsMetricBandwidth, RadarQualityIqiGetTimeseriesGroupsParamsMetricDNS, RadarQualityIqiGetTimeseriesGroupsParamsMetricLatency:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarQualityIqiGetTimeseriesGroupsParamsAggInterval string

const (
	RadarQualityIqiGetTimeseriesGroupsParamsAggInterval15m RadarQualityIqiGetTimeseriesGroupsParamsAggInterval = "15m"
	RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1h  RadarQualityIqiGetTimeseriesGroupsParamsAggInterval = "1h"
	RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1d  RadarQualityIqiGetTimeseriesGroupsParamsAggInterval = "1d"
	RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1w  RadarQualityIqiGetTimeseriesGroupsParamsAggInterval = "1w"
)

func (r RadarQualityIqiGetTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetTimeseriesGroupsParamsAggInterval15m, RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1h, RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1d, RadarQualityIqiGetTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarQualityIqiGetTimeseriesGroupsParamsFormat string

const (
	RadarQualityIqiGetTimeseriesGroupsParamsFormatJson RadarQualityIqiGetTimeseriesGroupsParamsFormat = "JSON"
	RadarQualityIqiGetTimeseriesGroupsParamsFormatCsv  RadarQualityIqiGetTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarQualityIqiGetTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualityIqiGetTimeseriesGroupsParamsFormatJson, RadarQualityIqiGetTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}
