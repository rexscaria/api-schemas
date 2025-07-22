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

// RadarDNSService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarDNSService] method instead.
type RadarDNSService struct {
	Options          []option.RequestOption
	Summary          *RadarDNSSummaryService
	TimeseriesGroups *RadarDNSTimeseriesGroupService
	Top              *RadarDNSTopService
}

// NewRadarDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDNSService(opts ...option.RequestOption) (r *RadarDNSService) {
	r = &RadarDNSService{}
	r.Options = opts
	r.Summary = NewRadarDNSSummaryService(opts...)
	r.TimeseriesGroups = NewRadarDNSTimeseriesGroupService(opts...)
	r.Top = NewRadarDNSTopService(opts...)
	return
}

// Retrieves normalized query volume to the 1.1.1.1 DNS resolver over time.
func (r *RadarDNSService) GetTimeseries(ctx context.Context, query RadarDNSGetTimeseriesParams, opts ...option.RequestOption) (res *RadarDNSGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSGetTimeseriesResponse struct {
	Result  RadarDNSGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarDNSGetTimeseriesResponseJSON   `json:"-"`
}

// radarDNSGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarDNSGetTimeseriesResponse]
type radarDNSGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResult struct {
	Meta   RadarDNSGetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarDNSGetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSGetTimeseriesResponseResultJSON   `json:"-"`
}

// radarDNSGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSGetTimeseriesResponseResult]
type radarDNSGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResultMeta struct {
	AggInterval    string                                                `json:"aggInterval,required"`
	DateRange      []RadarDNSGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                             `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarDNSGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSGetTimeseriesResponseResultMetaJSON           `json:"-"`
}

// radarDNSGetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarDNSGetTimeseriesResponseResultMeta]
type radarDNSGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarDNSGetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSGetTimeseriesResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarDNSGetTimeseriesResponseResultMetaDateRange]
type radarDNSGetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarDNSGetTimeseriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSGetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarDNSGetTimeseriesResponseResultMetaConfidenceInfo]
type radarDNSGetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                   `json:"timestamps,required" format:"date-time"`
	Values     []string                                      `json:"values,required"`
	JSON       radarDNSGetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarDNSGetTimeseriesResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarDNSGetTimeseriesResponseResultSerie0]
type radarDNSGetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSGetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSGetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSGetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSGetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarDNSGetTimeseriesParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSGetTimeseriesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSGetTimeseriesParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSGetTimeseriesParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarDNSGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSGetTimeseriesParamsAggInterval string

const (
	RadarDNSGetTimeseriesParamsAggInterval15m RadarDNSGetTimeseriesParamsAggInterval = "15m"
	RadarDNSGetTimeseriesParamsAggInterval1h  RadarDNSGetTimeseriesParamsAggInterval = "1h"
	RadarDNSGetTimeseriesParamsAggInterval1d  RadarDNSGetTimeseriesParamsAggInterval = "1d"
	RadarDNSGetTimeseriesParamsAggInterval1w  RadarDNSGetTimeseriesParamsAggInterval = "1w"
)

func (r RadarDNSGetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSGetTimeseriesParamsAggInterval15m, RadarDNSGetTimeseriesParamsAggInterval1h, RadarDNSGetTimeseriesParamsAggInterval1d, RadarDNSGetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSGetTimeseriesParamsFormat string

const (
	RadarDNSGetTimeseriesParamsFormatJson RadarDNSGetTimeseriesParamsFormat = "JSON"
	RadarDNSGetTimeseriesParamsFormatCsv  RadarDNSGetTimeseriesParamsFormat = "CSV"
)

func (r RadarDNSGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSGetTimeseriesParamsFormatJson, RadarDNSGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSGetTimeseriesParamsProtocol string

const (
	RadarDNSGetTimeseriesParamsProtocolUdp   RadarDNSGetTimeseriesParamsProtocol = "UDP"
	RadarDNSGetTimeseriesParamsProtocolTcp   RadarDNSGetTimeseriesParamsProtocol = "TCP"
	RadarDNSGetTimeseriesParamsProtocolHTTPS RadarDNSGetTimeseriesParamsProtocol = "HTTPS"
	RadarDNSGetTimeseriesParamsProtocolTls   RadarDNSGetTimeseriesParamsProtocol = "TLS"
)

func (r RadarDNSGetTimeseriesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSGetTimeseriesParamsProtocolUdp, RadarDNSGetTimeseriesParamsProtocolTcp, RadarDNSGetTimeseriesParamsProtocolHTTPS, RadarDNSGetTimeseriesParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSGetTimeseriesParamsQueryType string

const (
	RadarDNSGetTimeseriesParamsQueryTypeA          RadarDNSGetTimeseriesParamsQueryType = "A"
	RadarDNSGetTimeseriesParamsQueryTypeAaaa       RadarDNSGetTimeseriesParamsQueryType = "AAAA"
	RadarDNSGetTimeseriesParamsQueryTypeA6         RadarDNSGetTimeseriesParamsQueryType = "A6"
	RadarDNSGetTimeseriesParamsQueryTypeAfsdb      RadarDNSGetTimeseriesParamsQueryType = "AFSDB"
	RadarDNSGetTimeseriesParamsQueryTypeAny        RadarDNSGetTimeseriesParamsQueryType = "ANY"
	RadarDNSGetTimeseriesParamsQueryTypeApl        RadarDNSGetTimeseriesParamsQueryType = "APL"
	RadarDNSGetTimeseriesParamsQueryTypeAtma       RadarDNSGetTimeseriesParamsQueryType = "ATMA"
	RadarDNSGetTimeseriesParamsQueryTypeAxfr       RadarDNSGetTimeseriesParamsQueryType = "AXFR"
	RadarDNSGetTimeseriesParamsQueryTypeCaa        RadarDNSGetTimeseriesParamsQueryType = "CAA"
	RadarDNSGetTimeseriesParamsQueryTypeCdnskey    RadarDNSGetTimeseriesParamsQueryType = "CDNSKEY"
	RadarDNSGetTimeseriesParamsQueryTypeCds        RadarDNSGetTimeseriesParamsQueryType = "CDS"
	RadarDNSGetTimeseriesParamsQueryTypeCert       RadarDNSGetTimeseriesParamsQueryType = "CERT"
	RadarDNSGetTimeseriesParamsQueryTypeCname      RadarDNSGetTimeseriesParamsQueryType = "CNAME"
	RadarDNSGetTimeseriesParamsQueryTypeCsync      RadarDNSGetTimeseriesParamsQueryType = "CSYNC"
	RadarDNSGetTimeseriesParamsQueryTypeDhcid      RadarDNSGetTimeseriesParamsQueryType = "DHCID"
	RadarDNSGetTimeseriesParamsQueryTypeDlv        RadarDNSGetTimeseriesParamsQueryType = "DLV"
	RadarDNSGetTimeseriesParamsQueryTypeDname      RadarDNSGetTimeseriesParamsQueryType = "DNAME"
	RadarDNSGetTimeseriesParamsQueryTypeDnskey     RadarDNSGetTimeseriesParamsQueryType = "DNSKEY"
	RadarDNSGetTimeseriesParamsQueryTypeDoa        RadarDNSGetTimeseriesParamsQueryType = "DOA"
	RadarDNSGetTimeseriesParamsQueryTypeDs         RadarDNSGetTimeseriesParamsQueryType = "DS"
	RadarDNSGetTimeseriesParamsQueryTypeEid        RadarDNSGetTimeseriesParamsQueryType = "EID"
	RadarDNSGetTimeseriesParamsQueryTypeEui48      RadarDNSGetTimeseriesParamsQueryType = "EUI48"
	RadarDNSGetTimeseriesParamsQueryTypeEui64      RadarDNSGetTimeseriesParamsQueryType = "EUI64"
	RadarDNSGetTimeseriesParamsQueryTypeGpos       RadarDNSGetTimeseriesParamsQueryType = "GPOS"
	RadarDNSGetTimeseriesParamsQueryTypeGid        RadarDNSGetTimeseriesParamsQueryType = "GID"
	RadarDNSGetTimeseriesParamsQueryTypeHinfo      RadarDNSGetTimeseriesParamsQueryType = "HINFO"
	RadarDNSGetTimeseriesParamsQueryTypeHip        RadarDNSGetTimeseriesParamsQueryType = "HIP"
	RadarDNSGetTimeseriesParamsQueryTypeHTTPS      RadarDNSGetTimeseriesParamsQueryType = "HTTPS"
	RadarDNSGetTimeseriesParamsQueryTypeIpseckey   RadarDNSGetTimeseriesParamsQueryType = "IPSECKEY"
	RadarDNSGetTimeseriesParamsQueryTypeIsdn       RadarDNSGetTimeseriesParamsQueryType = "ISDN"
	RadarDNSGetTimeseriesParamsQueryTypeIxfr       RadarDNSGetTimeseriesParamsQueryType = "IXFR"
	RadarDNSGetTimeseriesParamsQueryTypeKey        RadarDNSGetTimeseriesParamsQueryType = "KEY"
	RadarDNSGetTimeseriesParamsQueryTypeKx         RadarDNSGetTimeseriesParamsQueryType = "KX"
	RadarDNSGetTimeseriesParamsQueryTypeL32        RadarDNSGetTimeseriesParamsQueryType = "L32"
	RadarDNSGetTimeseriesParamsQueryTypeL64        RadarDNSGetTimeseriesParamsQueryType = "L64"
	RadarDNSGetTimeseriesParamsQueryTypeLoc        RadarDNSGetTimeseriesParamsQueryType = "LOC"
	RadarDNSGetTimeseriesParamsQueryTypeLp         RadarDNSGetTimeseriesParamsQueryType = "LP"
	RadarDNSGetTimeseriesParamsQueryTypeMaila      RadarDNSGetTimeseriesParamsQueryType = "MAILA"
	RadarDNSGetTimeseriesParamsQueryTypeMailb      RadarDNSGetTimeseriesParamsQueryType = "MAILB"
	RadarDNSGetTimeseriesParamsQueryTypeMB         RadarDNSGetTimeseriesParamsQueryType = "MB"
	RadarDNSGetTimeseriesParamsQueryTypeMd         RadarDNSGetTimeseriesParamsQueryType = "MD"
	RadarDNSGetTimeseriesParamsQueryTypeMf         RadarDNSGetTimeseriesParamsQueryType = "MF"
	RadarDNSGetTimeseriesParamsQueryTypeMg         RadarDNSGetTimeseriesParamsQueryType = "MG"
	RadarDNSGetTimeseriesParamsQueryTypeMinfo      RadarDNSGetTimeseriesParamsQueryType = "MINFO"
	RadarDNSGetTimeseriesParamsQueryTypeMr         RadarDNSGetTimeseriesParamsQueryType = "MR"
	RadarDNSGetTimeseriesParamsQueryTypeMx         RadarDNSGetTimeseriesParamsQueryType = "MX"
	RadarDNSGetTimeseriesParamsQueryTypeNaptr      RadarDNSGetTimeseriesParamsQueryType = "NAPTR"
	RadarDNSGetTimeseriesParamsQueryTypeNb         RadarDNSGetTimeseriesParamsQueryType = "NB"
	RadarDNSGetTimeseriesParamsQueryTypeNbstat     RadarDNSGetTimeseriesParamsQueryType = "NBSTAT"
	RadarDNSGetTimeseriesParamsQueryTypeNid        RadarDNSGetTimeseriesParamsQueryType = "NID"
	RadarDNSGetTimeseriesParamsQueryTypeNimloc     RadarDNSGetTimeseriesParamsQueryType = "NIMLOC"
	RadarDNSGetTimeseriesParamsQueryTypeNinfo      RadarDNSGetTimeseriesParamsQueryType = "NINFO"
	RadarDNSGetTimeseriesParamsQueryTypeNs         RadarDNSGetTimeseriesParamsQueryType = "NS"
	RadarDNSGetTimeseriesParamsQueryTypeNsap       RadarDNSGetTimeseriesParamsQueryType = "NSAP"
	RadarDNSGetTimeseriesParamsQueryTypeNsec       RadarDNSGetTimeseriesParamsQueryType = "NSEC"
	RadarDNSGetTimeseriesParamsQueryTypeNsec3      RadarDNSGetTimeseriesParamsQueryType = "NSEC3"
	RadarDNSGetTimeseriesParamsQueryTypeNsec3Param RadarDNSGetTimeseriesParamsQueryType = "NSEC3PARAM"
	RadarDNSGetTimeseriesParamsQueryTypeNull       RadarDNSGetTimeseriesParamsQueryType = "NULL"
	RadarDNSGetTimeseriesParamsQueryTypeNxt        RadarDNSGetTimeseriesParamsQueryType = "NXT"
	RadarDNSGetTimeseriesParamsQueryTypeOpenpgpkey RadarDNSGetTimeseriesParamsQueryType = "OPENPGPKEY"
	RadarDNSGetTimeseriesParamsQueryTypeOpt        RadarDNSGetTimeseriesParamsQueryType = "OPT"
	RadarDNSGetTimeseriesParamsQueryTypePtr        RadarDNSGetTimeseriesParamsQueryType = "PTR"
	RadarDNSGetTimeseriesParamsQueryTypePx         RadarDNSGetTimeseriesParamsQueryType = "PX"
	RadarDNSGetTimeseriesParamsQueryTypeRkey       RadarDNSGetTimeseriesParamsQueryType = "RKEY"
	RadarDNSGetTimeseriesParamsQueryTypeRp         RadarDNSGetTimeseriesParamsQueryType = "RP"
	RadarDNSGetTimeseriesParamsQueryTypeRrsig      RadarDNSGetTimeseriesParamsQueryType = "RRSIG"
	RadarDNSGetTimeseriesParamsQueryTypeRt         RadarDNSGetTimeseriesParamsQueryType = "RT"
	RadarDNSGetTimeseriesParamsQueryTypeSig        RadarDNSGetTimeseriesParamsQueryType = "SIG"
	RadarDNSGetTimeseriesParamsQueryTypeSink       RadarDNSGetTimeseriesParamsQueryType = "SINK"
	RadarDNSGetTimeseriesParamsQueryTypeSmimea     RadarDNSGetTimeseriesParamsQueryType = "SMIMEA"
	RadarDNSGetTimeseriesParamsQueryTypeSoa        RadarDNSGetTimeseriesParamsQueryType = "SOA"
	RadarDNSGetTimeseriesParamsQueryTypeSpf        RadarDNSGetTimeseriesParamsQueryType = "SPF"
	RadarDNSGetTimeseriesParamsQueryTypeSrv        RadarDNSGetTimeseriesParamsQueryType = "SRV"
	RadarDNSGetTimeseriesParamsQueryTypeSshfp      RadarDNSGetTimeseriesParamsQueryType = "SSHFP"
	RadarDNSGetTimeseriesParamsQueryTypeSvcb       RadarDNSGetTimeseriesParamsQueryType = "SVCB"
	RadarDNSGetTimeseriesParamsQueryTypeTa         RadarDNSGetTimeseriesParamsQueryType = "TA"
	RadarDNSGetTimeseriesParamsQueryTypeTalink     RadarDNSGetTimeseriesParamsQueryType = "TALINK"
	RadarDNSGetTimeseriesParamsQueryTypeTkey       RadarDNSGetTimeseriesParamsQueryType = "TKEY"
	RadarDNSGetTimeseriesParamsQueryTypeTlsa       RadarDNSGetTimeseriesParamsQueryType = "TLSA"
	RadarDNSGetTimeseriesParamsQueryTypeTsig       RadarDNSGetTimeseriesParamsQueryType = "TSIG"
	RadarDNSGetTimeseriesParamsQueryTypeTxt        RadarDNSGetTimeseriesParamsQueryType = "TXT"
	RadarDNSGetTimeseriesParamsQueryTypeUinfo      RadarDNSGetTimeseriesParamsQueryType = "UINFO"
	RadarDNSGetTimeseriesParamsQueryTypeUid        RadarDNSGetTimeseriesParamsQueryType = "UID"
	RadarDNSGetTimeseriesParamsQueryTypeUnspec     RadarDNSGetTimeseriesParamsQueryType = "UNSPEC"
	RadarDNSGetTimeseriesParamsQueryTypeUri        RadarDNSGetTimeseriesParamsQueryType = "URI"
	RadarDNSGetTimeseriesParamsQueryTypeWks        RadarDNSGetTimeseriesParamsQueryType = "WKS"
	RadarDNSGetTimeseriesParamsQueryTypeX25        RadarDNSGetTimeseriesParamsQueryType = "X25"
	RadarDNSGetTimeseriesParamsQueryTypeZonemd     RadarDNSGetTimeseriesParamsQueryType = "ZONEMD"
)

func (r RadarDNSGetTimeseriesParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSGetTimeseriesParamsQueryTypeA, RadarDNSGetTimeseriesParamsQueryTypeAaaa, RadarDNSGetTimeseriesParamsQueryTypeA6, RadarDNSGetTimeseriesParamsQueryTypeAfsdb, RadarDNSGetTimeseriesParamsQueryTypeAny, RadarDNSGetTimeseriesParamsQueryTypeApl, RadarDNSGetTimeseriesParamsQueryTypeAtma, RadarDNSGetTimeseriesParamsQueryTypeAxfr, RadarDNSGetTimeseriesParamsQueryTypeCaa, RadarDNSGetTimeseriesParamsQueryTypeCdnskey, RadarDNSGetTimeseriesParamsQueryTypeCds, RadarDNSGetTimeseriesParamsQueryTypeCert, RadarDNSGetTimeseriesParamsQueryTypeCname, RadarDNSGetTimeseriesParamsQueryTypeCsync, RadarDNSGetTimeseriesParamsQueryTypeDhcid, RadarDNSGetTimeseriesParamsQueryTypeDlv, RadarDNSGetTimeseriesParamsQueryTypeDname, RadarDNSGetTimeseriesParamsQueryTypeDnskey, RadarDNSGetTimeseriesParamsQueryTypeDoa, RadarDNSGetTimeseriesParamsQueryTypeDs, RadarDNSGetTimeseriesParamsQueryTypeEid, RadarDNSGetTimeseriesParamsQueryTypeEui48, RadarDNSGetTimeseriesParamsQueryTypeEui64, RadarDNSGetTimeseriesParamsQueryTypeGpos, RadarDNSGetTimeseriesParamsQueryTypeGid, RadarDNSGetTimeseriesParamsQueryTypeHinfo, RadarDNSGetTimeseriesParamsQueryTypeHip, RadarDNSGetTimeseriesParamsQueryTypeHTTPS, RadarDNSGetTimeseriesParamsQueryTypeIpseckey, RadarDNSGetTimeseriesParamsQueryTypeIsdn, RadarDNSGetTimeseriesParamsQueryTypeIxfr, RadarDNSGetTimeseriesParamsQueryTypeKey, RadarDNSGetTimeseriesParamsQueryTypeKx, RadarDNSGetTimeseriesParamsQueryTypeL32, RadarDNSGetTimeseriesParamsQueryTypeL64, RadarDNSGetTimeseriesParamsQueryTypeLoc, RadarDNSGetTimeseriesParamsQueryTypeLp, RadarDNSGetTimeseriesParamsQueryTypeMaila, RadarDNSGetTimeseriesParamsQueryTypeMailb, RadarDNSGetTimeseriesParamsQueryTypeMB, RadarDNSGetTimeseriesParamsQueryTypeMd, RadarDNSGetTimeseriesParamsQueryTypeMf, RadarDNSGetTimeseriesParamsQueryTypeMg, RadarDNSGetTimeseriesParamsQueryTypeMinfo, RadarDNSGetTimeseriesParamsQueryTypeMr, RadarDNSGetTimeseriesParamsQueryTypeMx, RadarDNSGetTimeseriesParamsQueryTypeNaptr, RadarDNSGetTimeseriesParamsQueryTypeNb, RadarDNSGetTimeseriesParamsQueryTypeNbstat, RadarDNSGetTimeseriesParamsQueryTypeNid, RadarDNSGetTimeseriesParamsQueryTypeNimloc, RadarDNSGetTimeseriesParamsQueryTypeNinfo, RadarDNSGetTimeseriesParamsQueryTypeNs, RadarDNSGetTimeseriesParamsQueryTypeNsap, RadarDNSGetTimeseriesParamsQueryTypeNsec, RadarDNSGetTimeseriesParamsQueryTypeNsec3, RadarDNSGetTimeseriesParamsQueryTypeNsec3Param, RadarDNSGetTimeseriesParamsQueryTypeNull, RadarDNSGetTimeseriesParamsQueryTypeNxt, RadarDNSGetTimeseriesParamsQueryTypeOpenpgpkey, RadarDNSGetTimeseriesParamsQueryTypeOpt, RadarDNSGetTimeseriesParamsQueryTypePtr, RadarDNSGetTimeseriesParamsQueryTypePx, RadarDNSGetTimeseriesParamsQueryTypeRkey, RadarDNSGetTimeseriesParamsQueryTypeRp, RadarDNSGetTimeseriesParamsQueryTypeRrsig, RadarDNSGetTimeseriesParamsQueryTypeRt, RadarDNSGetTimeseriesParamsQueryTypeSig, RadarDNSGetTimeseriesParamsQueryTypeSink, RadarDNSGetTimeseriesParamsQueryTypeSmimea, RadarDNSGetTimeseriesParamsQueryTypeSoa, RadarDNSGetTimeseriesParamsQueryTypeSpf, RadarDNSGetTimeseriesParamsQueryTypeSrv, RadarDNSGetTimeseriesParamsQueryTypeSshfp, RadarDNSGetTimeseriesParamsQueryTypeSvcb, RadarDNSGetTimeseriesParamsQueryTypeTa, RadarDNSGetTimeseriesParamsQueryTypeTalink, RadarDNSGetTimeseriesParamsQueryTypeTkey, RadarDNSGetTimeseriesParamsQueryTypeTlsa, RadarDNSGetTimeseriesParamsQueryTypeTsig, RadarDNSGetTimeseriesParamsQueryTypeTxt, RadarDNSGetTimeseriesParamsQueryTypeUinfo, RadarDNSGetTimeseriesParamsQueryTypeUid, RadarDNSGetTimeseriesParamsQueryTypeUnspec, RadarDNSGetTimeseriesParamsQueryTypeUri, RadarDNSGetTimeseriesParamsQueryTypeWks, RadarDNSGetTimeseriesParamsQueryTypeX25, RadarDNSGetTimeseriesParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSGetTimeseriesParamsResponseCode string

const (
	RadarDNSGetTimeseriesParamsResponseCodeNoerror   RadarDNSGetTimeseriesParamsResponseCode = "NOERROR"
	RadarDNSGetTimeseriesParamsResponseCodeFormerr   RadarDNSGetTimeseriesParamsResponseCode = "FORMERR"
	RadarDNSGetTimeseriesParamsResponseCodeServfail  RadarDNSGetTimeseriesParamsResponseCode = "SERVFAIL"
	RadarDNSGetTimeseriesParamsResponseCodeNxdomain  RadarDNSGetTimeseriesParamsResponseCode = "NXDOMAIN"
	RadarDNSGetTimeseriesParamsResponseCodeNotimp    RadarDNSGetTimeseriesParamsResponseCode = "NOTIMP"
	RadarDNSGetTimeseriesParamsResponseCodeRefused   RadarDNSGetTimeseriesParamsResponseCode = "REFUSED"
	RadarDNSGetTimeseriesParamsResponseCodeYxdomain  RadarDNSGetTimeseriesParamsResponseCode = "YXDOMAIN"
	RadarDNSGetTimeseriesParamsResponseCodeYxrrset   RadarDNSGetTimeseriesParamsResponseCode = "YXRRSET"
	RadarDNSGetTimeseriesParamsResponseCodeNxrrset   RadarDNSGetTimeseriesParamsResponseCode = "NXRRSET"
	RadarDNSGetTimeseriesParamsResponseCodeNotauth   RadarDNSGetTimeseriesParamsResponseCode = "NOTAUTH"
	RadarDNSGetTimeseriesParamsResponseCodeNotzone   RadarDNSGetTimeseriesParamsResponseCode = "NOTZONE"
	RadarDNSGetTimeseriesParamsResponseCodeBadsig    RadarDNSGetTimeseriesParamsResponseCode = "BADSIG"
	RadarDNSGetTimeseriesParamsResponseCodeBadkey    RadarDNSGetTimeseriesParamsResponseCode = "BADKEY"
	RadarDNSGetTimeseriesParamsResponseCodeBadtime   RadarDNSGetTimeseriesParamsResponseCode = "BADTIME"
	RadarDNSGetTimeseriesParamsResponseCodeBadmode   RadarDNSGetTimeseriesParamsResponseCode = "BADMODE"
	RadarDNSGetTimeseriesParamsResponseCodeBadname   RadarDNSGetTimeseriesParamsResponseCode = "BADNAME"
	RadarDNSGetTimeseriesParamsResponseCodeBadalg    RadarDNSGetTimeseriesParamsResponseCode = "BADALG"
	RadarDNSGetTimeseriesParamsResponseCodeBadtrunc  RadarDNSGetTimeseriesParamsResponseCode = "BADTRUNC"
	RadarDNSGetTimeseriesParamsResponseCodeBadcookie RadarDNSGetTimeseriesParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSGetTimeseriesParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSGetTimeseriesParamsResponseCodeNoerror, RadarDNSGetTimeseriesParamsResponseCodeFormerr, RadarDNSGetTimeseriesParamsResponseCodeServfail, RadarDNSGetTimeseriesParamsResponseCodeNxdomain, RadarDNSGetTimeseriesParamsResponseCodeNotimp, RadarDNSGetTimeseriesParamsResponseCodeRefused, RadarDNSGetTimeseriesParamsResponseCodeYxdomain, RadarDNSGetTimeseriesParamsResponseCodeYxrrset, RadarDNSGetTimeseriesParamsResponseCodeNxrrset, RadarDNSGetTimeseriesParamsResponseCodeNotauth, RadarDNSGetTimeseriesParamsResponseCodeNotzone, RadarDNSGetTimeseriesParamsResponseCodeBadsig, RadarDNSGetTimeseriesParamsResponseCodeBadkey, RadarDNSGetTimeseriesParamsResponseCodeBadtime, RadarDNSGetTimeseriesParamsResponseCodeBadmode, RadarDNSGetTimeseriesParamsResponseCodeBadname, RadarDNSGetTimeseriesParamsResponseCodeBadalg, RadarDNSGetTimeseriesParamsResponseCodeBadtrunc, RadarDNSGetTimeseriesParamsResponseCodeBadcookie:
		return true
	}
	return false
}
