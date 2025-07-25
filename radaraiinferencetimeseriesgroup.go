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

// RadarAIInferenceTimeseriesGroupService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAIInferenceTimeseriesGroupService] method instead.
type RadarAIInferenceTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAIInferenceTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAIInferenceTimeseriesGroupService(opts ...option.RequestOption) (r *RadarAIInferenceTimeseriesGroupService) {
	r = &RadarAIInferenceTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model over time.
func (r *RadarAIInferenceTimeseriesGroupService) GetModel(ctx context.Context, query RadarAIInferenceTimeseriesGroupGetModelParams, opts ...option.RequestOption) (res *RadarAIInferenceTimeseriesGroupGetModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of unique accounts by task over time.
func (r *RadarAIInferenceTimeseriesGroupService) GetTask(ctx context.Context, query RadarAIInferenceTimeseriesGroupGetTaskParams, opts ...option.RequestOption) (res *RadarAIInferenceTimeseriesGroupGetTaskResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAIInferenceTimeseriesGroupGetModelResponse struct {
	Result  RadarAIInferenceTimeseriesGroupGetModelResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAIInferenceTimeseriesGroupGetModelResponseJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseJSON contains the JSON metadata
// for the struct [RadarAIInferenceTimeseriesGroupGetModelResponse]
type radarAIInferenceTimeseriesGroupGetModelResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetModelResponseResult struct {
	// Metadata for the results.
	Meta   RadarAIInferenceTimeseriesGroupGetModelResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAIInferenceTimeseriesGroupGetModelResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAIInferenceTimeseriesGroupGetModelResponseResultJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultJSON contains the JSON
// metadata for the struct [RadarAIInferenceTimeseriesGroupGetModelResponseResult]
type radarAIInferenceTimeseriesGroupGetModelResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIInferenceTimeseriesGroupGetModelResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIInferenceTimeseriesGroupGetModelResponseResultMetaJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultMeta]
type radarAIInferenceTimeseriesGroupGetModelResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval string

const (
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalFifteenMinutes RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneHour        RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneDay         RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval = "ONE_DAY"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneWeek        RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneMonth       RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalFifteenMinutes, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneHour, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneDay, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneWeek, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                       `json:"level,required"`
	JSON  radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfo]
type radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRange]
type radarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization string

const (
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationPercentage           RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationMin0Max              RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationMinMax               RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "MIN_MAX"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRawValues            RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationPercentageChange     RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRollingAverage       RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationOverlappedPercentage RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRatio                RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationPercentage, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationMin0Max, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationMinMax, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRawValues, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationPercentageChange, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRollingAverage, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationOverlappedPercentage, RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnitJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnit]
type radarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetModelResponseResultSerie0 struct {
	Timestamps  []time.Time                                                     `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                             `json:"-,extras"`
	JSON        radarAIInferenceTimeseriesGroupGetModelResponseResultSerie0JSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetModelResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetModelResponseResultSerie0]
type radarAIInferenceTimeseriesGroupGetModelResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetModelResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetModelResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetTaskResponse struct {
	Result  RadarAIInferenceTimeseriesGroupGetTaskResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAIInferenceTimeseriesGroupGetTaskResponseJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseJSON contains the JSON metadata
// for the struct [RadarAIInferenceTimeseriesGroupGetTaskResponse]
type radarAIInferenceTimeseriesGroupGetTaskResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetTaskResponseResult struct {
	// Metadata for the results.
	Meta   RadarAIInferenceTimeseriesGroupGetTaskResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAIInferenceTimeseriesGroupGetTaskResponseResultJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultJSON contains the JSON
// metadata for the struct [RadarAIInferenceTimeseriesGroupGetTaskResponseResult]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaJSON   `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultMeta]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval string

const (
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalFifteenMinutes RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneHour        RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneDay         RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval = "ONE_DAY"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneWeek        RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneMonth       RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalFifteenMinutes, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneHour, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneDay, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneWeek, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfo]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRange]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization string

const (
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationPercentage           RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationMin0Max              RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationMinMax               RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "MIN_MAX"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRawValues            RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationPercentageChange     RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRollingAverage       RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationOverlappedPercentage RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRatio                RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationPercentage, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationMin0Max, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationMinMax, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRawValues, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationPercentageChange, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRollingAverage, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationOverlappedPercentage, RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnitJSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnit]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0 struct {
	Timestamps  []time.Time                                                    `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                            `json:"-,extras"`
	JSON        radarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0JSON `json:"-"`
}

// radarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0]
type radarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceTimeseriesGroupGetModelParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceTimeseriesGroupGetModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIInferenceTimeseriesGroupGetModelParams]'s query
// parameters as `url.Values`.
func (r RadarAIInferenceTimeseriesGroupGetModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval string

const (
	RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval15m RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval = "15m"
	RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1h  RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval = "1h"
	RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1d  RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval = "1d"
	RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1w  RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval = "1w"
)

func (r RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval15m, RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1h, RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1d, RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAIInferenceTimeseriesGroupGetModelParamsFormat string

const (
	RadarAIInferenceTimeseriesGroupGetModelParamsFormatJson RadarAIInferenceTimeseriesGroupGetModelParamsFormat = "JSON"
	RadarAIInferenceTimeseriesGroupGetModelParamsFormatCsv  RadarAIInferenceTimeseriesGroupGetModelParamsFormat = "CSV"
)

func (r RadarAIInferenceTimeseriesGroupGetModelParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetModelParamsFormatJson, RadarAIInferenceTimeseriesGroupGetModelParamsFormatCsv:
		return true
	}
	return false
}

type RadarAIInferenceTimeseriesGroupGetTaskParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceTimeseriesGroupGetTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIInferenceTimeseriesGroupGetTaskParams]'s query
// parameters as `url.Values`.
func (r RadarAIInferenceTimeseriesGroupGetTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval string

const (
	RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval15m RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval = "15m"
	RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1h  RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval = "1h"
	RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1d  RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval = "1d"
	RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1w  RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval = "1w"
)

func (r RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval15m, RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1h, RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1d, RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAIInferenceTimeseriesGroupGetTaskParamsFormat string

const (
	RadarAIInferenceTimeseriesGroupGetTaskParamsFormatJson RadarAIInferenceTimeseriesGroupGetTaskParamsFormat = "JSON"
	RadarAIInferenceTimeseriesGroupGetTaskParamsFormatCsv  RadarAIInferenceTimeseriesGroupGetTaskParamsFormat = "CSV"
)

func (r RadarAIInferenceTimeseriesGroupGetTaskParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIInferenceTimeseriesGroupGetTaskParamsFormatJson, RadarAIInferenceTimeseriesGroupGetTaskParamsFormatCsv:
		return true
	}
	return false
}
