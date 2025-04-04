// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// RadarRankingInternetServiceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarRankingInternetServiceService] method instead.
type RadarRankingInternetServiceService struct {
	Options []option.RequestOption
}

// NewRadarRankingInternetServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarRankingInternetServiceService(opts ...option.RequestOption) (r *RadarRankingInternetServiceService) {
	r = &RadarRankingInternetServiceService{}
	r.Options = opts
	return
}

// Retrieves the list of Internet services categories.
func (r *RadarRankingInternetServiceService) ListCategories(ctx context.Context, query RadarRankingInternetServiceListCategoriesParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceListCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves Internet Services rank update changes over time.
func (r *RadarRankingInternetServiceService) GetTimeseriesGroups(ctx context.Context, query RadarRankingInternetServiceGetTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceGetTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves top Internet services based on their rank.
func (r *RadarRankingInternetServiceService) GetTopServices(ctx context.Context, query RadarRankingInternetServiceGetTopServicesParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceGetTopServicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingInternetServiceListCategoriesResponse struct {
	Result  RadarRankingInternetServiceListCategoriesResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarRankingInternetServiceListCategoriesResponseJSON   `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseJSON contains the JSON metadata
// for the struct [RadarRankingInternetServiceListCategoriesResponse]
type radarRankingInternetServiceListCategoriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesResponseResult struct {
	Categories0 []RadarRankingInternetServiceListCategoriesResponseResultCategories0 `json:"categories_0,required"`
	JSON        radarRankingInternetServiceListCategoriesResponseResultJSON          `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseResultJSON contains the JSON
// metadata for the struct
// [RadarRankingInternetServiceListCategoriesResponseResult]
type radarRankingInternetServiceListCategoriesResponseResultJSON struct {
	Categories0 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesResponseResultCategories0 struct {
	Name string                                                                 `json:"name,required"`
	JSON radarRankingInternetServiceListCategoriesResponseResultCategories0JSON `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseResultCategories0JSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceListCategoriesResponseResultCategories0]
type radarRankingInternetServiceListCategoriesResponseResultCategories0JSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponseResultCategories0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseResultCategories0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponse struct {
	Result  RadarRankingInternetServiceGetTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarRankingInternetServiceGetTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct [RadarRankingInternetServiceGetTimeseriesGroupsResponse]
type radarRankingInternetServiceGetTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResult struct {
	Meta   RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResult]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta struct {
	DateRange []RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON        `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0 struct {
	Timestamps  []string                                                                             `json:"timestamps,required"`
	ExtraFields map[string][]RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union `json:"-,extras"`
	JSON        radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON               `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union interface {
	ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union)(nil)).Elem(),
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

type RadarRankingInternetServiceGetTopServicesResponse struct {
	Result  RadarRankingInternetServiceGetTopServicesResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarRankingInternetServiceGetTopServicesResponseJSON   `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseJSON contains the JSON metadata
// for the struct [RadarRankingInternetServiceGetTopServicesResponse]
type radarRankingInternetServiceGetTopServicesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResult struct {
	Meta RadarRankingInternetServiceGetTopServicesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRankingInternetServiceGetTopServicesResponseResultTop0 `json:"top_0,required"`
	JSON radarRankingInternetServiceGetTopServicesResponseResultJSON   `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultJSON contains the JSON
// metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResult]
type radarRankingInternetServiceGetTopServicesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultMeta struct {
	Top0 RadarRankingInternetServiceGetTopServicesResponseResultMetaTop0 `json:"top_0,required"`
	JSON radarRankingInternetServiceGetTopServicesResponseResultMetaJSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMeta]
type radarRankingInternetServiceGetTopServicesResponseResultMetaJSON struct {
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultMetaTop0 struct {
	Date            string                                                              `json:"date,required"`
	ServiceCategory string                                                              `json:"serviceCategory,required"`
	JSON            radarRankingInternetServiceGetTopServicesResponseResultMetaTop0JSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaTop0JSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMetaTop0]
type radarRankingInternetServiceGetTopServicesResponseResultMetaTop0JSON struct {
	Date            apijson.Field
	ServiceCategory apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMetaTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultTop0 struct {
	Rank    int64                                                           `json:"rank,required"`
	Service string                                                          `json:"service,required"`
	JSON    radarRankingInternetServiceGetTopServicesResponseResultTop0JSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultTop0JSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultTop0]
type radarRankingInternetServiceGetTopServicesResponseResultTop0JSON struct {
	Rank        apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceListCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarRankingInternetServiceListCategoriesParams]'s query
// parameters as `url.Values`.
func (r RadarRankingInternetServiceListCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceListCategoriesParamsFormat string

const (
	RadarRankingInternetServiceListCategoriesParamsFormatJson RadarRankingInternetServiceListCategoriesParamsFormat = "JSON"
	RadarRankingInternetServiceListCategoriesParamsFormatCsv  RadarRankingInternetServiceListCategoriesParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceListCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceListCategoriesParamsFormatJson, RadarRankingInternetServiceListCategoriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RadarRankingInternetServiceGetTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarRankingInternetServiceGetTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat string

const (
	RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatJson RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat = "JSON"
	RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatCsv  RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatJson, RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTopServicesParams struct {
	// Array of dates to filter the results.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceGetTopServicesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RadarRankingInternetServiceGetTopServicesParams]'s query
// parameters as `url.Values`.
func (r RadarRankingInternetServiceGetTopServicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceGetTopServicesParamsFormat string

const (
	RadarRankingInternetServiceGetTopServicesParamsFormatJson RadarRankingInternetServiceGetTopServicesParamsFormat = "JSON"
	RadarRankingInternetServiceGetTopServicesParamsFormatCsv  RadarRankingInternetServiceGetTopServicesParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceGetTopServicesParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTopServicesParamsFormatJson, RadarRankingInternetServiceGetTopServicesParamsFormatCsv:
		return true
	}
	return false
}
