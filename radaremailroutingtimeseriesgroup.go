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

// RadarEmailRoutingTimeseriesGroupService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailRoutingTimeseriesGroupService] method instead.
type RadarEmailRoutingTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarEmailRoutingTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailRoutingTimeseriesGroupService(opts ...option.RequestOption) (r *RadarEmailRoutingTimeseriesGroupService) {
	r = &RadarEmailRoutingTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of emails by ARC (Authenticated Received Chain)
// validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetArc(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetArcParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetArcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DKIM (DomainKeys Identified Mail)
// validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetDkim(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetDkimParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetDkimResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by DMARC (Domain-based Message
// Authentication, Reporting and Conformance) validation over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetDmarc(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetDmarcParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetDmarcResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by encryption status (encrypted vs.
// not-encrypted) over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetEncrypted(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetEncryptedParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetEncryptedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/encrypted"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by IP version over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetIPVersion(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetIPVersionParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of emails by SPF (Sender Policy Framework) validation
// over time.
func (r *RadarEmailRoutingTimeseriesGroupService) GetSpf(ctx context.Context, query RadarEmailRoutingTimeseriesGroupGetSpfParams, opts ...option.RequestOption) (res *RadarEmailRoutingTimeseriesGroupGetSpfResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/routing/timeseries_groups/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailRoutingTimeseriesGroupGetArcResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetArcResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetArcResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetArcResponse]
type radarEmailRoutingTimeseriesGroupGetArcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetArcResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetArcResponseResult]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0 struct {
	Fail []string                                                       `json:"FAIL,required"`
	None []string                                                       `json:"NONE,required"`
	Pass []string                                                       `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetArcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetDkimResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetDkimResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetDkimResponse]
type radarEmailRoutingTimeseriesGroupGetDkimResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetDkimResponseResult]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                       `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0 struct {
	Fail []string                                                        `json:"FAIL,required"`
	None []string                                                        `json:"NONE,required"`
	Pass []string                                                        `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDkimResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetDmarcResponse]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                        `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnit struct {
	Name  string                                                             `json:"name,required"`
	Value string                                                             `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0 struct {
	Fail []string                                                         `json:"FAIL,required"`
	None []string                                                         `json:"NONE,required"`
	Pass []string                                                         `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetDmarcResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetEncryptedResponse]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                            `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                  `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnit struct {
	Name  string                                                                 `json:"name,required"`
	Value string                                                                 `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0 struct {
	Encrypted    []string                                                             `json:"ENCRYPTED,required"`
	NotEncrypted []string                                                             `json:"NOT_ENCRYPTED,required"`
	JSON         radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON struct {
	Encrypted    apijson.Field
	NotEncrypted apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetEncryptedResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetIPVersionResponse]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                            `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                  `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnit struct {
	Name  string                                                                 `json:"name,required"`
	Value string                                                                 `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0 struct {
	IPv4 []string                                                             `json:"IPv4,required"`
	IPv6 []string                                                             `json:"IPv6,required"`
	JSON radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetIPVersionResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponse struct {
	Result  RadarEmailRoutingTimeseriesGroupGetSpfResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailRoutingTimeseriesGroupGetSpfResponseJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseJSON contains the JSON metadata
// for the struct [RadarEmailRoutingTimeseriesGroupGetSpfResponse]
type radarEmailRoutingTimeseriesGroupGetSpfResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResult struct {
	// Metadata for the results.
	Meta   RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMeta   `json:"meta,required"`
	Serie0 RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailRoutingTimeseriesGroupGetSpfResponseResult]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnit `json:"units,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaJSON   `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMeta]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalFifteenMinutes RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneHour        RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_HOUR"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneDay         RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_DAY"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneWeek        RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_WEEK"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneMonth       RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalFifteenMinutes, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneHour, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneDay, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneWeek, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRange]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentage           RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationMin0Max              RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "MIN0_MAX"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationMinMax               RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "MIN_MAX"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRawValues            RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "RAW_VALUES"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentageChange     RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRollingAverage       RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationOverlappedPercentage RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRatio                RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization = "RATIO"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentage, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationMin0Max, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationMinMax, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRawValues, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationPercentageChange, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRollingAverage, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationOverlappedPercentage, RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnitJSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnit]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0 struct {
	Fail []string                                                       `json:"FAIL,required"`
	None []string                                                       `json:"NONE,required"`
	Pass []string                                                       `json:"PASS,required"`
	JSON radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON `json:"-"`
}

// radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0]
type radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailRoutingTimeseriesGroupGetSpfResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarEmailRoutingTimeseriesGroupGetArcParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetArcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetArcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetArcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetArcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetArcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimPass RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimNone RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDkimFail RadarEmailRoutingTimeseriesGroupGetArcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetArcParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetArcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetArcParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetArcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetArcParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsFormatJson RadarEmailRoutingTimeseriesGroupGetArcParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetArcParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetArcParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetArcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetArcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetArcParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfPass RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfNone RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetArcParamsSpfFail RadarEmailRoutingTimeseriesGroupGetArcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetArcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetArcParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetArcParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetArcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsArc] `query:"arc"`
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
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetDkimParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetDkimParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetDkimParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcPass RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcNone RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsArcFail RadarEmailRoutingTimeseriesGroupGetDkimParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsArcPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsArcNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetDkimParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatJson RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetDkimParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetDkimParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfPass RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfNone RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfFail RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDkimParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetDkimParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim] `query:"dkim"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetDmarcParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetDmarcParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetDmarcParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetDmarcParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatJson RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetDmarcParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetDmarcParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfPass RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfNone RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfFail RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetDmarcParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc] `query:"dmarc"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetEncryptedParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsDmarcFail:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatJson RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfPass RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfNone RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfFail RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetEncryptedParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by SPF (Sender Policy Framework) validation status.
	Spf param.Field[[]RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatJson RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf string

const (
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfPass RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "PASS"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfNone RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "NONE"
	RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfFail RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpf) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfPass, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfNone, RadarEmailRoutingTimeseriesGroupGetIPVersionParamsSpfFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval] `query:"aggInterval"`
	// Filters results by ARC (Authenticated Received Chain) validation.
	Arc param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsArc] `query:"arc"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by DKIM (DomainKeys Identified Mail) validation status.
	Dkim param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim] `query:"dkim"`
	// Filters results by DMARC (Domain-based Message Authentication, Reporting and
	// Conformance) validation status.
	Dmarc param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc] `query:"dmarc"`
	// Filters results by encryption status (encrypted vs. not-encrypted).
	Encrypted param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted] `query:"encrypted"`
	// Format in which results will be returned.
	Format param.Field[RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion] `query:"ipVersion"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailRoutingTimeseriesGroupGetSpfParams]'s query
