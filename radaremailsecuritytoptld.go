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

// RadarEmailSecurityTopTldService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTopTldService] method instead.
type RadarEmailSecurityTopTldService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopTldService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopTldService(opts ...option.RequestOption) (r *RadarEmailSecurityTopTldService) {
	r = &RadarEmailSecurityTopTldService{}
	r.Options = opts
	return
}

// Retrieves the top TLDs by number of email messages.
func (r *RadarEmailSecurityTopTldService) Get(ctx context.Context, query RadarEmailSecurityTopTldGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/top/tlds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top TLDs by emails classified as malicious or not.
func (r *RadarEmailSecurityTopTldService) GetMalicious(ctx context.Context, malicious RadarEmailSecurityTopTldGetMaliciousParamsMalicious, query RadarEmailSecurityTopTldGetMaliciousParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldGetMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top TLDs by emails classified as spam or not.
func (r *RadarEmailSecurityTopTldService) GetSpam(ctx context.Context, spam RadarEmailSecurityTopTldGetSpamParamsSpam, query RadarEmailSecurityTopTldGetSpamParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldGetSpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top TLDs by emails classified as spoof or not.
func (r *RadarEmailSecurityTopTldService) GetSpoof(ctx context.Context, spoof RadarEmailSecurityTopTldGetSpoofParamsSpoof, query RadarEmailSecurityTopTldGetSpoofParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopTldGetSpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/tlds/spoof/%v", spoof)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopTldGetResponse struct {
	Result  RadarEmailSecurityTopTldGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarEmailSecurityTopTldGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetResponse]
type radarEmailSecurityTopTldGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResult struct {
	Meta RadarEmailSecurityTopTldGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldGetResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetResponseResult]
type radarEmailSecurityTopTldGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopTldGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldGetResponseResultMeta]
type radarEmailSecurityTopTldGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetResponseResultMetaDateRange]
type radarEmailSecurityTopTldGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetResponseResultTop0 struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  radarEmailSecurityTopTldGetResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldGetResponseResultTop0]
type radarEmailSecurityTopTldGetResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponse struct {
	Result  RadarEmailSecurityTopTldGetMaliciousResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecurityTopTldGetMaliciousResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldGetMaliciousResponse]
type radarEmailSecurityTopTldGetMaliciousResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResult struct {
	Meta RadarEmailSecurityTopTldGetMaliciousResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldGetMaliciousResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldGetMaliciousResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetMaliciousResponseResult]
type radarEmailSecurityTopTldGetMaliciousResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldGetMaliciousResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetMaliciousResponseResultMeta]
type radarEmailSecurityTopTldGetMaliciousResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRange]
type radarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetMaliciousResponseResultTop0 struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  radarEmailSecurityTopTldGetMaliciousResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldGetMaliciousResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetMaliciousResponseResultTop0]
type radarEmailSecurityTopTldGetMaliciousResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetMaliciousResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetMaliciousResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponse struct {
	Result  RadarEmailSecurityTopTldGetSpamResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailSecurityTopTldGetSpamResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetSpamResponse]
type radarEmailSecurityTopTldGetSpamResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResult struct {
	Meta RadarEmailSecurityTopTldGetSpamResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldGetSpamResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldGetSpamResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopTldGetSpamResponseResult]
type radarEmailSecurityTopTldGetSpamResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopTldGetSpamResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldGetSpamResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldGetSpamResponseResultMeta]
type radarEmailSecurityTopTldGetSpamResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldGetSpamResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopTldGetSpamResponseResultMetaDateRange]
type radarEmailSecurityTopTldGetSpamResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpamResponseResultTop0 struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  radarEmailSecurityTopTldGetSpamResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpamResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldGetSpamResponseResultTop0]
type radarEmailSecurityTopTldGetSpamResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpamResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpamResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponse struct {
	Result  RadarEmailSecurityTopTldGetSpoofResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailSecurityTopTldGetSpoofResponseJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopTldGetSpoofResponse]
type radarEmailSecurityTopTldGetSpoofResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResult struct {
	Meta RadarEmailSecurityTopTldGetSpoofResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopTldGetSpoofResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopTldGetSpoofResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopTldGetSpoofResponseResult]
type radarEmailSecurityTopTldGetSpoofResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopTldGetSpoofResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopTldGetSpoofResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetSpoofResponseResultMeta]
type radarEmailSecurityTopTldGetSpoofResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopTldGetSpoofResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopTldGetSpoofResponseResultMetaDateRange]
type radarEmailSecurityTopTldGetSpoofResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetSpoofResponseResultTop0 struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  radarEmailSecurityTopTldGetSpoofResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopTldGetSpoofResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopTldGetSpoofResponseResultTop0]
type radarEmailSecurityTopTldGetSpoofResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopTldGetSpoofResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTopTldGetSpoofResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTopTldGetParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTopTldGetParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTopTldGetParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTopTldGetParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTopTldGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTopTldGetParamsSpf] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[RadarEmailSecurityTopTldGetParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTopTldGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopTldGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopTldGetParamsArc string

const (
	RadarEmailSecurityTopTldGetParamsArcPass RadarEmailSecurityTopTldGetParamsArc = "PASS"
	RadarEmailSecurityTopTldGetParamsArcNone RadarEmailSecurityTopTldGetParamsArc = "NONE"
	RadarEmailSecurityTopTldGetParamsArcFail RadarEmailSecurityTopTldGetParamsArc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsArcPass, RadarEmailSecurityTopTldGetParamsArcNone, RadarEmailSecurityTopTldGetParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetParamsDkim string

const (
	RadarEmailSecurityTopTldGetParamsDkimPass RadarEmailSecurityTopTldGetParamsDkim = "PASS"
	RadarEmailSecurityTopTldGetParamsDkimNone RadarEmailSecurityTopTldGetParamsDkim = "NONE"
	RadarEmailSecurityTopTldGetParamsDkimFail RadarEmailSecurityTopTldGetParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTopTldGetParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsDkimPass, RadarEmailSecurityTopTldGetParamsDkimNone, RadarEmailSecurityTopTldGetParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetParamsDmarc string

const (
	RadarEmailSecurityTopTldGetParamsDmarcPass RadarEmailSecurityTopTldGetParamsDmarc = "PASS"
	RadarEmailSecurityTopTldGetParamsDmarcNone RadarEmailSecurityTopTldGetParamsDmarc = "NONE"
	RadarEmailSecurityTopTldGetParamsDmarcFail RadarEmailSecurityTopTldGetParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsDmarcPass, RadarEmailSecurityTopTldGetParamsDmarcNone, RadarEmailSecurityTopTldGetParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTopTldGetParamsFormat string

const (
	RadarEmailSecurityTopTldGetParamsFormatJson RadarEmailSecurityTopTldGetParamsFormat = "JSON"
	RadarEmailSecurityTopTldGetParamsFormatCsv  RadarEmailSecurityTopTldGetParamsFormat = "CSV"
)

func (r RadarEmailSecurityTopTldGetParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsFormatJson, RadarEmailSecurityTopTldGetParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetParamsSpf string

const (
	RadarEmailSecurityTopTldGetParamsSpfPass RadarEmailSecurityTopTldGetParamsSpf = "PASS"
	RadarEmailSecurityTopTldGetParamsSpfNone RadarEmailSecurityTopTldGetParamsSpf = "NONE"
	RadarEmailSecurityTopTldGetParamsSpfFail RadarEmailSecurityTopTldGetParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTopTldGetParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsSpfPass, RadarEmailSecurityTopTldGetParamsSpfNone, RadarEmailSecurityTopTldGetParamsSpfFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type RadarEmailSecurityTopTldGetParamsTldCategory string

const (
	RadarEmailSecurityTopTldGetParamsTldCategoryClassic RadarEmailSecurityTopTldGetParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldGetParamsTldCategoryCountry RadarEmailSecurityTopTldGetParamsTldCategory = "COUNTRY"
)

func (r RadarEmailSecurityTopTldGetParamsTldCategory) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsTldCategoryClassic, RadarEmailSecurityTopTldGetParamsTldCategoryCountry:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetParamsTlsVersion string

const (
	RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_0 RadarEmailSecurityTopTldGetParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_1 RadarEmailSecurityTopTldGetParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_2 RadarEmailSecurityTopTldGetParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_3 RadarEmailSecurityTopTldGetParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTopTldGetParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_0, RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_1, RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_2, RadarEmailSecurityTopTldGetParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTopTldGetMaliciousParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTopTldGetMaliciousParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTopTldGetMaliciousParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTopTldGetMaliciousParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTopTldGetMaliciousParamsSpf] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[RadarEmailSecurityTopTldGetMaliciousParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldGetMaliciousParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopTldGetMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious classification.
type RadarEmailSecurityTopTldGetMaliciousParamsMalicious string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsMaliciousMalicious    RadarEmailSecurityTopTldGetMaliciousParamsMalicious = "MALICIOUS"
	RadarEmailSecurityTopTldGetMaliciousParamsMaliciousNotMalicious RadarEmailSecurityTopTldGetMaliciousParamsMalicious = "NOT_MALICIOUS"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsMalicious) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsMaliciousMalicious, RadarEmailSecurityTopTldGetMaliciousParamsMaliciousNotMalicious:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParamsArc string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsArcPass RadarEmailSecurityTopTldGetMaliciousParamsArc = "PASS"
	RadarEmailSecurityTopTldGetMaliciousParamsArcNone RadarEmailSecurityTopTldGetMaliciousParamsArc = "NONE"
	RadarEmailSecurityTopTldGetMaliciousParamsArcFail RadarEmailSecurityTopTldGetMaliciousParamsArc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsArcPass, RadarEmailSecurityTopTldGetMaliciousParamsArcNone, RadarEmailSecurityTopTldGetMaliciousParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParamsDkim string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsDkimPass RadarEmailSecurityTopTldGetMaliciousParamsDkim = "PASS"
	RadarEmailSecurityTopTldGetMaliciousParamsDkimNone RadarEmailSecurityTopTldGetMaliciousParamsDkim = "NONE"
	RadarEmailSecurityTopTldGetMaliciousParamsDkimFail RadarEmailSecurityTopTldGetMaliciousParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsDkimPass, RadarEmailSecurityTopTldGetMaliciousParamsDkimNone, RadarEmailSecurityTopTldGetMaliciousParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParamsDmarc string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsDmarcPass RadarEmailSecurityTopTldGetMaliciousParamsDmarc = "PASS"
	RadarEmailSecurityTopTldGetMaliciousParamsDmarcNone RadarEmailSecurityTopTldGetMaliciousParamsDmarc = "NONE"
	RadarEmailSecurityTopTldGetMaliciousParamsDmarcFail RadarEmailSecurityTopTldGetMaliciousParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsDmarcPass, RadarEmailSecurityTopTldGetMaliciousParamsDmarcNone, RadarEmailSecurityTopTldGetMaliciousParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTopTldGetMaliciousParamsFormat string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsFormatJson RadarEmailSecurityTopTldGetMaliciousParamsFormat = "JSON"
	RadarEmailSecurityTopTldGetMaliciousParamsFormatCsv  RadarEmailSecurityTopTldGetMaliciousParamsFormat = "CSV"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsFormatJson, RadarEmailSecurityTopTldGetMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParamsSpf string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsSpfPass RadarEmailSecurityTopTldGetMaliciousParamsSpf = "PASS"
	RadarEmailSecurityTopTldGetMaliciousParamsSpfNone RadarEmailSecurityTopTldGetMaliciousParamsSpf = "NONE"
	RadarEmailSecurityTopTldGetMaliciousParamsSpfFail RadarEmailSecurityTopTldGetMaliciousParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsSpfPass, RadarEmailSecurityTopTldGetMaliciousParamsSpfNone, RadarEmailSecurityTopTldGetMaliciousParamsSpfFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type RadarEmailSecurityTopTldGetMaliciousParamsTldCategory string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsTldCategoryClassic RadarEmailSecurityTopTldGetMaliciousParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldGetMaliciousParamsTldCategoryCountry RadarEmailSecurityTopTldGetMaliciousParamsTldCategory = "COUNTRY"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsTldCategory) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsTldCategoryClassic, RadarEmailSecurityTopTldGetMaliciousParamsTldCategoryCountry:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion string

const (
	RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_0 RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_1 RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_2 RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_3 RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTopTldGetMaliciousParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_0, RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_1, RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_2, RadarEmailSecurityTopTldGetMaliciousParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTopTldGetSpamParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTopTldGetSpamParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTopTldGetSpamParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTopTldGetSpamParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTopTldGetSpamParamsSpf] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[RadarEmailSecurityTopTldGetSpamParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTopTldGetSpamParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldGetSpamParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopTldGetSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spam classification.
type RadarEmailSecurityTopTldGetSpamParamsSpam string

const (
	RadarEmailSecurityTopTldGetSpamParamsSpamSpam    RadarEmailSecurityTopTldGetSpamParamsSpam = "SPAM"
	RadarEmailSecurityTopTldGetSpamParamsSpamNotSpam RadarEmailSecurityTopTldGetSpamParamsSpam = "NOT_SPAM"
)

func (r RadarEmailSecurityTopTldGetSpamParamsSpam) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsSpamSpam, RadarEmailSecurityTopTldGetSpamParamsSpamNotSpam:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParamsArc string

const (
	RadarEmailSecurityTopTldGetSpamParamsArcPass RadarEmailSecurityTopTldGetSpamParamsArc = "PASS"
	RadarEmailSecurityTopTldGetSpamParamsArcNone RadarEmailSecurityTopTldGetSpamParamsArc = "NONE"
	RadarEmailSecurityTopTldGetSpamParamsArcFail RadarEmailSecurityTopTldGetSpamParamsArc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpamParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsArcPass, RadarEmailSecurityTopTldGetSpamParamsArcNone, RadarEmailSecurityTopTldGetSpamParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParamsDkim string

const (
	RadarEmailSecurityTopTldGetSpamParamsDkimPass RadarEmailSecurityTopTldGetSpamParamsDkim = "PASS"
	RadarEmailSecurityTopTldGetSpamParamsDkimNone RadarEmailSecurityTopTldGetSpamParamsDkim = "NONE"
	RadarEmailSecurityTopTldGetSpamParamsDkimFail RadarEmailSecurityTopTldGetSpamParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpamParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsDkimPass, RadarEmailSecurityTopTldGetSpamParamsDkimNone, RadarEmailSecurityTopTldGetSpamParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParamsDmarc string

const (
	RadarEmailSecurityTopTldGetSpamParamsDmarcPass RadarEmailSecurityTopTldGetSpamParamsDmarc = "PASS"
	RadarEmailSecurityTopTldGetSpamParamsDmarcNone RadarEmailSecurityTopTldGetSpamParamsDmarc = "NONE"
	RadarEmailSecurityTopTldGetSpamParamsDmarcFail RadarEmailSecurityTopTldGetSpamParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpamParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsDmarcPass, RadarEmailSecurityTopTldGetSpamParamsDmarcNone, RadarEmailSecurityTopTldGetSpamParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTopTldGetSpamParamsFormat string

const (
	RadarEmailSecurityTopTldGetSpamParamsFormatJson RadarEmailSecurityTopTldGetSpamParamsFormat = "JSON"
	RadarEmailSecurityTopTldGetSpamParamsFormatCsv  RadarEmailSecurityTopTldGetSpamParamsFormat = "CSV"
)

func (r RadarEmailSecurityTopTldGetSpamParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsFormatJson, RadarEmailSecurityTopTldGetSpamParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParamsSpf string

const (
	RadarEmailSecurityTopTldGetSpamParamsSpfPass RadarEmailSecurityTopTldGetSpamParamsSpf = "PASS"
	RadarEmailSecurityTopTldGetSpamParamsSpfNone RadarEmailSecurityTopTldGetSpamParamsSpf = "NONE"
	RadarEmailSecurityTopTldGetSpamParamsSpfFail RadarEmailSecurityTopTldGetSpamParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpamParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsSpfPass, RadarEmailSecurityTopTldGetSpamParamsSpfNone, RadarEmailSecurityTopTldGetSpamParamsSpfFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type RadarEmailSecurityTopTldGetSpamParamsTldCategory string

const (
	RadarEmailSecurityTopTldGetSpamParamsTldCategoryClassic RadarEmailSecurityTopTldGetSpamParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldGetSpamParamsTldCategoryCountry RadarEmailSecurityTopTldGetSpamParamsTldCategory = "COUNTRY"
)

func (r RadarEmailSecurityTopTldGetSpamParamsTldCategory) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsTldCategoryClassic, RadarEmailSecurityTopTldGetSpamParamsTldCategoryCountry:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpamParamsTlsVersion string

const (
	RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_0 RadarEmailSecurityTopTldGetSpamParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_1 RadarEmailSecurityTopTldGetSpamParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_2 RadarEmailSecurityTopTldGetSpamParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_3 RadarEmailSecurityTopTldGetSpamParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTopTldGetSpamParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_0, RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_1, RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_2, RadarEmailSecurityTopTldGetSpamParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTopTldGetSpoofParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTopTldGetSpoofParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTopTldGetSpoofParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTopTldGetSpoofParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTopTldGetSpoofParamsSpf] `query:"spf"`
	// Filters results by TLD category.
	TldCategory param.Field[RadarEmailSecurityTopTldGetSpoofParamsTldCategory] `query:"tldCategory"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTopTldGetSpoofParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTopTldGetSpoofParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopTldGetSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spoof classification.
type RadarEmailSecurityTopTldGetSpoofParamsSpoof string

const (
	RadarEmailSecurityTopTldGetSpoofParamsSpoofSpoof    RadarEmailSecurityTopTldGetSpoofParamsSpoof = "SPOOF"
	RadarEmailSecurityTopTldGetSpoofParamsSpoofNotSpoof RadarEmailSecurityTopTldGetSpoofParamsSpoof = "NOT_SPOOF"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsSpoof) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsSpoofSpoof, RadarEmailSecurityTopTldGetSpoofParamsSpoofNotSpoof:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParamsArc string

const (
	RadarEmailSecurityTopTldGetSpoofParamsArcPass RadarEmailSecurityTopTldGetSpoofParamsArc = "PASS"
	RadarEmailSecurityTopTldGetSpoofParamsArcNone RadarEmailSecurityTopTldGetSpoofParamsArc = "NONE"
	RadarEmailSecurityTopTldGetSpoofParamsArcFail RadarEmailSecurityTopTldGetSpoofParamsArc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsArcPass, RadarEmailSecurityTopTldGetSpoofParamsArcNone, RadarEmailSecurityTopTldGetSpoofParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParamsDkim string

const (
	RadarEmailSecurityTopTldGetSpoofParamsDkimPass RadarEmailSecurityTopTldGetSpoofParamsDkim = "PASS"
	RadarEmailSecurityTopTldGetSpoofParamsDkimNone RadarEmailSecurityTopTldGetSpoofParamsDkim = "NONE"
	RadarEmailSecurityTopTldGetSpoofParamsDkimFail RadarEmailSecurityTopTldGetSpoofParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsDkimPass, RadarEmailSecurityTopTldGetSpoofParamsDkimNone, RadarEmailSecurityTopTldGetSpoofParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParamsDmarc string

const (
	RadarEmailSecurityTopTldGetSpoofParamsDmarcPass RadarEmailSecurityTopTldGetSpoofParamsDmarc = "PASS"
	RadarEmailSecurityTopTldGetSpoofParamsDmarcNone RadarEmailSecurityTopTldGetSpoofParamsDmarc = "NONE"
	RadarEmailSecurityTopTldGetSpoofParamsDmarcFail RadarEmailSecurityTopTldGetSpoofParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsDmarcPass, RadarEmailSecurityTopTldGetSpoofParamsDmarcNone, RadarEmailSecurityTopTldGetSpoofParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTopTldGetSpoofParamsFormat string

const (
	RadarEmailSecurityTopTldGetSpoofParamsFormatJson RadarEmailSecurityTopTldGetSpoofParamsFormat = "JSON"
	RadarEmailSecurityTopTldGetSpoofParamsFormatCsv  RadarEmailSecurityTopTldGetSpoofParamsFormat = "CSV"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsFormatJson, RadarEmailSecurityTopTldGetSpoofParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParamsSpf string

const (
	RadarEmailSecurityTopTldGetSpoofParamsSpfPass RadarEmailSecurityTopTldGetSpoofParamsSpf = "PASS"
	RadarEmailSecurityTopTldGetSpoofParamsSpfNone RadarEmailSecurityTopTldGetSpoofParamsSpf = "NONE"
	RadarEmailSecurityTopTldGetSpoofParamsSpfFail RadarEmailSecurityTopTldGetSpoofParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsSpfPass, RadarEmailSecurityTopTldGetSpoofParamsSpfNone, RadarEmailSecurityTopTldGetSpoofParamsSpfFail:
		return true
	}
	return false
}

// Filters results by TLD category.
type RadarEmailSecurityTopTldGetSpoofParamsTldCategory string

const (
	RadarEmailSecurityTopTldGetSpoofParamsTldCategoryClassic RadarEmailSecurityTopTldGetSpoofParamsTldCategory = "CLASSIC"
	RadarEmailSecurityTopTldGetSpoofParamsTldCategoryCountry RadarEmailSecurityTopTldGetSpoofParamsTldCategory = "COUNTRY"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsTldCategory) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsTldCategoryClassic, RadarEmailSecurityTopTldGetSpoofParamsTldCategoryCountry:
		return true
	}
	return false
}

type RadarEmailSecurityTopTldGetSpoofParamsTlsVersion string

const (
	RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_0 RadarEmailSecurityTopTldGetSpoofParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_1 RadarEmailSecurityTopTldGetSpoofParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_2 RadarEmailSecurityTopTldGetSpoofParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_3 RadarEmailSecurityTopTldGetSpoofParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTopTldGetSpoofParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_0, RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_1, RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_2, RadarEmailSecurityTopTldGetSpoofParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}
