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

// RadarHTTPTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTopService] method instead.
type RadarHTTPTopService struct {
	Options   []option.RequestOption
	Ases      *RadarHTTPTopAseService
	Locations *RadarHTTPTopLocationService
}

// NewRadarHTTPTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPTopService(opts ...option.RequestOption) (r *RadarHTTPTopService) {
	r = &RadarHTTPTopService{}
	r.Options = opts
	r.Ases = NewRadarHTTPTopAseService(opts...)
	r.Locations = NewRadarHTTPTopLocationService(opts...)
	return
}

// Retrieves the top user agents, aggregated in families, by HTTP requests.
func (r *RadarHTTPTopService) GetTopBrowserFamilies(ctx context.Context, query RadarHTTPTopGetTopBrowserFamiliesParams, opts ...option.RequestOption) (res *RadarHTTPTopGetTopBrowserFamiliesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top user agents by HTTP requests.
func (r *RadarHTTPTopService) GetTopBrowsers(ctx context.Context, query RadarHTTPTopGetTopBrowsersParams, opts ...option.RequestOption) (res *RadarHTTPTopGetTopBrowsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopGetTopBrowserFamiliesResponse struct {
	Result  RadarHTTPTopGetTopBrowserFamiliesResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTopGetTopBrowserFamiliesResponseJSON   `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopGetTopBrowserFamiliesResponse]
type radarHTTPTopGetTopBrowserFamiliesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResult struct {
	Meta RadarHTTPTopGetTopBrowserFamiliesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopGetTopBrowserFamiliesResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopGetTopBrowserFamiliesResponseResultJSON   `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopGetTopBrowserFamiliesResponseResult]
type radarHTTPTopGetTopBrowserFamiliesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResultMeta struct {
	DateRange      []RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopGetTopBrowserFamiliesResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopGetTopBrowserFamiliesResponseResultMeta]
type radarHTTPTopGetTopBrowserFamiliesResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRange]
type radarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfo]
type radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesResponseResultTop0 struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarHTTPTopGetTopBrowserFamiliesResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopGetTopBrowserFamiliesResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopGetTopBrowserFamiliesResponseResultTop0]
type radarHTTPTopGetTopBrowserFamiliesResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowserFamiliesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowserFamiliesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponse struct {
	Result  RadarHTTPTopGetTopBrowsersResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarHTTPTopGetTopBrowsersResponseJSON   `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopGetTopBrowsersResponse]
type radarHTTPTopGetTopBrowsersResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResult struct {
	Meta RadarHTTPTopGetTopBrowsersResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopGetTopBrowsersResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopGetTopBrowsersResponseResultJSON   `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopGetTopBrowsersResponseResult]
type radarHTTPTopGetTopBrowsersResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResultMeta struct {
	DateRange      []RadarHTTPTopGetTopBrowsersResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopGetTopBrowsersResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopGetTopBrowsersResponseResultMeta]
type radarHTTPTopGetTopBrowsersResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopGetTopBrowsersResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopGetTopBrowsersResponseResultMetaDateRange]
type radarHTTPTopGetTopBrowsersResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfo]
type radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowsersResponseResultTop0 struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarHTTPTopGetTopBrowsersResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopGetTopBrowsersResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopGetTopBrowsersResponseResultTop0]
type radarHTTPTopGetTopBrowsersResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopGetTopBrowsersResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopGetTopBrowsersResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopGetTopBrowserFamiliesParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopGetTopBrowserFamiliesParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopGetTopBrowserFamiliesParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopGetTopBrowserFamiliesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopGetTopBrowserFamiliesParamsBotClass string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsBotClassLikelyAutomated RadarHTTPTopGetTopBrowserFamiliesParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopGetTopBrowserFamiliesParamsBotClassLikelyHuman     RadarHTTPTopGetTopBrowserFamiliesParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsBotClassLikelyAutomated, RadarHTTPTopGetTopBrowserFamiliesParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeDesktop RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType = "DESKTOP"
	RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeMobile  RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType = "MOBILE"
	RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeOther   RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeDesktop, RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeMobile, RadarHTTPTopGetTopBrowserFamiliesParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopGetTopBrowserFamiliesParamsFormat string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsFormatJson RadarHTTPTopGetTopBrowserFamiliesParamsFormat = "JSON"
	RadarHTTPTopGetTopBrowserFamiliesParamsFormatCsv  RadarHTTPTopGetTopBrowserFamiliesParamsFormat = "CSV"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsFormatJson, RadarHTTPTopGetTopBrowserFamiliesParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocolHTTP  RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol = "HTTP"
	RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocolHTTPS RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocolHTTP, RadarHTTPTopGetTopBrowserFamiliesParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv1 RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv2 RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv3 RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv1, RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv2, RadarHTTPTopGetTopBrowserFamiliesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsIPVersionIPv4 RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion = "IPv4"
	RadarHTTPTopGetTopBrowserFamiliesParamsIPVersionIPv6 RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsIPVersionIPv4, RadarHTTPTopGetTopBrowserFamiliesParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsO string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsOWindows  RadarHTTPTopGetTopBrowserFamiliesParamsO = "WINDOWS"
	RadarHTTPTopGetTopBrowserFamiliesParamsOMacosx   RadarHTTPTopGetTopBrowserFamiliesParamsO = "MACOSX"
	RadarHTTPTopGetTopBrowserFamiliesParamsOIos      RadarHTTPTopGetTopBrowserFamiliesParamsO = "IOS"
	RadarHTTPTopGetTopBrowserFamiliesParamsOAndroid  RadarHTTPTopGetTopBrowserFamiliesParamsO = "ANDROID"
	RadarHTTPTopGetTopBrowserFamiliesParamsOChromeos RadarHTTPTopGetTopBrowserFamiliesParamsO = "CHROMEOS"
	RadarHTTPTopGetTopBrowserFamiliesParamsOLinux    RadarHTTPTopGetTopBrowserFamiliesParamsO = "LINUX"
	RadarHTTPTopGetTopBrowserFamiliesParamsOSmartTv  RadarHTTPTopGetTopBrowserFamiliesParamsO = "SMART_TV"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsOWindows, RadarHTTPTopGetTopBrowserFamiliesParamsOMacosx, RadarHTTPTopGetTopBrowserFamiliesParamsOIos, RadarHTTPTopGetTopBrowserFamiliesParamsOAndroid, RadarHTTPTopGetTopBrowserFamiliesParamsOChromeos, RadarHTTPTopGetTopBrowserFamiliesParamsOLinux, RadarHTTPTopGetTopBrowserFamiliesParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion string

