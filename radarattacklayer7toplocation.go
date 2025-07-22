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

// RadarAttackLayer7TopLocationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7TopLocationService] method instead.
type RadarAttackLayer7TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7TopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopLocationService(opts ...option.RequestOption) (r *RadarAttackLayer7TopLocationService) {
	r = &RadarAttackLayer7TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the top origin locations of layer 7 attacks. Values are percentages of
// the total layer 7 attacks, with the origin location determined by the client IP
// address.
func (r *RadarAttackLayer7TopLocationService) GetTopOriginLocations(ctx context.Context, query RadarAttackLayer7TopLocationGetTopOriginLocationsParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationGetTopOriginLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top target locations of and by layer 7 attacks. Values are a
// percentage out of the total layer 7 attacks. The target location is determined
// by the attacked zone's billing country, when available.
func (r *RadarAttackLayer7TopLocationService) GetTopTargetLocations(ctx context.Context, query RadarAttackLayer7TopLocationGetTopTargetLocationsParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopLocationGetTopTargetLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponse struct {
	Result  RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResult `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarAttackLayer7TopLocationGetTopOriginLocationsResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponse]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResult struct {
	Meta RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResult]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMeta]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRange]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                         `json:"level"`
	JSON        radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                          `json:"dataSource,required"`
	Description     string                                                                                          `json:"description,required"`
	EventType       string                                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                                  `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                                  `json:"originCountryName,required"`
	Rank                float64                                                                 `json:"rank,required"`
	Value               string                                                                  `json:"value,required"`
	JSON                radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0]
type radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopOriginLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponse struct {
	Result  RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResult `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarAttackLayer7TopLocationGetTopTargetLocationsResponseJSON   `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponse]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResult struct {
	Meta RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResult]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMeta]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRange]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                         `json:"level"`
	JSON        radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                          `json:"dataSource,required"`
	Description     string                                                                                          `json:"description,required"`
	EventType       string                                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0 struct {
	Rank                float64                                                                 `json:"rank,required"`
	TargetCountryAlpha2 string                                                                  `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                                  `json:"targetCountryName,required"`
	Value               string                                                                  `json:"value,required"`
	JSON                radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0]
type radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopLocationGetTopTargetLocationsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsParams struct {
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
	Format param.Field[RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationGetTopOriginLocationsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormat string

const (
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormatJson RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormat = "JSON"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormatCsv  RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormatJson, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod string

const (
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodGet             RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "GET"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPost            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "POST"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodDelete          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPut             RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "PUT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodHead            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPurge           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodOptions         RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPropfind        RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkcol           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPatch           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodACL             RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "ACL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBcopy           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBdelete         RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBmove           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBpropfind       RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBproppatch      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCheckin         RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCheckout        RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodConnect         RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCopy            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "COPY"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodLabel           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodLock            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMerge           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkactivity      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkworkspace     RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMove            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodNotify          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodOrderpatch      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPoll            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "POLL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodProppatch       RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodReport          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodSearch          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodSubscribe       RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodTrace           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUncheckout      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUnlock          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUnsubscribe     RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUpdate          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodVersioncontrol  RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBaselinecontrol RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodXmsenumatts     RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodRpcOutData      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodRpcInData       RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodJson            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "JSON"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCook            RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "COOK"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodTrack           RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodGet, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPost, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodDelete, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPut, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodHead, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPurge, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodOptions, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPropfind, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkcol, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPatch, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodACL, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBcopy, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBdelete, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBmove, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBpropfind, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBproppatch, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCheckin, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCheckout, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodConnect, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCopy, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodLabel, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodLock, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMerge, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkactivity, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMkworkspace, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodMove, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodNotify, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodOrderpatch, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodPoll, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodProppatch, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodReport, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodSearch, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodSubscribe, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodTrace, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUncheckout, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUnlock, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUnsubscribe, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodUpdate, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodVersioncontrol, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodXmsenumatts, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodRpcOutData, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodRpcInData, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodJson, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodCook, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion string

const (
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv1 RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv2 RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv3 RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv1, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv2, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion string

const (
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersionIPv4 RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion = "IPv4"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersionIPv6 RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersionIPv4, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct string

const (
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductDdos               RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductWaf                RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductBotManagement      RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductAccessRules        RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductIPReputation       RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductAPIShield          RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductDataLossPrevention RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductDdos, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductWaf, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductBotManagement, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductAccessRules, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductIPReputation, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductAPIShield, RadarAttackLayer7TopLocationGetTopOriginLocationsParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsParams struct {
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
	Format param.Field[RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopLocationGetTopTargetLocationsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7TopLocationGetTopTargetLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormat string

const (
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormatJson RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormat = "JSON"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormatCsv  RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormatJson, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct string

const (
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductDdos               RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductWaf                RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductBotManagement      RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductAccessRules        RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductIPReputation       RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductAPIShield          RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductDataLossPrevention RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductDdos, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductWaf, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductBotManagement, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductAccessRules, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductIPReputation, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductAPIShield, RadarAttackLayer7TopLocationGetTopTargetLocationsParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}
