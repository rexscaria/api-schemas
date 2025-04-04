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

// RadarTcpResetsTimeoutService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarTcpResetsTimeoutService] method instead.
type RadarTcpResetsTimeoutService struct {
	Options []option.RequestOption
}

// NewRadarTcpResetsTimeoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarTcpResetsTimeoutService(opts ...option.RequestOption) (r *RadarTcpResetsTimeoutService) {
	r = &RadarTcpResetsTimeoutService{}
	r.Options = opts
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout.
func (r *RadarTcpResetsTimeoutService) Summary(ctx context.Context, query RadarTcpResetsTimeoutSummaryParams, opts ...option.RequestOption) (res *RadarTcpResetsTimeoutSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout over time.
func (r *RadarTcpResetsTimeoutService) TimeseriesGroups(ctx context.Context, query RadarTcpResetsTimeoutTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarTcpResetsTimeoutTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarTcpResetsTimeoutSummaryResponse struct {
	Result  RadarTcpResetsTimeoutSummaryResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarTcpResetsTimeoutSummaryResponseJSON   `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseJSON contains the JSON metadata for the
// struct [RadarTcpResetsTimeoutSummaryResponse]
type radarTcpResetsTimeoutSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResult struct {
	Meta     RadarTcpResetsTimeoutSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarTcpResetsTimeoutSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarTcpResetsTimeoutSummaryResponseResultJSON     `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultJSON contains the JSON metadata for
// the struct [RadarTcpResetsTimeoutSummaryResponseResult]
type radarTcpResetsTimeoutSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMeta struct {
	DateRange      []RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarTcpResetsTimeoutSummaryResponseResultMetaJSON           `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarTcpResetsTimeoutSummaryResponseResultMeta]
type radarTcpResetsTimeoutSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange]
type radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo]
type radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultSummary0 struct {
	// Connection resets within the first 10 packets from the client, but after the
	// server has received multiple data packets.
	LaterInFlow string `json:"later_in_flow,required"`
	// All other connections.
	NoMatch string `json:"no_match,required"`
	// Connection resets or timeouts after the server received both a SYN packet and an
	// ACK packet, meaning the connection was successfully established.
	PostAck string `json:"post_ack,required"`
	// Connection resets or timeouts after the server received a packet with PSH flag
	// set, following connection establishment.
	PostPsh string `json:"post_psh,required"`
	// Connection resets or timeouts after the server received only a single SYN
	// packet.
	PostSyn string                                                 `json:"post_syn,required"`
	JSON    radarTcpResetsTimeoutSummaryResponseResultSummary0JSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarTcpResetsTimeoutSummaryResponseResultSummary0]
type radarTcpResetsTimeoutSummaryResponseResultSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponse struct {
	Result  RadarTcpResetsTimeoutTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarTcpResetsTimeoutTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseJSON contains the JSON metadata for
// the struct [RadarTcpResetsTimeoutTimeseriesGroupsResponse]
type radarTcpResetsTimeoutTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResult struct {
	Meta   RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON contains the JSON
// metadata for the struct [RadarTcpResetsTimeoutTimeseriesGroupsResponseResult]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta struct {
	AggInterval    string                                                                `json:"aggInterval,required"`
	DateRange      []RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                             `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON           `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                             `json:"level"`
	JSON        radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                              `json:"dataSource,required"`
	Description     string                                                                              `json:"description,required"`
	EventType       string                                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                                           `json:"startTime" format:"date-time"`
	JSON            radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0 struct {
	LaterInFlow []string                                                      `json:"later_in_flow,required"`
	NoMatch     []string                                                      `json:"no_match,required"`
	PostAck     []string                                                      `json:"post_ack,required"`
	PostPsh     []string                                                      `json:"post_psh,required"`
	PostSyn     []string                                                      `json:"post_syn,required"`
	Timestamps  []time.Time                                                   `json:"timestamps,required" format:"date-time"`
	JSON        radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryParams struct {
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
	Format param.Field[RadarTcpResetsTimeoutSummaryParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarTcpResetsTimeoutSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarTcpResetsTimeoutSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarTcpResetsTimeoutSummaryParamsFormat string

const (
	RadarTcpResetsTimeoutSummaryParamsFormatJson RadarTcpResetsTimeoutSummaryParamsFormat = "JSON"
	RadarTcpResetsTimeoutSummaryParamsFormatCsv  RadarTcpResetsTimeoutSummaryParamsFormat = "CSV"
)

func (r RadarTcpResetsTimeoutSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutSummaryParamsFormatJson, RadarTcpResetsTimeoutSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarTcpResetsTimeoutTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarTcpResetsTimeoutTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RadarTcpResetsTimeoutTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval15m RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "15m"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1h  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1h"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1d  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1d"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1w  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1w"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval15m, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1h, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1d, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatJson RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat = "JSON"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatCsv  RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatJson, RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}
