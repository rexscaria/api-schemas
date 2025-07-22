// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// RadarRankingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarRankingService] method instead.
type RadarRankingService struct {
	Options          []option.RequestOption
	InternetServices *RadarRankingInternetServiceService
}

// NewRadarRankingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarRankingService(opts ...option.RequestOption) (r *RadarRankingService) {
	r = &RadarRankingService{}
	r.Options = opts
	r.InternetServices = NewRadarRankingInternetServiceService(opts...)
	return
}

// Retrieves domain rank details. Cloudflare provides an ordered rank for the top
// 100 domains, but for the remainder it only provides ranking buckets like top 200
// thousand, top one million, etc.. These are available through Radar datasets
// endpoints.
func (r *RadarRankingService) GetDomainRank(ctx context.Context, domain string, query RadarRankingGetDomainRankParams, opts ...option.RequestOption) (res *RadarRankingGetDomainRankResponse, err error) {
	opts = append(r.Options[:], opts...)
	if domain == "" {
		err = errors.New("missing required domain parameter")
		return
	}
	path := fmt.Sprintf("radar/ranking/domain/%s", domain)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves domains rank over time.
func (r *RadarRankingService) GetTimeseriesGroups(ctx context.Context, query RadarRankingGetTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarRankingGetTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top or trending domains based on their rank. Popular domains are
// domains of broad appeal based on how people use the Internet. Trending domains
// are domains that are generating a surge in interest. For more information on top
// domains, see https://blog.cloudflare.com/radar-domain-rankings/.
func (r *RadarRankingService) GetTopDomains(ctx context.Context, query RadarRankingGetTopDomainsParams, opts ...option.RequestOption) (res *RadarRankingGetTopDomainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingGetDomainRankResponse struct {
	Result  RadarRankingGetDomainRankResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarRankingGetDomainRankResponseJSON   `json:"-"`
}

// radarRankingGetDomainRankResponseJSON contains the JSON metadata for the struct
// [RadarRankingGetDomainRankResponse]
type radarRankingGetDomainRankResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResult struct {
	Details0 RadarRankingGetDomainRankResponseResultDetails0 `json:"details_0,required"`
	Meta     RadarRankingGetDomainRankResponseResultMeta     `json:"meta,required"`
	JSON     radarRankingGetDomainRankResponseResultJSON     `json:"-"`
}

// radarRankingGetDomainRankResponseResultJSON contains the JSON metadata for the
// struct [RadarRankingGetDomainRankResponseResult]
type radarRankingGetDomainRankResponseResultJSON struct {
	Details0    apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResultDetails0 struct {
	Categories []RadarRankingGetDomainRankResponseResultDetails0Category `json:"categories,required"`
	// Only available in POPULAR ranking for the most recent ranking.
	Bucket       string                                                       `json:"bucket"`
	Rank         int64                                                        `json:"rank"`
	TopLocations []RadarRankingGetDomainRankResponseResultDetails0TopLocation `json:"top_locations"`
	JSON         radarRankingGetDomainRankResponseResultDetails0JSON          `json:"-"`
}

// radarRankingGetDomainRankResponseResultDetails0JSON contains the JSON metadata
// for the struct [RadarRankingGetDomainRankResponseResultDetails0]
type radarRankingGetDomainRankResponseResultDetails0JSON struct {
	Categories   apijson.Field
	Bucket       apijson.Field
	Rank         apijson.Field
	TopLocations apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResultDetails0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultDetails0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResultDetails0Category struct {
	ID              int64                                                       `json:"id,required"`
	Name            string                                                      `json:"name,required"`
	SuperCategoryID int64                                                       `json:"superCategoryId,required"`
	JSON            radarRankingGetDomainRankResponseResultDetails0CategoryJSON `json:"-"`
}

// radarRankingGetDomainRankResponseResultDetails0CategoryJSON contains the JSON
// metadata for the struct
// [RadarRankingGetDomainRankResponseResultDetails0Category]
type radarRankingGetDomainRankResponseResultDetails0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResultDetails0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultDetails0CategoryJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResultDetails0TopLocation struct {
	LocationCode string                                                         `json:"locationCode,required"`
	LocationName string                                                         `json:"locationName,required"`
	Rank         int64                                                          `json:"rank,required"`
	JSON         radarRankingGetDomainRankResponseResultDetails0TopLocationJSON `json:"-"`
}

// radarRankingGetDomainRankResponseResultDetails0TopLocationJSON contains the JSON
// metadata for the struct
// [RadarRankingGetDomainRankResponseResultDetails0TopLocation]
type radarRankingGetDomainRankResponseResultDetails0TopLocationJSON struct {
	LocationCode apijson.Field
	LocationName apijson.Field
	Rank         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResultDetails0TopLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultDetails0TopLocationJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResultMeta struct {
	DateRange []RadarRankingGetDomainRankResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingGetDomainRankResponseResultMetaJSON        `json:"-"`
}

// radarRankingGetDomainRankResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarRankingGetDomainRankResponseResultMeta]
type radarRankingGetDomainRankResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarRankingGetDomainRankResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingGetDomainRankResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarRankingGetDomainRankResponseResultMetaDateRange]
type radarRankingGetDomainRankResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetDomainRankResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetDomainRankResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTimeseriesGroupsResponse struct {
	Result  RadarRankingGetTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarRankingGetTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarRankingGetTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [RadarRankingGetTimeseriesGroupsResponse]
type radarRankingGetTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTimeseriesGroupsResponseResult struct {
	Meta   RadarRankingGetTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarRankingGetTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarRankingGetTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarRankingGetTimeseriesGroupsResponseResultJSON contains the JSON metadata for
// the struct [RadarRankingGetTimeseriesGroupsResponseResult]
type radarRankingGetTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTimeseriesGroupsResponseResultMeta struct {
	DateRange []RadarRankingGetTimeseriesGroupsResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingGetTimeseriesGroupsResponseResultMetaJSON        `json:"-"`
}

// radarRankingGetTimeseriesGroupsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarRankingGetTimeseriesGroupsResponseResultMeta]
type radarRankingGetTimeseriesGroupsResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarRankingGetTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingGetTimeseriesGroupsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarRankingGetTimeseriesGroupsResponseResultMetaDateRange]
type radarRankingGetTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTimeseriesGroupsResponseResultSerie0 struct {
	Timestamps  []string                                                              `json:"timestamps,required"`
	ExtraFields map[string][]RadarRankingGetTimeseriesGroupsResponseResultSerie0Union `json:"-,extras"`
	JSON        radarRankingGetTimeseriesGroupsResponseResultSerie0JSON               `json:"-"`
}

// radarRankingGetTimeseriesGroupsResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarRankingGetTimeseriesGroupsResponseResultSerie0]
type radarRankingGetTimeseriesGroupsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RadarRankingGetTimeseriesGroupsResponseResultSerie0Union interface {
	ImplementsRadarRankingGetTimeseriesGroupsResponseResultSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RadarRankingGetTimeseriesGroupsResponseResultSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type RadarRankingGetTopDomainsResponse struct {
	Result  RadarRankingGetTopDomainsResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarRankingGetTopDomainsResponseJSON   `json:"-"`
}

// radarRankingGetTopDomainsResponseJSON contains the JSON metadata for the struct
// [RadarRankingGetTopDomainsResponse]
type radarRankingGetTopDomainsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTopDomainsResponseResult struct {
	Meta RadarRankingGetTopDomainsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRankingGetTopDomainsResponseResultTop0 `json:"top_0,required"`
	JSON radarRankingGetTopDomainsResponseResultJSON   `json:"-"`
}

// radarRankingGetTopDomainsResponseResultJSON contains the JSON metadata for the
// struct [RadarRankingGetTopDomainsResponseResult]
type radarRankingGetTopDomainsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTopDomainsResponseResultMeta struct {
	Top0 RadarRankingGetTopDomainsResponseResultMetaTop0 `json:"top_0,required"`
	JSON radarRankingGetTopDomainsResponseResultMetaJSON `json:"-"`
}

// radarRankingGetTopDomainsResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarRankingGetTopDomainsResponseResultMeta]
type radarRankingGetTopDomainsResponseResultMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTopDomainsResponseResultMetaTop0 struct {
	Date string                                              `json:"date,required"`
	JSON radarRankingGetTopDomainsResponseResultMetaTop0JSON `json:"-"`
}

// radarRankingGetTopDomainsResponseResultMetaTop0JSON contains the JSON metadata
// for the struct [RadarRankingGetTopDomainsResponseResultMetaTop0]
type radarRankingGetTopDomainsResponseResultMetaTop0JSON struct {
	Date        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponseResultMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseResultMetaTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTopDomainsResponseResultTop0 struct {
	Categories []RadarRankingGetTopDomainsResponseResultTop0Category `json:"categories,required"`
	Domain     string                                                `json:"domain,required"`
	Rank       int64                                                 `json:"rank,required"`
	// Only available in TRENDING rankings.
	PctRankChange float64                                         `json:"pctRankChange"`
	JSON          radarRankingGetTopDomainsResponseResultTop0JSON `json:"-"`
}

// radarRankingGetTopDomainsResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarRankingGetTopDomainsResponseResultTop0]
type radarRankingGetTopDomainsResponseResultTop0JSON struct {
	Categories    apijson.Field
	Domain        apijson.Field
	Rank          apijson.Field
	PctRankChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetTopDomainsResponseResultTop0Category struct {
	ID              float64                                                 `json:"id,required"`
	Name            string                                                  `json:"name,required"`
	SuperCategoryID float64                                                 `json:"superCategoryId,required"`
	JSON            radarRankingGetTopDomainsResponseResultTop0CategoryJSON `json:"-"`
}

// radarRankingGetTopDomainsResponseResultTop0CategoryJSON contains the JSON
// metadata for the struct [RadarRankingGetTopDomainsResponseResultTop0Category]
type radarRankingGetTopDomainsResponseResultTop0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingGetTopDomainsResponseResultTop0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingGetTopDomainsResponseResultTop0CategoryJSON) RawJSON() string {
	return r.raw
}

type RadarRankingGetDomainRankParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingGetDomainRankParamsFormat] `query:"format"`
	// Includes top locations in the response.
	IncludeTopLocations param.Field[bool] `query:"includeTopLocations"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Ranking type.
	RankingType param.Field[RadarRankingGetDomainRankParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingGetDomainRankParams]'s query parameters as
// `url.Values`.
func (r RadarRankingGetDomainRankParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingGetDomainRankParamsFormat string

const (
	RadarRankingGetDomainRankParamsFormatJson RadarRankingGetDomainRankParamsFormat = "JSON"
	RadarRankingGetDomainRankParamsFormatCsv  RadarRankingGetDomainRankParamsFormat = "CSV"
)

func (r RadarRankingGetDomainRankParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingGetDomainRankParamsFormatJson, RadarRankingGetDomainRankParamsFormatCsv:
		return true
	}
	return false
}

// Ranking type.
type RadarRankingGetDomainRankParamsRankingType string

const (
	RadarRankingGetDomainRankParamsRankingTypePopular        RadarRankingGetDomainRankParamsRankingType = "POPULAR"
	RadarRankingGetDomainRankParamsRankingTypeTrendingRise   RadarRankingGetDomainRankParamsRankingType = "TRENDING_RISE"
	RadarRankingGetDomainRankParamsRankingTypeTrendingSteady RadarRankingGetDomainRankParamsRankingType = "TRENDING_STEADY"
)

func (r RadarRankingGetDomainRankParamsRankingType) IsKnown() bool {
	switch r {
	case RadarRankingGetDomainRankParamsRankingTypePopular, RadarRankingGetDomainRankParamsRankingTypeTrendingRise, RadarRankingGetDomainRankParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RadarRankingGetTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Comma-separated list of domain names.
	Domains param.Field[[]string] `query:"domains"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingGetTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Ranking type.
	RankingType param.Field[RadarRankingGetTimeseriesGroupsParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingGetTimeseriesGroupsParams]'s query parameters
// as `url.Values`.
func (r RadarRankingGetTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingGetTimeseriesGroupsParamsFormat string

const (
	RadarRankingGetTimeseriesGroupsParamsFormatJson RadarRankingGetTimeseriesGroupsParamsFormat = "JSON"
	RadarRankingGetTimeseriesGroupsParamsFormatCsv  RadarRankingGetTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarRankingGetTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingGetTimeseriesGroupsParamsFormatJson, RadarRankingGetTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// Ranking type.
type RadarRankingGetTimeseriesGroupsParamsRankingType string

const (
	RadarRankingGetTimeseriesGroupsParamsRankingTypePopular        RadarRankingGetTimeseriesGroupsParamsRankingType = "POPULAR"
	RadarRankingGetTimeseriesGroupsParamsRankingTypeTrendingRise   RadarRankingGetTimeseriesGroupsParamsRankingType = "TRENDING_RISE"
	RadarRankingGetTimeseriesGroupsParamsRankingTypeTrendingSteady RadarRankingGetTimeseriesGroupsParamsRankingType = "TRENDING_STEADY"
)

func (r RadarRankingGetTimeseriesGroupsParamsRankingType) IsKnown() bool {
	switch r {
	case RadarRankingGetTimeseriesGroupsParamsRankingTypePopular, RadarRankingGetTimeseriesGroupsParamsRankingTypeTrendingRise, RadarRankingGetTimeseriesGroupsParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RadarRankingGetTopDomainsParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingGetTopDomainsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Ranking type.
	RankingType param.Field[RadarRankingGetTopDomainsParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RadarRankingGetTopDomainsParams]'s query parameters as
// `url.Values`.
func (r RadarRankingGetTopDomainsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingGetTopDomainsParamsFormat string

const (
	RadarRankingGetTopDomainsParamsFormatJson RadarRankingGetTopDomainsParamsFormat = "JSON"
	RadarRankingGetTopDomainsParamsFormatCsv  RadarRankingGetTopDomainsParamsFormat = "CSV"
)

func (r RadarRankingGetTopDomainsParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingGetTopDomainsParamsFormatJson, RadarRankingGetTopDomainsParamsFormatCsv:
		return true
	}
	return false
}

// Ranking type.
type RadarRankingGetTopDomainsParamsRankingType string

const (
	RadarRankingGetTopDomainsParamsRankingTypePopular        RadarRankingGetTopDomainsParamsRankingType = "POPULAR"
	RadarRankingGetTopDomainsParamsRankingTypeTrendingRise   RadarRankingGetTopDomainsParamsRankingType = "TRENDING_RISE"
	RadarRankingGetTopDomainsParamsRankingTypeTrendingSteady RadarRankingGetTopDomainsParamsRankingType = "TRENDING_STEADY"
)

func (r RadarRankingGetTopDomainsParamsRankingType) IsKnown() bool {
	switch r {
	case RadarRankingGetTopDomainsParamsRankingTypePopular, RadarRankingGetTopDomainsParamsRankingTypeTrendingRise, RadarRankingGetTopDomainsParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}
