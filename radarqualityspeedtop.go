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

// RadarQualitySpeedTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarQualitySpeedTopService] method instead.
type RadarQualitySpeedTopService struct {
	Options []option.RequestOption
}

// NewRadarQualitySpeedTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedTopService(opts ...option.RequestOption) (r *RadarQualitySpeedTopService) {
	r = &RadarQualitySpeedTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by bandwidth, latency, jitter, or packet
// loss, from the previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedTopService) GetTopAs(ctx context.Context, query RadarQualitySpeedTopGetTopAsParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopGetTopAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations by bandwidth, latency, jitter, or packet loss, from
// the previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedTopService) GetTopLocations(ctx context.Context, query RadarQualitySpeedTopGetTopLocationsParams, opts ...option.RequestOption) (res *RadarQualitySpeedTopGetTopLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualitySpeedTopGetTopAsResponse struct {
	Result  RadarQualitySpeedTopGetTopAsResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarQualitySpeedTopGetTopAsResponseJSON   `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedTopGetTopAsResponse]
type radarQualitySpeedTopGetTopAsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResult struct {
	Meta RadarQualitySpeedTopGetTopAsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopGetTopAsResponseResultTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopGetTopAsResponseResultJSON   `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultJSON contains the JSON metadata for
// the struct [RadarQualitySpeedTopGetTopAsResponseResult]
type radarQualitySpeedTopGetTopAsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResultMeta struct {
	DateRange      []RadarQualitySpeedTopGetTopAsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopGetTopAsResponseResultMetaJSON           `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarQualitySpeedTopGetTopAsResponseResultMeta]
type radarQualitySpeedTopGetTopAsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopGetTopAsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarQualitySpeedTopGetTopAsResponseResultMetaDateRange]
type radarQualitySpeedTopGetTopAsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfo]
type radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsResponseResultTop0 struct {
	BandwidthDownload string                                             `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                             `json:"bandwidthUpload,required"`
	ClientAsn         float64                                            `json:"clientASN,required"`
	ClientAsName      string                                             `json:"clientASName,required"`
	JitterIdle        string                                             `json:"jitterIdle,required"`
	JitterLoaded      string                                             `json:"jitterLoaded,required"`
	LatencyIdle       string                                             `json:"latencyIdle,required"`
	LatencyLoaded     string                                             `json:"latencyLoaded,required"`
	NumTests          float64                                            `json:"numTests,required"`
	RankPower         float64                                            `json:"rankPower,required"`
	JSON              radarQualitySpeedTopGetTopAsResponseResultTop0JSON `json:"-"`
}

// radarQualitySpeedTopGetTopAsResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarQualitySpeedTopGetTopAsResponseResultTop0]
type radarQualitySpeedTopGetTopAsResponseResultTop0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	ClientAsn         apijson.Field
	ClientAsName      apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	NumTests          apijson.Field
	RankPower         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopAsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopAsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponse struct {
	Result  RadarQualitySpeedTopGetTopLocationsResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarQualitySpeedTopGetTopLocationsResponseJSON   `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseJSON contains the JSON metadata for
// the struct [RadarQualitySpeedTopGetTopLocationsResponse]
type radarQualitySpeedTopGetTopLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResult struct {
	Meta RadarQualitySpeedTopGetTopLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarQualitySpeedTopGetTopLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarQualitySpeedTopGetTopLocationsResponseResultJSON   `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultJSON contains the JSON metadata
// for the struct [RadarQualitySpeedTopGetTopLocationsResponseResult]
type radarQualitySpeedTopGetTopLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResultMeta struct {
	DateRange      []RadarQualitySpeedTopGetTopLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedTopGetTopLocationsResponseResultMetaJSON           `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopGetTopLocationsResponseResultMeta]
type radarQualitySpeedTopGetTopLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedTopGetTopLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedTopGetTopLocationsResponseResultMetaDateRange]
type radarQualitySpeedTopGetTopLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfo]
type radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopLocationsResponseResultTop0 struct {
	BandwidthDownload   string                                                    `json:"bandwidthDownload,required"`
	BandwidthUpload     string                                                    `json:"bandwidthUpload,required"`
	ClientCountryAlpha2 string                                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                    `json:"clientCountryName,required"`
	JitterIdle          string                                                    `json:"jitterIdle,required"`
	JitterLoaded        string                                                    `json:"jitterLoaded,required"`
	LatencyIdle         string                                                    `json:"latencyIdle,required"`
	LatencyLoaded       string                                                    `json:"latencyLoaded,required"`
	NumTests            float64                                                   `json:"numTests,required"`
	RankPower           float64                                                   `json:"rankPower,required"`
	JSON                radarQualitySpeedTopGetTopLocationsResponseResultTop0JSON `json:"-"`
}

// radarQualitySpeedTopGetTopLocationsResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarQualitySpeedTopGetTopLocationsResponseResultTop0]
type radarQualitySpeedTopGetTopLocationsResponseResultTop0JSON struct {
	BandwidthDownload   apijson.Field
	BandwidthUpload     apijson.Field
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	JitterIdle          apijson.Field
	JitterLoaded        apijson.Field
	LatencyIdle         apijson.Field
	LatencyLoaded       apijson.Field
	NumTests            apijson.Field
	RankPower           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarQualitySpeedTopGetTopLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedTopGetTopLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedTopGetTopAsParams struct {
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
	// Format in which results will be returned.
	Format param.Field[RadarQualitySpeedTopGetTopAsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopGetTopAsParamsOrderBy] `query:"orderBy"`
	// Reverses the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopGetTopAsParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedTopGetTopAsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarQualitySpeedTopGetTopAsParamsFormat string

const (
	RadarQualitySpeedTopGetTopAsParamsFormatJson RadarQualitySpeedTopGetTopAsParamsFormat = "JSON"
	RadarQualitySpeedTopGetTopAsParamsFormatCsv  RadarQualitySpeedTopGetTopAsParamsFormat = "CSV"
)

func (r RadarQualitySpeedTopGetTopAsParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualitySpeedTopGetTopAsParamsFormatJson, RadarQualitySpeedTopGetTopAsParamsFormatCsv:
		return true
	}
	return false
}

// Metric to order the results by.
type RadarQualitySpeedTopGetTopAsParamsOrderBy string

const (
	RadarQualitySpeedTopGetTopAsParamsOrderByBandwidthDownload RadarQualitySpeedTopGetTopAsParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopGetTopAsParamsOrderByBandwidthUpload   RadarQualitySpeedTopGetTopAsParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopGetTopAsParamsOrderByLatencyIdle       RadarQualitySpeedTopGetTopAsParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopGetTopAsParamsOrderByLatencyLoaded     RadarQualitySpeedTopGetTopAsParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopGetTopAsParamsOrderByJitterIdle        RadarQualitySpeedTopGetTopAsParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopGetTopAsParamsOrderByJitterLoaded      RadarQualitySpeedTopGetTopAsParamsOrderBy = "JITTER_LOADED"
)

func (r RadarQualitySpeedTopGetTopAsParamsOrderBy) IsKnown() bool {
	switch r {
	case RadarQualitySpeedTopGetTopAsParamsOrderByBandwidthDownload, RadarQualitySpeedTopGetTopAsParamsOrderByBandwidthUpload, RadarQualitySpeedTopGetTopAsParamsOrderByLatencyIdle, RadarQualitySpeedTopGetTopAsParamsOrderByLatencyLoaded, RadarQualitySpeedTopGetTopAsParamsOrderByJitterIdle, RadarQualitySpeedTopGetTopAsParamsOrderByJitterLoaded:
		return true
	}
	return false
}

type RadarQualitySpeedTopGetTopLocationsParams struct {
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
	// Format in which results will be returned.
	Format param.Field[RadarQualitySpeedTopGetTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Metric to order the results by.
	OrderBy param.Field[RadarQualitySpeedTopGetTopLocationsParamsOrderBy] `query:"orderBy"`
	// Reverses the order of results.
	Reverse param.Field[bool] `query:"reverse"`
}

// URLQuery serializes [RadarQualitySpeedTopGetTopLocationsParams]'s query
// parameters as `url.Values`.
func (r RadarQualitySpeedTopGetTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarQualitySpeedTopGetTopLocationsParamsFormat string

const (
	RadarQualitySpeedTopGetTopLocationsParamsFormatJson RadarQualitySpeedTopGetTopLocationsParamsFormat = "JSON"
	RadarQualitySpeedTopGetTopLocationsParamsFormatCsv  RadarQualitySpeedTopGetTopLocationsParamsFormat = "CSV"
)

func (r RadarQualitySpeedTopGetTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualitySpeedTopGetTopLocationsParamsFormatJson, RadarQualitySpeedTopGetTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

// Metric to order the results by.
type RadarQualitySpeedTopGetTopLocationsParamsOrderBy string

const (
	RadarQualitySpeedTopGetTopLocationsParamsOrderByBandwidthDownload RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "BANDWIDTH_DOWNLOAD"
	RadarQualitySpeedTopGetTopLocationsParamsOrderByBandwidthUpload   RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "BANDWIDTH_UPLOAD"
	RadarQualitySpeedTopGetTopLocationsParamsOrderByLatencyIdle       RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "LATENCY_IDLE"
	RadarQualitySpeedTopGetTopLocationsParamsOrderByLatencyLoaded     RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "LATENCY_LOADED"
	RadarQualitySpeedTopGetTopLocationsParamsOrderByJitterIdle        RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "JITTER_IDLE"
	RadarQualitySpeedTopGetTopLocationsParamsOrderByJitterLoaded      RadarQualitySpeedTopGetTopLocationsParamsOrderBy = "JITTER_LOADED"
)

func (r RadarQualitySpeedTopGetTopLocationsParamsOrderBy) IsKnown() bool {
	switch r {
	case RadarQualitySpeedTopGetTopLocationsParamsOrderByBandwidthDownload, RadarQualitySpeedTopGetTopLocationsParamsOrderByBandwidthUpload, RadarQualitySpeedTopGetTopLocationsParamsOrderByLatencyIdle, RadarQualitySpeedTopGetTopLocationsParamsOrderByLatencyLoaded, RadarQualitySpeedTopGetTopLocationsParamsOrderByJitterIdle, RadarQualitySpeedTopGetTopLocationsParamsOrderByJitterLoaded:
		return true
	}
	return false
}
