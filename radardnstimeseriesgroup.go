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

// RadarDNSTimeseriesGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarDNSTimeseriesGroupService] method instead.
type RadarDNSTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarDNSTimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarDNSTimeseriesGroupService(opts ...option.RequestOption) (r *RadarDNSTimeseriesGroupService) {
	r = &RadarDNSTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries by cache status over time.
func (r *RadarDNSTimeseriesGroupService) GetCacheHit(ctx context.Context, query RadarDNSTimeseriesGroupGetCacheHitParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetCacheHitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/cache_hit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS responses by DNSSEC (DNS Security Extensions)
// support over time.
func (r *RadarDNSTimeseriesGroupService) GetDnssec(ctx context.Context, query RadarDNSTimeseriesGroupGetDnssecParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by DNSSEC (DNS Security Extensions)
// client awareness over time.
func (r *RadarDNSTimeseriesGroupService) GetDnssecAware(ctx context.Context, query RadarDNSTimeseriesGroupGetDnssecAwareParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetDnssecAwareResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec_aware"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNSSEC-validated answers by end-to-end security
// status over time.
func (r *RadarDNSTimeseriesGroupService) GetDnssecE2E(ctx context.Context, query RadarDNSTimeseriesGroupGetDnssecE2EParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetDnssecE2EResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/dnssec_e2e"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by IP version over time.
func (r *RadarDNSTimeseriesGroupService) GetIPVersion(ctx context.Context, query RadarDNSTimeseriesGroupGetIPVersionParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by matching answers over time.
func (r *RadarDNSTimeseriesGroupService) GetMatchingAnswer(ctx context.Context, query RadarDNSTimeseriesGroupGetMatchingAnswerParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetMatchingAnswerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/matching_answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by DNS transport protocol over time.
func (r *RadarDNSTimeseriesGroupService) GetProtocol(ctx context.Context, query RadarDNSTimeseriesGroupGetProtocolParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by type over time.
func (r *RadarDNSTimeseriesGroupService) GetQueryType(ctx context.Context, query RadarDNSTimeseriesGroupGetQueryTypeParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by response code over time.
func (r *RadarDNSTimeseriesGroupService) GetResponseCode(ctx context.Context, query RadarDNSTimeseriesGroupGetResponseCodeParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetResponseCodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/response_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by minimum answer TTL over time.
func (r *RadarDNSTimeseriesGroupService) GetResponseTtl(ctx context.Context, query RadarDNSTimeseriesGroupGetResponseTtlParams, opts ...option.RequestOption) (res *RadarDNSTimeseriesGroupGetResponseTtlResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/timeseries_groups/response_ttl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSTimeseriesGroupGetCacheHitResponse struct {
	Result  RadarDNSTimeseriesGroupGetCacheHitResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetCacheHitResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetCacheHitResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetCacheHitResponse]
type radarDNSTimeseriesGroupGetCacheHitResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetCacheHitResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetCacheHitResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetCacheHitResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetCacheHitResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetCacheHitResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetCacheHitResponseResult]
type radarDNSTimeseriesGroupGetCacheHitResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetCacheHitResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetCacheHitResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetCacheHitResponseResultSerie0 struct {
	Negative []string                                                   `json:"NEGATIVE,required"`
	Positive []string                                                   `json:"POSITIVE,required"`
	JSON     radarDNSTimeseriesGroupGetCacheHitResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetCacheHitResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetCacheHitResponseResultSerie0]
type radarDNSTimeseriesGroupGetCacheHitResponseResultSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetCacheHitResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetCacheHitResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecResponse struct {
	Result  RadarDNSTimeseriesGroupGetDnssecResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetDnssecResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecResponseJSON contains the JSON metadata for the
// struct [RadarDNSTimeseriesGroupGetDnssecResponse]
type radarDNSTimeseriesGroupGetDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecResponseResult struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetDnssecResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetDnssecResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetDnssecResponseResult]
type radarDNSTimeseriesGroupGetDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecResponseResultSerie0 struct {
	Insecure []string                                                 `json:"INSECURE,required"`
	Invalid  []string                                                 `json:"INVALID,required"`
	Other    []string                                                 `json:"OTHER,required"`
	Secure   []string                                                 `json:"SECURE,required"`
	JSON     radarDNSTimeseriesGroupGetDnssecResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetDnssecResponseResultSerie0]
type radarDNSTimeseriesGroupGetDnssecResponseResultSerie0JSON struct {
	Insecure    apijson.Field
	Invalid     apijson.Field
	Other       apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecAwareResponse struct {
	Result  RadarDNSTimeseriesGroupGetDnssecAwareResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetDnssecAwareResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecAwareResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetDnssecAwareResponse]
type radarDNSTimeseriesGroupGetDnssecAwareResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecAwareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecAwareResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecAwareResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetDnssecAwareResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecAwareResponseResultJSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetDnssecAwareResponseResult]
type radarDNSTimeseriesGroupGetDnssecAwareResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecAwareResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecAwareResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0 struct {
	NotSupported []string                                                      `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                      `json:"SUPPORTED,required"`
	JSON         radarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0]
type radarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecAwareResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecE2EResponse struct {
	Result  RadarDNSTimeseriesGroupGetDnssecE2EResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetDnssecE2EResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecE2EResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetDnssecE2EResponse]
type radarDNSTimeseriesGroupGetDnssecE2EResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecE2EResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecE2EResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecE2EResponseResult struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetDnssecE2EResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecE2EResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetDnssecE2EResponseResult]
type radarDNSTimeseriesGroupGetDnssecE2EResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecE2EResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecE2EResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0 struct {
	Negative []string                                                    `json:"NEGATIVE,required"`
	Positive []string                                                    `json:"POSITIVE,required"`
	JSON     radarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0]
type radarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetDnssecE2EResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetIPVersionResponse struct {
	Result  RadarDNSTimeseriesGroupGetIPVersionResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetIPVersionResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetIPVersionResponse]
type radarDNSTimeseriesGroupGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetIPVersionResponseResult struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetIPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetIPVersionResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetIPVersionResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetIPVersionResponseResult]
type radarDNSTimeseriesGroupGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetIPVersionResponseResultSerie0 struct {
	IPv4 []string                                                    `json:"IPv4,required"`
	IPv6 []string                                                    `json:"IPv6,required"`
	JSON radarDNSTimeseriesGroupGetIPVersionResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetIPVersionResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetIPVersionResponseResultSerie0]
type radarDNSTimeseriesGroupGetIPVersionResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetIPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetIPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetMatchingAnswerResponse struct {
	Result  RadarDNSTimeseriesGroupGetMatchingAnswerResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetMatchingAnswerResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetMatchingAnswerResponseJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetMatchingAnswerResponse]
type radarDNSTimeseriesGroupGetMatchingAnswerResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetMatchingAnswerResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetMatchingAnswerResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetMatchingAnswerResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetMatchingAnswerResponseResultJSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetMatchingAnswerResponseResult]
type radarDNSTimeseriesGroupGetMatchingAnswerResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetMatchingAnswerResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetMatchingAnswerResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0 struct {
	Negative []string                                                         `json:"NEGATIVE,required"`
	Positive []string                                                         `json:"POSITIVE,required"`
	JSON     radarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0]
type radarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetMatchingAnswerResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetProtocolResponse struct {
	Result  RadarDNSTimeseriesGroupGetProtocolResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetProtocolResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetProtocolResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetProtocolResponse]
type radarDNSTimeseriesGroupGetProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetProtocolResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetProtocolResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetProtocolResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetProtocolResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetProtocolResponseResult]
type radarDNSTimeseriesGroupGetProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetProtocolResponseResultSerie0 struct {
	HTTPS []string                                                   `json:"HTTPS,required"`
	Tcp   []string                                                   `json:"TCP,required"`
	Tls   []string                                                   `json:"TLS,required"`
	Udp   []string                                                   `json:"UDP,required"`
	JSON  radarDNSTimeseriesGroupGetProtocolResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetProtocolResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetProtocolResponseResultSerie0]
type radarDNSTimeseriesGroupGetProtocolResponseResultSerie0JSON struct {
	HTTPS       apijson.Field
	Tcp         apijson.Field
	Tls         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetProtocolResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetProtocolResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetQueryTypeResponse struct {
	Result  RadarDNSTimeseriesGroupGetQueryTypeResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetQueryTypeResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetQueryTypeResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetQueryTypeResponse]
type radarDNSTimeseriesGroupGetQueryTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetQueryTypeResponseResult struct {
	Meta   interface{}                                             `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetQueryTypeResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetQueryTypeResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetQueryTypeResponseResult]
type radarDNSTimeseriesGroupGetQueryTypeResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetQueryTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetQueryTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0 struct {
	Timestamps  []string                                                    `json:"timestamps,required"`
	ExtraFields map[string][]string                                         `json:"-,extras"`
	JSON        radarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0]
type radarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetQueryTypeResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseCodeResponse struct {
	Result  RadarDNSTimeseriesGroupGetResponseCodeResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetResponseCodeResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseCodeResponseJSON contains the JSON metadata
// for the struct [RadarDNSTimeseriesGroupGetResponseCodeResponse]
type radarDNSTimeseriesGroupGetResponseCodeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseCodeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseCodeResponseResult struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetResponseCodeResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseCodeResponseResultJSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetResponseCodeResponseResult]
type radarDNSTimeseriesGroupGetResponseCodeResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseCodeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseCodeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0 struct {
	Timestamps  []string                                                       `json:"timestamps,required"`
	ExtraFields map[string][]string                                            `json:"-,extras"`
	JSON        radarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0]
type radarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseCodeResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseTtlResponse struct {
	Result  RadarDNSTimeseriesGroupGetResponseTtlResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarDNSTimeseriesGroupGetResponseTtlResponseJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseTtlResponseJSON contains the JSON metadata for
// the struct [RadarDNSTimeseriesGroupGetResponseTtlResponse]
type radarDNSTimeseriesGroupGetResponseTtlResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseTtlResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseTtlResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseTtlResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarDNSTimeseriesGroupGetResponseTtlResponseResultJSON   `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseTtlResponseResultJSON contains the JSON
// metadata for the struct [RadarDNSTimeseriesGroupGetResponseTtlResponseResult]
type radarDNSTimeseriesGroupGetResponseTtlResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseTtlResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseTtlResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0 struct {
	Gt15mLte1h []string                                                      `json:"gt_15m_lte_1h,required"`
	Gt1dLte1w  []string                                                      `json:"gt_1d_lte_1w,required"`
	Gt1hLte1d  []string                                                      `json:"gt_1h_lte_1d,required"`
	Gt1mLte5m  []string                                                      `json:"gt_1m_lte_5m,required"`
	Gt1w       []string                                                      `json:"gt_1w,required"`
	Gt5mLte15m []string                                                      `json:"gt_5m_lte_15m,required"`
	Lte1m      []string                                                      `json:"lte_1m,required"`
	JSON       radarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0JSON `json:"-"`
}

// radarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0]
type radarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0JSON struct {
	Gt15mLte1h  apijson.Field
	Gt1dLte1w   apijson.Field
	Gt1hLte1d   apijson.Field
	Gt1mLte5m   apijson.Field
	Gt1w        apijson.Field
	Gt5mLte15m  apijson.Field
	Lte1m       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSTimeseriesGroupGetResponseTtlResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSTimeseriesGroupGetCacheHitParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetCacheHitParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetCacheHitParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetCacheHitParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetCacheHitParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetCacheHitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval15m RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1h  RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1d  RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1w  RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval15m, RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1h, RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1d, RadarDNSTimeseriesGroupGetCacheHitParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetCacheHitParamsFormat string

const (
	RadarDNSTimeseriesGroupGetCacheHitParamsFormatJson RadarDNSTimeseriesGroupGetCacheHitParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetCacheHitParamsFormatCsv  RadarDNSTimeseriesGroupGetCacheHitParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetCacheHitParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetCacheHitParamsFormatJson, RadarDNSTimeseriesGroupGetCacheHitParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetCacheHitParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetCacheHitParamsProtocolUdp   RadarDNSTimeseriesGroupGetCacheHitParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetCacheHitParamsProtocolTcp   RadarDNSTimeseriesGroupGetCacheHitParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetCacheHitParamsProtocolHTTPS RadarDNSTimeseriesGroupGetCacheHitParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetCacheHitParamsProtocolTls   RadarDNSTimeseriesGroupGetCacheHitParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetCacheHitParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetCacheHitParamsProtocolUdp, RadarDNSTimeseriesGroupGetCacheHitParamsProtocolTcp, RadarDNSTimeseriesGroupGetCacheHitParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetCacheHitParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetCacheHitParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeA          RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeA6         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAny        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeApl        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCds        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCert       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCname      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDname      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDs         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEid        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeGid        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHip        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeKey        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeKx         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeL32        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeL64        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeLp         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMB         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMd         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMf         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMg         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMr         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMx         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNb         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNid        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNs         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNull       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypePtr        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypePx         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRp         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRt         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSig        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSink       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTa         RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUid        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUri        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeWks        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeX25        RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetCacheHitParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetCacheHitParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeA, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeA6, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAny, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeApl, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCds, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCert, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCname, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDname, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeDs, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEid, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeGid, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHip, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeKey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeKx, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeL32, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeL64, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeLp, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMB, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMd, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMf, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMg, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMr, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeMx, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNb, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNid, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNs, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNull, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypePtr, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypePx, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRp, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeRt, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSig, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSink, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUid, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeUri, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeWks, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeX25, RadarDNSTimeseriesGroupGetCacheHitParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetCacheHitParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetCacheHitParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetDnssecParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetDnssecParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetDnssecParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetDnssecParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetDnssecParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetDnssecParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetDnssecParams]'s query parameters
// as `url.Values`.
func (r RadarDNSTimeseriesGroupGetDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetDnssecParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetDnssecParamsAggInterval15m RadarDNSTimeseriesGroupGetDnssecParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1h  RadarDNSTimeseriesGroupGetDnssecParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1d  RadarDNSTimeseriesGroupGetDnssecParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1w  RadarDNSTimeseriesGroupGetDnssecParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetDnssecParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecParamsAggInterval15m, RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1h, RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1d, RadarDNSTimeseriesGroupGetDnssecParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetDnssecParamsFormat string

const (
	RadarDNSTimeseriesGroupGetDnssecParamsFormatJson RadarDNSTimeseriesGroupGetDnssecParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetDnssecParamsFormatCsv  RadarDNSTimeseriesGroupGetDnssecParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecParamsFormatJson, RadarDNSTimeseriesGroupGetDnssecParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetDnssecParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetDnssecParamsProtocolUdp   RadarDNSTimeseriesGroupGetDnssecParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetDnssecParamsProtocolTcp   RadarDNSTimeseriesGroupGetDnssecParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetDnssecParamsProtocolHTTPS RadarDNSTimeseriesGroupGetDnssecParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecParamsProtocolTls   RadarDNSTimeseriesGroupGetDnssecParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetDnssecParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecParamsProtocolUdp, RadarDNSTimeseriesGroupGetDnssecParamsProtocolTcp, RadarDNSTimeseriesGroupGetDnssecParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetDnssecParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetDnssecParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeA          RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeA6         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAny        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeApl        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCds        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCert       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCname      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDname      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDs         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEid        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeGid        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHip        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeKey        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeKx         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeL32        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeL64        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeLp         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMB         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMd         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMf         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMg         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMr         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMx         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNb         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNid        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNs         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNull       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypePtr        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypePx         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRp         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRt         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSig        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSink       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTa         RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUid        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUri        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeWks        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeX25        RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetDnssecParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetDnssecParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeA, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeA6, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAny, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeApl, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCds, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCert, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCname, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDname, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeDs, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEid, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeGid, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHip, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeKey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeKx, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeL32, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeL64, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeLp, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMB, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMd, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMf, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMg, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMr, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeMx, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNb, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNid, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNs, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNull, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypePtr, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypePx, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRp, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeRt, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSig, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSink, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUid, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeUri, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeWks, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeX25, RadarDNSTimeseriesGroupGetDnssecParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetDnssecParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetDnssecParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetDnssecParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetDnssecParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetDnssecAwareParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetDnssecAwareParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetDnssecAwareParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetDnssecAwareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval15m RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1h  RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1d  RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1w  RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval15m, RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1h, RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1d, RadarDNSTimeseriesGroupGetDnssecAwareParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetDnssecAwareParamsFormat string

const (
	RadarDNSTimeseriesGroupGetDnssecAwareParamsFormatJson RadarDNSTimeseriesGroupGetDnssecAwareParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsFormatCsv  RadarDNSTimeseriesGroupGetDnssecAwareParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetDnssecAwareParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecAwareParamsFormatJson, RadarDNSTimeseriesGroupGetDnssecAwareParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolUdp   RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolTcp   RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolHTTPS RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolTls   RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolUdp, RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolTcp, RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetDnssecAwareParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeA          RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeA6         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAny        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeApl        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCds        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCert       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCname      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDname      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDs         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEid        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeGid        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHip        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeKey        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeKx         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeL32        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeL64        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeLp         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMB         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMd         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMf         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMg         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMr         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMx         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNb         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNid        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNs         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNull       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypePtr        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypePx         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRp         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRt         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSig        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSink       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTa         RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUid        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUri        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeWks        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeX25        RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeA, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeA6, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAny, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeApl, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCds, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCert, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCname, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDname, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeDs, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEid, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeGid, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHip, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeKey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeKx, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeL32, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeL64, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeLp, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMB, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMd, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMf, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMg, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMr, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeMx, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNb, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNid, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNs, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNull, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypePtr, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypePx, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRp, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeRt, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSig, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSink, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUid, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeUri, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeWks, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeX25, RadarDNSTimeseriesGroupGetDnssecAwareParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetDnssecAwareParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetDnssecE2EParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetDnssecE2EParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetDnssecE2EParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetDnssecE2EParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval15m RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1h  RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1d  RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1w  RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval15m, RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1h, RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1d, RadarDNSTimeseriesGroupGetDnssecE2EParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetDnssecE2EParamsFormat string

const (
	RadarDNSTimeseriesGroupGetDnssecE2EParamsFormatJson RadarDNSTimeseriesGroupGetDnssecE2EParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsFormatCsv  RadarDNSTimeseriesGroupGetDnssecE2EParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetDnssecE2EParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecE2EParamsFormatJson, RadarDNSTimeseriesGroupGetDnssecE2EParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolUdp   RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolTcp   RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolHTTPS RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolTls   RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolUdp, RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolTcp, RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetDnssecE2EParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeA          RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeA6         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAny        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeApl        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCds        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCert       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCname      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDname      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDs         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEid        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeGid        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHip        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeKey        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeKx         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeL32        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeL64        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeLp         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMB         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMd         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMf         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMg         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMr         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMx         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNb         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNid        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNs         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNull       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypePtr        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypePx         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRp         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRt         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSig        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSink       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTa         RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUid        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUri        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeWks        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeX25        RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeA, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeA6, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAny, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeApl, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCds, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCert, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCname, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDname, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeDs, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEid, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeGid, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHip, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeKey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeKx, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeL32, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeL64, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeLp, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMB, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMd, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMf, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMg, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMr, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeMx, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNb, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNid, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNs, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNull, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypePtr, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypePx, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRp, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeRt, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSig, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSink, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUid, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeUri, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeWks, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeX25, RadarDNSTimeseriesGroupGetDnssecE2EParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetDnssecE2EParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetIPVersionParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetIPVersionParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval15m RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1h  RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1d  RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1w  RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval15m, RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1h, RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1d, RadarDNSTimeseriesGroupGetIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetIPVersionParamsFormat string

const (
	RadarDNSTimeseriesGroupGetIPVersionParamsFormatJson RadarDNSTimeseriesGroupGetIPVersionParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetIPVersionParamsFormatCsv  RadarDNSTimeseriesGroupGetIPVersionParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetIPVersionParamsFormatJson, RadarDNSTimeseriesGroupGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetIPVersionParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetIPVersionParamsProtocolUdp   RadarDNSTimeseriesGroupGetIPVersionParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetIPVersionParamsProtocolTcp   RadarDNSTimeseriesGroupGetIPVersionParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetIPVersionParamsProtocolHTTPS RadarDNSTimeseriesGroupGetIPVersionParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetIPVersionParamsProtocolTls   RadarDNSTimeseriesGroupGetIPVersionParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetIPVersionParamsProtocolUdp, RadarDNSTimeseriesGroupGetIPVersionParamsProtocolTcp, RadarDNSTimeseriesGroupGetIPVersionParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetIPVersionParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetIPVersionParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeA          RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeA6         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAny        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeApl        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCds        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCert       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCname      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDname      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDs         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEid        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeGid        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHip        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeKey        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeKx         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeL32        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeL64        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeLp         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMB         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMd         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMf         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMg         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMr         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMx         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNb         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNid        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNs         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNull       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypePtr        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypePx         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRp         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRt         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSig        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSink       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTa         RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUid        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUri        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeWks        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeX25        RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetIPVersionParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeA, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeA6, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAny, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeApl, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCds, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCert, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCname, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDname, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeDs, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEid, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeGid, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHip, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeKey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeKx, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeL32, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeL64, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeLp, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMB, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMd, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMf, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMg, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMr, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeMx, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNb, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNid, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNs, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNull, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypePtr, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypePx, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRp, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeRt, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSig, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSink, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUid, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeUri, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeWks, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeX25, RadarDNSTimeseriesGroupGetIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetMatchingAnswerParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetMatchingAnswerParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetMatchingAnswerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval15m RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1h  RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1d  RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1w  RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval15m, RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1h, RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1d, RadarDNSTimeseriesGroupGetMatchingAnswerParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormat string

const (
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormatJson RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormatCsv  RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormatJson, RadarDNSTimeseriesGroupGetMatchingAnswerParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolUdp   RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolTcp   RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolHTTPS RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolTls   RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolUdp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolTcp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetMatchingAnswerParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeA          RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeA6         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAny        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeApl        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCds        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCert       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCname      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDname      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDs         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEid        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeGid        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHip        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeKey        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeKx         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeL32        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeL64        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeLp         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMB         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMd         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMf         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMg         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMr         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMx         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNb         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNid        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNs         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNull       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypePtr        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypePx         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRp         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRt         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSig        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSink       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTa         RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUid        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUri        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeWks        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeX25        RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeA, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeA6, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAny, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeApl, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCds, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCert, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCname, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDname, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeDs, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEid, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeGid, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHip, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeKey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeKx, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeL32, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeL64, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeLp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMB, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMd, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMf, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMg, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeMx, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNb, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNid, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNs, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNull, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypePtr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypePx, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeRt, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSig, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSink, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUid, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeUri, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeWks, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeX25, RadarDNSTimeseriesGroupGetMatchingAnswerParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetMatchingAnswerParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetProtocolParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetProtocolParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetProtocolParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetProtocolParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetProtocolParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetProtocolParamsAggInterval15m RadarDNSTimeseriesGroupGetProtocolParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1h  RadarDNSTimeseriesGroupGetProtocolParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1d  RadarDNSTimeseriesGroupGetProtocolParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1w  RadarDNSTimeseriesGroupGetProtocolParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetProtocolParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetProtocolParamsAggInterval15m, RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1h, RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1d, RadarDNSTimeseriesGroupGetProtocolParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetProtocolParamsFormat string

const (
	RadarDNSTimeseriesGroupGetProtocolParamsFormatJson RadarDNSTimeseriesGroupGetProtocolParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetProtocolParamsFormatCsv  RadarDNSTimeseriesGroupGetProtocolParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetProtocolParamsFormatJson, RadarDNSTimeseriesGroupGetProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetProtocolParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeA          RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeA6         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAny        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeApl        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCds        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCert       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCname      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDname      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDs         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEid        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeGid        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHip        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeKey        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeKx         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeL32        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeL64        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeLp         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMB         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMd         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMf         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMg         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMr         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMx         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNb         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNid        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNs         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNull       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypePtr        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypePx         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRp         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRt         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSig        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSink       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTa         RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUid        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUri        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeWks        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeX25        RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetProtocolParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeA, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeA6, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAny, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeApl, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCds, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCert, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCname, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDname, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeDs, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEid, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeGid, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHip, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeKey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeKx, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeL32, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeL64, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeLp, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMB, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMd, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMf, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMg, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMr, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeMx, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNb, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNid, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNs, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNull, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypePtr, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypePx, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRp, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeRt, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSig, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSink, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUid, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeUri, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeWks, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeX25, RadarDNSTimeseriesGroupGetProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetProtocolParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetProtocolParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetQueryTypeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetQueryTypeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetQueryTypeParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval15m RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1h  RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1d  RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1w  RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval15m, RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1h, RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1d, RadarDNSTimeseriesGroupGetQueryTypeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetQueryTypeParamsFormat string

const (
	RadarDNSTimeseriesGroupGetQueryTypeParamsFormatJson RadarDNSTimeseriesGroupGetQueryTypeParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetQueryTypeParamsFormatCsv  RadarDNSTimeseriesGroupGetQueryTypeParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetQueryTypeParamsFormatJson, RadarDNSTimeseriesGroupGetQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolUdp   RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolTcp   RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolHTTPS RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolTls   RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolUdp, RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolTcp, RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetQueryTypeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetResponseCodeParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetResponseCodeParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType] `query:"queryType"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetResponseCodeParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetResponseCodeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval15m RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1h  RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1d  RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1w  RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval15m, RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1h, RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1d, RadarDNSTimeseriesGroupGetResponseCodeParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetResponseCodeParamsFormat string

const (
	RadarDNSTimeseriesGroupGetResponseCodeParamsFormatJson RadarDNSTimeseriesGroupGetResponseCodeParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetResponseCodeParamsFormatCsv  RadarDNSTimeseriesGroupGetResponseCodeParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetResponseCodeParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseCodeParamsFormatJson, RadarDNSTimeseriesGroupGetResponseCodeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolUdp   RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolTcp   RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolHTTPS RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolTls   RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetResponseCodeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolUdp, RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolTcp, RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetResponseCodeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeA          RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeA6         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAny        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeApl        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCds        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCert       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCname      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDname      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDs         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEid        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeGid        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHip        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeKey        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeKx         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeL32        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeL64        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeLp         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMB         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMd         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMf         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMg         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMr         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMx         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNb         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNid        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNs         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNull       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypePtr        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypePx         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRp         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRt         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSig        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSink       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTa         RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUid        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUri        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeWks        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeX25        RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetResponseCodeParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeA, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeA6, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAny, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeApl, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCds, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCert, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCname, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDname, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeDs, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEid, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeGid, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHip, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeKey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeKx, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeL32, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeL64, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeLp, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMB, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMd, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMf, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMg, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMr, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeMx, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNb, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNid, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNs, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNull, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypePtr, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypePx, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRp, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeRt, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSig, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSink, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUid, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeUri, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeWks, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeX25, RadarDNSTimeseriesGroupGetResponseCodeParamsQueryTypeZonemd:
		return true
	}
	return false
}

type RadarDNSTimeseriesGroupGetResponseTtlParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarDNSTimeseriesGroupGetResponseTtlParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSTimeseriesGroupGetResponseTtlParams]'s query
// parameters as `url.Values`.
func (r RadarDNSTimeseriesGroupGetResponseTtlParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval string

const (
	RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval15m RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval = "15m"
	RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1h  RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval = "1h"
	RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1d  RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval = "1d"
	RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1w  RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval = "1w"
)

func (r RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval15m, RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1h, RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1d, RadarDNSTimeseriesGroupGetResponseTtlParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarDNSTimeseriesGroupGetResponseTtlParamsFormat string

const (
	RadarDNSTimeseriesGroupGetResponseTtlParamsFormatJson RadarDNSTimeseriesGroupGetResponseTtlParamsFormat = "JSON"
	RadarDNSTimeseriesGroupGetResponseTtlParamsFormatCsv  RadarDNSTimeseriesGroupGetResponseTtlParamsFormat = "CSV"
)

func (r RadarDNSTimeseriesGroupGetResponseTtlParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseTtlParamsFormatJson, RadarDNSTimeseriesGroupGetResponseTtlParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol string

const (
	RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolUdp   RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol = "UDP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolTcp   RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol = "TCP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolHTTPS RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol = "HTTPS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolTls   RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol = "TLS"
)

func (r RadarDNSTimeseriesGroupGetResponseTtlParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolUdp, RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolTcp, RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolHTTPS, RadarDNSTimeseriesGroupGetResponseTtlParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType string

const (
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeA          RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "A"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAaaa       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "AAAA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeA6         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "A6"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAfsdb      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "AFSDB"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAny        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "ANY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeApl        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "APL"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAtma       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "ATMA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAxfr       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "AXFR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCaa        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CAA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCdnskey    RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CDNSKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCds        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CDS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCert       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CERT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCname      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CNAME"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCsync      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "CSYNC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDhcid      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DHCID"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDlv        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DLV"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDname      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DNAME"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDnskey     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DNSKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDoa        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DOA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDs         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "DS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEid        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "EID"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEui48      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "EUI48"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEui64      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "EUI64"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeGpos       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "GPOS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeGid        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "GID"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHinfo      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "HINFO"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHip        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "HIP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHTTPS      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "HTTPS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIpseckey   RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "IPSECKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIsdn       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "ISDN"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIxfr       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "IXFR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeKey        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "KEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeKx         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "KX"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeL32        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "L32"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeL64        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "L64"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeLoc        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "LOC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeLp         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "LP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMaila      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MAILA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMailb      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MAILB"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMB         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MB"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMd         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MD"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMf         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MF"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMg         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMinfo      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MINFO"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMr         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMx         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "MX"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNaptr      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NAPTR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNb         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NB"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNbstat     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NBSTAT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNid        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NID"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNimloc     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NIMLOC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNinfo      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NINFO"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNs         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsap       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NSAP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NSEC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec3      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NSEC3"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec3Param RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NSEC3PARAM"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNull       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NULL"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNxt        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "NXT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeOpenpgpkey RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "OPENPGPKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeOpt        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "OPT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypePtr        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "PTR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypePx         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "PX"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRkey       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "RKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRp         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "RP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRrsig      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "RRSIG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRt         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "RT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSig        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SIG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSink       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SINK"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSmimea     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SMIMEA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSoa        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SOA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSpf        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SPF"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSrv        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SRV"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSshfp      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SSHFP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSvcb       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "SVCB"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTa         RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTalink     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TALINK"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTkey       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTlsa       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TLSA"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTsig       RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TSIG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTxt        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "TXT"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUinfo      RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "UINFO"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUid        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "UID"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUnspec     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "UNSPEC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUri        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "URI"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeWks        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "WKS"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeX25        RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "X25"
	RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeZonemd     RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType = "ZONEMD"
)

func (r RadarDNSTimeseriesGroupGetResponseTtlParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeA, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAaaa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeA6, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAfsdb, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAny, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeApl, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAtma, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeAxfr, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCaa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCdnskey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCds, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCert, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCname, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeCsync, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDhcid, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDlv, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDname, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDnskey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDoa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeDs, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEid, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEui48, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeEui64, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeGpos, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeGid, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHinfo, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHip, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeHTTPS, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIpseckey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIsdn, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeIxfr, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeKey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeKx, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeL32, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeL64, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeLoc, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeLp, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMaila, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMailb, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMB, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMd, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMf, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMg, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMinfo, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMr, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeMx, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNaptr, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNb, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNbstat, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNid, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNimloc, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNinfo, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNs, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsap, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec3, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNsec3Param, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNull, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeNxt, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeOpenpgpkey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeOpt, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypePtr, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypePx, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRkey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRp, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRrsig, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeRt, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSig, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSink, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSmimea, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSoa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSpf, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSrv, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSshfp, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeSvcb, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTalink, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTkey, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTlsa, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTsig, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeTxt, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUinfo, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUid, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUnspec, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeUri, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeWks, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeX25, RadarDNSTimeseriesGroupGetResponseTtlParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode string

const (
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNoerror   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NOERROR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeFormerr   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "FORMERR"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeServfail  RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "SERVFAIL"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNxdomain  RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NXDOMAIN"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotimp    RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NOTIMP"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeRefused   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "REFUSED"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeYxdomain  RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "YXDOMAIN"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeYxrrset   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "YXRRSET"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNxrrset   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NXRRSET"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotauth   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NOTAUTH"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotzone   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "NOTZONE"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadsig    RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADSIG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadkey    RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADKEY"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadtime   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADTIME"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadmode   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADMODE"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadname   RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADNAME"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadalg    RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADALG"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadtrunc  RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADTRUNC"
	RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadcookie RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNoerror, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeFormerr, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeServfail, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNxdomain, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotimp, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeRefused, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeYxdomain, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeYxrrset, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNxrrset, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotauth, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeNotzone, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadsig, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadkey, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadtime, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadmode, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadname, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadalg, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadtrunc, RadarDNSTimeseriesGroupGetResponseTtlParamsResponseCodeBadcookie:
		return true
	}
	return false
}
