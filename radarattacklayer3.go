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
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarAttackLayer3GetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3GetTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3GetTimeseriesResponseResult]
type radarAttackLayer3GetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3GetTimeseriesResponseResultSerie0 struct {
	Timestamps []string                                               `json:"timestamps,required"`
	Values     []string                                               `json:"values,required"`
	JSON       radarAttackLayer3GetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3GetTimeseriesResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3GetTimeseriesResponseResultSerie0]
type radarAttackLayer3GetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3GetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3GetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3GetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3GetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3GetTimeseriesParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3GetTimeseriesParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3GetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[RadarAttackLayer3GetTimeseriesParamsMetric] `query:"metric"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3GetTimeseriesParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
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

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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

// Normalization method applied. Refer to
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