const (
	RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_0  RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_1  RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_2  RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_3  RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSvQuic RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_0, RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_1, RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_2, RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSv1_3, RadarHTTPTopGetTopBrowserFamiliesParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopGetTopBrowsersParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopGetTopBrowsersParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopGetTopBrowsersParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopGetTopBrowsersParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopGetTopBrowsersParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopGetTopBrowsersParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopGetTopBrowsersParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopGetTopBrowsersParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopGetTopBrowsersParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopGetTopBrowsersParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopGetTopBrowsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopGetTopBrowsersParamsBotClass string

const (
	RadarHTTPTopGetTopBrowsersParamsBotClassLikelyAutomated RadarHTTPTopGetTopBrowsersParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopGetTopBrowsersParamsBotClassLikelyHuman     RadarHTTPTopGetTopBrowsersParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopGetTopBrowsersParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsBotClassLikelyAutomated, RadarHTTPTopGetTopBrowsersParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsBrowserFamily string

const (
	RadarHTTPTopGetTopBrowsersParamsBrowserFamilyChrome  RadarHTTPTopGetTopBrowsersParamsBrowserFamily = "CHROME"
	RadarHTTPTopGetTopBrowsersParamsBrowserFamilyEdge    RadarHTTPTopGetTopBrowsersParamsBrowserFamily = "EDGE"
	RadarHTTPTopGetTopBrowsersParamsBrowserFamilyFirefox RadarHTTPTopGetTopBrowsersParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopGetTopBrowsersParamsBrowserFamilySafari  RadarHTTPTopGetTopBrowsersParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopGetTopBrowsersParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsBrowserFamilyChrome, RadarHTTPTopGetTopBrowsersParamsBrowserFamilyEdge, RadarHTTPTopGetTopBrowsersParamsBrowserFamilyFirefox, RadarHTTPTopGetTopBrowsersParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsDeviceType string

const (
	RadarHTTPTopGetTopBrowsersParamsDeviceTypeDesktop RadarHTTPTopGetTopBrowsersParamsDeviceType = "DESKTOP"
	RadarHTTPTopGetTopBrowsersParamsDeviceTypeMobile  RadarHTTPTopGetTopBrowsersParamsDeviceType = "MOBILE"
	RadarHTTPTopGetTopBrowsersParamsDeviceTypeOther   RadarHTTPTopGetTopBrowsersParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopGetTopBrowsersParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsDeviceTypeDesktop, RadarHTTPTopGetTopBrowsersParamsDeviceTypeMobile, RadarHTTPTopGetTopBrowsersParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopGetTopBrowsersParamsFormat string

const (
	RadarHTTPTopGetTopBrowsersParamsFormatJson RadarHTTPTopGetTopBrowsersParamsFormat = "JSON"
	RadarHTTPTopGetTopBrowsersParamsFormatCsv  RadarHTTPTopGetTopBrowsersParamsFormat = "CSV"
)

func (r RadarHTTPTopGetTopBrowsersParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsFormatJson, RadarHTTPTopGetTopBrowsersParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsHTTPProtocol string

const (
	RadarHTTPTopGetTopBrowsersParamsHTTPProtocolHTTP  RadarHTTPTopGetTopBrowsersParamsHTTPProtocol = "HTTP"
	RadarHTTPTopGetTopBrowsersParamsHTTPProtocolHTTPS RadarHTTPTopGetTopBrowsersParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopGetTopBrowsersParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsHTTPProtocolHTTP, RadarHTTPTopGetTopBrowsersParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsHTTPVersion string

const (
	RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv1 RadarHTTPTopGetTopBrowsersParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv2 RadarHTTPTopGetTopBrowsersParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv3 RadarHTTPTopGetTopBrowsersParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopGetTopBrowsersParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv1, RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv2, RadarHTTPTopGetTopBrowsersParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsIPVersion string

const (
	RadarHTTPTopGetTopBrowsersParamsIPVersionIPv4 RadarHTTPTopGetTopBrowsersParamsIPVersion = "IPv4"
	RadarHTTPTopGetTopBrowsersParamsIPVersionIPv6 RadarHTTPTopGetTopBrowsersParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopGetTopBrowsersParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsIPVersionIPv4, RadarHTTPTopGetTopBrowsersParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsO string

const (
	RadarHTTPTopGetTopBrowsersParamsOWindows  RadarHTTPTopGetTopBrowsersParamsO = "WINDOWS"
	RadarHTTPTopGetTopBrowsersParamsOMacosx   RadarHTTPTopGetTopBrowsersParamsO = "MACOSX"
	RadarHTTPTopGetTopBrowsersParamsOIos      RadarHTTPTopGetTopBrowsersParamsO = "IOS"
	RadarHTTPTopGetTopBrowsersParamsOAndroid  RadarHTTPTopGetTopBrowsersParamsO = "ANDROID"
	RadarHTTPTopGetTopBrowsersParamsOChromeos RadarHTTPTopGetTopBrowsersParamsO = "CHROMEOS"
	RadarHTTPTopGetTopBrowsersParamsOLinux    RadarHTTPTopGetTopBrowsersParamsO = "LINUX"
	RadarHTTPTopGetTopBrowsersParamsOSmartTv  RadarHTTPTopGetTopBrowsersParamsO = "SMART_TV"
)

func (r RadarHTTPTopGetTopBrowsersParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsOWindows, RadarHTTPTopGetTopBrowsersParamsOMacosx, RadarHTTPTopGetTopBrowsersParamsOIos, RadarHTTPTopGetTopBrowsersParamsOAndroid, RadarHTTPTopGetTopBrowsersParamsOChromeos, RadarHTTPTopGetTopBrowsersParamsOLinux, RadarHTTPTopGetTopBrowsersParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopGetTopBrowsersParamsTlsVersion string

const (
	RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_0  RadarHTTPTopGetTopBrowsersParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_1  RadarHTTPTopGetTopBrowsersParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_2  RadarHTTPTopGetTopBrowsersParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_3  RadarHTTPTopGetTopBrowsersParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSvQuic RadarHTTPTopGetTopBrowsersParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopGetTopBrowsersParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_0, RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_1, RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_2, RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSv1_3, RadarHTTPTopGetTopBrowsersParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}
