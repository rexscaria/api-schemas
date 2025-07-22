// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// RadarAs112TopLocationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAs112TopLocationService] method instead.
type RadarAs112TopLocationService struct {
	Options []option.RequestOption
}

// NewRadarAs112TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112TopLocationService(opts ...option.RequestOption) (r *RadarAs112TopLocationService) {
	r = &RadarAs112TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the top locations by AS112 DNS queries.
func (r *RadarAs112TopLocationService) Get(ctx context.Context, query RadarAs112TopLocationGetParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 with DNSSEC (DNS Security
// Extensions) support.
func (r *RadarAs112TopLocationService) GetByDnssec(ctx context.Context, dnssec RadarAs112TopLocationGetByDnssecParamsDnssec, query RadarAs112TopLocationGetByDnssecParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByDnssecResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/dnssec/%v", dnssec)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 with EDNS (Extension
// Mechanisms for DNS) support.
func (r *RadarAs112TopLocationService) GetByEdns(ctx context.Context, edns RadarAs112TopLocationGetByEdnsParamsEdns, query RadarAs112TopLocationGetByEdnsParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByEdnsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/edns/%v", edns)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations of DNS queries to AS112 for an IP version.
func (r *RadarAs112TopLocationService) GetByIPVersion(ctx context.Context, ipVersion RadarAs112TopLocationGetByIPVersionParamsIPVersion, query RadarAs112TopLocationGetByIPVersionParams, opts ...option.RequestOption) (res *RadarAs112TopLocationGetByIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/as112/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TopLocationGetResponse struct {
	Result  RadarAs112TopLocationGetResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarAs112TopLocationGetResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetResponseJSON contains the JSON metadata for the struct
// [RadarAs112TopLocationGetResponse]
type radarAs112TopLocationGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResult struct {
	Meta RadarAs112TopLocationGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetResponseResult]
type radarAs112TopLocationGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMeta struct {
	DateRange      []RadarAs112TopLocationGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopLocationGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopLocationGetResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetResponseResultMeta]
type radarAs112TopLocationGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetResponseResultMetaDateRange]
type radarAs112TopLocationGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationGetResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarAs112TopLocationGetResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetResponseResultTop0]
type radarAs112TopLocationGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponse struct {
	Result  RadarAs112TopLocationGetByDnssecResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarAs112TopLocationGetByDnssecResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetByDnssecResponse]
type radarAs112TopLocationGetByDnssecResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResult struct {
	Meta RadarAs112TopLocationGetByDnssecResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByDnssecResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByDnssecResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByDnssecResponseResult]
type radarAs112TopLocationGetByDnssecResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMeta struct {
	DateRange      []RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopLocationGetByDnssecResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByDnssecResponseResultMeta]
type radarAs112TopLocationGetByDnssecResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange]
type radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByDnssecResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                 `json:"clientCountryName,required"`
	Value               string                                                 `json:"value,required"`
	JSON                radarAs112TopLocationGetByDnssecResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByDnssecResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByDnssecResponseResultTop0]
type radarAs112TopLocationGetByDnssecResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByDnssecResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByDnssecResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponse struct {
	Result  RadarAs112TopLocationGetByEdnsResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112TopLocationGetByEdnsResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseJSON contains the JSON metadata for the
// struct [RadarAs112TopLocationGetByEdnsResponse]
type radarAs112TopLocationGetByEdnsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResult struct {
	Meta RadarAs112TopLocationGetByEdnsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByEdnsResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByEdnsResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetByEdnsResponseResult]
type radarAs112TopLocationGetByEdnsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMeta struct {
	DateRange      []RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopLocationGetByEdnsResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByEdnsResponseResultMeta]
type radarAs112TopLocationGetByEdnsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange]
type radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByEdnsResponseResultTop0 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarAs112TopLocationGetByEdnsResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByEdnsResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByEdnsResponseResultTop0]
type radarAs112TopLocationGetByEdnsResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByEdnsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByEdnsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponse struct {
	Result  RadarAs112TopLocationGetByIPVersionResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAs112TopLocationGetByIPVersionResponseJSON   `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarAs112TopLocationGetByIPVersionResponse]
type radarAs112TopLocationGetByIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResult struct {
	Meta RadarAs112TopLocationGetByIPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAs112TopLocationGetByIPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarAs112TopLocationGetByIPVersionResponseResultJSON   `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TopLocationGetByIPVersionResponseResult]
type radarAs112TopLocationGetByIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMeta struct {
	DateRange      []RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TopLocationGetByIPVersionResponseResultMetaJSON           `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByIPVersionResponseResultMeta]
type radarAs112TopLocationGetByIPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange]
type radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo]
type radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetByIPVersionResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                    `json:"clientCountryName,required"`
	Value               string                                                    `json:"value,required"`
	JSON                radarAs112TopLocationGetByIPVersionResponseResultTop0JSON `json:"-"`
}

// radarAs112TopLocationGetByIPVersionResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAs112TopLocationGetByIPVersionResponseResultTop0]
type radarAs112TopLocationGetByIPVersionResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAs112TopLocationGetByIPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAs112TopLocationGetByIPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAs112TopLocationGetParams struct {
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
	Format param.Field[RadarAs112TopLocationGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAs112TopLocationGetParamsFormat string

const (
	RadarAs112TopLocationGetParamsFormatJson RadarAs112TopLocationGetParamsFormat = "JSON"
	RadarAs112TopLocationGetParamsFormatCsv  RadarAs112TopLocationGetParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetParamsFormatJson, RadarAs112TopLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByDnssecParams struct {
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
	Format param.Field[RadarAs112TopLocationGetByDnssecParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByDnssecParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TopLocationGetByDnssecParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DNSSEC (DNS Security Extensions) status.
type RadarAs112TopLocationGetByDnssecParamsDnssec string

const (
	RadarAs112TopLocationGetByDnssecParamsDnssecSupported    RadarAs112TopLocationGetByDnssecParamsDnssec = "SUPPORTED"
	RadarAs112TopLocationGetByDnssecParamsDnssecNotSupported RadarAs112TopLocationGetByDnssecParamsDnssec = "NOT_SUPPORTED"
)

func (r RadarAs112TopLocationGetByDnssecParamsDnssec) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByDnssecParamsDnssecSupported, RadarAs112TopLocationGetByDnssecParamsDnssecNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByDnssecParamsFormat string

const (
	RadarAs112TopLocationGetByDnssecParamsFormatJson RadarAs112TopLocationGetByDnssecParamsFormat = "JSON"
	RadarAs112TopLocationGetByDnssecParamsFormatCsv  RadarAs112TopLocationGetByDnssecParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByDnssecParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByDnssecParamsFormatJson, RadarAs112TopLocationGetByDnssecParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByEdnsParams struct {
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
	Format param.Field[RadarAs112TopLocationGetByEdnsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByEdnsParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TopLocationGetByEdnsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// EDNS (Extension Mechanisms for DNS) status.
type RadarAs112TopLocationGetByEdnsParamsEdns string

const (
	RadarAs112TopLocationGetByEdnsParamsEdnsSupported    RadarAs112TopLocationGetByEdnsParamsEdns = "SUPPORTED"
	RadarAs112TopLocationGetByEdnsParamsEdnsNotSupported RadarAs112TopLocationGetByEdnsParamsEdns = "NOT_SUPPORTED"
)

func (r RadarAs112TopLocationGetByEdnsParamsEdns) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByEdnsParamsEdnsSupported, RadarAs112TopLocationGetByEdnsParamsEdnsNotSupported:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByEdnsParamsFormat string

const (
	RadarAs112TopLocationGetByEdnsParamsFormatJson RadarAs112TopLocationGetByEdnsParamsFormat = "JSON"
	RadarAs112TopLocationGetByEdnsParamsFormatCsv  RadarAs112TopLocationGetByEdnsParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByEdnsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByEdnsParamsFormatJson, RadarAs112TopLocationGetByEdnsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAs112TopLocationGetByIPVersionParams struct {
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
	Format param.Field[RadarAs112TopLocationGetByIPVersionParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TopLocationGetByIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TopLocationGetByIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarAs112TopLocationGetByIPVersionParamsIPVersion string

const (
	RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv4 RadarAs112TopLocationGetByIPVersionParamsIPVersion = "IPv4"
	RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv6 RadarAs112TopLocationGetByIPVersionParamsIPVersion = "IPv6"
)

func (r RadarAs112TopLocationGetByIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv4, RadarAs112TopLocationGetByIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAs112TopLocationGetByIPVersionParamsFormat string

const (
	RadarAs112TopLocationGetByIPVersionParamsFormatJson RadarAs112TopLocationGetByIPVersionParamsFormat = "JSON"
	RadarAs112TopLocationGetByIPVersionParamsFormatCsv  RadarAs112TopLocationGetByIPVersionParamsFormat = "CSV"
)

func (r RadarAs112TopLocationGetByIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarAs112TopLocationGetByIPVersionParamsFormatJson, RadarAs112TopLocationGetByIPVersionParamsFormatCsv:
		return true
	}
	return false
}
