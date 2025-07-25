// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
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

// RadarHTTPTopLocationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTopLocationService] method instead.
type RadarHTTPTopLocationService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationService(opts ...option.RequestOption) (r *RadarHTTPTopLocationService) {
	r = &RadarHTTPTopLocationService{}
	r.Options = opts
	return
}

// Retrieves the top locations by HTTP requests.
func (r *RadarHTTPTopLocationService) List(ctx context.Context, query RadarHTTPTopLocationListParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested bot class.
func (r *RadarHTTPTopLocationService) ListByBotClass(ctx context.Context, botClass RadarHTTPTopLocationListByBotClassParamsBotClass, query RadarHTTPTopLocationListByBotClassParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested browser family.
func (r *RadarHTTPTopLocationService) ListByBrowserFamily(ctx context.Context, browserFamily RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily, query RadarHTTPTopLocationListByBrowserFamilyParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByBrowserFamilyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/browser_family/%v", browserFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested device type.
func (r *RadarHTTPTopLocationService) ListByDeviceType(ctx context.Context, deviceType RadarHTTPTopLocationListByDeviceTypeParamsDeviceType, query RadarHTTPTopLocationListByDeviceTypeParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByDeviceTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested HTTP protocol.
func (r *RadarHTTPTopLocationService) ListByHTTPProtocol(ctx context.Context, httpProtocol RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocol, query RadarHTTPTopLocationListByHTTPProtocolParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByHTTPProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested HTTP version.
func (r *RadarHTTPTopLocationService) ListByHTTPVersion(ctx context.Context, httpVersion RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion, query RadarHTTPTopLocationListByHTTPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested IP version.
func (r *RadarHTTPTopLocationService) ListByIPVersion(ctx context.Context, ipVersion RadarHTTPTopLocationListByIPVersionParamsIPVersion, query RadarHTTPTopLocationListByIPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested operating
// system.
func (r *RadarHTTPTopLocationService) ListByOs(ctx context.Context, os RadarHTTPTopLocationListByOsParamsOs, query RadarHTTPTopLocationListByOsParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByOsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top locations, by HTTP requests, of the requested TLS protocol
// version.
func (r *RadarHTTPTopLocationService) ListByTlsVersion(ctx context.Context, tlsVersion RadarHTTPTopLocationListByTlsVersionParamsTlsVersion, query RadarHTTPTopLocationListByTlsVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListByTlsVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationListResponse struct {
	Result  RadarHTTPTopLocationListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarHTTPTopLocationListResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopLocationListResponse]
type radarHTTPTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationListResponseResult]
type radarHTTPTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListResponseResultMeta]
type radarHTTPTopLocationListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                        `json:"level,required"`
	JSON  radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                              `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListResponseResultMetaDateRange]
type radarHTTPTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListResponseResultMetaUnit struct {
	Name  string                                             `json:"name,required"`
	Value string                                             `json:"value,required"`
	JSON  radarHTTPTopLocationListResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaUnitJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListResponseResultMetaUnit]
type radarHTTPTopLocationListResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                         `json:"value,required"`
	JSON  radarHTTPTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListResponseResultTop0]
type radarHTTPTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBotClassResponse struct {
	Result  RadarHTTPTopLocationListByBotClassResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarHTTPTopLocationListByBotClassResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByBotClassResponse]
type radarHTTPTopLocationListByBotClassResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBotClassResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByBotClassResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByBotClassResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByBotClassResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByBotClassResponseResult]
type radarHTTPTopLocationListByBotClassResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByBotClassResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByBotClassResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByBotClassResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByBotClassResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByBotClassResponseResultMeta]
type radarHTTPTopLocationListByBotClassResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                  `json:"level,required"`
	JSON  radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                        `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBotClassResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByBotClassResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByBotClassResponseResultMetaDateRange]
type radarHTTPTopLocationListByBotClassResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByBotClassResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByBotClassResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassResponseResultMetaUnit struct {
	Name  string                                                       `json:"name,required"`
	Value string                                                       `json:"value,required"`
	JSON  radarHTTPTopLocationListByBotClassResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByBotClassResponseResultMetaUnit]
type radarHTTPTopLocationListByBotClassResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBotClassResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                   `json:"value,required"`
	JSON  radarHTTPTopLocationListByBotClassResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByBotClassResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByBotClassResponseResultTop0]
type radarHTTPTopLocationListByBotClassResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBotClassResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBotClassResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBrowserFamilyResponse struct {
	Result  RadarHTTPTopLocationListByBrowserFamilyResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarHTTPTopLocationListByBrowserFamilyResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByBrowserFamilyResponse]
type radarHTTPTopLocationListByBrowserFamilyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBrowserFamilyResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByBrowserFamilyResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByBrowserFamilyResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByBrowserFamilyResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByBrowserFamilyResponseResult]
type radarHTTPTopLocationListByBrowserFamilyResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByBrowserFamilyResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByBrowserFamilyResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultMeta]
type radarHTTPTopLocationListByBrowserFamilyResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                       `json:"level,required"`
	JSON  radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                             `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRange]
type radarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnit struct {
	Name  string                                                            `json:"name,required"`
	Value string                                                            `json:"value,required"`
	JSON  radarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnit]
type radarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByBrowserFamilyResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                        `json:"value,required"`
	JSON  radarHTTPTopLocationListByBrowserFamilyResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByBrowserFamilyResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByBrowserFamilyResponseResultTop0]
type radarHTTPTopLocationListByBrowserFamilyResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByBrowserFamilyResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByBrowserFamilyResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByDeviceTypeResponse struct {
	Result  RadarHTTPTopLocationListByDeviceTypeResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPTopLocationListByDeviceTypeResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByDeviceTypeResponse]
type radarHTTPTopLocationListByDeviceTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByDeviceTypeResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByDeviceTypeResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByDeviceTypeResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByDeviceTypeResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByDeviceTypeResponseResult]
type radarHTTPTopLocationListByDeviceTypeResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByDeviceTypeResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByDeviceTypeResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByDeviceTypeResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByDeviceTypeResponseResultMeta]
type radarHTTPTopLocationListByDeviceTypeResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                    `json:"level,required"`
	JSON  radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                          `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRange]
type radarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByDeviceTypeResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeResponseResultMetaUnit struct {
	Name  string                                                         `json:"name,required"`
	Value string                                                         `json:"value,required"`
	JSON  radarHTTPTopLocationListByDeviceTypeResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByDeviceTypeResponseResultMetaUnit]
type radarHTTPTopLocationListByDeviceTypeResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByDeviceTypeResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                     `json:"value,required"`
	JSON  radarHTTPTopLocationListByDeviceTypeResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByDeviceTypeResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByDeviceTypeResponseResultTop0]
type radarHTTPTopLocationListByDeviceTypeResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByDeviceTypeResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByDeviceTypeResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPProtocolResponse struct {
	Result  RadarHTTPTopLocationListByHTTPProtocolResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarHTTPTopLocationListByHTTPProtocolResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByHTTPProtocolResponse]
type radarHTTPTopLocationListByHTTPProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPProtocolResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByHTTPProtocolResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByHTTPProtocolResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByHTTPProtocolResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByHTTPProtocolResponseResult]
type radarHTTPTopLocationListByHTTPProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByHTTPProtocolResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByHTTPProtocolResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultMeta]
type radarHTTPTopLocationListByHTTPProtocolResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                      `json:"level,required"`
	JSON  radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                            `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRange]
type radarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnit struct {
	Name  string                                                           `json:"name,required"`
	Value string                                                           `json:"value,required"`
	JSON  radarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnit]
type radarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPProtocolResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                       `json:"value,required"`
	JSON  radarHTTPTopLocationListByHTTPProtocolResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPProtocolResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByHTTPProtocolResponseResultTop0]
type radarHTTPTopLocationListByHTTPProtocolResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPProtocolResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPProtocolResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPVersionResponse struct {
	Result  RadarHTTPTopLocationListByHTTPVersionResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarHTTPTopLocationListByHTTPVersionResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByHTTPVersionResponse]
type radarHTTPTopLocationListByHTTPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPVersionResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByHTTPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByHTTPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByHTTPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByHTTPVersionResponseResult]
type radarHTTPTopLocationListByHTTPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByHTTPVersionResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByHTTPVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByHTTPVersionResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultMeta]
type radarHTTPTopLocationListByHTTPVersionResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                     `json:"level,required"`
	JSON  radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                           `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRange]
type radarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByHTTPVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionResponseResultMetaUnit struct {
	Name  string                                                          `json:"name,required"`
	Value string                                                          `json:"value,required"`
	JSON  radarHTTPTopLocationListByHTTPVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultMetaUnitJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultMetaUnit]
type radarHTTPTopLocationListByHTTPVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByHTTPVersionResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                      `json:"value,required"`
	JSON  radarHTTPTopLocationListByHTTPVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByHTTPVersionResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByHTTPVersionResponseResultTop0]
type radarHTTPTopLocationListByHTTPVersionResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByHTTPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByHTTPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByIPVersionResponse struct {
	Result  RadarHTTPTopLocationListByIPVersionResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarHTTPTopLocationListByIPVersionResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByIPVersionResponse]
type radarHTTPTopLocationListByIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByIPVersionResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByIPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByIPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByIPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByIPVersionResponseResult]
type radarHTTPTopLocationListByIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByIPVersionResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByIPVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByIPVersionResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByIPVersionResponseResultMeta]
type radarHTTPTopLocationListByIPVersionResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                   `json:"level,required"`
	JSON  radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                         `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByIPVersionResponseResultMetaDateRange]
type radarHTTPTopLocationListByIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByIPVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionResponseResultMetaUnit struct {
	Name  string                                                        `json:"name,required"`
	Value string                                                        `json:"value,required"`
	JSON  radarHTTPTopLocationListByIPVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByIPVersionResponseResultMetaUnit]
type radarHTTPTopLocationListByIPVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByIPVersionResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                    `json:"value,required"`
	JSON  radarHTTPTopLocationListByIPVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByIPVersionResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByIPVersionResponseResultTop0]
type radarHTTPTopLocationListByIPVersionResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByIPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByIPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByOsResponse struct {
	Result  RadarHTTPTopLocationListByOsResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPTopLocationListByOsResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByOsResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationListByOsResponse]
type radarHTTPTopLocationListByOsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByOsResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByOsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByOsResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByOsResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByOsResponseResult]
type radarHTTPTopLocationListByOsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByOsResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByOsResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByOsResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByOsResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByOsResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByOsResponseResultMeta]
type radarHTTPTopLocationListByOsResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                            `json:"level,required"`
	JSON  radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                  `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByOsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByOsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByOsResponseResultMetaDateRange]
type radarHTTPTopLocationListByOsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByOsResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByOsResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByOsResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByOsResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsResponseResultMetaUnit struct {
	Name  string                                                 `json:"name,required"`
	Value string                                                 `json:"value,required"`
	JSON  radarHTTPTopLocationListByOsResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultMetaUnitJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByOsResponseResultMetaUnit]
type radarHTTPTopLocationListByOsResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByOsResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                             `json:"value,required"`
	JSON  radarHTTPTopLocationListByOsResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByOsResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationListByOsResponseResultTop0]
type radarHTTPTopLocationListByOsResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByOsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByOsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByTlsVersionResponse struct {
	Result  RadarHTTPTopLocationListByTlsVersionResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPTopLocationListByTlsVersionResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListByTlsVersionResponse]
type radarHTTPTopLocationListByTlsVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByTlsVersionResponseResult struct {
	// Metadata for the results.
	Meta RadarHTTPTopLocationListByTlsVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListByTlsVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListByTlsVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByTlsVersionResponseResult]
type radarHTTPTopLocationListByTlsVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RadarHTTPTopLocationListByTlsVersionResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RadarHTTPTopLocationListByTlsVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RadarHTTPTopLocationListByTlsVersionResponseResultMetaUnit `json:"units,required"`
	JSON  radarHTTPTopLocationListByTlsVersionResponseResultMetaJSON   `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByTlsVersionResponseResultMeta]
type radarHTTPTopLocationListByTlsVersionResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                    `json:"level,required"`
	JSON  radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                          `json:"startDate,required" format:"date-time"`
	JSON            radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByTlsVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListByTlsVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationListByTlsVersionResponseResultMetaDateRange]
type radarHTTPTopLocationListByTlsVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization string

const (
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationPercentage           RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "PERCENTAGE"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationMin0Max              RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "MIN0_MAX"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationMinMax               RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "MIN_MAX"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRawValues            RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "RAW_VALUES"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationPercentageChange     RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "PERCENTAGE_CHANGE"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRollingAverage       RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "ROLLING_AVERAGE"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationOverlappedPercentage RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRatio                RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization = "RATIO"
)

func (r RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalization) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationPercentage, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationMin0Max, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationMinMax, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRawValues, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationPercentageChange, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRollingAverage, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationOverlappedPercentage, RadarHTTPTopLocationListByTlsVersionResponseResultMetaNormalizationRatio:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionResponseResultMetaUnit struct {
	Name  string                                                         `json:"name,required"`
	Value string                                                         `json:"value,required"`
	JSON  radarHTTPTopLocationListByTlsVersionResponseResultMetaUnitJSON `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultMetaUnitJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListByTlsVersionResponseResultMetaUnit]
type radarHTTPTopLocationListByTlsVersionResponseResultMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListByTlsVersionResponseResultTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                                     `json:"value,required"`
	JSON  radarHTTPTopLocationListByTlsVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListByTlsVersionResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListByTlsVersionResponseResultTop0]
type radarHTTPTopLocationListByTlsVersionResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListByTlsVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopLocationListByTlsVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopLocationListParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationListParamsBotClass string

