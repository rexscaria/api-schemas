// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// ZoneAnalyticsService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAnalyticsService] method instead.
type ZoneAnalyticsService struct {
	Options []option.RequestOption
	Latency *ZoneAnalyticsLatencyService
}

// NewZoneAnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAnalyticsService(opts ...option.RequestOption) (r *ZoneAnalyticsService) {
	r = &ZoneAnalyticsService{}
	r.Options = opts
	r.Latency = NewZoneAnalyticsLatencyService(opts...)
	return
}

// This view provides a breakdown of analytics data by datacenter. Note: This is
// available to Enterprise customers only.
func (r *ZoneAnalyticsService) ListColos(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsListColosParams, opts ...option.RequestOption) (res *ZoneAnalyticsListColosResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/analytics/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The dashboard view provides both totals and timeseries data for the given zone
// and time period across the entire Cloudflare network.
func (r *ZoneAnalyticsService) GetDashboard(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsGetDashboardParams, opts ...option.RequestOption) (res *ZoneAnalyticsGetDashboardResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneIdentifier == "" {
		err = errors.New("missing required zone_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/analytics/dashboard", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAnalyticsAPIResponse struct {
	Errors   []ZoneAnalyticsMessage              `json:"errors,required"`
	Messages []ZoneAnalyticsMessage              `json:"messages,required"`
	Result   ZoneAnalyticsAPIResponseResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success ZoneAnalyticsAPIResponseSuccess `json:"success,required"`
	JSON    zoneAnalyticsAPIResponseJSON    `json:"-"`
}

// zoneAnalyticsAPIResponseJSON contains the JSON metadata for the struct
// [ZoneAnalyticsAPIResponse]
type zoneAnalyticsAPIResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsAPIResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsAPIResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZoneAnalyticsAPIResponseResultArray] or
// [shared.UnionString].
type ZoneAnalyticsAPIResponseResultUnion interface {
	ImplementsZoneAnalyticsAPIResponseResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsAPIResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZoneAnalyticsAPIResponseResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneAnalyticsAPIResponseResultArray []interface{}

func (r ZoneAnalyticsAPIResponseResultArray) ImplementsZoneAnalyticsAPIResponseResultUnion() {}

// Whether the API call was successful
type ZoneAnalyticsAPIResponseSuccess bool

const (
	ZoneAnalyticsAPIResponseSuccessTrue ZoneAnalyticsAPIResponseSuccess = true
)

func (r ZoneAnalyticsAPIResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneAnalyticsAPIResponseSuccessTrue:
		return true
	}
	return false
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number in bytes served.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of bytes served to that country.
	Country interface{} `json:"country"`
	// A break down of bytes served over HTTPS.
	Ssl ZoneAnalyticsBandwidthSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsBandwidthSslProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                      `json:"uncached"`
	JSON     zoneAnalyticsBandwidthJSON `json:"-"`
}

// zoneAnalyticsBandwidthJSON contains the JSON metadata for the struct
// [ZoneAnalyticsBandwidth]
type zoneAnalyticsBandwidthJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsBandwidthJSON) RawJSON() string {
	return r.raw
}

// A break down of bytes served over HTTPS.
type ZoneAnalyticsBandwidthSsl struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                         `json:"unencrypted"`
	JSON        zoneAnalyticsBandwidthSslJSON `json:"-"`
}

// zoneAnalyticsBandwidthSslJSON contains the JSON metadata for the struct
// [ZoneAnalyticsBandwidthSsl]
type zoneAnalyticsBandwidthSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsBandwidthSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsBandwidthSslJSON) RawJSON() string {
	return r.raw
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsBandwidthSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                  `json:"TLSv1.3"`
	JSON    zoneAnalyticsBandwidthSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsBandwidthSslProtocolsJSON contains the JSON metadata for the struct
// [ZoneAnalyticsBandwidthSslProtocols]
type zoneAnalyticsBandwidthSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsBandwidthSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsBandwidthSslProtocolsJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsBandwidthByColo struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                            `json:"uncached"`
	JSON     zoneAnalyticsBandwidthByColoJSON `json:"-"`
}

// zoneAnalyticsBandwidthByColoJSON contains the JSON metadata for the struct
// [ZoneAnalyticsBandwidthByColo]
type zoneAnalyticsBandwidthByColoJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsBandwidthByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsBandwidthByColoJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsMessage struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    zoneAnalyticsMessageJSON `json:"-"`
}

// zoneAnalyticsMessageJSON contains the JSON metadata for the struct
// [ZoneAnalyticsMessage]
type zoneAnalyticsMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsMessageJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals for pageviews.
type ZoneAnalyticsPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                `json:"search_engine"`
	JSON         zoneAnalyticsPageviewsJSON `json:"-"`
}

// zoneAnalyticsPageviewsJSON contains the JSON metadata for the struct
// [ZoneAnalyticsPageviews]
type zoneAnalyticsPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsPageviewsJSON) RawJSON() string {
	return r.raw
}

