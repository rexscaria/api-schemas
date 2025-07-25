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

// RadarAttackLayer3Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3Service] method instead.
type RadarAttackLayer3Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer3SummaryService
	TimeseriesGroups *RadarAttackLayer3TimeseriesGroupService
	Top              *RadarAttackLayer3TopService
}

// NewRadarAttackLayer3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3Service(opts ...option.RequestOption) (r *RadarAttackLayer3Service) {
	r = &RadarAttackLayer3Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer3SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer3TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer3TopService(opts...)
	return
}

// Retrieves layer 3 attacks over time.
func (r *RadarAttackLayer3Service) GetTimeseries(ctx context.Context, query RadarAttackLayer3GetTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3GetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3GetTimeseriesResponse struct {
	Result  RadarAttackLayer3GetTimeseriesResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAttackLayer3GetTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3GetTimeseriesResponse]
type radarAttackLayer3GetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3GetTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta        RadarAttackLayer3GetTimeseriesResponseResultMeta        `json:"meta,required"`
	ExtraFields map[string]RadarAttackLayer3GetTimeseriesResponseResult `json:"-,extras"`
	JSON        radarAttackLayer3GetTimeseriesResponseResultJSON        `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3GetTimeseriesResponseResult]
type radarAttackLayer3GetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3GetTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3GetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3GetTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3GetTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3GetTimeseriesResponseResultMeta]
type radarAttackLayer3GetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval string

const (
	RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneHour        RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneDay         RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneWeek        RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneMonth       RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarAttackLayer3GetTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneHour, RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneDay, RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneWeek, RadarAttackLayer3GetTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                              `json:"level,required"`
	JSON  radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3GetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3GetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3GetTimeseriesResponseResultMetaDateRange]
type radarAttackLayer3GetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization string

const (
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationPercentage           RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationMin0Max              RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationMinMax               RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRawValues            RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRatio                RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3GetTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationPercentage, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationMin0Max, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationMinMax, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRawValues, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3GetTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3GetTimeseriesResponseResultMetaUnit struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAttackLayer3GetTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAttackLayer3GetTimeseriesResponseResultMetaUnit]
type radarAttackLayer3GetTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3GetTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3GetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3GetTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3GetTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3GetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[RadarAttackLayer3GetTimeseriesParamsMetric] `query:"metric"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3GetTimeseriesParamsNormalization] `query:"normalization"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3GetTimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3GetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3GetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3GetTimeseriesParamsAggInterval string

const (
	RadarAttackLayer3GetTimeseriesParamsAggInterval15m RadarAttackLayer3GetTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3GetTimeseriesParamsAggInterval1h  RadarAttackLayer3GetTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3GetTimeseriesParamsAggInterval1d  RadarAttackLayer3GetTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3GetTimeseriesParamsAggInterval1w  RadarAttackLayer3GetTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer3GetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsAggInterval15m, RadarAttackLayer3GetTimeseriesParamsAggInterval1h, RadarAttackLayer3GetTimeseriesParamsAggInterval1d, RadarAttackLayer3GetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3GetTimeseriesParamsDirection string

const (
	RadarAttackLayer3GetTimeseriesParamsDirectionOrigin RadarAttackLayer3GetTimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3GetTimeseriesParamsDirectionTarget RadarAttackLayer3GetTimeseriesParamsDirection = "TARGET"
)

func (r RadarAttackLayer3GetTimeseriesParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsDirectionOrigin, RadarAttackLayer3GetTimeseriesParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3GetTimeseriesParamsFormat string

const (
	RadarAttackLayer3GetTimeseriesParamsFormatJson RadarAttackLayer3GetTimeseriesParamsFormat = "JSON"
	RadarAttackLayer3GetTimeseriesParamsFormatCsv  RadarAttackLayer3GetTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer3GetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsFormatJson, RadarAttackLayer3GetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3GetTimeseriesParamsIPVersion string

const (
	RadarAttackLayer3GetTimeseriesParamsIPVersionIPv4 RadarAttackLayer3GetTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3GetTimeseriesParamsIPVersionIPv6 RadarAttackLayer3GetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3GetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsIPVersionIPv4, RadarAttackLayer3GetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Measurement units, eg. bytes.
type RadarAttackLayer3GetTimeseriesParamsMetric string

const (
	RadarAttackLayer3GetTimeseriesParamsMetricBytes    RadarAttackLayer3GetTimeseriesParamsMetric = "BYTES"
	RadarAttackLayer3GetTimeseriesParamsMetricBytesOld RadarAttackLayer3GetTimeseriesParamsMetric = "BYTES_OLD"
)

func (r RadarAttackLayer3GetTimeseriesParamsMetric) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsMetricBytes, RadarAttackLayer3GetTimeseriesParamsMetricBytesOld:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3GetTimeseriesParamsNormalization string

const (
	RadarAttackLayer3GetTimeseriesParamsNormalizationPercentageChange RadarAttackLayer3GetTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3GetTimeseriesParamsNormalizationMin0Max          RadarAttackLayer3GetTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer3GetTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsNormalizationPercentageChange, RadarAttackLayer3GetTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarAttackLayer3GetTimeseriesParamsProtocol string

const (
	RadarAttackLayer3GetTimeseriesParamsProtocolUdp  RadarAttackLayer3GetTimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3GetTimeseriesParamsProtocolTcp  RadarAttackLayer3GetTimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3GetTimeseriesParamsProtocolIcmp RadarAttackLayer3GetTimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3GetTimeseriesParamsProtocolGre  RadarAttackLayer3GetTimeseriesParamsProtocol = "GRE"
)

func (r RadarAttackLayer3GetTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3GetTimeseriesParamsProtocolUdp, RadarAttackLayer3GetTimeseriesParamsProtocolTcp, RadarAttackLayer3GetTimeseriesParamsProtocolIcmp, RadarAttackLayer3GetTimeseriesParamsProtocolGre:
		return true
	}
	return false
}
