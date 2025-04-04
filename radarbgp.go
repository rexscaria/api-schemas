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

// RadarBgpService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpService] method instead.
type RadarBgpService struct {
	Options []option.RequestOption
	Hijacks *RadarBgpHijackService
	IPs     *RadarBgpIPService
	Leaks   *RadarBgpLeakService
	Routes  *RadarBgpRouteService
	Top     *RadarBgpTopService
}

// NewRadarBgpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpService(opts ...option.RequestOption) (r *RadarBgpService) {
	r = &RadarBgpService{}
	r.Options = opts
	r.Hijacks = NewRadarBgpHijackService(opts...)
	r.IPs = NewRadarBgpIPService(opts...)
	r.Leaks = NewRadarBgpLeakService(opts...)
	r.Routes = NewRadarBgpRouteService(opts...)
	r.Top = NewRadarBgpTopService(opts...)
	return
}

// Retrieves BGP updates over time. When requesting updates for an autonomous
// system, only BGP updates of type announcement are returned.
func (r *RadarBgpService) GetTimeseries(ctx context.Context, query RadarBgpGetTimeseriesParams, opts ...option.RequestOption) (res *RadarBgpGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpGetTimeseriesResponse struct {
	Result  RadarBgpGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpGetTimeseriesResponseJSON   `json:"-"`
}

// radarBgpGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarBgpGetTimeseriesResponse]
type radarBgpGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResult struct {
	Meta   RadarBgpGetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarBgpGetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarBgpGetTimeseriesResponseResultJSON   `json:"-"`
}

// radarBgpGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpGetTimeseriesResponseResult]
type radarBgpGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResultMeta struct {
	AggInterval    string                                                `json:"aggInterval,required"`
	DateRange      []RadarBgpGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                             `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarBgpGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarBgpGetTimeseriesResponseResultMetaJSON           `json:"-"`
}

// radarBgpGetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpGetTimeseriesResponseResultMeta]
type radarBgpGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarBgpGetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpGetTimeseriesResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarBgpGetTimeseriesResponseResultMetaDateRange]
type radarBgpGetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarBgpGetTimeseriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarBgpGetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarBgpGetTimeseriesResponseResultMetaConfidenceInfo]
type radarBgpGetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                   `json:"timestamps,required" format:"date-time"`
	Values     []string                                      `json:"values,required"`
	JSON       radarBgpGetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarBgpGetTimeseriesResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarBgpGetTimeseriesResponseResultSerie0]
type radarBgpGetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpGetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpGetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarBgpGetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarBgpGetTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpGetTimeseriesParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpGetTimeseriesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarBgpGetTimeseriesParamsAggInterval string

const (
	RadarBgpGetTimeseriesParamsAggInterval15m RadarBgpGetTimeseriesParamsAggInterval = "15m"
	RadarBgpGetTimeseriesParamsAggInterval1h  RadarBgpGetTimeseriesParamsAggInterval = "1h"
	RadarBgpGetTimeseriesParamsAggInterval1d  RadarBgpGetTimeseriesParamsAggInterval = "1d"
	RadarBgpGetTimeseriesParamsAggInterval1w  RadarBgpGetTimeseriesParamsAggInterval = "1w"
)

func (r RadarBgpGetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarBgpGetTimeseriesParamsAggInterval15m, RadarBgpGetTimeseriesParamsAggInterval1h, RadarBgpGetTimeseriesParamsAggInterval1d, RadarBgpGetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarBgpGetTimeseriesParamsFormat string

const (
	RadarBgpGetTimeseriesParamsFormatJson RadarBgpGetTimeseriesParamsFormat = "JSON"
	RadarBgpGetTimeseriesParamsFormatCsv  RadarBgpGetTimeseriesParamsFormat = "CSV"
)

func (r RadarBgpGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpGetTimeseriesParamsFormatJson, RadarBgpGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpGetTimeseriesParamsUpdateType string

const (
	RadarBgpGetTimeseriesParamsUpdateTypeAnnouncement RadarBgpGetTimeseriesParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpGetTimeseriesParamsUpdateTypeWithdrawal   RadarBgpGetTimeseriesParamsUpdateType = "WITHDRAWAL"
)

func (r RadarBgpGetTimeseriesParamsUpdateType) IsKnown() bool {
	switch r {
	case RadarBgpGetTimeseriesParamsUpdateTypeAnnouncement, RadarBgpGetTimeseriesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}
