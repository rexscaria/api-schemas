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

// RadarRobotsTxtTopUserAgentService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarRobotsTxtTopUserAgentService] method instead.
type RadarRobotsTxtTopUserAgentService struct {
	Options []option.RequestOption
}

// NewRadarRobotsTxtTopUserAgentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarRobotsTxtTopUserAgentService(opts ...option.RequestOption) (r *RadarRobotsTxtTopUserAgentService) {
	r = &RadarRobotsTxtTopUserAgentService{}
	r.Options = opts
	return
}

// Retrieves the top user agents on robots.txt files.
func (r *RadarRobotsTxtTopUserAgentService) GetTopUserAgents(ctx context.Context, query RadarRobotsTxtTopUserAgentGetTopUserAgentsParams, opts ...option.RequestOption) (res *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/user_agents/directive"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponse struct {
	Result  RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarRobotsTxtTopUserAgentGetTopUserAgentsResponseJSON   `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseJSON contains the JSON
// metadata for the struct [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponse]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResult struct {
	Meta RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0 `json:"top_0,required"`
	JSON radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultJSON   `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultJSON contains the JSON
// metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResult]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMeta struct {
	DateRange      []RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                     `json:"lastUpdated,required"`
	Normalization  string                                                                     `json:"normalization,required"`
	ConfidenceInfo RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	Units          []RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnit         `json:"units"`
	JSON           radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaJSON           `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMeta]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRange]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                  `json:"level"`
	JSON        radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfo]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                   `json:"dataSource,required"`
	Description     string                                                                                   `json:"description,required"`
	EventType       string                                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                                `json:"startTime" format:"date-time"`
	JSON            radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotation]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnitJSON `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnit]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0 struct {
	Name      string                                                           `json:"name,required"`
	Value     int64                                                            `json:"value,required"`
	Fully     int64                                                            `json:"fully"`
	Partially int64                                                            `json:"partially"`
	JSON      radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0JSON `json:"-"`
}

// radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0JSON contains the
// JSON metadata for the struct
// [RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0]
type radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Fully       apijson.Field
	Partially   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopUserAgentGetTopUserAgentsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopUserAgentGetTopUserAgentsParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Filters results by robots.txt directive.
	Directive param.Field[RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirective] `query:"directive"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format in which results will be returned.
	Format param.Field[RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by user agent category.
	UserAgentCategory param.Field[RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RadarRobotsTxtTopUserAgentGetTopUserAgentsParams]'s query
// parameters as `url.Values`.
func (r RadarRobotsTxtTopUserAgentGetTopUserAgentsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filters results by robots.txt directive.
type RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirective string

const (
	RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirectiveAllow    RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirective = "ALLOW"
	RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirectiveDisallow RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirective = "DISALLOW"
)

func (r RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirective) IsKnown() bool {
	switch r {
	case RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirectiveAllow, RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsDirectiveDisallow:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormat string

const (
	RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormatJson RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormat = "JSON"
	RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormatCsv  RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormat = "CSV"
)

func (r RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormat) IsKnown() bool {
	switch r {
	case RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormatJson, RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by user agent category.
type RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategory string

const (
	RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategoryAI RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategory = "AI"
)

func (r RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RadarRobotsTxtTopUserAgentGetTopUserAgentsParamsUserAgentCategoryAI:
		return true
	}
	return false
}