// The exact parameters/timestamps the analytics service used to return data.
type ZoneAnalyticsQueryResponse struct {
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsSinceUnion `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsUntilUnion        `json:"until"`
	JSON  zoneAnalyticsQueryResponseJSON `json:"-"`
}

// zoneAnalyticsQueryResponseJSON contains the JSON metadata for the struct
// [ZoneAnalyticsQueryResponse]
type zoneAnalyticsQueryResponseJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsQueryResponseJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals for requests.
type ZoneAnalyticsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number of requests.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of requests served to that country.
	Country interface{} `json:"country"`
	// Key/value pairs where the key is a HTTP status code and the value is the number
	// of requests served with that code.
	HTTPStatus map[string]interface{} `json:"http_status"`
	// A break down of requests served over HTTPS.
	Ssl ZoneAnalyticsRequestsSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsRequestsSslProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                     `json:"uncached"`
	JSON     zoneAnalyticsRequestsJSON `json:"-"`
}

// zoneAnalyticsRequestsJSON contains the JSON metadata for the struct
// [ZoneAnalyticsRequests]
type zoneAnalyticsRequestsJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	HTTPStatus   apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsRequestsJSON) RawJSON() string {
	return r.raw
}

// A break down of requests served over HTTPS.
type ZoneAnalyticsRequestsSsl struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                        `json:"unencrypted"`
	JSON        zoneAnalyticsRequestsSslJSON `json:"-"`
}

// zoneAnalyticsRequestsSslJSON contains the JSON metadata for the struct
// [ZoneAnalyticsRequestsSsl]
type zoneAnalyticsRequestsSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsRequestsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsRequestsSslJSON) RawJSON() string {
	return r.raw
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsRequestsSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                 `json:"TLSv1.3"`
	JSON    zoneAnalyticsRequestsSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsRequestsSslProtocolsJSON contains the JSON metadata for the struct
// [ZoneAnalyticsRequestsSslProtocols]
type zoneAnalyticsRequestsSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsRequestsSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsRequestsSslProtocolsJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals for requests.
type ZoneAnalyticsRequestsByColo struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country map[string]interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                           `json:"uncached"`
	JSON     zoneAnalyticsRequestsByColoJSON `json:"-"`
}

// zoneAnalyticsRequestsByColoJSON contains the JSON metadata for the struct
// [ZoneAnalyticsRequestsByColo]
type zoneAnalyticsRequestsByColoJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsRequestsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsRequestsByColoJSON) RawJSON() string {
	return r.raw
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsSinceUnion interface {
	ImplementsZoneAnalyticsSinceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsSinceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type ZoneAnalyticsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}              `json:"type"`
	JSON zoneAnalyticsThreatsJSON `json:"-"`
}

// zoneAnalyticsThreatsJSON contains the JSON metadata for the struct
// [ZoneAnalyticsThreats]
type zoneAnalyticsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsThreatsJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                    `json:"all"`
	JSON zoneAnalyticsUniquesJSON `json:"-"`
}

// zoneAnalyticsUniquesJSON contains the JSON metadata for the struct
// [ZoneAnalyticsUniques]
type zoneAnalyticsUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsUniquesJSON) RawJSON() string {
	return r.raw
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsUntilUnion interface {
	ImplementsZoneAnalyticsUntilUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsUntilUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsUntilUnionParam interface {
	ImplementsZoneAnalyticsUntilUnionParam()
}

type ZoneAnalyticsListColosResponse struct {
	// The exact parameters/timestamps the analytics service used to return data.
	Query ZoneAnalyticsQueryResponse `json:"query"`
	// A breakdown of all dashboard analytics data by co-locations. This is limited to
	// Enterprise zones only.
	Result []ZoneAnalyticsListColosResponseResult `json:"result"`
	JSON   zoneAnalyticsListColosResponseJSON     `json:"-"`
	ZoneAnalyticsAPIResponse
}

// zoneAnalyticsListColosResponseJSON contains the JSON metadata for the struct
// [ZoneAnalyticsListColosResponse]
type zoneAnalyticsListColosResponseJSON struct {
	Query       apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsListColosResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsListColosResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsListColosResponseResult struct {
	// The airport code identifer for the co-location.
	ColoID string `json:"colo_id"`
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []ZoneAnalyticsListColosResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals ZoneAnalyticsListColosResponseResultTotals `json:"totals"`
	JSON   zoneAnalyticsListColosResponseResultJSON   `json:"-"`
}

// zoneAnalyticsListColosResponseResultJSON contains the JSON metadata for the
// struct [ZoneAnalyticsListColosResponseResult]
type zoneAnalyticsListColosResponseResultJSON struct {
	ColoID      apijson.Field
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsListColosResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsListColosResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsListColosResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsBandwidthByColo `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsRequestsByColo `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsSinceUnion `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsUntilUnion                          `json:"until"`
	JSON  zoneAnalyticsListColosResponseResultTimeseryJSON `json:"-"`
}

