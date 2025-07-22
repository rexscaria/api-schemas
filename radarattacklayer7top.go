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

// RadarAttackLayer7TopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer7TopService] method instead.
type RadarAttackLayer7TopService struct {
	Options   []option.RequestOption
	Ases      *RadarAttackLayer7TopAseService
	Locations *RadarAttackLayer7TopLocationService
}

// NewRadarAttackLayer7TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7TopService(opts ...option.RequestOption) (r *RadarAttackLayer7TopService) {
	r = &RadarAttackLayer7TopService{}
	r.Options = opts
	r.Ases = NewRadarAttackLayer7TopAseService(opts...)
	r.Locations = NewRadarAttackLayer7TopLocationService(opts...)
	return
}

// Retrieves the top attacks from origin to target location. Values are percentages
// of the total layer 7 attacks (with billing country). The attack magnitude can be
// defined by the number of mitigated requests or by the number of zones affected.
// You can optionally limit the number of attacks by origin/target location (useful
// if all the top attacks are from or to the same location).
func (r *RadarAttackLayer7TopService) GetTopAttacks(ctx context.Context, query RadarAttackLayer7TopGetTopAttacksParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopGetTopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by industry endpoint.
//
// Deprecated: deprecated
func (r *RadarAttackLayer7TopService) GetTopIndustry(ctx context.Context, query RadarAttackLayer7TopGetTopIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopGetTopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by vertical endpoint.
//
// Deprecated: deprecated
func (r *RadarAttackLayer7TopService) GetTopVerticals(ctx context.Context, query RadarAttackLayer7TopGetTopVerticalsParams, opts ...option.RequestOption) (res *RadarAttackLayer7TopGetTopVerticalsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7TopGetTopAttacksResponse struct {
	Result  RadarAttackLayer7TopGetTopAttacksResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TopGetTopAttacksResponseJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TopGetTopAttacksResponse]
type radarAttackLayer7TopGetTopAttacksResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResult struct {
	Meta RadarAttackLayer7TopGetTopAttacksResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopGetTopAttacksResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopGetTopAttacksResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopGetTopAttacksResponseResult]
type radarAttackLayer7TopGetTopAttacksResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopGetTopAttacksResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopGetTopAttacksResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopAttacksResponseResultMeta]
type radarAttackLayer7TopGetTopAttacksResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopGetTopAttacksResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopGetTopAttacksResponseResultMetaDateRange]
type radarAttackLayer7TopGetTopAttacksResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                  `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                  `json:"originCountryName,required"`
	TargetCountryAlpha2 string                                                  `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                                  `json:"targetCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarAttackLayer7TopGetTopAttacksResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopGetTopAttacksResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopAttacksResponseResultTop0]
type radarAttackLayer7TopGetTopAttacksResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopAttacksResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopAttacksResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponse struct {
	Result  RadarAttackLayer7TopGetTopIndustryResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer7TopGetTopIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopGetTopIndustryResponse]
type radarAttackLayer7TopGetTopIndustryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResult struct {
	Meta RadarAttackLayer7TopGetTopIndustryResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopGetTopIndustryResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopGetTopIndustryResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopGetTopIndustryResponseResult]
type radarAttackLayer7TopGetTopIndustryResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopGetTopIndustryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopGetTopIndustryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopIndustryResponseResultMeta]
type radarAttackLayer7TopGetTopIndustryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopGetTopIndustryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopGetTopIndustryResponseResultMetaDateRange]
type radarAttackLayer7TopGetTopIndustryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopIndustryResponseResultTop0 struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAttackLayer7TopGetTopIndustryResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopGetTopIndustryResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopIndustryResponseResultTop0]
type radarAttackLayer7TopGetTopIndustryResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopIndustryResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopIndustryResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponse struct {
	Result  RadarAttackLayer7TopGetTopVerticalsResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAttackLayer7TopGetTopVerticalsResponseJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TopGetTopVerticalsResponse]
type radarAttackLayer7TopGetTopVerticalsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResult struct {
	Meta RadarAttackLayer7TopGetTopVerticalsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7TopGetTopVerticalsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7TopGetTopVerticalsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7TopGetTopVerticalsResponseResult]
type radarAttackLayer7TopGetTopVerticalsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7TopGetTopVerticalsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopVerticalsResponseResultMeta]
type radarAttackLayer7TopGetTopVerticalsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRange]
type radarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfo]
type radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopVerticalsResponseResultTop0 struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  radarAttackLayer7TopGetTopVerticalsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7TopGetTopVerticalsResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7TopGetTopVerticalsResponseResultTop0]
type radarAttackLayer7TopGetTopVerticalsResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TopGetTopVerticalsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer7TopGetTopVerticalsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer7TopGetTopAttacksParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7TopGetTopAttacksParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer7TopGetTopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limits the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// This parameter is deprecated. In the future, we will only support attack
	// magnitude defined by the total number of mitigated requests
	// (MITIGATED_REQUESTS).
	Magnitude param.Field[RadarAttackLayer7TopGetTopAttacksParamsMagnitude] `query:"magnitude"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TopGetTopAttacksParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7TopGetTopAttacksParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7TopGetTopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopGetTopAttacksParamsFormat string

const (
	RadarAttackLayer7TopGetTopAttacksParamsFormatJson RadarAttackLayer7TopGetTopAttacksParamsFormat = "JSON"
	RadarAttackLayer7TopGetTopAttacksParamsFormatCsv  RadarAttackLayer7TopGetTopAttacksParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopGetTopAttacksParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopAttacksParamsFormatJson, RadarAttackLayer7TopGetTopAttacksParamsFormatCsv:
		return true
	}
	return false
}

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer7TopGetTopAttacksParamsLimitDirection string

