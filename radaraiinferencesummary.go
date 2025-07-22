// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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

type RadarAIInferenceSummaryGetModelResponseResultMeta struct {
	DateRange      []RadarAIInferenceSummaryGetModelResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	Normalization  string                                                          `json:"normalization,required"`
	ConfidenceInfo RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAIInferenceSummaryGetModelResponseResultMetaJSON           `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAIInferenceSummaryGetModelResponseResultMeta]
type radarAIInferenceSummaryGetModelResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetModelResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaJSON) RawJSON() string {
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

type RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetModelResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

type RadarAIInferenceSummaryGetTaskResponseResultMeta struct {
	DateRange      []RadarAIInferenceSummaryGetTaskResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAIInferenceSummaryGetTaskResponseResultMetaJSON           `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAIInferenceSummaryGetTaskResponseResultMeta]
type radarAIInferenceSummaryGetTaskResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIInferenceSummaryGetTaskResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaJSON) RawJSON() string {
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

type RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation]
type radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIInferenceSummaryGetTaskResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIInferenceSummaryGetModelParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceSummaryGetModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
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
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceSummaryGetTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
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
