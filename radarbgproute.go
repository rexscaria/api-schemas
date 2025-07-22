// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarBgpRouteService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarBgpRouteService] method instead.
type RadarBgpRouteService struct {
	Options []option.RequestOption
}

// NewRadarBgpRouteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarBgpRouteService(opts ...option.RequestOption) (r *RadarBgpRouteService) {
	r = &RadarBgpRouteService{}
	r.Options = opts
	return
}

// Retrieves all ASes in the current global routing tables with routing statistics.
func (r *RadarBgpRouteService) ListAses(ctx context.Context, query RadarBgpRouteListAsesParams, opts ...option.RequestOption) (res *RadarBgpRouteListAsesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves all Multi-Origin AS (MOAS) prefixes in the global routing tables.
func (r *RadarBgpRouteService) ListMoas(ctx context.Context, query RadarBgpRouteListMoasParams, opts ...option.RequestOption) (res *RadarBgpRouteListMoasResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/moas"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the prefix-to-ASN mapping from global routing tables.
func (r *RadarBgpRouteService) GetPrefixToAs(ctx context.Context, query RadarBgpRouteGetPrefixToAsParams, opts ...option.RequestOption) (res *RadarBgpRouteGetPrefixToAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/pfx2as"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves realtime routes for prefixes using public realtime data collectors
// (RouteViews and RIPE RIS).
func (r *RadarBgpRouteService) GetRealtimeRoutes(ctx context.Context, query RadarBgpRouteGetRealtimeRoutesParams, opts ...option.RequestOption) (res *RadarBgpRouteGetRealtimeRoutesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/realtime"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the BGP routing table stats.
func (r *RadarBgpRouteService) GetStats(ctx context.Context, query RadarBgpRouteGetStatsParams, opts ...option.RequestOption) (res *RadarBgpRouteGetStatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/routes/stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarBgpRouteListAsesResponse struct {
	Result  RadarBgpRouteListAsesResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpRouteListAsesResponseJSON   `json:"-"`
}

// radarBgpRouteListAsesResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteListAsesResponse]
type radarBgpRouteListAsesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListAsesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListAsesResponseResult struct {
	Asns []RadarBgpRouteListAsesResponseResultAsn `json:"asns,required"`
	Meta RadarBgpRouteListAsesResponseResultMeta  `json:"meta,required"`
	JSON radarBgpRouteListAsesResponseResultJSON  `json:"-"`
}

// radarBgpRouteListAsesResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRouteListAsesResponseResult]
type radarBgpRouteListAsesResponseResultJSON struct {
	Asns        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListAsesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListAsesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListAsesResponseResultAsn struct {
	Asn int64 `json:"asn,required"`
	// AS's customer cone size
	ConeSize int64 `json:"coneSize,required"`
	// 2-letter country code for the AS's registration country
	Country string `json:"country,required"`
	// number of IPv4 addresses originated by the AS
	Ipv4Count int64 `json:"ipv4Count,required"`
	// number of IPv6 addresses originated by the AS
	Ipv6Count string `json:"ipv6Count,required"`
	// name of the AS
	Name string `json:"name,required"`
	// number of total IP prefixes originated by the AS
	PfxsCount int64 `json:"pfxsCount,required"`
	// number of RPKI invalid prefixes originated by the AS
	RpkiInvalid int64 `json:"rpkiInvalid,required"`
	// number of RPKI unknown prefixes originated by the AS
	RpkiUnknown int64 `json:"rpkiUnknown,required"`
	// number of RPKI valid prefixes originated by the AS
	RpkiValid int64                                      `json:"rpkiValid,required"`
	JSON      radarBgpRouteListAsesResponseResultAsnJSON `json:"-"`
}

// radarBgpRouteListAsesResponseResultAsnJSON contains the JSON metadata for the
// struct [RadarBgpRouteListAsesResponseResultAsn]
type radarBgpRouteListAsesResponseResultAsnJSON struct {
	Asn         apijson.Field
	ConeSize    apijson.Field
	Country     apijson.Field
	Ipv4Count   apijson.Field
	Ipv6Count   apijson.Field
	Name        apijson.Field
	PfxsCount   apijson.Field
	RpkiInvalid apijson.Field
	RpkiUnknown apijson.Field
	RpkiValid   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListAsesResponseResultAsn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListAsesResponseResultAsnJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListAsesResponseResultMeta struct {
	// the timestamp of when the data is generated
	DataTime string `json:"dataTime,required"`
	// the timestamp of the query
	QueryTime string `json:"queryTime,required"`
	// total number of route collector peers used to generate this data
	TotalPeers int64                                       `json:"totalPeers,required"`
	JSON       radarBgpRouteListAsesResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteListAsesResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRouteListAsesResponseResultMeta]
type radarBgpRouteListAsesResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListAsesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListAsesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListMoasResponse struct {
	Result  RadarBgpRouteListMoasResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpRouteListMoasResponseJSON   `json:"-"`
}

// radarBgpRouteListMoasResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteListMoasResponse]
type radarBgpRouteListMoasResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListMoasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListMoasResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListMoasResponseResult struct {
	Meta RadarBgpRouteListMoasResponseResultMeta  `json:"meta,required"`
	Moas []RadarBgpRouteListMoasResponseResultMoa `json:"moas,required"`
	JSON radarBgpRouteListMoasResponseResultJSON  `json:"-"`
}

// radarBgpRouteListMoasResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRouteListMoasResponseResult]
type radarBgpRouteListMoasResponseResultJSON struct {
	Meta        apijson.Field
	Moas        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListMoasResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListMoasResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListMoasResponseResultMeta struct {
	DataTime   string                                      `json:"data_time,required"`
	QueryTime  string                                      `json:"query_time,required"`
	TotalPeers int64                                       `json:"total_peers,required"`
	JSON       radarBgpRouteListMoasResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteListMoasResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRouteListMoasResponseResultMeta]
type radarBgpRouteListMoasResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListMoasResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListMoasResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListMoasResponseResultMoa struct {
	Origins []RadarBgpRouteListMoasResponseResultMoasOrigin `json:"origins,required"`
	Prefix  string                                          `json:"prefix,required"`
	JSON    radarBgpRouteListMoasResponseResultMoaJSON      `json:"-"`
}

// radarBgpRouteListMoasResponseResultMoaJSON contains the JSON metadata for the
// struct [RadarBgpRouteListMoasResponseResultMoa]
type radarBgpRouteListMoasResponseResultMoaJSON struct {
	Origins     apijson.Field
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteListMoasResponseResultMoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListMoasResponseResultMoaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListMoasResponseResultMoasOrigin struct {
	Origin         int64                                             `json:"origin,required"`
	PeerCount      int64                                             `json:"peer_count,required"`
	RpkiValidation string                                            `json:"rpki_validation,required"`
	JSON           radarBgpRouteListMoasResponseResultMoasOriginJSON `json:"-"`
}

// radarBgpRouteListMoasResponseResultMoasOriginJSON contains the JSON metadata for
// the struct [RadarBgpRouteListMoasResponseResultMoasOrigin]
type radarBgpRouteListMoasResponseResultMoasOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpRouteListMoasResponseResultMoasOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteListMoasResponseResultMoasOriginJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetPrefixToAsResponse struct {
	Result  RadarBgpRouteGetPrefixToAsResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarBgpRouteGetPrefixToAsResponseJSON   `json:"-"`
}

// radarBgpRouteGetPrefixToAsResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteGetPrefixToAsResponse]
type radarBgpRouteGetPrefixToAsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetPrefixToAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetPrefixToAsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetPrefixToAsResponseResult struct {
	Meta          RadarBgpRouteGetPrefixToAsResponseResultMeta           `json:"meta,required"`
	PrefixOrigins []RadarBgpRouteGetPrefixToAsResponseResultPrefixOrigin `json:"prefix_origins,required"`
	JSON          radarBgpRouteGetPrefixToAsResponseResultJSON           `json:"-"`
}

// radarBgpRouteGetPrefixToAsResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRouteGetPrefixToAsResponseResult]
type radarBgpRouteGetPrefixToAsResponseResultJSON struct {
	Meta          apijson.Field
	PrefixOrigins apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RadarBgpRouteGetPrefixToAsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetPrefixToAsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetPrefixToAsResponseResultMeta struct {
	DataTime   string                                           `json:"data_time,required"`
	QueryTime  string                                           `json:"query_time,required"`
	TotalPeers int64                                            `json:"total_peers,required"`
	JSON       radarBgpRouteGetPrefixToAsResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteGetPrefixToAsResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarBgpRouteGetPrefixToAsResponseResultMeta]
type radarBgpRouteGetPrefixToAsResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetPrefixToAsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetPrefixToAsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetPrefixToAsResponseResultPrefixOrigin struct {
	Origin         int64                                                    `json:"origin,required"`
	PeerCount      int64                                                    `json:"peer_count,required"`
	Prefix         string                                                   `json:"prefix,required"`
	RpkiValidation string                                                   `json:"rpki_validation,required"`
	JSON           radarBgpRouteGetPrefixToAsResponseResultPrefixOriginJSON `json:"-"`
}

// radarBgpRouteGetPrefixToAsResponseResultPrefixOriginJSON contains the JSON
// metadata for the struct [RadarBgpRouteGetPrefixToAsResponseResultPrefixOrigin]
type radarBgpRouteGetPrefixToAsResponseResultPrefixOriginJSON struct {
	Origin         apijson.Field
	PeerCount      apijson.Field
	Prefix         apijson.Field
	RpkiValidation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarBgpRouteGetPrefixToAsResponseResultPrefixOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetPrefixToAsResponseResultPrefixOriginJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetRealtimeRoutesResponse struct {
	Result  RadarBgpRouteGetRealtimeRoutesResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarBgpRouteGetRealtimeRoutesResponseJSON   `json:"-"`
}

// radarBgpRouteGetRealtimeRoutesResponseJSON contains the JSON metadata for the
// struct [RadarBgpRouteGetRealtimeRoutesResponse]
type radarBgpRouteGetRealtimeRoutesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetRealtimeRoutesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetRealtimeRoutesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetRealtimeRoutesResponseResult struct {
	Meta   RadarBgpRouteGetRealtimeRoutesResponseResultMeta    `json:"meta,required"`
	Routes []RadarBgpRouteGetRealtimeRoutesResponseResultRoute `json:"routes,required"`
	JSON   radarBgpRouteGetRealtimeRoutesResponseResultJSON    `json:"-"`
}

// radarBgpRouteGetRealtimeRoutesResponseResultJSON contains the JSON metadata for
// the struct [RadarBgpRouteGetRealtimeRoutesResponseResult]
type radarBgpRouteGetRealtimeRoutesResponseResultJSON struct {
	Meta        apijson.Field
	Routes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetRealtimeRoutesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetRealtimeRoutesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetRealtimeRoutesResponseResultMeta struct {
	Collectors []RadarBgpRouteGetRealtimeRoutesResponseResultMetaCollector `json:"collectors,required"`
	JSON       radarBgpRouteGetRealtimeRoutesResponseResultMetaJSON        `json:"-"`
}

// radarBgpRouteGetRealtimeRoutesResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarBgpRouteGetRealtimeRoutesResponseResultMeta]
type radarBgpRouteGetRealtimeRoutesResponseResultMetaJSON struct {
	Collectors  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetRealtimeRoutesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetRealtimeRoutesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetRealtimeRoutesResponseResultMetaCollector struct {
	// public route collector ID
	Collector string `json:"collector,required"`
	// latest realtime stream timestamp for this collector
	LatestRealtimeTs string `json:"latest_realtime_ts,required"`
	// latest RIB dump MRT file timestamp for this collector
	LatestRibTs string `json:"latest_rib_ts,required"`
	// latest BGP updates MRT file timestamp for this collector
	LatestUpdatesTs string                                                        `json:"latest_updates_ts,required"`
	JSON            radarBgpRouteGetRealtimeRoutesResponseResultMetaCollectorJSON `json:"-"`
}

// radarBgpRouteGetRealtimeRoutesResponseResultMetaCollectorJSON contains the JSON
// metadata for the struct
// [RadarBgpRouteGetRealtimeRoutesResponseResultMetaCollector]
type radarBgpRouteGetRealtimeRoutesResponseResultMetaCollectorJSON struct {
	Collector        apijson.Field
	LatestRealtimeTs apijson.Field
	LatestRibTs      apijson.Field
	LatestUpdatesTs  apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarBgpRouteGetRealtimeRoutesResponseResultMetaCollector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetRealtimeRoutesResponseResultMetaCollectorJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetRealtimeRoutesResponseResultRoute struct {
	// AS-level path for this route, from collector to origin
	AsPath []int64 `json:"as_path,required"`
	// public collector ID for this route
	Collector string `json:"collector,required"`
	// BGP community values
	Communities []string `json:"communities,required"`
	// IP prefix of this query
	Prefix string `json:"prefix,required"`
	// latest timestamp of change for this route
	Timestamp string                                                `json:"timestamp,required"`
	JSON      radarBgpRouteGetRealtimeRoutesResponseResultRouteJSON `json:"-"`
}

// radarBgpRouteGetRealtimeRoutesResponseResultRouteJSON contains the JSON metadata
// for the struct [RadarBgpRouteGetRealtimeRoutesResponseResultRoute]
type radarBgpRouteGetRealtimeRoutesResponseResultRouteJSON struct {
	AsPath      apijson.Field
	Collector   apijson.Field
	Communities apijson.Field
	Prefix      apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetRealtimeRoutesResponseResultRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetRealtimeRoutesResponseResultRouteJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetStatsResponse struct {
	Result  RadarBgpRouteGetStatsResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarBgpRouteGetStatsResponseJSON   `json:"-"`
}

// radarBgpRouteGetStatsResponseJSON contains the JSON metadata for the struct
// [RadarBgpRouteGetStatsResponse]
type radarBgpRouteGetStatsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetStatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetStatsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetStatsResponseResult struct {
	Meta  RadarBgpRouteGetStatsResponseResultMeta  `json:"meta,required"`
	Stats RadarBgpRouteGetStatsResponseResultStats `json:"stats,required"`
	JSON  radarBgpRouteGetStatsResponseResultJSON  `json:"-"`
}

// radarBgpRouteGetStatsResponseResultJSON contains the JSON metadata for the
// struct [RadarBgpRouteGetStatsResponseResult]
type radarBgpRouteGetStatsResponseResultJSON struct {
	Meta        apijson.Field
	Stats       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetStatsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetStatsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetStatsResponseResultMeta struct {
	DataTime   string                                      `json:"data_time,required"`
	QueryTime  string                                      `json:"query_time,required"`
	TotalPeers int64                                       `json:"total_peers,required"`
	JSON       radarBgpRouteGetStatsResponseResultMetaJSON `json:"-"`
}

// radarBgpRouteGetStatsResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarBgpRouteGetStatsResponseResultMeta]
type radarBgpRouteGetStatsResponseResultMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarBgpRouteGetStatsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetStatsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteGetStatsResponseResultStats struct {
	DistinctOrigins      int64                                        `json:"distinct_origins,required"`
	DistinctOriginsIpv4  int64                                        `json:"distinct_origins_ipv4,required"`
	DistinctOriginsIpv6  int64                                        `json:"distinct_origins_ipv6,required"`
	DistinctPrefixes     int64                                        `json:"distinct_prefixes,required"`
	DistinctPrefixesIpv4 int64                                        `json:"distinct_prefixes_ipv4,required"`
	DistinctPrefixesIpv6 int64                                        `json:"distinct_prefixes_ipv6,required"`
	RoutesInvalid        int64                                        `json:"routes_invalid,required"`
	RoutesInvalidIpv4    int64                                        `json:"routes_invalid_ipv4,required"`
	RoutesInvalidIpv6    int64                                        `json:"routes_invalid_ipv6,required"`
	RoutesTotal          int64                                        `json:"routes_total,required"`
	RoutesTotalIpv4      int64                                        `json:"routes_total_ipv4,required"`
	RoutesTotalIpv6      int64                                        `json:"routes_total_ipv6,required"`
	RoutesUnknown        int64                                        `json:"routes_unknown,required"`
	RoutesUnknownIpv4    int64                                        `json:"routes_unknown_ipv4,required"`
	RoutesUnknownIpv6    int64                                        `json:"routes_unknown_ipv6,required"`
	RoutesValid          int64                                        `json:"routes_valid,required"`
	RoutesValidIpv4      int64                                        `json:"routes_valid_ipv4,required"`
	RoutesValidIpv6      int64                                        `json:"routes_valid_ipv6,required"`
	JSON                 radarBgpRouteGetStatsResponseResultStatsJSON `json:"-"`
}

// radarBgpRouteGetStatsResponseResultStatsJSON contains the JSON metadata for the
// struct [RadarBgpRouteGetStatsResponseResultStats]
type radarBgpRouteGetStatsResponseResultStatsJSON struct {
	DistinctOrigins      apijson.Field
	DistinctOriginsIpv4  apijson.Field
	DistinctOriginsIpv6  apijson.Field
	DistinctPrefixes     apijson.Field
	DistinctPrefixesIpv4 apijson.Field
	DistinctPrefixesIpv6 apijson.Field
	RoutesInvalid        apijson.Field
	RoutesInvalidIpv4    apijson.Field
	RoutesInvalidIpv6    apijson.Field
	RoutesTotal          apijson.Field
	RoutesTotalIpv4      apijson.Field
	RoutesTotalIpv6      apijson.Field
	RoutesUnknown        apijson.Field
	RoutesUnknownIpv4    apijson.Field
	RoutesUnknownIpv6    apijson.Field
	RoutesValid          apijson.Field
	RoutesValidIpv4      apijson.Field
	RoutesValidIpv6      apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarBgpRouteGetStatsResponseResultStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarBgpRouteGetStatsResponseResultStatsJSON) RawJSON() string {
	return r.raw
}

type RadarBgpRouteListAsesParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarBgpRouteListAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Location alpha-2 code.
	Location param.Field[string] `query:"location"`
	// Sorts results by the specified field.
	SortBy param.Field[RadarBgpRouteListAsesParamsSortBy] `query:"sortBy"`
	// Sort order.
	SortOrder param.Field[RadarBgpRouteListAsesParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [RadarBgpRouteListAsesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteListAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpRouteListAsesParamsFormat string

const (
	RadarBgpRouteListAsesParamsFormatJson RadarBgpRouteListAsesParamsFormat = "JSON"
	RadarBgpRouteListAsesParamsFormatCsv  RadarBgpRouteListAsesParamsFormat = "CSV"
)

func (r RadarBgpRouteListAsesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpRouteListAsesParamsFormatJson, RadarBgpRouteListAsesParamsFormatCsv:
		return true
	}
	return false
}

// Sorts results by the specified field.
type RadarBgpRouteListAsesParamsSortBy string

const (
	RadarBgpRouteListAsesParamsSortByCone        RadarBgpRouteListAsesParamsSortBy = "cone"
	RadarBgpRouteListAsesParamsSortByPfxs        RadarBgpRouteListAsesParamsSortBy = "pfxs"
	RadarBgpRouteListAsesParamsSortByIpv4        RadarBgpRouteListAsesParamsSortBy = "ipv4"
	RadarBgpRouteListAsesParamsSortByIpv6        RadarBgpRouteListAsesParamsSortBy = "ipv6"
	RadarBgpRouteListAsesParamsSortByRpkiValid   RadarBgpRouteListAsesParamsSortBy = "rpki_valid"
	RadarBgpRouteListAsesParamsSortByRpkiInvalid RadarBgpRouteListAsesParamsSortBy = "rpki_invalid"
	RadarBgpRouteListAsesParamsSortByRpkiUnknown RadarBgpRouteListAsesParamsSortBy = "rpki_unknown"
)

func (r RadarBgpRouteListAsesParamsSortBy) IsKnown() bool {
	switch r {
	case RadarBgpRouteListAsesParamsSortByCone, RadarBgpRouteListAsesParamsSortByPfxs, RadarBgpRouteListAsesParamsSortByIpv4, RadarBgpRouteListAsesParamsSortByIpv6, RadarBgpRouteListAsesParamsSortByRpkiValid, RadarBgpRouteListAsesParamsSortByRpkiInvalid, RadarBgpRouteListAsesParamsSortByRpkiUnknown:
		return true
	}
	return false
}

// Sort order.
type RadarBgpRouteListAsesParamsSortOrder string

const (
	RadarBgpRouteListAsesParamsSortOrderAsc  RadarBgpRouteListAsesParamsSortOrder = "ASC"
	RadarBgpRouteListAsesParamsSortOrderDesc RadarBgpRouteListAsesParamsSortOrder = "DESC"
)

func (r RadarBgpRouteListAsesParamsSortOrder) IsKnown() bool {
	switch r {
	case RadarBgpRouteListAsesParamsSortOrderAsc, RadarBgpRouteListAsesParamsSortOrderDesc:
		return true
	}
	return false
}

type RadarBgpRouteListMoasParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarBgpRouteListMoasParamsFormat] `query:"format"`
	// Lookup only RPKI invalid MOASes.
	InvalidOnly param.Field[bool] `query:"invalid_only"`
	// Lookup MOASes originated by the given ASN.
	Origin param.Field[int64] `query:"origin"`
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [RadarBgpRouteListMoasParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteListMoasParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpRouteListMoasParamsFormat string

const (
	RadarBgpRouteListMoasParamsFormatJson RadarBgpRouteListMoasParamsFormat = "JSON"
	RadarBgpRouteListMoasParamsFormatCsv  RadarBgpRouteListMoasParamsFormat = "CSV"
)

func (r RadarBgpRouteListMoasParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpRouteListMoasParamsFormatJson, RadarBgpRouteListMoasParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpRouteGetPrefixToAsParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarBgpRouteGetPrefixToAsParamsFormat] `query:"format"`
	// Return only results with the longest prefix match for the given prefix. For
	// example, specify a /32 prefix to lookup the origin ASN for an IPv4 address.
	LongestPrefixMatch param.Field[bool] `query:"longestPrefixMatch"`
	// Lookup prefixes originated by the given ASN.
	Origin param.Field[int64] `query:"origin"`
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
	// Return only results with matching rpki status: valid, invalid or unknown.
	RpkiStatus param.Field[RadarBgpRouteGetPrefixToAsParamsRpkiStatus] `query:"rpkiStatus"`
}

// URLQuery serializes [RadarBgpRouteGetPrefixToAsParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteGetPrefixToAsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpRouteGetPrefixToAsParamsFormat string

const (
	RadarBgpRouteGetPrefixToAsParamsFormatJson RadarBgpRouteGetPrefixToAsParamsFormat = "JSON"
	RadarBgpRouteGetPrefixToAsParamsFormatCsv  RadarBgpRouteGetPrefixToAsParamsFormat = "CSV"
)

func (r RadarBgpRouteGetPrefixToAsParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpRouteGetPrefixToAsParamsFormatJson, RadarBgpRouteGetPrefixToAsParamsFormatCsv:
		return true
	}
	return false
}

// Return only results with matching rpki status: valid, invalid or unknown.
type RadarBgpRouteGetPrefixToAsParamsRpkiStatus string

const (
	RadarBgpRouteGetPrefixToAsParamsRpkiStatusValid   RadarBgpRouteGetPrefixToAsParamsRpkiStatus = "VALID"
	RadarBgpRouteGetPrefixToAsParamsRpkiStatusInvalid RadarBgpRouteGetPrefixToAsParamsRpkiStatus = "INVALID"
	RadarBgpRouteGetPrefixToAsParamsRpkiStatusUnknown RadarBgpRouteGetPrefixToAsParamsRpkiStatus = "UNKNOWN"
)

func (r RadarBgpRouteGetPrefixToAsParamsRpkiStatus) IsKnown() bool {
	switch r {
	case RadarBgpRouteGetPrefixToAsParamsRpkiStatusValid, RadarBgpRouteGetPrefixToAsParamsRpkiStatusInvalid, RadarBgpRouteGetPrefixToAsParamsRpkiStatusUnknown:
		return true
	}
	return false
}

type RadarBgpRouteGetRealtimeRoutesParams struct {
	// Format in which results will be returned.
	Format param.Field[RadarBgpRouteGetRealtimeRoutesParamsFormat] `query:"format"`
	// Network prefix, IPv4 or IPv6.
	Prefix param.Field[string] `query:"prefix"`
}

// URLQuery serializes [RadarBgpRouteGetRealtimeRoutesParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteGetRealtimeRoutesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpRouteGetRealtimeRoutesParamsFormat string

const (
	RadarBgpRouteGetRealtimeRoutesParamsFormatJson RadarBgpRouteGetRealtimeRoutesParamsFormat = "JSON"
	RadarBgpRouteGetRealtimeRoutesParamsFormatCsv  RadarBgpRouteGetRealtimeRoutesParamsFormat = "CSV"
)

func (r RadarBgpRouteGetRealtimeRoutesParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpRouteGetRealtimeRoutesParamsFormatJson, RadarBgpRouteGetRealtimeRoutesParamsFormatCsv:
		return true
	}
	return false
}

type RadarBgpRouteGetStatsParams struct {
	// Single Autonomous System Number (ASN) as integer.
	Asn param.Field[int64] `query:"asn"`
	// Format in which results will be returned.
	Format param.Field[RadarBgpRouteGetStatsParamsFormat] `query:"format"`
	// Location alpha-2 code.
	Location param.Field[string] `query:"location"`
}

// URLQuery serializes [RadarBgpRouteGetStatsParams]'s query parameters as
// `url.Values`.
func (r RadarBgpRouteGetStatsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarBgpRouteGetStatsParamsFormat string

const (
	RadarBgpRouteGetStatsParamsFormatJson RadarBgpRouteGetStatsParamsFormat = "JSON"
	RadarBgpRouteGetStatsParamsFormatCsv  RadarBgpRouteGetStatsParamsFormat = "CSV"
)

func (r RadarBgpRouteGetStatsParamsFormat) IsKnown() bool {
	switch r {
	case RadarBgpRouteGetStatsParamsFormatJson, RadarBgpRouteGetStatsParamsFormatCsv:
		return true
	}
	return false
}
