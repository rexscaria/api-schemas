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
	Meta   interface{}                                               `json:"meta,required"`
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

type RadarAIBotTimeseriesGroupGetUserAgentResponseResultSerie0 struct {
	Timestamps  []string                                                      `json:"timestamps,required"`
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
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAIBotTimeseriesGroupGetUserAgentParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAIBotTimeseriesGroupGetUserAgentParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
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

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
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
