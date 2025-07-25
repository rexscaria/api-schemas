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

// RadarEmailSecuritySummaryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecuritySummaryService] method instead.
type RadarEmailSecuritySummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryService) {
	r = &RadarEmailSecuritySummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation.
func (r *RadarEmailSecuritySummaryService) GetArc(ctx context.Context, query RadarEmailSecuritySummaryGetArcParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation.
func (r *RadarEmailSecuritySummaryService) GetDkim(ctx context.Context, query RadarEmailSecuritySummaryGetDkimParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetDkimResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation.
func (r *RadarEmailSecuritySummaryService) GetDmarc(ctx context.Context, query RadarEmailSecuritySummaryGetDmarcParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by malicious classification.
func (r *RadarEmailSecuritySummaryService) GetMalicious(ctx context.Context, query RadarEmailSecuritySummaryGetMaliciousParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the proportion of emails by spam classification (spam vs. non-spam).
func (r *RadarEmailSecuritySummaryService) GetSpam(ctx context.Context, query RadarEmailSecuritySummaryGetSpamParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetSpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework)
// validation.
func (r *RadarEmailSecuritySummaryService) GetSpf(ctx context.Context, query RadarEmailSecuritySummaryGetSpfParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetSpfResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the proportion of emails by spoof classification (spoof vs.
// non-spoof).
func (r *RadarEmailSecuritySummaryService) GetSpoof(ctx context.Context, query RadarEmailSecuritySummaryGetSpoofParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetSpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by threat categories.
func (r *RadarEmailSecuritySummaryService) GetThreatCategory(ctx context.Context, query RadarEmailSecuritySummaryGetThreatCategoryParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by TLS version.
func (r *RadarEmailSecuritySummaryService) GetTlsVersion(ctx context.Context, query RadarEmailSecuritySummaryGetTlsVersionParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryGetTlsVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySummaryGetArcResponse struct {
	Result  RadarEmailSecuritySummaryGetArcResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetArcResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetArcResponse]
type radarEmailSecuritySummaryGetArcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetArcResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetArcResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetArcResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetArcResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryGetArcResponseResult]
type radarEmailSecuritySummaryGetArcResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetArcResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetArcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetArcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetArcResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetArcResponseResultMeta]
type radarEmailSecuritySummaryGetArcResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetArcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetArcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetArcResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetArcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetArcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetArcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetArcResponseResultMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetArcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetArcResponseResultMetaUnit]
type radarEmailSecuritySummaryGetArcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetArcResponseResultSummary0 struct {
	// A numeric string.
	Fail string `json:"FAIL,required"`
	// A numeric string.
	None string `json:"NONE,required"`
	// A numeric string.
	Pass string                                                    `json:"PASS,required"`
	JSON radarEmailSecuritySummaryGetArcResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetArcResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetArcResponseResultSummary0]
type radarEmailSecuritySummaryGetArcResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetArcResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetArcResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDkimResponse struct {
	Result  RadarEmailSecuritySummaryGetDkimResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetDkimResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetDkimResponse]
type radarEmailSecuritySummaryGetDkimResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDkimResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetDkimResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetDkimResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetDkimResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetDkimResponseResult]
type radarEmailSecuritySummaryGetDkimResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetDkimResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetDkimResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetDkimResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetDkimResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetDkimResponseResultMeta]
type radarEmailSecuritySummaryGetDkimResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDkimResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetDkimResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDkimResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetDkimResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetDkimResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDkimResponseResultMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetDkimResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetDkimResponseResultMetaUnit]
type radarEmailSecuritySummaryGetDkimResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDkimResponseResultSummary0 struct {
	// A numeric string.
	Fail string `json:"FAIL,required"`
	// A numeric string.
	None string `json:"NONE,required"`
	// A numeric string.
	Pass string                                                     `json:"PASS,required"`
	JSON radarEmailSecuritySummaryGetDkimResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetDkimResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetDkimResponseResultSummary0]
type radarEmailSecuritySummaryGetDkimResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDkimResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDkimResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDmarcResponse struct {
	Result  RadarEmailSecuritySummaryGetDmarcResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetDmarcResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetDmarcResponse]
type radarEmailSecuritySummaryGetDmarcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDmarcResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetDmarcResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetDmarcResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetDmarcResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetDmarcResponseResult]
type radarEmailSecuritySummaryGetDmarcResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetDmarcResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetDmarcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetDmarcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetDmarcResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetDmarcResponseResultMeta]
type radarEmailSecuritySummaryGetDmarcResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                 `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                       `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDmarcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetDmarcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetDmarcResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetDmarcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetDmarcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDmarcResponseResultMetaUnit struct {
	Name  string                                                      `json:"name,required"`
	Value string                                                      `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetDmarcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetDmarcResponseResultMetaUnit]
type radarEmailSecuritySummaryGetDmarcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetDmarcResponseResultSummary0 struct {
	// A numeric string.
	Fail string `json:"FAIL,required"`
	// A numeric string.
	None string `json:"NONE,required"`
	// A numeric string.
	Pass string                                                      `json:"PASS,required"`
	JSON radarEmailSecuritySummaryGetDmarcResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetDmarcResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetDmarcResponseResultSummary0]
type radarEmailSecuritySummaryGetDmarcResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetDmarcResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetDmarcResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetMaliciousResponse struct {
	Result  RadarEmailSecuritySummaryGetMaliciousResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetMaliciousResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryGetMaliciousResponse]
type radarEmailSecuritySummaryGetMaliciousResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetMaliciousResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetMaliciousResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetMaliciousResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetMaliciousResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetMaliciousResponseResult]
type radarEmailSecuritySummaryGetMaliciousResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetMaliciousResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetMaliciousResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetMaliciousResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultMeta]
type radarEmailSecuritySummaryGetMaliciousResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                     `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetMaliciousResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousResponseResultMetaUnit struct {
	Name  string                                                          `json:"name,required"`
	Value string                                                          `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetMaliciousResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultMetaUnit]
type radarEmailSecuritySummaryGetMaliciousResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetMaliciousResponseResultSummary0 struct {
	// A numeric string.
	Malicious string `json:"MALICIOUS,required"`
	// A numeric string.
	NotMalicious string                                                          `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecuritySummaryGetMaliciousResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetMaliciousResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetMaliciousResponseResultSummary0]
type radarEmailSecuritySummaryGetMaliciousResponseResultSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetMaliciousResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetMaliciousResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpamResponse struct {
	Result  RadarEmailSecuritySummaryGetSpamResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetSpamResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetSpamResponse]
type radarEmailSecuritySummaryGetSpamResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpamResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetSpamResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetSpamResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetSpamResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetSpamResponseResult]
type radarEmailSecuritySummaryGetSpamResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetSpamResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetSpamResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetSpamResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetSpamResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpamResponseResultMeta]
type radarEmailSecuritySummaryGetSpamResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpamResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetSpamResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpamResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetSpamResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetSpamResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamResponseResultMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetSpamResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpamResponseResultMetaUnit]
type radarEmailSecuritySummaryGetSpamResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpamResponseResultSummary0 struct {
	// A numeric string.
	NotSpam string `json:"NOT_SPAM,required"`
	// A numeric string.
	Spam string                                                     `json:"SPAM,required"`
	JSON radarEmailSecuritySummaryGetSpamResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpamResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpamResponseResultSummary0]
