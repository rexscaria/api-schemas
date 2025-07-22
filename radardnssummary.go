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

// RadarDNSSummaryService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarDNSSummaryService] method instead.
type RadarDNSSummaryService struct {
	Options []option.RequestOption
}

// NewRadarDNSSummaryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarDNSSummaryService(opts ...option.RequestOption) (r *RadarDNSSummaryService) {
	r = &RadarDNSSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries by cache status.
func (r *RadarDNSSummaryService) GetCacheHit(ctx context.Context, query RadarDNSSummaryGetCacheHitParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetCacheHitResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/cache_hit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS responses by DNSSEC (DNS Security Extensions)
// support.
func (r *RadarDNSSummaryService) GetDnssec(ctx context.Context, query RadarDNSSummaryGetDnssecParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by DNSSEC (DNS Security Extensions)
// client awareness.
func (r *RadarDNSSummaryService) GetDnssecAware(ctx context.Context, query RadarDNSSummaryGetDnssecAwareParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetDnssecAwareResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec_aware"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNSSEC-validated answers by end-to-end security
// status.
func (r *RadarDNSSummaryService) GetDnssecE2E(ctx context.Context, query RadarDNSSummaryGetDnssecE2EParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetDnssecE2EResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/dnssec_e2e"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by IP version.
func (r *RadarDNSSummaryService) GetIPVersion(ctx context.Context, query RadarDNSSummaryGetIPVersionParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by matching answers.
func (r *RadarDNSSummaryService) GetMatchingAnswer(ctx context.Context, query RadarDNSSummaryGetMatchingAnswerParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetMatchingAnswerResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/matching_answer"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by DNS transport protocol.
func (r *RadarDNSSummaryService) GetProtocol(ctx context.Context, query RadarDNSSummaryGetProtocolParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by type.
func (r *RadarDNSSummaryService) GetQueryType(ctx context.Context, query RadarDNSSummaryGetQueryTypeParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by response code.
func (r *RadarDNSSummaryService) GetResponseCode(ctx context.Context, query RadarDNSSummaryGetResponseCodeParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetResponseCodeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/response_code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries by minimum response TTL.
func (r *RadarDNSSummaryService) GetResponseTtl(ctx context.Context, query RadarDNSSummaryGetResponseTtlParams, opts ...option.RequestOption) (res *RadarDNSSummaryGetResponseTtlResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/dns/summary/response_ttl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarDNSSummaryGetCacheHitResponse struct {
	Result  RadarDNSSummaryGetCacheHitResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarDNSSummaryGetCacheHitResponseJSON   `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseJSON contains the JSON metadata for the struct
// [RadarDNSSummaryGetCacheHitResponse]
type radarDNSSummaryGetCacheHitResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResult struct {
	Meta     RadarDNSSummaryGetCacheHitResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetCacheHitResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetCacheHitResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetCacheHitResponseResult]
type radarDNSSummaryGetCacheHitResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetCacheHitResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetCacheHitResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetCacheHitResponseResultMeta]
type radarDNSSummaryGetCacheHitResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetCacheHitResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetCacheHitResponseResultMetaDateRange]
type radarDNSSummaryGetCacheHitResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitResponseResultSummary0 struct {
	Negative string                                               `json:"NEGATIVE,required"`
	Positive string                                               `json:"POSITIVE,required"`
	JSON     radarDNSSummaryGetCacheHitResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetCacheHitResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetCacheHitResponseResultSummary0]
type radarDNSSummaryGetCacheHitResponseResultSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetCacheHitResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetCacheHitResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponse struct {
	Result  RadarDNSSummaryGetDnssecResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarDNSSummaryGetDnssecResponseJSON   `json:"-"`
}

// radarDNSSummaryGetDnssecResponseJSON contains the JSON metadata for the struct
// [RadarDNSSummaryGetDnssecResponse]
type radarDNSSummaryGetDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResult struct {
	Meta     RadarDNSSummaryGetDnssecResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetDnssecResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetDnssecResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetDnssecResponseResult]
type radarDNSSummaryGetDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetDnssecResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetDnssecResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetDnssecResponseResultMeta]
type radarDNSSummaryGetDnssecResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetDnssecResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetDnssecResponseResultMetaDateRange]
type radarDNSSummaryGetDnssecResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecResponseResultSummary0 struct {
	Insecure string                                             `json:"INSECURE,required"`
	Invalid  string                                             `json:"INVALID,required"`
	Other    string                                             `json:"OTHER,required"`
	Secure   string                                             `json:"SECURE,required"`
	JSON     radarDNSSummaryGetDnssecResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetDnssecResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetDnssecResponseResultSummary0]
type radarDNSSummaryGetDnssecResponseResultSummary0JSON struct {
	Insecure    apijson.Field
	Invalid     apijson.Field
	Other       apijson.Field
	Secure      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponse struct {
	Result  RadarDNSSummaryGetDnssecAwareResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarDNSSummaryGetDnssecAwareResponseJSON   `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetDnssecAwareResponse]
type radarDNSSummaryGetDnssecAwareResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResult struct {
	Meta     RadarDNSSummaryGetDnssecAwareResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetDnssecAwareResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetDnssecAwareResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetDnssecAwareResponseResult]
type radarDNSSummaryGetDnssecAwareResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetDnssecAwareResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetDnssecAwareResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetDnssecAwareResponseResultMeta]
type radarDNSSummaryGetDnssecAwareResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetDnssecAwareResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetDnssecAwareResponseResultMetaDateRange]
type radarDNSSummaryGetDnssecAwareResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecAwareResponseResultSummary0 struct {
	NotSupported string                                                  `json:"NOT_SUPPORTED,required"`
	Supported    string                                                  `json:"SUPPORTED,required"`
	JSON         radarDNSSummaryGetDnssecAwareResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetDnssecAwareResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetDnssecAwareResponseResultSummary0]
type radarDNSSummaryGetDnssecAwareResponseResultSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecAwareResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecAwareResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponse struct {
	Result  RadarDNSSummaryGetDnssecE2EResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarDNSSummaryGetDnssecE2EResponseJSON   `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetDnssecE2EResponse]
type radarDNSSummaryGetDnssecE2EResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResult struct {
	Meta     RadarDNSSummaryGetDnssecE2EResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetDnssecE2EResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetDnssecE2EResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetDnssecE2EResponseResult]
type radarDNSSummaryGetDnssecE2EResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetDnssecE2EResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetDnssecE2EResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetDnssecE2EResponseResultMeta]
type radarDNSSummaryGetDnssecE2EResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetDnssecE2EResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetDnssecE2EResponseResultMetaDateRange]
type radarDNSSummaryGetDnssecE2EResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetDnssecE2EResponseResultSummary0 struct {
	Negative string                                                `json:"NEGATIVE,required"`
	Positive string                                                `json:"POSITIVE,required"`
	JSON     radarDNSSummaryGetDnssecE2EResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetDnssecE2EResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetDnssecE2EResponseResultSummary0]
type radarDNSSummaryGetDnssecE2EResponseResultSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetDnssecE2EResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetDnssecE2EResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponse struct {
	Result  RadarDNSSummaryGetIPVersionResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarDNSSummaryGetIPVersionResponseJSON   `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetIPVersionResponse]
type radarDNSSummaryGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResult struct {
	Meta     RadarDNSSummaryGetIPVersionResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetIPVersionResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetIPVersionResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetIPVersionResponseResult]
type radarDNSSummaryGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetIPVersionResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetIPVersionResponseResultMeta]
type radarDNSSummaryGetIPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetIPVersionResponseResultMetaDateRange]
type radarDNSSummaryGetIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetIPVersionResponseResultSummary0 struct {
	IPv4 string                                                `json:"IPv4,required"`
	IPv6 string                                                `json:"IPv6,required"`
	JSON radarDNSSummaryGetIPVersionResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetIPVersionResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetIPVersionResponseResultSummary0]
type radarDNSSummaryGetIPVersionResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetIPVersionResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetIPVersionResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponse struct {
	Result  RadarDNSSummaryGetMatchingAnswerResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarDNSSummaryGetMatchingAnswerResponseJSON   `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetMatchingAnswerResponse]
type radarDNSSummaryGetMatchingAnswerResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResult struct {
	Meta     RadarDNSSummaryGetMatchingAnswerResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetMatchingAnswerResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetMatchingAnswerResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultJSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetMatchingAnswerResponseResult]
type radarDNSSummaryGetMatchingAnswerResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetMatchingAnswerResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	Normalization  string                                                           `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetMatchingAnswerResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetMatchingAnswerResponseResultMeta]
type radarDNSSummaryGetMatchingAnswerResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetMatchingAnswerResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetMatchingAnswerResponseResultMetaDateRange]
type radarDNSSummaryGetMatchingAnswerResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetMatchingAnswerResponseResultSummary0 struct {
	Negative string                                                     `json:"NEGATIVE,required"`
	Positive string                                                     `json:"POSITIVE,required"`
	JSON     radarDNSSummaryGetMatchingAnswerResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetMatchingAnswerResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetMatchingAnswerResponseResultSummary0]
type radarDNSSummaryGetMatchingAnswerResponseResultSummary0JSON struct {
	Negative    apijson.Field
	Positive    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetMatchingAnswerResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetMatchingAnswerResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponse struct {
	Result  RadarDNSSummaryGetProtocolResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarDNSSummaryGetProtocolResponseJSON   `json:"-"`
}

// radarDNSSummaryGetProtocolResponseJSON contains the JSON metadata for the struct
// [RadarDNSSummaryGetProtocolResponse]
type radarDNSSummaryGetProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResult struct {
	Meta     RadarDNSSummaryGetProtocolResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetProtocolResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetProtocolResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetProtocolResponseResult]
type radarDNSSummaryGetProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetProtocolResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetProtocolResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetProtocolResponseResultMeta]
type radarDNSSummaryGetProtocolResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetProtocolResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetProtocolResponseResultMetaDateRange]
type radarDNSSummaryGetProtocolResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetProtocolResponseResultSummary0 struct {
	HTTPS string                                               `json:"HTTPS,required"`
	Tcp   string                                               `json:"TCP,required"`
	Tls   string                                               `json:"TLS,required"`
	Udp   string                                               `json:"UDP,required"`
	JSON  radarDNSSummaryGetProtocolResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetProtocolResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetProtocolResponseResultSummary0]
type radarDNSSummaryGetProtocolResponseResultSummary0JSON struct {
	HTTPS       apijson.Field
	Tcp         apijson.Field
	Tls         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetProtocolResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetProtocolResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponse struct {
	Result  RadarDNSSummaryGetQueryTypeResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarDNSSummaryGetQueryTypeResponseJSON   `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetQueryTypeResponse]
type radarDNSSummaryGetQueryTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponseResult struct {
	Meta     RadarDNSSummaryGetQueryTypeResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                             `json:"summary_0,required"`
	JSON     radarDNSSummaryGetQueryTypeResponseResultJSON `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseResultJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetQueryTypeResponseResult]
type radarDNSSummaryGetQueryTypeResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetQueryTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetQueryTypeResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetQueryTypeResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetQueryTypeResponseResultMeta]
type radarDNSSummaryGetQueryTypeResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetQueryTypeResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetQueryTypeResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetQueryTypeResponseResultMetaDateRange]
type radarDNSSummaryGetQueryTypeResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetQueryTypeResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponse struct {
	Result  RadarDNSSummaryGetResponseCodeResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarDNSSummaryGetResponseCodeResponseJSON   `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetResponseCodeResponse]
type radarDNSSummaryGetResponseCodeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseCodeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponseResult struct {
	Meta     RadarDNSSummaryGetResponseCodeResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                `json:"summary_0,required"`
	JSON     radarDNSSummaryGetResponseCodeResponseResultJSON `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseResultJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetResponseCodeResponseResult]
type radarDNSSummaryGetResponseCodeResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseCodeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetResponseCodeResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetResponseCodeResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetResponseCodeResponseResultMeta]
type radarDNSSummaryGetResponseCodeResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseCodeResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetResponseCodeResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetResponseCodeResponseResultMetaDateRange]
type radarDNSSummaryGetResponseCodeResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseCodeResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseCodeResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponse struct {
	Result  RadarDNSSummaryGetResponseTtlResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarDNSSummaryGetResponseTtlResponseJSON   `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseJSON contains the JSON metadata for the
// struct [RadarDNSSummaryGetResponseTtlResponse]
type radarDNSSummaryGetResponseTtlResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseTtlResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResult struct {
	Meta     RadarDNSSummaryGetResponseTtlResponseResultMeta     `json:"meta,required"`
	Summary0 RadarDNSSummaryGetResponseTtlResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarDNSSummaryGetResponseTtlResponseResultJSON     `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultJSON contains the JSON metadata for
// the struct [RadarDNSSummaryGetResponseTtlResponseResult]
type radarDNSSummaryGetResponseTtlResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseTtlResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResultMeta struct {
	DateRange      []RadarDNSSummaryGetResponseTtlResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarDNSSummaryGetResponseTtlResponseResultMetaJSON           `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarDNSSummaryGetResponseTtlResponseResultMeta]
type radarDNSSummaryGetResponseTtlResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseTtlResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarDNSSummaryGetResponseTtlResponseResultMetaDateRangeJSON `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarDNSSummaryGetResponseTtlResponseResultMetaDateRange]
type radarDNSSummaryGetResponseTtlResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseTtlResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfo struct {
	Annotations []RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfo]
type radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotation]
type radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetResponseTtlResponseResultSummary0 struct {
	Gt15mLte1h string                                                  `json:"gt_15m_lte_1h,required"`
	Gt1dLte1w  string                                                  `json:"gt_1d_lte_1w,required"`
	Gt1hLte1d  string                                                  `json:"gt_1h_lte_1d,required"`
	Gt1mLte5m  string                                                  `json:"gt_1m_lte_5m,required"`
	Gt1w       string                                                  `json:"gt_1w,required"`
	Gt5mLte15m string                                                  `json:"gt_5m_lte_15m,required"`
	Lte1m      string                                                  `json:"lte_1m,required"`
	JSON       radarDNSSummaryGetResponseTtlResponseResultSummary0JSON `json:"-"`
}

// radarDNSSummaryGetResponseTtlResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarDNSSummaryGetResponseTtlResponseResultSummary0]
type radarDNSSummaryGetResponseTtlResponseResultSummary0JSON struct {
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

func (r *RadarDNSSummaryGetResponseTtlResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarDNSSummaryGetResponseTtlResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarDNSSummaryGetCacheHitParams struct {
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
	Format param.Field[RadarDNSSummaryGetCacheHitParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetCacheHitParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetCacheHitParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetCacheHitParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetCacheHitParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetCacheHitParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetCacheHitParamsFormat string

const (
	RadarDNSSummaryGetCacheHitParamsFormatJson RadarDNSSummaryGetCacheHitParamsFormat = "JSON"
	RadarDNSSummaryGetCacheHitParamsFormatCsv  RadarDNSSummaryGetCacheHitParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetCacheHitParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetCacheHitParamsFormatJson, RadarDNSSummaryGetCacheHitParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetCacheHitParamsProtocol string

const (
	RadarDNSSummaryGetCacheHitParamsProtocolUdp   RadarDNSSummaryGetCacheHitParamsProtocol = "UDP"
	RadarDNSSummaryGetCacheHitParamsProtocolTcp   RadarDNSSummaryGetCacheHitParamsProtocol = "TCP"
	RadarDNSSummaryGetCacheHitParamsProtocolHTTPS RadarDNSSummaryGetCacheHitParamsProtocol = "HTTPS"
	RadarDNSSummaryGetCacheHitParamsProtocolTls   RadarDNSSummaryGetCacheHitParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetCacheHitParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetCacheHitParamsProtocolUdp, RadarDNSSummaryGetCacheHitParamsProtocolTcp, RadarDNSSummaryGetCacheHitParamsProtocolHTTPS, RadarDNSSummaryGetCacheHitParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetCacheHitParamsQueryType string

const (
	RadarDNSSummaryGetCacheHitParamsQueryTypeA          RadarDNSSummaryGetCacheHitParamsQueryType = "A"
	RadarDNSSummaryGetCacheHitParamsQueryTypeAaaa       RadarDNSSummaryGetCacheHitParamsQueryType = "AAAA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeA6         RadarDNSSummaryGetCacheHitParamsQueryType = "A6"
	RadarDNSSummaryGetCacheHitParamsQueryTypeAfsdb      RadarDNSSummaryGetCacheHitParamsQueryType = "AFSDB"
	RadarDNSSummaryGetCacheHitParamsQueryTypeAny        RadarDNSSummaryGetCacheHitParamsQueryType = "ANY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeApl        RadarDNSSummaryGetCacheHitParamsQueryType = "APL"
	RadarDNSSummaryGetCacheHitParamsQueryTypeAtma       RadarDNSSummaryGetCacheHitParamsQueryType = "ATMA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeAxfr       RadarDNSSummaryGetCacheHitParamsQueryType = "AXFR"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCaa        RadarDNSSummaryGetCacheHitParamsQueryType = "CAA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCdnskey    RadarDNSSummaryGetCacheHitParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCds        RadarDNSSummaryGetCacheHitParamsQueryType = "CDS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCert       RadarDNSSummaryGetCacheHitParamsQueryType = "CERT"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCname      RadarDNSSummaryGetCacheHitParamsQueryType = "CNAME"
	RadarDNSSummaryGetCacheHitParamsQueryTypeCsync      RadarDNSSummaryGetCacheHitParamsQueryType = "CSYNC"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDhcid      RadarDNSSummaryGetCacheHitParamsQueryType = "DHCID"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDlv        RadarDNSSummaryGetCacheHitParamsQueryType = "DLV"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDname      RadarDNSSummaryGetCacheHitParamsQueryType = "DNAME"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDnskey     RadarDNSSummaryGetCacheHitParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDoa        RadarDNSSummaryGetCacheHitParamsQueryType = "DOA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeDs         RadarDNSSummaryGetCacheHitParamsQueryType = "DS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeEid        RadarDNSSummaryGetCacheHitParamsQueryType = "EID"
	RadarDNSSummaryGetCacheHitParamsQueryTypeEui48      RadarDNSSummaryGetCacheHitParamsQueryType = "EUI48"
	RadarDNSSummaryGetCacheHitParamsQueryTypeEui64      RadarDNSSummaryGetCacheHitParamsQueryType = "EUI64"
	RadarDNSSummaryGetCacheHitParamsQueryTypeGpos       RadarDNSSummaryGetCacheHitParamsQueryType = "GPOS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeGid        RadarDNSSummaryGetCacheHitParamsQueryType = "GID"
	RadarDNSSummaryGetCacheHitParamsQueryTypeHinfo      RadarDNSSummaryGetCacheHitParamsQueryType = "HINFO"
	RadarDNSSummaryGetCacheHitParamsQueryTypeHip        RadarDNSSummaryGetCacheHitParamsQueryType = "HIP"
	RadarDNSSummaryGetCacheHitParamsQueryTypeHTTPS      RadarDNSSummaryGetCacheHitParamsQueryType = "HTTPS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeIpseckey   RadarDNSSummaryGetCacheHitParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeIsdn       RadarDNSSummaryGetCacheHitParamsQueryType = "ISDN"
	RadarDNSSummaryGetCacheHitParamsQueryTypeIxfr       RadarDNSSummaryGetCacheHitParamsQueryType = "IXFR"
	RadarDNSSummaryGetCacheHitParamsQueryTypeKey        RadarDNSSummaryGetCacheHitParamsQueryType = "KEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeKx         RadarDNSSummaryGetCacheHitParamsQueryType = "KX"
	RadarDNSSummaryGetCacheHitParamsQueryTypeL32        RadarDNSSummaryGetCacheHitParamsQueryType = "L32"
	RadarDNSSummaryGetCacheHitParamsQueryTypeL64        RadarDNSSummaryGetCacheHitParamsQueryType = "L64"
	RadarDNSSummaryGetCacheHitParamsQueryTypeLoc        RadarDNSSummaryGetCacheHitParamsQueryType = "LOC"
	RadarDNSSummaryGetCacheHitParamsQueryTypeLp         RadarDNSSummaryGetCacheHitParamsQueryType = "LP"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMaila      RadarDNSSummaryGetCacheHitParamsQueryType = "MAILA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMailb      RadarDNSSummaryGetCacheHitParamsQueryType = "MAILB"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMB         RadarDNSSummaryGetCacheHitParamsQueryType = "MB"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMd         RadarDNSSummaryGetCacheHitParamsQueryType = "MD"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMf         RadarDNSSummaryGetCacheHitParamsQueryType = "MF"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMg         RadarDNSSummaryGetCacheHitParamsQueryType = "MG"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMinfo      RadarDNSSummaryGetCacheHitParamsQueryType = "MINFO"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMr         RadarDNSSummaryGetCacheHitParamsQueryType = "MR"
	RadarDNSSummaryGetCacheHitParamsQueryTypeMx         RadarDNSSummaryGetCacheHitParamsQueryType = "MX"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNaptr      RadarDNSSummaryGetCacheHitParamsQueryType = "NAPTR"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNb         RadarDNSSummaryGetCacheHitParamsQueryType = "NB"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNbstat     RadarDNSSummaryGetCacheHitParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNid        RadarDNSSummaryGetCacheHitParamsQueryType = "NID"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNimloc     RadarDNSSummaryGetCacheHitParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNinfo      RadarDNSSummaryGetCacheHitParamsQueryType = "NINFO"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNs         RadarDNSSummaryGetCacheHitParamsQueryType = "NS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNsap       RadarDNSSummaryGetCacheHitParamsQueryType = "NSAP"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNsec       RadarDNSSummaryGetCacheHitParamsQueryType = "NSEC"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNsec3      RadarDNSSummaryGetCacheHitParamsQueryType = "NSEC3"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNsec3Param RadarDNSSummaryGetCacheHitParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNull       RadarDNSSummaryGetCacheHitParamsQueryType = "NULL"
	RadarDNSSummaryGetCacheHitParamsQueryTypeNxt        RadarDNSSummaryGetCacheHitParamsQueryType = "NXT"
	RadarDNSSummaryGetCacheHitParamsQueryTypeOpenpgpkey RadarDNSSummaryGetCacheHitParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeOpt        RadarDNSSummaryGetCacheHitParamsQueryType = "OPT"
	RadarDNSSummaryGetCacheHitParamsQueryTypePtr        RadarDNSSummaryGetCacheHitParamsQueryType = "PTR"
	RadarDNSSummaryGetCacheHitParamsQueryTypePx         RadarDNSSummaryGetCacheHitParamsQueryType = "PX"
	RadarDNSSummaryGetCacheHitParamsQueryTypeRkey       RadarDNSSummaryGetCacheHitParamsQueryType = "RKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeRp         RadarDNSSummaryGetCacheHitParamsQueryType = "RP"
	RadarDNSSummaryGetCacheHitParamsQueryTypeRrsig      RadarDNSSummaryGetCacheHitParamsQueryType = "RRSIG"
	RadarDNSSummaryGetCacheHitParamsQueryTypeRt         RadarDNSSummaryGetCacheHitParamsQueryType = "RT"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSig        RadarDNSSummaryGetCacheHitParamsQueryType = "SIG"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSink       RadarDNSSummaryGetCacheHitParamsQueryType = "SINK"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSmimea     RadarDNSSummaryGetCacheHitParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSoa        RadarDNSSummaryGetCacheHitParamsQueryType = "SOA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSpf        RadarDNSSummaryGetCacheHitParamsQueryType = "SPF"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSrv        RadarDNSSummaryGetCacheHitParamsQueryType = "SRV"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSshfp      RadarDNSSummaryGetCacheHitParamsQueryType = "SSHFP"
	RadarDNSSummaryGetCacheHitParamsQueryTypeSvcb       RadarDNSSummaryGetCacheHitParamsQueryType = "SVCB"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTa         RadarDNSSummaryGetCacheHitParamsQueryType = "TA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTalink     RadarDNSSummaryGetCacheHitParamsQueryType = "TALINK"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTkey       RadarDNSSummaryGetCacheHitParamsQueryType = "TKEY"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTlsa       RadarDNSSummaryGetCacheHitParamsQueryType = "TLSA"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTsig       RadarDNSSummaryGetCacheHitParamsQueryType = "TSIG"
	RadarDNSSummaryGetCacheHitParamsQueryTypeTxt        RadarDNSSummaryGetCacheHitParamsQueryType = "TXT"
	RadarDNSSummaryGetCacheHitParamsQueryTypeUinfo      RadarDNSSummaryGetCacheHitParamsQueryType = "UINFO"
	RadarDNSSummaryGetCacheHitParamsQueryTypeUid        RadarDNSSummaryGetCacheHitParamsQueryType = "UID"
	RadarDNSSummaryGetCacheHitParamsQueryTypeUnspec     RadarDNSSummaryGetCacheHitParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetCacheHitParamsQueryTypeUri        RadarDNSSummaryGetCacheHitParamsQueryType = "URI"
	RadarDNSSummaryGetCacheHitParamsQueryTypeWks        RadarDNSSummaryGetCacheHitParamsQueryType = "WKS"
	RadarDNSSummaryGetCacheHitParamsQueryTypeX25        RadarDNSSummaryGetCacheHitParamsQueryType = "X25"
	RadarDNSSummaryGetCacheHitParamsQueryTypeZonemd     RadarDNSSummaryGetCacheHitParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetCacheHitParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetCacheHitParamsQueryTypeA, RadarDNSSummaryGetCacheHitParamsQueryTypeAaaa, RadarDNSSummaryGetCacheHitParamsQueryTypeA6, RadarDNSSummaryGetCacheHitParamsQueryTypeAfsdb, RadarDNSSummaryGetCacheHitParamsQueryTypeAny, RadarDNSSummaryGetCacheHitParamsQueryTypeApl, RadarDNSSummaryGetCacheHitParamsQueryTypeAtma, RadarDNSSummaryGetCacheHitParamsQueryTypeAxfr, RadarDNSSummaryGetCacheHitParamsQueryTypeCaa, RadarDNSSummaryGetCacheHitParamsQueryTypeCdnskey, RadarDNSSummaryGetCacheHitParamsQueryTypeCds, RadarDNSSummaryGetCacheHitParamsQueryTypeCert, RadarDNSSummaryGetCacheHitParamsQueryTypeCname, RadarDNSSummaryGetCacheHitParamsQueryTypeCsync, RadarDNSSummaryGetCacheHitParamsQueryTypeDhcid, RadarDNSSummaryGetCacheHitParamsQueryTypeDlv, RadarDNSSummaryGetCacheHitParamsQueryTypeDname, RadarDNSSummaryGetCacheHitParamsQueryTypeDnskey, RadarDNSSummaryGetCacheHitParamsQueryTypeDoa, RadarDNSSummaryGetCacheHitParamsQueryTypeDs, RadarDNSSummaryGetCacheHitParamsQueryTypeEid, RadarDNSSummaryGetCacheHitParamsQueryTypeEui48, RadarDNSSummaryGetCacheHitParamsQueryTypeEui64, RadarDNSSummaryGetCacheHitParamsQueryTypeGpos, RadarDNSSummaryGetCacheHitParamsQueryTypeGid, RadarDNSSummaryGetCacheHitParamsQueryTypeHinfo, RadarDNSSummaryGetCacheHitParamsQueryTypeHip, RadarDNSSummaryGetCacheHitParamsQueryTypeHTTPS, RadarDNSSummaryGetCacheHitParamsQueryTypeIpseckey, RadarDNSSummaryGetCacheHitParamsQueryTypeIsdn, RadarDNSSummaryGetCacheHitParamsQueryTypeIxfr, RadarDNSSummaryGetCacheHitParamsQueryTypeKey, RadarDNSSummaryGetCacheHitParamsQueryTypeKx, RadarDNSSummaryGetCacheHitParamsQueryTypeL32, RadarDNSSummaryGetCacheHitParamsQueryTypeL64, RadarDNSSummaryGetCacheHitParamsQueryTypeLoc, RadarDNSSummaryGetCacheHitParamsQueryTypeLp, RadarDNSSummaryGetCacheHitParamsQueryTypeMaila, RadarDNSSummaryGetCacheHitParamsQueryTypeMailb, RadarDNSSummaryGetCacheHitParamsQueryTypeMB, RadarDNSSummaryGetCacheHitParamsQueryTypeMd, RadarDNSSummaryGetCacheHitParamsQueryTypeMf, RadarDNSSummaryGetCacheHitParamsQueryTypeMg, RadarDNSSummaryGetCacheHitParamsQueryTypeMinfo, RadarDNSSummaryGetCacheHitParamsQueryTypeMr, RadarDNSSummaryGetCacheHitParamsQueryTypeMx, RadarDNSSummaryGetCacheHitParamsQueryTypeNaptr, RadarDNSSummaryGetCacheHitParamsQueryTypeNb, RadarDNSSummaryGetCacheHitParamsQueryTypeNbstat, RadarDNSSummaryGetCacheHitParamsQueryTypeNid, RadarDNSSummaryGetCacheHitParamsQueryTypeNimloc, RadarDNSSummaryGetCacheHitParamsQueryTypeNinfo, RadarDNSSummaryGetCacheHitParamsQueryTypeNs, RadarDNSSummaryGetCacheHitParamsQueryTypeNsap, RadarDNSSummaryGetCacheHitParamsQueryTypeNsec, RadarDNSSummaryGetCacheHitParamsQueryTypeNsec3, RadarDNSSummaryGetCacheHitParamsQueryTypeNsec3Param, RadarDNSSummaryGetCacheHitParamsQueryTypeNull, RadarDNSSummaryGetCacheHitParamsQueryTypeNxt, RadarDNSSummaryGetCacheHitParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetCacheHitParamsQueryTypeOpt, RadarDNSSummaryGetCacheHitParamsQueryTypePtr, RadarDNSSummaryGetCacheHitParamsQueryTypePx, RadarDNSSummaryGetCacheHitParamsQueryTypeRkey, RadarDNSSummaryGetCacheHitParamsQueryTypeRp, RadarDNSSummaryGetCacheHitParamsQueryTypeRrsig, RadarDNSSummaryGetCacheHitParamsQueryTypeRt, RadarDNSSummaryGetCacheHitParamsQueryTypeSig, RadarDNSSummaryGetCacheHitParamsQueryTypeSink, RadarDNSSummaryGetCacheHitParamsQueryTypeSmimea, RadarDNSSummaryGetCacheHitParamsQueryTypeSoa, RadarDNSSummaryGetCacheHitParamsQueryTypeSpf, RadarDNSSummaryGetCacheHitParamsQueryTypeSrv, RadarDNSSummaryGetCacheHitParamsQueryTypeSshfp, RadarDNSSummaryGetCacheHitParamsQueryTypeSvcb, RadarDNSSummaryGetCacheHitParamsQueryTypeTa, RadarDNSSummaryGetCacheHitParamsQueryTypeTalink, RadarDNSSummaryGetCacheHitParamsQueryTypeTkey, RadarDNSSummaryGetCacheHitParamsQueryTypeTlsa, RadarDNSSummaryGetCacheHitParamsQueryTypeTsig, RadarDNSSummaryGetCacheHitParamsQueryTypeTxt, RadarDNSSummaryGetCacheHitParamsQueryTypeUinfo, RadarDNSSummaryGetCacheHitParamsQueryTypeUid, RadarDNSSummaryGetCacheHitParamsQueryTypeUnspec, RadarDNSSummaryGetCacheHitParamsQueryTypeUri, RadarDNSSummaryGetCacheHitParamsQueryTypeWks, RadarDNSSummaryGetCacheHitParamsQueryTypeX25, RadarDNSSummaryGetCacheHitParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetCacheHitParamsResponseCode string

const (
	RadarDNSSummaryGetCacheHitParamsResponseCodeNoerror   RadarDNSSummaryGetCacheHitParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetCacheHitParamsResponseCodeFormerr   RadarDNSSummaryGetCacheHitParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetCacheHitParamsResponseCodeServfail  RadarDNSSummaryGetCacheHitParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetCacheHitParamsResponseCodeNxdomain  RadarDNSSummaryGetCacheHitParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetCacheHitParamsResponseCodeNotimp    RadarDNSSummaryGetCacheHitParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetCacheHitParamsResponseCodeRefused   RadarDNSSummaryGetCacheHitParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetCacheHitParamsResponseCodeYxdomain  RadarDNSSummaryGetCacheHitParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetCacheHitParamsResponseCodeYxrrset   RadarDNSSummaryGetCacheHitParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetCacheHitParamsResponseCodeNxrrset   RadarDNSSummaryGetCacheHitParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetCacheHitParamsResponseCodeNotauth   RadarDNSSummaryGetCacheHitParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetCacheHitParamsResponseCodeNotzone   RadarDNSSummaryGetCacheHitParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadsig    RadarDNSSummaryGetCacheHitParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadkey    RadarDNSSummaryGetCacheHitParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadtime   RadarDNSSummaryGetCacheHitParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadmode   RadarDNSSummaryGetCacheHitParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadname   RadarDNSSummaryGetCacheHitParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadalg    RadarDNSSummaryGetCacheHitParamsResponseCode = "BADALG"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadtrunc  RadarDNSSummaryGetCacheHitParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetCacheHitParamsResponseCodeBadcookie RadarDNSSummaryGetCacheHitParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetCacheHitParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetCacheHitParamsResponseCodeNoerror, RadarDNSSummaryGetCacheHitParamsResponseCodeFormerr, RadarDNSSummaryGetCacheHitParamsResponseCodeServfail, RadarDNSSummaryGetCacheHitParamsResponseCodeNxdomain, RadarDNSSummaryGetCacheHitParamsResponseCodeNotimp, RadarDNSSummaryGetCacheHitParamsResponseCodeRefused, RadarDNSSummaryGetCacheHitParamsResponseCodeYxdomain, RadarDNSSummaryGetCacheHitParamsResponseCodeYxrrset, RadarDNSSummaryGetCacheHitParamsResponseCodeNxrrset, RadarDNSSummaryGetCacheHitParamsResponseCodeNotauth, RadarDNSSummaryGetCacheHitParamsResponseCodeNotzone, RadarDNSSummaryGetCacheHitParamsResponseCodeBadsig, RadarDNSSummaryGetCacheHitParamsResponseCodeBadkey, RadarDNSSummaryGetCacheHitParamsResponseCodeBadtime, RadarDNSSummaryGetCacheHitParamsResponseCodeBadmode, RadarDNSSummaryGetCacheHitParamsResponseCodeBadname, RadarDNSSummaryGetCacheHitParamsResponseCodeBadalg, RadarDNSSummaryGetCacheHitParamsResponseCodeBadtrunc, RadarDNSSummaryGetCacheHitParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetDnssecParams struct {
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
	Format param.Field[RadarDNSSummaryGetDnssecParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetDnssecParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetDnssecParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetDnssecParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetDnssecParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetDnssecParamsFormat string

const (
	RadarDNSSummaryGetDnssecParamsFormatJson RadarDNSSummaryGetDnssecParamsFormat = "JSON"
	RadarDNSSummaryGetDnssecParamsFormatCsv  RadarDNSSummaryGetDnssecParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecParamsFormatJson, RadarDNSSummaryGetDnssecParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetDnssecParamsProtocol string

const (
	RadarDNSSummaryGetDnssecParamsProtocolUdp   RadarDNSSummaryGetDnssecParamsProtocol = "UDP"
	RadarDNSSummaryGetDnssecParamsProtocolTcp   RadarDNSSummaryGetDnssecParamsProtocol = "TCP"
	RadarDNSSummaryGetDnssecParamsProtocolHTTPS RadarDNSSummaryGetDnssecParamsProtocol = "HTTPS"
	RadarDNSSummaryGetDnssecParamsProtocolTls   RadarDNSSummaryGetDnssecParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetDnssecParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecParamsProtocolUdp, RadarDNSSummaryGetDnssecParamsProtocolTcp, RadarDNSSummaryGetDnssecParamsProtocolHTTPS, RadarDNSSummaryGetDnssecParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetDnssecParamsQueryType string

const (
	RadarDNSSummaryGetDnssecParamsQueryTypeA          RadarDNSSummaryGetDnssecParamsQueryType = "A"
	RadarDNSSummaryGetDnssecParamsQueryTypeAaaa       RadarDNSSummaryGetDnssecParamsQueryType = "AAAA"
	RadarDNSSummaryGetDnssecParamsQueryTypeA6         RadarDNSSummaryGetDnssecParamsQueryType = "A6"
	RadarDNSSummaryGetDnssecParamsQueryTypeAfsdb      RadarDNSSummaryGetDnssecParamsQueryType = "AFSDB"
	RadarDNSSummaryGetDnssecParamsQueryTypeAny        RadarDNSSummaryGetDnssecParamsQueryType = "ANY"
	RadarDNSSummaryGetDnssecParamsQueryTypeApl        RadarDNSSummaryGetDnssecParamsQueryType = "APL"
	RadarDNSSummaryGetDnssecParamsQueryTypeAtma       RadarDNSSummaryGetDnssecParamsQueryType = "ATMA"
	RadarDNSSummaryGetDnssecParamsQueryTypeAxfr       RadarDNSSummaryGetDnssecParamsQueryType = "AXFR"
	RadarDNSSummaryGetDnssecParamsQueryTypeCaa        RadarDNSSummaryGetDnssecParamsQueryType = "CAA"
	RadarDNSSummaryGetDnssecParamsQueryTypeCdnskey    RadarDNSSummaryGetDnssecParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeCds        RadarDNSSummaryGetDnssecParamsQueryType = "CDS"
	RadarDNSSummaryGetDnssecParamsQueryTypeCert       RadarDNSSummaryGetDnssecParamsQueryType = "CERT"
	RadarDNSSummaryGetDnssecParamsQueryTypeCname      RadarDNSSummaryGetDnssecParamsQueryType = "CNAME"
	RadarDNSSummaryGetDnssecParamsQueryTypeCsync      RadarDNSSummaryGetDnssecParamsQueryType = "CSYNC"
	RadarDNSSummaryGetDnssecParamsQueryTypeDhcid      RadarDNSSummaryGetDnssecParamsQueryType = "DHCID"
	RadarDNSSummaryGetDnssecParamsQueryTypeDlv        RadarDNSSummaryGetDnssecParamsQueryType = "DLV"
	RadarDNSSummaryGetDnssecParamsQueryTypeDname      RadarDNSSummaryGetDnssecParamsQueryType = "DNAME"
	RadarDNSSummaryGetDnssecParamsQueryTypeDnskey     RadarDNSSummaryGetDnssecParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeDoa        RadarDNSSummaryGetDnssecParamsQueryType = "DOA"
	RadarDNSSummaryGetDnssecParamsQueryTypeDs         RadarDNSSummaryGetDnssecParamsQueryType = "DS"
	RadarDNSSummaryGetDnssecParamsQueryTypeEid        RadarDNSSummaryGetDnssecParamsQueryType = "EID"
	RadarDNSSummaryGetDnssecParamsQueryTypeEui48      RadarDNSSummaryGetDnssecParamsQueryType = "EUI48"
	RadarDNSSummaryGetDnssecParamsQueryTypeEui64      RadarDNSSummaryGetDnssecParamsQueryType = "EUI64"
	RadarDNSSummaryGetDnssecParamsQueryTypeGpos       RadarDNSSummaryGetDnssecParamsQueryType = "GPOS"
	RadarDNSSummaryGetDnssecParamsQueryTypeGid        RadarDNSSummaryGetDnssecParamsQueryType = "GID"
	RadarDNSSummaryGetDnssecParamsQueryTypeHinfo      RadarDNSSummaryGetDnssecParamsQueryType = "HINFO"
	RadarDNSSummaryGetDnssecParamsQueryTypeHip        RadarDNSSummaryGetDnssecParamsQueryType = "HIP"
	RadarDNSSummaryGetDnssecParamsQueryTypeHTTPS      RadarDNSSummaryGetDnssecParamsQueryType = "HTTPS"
	RadarDNSSummaryGetDnssecParamsQueryTypeIpseckey   RadarDNSSummaryGetDnssecParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeIsdn       RadarDNSSummaryGetDnssecParamsQueryType = "ISDN"
	RadarDNSSummaryGetDnssecParamsQueryTypeIxfr       RadarDNSSummaryGetDnssecParamsQueryType = "IXFR"
	RadarDNSSummaryGetDnssecParamsQueryTypeKey        RadarDNSSummaryGetDnssecParamsQueryType = "KEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeKx         RadarDNSSummaryGetDnssecParamsQueryType = "KX"
	RadarDNSSummaryGetDnssecParamsQueryTypeL32        RadarDNSSummaryGetDnssecParamsQueryType = "L32"
	RadarDNSSummaryGetDnssecParamsQueryTypeL64        RadarDNSSummaryGetDnssecParamsQueryType = "L64"
	RadarDNSSummaryGetDnssecParamsQueryTypeLoc        RadarDNSSummaryGetDnssecParamsQueryType = "LOC"
	RadarDNSSummaryGetDnssecParamsQueryTypeLp         RadarDNSSummaryGetDnssecParamsQueryType = "LP"
	RadarDNSSummaryGetDnssecParamsQueryTypeMaila      RadarDNSSummaryGetDnssecParamsQueryType = "MAILA"
	RadarDNSSummaryGetDnssecParamsQueryTypeMailb      RadarDNSSummaryGetDnssecParamsQueryType = "MAILB"
	RadarDNSSummaryGetDnssecParamsQueryTypeMB         RadarDNSSummaryGetDnssecParamsQueryType = "MB"
	RadarDNSSummaryGetDnssecParamsQueryTypeMd         RadarDNSSummaryGetDnssecParamsQueryType = "MD"
	RadarDNSSummaryGetDnssecParamsQueryTypeMf         RadarDNSSummaryGetDnssecParamsQueryType = "MF"
	RadarDNSSummaryGetDnssecParamsQueryTypeMg         RadarDNSSummaryGetDnssecParamsQueryType = "MG"
	RadarDNSSummaryGetDnssecParamsQueryTypeMinfo      RadarDNSSummaryGetDnssecParamsQueryType = "MINFO"
	RadarDNSSummaryGetDnssecParamsQueryTypeMr         RadarDNSSummaryGetDnssecParamsQueryType = "MR"
	RadarDNSSummaryGetDnssecParamsQueryTypeMx         RadarDNSSummaryGetDnssecParamsQueryType = "MX"
	RadarDNSSummaryGetDnssecParamsQueryTypeNaptr      RadarDNSSummaryGetDnssecParamsQueryType = "NAPTR"
	RadarDNSSummaryGetDnssecParamsQueryTypeNb         RadarDNSSummaryGetDnssecParamsQueryType = "NB"
	RadarDNSSummaryGetDnssecParamsQueryTypeNbstat     RadarDNSSummaryGetDnssecParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetDnssecParamsQueryTypeNid        RadarDNSSummaryGetDnssecParamsQueryType = "NID"
	RadarDNSSummaryGetDnssecParamsQueryTypeNimloc     RadarDNSSummaryGetDnssecParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetDnssecParamsQueryTypeNinfo      RadarDNSSummaryGetDnssecParamsQueryType = "NINFO"
	RadarDNSSummaryGetDnssecParamsQueryTypeNs         RadarDNSSummaryGetDnssecParamsQueryType = "NS"
	RadarDNSSummaryGetDnssecParamsQueryTypeNsap       RadarDNSSummaryGetDnssecParamsQueryType = "NSAP"
	RadarDNSSummaryGetDnssecParamsQueryTypeNsec       RadarDNSSummaryGetDnssecParamsQueryType = "NSEC"
	RadarDNSSummaryGetDnssecParamsQueryTypeNsec3      RadarDNSSummaryGetDnssecParamsQueryType = "NSEC3"
	RadarDNSSummaryGetDnssecParamsQueryTypeNsec3Param RadarDNSSummaryGetDnssecParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetDnssecParamsQueryTypeNull       RadarDNSSummaryGetDnssecParamsQueryType = "NULL"
	RadarDNSSummaryGetDnssecParamsQueryTypeNxt        RadarDNSSummaryGetDnssecParamsQueryType = "NXT"
	RadarDNSSummaryGetDnssecParamsQueryTypeOpenpgpkey RadarDNSSummaryGetDnssecParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeOpt        RadarDNSSummaryGetDnssecParamsQueryType = "OPT"
	RadarDNSSummaryGetDnssecParamsQueryTypePtr        RadarDNSSummaryGetDnssecParamsQueryType = "PTR"
	RadarDNSSummaryGetDnssecParamsQueryTypePx         RadarDNSSummaryGetDnssecParamsQueryType = "PX"
	RadarDNSSummaryGetDnssecParamsQueryTypeRkey       RadarDNSSummaryGetDnssecParamsQueryType = "RKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeRp         RadarDNSSummaryGetDnssecParamsQueryType = "RP"
	RadarDNSSummaryGetDnssecParamsQueryTypeRrsig      RadarDNSSummaryGetDnssecParamsQueryType = "RRSIG"
	RadarDNSSummaryGetDnssecParamsQueryTypeRt         RadarDNSSummaryGetDnssecParamsQueryType = "RT"
	RadarDNSSummaryGetDnssecParamsQueryTypeSig        RadarDNSSummaryGetDnssecParamsQueryType = "SIG"
	RadarDNSSummaryGetDnssecParamsQueryTypeSink       RadarDNSSummaryGetDnssecParamsQueryType = "SINK"
	RadarDNSSummaryGetDnssecParamsQueryTypeSmimea     RadarDNSSummaryGetDnssecParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetDnssecParamsQueryTypeSoa        RadarDNSSummaryGetDnssecParamsQueryType = "SOA"
	RadarDNSSummaryGetDnssecParamsQueryTypeSpf        RadarDNSSummaryGetDnssecParamsQueryType = "SPF"
	RadarDNSSummaryGetDnssecParamsQueryTypeSrv        RadarDNSSummaryGetDnssecParamsQueryType = "SRV"
	RadarDNSSummaryGetDnssecParamsQueryTypeSshfp      RadarDNSSummaryGetDnssecParamsQueryType = "SSHFP"
	RadarDNSSummaryGetDnssecParamsQueryTypeSvcb       RadarDNSSummaryGetDnssecParamsQueryType = "SVCB"
	RadarDNSSummaryGetDnssecParamsQueryTypeTa         RadarDNSSummaryGetDnssecParamsQueryType = "TA"
	RadarDNSSummaryGetDnssecParamsQueryTypeTalink     RadarDNSSummaryGetDnssecParamsQueryType = "TALINK"
	RadarDNSSummaryGetDnssecParamsQueryTypeTkey       RadarDNSSummaryGetDnssecParamsQueryType = "TKEY"
	RadarDNSSummaryGetDnssecParamsQueryTypeTlsa       RadarDNSSummaryGetDnssecParamsQueryType = "TLSA"
	RadarDNSSummaryGetDnssecParamsQueryTypeTsig       RadarDNSSummaryGetDnssecParamsQueryType = "TSIG"
	RadarDNSSummaryGetDnssecParamsQueryTypeTxt        RadarDNSSummaryGetDnssecParamsQueryType = "TXT"
	RadarDNSSummaryGetDnssecParamsQueryTypeUinfo      RadarDNSSummaryGetDnssecParamsQueryType = "UINFO"
	RadarDNSSummaryGetDnssecParamsQueryTypeUid        RadarDNSSummaryGetDnssecParamsQueryType = "UID"
	RadarDNSSummaryGetDnssecParamsQueryTypeUnspec     RadarDNSSummaryGetDnssecParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetDnssecParamsQueryTypeUri        RadarDNSSummaryGetDnssecParamsQueryType = "URI"
	RadarDNSSummaryGetDnssecParamsQueryTypeWks        RadarDNSSummaryGetDnssecParamsQueryType = "WKS"
	RadarDNSSummaryGetDnssecParamsQueryTypeX25        RadarDNSSummaryGetDnssecParamsQueryType = "X25"
	RadarDNSSummaryGetDnssecParamsQueryTypeZonemd     RadarDNSSummaryGetDnssecParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetDnssecParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecParamsQueryTypeA, RadarDNSSummaryGetDnssecParamsQueryTypeAaaa, RadarDNSSummaryGetDnssecParamsQueryTypeA6, RadarDNSSummaryGetDnssecParamsQueryTypeAfsdb, RadarDNSSummaryGetDnssecParamsQueryTypeAny, RadarDNSSummaryGetDnssecParamsQueryTypeApl, RadarDNSSummaryGetDnssecParamsQueryTypeAtma, RadarDNSSummaryGetDnssecParamsQueryTypeAxfr, RadarDNSSummaryGetDnssecParamsQueryTypeCaa, RadarDNSSummaryGetDnssecParamsQueryTypeCdnskey, RadarDNSSummaryGetDnssecParamsQueryTypeCds, RadarDNSSummaryGetDnssecParamsQueryTypeCert, RadarDNSSummaryGetDnssecParamsQueryTypeCname, RadarDNSSummaryGetDnssecParamsQueryTypeCsync, RadarDNSSummaryGetDnssecParamsQueryTypeDhcid, RadarDNSSummaryGetDnssecParamsQueryTypeDlv, RadarDNSSummaryGetDnssecParamsQueryTypeDname, RadarDNSSummaryGetDnssecParamsQueryTypeDnskey, RadarDNSSummaryGetDnssecParamsQueryTypeDoa, RadarDNSSummaryGetDnssecParamsQueryTypeDs, RadarDNSSummaryGetDnssecParamsQueryTypeEid, RadarDNSSummaryGetDnssecParamsQueryTypeEui48, RadarDNSSummaryGetDnssecParamsQueryTypeEui64, RadarDNSSummaryGetDnssecParamsQueryTypeGpos, RadarDNSSummaryGetDnssecParamsQueryTypeGid, RadarDNSSummaryGetDnssecParamsQueryTypeHinfo, RadarDNSSummaryGetDnssecParamsQueryTypeHip, RadarDNSSummaryGetDnssecParamsQueryTypeHTTPS, RadarDNSSummaryGetDnssecParamsQueryTypeIpseckey, RadarDNSSummaryGetDnssecParamsQueryTypeIsdn, RadarDNSSummaryGetDnssecParamsQueryTypeIxfr, RadarDNSSummaryGetDnssecParamsQueryTypeKey, RadarDNSSummaryGetDnssecParamsQueryTypeKx, RadarDNSSummaryGetDnssecParamsQueryTypeL32, RadarDNSSummaryGetDnssecParamsQueryTypeL64, RadarDNSSummaryGetDnssecParamsQueryTypeLoc, RadarDNSSummaryGetDnssecParamsQueryTypeLp, RadarDNSSummaryGetDnssecParamsQueryTypeMaila, RadarDNSSummaryGetDnssecParamsQueryTypeMailb, RadarDNSSummaryGetDnssecParamsQueryTypeMB, RadarDNSSummaryGetDnssecParamsQueryTypeMd, RadarDNSSummaryGetDnssecParamsQueryTypeMf, RadarDNSSummaryGetDnssecParamsQueryTypeMg, RadarDNSSummaryGetDnssecParamsQueryTypeMinfo, RadarDNSSummaryGetDnssecParamsQueryTypeMr, RadarDNSSummaryGetDnssecParamsQueryTypeMx, RadarDNSSummaryGetDnssecParamsQueryTypeNaptr, RadarDNSSummaryGetDnssecParamsQueryTypeNb, RadarDNSSummaryGetDnssecParamsQueryTypeNbstat, RadarDNSSummaryGetDnssecParamsQueryTypeNid, RadarDNSSummaryGetDnssecParamsQueryTypeNimloc, RadarDNSSummaryGetDnssecParamsQueryTypeNinfo, RadarDNSSummaryGetDnssecParamsQueryTypeNs, RadarDNSSummaryGetDnssecParamsQueryTypeNsap, RadarDNSSummaryGetDnssecParamsQueryTypeNsec, RadarDNSSummaryGetDnssecParamsQueryTypeNsec3, RadarDNSSummaryGetDnssecParamsQueryTypeNsec3Param, RadarDNSSummaryGetDnssecParamsQueryTypeNull, RadarDNSSummaryGetDnssecParamsQueryTypeNxt, RadarDNSSummaryGetDnssecParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetDnssecParamsQueryTypeOpt, RadarDNSSummaryGetDnssecParamsQueryTypePtr, RadarDNSSummaryGetDnssecParamsQueryTypePx, RadarDNSSummaryGetDnssecParamsQueryTypeRkey, RadarDNSSummaryGetDnssecParamsQueryTypeRp, RadarDNSSummaryGetDnssecParamsQueryTypeRrsig, RadarDNSSummaryGetDnssecParamsQueryTypeRt, RadarDNSSummaryGetDnssecParamsQueryTypeSig, RadarDNSSummaryGetDnssecParamsQueryTypeSink, RadarDNSSummaryGetDnssecParamsQueryTypeSmimea, RadarDNSSummaryGetDnssecParamsQueryTypeSoa, RadarDNSSummaryGetDnssecParamsQueryTypeSpf, RadarDNSSummaryGetDnssecParamsQueryTypeSrv, RadarDNSSummaryGetDnssecParamsQueryTypeSshfp, RadarDNSSummaryGetDnssecParamsQueryTypeSvcb, RadarDNSSummaryGetDnssecParamsQueryTypeTa, RadarDNSSummaryGetDnssecParamsQueryTypeTalink, RadarDNSSummaryGetDnssecParamsQueryTypeTkey, RadarDNSSummaryGetDnssecParamsQueryTypeTlsa, RadarDNSSummaryGetDnssecParamsQueryTypeTsig, RadarDNSSummaryGetDnssecParamsQueryTypeTxt, RadarDNSSummaryGetDnssecParamsQueryTypeUinfo, RadarDNSSummaryGetDnssecParamsQueryTypeUid, RadarDNSSummaryGetDnssecParamsQueryTypeUnspec, RadarDNSSummaryGetDnssecParamsQueryTypeUri, RadarDNSSummaryGetDnssecParamsQueryTypeWks, RadarDNSSummaryGetDnssecParamsQueryTypeX25, RadarDNSSummaryGetDnssecParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetDnssecParamsResponseCode string

const (
	RadarDNSSummaryGetDnssecParamsResponseCodeNoerror   RadarDNSSummaryGetDnssecParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetDnssecParamsResponseCodeFormerr   RadarDNSSummaryGetDnssecParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetDnssecParamsResponseCodeServfail  RadarDNSSummaryGetDnssecParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetDnssecParamsResponseCodeNxdomain  RadarDNSSummaryGetDnssecParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetDnssecParamsResponseCodeNotimp    RadarDNSSummaryGetDnssecParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetDnssecParamsResponseCodeRefused   RadarDNSSummaryGetDnssecParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetDnssecParamsResponseCodeYxdomain  RadarDNSSummaryGetDnssecParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetDnssecParamsResponseCodeYxrrset   RadarDNSSummaryGetDnssecParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetDnssecParamsResponseCodeNxrrset   RadarDNSSummaryGetDnssecParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetDnssecParamsResponseCodeNotauth   RadarDNSSummaryGetDnssecParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetDnssecParamsResponseCodeNotzone   RadarDNSSummaryGetDnssecParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadsig    RadarDNSSummaryGetDnssecParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadkey    RadarDNSSummaryGetDnssecParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadtime   RadarDNSSummaryGetDnssecParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadmode   RadarDNSSummaryGetDnssecParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadname   RadarDNSSummaryGetDnssecParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadalg    RadarDNSSummaryGetDnssecParamsResponseCode = "BADALG"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadtrunc  RadarDNSSummaryGetDnssecParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetDnssecParamsResponseCodeBadcookie RadarDNSSummaryGetDnssecParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetDnssecParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecParamsResponseCodeNoerror, RadarDNSSummaryGetDnssecParamsResponseCodeFormerr, RadarDNSSummaryGetDnssecParamsResponseCodeServfail, RadarDNSSummaryGetDnssecParamsResponseCodeNxdomain, RadarDNSSummaryGetDnssecParamsResponseCodeNotimp, RadarDNSSummaryGetDnssecParamsResponseCodeRefused, RadarDNSSummaryGetDnssecParamsResponseCodeYxdomain, RadarDNSSummaryGetDnssecParamsResponseCodeYxrrset, RadarDNSSummaryGetDnssecParamsResponseCodeNxrrset, RadarDNSSummaryGetDnssecParamsResponseCodeNotauth, RadarDNSSummaryGetDnssecParamsResponseCodeNotzone, RadarDNSSummaryGetDnssecParamsResponseCodeBadsig, RadarDNSSummaryGetDnssecParamsResponseCodeBadkey, RadarDNSSummaryGetDnssecParamsResponseCodeBadtime, RadarDNSSummaryGetDnssecParamsResponseCodeBadmode, RadarDNSSummaryGetDnssecParamsResponseCodeBadname, RadarDNSSummaryGetDnssecParamsResponseCodeBadalg, RadarDNSSummaryGetDnssecParamsResponseCodeBadtrunc, RadarDNSSummaryGetDnssecParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetDnssecAwareParams struct {
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
	Format param.Field[RadarDNSSummaryGetDnssecAwareParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetDnssecAwareParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetDnssecAwareParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetDnssecAwareParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetDnssecAwareParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetDnssecAwareParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetDnssecAwareParamsFormat string

const (
	RadarDNSSummaryGetDnssecAwareParamsFormatJson RadarDNSSummaryGetDnssecAwareParamsFormat = "JSON"
	RadarDNSSummaryGetDnssecAwareParamsFormatCsv  RadarDNSSummaryGetDnssecAwareParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetDnssecAwareParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecAwareParamsFormatJson, RadarDNSSummaryGetDnssecAwareParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetDnssecAwareParamsProtocol string

const (
	RadarDNSSummaryGetDnssecAwareParamsProtocolUdp   RadarDNSSummaryGetDnssecAwareParamsProtocol = "UDP"
	RadarDNSSummaryGetDnssecAwareParamsProtocolTcp   RadarDNSSummaryGetDnssecAwareParamsProtocol = "TCP"
	RadarDNSSummaryGetDnssecAwareParamsProtocolHTTPS RadarDNSSummaryGetDnssecAwareParamsProtocol = "HTTPS"
	RadarDNSSummaryGetDnssecAwareParamsProtocolTls   RadarDNSSummaryGetDnssecAwareParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetDnssecAwareParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecAwareParamsProtocolUdp, RadarDNSSummaryGetDnssecAwareParamsProtocolTcp, RadarDNSSummaryGetDnssecAwareParamsProtocolHTTPS, RadarDNSSummaryGetDnssecAwareParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetDnssecAwareParamsQueryType string

const (
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeA          RadarDNSSummaryGetDnssecAwareParamsQueryType = "A"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeAaaa       RadarDNSSummaryGetDnssecAwareParamsQueryType = "AAAA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeA6         RadarDNSSummaryGetDnssecAwareParamsQueryType = "A6"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeAfsdb      RadarDNSSummaryGetDnssecAwareParamsQueryType = "AFSDB"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeAny        RadarDNSSummaryGetDnssecAwareParamsQueryType = "ANY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeApl        RadarDNSSummaryGetDnssecAwareParamsQueryType = "APL"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeAtma       RadarDNSSummaryGetDnssecAwareParamsQueryType = "ATMA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeAxfr       RadarDNSSummaryGetDnssecAwareParamsQueryType = "AXFR"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCaa        RadarDNSSummaryGetDnssecAwareParamsQueryType = "CAA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCdnskey    RadarDNSSummaryGetDnssecAwareParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCds        RadarDNSSummaryGetDnssecAwareParamsQueryType = "CDS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCert       RadarDNSSummaryGetDnssecAwareParamsQueryType = "CERT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCname      RadarDNSSummaryGetDnssecAwareParamsQueryType = "CNAME"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeCsync      RadarDNSSummaryGetDnssecAwareParamsQueryType = "CSYNC"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDhcid      RadarDNSSummaryGetDnssecAwareParamsQueryType = "DHCID"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDlv        RadarDNSSummaryGetDnssecAwareParamsQueryType = "DLV"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDname      RadarDNSSummaryGetDnssecAwareParamsQueryType = "DNAME"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDnskey     RadarDNSSummaryGetDnssecAwareParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDoa        RadarDNSSummaryGetDnssecAwareParamsQueryType = "DOA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeDs         RadarDNSSummaryGetDnssecAwareParamsQueryType = "DS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeEid        RadarDNSSummaryGetDnssecAwareParamsQueryType = "EID"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeEui48      RadarDNSSummaryGetDnssecAwareParamsQueryType = "EUI48"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeEui64      RadarDNSSummaryGetDnssecAwareParamsQueryType = "EUI64"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeGpos       RadarDNSSummaryGetDnssecAwareParamsQueryType = "GPOS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeGid        RadarDNSSummaryGetDnssecAwareParamsQueryType = "GID"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeHinfo      RadarDNSSummaryGetDnssecAwareParamsQueryType = "HINFO"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeHip        RadarDNSSummaryGetDnssecAwareParamsQueryType = "HIP"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeHTTPS      RadarDNSSummaryGetDnssecAwareParamsQueryType = "HTTPS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeIpseckey   RadarDNSSummaryGetDnssecAwareParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeIsdn       RadarDNSSummaryGetDnssecAwareParamsQueryType = "ISDN"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeIxfr       RadarDNSSummaryGetDnssecAwareParamsQueryType = "IXFR"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeKey        RadarDNSSummaryGetDnssecAwareParamsQueryType = "KEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeKx         RadarDNSSummaryGetDnssecAwareParamsQueryType = "KX"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeL32        RadarDNSSummaryGetDnssecAwareParamsQueryType = "L32"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeL64        RadarDNSSummaryGetDnssecAwareParamsQueryType = "L64"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeLoc        RadarDNSSummaryGetDnssecAwareParamsQueryType = "LOC"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeLp         RadarDNSSummaryGetDnssecAwareParamsQueryType = "LP"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMaila      RadarDNSSummaryGetDnssecAwareParamsQueryType = "MAILA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMailb      RadarDNSSummaryGetDnssecAwareParamsQueryType = "MAILB"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMB         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MB"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMd         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MD"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMf         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MF"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMg         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MG"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMinfo      RadarDNSSummaryGetDnssecAwareParamsQueryType = "MINFO"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMr         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MR"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeMx         RadarDNSSummaryGetDnssecAwareParamsQueryType = "MX"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNaptr      RadarDNSSummaryGetDnssecAwareParamsQueryType = "NAPTR"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNb         RadarDNSSummaryGetDnssecAwareParamsQueryType = "NB"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNbstat     RadarDNSSummaryGetDnssecAwareParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNid        RadarDNSSummaryGetDnssecAwareParamsQueryType = "NID"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNimloc     RadarDNSSummaryGetDnssecAwareParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNinfo      RadarDNSSummaryGetDnssecAwareParamsQueryType = "NINFO"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNs         RadarDNSSummaryGetDnssecAwareParamsQueryType = "NS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsap       RadarDNSSummaryGetDnssecAwareParamsQueryType = "NSAP"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec       RadarDNSSummaryGetDnssecAwareParamsQueryType = "NSEC"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec3      RadarDNSSummaryGetDnssecAwareParamsQueryType = "NSEC3"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec3Param RadarDNSSummaryGetDnssecAwareParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNull       RadarDNSSummaryGetDnssecAwareParamsQueryType = "NULL"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeNxt        RadarDNSSummaryGetDnssecAwareParamsQueryType = "NXT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeOpenpgpkey RadarDNSSummaryGetDnssecAwareParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeOpt        RadarDNSSummaryGetDnssecAwareParamsQueryType = "OPT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypePtr        RadarDNSSummaryGetDnssecAwareParamsQueryType = "PTR"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypePx         RadarDNSSummaryGetDnssecAwareParamsQueryType = "PX"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeRkey       RadarDNSSummaryGetDnssecAwareParamsQueryType = "RKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeRp         RadarDNSSummaryGetDnssecAwareParamsQueryType = "RP"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeRrsig      RadarDNSSummaryGetDnssecAwareParamsQueryType = "RRSIG"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeRt         RadarDNSSummaryGetDnssecAwareParamsQueryType = "RT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSig        RadarDNSSummaryGetDnssecAwareParamsQueryType = "SIG"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSink       RadarDNSSummaryGetDnssecAwareParamsQueryType = "SINK"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSmimea     RadarDNSSummaryGetDnssecAwareParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSoa        RadarDNSSummaryGetDnssecAwareParamsQueryType = "SOA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSpf        RadarDNSSummaryGetDnssecAwareParamsQueryType = "SPF"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSrv        RadarDNSSummaryGetDnssecAwareParamsQueryType = "SRV"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSshfp      RadarDNSSummaryGetDnssecAwareParamsQueryType = "SSHFP"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeSvcb       RadarDNSSummaryGetDnssecAwareParamsQueryType = "SVCB"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTa         RadarDNSSummaryGetDnssecAwareParamsQueryType = "TA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTalink     RadarDNSSummaryGetDnssecAwareParamsQueryType = "TALINK"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTkey       RadarDNSSummaryGetDnssecAwareParamsQueryType = "TKEY"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTlsa       RadarDNSSummaryGetDnssecAwareParamsQueryType = "TLSA"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTsig       RadarDNSSummaryGetDnssecAwareParamsQueryType = "TSIG"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeTxt        RadarDNSSummaryGetDnssecAwareParamsQueryType = "TXT"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeUinfo      RadarDNSSummaryGetDnssecAwareParamsQueryType = "UINFO"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeUid        RadarDNSSummaryGetDnssecAwareParamsQueryType = "UID"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeUnspec     RadarDNSSummaryGetDnssecAwareParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeUri        RadarDNSSummaryGetDnssecAwareParamsQueryType = "URI"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeWks        RadarDNSSummaryGetDnssecAwareParamsQueryType = "WKS"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeX25        RadarDNSSummaryGetDnssecAwareParamsQueryType = "X25"
	RadarDNSSummaryGetDnssecAwareParamsQueryTypeZonemd     RadarDNSSummaryGetDnssecAwareParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetDnssecAwareParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecAwareParamsQueryTypeA, RadarDNSSummaryGetDnssecAwareParamsQueryTypeAaaa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeA6, RadarDNSSummaryGetDnssecAwareParamsQueryTypeAfsdb, RadarDNSSummaryGetDnssecAwareParamsQueryTypeAny, RadarDNSSummaryGetDnssecAwareParamsQueryTypeApl, RadarDNSSummaryGetDnssecAwareParamsQueryTypeAtma, RadarDNSSummaryGetDnssecAwareParamsQueryTypeAxfr, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCaa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCdnskey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCds, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCert, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCname, RadarDNSSummaryGetDnssecAwareParamsQueryTypeCsync, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDhcid, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDlv, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDname, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDnskey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDoa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeDs, RadarDNSSummaryGetDnssecAwareParamsQueryTypeEid, RadarDNSSummaryGetDnssecAwareParamsQueryTypeEui48, RadarDNSSummaryGetDnssecAwareParamsQueryTypeEui64, RadarDNSSummaryGetDnssecAwareParamsQueryTypeGpos, RadarDNSSummaryGetDnssecAwareParamsQueryTypeGid, RadarDNSSummaryGetDnssecAwareParamsQueryTypeHinfo, RadarDNSSummaryGetDnssecAwareParamsQueryTypeHip, RadarDNSSummaryGetDnssecAwareParamsQueryTypeHTTPS, RadarDNSSummaryGetDnssecAwareParamsQueryTypeIpseckey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeIsdn, RadarDNSSummaryGetDnssecAwareParamsQueryTypeIxfr, RadarDNSSummaryGetDnssecAwareParamsQueryTypeKey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeKx, RadarDNSSummaryGetDnssecAwareParamsQueryTypeL32, RadarDNSSummaryGetDnssecAwareParamsQueryTypeL64, RadarDNSSummaryGetDnssecAwareParamsQueryTypeLoc, RadarDNSSummaryGetDnssecAwareParamsQueryTypeLp, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMaila, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMailb, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMB, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMd, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMf, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMg, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMinfo, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMr, RadarDNSSummaryGetDnssecAwareParamsQueryTypeMx, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNaptr, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNb, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNbstat, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNid, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNimloc, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNinfo, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNs, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsap, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec3, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNsec3Param, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNull, RadarDNSSummaryGetDnssecAwareParamsQueryTypeNxt, RadarDNSSummaryGetDnssecAwareParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeOpt, RadarDNSSummaryGetDnssecAwareParamsQueryTypePtr, RadarDNSSummaryGetDnssecAwareParamsQueryTypePx, RadarDNSSummaryGetDnssecAwareParamsQueryTypeRkey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeRp, RadarDNSSummaryGetDnssecAwareParamsQueryTypeRrsig, RadarDNSSummaryGetDnssecAwareParamsQueryTypeRt, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSig, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSink, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSmimea, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSoa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSpf, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSrv, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSshfp, RadarDNSSummaryGetDnssecAwareParamsQueryTypeSvcb, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTalink, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTkey, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTlsa, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTsig, RadarDNSSummaryGetDnssecAwareParamsQueryTypeTxt, RadarDNSSummaryGetDnssecAwareParamsQueryTypeUinfo, RadarDNSSummaryGetDnssecAwareParamsQueryTypeUid, RadarDNSSummaryGetDnssecAwareParamsQueryTypeUnspec, RadarDNSSummaryGetDnssecAwareParamsQueryTypeUri, RadarDNSSummaryGetDnssecAwareParamsQueryTypeWks, RadarDNSSummaryGetDnssecAwareParamsQueryTypeX25, RadarDNSSummaryGetDnssecAwareParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetDnssecAwareParamsResponseCode string

const (
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNoerror   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeFormerr   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeServfail  RadarDNSSummaryGetDnssecAwareParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNxdomain  RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotimp    RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeRefused   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeYxdomain  RadarDNSSummaryGetDnssecAwareParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeYxrrset   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNxrrset   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotauth   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotzone   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadsig    RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadkey    RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadtime   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadmode   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadname   RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadalg    RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADALG"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadtrunc  RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadcookie RadarDNSSummaryGetDnssecAwareParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetDnssecAwareParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecAwareParamsResponseCodeNoerror, RadarDNSSummaryGetDnssecAwareParamsResponseCodeFormerr, RadarDNSSummaryGetDnssecAwareParamsResponseCodeServfail, RadarDNSSummaryGetDnssecAwareParamsResponseCodeNxdomain, RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotimp, RadarDNSSummaryGetDnssecAwareParamsResponseCodeRefused, RadarDNSSummaryGetDnssecAwareParamsResponseCodeYxdomain, RadarDNSSummaryGetDnssecAwareParamsResponseCodeYxrrset, RadarDNSSummaryGetDnssecAwareParamsResponseCodeNxrrset, RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotauth, RadarDNSSummaryGetDnssecAwareParamsResponseCodeNotzone, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadsig, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadkey, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadtime, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadmode, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadname, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadalg, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadtrunc, RadarDNSSummaryGetDnssecAwareParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetDnssecE2EParams struct {
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
	Format param.Field[RadarDNSSummaryGetDnssecE2EParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetDnssecE2EParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetDnssecE2EParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetDnssecE2EParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetDnssecE2EParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetDnssecE2EParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetDnssecE2EParamsFormat string

const (
	RadarDNSSummaryGetDnssecE2EParamsFormatJson RadarDNSSummaryGetDnssecE2EParamsFormat = "JSON"
	RadarDNSSummaryGetDnssecE2EParamsFormatCsv  RadarDNSSummaryGetDnssecE2EParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetDnssecE2EParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecE2EParamsFormatJson, RadarDNSSummaryGetDnssecE2EParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetDnssecE2EParamsProtocol string

const (
	RadarDNSSummaryGetDnssecE2EParamsProtocolUdp   RadarDNSSummaryGetDnssecE2EParamsProtocol = "UDP"
	RadarDNSSummaryGetDnssecE2EParamsProtocolTcp   RadarDNSSummaryGetDnssecE2EParamsProtocol = "TCP"
	RadarDNSSummaryGetDnssecE2EParamsProtocolHTTPS RadarDNSSummaryGetDnssecE2EParamsProtocol = "HTTPS"
	RadarDNSSummaryGetDnssecE2EParamsProtocolTls   RadarDNSSummaryGetDnssecE2EParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetDnssecE2EParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecE2EParamsProtocolUdp, RadarDNSSummaryGetDnssecE2EParamsProtocolTcp, RadarDNSSummaryGetDnssecE2EParamsProtocolHTTPS, RadarDNSSummaryGetDnssecE2EParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetDnssecE2EParamsQueryType string

const (
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeA          RadarDNSSummaryGetDnssecE2EParamsQueryType = "A"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeAaaa       RadarDNSSummaryGetDnssecE2EParamsQueryType = "AAAA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeA6         RadarDNSSummaryGetDnssecE2EParamsQueryType = "A6"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeAfsdb      RadarDNSSummaryGetDnssecE2EParamsQueryType = "AFSDB"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeAny        RadarDNSSummaryGetDnssecE2EParamsQueryType = "ANY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeApl        RadarDNSSummaryGetDnssecE2EParamsQueryType = "APL"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeAtma       RadarDNSSummaryGetDnssecE2EParamsQueryType = "ATMA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeAxfr       RadarDNSSummaryGetDnssecE2EParamsQueryType = "AXFR"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCaa        RadarDNSSummaryGetDnssecE2EParamsQueryType = "CAA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCdnskey    RadarDNSSummaryGetDnssecE2EParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCds        RadarDNSSummaryGetDnssecE2EParamsQueryType = "CDS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCert       RadarDNSSummaryGetDnssecE2EParamsQueryType = "CERT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCname      RadarDNSSummaryGetDnssecE2EParamsQueryType = "CNAME"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeCsync      RadarDNSSummaryGetDnssecE2EParamsQueryType = "CSYNC"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDhcid      RadarDNSSummaryGetDnssecE2EParamsQueryType = "DHCID"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDlv        RadarDNSSummaryGetDnssecE2EParamsQueryType = "DLV"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDname      RadarDNSSummaryGetDnssecE2EParamsQueryType = "DNAME"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDnskey     RadarDNSSummaryGetDnssecE2EParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDoa        RadarDNSSummaryGetDnssecE2EParamsQueryType = "DOA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeDs         RadarDNSSummaryGetDnssecE2EParamsQueryType = "DS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeEid        RadarDNSSummaryGetDnssecE2EParamsQueryType = "EID"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeEui48      RadarDNSSummaryGetDnssecE2EParamsQueryType = "EUI48"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeEui64      RadarDNSSummaryGetDnssecE2EParamsQueryType = "EUI64"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeGpos       RadarDNSSummaryGetDnssecE2EParamsQueryType = "GPOS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeGid        RadarDNSSummaryGetDnssecE2EParamsQueryType = "GID"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeHinfo      RadarDNSSummaryGetDnssecE2EParamsQueryType = "HINFO"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeHip        RadarDNSSummaryGetDnssecE2EParamsQueryType = "HIP"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeHTTPS      RadarDNSSummaryGetDnssecE2EParamsQueryType = "HTTPS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeIpseckey   RadarDNSSummaryGetDnssecE2EParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeIsdn       RadarDNSSummaryGetDnssecE2EParamsQueryType = "ISDN"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeIxfr       RadarDNSSummaryGetDnssecE2EParamsQueryType = "IXFR"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeKey        RadarDNSSummaryGetDnssecE2EParamsQueryType = "KEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeKx         RadarDNSSummaryGetDnssecE2EParamsQueryType = "KX"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeL32        RadarDNSSummaryGetDnssecE2EParamsQueryType = "L32"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeL64        RadarDNSSummaryGetDnssecE2EParamsQueryType = "L64"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeLoc        RadarDNSSummaryGetDnssecE2EParamsQueryType = "LOC"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeLp         RadarDNSSummaryGetDnssecE2EParamsQueryType = "LP"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMaila      RadarDNSSummaryGetDnssecE2EParamsQueryType = "MAILA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMailb      RadarDNSSummaryGetDnssecE2EParamsQueryType = "MAILB"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMB         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MB"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMd         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MD"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMf         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MF"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMg         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MG"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMinfo      RadarDNSSummaryGetDnssecE2EParamsQueryType = "MINFO"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMr         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MR"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeMx         RadarDNSSummaryGetDnssecE2EParamsQueryType = "MX"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNaptr      RadarDNSSummaryGetDnssecE2EParamsQueryType = "NAPTR"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNb         RadarDNSSummaryGetDnssecE2EParamsQueryType = "NB"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNbstat     RadarDNSSummaryGetDnssecE2EParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNid        RadarDNSSummaryGetDnssecE2EParamsQueryType = "NID"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNimloc     RadarDNSSummaryGetDnssecE2EParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNinfo      RadarDNSSummaryGetDnssecE2EParamsQueryType = "NINFO"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNs         RadarDNSSummaryGetDnssecE2EParamsQueryType = "NS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsap       RadarDNSSummaryGetDnssecE2EParamsQueryType = "NSAP"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec       RadarDNSSummaryGetDnssecE2EParamsQueryType = "NSEC"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec3      RadarDNSSummaryGetDnssecE2EParamsQueryType = "NSEC3"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec3Param RadarDNSSummaryGetDnssecE2EParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNull       RadarDNSSummaryGetDnssecE2EParamsQueryType = "NULL"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeNxt        RadarDNSSummaryGetDnssecE2EParamsQueryType = "NXT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeOpenpgpkey RadarDNSSummaryGetDnssecE2EParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeOpt        RadarDNSSummaryGetDnssecE2EParamsQueryType = "OPT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypePtr        RadarDNSSummaryGetDnssecE2EParamsQueryType = "PTR"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypePx         RadarDNSSummaryGetDnssecE2EParamsQueryType = "PX"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeRkey       RadarDNSSummaryGetDnssecE2EParamsQueryType = "RKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeRp         RadarDNSSummaryGetDnssecE2EParamsQueryType = "RP"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeRrsig      RadarDNSSummaryGetDnssecE2EParamsQueryType = "RRSIG"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeRt         RadarDNSSummaryGetDnssecE2EParamsQueryType = "RT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSig        RadarDNSSummaryGetDnssecE2EParamsQueryType = "SIG"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSink       RadarDNSSummaryGetDnssecE2EParamsQueryType = "SINK"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSmimea     RadarDNSSummaryGetDnssecE2EParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSoa        RadarDNSSummaryGetDnssecE2EParamsQueryType = "SOA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSpf        RadarDNSSummaryGetDnssecE2EParamsQueryType = "SPF"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSrv        RadarDNSSummaryGetDnssecE2EParamsQueryType = "SRV"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSshfp      RadarDNSSummaryGetDnssecE2EParamsQueryType = "SSHFP"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeSvcb       RadarDNSSummaryGetDnssecE2EParamsQueryType = "SVCB"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTa         RadarDNSSummaryGetDnssecE2EParamsQueryType = "TA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTalink     RadarDNSSummaryGetDnssecE2EParamsQueryType = "TALINK"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTkey       RadarDNSSummaryGetDnssecE2EParamsQueryType = "TKEY"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTlsa       RadarDNSSummaryGetDnssecE2EParamsQueryType = "TLSA"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTsig       RadarDNSSummaryGetDnssecE2EParamsQueryType = "TSIG"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeTxt        RadarDNSSummaryGetDnssecE2EParamsQueryType = "TXT"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeUinfo      RadarDNSSummaryGetDnssecE2EParamsQueryType = "UINFO"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeUid        RadarDNSSummaryGetDnssecE2EParamsQueryType = "UID"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeUnspec     RadarDNSSummaryGetDnssecE2EParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeUri        RadarDNSSummaryGetDnssecE2EParamsQueryType = "URI"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeWks        RadarDNSSummaryGetDnssecE2EParamsQueryType = "WKS"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeX25        RadarDNSSummaryGetDnssecE2EParamsQueryType = "X25"
	RadarDNSSummaryGetDnssecE2EParamsQueryTypeZonemd     RadarDNSSummaryGetDnssecE2EParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetDnssecE2EParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecE2EParamsQueryTypeA, RadarDNSSummaryGetDnssecE2EParamsQueryTypeAaaa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeA6, RadarDNSSummaryGetDnssecE2EParamsQueryTypeAfsdb, RadarDNSSummaryGetDnssecE2EParamsQueryTypeAny, RadarDNSSummaryGetDnssecE2EParamsQueryTypeApl, RadarDNSSummaryGetDnssecE2EParamsQueryTypeAtma, RadarDNSSummaryGetDnssecE2EParamsQueryTypeAxfr, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCaa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCdnskey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCds, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCert, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCname, RadarDNSSummaryGetDnssecE2EParamsQueryTypeCsync, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDhcid, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDlv, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDname, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDnskey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDoa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeDs, RadarDNSSummaryGetDnssecE2EParamsQueryTypeEid, RadarDNSSummaryGetDnssecE2EParamsQueryTypeEui48, RadarDNSSummaryGetDnssecE2EParamsQueryTypeEui64, RadarDNSSummaryGetDnssecE2EParamsQueryTypeGpos, RadarDNSSummaryGetDnssecE2EParamsQueryTypeGid, RadarDNSSummaryGetDnssecE2EParamsQueryTypeHinfo, RadarDNSSummaryGetDnssecE2EParamsQueryTypeHip, RadarDNSSummaryGetDnssecE2EParamsQueryTypeHTTPS, RadarDNSSummaryGetDnssecE2EParamsQueryTypeIpseckey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeIsdn, RadarDNSSummaryGetDnssecE2EParamsQueryTypeIxfr, RadarDNSSummaryGetDnssecE2EParamsQueryTypeKey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeKx, RadarDNSSummaryGetDnssecE2EParamsQueryTypeL32, RadarDNSSummaryGetDnssecE2EParamsQueryTypeL64, RadarDNSSummaryGetDnssecE2EParamsQueryTypeLoc, RadarDNSSummaryGetDnssecE2EParamsQueryTypeLp, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMaila, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMailb, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMB, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMd, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMf, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMg, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMinfo, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMr, RadarDNSSummaryGetDnssecE2EParamsQueryTypeMx, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNaptr, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNb, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNbstat, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNid, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNimloc, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNinfo, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNs, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsap, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec3, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNsec3Param, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNull, RadarDNSSummaryGetDnssecE2EParamsQueryTypeNxt, RadarDNSSummaryGetDnssecE2EParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeOpt, RadarDNSSummaryGetDnssecE2EParamsQueryTypePtr, RadarDNSSummaryGetDnssecE2EParamsQueryTypePx, RadarDNSSummaryGetDnssecE2EParamsQueryTypeRkey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeRp, RadarDNSSummaryGetDnssecE2EParamsQueryTypeRrsig, RadarDNSSummaryGetDnssecE2EParamsQueryTypeRt, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSig, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSink, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSmimea, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSoa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSpf, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSrv, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSshfp, RadarDNSSummaryGetDnssecE2EParamsQueryTypeSvcb, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTalink, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTkey, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTlsa, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTsig, RadarDNSSummaryGetDnssecE2EParamsQueryTypeTxt, RadarDNSSummaryGetDnssecE2EParamsQueryTypeUinfo, RadarDNSSummaryGetDnssecE2EParamsQueryTypeUid, RadarDNSSummaryGetDnssecE2EParamsQueryTypeUnspec, RadarDNSSummaryGetDnssecE2EParamsQueryTypeUri, RadarDNSSummaryGetDnssecE2EParamsQueryTypeWks, RadarDNSSummaryGetDnssecE2EParamsQueryTypeX25, RadarDNSSummaryGetDnssecE2EParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetDnssecE2EParamsResponseCode string

const (
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNoerror   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeFormerr   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeServfail  RadarDNSSummaryGetDnssecE2EParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNxdomain  RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotimp    RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeRefused   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeYxdomain  RadarDNSSummaryGetDnssecE2EParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeYxrrset   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNxrrset   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotauth   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotzone   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadsig    RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadkey    RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadtime   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadmode   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadname   RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadalg    RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADALG"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadtrunc  RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadcookie RadarDNSSummaryGetDnssecE2EParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetDnssecE2EParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetDnssecE2EParamsResponseCodeNoerror, RadarDNSSummaryGetDnssecE2EParamsResponseCodeFormerr, RadarDNSSummaryGetDnssecE2EParamsResponseCodeServfail, RadarDNSSummaryGetDnssecE2EParamsResponseCodeNxdomain, RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotimp, RadarDNSSummaryGetDnssecE2EParamsResponseCodeRefused, RadarDNSSummaryGetDnssecE2EParamsResponseCodeYxdomain, RadarDNSSummaryGetDnssecE2EParamsResponseCodeYxrrset, RadarDNSSummaryGetDnssecE2EParamsResponseCodeNxrrset, RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotauth, RadarDNSSummaryGetDnssecE2EParamsResponseCodeNotzone, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadsig, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadkey, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadtime, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadmode, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadname, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadalg, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadtrunc, RadarDNSSummaryGetDnssecE2EParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetIPVersionParams struct {
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
	Format param.Field[RadarDNSSummaryGetIPVersionParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetIPVersionParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetIPVersionParamsFormat string

const (
	RadarDNSSummaryGetIPVersionParamsFormatJson RadarDNSSummaryGetIPVersionParamsFormat = "JSON"
	RadarDNSSummaryGetIPVersionParamsFormatCsv  RadarDNSSummaryGetIPVersionParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetIPVersionParamsFormatJson, RadarDNSSummaryGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetIPVersionParamsProtocol string

const (
	RadarDNSSummaryGetIPVersionParamsProtocolUdp   RadarDNSSummaryGetIPVersionParamsProtocol = "UDP"
	RadarDNSSummaryGetIPVersionParamsProtocolTcp   RadarDNSSummaryGetIPVersionParamsProtocol = "TCP"
	RadarDNSSummaryGetIPVersionParamsProtocolHTTPS RadarDNSSummaryGetIPVersionParamsProtocol = "HTTPS"
	RadarDNSSummaryGetIPVersionParamsProtocolTls   RadarDNSSummaryGetIPVersionParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetIPVersionParamsProtocolUdp, RadarDNSSummaryGetIPVersionParamsProtocolTcp, RadarDNSSummaryGetIPVersionParamsProtocolHTTPS, RadarDNSSummaryGetIPVersionParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetIPVersionParamsQueryType string

const (
	RadarDNSSummaryGetIPVersionParamsQueryTypeA          RadarDNSSummaryGetIPVersionParamsQueryType = "A"
	RadarDNSSummaryGetIPVersionParamsQueryTypeAaaa       RadarDNSSummaryGetIPVersionParamsQueryType = "AAAA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeA6         RadarDNSSummaryGetIPVersionParamsQueryType = "A6"
	RadarDNSSummaryGetIPVersionParamsQueryTypeAfsdb      RadarDNSSummaryGetIPVersionParamsQueryType = "AFSDB"
	RadarDNSSummaryGetIPVersionParamsQueryTypeAny        RadarDNSSummaryGetIPVersionParamsQueryType = "ANY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeApl        RadarDNSSummaryGetIPVersionParamsQueryType = "APL"
	RadarDNSSummaryGetIPVersionParamsQueryTypeAtma       RadarDNSSummaryGetIPVersionParamsQueryType = "ATMA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeAxfr       RadarDNSSummaryGetIPVersionParamsQueryType = "AXFR"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCaa        RadarDNSSummaryGetIPVersionParamsQueryType = "CAA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCdnskey    RadarDNSSummaryGetIPVersionParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCds        RadarDNSSummaryGetIPVersionParamsQueryType = "CDS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCert       RadarDNSSummaryGetIPVersionParamsQueryType = "CERT"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCname      RadarDNSSummaryGetIPVersionParamsQueryType = "CNAME"
	RadarDNSSummaryGetIPVersionParamsQueryTypeCsync      RadarDNSSummaryGetIPVersionParamsQueryType = "CSYNC"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDhcid      RadarDNSSummaryGetIPVersionParamsQueryType = "DHCID"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDlv        RadarDNSSummaryGetIPVersionParamsQueryType = "DLV"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDname      RadarDNSSummaryGetIPVersionParamsQueryType = "DNAME"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDnskey     RadarDNSSummaryGetIPVersionParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDoa        RadarDNSSummaryGetIPVersionParamsQueryType = "DOA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeDs         RadarDNSSummaryGetIPVersionParamsQueryType = "DS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeEid        RadarDNSSummaryGetIPVersionParamsQueryType = "EID"
	RadarDNSSummaryGetIPVersionParamsQueryTypeEui48      RadarDNSSummaryGetIPVersionParamsQueryType = "EUI48"
	RadarDNSSummaryGetIPVersionParamsQueryTypeEui64      RadarDNSSummaryGetIPVersionParamsQueryType = "EUI64"
	RadarDNSSummaryGetIPVersionParamsQueryTypeGpos       RadarDNSSummaryGetIPVersionParamsQueryType = "GPOS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeGid        RadarDNSSummaryGetIPVersionParamsQueryType = "GID"
	RadarDNSSummaryGetIPVersionParamsQueryTypeHinfo      RadarDNSSummaryGetIPVersionParamsQueryType = "HINFO"
	RadarDNSSummaryGetIPVersionParamsQueryTypeHip        RadarDNSSummaryGetIPVersionParamsQueryType = "HIP"
	RadarDNSSummaryGetIPVersionParamsQueryTypeHTTPS      RadarDNSSummaryGetIPVersionParamsQueryType = "HTTPS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeIpseckey   RadarDNSSummaryGetIPVersionParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeIsdn       RadarDNSSummaryGetIPVersionParamsQueryType = "ISDN"
	RadarDNSSummaryGetIPVersionParamsQueryTypeIxfr       RadarDNSSummaryGetIPVersionParamsQueryType = "IXFR"
	RadarDNSSummaryGetIPVersionParamsQueryTypeKey        RadarDNSSummaryGetIPVersionParamsQueryType = "KEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeKx         RadarDNSSummaryGetIPVersionParamsQueryType = "KX"
	RadarDNSSummaryGetIPVersionParamsQueryTypeL32        RadarDNSSummaryGetIPVersionParamsQueryType = "L32"
	RadarDNSSummaryGetIPVersionParamsQueryTypeL64        RadarDNSSummaryGetIPVersionParamsQueryType = "L64"
	RadarDNSSummaryGetIPVersionParamsQueryTypeLoc        RadarDNSSummaryGetIPVersionParamsQueryType = "LOC"
	RadarDNSSummaryGetIPVersionParamsQueryTypeLp         RadarDNSSummaryGetIPVersionParamsQueryType = "LP"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMaila      RadarDNSSummaryGetIPVersionParamsQueryType = "MAILA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMailb      RadarDNSSummaryGetIPVersionParamsQueryType = "MAILB"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMB         RadarDNSSummaryGetIPVersionParamsQueryType = "MB"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMd         RadarDNSSummaryGetIPVersionParamsQueryType = "MD"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMf         RadarDNSSummaryGetIPVersionParamsQueryType = "MF"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMg         RadarDNSSummaryGetIPVersionParamsQueryType = "MG"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMinfo      RadarDNSSummaryGetIPVersionParamsQueryType = "MINFO"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMr         RadarDNSSummaryGetIPVersionParamsQueryType = "MR"
	RadarDNSSummaryGetIPVersionParamsQueryTypeMx         RadarDNSSummaryGetIPVersionParamsQueryType = "MX"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNaptr      RadarDNSSummaryGetIPVersionParamsQueryType = "NAPTR"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNb         RadarDNSSummaryGetIPVersionParamsQueryType = "NB"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNbstat     RadarDNSSummaryGetIPVersionParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNid        RadarDNSSummaryGetIPVersionParamsQueryType = "NID"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNimloc     RadarDNSSummaryGetIPVersionParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNinfo      RadarDNSSummaryGetIPVersionParamsQueryType = "NINFO"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNs         RadarDNSSummaryGetIPVersionParamsQueryType = "NS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNsap       RadarDNSSummaryGetIPVersionParamsQueryType = "NSAP"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNsec       RadarDNSSummaryGetIPVersionParamsQueryType = "NSEC"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNsec3      RadarDNSSummaryGetIPVersionParamsQueryType = "NSEC3"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNsec3Param RadarDNSSummaryGetIPVersionParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNull       RadarDNSSummaryGetIPVersionParamsQueryType = "NULL"
	RadarDNSSummaryGetIPVersionParamsQueryTypeNxt        RadarDNSSummaryGetIPVersionParamsQueryType = "NXT"
	RadarDNSSummaryGetIPVersionParamsQueryTypeOpenpgpkey RadarDNSSummaryGetIPVersionParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeOpt        RadarDNSSummaryGetIPVersionParamsQueryType = "OPT"
	RadarDNSSummaryGetIPVersionParamsQueryTypePtr        RadarDNSSummaryGetIPVersionParamsQueryType = "PTR"
	RadarDNSSummaryGetIPVersionParamsQueryTypePx         RadarDNSSummaryGetIPVersionParamsQueryType = "PX"
	RadarDNSSummaryGetIPVersionParamsQueryTypeRkey       RadarDNSSummaryGetIPVersionParamsQueryType = "RKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeRp         RadarDNSSummaryGetIPVersionParamsQueryType = "RP"
	RadarDNSSummaryGetIPVersionParamsQueryTypeRrsig      RadarDNSSummaryGetIPVersionParamsQueryType = "RRSIG"
	RadarDNSSummaryGetIPVersionParamsQueryTypeRt         RadarDNSSummaryGetIPVersionParamsQueryType = "RT"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSig        RadarDNSSummaryGetIPVersionParamsQueryType = "SIG"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSink       RadarDNSSummaryGetIPVersionParamsQueryType = "SINK"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSmimea     RadarDNSSummaryGetIPVersionParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSoa        RadarDNSSummaryGetIPVersionParamsQueryType = "SOA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSpf        RadarDNSSummaryGetIPVersionParamsQueryType = "SPF"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSrv        RadarDNSSummaryGetIPVersionParamsQueryType = "SRV"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSshfp      RadarDNSSummaryGetIPVersionParamsQueryType = "SSHFP"
	RadarDNSSummaryGetIPVersionParamsQueryTypeSvcb       RadarDNSSummaryGetIPVersionParamsQueryType = "SVCB"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTa         RadarDNSSummaryGetIPVersionParamsQueryType = "TA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTalink     RadarDNSSummaryGetIPVersionParamsQueryType = "TALINK"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTkey       RadarDNSSummaryGetIPVersionParamsQueryType = "TKEY"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTlsa       RadarDNSSummaryGetIPVersionParamsQueryType = "TLSA"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTsig       RadarDNSSummaryGetIPVersionParamsQueryType = "TSIG"
	RadarDNSSummaryGetIPVersionParamsQueryTypeTxt        RadarDNSSummaryGetIPVersionParamsQueryType = "TXT"
	RadarDNSSummaryGetIPVersionParamsQueryTypeUinfo      RadarDNSSummaryGetIPVersionParamsQueryType = "UINFO"
	RadarDNSSummaryGetIPVersionParamsQueryTypeUid        RadarDNSSummaryGetIPVersionParamsQueryType = "UID"
	RadarDNSSummaryGetIPVersionParamsQueryTypeUnspec     RadarDNSSummaryGetIPVersionParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetIPVersionParamsQueryTypeUri        RadarDNSSummaryGetIPVersionParamsQueryType = "URI"
	RadarDNSSummaryGetIPVersionParamsQueryTypeWks        RadarDNSSummaryGetIPVersionParamsQueryType = "WKS"
	RadarDNSSummaryGetIPVersionParamsQueryTypeX25        RadarDNSSummaryGetIPVersionParamsQueryType = "X25"
	RadarDNSSummaryGetIPVersionParamsQueryTypeZonemd     RadarDNSSummaryGetIPVersionParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetIPVersionParamsQueryTypeA, RadarDNSSummaryGetIPVersionParamsQueryTypeAaaa, RadarDNSSummaryGetIPVersionParamsQueryTypeA6, RadarDNSSummaryGetIPVersionParamsQueryTypeAfsdb, RadarDNSSummaryGetIPVersionParamsQueryTypeAny, RadarDNSSummaryGetIPVersionParamsQueryTypeApl, RadarDNSSummaryGetIPVersionParamsQueryTypeAtma, RadarDNSSummaryGetIPVersionParamsQueryTypeAxfr, RadarDNSSummaryGetIPVersionParamsQueryTypeCaa, RadarDNSSummaryGetIPVersionParamsQueryTypeCdnskey, RadarDNSSummaryGetIPVersionParamsQueryTypeCds, RadarDNSSummaryGetIPVersionParamsQueryTypeCert, RadarDNSSummaryGetIPVersionParamsQueryTypeCname, RadarDNSSummaryGetIPVersionParamsQueryTypeCsync, RadarDNSSummaryGetIPVersionParamsQueryTypeDhcid, RadarDNSSummaryGetIPVersionParamsQueryTypeDlv, RadarDNSSummaryGetIPVersionParamsQueryTypeDname, RadarDNSSummaryGetIPVersionParamsQueryTypeDnskey, RadarDNSSummaryGetIPVersionParamsQueryTypeDoa, RadarDNSSummaryGetIPVersionParamsQueryTypeDs, RadarDNSSummaryGetIPVersionParamsQueryTypeEid, RadarDNSSummaryGetIPVersionParamsQueryTypeEui48, RadarDNSSummaryGetIPVersionParamsQueryTypeEui64, RadarDNSSummaryGetIPVersionParamsQueryTypeGpos, RadarDNSSummaryGetIPVersionParamsQueryTypeGid, RadarDNSSummaryGetIPVersionParamsQueryTypeHinfo, RadarDNSSummaryGetIPVersionParamsQueryTypeHip, RadarDNSSummaryGetIPVersionParamsQueryTypeHTTPS, RadarDNSSummaryGetIPVersionParamsQueryTypeIpseckey, RadarDNSSummaryGetIPVersionParamsQueryTypeIsdn, RadarDNSSummaryGetIPVersionParamsQueryTypeIxfr, RadarDNSSummaryGetIPVersionParamsQueryTypeKey, RadarDNSSummaryGetIPVersionParamsQueryTypeKx, RadarDNSSummaryGetIPVersionParamsQueryTypeL32, RadarDNSSummaryGetIPVersionParamsQueryTypeL64, RadarDNSSummaryGetIPVersionParamsQueryTypeLoc, RadarDNSSummaryGetIPVersionParamsQueryTypeLp, RadarDNSSummaryGetIPVersionParamsQueryTypeMaila, RadarDNSSummaryGetIPVersionParamsQueryTypeMailb, RadarDNSSummaryGetIPVersionParamsQueryTypeMB, RadarDNSSummaryGetIPVersionParamsQueryTypeMd, RadarDNSSummaryGetIPVersionParamsQueryTypeMf, RadarDNSSummaryGetIPVersionParamsQueryTypeMg, RadarDNSSummaryGetIPVersionParamsQueryTypeMinfo, RadarDNSSummaryGetIPVersionParamsQueryTypeMr, RadarDNSSummaryGetIPVersionParamsQueryTypeMx, RadarDNSSummaryGetIPVersionParamsQueryTypeNaptr, RadarDNSSummaryGetIPVersionParamsQueryTypeNb, RadarDNSSummaryGetIPVersionParamsQueryTypeNbstat, RadarDNSSummaryGetIPVersionParamsQueryTypeNid, RadarDNSSummaryGetIPVersionParamsQueryTypeNimloc, RadarDNSSummaryGetIPVersionParamsQueryTypeNinfo, RadarDNSSummaryGetIPVersionParamsQueryTypeNs, RadarDNSSummaryGetIPVersionParamsQueryTypeNsap, RadarDNSSummaryGetIPVersionParamsQueryTypeNsec, RadarDNSSummaryGetIPVersionParamsQueryTypeNsec3, RadarDNSSummaryGetIPVersionParamsQueryTypeNsec3Param, RadarDNSSummaryGetIPVersionParamsQueryTypeNull, RadarDNSSummaryGetIPVersionParamsQueryTypeNxt, RadarDNSSummaryGetIPVersionParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetIPVersionParamsQueryTypeOpt, RadarDNSSummaryGetIPVersionParamsQueryTypePtr, RadarDNSSummaryGetIPVersionParamsQueryTypePx, RadarDNSSummaryGetIPVersionParamsQueryTypeRkey, RadarDNSSummaryGetIPVersionParamsQueryTypeRp, RadarDNSSummaryGetIPVersionParamsQueryTypeRrsig, RadarDNSSummaryGetIPVersionParamsQueryTypeRt, RadarDNSSummaryGetIPVersionParamsQueryTypeSig, RadarDNSSummaryGetIPVersionParamsQueryTypeSink, RadarDNSSummaryGetIPVersionParamsQueryTypeSmimea, RadarDNSSummaryGetIPVersionParamsQueryTypeSoa, RadarDNSSummaryGetIPVersionParamsQueryTypeSpf, RadarDNSSummaryGetIPVersionParamsQueryTypeSrv, RadarDNSSummaryGetIPVersionParamsQueryTypeSshfp, RadarDNSSummaryGetIPVersionParamsQueryTypeSvcb, RadarDNSSummaryGetIPVersionParamsQueryTypeTa, RadarDNSSummaryGetIPVersionParamsQueryTypeTalink, RadarDNSSummaryGetIPVersionParamsQueryTypeTkey, RadarDNSSummaryGetIPVersionParamsQueryTypeTlsa, RadarDNSSummaryGetIPVersionParamsQueryTypeTsig, RadarDNSSummaryGetIPVersionParamsQueryTypeTxt, RadarDNSSummaryGetIPVersionParamsQueryTypeUinfo, RadarDNSSummaryGetIPVersionParamsQueryTypeUid, RadarDNSSummaryGetIPVersionParamsQueryTypeUnspec, RadarDNSSummaryGetIPVersionParamsQueryTypeUri, RadarDNSSummaryGetIPVersionParamsQueryTypeWks, RadarDNSSummaryGetIPVersionParamsQueryTypeX25, RadarDNSSummaryGetIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetIPVersionParamsResponseCode string

const (
	RadarDNSSummaryGetIPVersionParamsResponseCodeNoerror   RadarDNSSummaryGetIPVersionParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetIPVersionParamsResponseCodeFormerr   RadarDNSSummaryGetIPVersionParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetIPVersionParamsResponseCodeServfail  RadarDNSSummaryGetIPVersionParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetIPVersionParamsResponseCodeNxdomain  RadarDNSSummaryGetIPVersionParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetIPVersionParamsResponseCodeNotimp    RadarDNSSummaryGetIPVersionParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetIPVersionParamsResponseCodeRefused   RadarDNSSummaryGetIPVersionParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetIPVersionParamsResponseCodeYxdomain  RadarDNSSummaryGetIPVersionParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetIPVersionParamsResponseCodeYxrrset   RadarDNSSummaryGetIPVersionParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetIPVersionParamsResponseCodeNxrrset   RadarDNSSummaryGetIPVersionParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetIPVersionParamsResponseCodeNotauth   RadarDNSSummaryGetIPVersionParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetIPVersionParamsResponseCodeNotzone   RadarDNSSummaryGetIPVersionParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadsig    RadarDNSSummaryGetIPVersionParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadkey    RadarDNSSummaryGetIPVersionParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadtime   RadarDNSSummaryGetIPVersionParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadmode   RadarDNSSummaryGetIPVersionParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadname   RadarDNSSummaryGetIPVersionParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadalg    RadarDNSSummaryGetIPVersionParamsResponseCode = "BADALG"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadtrunc  RadarDNSSummaryGetIPVersionParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetIPVersionParamsResponseCodeBadcookie RadarDNSSummaryGetIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetIPVersionParamsResponseCodeNoerror, RadarDNSSummaryGetIPVersionParamsResponseCodeFormerr, RadarDNSSummaryGetIPVersionParamsResponseCodeServfail, RadarDNSSummaryGetIPVersionParamsResponseCodeNxdomain, RadarDNSSummaryGetIPVersionParamsResponseCodeNotimp, RadarDNSSummaryGetIPVersionParamsResponseCodeRefused, RadarDNSSummaryGetIPVersionParamsResponseCodeYxdomain, RadarDNSSummaryGetIPVersionParamsResponseCodeYxrrset, RadarDNSSummaryGetIPVersionParamsResponseCodeNxrrset, RadarDNSSummaryGetIPVersionParamsResponseCodeNotauth, RadarDNSSummaryGetIPVersionParamsResponseCodeNotzone, RadarDNSSummaryGetIPVersionParamsResponseCodeBadsig, RadarDNSSummaryGetIPVersionParamsResponseCodeBadkey, RadarDNSSummaryGetIPVersionParamsResponseCodeBadtime, RadarDNSSummaryGetIPVersionParamsResponseCodeBadmode, RadarDNSSummaryGetIPVersionParamsResponseCodeBadname, RadarDNSSummaryGetIPVersionParamsResponseCodeBadalg, RadarDNSSummaryGetIPVersionParamsResponseCodeBadtrunc, RadarDNSSummaryGetIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetMatchingAnswerParams struct {
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
	Format param.Field[RadarDNSSummaryGetMatchingAnswerParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetMatchingAnswerParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetMatchingAnswerParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetMatchingAnswerParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetMatchingAnswerParams]'s query parameters
// as `url.Values`.
func (r RadarDNSSummaryGetMatchingAnswerParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetMatchingAnswerParamsFormat string

const (
	RadarDNSSummaryGetMatchingAnswerParamsFormatJson RadarDNSSummaryGetMatchingAnswerParamsFormat = "JSON"
	RadarDNSSummaryGetMatchingAnswerParamsFormatCsv  RadarDNSSummaryGetMatchingAnswerParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetMatchingAnswerParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetMatchingAnswerParamsFormatJson, RadarDNSSummaryGetMatchingAnswerParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetMatchingAnswerParamsProtocol string

const (
	RadarDNSSummaryGetMatchingAnswerParamsProtocolUdp   RadarDNSSummaryGetMatchingAnswerParamsProtocol = "UDP"
	RadarDNSSummaryGetMatchingAnswerParamsProtocolTcp   RadarDNSSummaryGetMatchingAnswerParamsProtocol = "TCP"
	RadarDNSSummaryGetMatchingAnswerParamsProtocolHTTPS RadarDNSSummaryGetMatchingAnswerParamsProtocol = "HTTPS"
	RadarDNSSummaryGetMatchingAnswerParamsProtocolTls   RadarDNSSummaryGetMatchingAnswerParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetMatchingAnswerParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetMatchingAnswerParamsProtocolUdp, RadarDNSSummaryGetMatchingAnswerParamsProtocolTcp, RadarDNSSummaryGetMatchingAnswerParamsProtocolHTTPS, RadarDNSSummaryGetMatchingAnswerParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetMatchingAnswerParamsQueryType string

const (
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeA          RadarDNSSummaryGetMatchingAnswerParamsQueryType = "A"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAaaa       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "AAAA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeA6         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "A6"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAfsdb      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "AFSDB"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAny        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "ANY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeApl        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "APL"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAtma       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "ATMA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAxfr       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "AXFR"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCaa        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CAA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCdnskey    RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCds        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CDS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCert       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CERT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCname      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CNAME"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCsync      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "CSYNC"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDhcid      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DHCID"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDlv        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DLV"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDname      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DNAME"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDnskey     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDoa        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DOA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDs         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "DS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEid        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "EID"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEui48      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "EUI48"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEui64      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "EUI64"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeGpos       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "GPOS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeGid        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "GID"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHinfo      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "HINFO"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHip        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "HIP"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHTTPS      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "HTTPS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIpseckey   RadarDNSSummaryGetMatchingAnswerParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIsdn       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "ISDN"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIxfr       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "IXFR"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeKey        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "KEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeKx         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "KX"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeL32        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "L32"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeL64        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "L64"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeLoc        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "LOC"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeLp         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "LP"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMaila      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MAILA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMailb      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MAILB"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMB         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MB"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMd         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MD"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMf         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MF"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMg         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MG"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMinfo      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MINFO"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMr         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MR"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMx         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "MX"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNaptr      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NAPTR"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNb         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NB"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNbstat     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNid        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NID"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNimloc     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNinfo      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NINFO"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNs         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsap       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NSAP"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NSEC"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec3      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NSEC3"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec3Param RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNull       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NULL"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNxt        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "NXT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeOpenpgpkey RadarDNSSummaryGetMatchingAnswerParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeOpt        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "OPT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypePtr        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "PTR"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypePx         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "PX"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRkey       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "RKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRp         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "RP"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRrsig      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "RRSIG"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRt         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "RT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSig        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SIG"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSink       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SINK"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSmimea     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSoa        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SOA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSpf        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SPF"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSrv        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SRV"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSshfp      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SSHFP"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSvcb       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "SVCB"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTa         RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTalink     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TALINK"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTkey       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TKEY"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTlsa       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TLSA"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTsig       RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TSIG"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTxt        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "TXT"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUinfo      RadarDNSSummaryGetMatchingAnswerParamsQueryType = "UINFO"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUid        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "UID"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUnspec     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUri        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "URI"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeWks        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "WKS"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeX25        RadarDNSSummaryGetMatchingAnswerParamsQueryType = "X25"
	RadarDNSSummaryGetMatchingAnswerParamsQueryTypeZonemd     RadarDNSSummaryGetMatchingAnswerParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetMatchingAnswerParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetMatchingAnswerParamsQueryTypeA, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAaaa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeA6, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAfsdb, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAny, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeApl, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAtma, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeAxfr, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCaa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCdnskey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCds, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCert, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCname, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeCsync, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDhcid, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDlv, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDname, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDnskey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDoa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeDs, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEid, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEui48, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeEui64, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeGpos, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeGid, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHinfo, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHip, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeHTTPS, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIpseckey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIsdn, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeIxfr, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeKey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeKx, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeL32, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeL64, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeLoc, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeLp, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMaila, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMailb, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMB, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMd, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMf, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMg, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMinfo, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMr, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeMx, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNaptr, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNb, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNbstat, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNid, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNimloc, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNinfo, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNs, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsap, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec3, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNsec3Param, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNull, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeNxt, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeOpt, RadarDNSSummaryGetMatchingAnswerParamsQueryTypePtr, RadarDNSSummaryGetMatchingAnswerParamsQueryTypePx, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRkey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRp, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRrsig, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeRt, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSig, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSink, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSmimea, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSoa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSpf, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSrv, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSshfp, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeSvcb, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTalink, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTkey, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTlsa, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTsig, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeTxt, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUinfo, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUid, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUnspec, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeUri, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeWks, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeX25, RadarDNSSummaryGetMatchingAnswerParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetMatchingAnswerParamsResponseCode string

const (
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNoerror   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeFormerr   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeServfail  RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNxdomain  RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotimp    RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeRefused   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeYxdomain  RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeYxrrset   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNxrrset   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotauth   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotzone   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadsig    RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadkey    RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadtime   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadmode   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadname   RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadalg    RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADALG"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadtrunc  RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadcookie RadarDNSSummaryGetMatchingAnswerParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetMatchingAnswerParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNoerror, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeFormerr, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeServfail, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNxdomain, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotimp, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeRefused, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeYxdomain, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeYxrrset, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNxrrset, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotauth, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeNotzone, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadsig, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadkey, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadtime, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadmode, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadname, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadalg, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadtrunc, RadarDNSSummaryGetMatchingAnswerParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetProtocolParams struct {
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
	Format param.Field[RadarDNSSummaryGetProtocolParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetProtocolParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetProtocolParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetProtocolParamsFormat string

const (
	RadarDNSSummaryGetProtocolParamsFormatJson RadarDNSSummaryGetProtocolParamsFormat = "JSON"
	RadarDNSSummaryGetProtocolParamsFormatCsv  RadarDNSSummaryGetProtocolParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetProtocolParamsFormatJson, RadarDNSSummaryGetProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetProtocolParamsQueryType string

const (
	RadarDNSSummaryGetProtocolParamsQueryTypeA          RadarDNSSummaryGetProtocolParamsQueryType = "A"
	RadarDNSSummaryGetProtocolParamsQueryTypeAaaa       RadarDNSSummaryGetProtocolParamsQueryType = "AAAA"
	RadarDNSSummaryGetProtocolParamsQueryTypeA6         RadarDNSSummaryGetProtocolParamsQueryType = "A6"
	RadarDNSSummaryGetProtocolParamsQueryTypeAfsdb      RadarDNSSummaryGetProtocolParamsQueryType = "AFSDB"
	RadarDNSSummaryGetProtocolParamsQueryTypeAny        RadarDNSSummaryGetProtocolParamsQueryType = "ANY"
	RadarDNSSummaryGetProtocolParamsQueryTypeApl        RadarDNSSummaryGetProtocolParamsQueryType = "APL"
	RadarDNSSummaryGetProtocolParamsQueryTypeAtma       RadarDNSSummaryGetProtocolParamsQueryType = "ATMA"
	RadarDNSSummaryGetProtocolParamsQueryTypeAxfr       RadarDNSSummaryGetProtocolParamsQueryType = "AXFR"
	RadarDNSSummaryGetProtocolParamsQueryTypeCaa        RadarDNSSummaryGetProtocolParamsQueryType = "CAA"
	RadarDNSSummaryGetProtocolParamsQueryTypeCdnskey    RadarDNSSummaryGetProtocolParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeCds        RadarDNSSummaryGetProtocolParamsQueryType = "CDS"
	RadarDNSSummaryGetProtocolParamsQueryTypeCert       RadarDNSSummaryGetProtocolParamsQueryType = "CERT"
	RadarDNSSummaryGetProtocolParamsQueryTypeCname      RadarDNSSummaryGetProtocolParamsQueryType = "CNAME"
	RadarDNSSummaryGetProtocolParamsQueryTypeCsync      RadarDNSSummaryGetProtocolParamsQueryType = "CSYNC"
	RadarDNSSummaryGetProtocolParamsQueryTypeDhcid      RadarDNSSummaryGetProtocolParamsQueryType = "DHCID"
	RadarDNSSummaryGetProtocolParamsQueryTypeDlv        RadarDNSSummaryGetProtocolParamsQueryType = "DLV"
	RadarDNSSummaryGetProtocolParamsQueryTypeDname      RadarDNSSummaryGetProtocolParamsQueryType = "DNAME"
	RadarDNSSummaryGetProtocolParamsQueryTypeDnskey     RadarDNSSummaryGetProtocolParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeDoa        RadarDNSSummaryGetProtocolParamsQueryType = "DOA"
	RadarDNSSummaryGetProtocolParamsQueryTypeDs         RadarDNSSummaryGetProtocolParamsQueryType = "DS"
	RadarDNSSummaryGetProtocolParamsQueryTypeEid        RadarDNSSummaryGetProtocolParamsQueryType = "EID"
	RadarDNSSummaryGetProtocolParamsQueryTypeEui48      RadarDNSSummaryGetProtocolParamsQueryType = "EUI48"
	RadarDNSSummaryGetProtocolParamsQueryTypeEui64      RadarDNSSummaryGetProtocolParamsQueryType = "EUI64"
	RadarDNSSummaryGetProtocolParamsQueryTypeGpos       RadarDNSSummaryGetProtocolParamsQueryType = "GPOS"
	RadarDNSSummaryGetProtocolParamsQueryTypeGid        RadarDNSSummaryGetProtocolParamsQueryType = "GID"
	RadarDNSSummaryGetProtocolParamsQueryTypeHinfo      RadarDNSSummaryGetProtocolParamsQueryType = "HINFO"
	RadarDNSSummaryGetProtocolParamsQueryTypeHip        RadarDNSSummaryGetProtocolParamsQueryType = "HIP"
	RadarDNSSummaryGetProtocolParamsQueryTypeHTTPS      RadarDNSSummaryGetProtocolParamsQueryType = "HTTPS"
	RadarDNSSummaryGetProtocolParamsQueryTypeIpseckey   RadarDNSSummaryGetProtocolParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeIsdn       RadarDNSSummaryGetProtocolParamsQueryType = "ISDN"
	RadarDNSSummaryGetProtocolParamsQueryTypeIxfr       RadarDNSSummaryGetProtocolParamsQueryType = "IXFR"
	RadarDNSSummaryGetProtocolParamsQueryTypeKey        RadarDNSSummaryGetProtocolParamsQueryType = "KEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeKx         RadarDNSSummaryGetProtocolParamsQueryType = "KX"
	RadarDNSSummaryGetProtocolParamsQueryTypeL32        RadarDNSSummaryGetProtocolParamsQueryType = "L32"
	RadarDNSSummaryGetProtocolParamsQueryTypeL64        RadarDNSSummaryGetProtocolParamsQueryType = "L64"
	RadarDNSSummaryGetProtocolParamsQueryTypeLoc        RadarDNSSummaryGetProtocolParamsQueryType = "LOC"
	RadarDNSSummaryGetProtocolParamsQueryTypeLp         RadarDNSSummaryGetProtocolParamsQueryType = "LP"
	RadarDNSSummaryGetProtocolParamsQueryTypeMaila      RadarDNSSummaryGetProtocolParamsQueryType = "MAILA"
	RadarDNSSummaryGetProtocolParamsQueryTypeMailb      RadarDNSSummaryGetProtocolParamsQueryType = "MAILB"
	RadarDNSSummaryGetProtocolParamsQueryTypeMB         RadarDNSSummaryGetProtocolParamsQueryType = "MB"
	RadarDNSSummaryGetProtocolParamsQueryTypeMd         RadarDNSSummaryGetProtocolParamsQueryType = "MD"
	RadarDNSSummaryGetProtocolParamsQueryTypeMf         RadarDNSSummaryGetProtocolParamsQueryType = "MF"
	RadarDNSSummaryGetProtocolParamsQueryTypeMg         RadarDNSSummaryGetProtocolParamsQueryType = "MG"
	RadarDNSSummaryGetProtocolParamsQueryTypeMinfo      RadarDNSSummaryGetProtocolParamsQueryType = "MINFO"
	RadarDNSSummaryGetProtocolParamsQueryTypeMr         RadarDNSSummaryGetProtocolParamsQueryType = "MR"
	RadarDNSSummaryGetProtocolParamsQueryTypeMx         RadarDNSSummaryGetProtocolParamsQueryType = "MX"
	RadarDNSSummaryGetProtocolParamsQueryTypeNaptr      RadarDNSSummaryGetProtocolParamsQueryType = "NAPTR"
	RadarDNSSummaryGetProtocolParamsQueryTypeNb         RadarDNSSummaryGetProtocolParamsQueryType = "NB"
	RadarDNSSummaryGetProtocolParamsQueryTypeNbstat     RadarDNSSummaryGetProtocolParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetProtocolParamsQueryTypeNid        RadarDNSSummaryGetProtocolParamsQueryType = "NID"
	RadarDNSSummaryGetProtocolParamsQueryTypeNimloc     RadarDNSSummaryGetProtocolParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetProtocolParamsQueryTypeNinfo      RadarDNSSummaryGetProtocolParamsQueryType = "NINFO"
	RadarDNSSummaryGetProtocolParamsQueryTypeNs         RadarDNSSummaryGetProtocolParamsQueryType = "NS"
	RadarDNSSummaryGetProtocolParamsQueryTypeNsap       RadarDNSSummaryGetProtocolParamsQueryType = "NSAP"
	RadarDNSSummaryGetProtocolParamsQueryTypeNsec       RadarDNSSummaryGetProtocolParamsQueryType = "NSEC"
	RadarDNSSummaryGetProtocolParamsQueryTypeNsec3      RadarDNSSummaryGetProtocolParamsQueryType = "NSEC3"
	RadarDNSSummaryGetProtocolParamsQueryTypeNsec3Param RadarDNSSummaryGetProtocolParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetProtocolParamsQueryTypeNull       RadarDNSSummaryGetProtocolParamsQueryType = "NULL"
	RadarDNSSummaryGetProtocolParamsQueryTypeNxt        RadarDNSSummaryGetProtocolParamsQueryType = "NXT"
	RadarDNSSummaryGetProtocolParamsQueryTypeOpenpgpkey RadarDNSSummaryGetProtocolParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeOpt        RadarDNSSummaryGetProtocolParamsQueryType = "OPT"
	RadarDNSSummaryGetProtocolParamsQueryTypePtr        RadarDNSSummaryGetProtocolParamsQueryType = "PTR"
	RadarDNSSummaryGetProtocolParamsQueryTypePx         RadarDNSSummaryGetProtocolParamsQueryType = "PX"
	RadarDNSSummaryGetProtocolParamsQueryTypeRkey       RadarDNSSummaryGetProtocolParamsQueryType = "RKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeRp         RadarDNSSummaryGetProtocolParamsQueryType = "RP"
	RadarDNSSummaryGetProtocolParamsQueryTypeRrsig      RadarDNSSummaryGetProtocolParamsQueryType = "RRSIG"
	RadarDNSSummaryGetProtocolParamsQueryTypeRt         RadarDNSSummaryGetProtocolParamsQueryType = "RT"
	RadarDNSSummaryGetProtocolParamsQueryTypeSig        RadarDNSSummaryGetProtocolParamsQueryType = "SIG"
	RadarDNSSummaryGetProtocolParamsQueryTypeSink       RadarDNSSummaryGetProtocolParamsQueryType = "SINK"
	RadarDNSSummaryGetProtocolParamsQueryTypeSmimea     RadarDNSSummaryGetProtocolParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetProtocolParamsQueryTypeSoa        RadarDNSSummaryGetProtocolParamsQueryType = "SOA"
	RadarDNSSummaryGetProtocolParamsQueryTypeSpf        RadarDNSSummaryGetProtocolParamsQueryType = "SPF"
	RadarDNSSummaryGetProtocolParamsQueryTypeSrv        RadarDNSSummaryGetProtocolParamsQueryType = "SRV"
	RadarDNSSummaryGetProtocolParamsQueryTypeSshfp      RadarDNSSummaryGetProtocolParamsQueryType = "SSHFP"
	RadarDNSSummaryGetProtocolParamsQueryTypeSvcb       RadarDNSSummaryGetProtocolParamsQueryType = "SVCB"
	RadarDNSSummaryGetProtocolParamsQueryTypeTa         RadarDNSSummaryGetProtocolParamsQueryType = "TA"
	RadarDNSSummaryGetProtocolParamsQueryTypeTalink     RadarDNSSummaryGetProtocolParamsQueryType = "TALINK"
	RadarDNSSummaryGetProtocolParamsQueryTypeTkey       RadarDNSSummaryGetProtocolParamsQueryType = "TKEY"
	RadarDNSSummaryGetProtocolParamsQueryTypeTlsa       RadarDNSSummaryGetProtocolParamsQueryType = "TLSA"
	RadarDNSSummaryGetProtocolParamsQueryTypeTsig       RadarDNSSummaryGetProtocolParamsQueryType = "TSIG"
	RadarDNSSummaryGetProtocolParamsQueryTypeTxt        RadarDNSSummaryGetProtocolParamsQueryType = "TXT"
	RadarDNSSummaryGetProtocolParamsQueryTypeUinfo      RadarDNSSummaryGetProtocolParamsQueryType = "UINFO"
	RadarDNSSummaryGetProtocolParamsQueryTypeUid        RadarDNSSummaryGetProtocolParamsQueryType = "UID"
	RadarDNSSummaryGetProtocolParamsQueryTypeUnspec     RadarDNSSummaryGetProtocolParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetProtocolParamsQueryTypeUri        RadarDNSSummaryGetProtocolParamsQueryType = "URI"
	RadarDNSSummaryGetProtocolParamsQueryTypeWks        RadarDNSSummaryGetProtocolParamsQueryType = "WKS"
	RadarDNSSummaryGetProtocolParamsQueryTypeX25        RadarDNSSummaryGetProtocolParamsQueryType = "X25"
	RadarDNSSummaryGetProtocolParamsQueryTypeZonemd     RadarDNSSummaryGetProtocolParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetProtocolParamsQueryTypeA, RadarDNSSummaryGetProtocolParamsQueryTypeAaaa, RadarDNSSummaryGetProtocolParamsQueryTypeA6, RadarDNSSummaryGetProtocolParamsQueryTypeAfsdb, RadarDNSSummaryGetProtocolParamsQueryTypeAny, RadarDNSSummaryGetProtocolParamsQueryTypeApl, RadarDNSSummaryGetProtocolParamsQueryTypeAtma, RadarDNSSummaryGetProtocolParamsQueryTypeAxfr, RadarDNSSummaryGetProtocolParamsQueryTypeCaa, RadarDNSSummaryGetProtocolParamsQueryTypeCdnskey, RadarDNSSummaryGetProtocolParamsQueryTypeCds, RadarDNSSummaryGetProtocolParamsQueryTypeCert, RadarDNSSummaryGetProtocolParamsQueryTypeCname, RadarDNSSummaryGetProtocolParamsQueryTypeCsync, RadarDNSSummaryGetProtocolParamsQueryTypeDhcid, RadarDNSSummaryGetProtocolParamsQueryTypeDlv, RadarDNSSummaryGetProtocolParamsQueryTypeDname, RadarDNSSummaryGetProtocolParamsQueryTypeDnskey, RadarDNSSummaryGetProtocolParamsQueryTypeDoa, RadarDNSSummaryGetProtocolParamsQueryTypeDs, RadarDNSSummaryGetProtocolParamsQueryTypeEid, RadarDNSSummaryGetProtocolParamsQueryTypeEui48, RadarDNSSummaryGetProtocolParamsQueryTypeEui64, RadarDNSSummaryGetProtocolParamsQueryTypeGpos, RadarDNSSummaryGetProtocolParamsQueryTypeGid, RadarDNSSummaryGetProtocolParamsQueryTypeHinfo, RadarDNSSummaryGetProtocolParamsQueryTypeHip, RadarDNSSummaryGetProtocolParamsQueryTypeHTTPS, RadarDNSSummaryGetProtocolParamsQueryTypeIpseckey, RadarDNSSummaryGetProtocolParamsQueryTypeIsdn, RadarDNSSummaryGetProtocolParamsQueryTypeIxfr, RadarDNSSummaryGetProtocolParamsQueryTypeKey, RadarDNSSummaryGetProtocolParamsQueryTypeKx, RadarDNSSummaryGetProtocolParamsQueryTypeL32, RadarDNSSummaryGetProtocolParamsQueryTypeL64, RadarDNSSummaryGetProtocolParamsQueryTypeLoc, RadarDNSSummaryGetProtocolParamsQueryTypeLp, RadarDNSSummaryGetProtocolParamsQueryTypeMaila, RadarDNSSummaryGetProtocolParamsQueryTypeMailb, RadarDNSSummaryGetProtocolParamsQueryTypeMB, RadarDNSSummaryGetProtocolParamsQueryTypeMd, RadarDNSSummaryGetProtocolParamsQueryTypeMf, RadarDNSSummaryGetProtocolParamsQueryTypeMg, RadarDNSSummaryGetProtocolParamsQueryTypeMinfo, RadarDNSSummaryGetProtocolParamsQueryTypeMr, RadarDNSSummaryGetProtocolParamsQueryTypeMx, RadarDNSSummaryGetProtocolParamsQueryTypeNaptr, RadarDNSSummaryGetProtocolParamsQueryTypeNb, RadarDNSSummaryGetProtocolParamsQueryTypeNbstat, RadarDNSSummaryGetProtocolParamsQueryTypeNid, RadarDNSSummaryGetProtocolParamsQueryTypeNimloc, RadarDNSSummaryGetProtocolParamsQueryTypeNinfo, RadarDNSSummaryGetProtocolParamsQueryTypeNs, RadarDNSSummaryGetProtocolParamsQueryTypeNsap, RadarDNSSummaryGetProtocolParamsQueryTypeNsec, RadarDNSSummaryGetProtocolParamsQueryTypeNsec3, RadarDNSSummaryGetProtocolParamsQueryTypeNsec3Param, RadarDNSSummaryGetProtocolParamsQueryTypeNull, RadarDNSSummaryGetProtocolParamsQueryTypeNxt, RadarDNSSummaryGetProtocolParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetProtocolParamsQueryTypeOpt, RadarDNSSummaryGetProtocolParamsQueryTypePtr, RadarDNSSummaryGetProtocolParamsQueryTypePx, RadarDNSSummaryGetProtocolParamsQueryTypeRkey, RadarDNSSummaryGetProtocolParamsQueryTypeRp, RadarDNSSummaryGetProtocolParamsQueryTypeRrsig, RadarDNSSummaryGetProtocolParamsQueryTypeRt, RadarDNSSummaryGetProtocolParamsQueryTypeSig, RadarDNSSummaryGetProtocolParamsQueryTypeSink, RadarDNSSummaryGetProtocolParamsQueryTypeSmimea, RadarDNSSummaryGetProtocolParamsQueryTypeSoa, RadarDNSSummaryGetProtocolParamsQueryTypeSpf, RadarDNSSummaryGetProtocolParamsQueryTypeSrv, RadarDNSSummaryGetProtocolParamsQueryTypeSshfp, RadarDNSSummaryGetProtocolParamsQueryTypeSvcb, RadarDNSSummaryGetProtocolParamsQueryTypeTa, RadarDNSSummaryGetProtocolParamsQueryTypeTalink, RadarDNSSummaryGetProtocolParamsQueryTypeTkey, RadarDNSSummaryGetProtocolParamsQueryTypeTlsa, RadarDNSSummaryGetProtocolParamsQueryTypeTsig, RadarDNSSummaryGetProtocolParamsQueryTypeTxt, RadarDNSSummaryGetProtocolParamsQueryTypeUinfo, RadarDNSSummaryGetProtocolParamsQueryTypeUid, RadarDNSSummaryGetProtocolParamsQueryTypeUnspec, RadarDNSSummaryGetProtocolParamsQueryTypeUri, RadarDNSSummaryGetProtocolParamsQueryTypeWks, RadarDNSSummaryGetProtocolParamsQueryTypeX25, RadarDNSSummaryGetProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetProtocolParamsResponseCode string

const (
	RadarDNSSummaryGetProtocolParamsResponseCodeNoerror   RadarDNSSummaryGetProtocolParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetProtocolParamsResponseCodeFormerr   RadarDNSSummaryGetProtocolParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetProtocolParamsResponseCodeServfail  RadarDNSSummaryGetProtocolParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetProtocolParamsResponseCodeNxdomain  RadarDNSSummaryGetProtocolParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetProtocolParamsResponseCodeNotimp    RadarDNSSummaryGetProtocolParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetProtocolParamsResponseCodeRefused   RadarDNSSummaryGetProtocolParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetProtocolParamsResponseCodeYxdomain  RadarDNSSummaryGetProtocolParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetProtocolParamsResponseCodeYxrrset   RadarDNSSummaryGetProtocolParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetProtocolParamsResponseCodeNxrrset   RadarDNSSummaryGetProtocolParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetProtocolParamsResponseCodeNotauth   RadarDNSSummaryGetProtocolParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetProtocolParamsResponseCodeNotzone   RadarDNSSummaryGetProtocolParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadsig    RadarDNSSummaryGetProtocolParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadkey    RadarDNSSummaryGetProtocolParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadtime   RadarDNSSummaryGetProtocolParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadmode   RadarDNSSummaryGetProtocolParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadname   RadarDNSSummaryGetProtocolParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadalg    RadarDNSSummaryGetProtocolParamsResponseCode = "BADALG"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadtrunc  RadarDNSSummaryGetProtocolParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetProtocolParamsResponseCodeBadcookie RadarDNSSummaryGetProtocolParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetProtocolParamsResponseCodeNoerror, RadarDNSSummaryGetProtocolParamsResponseCodeFormerr, RadarDNSSummaryGetProtocolParamsResponseCodeServfail, RadarDNSSummaryGetProtocolParamsResponseCodeNxdomain, RadarDNSSummaryGetProtocolParamsResponseCodeNotimp, RadarDNSSummaryGetProtocolParamsResponseCodeRefused, RadarDNSSummaryGetProtocolParamsResponseCodeYxdomain, RadarDNSSummaryGetProtocolParamsResponseCodeYxrrset, RadarDNSSummaryGetProtocolParamsResponseCodeNxrrset, RadarDNSSummaryGetProtocolParamsResponseCodeNotauth, RadarDNSSummaryGetProtocolParamsResponseCodeNotzone, RadarDNSSummaryGetProtocolParamsResponseCodeBadsig, RadarDNSSummaryGetProtocolParamsResponseCodeBadkey, RadarDNSSummaryGetProtocolParamsResponseCodeBadtime, RadarDNSSummaryGetProtocolParamsResponseCodeBadmode, RadarDNSSummaryGetProtocolParamsResponseCodeBadname, RadarDNSSummaryGetProtocolParamsResponseCodeBadalg, RadarDNSSummaryGetProtocolParamsResponseCodeBadtrunc, RadarDNSSummaryGetProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetQueryTypeParams struct {
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
	Format param.Field[RadarDNSSummaryGetQueryTypeParamsFormat] `query:"format"`
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
	Protocol param.Field[RadarDNSSummaryGetQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetQueryTypeParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetQueryTypeParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetQueryTypeParamsFormat string

const (
	RadarDNSSummaryGetQueryTypeParamsFormatJson RadarDNSSummaryGetQueryTypeParamsFormat = "JSON"
	RadarDNSSummaryGetQueryTypeParamsFormatCsv  RadarDNSSummaryGetQueryTypeParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetQueryTypeParamsFormatJson, RadarDNSSummaryGetQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetQueryTypeParamsProtocol string

const (
	RadarDNSSummaryGetQueryTypeParamsProtocolUdp   RadarDNSSummaryGetQueryTypeParamsProtocol = "UDP"
	RadarDNSSummaryGetQueryTypeParamsProtocolTcp   RadarDNSSummaryGetQueryTypeParamsProtocol = "TCP"
	RadarDNSSummaryGetQueryTypeParamsProtocolHTTPS RadarDNSSummaryGetQueryTypeParamsProtocol = "HTTPS"
	RadarDNSSummaryGetQueryTypeParamsProtocolTls   RadarDNSSummaryGetQueryTypeParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetQueryTypeParamsProtocolUdp, RadarDNSSummaryGetQueryTypeParamsProtocolTcp, RadarDNSSummaryGetQueryTypeParamsProtocolHTTPS, RadarDNSSummaryGetQueryTypeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetQueryTypeParamsResponseCode string

const (
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNoerror   RadarDNSSummaryGetQueryTypeParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeFormerr   RadarDNSSummaryGetQueryTypeParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeServfail  RadarDNSSummaryGetQueryTypeParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNxdomain  RadarDNSSummaryGetQueryTypeParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNotimp    RadarDNSSummaryGetQueryTypeParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeRefused   RadarDNSSummaryGetQueryTypeParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeYxdomain  RadarDNSSummaryGetQueryTypeParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeYxrrset   RadarDNSSummaryGetQueryTypeParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNxrrset   RadarDNSSummaryGetQueryTypeParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNotauth   RadarDNSSummaryGetQueryTypeParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeNotzone   RadarDNSSummaryGetQueryTypeParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadsig    RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadkey    RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadtime   RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadmode   RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadname   RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadalg    RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADALG"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadtrunc  RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetQueryTypeParamsResponseCodeBadcookie RadarDNSSummaryGetQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetQueryTypeParamsResponseCodeNoerror, RadarDNSSummaryGetQueryTypeParamsResponseCodeFormerr, RadarDNSSummaryGetQueryTypeParamsResponseCodeServfail, RadarDNSSummaryGetQueryTypeParamsResponseCodeNxdomain, RadarDNSSummaryGetQueryTypeParamsResponseCodeNotimp, RadarDNSSummaryGetQueryTypeParamsResponseCodeRefused, RadarDNSSummaryGetQueryTypeParamsResponseCodeYxdomain, RadarDNSSummaryGetQueryTypeParamsResponseCodeYxrrset, RadarDNSSummaryGetQueryTypeParamsResponseCodeNxrrset, RadarDNSSummaryGetQueryTypeParamsResponseCodeNotauth, RadarDNSSummaryGetQueryTypeParamsResponseCodeNotzone, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadsig, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadkey, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadtime, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadmode, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadname, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadalg, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadtrunc, RadarDNSSummaryGetQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarDNSSummaryGetResponseCodeParams struct {
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
	Format param.Field[RadarDNSSummaryGetResponseCodeParamsFormat] `query:"format"`
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
	Protocol param.Field[RadarDNSSummaryGetResponseCodeParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetResponseCodeParamsQueryType] `query:"queryType"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetResponseCodeParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetResponseCodeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetResponseCodeParamsFormat string

const (
	RadarDNSSummaryGetResponseCodeParamsFormatJson RadarDNSSummaryGetResponseCodeParamsFormat = "JSON"
	RadarDNSSummaryGetResponseCodeParamsFormatCsv  RadarDNSSummaryGetResponseCodeParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetResponseCodeParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseCodeParamsFormatJson, RadarDNSSummaryGetResponseCodeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetResponseCodeParamsProtocol string

const (
	RadarDNSSummaryGetResponseCodeParamsProtocolUdp   RadarDNSSummaryGetResponseCodeParamsProtocol = "UDP"
	RadarDNSSummaryGetResponseCodeParamsProtocolTcp   RadarDNSSummaryGetResponseCodeParamsProtocol = "TCP"
	RadarDNSSummaryGetResponseCodeParamsProtocolHTTPS RadarDNSSummaryGetResponseCodeParamsProtocol = "HTTPS"
	RadarDNSSummaryGetResponseCodeParamsProtocolTls   RadarDNSSummaryGetResponseCodeParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetResponseCodeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseCodeParamsProtocolUdp, RadarDNSSummaryGetResponseCodeParamsProtocolTcp, RadarDNSSummaryGetResponseCodeParamsProtocolHTTPS, RadarDNSSummaryGetResponseCodeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetResponseCodeParamsQueryType string

const (
	RadarDNSSummaryGetResponseCodeParamsQueryTypeA          RadarDNSSummaryGetResponseCodeParamsQueryType = "A"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeAaaa       RadarDNSSummaryGetResponseCodeParamsQueryType = "AAAA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeA6         RadarDNSSummaryGetResponseCodeParamsQueryType = "A6"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeAfsdb      RadarDNSSummaryGetResponseCodeParamsQueryType = "AFSDB"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeAny        RadarDNSSummaryGetResponseCodeParamsQueryType = "ANY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeApl        RadarDNSSummaryGetResponseCodeParamsQueryType = "APL"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeAtma       RadarDNSSummaryGetResponseCodeParamsQueryType = "ATMA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeAxfr       RadarDNSSummaryGetResponseCodeParamsQueryType = "AXFR"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCaa        RadarDNSSummaryGetResponseCodeParamsQueryType = "CAA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCdnskey    RadarDNSSummaryGetResponseCodeParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCds        RadarDNSSummaryGetResponseCodeParamsQueryType = "CDS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCert       RadarDNSSummaryGetResponseCodeParamsQueryType = "CERT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCname      RadarDNSSummaryGetResponseCodeParamsQueryType = "CNAME"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeCsync      RadarDNSSummaryGetResponseCodeParamsQueryType = "CSYNC"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDhcid      RadarDNSSummaryGetResponseCodeParamsQueryType = "DHCID"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDlv        RadarDNSSummaryGetResponseCodeParamsQueryType = "DLV"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDname      RadarDNSSummaryGetResponseCodeParamsQueryType = "DNAME"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDnskey     RadarDNSSummaryGetResponseCodeParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDoa        RadarDNSSummaryGetResponseCodeParamsQueryType = "DOA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeDs         RadarDNSSummaryGetResponseCodeParamsQueryType = "DS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeEid        RadarDNSSummaryGetResponseCodeParamsQueryType = "EID"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeEui48      RadarDNSSummaryGetResponseCodeParamsQueryType = "EUI48"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeEui64      RadarDNSSummaryGetResponseCodeParamsQueryType = "EUI64"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeGpos       RadarDNSSummaryGetResponseCodeParamsQueryType = "GPOS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeGid        RadarDNSSummaryGetResponseCodeParamsQueryType = "GID"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeHinfo      RadarDNSSummaryGetResponseCodeParamsQueryType = "HINFO"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeHip        RadarDNSSummaryGetResponseCodeParamsQueryType = "HIP"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeHTTPS      RadarDNSSummaryGetResponseCodeParamsQueryType = "HTTPS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeIpseckey   RadarDNSSummaryGetResponseCodeParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeIsdn       RadarDNSSummaryGetResponseCodeParamsQueryType = "ISDN"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeIxfr       RadarDNSSummaryGetResponseCodeParamsQueryType = "IXFR"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeKey        RadarDNSSummaryGetResponseCodeParamsQueryType = "KEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeKx         RadarDNSSummaryGetResponseCodeParamsQueryType = "KX"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeL32        RadarDNSSummaryGetResponseCodeParamsQueryType = "L32"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeL64        RadarDNSSummaryGetResponseCodeParamsQueryType = "L64"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeLoc        RadarDNSSummaryGetResponseCodeParamsQueryType = "LOC"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeLp         RadarDNSSummaryGetResponseCodeParamsQueryType = "LP"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMaila      RadarDNSSummaryGetResponseCodeParamsQueryType = "MAILA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMailb      RadarDNSSummaryGetResponseCodeParamsQueryType = "MAILB"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMB         RadarDNSSummaryGetResponseCodeParamsQueryType = "MB"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMd         RadarDNSSummaryGetResponseCodeParamsQueryType = "MD"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMf         RadarDNSSummaryGetResponseCodeParamsQueryType = "MF"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMg         RadarDNSSummaryGetResponseCodeParamsQueryType = "MG"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMinfo      RadarDNSSummaryGetResponseCodeParamsQueryType = "MINFO"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMr         RadarDNSSummaryGetResponseCodeParamsQueryType = "MR"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeMx         RadarDNSSummaryGetResponseCodeParamsQueryType = "MX"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNaptr      RadarDNSSummaryGetResponseCodeParamsQueryType = "NAPTR"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNb         RadarDNSSummaryGetResponseCodeParamsQueryType = "NB"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNbstat     RadarDNSSummaryGetResponseCodeParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNid        RadarDNSSummaryGetResponseCodeParamsQueryType = "NID"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNimloc     RadarDNSSummaryGetResponseCodeParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNinfo      RadarDNSSummaryGetResponseCodeParamsQueryType = "NINFO"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNs         RadarDNSSummaryGetResponseCodeParamsQueryType = "NS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNsap       RadarDNSSummaryGetResponseCodeParamsQueryType = "NSAP"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec       RadarDNSSummaryGetResponseCodeParamsQueryType = "NSEC"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec3      RadarDNSSummaryGetResponseCodeParamsQueryType = "NSEC3"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec3Param RadarDNSSummaryGetResponseCodeParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNull       RadarDNSSummaryGetResponseCodeParamsQueryType = "NULL"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeNxt        RadarDNSSummaryGetResponseCodeParamsQueryType = "NXT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeOpenpgpkey RadarDNSSummaryGetResponseCodeParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeOpt        RadarDNSSummaryGetResponseCodeParamsQueryType = "OPT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypePtr        RadarDNSSummaryGetResponseCodeParamsQueryType = "PTR"
	RadarDNSSummaryGetResponseCodeParamsQueryTypePx         RadarDNSSummaryGetResponseCodeParamsQueryType = "PX"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeRkey       RadarDNSSummaryGetResponseCodeParamsQueryType = "RKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeRp         RadarDNSSummaryGetResponseCodeParamsQueryType = "RP"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeRrsig      RadarDNSSummaryGetResponseCodeParamsQueryType = "RRSIG"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeRt         RadarDNSSummaryGetResponseCodeParamsQueryType = "RT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSig        RadarDNSSummaryGetResponseCodeParamsQueryType = "SIG"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSink       RadarDNSSummaryGetResponseCodeParamsQueryType = "SINK"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSmimea     RadarDNSSummaryGetResponseCodeParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSoa        RadarDNSSummaryGetResponseCodeParamsQueryType = "SOA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSpf        RadarDNSSummaryGetResponseCodeParamsQueryType = "SPF"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSrv        RadarDNSSummaryGetResponseCodeParamsQueryType = "SRV"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSshfp      RadarDNSSummaryGetResponseCodeParamsQueryType = "SSHFP"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeSvcb       RadarDNSSummaryGetResponseCodeParamsQueryType = "SVCB"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTa         RadarDNSSummaryGetResponseCodeParamsQueryType = "TA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTalink     RadarDNSSummaryGetResponseCodeParamsQueryType = "TALINK"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTkey       RadarDNSSummaryGetResponseCodeParamsQueryType = "TKEY"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTlsa       RadarDNSSummaryGetResponseCodeParamsQueryType = "TLSA"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTsig       RadarDNSSummaryGetResponseCodeParamsQueryType = "TSIG"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeTxt        RadarDNSSummaryGetResponseCodeParamsQueryType = "TXT"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeUinfo      RadarDNSSummaryGetResponseCodeParamsQueryType = "UINFO"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeUid        RadarDNSSummaryGetResponseCodeParamsQueryType = "UID"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeUnspec     RadarDNSSummaryGetResponseCodeParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeUri        RadarDNSSummaryGetResponseCodeParamsQueryType = "URI"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeWks        RadarDNSSummaryGetResponseCodeParamsQueryType = "WKS"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeX25        RadarDNSSummaryGetResponseCodeParamsQueryType = "X25"
	RadarDNSSummaryGetResponseCodeParamsQueryTypeZonemd     RadarDNSSummaryGetResponseCodeParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetResponseCodeParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseCodeParamsQueryTypeA, RadarDNSSummaryGetResponseCodeParamsQueryTypeAaaa, RadarDNSSummaryGetResponseCodeParamsQueryTypeA6, RadarDNSSummaryGetResponseCodeParamsQueryTypeAfsdb, RadarDNSSummaryGetResponseCodeParamsQueryTypeAny, RadarDNSSummaryGetResponseCodeParamsQueryTypeApl, RadarDNSSummaryGetResponseCodeParamsQueryTypeAtma, RadarDNSSummaryGetResponseCodeParamsQueryTypeAxfr, RadarDNSSummaryGetResponseCodeParamsQueryTypeCaa, RadarDNSSummaryGetResponseCodeParamsQueryTypeCdnskey, RadarDNSSummaryGetResponseCodeParamsQueryTypeCds, RadarDNSSummaryGetResponseCodeParamsQueryTypeCert, RadarDNSSummaryGetResponseCodeParamsQueryTypeCname, RadarDNSSummaryGetResponseCodeParamsQueryTypeCsync, RadarDNSSummaryGetResponseCodeParamsQueryTypeDhcid, RadarDNSSummaryGetResponseCodeParamsQueryTypeDlv, RadarDNSSummaryGetResponseCodeParamsQueryTypeDname, RadarDNSSummaryGetResponseCodeParamsQueryTypeDnskey, RadarDNSSummaryGetResponseCodeParamsQueryTypeDoa, RadarDNSSummaryGetResponseCodeParamsQueryTypeDs, RadarDNSSummaryGetResponseCodeParamsQueryTypeEid, RadarDNSSummaryGetResponseCodeParamsQueryTypeEui48, RadarDNSSummaryGetResponseCodeParamsQueryTypeEui64, RadarDNSSummaryGetResponseCodeParamsQueryTypeGpos, RadarDNSSummaryGetResponseCodeParamsQueryTypeGid, RadarDNSSummaryGetResponseCodeParamsQueryTypeHinfo, RadarDNSSummaryGetResponseCodeParamsQueryTypeHip, RadarDNSSummaryGetResponseCodeParamsQueryTypeHTTPS, RadarDNSSummaryGetResponseCodeParamsQueryTypeIpseckey, RadarDNSSummaryGetResponseCodeParamsQueryTypeIsdn, RadarDNSSummaryGetResponseCodeParamsQueryTypeIxfr, RadarDNSSummaryGetResponseCodeParamsQueryTypeKey, RadarDNSSummaryGetResponseCodeParamsQueryTypeKx, RadarDNSSummaryGetResponseCodeParamsQueryTypeL32, RadarDNSSummaryGetResponseCodeParamsQueryTypeL64, RadarDNSSummaryGetResponseCodeParamsQueryTypeLoc, RadarDNSSummaryGetResponseCodeParamsQueryTypeLp, RadarDNSSummaryGetResponseCodeParamsQueryTypeMaila, RadarDNSSummaryGetResponseCodeParamsQueryTypeMailb, RadarDNSSummaryGetResponseCodeParamsQueryTypeMB, RadarDNSSummaryGetResponseCodeParamsQueryTypeMd, RadarDNSSummaryGetResponseCodeParamsQueryTypeMf, RadarDNSSummaryGetResponseCodeParamsQueryTypeMg, RadarDNSSummaryGetResponseCodeParamsQueryTypeMinfo, RadarDNSSummaryGetResponseCodeParamsQueryTypeMr, RadarDNSSummaryGetResponseCodeParamsQueryTypeMx, RadarDNSSummaryGetResponseCodeParamsQueryTypeNaptr, RadarDNSSummaryGetResponseCodeParamsQueryTypeNb, RadarDNSSummaryGetResponseCodeParamsQueryTypeNbstat, RadarDNSSummaryGetResponseCodeParamsQueryTypeNid, RadarDNSSummaryGetResponseCodeParamsQueryTypeNimloc, RadarDNSSummaryGetResponseCodeParamsQueryTypeNinfo, RadarDNSSummaryGetResponseCodeParamsQueryTypeNs, RadarDNSSummaryGetResponseCodeParamsQueryTypeNsap, RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec, RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec3, RadarDNSSummaryGetResponseCodeParamsQueryTypeNsec3Param, RadarDNSSummaryGetResponseCodeParamsQueryTypeNull, RadarDNSSummaryGetResponseCodeParamsQueryTypeNxt, RadarDNSSummaryGetResponseCodeParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetResponseCodeParamsQueryTypeOpt, RadarDNSSummaryGetResponseCodeParamsQueryTypePtr, RadarDNSSummaryGetResponseCodeParamsQueryTypePx, RadarDNSSummaryGetResponseCodeParamsQueryTypeRkey, RadarDNSSummaryGetResponseCodeParamsQueryTypeRp, RadarDNSSummaryGetResponseCodeParamsQueryTypeRrsig, RadarDNSSummaryGetResponseCodeParamsQueryTypeRt, RadarDNSSummaryGetResponseCodeParamsQueryTypeSig, RadarDNSSummaryGetResponseCodeParamsQueryTypeSink, RadarDNSSummaryGetResponseCodeParamsQueryTypeSmimea, RadarDNSSummaryGetResponseCodeParamsQueryTypeSoa, RadarDNSSummaryGetResponseCodeParamsQueryTypeSpf, RadarDNSSummaryGetResponseCodeParamsQueryTypeSrv, RadarDNSSummaryGetResponseCodeParamsQueryTypeSshfp, RadarDNSSummaryGetResponseCodeParamsQueryTypeSvcb, RadarDNSSummaryGetResponseCodeParamsQueryTypeTa, RadarDNSSummaryGetResponseCodeParamsQueryTypeTalink, RadarDNSSummaryGetResponseCodeParamsQueryTypeTkey, RadarDNSSummaryGetResponseCodeParamsQueryTypeTlsa, RadarDNSSummaryGetResponseCodeParamsQueryTypeTsig, RadarDNSSummaryGetResponseCodeParamsQueryTypeTxt, RadarDNSSummaryGetResponseCodeParamsQueryTypeUinfo, RadarDNSSummaryGetResponseCodeParamsQueryTypeUid, RadarDNSSummaryGetResponseCodeParamsQueryTypeUnspec, RadarDNSSummaryGetResponseCodeParamsQueryTypeUri, RadarDNSSummaryGetResponseCodeParamsQueryTypeWks, RadarDNSSummaryGetResponseCodeParamsQueryTypeX25, RadarDNSSummaryGetResponseCodeParamsQueryTypeZonemd:
		return true
	}
	return false
}

type RadarDNSSummaryGetResponseTtlParams struct {
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
	Format param.Field[RadarDNSSummaryGetResponseTtlParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Includes empty DNS responses (NODATA).
	Nodata param.Field[bool] `query:"nodata"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarDNSSummaryGetResponseTtlParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarDNSSummaryGetResponseTtlParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarDNSSummaryGetResponseTtlParamsResponseCode] `query:"responseCode"`
	// Filters results by country code top-level domain (ccTLD).
	Tld param.Field[[]string] `query:"tld"`
}

// URLQuery serializes [RadarDNSSummaryGetResponseTtlParams]'s query parameters as
// `url.Values`.
func (r RadarDNSSummaryGetResponseTtlParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarDNSSummaryGetResponseTtlParamsFormat string

const (
	RadarDNSSummaryGetResponseTtlParamsFormatJson RadarDNSSummaryGetResponseTtlParamsFormat = "JSON"
	RadarDNSSummaryGetResponseTtlParamsFormatCsv  RadarDNSSummaryGetResponseTtlParamsFormat = "CSV"
)

func (r RadarDNSSummaryGetResponseTtlParamsFormat) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseTtlParamsFormatJson, RadarDNSSummaryGetResponseTtlParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarDNSSummaryGetResponseTtlParamsProtocol string

const (
	RadarDNSSummaryGetResponseTtlParamsProtocolUdp   RadarDNSSummaryGetResponseTtlParamsProtocol = "UDP"
	RadarDNSSummaryGetResponseTtlParamsProtocolTcp   RadarDNSSummaryGetResponseTtlParamsProtocol = "TCP"
	RadarDNSSummaryGetResponseTtlParamsProtocolHTTPS RadarDNSSummaryGetResponseTtlParamsProtocol = "HTTPS"
	RadarDNSSummaryGetResponseTtlParamsProtocolTls   RadarDNSSummaryGetResponseTtlParamsProtocol = "TLS"
)

func (r RadarDNSSummaryGetResponseTtlParamsProtocol) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseTtlParamsProtocolUdp, RadarDNSSummaryGetResponseTtlParamsProtocolTcp, RadarDNSSummaryGetResponseTtlParamsProtocolHTTPS, RadarDNSSummaryGetResponseTtlParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarDNSSummaryGetResponseTtlParamsQueryType string

const (
	RadarDNSSummaryGetResponseTtlParamsQueryTypeA          RadarDNSSummaryGetResponseTtlParamsQueryType = "A"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeAaaa       RadarDNSSummaryGetResponseTtlParamsQueryType = "AAAA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeA6         RadarDNSSummaryGetResponseTtlParamsQueryType = "A6"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeAfsdb      RadarDNSSummaryGetResponseTtlParamsQueryType = "AFSDB"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeAny        RadarDNSSummaryGetResponseTtlParamsQueryType = "ANY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeApl        RadarDNSSummaryGetResponseTtlParamsQueryType = "APL"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeAtma       RadarDNSSummaryGetResponseTtlParamsQueryType = "ATMA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeAxfr       RadarDNSSummaryGetResponseTtlParamsQueryType = "AXFR"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCaa        RadarDNSSummaryGetResponseTtlParamsQueryType = "CAA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCdnskey    RadarDNSSummaryGetResponseTtlParamsQueryType = "CDNSKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCds        RadarDNSSummaryGetResponseTtlParamsQueryType = "CDS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCert       RadarDNSSummaryGetResponseTtlParamsQueryType = "CERT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCname      RadarDNSSummaryGetResponseTtlParamsQueryType = "CNAME"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeCsync      RadarDNSSummaryGetResponseTtlParamsQueryType = "CSYNC"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDhcid      RadarDNSSummaryGetResponseTtlParamsQueryType = "DHCID"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDlv        RadarDNSSummaryGetResponseTtlParamsQueryType = "DLV"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDname      RadarDNSSummaryGetResponseTtlParamsQueryType = "DNAME"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDnskey     RadarDNSSummaryGetResponseTtlParamsQueryType = "DNSKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDoa        RadarDNSSummaryGetResponseTtlParamsQueryType = "DOA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeDs         RadarDNSSummaryGetResponseTtlParamsQueryType = "DS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeEid        RadarDNSSummaryGetResponseTtlParamsQueryType = "EID"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeEui48      RadarDNSSummaryGetResponseTtlParamsQueryType = "EUI48"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeEui64      RadarDNSSummaryGetResponseTtlParamsQueryType = "EUI64"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeGpos       RadarDNSSummaryGetResponseTtlParamsQueryType = "GPOS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeGid        RadarDNSSummaryGetResponseTtlParamsQueryType = "GID"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeHinfo      RadarDNSSummaryGetResponseTtlParamsQueryType = "HINFO"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeHip        RadarDNSSummaryGetResponseTtlParamsQueryType = "HIP"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeHTTPS      RadarDNSSummaryGetResponseTtlParamsQueryType = "HTTPS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeIpseckey   RadarDNSSummaryGetResponseTtlParamsQueryType = "IPSECKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeIsdn       RadarDNSSummaryGetResponseTtlParamsQueryType = "ISDN"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeIxfr       RadarDNSSummaryGetResponseTtlParamsQueryType = "IXFR"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeKey        RadarDNSSummaryGetResponseTtlParamsQueryType = "KEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeKx         RadarDNSSummaryGetResponseTtlParamsQueryType = "KX"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeL32        RadarDNSSummaryGetResponseTtlParamsQueryType = "L32"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeL64        RadarDNSSummaryGetResponseTtlParamsQueryType = "L64"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeLoc        RadarDNSSummaryGetResponseTtlParamsQueryType = "LOC"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeLp         RadarDNSSummaryGetResponseTtlParamsQueryType = "LP"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMaila      RadarDNSSummaryGetResponseTtlParamsQueryType = "MAILA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMailb      RadarDNSSummaryGetResponseTtlParamsQueryType = "MAILB"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMB         RadarDNSSummaryGetResponseTtlParamsQueryType = "MB"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMd         RadarDNSSummaryGetResponseTtlParamsQueryType = "MD"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMf         RadarDNSSummaryGetResponseTtlParamsQueryType = "MF"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMg         RadarDNSSummaryGetResponseTtlParamsQueryType = "MG"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMinfo      RadarDNSSummaryGetResponseTtlParamsQueryType = "MINFO"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMr         RadarDNSSummaryGetResponseTtlParamsQueryType = "MR"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeMx         RadarDNSSummaryGetResponseTtlParamsQueryType = "MX"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNaptr      RadarDNSSummaryGetResponseTtlParamsQueryType = "NAPTR"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNb         RadarDNSSummaryGetResponseTtlParamsQueryType = "NB"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNbstat     RadarDNSSummaryGetResponseTtlParamsQueryType = "NBSTAT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNid        RadarDNSSummaryGetResponseTtlParamsQueryType = "NID"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNimloc     RadarDNSSummaryGetResponseTtlParamsQueryType = "NIMLOC"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNinfo      RadarDNSSummaryGetResponseTtlParamsQueryType = "NINFO"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNs         RadarDNSSummaryGetResponseTtlParamsQueryType = "NS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNsap       RadarDNSSummaryGetResponseTtlParamsQueryType = "NSAP"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec       RadarDNSSummaryGetResponseTtlParamsQueryType = "NSEC"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec3      RadarDNSSummaryGetResponseTtlParamsQueryType = "NSEC3"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec3Param RadarDNSSummaryGetResponseTtlParamsQueryType = "NSEC3PARAM"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNull       RadarDNSSummaryGetResponseTtlParamsQueryType = "NULL"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeNxt        RadarDNSSummaryGetResponseTtlParamsQueryType = "NXT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeOpenpgpkey RadarDNSSummaryGetResponseTtlParamsQueryType = "OPENPGPKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeOpt        RadarDNSSummaryGetResponseTtlParamsQueryType = "OPT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypePtr        RadarDNSSummaryGetResponseTtlParamsQueryType = "PTR"
	RadarDNSSummaryGetResponseTtlParamsQueryTypePx         RadarDNSSummaryGetResponseTtlParamsQueryType = "PX"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeRkey       RadarDNSSummaryGetResponseTtlParamsQueryType = "RKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeRp         RadarDNSSummaryGetResponseTtlParamsQueryType = "RP"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeRrsig      RadarDNSSummaryGetResponseTtlParamsQueryType = "RRSIG"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeRt         RadarDNSSummaryGetResponseTtlParamsQueryType = "RT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSig        RadarDNSSummaryGetResponseTtlParamsQueryType = "SIG"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSink       RadarDNSSummaryGetResponseTtlParamsQueryType = "SINK"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSmimea     RadarDNSSummaryGetResponseTtlParamsQueryType = "SMIMEA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSoa        RadarDNSSummaryGetResponseTtlParamsQueryType = "SOA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSpf        RadarDNSSummaryGetResponseTtlParamsQueryType = "SPF"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSrv        RadarDNSSummaryGetResponseTtlParamsQueryType = "SRV"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSshfp      RadarDNSSummaryGetResponseTtlParamsQueryType = "SSHFP"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeSvcb       RadarDNSSummaryGetResponseTtlParamsQueryType = "SVCB"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTa         RadarDNSSummaryGetResponseTtlParamsQueryType = "TA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTalink     RadarDNSSummaryGetResponseTtlParamsQueryType = "TALINK"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTkey       RadarDNSSummaryGetResponseTtlParamsQueryType = "TKEY"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTlsa       RadarDNSSummaryGetResponseTtlParamsQueryType = "TLSA"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTsig       RadarDNSSummaryGetResponseTtlParamsQueryType = "TSIG"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeTxt        RadarDNSSummaryGetResponseTtlParamsQueryType = "TXT"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeUinfo      RadarDNSSummaryGetResponseTtlParamsQueryType = "UINFO"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeUid        RadarDNSSummaryGetResponseTtlParamsQueryType = "UID"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeUnspec     RadarDNSSummaryGetResponseTtlParamsQueryType = "UNSPEC"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeUri        RadarDNSSummaryGetResponseTtlParamsQueryType = "URI"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeWks        RadarDNSSummaryGetResponseTtlParamsQueryType = "WKS"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeX25        RadarDNSSummaryGetResponseTtlParamsQueryType = "X25"
	RadarDNSSummaryGetResponseTtlParamsQueryTypeZonemd     RadarDNSSummaryGetResponseTtlParamsQueryType = "ZONEMD"
)

func (r RadarDNSSummaryGetResponseTtlParamsQueryType) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseTtlParamsQueryTypeA, RadarDNSSummaryGetResponseTtlParamsQueryTypeAaaa, RadarDNSSummaryGetResponseTtlParamsQueryTypeA6, RadarDNSSummaryGetResponseTtlParamsQueryTypeAfsdb, RadarDNSSummaryGetResponseTtlParamsQueryTypeAny, RadarDNSSummaryGetResponseTtlParamsQueryTypeApl, RadarDNSSummaryGetResponseTtlParamsQueryTypeAtma, RadarDNSSummaryGetResponseTtlParamsQueryTypeAxfr, RadarDNSSummaryGetResponseTtlParamsQueryTypeCaa, RadarDNSSummaryGetResponseTtlParamsQueryTypeCdnskey, RadarDNSSummaryGetResponseTtlParamsQueryTypeCds, RadarDNSSummaryGetResponseTtlParamsQueryTypeCert, RadarDNSSummaryGetResponseTtlParamsQueryTypeCname, RadarDNSSummaryGetResponseTtlParamsQueryTypeCsync, RadarDNSSummaryGetResponseTtlParamsQueryTypeDhcid, RadarDNSSummaryGetResponseTtlParamsQueryTypeDlv, RadarDNSSummaryGetResponseTtlParamsQueryTypeDname, RadarDNSSummaryGetResponseTtlParamsQueryTypeDnskey, RadarDNSSummaryGetResponseTtlParamsQueryTypeDoa, RadarDNSSummaryGetResponseTtlParamsQueryTypeDs, RadarDNSSummaryGetResponseTtlParamsQueryTypeEid, RadarDNSSummaryGetResponseTtlParamsQueryTypeEui48, RadarDNSSummaryGetResponseTtlParamsQueryTypeEui64, RadarDNSSummaryGetResponseTtlParamsQueryTypeGpos, RadarDNSSummaryGetResponseTtlParamsQueryTypeGid, RadarDNSSummaryGetResponseTtlParamsQueryTypeHinfo, RadarDNSSummaryGetResponseTtlParamsQueryTypeHip, RadarDNSSummaryGetResponseTtlParamsQueryTypeHTTPS, RadarDNSSummaryGetResponseTtlParamsQueryTypeIpseckey, RadarDNSSummaryGetResponseTtlParamsQueryTypeIsdn, RadarDNSSummaryGetResponseTtlParamsQueryTypeIxfr, RadarDNSSummaryGetResponseTtlParamsQueryTypeKey, RadarDNSSummaryGetResponseTtlParamsQueryTypeKx, RadarDNSSummaryGetResponseTtlParamsQueryTypeL32, RadarDNSSummaryGetResponseTtlParamsQueryTypeL64, RadarDNSSummaryGetResponseTtlParamsQueryTypeLoc, RadarDNSSummaryGetResponseTtlParamsQueryTypeLp, RadarDNSSummaryGetResponseTtlParamsQueryTypeMaila, RadarDNSSummaryGetResponseTtlParamsQueryTypeMailb, RadarDNSSummaryGetResponseTtlParamsQueryTypeMB, RadarDNSSummaryGetResponseTtlParamsQueryTypeMd, RadarDNSSummaryGetResponseTtlParamsQueryTypeMf, RadarDNSSummaryGetResponseTtlParamsQueryTypeMg, RadarDNSSummaryGetResponseTtlParamsQueryTypeMinfo, RadarDNSSummaryGetResponseTtlParamsQueryTypeMr, RadarDNSSummaryGetResponseTtlParamsQueryTypeMx, RadarDNSSummaryGetResponseTtlParamsQueryTypeNaptr, RadarDNSSummaryGetResponseTtlParamsQueryTypeNb, RadarDNSSummaryGetResponseTtlParamsQueryTypeNbstat, RadarDNSSummaryGetResponseTtlParamsQueryTypeNid, RadarDNSSummaryGetResponseTtlParamsQueryTypeNimloc, RadarDNSSummaryGetResponseTtlParamsQueryTypeNinfo, RadarDNSSummaryGetResponseTtlParamsQueryTypeNs, RadarDNSSummaryGetResponseTtlParamsQueryTypeNsap, RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec, RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec3, RadarDNSSummaryGetResponseTtlParamsQueryTypeNsec3Param, RadarDNSSummaryGetResponseTtlParamsQueryTypeNull, RadarDNSSummaryGetResponseTtlParamsQueryTypeNxt, RadarDNSSummaryGetResponseTtlParamsQueryTypeOpenpgpkey, RadarDNSSummaryGetResponseTtlParamsQueryTypeOpt, RadarDNSSummaryGetResponseTtlParamsQueryTypePtr, RadarDNSSummaryGetResponseTtlParamsQueryTypePx, RadarDNSSummaryGetResponseTtlParamsQueryTypeRkey, RadarDNSSummaryGetResponseTtlParamsQueryTypeRp, RadarDNSSummaryGetResponseTtlParamsQueryTypeRrsig, RadarDNSSummaryGetResponseTtlParamsQueryTypeRt, RadarDNSSummaryGetResponseTtlParamsQueryTypeSig, RadarDNSSummaryGetResponseTtlParamsQueryTypeSink, RadarDNSSummaryGetResponseTtlParamsQueryTypeSmimea, RadarDNSSummaryGetResponseTtlParamsQueryTypeSoa, RadarDNSSummaryGetResponseTtlParamsQueryTypeSpf, RadarDNSSummaryGetResponseTtlParamsQueryTypeSrv, RadarDNSSummaryGetResponseTtlParamsQueryTypeSshfp, RadarDNSSummaryGetResponseTtlParamsQueryTypeSvcb, RadarDNSSummaryGetResponseTtlParamsQueryTypeTa, RadarDNSSummaryGetResponseTtlParamsQueryTypeTalink, RadarDNSSummaryGetResponseTtlParamsQueryTypeTkey, RadarDNSSummaryGetResponseTtlParamsQueryTypeTlsa, RadarDNSSummaryGetResponseTtlParamsQueryTypeTsig, RadarDNSSummaryGetResponseTtlParamsQueryTypeTxt, RadarDNSSummaryGetResponseTtlParamsQueryTypeUinfo, RadarDNSSummaryGetResponseTtlParamsQueryTypeUid, RadarDNSSummaryGetResponseTtlParamsQueryTypeUnspec, RadarDNSSummaryGetResponseTtlParamsQueryTypeUri, RadarDNSSummaryGetResponseTtlParamsQueryTypeWks, RadarDNSSummaryGetResponseTtlParamsQueryTypeX25, RadarDNSSummaryGetResponseTtlParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarDNSSummaryGetResponseTtlParamsResponseCode string

const (
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNoerror   RadarDNSSummaryGetResponseTtlParamsResponseCode = "NOERROR"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeFormerr   RadarDNSSummaryGetResponseTtlParamsResponseCode = "FORMERR"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeServfail  RadarDNSSummaryGetResponseTtlParamsResponseCode = "SERVFAIL"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNxdomain  RadarDNSSummaryGetResponseTtlParamsResponseCode = "NXDOMAIN"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNotimp    RadarDNSSummaryGetResponseTtlParamsResponseCode = "NOTIMP"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeRefused   RadarDNSSummaryGetResponseTtlParamsResponseCode = "REFUSED"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeYxdomain  RadarDNSSummaryGetResponseTtlParamsResponseCode = "YXDOMAIN"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeYxrrset   RadarDNSSummaryGetResponseTtlParamsResponseCode = "YXRRSET"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNxrrset   RadarDNSSummaryGetResponseTtlParamsResponseCode = "NXRRSET"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNotauth   RadarDNSSummaryGetResponseTtlParamsResponseCode = "NOTAUTH"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeNotzone   RadarDNSSummaryGetResponseTtlParamsResponseCode = "NOTZONE"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadsig    RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADSIG"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadkey    RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADKEY"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadtime   RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADTIME"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadmode   RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADMODE"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadname   RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADNAME"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadalg    RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADALG"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadtrunc  RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADTRUNC"
	RadarDNSSummaryGetResponseTtlParamsResponseCodeBadcookie RadarDNSSummaryGetResponseTtlParamsResponseCode = "BADCOOKIE"
)

func (r RadarDNSSummaryGetResponseTtlParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarDNSSummaryGetResponseTtlParamsResponseCodeNoerror, RadarDNSSummaryGetResponseTtlParamsResponseCodeFormerr, RadarDNSSummaryGetResponseTtlParamsResponseCodeServfail, RadarDNSSummaryGetResponseTtlParamsResponseCodeNxdomain, RadarDNSSummaryGetResponseTtlParamsResponseCodeNotimp, RadarDNSSummaryGetResponseTtlParamsResponseCodeRefused, RadarDNSSummaryGetResponseTtlParamsResponseCodeYxdomain, RadarDNSSummaryGetResponseTtlParamsResponseCodeYxrrset, RadarDNSSummaryGetResponseTtlParamsResponseCodeNxrrset, RadarDNSSummaryGetResponseTtlParamsResponseCodeNotauth, RadarDNSSummaryGetResponseTtlParamsResponseCodeNotzone, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadsig, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadkey, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadtime, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadmode, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadname, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadalg, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadtrunc, RadarDNSSummaryGetResponseTtlParamsResponseCodeBadcookie:
		return true
	}
	return false
}
