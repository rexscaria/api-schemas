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

// RadarAttackLayer3SummaryService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3SummaryService] method instead.
type RadarAttackLayer3SummaryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer3SummaryService) {
	r = &RadarAttackLayer3SummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of layer 3 attacks by bitrate.
func (r *RadarAttackLayer3SummaryService) GetBitrateSummary(ctx context.Context, query RadarAttackLayer3SummaryGetBitrateSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetBitrateSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by duration.
func (r *RadarAttackLayer3SummaryService) GetDurationSummary(ctx context.Context, query RadarAttackLayer3SummaryGetDurationSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetDurationSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by targeted industry.
func (r *RadarAttackLayer3SummaryService) GetIndustrySummary(ctx context.Context, query RadarAttackLayer3SummaryGetIndustrySummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetIndustrySummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by IP version.
func (r *RadarAttackLayer3SummaryService) GetIPVersionSummary(ctx context.Context, query RadarAttackLayer3SummaryGetIPVersionSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetIPVersionSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by protocol.
func (r *RadarAttackLayer3SummaryService) GetProtocolSummary(ctx context.Context, query RadarAttackLayer3SummaryGetProtocolSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetProtocolSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by vector.
func (r *RadarAttackLayer3SummaryService) GetVectorSummary(ctx context.Context, query RadarAttackLayer3SummaryGetVectorSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetVectorSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the distribution of layer 3 attacks by targeted vertical.
func (r *RadarAttackLayer3SummaryService) GetVerticalSummary(ctx context.Context, query RadarAttackLayer3SummaryGetVerticalSummaryParams, opts ...option.RequestOption) (res *RadarAttackLayer3SummaryGetVerticalSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/summary/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetBitrateSummaryResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetBitrateSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryGetBitrateSummaryResponse]
type radarAttackLayer3SummaryGetBitrateSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetBitrateSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResult]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                         `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                               `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnit struct {
	Name  string                                                              `json:"name,required"`
	Value string                                                              `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0 struct {
	// A numeric string.
	OneGBPSToTenGBPS string `json:"_1_GBPS_TO_10_GBPS,required"`
	// A numeric string.
	TenGBPSToOneHundredGBPS string `json:"_10_GBPS_TO_100_GBPS,required"`
	// A numeric string.
	FiveHundredMBPSToOneGBPS string `json:"_500_MBPS_TO_1_GBPS,required"`
	// A numeric string.
	Over100Gbps string `json:"OVER_100_GBPS,required"`
	// A numeric string.
	Under500Mbps string                                                              `json:"UNDER_500_MBPS,required"`
	JSON         radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON struct {
	OneGBPSToTenGBPS         apijson.Field
	TenGBPSToOneHundredGBPS  apijson.Field
	FiveHundredMBPSToOneGBPS apijson.Field
	Over100Gbps              apijson.Field
	Under500Mbps             apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetDurationSummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetDurationSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetDurationSummaryResponse]
type radarAttackLayer3SummaryGetDurationSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetDurationSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResult]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0 struct {
	// A numeric string.
	OneHourToThreeHours string `json:"_1_HOUR_TO_3_HOURS,required"`
	// A numeric string.
	TenMinsToTwentyMins string `json:"_10_MINS_TO_20_MINS,required"`
	// A numeric string.
	TwentyMinsToFortyMins string `json:"_20_MINS_TO_40_MINS,required"`
	// A numeric string.
	FortyMinsToOneHour string `json:"_40_MINS_TO_1_HOUR,required"`
	// A numeric string.
	Over3Hours string `json:"OVER_3_HOURS,required"`
	// A numeric string.
	Under10Mins string                                                               `json:"UNDER_10_MINS,required"`
	JSON        radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON struct {
	OneHourToThreeHours   apijson.Field
	TenMinsToTwentyMins   apijson.Field
	TwentyMinsToFortyMins apijson.Field
	FortyMinsToOneHour    apijson.Field
	Over3Hours            apijson.Field
	Under10Mins           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIndustrySummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetIndustrySummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetIndustrySummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetIndustrySummaryResponse]
type radarAttackLayer3SummaryGetIndustrySummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                            `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetIndustrySummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResult]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetIPVersionSummaryResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetIPVersionSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetIPVersionSummaryResponse]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetIPVersionSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResult]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                           `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                 `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnit struct {
	Name  string                                                                `json:"name,required"`
	Value string                                                                `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0 struct {
	// A numeric string.
	IPv4 string `json:"IPv4,required"`
	// A numeric string.
	IPv6 string                                                                `json:"IPv6,required"`
	JSON radarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetProtocolSummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetProtocolSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetProtocolSummaryResponse]
type radarAttackLayer3SummaryGetProtocolSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetProtocolSummaryResponseResultJSON     `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResult]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0 struct {
	// A numeric string.
	Gre string `json:"GRE,required"`
	// A numeric string.
	Icmp string `json:"ICMP,required"`
	// A numeric string.
	Tcp string `json:"TCP,required"`
	// A numeric string.
	Udp  string                                                               `json:"UDP,required"`
	JSON radarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVectorSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetVectorSummaryResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetVectorSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer3SummaryGetVectorSummaryResponse]
type radarAttackLayer3SummaryGetVectorSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVectorSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                          `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetVectorSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetVectorSummaryResponseResult]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                        `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnit struct {
	Name  string                                                             `json:"name,required"`
	Value string                                                             `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVerticalSummaryResponse struct {
	Result  RadarAttackLayer3SummaryGetVerticalSummaryResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer3SummaryGetVerticalSummaryResponseJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3SummaryGetVerticalSummaryResponse]
type radarAttackLayer3SummaryGetVerticalSummaryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResult struct {
	// Metadata for the results.
	Meta     RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta `json:"meta,required"`
	Summary0 map[string]string                                            `json:"summary_0,required"`
	JSON     radarAttackLayer3SummaryGetVerticalSummaryResponseResultJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResult]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta struct {
	ConfidenceInfo RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnit `json:"units,required"`
	JSON  radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON   `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                          `json:"level,required"`
	JSON  radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                   `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                                `json:"startDate,required" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRange]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization string

const (
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentage           RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "PERCENTAGE"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationMin0Max              RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "MIN0_MAX"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationMinMax               RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "MIN_MAX"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRawValues            RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "RAW_VALUES"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentageChange     RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRollingAverage       RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationOverlappedPercentage RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRatio                RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization = "RATIO"
)

func (r RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentage, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationMin0Max, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationMinMax, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRawValues, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationPercentageChange, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRollingAverage, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationOverlappedPercentage, RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnit struct {
	Name  string                                                               `json:"name,required"`
	Value string                                                               `json:"value,required"`
	JSON  radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnitJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnitJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnit]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetBitrateSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetBitrateSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetBitrateSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetBitrateSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetBitrateSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetBitrateSummaryParamsFormatJson RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetBitrateSummaryParamsFormatJson, RadarAttackLayer3SummaryGetBitrateSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolGre  RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetBitrateSummaryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetDurationSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetDurationSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetDurationSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetDurationSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetDurationSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetDurationSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetDurationSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetDurationSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetDurationSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetDurationSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetDurationSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetDurationSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetDurationSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetDurationSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetDurationSummaryParamsFormatJson RadarAttackLayer3SummaryGetDurationSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetDurationSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetDurationSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetDurationSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetDurationSummaryParamsFormatJson, RadarAttackLayer3SummaryGetDurationSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolGre  RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetDurationSummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetDurationSummaryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIndustrySummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetIndustrySummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetIndustrySummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetIndustrySummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsDirectionTarget RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIndustrySummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetIndustrySummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetIndustrySummaryParamsFormatJson RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsFormatCsv  RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIndustrySummaryParamsFormatJson, RadarAttackLayer3SummaryGetIndustrySummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolGre  RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetIndustrySummaryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIPVersionSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetIPVersionSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetIPVersionSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormatJson RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormatJson, RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolGre  RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetIPVersionSummaryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetProtocolSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion] `query:"ipVersion"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetProtocolSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetProtocolSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetProtocolSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetProtocolSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetProtocolSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetProtocolSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetProtocolSummaryParamsFormatJson RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetProtocolSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetProtocolSummaryParamsFormatJson, RadarAttackLayer3SummaryGetProtocolSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVectorSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetVectorSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetVectorSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetVectorSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetVectorSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetVectorSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetVectorSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetVectorSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetVectorSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetVectorSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetVectorSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVectorSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetVectorSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetVectorSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetVectorSummaryParamsFormatJson RadarAttackLayer3SummaryGetVectorSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetVectorSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetVectorSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetVectorSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVectorSummaryParamsFormatJson, RadarAttackLayer3SummaryGetVectorSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolGre  RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetVectorSummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetVectorSummaryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVerticalSummaryParams struct {
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
	// Specifies whether the `location` filter applies to the source or target
	// location.
	Direction param.Field[RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion] `query:"ipVersion"`
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
	// Filters the results by layer 3/4 protocol.
	Protocol param.Field[[]RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3SummaryGetVerticalSummaryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3SummaryGetVerticalSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies whether the `location` filter applies to the source or target
// location.
type RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection string

const (
	RadarAttackLayer3SummaryGetVerticalSummaryParamsDirectionOrigin RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection = "ORIGIN"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsDirectionTarget RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection = "TARGET"
)

func (r RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVerticalSummaryParamsDirectionOrigin, RadarAttackLayer3SummaryGetVerticalSummaryParamsDirectionTarget:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat string

const (
	RadarAttackLayer3SummaryGetVerticalSummaryParamsFormatJson RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat = "JSON"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsFormatCsv  RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat = "CSV"
)

func (r RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVerticalSummaryParamsFormatJson, RadarAttackLayer3SummaryGetVerticalSummaryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion string

const (
	RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersionIPv4 RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion = "IPv4"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersionIPv6 RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersionIPv4, RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol string

const (
	RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolUdp  RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol = "UDP"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolTcp  RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol = "TCP"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolIcmp RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol = "ICMP"
	RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolGre  RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolUdp, RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolTcp, RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolIcmp, RadarAttackLayer3SummaryGetVerticalSummaryParamsProtocolGre:
		return true
	}
	return false
}
