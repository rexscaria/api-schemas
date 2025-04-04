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

// RadarAttackLayer7Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7Service] method instead.
type RadarAttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer7SummaryService
	TimeseriesGroups *RadarAttackLayer7TimeseriesGroupService
	Top              *RadarAttackLayer7TopService
}

// NewRadarAttackLayer7Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7Service(opts ...option.RequestOption) (r *RadarAttackLayer7Service) {
	r = &RadarAttackLayer7Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer7TopService(opts...)
	return
}

// Retrieves layer 7 attacks over time.
func (r *RadarAttackLayer7Service) GetTimeseries(ctx context.Context, query RadarAttackLayer7GetTimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7GetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7GetTimeseriesResponse struct {
	Result  RadarAttackLayer7GetTimeseriesResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAttackLayer7GetTimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7GetTimeseriesResponse]
type radarAttackLayer7GetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResult struct {
	Meta   RadarAttackLayer7GetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer7GetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7GetTimeseriesResponseResultJSON   `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer7GetTimeseriesResponseResult]
type radarAttackLayer7GetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultMeta struct {
	AggInterval    string                                                         `json:"aggInterval,required"`
	DateRange      []RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                                      `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7GetTimeseriesResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7GetTimeseriesResponseResultMeta]
type radarAttackLayer7GetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange]
type radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo]
type radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesResponseResultSerie0 struct {
	Timestamps []time.Time                                            `json:"timestamps,required" format:"date-time"`
	Values     []string                                               `json:"values,required"`
	JSON       radarAttackLayer7GetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7GetTimeseriesResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7GetTimeseriesResponseResultSerie0]
type radarAttackLayer7GetTimeseriesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7GetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7GetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7GetTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7GetTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[RadarAttackLayer7GetTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7GetTimeseriesParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7GetTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7GetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7GetTimeseriesParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7GetTimeseriesParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7GetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7GetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7GetTimeseriesParamsAggInterval string

const (
	RadarAttackLayer7GetTimeseriesParamsAggInterval15m RadarAttackLayer7GetTimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1h  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1d  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7GetTimeseriesParamsAggInterval1w  RadarAttackLayer7GetTimeseriesParamsAggInterval = "1w"
)

func (r RadarAttackLayer7GetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsAggInterval15m, RadarAttackLayer7GetTimeseriesParamsAggInterval1h, RadarAttackLayer7GetTimeseriesParamsAggInterval1d, RadarAttackLayer7GetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer7GetTimeseriesParamsFormat string

const (
	RadarAttackLayer7GetTimeseriesParamsFormatJson RadarAttackLayer7GetTimeseriesParamsFormat = "JSON"
	RadarAttackLayer7GetTimeseriesParamsFormatCsv  RadarAttackLayer7GetTimeseriesParamsFormat = "CSV"
)

func (r RadarAttackLayer7GetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsFormatJson, RadarAttackLayer7GetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsHTTPMethod string

const (
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodGet             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "GET"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPost            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "POST"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodDelete          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "DELETE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPut             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodHead            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "HEAD"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPurge           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PURGE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodOptions         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPropfind        RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkcol           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPatch           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodACL             RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "ACL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBcopy           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBdelete         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBmove           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBpropfind       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBproppatch      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckin         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckout        RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodConnect         RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCopy            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "COPY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodLabel           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "LABEL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodLock            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "LOCK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMerge           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MERGE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkactivity      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkworkspace     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodMove            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "MOVE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodNotify          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodOrderpatch      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodPoll            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "POLL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodProppatch       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodReport          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "REPORT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodSearch          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodSubscribe       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrace           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "TRACE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUncheckout      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnlock          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnsubscribe     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodUpdate          RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodVersioncontrol  RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodBaselinecontrol RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodXmsenumatts     RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcOutData      RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcInData       RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodJson            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "JSON"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodCook            RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "COOK"
	RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrack           RadarAttackLayer7GetTimeseriesParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7GetTimeseriesParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsHTTPMethodGet, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPost, RadarAttackLayer7GetTimeseriesParamsHTTPMethodDelete, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPut, RadarAttackLayer7GetTimeseriesParamsHTTPMethodHead, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPurge, RadarAttackLayer7GetTimeseriesParamsHTTPMethodOptions, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPropfind, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkcol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodACL, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBcopy, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBdelete, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBmove, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBpropfind, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBproppatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckin, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCheckout, RadarAttackLayer7GetTimeseriesParamsHTTPMethodConnect, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCopy, RadarAttackLayer7GetTimeseriesParamsHTTPMethodLabel, RadarAttackLayer7GetTimeseriesParamsHTTPMethodLock, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMerge, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkactivity, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMkworkspace, RadarAttackLayer7GetTimeseriesParamsHTTPMethodMove, RadarAttackLayer7GetTimeseriesParamsHTTPMethodNotify, RadarAttackLayer7GetTimeseriesParamsHTTPMethodOrderpatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodPoll, RadarAttackLayer7GetTimeseriesParamsHTTPMethodProppatch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodReport, RadarAttackLayer7GetTimeseriesParamsHTTPMethodSearch, RadarAttackLayer7GetTimeseriesParamsHTTPMethodSubscribe, RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrace, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUncheckout, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnlock, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUnsubscribe, RadarAttackLayer7GetTimeseriesParamsHTTPMethodUpdate, RadarAttackLayer7GetTimeseriesParamsHTTPMethodVersioncontrol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodBaselinecontrol, RadarAttackLayer7GetTimeseriesParamsHTTPMethodXmsenumatts, RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcOutData, RadarAttackLayer7GetTimeseriesParamsHTTPMethodRpcInData, RadarAttackLayer7GetTimeseriesParamsHTTPMethodJson, RadarAttackLayer7GetTimeseriesParamsHTTPMethodCook, RadarAttackLayer7GetTimeseriesParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsHTTPVersion string

const (
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv1 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv2 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv3 RadarAttackLayer7GetTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7GetTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv1, RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv2, RadarAttackLayer7GetTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsIPVersion string

const (
	RadarAttackLayer7GetTimeseriesParamsIPVersionIPv4 RadarAttackLayer7GetTimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7GetTimeseriesParamsIPVersionIPv6 RadarAttackLayer7GetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7GetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsIPVersionIPv4, RadarAttackLayer7GetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7GetTimeseriesParamsMitigationProduct string

const (
	RadarAttackLayer7GetTimeseriesParamsMitigationProductDdos               RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "DDOS"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductWaf                RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "WAF"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductBotManagement      RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductAccessRules        RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductIPReputation       RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductAPIShield          RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7GetTimeseriesParamsMitigationProductDataLossPrevention RadarAttackLayer7GetTimeseriesParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7GetTimeseriesParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsMitigationProductDdos, RadarAttackLayer7GetTimeseriesParamsMitigationProductWaf, RadarAttackLayer7GetTimeseriesParamsMitigationProductBotManagement, RadarAttackLayer7GetTimeseriesParamsMitigationProductAccessRules, RadarAttackLayer7GetTimeseriesParamsMitigationProductIPReputation, RadarAttackLayer7GetTimeseriesParamsMitigationProductAPIShield, RadarAttackLayer7GetTimeseriesParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7GetTimeseriesParamsNormalization string

const (
	RadarAttackLayer7GetTimeseriesParamsNormalizationPercentageChange RadarAttackLayer7GetTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7GetTimeseriesParamsNormalizationMin0Max          RadarAttackLayer7GetTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarAttackLayer7GetTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7GetTimeseriesParamsNormalizationPercentageChange, RadarAttackLayer7GetTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}
