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

// RadarAttackLayer7TopAseService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7TopAseService] method instead.
type RadarAttackLayer7TopAseService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopAseService(opts ...option.RequestOption) (r *RadarAttackLayer7TopAseService) {
	r = &RadarAttackLayer7TopAseService{}
	r.Options = opts
	return
}

// Retrieves the top origin autonomous systems of layer 7 attacks. Values are
// percentages of the total layer 7 attacks, with the origin autonomous systems
// determined by the client IP address.
func (r *RadarAttackLayer7TopAseService) GetTopOriginAs(ctx context.Context, query RadarAttackLayer7TopAseGetTopOriginAsParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopAseGetTopOriginAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopAseGetTopOriginAsResponse struct {
	Result  RadarAttackLayer7TopAseGetTopOriginAsResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAttackLayer7TopAseGetTopOriginAsResponseJSON   `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopAseGetTopOriginAsResponse]
type radarAttackLayer7TopAseGetTopOriginAsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResult struct {
	Meta RadarAttackLayer7TopAseGetTopOriginAsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopAseGetTopOriginAsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopAseGetTopOriginAsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopAseGetTopOriginAsResponseResult]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopAseGetTopOriginAsResponseResultMeta]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRange]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                             `json:"level"`
	JSON        radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                              `json:"dataSource,required"`
	Description     string                                                                              `json:"description,required"`
	EventType       string                                                                              `json:"eventType,required"`
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsResponseResultTop0 struct {
	OriginAsn     string                                                      `json:"originAsn,required"`
	OriginAsnName string                                                      `json:"originAsnName,required"`
	Rank          float64                                                     `json:"rank,required"`
	Value         string                                                      `json:"value,required"`
	JSON          radarAttackLayer7TopAseGetTopOriginAsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopAseGetTopOriginAsResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopAseGetTopOriginAsResponseResultTop0]
type radarAttackLayer7TopAseGetTopOriginAsResponseResultTop0JSON struct {
	OriginAsn     apijson.Field
	OriginAsnName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarAttackLayer7TopAseGetTopOriginAsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopAseGetTopOriginAsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopAseGetTopOriginAsParams struct {
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
	Format param.Field[RadarAttackLayer7TopAseGetTopOriginAsParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopAseGetTopOriginAsParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopAseGetTopOriginAsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopAseGetTopOriginAsParamsFormat string

const (
	RadarAttackLayer7TopAseGetTopOriginAsParamsFormatJson RadarAttackLayer7TopAseGetTopOriginAsParamsFormat = "JSON"
	RadarAttackLayer7TopAseGetTopOriginAsParamsFormatCsv  RadarAttackLayer7TopAseGetTopOriginAsParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopAseGetTopOriginAsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopAseGetTopOriginAsParamsFormatJson, RadarAttackLayer7TopAseGetTopOriginAsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod string

const (
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodGet             RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "GET"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPost            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "POST"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodDelete          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPut             RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "PUT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodHead            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPurge           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodOptions         RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPropfind        RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkcol           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPatch           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodACL             RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "ACL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBcopy           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBdelete         RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBmove           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBpropfind       RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBproppatch      RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCheckin         RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCheckout        RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodConnect         RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCopy            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "COPY"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodLabel           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodLock            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMerge           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkactivity      RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkworkspace     RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMove            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodNotify          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodOrderpatch      RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPoll            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "POLL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodProppatch       RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodReport          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodSearch          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodSubscribe       RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodTrace           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUncheckout      RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUnlock          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUnsubscribe     RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUpdate          RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodVersioncontrol  RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBaselinecontrol RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodXmsenumatts     RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodRpcOutData      RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodRpcInData       RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodJson            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "JSON"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCook            RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "COOK"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodTrack           RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodGet, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPost, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodDelete, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPut, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodHead, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPurge, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodOptions, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPropfind, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkcol, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPatch, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodACL, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBcopy, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBdelete, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBmove, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBpropfind, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBproppatch, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCheckin, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCheckout, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodConnect, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCopy, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodLabel, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodLock, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMerge, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkactivity, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMkworkspace, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodMove, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodNotify, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodOrderpatch, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodPoll, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodProppatch, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodReport, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodSearch, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodSubscribe, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodTrace, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUncheckout, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUnlock, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUnsubscribe, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodUpdate, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodVersioncontrol, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodXmsenumatts, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodRpcOutData, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodRpcInData, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodJson, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodCook, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion string

const (
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv1 RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv2 RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv3 RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv1, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv2, RadarAttackLayer7TopAseGetTopOriginAsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersion string

const (
	RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersionIPv4 RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersion = "IPv4"
	RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersionIPv6 RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersionIPv4, RadarAttackLayer7TopAseGetTopOriginAsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct string

const (
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductDdos               RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductWaf                RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductBotManagement      RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductAccessRules        RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductIPReputation       RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductAPIShield          RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductDataLossPrevention RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductDdos, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductWaf, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductBotManagement, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductAccessRules, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductIPReputation, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductAPIShield, RadarAttackLayer7TopAseGetTopOriginAsParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}
