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

// RadarLeakedCredentialCheckTimeseriesGroupService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarLeakedCredentialCheckTimeseriesGroupService] method instead.
type RadarLeakedCredentialCheckTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarLeakedCredentialCheckTimeseriesGroupService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarLeakedCredentialCheckTimeseriesGroupService(opts ...option.RequestOption) (r *RadarLeakedCredentialCheckTimeseriesGroupService) {
	r = &RadarLeakedCredentialCheckTimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves the distribution of HTTP authentication requests by bot class over
// time.
func (r *RadarLeakedCredentialCheckTimeseriesGroupService) GetByBotClass(ctx context.Context, query RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParams, opts ...option.RequestOption) (res *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of HTTP authentication requests by compromised
// credential status over time.
func (r *RadarLeakedCredentialCheckTimeseriesGroupService) GetByCompromisedStatus(ctx context.Context, query RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParams, opts ...option.RequestOption) (res *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/leaked_credential_checks/timeseries_groups/compromised"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponse struct {
	Result  RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResult `json:"result,required"`
	Success bool                                                                 `json:"success,required"`
	JSON    radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseJSON contains the
// JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponse]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResult struct {
	// Metadata for the results.
	Meta   RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMeta   `json:"meta,required"`
	Serie0 RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResult]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnit `json:"units,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMeta]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalFifteenMinutes RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneHour        RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval = "ONE_HOUR"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneDay         RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval = "ONE_DAY"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneWeek        RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval = "ONE_WEEK"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneMonth       RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalFifteenMinutes, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneHour, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneDay, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneWeek, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfo struct {
	Annotations []RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                      `json:"level,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfo]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotation]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRangeJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRange]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationPercentage           RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "PERCENTAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationMin0Max              RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "MIN0_MAX"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationMinMax               RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "MIN_MAX"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRawValues            RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "RAW_VALUES"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationPercentageChange     RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRollingAverage       RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationOverlappedPercentage RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRatio                RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization = "RATIO"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationPercentage, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationMin0Max, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationMinMax, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRawValues, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationPercentageChange, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRollingAverage, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationOverlappedPercentage, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnit struct {
	Name  string                                                                           `json:"name,required"`
	Value string                                                                           `json:"value,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnitJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnit]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0 struct {
	Bot        []string                                                                       `json:"bot,required"`
	Human      []string                                                                       `json:"human,required"`
	Timestamps []time.Time                                                                    `json:"timestamps,required" format:"date-time"`
	JSON       radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0JSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0]
type radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByBotClassResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponse struct {
	Result  RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResult `json:"result,required"`
	Success bool                                                                          `json:"success,required"`
	JSON    radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponse]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResult struct {
	// Metadata for the results.
	Meta   RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMeta   `json:"meta,required"`
	Serie0 RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResult]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnit `json:"units,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaJSON   `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMeta]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalFifteenMinutes RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneHour        RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval = "ONE_HOUR"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneDay         RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval = "ONE_DAY"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneWeek        RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval = "ONE_WEEK"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneMonth       RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalFifteenMinutes, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneHour, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneDay, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneWeek, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfo struct {
	Annotations []RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                               `json:"level,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfo]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                                     `json:"startDate,required" format:"date-time"`
	JSON            radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRangeJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRange]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationPercentage           RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "PERCENTAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationMin0Max              RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "MIN0_MAX"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationMinMax               RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "MIN_MAX"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRawValues            RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "RAW_VALUES"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationPercentageChange     RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRollingAverage       RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationOverlappedPercentage RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRatio                RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization = "RATIO"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationPercentage, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationMin0Max, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationMinMax, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRawValues, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationPercentageChange, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRollingAverage, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationOverlappedPercentage, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnit struct {
	Name  string                                                                                    `json:"name,required"`
	Value string                                                                                    `json:"value,required"`
	JSON  radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnitJSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnit]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0 struct {
	Clean       []string                                                                                `json:"CLEAN,required"`
	Compromised []string                                                                                `json:"COMPROMISED,required"`
	Timestamps  []time.Time                                                                             `json:"timestamps,required" format:"date-time"`
	JSON        radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0JSON `json:"-"`
}

// radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0]
type radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0JSON struct {
	Clean       apijson.Field
	Compromised apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval] `query:"aggInterval"`
	// Filters results by compromised credential status (clean vs. compromised).
	Compromised param.Field[[]RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromised] `query:"compromised"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes
// [RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParams]'s query
// parameters as `url.Values`.
func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval15m RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval = "15m"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1h  RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval = "1h"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1d  RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval = "1d"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1w  RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval = "1w"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval15m, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1h, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1d, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsAggInterval1w:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromised string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromisedClean       RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromised = "CLEAN"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromisedCompromised RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromised = "COMPROMISED"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromised) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromisedClean, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsCompromisedCompromised:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormat string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormatJson RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormat = "JSON"
	RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormatCsv  RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormat = "CSV"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormat) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormatJson, RadarLeakedCredentialCheckTimeseriesGroupGetByBotClassParamsFormatCsv:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval] `query:"aggInterval"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes
// [RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParams]'s query
// parameters as `url.Values`.
func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval15m RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval = "15m"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1h  RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval = "1h"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1d  RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval = "1d"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1w  RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval = "1w"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval15m, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1h, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1d, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsAggInterval1w:
		return true
	}
	return false
}

type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClass string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClassLikelyAutomated RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClass = "LIKELY_AUTOMATED"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClassLikelyHuman     RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClass) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClassLikelyAutomated, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsBotClassLikelyHuman:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormat string

const (
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormatJson RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormat = "JSON"
	RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormatCsv  RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormat = "CSV"
)

func (r RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormat) IsKnown() bool {
	switch r {
	case RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormatJson, RadarLeakedCredentialCheckTimeseriesGroupGetByCompromisedStatusParamsFormatCsv:
		return true
	}
	return false
}
