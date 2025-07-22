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

// RadarAs112TimeseriesGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAs112TimeseriesGroupService] method instead.
type RadarAs112TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupService) {
	r = &RadarAs112TimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of AS112 DNS queries by DNSSEC (DNS Security
// Extensions) support over time.
func (r *RadarAs112TimeseriesGroupService) GetDnssec(ctx context.Context, query RadarAs112TimeseriesGroupGetDnssecParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS queries by EDNS (Extension Mechanisms
// for DNS) support over time.
func (r *RadarAs112TimeseriesGroupService) GetEdns(ctx context.Context, query RadarAs112TimeseriesGroupGetEdnsParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS queries by IP version over time.
func (r *RadarAs112TimeseriesGroupService) GetIPVersion(ctx context.Context, query RadarAs112TimeseriesGroupGetIPVersionParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS requests classified by protocol over
// time.
func (r *RadarAs112TimeseriesGroupService) GetProtocol(ctx context.Context, query RadarAs112TimeseriesGroupGetProtocolParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS queries by type over time.
func (r *RadarAs112TimeseriesGroupService) GetQueryType(ctx context.Context, query RadarAs112TimeseriesGroupGetQueryTypeParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS requests classified by response code
// over time.
func (r *RadarAs112TimeseriesGroupService) GetResponseCodes(ctx context.Context, query RadarAs112TimeseriesGroupGetResponseCodesParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupGetResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupGetDnssecResponse struct {
	Result  RadarAs112TimeseriesGroupGetDnssecResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetDnssecResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetDnssecResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupGetDnssecResponse]
type radarAs112TimeseriesGroupGetDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetDnssecResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetDnssecResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetDnssecResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetDnssecResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupGetDnssecResponseResult]
type radarAs112TimeseriesGroupGetDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetDnssecResponseResultSerie0 struct {
	NotSupported []string                                                   `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                   `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupGetDnssecResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetDnssecResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupGetDnssecResponseResultSerie0]
type radarAs112TimeseriesGroupGetDnssecResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetDnssecResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetDnssecResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetEdnsResponse struct {
	Result  RadarAs112TimeseriesGroupGetEdnsResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetEdnsResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetEdnsResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupGetEdnsResponse]
type radarAs112TimeseriesGroupGetEdnsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetEdnsResponseResult struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetEdnsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetEdnsResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetEdnsResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupGetEdnsResponseResult]
type radarAs112TimeseriesGroupGetEdnsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetEdnsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetEdnsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetEdnsResponseResultSerie0 struct {
	NotSupported []string                                                 `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                 `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupGetEdnsResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetEdnsResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupGetEdnsResponseResultSerie0]
type radarAs112TimeseriesGroupGetEdnsResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetEdnsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetEdnsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetIPVersionResponse struct {
	Result  RadarAs112TimeseriesGroupGetIPVersionResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetIPVersionResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupGetIPVersionResponse]
type radarAs112TimeseriesGroupGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetIPVersionResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetIPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetIPVersionResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetIPVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupGetIPVersionResponseResult]
type radarAs112TimeseriesGroupGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetIPVersionResponseResultSerie0 struct {
	IPv4 []string                                                      `json:"IPv4,required"`
	IPv6 []string                                                      `json:"IPv6,required"`
	JSON radarAs112TimeseriesGroupGetIPVersionResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetIPVersionResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupGetIPVersionResponseResultSerie0]
type radarAs112TimeseriesGroupGetIPVersionResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetIPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetIPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetProtocolResponse struct {
	Result  RadarAs112TimeseriesGroupGetProtocolResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetProtocolResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetProtocolResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupGetProtocolResponse]
type radarAs112TimeseriesGroupGetProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetProtocolResponseResult struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetProtocolResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetProtocolResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetProtocolResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupGetProtocolResponseResult]
type radarAs112TimeseriesGroupGetProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetProtocolResponseResultSerie0 struct {
	HTTPS []string                                                     `json:"HTTPS,required"`
	Tcp   []string                                                     `json:"TCP,required"`
	Tls   []string                                                     `json:"TLS,required"`
	Udp   []string                                                     `json:"UDP,required"`
	JSON  radarAs112TimeseriesGroupGetProtocolResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetProtocolResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupGetProtocolResponseResultSerie0]
type radarAs112TimeseriesGroupGetProtocolResponseResultSerie0JSON struct {
	HTTPS       apijson.Field
	Tcp         apijson.Field
	Tls         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetProtocolResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetProtocolResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetQueryTypeResponse struct {
	Result  RadarAs112TimeseriesGroupGetQueryTypeResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetQueryTypeResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetQueryTypeResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupGetQueryTypeResponse]
type radarAs112TimeseriesGroupGetQueryTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetQueryTypeResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetQueryTypeResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetQueryTypeResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupGetQueryTypeResponseResult]
type radarAs112TimeseriesGroupGetQueryTypeResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetQueryTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetQueryTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0 struct {
	Timestamps  []string                                                      `json:"timestamps,required"`
	ExtraFields map[string][]string                                           `json:"-,extras"`
	JSON        radarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0]
type radarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetQueryTypeResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetResponseCodesResponse struct {
	Result  RadarAs112TimeseriesGroupGetResponseCodesResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAs112TimeseriesGroupGetResponseCodesResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetResponseCodesResponseJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupGetResponseCodesResponse]
type radarAs112TimeseriesGroupGetResponseCodesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetResponseCodesResponseResult struct {
	Meta   interface{}                                                   `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupGetResponseCodesResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupGetResponseCodesResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupGetResponseCodesResponseResult]
type radarAs112TimeseriesGroupGetResponseCodesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetResponseCodesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetResponseCodesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0 struct {
	Timestamps  []string                                                          `json:"timestamps,required"`
	ExtraFields map[string][]string                                               `json:"-,extras"`
	JSON        radarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0]
type radarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TimeseriesGroupGetResponseCodesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TimeseriesGroupGetDnssecParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetDnssecParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetDnssecParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112TimeseriesGroupGetDnssecParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112TimeseriesGroupGetDnssecParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112TimeseriesGroupGetDnssecParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetDnssecParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupGetDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetDnssecParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetDnssecParamsAggInterval15m RadarAs112TimeseriesGroupGetDnssecParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1h  RadarAs112TimeseriesGroupGetDnssecParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1d  RadarAs112TimeseriesGroupGetDnssecParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1w  RadarAs112TimeseriesGroupGetDnssecParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetDnssecParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetDnssecParamsAggInterval15m, RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1h, RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1d, RadarAs112TimeseriesGroupGetDnssecParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetDnssecParamsFormat string

const (
	RadarAs112TimeseriesGroupGetDnssecParamsFormatJson RadarAs112TimeseriesGroupGetDnssecParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetDnssecParamsFormatCsv  RadarAs112TimeseriesGroupGetDnssecParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetDnssecParamsFormatJson, RadarAs112TimeseriesGroupGetDnssecParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112TimeseriesGroupGetDnssecParamsProtocol string

const (
	RadarAs112TimeseriesGroupGetDnssecParamsProtocolUdp   RadarAs112TimeseriesGroupGetDnssecParamsProtocol = "UDP"
	RadarAs112TimeseriesGroupGetDnssecParamsProtocolTcp   RadarAs112TimeseriesGroupGetDnssecParamsProtocol = "TCP"
	RadarAs112TimeseriesGroupGetDnssecParamsProtocolHTTPS RadarAs112TimeseriesGroupGetDnssecParamsProtocol = "HTTPS"
	RadarAs112TimeseriesGroupGetDnssecParamsProtocolTls   RadarAs112TimeseriesGroupGetDnssecParamsProtocol = "TLS"
)

func (r RadarAs112TimeseriesGroupGetDnssecParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetDnssecParamsProtocolUdp, RadarAs112TimeseriesGroupGetDnssecParamsProtocolTcp, RadarAs112TimeseriesGroupGetDnssecParamsProtocolHTTPS, RadarAs112TimeseriesGroupGetDnssecParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112TimeseriesGroupGetDnssecParamsQueryType string

const (
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeA          RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "A"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAaaa       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "AAAA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeA6         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "A6"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAfsdb      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "AFSDB"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAny        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "ANY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeApl        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "APL"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAtma       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "ATMA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAxfr       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "AXFR"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCaa        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CAA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCdnskey    RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CDNSKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCds        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CDS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCert       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CERT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCname      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CNAME"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCsync      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "CSYNC"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDhcid      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DHCID"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDlv        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DLV"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDname      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DNAME"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDnskey     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DNSKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDoa        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DOA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDs         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "DS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEid        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "EID"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEui48      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "EUI48"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEui64      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "EUI64"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeGpos       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "GPOS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeGid        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "GID"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHinfo      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "HINFO"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHip        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "HIP"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHTTPS      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "HTTPS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIpseckey   RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "IPSECKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIsdn       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "ISDN"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIxfr       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "IXFR"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeKey        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "KEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeKx         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "KX"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeL32        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "L32"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeL64        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "L64"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeLoc        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "LOC"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeLp         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "LP"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMaila      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MAILA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMailb      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MAILB"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMB         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MB"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMd         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MD"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMf         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MF"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMg         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MG"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMinfo      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MINFO"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMr         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MR"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMx         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "MX"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNaptr      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NAPTR"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNb         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NB"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNbstat     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NBSTAT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNid        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NID"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNimloc     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NIMLOC"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNinfo      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NINFO"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNs         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsap       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NSAP"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NSEC"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec3      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NSEC3"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec3Param RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NSEC3PARAM"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNull       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NULL"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNxt        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "NXT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeOpenpgpkey RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "OPENPGPKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeOpt        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "OPT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypePtr        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "PTR"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypePx         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "PX"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRkey       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "RKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRp         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "RP"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRrsig      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "RRSIG"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRt         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "RT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSig        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SIG"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSink       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SINK"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSmimea     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SMIMEA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSoa        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SOA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSpf        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SPF"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSrv        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SRV"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSshfp      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SSHFP"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSvcb       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "SVCB"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTa         RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTalink     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TALINK"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTkey       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTlsa       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TLSA"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTsig       RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TSIG"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTxt        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "TXT"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUinfo      RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "UINFO"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUid        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "UID"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUnspec     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "UNSPEC"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUri        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "URI"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeWks        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "WKS"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeX25        RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "X25"
	RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeZonemd     RadarAs112TimeseriesGroupGetDnssecParamsQueryType = "ZONEMD"
)

func (r RadarAs112TimeseriesGroupGetDnssecParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeA, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAaaa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeA6, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAfsdb, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAny, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeApl, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAtma, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeAxfr, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCaa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCdnskey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCds, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCert, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCname, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeCsync, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDhcid, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDlv, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDname, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDnskey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDoa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeDs, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEid, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEui48, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeEui64, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeGpos, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeGid, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHinfo, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHip, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeHTTPS, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIpseckey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIsdn, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeIxfr, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeKey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeKx, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeL32, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeL64, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeLoc, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeLp, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMaila, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMailb, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMB, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMd, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMf, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMg, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMinfo, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMr, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeMx, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNaptr, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNb, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNbstat, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNid, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNimloc, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNinfo, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNs, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsap, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec3, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNsec3Param, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNull, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeNxt, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeOpenpgpkey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeOpt, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypePtr, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypePx, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRkey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRp, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRrsig, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeRt, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSig, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSink, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSmimea, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSoa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSpf, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSrv, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSshfp, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeSvcb, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTalink, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTkey, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTlsa, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTsig, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeTxt, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUinfo, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUid, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUnspec, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeUri, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeWks, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeX25, RadarAs112TimeseriesGroupGetDnssecParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112TimeseriesGroupGetDnssecParamsResponseCode string

const (
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNoerror   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NOERROR"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeFormerr   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "FORMERR"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeServfail  RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "SERVFAIL"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNxdomain  RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NXDOMAIN"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotimp    RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NOTIMP"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeRefused   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "REFUSED"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeYxdomain  RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "YXDOMAIN"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeYxrrset   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "YXRRSET"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNxrrset   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NXRRSET"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotauth   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NOTAUTH"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotzone   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "NOTZONE"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadsig    RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADSIG"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadkey    RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADKEY"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadtime   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADTIME"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadmode   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADMODE"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadname   RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADNAME"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadalg    RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADALG"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadtrunc  RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADTRUNC"
	RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadcookie RadarAs112TimeseriesGroupGetDnssecParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112TimeseriesGroupGetDnssecParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNoerror, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeFormerr, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeServfail, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNxdomain, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotimp, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeRefused, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeYxdomain, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeYxrrset, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNxrrset, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotauth, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeNotzone, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadsig, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadkey, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadtime, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadmode, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadname, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadalg, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadtrunc, RadarAs112TimeseriesGroupGetDnssecParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112TimeseriesGroupGetEdnsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetEdnsParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetEdnsParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112TimeseriesGroupGetEdnsParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112TimeseriesGroupGetEdnsParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112TimeseriesGroupGetEdnsParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetEdnsParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseriesGroupGetEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetEdnsParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetEdnsParamsAggInterval15m RadarAs112TimeseriesGroupGetEdnsParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1h  RadarAs112TimeseriesGroupGetEdnsParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1d  RadarAs112TimeseriesGroupGetEdnsParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1w  RadarAs112TimeseriesGroupGetEdnsParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetEdnsParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetEdnsParamsAggInterval15m, RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1h, RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1d, RadarAs112TimeseriesGroupGetEdnsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetEdnsParamsFormat string

const (
	RadarAs112TimeseriesGroupGetEdnsParamsFormatJson RadarAs112TimeseriesGroupGetEdnsParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetEdnsParamsFormatCsv  RadarAs112TimeseriesGroupGetEdnsParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetEdnsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetEdnsParamsFormatJson, RadarAs112TimeseriesGroupGetEdnsParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112TimeseriesGroupGetEdnsParamsProtocol string

const (
	RadarAs112TimeseriesGroupGetEdnsParamsProtocolUdp   RadarAs112TimeseriesGroupGetEdnsParamsProtocol = "UDP"
	RadarAs112TimeseriesGroupGetEdnsParamsProtocolTcp   RadarAs112TimeseriesGroupGetEdnsParamsProtocol = "TCP"
	RadarAs112TimeseriesGroupGetEdnsParamsProtocolHTTPS RadarAs112TimeseriesGroupGetEdnsParamsProtocol = "HTTPS"
	RadarAs112TimeseriesGroupGetEdnsParamsProtocolTls   RadarAs112TimeseriesGroupGetEdnsParamsProtocol = "TLS"
)

func (r RadarAs112TimeseriesGroupGetEdnsParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetEdnsParamsProtocolUdp, RadarAs112TimeseriesGroupGetEdnsParamsProtocolTcp, RadarAs112TimeseriesGroupGetEdnsParamsProtocolHTTPS, RadarAs112TimeseriesGroupGetEdnsParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112TimeseriesGroupGetEdnsParamsQueryType string

const (
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeA          RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "A"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAaaa       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "AAAA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeA6         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "A6"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAfsdb      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "AFSDB"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAny        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "ANY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeApl        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "APL"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAtma       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "ATMA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAxfr       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "AXFR"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCaa        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CAA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCdnskey    RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CDNSKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCds        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CDS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCert       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CERT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCname      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CNAME"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCsync      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "CSYNC"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDhcid      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DHCID"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDlv        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DLV"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDname      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DNAME"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDnskey     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DNSKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDoa        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DOA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDs         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "DS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEid        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "EID"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEui48      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "EUI48"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEui64      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "EUI64"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeGpos       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "GPOS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeGid        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "GID"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHinfo      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "HINFO"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHip        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "HIP"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHTTPS      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "HTTPS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIpseckey   RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "IPSECKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIsdn       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "ISDN"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIxfr       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "IXFR"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeKey        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "KEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeKx         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "KX"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeL32        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "L32"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeL64        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "L64"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeLoc        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "LOC"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeLp         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "LP"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMaila      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MAILA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMailb      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MAILB"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMB         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MB"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMd         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MD"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMf         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MF"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMg         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MG"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMinfo      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MINFO"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMr         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MR"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMx         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "MX"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNaptr      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NAPTR"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNb         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NB"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNbstat     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NBSTAT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNid        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NID"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNimloc     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NIMLOC"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNinfo      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NINFO"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNs         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsap       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NSAP"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NSEC"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec3      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NSEC3"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec3Param RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NSEC3PARAM"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNull       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NULL"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNxt        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "NXT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeOpenpgpkey RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "OPENPGPKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeOpt        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "OPT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypePtr        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "PTR"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypePx         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "PX"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRkey       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "RKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRp         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "RP"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRrsig      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "RRSIG"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRt         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "RT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSig        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SIG"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSink       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SINK"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSmimea     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SMIMEA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSoa        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SOA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSpf        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SPF"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSrv        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SRV"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSshfp      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SSHFP"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSvcb       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "SVCB"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTa         RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTalink     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TALINK"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTkey       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTlsa       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TLSA"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTsig       RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TSIG"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTxt        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "TXT"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUinfo      RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "UINFO"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUid        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "UID"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUnspec     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "UNSPEC"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUri        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "URI"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeWks        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "WKS"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeX25        RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "X25"
	RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeZonemd     RadarAs112TimeseriesGroupGetEdnsParamsQueryType = "ZONEMD"
)

func (r RadarAs112TimeseriesGroupGetEdnsParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeA, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAaaa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeA6, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAfsdb, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAny, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeApl, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAtma, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeAxfr, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCaa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCdnskey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCds, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCert, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCname, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeCsync, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDhcid, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDlv, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDname, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDnskey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDoa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeDs, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEid, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEui48, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeEui64, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeGpos, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeGid, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHinfo, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHip, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeHTTPS, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIpseckey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIsdn, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeIxfr, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeKey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeKx, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeL32, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeL64, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeLoc, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeLp, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMaila, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMailb, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMB, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMd, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMf, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMg, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMinfo, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMr, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeMx, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNaptr, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNb, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNbstat, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNid, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNimloc, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNinfo, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNs, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsap, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec3, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNsec3Param, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNull, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeNxt, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeOpenpgpkey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeOpt, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypePtr, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypePx, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRkey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRp, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRrsig, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeRt, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSig, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSink, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSmimea, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSoa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSpf, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSrv, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSshfp, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeSvcb, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTalink, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTkey, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTlsa, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTsig, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeTxt, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUinfo, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUid, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUnspec, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeUri, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeWks, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeX25, RadarAs112TimeseriesGroupGetEdnsParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112TimeseriesGroupGetEdnsParamsResponseCode string

const (
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNoerror   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NOERROR"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeFormerr   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "FORMERR"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeServfail  RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "SERVFAIL"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNxdomain  RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NXDOMAIN"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotimp    RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NOTIMP"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeRefused   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "REFUSED"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeYxdomain  RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "YXDOMAIN"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeYxrrset   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "YXRRSET"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNxrrset   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NXRRSET"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotauth   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NOTAUTH"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotzone   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "NOTZONE"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadsig    RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADSIG"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadkey    RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADKEY"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadtime   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADTIME"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadmode   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADMODE"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadname   RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADNAME"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadalg    RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADALG"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadtrunc  RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADTRUNC"
	RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadcookie RadarAs112TimeseriesGroupGetEdnsParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112TimeseriesGroupGetEdnsParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNoerror, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeFormerr, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeServfail, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNxdomain, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotimp, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeRefused, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeYxdomain, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeYxrrset, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNxrrset, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotauth, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeNotzone, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadsig, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadkey, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadtime, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadmode, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadname, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadalg, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadtrunc, RadarAs112TimeseriesGroupGetEdnsParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112TimeseriesGroupGetIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetIPVersionParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112TimeseriesGroupGetIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112TimeseriesGroupGetIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval15m RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1h  RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1d  RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1w  RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval15m, RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1h, RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1d, RadarAs112TimeseriesGroupGetIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetIPVersionParamsFormat string

const (
	RadarAs112TimeseriesGroupGetIPVersionParamsFormatJson RadarAs112TimeseriesGroupGetIPVersionParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetIPVersionParamsFormatCsv  RadarAs112TimeseriesGroupGetIPVersionParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetIPVersionParamsFormatJson, RadarAs112TimeseriesGroupGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112TimeseriesGroupGetIPVersionParamsProtocol string

const (
	RadarAs112TimeseriesGroupGetIPVersionParamsProtocolUdp   RadarAs112TimeseriesGroupGetIPVersionParamsProtocol = "UDP"
	RadarAs112TimeseriesGroupGetIPVersionParamsProtocolTcp   RadarAs112TimeseriesGroupGetIPVersionParamsProtocol = "TCP"
	RadarAs112TimeseriesGroupGetIPVersionParamsProtocolHTTPS RadarAs112TimeseriesGroupGetIPVersionParamsProtocol = "HTTPS"
	RadarAs112TimeseriesGroupGetIPVersionParamsProtocolTls   RadarAs112TimeseriesGroupGetIPVersionParamsProtocol = "TLS"
)

func (r RadarAs112TimeseriesGroupGetIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetIPVersionParamsProtocolUdp, RadarAs112TimeseriesGroupGetIPVersionParamsProtocolTcp, RadarAs112TimeseriesGroupGetIPVersionParamsProtocolHTTPS, RadarAs112TimeseriesGroupGetIPVersionParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112TimeseriesGroupGetIPVersionParamsQueryType string

const (
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeA          RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "A"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAaaa       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "AAAA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeA6         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "A6"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAfsdb      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "AFSDB"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAny        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "ANY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeApl        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "APL"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAtma       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "ATMA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAxfr       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "AXFR"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCaa        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CAA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCdnskey    RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CDNSKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCds        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CDS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCert       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CERT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCname      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CNAME"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCsync      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "CSYNC"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDhcid      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DHCID"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDlv        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DLV"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDname      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DNAME"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDnskey     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DNSKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDoa        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DOA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDs         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "DS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEid        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "EID"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEui48      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "EUI48"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEui64      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "EUI64"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeGpos       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "GPOS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeGid        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "GID"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHinfo      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "HINFO"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHip        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "HIP"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHTTPS      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "HTTPS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIpseckey   RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "IPSECKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIsdn       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "ISDN"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIxfr       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "IXFR"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeKey        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "KEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeKx         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "KX"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeL32        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "L32"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeL64        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "L64"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeLoc        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "LOC"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeLp         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "LP"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMaila      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MAILA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMailb      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MAILB"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMB         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MB"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMd         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MD"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMf         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MF"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMg         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MG"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMinfo      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MINFO"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMr         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MR"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMx         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "MX"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNaptr      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NAPTR"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNb         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NB"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNbstat     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NBSTAT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNid        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NID"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNimloc     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NIMLOC"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNinfo      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NINFO"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNs         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsap       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NSAP"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NSEC"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec3      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NSEC3"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec3Param RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NSEC3PARAM"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNull       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NULL"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNxt        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "NXT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeOpenpgpkey RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "OPENPGPKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeOpt        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "OPT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypePtr        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "PTR"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypePx         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "PX"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRkey       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "RKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRp         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "RP"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRrsig      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "RRSIG"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRt         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "RT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSig        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SIG"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSink       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SINK"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSmimea     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SMIMEA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSoa        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SOA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSpf        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SPF"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSrv        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SRV"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSshfp      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SSHFP"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSvcb       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "SVCB"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTa         RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTalink     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TALINK"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTkey       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTlsa       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TLSA"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTsig       RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TSIG"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTxt        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "TXT"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUinfo      RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "UINFO"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUid        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "UID"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUnspec     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "UNSPEC"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUri        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "URI"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeWks        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "WKS"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeX25        RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "X25"
	RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeZonemd     RadarAs112TimeseriesGroupGetIPVersionParamsQueryType = "ZONEMD"
)

func (r RadarAs112TimeseriesGroupGetIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeA, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAaaa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeA6, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAfsdb, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAny, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeApl, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAtma, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeAxfr, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCaa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCdnskey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCds, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCert, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCname, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeCsync, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDhcid, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDlv, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDname, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDnskey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDoa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeDs, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEid, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEui48, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeEui64, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeGpos, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeGid, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHinfo, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHip, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeHTTPS, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIpseckey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIsdn, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeIxfr, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeKey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeKx, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeL32, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeL64, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeLoc, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeLp, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMaila, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMailb, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMB, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMd, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMf, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMg, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMinfo, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMr, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeMx, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNaptr, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNb, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNbstat, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNid, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNimloc, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNinfo, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNs, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsap, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec3, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNsec3Param, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNull, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeNxt, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeOpenpgpkey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeOpt, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypePtr, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypePx, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRkey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRp, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRrsig, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeRt, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSig, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSink, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSmimea, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSoa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSpf, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSrv, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSshfp, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeSvcb, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTalink, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTkey, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTlsa, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTsig, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeTxt, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUinfo, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUid, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUnspec, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeUri, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeWks, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeX25, RadarAs112TimeseriesGroupGetIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode string

const (
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNoerror   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NOERROR"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeFormerr   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "FORMERR"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeServfail  RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "SERVFAIL"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNxdomain  RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NXDOMAIN"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotimp    RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NOTIMP"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeRefused   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "REFUSED"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeYxdomain  RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "YXDOMAIN"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeYxrrset   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "YXRRSET"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNxrrset   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NXRRSET"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotauth   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NOTAUTH"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotzone   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "NOTZONE"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadsig    RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADSIG"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadkey    RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADKEY"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadtime   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADTIME"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadmode   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADMODE"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadname   RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADNAME"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadalg    RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADALG"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadtrunc  RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADTRUNC"
	RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadcookie RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112TimeseriesGroupGetIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNoerror, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeFormerr, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeServfail, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNxdomain, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotimp, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeRefused, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeYxdomain, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeYxrrset, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNxrrset, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotauth, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeNotzone, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadsig, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadkey, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadtime, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadmode, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadname, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadalg, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadtrunc, RadarAs112TimeseriesGroupGetIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112TimeseriesGroupGetProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetProtocolParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetProtocolParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112TimeseriesGroupGetProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112TimeseriesGroupGetProtocolParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupGetProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetProtocolParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetProtocolParamsAggInterval15m RadarAs112TimeseriesGroupGetProtocolParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1h  RadarAs112TimeseriesGroupGetProtocolParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1d  RadarAs112TimeseriesGroupGetProtocolParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1w  RadarAs112TimeseriesGroupGetProtocolParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetProtocolParamsAggInterval15m, RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1h, RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1d, RadarAs112TimeseriesGroupGetProtocolParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetProtocolParamsFormat string

const (
	RadarAs112TimeseriesGroupGetProtocolParamsFormatJson RadarAs112TimeseriesGroupGetProtocolParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetProtocolParamsFormatCsv  RadarAs112TimeseriesGroupGetProtocolParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetProtocolParamsFormatJson, RadarAs112TimeseriesGroupGetProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112TimeseriesGroupGetProtocolParamsQueryType string

const (
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeA          RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "A"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAaaa       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "AAAA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeA6         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "A6"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAfsdb      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "AFSDB"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAny        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "ANY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeApl        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "APL"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAtma       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "ATMA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAxfr       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "AXFR"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCaa        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CAA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCdnskey    RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CDNSKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCds        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CDS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCert       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CERT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCname      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CNAME"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCsync      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "CSYNC"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDhcid      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DHCID"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDlv        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DLV"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDname      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DNAME"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDnskey     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DNSKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDoa        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DOA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDs         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "DS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEid        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "EID"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEui48      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "EUI48"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEui64      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "EUI64"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeGpos       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "GPOS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeGid        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "GID"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHinfo      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "HINFO"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHip        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "HIP"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHTTPS      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "HTTPS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIpseckey   RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "IPSECKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIsdn       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "ISDN"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIxfr       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "IXFR"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeKey        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "KEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeKx         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "KX"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeL32        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "L32"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeL64        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "L64"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeLoc        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "LOC"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeLp         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "LP"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMaila      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MAILA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMailb      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MAILB"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMB         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MB"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMd         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MD"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMf         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MF"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMg         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MG"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMinfo      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MINFO"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMr         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MR"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMx         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "MX"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNaptr      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NAPTR"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNb         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NB"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNbstat     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NBSTAT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNid        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NID"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNimloc     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NIMLOC"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNinfo      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NINFO"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNs         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsap       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NSAP"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NSEC"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec3      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NSEC3"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec3Param RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NSEC3PARAM"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNull       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NULL"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNxt        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "NXT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeOpenpgpkey RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "OPENPGPKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeOpt        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "OPT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypePtr        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "PTR"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypePx         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "PX"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRkey       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "RKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRp         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "RP"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRrsig      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "RRSIG"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRt         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "RT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSig        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SIG"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSink       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SINK"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSmimea     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SMIMEA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSoa        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SOA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSpf        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SPF"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSrv        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SRV"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSshfp      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SSHFP"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSvcb       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "SVCB"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTa         RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTalink     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TALINK"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTkey       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTlsa       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TLSA"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTsig       RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TSIG"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTxt        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "TXT"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUinfo      RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "UINFO"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUid        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "UID"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUnspec     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "UNSPEC"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUri        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "URI"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeWks        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "WKS"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeX25        RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "X25"
	RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeZonemd     RadarAs112TimeseriesGroupGetProtocolParamsQueryType = "ZONEMD"
)

func (r RadarAs112TimeseriesGroupGetProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeA, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAaaa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeA6, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAfsdb, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAny, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeApl, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAtma, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeAxfr, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCaa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCdnskey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCds, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCert, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCname, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeCsync, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDhcid, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDlv, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDname, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDnskey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDoa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeDs, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEid, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEui48, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeEui64, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeGpos, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeGid, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHinfo, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHip, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeHTTPS, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIpseckey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIsdn, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeIxfr, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeKey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeKx, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeL32, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeL64, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeLoc, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeLp, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMaila, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMailb, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMB, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMd, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMf, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMg, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMinfo, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMr, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeMx, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNaptr, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNb, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNbstat, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNid, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNimloc, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNinfo, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNs, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsap, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec3, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNsec3Param, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNull, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeNxt, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeOpenpgpkey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeOpt, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypePtr, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypePx, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRkey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRp, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRrsig, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeRt, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSig, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSink, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSmimea, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSoa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSpf, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSrv, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSshfp, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeSvcb, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTalink, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTkey, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTlsa, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTsig, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeTxt, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUinfo, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUid, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUnspec, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeUri, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeWks, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeX25, RadarAs112TimeseriesGroupGetProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112TimeseriesGroupGetProtocolParamsResponseCode string

const (
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNoerror   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NOERROR"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeFormerr   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "FORMERR"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeServfail  RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "SERVFAIL"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNxdomain  RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NXDOMAIN"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotimp    RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NOTIMP"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeRefused   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "REFUSED"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeYxdomain  RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "YXDOMAIN"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeYxrrset   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "YXRRSET"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNxrrset   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NXRRSET"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotauth   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NOTAUTH"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotzone   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "NOTZONE"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadsig    RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADSIG"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadkey    RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADKEY"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadtime   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADTIME"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadmode   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADMODE"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadname   RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADNAME"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadalg    RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADALG"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadtrunc  RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADTRUNC"
	RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadcookie RadarAs112TimeseriesGroupGetProtocolParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112TimeseriesGroupGetProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNoerror, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeFormerr, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeServfail, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNxdomain, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotimp, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeRefused, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeYxdomain, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeYxrrset, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNxrrset, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotauth, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeNotzone, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadsig, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadkey, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadtime, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadmode, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadname, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadalg, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadtrunc, RadarAs112TimeseriesGroupGetProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112TimeseriesGroupGetQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetQueryTypeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetQueryTypeParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupGetQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval15m RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1h  RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1d  RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1w  RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval15m, RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1h, RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1d, RadarAs112TimeseriesGroupGetQueryTypeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetQueryTypeParamsFormat string

const (
	RadarAs112TimeseriesGroupGetQueryTypeParamsFormatJson RadarAs112TimeseriesGroupGetQueryTypeParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetQueryTypeParamsFormatCsv  RadarAs112TimeseriesGroupGetQueryTypeParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetQueryTypeParamsFormatJson, RadarAs112TimeseriesGroupGetQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol string

const (
	RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolUdp   RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol = "UDP"
	RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolTcp   RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol = "TCP"
	RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolHTTPS RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol = "HTTPS"
	RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolTls   RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol = "TLS"
)

func (r RadarAs112TimeseriesGroupGetQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolUdp, RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolTcp, RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolHTTPS, RadarAs112TimeseriesGroupGetQueryTypeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode string

const (
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNoerror   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NOERROR"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeFormerr   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "FORMERR"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeServfail  RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "SERVFAIL"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNxdomain  RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NXDOMAIN"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotimp    RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NOTIMP"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeRefused   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "REFUSED"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeYxdomain  RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "YXDOMAIN"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeYxrrset   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "YXRRSET"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNxrrset   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NXRRSET"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotauth   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NOTAUTH"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotzone   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "NOTZONE"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadsig    RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADSIG"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadkey    RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADKEY"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadtime   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADTIME"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadmode   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADMODE"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadname   RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADNAME"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadalg    RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADALG"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadtrunc  RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADTRUNC"
	RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadcookie RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNoerror, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeFormerr, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeServfail, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNxdomain, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotimp, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeRefused, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeYxdomain, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeYxrrset, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNxrrset, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotauth, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeNotzone, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadsig, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadkey, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadtime, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadmode, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadname, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadalg, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadtrunc, RadarAs112TimeseriesGroupGetQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112TimeseriesGroupGetResponseCodesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval] `query:"aggInterval"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAs112TimeseriesGroupGetResponseCodesParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType] `query:"queryType"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupGetResponseCodesParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupGetResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval string

const (
	RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval15m RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1h  RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1d  RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1w  RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval = "1w"
)

func (r RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval15m, RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1h, RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1d, RadarAs112TimeseriesGroupGetResponseCodesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TimeseriesGroupGetResponseCodesParamsFormat string

const (
	RadarAs112TimeseriesGroupGetResponseCodesParamsFormatJson RadarAs112TimeseriesGroupGetResponseCodesParamsFormat = "JSON"
	RadarAs112TimeseriesGroupGetResponseCodesParamsFormatCsv  RadarAs112TimeseriesGroupGetResponseCodesParamsFormat = "CSV"
)

func (r RadarAs112TimeseriesGroupGetResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetResponseCodesParamsFormatJson, RadarAs112TimeseriesGroupGetResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol string

const (
	RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolUdp   RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol = "UDP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolTcp   RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol = "TCP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolHTTPS RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol = "HTTPS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolTls   RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol = "TLS"
)

func (r RadarAs112TimeseriesGroupGetResponseCodesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolUdp, RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolTcp, RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolHTTPS, RadarAs112TimeseriesGroupGetResponseCodesParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType string

const (
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeA          RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "A"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAaaa       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "AAAA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeA6         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "A6"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAfsdb      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "AFSDB"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAny        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "ANY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeApl        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "APL"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAtma       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "ATMA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAxfr       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "AXFR"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCaa        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CAA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCdnskey    RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CDNSKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCds        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CDS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCert       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CERT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCname      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CNAME"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCsync      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "CSYNC"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDhcid      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DHCID"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDlv        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DLV"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDname      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DNAME"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDnskey     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DNSKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDoa        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DOA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDs         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "DS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEid        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "EID"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEui48      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "EUI48"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEui64      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "EUI64"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeGpos       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "GPOS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeGid        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "GID"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHinfo      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "HINFO"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHip        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "HIP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHTTPS      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "HTTPS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIpseckey   RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "IPSECKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIsdn       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "ISDN"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIxfr       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "IXFR"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeKey        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "KEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeKx         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "KX"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeL32        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "L32"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeL64        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "L64"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeLoc        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "LOC"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeLp         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "LP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMaila      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MAILA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMailb      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MAILB"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMB         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MB"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMd         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MD"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMf         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MF"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMg         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MG"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMinfo      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MINFO"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMr         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MR"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMx         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "MX"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNaptr      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NAPTR"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNb         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NB"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNbstat     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NBSTAT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNid        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NID"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNimloc     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NIMLOC"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNinfo      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NINFO"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNs         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsap       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NSAP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NSEC"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec3      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NSEC3"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec3Param RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NSEC3PARAM"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNull       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NULL"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNxt        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "NXT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeOpenpgpkey RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "OPENPGPKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeOpt        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "OPT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypePtr        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "PTR"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypePx         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "PX"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRkey       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "RKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRp         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "RP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRrsig      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "RRSIG"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRt         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "RT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSig        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SIG"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSink       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SINK"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSmimea     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SMIMEA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSoa        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SOA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSpf        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SPF"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSrv        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SRV"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSshfp      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SSHFP"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSvcb       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "SVCB"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTa         RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTalink     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TALINK"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTkey       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TKEY"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTlsa       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TLSA"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTsig       RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TSIG"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTxt        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "TXT"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUinfo      RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "UINFO"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUid        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "UID"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUnspec     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "UNSPEC"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUri        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "URI"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeWks        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "WKS"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeX25        RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "X25"
	RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeZonemd     RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType = "ZONEMD"
)

func (r RadarAs112TimeseriesGroupGetResponseCodesParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeA, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAaaa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeA6, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAfsdb, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAny, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeApl, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAtma, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeAxfr, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCaa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCdnskey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCds, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCert, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCname, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeCsync, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDhcid, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDlv, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDname, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDnskey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDoa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeDs, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEid, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEui48, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeEui64, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeGpos, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeGid, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHinfo, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHip, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeHTTPS, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIpseckey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIsdn, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeIxfr, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeKey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeKx, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeL32, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeL64, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeLoc, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeLp, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMaila, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMailb, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMB, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMd, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMf, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMg, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMinfo, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMr, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeMx, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNaptr, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNb, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNbstat, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNid, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNimloc, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNinfo, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNs, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsap, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec3, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNsec3Param, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNull, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeNxt, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeOpenpgpkey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeOpt, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypePtr, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypePx, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRkey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRp, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRrsig, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeRt, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSig, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSink, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSmimea, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSoa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSpf, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSrv, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSshfp, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeSvcb, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTalink, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTkey, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTlsa, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTsig, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeTxt, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUinfo, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUid, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUnspec, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeUri, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeWks, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeX25, RadarAs112TimeseriesGroupGetResponseCodesParamsQueryTypeZonemd:
		return true
	}
	return false
}
