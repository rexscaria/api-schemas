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

// RadarQualitySpeedService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarQualitySpeedService] method instead.
type RadarQualitySpeedService struct {
	Options []option.RequestOption
	Top     *RadarQualitySpeedTopService
}

// NewRadarQualitySpeedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualitySpeedService(opts ...option.RequestOption) (r *RadarQualitySpeedService) {
	r = &RadarQualitySpeedService{}
	r.Options = opts
	r.Top = NewRadarQualitySpeedTopService(opts...)
	return
}

// Retrieves a histogram from the previous 90 days of Cloudflare Speed Test data,
// split into fixed bandwidth (Mbps), latency (ms), or jitter (ms) buckets.
func (r *RadarQualitySpeedService) GetHistogram(ctx context.Context, query RadarQualitySpeedGetHistogramParams, opts ...option.RequestOption) (res *RadarQualitySpeedGetHistogramResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/histogram"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a summary of bandwidth, latency, jitter, and packet loss, from the
// previous 90 days of Cloudflare Speed Test data.
func (r *RadarQualitySpeedService) GetSummary(ctx context.Context, query RadarQualitySpeedGetSummaryParams, opts ...option.RequestOption) (res *RadarQualitySpeedGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/speed/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualitySpeedGetHistogramResponse struct {
	Result  RadarQualitySpeedGetHistogramResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarQualitySpeedGetHistogramResponseJSON   `json:"-"`
}

// radarQualitySpeedGetHistogramResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedGetHistogramResponse]
type radarQualitySpeedGetHistogramResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetHistogramResponseResult struct {
	Histogram0 RadarQualitySpeedGetHistogramResponseResultHistogram0 `json:"histogram_0,required"`
	Meta       interface{}                                           `json:"meta,required"`
	JSON       radarQualitySpeedGetHistogramResponseResultJSON       `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultJSON contains the JSON metadata for
// the struct [RadarQualitySpeedGetHistogramResponseResult]
type radarQualitySpeedGetHistogramResponseResultJSON struct {
	Histogram0  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetHistogramResponseResultHistogram0 struct {
	BandwidthDownload []string                                                  `json:"bandwidthDownload,required"`
	BandwidthUpload   []string                                                  `json:"bandwidthUpload,required"`
	BucketMin         []string                                                  `json:"bucketMin,required"`
	JSON              radarQualitySpeedGetHistogramResponseResultHistogram0JSON `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultHistogram0JSON contains the JSON
// metadata for the struct [RadarQualitySpeedGetHistogramResponseResultHistogram0]
type radarQualitySpeedGetHistogramResponseResultHistogram0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	BucketMin         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResultHistogram0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultHistogram0JSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponse struct {
	Result  RadarQualitySpeedGetSummaryResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarQualitySpeedGetSummaryResponseJSON   `json:"-"`
}

// radarQualitySpeedGetSummaryResponseJSON contains the JSON metadata for the
// struct [RadarQualitySpeedGetSummaryResponse]
type radarQualitySpeedGetSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResult struct {
	Meta     RadarQualitySpeedGetSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarQualitySpeedGetSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarQualitySpeedGetSummaryResponseResultJSON     `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultJSON contains the JSON metadata for the
// struct [RadarQualitySpeedGetSummaryResponseResult]
type radarQualitySpeedGetSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultMeta struct {
	DateRange      []RadarQualitySpeedGetSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualitySpeedGetSummaryResponseResultMetaJSON           `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarQualitySpeedGetSummaryResponseResultMeta]
type radarQualitySpeedGetSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedGetSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarQualitySpeedGetSummaryResponseResultMetaDateRange]
type radarQualitySpeedGetSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo]
type radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultSummary0 struct {
	BandwidthDownload string                                                `json:"bandwidthDownload,required"`
	BandwidthUpload   string                                                `json:"bandwidthUpload,required"`
	JitterIdle        string                                                `json:"jitterIdle,required"`
	JitterLoaded      string                                                `json:"jitterLoaded,required"`
	LatencyIdle       string                                                `json:"latencyIdle,required"`
	LatencyLoaded     string                                                `json:"latencyLoaded,required"`
	PacketLoss        string                                                `json:"packetLoss,required"`
	JSON              radarQualitySpeedGetSummaryResponseResultSummary0JSON `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarQualitySpeedGetSummaryResponseResultSummary0]
type radarQualitySpeedGetSummaryResponseResultSummary0JSON struct {
	BandwidthDownload apijson.Field
	BandwidthUpload   apijson.Field
	JitterIdle        apijson.Field
	JitterLoaded      apijson.Field
	LatencyIdle       apijson.Field
	LatencyLoaded     apijson.Field
	PacketLoss        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetHistogramParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarQualitySpeedGetHistogramParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Metrics to be returned.
	MetricGroup param.Field[RadarQualitySpeedGetHistogramParamsMetricGroup] `query:"metricGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedGetHistogramParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedGetHistogramParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarQualitySpeedGetHistogramParamsFormat string

const (
	RadarQualitySpeedGetHistogramParamsFormatJson RadarQualitySpeedGetHistogramParamsFormat = "JSON"
	RadarQualitySpeedGetHistogramParamsFormatCsv  RadarQualitySpeedGetHistogramParamsFormat = "CSV"
)

func (r RadarQualitySpeedGetHistogramParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualitySpeedGetHistogramParamsFormatJson, RadarQualitySpeedGetHistogramParamsFormatCsv:
		return true
	}
	return false
}

// Metrics to be returned.
type RadarQualitySpeedGetHistogramParamsMetricGroup string

const (
	RadarQualitySpeedGetHistogramParamsMetricGroupBandwidth RadarQualitySpeedGetHistogramParamsMetricGroup = "BANDWIDTH"
	RadarQualitySpeedGetHistogramParamsMetricGroupLatency   RadarQualitySpeedGetHistogramParamsMetricGroup = "LATENCY"
	RadarQualitySpeedGetHistogramParamsMetricGroupJitter    RadarQualitySpeedGetHistogramParamsMetricGroup = "JITTER"
)

func (r RadarQualitySpeedGetHistogramParamsMetricGroup) IsKnown() bool {
	switch r {
	case RadarQualitySpeedGetHistogramParamsMetricGroupBandwidth, RadarQualitySpeedGetHistogramParamsMetricGroupLatency, RadarQualitySpeedGetHistogramParamsMetricGroupJitter:
		return true
	}
	return false
}

type RadarQualitySpeedGetSummaryParams struct {
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
	Format param.Field[RadarQualitySpeedGetSummaryParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualitySpeedGetSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarQualitySpeedGetSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarQualitySpeedGetSummaryParamsFormat string

const (
	RadarQualitySpeedGetSummaryParamsFormatJson RadarQualitySpeedGetSummaryParamsFormat = "JSON"
	RadarQualitySpeedGetSummaryParamsFormatCsv  RadarQualitySpeedGetSummaryParamsFormat = "CSV"
)

func (r RadarQualitySpeedGetSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarQualitySpeedGetSummaryParamsFormatJson, RadarQualitySpeedGetSummaryParamsFormatCsv:
		return true
	}
	return false
}