const (
	RadarAttackLayer7TopGetTopAttacksParamsLimitDirectionOrigin RadarAttackLayer7TopGetTopAttacksParamsLimitDirection = "ORIGIN"
	RadarAttackLayer7TopGetTopAttacksParamsLimitDirectionTarget RadarAttackLayer7TopGetTopAttacksParamsLimitDirection = "TARGET"
)

func (r RadarAttackLayer7TopGetTopAttacksParamsLimitDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopAttacksParamsLimitDirectionOrigin, RadarAttackLayer7TopGetTopAttacksParamsLimitDirectionTarget:
		return true
	}
	return false
}

// This parameter is deprecated. In the future, we will only support attack
// magnitude defined by the total number of mitigated requests
// (MITIGATED_REQUESTS).
type RadarAttackLayer7TopGetTopAttacksParamsMagnitude string

const (
	RadarAttackLayer7TopGetTopAttacksParamsMagnitudeAffectedZones     RadarAttackLayer7TopGetTopAttacksParamsMagnitude = "AFFECTED_ZONES"
	RadarAttackLayer7TopGetTopAttacksParamsMagnitudeMitigatedRequests RadarAttackLayer7TopGetTopAttacksParamsMagnitude = "MITIGATED_REQUESTS"
)

func (r RadarAttackLayer7TopGetTopAttacksParamsMagnitude) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopAttacksParamsMagnitudeAffectedZones, RadarAttackLayer7TopGetTopAttacksParamsMagnitudeMitigatedRequests:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct string

