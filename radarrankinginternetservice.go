// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// RadarRankingInternetServiceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarRankingInternetServiceService] method instead.
type RadarRankingInternetServiceService struct {
	Options []option.RequestOption
}

// NewRadarRankingInternetServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarRankingInternetServiceService(opts ...option.RequestOption) (r *RadarRankingInternetServiceService) {
	r = &RadarRankingInternetServiceService{}
	r.Options = opts
	return
}

// Retrieves the list of Internet services categories.
func (r *RadarRankingInternetServiceService) ListCategories(ctx context.Context, query RadarRankingInternetServiceListCategoriesParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceListCategoriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves Internet Services rank update changes over time.
func (r *RadarRankingInternetServiceService) GetTimeseriesGroups(ctx context.Context, query RadarRankingInternetServiceGetTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceGetTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves top Internet services based on their rank.
func (r *RadarRankingInternetServiceService) GetTopServices(ctx context.Context, query RadarRankingInternetServiceGetTopServicesParams, opts ...option.RequestOption) (res *RadarRankingInternetServiceGetTopServicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarRankingInternetServiceListCategoriesResponse struct {
	Result  RadarRankingInternetServiceListCategoriesResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarRankingInternetServiceListCategoriesResponseJSON   `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseJSON contains the JSON metadata
// for the struct [RadarRankingInternetServiceListCategoriesResponse]
type radarRankingInternetServiceListCategoriesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesResponseResult struct {
	Categories0 []RadarRankingInternetServiceListCategoriesResponseResultCategories0 `json:"categories_0,required"`
	JSON        radarRankingInternetServiceListCategoriesResponseResultJSON          `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseResultJSON contains the JSON
// metadata for the struct
// [RadarRankingInternetServiceListCategoriesResponseResult]
type radarRankingInternetServiceListCategoriesResponseResultJSON struct {
	Categories0 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesResponseResultCategories0 struct {
	Name string                                                                 `json:"name,required"`
	JSON radarRankingInternetServiceListCategoriesResponseResultCategories0JSON `json:"-"`
}

// radarRankingInternetServiceListCategoriesResponseResultCategories0JSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceListCategoriesResponseResultCategories0]
type radarRankingInternetServiceListCategoriesResponseResultCategories0JSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceListCategoriesResponseResultCategories0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceListCategoriesResponseResultCategories0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponse struct {
	Result  RadarRankingInternetServiceGetTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarRankingInternetServiceGetTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct [RadarRankingInternetServiceGetTimeseriesGroupsResponse]
type radarRankingInternetServiceGetTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResult struct {
	// Metadata for the results.
	Meta   RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta   `json:"meta,required"`
	Serie0 RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResult]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnit `json:"units,required"`
	JSON  radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON   `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval string

const (
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval = "FIFTEEN_MINUTES"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneHour        RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_HOUR"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneDay         RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_DAY"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneWeek        RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_WEEK"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneMonth       RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval = "ONE_MONTH"
)

func (r RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggInterval) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalFifteenMinutes, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneHour, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneDay, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneWeek, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                              `json:"level,required"`
	JSON  radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfo]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization string

const (
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationPercentage           RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationMin0Max              RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "MIN0_MAX"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationMinMax               RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "MIN_MAX"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRawValues            RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "RAW_VALUES"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationPercentageChange     RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRollingAverage       RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRatio                RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization = "RATIO"
)

func (r RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationPercentage, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationMin0Max, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationMinMax, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRawValues, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationPercentageChange, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRollingAverage, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationOverlappedPercentage, RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnit struct {
	Name  string                                                                   `json:"name,required"`
	Value string                                                                   `json:"value,required"`
	JSON  radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnitJSON `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnit]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0 struct {
	Timestamps  []time.Time                                                                          `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union `json:"-,extras"`
	JSON        radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON               `json:"-"`
}

// radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0]
type radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0JSON) RawJSON() string {
	return r.raw
}

// A numeric string.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union interface {
	ImplementsRadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RadarRankingInternetServiceGetTimeseriesGroupsResponseResultSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type RadarRankingInternetServiceGetTopServicesResponse struct {
	Result  RadarRankingInternetServiceGetTopServicesResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarRankingInternetServiceGetTopServicesResponseJSON   `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseJSON contains the JSON metadata
// for the struct [RadarRankingInternetServiceGetTopServicesResponse]
type radarRankingInternetServiceGetTopServicesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResult struct {
	Meta RadarRankingInternetServiceGetTopServicesResponseResultMeta   `json:"meta,required"`
	Top0 []RadarRankingInternetServiceGetTopServicesResponseResultTop0 `json:"top_0,required"`
	JSON radarRankingInternetServiceGetTopServicesResponseResultJSON   `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultJSON contains the JSON
// metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResult]
type radarRankingInternetServiceGetTopServicesResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultMeta struct {
	ConfidenceInfo RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarRankingInternetServiceGetTopServicesResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarRankingInternetServiceGetTopServicesResponseResultMetaUnit `json:"units,required"`
	JSON  radarRankingInternetServiceGetTopServicesResponseResultMetaJSON   `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMeta]
type radarRankingInternetServiceGetTopServicesResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfo struct {
	Annotations []RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                         `json:"level,required"`
	JSON  radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfo]
type radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotation]
type radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                `json:"startTime,required" format:"date-time"`
	JSON      radarRankingInternetServiceGetTopServicesResponseResultMetaDateRangeJSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMetaDateRange]
type radarRankingInternetServiceGetTopServicesResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization string

const (
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationPercentage           RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "PERCENTAGE"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationMin0Max              RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "MIN0_MAX"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationMinMax               RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "MIN_MAX"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRawValues            RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "RAW_VALUES"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationPercentageChange     RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRollingAverage       RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationOverlappedPercentage RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRatio                RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization = "RATIO"
)

func (r RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationPercentage, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationMin0Max, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationMinMax, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRawValues, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationPercentageChange, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRollingAverage, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationOverlappedPercentage, RadarRankingInternetServiceGetTopServicesResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTopServicesResponseResultMetaUnit struct {
	Name  string                                                              `json:"name,required"`
	Value string                                                              `json:"value,required"`
	JSON  radarRankingInternetServiceGetTopServicesResponseResultMetaUnitJSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultMetaUnit]
type radarRankingInternetServiceGetTopServicesResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceGetTopServicesResponseResultTop0 struct {
	Rank    int64                                                           `json:"rank,required"`
	Service string                                                          `json:"service,required"`
	JSON    radarRankingInternetServiceGetTopServicesResponseResultTop0JSON `json:"-"`
}

// radarRankingInternetServiceGetTopServicesResponseResultTop0JSON contains the
// JSON metadata for the struct
// [RadarRankingInternetServiceGetTopServicesResponseResultTop0]
type radarRankingInternetServiceGetTopServicesResponseResultTop0JSON struct {
	Rank        apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarRankingInternetServiceGetTopServicesResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarRankingInternetServiceGetTopServicesResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarRankingInternetServiceListCategoriesParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceListCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarRankingInternetServiceListCategoriesParams]'s query
// parameters as `url.Values`.
func (r RadarRankingInternetServiceListCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceListCategoriesParamsFormat string

const (
	RadarRankingInternetServiceListCategoriesParamsFormatJson RadarRankingInternetServiceListCategoriesParamsFormat = "JSON"
	RadarRankingInternetServiceListCategoriesParamsFormatCsv  RadarRankingInternetServiceListCategoriesParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceListCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceListCategoriesParamsFormatJson, RadarRankingInternetServiceListCategoriesParamsFormatCsv:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RadarRankingInternetServiceGetTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarRankingInternetServiceGetTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat string

const (
	RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatJson RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat = "JSON"
	RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatCsv  RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceGetTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatJson, RadarRankingInternetServiceGetTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type RadarRankingInternetServiceGetTopServicesParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RadarRankingInternetServiceGetTopServicesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RadarRankingInternetServiceGetTopServicesParams]'s query
// parameters as `url.Values`.
func (r RadarRankingInternetServiceGetTopServicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarRankingInternetServiceGetTopServicesParamsFormat string

const (
	RadarRankingInternetServiceGetTopServicesParamsFormatJson RadarRankingInternetServiceGetTopServicesParamsFormat = "JSON"
	RadarRankingInternetServiceGetTopServicesParamsFormatCsv  RadarRankingInternetServiceGetTopServicesParamsFormat = "CSV"
)

func (r RadarRankingInternetServiceGetTopServicesParamsFormat) IsKnown() bool {
	switch r {
	case RadarRankingInternetServiceGetTopServicesParamsFormatJson, RadarRankingInternetServiceGetTopServicesParamsFormatCsv:
		return true
	}
	return false
}
