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

// RadarAttackLayer3TopService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TopService] method instead.
type RadarAttackLayer3TopService struct {
	Options   []option.RequestOption
	Locations *RadarAttackLayer3TopLocationService
}

// NewRadarAttackLayer3TopService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TopService(opts ...option.RequestOption) (r *RadarAttackLayer3TopService) {
	r = &RadarAttackLayer3TopService{}
	r.Options = opts
	r.Locations = NewRadarAttackLayer3TopLocationService(opts...)
	return
}

// Retrieves the top layer 3 attacks from origin to target location. Values are a
// percentage out of the total layer 3 attacks (with billing country). You can
// optionally limit the number of attacks by origin/target location (useful if all
// the top attacks are from or to the same location).
func (r *RadarAttackLayer3TopService) GetTopAttacks(ctx context.Context, query RadarAttackLayer3TopGetTopAttacksParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopGetTopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by industry endpoint.
func (r *RadarAttackLayer3TopService) GetTopIndustry(ctx context.Context, query RadarAttackLayer3TopGetTopIndustryParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopGetTopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This endpoint is deprecated. To continue getting this data, switch to the
// summary by vertical endpoint.
func (r *RadarAttackLayer3TopService) GetTopVerticals(ctx context.Context, query RadarAttackLayer3TopGetTopVerticalsParams, opts ...option.RequestOption) (res *RadarAttackLayer3TopGetTopVerticalsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TopGetTopAttacksResponse struct {
	Result  RadarAttackLayer3TopGetTopAttacksResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer3TopGetTopAttacksResponseJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TopGetTopAttacksResponse]
type radarAttackLayer3TopGetTopAttacksResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResult struct {
	Meta RadarAttackLayer3TopGetTopAttacksResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopGetTopAttacksResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopGetTopAttacksResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopGetTopAttacksResponseResult]
type radarAttackLayer3TopGetTopAttacksResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopGetTopAttacksResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopGetTopAttacksResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopAttacksResponseResultMeta]
type radarAttackLayer3TopGetTopAttacksResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopGetTopAttacksResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopGetTopAttacksResponseResultMetaDateRange]
type radarAttackLayer3TopGetTopAttacksResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksResponseResultTop0 struct {
	OriginCountryAlpha2 string                                                  `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                                  `json:"originCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarAttackLayer3TopGetTopAttacksResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopGetTopAttacksResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopAttacksResponseResultTop0]
type radarAttackLayer3TopGetTopAttacksResponseResultTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopAttacksResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopAttacksResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponse struct {
	Result  RadarAttackLayer3TopGetTopIndustryResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAttackLayer3TopGetTopIndustryResponseJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopGetTopIndustryResponse]
type radarAttackLayer3TopGetTopIndustryResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResult struct {
	Meta RadarAttackLayer3TopGetTopIndustryResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopGetTopIndustryResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopGetTopIndustryResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopGetTopIndustryResponseResult]
type radarAttackLayer3TopGetTopIndustryResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopGetTopIndustryResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopGetTopIndustryResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopIndustryResponseResultMeta]
type radarAttackLayer3TopGetTopIndustryResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopGetTopIndustryResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopGetTopIndustryResponseResultMetaDateRange]
type radarAttackLayer3TopGetTopIndustryResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopIndustryResponseResultTop0 struct {
	Name  string                                                   `json:"name,required"`
	Value string                                                   `json:"value,required"`
	JSON  radarAttackLayer3TopGetTopIndustryResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopGetTopIndustryResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopIndustryResponseResultTop0]
type radarAttackLayer3TopGetTopIndustryResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopIndustryResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopIndustryResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponse struct {
	Result  RadarAttackLayer3TopGetTopVerticalsResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarAttackLayer3TopGetTopVerticalsResponseJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TopGetTopVerticalsResponse]
type radarAttackLayer3TopGetTopVerticalsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResult struct {
	Meta RadarAttackLayer3TopGetTopVerticalsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer3TopGetTopVerticalsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer3TopGetTopVerticalsResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TopGetTopVerticalsResponseResult]
type radarAttackLayer3TopGetTopVerticalsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResultMeta struct {
	DateRange      []RadarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer3TopGetTopVerticalsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopVerticalsResponseResultMeta]
