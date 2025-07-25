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

// RadarHTTPService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPService] method instead.
type RadarHTTPService struct {
	Options          []option.RequestOption
	Summary          *RadarHTTPSummaryService
	TimeseriesGroups *RadarHTTPTimeseriesGroupService
	Top              *RadarHTTPTopService
}

// NewRadarHTTPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPService(opts ...option.RequestOption) (r *RadarHTTPService) {
	r = &RadarHTTPService{}
	r.Options = opts
	r.Summary = NewRadarHTTPSummaryService(opts...)
	r.TimeseriesGroups = NewRadarHTTPTimeseriesGroupService(opts...)
	r.Top = NewRadarHTTPTopService(opts...)
	return
}

// Retrieves the HTTP requests over time.
func (r *RadarHTTPService) GetTimeseries(ctx context.Context, query RadarHTTPGetTimeseriesParams, opts ...option.RequestOption) (res *RadarHTTPGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPGetTimeseriesResponse struct {
	Result  RadarHTTPGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarHTTPGetTimeseriesResponseJSON   `json:"-"`
}

// radarHTTPGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarHTTPGetTimeseriesResponse]
type radarHTTPGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPGetTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta        RadarHTTPGetTimeseriesResponseResultMeta        `json:"meta,required"`
	ExtraFields map[string]RadarHTTPGetTimeseriesResponseResult `json:"-,extras"`
	JSON        radarHTTPGetTimeseriesResponseResultJSON        `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPGetTimeseriesResponseResult]
type radarHTTPGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPGetTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarHTTPGetTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarHTTPGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPGetTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPGetTimeseriesResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPGetTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPGetTimeseriesResponseResultMeta]
type radarHTTPGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPGetTimeseriesResponseResultMetaAggInterval string

const (
	RadarHTTPGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarHTTPGetTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneHour        RadarHTTPGetTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneDay         RadarHTTPGetTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneWeek        RadarHTTPGetTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneMonth       RadarHTTPGetTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarHTTPGetTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneHour, RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneDay, RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneWeek, RadarHTTPGetTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                      `json:"level,required"`
	JSON  radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfo]
type radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPGetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPGetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPGetTimeseriesResponseResultMetaDateRange]
type radarHTTPGetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPGetTimeseriesResponseResultMetaNormalization string

const (
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationPercentage           RadarHTTPGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationMin0Max              RadarHTTPGetTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationMinMax               RadarHTTPGetTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationRawValues            RadarHTTPGetTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationPercentageChange     RadarHTTPGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationRollingAverage       RadarHTTPGetTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarHTTPGetTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPGetTimeseriesResponseResultMetaNormalizationRatio                RadarHTTPGetTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPGetTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesResponseResultMetaNormalizationPercentage, RadarHTTPGetTimeseriesResponseResultMetaNormalizationMin0Max, RadarHTTPGetTimeseriesResponseResultMetaNormalizationMinMax, RadarHTTPGetTimeseriesResponseResultMetaNormalizationRawValues, RadarHTTPGetTimeseriesResponseResultMetaNormalizationPercentageChange, RadarHTTPGetTimeseriesResponseResultMetaNormalizationRollingAverage, RadarHTTPGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPGetTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesResponseResultMetaUnit struct {
	Name  string                                           `json:"name,required"`
	Value string                                           `json:"value,required"`
	JSON  radarHTTPGetTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPGetTimeseriesResponseResultMetaUnitJSON contains the JSON metadata for
// the struct [RadarHTTPGetTimeseriesResponseResultMetaUnit]
type radarHTTPGetTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPGetTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPGetTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPGetTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPGetTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPGetTimeseriesParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPGetTimeseriesParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPGetTimeseriesParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPGetTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPGetTimeseriesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPGetTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPGetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarHTTPGetTimeseriesParamsNormalization] `query:"normalization"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPGetTimeseriesParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPGetTimeseriesParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPGetTimeseriesParamsAggInterval string

const (
	RadarHTTPGetTimeseriesParamsAggInterval15m RadarHTTPGetTimeseriesParamsAggInterval = "15m"
	RadarHTTPGetTimeseriesParamsAggInterval1h  RadarHTTPGetTimeseriesParamsAggInterval = "1h"
	RadarHTTPGetTimeseriesParamsAggInterval1d  RadarHTTPGetTimeseriesParamsAggInterval = "1d"
	RadarHTTPGetTimeseriesParamsAggInterval1w  RadarHTTPGetTimeseriesParamsAggInterval = "1w"
)

func (r RadarHTTPGetTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsAggInterval15m, RadarHTTPGetTimeseriesParamsAggInterval1h, RadarHTTPGetTimeseriesParamsAggInterval1d, RadarHTTPGetTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsBotClass string

const (
	RadarHTTPGetTimeseriesParamsBotClassLikelyAutomated RadarHTTPGetTimeseriesParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPGetTimeseriesParamsBotClassLikelyHuman     RadarHTTPGetTimeseriesParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPGetTimeseriesParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsBotClassLikelyAutomated, RadarHTTPGetTimeseriesParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsBrowserFamily string

const (
	RadarHTTPGetTimeseriesParamsBrowserFamilyChrome  RadarHTTPGetTimeseriesParamsBrowserFamily = "CHROME"
	RadarHTTPGetTimeseriesParamsBrowserFamilyEdge    RadarHTTPGetTimeseriesParamsBrowserFamily = "EDGE"
	RadarHTTPGetTimeseriesParamsBrowserFamilyFirefox RadarHTTPGetTimeseriesParamsBrowserFamily = "FIREFOX"
	RadarHTTPGetTimeseriesParamsBrowserFamilySafari  RadarHTTPGetTimeseriesParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPGetTimeseriesParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsBrowserFamilyChrome, RadarHTTPGetTimeseriesParamsBrowserFamilyEdge, RadarHTTPGetTimeseriesParamsBrowserFamilyFirefox, RadarHTTPGetTimeseriesParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsDeviceType string

const (
	RadarHTTPGetTimeseriesParamsDeviceTypeDesktop RadarHTTPGetTimeseriesParamsDeviceType = "DESKTOP"
	RadarHTTPGetTimeseriesParamsDeviceTypeMobile  RadarHTTPGetTimeseriesParamsDeviceType = "MOBILE"
	RadarHTTPGetTimeseriesParamsDeviceTypeOther   RadarHTTPGetTimeseriesParamsDeviceType = "OTHER"
)

func (r RadarHTTPGetTimeseriesParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsDeviceTypeDesktop, RadarHTTPGetTimeseriesParamsDeviceTypeMobile, RadarHTTPGetTimeseriesParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPGetTimeseriesParamsFormat string

const (
	RadarHTTPGetTimeseriesParamsFormatJson RadarHTTPGetTimeseriesParamsFormat = "JSON"
	RadarHTTPGetTimeseriesParamsFormatCsv  RadarHTTPGetTimeseriesParamsFormat = "CSV"
)

func (r RadarHTTPGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsFormatJson, RadarHTTPGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsHTTPProtocol string

const (
	RadarHTTPGetTimeseriesParamsHTTPProtocolHTTP  RadarHTTPGetTimeseriesParamsHTTPProtocol = "HTTP"
	RadarHTTPGetTimeseriesParamsHTTPProtocolHTTPS RadarHTTPGetTimeseriesParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPGetTimeseriesParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsHTTPProtocolHTTP, RadarHTTPGetTimeseriesParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsHTTPVersion string

const (
	RadarHTTPGetTimeseriesParamsHTTPVersionHttPv1 RadarHTTPGetTimeseriesParamsHTTPVersion = "HTTPv1"
	RadarHTTPGetTimeseriesParamsHTTPVersionHttPv2 RadarHTTPGetTimeseriesParamsHTTPVersion = "HTTPv2"
	RadarHTTPGetTimeseriesParamsHTTPVersionHttPv3 RadarHTTPGetTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPGetTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsHTTPVersionHttPv1, RadarHTTPGetTimeseriesParamsHTTPVersionHttPv2, RadarHTTPGetTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsIPVersion string

const (
	RadarHTTPGetTimeseriesParamsIPVersionIPv4 RadarHTTPGetTimeseriesParamsIPVersion = "IPv4"
	RadarHTTPGetTimeseriesParamsIPVersionIPv6 RadarHTTPGetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarHTTPGetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsIPVersionIPv4, RadarHTTPGetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPGetTimeseriesParamsNormalization string

const (
	RadarHTTPGetTimeseriesParamsNormalizationPercentageChange RadarHTTPGetTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPGetTimeseriesParamsNormalizationMin0Max          RadarHTTPGetTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r RadarHTTPGetTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsNormalizationPercentageChange, RadarHTTPGetTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsO string

const (
	RadarHTTPGetTimeseriesParamsOWindows  RadarHTTPGetTimeseriesParamsO = "WINDOWS"
	RadarHTTPGetTimeseriesParamsOMacosx   RadarHTTPGetTimeseriesParamsO = "MACOSX"
	RadarHTTPGetTimeseriesParamsOIos      RadarHTTPGetTimeseriesParamsO = "IOS"
	RadarHTTPGetTimeseriesParamsOAndroid  RadarHTTPGetTimeseriesParamsO = "ANDROID"
	RadarHTTPGetTimeseriesParamsOChromeos RadarHTTPGetTimeseriesParamsO = "CHROMEOS"
	RadarHTTPGetTimeseriesParamsOLinux    RadarHTTPGetTimeseriesParamsO = "LINUX"
	RadarHTTPGetTimeseriesParamsOSmartTv  RadarHTTPGetTimeseriesParamsO = "SMART_TV"
)

func (r RadarHTTPGetTimeseriesParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsOWindows, RadarHTTPGetTimeseriesParamsOMacosx, RadarHTTPGetTimeseriesParamsOIos, RadarHTTPGetTimeseriesParamsOAndroid, RadarHTTPGetTimeseriesParamsOChromeos, RadarHTTPGetTimeseriesParamsOLinux, RadarHTTPGetTimeseriesParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPGetTimeseriesParamsTlsVersion string

const (
	RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_0  RadarHTTPGetTimeseriesParamsTlsVersion = "TLSv1_0"
	RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_1  RadarHTTPGetTimeseriesParamsTlsVersion = "TLSv1_1"
	RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_2  RadarHTTPGetTimeseriesParamsTlsVersion = "TLSv1_2"
	RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_3  RadarHTTPGetTimeseriesParamsTlsVersion = "TLSv1_3"
	RadarHTTPGetTimeseriesParamsTlsVersionTlSvQuic RadarHTTPGetTimeseriesParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPGetTimeseriesParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_0, RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_1, RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_2, RadarHTTPGetTimeseriesParamsTlsVersionTlSv1_3, RadarHTTPGetTimeseriesParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}
