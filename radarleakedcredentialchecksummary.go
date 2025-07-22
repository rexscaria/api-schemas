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

// RadarLeakedCredentialCheckSummaryService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarLeakedCredentialCheckSummaryService] method instead.
type RadarLeakedCredentialCheckSummaryService struct {
	Options []option.RequestOption
}

// NewRadarLeakedCredentialCheckSummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarLeakedCredentialCheckSummaryService(opts ...option.RequestOption) (r *RadarLeakedCredentialCheckSummaryService) {
	r = &RadarLeakedCredentialCheckSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP authentication requests by bot class.
func (r *RadarLeakedCredentialCheckSummaryService) GetByBotClass(ctx context.Context, query RadarLeakedCredentialCheckSummaryGetByBotClassParams, opts ...option.RequestOption) (res *RadarLeakedCredentialCheckSummaryGetByBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/summary/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP authentication requests by compromised
// credential status.
func (r *RadarLeakedCredentialCheckSummaryService) GetByCompromisedStatus(ctx context.Context, query RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParams, opts ...option.RequestOption) (res *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/summary/compromised"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponse struct {
	Result  RadarLeakedCredentialCheckSummaryGetByBotClassResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarLeakedCredentialCheckSummaryGetByBotClassResponseJSON   `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseJSON contains the JSON
// metadata for the struct [RadarLeakedCredentialCheckSummaryGetByBotClassResponse]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResult struct {
	Meta     RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMeta     `json:"meta,required"`
	Summary0 RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarLeakedCredentialCheckSummaryGetByBotClassResponseResultJSON     `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultJSON contains the
// JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResult]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMeta struct {
	DateRange      []RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                         `json:"lastUpdated,required"`
	Normalization  string                                                                         `json:"normalization,required"`
	ConfidenceInfo RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaJSON           `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMeta]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRangeJSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRange]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfo struct {
	Annotations []RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                      `json:"level"`
	JSON        radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfo]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                       `json:"dataSource,required"`
	Description     string                                                                                       `json:"description,required"`
	EventType       string                                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                                    `json:"startTime" format:"date-time"`
	JSON            radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotation]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0 struct {
	Bot   string                                                                   `json:"bot,required"`
	Human string                                                                   `json:"human,required"`
	JSON  radarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0JSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0JSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0]
type radarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByBotClassResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponse struct {
	Result  RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResult `json:"result,required"`
	Success bool                                                                  `json:"success,required"`
	JSON    radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseJSON   `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseJSON contains the
// JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponse]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResult struct {
	Meta     RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMeta     `json:"meta,required"`
	Summary0 RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultJSON     `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResult]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMeta struct {
	DateRange      []RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                                  `json:"lastUpdated,required"`
	Normalization  string                                                                                  `json:"normalization,required"`
	ConfidenceInfo RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaJSON           `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMeta]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRangeJSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRange]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfo struct {
	Annotations []RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                               `json:"level"`
	JSON        radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfo]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                                `json:"dataSource,required"`
	Description     string                                                                                                `json:"description,required"`
	EventType       string                                                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                                             `json:"startTime" format:"date-time"`
	JSON            radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0 struct {
	Clean       string                                                                            `json:"CLEAN,required"`
	Compromised string                                                                            `json:"COMPROMISED,required"`
	JSON        radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0JSON `json:"-"`
}

// radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0JSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0]
type radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0JSON struct {
	Clean       apijson.Field
	Compromised apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckSummaryGetByCompromisedStatusResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckSummaryGetByBotClassParams struct {
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromised] `query:"compromised"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarLeakedCredentialCheckSummaryGetByBotClassParams]'s
// query parameters as `url.Values`.
func (r RadarLeakedCredentialCheckSummaryGetByBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromised string

const (
	RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromisedClean       RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromised = "CLEAN"
	RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromisedCompromised RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromised = "COMPROMISED"
)

func (r RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromised) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromisedClean, RadarLeakedCredentialCheckSummaryGetByBotClassParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormat string

const (
	RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormatJson RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormat = "JSON"
	RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormatCsv  RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormat = "CSV"
)

func (r RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormat) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormatJson, RadarLeakedCredentialCheckSummaryGetByBotClassParamsFormatCsv:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParams struct {
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes
// [RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParams]'s query
// parameters as `url.Values`.
func (r RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClass string

const (
	RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClassLikelyAutomated RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClass = "LIKELY_AUTOMATED"
	RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClassLikelyHuman     RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClass) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClassLikelyAutomated, RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsBotClassLikelyHuman:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormat string

const (
	RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormatJson RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormat = "JSON"
	RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormatCsv  RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormat = "CSV"
)

func (r RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormat) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormatJson, RadarLeakedCredentialCheckSummaryGetByCompromisedStatusParamsFormatCsv:
		return true
	}
	return false
}
