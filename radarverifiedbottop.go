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
//
// Deprecated: Use
// [Radar Bots API](https://developers.cloudflare.com/api/resources/radar/subresources/bots/)
// instead.
func (r *RadarVerifiedBotTopService) Bots(ctx context.Context, query RadarVerifiedBotTopBotsParams, opts ...option.RequestOption) (res *RadarVerifiedBotTopBotsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/verified_bots/top/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top verified bot categories by HTTP requests, along with their
// corresponding percentage, over the total verified bot HTTP requests.
//
// Deprecated: Use
// [Radar Bots API](https://developers.cloudflare.com/api/resources/radar/subresources/bots/)
// instead.
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
	// Metadata for the results.
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

// Metadata for the results.
type RadarVerifiedBotTopBotsResponseResultMeta struct {
	ConfidenceInfo RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarVerifiedBotTopBotsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarVerifiedBotTopBotsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarVerifiedBotTopBotsResponseResultMetaUnit `json:"units,required"`
	JSON  radarVerifiedBotTopBotsResponseResultMetaJSON   `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarVerifiedBotTopBotsResponseResultMeta]
type radarVerifiedBotTopBotsResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarVerifiedBotTopBotsResponseResultMetaNormalization string

const (
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationPercentage           RadarVerifiedBotTopBotsResponseResultMetaNormalization = "PERCENTAGE"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationMin0Max              RadarVerifiedBotTopBotsResponseResultMetaNormalization = "MIN0_MAX"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationMinMax               RadarVerifiedBotTopBotsResponseResultMetaNormalization = "MIN_MAX"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationRawValues            RadarVerifiedBotTopBotsResponseResultMetaNormalization = "RAW_VALUES"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationPercentageChange     RadarVerifiedBotTopBotsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationRollingAverage       RadarVerifiedBotTopBotsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationOverlappedPercentage RadarVerifiedBotTopBotsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarVerifiedBotTopBotsResponseResultMetaNormalizationRatio                RadarVerifiedBotTopBotsResponseResultMetaNormalization = "RATIO"
)

func (r RadarVerifiedBotTopBotsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarVerifiedBotTopBotsResponseResultMetaNormalizationPercentage, RadarVerifiedBotTopBotsResponseResultMetaNormalizationMin0Max, RadarVerifiedBotTopBotsResponseResultMetaNormalizationMinMax, RadarVerifiedBotTopBotsResponseResultMetaNormalizationRawValues, RadarVerifiedBotTopBotsResponseResultMetaNormalizationPercentageChange, RadarVerifiedBotTopBotsResponseResultMetaNormalizationRollingAverage, RadarVerifiedBotTopBotsResponseResultMetaNormalizationOverlappedPercentage, RadarVerifiedBotTopBotsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarVerifiedBotTopBotsResponseResultMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  radarVerifiedBotTopBotsResponseResultMetaUnitJSON `json:"-"`
}

// radarVerifiedBotTopBotsResponseResultMetaUnitJSON contains the JSON metadata for
// the struct [RadarVerifiedBotTopBotsResponseResultMetaUnit]
type radarVerifiedBotTopBotsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopBotsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopBotsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopBotsResponseResultTop0 struct {
	BotCategory string `json:"botCategory,required"`
	BotName     string `json:"botName,required"`
	BotOwner    string `json:"botOwner,required"`
	// A numeric string.
	Value string                                        `json:"value,required"`
	JSON  radarVerifiedBotTopBotsResponseResultTop0JSON `json:"-"`
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
	// Metadata for the results.
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

// Metadata for the results.
type RadarVerifiedBotTopCategoriesResponseResultMeta struct {
	ConfidenceInfo RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarVerifiedBotTopCategoriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarVerifiedBotTopCategoriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarVerifiedBotTopCategoriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarVerifiedBotTopCategoriesResponseResultMetaJSON   `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarVerifiedBotTopCategoriesResponseResultMeta]
type radarVerifiedBotTopCategoriesResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation]
type radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarVerifiedBotTopCategoriesResponseResultMetaNormalization string

const (
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationPercentage           RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationMin0Max              RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationMinMax               RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "MIN_MAX"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRawValues            RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationPercentageChange     RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRollingAverage       RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationOverlappedPercentage RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRatio                RadarVerifiedBotTopCategoriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarVerifiedBotTopCategoriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationPercentage, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationMin0Max, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationMinMax, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRawValues, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationPercentageChange, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRollingAverage, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationOverlappedPercentage, RadarVerifiedBotTopCategoriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarVerifiedBotTopCategoriesResponseResultMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarVerifiedBotTopCategoriesResponseResultMetaUnitJSON `json:"-"`
}

// radarVerifiedBotTopCategoriesResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarVerifiedBotTopCategoriesResponseResultMetaUnit]
type radarVerifiedBotTopCategoriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarVerifiedBotTopCategoriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarVerifiedBotTopCategoriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarVerifiedBotTopCategoriesResponseResultTop0 struct {
	BotCategory string `json:"botCategory,required"`
	// A numeric string.
	Value string                                              `json:"value,required"`
	JSON  radarVerifiedBotTopCategoriesResponseResultTop0JSON `json:"-"`
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
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarVerifiedBotTopBotsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarVerifiedBotTopCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