type radarAttackLayer3TopGetTopVerticalsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRange]
type radarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfo]
type radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopVerticalsResponseResultTop0 struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  radarAttackLayer3TopGetTopVerticalsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer3TopGetTopVerticalsResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer3TopGetTopVerticalsResponseResultTop0]
type radarAttackLayer3TopGetTopVerticalsResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TopGetTopVerticalsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TopGetTopVerticalsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TopGetTopAttacksParams struct {
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
	Format param.Field[RadarAttackLayer3TopGetTopAttacksParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TopGetTopAttacksParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[RadarAttackLayer3TopGetTopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limits the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Attack magnitude can be defined by total mitigated bytes or by total mitigated
	// attacks.
	Magnitude param.Field[RadarAttackLayer3TopGetTopAttacksParamsMagnitude] `query:"magnitude"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TopGetTopAttacksParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopGetTopAttacksParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopGetTopAttacksParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer3TopGetTopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer3TopGetTopAttacksParamsFormat string

const (
	RadarAttackLayer3TopGetTopAttacksParamsFormatJson RadarAttackLayer3TopGetTopAttacksParamsFormat = "JSON"
	RadarAttackLayer3TopGetTopAttacksParamsFormatCsv  RadarAttackLayer3TopGetTopAttacksParamsFormat = "CSV"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsFormatJson, RadarAttackLayer3TopGetTopAttacksParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopAttacksParamsIPVersion string

const (
	RadarAttackLayer3TopGetTopAttacksParamsIPVersionIPv4 RadarAttackLayer3TopGetTopAttacksParamsIPVersion = "IPv4"
	RadarAttackLayer3TopGetTopAttacksParamsIPVersionIPv6 RadarAttackLayer3TopGetTopAttacksParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsIPVersionIPv4, RadarAttackLayer3TopGetTopAttacksParamsIPVersionIPv6:
		return true
	}
	return false
}

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type RadarAttackLayer3TopGetTopAttacksParamsLimitDirection string

const (
	RadarAttackLayer3TopGetTopAttacksParamsLimitDirectionOrigin RadarAttackLayer3TopGetTopAttacksParamsLimitDirection = "ORIGIN"
	RadarAttackLayer3TopGetTopAttacksParamsLimitDirectionTarget RadarAttackLayer3TopGetTopAttacksParamsLimitDirection = "TARGET"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsLimitDirection) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsLimitDirectionOrigin, RadarAttackLayer3TopGetTopAttacksParamsLimitDirectionTarget:
		return true
	}
	return false
}

// Attack magnitude can be defined by total mitigated bytes or by total mitigated
// attacks.
type RadarAttackLayer3TopGetTopAttacksParamsMagnitude string

const (
	RadarAttackLayer3TopGetTopAttacksParamsMagnitudeMitigatedBytes   RadarAttackLayer3TopGetTopAttacksParamsMagnitude = "MITIGATED_BYTES"
	RadarAttackLayer3TopGetTopAttacksParamsMagnitudeMitigatedAttacks RadarAttackLayer3TopGetTopAttacksParamsMagnitude = "MITIGATED_ATTACKS"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsMagnitude) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsMagnitudeMitigatedBytes, RadarAttackLayer3TopGetTopAttacksParamsMagnitudeMitigatedAttacks:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TopGetTopAttacksParamsNormalization string

const (
	RadarAttackLayer3TopGetTopAttacksParamsNormalizationPercentage RadarAttackLayer3TopGetTopAttacksParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TopGetTopAttacksParamsNormalizationMinMax     RadarAttackLayer3TopGetTopAttacksParamsNormalization = "MIN_MAX"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsNormalization) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsNormalizationPercentage, RadarAttackLayer3TopGetTopAttacksParamsNormalizationMinMax:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopAttacksParamsProtocol string

const (
	RadarAttackLayer3TopGetTopAttacksParamsProtocolUdp  RadarAttackLayer3TopGetTopAttacksParamsProtocol = "UDP"
	RadarAttackLayer3TopGetTopAttacksParamsProtocolTcp  RadarAttackLayer3TopGetTopAttacksParamsProtocol = "TCP"
	RadarAttackLayer3TopGetTopAttacksParamsProtocolIcmp RadarAttackLayer3TopGetTopAttacksParamsProtocol = "ICMP"
	RadarAttackLayer3TopGetTopAttacksParamsProtocolGre  RadarAttackLayer3TopGetTopAttacksParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TopGetTopAttacksParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopAttacksParamsProtocolUdp, RadarAttackLayer3TopGetTopAttacksParamsProtocolTcp, RadarAttackLayer3TopGetTopAttacksParamsProtocolIcmp, RadarAttackLayer3TopGetTopAttacksParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopIndustryParams struct {
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
	Format param.Field[RadarAttackLayer3TopGetTopIndustryParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TopGetTopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopGetTopIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopGetTopIndustryParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopGetTopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer3TopGetTopIndustryParamsFormat string

const (
	RadarAttackLayer3TopGetTopIndustryParamsFormatJson RadarAttackLayer3TopGetTopIndustryParamsFormat = "JSON"
	RadarAttackLayer3TopGetTopIndustryParamsFormatCsv  RadarAttackLayer3TopGetTopIndustryParamsFormat = "CSV"
)

func (r RadarAttackLayer3TopGetTopIndustryParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopIndustryParamsFormatJson, RadarAttackLayer3TopGetTopIndustryParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopIndustryParamsIPVersion string

const (
	RadarAttackLayer3TopGetTopIndustryParamsIPVersionIPv4 RadarAttackLayer3TopGetTopIndustryParamsIPVersion = "IPv4"
	RadarAttackLayer3TopGetTopIndustryParamsIPVersionIPv6 RadarAttackLayer3TopGetTopIndustryParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TopGetTopIndustryParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopIndustryParamsIPVersionIPv4, RadarAttackLayer3TopGetTopIndustryParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopIndustryParamsProtocol string

const (
	RadarAttackLayer3TopGetTopIndustryParamsProtocolUdp  RadarAttackLayer3TopGetTopIndustryParamsProtocol = "UDP"
	RadarAttackLayer3TopGetTopIndustryParamsProtocolTcp  RadarAttackLayer3TopGetTopIndustryParamsProtocol = "TCP"
	RadarAttackLayer3TopGetTopIndustryParamsProtocolIcmp RadarAttackLayer3TopGetTopIndustryParamsProtocol = "ICMP"
	RadarAttackLayer3TopGetTopIndustryParamsProtocolGre  RadarAttackLayer3TopGetTopIndustryParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TopGetTopIndustryParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopIndustryParamsProtocolUdp, RadarAttackLayer3TopGetTopIndustryParamsProtocolTcp, RadarAttackLayer3TopGetTopIndustryParamsProtocolIcmp, RadarAttackLayer3TopGetTopIndustryParamsProtocolGre:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopVerticalsParams struct {
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
	Format param.Field[RadarAttackLayer3TopGetTopVerticalsParamsFormat] `query:"format"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarAttackLayer3TopGetTopVerticalsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TopGetTopVerticalsParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TopGetTopVerticalsParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TopGetTopVerticalsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format in which results will be returned.
type RadarAttackLayer3TopGetTopVerticalsParamsFormat string

const (
	RadarAttackLayer3TopGetTopVerticalsParamsFormatJson RadarAttackLayer3TopGetTopVerticalsParamsFormat = "JSON"
	RadarAttackLayer3TopGetTopVerticalsParamsFormatCsv  RadarAttackLayer3TopGetTopVerticalsParamsFormat = "CSV"
)

func (r RadarAttackLayer3TopGetTopVerticalsParamsFormat) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopVerticalsParamsFormatJson, RadarAttackLayer3TopGetTopVerticalsParamsFormatCsv:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopVerticalsParamsIPVersion string

const (
	RadarAttackLayer3TopGetTopVerticalsParamsIPVersionIPv4 RadarAttackLayer3TopGetTopVerticalsParamsIPVersion = "IPv4"
	RadarAttackLayer3TopGetTopVerticalsParamsIPVersionIPv6 RadarAttackLayer3TopGetTopVerticalsParamsIPVersion = "IPv6"
)

func (r RadarAttackLayer3TopGetTopVerticalsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopVerticalsParamsIPVersionIPv4, RadarAttackLayer3TopGetTopVerticalsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarAttackLayer3TopGetTopVerticalsParamsProtocol string

const (
	RadarAttackLayer3TopGetTopVerticalsParamsProtocolUdp  RadarAttackLayer3TopGetTopVerticalsParamsProtocol = "UDP"
	RadarAttackLayer3TopGetTopVerticalsParamsProtocolTcp  RadarAttackLayer3TopGetTopVerticalsParamsProtocol = "TCP"
	RadarAttackLayer3TopGetTopVerticalsParamsProtocolIcmp RadarAttackLayer3TopGetTopVerticalsParamsProtocol = "ICMP"
	RadarAttackLayer3TopGetTopVerticalsParamsProtocolGre  RadarAttackLayer3TopGetTopVerticalsParamsProtocol = "GRE"
)

func (r RadarAttackLayer3TopGetTopVerticalsParamsProtocol) IsKnown() bool {
	switch r {
	case RadarAttackLayer3TopGetTopVerticalsParamsProtocolUdp, RadarAttackLayer3TopGetTopVerticalsParamsProtocolTcp, RadarAttackLayer3TopGetTopVerticalsParamsProtocolIcmp, RadarAttackLayer3TopGetTopVerticalsParamsProtocolGre:
		return true
	}
	return false
}
