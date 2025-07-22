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

// RadarBgpIPService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpIPService] method instead.
type RadarBgpIPService struct {
	Options []option.RequestOption
}

// NewRadarBgpIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpIPService(opts ...option.RequestOption) (r *RadarBgpIPService) {
	r = &RadarBgpIPService{}
	r.Options = opts
	return
}

// Retrieves time series data for the announced IP space count, represented as the
// number of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *RadarBgpIPService) GetTimeseries(ctx context.Context, query RadarBgpIPGetTimeseriesParams, opts ...option.RequestOption) (res *RadarBgpIPGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/ips/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpIPGetTimeseriesResponse struct {
	Result  RadarBgpIPGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarBgpIPGetTimeseriesResponseJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarBgpIPGetTimeseriesResponse]
type radarBgpIPGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResult struct {
	Meta     RadarBgpIPGetTimeseriesResponseResultMeta     `json:"meta,required"`
	Serie174 RadarBgpIPGetTimeseriesResponseResultSerie174 `json:"serie_174,required"`
	JSON     radarBgpIPGetTimeseriesResponseResultJSON     `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpIPGetTimeseriesResponseResult]
type radarBgpIPGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie174    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMeta struct {
	Queries []RadarBgpIPGetTimeseriesResponseResultMetaQuery `json:"queries,required"`
	JSON    radarBgpIPGetTimeseriesResponseResultMetaJSON    `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpIPGetTimeseriesResponseResultMeta]
type radarBgpIPGetTimeseriesResponseResultMetaJSON struct {
	Queries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaQuery struct {
	DateRange RadarBgpIPGetTimeseriesResponseResultMetaQueriesDateRange `json:"dateRange,required"`
	Entity    string                                                    `json:"entity,required"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaQueryJSON        `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaQueryJSON contains the JSON metadata
// for the struct [RadarBgpIPGetTimeseriesResponseResultMetaQuery]
type radarBgpIPGetTimeseriesResponseResultMetaQueryJSON struct {
	DateRange   apijson.Field
	Entity      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaQueryJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaQueriesDateRange struct {
	EndTime   string                                                        `json:"endTime,required"`
	StartTime string                                                        `json:"startTime,required"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaQueriesDateRangeJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaQueriesDateRangeJSON contains the JSON
// metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaQueriesDateRange]
type radarBgpIPGetTimeseriesResponseResultMetaQueriesDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaQueriesDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaQueriesDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultSerie174 struct {
	Ipv4       []string                                          `json:"ipv4,required"`
	Ipv6       []string                                          `json:"ipv6,required"`
	Timestamps []time.Time                                       `json:"timestamps,required" format:"date-time"`
	JSON       radarBgpIPGetTimeseriesResponseResultSerie174JSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultSerie174JSON contains the JSON metadata for
// the struct [RadarBgpIPGetTimeseriesResponseResultSerie174]
type radarBgpIPGetTimeseriesResponseResultSerie174JSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultSerie174) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultSerie174JSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesParams struct {
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
	Format param.Field[RadarBgpIPGetTimeseriesParamsFormat] `query:"format"`
	// Include data delay meta information.
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarBgpIPGetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarBgpIPGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpIPGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpIPGetTimeseriesParamsFormat string

const (
	RadarBgpIPGetTimeseriesParamsFormatJson RadarBgpIPGetTimeseriesParamsFormat = "JSON"
	RadarBgpIPGetTimeseriesParamsFormatCsv  RadarBgpIPGetTimeseriesParamsFormat = "CSV"
)

func (r RadarBgpIPGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesParamsFormatJson, RadarBgpIPGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpIPGetTimeseriesParamsIPVersion string

const (
	RadarBgpIPGetTimeseriesParamsIPVersionIPv4 RadarBgpIPGetTimeseriesParamsIPVersion = "IPv4"
	RadarBgpIPGetTimeseriesParamsIPVersionIPv6 RadarBgpIPGetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarBgpIPGetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesParamsIPVersionIPv4, RadarBgpIPGetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}