// parameters as `url.Values`.
func (r RadarEmailRoutingTimeseriesGroupGetSpfParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval15m RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "15m"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1h  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1h"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1d  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1d"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1w  RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval = "1w"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval15m, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1h, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1d, RadarEmailRoutingTimeseriesGroupGetSpfParamsAggInterval1w:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsArc string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcPass RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcNone RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsArcFail RadarEmailRoutingTimeseriesGroupGetSpfParamsArc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsArc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsArcPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsArcNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsArcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimPass RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimNone RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimFail RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsDkim) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsDkimFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcPass RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "PASS"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcNone RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "NONE"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcFail RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc = "FAIL"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarc) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcPass, RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcNone, RadarEmailRoutingTimeseriesGroupGetSpfParamsDmarcFail:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedEncrypted    RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted = "ENCRYPTED"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedNotEncrypted RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted = "NOT_ENCRYPTED"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsEncrypted) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedEncrypted, RadarEmailRoutingTimeseriesGroupGetSpfParamsEncryptedNotEncrypted:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatJson RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat = "JSON"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatCsv  RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat = "CSV"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsFormat) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatJson, RadarEmailRoutingTimeseriesGroupGetSpfParamsFormatCsv:
		return true
	}
	return false
}

type RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion string

const (
	RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv4 RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion = "IPv4"
	RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv6 RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion = "IPv6"
)

func (r RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv4, RadarEmailRoutingTimeseriesGroupGetSpfParamsIPVersionIPv6:
		return true
	}
	return false
}