const (
	RadarHTTPTopLocationListParamsBotClassLikelyAutomated RadarHTTPTopLocationListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListParamsBotClassLikelyHuman     RadarHTTPTopLocationListParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsBotClassLikelyAutomated, RadarHTTPTopLocationListParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsBrowserFamily string

const (
	RadarHTTPTopLocationListParamsBrowserFamilyChrome  RadarHTTPTopLocationListParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListParamsBrowserFamilyEdge    RadarHTTPTopLocationListParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListParamsBrowserFamilyFirefox RadarHTTPTopLocationListParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListParamsBrowserFamilySafari  RadarHTTPTopLocationListParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsBrowserFamilyChrome, RadarHTTPTopLocationListParamsBrowserFamilyEdge, RadarHTTPTopLocationListParamsBrowserFamilyFirefox, RadarHTTPTopLocationListParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsDeviceType string

const (
	RadarHTTPTopLocationListParamsDeviceTypeDesktop RadarHTTPTopLocationListParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListParamsDeviceTypeMobile  RadarHTTPTopLocationListParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListParamsDeviceTypeOther   RadarHTTPTopLocationListParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsDeviceTypeDesktop, RadarHTTPTopLocationListParamsDeviceTypeMobile, RadarHTTPTopLocationListParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListParamsFormat string

const (
	RadarHTTPTopLocationListParamsFormatJson RadarHTTPTopLocationListParamsFormat = "JSON"
	RadarHTTPTopLocationListParamsFormatCsv  RadarHTTPTopLocationListParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsFormatJson, RadarHTTPTopLocationListParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListParamsHTTPProtocolHTTP  RadarHTTPTopLocationListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListParamsHTTPProtocolHTTPS RadarHTTPTopLocationListParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsHTTPProtocolHTTP, RadarHTTPTopLocationListParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsHTTPVersion string

const (
	RadarHTTPTopLocationListParamsHTTPVersionHttPv1 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListParamsHTTPVersionHttPv2 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListParamsHTTPVersionHttPv3 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsHTTPVersionHttPv1, RadarHTTPTopLocationListParamsHTTPVersionHttPv2, RadarHTTPTopLocationListParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsIPVersion string

const (
	RadarHTTPTopLocationListParamsIPVersionIPv4 RadarHTTPTopLocationListParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListParamsIPVersionIPv6 RadarHTTPTopLocationListParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsIPVersionIPv4, RadarHTTPTopLocationListParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsO string

const (
	RadarHTTPTopLocationListParamsOWindows  RadarHTTPTopLocationListParamsO = "WINDOWS"
	RadarHTTPTopLocationListParamsOMacosx   RadarHTTPTopLocationListParamsO = "MACOSX"
	RadarHTTPTopLocationListParamsOIos      RadarHTTPTopLocationListParamsO = "IOS"
	RadarHTTPTopLocationListParamsOAndroid  RadarHTTPTopLocationListParamsO = "ANDROID"
	RadarHTTPTopLocationListParamsOChromeos RadarHTTPTopLocationListParamsO = "CHROMEOS"
	RadarHTTPTopLocationListParamsOLinux    RadarHTTPTopLocationListParamsO = "LINUX"
	RadarHTTPTopLocationListParamsOSmartTv  RadarHTTPTopLocationListParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsOWindows, RadarHTTPTopLocationListParamsOMacosx, RadarHTTPTopLocationListParamsOIos, RadarHTTPTopLocationListParamsOAndroid, RadarHTTPTopLocationListParamsOChromeos, RadarHTTPTopLocationListParamsOLinux, RadarHTTPTopLocationListParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListParamsTlsVersion string

const (
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListParamsTlsVersionTlSvQuic RadarHTTPTopLocationListParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByBotClassParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByBotClassParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByBotClassParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByBotClassParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByBotClassParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByBotClassParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByBotClassParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bot class. Refer to
// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
type RadarHTTPTopLocationListByBotClassParamsBotClass string

const (
	RadarHTTPTopLocationListByBotClassParamsBotClassLikelyAutomated RadarHTTPTopLocationListByBotClassParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByBotClassParamsBotClassLikelyHuman     RadarHTTPTopLocationListByBotClassParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByBotClassParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByBotClassParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByBotClassParamsBrowserFamilyChrome  RadarHTTPTopLocationListByBotClassParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByBotClassParamsBrowserFamilyEdge    RadarHTTPTopLocationListByBotClassParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByBotClassParamsBrowserFamilyFirefox RadarHTTPTopLocationListByBotClassParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByBotClassParamsBrowserFamilySafari  RadarHTTPTopLocationListByBotClassParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByBotClassParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsBrowserFamilyChrome, RadarHTTPTopLocationListByBotClassParamsBrowserFamilyEdge, RadarHTTPTopLocationListByBotClassParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByBotClassParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsDeviceType string

const (
	RadarHTTPTopLocationListByBotClassParamsDeviceTypeDesktop RadarHTTPTopLocationListByBotClassParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByBotClassParamsDeviceTypeMobile  RadarHTTPTopLocationListByBotClassParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByBotClassParamsDeviceTypeOther   RadarHTTPTopLocationListByBotClassParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByBotClassParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsDeviceTypeDesktop, RadarHTTPTopLocationListByBotClassParamsDeviceTypeMobile, RadarHTTPTopLocationListByBotClassParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByBotClassParamsFormat string

const (
	RadarHTTPTopLocationListByBotClassParamsFormatJson RadarHTTPTopLocationListByBotClassParamsFormat = "JSON"
	RadarHTTPTopLocationListByBotClassParamsFormatCsv  RadarHTTPTopLocationListByBotClassParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByBotClassParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsFormatJson, RadarHTTPTopLocationListByBotClassParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByBotClassParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByBotClassParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByBotClassParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByBotClassParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByBotClassParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByBotClassParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByBotClassParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByBotClassParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByBotClassParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByBotClassParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByBotClassParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsIPVersion string

const (
	RadarHTTPTopLocationListByBotClassParamsIPVersionIPv4 RadarHTTPTopLocationListByBotClassParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByBotClassParamsIPVersionIPv6 RadarHTTPTopLocationListByBotClassParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByBotClassParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsIPVersionIPv4, RadarHTTPTopLocationListByBotClassParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsO string

const (
	RadarHTTPTopLocationListByBotClassParamsOWindows  RadarHTTPTopLocationListByBotClassParamsO = "WINDOWS"
	RadarHTTPTopLocationListByBotClassParamsOMacosx   RadarHTTPTopLocationListByBotClassParamsO = "MACOSX"
	RadarHTTPTopLocationListByBotClassParamsOIos      RadarHTTPTopLocationListByBotClassParamsO = "IOS"
	RadarHTTPTopLocationListByBotClassParamsOAndroid  RadarHTTPTopLocationListByBotClassParamsO = "ANDROID"
	RadarHTTPTopLocationListByBotClassParamsOChromeos RadarHTTPTopLocationListByBotClassParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByBotClassParamsOLinux    RadarHTTPTopLocationListByBotClassParamsO = "LINUX"
	RadarHTTPTopLocationListByBotClassParamsOSmartTv  RadarHTTPTopLocationListByBotClassParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByBotClassParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsOWindows, RadarHTTPTopLocationListByBotClassParamsOMacosx, RadarHTTPTopLocationListByBotClassParamsOIos, RadarHTTPTopLocationListByBotClassParamsOAndroid, RadarHTTPTopLocationListByBotClassParamsOChromeos, RadarHTTPTopLocationListByBotClassParamsOLinux, RadarHTTPTopLocationListByBotClassParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBotClassParamsTlsVersion string

const (
	RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByBotClassParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByBotClassParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByBotClassParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByBotClassParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByBotClassParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByBotClassParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByBotClassParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByBrowserFamilyParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByBrowserFamilyParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Browser family.
type RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyChrome  RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyEdge    RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyFirefox RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilySafari  RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyChrome, RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyEdge, RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByBrowserFamilyParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsBotClass string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsBotClassLikelyAutomated RadarHTTPTopLocationListByBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByBrowserFamilyParamsBotClassLikelyHuman     RadarHTTPTopLocationListByBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeDesktop RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeMobile  RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeOther   RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeDesktop, RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeMobile, RadarHTTPTopLocationListByBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByBrowserFamilyParamsFormat string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsFormatJson RadarHTTPTopLocationListByBrowserFamilyParamsFormat = "JSON"
	RadarHTTPTopLocationListByBrowserFamilyParamsFormatCsv  RadarHTTPTopLocationListByBrowserFamilyParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsFormatJson, RadarHTTPTopLocationListByBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsIPVersionIPv4 RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByBrowserFamilyParamsIPVersionIPv6 RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsIPVersionIPv4, RadarHTTPTopLocationListByBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsO string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsOWindows  RadarHTTPTopLocationListByBrowserFamilyParamsO = "WINDOWS"
	RadarHTTPTopLocationListByBrowserFamilyParamsOMacosx   RadarHTTPTopLocationListByBrowserFamilyParamsO = "MACOSX"
	RadarHTTPTopLocationListByBrowserFamilyParamsOIos      RadarHTTPTopLocationListByBrowserFamilyParamsO = "IOS"
	RadarHTTPTopLocationListByBrowserFamilyParamsOAndroid  RadarHTTPTopLocationListByBrowserFamilyParamsO = "ANDROID"
	RadarHTTPTopLocationListByBrowserFamilyParamsOChromeos RadarHTTPTopLocationListByBrowserFamilyParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByBrowserFamilyParamsOLinux    RadarHTTPTopLocationListByBrowserFamilyParamsO = "LINUX"
	RadarHTTPTopLocationListByBrowserFamilyParamsOSmartTv  RadarHTTPTopLocationListByBrowserFamilyParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsOWindows, RadarHTTPTopLocationListByBrowserFamilyParamsOMacosx, RadarHTTPTopLocationListByBrowserFamilyParamsOIos, RadarHTTPTopLocationListByBrowserFamilyParamsOAndroid, RadarHTTPTopLocationListByBrowserFamilyParamsOChromeos, RadarHTTPTopLocationListByBrowserFamilyParamsOLinux, RadarHTTPTopLocationListByBrowserFamilyParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion string

const (
	RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByBrowserFamilyParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily] `query:"browserFamily"`
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
	Format param.Field[RadarHTTPTopLocationListByDeviceTypeParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByDeviceTypeParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Device type.
type RadarHTTPTopLocationListByDeviceTypeParamsDeviceType string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeDesktop RadarHTTPTopLocationListByDeviceTypeParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeMobile  RadarHTTPTopLocationListByDeviceTypeParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeOther   RadarHTTPTopLocationListByDeviceTypeParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeDesktop, RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeMobile, RadarHTTPTopLocationListByDeviceTypeParamsDeviceTypeOther:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsBotClass string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsBotClassLikelyAutomated RadarHTTPTopLocationListByDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByDeviceTypeParamsBotClassLikelyHuman     RadarHTTPTopLocationListByDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByDeviceTypeParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyChrome  RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyEdge    RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyFirefox RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilySafari  RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyChrome, RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyEdge, RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByDeviceTypeParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByDeviceTypeParamsFormat string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsFormatJson RadarHTTPTopLocationListByDeviceTypeParamsFormat = "JSON"
	RadarHTTPTopLocationListByDeviceTypeParamsFormatCsv  RadarHTTPTopLocationListByDeviceTypeParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsFormatJson, RadarHTTPTopLocationListByDeviceTypeParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByDeviceTypeParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByDeviceTypeParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsIPVersion string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsIPVersionIPv4 RadarHTTPTopLocationListByDeviceTypeParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByDeviceTypeParamsIPVersionIPv6 RadarHTTPTopLocationListByDeviceTypeParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsIPVersionIPv4, RadarHTTPTopLocationListByDeviceTypeParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsO string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsOWindows  RadarHTTPTopLocationListByDeviceTypeParamsO = "WINDOWS"
	RadarHTTPTopLocationListByDeviceTypeParamsOMacosx   RadarHTTPTopLocationListByDeviceTypeParamsO = "MACOSX"
	RadarHTTPTopLocationListByDeviceTypeParamsOIos      RadarHTTPTopLocationListByDeviceTypeParamsO = "IOS"
	RadarHTTPTopLocationListByDeviceTypeParamsOAndroid  RadarHTTPTopLocationListByDeviceTypeParamsO = "ANDROID"
	RadarHTTPTopLocationListByDeviceTypeParamsOChromeos RadarHTTPTopLocationListByDeviceTypeParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByDeviceTypeParamsOLinux    RadarHTTPTopLocationListByDeviceTypeParamsO = "LINUX"
	RadarHTTPTopLocationListByDeviceTypeParamsOSmartTv  RadarHTTPTopLocationListByDeviceTypeParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsOWindows, RadarHTTPTopLocationListByDeviceTypeParamsOMacosx, RadarHTTPTopLocationListByDeviceTypeParamsOIos, RadarHTTPTopLocationListByDeviceTypeParamsOAndroid, RadarHTTPTopLocationListByDeviceTypeParamsOChromeos, RadarHTTPTopLocationListByDeviceTypeParamsOLinux, RadarHTTPTopLocationListByDeviceTypeParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion string

const (
	RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByDeviceTypeParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByDeviceTypeParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByHTTPProtocolParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByHTTPProtocolParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP protocol (HTTP vs. HTTPS).
type RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByHTTPProtocolParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsBotClass string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsBotClassLikelyAutomated RadarHTTPTopLocationListByHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByHTTPProtocolParamsBotClassLikelyHuman     RadarHTTPTopLocationListByHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByHTTPProtocolParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyChrome  RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyEdge    RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyFirefox RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilySafari  RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyChrome, RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyEdge, RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByHTTPProtocolParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeDesktop RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeMobile  RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeOther   RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeDesktop, RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeMobile, RadarHTTPTopLocationListByHTTPProtocolParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByHTTPProtocolParamsFormat string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsFormatJson RadarHTTPTopLocationListByHTTPProtocolParamsFormat = "JSON"
	RadarHTTPTopLocationListByHTTPProtocolParamsFormatCsv  RadarHTTPTopLocationListByHTTPProtocolParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsFormatJson, RadarHTTPTopLocationListByHTTPProtocolParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByHTTPProtocolParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsIPVersionIPv4 RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByHTTPProtocolParamsIPVersionIPv6 RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsIPVersionIPv4, RadarHTTPTopLocationListByHTTPProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsO string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsOWindows  RadarHTTPTopLocationListByHTTPProtocolParamsO = "WINDOWS"
	RadarHTTPTopLocationListByHTTPProtocolParamsOMacosx   RadarHTTPTopLocationListByHTTPProtocolParamsO = "MACOSX"
	RadarHTTPTopLocationListByHTTPProtocolParamsOIos      RadarHTTPTopLocationListByHTTPProtocolParamsO = "IOS"
	RadarHTTPTopLocationListByHTTPProtocolParamsOAndroid  RadarHTTPTopLocationListByHTTPProtocolParamsO = "ANDROID"
	RadarHTTPTopLocationListByHTTPProtocolParamsOChromeos RadarHTTPTopLocationListByHTTPProtocolParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByHTTPProtocolParamsOLinux    RadarHTTPTopLocationListByHTTPProtocolParamsO = "LINUX"
	RadarHTTPTopLocationListByHTTPProtocolParamsOSmartTv  RadarHTTPTopLocationListByHTTPProtocolParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsOWindows, RadarHTTPTopLocationListByHTTPProtocolParamsOMacosx, RadarHTTPTopLocationListByHTTPProtocolParamsOIos, RadarHTTPTopLocationListByHTTPProtocolParamsOAndroid, RadarHTTPTopLocationListByHTTPProtocolParamsOChromeos, RadarHTTPTopLocationListByHTTPProtocolParamsOLinux, RadarHTTPTopLocationListByHTTPProtocolParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion string

const (
	RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByHTTPProtocolParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByHTTPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByHTTPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsBotClass string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsBotClassLikelyAutomated RadarHTTPTopLocationListByHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByHTTPVersionParamsBotClassLikelyHuman     RadarHTTPTopLocationListByHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByHTTPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyChrome  RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyEdge    RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyFirefox RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilySafari  RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyChrome, RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyEdge, RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByHTTPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsDeviceType string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeDesktop RadarHTTPTopLocationListByHTTPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeMobile  RadarHTTPTopLocationListByHTTPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeOther   RadarHTTPTopLocationListByHTTPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeDesktop, RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeMobile, RadarHTTPTopLocationListByHTTPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByHTTPVersionParamsFormat string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsFormatJson RadarHTTPTopLocationListByHTTPVersionParamsFormat = "JSON"
	RadarHTTPTopLocationListByHTTPVersionParamsFormatCsv  RadarHTTPTopLocationListByHTTPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsFormatJson, RadarHTTPTopLocationListByHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByHTTPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsIPVersion string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsIPVersionIPv4 RadarHTTPTopLocationListByHTTPVersionParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByHTTPVersionParamsIPVersionIPv6 RadarHTTPTopLocationListByHTTPVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsIPVersionIPv4, RadarHTTPTopLocationListByHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsO string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsOWindows  RadarHTTPTopLocationListByHTTPVersionParamsO = "WINDOWS"
	RadarHTTPTopLocationListByHTTPVersionParamsOMacosx   RadarHTTPTopLocationListByHTTPVersionParamsO = "MACOSX"
	RadarHTTPTopLocationListByHTTPVersionParamsOIos      RadarHTTPTopLocationListByHTTPVersionParamsO = "IOS"
	RadarHTTPTopLocationListByHTTPVersionParamsOAndroid  RadarHTTPTopLocationListByHTTPVersionParamsO = "ANDROID"
	RadarHTTPTopLocationListByHTTPVersionParamsOChromeos RadarHTTPTopLocationListByHTTPVersionParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByHTTPVersionParamsOLinux    RadarHTTPTopLocationListByHTTPVersionParamsO = "LINUX"
	RadarHTTPTopLocationListByHTTPVersionParamsOSmartTv  RadarHTTPTopLocationListByHTTPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsOWindows, RadarHTTPTopLocationListByHTTPVersionParamsOMacosx, RadarHTTPTopLocationListByHTTPVersionParamsOIos, RadarHTTPTopLocationListByHTTPVersionParamsOAndroid, RadarHTTPTopLocationListByHTTPVersionParamsOChromeos, RadarHTTPTopLocationListByHTTPVersionParamsOLinux, RadarHTTPTopLocationListByHTTPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion string

const (
	RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByHTTPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByHTTPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByIPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByIPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByIPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByIPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByIPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByIPVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPTopLocationListByIPVersionParamsIPVersion string

const (
	RadarHTTPTopLocationListByIPVersionParamsIPVersionIPv4 RadarHTTPTopLocationListByIPVersionParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByIPVersionParamsIPVersionIPv6 RadarHTTPTopLocationListByIPVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsIPVersionIPv4, RadarHTTPTopLocationListByIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsBotClass string

const (
	RadarHTTPTopLocationListByIPVersionParamsBotClassLikelyAutomated RadarHTTPTopLocationListByIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByIPVersionParamsBotClassLikelyHuman     RadarHTTPTopLocationListByIPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByIPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByIPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyChrome  RadarHTTPTopLocationListByIPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyEdge    RadarHTTPTopLocationListByIPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyFirefox RadarHTTPTopLocationListByIPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByIPVersionParamsBrowserFamilySafari  RadarHTTPTopLocationListByIPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByIPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyChrome, RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyEdge, RadarHTTPTopLocationListByIPVersionParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByIPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsDeviceType string

const (
	RadarHTTPTopLocationListByIPVersionParamsDeviceTypeDesktop RadarHTTPTopLocationListByIPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByIPVersionParamsDeviceTypeMobile  RadarHTTPTopLocationListByIPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByIPVersionParamsDeviceTypeOther   RadarHTTPTopLocationListByIPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByIPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsDeviceTypeDesktop, RadarHTTPTopLocationListByIPVersionParamsDeviceTypeMobile, RadarHTTPTopLocationListByIPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByIPVersionParamsFormat string

const (
	RadarHTTPTopLocationListByIPVersionParamsFormatJson RadarHTTPTopLocationListByIPVersionParamsFormat = "JSON"
	RadarHTTPTopLocationListByIPVersionParamsFormatCsv  RadarHTTPTopLocationListByIPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsFormatJson, RadarHTTPTopLocationListByIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByIPVersionParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByIPVersionParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByIPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByIPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByIPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByIPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsO string

const (
	RadarHTTPTopLocationListByIPVersionParamsOWindows  RadarHTTPTopLocationListByIPVersionParamsO = "WINDOWS"
	RadarHTTPTopLocationListByIPVersionParamsOMacosx   RadarHTTPTopLocationListByIPVersionParamsO = "MACOSX"
	RadarHTTPTopLocationListByIPVersionParamsOIos      RadarHTTPTopLocationListByIPVersionParamsO = "IOS"
	RadarHTTPTopLocationListByIPVersionParamsOAndroid  RadarHTTPTopLocationListByIPVersionParamsO = "ANDROID"
	RadarHTTPTopLocationListByIPVersionParamsOChromeos RadarHTTPTopLocationListByIPVersionParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByIPVersionParamsOLinux    RadarHTTPTopLocationListByIPVersionParamsO = "LINUX"
	RadarHTTPTopLocationListByIPVersionParamsOSmartTv  RadarHTTPTopLocationListByIPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByIPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsOWindows, RadarHTTPTopLocationListByIPVersionParamsOMacosx, RadarHTTPTopLocationListByIPVersionParamsOIos, RadarHTTPTopLocationListByIPVersionParamsOAndroid, RadarHTTPTopLocationListByIPVersionParamsOChromeos, RadarHTTPTopLocationListByIPVersionParamsOLinux, RadarHTTPTopLocationListByIPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByIPVersionParamsTlsVersion string

const (
	RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByIPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByIPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByIPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByIPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByIPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByIPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByIPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByOsParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByOsParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByOsParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByOsParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByOsParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByOsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByOsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListByOsParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListByOsParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopLocationListByOsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operating system.
type RadarHTTPTopLocationListByOsParamsOs string

const (
	RadarHTTPTopLocationListByOsParamsOsWindows  RadarHTTPTopLocationListByOsParamsOs = "WINDOWS"
	RadarHTTPTopLocationListByOsParamsOsMacosx   RadarHTTPTopLocationListByOsParamsOs = "MACOSX"
	RadarHTTPTopLocationListByOsParamsOsIos      RadarHTTPTopLocationListByOsParamsOs = "IOS"
	RadarHTTPTopLocationListByOsParamsOsAndroid  RadarHTTPTopLocationListByOsParamsOs = "ANDROID"
	RadarHTTPTopLocationListByOsParamsOsChromeos RadarHTTPTopLocationListByOsParamsOs = "CHROMEOS"
	RadarHTTPTopLocationListByOsParamsOsLinux    RadarHTTPTopLocationListByOsParamsOs = "LINUX"
	RadarHTTPTopLocationListByOsParamsOsSmartTv  RadarHTTPTopLocationListByOsParamsOs = "SMART_TV"
)

func (r RadarHTTPTopLocationListByOsParamsOs) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsOsWindows, RadarHTTPTopLocationListByOsParamsOsMacosx, RadarHTTPTopLocationListByOsParamsOsIos, RadarHTTPTopLocationListByOsParamsOsAndroid, RadarHTTPTopLocationListByOsParamsOsChromeos, RadarHTTPTopLocationListByOsParamsOsLinux, RadarHTTPTopLocationListByOsParamsOsSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsBotClass string

const (
	RadarHTTPTopLocationListByOsParamsBotClassLikelyAutomated RadarHTTPTopLocationListByOsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByOsParamsBotClassLikelyHuman     RadarHTTPTopLocationListByOsParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByOsParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByOsParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByOsParamsBrowserFamilyChrome  RadarHTTPTopLocationListByOsParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByOsParamsBrowserFamilyEdge    RadarHTTPTopLocationListByOsParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByOsParamsBrowserFamilyFirefox RadarHTTPTopLocationListByOsParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByOsParamsBrowserFamilySafari  RadarHTTPTopLocationListByOsParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByOsParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsBrowserFamilyChrome, RadarHTTPTopLocationListByOsParamsBrowserFamilyEdge, RadarHTTPTopLocationListByOsParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByOsParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsDeviceType string

const (
	RadarHTTPTopLocationListByOsParamsDeviceTypeDesktop RadarHTTPTopLocationListByOsParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByOsParamsDeviceTypeMobile  RadarHTTPTopLocationListByOsParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByOsParamsDeviceTypeOther   RadarHTTPTopLocationListByOsParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByOsParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsDeviceTypeDesktop, RadarHTTPTopLocationListByOsParamsDeviceTypeMobile, RadarHTTPTopLocationListByOsParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByOsParamsFormat string

const (
	RadarHTTPTopLocationListByOsParamsFormatJson RadarHTTPTopLocationListByOsParamsFormat = "JSON"
	RadarHTTPTopLocationListByOsParamsFormatCsv  RadarHTTPTopLocationListByOsParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByOsParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsFormatJson, RadarHTTPTopLocationListByOsParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByOsParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByOsParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByOsParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByOsParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByOsParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByOsParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByOsParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByOsParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByOsParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByOsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByOsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsIPVersion string

const (
	RadarHTTPTopLocationListByOsParamsIPVersionIPv4 RadarHTTPTopLocationListByOsParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByOsParamsIPVersionIPv6 RadarHTTPTopLocationListByOsParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByOsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsIPVersionIPv4, RadarHTTPTopLocationListByOsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByOsParamsTlsVersion string

const (
	RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByOsParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByOsParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByOsParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByOsParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByOsParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByOsParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByOsParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByOsParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByOsParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopLocationListByTlsVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopLocationListByTlsVersionParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTopLocationListByTlsVersionParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationListByTlsVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type RadarHTTPTopLocationListByTlsVersionParamsTlsVersion string

const (
	RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListByTlsVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListByTlsVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListByTlsVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListByTlsVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSvQuic RadarHTTPTopLocationListByTlsVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_0, RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_1, RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_2, RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSv1_3, RadarHTTPTopLocationListByTlsVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsBotClass string

const (
	RadarHTTPTopLocationListByTlsVersionParamsBotClassLikelyAutomated RadarHTTPTopLocationListByTlsVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListByTlsVersionParamsBotClassLikelyHuman     RadarHTTPTopLocationListByTlsVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsBotClassLikelyAutomated, RadarHTTPTopLocationListByTlsVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily string

const (
	RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyChrome  RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyEdge    RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyFirefox RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilySafari  RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyChrome, RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyEdge, RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilyFirefox, RadarHTTPTopLocationListByTlsVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsDeviceType string

const (
	RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeDesktop RadarHTTPTopLocationListByTlsVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeMobile  RadarHTTPTopLocationListByTlsVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeOther   RadarHTTPTopLocationListByTlsVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeDesktop, RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeMobile, RadarHTTPTopLocationListByTlsVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopLocationListByTlsVersionParamsFormat string

const (
	RadarHTTPTopLocationListByTlsVersionParamsFormatJson RadarHTTPTopLocationListByTlsVersionParamsFormat = "JSON"
	RadarHTTPTopLocationListByTlsVersionParamsFormatCsv  RadarHTTPTopLocationListByTlsVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsFormatJson, RadarHTTPTopLocationListByTlsVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocolHTTP  RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocolHTTPS RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocolHTTP, RadarHTTPTopLocationListByTlsVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion string

const (
	RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv1 RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv2 RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv3 RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv1, RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv2, RadarHTTPTopLocationListByTlsVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsIPVersion string

const (
	RadarHTTPTopLocationListByTlsVersionParamsIPVersionIPv4 RadarHTTPTopLocationListByTlsVersionParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListByTlsVersionParamsIPVersionIPv6 RadarHTTPTopLocationListByTlsVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsIPVersionIPv4, RadarHTTPTopLocationListByTlsVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopLocationListByTlsVersionParamsO string

const (
	RadarHTTPTopLocationListByTlsVersionParamsOWindows  RadarHTTPTopLocationListByTlsVersionParamsO = "WINDOWS"
	RadarHTTPTopLocationListByTlsVersionParamsOMacosx   RadarHTTPTopLocationListByTlsVersionParamsO = "MACOSX"
	RadarHTTPTopLocationListByTlsVersionParamsOIos      RadarHTTPTopLocationListByTlsVersionParamsO = "IOS"
	RadarHTTPTopLocationListByTlsVersionParamsOAndroid  RadarHTTPTopLocationListByTlsVersionParamsO = "ANDROID"
	RadarHTTPTopLocationListByTlsVersionParamsOChromeos RadarHTTPTopLocationListByTlsVersionParamsO = "CHROMEOS"
	RadarHTTPTopLocationListByTlsVersionParamsOLinux    RadarHTTPTopLocationListByTlsVersionParamsO = "LINUX"
	RadarHTTPTopLocationListByTlsVersionParamsOSmartTv  RadarHTTPTopLocationListByTlsVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopLocationListByTlsVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopLocationListByTlsVersionParamsOWindows, RadarHTTPTopLocationListByTlsVersionParamsOMacosx, RadarHTTPTopLocationListByTlsVersionParamsOIos, RadarHTTPTopLocationListByTlsVersionParamsOAndroid, RadarHTTPTopLocationListByTlsVersionParamsOChromeos, RadarHTTPTopLocationListByTlsVersionParamsOLinux, RadarHTTPTopLocationListByTlsVersionParamsOSmartTv:
		return true
	}
	return false
}