type radarEmailSecuritySummaryGetSpamResponseResultSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpamResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpamResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpfResponse struct {
	Result  RadarEmailSecuritySummaryGetSpfResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetSpfResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetSpfResponse]
type radarEmailSecuritySummaryGetSpfResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpfResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetSpfResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetSpfResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetSpfResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryGetSpfResponseResult]
type radarEmailSecuritySummaryGetSpfResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetSpfResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetSpfResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetSpfResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetSpfResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetSpfResponseResultMeta]
type radarEmailSecuritySummaryGetSpfResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpfResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetSpfResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetSpfResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetSpfResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetSpfResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpfResponseResultMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetSpfResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpfResponseResultMetaUnit]
type radarEmailSecuritySummaryGetSpfResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpfResponseResultSummary0 struct {
	// A numeric string.
	Fail string `json:"FAIL,required"`
	// A numeric string.
	None string `json:"NONE,required"`
	// A numeric string.
	Pass string                                                    `json:"PASS,required"`
	JSON radarEmailSecuritySummaryGetSpfResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpfResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpfResponseResultSummary0]
type radarEmailSecuritySummaryGetSpfResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpfResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpfResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpoofResponse struct {
	Result  RadarEmailSecuritySummaryGetSpoofResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetSpoofResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryGetSpoofResponse]
type radarEmailSecuritySummaryGetSpoofResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpoofResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetSpoofResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetSpoofResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetSpoofResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetSpoofResponseResult]
type radarEmailSecuritySummaryGetSpoofResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetSpoofResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetSpoofResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetSpoofResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetSpoofResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetSpoofResponseResultMeta]
type radarEmailSecuritySummaryGetSpoofResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                 `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                       `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpoofResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetSpoofResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetSpoofResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetSpoofResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetSpoofResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofResponseResultMetaUnit struct {
	Name  string                                                      `json:"name,required"`
	Value string                                                      `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetSpoofResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetSpoofResponseResultMetaUnit]
type radarEmailSecuritySummaryGetSpoofResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetSpoofResponseResultSummary0 struct {
	// A numeric string.
	NotSpoof string `json:"NOT_SPOOF,required"`
	// A numeric string.
	Spoof string                                                      `json:"SPOOF,required"`
	JSON  radarEmailSecuritySummaryGetSpoofResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetSpoofResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetSpoofResponseResultSummary0]
type radarEmailSecuritySummaryGetSpoofResponseResultSummary0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetSpoofResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetSpoofResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetThreatCategoryResponse struct {
	Result  RadarEmailSecuritySummaryGetThreatCategoryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetThreatCategoryResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetThreatCategoryResponse]
type radarEmailSecuritySummaryGetThreatCategoryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetThreatCategoryResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetThreatCategoryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetThreatCategoryResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResult]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultMeta]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnit]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0 struct {
	// A numeric string.
	BrandImpersonation string `json:"BrandImpersonation,required"`
	// A numeric string.
	CredentialHarvester string `json:"CredentialHarvester,required"`
	// A numeric string.
	IdentityDeception string `json:"IdentityDeception,required"`
	// A numeric string.
	Link string                                                               `json:"Link,required"`
	JSON radarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0]
type radarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetThreatCategoryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetTlsVersionResponse struct {
	Result  RadarEmailSecuritySummaryGetTlsVersionResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecuritySummaryGetTlsVersionResponseJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryGetTlsVersionResponse]
type radarEmailSecuritySummaryGetTlsVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetTlsVersionResponseResult struct {
	// Metadata for the results.
	Meta     RadarEmailSecuritySummaryGetTlsVersionResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryGetTlsVersionResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryGetTlsVersionResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryGetTlsVersionResponseResult]
type radarEmailSecuritySummaryGetTlsVersionResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecuritySummaryGetTlsVersionResponseResultMeta struct {
	ConfidenceInfo RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecuritySummaryGetTlsVersionResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultMeta]
type radarEmailSecuritySummaryGetTlsVersionResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfo]
type radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRange]
type radarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization string

const (
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationPercentage           RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationMin0Max              RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationMinMax               RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRawValues            RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationPercentageChange     RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRollingAverage       RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRatio                RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationPercentage, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationMin0Max, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationMinMax, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRawValues, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationPercentageChange, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRollingAverage, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  radarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnit]
type radarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetTlsVersionResponseResultSummary0 struct {
	// A numeric string.
	Tls1_0 string `json:"TLS 1.0,required"`
	// A numeric string.
	Tls1_1 string `json:"TLS 1.1,required"`
	// A numeric string.
	Tls1_2 string `json:"TLS 1.2,required"`
	// A numeric string.
	Tls1_3 string                                                           `json:"TLS 1.3,required"`
	JSON   radarEmailSecuritySummaryGetTlsVersionResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryGetTlsVersionResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryGetTlsVersionResponseResultSummary0]
type radarEmailSecuritySummaryGetTlsVersionResponseResultSummary0JSON struct {
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryGetTlsVersionResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecuritySummaryGetTlsVersionResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecuritySummaryGetArcParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetArcParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetArcParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetArcParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetArcParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetArcParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetArcParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetArcParamsDkim string

const (
	RadarEmailSecuritySummaryGetArcParamsDkimPass RadarEmailSecuritySummaryGetArcParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetArcParamsDkimNone RadarEmailSecuritySummaryGetArcParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetArcParamsDkimFail RadarEmailSecuritySummaryGetArcParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetArcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcParamsDkimPass, RadarEmailSecuritySummaryGetArcParamsDkimNone, RadarEmailSecuritySummaryGetArcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetArcParamsDmarc string

const (
	RadarEmailSecuritySummaryGetArcParamsDmarcPass RadarEmailSecuritySummaryGetArcParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetArcParamsDmarcNone RadarEmailSecuritySummaryGetArcParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetArcParamsDmarcFail RadarEmailSecuritySummaryGetArcParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetArcParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcParamsDmarcPass, RadarEmailSecuritySummaryGetArcParamsDmarcNone, RadarEmailSecuritySummaryGetArcParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetArcParamsFormat string

const (
	RadarEmailSecuritySummaryGetArcParamsFormatJson RadarEmailSecuritySummaryGetArcParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetArcParamsFormatCsv  RadarEmailSecuritySummaryGetArcParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetArcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcParamsFormatJson, RadarEmailSecuritySummaryGetArcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetArcParamsSpf string

const (
	RadarEmailSecuritySummaryGetArcParamsSpfPass RadarEmailSecuritySummaryGetArcParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetArcParamsSpfNone RadarEmailSecuritySummaryGetArcParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetArcParamsSpfFail RadarEmailSecuritySummaryGetArcParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetArcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcParamsSpfPass, RadarEmailSecuritySummaryGetArcParamsSpfNone, RadarEmailSecuritySummaryGetArcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetArcParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetArcParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetArcParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetArcParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetArcParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetArcParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetArcParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDkimParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetDkimParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetDkimParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetDkimParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetDkimParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetDkimParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetDkimParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetDkimParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetDkimParamsArc string

const (
	RadarEmailSecuritySummaryGetDkimParamsArcPass RadarEmailSecuritySummaryGetDkimParamsArc = "PASS"
	RadarEmailSecuritySummaryGetDkimParamsArcNone RadarEmailSecuritySummaryGetDkimParamsArc = "NONE"
	RadarEmailSecuritySummaryGetDkimParamsArcFail RadarEmailSecuritySummaryGetDkimParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDkimParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimParamsArcPass, RadarEmailSecuritySummaryGetDkimParamsArcNone, RadarEmailSecuritySummaryGetDkimParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDkimParamsDmarc string

const (
	RadarEmailSecuritySummaryGetDkimParamsDmarcPass RadarEmailSecuritySummaryGetDkimParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetDkimParamsDmarcNone RadarEmailSecuritySummaryGetDkimParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetDkimParamsDmarcFail RadarEmailSecuritySummaryGetDkimParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDkimParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimParamsDmarcPass, RadarEmailSecuritySummaryGetDkimParamsDmarcNone, RadarEmailSecuritySummaryGetDkimParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetDkimParamsFormat string

const (
	RadarEmailSecuritySummaryGetDkimParamsFormatJson RadarEmailSecuritySummaryGetDkimParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetDkimParamsFormatCsv  RadarEmailSecuritySummaryGetDkimParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetDkimParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimParamsFormatJson, RadarEmailSecuritySummaryGetDkimParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDkimParamsSpf string

const (
	RadarEmailSecuritySummaryGetDkimParamsSpfPass RadarEmailSecuritySummaryGetDkimParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetDkimParamsSpfNone RadarEmailSecuritySummaryGetDkimParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetDkimParamsSpfFail RadarEmailSecuritySummaryGetDkimParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDkimParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimParamsSpfPass, RadarEmailSecuritySummaryGetDkimParamsSpfNone, RadarEmailSecuritySummaryGetDkimParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDkimParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetDkimParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetDkimParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetDkimParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetDkimParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetDkimParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetDkimParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDmarcParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetDmarcParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetDmarcParamsDkim] `query:"dkim"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetDmarcParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetDmarcParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetDmarcParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetDmarcParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetDmarcParamsArc string

const (
	RadarEmailSecuritySummaryGetDmarcParamsArcPass RadarEmailSecuritySummaryGetDmarcParamsArc = "PASS"
	RadarEmailSecuritySummaryGetDmarcParamsArcNone RadarEmailSecuritySummaryGetDmarcParamsArc = "NONE"
	RadarEmailSecuritySummaryGetDmarcParamsArcFail RadarEmailSecuritySummaryGetDmarcParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDmarcParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcParamsArcPass, RadarEmailSecuritySummaryGetDmarcParamsArcNone, RadarEmailSecuritySummaryGetDmarcParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDmarcParamsDkim string

const (
	RadarEmailSecuritySummaryGetDmarcParamsDkimPass RadarEmailSecuritySummaryGetDmarcParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetDmarcParamsDkimNone RadarEmailSecuritySummaryGetDmarcParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetDmarcParamsDkimFail RadarEmailSecuritySummaryGetDmarcParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDmarcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcParamsDkimPass, RadarEmailSecuritySummaryGetDmarcParamsDkimNone, RadarEmailSecuritySummaryGetDmarcParamsDkimFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetDmarcParamsFormat string

const (
	RadarEmailSecuritySummaryGetDmarcParamsFormatJson RadarEmailSecuritySummaryGetDmarcParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetDmarcParamsFormatCsv  RadarEmailSecuritySummaryGetDmarcParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetDmarcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcParamsFormatJson, RadarEmailSecuritySummaryGetDmarcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDmarcParamsSpf string

const (
	RadarEmailSecuritySummaryGetDmarcParamsSpfPass RadarEmailSecuritySummaryGetDmarcParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetDmarcParamsSpfNone RadarEmailSecuritySummaryGetDmarcParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetDmarcParamsSpfFail RadarEmailSecuritySummaryGetDmarcParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetDmarcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcParamsSpfPass, RadarEmailSecuritySummaryGetDmarcParamsSpfNone, RadarEmailSecuritySummaryGetDmarcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetDmarcParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetDmarcParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetDmarcParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetDmarcParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetDmarcParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetDmarcParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetDmarcParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetMaliciousParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetMaliciousParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetMaliciousParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetMaliciousParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetMaliciousParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetMaliciousParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryGetMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetMaliciousParamsArc string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsArcPass RadarEmailSecuritySummaryGetMaliciousParamsArc = "PASS"
	RadarEmailSecuritySummaryGetMaliciousParamsArcNone RadarEmailSecuritySummaryGetMaliciousParamsArc = "NONE"
	RadarEmailSecuritySummaryGetMaliciousParamsArcFail RadarEmailSecuritySummaryGetMaliciousParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsArcPass, RadarEmailSecuritySummaryGetMaliciousParamsArcNone, RadarEmailSecuritySummaryGetMaliciousParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousParamsDkim string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsDkimPass RadarEmailSecuritySummaryGetMaliciousParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetMaliciousParamsDkimNone RadarEmailSecuritySummaryGetMaliciousParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetMaliciousParamsDkimFail RadarEmailSecuritySummaryGetMaliciousParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsDkimPass, RadarEmailSecuritySummaryGetMaliciousParamsDkimNone, RadarEmailSecuritySummaryGetMaliciousParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousParamsDmarc string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsDmarcPass RadarEmailSecuritySummaryGetMaliciousParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetMaliciousParamsDmarcNone RadarEmailSecuritySummaryGetMaliciousParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetMaliciousParamsDmarcFail RadarEmailSecuritySummaryGetMaliciousParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsDmarcPass, RadarEmailSecuritySummaryGetMaliciousParamsDmarcNone, RadarEmailSecuritySummaryGetMaliciousParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetMaliciousParamsFormat string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsFormatJson RadarEmailSecuritySummaryGetMaliciousParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetMaliciousParamsFormatCsv  RadarEmailSecuritySummaryGetMaliciousParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsFormatJson, RadarEmailSecuritySummaryGetMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousParamsSpf string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsSpfPass RadarEmailSecuritySummaryGetMaliciousParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetMaliciousParamsSpfNone RadarEmailSecuritySummaryGetMaliciousParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetMaliciousParamsSpfFail RadarEmailSecuritySummaryGetMaliciousParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsSpfPass, RadarEmailSecuritySummaryGetMaliciousParamsSpfNone, RadarEmailSecuritySummaryGetMaliciousParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetMaliciousParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetMaliciousParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetSpamParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetSpamParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetSpamParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetSpamParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetSpamParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetSpamParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetSpamParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetSpamParamsArc string

const (
	RadarEmailSecuritySummaryGetSpamParamsArcPass RadarEmailSecuritySummaryGetSpamParamsArc = "PASS"
	RadarEmailSecuritySummaryGetSpamParamsArcNone RadarEmailSecuritySummaryGetSpamParamsArc = "NONE"
	RadarEmailSecuritySummaryGetSpamParamsArcFail RadarEmailSecuritySummaryGetSpamParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpamParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsArcPass, RadarEmailSecuritySummaryGetSpamParamsArcNone, RadarEmailSecuritySummaryGetSpamParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamParamsDkim string

const (
	RadarEmailSecuritySummaryGetSpamParamsDkimPass RadarEmailSecuritySummaryGetSpamParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetSpamParamsDkimNone RadarEmailSecuritySummaryGetSpamParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetSpamParamsDkimFail RadarEmailSecuritySummaryGetSpamParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpamParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsDkimPass, RadarEmailSecuritySummaryGetSpamParamsDkimNone, RadarEmailSecuritySummaryGetSpamParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamParamsDmarc string

const (
	RadarEmailSecuritySummaryGetSpamParamsDmarcPass RadarEmailSecuritySummaryGetSpamParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetSpamParamsDmarcNone RadarEmailSecuritySummaryGetSpamParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetSpamParamsDmarcFail RadarEmailSecuritySummaryGetSpamParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpamParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsDmarcPass, RadarEmailSecuritySummaryGetSpamParamsDmarcNone, RadarEmailSecuritySummaryGetSpamParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetSpamParamsFormat string

const (
	RadarEmailSecuritySummaryGetSpamParamsFormatJson RadarEmailSecuritySummaryGetSpamParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetSpamParamsFormatCsv  RadarEmailSecuritySummaryGetSpamParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetSpamParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsFormatJson, RadarEmailSecuritySummaryGetSpamParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamParamsSpf string

const (
	RadarEmailSecuritySummaryGetSpamParamsSpfPass RadarEmailSecuritySummaryGetSpamParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetSpamParamsSpfNone RadarEmailSecuritySummaryGetSpamParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetSpamParamsSpfFail RadarEmailSecuritySummaryGetSpamParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpamParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsSpfPass, RadarEmailSecuritySummaryGetSpamParamsSpfNone, RadarEmailSecuritySummaryGetSpamParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpamParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetSpamParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetSpamParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetSpamParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetSpamParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetSpamParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetSpamParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpfParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetSpfParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetSpfParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetSpfParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetSpfParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetSpfParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetSpfParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetSpfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetSpfParamsArc string

const (
	RadarEmailSecuritySummaryGetSpfParamsArcPass RadarEmailSecuritySummaryGetSpfParamsArc = "PASS"
	RadarEmailSecuritySummaryGetSpfParamsArcNone RadarEmailSecuritySummaryGetSpfParamsArc = "NONE"
	RadarEmailSecuritySummaryGetSpfParamsArcFail RadarEmailSecuritySummaryGetSpfParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpfParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfParamsArcPass, RadarEmailSecuritySummaryGetSpfParamsArcNone, RadarEmailSecuritySummaryGetSpfParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpfParamsDkim string

const (
	RadarEmailSecuritySummaryGetSpfParamsDkimPass RadarEmailSecuritySummaryGetSpfParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetSpfParamsDkimNone RadarEmailSecuritySummaryGetSpfParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetSpfParamsDkimFail RadarEmailSecuritySummaryGetSpfParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpfParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfParamsDkimPass, RadarEmailSecuritySummaryGetSpfParamsDkimNone, RadarEmailSecuritySummaryGetSpfParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpfParamsDmarc string

const (
	RadarEmailSecuritySummaryGetSpfParamsDmarcPass RadarEmailSecuritySummaryGetSpfParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetSpfParamsDmarcNone RadarEmailSecuritySummaryGetSpfParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetSpfParamsDmarcFail RadarEmailSecuritySummaryGetSpfParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpfParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfParamsDmarcPass, RadarEmailSecuritySummaryGetSpfParamsDmarcNone, RadarEmailSecuritySummaryGetSpfParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetSpfParamsFormat string

const (
	RadarEmailSecuritySummaryGetSpfParamsFormatJson RadarEmailSecuritySummaryGetSpfParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetSpfParamsFormatCsv  RadarEmailSecuritySummaryGetSpfParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetSpfParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfParamsFormatJson, RadarEmailSecuritySummaryGetSpfParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpfParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetSpfParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetSpfParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetSpfParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetSpfParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetSpfParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetSpfParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetSpoofParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetSpoofParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetSpoofParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetSpoofParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetSpoofParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetSpoofParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetSpoofParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryGetSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetSpoofParamsArc string

const (
	RadarEmailSecuritySummaryGetSpoofParamsArcPass RadarEmailSecuritySummaryGetSpoofParamsArc = "PASS"
	RadarEmailSecuritySummaryGetSpoofParamsArcNone RadarEmailSecuritySummaryGetSpoofParamsArc = "NONE"
	RadarEmailSecuritySummaryGetSpoofParamsArcFail RadarEmailSecuritySummaryGetSpoofParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsArcPass, RadarEmailSecuritySummaryGetSpoofParamsArcNone, RadarEmailSecuritySummaryGetSpoofParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofParamsDkim string

const (
	RadarEmailSecuritySummaryGetSpoofParamsDkimPass RadarEmailSecuritySummaryGetSpoofParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetSpoofParamsDkimNone RadarEmailSecuritySummaryGetSpoofParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetSpoofParamsDkimFail RadarEmailSecuritySummaryGetSpoofParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsDkimPass, RadarEmailSecuritySummaryGetSpoofParamsDkimNone, RadarEmailSecuritySummaryGetSpoofParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofParamsDmarc string

const (
	RadarEmailSecuritySummaryGetSpoofParamsDmarcPass RadarEmailSecuritySummaryGetSpoofParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetSpoofParamsDmarcNone RadarEmailSecuritySummaryGetSpoofParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetSpoofParamsDmarcFail RadarEmailSecuritySummaryGetSpoofParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsDmarcPass, RadarEmailSecuritySummaryGetSpoofParamsDmarcNone, RadarEmailSecuritySummaryGetSpoofParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetSpoofParamsFormat string

const (
	RadarEmailSecuritySummaryGetSpoofParamsFormatJson RadarEmailSecuritySummaryGetSpoofParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetSpoofParamsFormatCsv  RadarEmailSecuritySummaryGetSpoofParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsFormatJson, RadarEmailSecuritySummaryGetSpoofParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofParamsSpf string

const (
	RadarEmailSecuritySummaryGetSpoofParamsSpfPass RadarEmailSecuritySummaryGetSpoofParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetSpoofParamsSpfNone RadarEmailSecuritySummaryGetSpoofParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetSpoofParamsSpfFail RadarEmailSecuritySummaryGetSpoofParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsSpfPass, RadarEmailSecuritySummaryGetSpoofParamsSpfNone, RadarEmailSecuritySummaryGetSpoofParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetSpoofParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetSpoofParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetSpoofParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetSpoofParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetSpoofParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetSpoofParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetSpoofParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetThreatCategoryParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetThreatCategoryParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetThreatCategoryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetThreatCategoryParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetThreatCategoryParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryGetThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetThreatCategoryParamsArc string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsArcPass RadarEmailSecuritySummaryGetThreatCategoryParamsArc = "PASS"
	RadarEmailSecuritySummaryGetThreatCategoryParamsArcNone RadarEmailSecuritySummaryGetThreatCategoryParamsArc = "NONE"
	RadarEmailSecuritySummaryGetThreatCategoryParamsArcFail RadarEmailSecuritySummaryGetThreatCategoryParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsArcPass, RadarEmailSecuritySummaryGetThreatCategoryParamsArcNone, RadarEmailSecuritySummaryGetThreatCategoryParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryParamsDkim string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsDkimPass RadarEmailSecuritySummaryGetThreatCategoryParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetThreatCategoryParamsDkimNone RadarEmailSecuritySummaryGetThreatCategoryParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetThreatCategoryParamsDkimFail RadarEmailSecuritySummaryGetThreatCategoryParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsDkimPass, RadarEmailSecuritySummaryGetThreatCategoryParamsDkimNone, RadarEmailSecuritySummaryGetThreatCategoryParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcPass RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcNone RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcFail RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcPass, RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcNone, RadarEmailSecuritySummaryGetThreatCategoryParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetThreatCategoryParamsFormat string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsFormatJson RadarEmailSecuritySummaryGetThreatCategoryParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetThreatCategoryParamsFormatCsv  RadarEmailSecuritySummaryGetThreatCategoryParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsFormatJson, RadarEmailSecuritySummaryGetThreatCategoryParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryParamsSpf string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsSpfPass RadarEmailSecuritySummaryGetThreatCategoryParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetThreatCategoryParamsSpfNone RadarEmailSecuritySummaryGetThreatCategoryParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetThreatCategoryParamsSpfFail RadarEmailSecuritySummaryGetThreatCategoryParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsSpfPass, RadarEmailSecuritySummaryGetThreatCategoryParamsSpfNone, RadarEmailSecuritySummaryGetThreatCategoryParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion string

const (
	RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_0 RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion = "TLSv1_0"
	RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_1 RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion = "TLSv1_1"
	RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_2 RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion = "TLSv1_2"
	RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_3 RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_0, RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_1, RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_2, RadarEmailSecuritySummaryGetThreatCategoryParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetTlsVersionParams struct {
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecuritySummaryGetTlsVersionParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecuritySummaryGetTlsVersionParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecuritySummaryGetTlsVersionParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecuritySummaryGetTlsVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecuritySummaryGetTlsVersionParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryGetTlsVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryGetTlsVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryGetTlsVersionParamsArc string

const (
	RadarEmailSecuritySummaryGetTlsVersionParamsArcPass RadarEmailSecuritySummaryGetTlsVersionParamsArc = "PASS"
	RadarEmailSecuritySummaryGetTlsVersionParamsArcNone RadarEmailSecuritySummaryGetTlsVersionParamsArc = "NONE"
	RadarEmailSecuritySummaryGetTlsVersionParamsArcFail RadarEmailSecuritySummaryGetTlsVersionParamsArc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetTlsVersionParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionParamsArcPass, RadarEmailSecuritySummaryGetTlsVersionParamsArcNone, RadarEmailSecuritySummaryGetTlsVersionParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetTlsVersionParamsDkim string

const (
	RadarEmailSecuritySummaryGetTlsVersionParamsDkimPass RadarEmailSecuritySummaryGetTlsVersionParamsDkim = "PASS"
	RadarEmailSecuritySummaryGetTlsVersionParamsDkimNone RadarEmailSecuritySummaryGetTlsVersionParamsDkim = "NONE"
	RadarEmailSecuritySummaryGetTlsVersionParamsDkimFail RadarEmailSecuritySummaryGetTlsVersionParamsDkim = "FAIL"
)

func (r RadarEmailSecuritySummaryGetTlsVersionParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionParamsDkimPass, RadarEmailSecuritySummaryGetTlsVersionParamsDkimNone, RadarEmailSecuritySummaryGetTlsVersionParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetTlsVersionParamsDmarc string

const (
	RadarEmailSecuritySummaryGetTlsVersionParamsDmarcPass RadarEmailSecuritySummaryGetTlsVersionParamsDmarc = "PASS"
	RadarEmailSecuritySummaryGetTlsVersionParamsDmarcNone RadarEmailSecuritySummaryGetTlsVersionParamsDmarc = "NONE"
	RadarEmailSecuritySummaryGetTlsVersionParamsDmarcFail RadarEmailSecuritySummaryGetTlsVersionParamsDmarc = "FAIL"
)

func (r RadarEmailSecuritySummaryGetTlsVersionParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionParamsDmarcPass, RadarEmailSecuritySummaryGetTlsVersionParamsDmarcNone, RadarEmailSecuritySummaryGetTlsVersionParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecuritySummaryGetTlsVersionParamsFormat string

const (
	RadarEmailSecuritySummaryGetTlsVersionParamsFormatJson RadarEmailSecuritySummaryGetTlsVersionParamsFormat = "JSON"
	RadarEmailSecuritySummaryGetTlsVersionParamsFormatCsv  RadarEmailSecuritySummaryGetTlsVersionParamsFormat = "CSV"
)

func (r RadarEmailSecuritySummaryGetTlsVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionParamsFormatJson, RadarEmailSecuritySummaryGetTlsVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecuritySummaryGetTlsVersionParamsSpf string

const (
	RadarEmailSecuritySummaryGetTlsVersionParamsSpfPass RadarEmailSecuritySummaryGetTlsVersionParamsSpf = "PASS"
	RadarEmailSecuritySummaryGetTlsVersionParamsSpfNone RadarEmailSecuritySummaryGetTlsVersionParamsSpf = "NONE"
	RadarEmailSecuritySummaryGetTlsVersionParamsSpfFail RadarEmailSecuritySummaryGetTlsVersionParamsSpf = "FAIL"
)

func (r RadarEmailSecuritySummaryGetTlsVersionParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecuritySummaryGetTlsVersionParamsSpfPass, RadarEmailSecuritySummaryGetTlsVersionParamsSpfNone, RadarEmailSecuritySummaryGetTlsVersionParamsSpfFail:
		return true
	}
	return false
}
