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

// RadarDNSTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarDNSTopService] method instead.
type RadarDNSTopService struct {
	Options []option.RequestOption
}

// NewRadarDNSTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDNSTopService(opts ...option.RequestOption) (r *RadarDNSTopService) {
	r = &RadarDNSTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by DNS queries made to 1.1.1.1 DNS
// resolver.
func (r *RadarDNSTopService) GetTopAses(ctx context.Context, query RadarDNSTopGetTopAsesParams, opts ...option.RequestOption) (res *RadarDNSTopGetTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations by DNS queries made to 1.1.1.1 DNS resolver.
func (r *RadarDNSTopService) GetTopLocations(ctx context.Context, query RadarDNSTopGetTopLocationsParams, opts ...option.RequestOption) (res *RadarDNSTopGetTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSTopGetTopAsesResponse struct {
	Result  RadarDNSTopGetTopAsesResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarDNSTopGetTopAsesResponseJSON   `json:"-"`
}

// radarDNSTopGetTopAsesResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopGetTopAsesResponse]
type radarDNSTopGetTopAsesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResult struct {
	Meta RadarDNSTopGetTopAsesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarDNSTopGetTopAsesResponseResultTop0 `json:"top_0,required"`
	JSON radarDNSTopGetTopAsesResponseResultJSON   `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSTopGetTopAsesResponseResult]
type radarDNSTopGetTopAsesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResultMeta struct {
	DateRange      []RadarDNSTopGetTopAsesResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSTopGetTopAsesResponseResultMetaJSON           `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarDNSTopGetTopAsesResponseResultMeta]
type radarDNSTopGetTopAsesResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopGetTopAsesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarDNSTopGetTopAsesResponseResultMetaDateRange]
type radarDNSTopGetTopAsesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfo]
type radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotation]
type radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesResponseResultTop0 struct {
	ClientAsn    int64                                       `json:"clientASN,required"`
	ClientAsName string                                      `json:"clientASName,required"`
	Value        string                                      `json:"value,required"`
	JSON         radarDNSTopGetTopAsesResponseResultTop0JSON `json:"-"`
}

// radarDNSTopGetTopAsesResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarDNSTopGetTopAsesResponseResultTop0]
type radarDNSTopGetTopAsesResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSTopGetTopAsesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopAsesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponse struct {
	Result  RadarDNSTopGetTopLocationsResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarDNSTopGetTopLocationsResponseJSON   `json:"-"`
}

// radarDNSTopGetTopLocationsResponseJSON contains the JSON metadata for the struct
// [RadarDNSTopGetTopLocationsResponse]
type radarDNSTopGetTopLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResult struct {
	Meta RadarDNSTopGetTopLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarDNSTopGetTopLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarDNSTopGetTopLocationsResponseResultJSON   `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSTopGetTopLocationsResponseResult]
type radarDNSTopGetTopLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResultMeta struct {
	DateRange      []RadarDNSTopGetTopLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSTopGetTopLocationsResponseResultMetaJSON           `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSTopGetTopLocationsResponseResultMeta]
type radarDNSTopGetTopLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarDNSTopGetTopLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSTopGetTopLocationsResponseResultMetaDateRange]
type radarDNSTopGetTopLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfo]
type radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopLocationsResponseResultTop0 struct {
	ClientCountryAlpha2 string                                           `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                           `json:"clientCountryName,required"`
	Value               string                                           `json:"value,required"`
	JSON                radarDNSTopGetTopLocationsResponseResultTop0JSON `json:"-"`
}

// radarDNSTopGetTopLocationsResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarDNSTopGetTopLocationsResponseResultTop0]
type radarDNSTopGetTopLocationsResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarDNSTopGetTopLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTopGetTopLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTopGetTopAsesParams struct {
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
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTopGetTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopGetTopAsesParams]'s query parameters as
// `url.Values`.
func (r RadarDNSTopGetTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSTopGetTopAsesParamsFormat string

const (
	RadarDNSTopGetTopAsesParamsFormatJson RadarDNSTopGetTopAsesParamsFormat = "JSON"
	RadarDNSTopGetTopAsesParamsFormatCsv  RadarDNSTopGetTopAsesParamsFormat = "CSV"
)

func (r RadarDNSTopGetTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTopGetTopAsesParamsFormatJson, RadarDNSTopGetTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type RadarDNSTopGetTopLocationsParams struct {
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
	// Array of domain names.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTopGetTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarDNSTopGetTopLocationsParams]'s query parameters as
// `url.Values`.
func (r RadarDNSTopGetTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSTopGetTopLocationsParamsFormat string

const (
	RadarDNSTopGetTopLocationsParamsFormatJson RadarDNSTopGetTopLocationsParamsFormat = "JSON"
	RadarDNSTopGetTopLocationsParamsFormatCsv  RadarDNSTopGetTopLocationsParamsFormat = "CSV"
)

func (r RadarDNSTopGetTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTopGetTopLocationsParamsFormatJson, RadarDNSTopGetTopLocationsParamsFormatCsv:
		return true
	}
	return false
}
