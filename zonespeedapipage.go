// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSpeedAPIPageService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpeedAPIPageService] method instead.
type ZoneSpeedAPIPageService struct {
	Options []option.RequestOption
	Tests   *ZoneSpeedAPIPageTestService
}

// NewZoneSpeedAPIPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIPageService(opts ...option.RequestOption) (r *ZoneSpeedAPIPageService) {
	r = &ZoneSpeedAPIPageService{}
	r.Options = opts
	r.Tests = NewZoneSpeedAPIPageTestService(opts...)
	return
}

// Lists all webpages which have been tested.
func (r *ZoneSpeedAPIPageService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneSpeedAPIPageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the core web vital metrics trend over time for a specific page.
func (r *ZoneSpeedAPIPageService) GetTrend(ctx context.Context, zoneID string, url string, query ZoneSpeedAPIPageGetTrendParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageGetTrendResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/trend", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ObservatoryAPIResponseCollection struct {
	Errors   []ObservatoryMessagesItem `json:"errors,required"`
	Messages []ObservatoryMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                 `json:"success,required"`
	JSON    observatoryAPIResponseCollectionJSON `json:"-"`
}

// observatoryAPIResponseCollectionJSON contains the JSON metadata for the struct
// [ObservatoryAPIResponseCollection]
type observatoryAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

// The type of device.
type ObservatoryDeviceType string

const (
	ObservatoryDeviceTypeDesktop ObservatoryDeviceType = "DESKTOP"
	ObservatoryDeviceTypeMobile  ObservatoryDeviceType = "MOBILE"
)

func (r ObservatoryDeviceType) IsKnown() bool {
	switch r {
	case ObservatoryDeviceTypeDesktop, ObservatoryDeviceTypeMobile:
		return true
	}
	return false
}

// A test region with a label.
type ObservatoryLabeledRegion struct {
	Label string `json:"label"`
	// A test region.
	Value ObservatoryRegion            `json:"value"`
	JSON  observatoryLabeledRegionJSON `json:"-"`
}

// observatoryLabeledRegionJSON contains the JSON metadata for the struct
// [ObservatoryLabeledRegion]
type observatoryLabeledRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryLabeledRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryLabeledRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type ObservatoryRegion string

const (
	ObservatoryRegionAsiaEast1           ObservatoryRegion = "asia-east1"
	ObservatoryRegionAsiaNortheast1      ObservatoryRegion = "asia-northeast1"
	ObservatoryRegionAsiaNortheast2      ObservatoryRegion = "asia-northeast2"
	ObservatoryRegionAsiaSouth1          ObservatoryRegion = "asia-south1"
	ObservatoryRegionAsiaSoutheast1      ObservatoryRegion = "asia-southeast1"
	ObservatoryRegionAustraliaSoutheast1 ObservatoryRegion = "australia-southeast1"
	ObservatoryRegionEuropeNorth1        ObservatoryRegion = "europe-north1"
	ObservatoryRegionEuropeSouthwest1    ObservatoryRegion = "europe-southwest1"
	ObservatoryRegionEuropeWest1         ObservatoryRegion = "europe-west1"
	ObservatoryRegionEuropeWest2         ObservatoryRegion = "europe-west2"
	ObservatoryRegionEuropeWest3         ObservatoryRegion = "europe-west3"
	ObservatoryRegionEuropeWest4         ObservatoryRegion = "europe-west4"
	ObservatoryRegionEuropeWest8         ObservatoryRegion = "europe-west8"
	ObservatoryRegionEuropeWest9         ObservatoryRegion = "europe-west9"
	ObservatoryRegionMeWest1             ObservatoryRegion = "me-west1"
	ObservatoryRegionSouthamericaEast1   ObservatoryRegion = "southamerica-east1"
	ObservatoryRegionUsCentral1          ObservatoryRegion = "us-central1"
	ObservatoryRegionUsEast1             ObservatoryRegion = "us-east1"
	ObservatoryRegionUsEast4             ObservatoryRegion = "us-east4"
	ObservatoryRegionUsSouth1            ObservatoryRegion = "us-south1"
	ObservatoryRegionUsWest1             ObservatoryRegion = "us-west1"
)

func (r ObservatoryRegion) IsKnown() bool {
	switch r {
	case ObservatoryRegionAsiaEast1, ObservatoryRegionAsiaNortheast1, ObservatoryRegionAsiaNortheast2, ObservatoryRegionAsiaSouth1, ObservatoryRegionAsiaSoutheast1, ObservatoryRegionAustraliaSoutheast1, ObservatoryRegionEuropeNorth1, ObservatoryRegionEuropeSouthwest1, ObservatoryRegionEuropeWest1, ObservatoryRegionEuropeWest2, ObservatoryRegionEuropeWest3, ObservatoryRegionEuropeWest4, ObservatoryRegionEuropeWest8, ObservatoryRegionEuropeWest9, ObservatoryRegionMeWest1, ObservatoryRegionSouthamericaEast1, ObservatoryRegionUsCentral1, ObservatoryRegionUsEast1, ObservatoryRegionUsEast4, ObservatoryRegionUsSouth1, ObservatoryRegionUsWest1:
		return true
	}
	return false
}

// The frequency of the test.
type ObservatoryScheduleFrequency string

const (
	ObservatoryScheduleFrequencyDaily  ObservatoryScheduleFrequency = "DAILY"
	ObservatoryScheduleFrequencyWeekly ObservatoryScheduleFrequency = "WEEKLY"
)

func (r ObservatoryScheduleFrequency) IsKnown() bool {
	switch r {
	case ObservatoryScheduleFrequencyDaily, ObservatoryScheduleFrequencyWeekly:
		return true
	}
	return false
}

type ZoneSpeedAPIPageListResponse struct {
	Result []ZoneSpeedAPIPageListResponseResult `json:"result"`
	JSON   zoneSpeedAPIPageListResponseJSON     `json:"-"`
	ObservatoryAPIResponseCollection
}

// zoneSpeedAPIPageListResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageListResponse]
type zoneSpeedAPIPageListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageListResponseResult struct {
	// A test region with a label.
	Region ObservatoryLabeledRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ObservatoryScheduleFrequency `json:"scheduleFrequency"`
	Tests             []ObservatoryPageTest        `json:"tests"`
	// A URL.
	URL  string                                 `json:"url"`
	JSON zoneSpeedAPIPageListResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageListResponseResult]
type zoneSpeedAPIPageListResponseResultJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageGetTrendResponse struct {
	Result ZoneSpeedAPIPageGetTrendResponseResult `json:"result"`
	JSON   zoneSpeedAPIPageGetTrendResponseJSON   `json:"-"`
	ObservatoryAPIResponseSingle
}

// zoneSpeedAPIPageGetTrendResponseJSON contains the JSON metadata for the struct
// [ZoneSpeedAPIPageGetTrendResponse]
type zoneSpeedAPIPageGetTrendResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageGetTrendResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageGetTrendResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageGetTrendResponseResult struct {
	// Cumulative Layout Shift trend.
	Cls []float64 `json:"cls"`
	// First Contentful Paint trend.
	Fcp []float64 `json:"fcp"`
	// Largest Contentful Paint trend.
	Lcp []float64 `json:"lcp"`
	// The Lighthouse score trend.
	PerformanceScore []float64 `json:"performanceScore"`
	// Speed Index trend.
	Si []float64 `json:"si"`
	// Total Blocking Time trend.
	Tbt []float64 `json:"tbt"`
	// Time To First Byte trend.
	Ttfb []float64 `json:"ttfb"`
	// Time To Interactive trend.
	Tti  []float64                                  `json:"tti"`
	JSON zoneSpeedAPIPageGetTrendResponseResultJSON `json:"-"`
}

// zoneSpeedAPIPageGetTrendResponseResultJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageGetTrendResponseResult]
type zoneSpeedAPIPageGetTrendResponseResultJSON struct {
	Cls              apijson.Field
	Fcp              apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageGetTrendResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageGetTrendResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageGetTrendParams struct {
	// The type of device.
	DeviceType param.Field[ObservatoryDeviceType] `query:"deviceType,required"`
	// A comma-separated list of metrics to include in the results.
	Metrics param.Field[string] `query:"metrics,required"`
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region,required"`
	Start  param.Field[time.Time]         `query:"start,required" format:"date-time"`
	// The timezone of the start and end timestamps.
	Tz  param.Field[string]    `query:"tz,required"`
	End param.Field[time.Time] `query:"end" format:"date-time"`
}

// URLQuery serializes [ZoneSpeedAPIPageGetTrendParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIPageGetTrendParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
