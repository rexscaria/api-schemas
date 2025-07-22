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

// RadarNetflowTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarNetflowTopService] method instead.
type RadarNetflowTopService struct {
	Options []option.RequestOption
}

// NewRadarNetflowTopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarNetflowTopService(opts ...option.RequestOption) (r *RadarNetflowTopService) {
	r = &RadarNetflowTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by network traffic (NetFlows).
func (r *RadarNetflowTopService) GetTopAs(ctx context.Context, query RadarNetflowTopGetTopAsParams, opts ...option.RequestOption) (res *RadarNetflowTopGetTopAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations by network traffic (NetFlows).
func (r *RadarNetflowTopService) GetTopLocations(ctx context.Context, query RadarNetflowTopGetTopLocationsParams, opts ...option.RequestOption) (res *RadarNetflowTopGetTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarNetflowTopGetTopAsResponse struct {
	Result  RadarNetflowTopGetTopAsResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarNetflowTopGetTopAsResponseJSON   `json:"-"`
}

// radarNetflowTopGetTopAsResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTopGetTopAsResponse]
type radarNetflowTopGetTopAsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsResponseResult struct {
	Top0 []RadarNetflowTopGetTopAsResponseResultTop0 `json:"top_0,required"`
	JSON radarNetflowTopGetTopAsResponseResultJSON   `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowTopGetTopAsResponseResult]
type radarNetflowTopGetTopAsResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsResponseResultTop0 struct {
	ClientAsn    float64                                       `json:"clientASN,required"`
	ClientAsName string                                        `json:"clientASName,required"`
	Value        string                                        `json:"value,required"`
	JSON         radarNetflowTopGetTopAsResponseResultTop0JSON `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarNetflowTopGetTopAsResponseResultTop0]
type radarNetflowTopGetTopAsResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponse struct {
	Result  RadarNetflowTopGetTopLocationsResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarNetflowTopGetTopLocationsResponseJSON   `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseJSON contains the JSON metadata for the
// struct [RadarNetflowTopGetTopLocationsResponse]
type radarNetflowTopGetTopLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponseResult struct {
	Top0 []RadarNetflowTopGetTopLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarNetflowTopGetTopLocationsResponseResultJSON   `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultJSON contains the JSON metadata for
// the struct [RadarNetflowTopGetTopLocationsResponseResult]
type radarNetflowTopGetTopLocationsResponseResultJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponseResultTop0 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarNetflowTopGetTopLocationsResponseResultTop0JSON `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarNetflowTopGetTopLocationsResponseResultTop0]
type radarNetflowTopGetTopLocationsResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsParams struct {
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
	Format param.Field[RadarNetflowTopGetTopAsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopGetTopAsParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopGetTopAsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarNetflowTopGetTopAsParamsFormat string

const (
	RadarNetflowTopGetTopAsParamsFormatJson RadarNetflowTopGetTopAsParamsFormat = "JSON"
	RadarNetflowTopGetTopAsParamsFormatCsv  RadarNetflowTopGetTopAsParamsFormat = "CSV"
)

func (r RadarNetflowTopGetTopAsParamsFormat) IsKnown() bool {
	switch r {
	case RadarNetflowTopGetTopAsParamsFormatJson, RadarNetflowTopGetTopAsParamsFormatCsv:
		return true
	}
	return false
}

type RadarNetflowTopGetTopLocationsParams struct {
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
	Format param.Field[RadarNetflowTopGetTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarNetflowTopGetTopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTopGetTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarNetflowTopGetTopLocationsParamsFormat string

const (
	RadarNetflowTopGetTopLocationsParamsFormatJson RadarNetflowTopGetTopLocationsParamsFormat = "JSON"
	RadarNetflowTopGetTopLocationsParamsFormatCsv  RadarNetflowTopGetTopLocationsParamsFormat = "CSV"
)

func (r RadarNetflowTopGetTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarNetflowTopGetTopLocationsParamsFormatJson, RadarNetflowTopGetTopLocationsParamsFormatCsv:
		return true
	}
	return false
}
