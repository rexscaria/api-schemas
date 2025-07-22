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

// RadarAttackLayer3TopLocationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TopLocationService] method instead.
type RadarAttackLayer3TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer3TopLocationService) {
	r = &RadarAttackLayer3TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the origin locations of layer 3 attacks.
func (r *RadarAttackLayer3TopLocationService) GetTopOriginLocations(ctx context.Context, query RadarAttackLayer3TopLocationGetTopOriginLocationsParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationGetTopOriginLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the target locations of layer 3 attacks.
func (r *RadarAttackLayer3TopLocationService) GetTopTargetLocations(ctx context.Context, query RadarAttackLayer3TopLocationGetTopTargetLocationsParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopLocationGetTopTargetLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponse struct {
	Result  RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResult `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarAttackLayer3TopLocationGetTopOriginLocationsResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponse]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResult struct {
	Meta RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResult]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMeta]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRange]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                         `json:"level"`
	JSON        radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                          `json:"dataSource,required"`
	Description     string                                                                                          `json:"description,required"`
	EventType       string                                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                                  `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                                  `json:"originCountryName,required"`
	Rank                float64                                                                 `json:"rank,required"`
	Value               string                                                                  `json:"value,required"`
	JSON                radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0]
type radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopOriginLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponse struct {
	Result  RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResult `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarAttackLayer3TopLocationGetTopTargetLocationsResponseJSON   `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponse]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResult struct {
	Meta RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResult]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMeta]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRange]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                         `json:"level"`
	JSON        radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                          `json:"dataSource,required"`
	Description     string                                                                                          `json:"description,required"`
	EventType       string                                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0 struct {
	Rank                float64                                                                 `json:"rank,required"`
	TargetCountryAlpha2 string                                                                  `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                                  `json:"targetCountryName,required"`
	Value               string                                                                  `json:"value,required"`
	JSON                radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0]
type radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopLocationGetTopTargetLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsParams struct {
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
	Format param.Field[RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationGetTopOriginLocationsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationGetTopOriginLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormat string

const (
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormatJson RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormat = "JSON"
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormatCsv  RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormat = "CSV"
)

func (r RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormatJson, RadarAttackLayer3TopLocationGetTopOriginLocationsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion string

const (
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersionIPv4 RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersionIPv6 RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersionIPv4, RadarAttackLayer3TopLocationGetTopOriginLocationsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol string

const (
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolUdp  RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolTcp  RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolIcmp RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolGre  RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolUdp, RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolTcp, RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolIcmp, RadarAttackLayer3TopLocationGetTopOriginLocationsParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsParams struct {
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
	Format param.Field[RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopLocationGetTopTargetLocationsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer3TopLocationGetTopTargetLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormat string

const (
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormatJson RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormat = "JSON"
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormatCsv  RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormat = "CSV"
)

func (r RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormatJson, RadarAttackLayer3TopLocationGetTopTargetLocationsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion string

const (
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersionIPv4 RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion = "IPv4"
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersionIPv6 RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersionIPv4, RadarAttackLayer3TopLocationGetTopTargetLocationsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol string

const (
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolUdp  RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol = "UDP"
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolTcp  RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol = "TCP"
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolIcmp RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol = "ICMP"
	RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolGre  RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolUdp, RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolTcp, RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolIcmp, RadarAttackLayer3TopLocationGetTopTargetLocationsParamsProtocolGre:
		return true
	}
	return false
}
