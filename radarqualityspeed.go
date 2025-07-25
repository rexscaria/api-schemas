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
	// Metadata for the results.
	Meta RadarQualitySpeedGetHistogramResponseResultMeta `json:"meta,required"`
	JSON radarQualitySpeedGetHistogramResponseResultJSON `json:"-"`
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

// Metadata for the results.
type RadarQualitySpeedGetHistogramResponseResultMeta struct {
	// The width for every bucket in the histogram.
	BucketSize     int64                                                         `json:"bucketSize,required"`
	ConfidenceInfo RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarQualitySpeedGetHistogramResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarQualitySpeedGetHistogramResponseResultMetaNormalization `json:"normalization,required"`
	TotalTests    []int64                                                      `json:"totalTests,required"`
	// Measurement units for the results.
	Units []RadarQualitySpeedGetHistogramResponseResultMetaUnit `json:"units,required"`
	JSON  radarQualitySpeedGetHistogramResponseResultMetaJSON   `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarQualitySpeedGetHistogramResponseResultMeta]
type radarQualitySpeedGetHistogramResponseResultMetaJSON struct {
	BucketSize     apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	TotalTests     apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfo]
type radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetHistogramResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarQualitySpeedGetHistogramResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarQualitySpeedGetHistogramResponseResultMetaDateRange]
type radarQualitySpeedGetHistogramResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarQualitySpeedGetHistogramResponseResultMetaNormalization string

const (
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationPercentage           RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "PERCENTAGE"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationMin0Max              RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "MIN0_MAX"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationMinMax               RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "MIN_MAX"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRawValues            RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "RAW_VALUES"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationPercentageChange     RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRollingAverage       RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationOverlappedPercentage RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRatio                RadarQualitySpeedGetHistogramResponseResultMetaNormalization = "RATIO"
)

func (r RadarQualitySpeedGetHistogramResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarQualitySpeedGetHistogramResponseResultMetaNormalizationPercentage, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationMin0Max, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationMinMax, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRawValues, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationPercentageChange, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRollingAverage, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationOverlappedPercentage, RadarQualitySpeedGetHistogramResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarQualitySpeedGetHistogramResponseResultMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarQualitySpeedGetHistogramResponseResultMetaUnitJSON `json:"-"`
}

// radarQualitySpeedGetHistogramResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarQualitySpeedGetHistogramResponseResultMetaUnit]
type radarQualitySpeedGetHistogramResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetHistogramResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetHistogramResponseResultMetaUnitJSON) RawJSON() string {
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
	// Metadata for the results.
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

// Metadata for the results.
type RadarQualitySpeedGetSummaryResponseResultMeta struct {
	ConfidenceInfo RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarQualitySpeedGetSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarQualitySpeedGetSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarQualitySpeedGetSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarQualitySpeedGetSummaryResponseResultMetaJSON   `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarQualitySpeedGetSummaryResponseResultMeta]
type radarQualitySpeedGetSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
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

// Annotation associated with the result (e.g. outage or other type of event).
type RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarQualitySpeedGetSummaryResponseResultMetaNormalization string

const (
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationPercentage           RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationMin0Max              RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationMinMax               RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRawValues            RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationPercentageChange     RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRollingAverage       RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationOverlappedPercentage RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRatio                RadarQualitySpeedGetSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarQualitySpeedGetSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarQualitySpeedGetSummaryResponseResultMetaNormalizationPercentage, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationMin0Max, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationMinMax, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRawValues, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationPercentageChange, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRollingAverage, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarQualitySpeedGetSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarQualitySpeedGetSummaryResponseResultMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  radarQualitySpeedGetSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarQualitySpeedGetSummaryResponseResultMetaUnitJSON contains the JSON metadata
// for the struct [RadarQualitySpeedGetSummaryResponseResultMetaUnit]
type radarQualitySpeedGetSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualitySpeedGetSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarQualitySpeedGetSummaryResponseResultMetaUnitJSON) RawJSON() string {
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
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Specifies the width for every bucket in the histogram.
	BucketSize param.Field[int64] `query:"bucketSize"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarQualitySpeedGetHistogramParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
	// Format in which results will be returned.
	Format param.Field[RadarQualitySpeedGetSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
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
