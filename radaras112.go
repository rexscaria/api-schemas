// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// RadarAs112Service contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAs112Service] method instead.
type RadarAs112Service struct {
	Options          []option.RequestOption
	Summary          *RadarAs112SummaryService
	TimeseriesGroups *RadarAs112TimeseriesGroupService
	Top              *RadarAs112TopService
}

// NewRadarAs112Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAs112Service(opts ...option.RequestOption) (r *RadarAs112Service) {
	r = &RadarAs112Service{}
	r.Options = opts
	r.Summary = NewRadarAs112SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAs112TimeseriesGroupService(opts...)
	r.Top = NewRadarAs112TopService(opts...)
	return
}

// Retrieves the AS112 DNS queries over time.
func (r *RadarAs112Service) GetTimeseries(ctx context.Context, query RadarAs112GetTimeseriesParams, opts ...option.RequestOption) (res *RadarAs112GetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112GetTimeseriesResponse struct {
	Result  RadarAs112GetTimeseriesResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarAs112GetTimeseriesResponseJSON   `json:"-"`
}

// radarAs112GetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarAs112GetTimeseriesResponse]
type radarAs112GetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResult struct {
	Meta   RadarAs112GetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAs112GetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112GetTimeseriesResponseResultJSON   `json:"-"`
}

// radarAs112GetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112GetTimeseriesResponseResult]
type radarAs112GetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResultMeta struct {
	AggInterval    string                                                  `json:"aggInterval,required"`
	DateRange      []RadarAs112GetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                               `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAs112GetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112GetTimeseriesResponseResultMetaJSON           `json:"-"`
}

// radarAs112GetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarAs112GetTimeseriesResponseResultMeta]
type radarAs112GetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAs112GetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112GetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112GetTimeseriesResponseResultMetaDateRange]
type radarAs112GetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarAs112GetTimeseriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112GetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112GetTimeseriesResponseResultMetaConfidenceInfo]
type radarAs112GetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                     `json:"timestamps,required" format:"date-time"`
	Values     []string                                        `json:"values,required"`
	JSON       radarAs112GetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAs112GetTimeseriesResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarAs112GetTimeseriesResponseResultSerie0]
type radarAs112GetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112GetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112GetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112GetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112GetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAs112GetTimeseriesParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112GetTimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112GetTimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112GetTimeseriesParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112GetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAs112GetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112GetTimeseriesParamsAggInterval string

const (
	RadarAs112GetTimeseriesParamsAggInterval15m RadarAs112GetTimeseriesParamsAggInterval = "15m"
	RadarAs112GetTimeseriesParamsAggInterval1h  RadarAs112GetTimeseriesParamsAggInterval = "1h"
	RadarAs112GetTimeseriesParamsAggInterval1d  RadarAs112GetTimeseriesParamsAggInterval = "1d"
	RadarAs112GetTimeseriesParamsAggInterval1w  RadarAs112GetTimeseriesParamsAggInterval = "1w"
)

func (r RadarAs112GetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAs112GetTimeseriesParamsAggInterval15m, RadarAs112GetTimeseriesParamsAggInterval1h, RadarAs112GetTimeseriesParamsAggInterval1d, RadarAs112GetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112GetTimeseriesParamsFormat string

const (
	RadarAs112GetTimeseriesParamsFormatJson RadarAs112GetTimeseriesParamsFormat = "JSON"
	RadarAs112GetTimeseriesParamsFormatCsv  RadarAs112GetTimeseriesParamsFormat = "CSV"
)

func (r RadarAs112GetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112GetTimeseriesParamsFormatJson, RadarAs112GetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112GetTimeseriesParamsProtocol string

const (
	RadarAs112GetTimeseriesParamsProtocolUdp   RadarAs112GetTimeseriesParamsProtocol = "UDP"
	RadarAs112GetTimeseriesParamsProtocolTcp   RadarAs112GetTimeseriesParamsProtocol = "TCP"
	RadarAs112GetTimeseriesParamsProtocolHTTPS RadarAs112GetTimeseriesParamsProtocol = "HTTPS"
	RadarAs112GetTimeseriesParamsProtocolTls   RadarAs112GetTimeseriesParamsProtocol = "TLS"
)

func (r RadarAs112GetTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112GetTimeseriesParamsProtocolUdp, RadarAs112GetTimeseriesParamsProtocolTcp, RadarAs112GetTimeseriesParamsProtocolHTTPS, RadarAs112GetTimeseriesParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112GetTimeseriesParamsQueryType string

const (
	RadarAs112GetTimeseriesParamsQueryTypeA          RadarAs112GetTimeseriesParamsQueryType = "A"
	RadarAs112GetTimeseriesParamsQueryTypeAaaa       RadarAs112GetTimeseriesParamsQueryType = "AAAA"
	RadarAs112GetTimeseriesParamsQueryTypeA6         RadarAs112GetTimeseriesParamsQueryType = "A6"
	RadarAs112GetTimeseriesParamsQueryTypeAfsdb      RadarAs112GetTimeseriesParamsQueryType = "AFSDB"
	RadarAs112GetTimeseriesParamsQueryTypeAny        RadarAs112GetTimeseriesParamsQueryType = "ANY"
	RadarAs112GetTimeseriesParamsQueryTypeApl        RadarAs112GetTimeseriesParamsQueryType = "APL"
	RadarAs112GetTimeseriesParamsQueryTypeAtma       RadarAs112GetTimeseriesParamsQueryType = "ATMA"
	RadarAs112GetTimeseriesParamsQueryTypeAxfr       RadarAs112GetTimeseriesParamsQueryType = "AXFR"
	RadarAs112GetTimeseriesParamsQueryTypeCaa        RadarAs112GetTimeseriesParamsQueryType = "CAA"
	RadarAs112GetTimeseriesParamsQueryTypeCdnskey    RadarAs112GetTimeseriesParamsQueryType = "CDNSKEY"
	RadarAs112GetTimeseriesParamsQueryTypeCds        RadarAs112GetTimeseriesParamsQueryType = "CDS"
	RadarAs112GetTimeseriesParamsQueryTypeCert       RadarAs112GetTimeseriesParamsQueryType = "CERT"
	RadarAs112GetTimeseriesParamsQueryTypeCname      RadarAs112GetTimeseriesParamsQueryType = "CNAME"
	RadarAs112GetTimeseriesParamsQueryTypeCsync      RadarAs112GetTimeseriesParamsQueryType = "CSYNC"
	RadarAs112GetTimeseriesParamsQueryTypeDhcid      RadarAs112GetTimeseriesParamsQueryType = "DHCID"
	RadarAs112GetTimeseriesParamsQueryTypeDlv        RadarAs112GetTimeseriesParamsQueryType = "DLV"
	RadarAs112GetTimeseriesParamsQueryTypeDname      RadarAs112GetTimeseriesParamsQueryType = "DNAME"
	RadarAs112GetTimeseriesParamsQueryTypeDnskey     RadarAs112GetTimeseriesParamsQueryType = "DNSKEY"
	RadarAs112GetTimeseriesParamsQueryTypeDoa        RadarAs112GetTimeseriesParamsQueryType = "DOA"
	RadarAs112GetTimeseriesParamsQueryTypeDs         RadarAs112GetTimeseriesParamsQueryType = "DS"
	RadarAs112GetTimeseriesParamsQueryTypeEid        RadarAs112GetTimeseriesParamsQueryType = "EID"
	RadarAs112GetTimeseriesParamsQueryTypeEui48      RadarAs112GetTimeseriesParamsQueryType = "EUI48"
	RadarAs112GetTimeseriesParamsQueryTypeEui64      RadarAs112GetTimeseriesParamsQueryType = "EUI64"
	RadarAs112GetTimeseriesParamsQueryTypeGpos       RadarAs112GetTimeseriesParamsQueryType = "GPOS"
	RadarAs112GetTimeseriesParamsQueryTypeGid        RadarAs112GetTimeseriesParamsQueryType = "GID"
	RadarAs112GetTimeseriesParamsQueryTypeHinfo      RadarAs112GetTimeseriesParamsQueryType = "HINFO"
	RadarAs112GetTimeseriesParamsQueryTypeHip        RadarAs112GetTimeseriesParamsQueryType = "HIP"
	RadarAs112GetTimeseriesParamsQueryTypeHTTPS      RadarAs112GetTimeseriesParamsQueryType = "HTTPS"
	RadarAs112GetTimeseriesParamsQueryTypeIpseckey   RadarAs112GetTimeseriesParamsQueryType = "IPSECKEY"
	RadarAs112GetTimeseriesParamsQueryTypeIsdn       RadarAs112GetTimeseriesParamsQueryType = "ISDN"
	RadarAs112GetTimeseriesParamsQueryTypeIxfr       RadarAs112GetTimeseriesParamsQueryType = "IXFR"
	RadarAs112GetTimeseriesParamsQueryTypeKey        RadarAs112GetTimeseriesParamsQueryType = "KEY"
	RadarAs112GetTimeseriesParamsQueryTypeKx         RadarAs112GetTimeseriesParamsQueryType = "KX"
	RadarAs112GetTimeseriesParamsQueryTypeL32        RadarAs112GetTimeseriesParamsQueryType = "L32"
	RadarAs112GetTimeseriesParamsQueryTypeL64        RadarAs112GetTimeseriesParamsQueryType = "L64"
	RadarAs112GetTimeseriesParamsQueryTypeLoc        RadarAs112GetTimeseriesParamsQueryType = "LOC"
	RadarAs112GetTimeseriesParamsQueryTypeLp         RadarAs112GetTimeseriesParamsQueryType = "LP"
	RadarAs112GetTimeseriesParamsQueryTypeMaila      RadarAs112GetTimeseriesParamsQueryType = "MAILA"
	RadarAs112GetTimeseriesParamsQueryTypeMailb      RadarAs112GetTimeseriesParamsQueryType = "MAILB"
	RadarAs112GetTimeseriesParamsQueryTypeMB         RadarAs112GetTimeseriesParamsQueryType = "MB"
	RadarAs112GetTimeseriesParamsQueryTypeMd         RadarAs112GetTimeseriesParamsQueryType = "MD"
	RadarAs112GetTimeseriesParamsQueryTypeMf         RadarAs112GetTimeseriesParamsQueryType = "MF"
	RadarAs112GetTimeseriesParamsQueryTypeMg         RadarAs112GetTimeseriesParamsQueryType = "MG"
	RadarAs112GetTimeseriesParamsQueryTypeMinfo      RadarAs112GetTimeseriesParamsQueryType = "MINFO"
	RadarAs112GetTimeseriesParamsQueryTypeMr         RadarAs112GetTimeseriesParamsQueryType = "MR"
	RadarAs112GetTimeseriesParamsQueryTypeMx         RadarAs112GetTimeseriesParamsQueryType = "MX"
	RadarAs112GetTimeseriesParamsQueryTypeNaptr      RadarAs112GetTimeseriesParamsQueryType = "NAPTR"
	RadarAs112GetTimeseriesParamsQueryTypeNb         RadarAs112GetTimeseriesParamsQueryType = "NB"
	RadarAs112GetTimeseriesParamsQueryTypeNbstat     RadarAs112GetTimeseriesParamsQueryType = "NBSTAT"
	RadarAs112GetTimeseriesParamsQueryTypeNid        RadarAs112GetTimeseriesParamsQueryType = "NID"
	RadarAs112GetTimeseriesParamsQueryTypeNimloc     RadarAs112GetTimeseriesParamsQueryType = "NIMLOC"
	RadarAs112GetTimeseriesParamsQueryTypeNinfo      RadarAs112GetTimeseriesParamsQueryType = "NINFO"
	RadarAs112GetTimeseriesParamsQueryTypeNs         RadarAs112GetTimeseriesParamsQueryType = "NS"
	RadarAs112GetTimeseriesParamsQueryTypeNsap       RadarAs112GetTimeseriesParamsQueryType = "NSAP"
	RadarAs112GetTimeseriesParamsQueryTypeNsec       RadarAs112GetTimeseriesParamsQueryType = "NSEC"
	RadarAs112GetTimeseriesParamsQueryTypeNsec3      RadarAs112GetTimeseriesParamsQueryType = "NSEC3"
	RadarAs112GetTimeseriesParamsQueryTypeNsec3Param RadarAs112GetTimeseriesParamsQueryType = "NSEC3PARAM"
	RadarAs112GetTimeseriesParamsQueryTypeNull       RadarAs112GetTimeseriesParamsQueryType = "NULL"
	RadarAs112GetTimeseriesParamsQueryTypeNxt        RadarAs112GetTimeseriesParamsQueryType = "NXT"
	RadarAs112GetTimeseriesParamsQueryTypeOpenpgpkey RadarAs112GetTimeseriesParamsQueryType = "OPENPGPKEY"
	RadarAs112GetTimeseriesParamsQueryTypeOpt        RadarAs112GetTimeseriesParamsQueryType = "OPT"
	RadarAs112GetTimeseriesParamsQueryTypePtr        RadarAs112GetTimeseriesParamsQueryType = "PTR"
	RadarAs112GetTimeseriesParamsQueryTypePx         RadarAs112GetTimeseriesParamsQueryType = "PX"
	RadarAs112GetTimeseriesParamsQueryTypeRkey       RadarAs112GetTimeseriesParamsQueryType = "RKEY"
	RadarAs112GetTimeseriesParamsQueryTypeRp         RadarAs112GetTimeseriesParamsQueryType = "RP"
	RadarAs112GetTimeseriesParamsQueryTypeRrsig      RadarAs112GetTimeseriesParamsQueryType = "RRSIG"
	RadarAs112GetTimeseriesParamsQueryTypeRt         RadarAs112GetTimeseriesParamsQueryType = "RT"
	RadarAs112GetTimeseriesParamsQueryTypeSig        RadarAs112GetTimeseriesParamsQueryType = "SIG"
	RadarAs112GetTimeseriesParamsQueryTypeSink       RadarAs112GetTimeseriesParamsQueryType = "SINK"
	RadarAs112GetTimeseriesParamsQueryTypeSmimea     RadarAs112GetTimeseriesParamsQueryType = "SMIMEA"
	RadarAs112GetTimeseriesParamsQueryTypeSoa        RadarAs112GetTimeseriesParamsQueryType = "SOA"
	RadarAs112GetTimeseriesParamsQueryTypeSpf        RadarAs112GetTimeseriesParamsQueryType = "SPF"
	RadarAs112GetTimeseriesParamsQueryTypeSrv        RadarAs112GetTimeseriesParamsQueryType = "SRV"
	RadarAs112GetTimeseriesParamsQueryTypeSshfp      RadarAs112GetTimeseriesParamsQueryType = "SSHFP"
	RadarAs112GetTimeseriesParamsQueryTypeSvcb       RadarAs112GetTimeseriesParamsQueryType = "SVCB"
	RadarAs112GetTimeseriesParamsQueryTypeTa         RadarAs112GetTimeseriesParamsQueryType = "TA"
	RadarAs112GetTimeseriesParamsQueryTypeTalink     RadarAs112GetTimeseriesParamsQueryType = "TALINK"
	RadarAs112GetTimeseriesParamsQueryTypeTkey       RadarAs112GetTimeseriesParamsQueryType = "TKEY"
	RadarAs112GetTimeseriesParamsQueryTypeTlsa       RadarAs112GetTimeseriesParamsQueryType = "TLSA"
	RadarAs112GetTimeseriesParamsQueryTypeTsig       RadarAs112GetTimeseriesParamsQueryType = "TSIG"
	RadarAs112GetTimeseriesParamsQueryTypeTxt        RadarAs112GetTimeseriesParamsQueryType = "TXT"
	RadarAs112GetTimeseriesParamsQueryTypeUinfo      RadarAs112GetTimeseriesParamsQueryType = "UINFO"
	RadarAs112GetTimeseriesParamsQueryTypeUid        RadarAs112GetTimeseriesParamsQueryType = "UID"
	RadarAs112GetTimeseriesParamsQueryTypeUnspec     RadarAs112GetTimeseriesParamsQueryType = "UNSPEC"
	RadarAs112GetTimeseriesParamsQueryTypeUri        RadarAs112GetTimeseriesParamsQueryType = "URI"
	RadarAs112GetTimeseriesParamsQueryTypeWks        RadarAs112GetTimeseriesParamsQueryType = "WKS"
	RadarAs112GetTimeseriesParamsQueryTypeX25        RadarAs112GetTimeseriesParamsQueryType = "X25"
	RadarAs112GetTimeseriesParamsQueryTypeZonemd     RadarAs112GetTimeseriesParamsQueryType = "ZONEMD"
)

func (r RadarAs112GetTimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112GetTimeseriesParamsQueryTypeA, RadarAs112GetTimeseriesParamsQueryTypeAaaa, RadarAs112GetTimeseriesParamsQueryTypeA6, RadarAs112GetTimeseriesParamsQueryTypeAfsdb, RadarAs112GetTimeseriesParamsQueryTypeAny, RadarAs112GetTimeseriesParamsQueryTypeApl, RadarAs112GetTimeseriesParamsQueryTypeAtma, RadarAs112GetTimeseriesParamsQueryTypeAxfr, RadarAs112GetTimeseriesParamsQueryTypeCaa, RadarAs112GetTimeseriesParamsQueryTypeCdnskey, RadarAs112GetTimeseriesParamsQueryTypeCds, RadarAs112GetTimeseriesParamsQueryTypeCert, RadarAs112GetTimeseriesParamsQueryTypeCname, RadarAs112GetTimeseriesParamsQueryTypeCsync, RadarAs112GetTimeseriesParamsQueryTypeDhcid, RadarAs112GetTimeseriesParamsQueryTypeDlv, RadarAs112GetTimeseriesParamsQueryTypeDname, RadarAs112GetTimeseriesParamsQueryTypeDnskey, RadarAs112GetTimeseriesParamsQueryTypeDoa, RadarAs112GetTimeseriesParamsQueryTypeDs, RadarAs112GetTimeseriesParamsQueryTypeEid, RadarAs112GetTimeseriesParamsQueryTypeEui48, RadarAs112GetTimeseriesParamsQueryTypeEui64, RadarAs112GetTimeseriesParamsQueryTypeGpos, RadarAs112GetTimeseriesParamsQueryTypeGid, RadarAs112GetTimeseriesParamsQueryTypeHinfo, RadarAs112GetTimeseriesParamsQueryTypeHip, RadarAs112GetTimeseriesParamsQueryTypeHTTPS, RadarAs112GetTimeseriesParamsQueryTypeIpseckey, RadarAs112GetTimeseriesParamsQueryTypeIsdn, RadarAs112GetTimeseriesParamsQueryTypeIxfr, RadarAs112GetTimeseriesParamsQueryTypeKey, RadarAs112GetTimeseriesParamsQueryTypeKx, RadarAs112GetTimeseriesParamsQueryTypeL32, RadarAs112GetTimeseriesParamsQueryTypeL64, RadarAs112GetTimeseriesParamsQueryTypeLoc, RadarAs112GetTimeseriesParamsQueryTypeLp, RadarAs112GetTimeseriesParamsQueryTypeMaila, RadarAs112GetTimeseriesParamsQueryTypeMailb, RadarAs112GetTimeseriesParamsQueryTypeMB, RadarAs112GetTimeseriesParamsQueryTypeMd, RadarAs112GetTimeseriesParamsQueryTypeMf, RadarAs112GetTimeseriesParamsQueryTypeMg, RadarAs112GetTimeseriesParamsQueryTypeMinfo, RadarAs112GetTimeseriesParamsQueryTypeMr, RadarAs112GetTimeseriesParamsQueryTypeMx, RadarAs112GetTimeseriesParamsQueryTypeNaptr, RadarAs112GetTimeseriesParamsQueryTypeNb, RadarAs112GetTimeseriesParamsQueryTypeNbstat, RadarAs112GetTimeseriesParamsQueryTypeNid, RadarAs112GetTimeseriesParamsQueryTypeNimloc, RadarAs112GetTimeseriesParamsQueryTypeNinfo, RadarAs112GetTimeseriesParamsQueryTypeNs, RadarAs112GetTimeseriesParamsQueryTypeNsap, RadarAs112GetTimeseriesParamsQueryTypeNsec, RadarAs112GetTimeseriesParamsQueryTypeNsec3, RadarAs112GetTimeseriesParamsQueryTypeNsec3Param, RadarAs112GetTimeseriesParamsQueryTypeNull, RadarAs112GetTimeseriesParamsQueryTypeNxt, RadarAs112GetTimeseriesParamsQueryTypeOpenpgpkey, RadarAs112GetTimeseriesParamsQueryTypeOpt, RadarAs112GetTimeseriesParamsQueryTypePtr, RadarAs112GetTimeseriesParamsQueryTypePx, RadarAs112GetTimeseriesParamsQueryTypeRkey, RadarAs112GetTimeseriesParamsQueryTypeRp, RadarAs112GetTimeseriesParamsQueryTypeRrsig, RadarAs112GetTimeseriesParamsQueryTypeRt, RadarAs112GetTimeseriesParamsQueryTypeSig, RadarAs112GetTimeseriesParamsQueryTypeSink, RadarAs112GetTimeseriesParamsQueryTypeSmimea, RadarAs112GetTimeseriesParamsQueryTypeSoa, RadarAs112GetTimeseriesParamsQueryTypeSpf, RadarAs112GetTimeseriesParamsQueryTypeSrv, RadarAs112GetTimeseriesParamsQueryTypeSshfp, RadarAs112GetTimeseriesParamsQueryTypeSvcb, RadarAs112GetTimeseriesParamsQueryTypeTa, RadarAs112GetTimeseriesParamsQueryTypeTalink, RadarAs112GetTimeseriesParamsQueryTypeTkey, RadarAs112GetTimeseriesParamsQueryTypeTlsa, RadarAs112GetTimeseriesParamsQueryTypeTsig, RadarAs112GetTimeseriesParamsQueryTypeTxt, RadarAs112GetTimeseriesParamsQueryTypeUinfo, RadarAs112GetTimeseriesParamsQueryTypeUid, RadarAs112GetTimeseriesParamsQueryTypeUnspec, RadarAs112GetTimeseriesParamsQueryTypeUri, RadarAs112GetTimeseriesParamsQueryTypeWks, RadarAs112GetTimeseriesParamsQueryTypeX25, RadarAs112GetTimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112GetTimeseriesParamsResponseCode string

const (
	RadarAs112GetTimeseriesParamsResponseCodeNoerror   RadarAs112GetTimeseriesParamsResponseCode = "NOERROR"
	RadarAs112GetTimeseriesParamsResponseCodeFormerr   RadarAs112GetTimeseriesParamsResponseCode = "FORMERR"
	RadarAs112GetTimeseriesParamsResponseCodeServfail  RadarAs112GetTimeseriesParamsResponseCode = "SERVFAIL"
	RadarAs112GetTimeseriesParamsResponseCodeNxdomain  RadarAs112GetTimeseriesParamsResponseCode = "NXDOMAIN"
	RadarAs112GetTimeseriesParamsResponseCodeNotimp    RadarAs112GetTimeseriesParamsResponseCode = "NOTIMP"
	RadarAs112GetTimeseriesParamsResponseCodeRefused   RadarAs112GetTimeseriesParamsResponseCode = "REFUSED"
	RadarAs112GetTimeseriesParamsResponseCodeYxdomain  RadarAs112GetTimeseriesParamsResponseCode = "YXDOMAIN"
	RadarAs112GetTimeseriesParamsResponseCodeYxrrset   RadarAs112GetTimeseriesParamsResponseCode = "YXRRSET"
	RadarAs112GetTimeseriesParamsResponseCodeNxrrset   RadarAs112GetTimeseriesParamsResponseCode = "NXRRSET"
	RadarAs112GetTimeseriesParamsResponseCodeNotauth   RadarAs112GetTimeseriesParamsResponseCode = "NOTAUTH"
	RadarAs112GetTimeseriesParamsResponseCodeNotzone   RadarAs112GetTimeseriesParamsResponseCode = "NOTZONE"
	RadarAs112GetTimeseriesParamsResponseCodeBadsig    RadarAs112GetTimeseriesParamsResponseCode = "BADSIG"
	RadarAs112GetTimeseriesParamsResponseCodeBadkey    RadarAs112GetTimeseriesParamsResponseCode = "BADKEY"
	RadarAs112GetTimeseriesParamsResponseCodeBadtime   RadarAs112GetTimeseriesParamsResponseCode = "BADTIME"
	RadarAs112GetTimeseriesParamsResponseCodeBadmode   RadarAs112GetTimeseriesParamsResponseCode = "BADMODE"
	RadarAs112GetTimeseriesParamsResponseCodeBadname   RadarAs112GetTimeseriesParamsResponseCode = "BADNAME"
	RadarAs112GetTimeseriesParamsResponseCodeBadalg    RadarAs112GetTimeseriesParamsResponseCode = "BADALG"
	RadarAs112GetTimeseriesParamsResponseCodeBadtrunc  RadarAs112GetTimeseriesParamsResponseCode = "BADTRUNC"
	RadarAs112GetTimeseriesParamsResponseCodeBadcookie RadarAs112GetTimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112GetTimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112GetTimeseriesParamsResponseCodeNoerror, RadarAs112GetTimeseriesParamsResponseCodeFormerr, RadarAs112GetTimeseriesParamsResponseCodeServfail, RadarAs112GetTimeseriesParamsResponseCodeNxdomain, RadarAs112GetTimeseriesParamsResponseCodeNotimp, RadarAs112GetTimeseriesParamsResponseCodeRefused, RadarAs112GetTimeseriesParamsResponseCodeYxdomain, RadarAs112GetTimeseriesParamsResponseCodeYxrrset, RadarAs112GetTimeseriesParamsResponseCodeNxrrset, RadarAs112GetTimeseriesParamsResponseCodeNotauth, RadarAs112GetTimeseriesParamsResponseCodeNotzone, RadarAs112GetTimeseriesParamsResponseCodeBadsig, RadarAs112GetTimeseriesParamsResponseCodeBadkey, RadarAs112GetTimeseriesParamsResponseCodeBadtime, RadarAs112GetTimeseriesParamsResponseCodeBadmode, RadarAs112GetTimeseriesParamsResponseCodeBadname, RadarAs112GetTimeseriesParamsResponseCodeBadalg, RadarAs112GetTimeseriesParamsResponseCodeBadtrunc, RadarAs112GetTimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}
