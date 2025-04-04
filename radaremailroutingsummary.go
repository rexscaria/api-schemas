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

// RadarEmailRoutingSummaryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailRoutingSummaryService] method instead.
type RadarEmailRoutingSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailRoutingSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingSummaryService(opts ...option.RequestOption) (r *RadarEmailRoutingSummaryService) {
	r = &RadarEmailRoutingSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation.
func (r *RadarEmailRoutingSummaryService) GetArc(ctx context.Context, query RadarEmailRoutingSummaryGetArcParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation.
func (r *RadarEmailRoutingSummaryService) GetDkim(ctx context.Context, query RadarEmailRoutingSummaryGetDkimParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetDkimResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation.
func (r *RadarEmailRoutingSummaryService) GetDmarc(ctx context.Context, query RadarEmailRoutingSummaryGetDmarcParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted).
func (r *RadarEmailRoutingSummaryService) GetEncrypted(ctx context.Context, query RadarEmailRoutingSummaryGetEncryptedParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by IP version.
func (r *RadarEmailRoutingSummaryService) GetIPVersion(ctx context.Context, query RadarEmailRoutingSummaryGetIPVersionParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework)
// validation.
func (r *RadarEmailRoutingSummaryService) GetSpf(ctx context.Context, query RadarEmailRoutingSummaryGetSpfParams, opts ...option.RequestOption) (res *RadarEmailRoutingSummaryGetSpfResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailRoutingSummaryGetArcResponse struct {
	Result  RadarEmailRoutingSummaryGetArcResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetArcResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryGetArcResponse]
type radarEmailRoutingSummaryGetArcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetArcResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetArcResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetArcResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryGetArcResponseResult]
type radarEmailRoutingSummaryGetArcResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetArcResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetArcResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryGetArcResponseResultMeta]
type radarEmailRoutingSummaryGetArcResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetArcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryGetArcResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetArcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcResponseResultSummary0 struct {
	Fail string                                                   `json:"FAIL,required"`
	None string                                                   `json:"NONE,required"`
	Pass string                                                   `json:"PASS,required"`
	JSON radarEmailRoutingSummaryGetArcResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetArcResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetArcResponseResultSummary0]
type radarEmailRoutingSummaryGetArcResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetArcResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetArcResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponse struct {
	Result  RadarEmailRoutingSummaryGetDkimResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetDkimResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryGetDkimResponse]
type radarEmailRoutingSummaryGetDkimResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetDkimResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetDkimResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetDkimResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryGetDkimResponseResult]
type radarEmailRoutingSummaryGetDkimResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetDkimResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	Normalization  string                                                          `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetDkimResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryGetDkimResponseResultMeta]
type radarEmailRoutingSummaryGetDkimResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetDkimResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryGetDkimResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetDkimResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDkimResponseResultSummary0 struct {
	Fail string                                                    `json:"FAIL,required"`
	None string                                                    `json:"NONE,required"`
	Pass string                                                    `json:"PASS,required"`
	JSON radarEmailRoutingSummaryGetDkimResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetDkimResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetDkimResponseResultSummary0]
type radarEmailRoutingSummaryGetDkimResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDkimResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDkimResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponse struct {
	Result  RadarEmailRoutingSummaryGetDmarcResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetDmarcResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryGetDmarcResponse]
type radarEmailRoutingSummaryGetDmarcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetDmarcResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetDmarcResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetDmarcResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryGetDmarcResponseResult]
type radarEmailRoutingSummaryGetDmarcResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetDmarcResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	Normalization  string                                                           `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetDmarcResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetDmarcResponseResultMeta]
type radarEmailRoutingSummaryGetDmarcResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetDmarcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetDmarcResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetDmarcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetDmarcResponseResultSummary0 struct {
	Fail string                                                     `json:"FAIL,required"`
	None string                                                     `json:"NONE,required"`
	Pass string                                                     `json:"PASS,required"`
	JSON radarEmailRoutingSummaryGetDmarcResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetDmarcResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetDmarcResponseResultSummary0]
type radarEmailRoutingSummaryGetDmarcResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetDmarcResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetDmarcResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponse struct {
	Result  RadarEmailRoutingSummaryGetEncryptedResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetEncryptedResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryGetEncryptedResponse]
type radarEmailRoutingSummaryGetEncryptedResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetEncryptedResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetEncryptedResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetEncryptedResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetEncryptedResponseResult]
type radarEmailRoutingSummaryGetEncryptedResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	Normalization  string                                                               `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetEncryptedResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetEncryptedResponseResultMeta]
type radarEmailRoutingSummaryGetEncryptedResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetEncryptedResponseResultSummary0 struct {
	Encrypted    string                                                         `json:"ENCRYPTED,required"`
	NotEncrypted string                                                         `json:"NOT_ENCRYPTED,required"`
	JSON         radarEmailRoutingSummaryGetEncryptedResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetEncryptedResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryGetEncryptedResponseResultSummary0]
type radarEmailRoutingSummaryGetEncryptedResponseResultSummary0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetEncryptedResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetEncryptedResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponse struct {
	Result  RadarEmailRoutingSummaryGetIPVersionResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetIPVersionResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryGetIPVersionResponse]
type radarEmailRoutingSummaryGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetIPVersionResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetIPVersionResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetIPVersionResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetIPVersionResponseResult]
type radarEmailRoutingSummaryGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	Normalization  string                                                               `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetIPVersionResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetIPVersionResponseResultMeta]
type radarEmailRoutingSummaryGetIPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetIPVersionResponseResultSummary0 struct {
	IPv4 string                                                         `json:"IPv4,required"`
	IPv6 string                                                         `json:"IPv6,required"`
	JSON radarEmailRoutingSummaryGetIPVersionResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetIPVersionResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryGetIPVersionResponseResultSummary0]
type radarEmailRoutingSummaryGetIPVersionResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetIPVersionResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetIPVersionResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponse struct {
	Result  RadarEmailRoutingSummaryGetSpfResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEmailRoutingSummaryGetSpfResponseJSON   `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseJSON contains the JSON metadata for the
// struct [RadarEmailRoutingSummaryGetSpfResponse]
type radarEmailRoutingSummaryGetSpfResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResult struct {
	Meta     RadarEmailRoutingSummaryGetSpfResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailRoutingSummaryGetSpfResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailRoutingSummaryGetSpfResponseResultJSON     `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailRoutingSummaryGetSpfResponseResult]
type radarEmailRoutingSummaryGetSpfResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResultMeta struct {
	DateRange      []RadarEmailRoutingSummaryGetSpfResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	Normalization  string                                                         `json:"normalization,required"`
	ConfidenceInfo RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailRoutingSummaryGetSpfResponseResultMetaJSON           `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailRoutingSummaryGetSpfResponseResultMeta]
type radarEmailRoutingSummaryGetSpfResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingSummaryGetSpfResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingSummaryGetSpfResponseResultMetaDateRange]
type radarEmailRoutingSummaryGetSpfResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfo]
type radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetSpfResponseResultSummary0 struct {
	Fail string                                                   `json:"FAIL,required"`
	None string                                                   `json:"NONE,required"`
	Pass string                                                   `json:"PASS,required"`
	JSON radarEmailRoutingSummaryGetSpfResponseResultSummary0JSON `json:"-"`
}

// radarEmailRoutingSummaryGetSpfResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailRoutingSummaryGetSpfResponseResultSummary0]
type radarEmailRoutingSummaryGetSpfResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingSummaryGetSpfResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingSummaryGetSpfResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingSummaryGetArcParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingSummaryGetArcParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingSummaryGetArcParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingSummaryGetArcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetArcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingSummaryGetArcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingSummaryGetArcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetArcParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummaryGetArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetArcParamsDkim string

const (
	RadarEmailRoutingSummaryGetArcParamsDkimPass RadarEmailRoutingSummaryGetArcParamsDkim = "PASS"
	RadarEmailRoutingSummaryGetArcParamsDkimNone RadarEmailRoutingSummaryGetArcParamsDkim = "NONE"
	RadarEmailRoutingSummaryGetArcParamsDkimFail RadarEmailRoutingSummaryGetArcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingSummaryGetArcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsDkimPass, RadarEmailRoutingSummaryGetArcParamsDkimNone, RadarEmailRoutingSummaryGetArcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetArcParamsDmarc string

const (
	RadarEmailRoutingSummaryGetArcParamsDmarcPass RadarEmailRoutingSummaryGetArcParamsDmarc = "PASS"
	RadarEmailRoutingSummaryGetArcParamsDmarcNone RadarEmailRoutingSummaryGetArcParamsDmarc = "NONE"
	RadarEmailRoutingSummaryGetArcParamsDmarcFail RadarEmailRoutingSummaryGetArcParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetArcParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsDmarcPass, RadarEmailRoutingSummaryGetArcParamsDmarcNone, RadarEmailRoutingSummaryGetArcParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetArcParamsEncrypted string

const (
	RadarEmailRoutingSummaryGetArcParamsEncryptedEncrypted    RadarEmailRoutingSummaryGetArcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryGetArcParamsEncryptedNotEncrypted RadarEmailRoutingSummaryGetArcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingSummaryGetArcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsEncryptedEncrypted, RadarEmailRoutingSummaryGetArcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetArcParamsFormat string

const (
	RadarEmailRoutingSummaryGetArcParamsFormatJson RadarEmailRoutingSummaryGetArcParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetArcParamsFormatCsv  RadarEmailRoutingSummaryGetArcParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetArcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsFormatJson, RadarEmailRoutingSummaryGetArcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetArcParamsIPVersion string

const (
	RadarEmailRoutingSummaryGetArcParamsIPVersionIPv4 RadarEmailRoutingSummaryGetArcParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryGetArcParamsIPVersionIPv6 RadarEmailRoutingSummaryGetArcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingSummaryGetArcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsIPVersionIPv4, RadarEmailRoutingSummaryGetArcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetArcParamsSpf string

const (
	RadarEmailRoutingSummaryGetArcParamsSpfPass RadarEmailRoutingSummaryGetArcParamsSpf = "PASS"
	RadarEmailRoutingSummaryGetArcParamsSpfNone RadarEmailRoutingSummaryGetArcParamsSpf = "NONE"
	RadarEmailRoutingSummaryGetArcParamsSpfFail RadarEmailRoutingSummaryGetArcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingSummaryGetArcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetArcParamsSpfPass, RadarEmailRoutingSummaryGetArcParamsSpfNone, RadarEmailRoutingSummaryGetArcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDkimParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingSummaryGetDkimParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingSummaryGetDkimParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingSummaryGetDkimParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetDkimParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingSummaryGetDkimParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingSummaryGetDkimParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetDkimParams]'s query parameters
// as `url.Values`.
func (r RadarEmailRoutingSummaryGetDkimParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetDkimParamsArc string

const (
	RadarEmailRoutingSummaryGetDkimParamsArcPass RadarEmailRoutingSummaryGetDkimParamsArc = "PASS"
	RadarEmailRoutingSummaryGetDkimParamsArcNone RadarEmailRoutingSummaryGetDkimParamsArc = "NONE"
	RadarEmailRoutingSummaryGetDkimParamsArcFail RadarEmailRoutingSummaryGetDkimParamsArc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDkimParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsArcPass, RadarEmailRoutingSummaryGetDkimParamsArcNone, RadarEmailRoutingSummaryGetDkimParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDkimParamsDmarc string

const (
	RadarEmailRoutingSummaryGetDkimParamsDmarcPass RadarEmailRoutingSummaryGetDkimParamsDmarc = "PASS"
	RadarEmailRoutingSummaryGetDkimParamsDmarcNone RadarEmailRoutingSummaryGetDkimParamsDmarc = "NONE"
	RadarEmailRoutingSummaryGetDkimParamsDmarcFail RadarEmailRoutingSummaryGetDkimParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDkimParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsDmarcPass, RadarEmailRoutingSummaryGetDkimParamsDmarcNone, RadarEmailRoutingSummaryGetDkimParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDkimParamsEncrypted string

const (
	RadarEmailRoutingSummaryGetDkimParamsEncryptedEncrypted    RadarEmailRoutingSummaryGetDkimParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryGetDkimParamsEncryptedNotEncrypted RadarEmailRoutingSummaryGetDkimParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingSummaryGetDkimParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsEncryptedEncrypted, RadarEmailRoutingSummaryGetDkimParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetDkimParamsFormat string

const (
	RadarEmailRoutingSummaryGetDkimParamsFormatJson RadarEmailRoutingSummaryGetDkimParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetDkimParamsFormatCsv  RadarEmailRoutingSummaryGetDkimParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetDkimParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsFormatJson, RadarEmailRoutingSummaryGetDkimParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDkimParamsIPVersion string

const (
	RadarEmailRoutingSummaryGetDkimParamsIPVersionIPv4 RadarEmailRoutingSummaryGetDkimParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryGetDkimParamsIPVersionIPv6 RadarEmailRoutingSummaryGetDkimParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingSummaryGetDkimParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsIPVersionIPv4, RadarEmailRoutingSummaryGetDkimParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDkimParamsSpf string

const (
	RadarEmailRoutingSummaryGetDkimParamsSpfPass RadarEmailRoutingSummaryGetDkimParamsSpf = "PASS"
	RadarEmailRoutingSummaryGetDkimParamsSpfNone RadarEmailRoutingSummaryGetDkimParamsSpf = "NONE"
	RadarEmailRoutingSummaryGetDkimParamsSpfFail RadarEmailRoutingSummaryGetDkimParamsSpf = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDkimParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDkimParamsSpfPass, RadarEmailRoutingSummaryGetDkimParamsSpfNone, RadarEmailRoutingSummaryGetDkimParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDmarcParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingSummaryGetDmarcParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingSummaryGetDmarcParamsDkim] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingSummaryGetDmarcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetDmarcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingSummaryGetDmarcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingSummaryGetDmarcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetDmarcParams]'s query parameters
// as `url.Values`.
func (r RadarEmailRoutingSummaryGetDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetDmarcParamsArc string

const (
	RadarEmailRoutingSummaryGetDmarcParamsArcPass RadarEmailRoutingSummaryGetDmarcParamsArc = "PASS"
	RadarEmailRoutingSummaryGetDmarcParamsArcNone RadarEmailRoutingSummaryGetDmarcParamsArc = "NONE"
	RadarEmailRoutingSummaryGetDmarcParamsArcFail RadarEmailRoutingSummaryGetDmarcParamsArc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsArcPass, RadarEmailRoutingSummaryGetDmarcParamsArcNone, RadarEmailRoutingSummaryGetDmarcParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDmarcParamsDkim string

const (
	RadarEmailRoutingSummaryGetDmarcParamsDkimPass RadarEmailRoutingSummaryGetDmarcParamsDkim = "PASS"
	RadarEmailRoutingSummaryGetDmarcParamsDkimNone RadarEmailRoutingSummaryGetDmarcParamsDkim = "NONE"
	RadarEmailRoutingSummaryGetDmarcParamsDkimFail RadarEmailRoutingSummaryGetDmarcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsDkimPass, RadarEmailRoutingSummaryGetDmarcParamsDkimNone, RadarEmailRoutingSummaryGetDmarcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDmarcParamsEncrypted string

const (
	RadarEmailRoutingSummaryGetDmarcParamsEncryptedEncrypted    RadarEmailRoutingSummaryGetDmarcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryGetDmarcParamsEncryptedNotEncrypted RadarEmailRoutingSummaryGetDmarcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsEncryptedEncrypted, RadarEmailRoutingSummaryGetDmarcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetDmarcParamsFormat string

const (
	RadarEmailRoutingSummaryGetDmarcParamsFormatJson RadarEmailRoutingSummaryGetDmarcParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetDmarcParamsFormatCsv  RadarEmailRoutingSummaryGetDmarcParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsFormatJson, RadarEmailRoutingSummaryGetDmarcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDmarcParamsIPVersion string

const (
	RadarEmailRoutingSummaryGetDmarcParamsIPVersionIPv4 RadarEmailRoutingSummaryGetDmarcParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryGetDmarcParamsIPVersionIPv6 RadarEmailRoutingSummaryGetDmarcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsIPVersionIPv4, RadarEmailRoutingSummaryGetDmarcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetDmarcParamsSpf string

const (
	RadarEmailRoutingSummaryGetDmarcParamsSpfPass RadarEmailRoutingSummaryGetDmarcParamsSpf = "PASS"
	RadarEmailRoutingSummaryGetDmarcParamsSpfNone RadarEmailRoutingSummaryGetDmarcParamsSpf = "NONE"
	RadarEmailRoutingSummaryGetDmarcParamsSpfFail RadarEmailRoutingSummaryGetDmarcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingSummaryGetDmarcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetDmarcParamsSpfPass, RadarEmailRoutingSummaryGetDmarcParamsSpfNone, RadarEmailRoutingSummaryGetDmarcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetEncryptedParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingSummaryGetEncryptedParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingSummaryGetEncryptedParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingSummaryGetEncryptedParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingSummaryGetEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingSummaryGetEncryptedParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetEncryptedParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingSummaryGetEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetEncryptedParamsArc string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsArcPass RadarEmailRoutingSummaryGetEncryptedParamsArc = "PASS"
	RadarEmailRoutingSummaryGetEncryptedParamsArcNone RadarEmailRoutingSummaryGetEncryptedParamsArc = "NONE"
	RadarEmailRoutingSummaryGetEncryptedParamsArcFail RadarEmailRoutingSummaryGetEncryptedParamsArc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsArcPass, RadarEmailRoutingSummaryGetEncryptedParamsArcNone, RadarEmailRoutingSummaryGetEncryptedParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetEncryptedParamsDkim string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsDkimPass RadarEmailRoutingSummaryGetEncryptedParamsDkim = "PASS"
	RadarEmailRoutingSummaryGetEncryptedParamsDkimNone RadarEmailRoutingSummaryGetEncryptedParamsDkim = "NONE"
	RadarEmailRoutingSummaryGetEncryptedParamsDkimFail RadarEmailRoutingSummaryGetEncryptedParamsDkim = "FAIL"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsDkimPass, RadarEmailRoutingSummaryGetEncryptedParamsDkimNone, RadarEmailRoutingSummaryGetEncryptedParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetEncryptedParamsDmarc string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsDmarcPass RadarEmailRoutingSummaryGetEncryptedParamsDmarc = "PASS"
	RadarEmailRoutingSummaryGetEncryptedParamsDmarcNone RadarEmailRoutingSummaryGetEncryptedParamsDmarc = "NONE"
	RadarEmailRoutingSummaryGetEncryptedParamsDmarcFail RadarEmailRoutingSummaryGetEncryptedParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsDmarcPass, RadarEmailRoutingSummaryGetEncryptedParamsDmarcNone, RadarEmailRoutingSummaryGetEncryptedParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetEncryptedParamsFormat string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsFormatJson RadarEmailRoutingSummaryGetEncryptedParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetEncryptedParamsFormatCsv  RadarEmailRoutingSummaryGetEncryptedParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsFormatJson, RadarEmailRoutingSummaryGetEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetEncryptedParamsIPVersion string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsIPVersionIPv4 RadarEmailRoutingSummaryGetEncryptedParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryGetEncryptedParamsIPVersionIPv6 RadarEmailRoutingSummaryGetEncryptedParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsIPVersionIPv4, RadarEmailRoutingSummaryGetEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetEncryptedParamsSpf string

const (
	RadarEmailRoutingSummaryGetEncryptedParamsSpfPass RadarEmailRoutingSummaryGetEncryptedParamsSpf = "PASS"
	RadarEmailRoutingSummaryGetEncryptedParamsSpfNone RadarEmailRoutingSummaryGetEncryptedParamsSpf = "NONE"
	RadarEmailRoutingSummaryGetEncryptedParamsSpfFail RadarEmailRoutingSummaryGetEncryptedParamsSpf = "FAIL"
)

func (r RadarEmailRoutingSummaryGetEncryptedParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetEncryptedParamsSpfPass, RadarEmailRoutingSummaryGetEncryptedParamsSpfNone, RadarEmailRoutingSummaryGetEncryptedParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetIPVersionParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingSummaryGetIPVersionParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingSummaryGetIPVersionParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingSummaryGetIPVersionParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingSummaryGetIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingSummaryGetIPVersionParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingSummaryGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetIPVersionParamsArc string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsArcPass RadarEmailRoutingSummaryGetIPVersionParamsArc = "PASS"
	RadarEmailRoutingSummaryGetIPVersionParamsArcNone RadarEmailRoutingSummaryGetIPVersionParamsArc = "NONE"
	RadarEmailRoutingSummaryGetIPVersionParamsArcFail RadarEmailRoutingSummaryGetIPVersionParamsArc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsArcPass, RadarEmailRoutingSummaryGetIPVersionParamsArcNone, RadarEmailRoutingSummaryGetIPVersionParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetIPVersionParamsDkim string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsDkimPass RadarEmailRoutingSummaryGetIPVersionParamsDkim = "PASS"
	RadarEmailRoutingSummaryGetIPVersionParamsDkimNone RadarEmailRoutingSummaryGetIPVersionParamsDkim = "NONE"
	RadarEmailRoutingSummaryGetIPVersionParamsDkimFail RadarEmailRoutingSummaryGetIPVersionParamsDkim = "FAIL"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsDkimPass, RadarEmailRoutingSummaryGetIPVersionParamsDkimNone, RadarEmailRoutingSummaryGetIPVersionParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetIPVersionParamsDmarc string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsDmarcPass RadarEmailRoutingSummaryGetIPVersionParamsDmarc = "PASS"
	RadarEmailRoutingSummaryGetIPVersionParamsDmarcNone RadarEmailRoutingSummaryGetIPVersionParamsDmarc = "NONE"
	RadarEmailRoutingSummaryGetIPVersionParamsDmarcFail RadarEmailRoutingSummaryGetIPVersionParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsDmarcPass, RadarEmailRoutingSummaryGetIPVersionParamsDmarcNone, RadarEmailRoutingSummaryGetIPVersionParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetIPVersionParamsEncrypted string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsEncryptedEncrypted    RadarEmailRoutingSummaryGetIPVersionParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryGetIPVersionParamsEncryptedNotEncrypted RadarEmailRoutingSummaryGetIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsEncryptedEncrypted, RadarEmailRoutingSummaryGetIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetIPVersionParamsFormat string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsFormatJson RadarEmailRoutingSummaryGetIPVersionParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetIPVersionParamsFormatCsv  RadarEmailRoutingSummaryGetIPVersionParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsFormatJson, RadarEmailRoutingSummaryGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetIPVersionParamsSpf string

const (
	RadarEmailRoutingSummaryGetIPVersionParamsSpfPass RadarEmailRoutingSummaryGetIPVersionParamsSpf = "PASS"
	RadarEmailRoutingSummaryGetIPVersionParamsSpfNone RadarEmailRoutingSummaryGetIPVersionParamsSpf = "NONE"
	RadarEmailRoutingSummaryGetIPVersionParamsSpfFail RadarEmailRoutingSummaryGetIPVersionParamsSpf = "FAIL"
)

func (r RadarEmailRoutingSummaryGetIPVersionParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetIPVersionParamsSpfPass, RadarEmailRoutingSummaryGetIPVersionParamsSpfNone, RadarEmailRoutingSummaryGetIPVersionParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetSpfParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingSummaryGetSpfParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingSummaryGetSpfParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingSummaryGetSpfParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingSummaryGetSpfParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingSummaryGetSpfParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingSummaryGetSpfParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailRoutingSummaryGetSpfParams]'s query parameters as
// `url.Values`.
func (r RadarEmailRoutingSummaryGetSpfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailRoutingSummaryGetSpfParamsArc string

const (
	RadarEmailRoutingSummaryGetSpfParamsArcPass RadarEmailRoutingSummaryGetSpfParamsArc = "PASS"
	RadarEmailRoutingSummaryGetSpfParamsArcNone RadarEmailRoutingSummaryGetSpfParamsArc = "NONE"
	RadarEmailRoutingSummaryGetSpfParamsArcFail RadarEmailRoutingSummaryGetSpfParamsArc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetSpfParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsArcPass, RadarEmailRoutingSummaryGetSpfParamsArcNone, RadarEmailRoutingSummaryGetSpfParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetSpfParamsDkim string

const (
	RadarEmailRoutingSummaryGetSpfParamsDkimPass RadarEmailRoutingSummaryGetSpfParamsDkim = "PASS"
	RadarEmailRoutingSummaryGetSpfParamsDkimNone RadarEmailRoutingSummaryGetSpfParamsDkim = "NONE"
	RadarEmailRoutingSummaryGetSpfParamsDkimFail RadarEmailRoutingSummaryGetSpfParamsDkim = "FAIL"
)

func (r RadarEmailRoutingSummaryGetSpfParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsDkimPass, RadarEmailRoutingSummaryGetSpfParamsDkimNone, RadarEmailRoutingSummaryGetSpfParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetSpfParamsDmarc string

const (
	RadarEmailRoutingSummaryGetSpfParamsDmarcPass RadarEmailRoutingSummaryGetSpfParamsDmarc = "PASS"
	RadarEmailRoutingSummaryGetSpfParamsDmarcNone RadarEmailRoutingSummaryGetSpfParamsDmarc = "NONE"
	RadarEmailRoutingSummaryGetSpfParamsDmarcFail RadarEmailRoutingSummaryGetSpfParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingSummaryGetSpfParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsDmarcPass, RadarEmailRoutingSummaryGetSpfParamsDmarcNone, RadarEmailRoutingSummaryGetSpfParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetSpfParamsEncrypted string

const (
	RadarEmailRoutingSummaryGetSpfParamsEncryptedEncrypted    RadarEmailRoutingSummaryGetSpfParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingSummaryGetSpfParamsEncryptedNotEncrypted RadarEmailRoutingSummaryGetSpfParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingSummaryGetSpfParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsEncryptedEncrypted, RadarEmailRoutingSummaryGetSpfParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingSummaryGetSpfParamsFormat string

const (
	RadarEmailRoutingSummaryGetSpfParamsFormatJson RadarEmailRoutingSummaryGetSpfParamsFormat = "JSON"
	RadarEmailRoutingSummaryGetSpfParamsFormatCsv  RadarEmailRoutingSummaryGetSpfParamsFormat = "CSV"
)

func (r RadarEmailRoutingSummaryGetSpfParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsFormatJson, RadarEmailRoutingSummaryGetSpfParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingSummaryGetSpfParamsIPVersion string

const (
	RadarEmailRoutingSummaryGetSpfParamsIPVersionIPv4 RadarEmailRoutingSummaryGetSpfParamsIPVersion = "IPv4"
	RadarEmailRoutingSummaryGetSpfParamsIPVersionIPv6 RadarEmailRoutingSummaryGetSpfParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingSummaryGetSpfParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingSummaryGetSpfParamsIPVersionIPv4, RadarEmailRoutingSummaryGetSpfParamsIPVersionIPv6:
		return true
	}
	return false
}
