// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// RadarBgpTopAseService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpTopAseService] method instead.
type RadarBgpTopAseService struct {
	Options []option.RequestOption
}

// NewRadarBgpTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpTopAseService(opts ...option.RequestOption) (r *RadarBgpTopAseService) {
	r = &RadarBgpTopAseService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by BGP updates (announcements only).
func (r *RadarBgpTopAseService) ListTopAses(ctx context.Context, query RadarBgpTopAseListTopAsesParams, opts ...option.RequestOption) (res *RadarBgpTopAseListTopAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the full list of autonomous systems on the global routing table
// ordered by announced prefixes count. The data comes from public BGP MRT data
// archives and updates every 2 hours.
func (r *RadarBgpTopAseService) ListTopPrefixes(ctx context.Context, query RadarBgpTopAseListTopPrefixesParams, opts ...option.RequestOption) (res *RadarBgpTopAseListTopPrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/ases/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTopAseListTopAsesResponse struct {
	Result  RadarBgpTopAseListTopAsesResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarBgpTopAseListTopAsesResponseJSON   `json:"-"`
}

// radarBgpTopAseListTopAsesResponseJSON contains the JSON metadata for the struct
// [RadarBgpTopAseListTopAsesResponse]
type radarBgpTopAseListTopAsesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopAsesResponseResult struct {
	Meta RadarBgpTopAseListTopAsesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarBgpTopAseListTopAsesResponseResultTop0 `json:"top_0,required"`
	JSON radarBgpTopAseListTopAsesResponseResultJSON   `json:"-"`
}

// radarBgpTopAseListTopAsesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpTopAseListTopAsesResponseResult]
type radarBgpTopAseListTopAsesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopAsesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopAsesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopAsesResponseResultMeta struct {
	DateRange []RadarBgpTopAseListTopAsesResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarBgpTopAseListTopAsesResponseResultMetaJSON        `json:"-"`
}

// radarBgpTopAseListTopAsesResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarBgpTopAseListTopAsesResponseResultMeta]
type radarBgpTopAseListTopAsesResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopAsesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopAsesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopAsesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarBgpTopAseListTopAsesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpTopAseListTopAsesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarBgpTopAseListTopAsesResponseResultMetaDateRange]
type radarBgpTopAseListTopAsesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopAsesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopAsesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopAsesResponseResultTop0 struct {
	Asn    int64  `json:"asn,required"`
	AsName string `json:"ASName,required"`
	// Percentage of updates by this AS out of the total updates by all autonomous
	// systems.
	Value string                                          `json:"value,required"`
	JSON  radarBgpTopAseListTopAsesResponseResultTop0JSON `json:"-"`
}

// radarBgpTopAseListTopAsesResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarBgpTopAseListTopAsesResponseResultTop0]
type radarBgpTopAseListTopAsesResponseResultTop0JSON struct {
	Asn         apijson.Field
	AsName      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopAsesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopAsesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopPrefixesResponse struct {
	Result  RadarBgpTopAseListTopPrefixesResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarBgpTopAseListTopPrefixesResponseJSON   `json:"-"`
}

// radarBgpTopAseListTopPrefixesResponseJSON contains the JSON metadata for the
// struct [RadarBgpTopAseListTopPrefixesResponse]
type radarBgpTopAseListTopPrefixesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopPrefixesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopPrefixesResponseResult struct {
	Asns []RadarBgpTopAseListTopPrefixesResponseResultAsn `json:"asns,required"`
	Meta RadarBgpTopAseListTopPrefixesResponseResultMeta  `json:"meta,required"`
	JSON radarBgpTopAseListTopPrefixesResponseResultJSON  `json:"-"`
}

// radarBgpTopAseListTopPrefixesResponseResultJSON contains the JSON metadata for
// the struct [RadarBgpTopAseListTopPrefixesResponseResult]
type radarBgpTopAseListTopPrefixesResponseResultJSON struct {
	Asns        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopPrefixesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopPrefixesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopPrefixesResponseResultAsn struct {
	Asn       int64                                              `json:"asn,required"`
	Country   string                                             `json:"country,required"`
	Name      string                                             `json:"name,required"`
	PfxsCount int64                                              `json:"pfxs_count,required"`
	JSON      radarBgpTopAseListTopPrefixesResponseResultAsnJSON `json:"-"`
}

// radarBgpTopAseListTopPrefixesResponseResultAsnJSON contains the JSON metadata
// for the struct [RadarBgpTopAseListTopPrefixesResponseResultAsn]
type radarBgpTopAseListTopPrefixesResponseResultAsnJSON struct {
	Asn         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopPrefixesResponseResultAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopPrefixesResponseResultAsnJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopPrefixesResponseResultMeta struct {
	DataTime   string                                              `json:"data_time,required"`
	QueryTime  string                                              `json:"query_time,required"`
	TotalPeers int64                                               `json:"total_peers,required"`
	JSON       radarBgpTopAseListTopPrefixesResponseResultMetaJSON `json:"-"`
}

// radarBgpTopAseListTopPrefixesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarBgpTopAseListTopPrefixesResponseResultMeta]
type radarBgpTopAseListTopPrefixesResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopAseListTopPrefixesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopAseListTopPrefixesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopAseListTopAsesParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpTopAseListTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP network prefixes.
	Prefix param.Field[[]string] `query:"prefix"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpTopAseListTopAsesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpTopAseListTopAsesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopAseListTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpTopAseListTopAsesParamsFormat string

const (
	RadarBgpTopAseListTopAsesParamsFormatJson RadarBgpTopAseListTopAsesParamsFormat = "JSON"
	RadarBgpTopAseListTopAsesParamsFormatCsv  RadarBgpTopAseListTopAsesParamsFormat = "CSV"
)

func (r RadarBgpTopAseListTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpTopAseListTopAsesParamsFormatJson, RadarBgpTopAseListTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpTopAseListTopAsesParamsUpdateType string

const (
	RadarBgpTopAseListTopAsesParamsUpdateTypeAnnouncement RadarBgpTopAseListTopAsesParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpTopAseListTopAsesParamsUpdateTypeWithdrawal   RadarBgpTopAseListTopAsesParamsUpdateType = "WITHDRAWAL"
)

func (r RadarBgpTopAseListTopAsesParamsUpdateType) IsKnown() bool {
	switch r {
	case RadarBgpTopAseListTopAsesParamsUpdateTypeAnnouncement, RadarBgpTopAseListTopAsesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}

type RadarBgpTopAseListTopPrefixesParams struct {
	// Alpha-2 country code.
	Country param.Field[string] `query:"country"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpTopAseListTopPrefixesParamsFormat] `query:"format"`
	// Maximum number of ASes to return.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [RadarBgpTopAseListTopPrefixesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopAseListTopPrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpTopAseListTopPrefixesParamsFormat string

const (
	RadarBgpTopAseListTopPrefixesParamsFormatJson RadarBgpTopAseListTopPrefixesParamsFormat = "JSON"
	RadarBgpTopAseListTopPrefixesParamsFormatCsv  RadarBgpTopAseListTopPrefixesParamsFormat = "CSV"
)

func (r RadarBgpTopAseListTopPrefixesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpTopAseListTopPrefixesParamsFormatJson, RadarBgpTopAseListTopPrefixesParamsFormatCsv:
		return true
	}
	return false
}
