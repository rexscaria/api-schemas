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
	// Metadata for the results.
	Meta RadarNetflowTopGetTopAsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarNetflowTopGetTopAsResponseResultTop0 `json:"top_0,required"`
	JSON radarNetflowTopGetTopAsResponseResultJSON   `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowTopGetTopAsResponseResult]
type radarNetflowTopGetTopAsResponseResultJSON struct {
	Meta        apijson.Field
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

// Metadata for the results.
type RadarNetflowTopGetTopAsResponseResultMeta struct {
	ConfidenceInfo RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarNetflowTopGetTopAsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarNetflowTopGetTopAsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarNetflowTopGetTopAsResponseResultMetaUnit `json:"units,required"`
	JSON  radarNetflowTopGetTopAsResponseResultMetaJSON   `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarNetflowTopGetTopAsResponseResultMeta]
type radarNetflowTopGetTopAsResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfo]
type radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowTopGetTopAsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarNetflowTopGetTopAsResponseResultMetaDateRange]
type radarNetflowTopGetTopAsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowTopGetTopAsResponseResultMetaNormalization string

const (
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationPercentage           RadarNetflowTopGetTopAsResponseResultMetaNormalization = "PERCENTAGE"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationMin0Max              RadarNetflowTopGetTopAsResponseResultMetaNormalization = "MIN0_MAX"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationMinMax               RadarNetflowTopGetTopAsResponseResultMetaNormalization = "MIN_MAX"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationRawValues            RadarNetflowTopGetTopAsResponseResultMetaNormalization = "RAW_VALUES"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationPercentageChange     RadarNetflowTopGetTopAsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationRollingAverage       RadarNetflowTopGetTopAsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationOverlappedPercentage RadarNetflowTopGetTopAsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarNetflowTopGetTopAsResponseResultMetaNormalizationRatio                RadarNetflowTopGetTopAsResponseResultMetaNormalization = "RATIO"
)

func (r RadarNetflowTopGetTopAsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarNetflowTopGetTopAsResponseResultMetaNormalizationPercentage, RadarNetflowTopGetTopAsResponseResultMetaNormalizationMin0Max, RadarNetflowTopGetTopAsResponseResultMetaNormalizationMinMax, RadarNetflowTopGetTopAsResponseResultMetaNormalizationRawValues, RadarNetflowTopGetTopAsResponseResultMetaNormalizationPercentageChange, RadarNetflowTopGetTopAsResponseResultMetaNormalizationRollingAverage, RadarNetflowTopGetTopAsResponseResultMetaNormalizationOverlappedPercentage, RadarNetflowTopGetTopAsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarNetflowTopGetTopAsResponseResultMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  radarNetflowTopGetTopAsResponseResultMetaUnitJSON `json:"-"`
}

// radarNetflowTopGetTopAsResponseResultMetaUnitJSON contains the JSON metadata for
// the struct [RadarNetflowTopGetTopAsResponseResultMetaUnit]
type radarNetflowTopGetTopAsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopAsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopAsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopAsResponseResultTop0 struct {
	ClientAsn    float64 `json:"clientASN,required"`
	ClientAsName string  `json:"clientASName,required"`
	// A numeric string.
	Value string                                        `json:"value,required"`
	JSON  radarNetflowTopGetTopAsResponseResultTop0JSON `json:"-"`
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
	// Metadata for the results.
	Meta RadarNetflowTopGetTopLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarNetflowTopGetTopLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarNetflowTopGetTopLocationsResponseResultJSON   `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultJSON contains the JSON metadata for
// the struct [RadarNetflowTopGetTopLocationsResponseResult]
type radarNetflowTopGetTopLocationsResponseResultJSON struct {
	Meta        apijson.Field
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

// Metadata for the results.
type RadarNetflowTopGetTopLocationsResponseResultMeta struct {
	ConfidenceInfo RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarNetflowTopGetTopLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarNetflowTopGetTopLocationsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarNetflowTopGetTopLocationsResponseResultMetaUnit `json:"units,required"`
	JSON  radarNetflowTopGetTopLocationsResponseResultMetaJSON   `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarNetflowTopGetTopLocationsResponseResultMeta]
type radarNetflowTopGetTopLocationsResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfo]
type radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowTopGetTopLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarNetflowTopGetTopLocationsResponseResultMetaDateRange]
type radarNetflowTopGetTopLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowTopGetTopLocationsResponseResultMetaNormalization string

const (
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationPercentage           RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "PERCENTAGE"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationMin0Max              RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "MIN0_MAX"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationMinMax               RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "MIN_MAX"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRawValues            RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "RAW_VALUES"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationPercentageChange     RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRollingAverage       RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationOverlappedPercentage RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRatio                RadarNetflowTopGetTopLocationsResponseResultMetaNormalization = "RATIO"
)

func (r RadarNetflowTopGetTopLocationsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationPercentage, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationMin0Max, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationMinMax, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRawValues, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationPercentageChange, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRollingAverage, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationOverlappedPercentage, RadarNetflowTopGetTopLocationsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarNetflowTopGetTopLocationsResponseResultMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarNetflowTopGetTopLocationsResponseResultMetaUnitJSON `json:"-"`
}

// radarNetflowTopGetTopLocationsResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarNetflowTopGetTopLocationsResponseResultMetaUnit]
type radarNetflowTopGetTopLocationsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTopGetTopLocationsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTopGetTopLocationsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTopGetTopLocationsResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                               `json:"value,required"`
	JSON  radarNetflowTopGetTopLocationsResponseResultTop0JSON `json:"-"`
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
	Format param.Field[RadarNetflowTopGetTopAsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
	Format param.Field[RadarNetflowTopGetTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
