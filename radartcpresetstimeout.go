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

// RadarTcpResetsTimeoutService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarTcpResetsTimeoutService] method instead.
type RadarTcpResetsTimeoutService struct {
	Options []option.RequestOption
}

// NewRadarTcpResetsTimeoutService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarTcpResetsTimeoutService(opts ...option.RequestOption) (r *RadarTcpResetsTimeoutService) {
	r = &RadarTcpResetsTimeoutService{}
	r.Options = opts
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout.
func (r *RadarTcpResetsTimeoutService) Summary(ctx context.Context, query RadarTcpResetsTimeoutSummaryParams, opts ...option.RequestOption) (res *RadarTcpResetsTimeoutSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of connection stage by TCP connections terminated
// within the first 10 packets by a reset or timeout over time.
func (r *RadarTcpResetsTimeoutService) TimeseriesGroups(ctx context.Context, query RadarTcpResetsTimeoutTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarTcpResetsTimeoutTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/tcp_resets_timeouts/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarTcpResetsTimeoutSummaryResponse struct {
	Result  RadarTcpResetsTimeoutSummaryResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarTcpResetsTimeoutSummaryResponseJSON   `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseJSON contains the JSON metadata for the
// struct [RadarTcpResetsTimeoutSummaryResponse]
type radarTcpResetsTimeoutSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarTcpResetsTimeoutSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarTcpResetsTimeoutSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarTcpResetsTimeoutSummaryResponseResultJSON     `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultJSON contains the JSON metadata for
// the struct [RadarTcpResetsTimeoutSummaryResponseResult]
type radarTcpResetsTimeoutSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarTcpResetsTimeoutSummaryResponseResultMeta struct {
	ConfidenceInfo RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarTcpResetsTimeoutSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarTcpResetsTimeoutSummaryResponseResultMetaJSON   `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarTcpResetsTimeoutSummaryResponseResultMeta]
type radarTcpResetsTimeoutSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo]
type radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange]
type radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization string

const (
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationPercentage           RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationMin0Max              RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationMinMax               RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRawValues            RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationPercentageChange     RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRollingAverage       RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationOverlappedPercentage RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRatio                RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarTcpResetsTimeoutSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationPercentage, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationMin0Max, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationMinMax, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRawValues, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationPercentageChange, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRollingAverage, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarTcpResetsTimeoutSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarTcpResetsTimeoutSummaryResponseResultMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  radarTcpResetsTimeoutSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarTcpResetsTimeoutSummaryResponseResultMetaUnit]
type radarTcpResetsTimeoutSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryResponseResultSummary0 struct {
	// Connection resets within the first 10 packets from the client, but after the
	// server has received multiple data packets.
	LaterInFlow string `json:"later_in_flow,required"`
	// All other connections.
	NoMatch string `json:"no_match,required"`
	// Connection resets or timeouts after the server received both a SYN packet and an
	// ACK packet, meaning the connection was successfully established.
	PostAck string `json:"post_ack,required"`
	// Connection resets or timeouts after the server received a packet with PSH flag
	// set, following connection establishment.
	PostPsh string `json:"post_psh,required"`
	// Connection resets or timeouts after the server received only a single SYN
	// packet.
	PostSyn string                                                 `json:"post_syn,required"`
	JSON    radarTcpResetsTimeoutSummaryResponseResultSummary0JSON `json:"-"`
}

// radarTcpResetsTimeoutSummaryResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarTcpResetsTimeoutSummaryResponseResultSummary0]
type radarTcpResetsTimeoutSummaryResponseResultSummary0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponse struct {
	Result  RadarTcpResetsTimeoutTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarTcpResetsTimeoutTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseJSON contains the JSON metadata for
// the struct [RadarTcpResetsTimeoutTimeseriesGroupsResponse]
type radarTcpResetsTimeoutTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResult struct {
	// Metadata for the results.
	Meta   RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON contains the JSON
// metadata for the struct [RadarTcpResetsTimeoutTimeseriesGroupsResponseResult]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnit `json:"units,required"`
	JSON  radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON   `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneHour        RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval = "ONE_HOUR"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneDay         RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval = "ONE_DAY"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneWeek        RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval = "ONE_WEEK"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneMonth       RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneHour, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneDay, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneWeek, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                     `json:"level,required"`
	JSON  radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationPercentage           RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationMin0Max              RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "MIN0_MAX"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationMinMax               RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "MIN_MAX"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRawValues            RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "RAW_VALUES"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationPercentageChange     RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRollingAverage       RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRatio                RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization = "RATIO"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationPercentage, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationMin0Max, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationMinMax, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRawValues, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationPercentageChange, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRollingAverage, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage, RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnit struct {
	Name  string                                                          `json:"name,required"`
	Value string                                                          `json:"value,required"`
	JSON  radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnitJSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnit]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0 struct {
	LaterInFlow []string                                                      `json:"later_in_flow,required"`
	NoMatch     []string                                                      `json:"no_match,required"`
	PostAck     []string                                                      `json:"post_ack,required"`
	PostPsh     []string                                                      `json:"post_psh,required"`
	PostSyn     []string                                                      `json:"post_syn,required"`
	Timestamps  []time.Time                                                   `json:"timestamps,required" format:"date-time"`
	JSON        radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0]
type radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON struct {
	LaterInFlow apijson.Field
	NoMatch     apijson.Field
	PostAck     apijson.Field
	PostPsh     apijson.Field
	PostSyn     apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarTcpResetsTimeoutTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarTcpResetsTimeoutSummaryParams struct {
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
	Format param.Field[RadarTcpResetsTimeoutSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarTcpResetsTimeoutSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarTcpResetsTimeoutSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarTcpResetsTimeoutSummaryParamsFormat string

const (
	RadarTcpResetsTimeoutSummaryParamsFormatJson RadarTcpResetsTimeoutSummaryParamsFormat = "JSON"
	RadarTcpResetsTimeoutSummaryParamsFormatCsv  RadarTcpResetsTimeoutSummaryParamsFormat = "CSV"
)

func (r RadarTcpResetsTimeoutSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutSummaryParamsFormatJson, RadarTcpResetsTimeoutSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarTcpResetsTimeoutTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarTcpResetsTimeoutTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RadarTcpResetsTimeoutTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval15m RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "15m"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1h  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1h"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1d  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1d"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1w  RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval = "1w"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval15m, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1h, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1d, RadarTcpResetsTimeoutTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat string

const (
	RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatJson RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat = "JSON"
	RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatCsv  RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarTcpResetsTimeoutTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatJson, RadarTcpResetsTimeoutTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}
