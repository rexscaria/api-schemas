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

// RadarEmailSecurityTimeseriesGroupService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTimeseriesGroupService] method instead.
type RadarEmailSecurityTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailSecurityTimeseriesGroupService) {
	r = &RadarEmailSecurityTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetArc(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetArcParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetDkim(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetDkimParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetDkimResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetDmarc(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetDmarcParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by malicious classification over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetMalicious(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetMaliciousParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetMaliciousResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by spam classification (spam vs. non-spam)
// over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetSpam(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetSpamParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetSpamResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetSpf(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetSpfParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetSpfResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by spoof classification (spoof vs.
// non-spoof) over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetSpoof(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetSpoofParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetSpoofResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/spoof"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by threat category over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetThreatCategory(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetThreatCategoryParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by TLS version over time.
func (r *RadarEmailSecurityTimeseriesGroupService) GetTlsVersion(ctx context.Context, query RadarEmailSecurityTimeseriesGroupGetTlsVersionParams, opts ...option.RequestOption) (res *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTimeseriesGroupGetArcResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetArcResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetArcResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetArcResponse]
type radarEmailSecurityTimeseriesGroupGetArcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetArcResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetArcResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetArcResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetArcResponseResult]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                       `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0 struct {
	Fail []string                                                        `json:"FAIL,required"`
	None []string                                                        `json:"NONE,required"`
	Pass []string                                                        `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetArcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetDkimResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetDkimResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetDkimResponse]
type radarEmailSecurityTimeseriesGroupGetDkimResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetDkimResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetDkimResponseResult]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                        `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnit struct {
	Name  string                                                             `json:"name,required"`
	Value string                                                             `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0 struct {
	Fail []string                                                         `json:"FAIL,required"`
	None []string                                                         `json:"NONE,required"`
	Pass []string                                                         `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDkimResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetDmarcResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetDmarcResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetDmarcResponse]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetDmarcResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResult]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                         `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnit struct {
	Name  string                                                              `json:"name,required"`
	Value string                                                              `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0 struct {
	Fail []string                                                          `json:"FAIL,required"`
	None []string                                                          `json:"NONE,required"`
	Pass []string                                                          `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetDmarcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResult `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetMaliciousResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetMaliciousResponse]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResult]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                             `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                   `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnit struct {
	Name  string                                                                  `json:"name,required"`
	Value string                                                                  `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0 struct {
	Malicious    []string                                                              `json:"MALICIOUS,required"`
	NotMalicious []string                                                              `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetMaliciousResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetSpamResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetSpamResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetSpamResponse]
type radarEmailSecurityTimeseriesGroupGetSpamResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetSpamResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetSpamResponseResult]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                        `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnit struct {
	Name  string                                                             `json:"name,required"`
	Value string                                                             `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0 struct {
	NotSpam []string                                                         `json:"NOT_SPAM,required"`
	Spam    []string                                                         `json:"SPAM,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpamResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetSpfResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetSpfResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetSpfResponse]
type radarEmailSecurityTimeseriesGroupGetSpfResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetSpfResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetSpfResponseResult]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                       `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0 struct {
	Fail []string                                                        `json:"FAIL,required"`
	None []string                                                        `json:"NONE,required"`
	Pass []string                                                        `json:"PASS,required"`
	JSON radarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpfResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetSpoofResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetSpoofResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTimeseriesGroupGetSpoofResponse]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetSpoofResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResult]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                         `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnit struct {
	Name  string                                                              `json:"name,required"`
	Value string                                                              `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0 struct {
	NotSpoof []string                                                          `json:"NOT_SPOOF,required"`
	Spoof    []string                                                          `json:"SPOOF,required"`
	JSON     radarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0JSON struct {
	NotSpoof    apijson.Field
	Spoof       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetSpoofResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResult `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponse]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResult]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                  `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                        `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnit struct {
	Name  string                                                                       `json:"name,required"`
	Value string                                                                       `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0 struct {
	BrandImpersonation  []string                                                                   `json:"BrandImpersonation,required"`
	CredentialHarvester []string                                                                   `json:"CredentialHarvester,required"`
	IdentityDeception   []string                                                                   `json:"IdentityDeception,required"`
	Link                []string                                                                   `json:"Link,required"`
	JSON                radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetThreatCategoryResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponse struct {
	Result  RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarEmailSecurityTimeseriesGroupGetTlsVersionResponseJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponse]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResult]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaJSON   `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMeta]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalFifteenMinutes RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneHour        RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneDay         RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneWeek        RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneMonth       RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalFifteenMinutes, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneHour, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneDay, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneWeek, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                              `json:"level,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfo]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRange]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationPercentage           RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationMin0Max              RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationMinMax               RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRawValues            RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationPercentageChange     RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRollingAverage       RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationOverlappedPercentage RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRatio                RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationPercentage, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationMin0Max, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationMinMax, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRawValues, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationPercentageChange, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRollingAverage, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationOverlappedPercentage, RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnit struct {
	Name  string                                                                   `json:"name,required"`
	Value string                                                                   `json:"value,required"`
	JSON  radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnit]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0 struct {
	Tls1_0 []string                                                               `json:"TLS 1.0,required"`
	Tls1_1 []string                                                               `json:"TLS 1.1,required"`
	Tls1_2 []string                                                               `json:"TLS 1.2,required"`
	Tls1_3 []string                                                               `json:"TLS 1.3,required"`
	JSON   radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0JSON `json:"-"`
}

// radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0]
type radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0JSON struct {
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSecurityTimeseriesGroupGetTlsVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailSecurityTimeseriesGroupGetArcParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetArcParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetArcParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetArcParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetArcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetArcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsDkimPass RadarEmailSecurityTimeseriesGroupGetArcParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetArcParamsDkimNone RadarEmailSecurityTimeseriesGroupGetArcParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetArcParamsDkimFail RadarEmailSecurityTimeseriesGroupGetArcParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetArcParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetArcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetArcParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetArcParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsFormatJson RadarEmailSecurityTimeseriesGroupGetArcParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetArcParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetArcParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetArcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsSpfPass RadarEmailSecurityTimeseriesGroupGetArcParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetArcParamsSpfNone RadarEmailSecurityTimeseriesGroupGetArcParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetArcParamsSpfFail RadarEmailSecurityTimeseriesGroupGetArcParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetArcParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetArcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetArcParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetDkimParamsArc] `query:"arc"`
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
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetDkimParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetDkimParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetDkimParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetDkimParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsArcPass RadarEmailSecurityTimeseriesGroupGetDkimParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsArcNone RadarEmailSecurityTimeseriesGroupGetDkimParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsArcFail RadarEmailSecurityTimeseriesGroupGetDkimParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsArcPass, RadarEmailSecurityTimeseriesGroupGetDkimParamsArcNone, RadarEmailSecurityTimeseriesGroupGetDkimParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetDkimParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetDkimParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsFormatJson RadarEmailSecurityTimeseriesGroupGetDkimParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetDkimParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetDkimParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfPass RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfNone RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfFail RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetDkimParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetDkimParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim] `query:"dkim"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetDmarcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetDmarcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcPass RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcNone RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcFail RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcPass, RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcNone, RadarEmailSecurityTimeseriesGroupGetDmarcParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimPass RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimNone RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimFail RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetDmarcParamsDkimFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormatJson RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetDmarcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfPass RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfNone RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfFail RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetDmarcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetDmarcParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetMaliciousParams]'s
// query parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcPass RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcNone RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcFail RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcPass, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcNone, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimPass RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimNone RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimFail RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormatJson RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfPass RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfNone RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfFail RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetMaliciousParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpamParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetSpamParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetSpamParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetSpamParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetSpamParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsArcPass RadarEmailSecurityTimeseriesGroupGetSpamParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsArcNone RadarEmailSecurityTimeseriesGroupGetSpamParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsArcFail RadarEmailSecurityTimeseriesGroupGetSpamParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsArcPass, RadarEmailSecurityTimeseriesGroupGetSpamParamsArcNone, RadarEmailSecurityTimeseriesGroupGetSpamParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimPass RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimNone RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimFail RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetSpamParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetSpamParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetSpamParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsFormatJson RadarEmailSecurityTimeseriesGroupGetSpamParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetSpamParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetSpamParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfPass RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfNone RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfFail RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetSpamParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetSpamParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpfParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetSpfParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetSpfParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetSpfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetSpfParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsArcPass RadarEmailSecurityTimeseriesGroupGetSpfParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsArcNone RadarEmailSecurityTimeseriesGroupGetSpfParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsArcFail RadarEmailSecurityTimeseriesGroupGetSpfParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsArcPass, RadarEmailSecurityTimeseriesGroupGetSpfParamsArcNone, RadarEmailSecurityTimeseriesGroupGetSpfParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimPass RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimNone RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimFail RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetSpfParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetSpfParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetSpfParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsFormatJson RadarEmailSecurityTimeseriesGroupGetSpfParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetSpfParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetSpfParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetSpfParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetSpoofParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetSpoofParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetSpoofParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcPass RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcNone RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcFail RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcPass, RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcNone, RadarEmailSecurityTimeseriesGroupGetSpoofParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimPass RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimNone RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimFail RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetSpoofParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetSpoofParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormatJson RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetSpoofParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfPass RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfNone RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfFail RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetSpoofParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetSpoofParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf] `query:"spf"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetThreatCategoryParams]'s
// query parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcPass RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcNone RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcFail RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcPass, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcNone, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimPass RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimNone RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimFail RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormatJson RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfPass RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfNone RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfFail RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion string

const (
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_0 RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion = "TLSv1_0"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_1 RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion = "TLSv1_1"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_2 RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion = "TLSv1_2"
	RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_3 RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion = "TLSv1_3"
)

func (r RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_0, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_1, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_2, RadarEmailSecurityTimeseriesGroupGetThreatCategoryParamsTlsVersionTlSv1_3:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTimeseriesGroupGetTlsVersionParams]'s
// query parameters as `url.Values`.
func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval15m RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval = "15m"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1h  RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval = "1h"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1d  RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval = "1d"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1w  RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval = "1w"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval15m, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1h, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1d, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcPass RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcNone RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcFail RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcPass, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcNone, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsArcFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimPass RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim = "PASS"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimNone RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim = "NONE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimFail RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimPass, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimNone, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcPass RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc = "PASS"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcNone RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc = "NONE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcFail RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcPass, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcNone, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormat string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormatJson RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormat = "JSON"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormatCsv  RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormat = "CSV"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormatJson, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf string

const (
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfPass RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf = "PASS"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfNone RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf = "NONE"
	RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfFail RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf = "FAIL"
)

func (r RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfPass, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfNone, RadarEmailSecurityTimeseriesGroupGetTlsVersionParamsSpfFail:
		return true
	}
	return false
}
