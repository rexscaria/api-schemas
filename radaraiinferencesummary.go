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

// RadarAIInferenceSummaryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAIInferenceSummaryService] method instead.
type RadarAIInferenceSummaryService struct {
	Options []option.RequestOption
}

// NewRadarAIInferenceSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAIInferenceSummaryService(opts ...option.RequestOption) (r *RadarAIInferenceSummaryService) {
	r = &RadarAIInferenceSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model.
func (r *RadarAIInferenceSummaryService) GetModel(ctx context.Context, query RadarAIInferenceSummaryGetModelParams, opts ...option.RequestOption) (res *RadarAIInferenceSummaryGetModelResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of unique accounts by task.
func (r *RadarAIInferenceSummaryService) GetTask(ctx context.Context, query RadarAIInferenceSummaryGetTaskParams, opts ...option.RequestOption) (res *RadarAIInferenceSummaryGetTaskResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAIInferenceSummaryGetModelResponse struct {
	Result  RadarAIInferenceSummaryGetModelResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAIInferenceSummaryGetModelResponseJSON   `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseJSON contains the JSON metadata for the
// struct [RadarAIInferenceSummaryGetModelResponse]
type radarAIInferenceSummaryGetModelResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetModelResponseResult struct {
	// Metadata for the results.
	Meta     RadarAIInferenceSummaryGetModelResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                 `json:"summary_0,required"`
	JSON     radarAIInferenceSummaryGetModelResponseResultJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultJSON contains the JSON metadata for
// the struct [RadarAIInferenceSummaryGetModelResponseResult]
type radarAIInferenceSummaryGetModelResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIInferenceSummaryGetModelResponseResultMeta struct {
	ConfidenceInfo RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIInferenceSummaryGetModelResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIInferenceSummaryGetModelResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIInferenceSummaryGetModelResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIInferenceSummaryGetModelResponseResultMetaJSON   `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAIInferenceSummaryGetModelResponseResultMeta]
type radarAIInferenceSummaryGetModelResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo]
type radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetModelResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarAIInferenceSummaryGetModelResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAIInferenceSummaryGetModelResponseResultMetaDateRange]
type radarAIInferenceSummaryGetModelResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIInferenceSummaryGetModelResponseResultMetaNormalization string

const (
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationPercentage           RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationMin0Max              RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationMinMax               RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "MIN_MAX"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRawValues            RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationPercentageChange     RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRollingAverage       RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationOverlappedPercentage RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRatio                RadarAIInferenceSummaryGetModelResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIInferenceSummaryGetModelResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationPercentage, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationMin0Max, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationMinMax, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRawValues, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationPercentageChange, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRollingAverage, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationOverlappedPercentage, RadarAIInferenceSummaryGetModelResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIInferenceSummaryGetModelResponseResultMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  radarAIInferenceSummaryGetModelResponseResultMetaUnitJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAIInferenceSummaryGetModelResponseResultMetaUnit]
type radarAIInferenceSummaryGetModelResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetTaskResponse struct {
	Result  RadarAIInferenceSummaryGetTaskResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAIInferenceSummaryGetTaskResponseJSON   `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseJSON contains the JSON metadata for the
// struct [RadarAIInferenceSummaryGetTaskResponse]
type radarAIInferenceSummaryGetTaskResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetTaskResponseResult struct {
	// Metadata for the results.
	Meta     RadarAIInferenceSummaryGetTaskResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                `json:"summary_0,required"`
	JSON     radarAIInferenceSummaryGetTaskResponseResultJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultJSON contains the JSON metadata for
// the struct [RadarAIInferenceSummaryGetTaskResponseResult]
type radarAIInferenceSummaryGetTaskResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIInferenceSummaryGetTaskResponseResultMeta struct {
	ConfidenceInfo RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIInferenceSummaryGetTaskResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIInferenceSummaryGetTaskResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIInferenceSummaryGetTaskResponseResultMetaJSON   `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAIInferenceSummaryGetTaskResponseResultMeta]
type radarAIInferenceSummaryGetTaskResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo]
type radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetTaskResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAIInferenceSummaryGetTaskResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAIInferenceSummaryGetTaskResponseResultMetaDateRange]
type radarAIInferenceSummaryGetTaskResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization string

const (
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationPercentage           RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationMin0Max              RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationMinMax               RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "MIN_MAX"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRawValues            RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationPercentageChange     RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRollingAverage       RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationOverlappedPercentage RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRatio                RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIInferenceSummaryGetTaskResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationPercentage, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationMin0Max, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationMinMax, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRawValues, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationPercentageChange, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRollingAverage, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationOverlappedPercentage, RadarAIInferenceSummaryGetTaskResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIInferenceSummaryGetTaskResponseResultMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAIInferenceSummaryGetTaskResponseResultMetaUnitJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAIInferenceSummaryGetTaskResponseResultMetaUnit]
type radarAIInferenceSummaryGetTaskResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetModelParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceSummaryGetModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIInferenceSummaryGetModelParams]'s query parameters
// as `url.Values`.
func (r RadarAIInferenceSummaryGetModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAIInferenceSummaryGetModelParamsFormat string

const (
	RadarAIInferenceSummaryGetModelParamsFormatJson RadarAIInferenceSummaryGetModelParamsFormat = "JSON"
	RadarAIInferenceSummaryGetModelParamsFormatCsv  RadarAIInferenceSummaryGetModelParamsFormat = "CSV"
)

func (r RadarAIInferenceSummaryGetModelParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIInferenceSummaryGetModelParamsFormatJson, RadarAIInferenceSummaryGetModelParamsFormatCsv:
		return true
	}
	return false
}

type RadarAIInferenceSummaryGetTaskParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceSummaryGetTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIInferenceSummaryGetTaskParams]'s query parameters as
// `url.Values`.
func (r RadarAIInferenceSummaryGetTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAIInferenceSummaryGetTaskParamsFormat string

const (
	RadarAIInferenceSummaryGetTaskParamsFormatJson RadarAIInferenceSummaryGetTaskParamsFormat = "JSON"
	RadarAIInferenceSummaryGetTaskParamsFormatCsv  RadarAIInferenceSummaryGetTaskParamsFormat = "CSV"
)

func (r RadarAIInferenceSummaryGetTaskParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIInferenceSummaryGetTaskParamsFormatJson, RadarAIInferenceSummaryGetTaskParamsFormatCsv:
		return true
	}
	return false
}
