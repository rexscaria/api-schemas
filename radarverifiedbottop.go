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

// RadarVerifiedBotTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarVerifiedBotTopService] method instead.
type RadarVerifiedBotTopService struct {
	Options []option.RequestOption
}

// NewRadarVerifiedBotTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarVerifiedBotTopService(opts ...option.RequestOption) (r *RadarVerifiedBotTopService) {
	r = &RadarVerifiedBotTopService{}
	r.Options = opts
	return
}

// Retrieves the top verified bots by HTTP requests, with owner and category.
func (r *RadarVerifiedBotTopService) Bots(ctx context.Context, query RadarVerifiedBotTopBotsParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopBotsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top verified bot categories by HTTP requests, along with their
// corresponding percentage, over the total verified bot HTTP requests.
func (r *RadarVerifiedBotTopService) Categories(ctx context.Context, query RadarVerifiedBotTopCategoriesParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarVerifiedBotTopBotsResponse struct {
	Result  RadarVerifiedBotTopBotsResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarVerifiedBotTopBotsResponseJSON   `json:"-"`
}

// radarVerifiedBotTopBotsResponseJSON contains the JSON metadata for the struct
// [RadarVerifiedBotTopBotsResponse]
type radarVerifiedBotTopBotsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResult struct {
	Meta RadarVerifiedBotTopBotsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopBotsResponseResultTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopBotsResponseResultJSON   `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseResult]
type radarVerifiedBotTopBotsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultMeta struct {
	DateRange      []RadarVerifiedBotTopBotsResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopBotsResponseResultMetaJSON           `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseResultMeta]
type radarVerifiedBotTopBotsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopBotsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarVerifiedBotTopBotsResponseResultMetaDateRange]
type radarVerifiedBotTopBotsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo]
type radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultTop0 struct {
	BotCategory string                                        `json:"botCategory,required"`
	BotName     string                                        `json:"botName,required"`
	BotOwner    string                                        `json:"botOwner,required"`
	Value       string                                        `json:"value,required"`
	JSON        radarVerifiedBotTopBotsResponseResultTop0JSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseResultTop0]
type radarVerifiedBotTopBotsResponseResultTop0JSON struct {
	BotCategory apijson.Field
	BotName     apijson.Field
	BotOwner    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponse struct {
	Result  RadarVerifiedBotTopCategoriesResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarVerifiedBotTopCategoriesResponseJSON   `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopCategoriesResponse]
type radarVerifiedBotTopCategoriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResult struct {
	Meta RadarVerifiedBotTopCategoriesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarVerifiedBotTopCategoriesResponseResultTop0 `json:"top_0,required"`
	JSON radarVerifiedBotTopCategoriesResponseResultJSON   `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopCategoriesResponseResult]
type radarVerifiedBotTopCategoriesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultMeta struct {
	DateRange      []RadarVerifiedBotTopCategoriesResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarVerifiedBotTopCategoriesResponseResultMetaJSON           `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoriesResponseResultMeta]
type radarVerifiedBotTopCategoriesResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarVerifiedBotTopCategoriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseResultMetaDateRange]
type radarVerifiedBotTopCategoriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo]
type radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultTop0 struct {
	BotCategory string                                              `json:"botCategory,required"`
	Value       string                                              `json:"value,required"`
	JSON        radarVerifiedBotTopCategoriesResponseResultTop0JSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoriesResponseResultTop0]
type radarVerifiedBotTopCategoriesResponseResultTop0JSON struct {
	BotCategory apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsParams struct {
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
	Format param.Field[RadarVerifiedBotTopBotsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopBotsParams]'s query parameters as
// `url.Values`.
func (r RadarVerifiedBotTopBotsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarVerifiedBotTopBotsParamsFormat string

const (
	RadarVerifiedBotTopBotsParamsFormatJson RadarVerifiedBotTopBotsParamsFormat = "JSON"
	RadarVerifiedBotTopBotsParamsFormatCsv  RadarVerifiedBotTopBotsParamsFormat = "CSV"
)

func (r RadarVerifiedBotTopBotsParamsFormat) IsKnown() bool {
	switch r {
	case RadarVerifiedBotTopBotsParamsFormatJson, RadarVerifiedBotTopBotsParamsFormatCsv:
		return true
	}
	return false
}

type RadarVerifiedBotTopCategoriesParams struct {
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
	Format param.Field[RadarVerifiedBotTopCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarVerifiedBotTopCategoriesParams]'s query parameters as
// `url.Values`.
func (r RadarVerifiedBotTopCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarVerifiedBotTopCategoriesParamsFormat string

const (
	RadarVerifiedBotTopCategoriesParamsFormatJson RadarVerifiedBotTopCategoriesParamsFormat = "JSON"
	RadarVerifiedBotTopCategoriesParamsFormatCsv  RadarVerifiedBotTopCategoriesParamsFormat = "CSV"
)

func (r RadarVerifiedBotTopCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarVerifiedBotTopCategoriesParamsFormatJson, RadarVerifiedBotTopCategoriesParamsFormatCsv:
		return true
	}
	return false
}
