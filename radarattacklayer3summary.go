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

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                    `json:"lastUpdated,required"`
	Normalization  string                                                                    `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                 `json:"level"`
	JSON        radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                  `json:"dataSource,required"`
	Description     string                                                                                  `json:"description,required"`
	EventType       string                                                                                  `json:"eventType,required"`
	IsInstantaneous bool                                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                                               `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetBitrateSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0 struct {
	Number1GbpsTo10Gbps   string                                                              `json:"_1_GBPS_TO_10_GBPS,required"`
	Number10GbpsTo100Gbps string                                                              `json:"_10_GBPS_TO_100_GBPS,required"`
	Number500MbpsTo1Gbps  string                                                              `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps           string                                                              `json:"OVER_100_GBPS,required"`
	Under500Mbps          string                                                              `json:"UNDER_500_MBPS,required"`
	JSON                  radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetBitrateSummaryResponseResultSummary0JSON struct {
	Number1GbpsTo10Gbps   apijson.Field
	Number10GbpsTo100Gbps apijson.Field
	Number500MbpsTo1Gbps  apijson.Field
	Over100Gbps           apijson.Field
	Under500Mbps          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
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

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                     `json:"lastUpdated,required"`
	Normalization  string                                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                   `json:"dataSource,required"`
	Description     string                                                                                   `json:"description,required"`
	EventType       string                                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetDurationSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0 struct {
	Number1HourTo3Hours  string                                                               `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins string                                                               `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins string                                                               `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  string                                                               `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           string                                                               `json:"OVER_3_HOURS,required"`
	Under10Mins          string                                                               `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0]
type radarAttackLayer3SummaryGetDurationSummaryResponseResultSummary0JSON struct {
	Number1HourTo3Hours  apijson.Field
	Number10MinsTo20Mins apijson.Field
	Number20MinsTo40Mins apijson.Field
	Number40MinsTo1Hour  apijson.Field
	Over3Hours           apijson.Field
	Under10Mins          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
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

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                     `json:"lastUpdated,required"`
	Normalization  string                                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                   `json:"dataSource,required"`
	Description     string                                                                                   `json:"description,required"`
	EventType       string                                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIndustrySummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                      `json:"lastUpdated,required"`
	Normalization  string                                                                      `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                   `json:"level"`
	JSON        radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                    `json:"dataSource,required"`
	Description     string                                                                                    `json:"description,required"`
	EventType       string                                                                                    `json:"eventType,required"`
	IsInstantaneous bool                                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                                 `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetIPVersionSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetIPVersionSummaryResponseResultSummary0 struct {
	IPv4 string                                                                `json:"IPv4,required"`
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

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                     `json:"lastUpdated,required"`
	Normalization  string                                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                   `json:"dataSource,required"`
	Description     string                                                                                   `json:"description,required"`
	EventType       string                                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetProtocolSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetProtocolSummaryResponseResultSummary0 struct {
	Gre  string                                                               `json:"GRE,required"`
	Icmp string                                                               `json:"ICMP,required"`
	Tcp  string                                                               `json:"TCP,required"`
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

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                   `json:"lastUpdated,required"`
	Normalization  string                                                                   `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                `json:"level"`
	JSON        radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                 `json:"dataSource,required"`
	Description     string                                                                                 `json:"description,required"`
	EventType       string                                                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                                              `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVectorSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                     `json:"lastUpdated,required"`
	Normalization  string                                                                     `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaJSON) RawJSON() string {
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

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                  `json:"level"`
	JSON        radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoJSON         `json:"-"`
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

type RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                   `json:"dataSource,required"`
	Description     string                                                                                   `json:"description,required"`
	EventType       string                                                                                   `json:"eventType,required"`
	IsInstantaneous bool                                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                                `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3SummaryGetVerticalSummaryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3SummaryGetBitrateSummaryParams struct {
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetBitrateSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetBitrateSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetBitrateSummaryParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetDurationSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetDurationSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetDurationSummaryParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetIndustrySummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetIndustrySummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetIndustrySummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetIPVersionSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetIPVersionSummaryParamsFormat] `query:"format"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetProtocolSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetProtocolSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetProtocolSummaryParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetVectorSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetVectorSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetVectorSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3SummaryGetVerticalSummaryParamsDirection] `query:"direction"`
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer3SummaryGetVerticalSummaryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3SummaryGetVerticalSummaryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects per group to the top items within the specified
	// time range. If there are more items than the limit, the response will include
	// the count of items, with any remaining items grouped together under an "other"
	// category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
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

// Together with the `location` parameter, will apply the filter to origin or
// target location.
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
