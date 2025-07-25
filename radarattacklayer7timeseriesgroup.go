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

// RadarAttackLayer7TimeseriesGroupService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7TimeseriesGroupService] method instead.
type RadarAttackLayer7TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer7TimeseriesGroupService) {
	r = &RadarAttackLayer7TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP method over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetHTTPMethodTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP version over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetHTTPVersionTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by targeted industry over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetIndustryTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by IP version used over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetIPVersionTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by managed rules over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetManagedRulesTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by mitigation product over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetMitigationProductTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by targeted vertical over time.
func (r *RadarAttackLayer7TimeseriesGroupService) GetVerticalTimeseries(ctx context.Context, query RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                  `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                       `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                            `json:"name,required"`
	Value string                                                                            `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                     `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                             `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                   `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                        `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                             `json:"name,required"`
	Value string                                                                             `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0 struct {
	HTTP1X     []string                                                                         `json:"HTTP/1.x,required"`
	HTTP2      []string                                                                         `json:"HTTP/2,required"`
	HTTP3      []string                                                                         `json:"HTTP/3,required"`
	Timestamps []time.Time                                                                      `json:"timestamps,required" format:"date-time"`
	JSON       radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                           `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIndustryTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                 `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                      `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                           `json:"name,required"`
	Value string                                                                           `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0 struct {
	IPv4       []string                                                                       `json:"IPv4,required"`
	IPv6       []string                                                                       `json:"IPv6,required"`
	Timestamps []time.Time                                                                    `json:"timestamps,required" format:"date-time"`
	JSON       radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                    `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                         `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                                `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                              `json:"name,required"`
	Value string                                                                              `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                       `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                               `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                         `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                              `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                                   `json:"name,required"`
	Value string                                                                                   `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                            `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                                    `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponse struct {
	Result  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResult `json:"result,required"`
	Success bool                                                                `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponse]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResult]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMeta]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                     `json:"level,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit struct {
	Name  string                                                                          `json:"name,required"`
	Value string                                                                          `json:"value,required"`
	JSON  radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                   `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                           `json:"-,extras"`
	JSON        radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0]
type radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TimeseriesGroupGetVerticalTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPMethodTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetHTTPVersionTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetIndustryTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetIPVersionTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetManagedRulesTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetMitigationProductTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval15m, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1h, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1d, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormatJson RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormatJson, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodGet             RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPost            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPut             RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodHead            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodACL             RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodLock            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMove            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodReport          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodJson            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCook            RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodGet, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPost, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPut, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodHead, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodACL, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodLock, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodMove, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodReport, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodJson, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodCook, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv4, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductDdos               RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductWaf                RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductDdos, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductWaf, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalizationMin0Max    RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalizationPercentage, RadarAttackLayer7TimeseriesGroupGetVerticalTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}
