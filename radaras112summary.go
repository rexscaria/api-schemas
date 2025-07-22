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

// RadarAs112SummaryService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAs112SummaryService] method instead.
type RadarAs112SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAs112SummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112SummaryService(opts ...option.RequestOption) (r *RadarAs112SummaryService) {
	r = &RadarAs112SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of DNS queries to AS112 by DNSSEC (DNS Security
// Extensions) support.
func (r *RadarAs112SummaryService) GetDnssec(ctx context.Context, query RadarAs112SummaryGetDnssecParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries to AS112 by EDNS (Extension Mechanisms
// for DNS) support.
func (r *RadarAs112SummaryService) GetEdns(ctx context.Context, query RadarAs112SummaryGetEdnsParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries to AS112 by IP version.
func (r *RadarAs112SummaryService) GetIPVersion(ctx context.Context, query RadarAs112SummaryGetIPVersionParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries to AS112 by protocol.
func (r *RadarAs112SummaryService) GetProtocol(ctx context.Context, query RadarAs112SummaryGetProtocolParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of DNS queries to AS112 by type.
func (r *RadarAs112SummaryService) GetQueryType(ctx context.Context, query RadarAs112SummaryGetQueryTypeParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetQueryTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of AS112 DNS requests classified by response code.
func (r *RadarAs112SummaryService) GetResponseCodes(ctx context.Context, query RadarAs112SummaryGetResponseCodesParams, opts ...option.RequestOption) (res *RadarAs112SummaryGetResponseCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/summary/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112SummaryGetDnssecResponse struct {
	Result  RadarAs112SummaryGetDnssecResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarAs112SummaryGetDnssecResponseJSON   `json:"-"`
}

// radarAs112SummaryGetDnssecResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryGetDnssecResponse]
type radarAs112SummaryGetDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResult struct {
	Meta     RadarAs112SummaryGetDnssecResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryGetDnssecResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryGetDnssecResponseResultJSON     `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetDnssecResponseResult]
type radarAs112SummaryGetDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetDnssecResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetDnssecResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112SummaryGetDnssecResponseResultMeta]
type radarAs112SummaryGetDnssecResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetDnssecResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112SummaryGetDnssecResponseResultMetaDateRange]
type radarAs112SummaryGetDnssecResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecResponseResultSummary0 struct {
	NotSupported string                                               `json:"NOT_SUPPORTED,required"`
	Supported    string                                               `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryGetDnssecResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryGetDnssecResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetDnssecResponseResultSummary0]
type radarAs112SummaryGetDnssecResponseResultSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryGetDnssecResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetDnssecResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponse struct {
	Result  RadarAs112SummaryGetEdnsResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarAs112SummaryGetEdnsResponseJSON   `json:"-"`
}

// radarAs112SummaryGetEdnsResponseJSON contains the JSON metadata for the struct
// [RadarAs112SummaryGetEdnsResponse]
type radarAs112SummaryGetEdnsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResult struct {
	Meta     RadarAs112SummaryGetEdnsResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryGetEdnsResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryGetEdnsResponseResultJSON     `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetEdnsResponseResult]
type radarAs112SummaryGetEdnsResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetEdnsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	Normalization  string                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetEdnsResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112SummaryGetEdnsResponseResultMeta]
type radarAs112SummaryGetEdnsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetEdnsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112SummaryGetEdnsResponseResultMetaDateRange]
type radarAs112SummaryGetEdnsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetEdnsResponseResultSummary0 struct {
	NotSupported string                                             `json:"NOT_SUPPORTED,required"`
	Supported    string                                             `json:"SUPPORTED,required"`
	JSON         radarAs112SummaryGetEdnsResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryGetEdnsResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetEdnsResponseResultSummary0]
type radarAs112SummaryGetEdnsResponseResultSummary0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112SummaryGetEdnsResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetEdnsResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponse struct {
	Result  RadarAs112SummaryGetIPVersionResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAs112SummaryGetIPVersionResponseJSON   `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetIPVersionResponse]
type radarAs112SummaryGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResult struct {
	Meta     RadarAs112SummaryGetIPVersionResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryGetIPVersionResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryGetIPVersionResponseResultJSON     `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryGetIPVersionResponseResult]
type radarAs112SummaryGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetIPVersionResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetIPVersionResponseResultMeta]
type radarAs112SummaryGetIPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryGetIPVersionResponseResultMetaDateRange]
type radarAs112SummaryGetIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetIPVersionResponseResultSummary0 struct {
	IPv4 string                                                  `json:"IPv4,required"`
	IPv6 string                                                  `json:"IPv6,required"`
	JSON radarAs112SummaryGetIPVersionResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryGetIPVersionResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAs112SummaryGetIPVersionResponseResultSummary0]
type radarAs112SummaryGetIPVersionResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetIPVersionResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetIPVersionResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponse struct {
	Result  RadarAs112SummaryGetProtocolResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAs112SummaryGetProtocolResponseJSON   `json:"-"`
}

// radarAs112SummaryGetProtocolResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetProtocolResponse]
type radarAs112SummaryGetProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResult struct {
	Meta     RadarAs112SummaryGetProtocolResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAs112SummaryGetProtocolResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAs112SummaryGetProtocolResponseResultJSON     `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryGetProtocolResponseResult]
type radarAs112SummaryGetProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetProtocolResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetProtocolResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetProtocolResponseResultMeta]
type radarAs112SummaryGetProtocolResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetProtocolResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryGetProtocolResponseResultMetaDateRange]
type radarAs112SummaryGetProtocolResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetProtocolResponseResultSummary0 struct {
	HTTPS string                                                 `json:"HTTPS,required"`
	Tcp   string                                                 `json:"TCP,required"`
	Tls   string                                                 `json:"TLS,required"`
	Udp   string                                                 `json:"UDP,required"`
	JSON  radarAs112SummaryGetProtocolResponseResultSummary0JSON `json:"-"`
}

// radarAs112SummaryGetProtocolResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAs112SummaryGetProtocolResponseResultSummary0]
type radarAs112SummaryGetProtocolResponseResultSummary0JSON struct {
	HTTPS       apijson.Field
	Tcp         apijson.Field
	Tls         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetProtocolResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetProtocolResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponse struct {
	Result  RadarAs112SummaryGetQueryTypeResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAs112SummaryGetQueryTypeResponseJSON   `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetQueryTypeResponse]
type radarAs112SummaryGetQueryTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetQueryTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponseResult struct {
	Meta     RadarAs112SummaryGetQueryTypeResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                               `json:"summary_0,required"`
	JSON     radarAs112SummaryGetQueryTypeResponseResultJSON `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112SummaryGetQueryTypeResponseResult]
type radarAs112SummaryGetQueryTypeResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetQueryTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetQueryTypeResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	Normalization  string                                                        `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetQueryTypeResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetQueryTypeResponseResultMeta]
type radarAs112SummaryGetQueryTypeResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetQueryTypeResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetQueryTypeResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112SummaryGetQueryTypeResponseResultMetaDateRange]
type radarAs112SummaryGetQueryTypeResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetQueryTypeResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetQueryTypeResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponse struct {
	Result  RadarAs112SummaryGetResponseCodesResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAs112SummaryGetResponseCodesResponseJSON   `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseJSON contains the JSON metadata for the
// struct [RadarAs112SummaryGetResponseCodesResponse]
type radarAs112SummaryGetResponseCodesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetResponseCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponseResult struct {
	Meta     RadarAs112SummaryGetResponseCodesResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                   `json:"summary_0,required"`
	JSON     radarAs112SummaryGetResponseCodesResponseResultJSON `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112SummaryGetResponseCodesResponseResult]
type radarAs112SummaryGetResponseCodesResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetResponseCodesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponseResultMeta struct {
	DateRange      []RadarAs112SummaryGetResponseCodesResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	Normalization  string                                                            `json:"normalization,required"`
	ConfidenceInfo RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112SummaryGetResponseCodesResponseResultMetaJSON           `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112SummaryGetResponseCodesResponseResultMeta]
type radarAs112SummaryGetResponseCodesResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112SummaryGetResponseCodesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAs112SummaryGetResponseCodesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112SummaryGetResponseCodesResponseResultMetaDateRange]
type radarAs112SummaryGetResponseCodesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetResponseCodesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfo]
type radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotation]
type radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112SummaryGetResponseCodesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112SummaryGetDnssecParams struct {
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
	Format param.Field[RadarAs112SummaryGetDnssecParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112SummaryGetDnssecParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112SummaryGetDnssecParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112SummaryGetDnssecParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112SummaryGetDnssecParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryGetDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetDnssecParamsFormat string

const (
	RadarAs112SummaryGetDnssecParamsFormatJson RadarAs112SummaryGetDnssecParamsFormat = "JSON"
	RadarAs112SummaryGetDnssecParamsFormatCsv  RadarAs112SummaryGetDnssecParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetDnssecParamsFormatJson, RadarAs112SummaryGetDnssecParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112SummaryGetDnssecParamsProtocol string

const (
	RadarAs112SummaryGetDnssecParamsProtocolUdp   RadarAs112SummaryGetDnssecParamsProtocol = "UDP"
	RadarAs112SummaryGetDnssecParamsProtocolTcp   RadarAs112SummaryGetDnssecParamsProtocol = "TCP"
	RadarAs112SummaryGetDnssecParamsProtocolHTTPS RadarAs112SummaryGetDnssecParamsProtocol = "HTTPS"
	RadarAs112SummaryGetDnssecParamsProtocolTls   RadarAs112SummaryGetDnssecParamsProtocol = "TLS"
)

func (r RadarAs112SummaryGetDnssecParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetDnssecParamsProtocolUdp, RadarAs112SummaryGetDnssecParamsProtocolTcp, RadarAs112SummaryGetDnssecParamsProtocolHTTPS, RadarAs112SummaryGetDnssecParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112SummaryGetDnssecParamsQueryType string

const (
	RadarAs112SummaryGetDnssecParamsQueryTypeA          RadarAs112SummaryGetDnssecParamsQueryType = "A"
	RadarAs112SummaryGetDnssecParamsQueryTypeAaaa       RadarAs112SummaryGetDnssecParamsQueryType = "AAAA"
	RadarAs112SummaryGetDnssecParamsQueryTypeA6         RadarAs112SummaryGetDnssecParamsQueryType = "A6"
	RadarAs112SummaryGetDnssecParamsQueryTypeAfsdb      RadarAs112SummaryGetDnssecParamsQueryType = "AFSDB"
	RadarAs112SummaryGetDnssecParamsQueryTypeAny        RadarAs112SummaryGetDnssecParamsQueryType = "ANY"
	RadarAs112SummaryGetDnssecParamsQueryTypeApl        RadarAs112SummaryGetDnssecParamsQueryType = "APL"
	RadarAs112SummaryGetDnssecParamsQueryTypeAtma       RadarAs112SummaryGetDnssecParamsQueryType = "ATMA"
	RadarAs112SummaryGetDnssecParamsQueryTypeAxfr       RadarAs112SummaryGetDnssecParamsQueryType = "AXFR"
	RadarAs112SummaryGetDnssecParamsQueryTypeCaa        RadarAs112SummaryGetDnssecParamsQueryType = "CAA"
	RadarAs112SummaryGetDnssecParamsQueryTypeCdnskey    RadarAs112SummaryGetDnssecParamsQueryType = "CDNSKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeCds        RadarAs112SummaryGetDnssecParamsQueryType = "CDS"
	RadarAs112SummaryGetDnssecParamsQueryTypeCert       RadarAs112SummaryGetDnssecParamsQueryType = "CERT"
	RadarAs112SummaryGetDnssecParamsQueryTypeCname      RadarAs112SummaryGetDnssecParamsQueryType = "CNAME"
	RadarAs112SummaryGetDnssecParamsQueryTypeCsync      RadarAs112SummaryGetDnssecParamsQueryType = "CSYNC"
	RadarAs112SummaryGetDnssecParamsQueryTypeDhcid      RadarAs112SummaryGetDnssecParamsQueryType = "DHCID"
	RadarAs112SummaryGetDnssecParamsQueryTypeDlv        RadarAs112SummaryGetDnssecParamsQueryType = "DLV"
	RadarAs112SummaryGetDnssecParamsQueryTypeDname      RadarAs112SummaryGetDnssecParamsQueryType = "DNAME"
	RadarAs112SummaryGetDnssecParamsQueryTypeDnskey     RadarAs112SummaryGetDnssecParamsQueryType = "DNSKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeDoa        RadarAs112SummaryGetDnssecParamsQueryType = "DOA"
	RadarAs112SummaryGetDnssecParamsQueryTypeDs         RadarAs112SummaryGetDnssecParamsQueryType = "DS"
	RadarAs112SummaryGetDnssecParamsQueryTypeEid        RadarAs112SummaryGetDnssecParamsQueryType = "EID"
	RadarAs112SummaryGetDnssecParamsQueryTypeEui48      RadarAs112SummaryGetDnssecParamsQueryType = "EUI48"
	RadarAs112SummaryGetDnssecParamsQueryTypeEui64      RadarAs112SummaryGetDnssecParamsQueryType = "EUI64"
	RadarAs112SummaryGetDnssecParamsQueryTypeGpos       RadarAs112SummaryGetDnssecParamsQueryType = "GPOS"
	RadarAs112SummaryGetDnssecParamsQueryTypeGid        RadarAs112SummaryGetDnssecParamsQueryType = "GID"
	RadarAs112SummaryGetDnssecParamsQueryTypeHinfo      RadarAs112SummaryGetDnssecParamsQueryType = "HINFO"
	RadarAs112SummaryGetDnssecParamsQueryTypeHip        RadarAs112SummaryGetDnssecParamsQueryType = "HIP"
	RadarAs112SummaryGetDnssecParamsQueryTypeHTTPS      RadarAs112SummaryGetDnssecParamsQueryType = "HTTPS"
	RadarAs112SummaryGetDnssecParamsQueryTypeIpseckey   RadarAs112SummaryGetDnssecParamsQueryType = "IPSECKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeIsdn       RadarAs112SummaryGetDnssecParamsQueryType = "ISDN"
	RadarAs112SummaryGetDnssecParamsQueryTypeIxfr       RadarAs112SummaryGetDnssecParamsQueryType = "IXFR"
	RadarAs112SummaryGetDnssecParamsQueryTypeKey        RadarAs112SummaryGetDnssecParamsQueryType = "KEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeKx         RadarAs112SummaryGetDnssecParamsQueryType = "KX"
	RadarAs112SummaryGetDnssecParamsQueryTypeL32        RadarAs112SummaryGetDnssecParamsQueryType = "L32"
	RadarAs112SummaryGetDnssecParamsQueryTypeL64        RadarAs112SummaryGetDnssecParamsQueryType = "L64"
	RadarAs112SummaryGetDnssecParamsQueryTypeLoc        RadarAs112SummaryGetDnssecParamsQueryType = "LOC"
	RadarAs112SummaryGetDnssecParamsQueryTypeLp         RadarAs112SummaryGetDnssecParamsQueryType = "LP"
	RadarAs112SummaryGetDnssecParamsQueryTypeMaila      RadarAs112SummaryGetDnssecParamsQueryType = "MAILA"
	RadarAs112SummaryGetDnssecParamsQueryTypeMailb      RadarAs112SummaryGetDnssecParamsQueryType = "MAILB"
	RadarAs112SummaryGetDnssecParamsQueryTypeMB         RadarAs112SummaryGetDnssecParamsQueryType = "MB"
	RadarAs112SummaryGetDnssecParamsQueryTypeMd         RadarAs112SummaryGetDnssecParamsQueryType = "MD"
	RadarAs112SummaryGetDnssecParamsQueryTypeMf         RadarAs112SummaryGetDnssecParamsQueryType = "MF"
	RadarAs112SummaryGetDnssecParamsQueryTypeMg         RadarAs112SummaryGetDnssecParamsQueryType = "MG"
	RadarAs112SummaryGetDnssecParamsQueryTypeMinfo      RadarAs112SummaryGetDnssecParamsQueryType = "MINFO"
	RadarAs112SummaryGetDnssecParamsQueryTypeMr         RadarAs112SummaryGetDnssecParamsQueryType = "MR"
	RadarAs112SummaryGetDnssecParamsQueryTypeMx         RadarAs112SummaryGetDnssecParamsQueryType = "MX"
	RadarAs112SummaryGetDnssecParamsQueryTypeNaptr      RadarAs112SummaryGetDnssecParamsQueryType = "NAPTR"
	RadarAs112SummaryGetDnssecParamsQueryTypeNb         RadarAs112SummaryGetDnssecParamsQueryType = "NB"
	RadarAs112SummaryGetDnssecParamsQueryTypeNbstat     RadarAs112SummaryGetDnssecParamsQueryType = "NBSTAT"
	RadarAs112SummaryGetDnssecParamsQueryTypeNid        RadarAs112SummaryGetDnssecParamsQueryType = "NID"
	RadarAs112SummaryGetDnssecParamsQueryTypeNimloc     RadarAs112SummaryGetDnssecParamsQueryType = "NIMLOC"
	RadarAs112SummaryGetDnssecParamsQueryTypeNinfo      RadarAs112SummaryGetDnssecParamsQueryType = "NINFO"
	RadarAs112SummaryGetDnssecParamsQueryTypeNs         RadarAs112SummaryGetDnssecParamsQueryType = "NS"
	RadarAs112SummaryGetDnssecParamsQueryTypeNsap       RadarAs112SummaryGetDnssecParamsQueryType = "NSAP"
	RadarAs112SummaryGetDnssecParamsQueryTypeNsec       RadarAs112SummaryGetDnssecParamsQueryType = "NSEC"
	RadarAs112SummaryGetDnssecParamsQueryTypeNsec3      RadarAs112SummaryGetDnssecParamsQueryType = "NSEC3"
	RadarAs112SummaryGetDnssecParamsQueryTypeNsec3Param RadarAs112SummaryGetDnssecParamsQueryType = "NSEC3PARAM"
	RadarAs112SummaryGetDnssecParamsQueryTypeNull       RadarAs112SummaryGetDnssecParamsQueryType = "NULL"
	RadarAs112SummaryGetDnssecParamsQueryTypeNxt        RadarAs112SummaryGetDnssecParamsQueryType = "NXT"
	RadarAs112SummaryGetDnssecParamsQueryTypeOpenpgpkey RadarAs112SummaryGetDnssecParamsQueryType = "OPENPGPKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeOpt        RadarAs112SummaryGetDnssecParamsQueryType = "OPT"
	RadarAs112SummaryGetDnssecParamsQueryTypePtr        RadarAs112SummaryGetDnssecParamsQueryType = "PTR"
	RadarAs112SummaryGetDnssecParamsQueryTypePx         RadarAs112SummaryGetDnssecParamsQueryType = "PX"
	RadarAs112SummaryGetDnssecParamsQueryTypeRkey       RadarAs112SummaryGetDnssecParamsQueryType = "RKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeRp         RadarAs112SummaryGetDnssecParamsQueryType = "RP"
	RadarAs112SummaryGetDnssecParamsQueryTypeRrsig      RadarAs112SummaryGetDnssecParamsQueryType = "RRSIG"
	RadarAs112SummaryGetDnssecParamsQueryTypeRt         RadarAs112SummaryGetDnssecParamsQueryType = "RT"
	RadarAs112SummaryGetDnssecParamsQueryTypeSig        RadarAs112SummaryGetDnssecParamsQueryType = "SIG"
	RadarAs112SummaryGetDnssecParamsQueryTypeSink       RadarAs112SummaryGetDnssecParamsQueryType = "SINK"
	RadarAs112SummaryGetDnssecParamsQueryTypeSmimea     RadarAs112SummaryGetDnssecParamsQueryType = "SMIMEA"
	RadarAs112SummaryGetDnssecParamsQueryTypeSoa        RadarAs112SummaryGetDnssecParamsQueryType = "SOA"
	RadarAs112SummaryGetDnssecParamsQueryTypeSpf        RadarAs112SummaryGetDnssecParamsQueryType = "SPF"
	RadarAs112SummaryGetDnssecParamsQueryTypeSrv        RadarAs112SummaryGetDnssecParamsQueryType = "SRV"
	RadarAs112SummaryGetDnssecParamsQueryTypeSshfp      RadarAs112SummaryGetDnssecParamsQueryType = "SSHFP"
	RadarAs112SummaryGetDnssecParamsQueryTypeSvcb       RadarAs112SummaryGetDnssecParamsQueryType = "SVCB"
	RadarAs112SummaryGetDnssecParamsQueryTypeTa         RadarAs112SummaryGetDnssecParamsQueryType = "TA"
	RadarAs112SummaryGetDnssecParamsQueryTypeTalink     RadarAs112SummaryGetDnssecParamsQueryType = "TALINK"
	RadarAs112SummaryGetDnssecParamsQueryTypeTkey       RadarAs112SummaryGetDnssecParamsQueryType = "TKEY"
	RadarAs112SummaryGetDnssecParamsQueryTypeTlsa       RadarAs112SummaryGetDnssecParamsQueryType = "TLSA"
	RadarAs112SummaryGetDnssecParamsQueryTypeTsig       RadarAs112SummaryGetDnssecParamsQueryType = "TSIG"
	RadarAs112SummaryGetDnssecParamsQueryTypeTxt        RadarAs112SummaryGetDnssecParamsQueryType = "TXT"
	RadarAs112SummaryGetDnssecParamsQueryTypeUinfo      RadarAs112SummaryGetDnssecParamsQueryType = "UINFO"
	RadarAs112SummaryGetDnssecParamsQueryTypeUid        RadarAs112SummaryGetDnssecParamsQueryType = "UID"
	RadarAs112SummaryGetDnssecParamsQueryTypeUnspec     RadarAs112SummaryGetDnssecParamsQueryType = "UNSPEC"
	RadarAs112SummaryGetDnssecParamsQueryTypeUri        RadarAs112SummaryGetDnssecParamsQueryType = "URI"
	RadarAs112SummaryGetDnssecParamsQueryTypeWks        RadarAs112SummaryGetDnssecParamsQueryType = "WKS"
	RadarAs112SummaryGetDnssecParamsQueryTypeX25        RadarAs112SummaryGetDnssecParamsQueryType = "X25"
	RadarAs112SummaryGetDnssecParamsQueryTypeZonemd     RadarAs112SummaryGetDnssecParamsQueryType = "ZONEMD"
)

func (r RadarAs112SummaryGetDnssecParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetDnssecParamsQueryTypeA, RadarAs112SummaryGetDnssecParamsQueryTypeAaaa, RadarAs112SummaryGetDnssecParamsQueryTypeA6, RadarAs112SummaryGetDnssecParamsQueryTypeAfsdb, RadarAs112SummaryGetDnssecParamsQueryTypeAny, RadarAs112SummaryGetDnssecParamsQueryTypeApl, RadarAs112SummaryGetDnssecParamsQueryTypeAtma, RadarAs112SummaryGetDnssecParamsQueryTypeAxfr, RadarAs112SummaryGetDnssecParamsQueryTypeCaa, RadarAs112SummaryGetDnssecParamsQueryTypeCdnskey, RadarAs112SummaryGetDnssecParamsQueryTypeCds, RadarAs112SummaryGetDnssecParamsQueryTypeCert, RadarAs112SummaryGetDnssecParamsQueryTypeCname, RadarAs112SummaryGetDnssecParamsQueryTypeCsync, RadarAs112SummaryGetDnssecParamsQueryTypeDhcid, RadarAs112SummaryGetDnssecParamsQueryTypeDlv, RadarAs112SummaryGetDnssecParamsQueryTypeDname, RadarAs112SummaryGetDnssecParamsQueryTypeDnskey, RadarAs112SummaryGetDnssecParamsQueryTypeDoa, RadarAs112SummaryGetDnssecParamsQueryTypeDs, RadarAs112SummaryGetDnssecParamsQueryTypeEid, RadarAs112SummaryGetDnssecParamsQueryTypeEui48, RadarAs112SummaryGetDnssecParamsQueryTypeEui64, RadarAs112SummaryGetDnssecParamsQueryTypeGpos, RadarAs112SummaryGetDnssecParamsQueryTypeGid, RadarAs112SummaryGetDnssecParamsQueryTypeHinfo, RadarAs112SummaryGetDnssecParamsQueryTypeHip, RadarAs112SummaryGetDnssecParamsQueryTypeHTTPS, RadarAs112SummaryGetDnssecParamsQueryTypeIpseckey, RadarAs112SummaryGetDnssecParamsQueryTypeIsdn, RadarAs112SummaryGetDnssecParamsQueryTypeIxfr, RadarAs112SummaryGetDnssecParamsQueryTypeKey, RadarAs112SummaryGetDnssecParamsQueryTypeKx, RadarAs112SummaryGetDnssecParamsQueryTypeL32, RadarAs112SummaryGetDnssecParamsQueryTypeL64, RadarAs112SummaryGetDnssecParamsQueryTypeLoc, RadarAs112SummaryGetDnssecParamsQueryTypeLp, RadarAs112SummaryGetDnssecParamsQueryTypeMaila, RadarAs112SummaryGetDnssecParamsQueryTypeMailb, RadarAs112SummaryGetDnssecParamsQueryTypeMB, RadarAs112SummaryGetDnssecParamsQueryTypeMd, RadarAs112SummaryGetDnssecParamsQueryTypeMf, RadarAs112SummaryGetDnssecParamsQueryTypeMg, RadarAs112SummaryGetDnssecParamsQueryTypeMinfo, RadarAs112SummaryGetDnssecParamsQueryTypeMr, RadarAs112SummaryGetDnssecParamsQueryTypeMx, RadarAs112SummaryGetDnssecParamsQueryTypeNaptr, RadarAs112SummaryGetDnssecParamsQueryTypeNb, RadarAs112SummaryGetDnssecParamsQueryTypeNbstat, RadarAs112SummaryGetDnssecParamsQueryTypeNid, RadarAs112SummaryGetDnssecParamsQueryTypeNimloc, RadarAs112SummaryGetDnssecParamsQueryTypeNinfo, RadarAs112SummaryGetDnssecParamsQueryTypeNs, RadarAs112SummaryGetDnssecParamsQueryTypeNsap, RadarAs112SummaryGetDnssecParamsQueryTypeNsec, RadarAs112SummaryGetDnssecParamsQueryTypeNsec3, RadarAs112SummaryGetDnssecParamsQueryTypeNsec3Param, RadarAs112SummaryGetDnssecParamsQueryTypeNull, RadarAs112SummaryGetDnssecParamsQueryTypeNxt, RadarAs112SummaryGetDnssecParamsQueryTypeOpenpgpkey, RadarAs112SummaryGetDnssecParamsQueryTypeOpt, RadarAs112SummaryGetDnssecParamsQueryTypePtr, RadarAs112SummaryGetDnssecParamsQueryTypePx, RadarAs112SummaryGetDnssecParamsQueryTypeRkey, RadarAs112SummaryGetDnssecParamsQueryTypeRp, RadarAs112SummaryGetDnssecParamsQueryTypeRrsig, RadarAs112SummaryGetDnssecParamsQueryTypeRt, RadarAs112SummaryGetDnssecParamsQueryTypeSig, RadarAs112SummaryGetDnssecParamsQueryTypeSink, RadarAs112SummaryGetDnssecParamsQueryTypeSmimea, RadarAs112SummaryGetDnssecParamsQueryTypeSoa, RadarAs112SummaryGetDnssecParamsQueryTypeSpf, RadarAs112SummaryGetDnssecParamsQueryTypeSrv, RadarAs112SummaryGetDnssecParamsQueryTypeSshfp, RadarAs112SummaryGetDnssecParamsQueryTypeSvcb, RadarAs112SummaryGetDnssecParamsQueryTypeTa, RadarAs112SummaryGetDnssecParamsQueryTypeTalink, RadarAs112SummaryGetDnssecParamsQueryTypeTkey, RadarAs112SummaryGetDnssecParamsQueryTypeTlsa, RadarAs112SummaryGetDnssecParamsQueryTypeTsig, RadarAs112SummaryGetDnssecParamsQueryTypeTxt, RadarAs112SummaryGetDnssecParamsQueryTypeUinfo, RadarAs112SummaryGetDnssecParamsQueryTypeUid, RadarAs112SummaryGetDnssecParamsQueryTypeUnspec, RadarAs112SummaryGetDnssecParamsQueryTypeUri, RadarAs112SummaryGetDnssecParamsQueryTypeWks, RadarAs112SummaryGetDnssecParamsQueryTypeX25, RadarAs112SummaryGetDnssecParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112SummaryGetDnssecParamsResponseCode string

const (
	RadarAs112SummaryGetDnssecParamsResponseCodeNoerror   RadarAs112SummaryGetDnssecParamsResponseCode = "NOERROR"
	RadarAs112SummaryGetDnssecParamsResponseCodeFormerr   RadarAs112SummaryGetDnssecParamsResponseCode = "FORMERR"
	RadarAs112SummaryGetDnssecParamsResponseCodeServfail  RadarAs112SummaryGetDnssecParamsResponseCode = "SERVFAIL"
	RadarAs112SummaryGetDnssecParamsResponseCodeNxdomain  RadarAs112SummaryGetDnssecParamsResponseCode = "NXDOMAIN"
	RadarAs112SummaryGetDnssecParamsResponseCodeNotimp    RadarAs112SummaryGetDnssecParamsResponseCode = "NOTIMP"
	RadarAs112SummaryGetDnssecParamsResponseCodeRefused   RadarAs112SummaryGetDnssecParamsResponseCode = "REFUSED"
	RadarAs112SummaryGetDnssecParamsResponseCodeYxdomain  RadarAs112SummaryGetDnssecParamsResponseCode = "YXDOMAIN"
	RadarAs112SummaryGetDnssecParamsResponseCodeYxrrset   RadarAs112SummaryGetDnssecParamsResponseCode = "YXRRSET"
	RadarAs112SummaryGetDnssecParamsResponseCodeNxrrset   RadarAs112SummaryGetDnssecParamsResponseCode = "NXRRSET"
	RadarAs112SummaryGetDnssecParamsResponseCodeNotauth   RadarAs112SummaryGetDnssecParamsResponseCode = "NOTAUTH"
	RadarAs112SummaryGetDnssecParamsResponseCodeNotzone   RadarAs112SummaryGetDnssecParamsResponseCode = "NOTZONE"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadsig    RadarAs112SummaryGetDnssecParamsResponseCode = "BADSIG"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadkey    RadarAs112SummaryGetDnssecParamsResponseCode = "BADKEY"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadtime   RadarAs112SummaryGetDnssecParamsResponseCode = "BADTIME"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadmode   RadarAs112SummaryGetDnssecParamsResponseCode = "BADMODE"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadname   RadarAs112SummaryGetDnssecParamsResponseCode = "BADNAME"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadalg    RadarAs112SummaryGetDnssecParamsResponseCode = "BADALG"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadtrunc  RadarAs112SummaryGetDnssecParamsResponseCode = "BADTRUNC"
	RadarAs112SummaryGetDnssecParamsResponseCodeBadcookie RadarAs112SummaryGetDnssecParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112SummaryGetDnssecParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetDnssecParamsResponseCodeNoerror, RadarAs112SummaryGetDnssecParamsResponseCodeFormerr, RadarAs112SummaryGetDnssecParamsResponseCodeServfail, RadarAs112SummaryGetDnssecParamsResponseCodeNxdomain, RadarAs112SummaryGetDnssecParamsResponseCodeNotimp, RadarAs112SummaryGetDnssecParamsResponseCodeRefused, RadarAs112SummaryGetDnssecParamsResponseCodeYxdomain, RadarAs112SummaryGetDnssecParamsResponseCodeYxrrset, RadarAs112SummaryGetDnssecParamsResponseCodeNxrrset, RadarAs112SummaryGetDnssecParamsResponseCodeNotauth, RadarAs112SummaryGetDnssecParamsResponseCodeNotzone, RadarAs112SummaryGetDnssecParamsResponseCodeBadsig, RadarAs112SummaryGetDnssecParamsResponseCodeBadkey, RadarAs112SummaryGetDnssecParamsResponseCodeBadtime, RadarAs112SummaryGetDnssecParamsResponseCodeBadmode, RadarAs112SummaryGetDnssecParamsResponseCodeBadname, RadarAs112SummaryGetDnssecParamsResponseCodeBadalg, RadarAs112SummaryGetDnssecParamsResponseCodeBadtrunc, RadarAs112SummaryGetDnssecParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112SummaryGetEdnsParams struct {
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
	Format param.Field[RadarAs112SummaryGetEdnsParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112SummaryGetEdnsParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112SummaryGetEdnsParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112SummaryGetEdnsParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112SummaryGetEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryGetEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetEdnsParamsFormat string

const (
	RadarAs112SummaryGetEdnsParamsFormatJson RadarAs112SummaryGetEdnsParamsFormat = "JSON"
	RadarAs112SummaryGetEdnsParamsFormatCsv  RadarAs112SummaryGetEdnsParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetEdnsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetEdnsParamsFormatJson, RadarAs112SummaryGetEdnsParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112SummaryGetEdnsParamsProtocol string

const (
	RadarAs112SummaryGetEdnsParamsProtocolUdp   RadarAs112SummaryGetEdnsParamsProtocol = "UDP"
	RadarAs112SummaryGetEdnsParamsProtocolTcp   RadarAs112SummaryGetEdnsParamsProtocol = "TCP"
	RadarAs112SummaryGetEdnsParamsProtocolHTTPS RadarAs112SummaryGetEdnsParamsProtocol = "HTTPS"
	RadarAs112SummaryGetEdnsParamsProtocolTls   RadarAs112SummaryGetEdnsParamsProtocol = "TLS"
)

func (r RadarAs112SummaryGetEdnsParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetEdnsParamsProtocolUdp, RadarAs112SummaryGetEdnsParamsProtocolTcp, RadarAs112SummaryGetEdnsParamsProtocolHTTPS, RadarAs112SummaryGetEdnsParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112SummaryGetEdnsParamsQueryType string

const (
	RadarAs112SummaryGetEdnsParamsQueryTypeA          RadarAs112SummaryGetEdnsParamsQueryType = "A"
	RadarAs112SummaryGetEdnsParamsQueryTypeAaaa       RadarAs112SummaryGetEdnsParamsQueryType = "AAAA"
	RadarAs112SummaryGetEdnsParamsQueryTypeA6         RadarAs112SummaryGetEdnsParamsQueryType = "A6"
	RadarAs112SummaryGetEdnsParamsQueryTypeAfsdb      RadarAs112SummaryGetEdnsParamsQueryType = "AFSDB"
	RadarAs112SummaryGetEdnsParamsQueryTypeAny        RadarAs112SummaryGetEdnsParamsQueryType = "ANY"
	RadarAs112SummaryGetEdnsParamsQueryTypeApl        RadarAs112SummaryGetEdnsParamsQueryType = "APL"
	RadarAs112SummaryGetEdnsParamsQueryTypeAtma       RadarAs112SummaryGetEdnsParamsQueryType = "ATMA"
	RadarAs112SummaryGetEdnsParamsQueryTypeAxfr       RadarAs112SummaryGetEdnsParamsQueryType = "AXFR"
	RadarAs112SummaryGetEdnsParamsQueryTypeCaa        RadarAs112SummaryGetEdnsParamsQueryType = "CAA"
	RadarAs112SummaryGetEdnsParamsQueryTypeCdnskey    RadarAs112SummaryGetEdnsParamsQueryType = "CDNSKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeCds        RadarAs112SummaryGetEdnsParamsQueryType = "CDS"
	RadarAs112SummaryGetEdnsParamsQueryTypeCert       RadarAs112SummaryGetEdnsParamsQueryType = "CERT"
	RadarAs112SummaryGetEdnsParamsQueryTypeCname      RadarAs112SummaryGetEdnsParamsQueryType = "CNAME"
	RadarAs112SummaryGetEdnsParamsQueryTypeCsync      RadarAs112SummaryGetEdnsParamsQueryType = "CSYNC"
	RadarAs112SummaryGetEdnsParamsQueryTypeDhcid      RadarAs112SummaryGetEdnsParamsQueryType = "DHCID"
	RadarAs112SummaryGetEdnsParamsQueryTypeDlv        RadarAs112SummaryGetEdnsParamsQueryType = "DLV"
	RadarAs112SummaryGetEdnsParamsQueryTypeDname      RadarAs112SummaryGetEdnsParamsQueryType = "DNAME"
	RadarAs112SummaryGetEdnsParamsQueryTypeDnskey     RadarAs112SummaryGetEdnsParamsQueryType = "DNSKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeDoa        RadarAs112SummaryGetEdnsParamsQueryType = "DOA"
	RadarAs112SummaryGetEdnsParamsQueryTypeDs         RadarAs112SummaryGetEdnsParamsQueryType = "DS"
	RadarAs112SummaryGetEdnsParamsQueryTypeEid        RadarAs112SummaryGetEdnsParamsQueryType = "EID"
	RadarAs112SummaryGetEdnsParamsQueryTypeEui48      RadarAs112SummaryGetEdnsParamsQueryType = "EUI48"
	RadarAs112SummaryGetEdnsParamsQueryTypeEui64      RadarAs112SummaryGetEdnsParamsQueryType = "EUI64"
	RadarAs112SummaryGetEdnsParamsQueryTypeGpos       RadarAs112SummaryGetEdnsParamsQueryType = "GPOS"
	RadarAs112SummaryGetEdnsParamsQueryTypeGid        RadarAs112SummaryGetEdnsParamsQueryType = "GID"
	RadarAs112SummaryGetEdnsParamsQueryTypeHinfo      RadarAs112SummaryGetEdnsParamsQueryType = "HINFO"
	RadarAs112SummaryGetEdnsParamsQueryTypeHip        RadarAs112SummaryGetEdnsParamsQueryType = "HIP"
	RadarAs112SummaryGetEdnsParamsQueryTypeHTTPS      RadarAs112SummaryGetEdnsParamsQueryType = "HTTPS"
	RadarAs112SummaryGetEdnsParamsQueryTypeIpseckey   RadarAs112SummaryGetEdnsParamsQueryType = "IPSECKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeIsdn       RadarAs112SummaryGetEdnsParamsQueryType = "ISDN"
	RadarAs112SummaryGetEdnsParamsQueryTypeIxfr       RadarAs112SummaryGetEdnsParamsQueryType = "IXFR"
	RadarAs112SummaryGetEdnsParamsQueryTypeKey        RadarAs112SummaryGetEdnsParamsQueryType = "KEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeKx         RadarAs112SummaryGetEdnsParamsQueryType = "KX"
	RadarAs112SummaryGetEdnsParamsQueryTypeL32        RadarAs112SummaryGetEdnsParamsQueryType = "L32"
	RadarAs112SummaryGetEdnsParamsQueryTypeL64        RadarAs112SummaryGetEdnsParamsQueryType = "L64"
	RadarAs112SummaryGetEdnsParamsQueryTypeLoc        RadarAs112SummaryGetEdnsParamsQueryType = "LOC"
	RadarAs112SummaryGetEdnsParamsQueryTypeLp         RadarAs112SummaryGetEdnsParamsQueryType = "LP"
	RadarAs112SummaryGetEdnsParamsQueryTypeMaila      RadarAs112SummaryGetEdnsParamsQueryType = "MAILA"
	RadarAs112SummaryGetEdnsParamsQueryTypeMailb      RadarAs112SummaryGetEdnsParamsQueryType = "MAILB"
	RadarAs112SummaryGetEdnsParamsQueryTypeMB         RadarAs112SummaryGetEdnsParamsQueryType = "MB"
	RadarAs112SummaryGetEdnsParamsQueryTypeMd         RadarAs112SummaryGetEdnsParamsQueryType = "MD"
	RadarAs112SummaryGetEdnsParamsQueryTypeMf         RadarAs112SummaryGetEdnsParamsQueryType = "MF"
	RadarAs112SummaryGetEdnsParamsQueryTypeMg         RadarAs112SummaryGetEdnsParamsQueryType = "MG"
	RadarAs112SummaryGetEdnsParamsQueryTypeMinfo      RadarAs112SummaryGetEdnsParamsQueryType = "MINFO"
	RadarAs112SummaryGetEdnsParamsQueryTypeMr         RadarAs112SummaryGetEdnsParamsQueryType = "MR"
	RadarAs112SummaryGetEdnsParamsQueryTypeMx         RadarAs112SummaryGetEdnsParamsQueryType = "MX"
	RadarAs112SummaryGetEdnsParamsQueryTypeNaptr      RadarAs112SummaryGetEdnsParamsQueryType = "NAPTR"
	RadarAs112SummaryGetEdnsParamsQueryTypeNb         RadarAs112SummaryGetEdnsParamsQueryType = "NB"
	RadarAs112SummaryGetEdnsParamsQueryTypeNbstat     RadarAs112SummaryGetEdnsParamsQueryType = "NBSTAT"
	RadarAs112SummaryGetEdnsParamsQueryTypeNid        RadarAs112SummaryGetEdnsParamsQueryType = "NID"
	RadarAs112SummaryGetEdnsParamsQueryTypeNimloc     RadarAs112SummaryGetEdnsParamsQueryType = "NIMLOC"
	RadarAs112SummaryGetEdnsParamsQueryTypeNinfo      RadarAs112SummaryGetEdnsParamsQueryType = "NINFO"
	RadarAs112SummaryGetEdnsParamsQueryTypeNs         RadarAs112SummaryGetEdnsParamsQueryType = "NS"
	RadarAs112SummaryGetEdnsParamsQueryTypeNsap       RadarAs112SummaryGetEdnsParamsQueryType = "NSAP"
	RadarAs112SummaryGetEdnsParamsQueryTypeNsec       RadarAs112SummaryGetEdnsParamsQueryType = "NSEC"
	RadarAs112SummaryGetEdnsParamsQueryTypeNsec3      RadarAs112SummaryGetEdnsParamsQueryType = "NSEC3"
	RadarAs112SummaryGetEdnsParamsQueryTypeNsec3Param RadarAs112SummaryGetEdnsParamsQueryType = "NSEC3PARAM"
	RadarAs112SummaryGetEdnsParamsQueryTypeNull       RadarAs112SummaryGetEdnsParamsQueryType = "NULL"
	RadarAs112SummaryGetEdnsParamsQueryTypeNxt        RadarAs112SummaryGetEdnsParamsQueryType = "NXT"
	RadarAs112SummaryGetEdnsParamsQueryTypeOpenpgpkey RadarAs112SummaryGetEdnsParamsQueryType = "OPENPGPKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeOpt        RadarAs112SummaryGetEdnsParamsQueryType = "OPT"
	RadarAs112SummaryGetEdnsParamsQueryTypePtr        RadarAs112SummaryGetEdnsParamsQueryType = "PTR"
	RadarAs112SummaryGetEdnsParamsQueryTypePx         RadarAs112SummaryGetEdnsParamsQueryType = "PX"
	RadarAs112SummaryGetEdnsParamsQueryTypeRkey       RadarAs112SummaryGetEdnsParamsQueryType = "RKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeRp         RadarAs112SummaryGetEdnsParamsQueryType = "RP"
	RadarAs112SummaryGetEdnsParamsQueryTypeRrsig      RadarAs112SummaryGetEdnsParamsQueryType = "RRSIG"
	RadarAs112SummaryGetEdnsParamsQueryTypeRt         RadarAs112SummaryGetEdnsParamsQueryType = "RT"
	RadarAs112SummaryGetEdnsParamsQueryTypeSig        RadarAs112SummaryGetEdnsParamsQueryType = "SIG"
	RadarAs112SummaryGetEdnsParamsQueryTypeSink       RadarAs112SummaryGetEdnsParamsQueryType = "SINK"
	RadarAs112SummaryGetEdnsParamsQueryTypeSmimea     RadarAs112SummaryGetEdnsParamsQueryType = "SMIMEA"
	RadarAs112SummaryGetEdnsParamsQueryTypeSoa        RadarAs112SummaryGetEdnsParamsQueryType = "SOA"
	RadarAs112SummaryGetEdnsParamsQueryTypeSpf        RadarAs112SummaryGetEdnsParamsQueryType = "SPF"
	RadarAs112SummaryGetEdnsParamsQueryTypeSrv        RadarAs112SummaryGetEdnsParamsQueryType = "SRV"
	RadarAs112SummaryGetEdnsParamsQueryTypeSshfp      RadarAs112SummaryGetEdnsParamsQueryType = "SSHFP"
	RadarAs112SummaryGetEdnsParamsQueryTypeSvcb       RadarAs112SummaryGetEdnsParamsQueryType = "SVCB"
	RadarAs112SummaryGetEdnsParamsQueryTypeTa         RadarAs112SummaryGetEdnsParamsQueryType = "TA"
	RadarAs112SummaryGetEdnsParamsQueryTypeTalink     RadarAs112SummaryGetEdnsParamsQueryType = "TALINK"
	RadarAs112SummaryGetEdnsParamsQueryTypeTkey       RadarAs112SummaryGetEdnsParamsQueryType = "TKEY"
	RadarAs112SummaryGetEdnsParamsQueryTypeTlsa       RadarAs112SummaryGetEdnsParamsQueryType = "TLSA"
	RadarAs112SummaryGetEdnsParamsQueryTypeTsig       RadarAs112SummaryGetEdnsParamsQueryType = "TSIG"
	RadarAs112SummaryGetEdnsParamsQueryTypeTxt        RadarAs112SummaryGetEdnsParamsQueryType = "TXT"
	RadarAs112SummaryGetEdnsParamsQueryTypeUinfo      RadarAs112SummaryGetEdnsParamsQueryType = "UINFO"
	RadarAs112SummaryGetEdnsParamsQueryTypeUid        RadarAs112SummaryGetEdnsParamsQueryType = "UID"
	RadarAs112SummaryGetEdnsParamsQueryTypeUnspec     RadarAs112SummaryGetEdnsParamsQueryType = "UNSPEC"
	RadarAs112SummaryGetEdnsParamsQueryTypeUri        RadarAs112SummaryGetEdnsParamsQueryType = "URI"
	RadarAs112SummaryGetEdnsParamsQueryTypeWks        RadarAs112SummaryGetEdnsParamsQueryType = "WKS"
	RadarAs112SummaryGetEdnsParamsQueryTypeX25        RadarAs112SummaryGetEdnsParamsQueryType = "X25"
	RadarAs112SummaryGetEdnsParamsQueryTypeZonemd     RadarAs112SummaryGetEdnsParamsQueryType = "ZONEMD"
)

func (r RadarAs112SummaryGetEdnsParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetEdnsParamsQueryTypeA, RadarAs112SummaryGetEdnsParamsQueryTypeAaaa, RadarAs112SummaryGetEdnsParamsQueryTypeA6, RadarAs112SummaryGetEdnsParamsQueryTypeAfsdb, RadarAs112SummaryGetEdnsParamsQueryTypeAny, RadarAs112SummaryGetEdnsParamsQueryTypeApl, RadarAs112SummaryGetEdnsParamsQueryTypeAtma, RadarAs112SummaryGetEdnsParamsQueryTypeAxfr, RadarAs112SummaryGetEdnsParamsQueryTypeCaa, RadarAs112SummaryGetEdnsParamsQueryTypeCdnskey, RadarAs112SummaryGetEdnsParamsQueryTypeCds, RadarAs112SummaryGetEdnsParamsQueryTypeCert, RadarAs112SummaryGetEdnsParamsQueryTypeCname, RadarAs112SummaryGetEdnsParamsQueryTypeCsync, RadarAs112SummaryGetEdnsParamsQueryTypeDhcid, RadarAs112SummaryGetEdnsParamsQueryTypeDlv, RadarAs112SummaryGetEdnsParamsQueryTypeDname, RadarAs112SummaryGetEdnsParamsQueryTypeDnskey, RadarAs112SummaryGetEdnsParamsQueryTypeDoa, RadarAs112SummaryGetEdnsParamsQueryTypeDs, RadarAs112SummaryGetEdnsParamsQueryTypeEid, RadarAs112SummaryGetEdnsParamsQueryTypeEui48, RadarAs112SummaryGetEdnsParamsQueryTypeEui64, RadarAs112SummaryGetEdnsParamsQueryTypeGpos, RadarAs112SummaryGetEdnsParamsQueryTypeGid, RadarAs112SummaryGetEdnsParamsQueryTypeHinfo, RadarAs112SummaryGetEdnsParamsQueryTypeHip, RadarAs112SummaryGetEdnsParamsQueryTypeHTTPS, RadarAs112SummaryGetEdnsParamsQueryTypeIpseckey, RadarAs112SummaryGetEdnsParamsQueryTypeIsdn, RadarAs112SummaryGetEdnsParamsQueryTypeIxfr, RadarAs112SummaryGetEdnsParamsQueryTypeKey, RadarAs112SummaryGetEdnsParamsQueryTypeKx, RadarAs112SummaryGetEdnsParamsQueryTypeL32, RadarAs112SummaryGetEdnsParamsQueryTypeL64, RadarAs112SummaryGetEdnsParamsQueryTypeLoc, RadarAs112SummaryGetEdnsParamsQueryTypeLp, RadarAs112SummaryGetEdnsParamsQueryTypeMaila, RadarAs112SummaryGetEdnsParamsQueryTypeMailb, RadarAs112SummaryGetEdnsParamsQueryTypeMB, RadarAs112SummaryGetEdnsParamsQueryTypeMd, RadarAs112SummaryGetEdnsParamsQueryTypeMf, RadarAs112SummaryGetEdnsParamsQueryTypeMg, RadarAs112SummaryGetEdnsParamsQueryTypeMinfo, RadarAs112SummaryGetEdnsParamsQueryTypeMr, RadarAs112SummaryGetEdnsParamsQueryTypeMx, RadarAs112SummaryGetEdnsParamsQueryTypeNaptr, RadarAs112SummaryGetEdnsParamsQueryTypeNb, RadarAs112SummaryGetEdnsParamsQueryTypeNbstat, RadarAs112SummaryGetEdnsParamsQueryTypeNid, RadarAs112SummaryGetEdnsParamsQueryTypeNimloc, RadarAs112SummaryGetEdnsParamsQueryTypeNinfo, RadarAs112SummaryGetEdnsParamsQueryTypeNs, RadarAs112SummaryGetEdnsParamsQueryTypeNsap, RadarAs112SummaryGetEdnsParamsQueryTypeNsec, RadarAs112SummaryGetEdnsParamsQueryTypeNsec3, RadarAs112SummaryGetEdnsParamsQueryTypeNsec3Param, RadarAs112SummaryGetEdnsParamsQueryTypeNull, RadarAs112SummaryGetEdnsParamsQueryTypeNxt, RadarAs112SummaryGetEdnsParamsQueryTypeOpenpgpkey, RadarAs112SummaryGetEdnsParamsQueryTypeOpt, RadarAs112SummaryGetEdnsParamsQueryTypePtr, RadarAs112SummaryGetEdnsParamsQueryTypePx, RadarAs112SummaryGetEdnsParamsQueryTypeRkey, RadarAs112SummaryGetEdnsParamsQueryTypeRp, RadarAs112SummaryGetEdnsParamsQueryTypeRrsig, RadarAs112SummaryGetEdnsParamsQueryTypeRt, RadarAs112SummaryGetEdnsParamsQueryTypeSig, RadarAs112SummaryGetEdnsParamsQueryTypeSink, RadarAs112SummaryGetEdnsParamsQueryTypeSmimea, RadarAs112SummaryGetEdnsParamsQueryTypeSoa, RadarAs112SummaryGetEdnsParamsQueryTypeSpf, RadarAs112SummaryGetEdnsParamsQueryTypeSrv, RadarAs112SummaryGetEdnsParamsQueryTypeSshfp, RadarAs112SummaryGetEdnsParamsQueryTypeSvcb, RadarAs112SummaryGetEdnsParamsQueryTypeTa, RadarAs112SummaryGetEdnsParamsQueryTypeTalink, RadarAs112SummaryGetEdnsParamsQueryTypeTkey, RadarAs112SummaryGetEdnsParamsQueryTypeTlsa, RadarAs112SummaryGetEdnsParamsQueryTypeTsig, RadarAs112SummaryGetEdnsParamsQueryTypeTxt, RadarAs112SummaryGetEdnsParamsQueryTypeUinfo, RadarAs112SummaryGetEdnsParamsQueryTypeUid, RadarAs112SummaryGetEdnsParamsQueryTypeUnspec, RadarAs112SummaryGetEdnsParamsQueryTypeUri, RadarAs112SummaryGetEdnsParamsQueryTypeWks, RadarAs112SummaryGetEdnsParamsQueryTypeX25, RadarAs112SummaryGetEdnsParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112SummaryGetEdnsParamsResponseCode string

const (
	RadarAs112SummaryGetEdnsParamsResponseCodeNoerror   RadarAs112SummaryGetEdnsParamsResponseCode = "NOERROR"
	RadarAs112SummaryGetEdnsParamsResponseCodeFormerr   RadarAs112SummaryGetEdnsParamsResponseCode = "FORMERR"
	RadarAs112SummaryGetEdnsParamsResponseCodeServfail  RadarAs112SummaryGetEdnsParamsResponseCode = "SERVFAIL"
	RadarAs112SummaryGetEdnsParamsResponseCodeNxdomain  RadarAs112SummaryGetEdnsParamsResponseCode = "NXDOMAIN"
	RadarAs112SummaryGetEdnsParamsResponseCodeNotimp    RadarAs112SummaryGetEdnsParamsResponseCode = "NOTIMP"
	RadarAs112SummaryGetEdnsParamsResponseCodeRefused   RadarAs112SummaryGetEdnsParamsResponseCode = "REFUSED"
	RadarAs112SummaryGetEdnsParamsResponseCodeYxdomain  RadarAs112SummaryGetEdnsParamsResponseCode = "YXDOMAIN"
	RadarAs112SummaryGetEdnsParamsResponseCodeYxrrset   RadarAs112SummaryGetEdnsParamsResponseCode = "YXRRSET"
	RadarAs112SummaryGetEdnsParamsResponseCodeNxrrset   RadarAs112SummaryGetEdnsParamsResponseCode = "NXRRSET"
	RadarAs112SummaryGetEdnsParamsResponseCodeNotauth   RadarAs112SummaryGetEdnsParamsResponseCode = "NOTAUTH"
	RadarAs112SummaryGetEdnsParamsResponseCodeNotzone   RadarAs112SummaryGetEdnsParamsResponseCode = "NOTZONE"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadsig    RadarAs112SummaryGetEdnsParamsResponseCode = "BADSIG"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadkey    RadarAs112SummaryGetEdnsParamsResponseCode = "BADKEY"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadtime   RadarAs112SummaryGetEdnsParamsResponseCode = "BADTIME"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadmode   RadarAs112SummaryGetEdnsParamsResponseCode = "BADMODE"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadname   RadarAs112SummaryGetEdnsParamsResponseCode = "BADNAME"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadalg    RadarAs112SummaryGetEdnsParamsResponseCode = "BADALG"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadtrunc  RadarAs112SummaryGetEdnsParamsResponseCode = "BADTRUNC"
	RadarAs112SummaryGetEdnsParamsResponseCodeBadcookie RadarAs112SummaryGetEdnsParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112SummaryGetEdnsParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetEdnsParamsResponseCodeNoerror, RadarAs112SummaryGetEdnsParamsResponseCodeFormerr, RadarAs112SummaryGetEdnsParamsResponseCodeServfail, RadarAs112SummaryGetEdnsParamsResponseCodeNxdomain, RadarAs112SummaryGetEdnsParamsResponseCodeNotimp, RadarAs112SummaryGetEdnsParamsResponseCodeRefused, RadarAs112SummaryGetEdnsParamsResponseCodeYxdomain, RadarAs112SummaryGetEdnsParamsResponseCodeYxrrset, RadarAs112SummaryGetEdnsParamsResponseCodeNxrrset, RadarAs112SummaryGetEdnsParamsResponseCodeNotauth, RadarAs112SummaryGetEdnsParamsResponseCodeNotzone, RadarAs112SummaryGetEdnsParamsResponseCodeBadsig, RadarAs112SummaryGetEdnsParamsResponseCodeBadkey, RadarAs112SummaryGetEdnsParamsResponseCodeBadtime, RadarAs112SummaryGetEdnsParamsResponseCodeBadmode, RadarAs112SummaryGetEdnsParamsResponseCodeBadname, RadarAs112SummaryGetEdnsParamsResponseCodeBadalg, RadarAs112SummaryGetEdnsParamsResponseCodeBadtrunc, RadarAs112SummaryGetEdnsParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112SummaryGetIPVersionParams struct {
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
	Format param.Field[RadarAs112SummaryGetIPVersionParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112SummaryGetIPVersionParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112SummaryGetIPVersionParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112SummaryGetIPVersionParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112SummaryGetIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetIPVersionParamsFormat string

const (
	RadarAs112SummaryGetIPVersionParamsFormatJson RadarAs112SummaryGetIPVersionParamsFormat = "JSON"
	RadarAs112SummaryGetIPVersionParamsFormatCsv  RadarAs112SummaryGetIPVersionParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetIPVersionParamsFormatJson, RadarAs112SummaryGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112SummaryGetIPVersionParamsProtocol string

const (
	RadarAs112SummaryGetIPVersionParamsProtocolUdp   RadarAs112SummaryGetIPVersionParamsProtocol = "UDP"
	RadarAs112SummaryGetIPVersionParamsProtocolTcp   RadarAs112SummaryGetIPVersionParamsProtocol = "TCP"
	RadarAs112SummaryGetIPVersionParamsProtocolHTTPS RadarAs112SummaryGetIPVersionParamsProtocol = "HTTPS"
	RadarAs112SummaryGetIPVersionParamsProtocolTls   RadarAs112SummaryGetIPVersionParamsProtocol = "TLS"
)

func (r RadarAs112SummaryGetIPVersionParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetIPVersionParamsProtocolUdp, RadarAs112SummaryGetIPVersionParamsProtocolTcp, RadarAs112SummaryGetIPVersionParamsProtocolHTTPS, RadarAs112SummaryGetIPVersionParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112SummaryGetIPVersionParamsQueryType string

const (
	RadarAs112SummaryGetIPVersionParamsQueryTypeA          RadarAs112SummaryGetIPVersionParamsQueryType = "A"
	RadarAs112SummaryGetIPVersionParamsQueryTypeAaaa       RadarAs112SummaryGetIPVersionParamsQueryType = "AAAA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeA6         RadarAs112SummaryGetIPVersionParamsQueryType = "A6"
	RadarAs112SummaryGetIPVersionParamsQueryTypeAfsdb      RadarAs112SummaryGetIPVersionParamsQueryType = "AFSDB"
	RadarAs112SummaryGetIPVersionParamsQueryTypeAny        RadarAs112SummaryGetIPVersionParamsQueryType = "ANY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeApl        RadarAs112SummaryGetIPVersionParamsQueryType = "APL"
	RadarAs112SummaryGetIPVersionParamsQueryTypeAtma       RadarAs112SummaryGetIPVersionParamsQueryType = "ATMA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeAxfr       RadarAs112SummaryGetIPVersionParamsQueryType = "AXFR"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCaa        RadarAs112SummaryGetIPVersionParamsQueryType = "CAA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCdnskey    RadarAs112SummaryGetIPVersionParamsQueryType = "CDNSKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCds        RadarAs112SummaryGetIPVersionParamsQueryType = "CDS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCert       RadarAs112SummaryGetIPVersionParamsQueryType = "CERT"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCname      RadarAs112SummaryGetIPVersionParamsQueryType = "CNAME"
	RadarAs112SummaryGetIPVersionParamsQueryTypeCsync      RadarAs112SummaryGetIPVersionParamsQueryType = "CSYNC"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDhcid      RadarAs112SummaryGetIPVersionParamsQueryType = "DHCID"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDlv        RadarAs112SummaryGetIPVersionParamsQueryType = "DLV"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDname      RadarAs112SummaryGetIPVersionParamsQueryType = "DNAME"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDnskey     RadarAs112SummaryGetIPVersionParamsQueryType = "DNSKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDoa        RadarAs112SummaryGetIPVersionParamsQueryType = "DOA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeDs         RadarAs112SummaryGetIPVersionParamsQueryType = "DS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeEid        RadarAs112SummaryGetIPVersionParamsQueryType = "EID"
	RadarAs112SummaryGetIPVersionParamsQueryTypeEui48      RadarAs112SummaryGetIPVersionParamsQueryType = "EUI48"
	RadarAs112SummaryGetIPVersionParamsQueryTypeEui64      RadarAs112SummaryGetIPVersionParamsQueryType = "EUI64"
	RadarAs112SummaryGetIPVersionParamsQueryTypeGpos       RadarAs112SummaryGetIPVersionParamsQueryType = "GPOS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeGid        RadarAs112SummaryGetIPVersionParamsQueryType = "GID"
	RadarAs112SummaryGetIPVersionParamsQueryTypeHinfo      RadarAs112SummaryGetIPVersionParamsQueryType = "HINFO"
	RadarAs112SummaryGetIPVersionParamsQueryTypeHip        RadarAs112SummaryGetIPVersionParamsQueryType = "HIP"
	RadarAs112SummaryGetIPVersionParamsQueryTypeHTTPS      RadarAs112SummaryGetIPVersionParamsQueryType = "HTTPS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeIpseckey   RadarAs112SummaryGetIPVersionParamsQueryType = "IPSECKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeIsdn       RadarAs112SummaryGetIPVersionParamsQueryType = "ISDN"
	RadarAs112SummaryGetIPVersionParamsQueryTypeIxfr       RadarAs112SummaryGetIPVersionParamsQueryType = "IXFR"
	RadarAs112SummaryGetIPVersionParamsQueryTypeKey        RadarAs112SummaryGetIPVersionParamsQueryType = "KEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeKx         RadarAs112SummaryGetIPVersionParamsQueryType = "KX"
	RadarAs112SummaryGetIPVersionParamsQueryTypeL32        RadarAs112SummaryGetIPVersionParamsQueryType = "L32"
	RadarAs112SummaryGetIPVersionParamsQueryTypeL64        RadarAs112SummaryGetIPVersionParamsQueryType = "L64"
	RadarAs112SummaryGetIPVersionParamsQueryTypeLoc        RadarAs112SummaryGetIPVersionParamsQueryType = "LOC"
	RadarAs112SummaryGetIPVersionParamsQueryTypeLp         RadarAs112SummaryGetIPVersionParamsQueryType = "LP"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMaila      RadarAs112SummaryGetIPVersionParamsQueryType = "MAILA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMailb      RadarAs112SummaryGetIPVersionParamsQueryType = "MAILB"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMB         RadarAs112SummaryGetIPVersionParamsQueryType = "MB"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMd         RadarAs112SummaryGetIPVersionParamsQueryType = "MD"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMf         RadarAs112SummaryGetIPVersionParamsQueryType = "MF"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMg         RadarAs112SummaryGetIPVersionParamsQueryType = "MG"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMinfo      RadarAs112SummaryGetIPVersionParamsQueryType = "MINFO"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMr         RadarAs112SummaryGetIPVersionParamsQueryType = "MR"
	RadarAs112SummaryGetIPVersionParamsQueryTypeMx         RadarAs112SummaryGetIPVersionParamsQueryType = "MX"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNaptr      RadarAs112SummaryGetIPVersionParamsQueryType = "NAPTR"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNb         RadarAs112SummaryGetIPVersionParamsQueryType = "NB"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNbstat     RadarAs112SummaryGetIPVersionParamsQueryType = "NBSTAT"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNid        RadarAs112SummaryGetIPVersionParamsQueryType = "NID"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNimloc     RadarAs112SummaryGetIPVersionParamsQueryType = "NIMLOC"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNinfo      RadarAs112SummaryGetIPVersionParamsQueryType = "NINFO"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNs         RadarAs112SummaryGetIPVersionParamsQueryType = "NS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNsap       RadarAs112SummaryGetIPVersionParamsQueryType = "NSAP"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNsec       RadarAs112SummaryGetIPVersionParamsQueryType = "NSEC"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNsec3      RadarAs112SummaryGetIPVersionParamsQueryType = "NSEC3"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNsec3Param RadarAs112SummaryGetIPVersionParamsQueryType = "NSEC3PARAM"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNull       RadarAs112SummaryGetIPVersionParamsQueryType = "NULL"
	RadarAs112SummaryGetIPVersionParamsQueryTypeNxt        RadarAs112SummaryGetIPVersionParamsQueryType = "NXT"
	RadarAs112SummaryGetIPVersionParamsQueryTypeOpenpgpkey RadarAs112SummaryGetIPVersionParamsQueryType = "OPENPGPKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeOpt        RadarAs112SummaryGetIPVersionParamsQueryType = "OPT"
	RadarAs112SummaryGetIPVersionParamsQueryTypePtr        RadarAs112SummaryGetIPVersionParamsQueryType = "PTR"
	RadarAs112SummaryGetIPVersionParamsQueryTypePx         RadarAs112SummaryGetIPVersionParamsQueryType = "PX"
	RadarAs112SummaryGetIPVersionParamsQueryTypeRkey       RadarAs112SummaryGetIPVersionParamsQueryType = "RKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeRp         RadarAs112SummaryGetIPVersionParamsQueryType = "RP"
	RadarAs112SummaryGetIPVersionParamsQueryTypeRrsig      RadarAs112SummaryGetIPVersionParamsQueryType = "RRSIG"
	RadarAs112SummaryGetIPVersionParamsQueryTypeRt         RadarAs112SummaryGetIPVersionParamsQueryType = "RT"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSig        RadarAs112SummaryGetIPVersionParamsQueryType = "SIG"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSink       RadarAs112SummaryGetIPVersionParamsQueryType = "SINK"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSmimea     RadarAs112SummaryGetIPVersionParamsQueryType = "SMIMEA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSoa        RadarAs112SummaryGetIPVersionParamsQueryType = "SOA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSpf        RadarAs112SummaryGetIPVersionParamsQueryType = "SPF"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSrv        RadarAs112SummaryGetIPVersionParamsQueryType = "SRV"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSshfp      RadarAs112SummaryGetIPVersionParamsQueryType = "SSHFP"
	RadarAs112SummaryGetIPVersionParamsQueryTypeSvcb       RadarAs112SummaryGetIPVersionParamsQueryType = "SVCB"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTa         RadarAs112SummaryGetIPVersionParamsQueryType = "TA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTalink     RadarAs112SummaryGetIPVersionParamsQueryType = "TALINK"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTkey       RadarAs112SummaryGetIPVersionParamsQueryType = "TKEY"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTlsa       RadarAs112SummaryGetIPVersionParamsQueryType = "TLSA"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTsig       RadarAs112SummaryGetIPVersionParamsQueryType = "TSIG"
	RadarAs112SummaryGetIPVersionParamsQueryTypeTxt        RadarAs112SummaryGetIPVersionParamsQueryType = "TXT"
	RadarAs112SummaryGetIPVersionParamsQueryTypeUinfo      RadarAs112SummaryGetIPVersionParamsQueryType = "UINFO"
	RadarAs112SummaryGetIPVersionParamsQueryTypeUid        RadarAs112SummaryGetIPVersionParamsQueryType = "UID"
	RadarAs112SummaryGetIPVersionParamsQueryTypeUnspec     RadarAs112SummaryGetIPVersionParamsQueryType = "UNSPEC"
	RadarAs112SummaryGetIPVersionParamsQueryTypeUri        RadarAs112SummaryGetIPVersionParamsQueryType = "URI"
	RadarAs112SummaryGetIPVersionParamsQueryTypeWks        RadarAs112SummaryGetIPVersionParamsQueryType = "WKS"
	RadarAs112SummaryGetIPVersionParamsQueryTypeX25        RadarAs112SummaryGetIPVersionParamsQueryType = "X25"
	RadarAs112SummaryGetIPVersionParamsQueryTypeZonemd     RadarAs112SummaryGetIPVersionParamsQueryType = "ZONEMD"
)

func (r RadarAs112SummaryGetIPVersionParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetIPVersionParamsQueryTypeA, RadarAs112SummaryGetIPVersionParamsQueryTypeAaaa, RadarAs112SummaryGetIPVersionParamsQueryTypeA6, RadarAs112SummaryGetIPVersionParamsQueryTypeAfsdb, RadarAs112SummaryGetIPVersionParamsQueryTypeAny, RadarAs112SummaryGetIPVersionParamsQueryTypeApl, RadarAs112SummaryGetIPVersionParamsQueryTypeAtma, RadarAs112SummaryGetIPVersionParamsQueryTypeAxfr, RadarAs112SummaryGetIPVersionParamsQueryTypeCaa, RadarAs112SummaryGetIPVersionParamsQueryTypeCdnskey, RadarAs112SummaryGetIPVersionParamsQueryTypeCds, RadarAs112SummaryGetIPVersionParamsQueryTypeCert, RadarAs112SummaryGetIPVersionParamsQueryTypeCname, RadarAs112SummaryGetIPVersionParamsQueryTypeCsync, RadarAs112SummaryGetIPVersionParamsQueryTypeDhcid, RadarAs112SummaryGetIPVersionParamsQueryTypeDlv, RadarAs112SummaryGetIPVersionParamsQueryTypeDname, RadarAs112SummaryGetIPVersionParamsQueryTypeDnskey, RadarAs112SummaryGetIPVersionParamsQueryTypeDoa, RadarAs112SummaryGetIPVersionParamsQueryTypeDs, RadarAs112SummaryGetIPVersionParamsQueryTypeEid, RadarAs112SummaryGetIPVersionParamsQueryTypeEui48, RadarAs112SummaryGetIPVersionParamsQueryTypeEui64, RadarAs112SummaryGetIPVersionParamsQueryTypeGpos, RadarAs112SummaryGetIPVersionParamsQueryTypeGid, RadarAs112SummaryGetIPVersionParamsQueryTypeHinfo, RadarAs112SummaryGetIPVersionParamsQueryTypeHip, RadarAs112SummaryGetIPVersionParamsQueryTypeHTTPS, RadarAs112SummaryGetIPVersionParamsQueryTypeIpseckey, RadarAs112SummaryGetIPVersionParamsQueryTypeIsdn, RadarAs112SummaryGetIPVersionParamsQueryTypeIxfr, RadarAs112SummaryGetIPVersionParamsQueryTypeKey, RadarAs112SummaryGetIPVersionParamsQueryTypeKx, RadarAs112SummaryGetIPVersionParamsQueryTypeL32, RadarAs112SummaryGetIPVersionParamsQueryTypeL64, RadarAs112SummaryGetIPVersionParamsQueryTypeLoc, RadarAs112SummaryGetIPVersionParamsQueryTypeLp, RadarAs112SummaryGetIPVersionParamsQueryTypeMaila, RadarAs112SummaryGetIPVersionParamsQueryTypeMailb, RadarAs112SummaryGetIPVersionParamsQueryTypeMB, RadarAs112SummaryGetIPVersionParamsQueryTypeMd, RadarAs112SummaryGetIPVersionParamsQueryTypeMf, RadarAs112SummaryGetIPVersionParamsQueryTypeMg, RadarAs112SummaryGetIPVersionParamsQueryTypeMinfo, RadarAs112SummaryGetIPVersionParamsQueryTypeMr, RadarAs112SummaryGetIPVersionParamsQueryTypeMx, RadarAs112SummaryGetIPVersionParamsQueryTypeNaptr, RadarAs112SummaryGetIPVersionParamsQueryTypeNb, RadarAs112SummaryGetIPVersionParamsQueryTypeNbstat, RadarAs112SummaryGetIPVersionParamsQueryTypeNid, RadarAs112SummaryGetIPVersionParamsQueryTypeNimloc, RadarAs112SummaryGetIPVersionParamsQueryTypeNinfo, RadarAs112SummaryGetIPVersionParamsQueryTypeNs, RadarAs112SummaryGetIPVersionParamsQueryTypeNsap, RadarAs112SummaryGetIPVersionParamsQueryTypeNsec, RadarAs112SummaryGetIPVersionParamsQueryTypeNsec3, RadarAs112SummaryGetIPVersionParamsQueryTypeNsec3Param, RadarAs112SummaryGetIPVersionParamsQueryTypeNull, RadarAs112SummaryGetIPVersionParamsQueryTypeNxt, RadarAs112SummaryGetIPVersionParamsQueryTypeOpenpgpkey, RadarAs112SummaryGetIPVersionParamsQueryTypeOpt, RadarAs112SummaryGetIPVersionParamsQueryTypePtr, RadarAs112SummaryGetIPVersionParamsQueryTypePx, RadarAs112SummaryGetIPVersionParamsQueryTypeRkey, RadarAs112SummaryGetIPVersionParamsQueryTypeRp, RadarAs112SummaryGetIPVersionParamsQueryTypeRrsig, RadarAs112SummaryGetIPVersionParamsQueryTypeRt, RadarAs112SummaryGetIPVersionParamsQueryTypeSig, RadarAs112SummaryGetIPVersionParamsQueryTypeSink, RadarAs112SummaryGetIPVersionParamsQueryTypeSmimea, RadarAs112SummaryGetIPVersionParamsQueryTypeSoa, RadarAs112SummaryGetIPVersionParamsQueryTypeSpf, RadarAs112SummaryGetIPVersionParamsQueryTypeSrv, RadarAs112SummaryGetIPVersionParamsQueryTypeSshfp, RadarAs112SummaryGetIPVersionParamsQueryTypeSvcb, RadarAs112SummaryGetIPVersionParamsQueryTypeTa, RadarAs112SummaryGetIPVersionParamsQueryTypeTalink, RadarAs112SummaryGetIPVersionParamsQueryTypeTkey, RadarAs112SummaryGetIPVersionParamsQueryTypeTlsa, RadarAs112SummaryGetIPVersionParamsQueryTypeTsig, RadarAs112SummaryGetIPVersionParamsQueryTypeTxt, RadarAs112SummaryGetIPVersionParamsQueryTypeUinfo, RadarAs112SummaryGetIPVersionParamsQueryTypeUid, RadarAs112SummaryGetIPVersionParamsQueryTypeUnspec, RadarAs112SummaryGetIPVersionParamsQueryTypeUri, RadarAs112SummaryGetIPVersionParamsQueryTypeWks, RadarAs112SummaryGetIPVersionParamsQueryTypeX25, RadarAs112SummaryGetIPVersionParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112SummaryGetIPVersionParamsResponseCode string

const (
	RadarAs112SummaryGetIPVersionParamsResponseCodeNoerror   RadarAs112SummaryGetIPVersionParamsResponseCode = "NOERROR"
	RadarAs112SummaryGetIPVersionParamsResponseCodeFormerr   RadarAs112SummaryGetIPVersionParamsResponseCode = "FORMERR"
	RadarAs112SummaryGetIPVersionParamsResponseCodeServfail  RadarAs112SummaryGetIPVersionParamsResponseCode = "SERVFAIL"
	RadarAs112SummaryGetIPVersionParamsResponseCodeNxdomain  RadarAs112SummaryGetIPVersionParamsResponseCode = "NXDOMAIN"
	RadarAs112SummaryGetIPVersionParamsResponseCodeNotimp    RadarAs112SummaryGetIPVersionParamsResponseCode = "NOTIMP"
	RadarAs112SummaryGetIPVersionParamsResponseCodeRefused   RadarAs112SummaryGetIPVersionParamsResponseCode = "REFUSED"
	RadarAs112SummaryGetIPVersionParamsResponseCodeYxdomain  RadarAs112SummaryGetIPVersionParamsResponseCode = "YXDOMAIN"
	RadarAs112SummaryGetIPVersionParamsResponseCodeYxrrset   RadarAs112SummaryGetIPVersionParamsResponseCode = "YXRRSET"
	RadarAs112SummaryGetIPVersionParamsResponseCodeNxrrset   RadarAs112SummaryGetIPVersionParamsResponseCode = "NXRRSET"
	RadarAs112SummaryGetIPVersionParamsResponseCodeNotauth   RadarAs112SummaryGetIPVersionParamsResponseCode = "NOTAUTH"
	RadarAs112SummaryGetIPVersionParamsResponseCodeNotzone   RadarAs112SummaryGetIPVersionParamsResponseCode = "NOTZONE"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadsig    RadarAs112SummaryGetIPVersionParamsResponseCode = "BADSIG"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadkey    RadarAs112SummaryGetIPVersionParamsResponseCode = "BADKEY"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadtime   RadarAs112SummaryGetIPVersionParamsResponseCode = "BADTIME"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadmode   RadarAs112SummaryGetIPVersionParamsResponseCode = "BADMODE"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadname   RadarAs112SummaryGetIPVersionParamsResponseCode = "BADNAME"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadalg    RadarAs112SummaryGetIPVersionParamsResponseCode = "BADALG"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadtrunc  RadarAs112SummaryGetIPVersionParamsResponseCode = "BADTRUNC"
	RadarAs112SummaryGetIPVersionParamsResponseCodeBadcookie RadarAs112SummaryGetIPVersionParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112SummaryGetIPVersionParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetIPVersionParamsResponseCodeNoerror, RadarAs112SummaryGetIPVersionParamsResponseCodeFormerr, RadarAs112SummaryGetIPVersionParamsResponseCodeServfail, RadarAs112SummaryGetIPVersionParamsResponseCodeNxdomain, RadarAs112SummaryGetIPVersionParamsResponseCodeNotimp, RadarAs112SummaryGetIPVersionParamsResponseCodeRefused, RadarAs112SummaryGetIPVersionParamsResponseCodeYxdomain, RadarAs112SummaryGetIPVersionParamsResponseCodeYxrrset, RadarAs112SummaryGetIPVersionParamsResponseCodeNxrrset, RadarAs112SummaryGetIPVersionParamsResponseCodeNotauth, RadarAs112SummaryGetIPVersionParamsResponseCodeNotzone, RadarAs112SummaryGetIPVersionParamsResponseCodeBadsig, RadarAs112SummaryGetIPVersionParamsResponseCodeBadkey, RadarAs112SummaryGetIPVersionParamsResponseCodeBadtime, RadarAs112SummaryGetIPVersionParamsResponseCodeBadmode, RadarAs112SummaryGetIPVersionParamsResponseCodeBadname, RadarAs112SummaryGetIPVersionParamsResponseCodeBadalg, RadarAs112SummaryGetIPVersionParamsResponseCodeBadtrunc, RadarAs112SummaryGetIPVersionParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112SummaryGetProtocolParams struct {
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
	Format param.Field[RadarAs112SummaryGetProtocolParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112SummaryGetProtocolParamsQueryType] `query:"queryType"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112SummaryGetProtocolParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112SummaryGetProtocolParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryGetProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetProtocolParamsFormat string

const (
	RadarAs112SummaryGetProtocolParamsFormatJson RadarAs112SummaryGetProtocolParamsFormat = "JSON"
	RadarAs112SummaryGetProtocolParamsFormatCsv  RadarAs112SummaryGetProtocolParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetProtocolParamsFormatJson, RadarAs112SummaryGetProtocolParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112SummaryGetProtocolParamsQueryType string

const (
	RadarAs112SummaryGetProtocolParamsQueryTypeA          RadarAs112SummaryGetProtocolParamsQueryType = "A"
	RadarAs112SummaryGetProtocolParamsQueryTypeAaaa       RadarAs112SummaryGetProtocolParamsQueryType = "AAAA"
	RadarAs112SummaryGetProtocolParamsQueryTypeA6         RadarAs112SummaryGetProtocolParamsQueryType = "A6"
	RadarAs112SummaryGetProtocolParamsQueryTypeAfsdb      RadarAs112SummaryGetProtocolParamsQueryType = "AFSDB"
	RadarAs112SummaryGetProtocolParamsQueryTypeAny        RadarAs112SummaryGetProtocolParamsQueryType = "ANY"
	RadarAs112SummaryGetProtocolParamsQueryTypeApl        RadarAs112SummaryGetProtocolParamsQueryType = "APL"
	RadarAs112SummaryGetProtocolParamsQueryTypeAtma       RadarAs112SummaryGetProtocolParamsQueryType = "ATMA"
	RadarAs112SummaryGetProtocolParamsQueryTypeAxfr       RadarAs112SummaryGetProtocolParamsQueryType = "AXFR"
	RadarAs112SummaryGetProtocolParamsQueryTypeCaa        RadarAs112SummaryGetProtocolParamsQueryType = "CAA"
	RadarAs112SummaryGetProtocolParamsQueryTypeCdnskey    RadarAs112SummaryGetProtocolParamsQueryType = "CDNSKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeCds        RadarAs112SummaryGetProtocolParamsQueryType = "CDS"
	RadarAs112SummaryGetProtocolParamsQueryTypeCert       RadarAs112SummaryGetProtocolParamsQueryType = "CERT"
	RadarAs112SummaryGetProtocolParamsQueryTypeCname      RadarAs112SummaryGetProtocolParamsQueryType = "CNAME"
	RadarAs112SummaryGetProtocolParamsQueryTypeCsync      RadarAs112SummaryGetProtocolParamsQueryType = "CSYNC"
	RadarAs112SummaryGetProtocolParamsQueryTypeDhcid      RadarAs112SummaryGetProtocolParamsQueryType = "DHCID"
	RadarAs112SummaryGetProtocolParamsQueryTypeDlv        RadarAs112SummaryGetProtocolParamsQueryType = "DLV"
	RadarAs112SummaryGetProtocolParamsQueryTypeDname      RadarAs112SummaryGetProtocolParamsQueryType = "DNAME"
	RadarAs112SummaryGetProtocolParamsQueryTypeDnskey     RadarAs112SummaryGetProtocolParamsQueryType = "DNSKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeDoa        RadarAs112SummaryGetProtocolParamsQueryType = "DOA"
	RadarAs112SummaryGetProtocolParamsQueryTypeDs         RadarAs112SummaryGetProtocolParamsQueryType = "DS"
	RadarAs112SummaryGetProtocolParamsQueryTypeEid        RadarAs112SummaryGetProtocolParamsQueryType = "EID"
	RadarAs112SummaryGetProtocolParamsQueryTypeEui48      RadarAs112SummaryGetProtocolParamsQueryType = "EUI48"
	RadarAs112SummaryGetProtocolParamsQueryTypeEui64      RadarAs112SummaryGetProtocolParamsQueryType = "EUI64"
	RadarAs112SummaryGetProtocolParamsQueryTypeGpos       RadarAs112SummaryGetProtocolParamsQueryType = "GPOS"
	RadarAs112SummaryGetProtocolParamsQueryTypeGid        RadarAs112SummaryGetProtocolParamsQueryType = "GID"
	RadarAs112SummaryGetProtocolParamsQueryTypeHinfo      RadarAs112SummaryGetProtocolParamsQueryType = "HINFO"
	RadarAs112SummaryGetProtocolParamsQueryTypeHip        RadarAs112SummaryGetProtocolParamsQueryType = "HIP"
	RadarAs112SummaryGetProtocolParamsQueryTypeHTTPS      RadarAs112SummaryGetProtocolParamsQueryType = "HTTPS"
	RadarAs112SummaryGetProtocolParamsQueryTypeIpseckey   RadarAs112SummaryGetProtocolParamsQueryType = "IPSECKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeIsdn       RadarAs112SummaryGetProtocolParamsQueryType = "ISDN"
	RadarAs112SummaryGetProtocolParamsQueryTypeIxfr       RadarAs112SummaryGetProtocolParamsQueryType = "IXFR"
	RadarAs112SummaryGetProtocolParamsQueryTypeKey        RadarAs112SummaryGetProtocolParamsQueryType = "KEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeKx         RadarAs112SummaryGetProtocolParamsQueryType = "KX"
	RadarAs112SummaryGetProtocolParamsQueryTypeL32        RadarAs112SummaryGetProtocolParamsQueryType = "L32"
	RadarAs112SummaryGetProtocolParamsQueryTypeL64        RadarAs112SummaryGetProtocolParamsQueryType = "L64"
	RadarAs112SummaryGetProtocolParamsQueryTypeLoc        RadarAs112SummaryGetProtocolParamsQueryType = "LOC"
	RadarAs112SummaryGetProtocolParamsQueryTypeLp         RadarAs112SummaryGetProtocolParamsQueryType = "LP"
	RadarAs112SummaryGetProtocolParamsQueryTypeMaila      RadarAs112SummaryGetProtocolParamsQueryType = "MAILA"
	RadarAs112SummaryGetProtocolParamsQueryTypeMailb      RadarAs112SummaryGetProtocolParamsQueryType = "MAILB"
	RadarAs112SummaryGetProtocolParamsQueryTypeMB         RadarAs112SummaryGetProtocolParamsQueryType = "MB"
	RadarAs112SummaryGetProtocolParamsQueryTypeMd         RadarAs112SummaryGetProtocolParamsQueryType = "MD"
	RadarAs112SummaryGetProtocolParamsQueryTypeMf         RadarAs112SummaryGetProtocolParamsQueryType = "MF"
	RadarAs112SummaryGetProtocolParamsQueryTypeMg         RadarAs112SummaryGetProtocolParamsQueryType = "MG"
	RadarAs112SummaryGetProtocolParamsQueryTypeMinfo      RadarAs112SummaryGetProtocolParamsQueryType = "MINFO"
	RadarAs112SummaryGetProtocolParamsQueryTypeMr         RadarAs112SummaryGetProtocolParamsQueryType = "MR"
	RadarAs112SummaryGetProtocolParamsQueryTypeMx         RadarAs112SummaryGetProtocolParamsQueryType = "MX"
	RadarAs112SummaryGetProtocolParamsQueryTypeNaptr      RadarAs112SummaryGetProtocolParamsQueryType = "NAPTR"
	RadarAs112SummaryGetProtocolParamsQueryTypeNb         RadarAs112SummaryGetProtocolParamsQueryType = "NB"
	RadarAs112SummaryGetProtocolParamsQueryTypeNbstat     RadarAs112SummaryGetProtocolParamsQueryType = "NBSTAT"
	RadarAs112SummaryGetProtocolParamsQueryTypeNid        RadarAs112SummaryGetProtocolParamsQueryType = "NID"
	RadarAs112SummaryGetProtocolParamsQueryTypeNimloc     RadarAs112SummaryGetProtocolParamsQueryType = "NIMLOC"
	RadarAs112SummaryGetProtocolParamsQueryTypeNinfo      RadarAs112SummaryGetProtocolParamsQueryType = "NINFO"
	RadarAs112SummaryGetProtocolParamsQueryTypeNs         RadarAs112SummaryGetProtocolParamsQueryType = "NS"
	RadarAs112SummaryGetProtocolParamsQueryTypeNsap       RadarAs112SummaryGetProtocolParamsQueryType = "NSAP"
	RadarAs112SummaryGetProtocolParamsQueryTypeNsec       RadarAs112SummaryGetProtocolParamsQueryType = "NSEC"
	RadarAs112SummaryGetProtocolParamsQueryTypeNsec3      RadarAs112SummaryGetProtocolParamsQueryType = "NSEC3"
	RadarAs112SummaryGetProtocolParamsQueryTypeNsec3Param RadarAs112SummaryGetProtocolParamsQueryType = "NSEC3PARAM"
	RadarAs112SummaryGetProtocolParamsQueryTypeNull       RadarAs112SummaryGetProtocolParamsQueryType = "NULL"
	RadarAs112SummaryGetProtocolParamsQueryTypeNxt        RadarAs112SummaryGetProtocolParamsQueryType = "NXT"
	RadarAs112SummaryGetProtocolParamsQueryTypeOpenpgpkey RadarAs112SummaryGetProtocolParamsQueryType = "OPENPGPKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeOpt        RadarAs112SummaryGetProtocolParamsQueryType = "OPT"
	RadarAs112SummaryGetProtocolParamsQueryTypePtr        RadarAs112SummaryGetProtocolParamsQueryType = "PTR"
	RadarAs112SummaryGetProtocolParamsQueryTypePx         RadarAs112SummaryGetProtocolParamsQueryType = "PX"
	RadarAs112SummaryGetProtocolParamsQueryTypeRkey       RadarAs112SummaryGetProtocolParamsQueryType = "RKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeRp         RadarAs112SummaryGetProtocolParamsQueryType = "RP"
	RadarAs112SummaryGetProtocolParamsQueryTypeRrsig      RadarAs112SummaryGetProtocolParamsQueryType = "RRSIG"
	RadarAs112SummaryGetProtocolParamsQueryTypeRt         RadarAs112SummaryGetProtocolParamsQueryType = "RT"
	RadarAs112SummaryGetProtocolParamsQueryTypeSig        RadarAs112SummaryGetProtocolParamsQueryType = "SIG"
	RadarAs112SummaryGetProtocolParamsQueryTypeSink       RadarAs112SummaryGetProtocolParamsQueryType = "SINK"
	RadarAs112SummaryGetProtocolParamsQueryTypeSmimea     RadarAs112SummaryGetProtocolParamsQueryType = "SMIMEA"
	RadarAs112SummaryGetProtocolParamsQueryTypeSoa        RadarAs112SummaryGetProtocolParamsQueryType = "SOA"
	RadarAs112SummaryGetProtocolParamsQueryTypeSpf        RadarAs112SummaryGetProtocolParamsQueryType = "SPF"
	RadarAs112SummaryGetProtocolParamsQueryTypeSrv        RadarAs112SummaryGetProtocolParamsQueryType = "SRV"
	RadarAs112SummaryGetProtocolParamsQueryTypeSshfp      RadarAs112SummaryGetProtocolParamsQueryType = "SSHFP"
	RadarAs112SummaryGetProtocolParamsQueryTypeSvcb       RadarAs112SummaryGetProtocolParamsQueryType = "SVCB"
	RadarAs112SummaryGetProtocolParamsQueryTypeTa         RadarAs112SummaryGetProtocolParamsQueryType = "TA"
	RadarAs112SummaryGetProtocolParamsQueryTypeTalink     RadarAs112SummaryGetProtocolParamsQueryType = "TALINK"
	RadarAs112SummaryGetProtocolParamsQueryTypeTkey       RadarAs112SummaryGetProtocolParamsQueryType = "TKEY"
	RadarAs112SummaryGetProtocolParamsQueryTypeTlsa       RadarAs112SummaryGetProtocolParamsQueryType = "TLSA"
	RadarAs112SummaryGetProtocolParamsQueryTypeTsig       RadarAs112SummaryGetProtocolParamsQueryType = "TSIG"
	RadarAs112SummaryGetProtocolParamsQueryTypeTxt        RadarAs112SummaryGetProtocolParamsQueryType = "TXT"
	RadarAs112SummaryGetProtocolParamsQueryTypeUinfo      RadarAs112SummaryGetProtocolParamsQueryType = "UINFO"
	RadarAs112SummaryGetProtocolParamsQueryTypeUid        RadarAs112SummaryGetProtocolParamsQueryType = "UID"
	RadarAs112SummaryGetProtocolParamsQueryTypeUnspec     RadarAs112SummaryGetProtocolParamsQueryType = "UNSPEC"
	RadarAs112SummaryGetProtocolParamsQueryTypeUri        RadarAs112SummaryGetProtocolParamsQueryType = "URI"
	RadarAs112SummaryGetProtocolParamsQueryTypeWks        RadarAs112SummaryGetProtocolParamsQueryType = "WKS"
	RadarAs112SummaryGetProtocolParamsQueryTypeX25        RadarAs112SummaryGetProtocolParamsQueryType = "X25"
	RadarAs112SummaryGetProtocolParamsQueryTypeZonemd     RadarAs112SummaryGetProtocolParamsQueryType = "ZONEMD"
)

func (r RadarAs112SummaryGetProtocolParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetProtocolParamsQueryTypeA, RadarAs112SummaryGetProtocolParamsQueryTypeAaaa, RadarAs112SummaryGetProtocolParamsQueryTypeA6, RadarAs112SummaryGetProtocolParamsQueryTypeAfsdb, RadarAs112SummaryGetProtocolParamsQueryTypeAny, RadarAs112SummaryGetProtocolParamsQueryTypeApl, RadarAs112SummaryGetProtocolParamsQueryTypeAtma, RadarAs112SummaryGetProtocolParamsQueryTypeAxfr, RadarAs112SummaryGetProtocolParamsQueryTypeCaa, RadarAs112SummaryGetProtocolParamsQueryTypeCdnskey, RadarAs112SummaryGetProtocolParamsQueryTypeCds, RadarAs112SummaryGetProtocolParamsQueryTypeCert, RadarAs112SummaryGetProtocolParamsQueryTypeCname, RadarAs112SummaryGetProtocolParamsQueryTypeCsync, RadarAs112SummaryGetProtocolParamsQueryTypeDhcid, RadarAs112SummaryGetProtocolParamsQueryTypeDlv, RadarAs112SummaryGetProtocolParamsQueryTypeDname, RadarAs112SummaryGetProtocolParamsQueryTypeDnskey, RadarAs112SummaryGetProtocolParamsQueryTypeDoa, RadarAs112SummaryGetProtocolParamsQueryTypeDs, RadarAs112SummaryGetProtocolParamsQueryTypeEid, RadarAs112SummaryGetProtocolParamsQueryTypeEui48, RadarAs112SummaryGetProtocolParamsQueryTypeEui64, RadarAs112SummaryGetProtocolParamsQueryTypeGpos, RadarAs112SummaryGetProtocolParamsQueryTypeGid, RadarAs112SummaryGetProtocolParamsQueryTypeHinfo, RadarAs112SummaryGetProtocolParamsQueryTypeHip, RadarAs112SummaryGetProtocolParamsQueryTypeHTTPS, RadarAs112SummaryGetProtocolParamsQueryTypeIpseckey, RadarAs112SummaryGetProtocolParamsQueryTypeIsdn, RadarAs112SummaryGetProtocolParamsQueryTypeIxfr, RadarAs112SummaryGetProtocolParamsQueryTypeKey, RadarAs112SummaryGetProtocolParamsQueryTypeKx, RadarAs112SummaryGetProtocolParamsQueryTypeL32, RadarAs112SummaryGetProtocolParamsQueryTypeL64, RadarAs112SummaryGetProtocolParamsQueryTypeLoc, RadarAs112SummaryGetProtocolParamsQueryTypeLp, RadarAs112SummaryGetProtocolParamsQueryTypeMaila, RadarAs112SummaryGetProtocolParamsQueryTypeMailb, RadarAs112SummaryGetProtocolParamsQueryTypeMB, RadarAs112SummaryGetProtocolParamsQueryTypeMd, RadarAs112SummaryGetProtocolParamsQueryTypeMf, RadarAs112SummaryGetProtocolParamsQueryTypeMg, RadarAs112SummaryGetProtocolParamsQueryTypeMinfo, RadarAs112SummaryGetProtocolParamsQueryTypeMr, RadarAs112SummaryGetProtocolParamsQueryTypeMx, RadarAs112SummaryGetProtocolParamsQueryTypeNaptr, RadarAs112SummaryGetProtocolParamsQueryTypeNb, RadarAs112SummaryGetProtocolParamsQueryTypeNbstat, RadarAs112SummaryGetProtocolParamsQueryTypeNid, RadarAs112SummaryGetProtocolParamsQueryTypeNimloc, RadarAs112SummaryGetProtocolParamsQueryTypeNinfo, RadarAs112SummaryGetProtocolParamsQueryTypeNs, RadarAs112SummaryGetProtocolParamsQueryTypeNsap, RadarAs112SummaryGetProtocolParamsQueryTypeNsec, RadarAs112SummaryGetProtocolParamsQueryTypeNsec3, RadarAs112SummaryGetProtocolParamsQueryTypeNsec3Param, RadarAs112SummaryGetProtocolParamsQueryTypeNull, RadarAs112SummaryGetProtocolParamsQueryTypeNxt, RadarAs112SummaryGetProtocolParamsQueryTypeOpenpgpkey, RadarAs112SummaryGetProtocolParamsQueryTypeOpt, RadarAs112SummaryGetProtocolParamsQueryTypePtr, RadarAs112SummaryGetProtocolParamsQueryTypePx, RadarAs112SummaryGetProtocolParamsQueryTypeRkey, RadarAs112SummaryGetProtocolParamsQueryTypeRp, RadarAs112SummaryGetProtocolParamsQueryTypeRrsig, RadarAs112SummaryGetProtocolParamsQueryTypeRt, RadarAs112SummaryGetProtocolParamsQueryTypeSig, RadarAs112SummaryGetProtocolParamsQueryTypeSink, RadarAs112SummaryGetProtocolParamsQueryTypeSmimea, RadarAs112SummaryGetProtocolParamsQueryTypeSoa, RadarAs112SummaryGetProtocolParamsQueryTypeSpf, RadarAs112SummaryGetProtocolParamsQueryTypeSrv, RadarAs112SummaryGetProtocolParamsQueryTypeSshfp, RadarAs112SummaryGetProtocolParamsQueryTypeSvcb, RadarAs112SummaryGetProtocolParamsQueryTypeTa, RadarAs112SummaryGetProtocolParamsQueryTypeTalink, RadarAs112SummaryGetProtocolParamsQueryTypeTkey, RadarAs112SummaryGetProtocolParamsQueryTypeTlsa, RadarAs112SummaryGetProtocolParamsQueryTypeTsig, RadarAs112SummaryGetProtocolParamsQueryTypeTxt, RadarAs112SummaryGetProtocolParamsQueryTypeUinfo, RadarAs112SummaryGetProtocolParamsQueryTypeUid, RadarAs112SummaryGetProtocolParamsQueryTypeUnspec, RadarAs112SummaryGetProtocolParamsQueryTypeUri, RadarAs112SummaryGetProtocolParamsQueryTypeWks, RadarAs112SummaryGetProtocolParamsQueryTypeX25, RadarAs112SummaryGetProtocolParamsQueryTypeZonemd:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112SummaryGetProtocolParamsResponseCode string

const (
	RadarAs112SummaryGetProtocolParamsResponseCodeNoerror   RadarAs112SummaryGetProtocolParamsResponseCode = "NOERROR"
	RadarAs112SummaryGetProtocolParamsResponseCodeFormerr   RadarAs112SummaryGetProtocolParamsResponseCode = "FORMERR"
	RadarAs112SummaryGetProtocolParamsResponseCodeServfail  RadarAs112SummaryGetProtocolParamsResponseCode = "SERVFAIL"
	RadarAs112SummaryGetProtocolParamsResponseCodeNxdomain  RadarAs112SummaryGetProtocolParamsResponseCode = "NXDOMAIN"
	RadarAs112SummaryGetProtocolParamsResponseCodeNotimp    RadarAs112SummaryGetProtocolParamsResponseCode = "NOTIMP"
	RadarAs112SummaryGetProtocolParamsResponseCodeRefused   RadarAs112SummaryGetProtocolParamsResponseCode = "REFUSED"
	RadarAs112SummaryGetProtocolParamsResponseCodeYxdomain  RadarAs112SummaryGetProtocolParamsResponseCode = "YXDOMAIN"
	RadarAs112SummaryGetProtocolParamsResponseCodeYxrrset   RadarAs112SummaryGetProtocolParamsResponseCode = "YXRRSET"
	RadarAs112SummaryGetProtocolParamsResponseCodeNxrrset   RadarAs112SummaryGetProtocolParamsResponseCode = "NXRRSET"
	RadarAs112SummaryGetProtocolParamsResponseCodeNotauth   RadarAs112SummaryGetProtocolParamsResponseCode = "NOTAUTH"
	RadarAs112SummaryGetProtocolParamsResponseCodeNotzone   RadarAs112SummaryGetProtocolParamsResponseCode = "NOTZONE"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadsig    RadarAs112SummaryGetProtocolParamsResponseCode = "BADSIG"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadkey    RadarAs112SummaryGetProtocolParamsResponseCode = "BADKEY"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadtime   RadarAs112SummaryGetProtocolParamsResponseCode = "BADTIME"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadmode   RadarAs112SummaryGetProtocolParamsResponseCode = "BADMODE"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadname   RadarAs112SummaryGetProtocolParamsResponseCode = "BADNAME"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadalg    RadarAs112SummaryGetProtocolParamsResponseCode = "BADALG"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadtrunc  RadarAs112SummaryGetProtocolParamsResponseCode = "BADTRUNC"
	RadarAs112SummaryGetProtocolParamsResponseCodeBadcookie RadarAs112SummaryGetProtocolParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112SummaryGetProtocolParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetProtocolParamsResponseCodeNoerror, RadarAs112SummaryGetProtocolParamsResponseCodeFormerr, RadarAs112SummaryGetProtocolParamsResponseCodeServfail, RadarAs112SummaryGetProtocolParamsResponseCodeNxdomain, RadarAs112SummaryGetProtocolParamsResponseCodeNotimp, RadarAs112SummaryGetProtocolParamsResponseCodeRefused, RadarAs112SummaryGetProtocolParamsResponseCodeYxdomain, RadarAs112SummaryGetProtocolParamsResponseCodeYxrrset, RadarAs112SummaryGetProtocolParamsResponseCodeNxrrset, RadarAs112SummaryGetProtocolParamsResponseCodeNotauth, RadarAs112SummaryGetProtocolParamsResponseCodeNotzone, RadarAs112SummaryGetProtocolParamsResponseCodeBadsig, RadarAs112SummaryGetProtocolParamsResponseCodeBadkey, RadarAs112SummaryGetProtocolParamsResponseCodeBadtime, RadarAs112SummaryGetProtocolParamsResponseCodeBadmode, RadarAs112SummaryGetProtocolParamsResponseCodeBadname, RadarAs112SummaryGetProtocolParamsResponseCodeBadalg, RadarAs112SummaryGetProtocolParamsResponseCodeBadtrunc, RadarAs112SummaryGetProtocolParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112SummaryGetQueryTypeParams struct {
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
	Format param.Field[RadarAs112SummaryGetQueryTypeParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112SummaryGetQueryTypeParamsProtocol] `query:"protocol"`
	// Filters results by DNS response code.
	ResponseCode param.Field[RadarAs112SummaryGetQueryTypeParamsResponseCode] `query:"responseCode"`
}

// URLQuery serializes [RadarAs112SummaryGetQueryTypeParams]'s query parameters as
// `url.Values`.
func (r RadarAs112SummaryGetQueryTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetQueryTypeParamsFormat string

const (
	RadarAs112SummaryGetQueryTypeParamsFormatJson RadarAs112SummaryGetQueryTypeParamsFormat = "JSON"
	RadarAs112SummaryGetQueryTypeParamsFormatCsv  RadarAs112SummaryGetQueryTypeParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetQueryTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetQueryTypeParamsFormatJson, RadarAs112SummaryGetQueryTypeParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112SummaryGetQueryTypeParamsProtocol string

const (
	RadarAs112SummaryGetQueryTypeParamsProtocolUdp   RadarAs112SummaryGetQueryTypeParamsProtocol = "UDP"
	RadarAs112SummaryGetQueryTypeParamsProtocolTcp   RadarAs112SummaryGetQueryTypeParamsProtocol = "TCP"
	RadarAs112SummaryGetQueryTypeParamsProtocolHTTPS RadarAs112SummaryGetQueryTypeParamsProtocol = "HTTPS"
	RadarAs112SummaryGetQueryTypeParamsProtocolTls   RadarAs112SummaryGetQueryTypeParamsProtocol = "TLS"
)

func (r RadarAs112SummaryGetQueryTypeParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetQueryTypeParamsProtocolUdp, RadarAs112SummaryGetQueryTypeParamsProtocolTcp, RadarAs112SummaryGetQueryTypeParamsProtocolHTTPS, RadarAs112SummaryGetQueryTypeParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS response code.
type RadarAs112SummaryGetQueryTypeParamsResponseCode string

const (
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNoerror   RadarAs112SummaryGetQueryTypeParamsResponseCode = "NOERROR"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeFormerr   RadarAs112SummaryGetQueryTypeParamsResponseCode = "FORMERR"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeServfail  RadarAs112SummaryGetQueryTypeParamsResponseCode = "SERVFAIL"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNxdomain  RadarAs112SummaryGetQueryTypeParamsResponseCode = "NXDOMAIN"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNotimp    RadarAs112SummaryGetQueryTypeParamsResponseCode = "NOTIMP"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeRefused   RadarAs112SummaryGetQueryTypeParamsResponseCode = "REFUSED"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeYxdomain  RadarAs112SummaryGetQueryTypeParamsResponseCode = "YXDOMAIN"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeYxrrset   RadarAs112SummaryGetQueryTypeParamsResponseCode = "YXRRSET"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNxrrset   RadarAs112SummaryGetQueryTypeParamsResponseCode = "NXRRSET"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNotauth   RadarAs112SummaryGetQueryTypeParamsResponseCode = "NOTAUTH"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeNotzone   RadarAs112SummaryGetQueryTypeParamsResponseCode = "NOTZONE"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadsig    RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADSIG"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadkey    RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADKEY"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadtime   RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADTIME"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadmode   RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADMODE"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadname   RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADNAME"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadalg    RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADALG"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadtrunc  RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADTRUNC"
	RadarAs112SummaryGetQueryTypeParamsResponseCodeBadcookie RadarAs112SummaryGetQueryTypeParamsResponseCode = "BADCOOKIE"
)

func (r RadarAs112SummaryGetQueryTypeParamsResponseCode) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetQueryTypeParamsResponseCodeNoerror, RadarAs112SummaryGetQueryTypeParamsResponseCodeFormerr, RadarAs112SummaryGetQueryTypeParamsResponseCodeServfail, RadarAs112SummaryGetQueryTypeParamsResponseCodeNxdomain, RadarAs112SummaryGetQueryTypeParamsResponseCodeNotimp, RadarAs112SummaryGetQueryTypeParamsResponseCodeRefused, RadarAs112SummaryGetQueryTypeParamsResponseCodeYxdomain, RadarAs112SummaryGetQueryTypeParamsResponseCodeYxrrset, RadarAs112SummaryGetQueryTypeParamsResponseCodeNxrrset, RadarAs112SummaryGetQueryTypeParamsResponseCodeNotauth, RadarAs112SummaryGetQueryTypeParamsResponseCodeNotzone, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadsig, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadkey, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadtime, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadmode, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadname, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadalg, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadtrunc, RadarAs112SummaryGetQueryTypeParamsResponseCodeBadcookie:
		return true
	}
	return false
}

type RadarAs112SummaryGetResponseCodesParams struct {
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
	Format param.Field[RadarAs112SummaryGetResponseCodesParamsFormat] `query:"format"`
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
	// Filters results by DNS transport protocol.
	Protocol param.Field[RadarAs112SummaryGetResponseCodesParamsProtocol] `query:"protocol"`
	// Filters results by DNS query type.
	QueryType param.Field[RadarAs112SummaryGetResponseCodesParamsQueryType] `query:"queryType"`
}

// URLQuery serializes [RadarAs112SummaryGetResponseCodesParams]'s query parameters
// as `url.Values`.
func (r RadarAs112SummaryGetResponseCodesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112SummaryGetResponseCodesParamsFormat string

const (
	RadarAs112SummaryGetResponseCodesParamsFormatJson RadarAs112SummaryGetResponseCodesParamsFormat = "JSON"
	RadarAs112SummaryGetResponseCodesParamsFormatCsv  RadarAs112SummaryGetResponseCodesParamsFormat = "CSV"
)

func (r RadarAs112SummaryGetResponseCodesParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetResponseCodesParamsFormatJson, RadarAs112SummaryGetResponseCodesParamsFormatCsv:
		return true
	}
	return false
}

// Filters results by DNS transport protocol.
type RadarAs112SummaryGetResponseCodesParamsProtocol string

const (
	RadarAs112SummaryGetResponseCodesParamsProtocolUdp   RadarAs112SummaryGetResponseCodesParamsProtocol = "UDP"
	RadarAs112SummaryGetResponseCodesParamsProtocolTcp   RadarAs112SummaryGetResponseCodesParamsProtocol = "TCP"
	RadarAs112SummaryGetResponseCodesParamsProtocolHTTPS RadarAs112SummaryGetResponseCodesParamsProtocol = "HTTPS"
	RadarAs112SummaryGetResponseCodesParamsProtocolTls   RadarAs112SummaryGetResponseCodesParamsProtocol = "TLS"
)

func (r RadarAs112SummaryGetResponseCodesParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetResponseCodesParamsProtocolUdp, RadarAs112SummaryGetResponseCodesParamsProtocolTcp, RadarAs112SummaryGetResponseCodesParamsProtocolHTTPS, RadarAs112SummaryGetResponseCodesParamsProtocolTls:
		return true
	}
	return false
}

// Filters results by DNS query type.
type RadarAs112SummaryGetResponseCodesParamsQueryType string

const (
	RadarAs112SummaryGetResponseCodesParamsQueryTypeA          RadarAs112SummaryGetResponseCodesParamsQueryType = "A"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeAaaa       RadarAs112SummaryGetResponseCodesParamsQueryType = "AAAA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeA6         RadarAs112SummaryGetResponseCodesParamsQueryType = "A6"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeAfsdb      RadarAs112SummaryGetResponseCodesParamsQueryType = "AFSDB"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeAny        RadarAs112SummaryGetResponseCodesParamsQueryType = "ANY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeApl        RadarAs112SummaryGetResponseCodesParamsQueryType = "APL"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeAtma       RadarAs112SummaryGetResponseCodesParamsQueryType = "ATMA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeAxfr       RadarAs112SummaryGetResponseCodesParamsQueryType = "AXFR"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCaa        RadarAs112SummaryGetResponseCodesParamsQueryType = "CAA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCdnskey    RadarAs112SummaryGetResponseCodesParamsQueryType = "CDNSKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCds        RadarAs112SummaryGetResponseCodesParamsQueryType = "CDS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCert       RadarAs112SummaryGetResponseCodesParamsQueryType = "CERT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCname      RadarAs112SummaryGetResponseCodesParamsQueryType = "CNAME"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeCsync      RadarAs112SummaryGetResponseCodesParamsQueryType = "CSYNC"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDhcid      RadarAs112SummaryGetResponseCodesParamsQueryType = "DHCID"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDlv        RadarAs112SummaryGetResponseCodesParamsQueryType = "DLV"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDname      RadarAs112SummaryGetResponseCodesParamsQueryType = "DNAME"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDnskey     RadarAs112SummaryGetResponseCodesParamsQueryType = "DNSKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDoa        RadarAs112SummaryGetResponseCodesParamsQueryType = "DOA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeDs         RadarAs112SummaryGetResponseCodesParamsQueryType = "DS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeEid        RadarAs112SummaryGetResponseCodesParamsQueryType = "EID"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeEui48      RadarAs112SummaryGetResponseCodesParamsQueryType = "EUI48"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeEui64      RadarAs112SummaryGetResponseCodesParamsQueryType = "EUI64"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeGpos       RadarAs112SummaryGetResponseCodesParamsQueryType = "GPOS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeGid        RadarAs112SummaryGetResponseCodesParamsQueryType = "GID"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeHinfo      RadarAs112SummaryGetResponseCodesParamsQueryType = "HINFO"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeHip        RadarAs112SummaryGetResponseCodesParamsQueryType = "HIP"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeHTTPS      RadarAs112SummaryGetResponseCodesParamsQueryType = "HTTPS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeIpseckey   RadarAs112SummaryGetResponseCodesParamsQueryType = "IPSECKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeIsdn       RadarAs112SummaryGetResponseCodesParamsQueryType = "ISDN"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeIxfr       RadarAs112SummaryGetResponseCodesParamsQueryType = "IXFR"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeKey        RadarAs112SummaryGetResponseCodesParamsQueryType = "KEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeKx         RadarAs112SummaryGetResponseCodesParamsQueryType = "KX"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeL32        RadarAs112SummaryGetResponseCodesParamsQueryType = "L32"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeL64        RadarAs112SummaryGetResponseCodesParamsQueryType = "L64"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeLoc        RadarAs112SummaryGetResponseCodesParamsQueryType = "LOC"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeLp         RadarAs112SummaryGetResponseCodesParamsQueryType = "LP"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMaila      RadarAs112SummaryGetResponseCodesParamsQueryType = "MAILA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMailb      RadarAs112SummaryGetResponseCodesParamsQueryType = "MAILB"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMB         RadarAs112SummaryGetResponseCodesParamsQueryType = "MB"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMd         RadarAs112SummaryGetResponseCodesParamsQueryType = "MD"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMf         RadarAs112SummaryGetResponseCodesParamsQueryType = "MF"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMg         RadarAs112SummaryGetResponseCodesParamsQueryType = "MG"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMinfo      RadarAs112SummaryGetResponseCodesParamsQueryType = "MINFO"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMr         RadarAs112SummaryGetResponseCodesParamsQueryType = "MR"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeMx         RadarAs112SummaryGetResponseCodesParamsQueryType = "MX"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNaptr      RadarAs112SummaryGetResponseCodesParamsQueryType = "NAPTR"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNb         RadarAs112SummaryGetResponseCodesParamsQueryType = "NB"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNbstat     RadarAs112SummaryGetResponseCodesParamsQueryType = "NBSTAT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNid        RadarAs112SummaryGetResponseCodesParamsQueryType = "NID"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNimloc     RadarAs112SummaryGetResponseCodesParamsQueryType = "NIMLOC"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNinfo      RadarAs112SummaryGetResponseCodesParamsQueryType = "NINFO"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNs         RadarAs112SummaryGetResponseCodesParamsQueryType = "NS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNsap       RadarAs112SummaryGetResponseCodesParamsQueryType = "NSAP"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec       RadarAs112SummaryGetResponseCodesParamsQueryType = "NSEC"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec3      RadarAs112SummaryGetResponseCodesParamsQueryType = "NSEC3"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec3Param RadarAs112SummaryGetResponseCodesParamsQueryType = "NSEC3PARAM"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNull       RadarAs112SummaryGetResponseCodesParamsQueryType = "NULL"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeNxt        RadarAs112SummaryGetResponseCodesParamsQueryType = "NXT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeOpenpgpkey RadarAs112SummaryGetResponseCodesParamsQueryType = "OPENPGPKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeOpt        RadarAs112SummaryGetResponseCodesParamsQueryType = "OPT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypePtr        RadarAs112SummaryGetResponseCodesParamsQueryType = "PTR"
	RadarAs112SummaryGetResponseCodesParamsQueryTypePx         RadarAs112SummaryGetResponseCodesParamsQueryType = "PX"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeRkey       RadarAs112SummaryGetResponseCodesParamsQueryType = "RKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeRp         RadarAs112SummaryGetResponseCodesParamsQueryType = "RP"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeRrsig      RadarAs112SummaryGetResponseCodesParamsQueryType = "RRSIG"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeRt         RadarAs112SummaryGetResponseCodesParamsQueryType = "RT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSig        RadarAs112SummaryGetResponseCodesParamsQueryType = "SIG"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSink       RadarAs112SummaryGetResponseCodesParamsQueryType = "SINK"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSmimea     RadarAs112SummaryGetResponseCodesParamsQueryType = "SMIMEA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSoa        RadarAs112SummaryGetResponseCodesParamsQueryType = "SOA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSpf        RadarAs112SummaryGetResponseCodesParamsQueryType = "SPF"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSrv        RadarAs112SummaryGetResponseCodesParamsQueryType = "SRV"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSshfp      RadarAs112SummaryGetResponseCodesParamsQueryType = "SSHFP"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeSvcb       RadarAs112SummaryGetResponseCodesParamsQueryType = "SVCB"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTa         RadarAs112SummaryGetResponseCodesParamsQueryType = "TA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTalink     RadarAs112SummaryGetResponseCodesParamsQueryType = "TALINK"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTkey       RadarAs112SummaryGetResponseCodesParamsQueryType = "TKEY"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTlsa       RadarAs112SummaryGetResponseCodesParamsQueryType = "TLSA"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTsig       RadarAs112SummaryGetResponseCodesParamsQueryType = "TSIG"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeTxt        RadarAs112SummaryGetResponseCodesParamsQueryType = "TXT"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeUinfo      RadarAs112SummaryGetResponseCodesParamsQueryType = "UINFO"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeUid        RadarAs112SummaryGetResponseCodesParamsQueryType = "UID"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeUnspec     RadarAs112SummaryGetResponseCodesParamsQueryType = "UNSPEC"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeUri        RadarAs112SummaryGetResponseCodesParamsQueryType = "URI"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeWks        RadarAs112SummaryGetResponseCodesParamsQueryType = "WKS"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeX25        RadarAs112SummaryGetResponseCodesParamsQueryType = "X25"
	RadarAs112SummaryGetResponseCodesParamsQueryTypeZonemd     RadarAs112SummaryGetResponseCodesParamsQueryType = "ZONEMD"
)

func (r RadarAs112SummaryGetResponseCodesParamsQueryType) IsKnown() bool {
	switch r {
	case RadarAs112SummaryGetResponseCodesParamsQueryTypeA, RadarAs112SummaryGetResponseCodesParamsQueryTypeAaaa, RadarAs112SummaryGetResponseCodesParamsQueryTypeA6, RadarAs112SummaryGetResponseCodesParamsQueryTypeAfsdb, RadarAs112SummaryGetResponseCodesParamsQueryTypeAny, RadarAs112SummaryGetResponseCodesParamsQueryTypeApl, RadarAs112SummaryGetResponseCodesParamsQueryTypeAtma, RadarAs112SummaryGetResponseCodesParamsQueryTypeAxfr, RadarAs112SummaryGetResponseCodesParamsQueryTypeCaa, RadarAs112SummaryGetResponseCodesParamsQueryTypeCdnskey, RadarAs112SummaryGetResponseCodesParamsQueryTypeCds, RadarAs112SummaryGetResponseCodesParamsQueryTypeCert, RadarAs112SummaryGetResponseCodesParamsQueryTypeCname, RadarAs112SummaryGetResponseCodesParamsQueryTypeCsync, RadarAs112SummaryGetResponseCodesParamsQueryTypeDhcid, RadarAs112SummaryGetResponseCodesParamsQueryTypeDlv, RadarAs112SummaryGetResponseCodesParamsQueryTypeDname, RadarAs112SummaryGetResponseCodesParamsQueryTypeDnskey, RadarAs112SummaryGetResponseCodesParamsQueryTypeDoa, RadarAs112SummaryGetResponseCodesParamsQueryTypeDs, RadarAs112SummaryGetResponseCodesParamsQueryTypeEid, RadarAs112SummaryGetResponseCodesParamsQueryTypeEui48, RadarAs112SummaryGetResponseCodesParamsQueryTypeEui64, RadarAs112SummaryGetResponseCodesParamsQueryTypeGpos, RadarAs112SummaryGetResponseCodesParamsQueryTypeGid, RadarAs112SummaryGetResponseCodesParamsQueryTypeHinfo, RadarAs112SummaryGetResponseCodesParamsQueryTypeHip, RadarAs112SummaryGetResponseCodesParamsQueryTypeHTTPS, RadarAs112SummaryGetResponseCodesParamsQueryTypeIpseckey, RadarAs112SummaryGetResponseCodesParamsQueryTypeIsdn, RadarAs112SummaryGetResponseCodesParamsQueryTypeIxfr, RadarAs112SummaryGetResponseCodesParamsQueryTypeKey, RadarAs112SummaryGetResponseCodesParamsQueryTypeKx, RadarAs112SummaryGetResponseCodesParamsQueryTypeL32, RadarAs112SummaryGetResponseCodesParamsQueryTypeL64, RadarAs112SummaryGetResponseCodesParamsQueryTypeLoc, RadarAs112SummaryGetResponseCodesParamsQueryTypeLp, RadarAs112SummaryGetResponseCodesParamsQueryTypeMaila, RadarAs112SummaryGetResponseCodesParamsQueryTypeMailb, RadarAs112SummaryGetResponseCodesParamsQueryTypeMB, RadarAs112SummaryGetResponseCodesParamsQueryTypeMd, RadarAs112SummaryGetResponseCodesParamsQueryTypeMf, RadarAs112SummaryGetResponseCodesParamsQueryTypeMg, RadarAs112SummaryGetResponseCodesParamsQueryTypeMinfo, RadarAs112SummaryGetResponseCodesParamsQueryTypeMr, RadarAs112SummaryGetResponseCodesParamsQueryTypeMx, RadarAs112SummaryGetResponseCodesParamsQueryTypeNaptr, RadarAs112SummaryGetResponseCodesParamsQueryTypeNb, RadarAs112SummaryGetResponseCodesParamsQueryTypeNbstat, RadarAs112SummaryGetResponseCodesParamsQueryTypeNid, RadarAs112SummaryGetResponseCodesParamsQueryTypeNimloc, RadarAs112SummaryGetResponseCodesParamsQueryTypeNinfo, RadarAs112SummaryGetResponseCodesParamsQueryTypeNs, RadarAs112SummaryGetResponseCodesParamsQueryTypeNsap, RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec, RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec3, RadarAs112SummaryGetResponseCodesParamsQueryTypeNsec3Param, RadarAs112SummaryGetResponseCodesParamsQueryTypeNull, RadarAs112SummaryGetResponseCodesParamsQueryTypeNxt, RadarAs112SummaryGetResponseCodesParamsQueryTypeOpenpgpkey, RadarAs112SummaryGetResponseCodesParamsQueryTypeOpt, RadarAs112SummaryGetResponseCodesParamsQueryTypePtr, RadarAs112SummaryGetResponseCodesParamsQueryTypePx, RadarAs112SummaryGetResponseCodesParamsQueryTypeRkey, RadarAs112SummaryGetResponseCodesParamsQueryTypeRp, RadarAs112SummaryGetResponseCodesParamsQueryTypeRrsig, RadarAs112SummaryGetResponseCodesParamsQueryTypeRt, RadarAs112SummaryGetResponseCodesParamsQueryTypeSig, RadarAs112SummaryGetResponseCodesParamsQueryTypeSink, RadarAs112SummaryGetResponseCodesParamsQueryTypeSmimea, RadarAs112SummaryGetResponseCodesParamsQueryTypeSoa, RadarAs112SummaryGetResponseCodesParamsQueryTypeSpf, RadarAs112SummaryGetResponseCodesParamsQueryTypeSrv, RadarAs112SummaryGetResponseCodesParamsQueryTypeSshfp, RadarAs112SummaryGetResponseCodesParamsQueryTypeSvcb, RadarAs112SummaryGetResponseCodesParamsQueryTypeTa, RadarAs112SummaryGetResponseCodesParamsQueryTypeTalink, RadarAs112SummaryGetResponseCodesParamsQueryTypeTkey, RadarAs112SummaryGetResponseCodesParamsQueryTypeTlsa, RadarAs112SummaryGetResponseCodesParamsQueryTypeTsig, RadarAs112SummaryGetResponseCodesParamsQueryTypeTxt, RadarAs112SummaryGetResponseCodesParamsQueryTypeUinfo, RadarAs112SummaryGetResponseCodesParamsQueryTypeUid, RadarAs112SummaryGetResponseCodesParamsQueryTypeUnspec, RadarAs112SummaryGetResponseCodesParamsQueryTypeUri, RadarAs112SummaryGetResponseCodesParamsQueryTypeWks, RadarAs112SummaryGetResponseCodesParamsQueryTypeX25, RadarAs112SummaryGetResponseCodesParamsQueryTypeZonemd:
		return true
	}
	return false
}
