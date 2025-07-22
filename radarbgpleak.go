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

// RadarBgpLeakService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpLeakService] method instead.
type RadarBgpLeakService struct {
	Options []option.RequestOption
}

// NewRadarBgpLeakService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpLeakService(opts ...option.RequestOption) (r *RadarBgpLeakService) {
	r = &RadarBgpLeakService{}
	r.Options = opts
	return
}

// Retrieves the BGP route leak events.
func (r *RadarBgpLeakService) ListEvents(ctx context.Context, query RadarBgpLeakListEventsParams, opts ...option.RequestOption) (res *RadarBgpLeakListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/leaks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpLeakListEventsResponse struct {
	Result     RadarBgpLeakListEventsResponseResult     `json:"result,required"`
	ResultInfo RadarBgpLeakListEventsResponseResultInfo `json:"result_info,required"`
	Success    bool                                     `json:"success,required"`
	JSON       radarBgpLeakListEventsResponseJSON       `json:"-"`
}

// radarBgpLeakListEventsResponseJSON contains the JSON metadata for the struct
// [RadarBgpLeakListEventsResponse]
type radarBgpLeakListEventsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpLeakListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpLeakListEventsResponseResult struct {
	AsnInfo []RadarBgpLeakListEventsResponseResultAsnInfo `json:"asn_info,required"`
	Events  []RadarBgpLeakListEventsResponseResultEvent   `json:"events,required"`
	JSON    radarBgpLeakListEventsResponseResultJSON      `json:"-"`
}

// radarBgpLeakListEventsResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpLeakListEventsResponseResult]
type radarBgpLeakListEventsResponseResultJSON struct {
	AsnInfo     apijson.Field
	Events      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakListEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpLeakListEventsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpLeakListEventsResponseResultAsnInfo struct {
	Asn         int64                                           `json:"asn,required"`
	CountryCode string                                          `json:"country_code,required"`
	OrgName     string                                          `json:"org_name,required"`
	JSON        radarBgpLeakListEventsResponseResultAsnInfoJSON `json:"-"`
}

// radarBgpLeakListEventsResponseResultAsnInfoJSON contains the JSON metadata for
// the struct [RadarBgpLeakListEventsResponseResultAsnInfo]
type radarBgpLeakListEventsResponseResultAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakListEventsResponseResultAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpLeakListEventsResponseResultAsnInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBgpLeakListEventsResponseResultEvent struct {
	ID          int64                                         `json:"id,required"`
	Countries   []string                                      `json:"countries,required"`
	DetectedTs  string                                        `json:"detected_ts,required"`
	Finished    bool                                          `json:"finished,required"`
	LeakAsn     int64                                         `json:"leak_asn,required"`
	LeakCount   int64                                         `json:"leak_count,required"`
	LeakSeg     []int64                                       `json:"leak_seg,required"`
	LeakType    int64                                         `json:"leak_type,required"`
	MaxTs       string                                        `json:"max_ts,required"`
	MinTs       string                                        `json:"min_ts,required"`
	OriginCount int64                                         `json:"origin_count,required"`
	PeerCount   int64                                         `json:"peer_count,required"`
	PrefixCount int64                                         `json:"prefix_count,required"`
	JSON        radarBgpLeakListEventsResponseResultEventJSON `json:"-"`
}

// radarBgpLeakListEventsResponseResultEventJSON contains the JSON metadata for the
// struct [RadarBgpLeakListEventsResponseResultEvent]
type radarBgpLeakListEventsResponseResultEventJSON struct {
	ID          apijson.Field
	Countries   apijson.Field
	DetectedTs  apijson.Field
	Finished    apijson.Field
	LeakAsn     apijson.Field
	LeakCount   apijson.Field
	LeakSeg     apijson.Field
	LeakType    apijson.Field
	MaxTs       apijson.Field
	MinTs       apijson.Field
	OriginCount apijson.Field
	PeerCount   apijson.Field
	PrefixCount apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakListEventsResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpLeakListEventsResponseResultEventJSON) RawJSON() string {
	return r.raw
}

type RadarBgpLeakListEventsResponseResultInfo struct {
	Count      int64                                        `json:"count,required"`
	Page       int64                                        `json:"page,required"`
	PerPage    int64                                        `json:"per_page,required"`
	TotalCount int64                                        `json:"total_count,required"`
	JSON       radarBgpLeakListEventsResponseResultInfoJSON `json:"-"`
}

// radarBgpLeakListEventsResponseResultInfoJSON contains the JSON metadata for the
// struct [RadarBgpLeakListEventsResponseResultInfo]
type radarBgpLeakListEventsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpLeakListEventsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpLeakListEventsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBgpLeakListEventsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Shorthand date ranges for the last X days - use when you don't need specific
	// start and end dates.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// The unique identifier of a event.
	EventID param.Field[int64] `query:"eventId"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpLeakListEventsParamsFormat] `query:"format"`
	// ASN that is causing or affected by a route leak event.
	InvolvedAsn param.Field[int64] `query:"involvedAsn"`
	// Country code of a involved ASN in a route leak event.
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The leaking AS of a route leak event.
	LeakAsn param.Field[int64] `query:"leakAsn"`
	// Current page number, starting from 1.
	Page param.Field[int64] `query:"page"`
	// Number of entries per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Sorts results by the specified field.
	SortBy param.Field[RadarBgpLeakListEventsParamsSortBy] `query:"sortBy"`
	// Sort order.
	SortOrder param.Field[RadarBgpLeakListEventsParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [RadarBgpLeakListEventsParams]'s query parameters as
// `url.Values`.
func (r RadarBgpLeakListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpLeakListEventsParamsFormat string

const (
	RadarBgpLeakListEventsParamsFormatJson RadarBgpLeakListEventsParamsFormat = "JSON"
	RadarBgpLeakListEventsParamsFormatCsv  RadarBgpLeakListEventsParamsFormat = "CSV"
)

func (r RadarBgpLeakListEventsParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpLeakListEventsParamsFormatJson, RadarBgpLeakListEventsParamsFormatCsv:
		return true
	}
	return false
}

// Sorts results by the specified field.
type RadarBgpLeakListEventsParamsSortBy string

const (
	RadarBgpLeakListEventsParamsSortByID       RadarBgpLeakListEventsParamsSortBy = "ID"
	RadarBgpLeakListEventsParamsSortByLeaks    RadarBgpLeakListEventsParamsSortBy = "LEAKS"
	RadarBgpLeakListEventsParamsSortByPeers    RadarBgpLeakListEventsParamsSortBy = "PEERS"
	RadarBgpLeakListEventsParamsSortByPrefixes RadarBgpLeakListEventsParamsSortBy = "PREFIXES"
	RadarBgpLeakListEventsParamsSortByOrigins  RadarBgpLeakListEventsParamsSortBy = "ORIGINS"
	RadarBgpLeakListEventsParamsSortByTime     RadarBgpLeakListEventsParamsSortBy = "TIME"
)

func (r RadarBgpLeakListEventsParamsSortBy) IsKnown() bool {
	switch r {
	case RadarBgpLeakListEventsParamsSortByID, RadarBgpLeakListEventsParamsSortByLeaks, RadarBgpLeakListEventsParamsSortByPeers, RadarBgpLeakListEventsParamsSortByPrefixes, RadarBgpLeakListEventsParamsSortByOrigins, RadarBgpLeakListEventsParamsSortByTime:
		return true
	}
	return false
}

// Sort order.
type RadarBgpLeakListEventsParamsSortOrder string

const (
	RadarBgpLeakListEventsParamsSortOrderAsc  RadarBgpLeakListEventsParamsSortOrder = "ASC"
	RadarBgpLeakListEventsParamsSortOrderDesc RadarBgpLeakListEventsParamsSortOrder = "DESC"
)

func (r RadarBgpLeakListEventsParamsSortOrder) IsKnown() bool {
	switch r {
	case RadarBgpLeakListEventsParamsSortOrderAsc, RadarBgpLeakListEventsParamsSortOrderDesc:
		return true
	}
	return false
}
