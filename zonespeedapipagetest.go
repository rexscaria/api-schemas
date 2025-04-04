// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSpeedAPIPageTestService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpeedAPIPageTestService] method instead.
type ZoneSpeedAPIPageTestService struct {
	Options []option.RequestOption
}

// NewZoneSpeedAPIPageTestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneSpeedAPIPageTestService(opts ...option.RequestOption) (r *ZoneSpeedAPIPageTestService) {
	r = &ZoneSpeedAPIPageTestService{}
	r.Options = opts
	return
}

// Deletes all tests for a specific webpage from a specific region. Deleted tests
// are still counted as part of the quota.
func (r *ZoneSpeedAPIPageTestService) DeleteAll(ctx context.Context, zoneID string, url string, body ZoneSpeedAPIPageTestDeleteAllParams, opts ...option.RequestOption) (res *ObservatoryCountResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Retrieves the result of a specific test.
func (r *ZoneSpeedAPIPageTestService) GetResult(ctx context.Context, zoneID string, url string, testID string, opts ...option.RequestOption) (res *ObservatoryPageTestResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests/%s", zoneID, url, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Test history (list of tests) for a specific webpage.
func (r *ZoneSpeedAPIPageTestService) ListHistory(ctx context.Context, zoneID string, url string, query ZoneSpeedAPIPageTestListHistoryParams, opts ...option.RequestOption) (res *ZoneSpeedAPIPageTestListHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Starts a test for a specific webpage, in a specific region.
func (r *ZoneSpeedAPIPageTestService) Start(ctx context.Context, zoneID string, url string, body ZoneSpeedAPIPageTestStartParams, opts ...option.RequestOption) (res *ObservatoryPageTestResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if url == "" {
		err = errors.New("missing required url parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ObservatoryCountResponse struct {
	Result ObservatoryCountResponseResult `json:"result"`
	JSON   observatoryCountResponseJSON   `json:"-"`
	ObservatoryAPIResponseSingle
}

// observatoryCountResponseJSON contains the JSON metadata for the struct
// [ObservatoryCountResponse]
type observatoryCountResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryCountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryCountResponseJSON) RawJSON() string {
	return r.raw
}

type ObservatoryCountResponseResult struct {
	// Number of items affected.
	Count float64                            `json:"count"`
	JSON  observatoryCountResponseResultJSON `json:"-"`
}

// observatoryCountResponseResultJSON contains the JSON metadata for the struct
// [ObservatoryCountResponseResult]
type observatoryCountResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryCountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryCountResponseResultJSON) RawJSON() string {
	return r.raw
}

// The Lighthouse report.
type ObservatoryLighthouseReport struct {
	// Cumulative Layout Shift.
	Cls float64 `json:"cls"`
	// The type of device.
	DeviceType ObservatoryDeviceType            `json:"deviceType"`
	Error      ObservatoryLighthouseReportError `json:"error"`
	// First Contentful Paint.
	Fcp float64 `json:"fcp"`
	// The URL to the full Lighthouse JSON report.
	JsonReportURL string `json:"jsonReportUrl"`
	// Largest Contentful Paint.
	Lcp float64 `json:"lcp"`
	// The Lighthouse performance score.
	PerformanceScore float64 `json:"performanceScore"`
	// Speed Index.
	Si float64 `json:"si"`
	// The state of the Lighthouse report.
	State ObservatoryLighthouseReportState `json:"state"`
	// Total Blocking Time.
	Tbt float64 `json:"tbt"`
	// Time To First Byte.
	Ttfb float64 `json:"ttfb"`
	// Time To Interactive.
	Tti  float64                         `json:"tti"`
	JSON observatoryLighthouseReportJSON `json:"-"`
}

// observatoryLighthouseReportJSON contains the JSON metadata for the struct
// [ObservatoryLighthouseReport]
type observatoryLighthouseReportJSON struct {
	Cls              apijson.Field
	DeviceType       apijson.Field
	Error            apijson.Field
	Fcp              apijson.Field
	JsonReportURL    apijson.Field
	Lcp              apijson.Field
	PerformanceScore apijson.Field
	Si               apijson.Field
	State            apijson.Field
	Tbt              apijson.Field
	Ttfb             apijson.Field
	Tti              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ObservatoryLighthouseReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryLighthouseReportJSON) RawJSON() string {
	return r.raw
}

type ObservatoryLighthouseReportError struct {
	// The error code of the Lighthouse result.
	Code ObservatoryLighthouseReportErrorCode `json:"code"`
	// Detailed error message.
	Detail string `json:"detail"`
	// The final URL displayed to the user.
	FinalDisplayedURL string                               `json:"finalDisplayedUrl"`
	JSON              observatoryLighthouseReportErrorJSON `json:"-"`
}

// observatoryLighthouseReportErrorJSON contains the JSON metadata for the struct
// [ObservatoryLighthouseReportError]
type observatoryLighthouseReportErrorJSON struct {
	Code              apijson.Field
	Detail            apijson.Field
	FinalDisplayedURL apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservatoryLighthouseReportError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryLighthouseReportErrorJSON) RawJSON() string {
	return r.raw
}

// The error code of the Lighthouse result.
type ObservatoryLighthouseReportErrorCode string

const (
	ObservatoryLighthouseReportErrorCodeNotReachable      ObservatoryLighthouseReportErrorCode = "NOT_REACHABLE"
	ObservatoryLighthouseReportErrorCodeDNSFailure        ObservatoryLighthouseReportErrorCode = "DNS_FAILURE"
	ObservatoryLighthouseReportErrorCodeNotHTML           ObservatoryLighthouseReportErrorCode = "NOT_HTML"
	ObservatoryLighthouseReportErrorCodeLighthouseTimeout ObservatoryLighthouseReportErrorCode = "LIGHTHOUSE_TIMEOUT"
	ObservatoryLighthouseReportErrorCodeUnknown           ObservatoryLighthouseReportErrorCode = "UNKNOWN"
)

func (r ObservatoryLighthouseReportErrorCode) IsKnown() bool {
	switch r {
	case ObservatoryLighthouseReportErrorCodeNotReachable, ObservatoryLighthouseReportErrorCodeDNSFailure, ObservatoryLighthouseReportErrorCodeNotHTML, ObservatoryLighthouseReportErrorCodeLighthouseTimeout, ObservatoryLighthouseReportErrorCodeUnknown:
		return true
	}
	return false
}

// The state of the Lighthouse report.
type ObservatoryLighthouseReportState string

const (
	ObservatoryLighthouseReportStateRunning  ObservatoryLighthouseReportState = "RUNNING"
	ObservatoryLighthouseReportStateComplete ObservatoryLighthouseReportState = "COMPLETE"
	ObservatoryLighthouseReportStateFailed   ObservatoryLighthouseReportState = "FAILED"
)

func (r ObservatoryLighthouseReportState) IsKnown() bool {
	switch r {
	case ObservatoryLighthouseReportStateRunning, ObservatoryLighthouseReportStateComplete, ObservatoryLighthouseReportStateFailed:
		return true
	}
	return false
}

type ObservatoryPageTest struct {
	// UUID
	ID   string    `json:"id"`
	Date time.Time `json:"date" format:"date-time"`
	// The Lighthouse report.
	DesktopReport ObservatoryLighthouseReport `json:"desktopReport"`
	// The Lighthouse report.
	MobileReport ObservatoryLighthouseReport `json:"mobileReport"`
	// A test region with a label.
	Region ObservatoryLabeledRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency ObservatoryScheduleFrequency `json:"scheduleFrequency"`
	// A URL.
	URL  string                  `json:"url"`
	JSON observatoryPageTestJSON `json:"-"`
}

// observatoryPageTestJSON contains the JSON metadata for the struct
// [ObservatoryPageTest]
type observatoryPageTestJSON struct {
	ID                apijson.Field
	Date              apijson.Field
	DesktopReport     apijson.Field
	MobileReport      apijson.Field
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservatoryPageTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryPageTestJSON) RawJSON() string {
	return r.raw
}

type ObservatoryPageTestResponseSingle struct {
	Result ObservatoryPageTest                   `json:"result"`
	JSON   observatoryPageTestResponseSingleJSON `json:"-"`
	ObservatoryAPIResponseSingle
}

// observatoryPageTestResponseSingleJSON contains the JSON metadata for the struct
// [ObservatoryPageTestResponseSingle]
type observatoryPageTestResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservatoryPageTestResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observatoryPageTestResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageTestListHistoryResponse struct {
	Result     []ObservatoryPageTest                             `json:"result"`
	ResultInfo ZoneSpeedAPIPageTestListHistoryResponseResultInfo `json:"result_info"`
	JSON       zoneSpeedAPIPageTestListHistoryResponseJSON       `json:"-"`
	ObservatoryAPIResponseCollection
}

// zoneSpeedAPIPageTestListHistoryResponseJSON contains the JSON metadata for the
// struct [ZoneSpeedAPIPageTestListHistoryResponse]
type zoneSpeedAPIPageTestListHistoryResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageTestListHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageTestListHistoryResponseResultInfo struct {
	Count      int64                                                 `json:"count"`
	Page       int64                                                 `json:"page"`
	PerPage    int64                                                 `json:"per_page"`
	TotalCount int64                                                 `json:"total_count"`
	JSON       zoneSpeedAPIPageTestListHistoryResponseResultInfoJSON `json:"-"`
}

// zoneSpeedAPIPageTestListHistoryResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneSpeedAPIPageTestListHistoryResponseResultInfo]
type zoneSpeedAPIPageTestListHistoryResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSpeedAPIPageTestListHistoryResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneSpeedAPIPageTestListHistoryResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneSpeedAPIPageTestDeleteAllParams struct {
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIPageTestDeleteAllParams]'s query parameters as
// `url.Values`.
func (r ZoneSpeedAPIPageTestDeleteAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSpeedAPIPageTestListHistoryParams struct {
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// A test region.
	Region param.Field[ObservatoryRegion] `query:"region"`
}

// URLQuery serializes [ZoneSpeedAPIPageTestListHistoryParams]'s query parameters
// as `url.Values`.
func (r ZoneSpeedAPIPageTestListHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSpeedAPIPageTestStartParams struct {
	// A test region.
	Region param.Field[ObservatoryRegion] `json:"region"`
}

func (r ZoneSpeedAPIPageTestStartParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
