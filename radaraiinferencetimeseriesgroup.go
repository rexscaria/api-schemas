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
	Meta   interface{}                                                 `json:"meta,required"`
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

type RadarAIInferenceTimeseriesGroupGetModelResponseResultSerie0 struct {
	Timestamps  []string                                                        `json:"timestamps,required"`
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
	Meta   interface{}                                                `json:"meta,required"`
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

type RadarAIInferenceTimeseriesGroupGetTaskResponseResultSerie0 struct {
	Timestamps  []string                                                       `json:"timestamps,required"`
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
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIInferenceTimeseriesGroupGetModelParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceTimeseriesGroupGetModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
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

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
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
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIInferenceTimeseriesGroupGetTaskParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAIInferenceTimeseriesGroupGetTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
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

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
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
