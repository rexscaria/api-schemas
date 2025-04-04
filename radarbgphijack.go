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

// RadarBgpHijackService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpHijackService] method instead.
type RadarBgpHijackService struct {
	Options []option.RequestOption
}

// NewRadarBgpHijackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpHijackService(opts ...option.RequestOption) (r *RadarBgpHijackService) {
	r = &RadarBgpHijackService{}
	r.Options = opts
	return
}

// Retrieves the BGP hijack events.
func (r *RadarBgpHijackService) ListEvents(ctx context.Context, query RadarBgpHijackListEventsParams, opts ...option.RequestOption) (res *RadarBgpHijackListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/hijacks/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpHijackListEventsResponse struct {
	Result     RadarBgpHijackListEventsResponseResult     `json:"result,required"`
	ResultInfo RadarBgpHijackListEventsResponseResultInfo `json:"result_info,required"`
	Success    bool                                       `json:"success,required"`
	JSON       radarBgpHijackListEventsResponseJSON       `json:"-"`
}

// radarBgpHijackListEventsResponseJSON contains the JSON metadata for the struct
// [RadarBgpHijackListEventsResponse]
type radarBgpHijackListEventsResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsResponseResult struct {
	AsnInfo       []RadarBgpHijackListEventsResponseResultAsnInfo `json:"asn_info,required"`
	Events        []RadarBgpHijackListEventsResponseResultEvent   `json:"events,required"`
	TotalMonitors int64                                           `json:"total_monitors,required"`
	JSON          radarBgpHijackListEventsResponseResultJSON      `json:"-"`
}

// radarBgpHijackListEventsResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpHijackListEventsResponseResult]
type radarBgpHijackListEventsResponseResultJSON struct {
	AsnInfo       apijson.Field
	Events        apijson.Field
	TotalMonitors apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsResponseResultAsnInfo struct {
	Asn         int64                                             `json:"asn,required"`
	CountryCode string                                            `json:"country_code,required"`
	OrgName     string                                            `json:"org_name,required"`
	JSON        radarBgpHijackListEventsResponseResultAsnInfoJSON `json:"-"`
}

// radarBgpHijackListEventsResponseResultAsnInfoJSON contains the JSON metadata for
// the struct [RadarBgpHijackListEventsResponseResultAsnInfo]
type radarBgpHijackListEventsResponseResultAsnInfoJSON struct {
	Asn         apijson.Field
	CountryCode apijson.Field
	OrgName     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponseResultAsnInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseResultAsnInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsResponseResultEvent struct {
	ID              int64                                             `json:"id,required"`
	ConfidenceScore int64                                             `json:"confidence_score,required"`
	Duration        int64                                             `json:"duration,required"`
	EventType       int64                                             `json:"event_type,required"`
	HijackMsgsCount int64                                             `json:"hijack_msgs_count,required"`
	HijackerAsn     int64                                             `json:"hijacker_asn,required"`
	HijackerCountry string                                            `json:"hijacker_country,required"`
	IsStale         bool                                              `json:"is_stale,required"`
	MaxHijackTs     string                                            `json:"max_hijack_ts,required"`
	MaxMsgTs        string                                            `json:"max_msg_ts,required"`
	MinHijackTs     string                                            `json:"min_hijack_ts,required"`
	OnGoingCount    int64                                             `json:"on_going_count,required"`
	PeerAsns        []int64                                           `json:"peer_asns,required"`
	PeerIPCount     int64                                             `json:"peer_ip_count,required"`
	Prefixes        []string                                          `json:"prefixes,required"`
	Tags            []RadarBgpHijackListEventsResponseResultEventsTag `json:"tags,required"`
	VictimAsns      []int64                                           `json:"victim_asns,required"`
	VictimCountries []string                                          `json:"victim_countries,required"`
	JSON            radarBgpHijackListEventsResponseResultEventJSON   `json:"-"`
}

// radarBgpHijackListEventsResponseResultEventJSON contains the JSON metadata for
// the struct [RadarBgpHijackListEventsResponseResultEvent]
type radarBgpHijackListEventsResponseResultEventJSON struct {
	ID              apijson.Field
	ConfidenceScore apijson.Field
	Duration        apijson.Field
	EventType       apijson.Field
	HijackMsgsCount apijson.Field
	HijackerAsn     apijson.Field
	HijackerCountry apijson.Field
	IsStale         apijson.Field
	MaxHijackTs     apijson.Field
	MaxMsgTs        apijson.Field
	MinHijackTs     apijson.Field
	OnGoingCount    apijson.Field
	PeerAsns        apijson.Field
	PeerIPCount     apijson.Field
	Prefixes        apijson.Field
	Tags            apijson.Field
	VictimAsns      apijson.Field
	VictimCountries apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseResultEventJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsResponseResultEventsTag struct {
	Name  string                                              `json:"name,required"`
	Score int64                                               `json:"score,required"`
	JSON  radarBgpHijackListEventsResponseResultEventsTagJSON `json:"-"`
}

// radarBgpHijackListEventsResponseResultEventsTagJSON contains the JSON metadata
// for the struct [RadarBgpHijackListEventsResponseResultEventsTag]
type radarBgpHijackListEventsResponseResultEventsTagJSON struct {
	Name        apijson.Field
	Score       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponseResultEventsTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseResultEventsTagJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsResponseResultInfo struct {
	Count      int64                                          `json:"count,required"`
	Page       int64                                          `json:"page,required"`
	PerPage    int64                                          `json:"per_page,required"`
	TotalCount int64                                          `json:"total_count,required"`
	JSON       radarBgpHijackListEventsResponseResultInfoJSON `json:"-"`
}

// radarBgpHijackListEventsResponseResultInfoJSON contains the JSON metadata for
// the struct [RadarBgpHijackListEventsResponseResultInfo]
type radarBgpHijackListEventsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpHijackListEventsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpHijackListEventsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type RadarBgpHijackListEventsParams struct {
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
	Format param.Field[RadarBgpHijackListEventsParamsFormat] `query:"format"`
	// The potential hijacker AS of a BGP hijack event.
	HijackerAsn param.Field[int64] `query:"hijackerAsn"`
	// The potential hijacker or victim AS of a BGP hijack event.
	InvolvedAsn param.Field[int64] `query:"involvedAsn"`
	// The country code of the potential hijacker or victim AS of a BGP hijack event.
	InvolvedCountry param.Field[string] `query:"involvedCountry"`
	// The maximum confidence score to filter events (1-4 low, 5-7 mid, 8+ high).
	MaxConfidence param.Field[int64] `query:"maxConfidence"`
	// The minimum confidence score to filter events (1-4 low, 5-7 mid, 8+ high).
	MinConfidence param.Field[int64] `query:"minConfidence"`
	// Current page number, starting from 1.
	Page param.Field[int64] `query:"page"`
	// Number of entries per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
	// Sorts results by the specified field.
	SortBy param.Field[RadarBgpHijackListEventsParamsSortBy] `query:"sortBy"`
	// Sort order.
	SortOrder param.Field[RadarBgpHijackListEventsParamsSortOrder] `query:"sortOrder"`
	// The potential victim AS of a BGP hijack event.
	VictimAsn param.Field[int64] `query:"victimAsn"`
}

// URLQuery serializes [RadarBgpHijackListEventsParams]'s query parameters as
// `url.Values`.
func (r RadarBgpHijackListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpHijackListEventsParamsFormat string

const (
	RadarBgpHijackListEventsParamsFormatJson RadarBgpHijackListEventsParamsFormat = "JSON"
	RadarBgpHijackListEventsParamsFormatCsv  RadarBgpHijackListEventsParamsFormat = "CSV"
)

func (r RadarBgpHijackListEventsParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpHijackListEventsParamsFormatJson, RadarBgpHijackListEventsParamsFormatCsv:
		return true
	}
	return false
}

// Sorts results by the specified field.
type RadarBgpHijackListEventsParamsSortBy string

const (
	RadarBgpHijackListEventsParamsSortByID         RadarBgpHijackListEventsParamsSortBy = "ID"
	RadarBgpHijackListEventsParamsSortByTime       RadarBgpHijackListEventsParamsSortBy = "TIME"
	RadarBgpHijackListEventsParamsSortByConfidence RadarBgpHijackListEventsParamsSortBy = "CONFIDENCE"
)

func (r RadarBgpHijackListEventsParamsSortBy) IsKnown() bool {
	switch r {
	case RadarBgpHijackListEventsParamsSortByID, RadarBgpHijackListEventsParamsSortByTime, RadarBgpHijackListEventsParamsSortByConfidence:
		return true
	}
	return false
}

// Sort order.
type RadarBgpHijackListEventsParamsSortOrder string

const (
	RadarBgpHijackListEventsParamsSortOrderAsc  RadarBgpHijackListEventsParamsSortOrder = "ASC"
	RadarBgpHijackListEventsParamsSortOrderDesc RadarBgpHijackListEventsParamsSortOrder = "DESC"
)

func (r RadarBgpHijackListEventsParamsSortOrder) IsKnown() bool {
	switch r {
	case RadarBgpHijackListEventsParamsSortOrderAsc, RadarBgpHijackListEventsParamsSortOrderDesc:
		return true
	}
	return false
}
