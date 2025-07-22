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

// RadarRobotsTxtTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarRobotsTxtTopService] method instead.
type RadarRobotsTxtTopService struct {
	Options    []option.RequestOption
	UserAgents *RadarRobotsTxtTopUserAgentService
}

// NewRadarRobotsTxtTopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarRobotsTxtTopService(opts ...option.RequestOption) (r *RadarRobotsTxtTopService) {
	r = &RadarRobotsTxtTopService{}
	r.Options = opts
	r.UserAgents = NewRadarRobotsTxtTopUserAgentService(opts...)
	return
}

// Retrieves the top domain categories by the number of robots.txt files parsed.
func (r *RadarRobotsTxtTopService) GetTopDomainCategories(ctx context.Context, query RadarRobotsTxtTopGetTopDomainCategoriesParams, opts ...option.RequestOption) (res *RadarRobotsTxtTopGetTopDomainCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/robots_txt/top/domain_categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponse struct {
	Result  RadarRobotsTxtTopGetTopDomainCategoriesResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarRobotsTxtTopGetTopDomainCategoriesResponseJSON   `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseJSON contains the JSON metadata
// for the struct [RadarRobotsTxtTopGetTopDomainCategoriesResponse]
type radarRobotsTxtTopGetTopDomainCategoriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResult struct {
	Meta RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0 `json:"top_0,required"`
	JSON radarRobotsTxtTopGetTopDomainCategoriesResponseResultJSON   `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultJSON contains the JSON
// metadata for the struct [RadarRobotsTxtTopGetTopDomainCategoriesResponseResult]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMeta struct {
	DateRange      []RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                  `json:"lastUpdated,required"`
	Normalization  string                                                                  `json:"normalization,required"`
	ConfidenceInfo RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	Units          []RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnit         `json:"units"`
	JSON           radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaJSON           `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMeta]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRange]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                               `json:"level"`
	JSON        radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfo]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                `json:"dataSource,required"`
	Description     string                                                                                `json:"description,required"`
	EventType       string                                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                             `json:"startTime" format:"date-time"`
	JSON            radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotation]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnitJSON `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnit]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0 struct {
	Name  string                                                        `json:"name,required"`
	Value int64                                                         `json:"value,required"`
	JSON  radarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0JSON `json:"-"`
}

// radarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0]
type radarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRobotsTxtTopGetTopDomainCategoriesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRobotsTxtTopGetTopDomainCategoriesParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRobotsTxtTopGetTopDomainCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by user agent category.
	UserAgentCategory param.Field[RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategory] `query:"userAgentCategory"`
}

// URLQuery serializes [RadarRobotsTxtTopGetTopDomainCategoriesParams]'s query
// parameters as `url.Values`.
func (r RadarRobotsTxtTopGetTopDomainCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRobotsTxtTopGetTopDomainCategoriesParamsFormat string

const (
	RadarRobotsTxtTopGetTopDomainCategoriesParamsFormatJson RadarRobotsTxtTopGetTopDomainCategoriesParamsFormat = "JSON"
	RadarRobotsTxtTopGetTopDomainCategoriesParamsFormatCsv  RadarRobotsTxtTopGetTopDomainCategoriesParamsFormat = "CSV"
)

func (r RadarRobotsTxtTopGetTopDomainCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarRobotsTxtTopGetTopDomainCategoriesParamsFormatJson, RadarRobotsTxtTopGetTopDomainCategoriesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by user agent category.
type RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategory string

const (
	RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategoryAI RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategory = "AI"
)

func (r RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategory) IsKnown() bool {
	switch r {
	case RadarRobotsTxtTopGetTopDomainCategoriesParamsUserAgentCategoryAI:
		return true
	}
	return false
}
