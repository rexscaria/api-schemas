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

// RadarAttackLayer3TimeseriesGroupService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupService] method instead.
type RadarAttackLayer3TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupService) {
	r = &RadarAttackLayer3TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 3 attacks by bitrate over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetBitrateTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by duration over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetDurationTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by targeted industry over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetIndustryTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by IP version over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetIPVersionTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by protocol over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetProtocolTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by vector over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetVectorTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by targeted vertical over time.
func (r *RadarAttackLayer3TimeseriesGroupService) GetVerticalTimeseries(ctx context.Context, query RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResult `json:"result,required"`
	Success bool                                                               `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                    `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                          `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                         `json:"name,required"`
	Value string                                                                         `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0 struct {
	OneGBPSToTenGBPS         []string                                                                     `json:"_1_GBPS_TO_10_GBPS,required"`
	TenGBPSToOneHundredGBPS  []string                                                                     `json:"_10_GBPS_TO_100_GBPS,required"`
	FiveHundredMBPSToOneGBPS []string                                                                     `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps              []string                                                                     `json:"OVER_100_GBPS,required"`
	Timestamps               []time.Time                                                                  `json:"timestamps,required" format:"date-time"`
	Under500Mbps             []string                                                                     `json:"UNDER_500_MBPS,required"`
	JSON                     radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0JSON struct {
	OneGBPSToTenGBPS         apijson.Field
	TenGBPSToOneHundredGBPS  apijson.Field
	FiveHundredMBPSToOneGBPS apijson.Field
	Over100Gbps              apijson.Field
	Timestamps               apijson.Field
	Under500Mbps             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetBitrateTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0 struct {
	OneHourToThreeHours   []string                                                                      `json:"_1_HOUR_TO_3_HOURS,required"`
	TenMinsToTwentyMins   []string                                                                      `json:"_10_MINS_TO_20_MINS,required"`
	TwentyMinsToFortyMins []string                                                                      `json:"_20_MINS_TO_40_MINS,required"`
	FortyMinsToOneHour    []string                                                                      `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours            []string                                                                      `json:"OVER_3_HOURS,required"`
	Timestamps            []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	Under10Mins           []string                                                                      `json:"UNDER_10_MINS,required"`
	JSON                  radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0JSON struct {
	OneHourToThreeHours   apijson.Field
	TenMinsToTwentyMins   apijson.Field
	TwentyMinsToFortyMins apijson.Field
	FortyMinsToOneHour    apijson.Field
	Over3Hours            apijson.Field
	Timestamps            apijson.Field
	Under10Mins           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetDurationTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                           `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                 `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                      `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                           `json:"name,required"`
	Value string                                                                           `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0 struct {
	IPv4       []string                                                                       `json:"IPv4,required"`
	IPv6       []string                                                                       `json:"IPv6,required"`
	Timestamps []time.Time                                                                    `json:"timestamps,required" format:"date-time"`
	JSON       radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0 struct {
	Gre        []string                                                                      `json:"GRE,required"`
	Icmp       []string                                                                      `json:"ICMP,required"`
	Tcp        []string                                                                      `json:"TCP,required"`
	Timestamps []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	Udp        []string                                                                      `json:"UDP,required"`
	JSON       radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetProtocolTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResult `json:"result,required"`
	Success bool                                                              `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                   `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                         `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                        `json:"name,required"`
	Value string                                                                        `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                 `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                         `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVectorTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponse]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResult]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMeta]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                           `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetBitrateTimeseriesParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetDurationTimeseriesParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetIndustryTimeseriesParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetIPVersionTimeseriesParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetProtocolTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetVectorTimeseriesParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes
// [RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval15m, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1h, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1d, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirectionOrigin, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormatJson RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormatJson, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolGre  RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolUdp, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolTcp, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolIcmp, RadarAttackLayer3TimeseriesGroupGetVerticalTimeseriesParamsProtocolGre:
		return true
	}
	return false
}