// zoneAnalyticsListColosResponseResultTimeseryJSON contains the JSON metadata for
// the struct [ZoneAnalyticsListColosResponseResultTimesery]
type zoneAnalyticsListColosResponseResultTimeseryJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsListColosResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsListColosResponseResultTimeseryJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals by data type.
type ZoneAnalyticsListColosResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsBandwidthByColo `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsRequestsByColo `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsSinceUnion `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsUntilUnion                        `json:"until"`
	JSON  zoneAnalyticsListColosResponseResultTotalsJSON `json:"-"`
}

// zoneAnalyticsListColosResponseResultTotalsJSON contains the JSON metadata for
// the struct [ZoneAnalyticsListColosResponseResultTotals]
type zoneAnalyticsListColosResponseResultTotalsJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsListColosResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsListColosResponseResultTotalsJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsGetDashboardResponse struct {
	// The exact parameters/timestamps the analytics service used to return data.
	Query ZoneAnalyticsQueryResponse `json:"query"`
	// Totals and timeseries data.
	Result ZoneAnalyticsGetDashboardResponseResult `json:"result"`
	JSON   zoneAnalyticsGetDashboardResponseJSON   `json:"-"`
	ZoneAnalyticsAPIResponse
}

// zoneAnalyticsGetDashboardResponseJSON contains the JSON metadata for the struct
// [ZoneAnalyticsGetDashboardResponse]
type zoneAnalyticsGetDashboardResponseJSON struct {
	Query       apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsGetDashboardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsGetDashboardResponseJSON) RawJSON() string {
	return r.raw
}

// Totals and timeseries data.
type ZoneAnalyticsGetDashboardResponseResult struct {
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []ZoneAnalyticsGetDashboardResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals ZoneAnalyticsGetDashboardResponseResultTotals `json:"totals"`
	JSON   zoneAnalyticsGetDashboardResponseResultJSON   `json:"-"`
}

// zoneAnalyticsGetDashboardResponseResultJSON contains the JSON metadata for the
// struct [ZoneAnalyticsGetDashboardResponseResult]
type zoneAnalyticsGetDashboardResponseResultJSON struct {
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsGetDashboardResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsGetDashboardResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsGetDashboardResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews ZoneAnalyticsPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsSinceUnion `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsThreats `json:"threats"`
	Uniques ZoneAnalyticsUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsUntilUnion                             `json:"until"`
	JSON  zoneAnalyticsGetDashboardResponseResultTimeseryJSON `json:"-"`
}

// zoneAnalyticsGetDashboardResponseResultTimeseryJSON contains the JSON metadata
// for the struct [ZoneAnalyticsGetDashboardResponseResultTimesery]
type zoneAnalyticsGetDashboardResponseResultTimeseryJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsGetDashboardResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsGetDashboardResponseResultTimeseryJSON) RawJSON() string {
	return r.raw
}

// Breakdown of totals by data type.
type ZoneAnalyticsGetDashboardResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews ZoneAnalyticsPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsSinceUnion `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsThreats `json:"threats"`
	Uniques ZoneAnalyticsUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsUntilUnion                           `json:"until"`
	JSON  zoneAnalyticsGetDashboardResponseResultTotalsJSON `json:"-"`
}

// zoneAnalyticsGetDashboardResponseResultTotalsJSON contains the JSON metadata for
// the struct [ZoneAnalyticsGetDashboardResponseResultTotals]
type zoneAnalyticsGetDashboardResponseResultTotalsJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsGetDashboardResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAnalyticsGetDashboardResponseResultTotalsJSON) RawJSON() string {
	return r.raw
}

type ZoneAnalyticsListColosParams struct {
	// When set to true, the API will move the requested time window backward, until it
	// finds a region with completely aggregated data.
	//
	// The API response _may not represent the requested time window_.
	Continuous param.Field[bool] `query:"continuous"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since param.Field[ZoneAnalyticsListColosParamsSinceUnion] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[ZoneAnalyticsUntilUnionParam] `query:"until"`
}

// URLQuery serializes [ZoneAnalyticsListColosParams]'s query parameters as
// `url.Values`.
func (r ZoneAnalyticsListColosParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsListColosParamsSinceUnion interface {
	ImplementsZoneAnalyticsListColosParamsSinceUnion()
}

type ZoneAnalyticsGetDashboardParams struct {
	// When set to true, the API will move the requested time window backward, until it
	// finds a region with completely aggregated data.
	//
	// The API response _may not represent the requested time window_.
	Continuous param.Field[bool] `query:"continuous"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since param.Field[ZoneAnalyticsGetDashboardParamsSinceUnion] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[ZoneAnalyticsUntilUnionParam] `query:"until"`
}

// URLQuery serializes [ZoneAnalyticsGetDashboardParams]'s query parameters as
// `url.Values`.
func (r ZoneAnalyticsGetDashboardParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsGetDashboardParamsSinceUnion interface {
	ImplementsZoneAnalyticsGetDashboardParamsSinceUnion()
}
