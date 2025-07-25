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

// RadarAttackLayer7SummaryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7SummaryService] method instead.
type RadarAttackLayer7SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryService) {
	r = &RadarAttackLayer7SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP method.
func (r *RadarAttackLayer7SummaryService) GetHTTPMethodSummary(ctx context.Context, query RadarAttackLayer7SummaryGetHTTPMethodSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by HTTP version.
func (r *RadarAttackLayer7SummaryService) GetHTTPVersionSummary(ctx context.Context, query RadarAttackLayer7SummaryGetHTTPVersionSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by targeted industry.
func (r *RadarAttackLayer7SummaryService) GetIndustrySummary(ctx context.Context, query RadarAttackLayer7SummaryGetIndustrySummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetIndustrySummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by IP version.
func (r *RadarAttackLayer7SummaryService) GetIPVersionSummary(ctx context.Context, query RadarAttackLayer7SummaryGetIPVersionSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetIPVersionSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by managed rules.
func (r *RadarAttackLayer7SummaryService) GetManagedRulesSummary(ctx context.Context, query RadarAttackLayer7SummaryGetManagedRulesSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetManagedRulesSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by mitigation product.
func (r *RadarAttackLayer7SummaryService) GetMitigationProductSummary(ctx context.Context, query RadarAttackLayer7SummaryGetMitigationProductSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetMitigationProductSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 7 attacks by targeted vertical.
func (r *RadarAttackLayer7SummaryService) GetVerticalSummary(ctx context.Context, query RadarAttackLayer7SummaryGetVerticalSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryGetVerticalSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetHTTPMethodSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponse]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                              `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResult]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                            `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                  `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnit struct {
	Name  string                                                                 `json:"name,required"`
	Value string                                                                 `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPMethodSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResult `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetHTTPVersionSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponse]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResult]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                             `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                   `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnit struct {
	Name  string                                                                  `json:"name,required"`
	Value string                                                                  `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0 struct {
	HTTP1X string                                                                  `json:"HTTP/1.x,required"`
	HTTP2  string                                                                  `json:"HTTP/2,required"`
	HTTP3  string                                                                  `json:"HTTP/3,required"`
	JSON   radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0]
type radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetHTTPVersionSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIndustrySummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetIndustrySummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetIndustrySummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetIndustrySummaryResponse]
type radarAttackLayer7SummaryGetIndustrySummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIndustrySummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                            `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetIndustrySummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResult]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMeta]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIndustrySummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetIPVersionSummaryResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetIPVersionSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetIPVersionSummaryResponse]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetIPVersionSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResult]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                           `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                 `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnit struct {
	Name  string                                                                `json:"name,required"`
	Value string                                                                `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0 struct {
	IPv4 string                                                                `json:"IPv4,required"`
	IPv6 string                                                                `json:"IPv6,required"`
	JSON radarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0]
type radarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetIPVersionSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetManagedRulesSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetManagedRulesSummaryResponse]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                                `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResult]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                              `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                    `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnit struct {
	Name  string                                                                   `json:"name,required"`
	Value string                                                                   `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetManagedRulesSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResult `json:"result,required"`
	Success bool                                                              `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetMitigationProductSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponse]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                                     `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResult]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                                   `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                         `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnit struct {
	Name  string                                                                        `json:"name,required"`
	Value string                                                                        `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnitJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetMitigationProductSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetVerticalSummaryResponse struct {
	Result  RadarAttackLayer7SummaryGetVerticalSummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer7SummaryGetVerticalSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryGetVerticalSummaryResponse]
type radarAttackLayer7SummaryGetVerticalSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetVerticalSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                            `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryGetVerticalSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResult]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMeta]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRange]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnit]
type radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7SummaryGetVerticalSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetHTTPMethodSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormatJson RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormatJson, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetHTTPMethodSummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetHTTPVersionSummaryParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormatJson RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormatJson, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetHTTPVersionSummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetIndustrySummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetIndustrySummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetIndustrySummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetIndustrySummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryParamsFormatJson RadarAttackLayer7SummaryGetIndustrySummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsFormatCsv  RadarAttackLayer7SummaryGetIndustrySummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryParamsFormatJson, RadarAttackLayer7SummaryGetIndustrySummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetIndustrySummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetIndustrySummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetIndustrySummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIPVersionSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetIPVersionSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetIPVersionSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormatJson RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormatJson, RadarAttackLayer7SummaryGetIPVersionSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetIPVersionSummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetIPVersionSummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetManagedRulesSummaryParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormatJson RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormatJson, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetManagedRulesSummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes
// [RadarAttackLayer7SummaryGetMitigationProductSummaryParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7SummaryGetMitigationProductSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormatJson RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormatJson, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetMitigationProductSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7SummaryGetVerticalSummaryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryGetVerticalSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryGetVerticalSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7SummaryGetVerticalSummaryParamsFormat string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryParamsFormatJson RadarAttackLayer7SummaryGetVerticalSummaryParamsFormat = "JSON"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsFormatCsv  RadarAttackLayer7SummaryGetVerticalSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryParamsFormatJson, RadarAttackLayer7SummaryGetVerticalSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodGet             RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPost            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodDelete          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPut             RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodHead            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPurge           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodOptions         RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPropfind        RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkcol           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPatch           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodACL             RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBcopy           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBdelete         RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBmove           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCheckin         RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCheckout        RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodConnect         RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCopy            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodLabel           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodLock            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMerge           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMove            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodNotify          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPoll            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodProppatch       RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodReport          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodSearch          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodTrace           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUnlock          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUpdate          RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodJson            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCook            RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodTrack           RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodGet, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPost, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodDelete, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPut, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodHead, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPurge, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodOptions, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPropfind, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkcol, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPatch, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodACL, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBcopy, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBdelete, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBmove, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBpropfind, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBproppatch, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCheckin, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCheckout, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodConnect, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCopy, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodLabel, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodLock, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMerge, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkactivity, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMkworkspace, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodMove, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodNotify, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodOrderpatch, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodPoll, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodProppatch, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodReport, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodSearch, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodSubscribe, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodTrace, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUncheckout, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUnlock, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUnsubscribe, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodUpdate, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodVersioncontrol, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodXmsenumatts, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodRpcOutData, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodRpcInData, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodJson, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodCook, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv1, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv2, RadarAttackLayer7SummaryGetVerticalSummaryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersionIPv4 RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersionIPv6 RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersionIPv4, RadarAttackLayer7SummaryGetVerticalSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductDdos               RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductWaf                RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductBotManagement      RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductAccessRules        RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductIPReputation       RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductAPIShield          RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductDdos, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductWaf, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductBotManagement, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductAccessRules, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductIPReputation, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductAPIShield, RadarAttackLayer7SummaryGetVerticalSummaryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}
