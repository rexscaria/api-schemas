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

// RadarBgpTopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpTopService] method instead.
type RadarBgpTopService struct {
	Options []option.RequestOption
	Ases    *RadarBgpTopAseService
}

// NewRadarBgpTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpTopService(opts ...option.RequestOption) (r *RadarBgpTopService) {
	r = &RadarBgpTopService{}
	r.Options = opts
	r.Ases = NewRadarBgpTopAseService(opts...)
	return
}

// Retrieves the top network prefixes by BGP updates.
func (r *RadarBgpTopService) ListTopPrefixes(ctx context.Context, query RadarBgpTopListTopPrefixesParams, opts ...option.RequestOption) (res *RadarBgpTopListTopPrefixesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/top/prefixes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpTopListTopPrefixesResponse struct {
	Result  RadarBgpTopListTopPrefixesResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarBgpTopListTopPrefixesResponseJSON   `json:"-"`
}

// radarBgpTopListTopPrefixesResponseJSON contains the JSON metadata for the struct
// [RadarBgpTopListTopPrefixesResponse]
type radarBgpTopListTopPrefixesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopListTopPrefixesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopListTopPrefixesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopListTopPrefixesResponseResult struct {
	Meta RadarBgpTopListTopPrefixesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarBgpTopListTopPrefixesResponseResultTop0 `json:"top_0,required"`
	JSON radarBgpTopListTopPrefixesResponseResultJSON   `json:"-"`
}

// radarBgpTopListTopPrefixesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpTopListTopPrefixesResponseResult]
type radarBgpTopListTopPrefixesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopListTopPrefixesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopListTopPrefixesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopListTopPrefixesResponseResultMeta struct {
	DateRange []RadarBgpTopListTopPrefixesResponseResultMetaDateRange `json:"dateRange,required"`
	JSON      radarBgpTopListTopPrefixesResponseResultMetaJSON        `json:"-"`
}

// radarBgpTopListTopPrefixesResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarBgpTopListTopPrefixesResponseResultMeta]
type radarBgpTopListTopPrefixesResponseResultMetaJSON struct {
	DateRange   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopListTopPrefixesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopListTopPrefixesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopListTopPrefixesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarBgpTopListTopPrefixesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarBgpTopListTopPrefixesResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarBgpTopListTopPrefixesResponseResultMetaDateRange]
type radarBgpTopListTopPrefixesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopListTopPrefixesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopListTopPrefixesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopListTopPrefixesResponseResultTop0 struct {
	Prefix string                                           `json:"prefix,required"`
	Value  string                                           `json:"value,required"`
	JSON   radarBgpTopListTopPrefixesResponseResultTop0JSON `json:"-"`
}

// radarBgpTopListTopPrefixesResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarBgpTopListTopPrefixesResponseResultTop0]
type radarBgpTopListTopPrefixesResponseResultTop0JSON struct {
	Prefix      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpTopListTopPrefixesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpTopListTopPrefixesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarBgpTopListTopPrefixesParams struct {
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
	Format param.Field[RadarBgpTopListTopPrefixesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of BGP update types.
	UpdateType param.Field[[]RadarBgpTopListTopPrefixesParamsUpdateType] `query:"updateType"`
}

// URLQuery serializes [RadarBgpTopListTopPrefixesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpTopListTopPrefixesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpTopListTopPrefixesParamsFormat string

const (
	RadarBgpTopListTopPrefixesParamsFormatJson RadarBgpTopListTopPrefixesParamsFormat = "JSON"
	RadarBgpTopListTopPrefixesParamsFormatCsv  RadarBgpTopListTopPrefixesParamsFormat = "CSV"
)

func (r RadarBgpTopListTopPrefixesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpTopListTopPrefixesParamsFormatJson, RadarBgpTopListTopPrefixesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpTopListTopPrefixesParamsUpdateType string

const (
	RadarBgpTopListTopPrefixesParamsUpdateTypeAnnouncement RadarBgpTopListTopPrefixesParamsUpdateType = "ANNOUNCEMENT"
	RadarBgpTopListTopPrefixesParamsUpdateTypeWithdrawal   RadarBgpTopListTopPrefixesParamsUpdateType = "WITHDRAWAL"
)

func (r RadarBgpTopListTopPrefixesParamsUpdateType) IsKnown() bool {
	switch r {
	case RadarBgpTopListTopPrefixesParamsUpdateTypeAnnouncement, RadarBgpTopListTopPrefixesParamsUpdateTypeWithdrawal:
		return true
	}
	return false
}