const (
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductDdos               RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductWaf                RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductBotManagement      RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductAccessRules        RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductIPReputation       RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductAPIShield          RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopGetTopAttacksParamsMitigationProductDataLossPrevention RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopGetTopAttacksParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopAttacksParamsMitigationProductDdos, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductWaf, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductBotManagement, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductAccessRules, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductIPReputation, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductAPIShield, RadarAttackLayer7TopGetTopAttacksParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TopGetTopAttacksParamsNormalization string

const (
	RadarAttackLayer7TopGetTopAttacksParamsNormalizationPercentage RadarAttackLayer7TopGetTopAttacksParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7TopGetTopAttacksParamsNormalizationMinMax     RadarAttackLayer7TopGetTopAttacksParamsNormalization = "MIN_MAX"
)

func (r RadarAttackLayer7TopGetTopAttacksParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopAttacksParamsNormalizationPercentage, RadarAttackLayer7TopGetTopAttacksParamsNormalizationMinMax:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopIndustryParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7TopGetTopIndustryParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TopGetTopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopGetTopIndustryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopGetTopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopGetTopIndustryParamsFormat string

const (
	RadarAttackLayer7TopGetTopIndustryParamsFormatJson RadarAttackLayer7TopGetTopIndustryParamsFormat = "JSON"
	RadarAttackLayer7TopGetTopIndustryParamsFormatCsv  RadarAttackLayer7TopGetTopIndustryParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopGetTopIndustryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopIndustryParamsFormatJson, RadarAttackLayer7TopGetTopIndustryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod string

const (
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodGet             RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "GET"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPost            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "POST"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodDelete          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPut             RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "PUT"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodHead            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPurge           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodOptions         RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPropfind        RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkcol           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPatch           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodACL             RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "ACL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBcopy           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBdelete         RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBmove           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBpropfind       RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBproppatch      RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCheckin         RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCheckout        RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodConnect         RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCopy            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "COPY"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodLabel           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodLock            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMerge           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkactivity      RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkworkspace     RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMove            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodNotify          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodOrderpatch      RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPoll            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "POLL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodProppatch       RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodReport          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodSearch          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodSubscribe       RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodTrace           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUncheckout      RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUnlock          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUnsubscribe     RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUpdate          RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodVersioncontrol  RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBaselinecontrol RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodXmsenumatts     RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodRpcOutData      RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodRpcInData       RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodJson            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "JSON"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCook            RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "COOK"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodTrack           RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TopGetTopIndustryParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodGet, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPost, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodDelete, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPut, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodHead, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPurge, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodOptions, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPropfind, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkcol, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPatch, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodACL, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBcopy, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBdelete, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBmove, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBpropfind, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBproppatch, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCheckin, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCheckout, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodConnect, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCopy, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodLabel, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodLock, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMerge, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkactivity, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMkworkspace, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodMove, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodNotify, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodOrderpatch, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodPoll, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodProppatch, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodReport, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodSearch, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodSubscribe, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodTrace, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUncheckout, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUnlock, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUnsubscribe, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodUpdate, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodVersioncontrol, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodXmsenumatts, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodRpcOutData, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodRpcInData, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodJson, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodCook, RadarAttackLayer7TopGetTopIndustryParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion string

const (
	RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv1 RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv2 RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv3 RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TopGetTopIndustryParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv1, RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv2, RadarAttackLayer7TopGetTopIndustryParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopIndustryParamsIPVersion string

const (
	RadarAttackLayer7TopGetTopIndustryParamsIPVersionIPv4 RadarAttackLayer7TopGetTopIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer7TopGetTopIndustryParamsIPVersionIPv6 RadarAttackLayer7TopGetTopIndustryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TopGetTopIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopIndustryParamsIPVersionIPv4, RadarAttackLayer7TopGetTopIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct string

const (
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductDdos               RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductWaf                RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductBotManagement      RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductAccessRules        RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductIPReputation       RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductAPIShield          RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopGetTopIndustryParamsMitigationProductDataLossPrevention RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopGetTopIndustryParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopIndustryParamsMitigationProductDdos, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductWaf, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductBotManagement, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductAccessRules, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductIPReputation, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductAPIShield, RadarAttackLayer7TopGetTopIndustryParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopVerticalsParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
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
	// Format in which results will be returned.
	Format param.Field[RadarAttackLayer7TopGetTopVerticalsParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer7TopGetTopVerticalsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7TopGetTopVerticalsParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7TopGetTopVerticalsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer7TopGetTopVerticalsParamsFormat string

const (
	RadarAttackLayer7TopGetTopVerticalsParamsFormatJson RadarAttackLayer7TopGetTopVerticalsParamsFormat = "JSON"
	RadarAttackLayer7TopGetTopVerticalsParamsFormatCsv  RadarAttackLayer7TopGetTopVerticalsParamsFormat = "CSV"
)

func (r RadarAttackLayer7TopGetTopVerticalsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopVerticalsParamsFormatJson, RadarAttackLayer7TopGetTopVerticalsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod string

const (
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodGet             RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "GET"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPost            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "POST"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodDelete          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPut             RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "PUT"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodHead            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPurge           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodOptions         RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPropfind        RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkcol           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPatch           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodACL             RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "ACL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBcopy           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBdelete         RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBmove           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBpropfind       RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBproppatch      RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCheckin         RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCheckout        RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodConnect         RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCopy            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "COPY"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodLabel           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodLock            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMerge           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkactivity      RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkworkspace     RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMove            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodNotify          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodOrderpatch      RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPoll            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "POLL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodProppatch       RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodReport          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodSearch          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodSubscribe       RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodTrace           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUncheckout      RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUnlock          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUnsubscribe     RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUpdate          RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodVersioncontrol  RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBaselinecontrol RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodXmsenumatts     RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodRpcOutData      RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodRpcInData       RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodJson            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "JSON"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCook            RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "COOK"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodTrack           RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod = "TRACK"
)

func (r RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethod) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodGet, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPost, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodDelete, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPut, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodHead, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPurge, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodOptions, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPropfind, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkcol, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPatch, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodACL, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBcopy, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBdelete, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBmove, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBpropfind, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBproppatch, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCheckin, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCheckout, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodConnect, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCopy, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodLabel, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodLock, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMerge, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkactivity, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMkworkspace, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodMove, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodNotify, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodOrderpatch, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodPoll, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodProppatch, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodReport, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodSearch, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodSubscribe, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodTrace, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUncheckout, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUnlock, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUnsubscribe, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodUpdate, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodVersioncontrol, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodBaselinecontrol, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodXmsenumatts, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodRpcOutData, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodRpcInData, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodJson, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodCook, RadarAttackLayer7TopGetTopVerticalsParamsHTTPMethodTrack:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion string

const (
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv1 RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv2 RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv3 RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion = "HTTPv3"
)

func (r RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv1, RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv2, RadarAttackLayer7TopGetTopVerticalsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopVerticalsParamsIPVersion string

const (
	RadarAttackLayer7TopGetTopVerticalsParamsIPVersionIPv4 RadarAttackLayer7TopGetTopVerticalsParamsIPVersion = "IPv4"
	RadarAttackLayer7TopGetTopVerticalsParamsIPVersionIPv6 RadarAttackLayer7TopGetTopVerticalsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer7TopGetTopVerticalsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopVerticalsParamsIPVersionIPv4, RadarAttackLayer7TopGetTopVerticalsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct string

const (
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductDdos               RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductWaf                RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "WAF"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductBotManagement      RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductAccessRules        RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductIPReputation       RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductAPIShield          RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductDataLossPrevention RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r RadarAttackLayer7TopGetTopVerticalsParamsMitigationProduct) IsKnown() bool {
	switch r {
	case RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductDdos, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductWaf, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductBotManagement, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductAccessRules, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductIPReputation, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductAPIShield, RadarAttackLayer7TopGetTopVerticalsParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}
