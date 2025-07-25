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

// RadarBgpIPService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpIPService] method instead.
type RadarBgpIPService struct {
	Options []option.RequestOption
}

// NewRadarBgpIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpIPService(opts ...option.RequestOption) (r *RadarBgpIPService) {
	r = &RadarBgpIPService{}
	r.Options = opts
	return
}

// Retrieves time series data for the announced IP space count, represented as the
// number of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *RadarBgpIPService) GetTimeseries(ctx context.Context, query RadarBgpIPGetTimeseriesParams, opts ...option.RequestOption) (res *RadarBgpIPGetTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/ips/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpIPGetTimeseriesResponse struct {
	Result  RadarBgpIPGetTimeseriesResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarBgpIPGetTimeseriesResponseJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarBgpIPGetTimeseriesResponse]
type radarBgpIPGetTimeseriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResult struct {
	// Metadata for the results.
	Meta   RadarBgpIPGetTimeseriesResponseResultMeta   `json:"meta,required"`
	Serie0 RadarBgpIPGetTimeseriesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarBgpIPGetTimeseriesResponseResultJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpIPGetTimeseriesResponseResult]
type radarBgpIPGetTimeseriesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarBgpIPGetTimeseriesResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarBgpIPGetTimeseriesResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarBgpIPGetTimeseriesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarBgpIPGetTimeseriesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarBgpIPGetTimeseriesResponseResultMetaUnit `json:"units,required"`
	Delay RadarBgpIPGetTimeseriesResponseResultMetaDelay  `json:"delay"`
	JSON  radarBgpIPGetTimeseriesResponseResultMetaJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpIPGetTimeseriesResponseResultMeta]
type radarBgpIPGetTimeseriesResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	Delay          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarBgpIPGetTimeseriesResponseResultMetaAggInterval string

const (
	RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes RadarBgpIPGetTimeseriesResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneHour        RadarBgpIPGetTimeseriesResponseResultMetaAggInterval = "ONE_HOUR"
	RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneDay         RadarBgpIPGetTimeseriesResponseResultMetaAggInterval = "ONE_DAY"
	RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneWeek        RadarBgpIPGetTimeseriesResponseResultMetaAggInterval = "ONE_WEEK"
	RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneMonth       RadarBgpIPGetTimeseriesResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarBgpIPGetTimeseriesResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalFifteenMinutes, RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneHour, RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneDay, RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneWeek, RadarBgpIPGetTimeseriesResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfo]
type radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation]
type radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarBgpIPGetTimeseriesResponseResultMetaDateRange]
type radarBgpIPGetTimeseriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarBgpIPGetTimeseriesResponseResultMetaNormalization string

const (
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationPercentage           RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationMin0Max              RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "MIN0_MAX"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationMinMax               RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "MIN_MAX"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRawValues            RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "RAW_VALUES"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationPercentageChange     RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRollingAverage       RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRatio                RadarBgpIPGetTimeseriesResponseResultMetaNormalization = "RATIO"
)

func (r RadarBgpIPGetTimeseriesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesResponseResultMetaNormalizationPercentage, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationMin0Max, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationMinMax, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRawValues, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationPercentageChange, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRollingAverage, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationOverlappedPercentage, RadarBgpIPGetTimeseriesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarBgpIPGetTimeseriesResponseResultMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  radarBgpIPGetTimeseriesResponseResultMetaUnitJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaUnitJSON contains the JSON metadata for
// the struct [RadarBgpIPGetTimeseriesResponseResultMetaUnit]
type radarBgpIPGetTimeseriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDelay struct {
	AsnData     RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnData     `json:"asn_data,required"`
	CountryData RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryData `json:"country_data,required"`
	Healthy     bool                                                      `json:"healthy,required"`
	NowTs       float64                                                   `json:"nowTs,required"`
	JSON        radarBgpIPGetTimeseriesResponseResultMetaDelayJSON        `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDelayJSON contains the JSON metadata
// for the struct [RadarBgpIPGetTimeseriesResponseResultMetaDelay]
type radarBgpIPGetTimeseriesResponseResultMetaDelayJSON struct {
	AsnData     apijson.Field
	CountryData apijson.Field
	Healthy     apijson.Field
	NowTs       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDelay) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDelayJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnData struct {
	DelaySecs float64                                                     `json:"delaySecs,required"`
	DelayStr  string                                                      `json:"delayStr,required"`
	Healthy   bool                                                        `json:"healthy,required"`
	Latest    RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatest `json:"latest,required"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataJSON contains the JSON
// metadata for the struct [RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnData]
type radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataJSON struct {
	DelaySecs   apijson.Field
	DelayStr    apijson.Field
	Healthy     apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatest struct {
	EntriesCount float64                                                         `json:"entries_count,required"`
	Path         string                                                          `json:"path,required"`
	Timestamp    float64                                                         `json:"timestamp,required"`
	JSON         radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatestJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatestJSON contains the
// JSON metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatest]
type radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatestJSON struct {
	EntriesCount apijson.Field
	Path         apijson.Field
	Timestamp    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDelayAsnDataLatestJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryData struct {
	DelaySecs float64                                                         `json:"delaySecs,required"`
	DelayStr  string                                                          `json:"delayStr,required"`
	Healthy   bool                                                            `json:"healthy,required"`
	Latest    RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatest `json:"latest,required"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataJSON   `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataJSON contains the JSON
// metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryData]
type radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataJSON struct {
	DelaySecs   apijson.Field
	DelayStr    apijson.Field
	Healthy     apijson.Field
	Latest      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatest struct {
	Count     float64                                                             `json:"count,required"`
	Timestamp float64                                                             `json:"timestamp,required"`
	JSON      radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatestJSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatestJSON contains the
// JSON metadata for the struct
// [RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatest]
type radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatestJSON struct {
	Count       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultMetaDelayCountryDataLatestJSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesResponseResultSerie0 struct {
	Ipv4       []string                                        `json:"ipv4,required"`
	Ipv6       []string                                        `json:"ipv6,required"`
	Timestamps []time.Time                                     `json:"timestamps,required" format:"date-time"`
	JSON       radarBgpIPGetTimeseriesResponseResultSerie0JSON `json:"-"`
}

// radarBgpIPGetTimeseriesResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarBgpIPGetTimeseriesResponseResultSerie0]
type radarBgpIPGetTimeseriesResponseResultSerie0JSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpIPGetTimeseriesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpIPGetTimeseriesResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarBgpIPGetTimeseriesParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpIPGetTimeseriesParamsFormat] `query:"format"`
	// Includes data delay meta information.
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarBgpIPGetTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarBgpIPGetTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpIPGetTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpIPGetTimeseriesParamsFormat string

const (
	RadarBgpIPGetTimeseriesParamsFormatJson RadarBgpIPGetTimeseriesParamsFormat = "JSON"
	RadarBgpIPGetTimeseriesParamsFormatCsv  RadarBgpIPGetTimeseriesParamsFormat = "CSV"
)

func (r RadarBgpIPGetTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesParamsFormatJson, RadarBgpIPGetTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpIPGetTimeseriesParamsIPVersion string

const (
	RadarBgpIPGetTimeseriesParamsIPVersionIPv4 RadarBgpIPGetTimeseriesParamsIPVersion = "IPv4"
	RadarBgpIPGetTimeseriesParamsIPVersionIPv6 RadarBgpIPGetTimeseriesParamsIPVersion = "IPv6"
)

func (r RadarBgpIPGetTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarBgpIPGetTimeseriesParamsIPVersionIPv4, RadarBgpIPGetTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}
