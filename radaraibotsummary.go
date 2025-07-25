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

// RadarAIBotSummaryService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAIBotSummaryService] method instead.
type RadarAIBotSummaryService struct {
	Options []option.RequestOption
}

// NewRadarAIBotSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAIBotSummaryService(opts ...option.RequestOption) (r *RadarAIBotSummaryService) {
	r = &RadarAIBotSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of traffic by AI user agent.
func (r *RadarAIBotSummaryService) GetUserAgent(ctx context.Context, query RadarAIBotSummaryGetUserAgentParams, opts ...option.RequestOption) (res *RadarAIBotSummaryGetUserAgentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/summary/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAIBotSummaryGetUserAgentResponse struct {
	Result  RadarAIBotSummaryGetUserAgentResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAIBotSummaryGetUserAgentResponseJSON   `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseJSON contains the JSON metadata for the
// struct [RadarAIBotSummaryGetUserAgentResponse]
type radarAIBotSummaryGetUserAgentResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotSummaryGetUserAgentResponseResult struct {
	// Metadata for the results.
	Meta     RadarAIBotSummaryGetUserAgentResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                               `json:"summary_0,required"`
	JSON     radarAIBotSummaryGetUserAgentResponseResultJSON `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultJSON contains the JSON metadata for
// the struct [RadarAIBotSummaryGetUserAgentResponseResult]
type radarAIBotSummaryGetUserAgentResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAIBotSummaryGetUserAgentResponseResultMeta struct {
	ConfidenceInfo RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAIBotSummaryGetUserAgentResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAIBotSummaryGetUserAgentResponseResultMetaUnit `json:"units,required"`
	JSON  radarAIBotSummaryGetUserAgentResponseResultMetaJSON   `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAIBotSummaryGetUserAgentResponseResultMeta]
type radarAIBotSummaryGetUserAgentResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                             `json:"level,required"`
	JSON  radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfo]
type radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                   `json:"startDate,required" format:"date-time"`
	JSON            radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotation]
type radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotSummaryGetUserAgentResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAIBotSummaryGetUserAgentResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAIBotSummaryGetUserAgentResponseResultMetaDateRange]
type radarAIBotSummaryGetUserAgentResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization string

const (
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationPercentage           RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "PERCENTAGE"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationMin0Max              RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "MIN0_MAX"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationMinMax               RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "MIN_MAX"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRawValues            RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "RAW_VALUES"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationPercentageChange     RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRollingAverage       RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationOverlappedPercentage RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRatio                RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization = "RATIO"
)

func (r RadarAIBotSummaryGetUserAgentResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationPercentage, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationMin0Max, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationMinMax, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRawValues, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationPercentageChange, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRollingAverage, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationOverlappedPercentage, RadarAIBotSummaryGetUserAgentResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAIBotSummaryGetUserAgentResponseResultMetaUnit struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarAIBotSummaryGetUserAgentResponseResultMetaUnitJSON `json:"-"`
}

// radarAIBotSummaryGetUserAgentResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarAIBotSummaryGetUserAgentResponseResultMetaUnit]
type radarAIBotSummaryGetUserAgentResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAIBotSummaryGetUserAgentResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAIBotSummaryGetUserAgentResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAIBotSummaryGetUserAgentParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAIBotSummaryGetUserAgentParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAIBotSummaryGetUserAgentParams]'s query parameters as
// `url.Values`.
func (r RadarAIBotSummaryGetUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAIBotSummaryGetUserAgentParamsFormat string

const (
	RadarAIBotSummaryGetUserAgentParamsFormatJson RadarAIBotSummaryGetUserAgentParamsFormat = "JSON"
	RadarAIBotSummaryGetUserAgentParamsFormatCsv  RadarAIBotSummaryGetUserAgentParamsFormat = "CSV"
)

func (r RadarAIBotSummaryGetUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case RadarAIBotSummaryGetUserAgentParamsFormatJson, RadarAIBotSummaryGetUserAgentParamsFormatCsv:
		return true
	}
	return false
}
