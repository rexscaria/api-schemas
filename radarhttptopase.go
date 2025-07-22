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

// RadarHTTPTopAseService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTopAseService] method instead.
type RadarHTTPTopAseService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPTopAseService(opts ...option.RequestOption) (r *RadarHTTPTopAseService) {
	r = &RadarHTTPTopAseService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by HTTP requests.
func (r *RadarHTTPTopAseService) List(ctx context.Context, query RadarHTTPTopAseListParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested bot
// class.
func (r *RadarHTTPTopAseService) ListByBotClass(ctx context.Context, botClass RadarHTTPTopAseListByBotClassParamsBotClass, query RadarHTTPTopAseListByBotClassParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByBotClassResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested browser
// family.
func (r *RadarHTTPTopAseService) ListByBrowserFamily(ctx context.Context, browserFamily RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily, query RadarHTTPTopAseListByBrowserFamilyParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByBrowserFamilyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/browser_family/%v", browserFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested device
// type.
func (r *RadarHTTPTopAseService) ListByDeviceType(ctx context.Context, deviceType RadarHTTPTopAseListByDeviceTypeParamsDeviceType, query RadarHTTPTopAseListByDeviceTypeParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByDeviceTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested HTTP
// protocol.
func (r *RadarHTTPTopAseService) ListByHTTPProtocol(ctx context.Context, httpProtocol RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocol, query RadarHTTPTopAseListByHTTPProtocolParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByHTTPProtocolResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested HTTP
// version.
func (r *RadarHTTPTopAseService) ListByHTTPVersion(ctx context.Context, httpVersion RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion, query RadarHTTPTopAseListByHTTPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByHTTPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested IP
// version.
func (r *RadarHTTPTopAseService) ListByIPVersion(ctx context.Context, ipVersion RadarHTTPTopAseListByIPVersionParamsIPVersion, query RadarHTTPTopAseListByIPVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByIPVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested
// operating system.
func (r *RadarHTTPTopAseService) ListByOs(ctx context.Context, os RadarHTTPTopAseListByOsParamsOs, query RadarHTTPTopAseListByOsParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByOsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested TLS
// protocol version.
func (r *RadarHTTPTopAseService) ListByTlsVersion(ctx context.Context, tlsVersion RadarHTTPTopAseListByTlsVersionParamsTlsVersion, query RadarHTTPTopAseListByTlsVersionParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListByTlsVersionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseListResponse struct {
	Result  RadarHTTPTopAseListResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarHTTPTopAseListResponseJSON   `json:"-"`
}

// radarHTTPTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseListResponse]
type radarHTTPTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResult struct {
	Meta RadarHTTPTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseListResponseResult]
type radarHTTPTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListResponseResultMeta]
type radarHTTPTopAseListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListResponseResultMetaDateRange]
type radarHTTPTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListResponseResultTop0 struct {
	ClientAsn    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPTopAseListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListResponseResultTop0]
type radarHTTPTopAseListResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponse struct {
	Result  RadarHTTPTopAseListByBotClassResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPTopAseListByBotClassResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByBotClassResponse]
type radarHTTPTopAseListByBotClassResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResult struct {
	Meta RadarHTTPTopAseListByBotClassResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByBotClassResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByBotClassResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseListByBotClassResponseResult]
type radarHTTPTopAseListByBotClassResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByBotClassResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByBotClassResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByBotClassResponseResultMeta]
type radarHTTPTopAseListByBotClassResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByBotClassResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseListByBotClassResponseResultMetaDateRange]
type radarHTTPTopAseListByBotClassResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBotClassResponseResultTop0 struct {
	ClientAsn    int64                                               `json:"clientASN,required"`
	ClientAsName string                                              `json:"clientASName,required"`
	Value        string                                              `json:"value,required"`
	JSON         radarHTTPTopAseListByBotClassResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByBotClassResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByBotClassResponseResultTop0]
type radarHTTPTopAseListByBotClassResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBotClassResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBotClassResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponse struct {
	Result  RadarHTTPTopAseListByBrowserFamilyResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarHTTPTopAseListByBrowserFamilyResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseListByBrowserFamilyResponse]
type radarHTTPTopAseListByBrowserFamilyResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResult struct {
	Meta RadarHTTPTopAseListByBrowserFamilyResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByBrowserFamilyResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByBrowserFamilyResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByBrowserFamilyResponseResult]
type radarHTTPTopAseListByBrowserFamilyResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByBrowserFamilyResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByBrowserFamilyResponseResultMeta]
type radarHTTPTopAseListByBrowserFamilyResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRange]
type radarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByBrowserFamilyResponseResultTop0 struct {
	ClientAsn    int64                                                    `json:"clientASN,required"`
	ClientAsName string                                                   `json:"clientASName,required"`
	Value        string                                                   `json:"value,required"`
	JSON         radarHTTPTopAseListByBrowserFamilyResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByBrowserFamilyResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByBrowserFamilyResponseResultTop0]
type radarHTTPTopAseListByBrowserFamilyResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByBrowserFamilyResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByBrowserFamilyResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponse struct {
	Result  RadarHTTPTopAseListByDeviceTypeResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTopAseListByDeviceTypeResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByDeviceTypeResponse]
type radarHTTPTopAseListByDeviceTypeResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResult struct {
	Meta RadarHTTPTopAseListByDeviceTypeResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByDeviceTypeResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByDeviceTypeResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseListByDeviceTypeResponseResult]
type radarHTTPTopAseListByDeviceTypeResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByDeviceTypeResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByDeviceTypeResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByDeviceTypeResponseResultMeta]
type radarHTTPTopAseListByDeviceTypeResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByDeviceTypeResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseListByDeviceTypeResponseResultMetaDateRange]
type radarHTTPTopAseListByDeviceTypeResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByDeviceTypeResponseResultTop0 struct {
	ClientAsn    int64                                                 `json:"clientASN,required"`
	ClientAsName string                                                `json:"clientASName,required"`
	Value        string                                                `json:"value,required"`
	JSON         radarHTTPTopAseListByDeviceTypeResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByDeviceTypeResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByDeviceTypeResponseResultTop0]
type radarHTTPTopAseListByDeviceTypeResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByDeviceTypeResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByDeviceTypeResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponse struct {
	Result  RadarHTTPTopAseListByHTTPProtocolResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTopAseListByHTTPProtocolResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByHTTPProtocolResponse]
type radarHTTPTopAseListByHTTPProtocolResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResult struct {
	Meta RadarHTTPTopAseListByHTTPProtocolResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByHTTPProtocolResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByHTTPProtocolResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByHTTPProtocolResponseResult]
type radarHTTPTopAseListByHTTPProtocolResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByHTTPProtocolResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByHTTPProtocolResponseResultMeta]
type radarHTTPTopAseListByHTTPProtocolResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRange]
type radarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPProtocolResponseResultTop0 struct {
	ClientAsn    int64                                                   `json:"clientASN,required"`
	ClientAsName string                                                  `json:"clientASName,required"`
	Value        string                                                  `json:"value,required"`
	JSON         radarHTTPTopAseListByHTTPProtocolResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByHTTPProtocolResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByHTTPProtocolResponseResultTop0]
type radarHTTPTopAseListByHTTPProtocolResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPProtocolResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPProtocolResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponse struct {
	Result  RadarHTTPTopAseListByHTTPVersionResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarHTTPTopAseListByHTTPVersionResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByHTTPVersionResponse]
type radarHTTPTopAseListByHTTPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResult struct {
	Meta RadarHTTPTopAseListByHTTPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByHTTPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByHTTPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByHTTPVersionResponseResult]
type radarHTTPTopAseListByHTTPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByHTTPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByHTTPVersionResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByHTTPVersionResponseResultMeta]
type radarHTTPTopAseListByHTTPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByHTTPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPVersionResponseResultMetaDateRange]
type radarHTTPTopAseListByHTTPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByHTTPVersionResponseResultTop0 struct {
	ClientAsn    int64                                                  `json:"clientASN,required"`
	ClientAsName string                                                 `json:"clientASName,required"`
	Value        string                                                 `json:"value,required"`
	JSON         radarHTTPTopAseListByHTTPVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByHTTPVersionResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByHTTPVersionResponseResultTop0]
type radarHTTPTopAseListByHTTPVersionResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByHTTPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByHTTPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponse struct {
	Result  RadarHTTPTopAseListByIPVersionResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPTopAseListByIPVersionResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByIPVersionResponse]
type radarHTTPTopAseListByIPVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResult struct {
	Meta RadarHTTPTopAseListByIPVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByIPVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByIPVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseListByIPVersionResponseResult]
type radarHTTPTopAseListByIPVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByIPVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByIPVersionResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByIPVersionResponseResultMeta]
type radarHTTPTopAseListByIPVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByIPVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseListByIPVersionResponseResultMetaDateRange]
type radarHTTPTopAseListByIPVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByIPVersionResponseResultTop0 struct {
	ClientAsn    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         radarHTTPTopAseListByIPVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByIPVersionResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByIPVersionResponseResultTop0]
type radarHTTPTopAseListByIPVersionResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByIPVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByIPVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponse struct {
	Result  RadarHTTPTopAseListByOsResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarHTTPTopAseListByOsResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByOsResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseListByOsResponse]
type radarHTTPTopAseListByOsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResult struct {
	Meta RadarHTTPTopAseListByOsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByOsResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByOsResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByOsResponseResult]
type radarHTTPTopAseListByOsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByOsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByOsResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByOsResponseResultMeta]
type radarHTTPTopAseListByOsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByOsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListByOsResponseResultMetaDateRange]
type radarHTTPTopAseListByOsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByOsResponseResultTop0 struct {
	ClientAsn    int64                                         `json:"clientASN,required"`
	ClientAsName string                                        `json:"clientASName,required"`
	Value        string                                        `json:"value,required"`
	JSON         radarHTTPTopAseListByOsResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByOsResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByOsResponseResultTop0]
type radarHTTPTopAseListByOsResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByOsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByOsResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponse struct {
	Result  RadarHTTPTopAseListByTlsVersionResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTopAseListByTlsVersionResponseJSON   `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListByTlsVersionResponse]
type radarHTTPTopAseListByTlsVersionResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResult struct {
	Meta RadarHTTPTopAseListByTlsVersionResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListByTlsVersionResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListByTlsVersionResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseListByTlsVersionResponseResult]
type radarHTTPTopAseListByTlsVersionResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListByTlsVersionResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListByTlsVersionResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByTlsVersionResponseResultMeta]
type radarHTTPTopAseListByTlsVersionResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListByTlsVersionResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseListByTlsVersionResponseResultMetaDateRange]
type radarHTTPTopAseListByTlsVersionResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListByTlsVersionResponseResultTop0 struct {
	ClientAsn    int64                                                 `json:"clientASN,required"`
	ClientAsName string                                                `json:"clientASName,required"`
	Value        string                                                `json:"value,required"`
	JSON         radarHTTPTopAseListByTlsVersionResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListByTlsVersionResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListByTlsVersionResponseResultTop0]
type radarHTTPTopAseListByTlsVersionResponseResultTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListByTlsVersionResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPTopAseListByTlsVersionResponseResultTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPTopAseListParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopAseListParamsBotClass string

const (
	RadarHTTPTopAseListParamsBotClassLikelyAutomated RadarHTTPTopAseListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListParamsBotClassLikelyHuman     RadarHTTPTopAseListParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsBotClassLikelyAutomated, RadarHTTPTopAseListParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsBrowserFamily string

const (
	RadarHTTPTopAseListParamsBrowserFamilyChrome  RadarHTTPTopAseListParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListParamsBrowserFamilyEdge    RadarHTTPTopAseListParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListParamsBrowserFamilyFirefox RadarHTTPTopAseListParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListParamsBrowserFamilySafari  RadarHTTPTopAseListParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsBrowserFamilyChrome, RadarHTTPTopAseListParamsBrowserFamilyEdge, RadarHTTPTopAseListParamsBrowserFamilyFirefox, RadarHTTPTopAseListParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsDeviceType string

const (
	RadarHTTPTopAseListParamsDeviceTypeDesktop RadarHTTPTopAseListParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListParamsDeviceTypeMobile  RadarHTTPTopAseListParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListParamsDeviceTypeOther   RadarHTTPTopAseListParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsDeviceTypeDesktop, RadarHTTPTopAseListParamsDeviceTypeMobile, RadarHTTPTopAseListParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListParamsFormat string

const (
	RadarHTTPTopAseListParamsFormatJson RadarHTTPTopAseListParamsFormat = "JSON"
	RadarHTTPTopAseListParamsFormatCsv  RadarHTTPTopAseListParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsFormatJson, RadarHTTPTopAseListParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsHTTPProtocol string

const (
	RadarHTTPTopAseListParamsHTTPProtocolHTTP  RadarHTTPTopAseListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListParamsHTTPProtocolHTTPS RadarHTTPTopAseListParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsHTTPProtocolHTTP, RadarHTTPTopAseListParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsHTTPVersion string

const (
	RadarHTTPTopAseListParamsHTTPVersionHttPv1 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListParamsHTTPVersionHttPv2 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListParamsHTTPVersionHttPv3 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsHTTPVersionHttPv1, RadarHTTPTopAseListParamsHTTPVersionHttPv2, RadarHTTPTopAseListParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsIPVersion string

const (
	RadarHTTPTopAseListParamsIPVersionIPv4 RadarHTTPTopAseListParamsIPVersion = "IPv4"
	RadarHTTPTopAseListParamsIPVersionIPv6 RadarHTTPTopAseListParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsIPVersionIPv4, RadarHTTPTopAseListParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsO string

const (
	RadarHTTPTopAseListParamsOWindows  RadarHTTPTopAseListParamsO = "WINDOWS"
	RadarHTTPTopAseListParamsOMacosx   RadarHTTPTopAseListParamsO = "MACOSX"
	RadarHTTPTopAseListParamsOIos      RadarHTTPTopAseListParamsO = "IOS"
	RadarHTTPTopAseListParamsOAndroid  RadarHTTPTopAseListParamsO = "ANDROID"
	RadarHTTPTopAseListParamsOChromeos RadarHTTPTopAseListParamsO = "CHROMEOS"
	RadarHTTPTopAseListParamsOLinux    RadarHTTPTopAseListParamsO = "LINUX"
	RadarHTTPTopAseListParamsOSmartTv  RadarHTTPTopAseListParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsOWindows, RadarHTTPTopAseListParamsOMacosx, RadarHTTPTopAseListParamsOIos, RadarHTTPTopAseListParamsOAndroid, RadarHTTPTopAseListParamsOChromeos, RadarHTTPTopAseListParamsOLinux, RadarHTTPTopAseListParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListParamsTlsVersion string

const (
	RadarHTTPTopAseListParamsTlsVersionTlSv1_0  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_1  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_2  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_3  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListParamsTlsVersionTlSvQuic RadarHTTPTopAseListParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListParamsTlsVersionTlSv1_0, RadarHTTPTopAseListParamsTlsVersionTlSv1_1, RadarHTTPTopAseListParamsTlsVersionTlSv1_2, RadarHTTPTopAseListParamsTlsVersionTlSv1_3, RadarHTTPTopAseListParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByBotClassParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByBotClassParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByBotClassParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByBotClassParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByBotClassParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByBotClassParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByBotClassParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByBotClassParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByBotClassParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseListByBotClassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bot class. Refer to
// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
type RadarHTTPTopAseListByBotClassParamsBotClass string

const (
	RadarHTTPTopAseListByBotClassParamsBotClassLikelyAutomated RadarHTTPTopAseListByBotClassParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByBotClassParamsBotClassLikelyHuman     RadarHTTPTopAseListByBotClassParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByBotClassParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsBotClassLikelyAutomated, RadarHTTPTopAseListByBotClassParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsBrowserFamily string

const (
	RadarHTTPTopAseListByBotClassParamsBrowserFamilyChrome  RadarHTTPTopAseListByBotClassParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByBotClassParamsBrowserFamilyEdge    RadarHTTPTopAseListByBotClassParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByBotClassParamsBrowserFamilyFirefox RadarHTTPTopAseListByBotClassParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByBotClassParamsBrowserFamilySafari  RadarHTTPTopAseListByBotClassParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByBotClassParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsBrowserFamilyChrome, RadarHTTPTopAseListByBotClassParamsBrowserFamilyEdge, RadarHTTPTopAseListByBotClassParamsBrowserFamilyFirefox, RadarHTTPTopAseListByBotClassParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsDeviceType string

const (
	RadarHTTPTopAseListByBotClassParamsDeviceTypeDesktop RadarHTTPTopAseListByBotClassParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByBotClassParamsDeviceTypeMobile  RadarHTTPTopAseListByBotClassParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByBotClassParamsDeviceTypeOther   RadarHTTPTopAseListByBotClassParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByBotClassParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsDeviceTypeDesktop, RadarHTTPTopAseListByBotClassParamsDeviceTypeMobile, RadarHTTPTopAseListByBotClassParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByBotClassParamsFormat string

const (
	RadarHTTPTopAseListByBotClassParamsFormatJson RadarHTTPTopAseListByBotClassParamsFormat = "JSON"
	RadarHTTPTopAseListByBotClassParamsFormatCsv  RadarHTTPTopAseListByBotClassParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByBotClassParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsFormatJson, RadarHTTPTopAseListByBotClassParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByBotClassParamsHTTPProtocolHTTP  RadarHTTPTopAseListByBotClassParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByBotClassParamsHTTPProtocolHTTPS RadarHTTPTopAseListByBotClassParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByBotClassParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsHTTPProtocolHTTP, RadarHTTPTopAseListByBotClassParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsHTTPVersion string

const (
	RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv1 RadarHTTPTopAseListByBotClassParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv2 RadarHTTPTopAseListByBotClassParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv3 RadarHTTPTopAseListByBotClassParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByBotClassParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv1, RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv2, RadarHTTPTopAseListByBotClassParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsIPVersion string

const (
	RadarHTTPTopAseListByBotClassParamsIPVersionIPv4 RadarHTTPTopAseListByBotClassParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByBotClassParamsIPVersionIPv6 RadarHTTPTopAseListByBotClassParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByBotClassParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsIPVersionIPv4, RadarHTTPTopAseListByBotClassParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsO string

const (
	RadarHTTPTopAseListByBotClassParamsOWindows  RadarHTTPTopAseListByBotClassParamsO = "WINDOWS"
	RadarHTTPTopAseListByBotClassParamsOMacosx   RadarHTTPTopAseListByBotClassParamsO = "MACOSX"
	RadarHTTPTopAseListByBotClassParamsOIos      RadarHTTPTopAseListByBotClassParamsO = "IOS"
	RadarHTTPTopAseListByBotClassParamsOAndroid  RadarHTTPTopAseListByBotClassParamsO = "ANDROID"
	RadarHTTPTopAseListByBotClassParamsOChromeos RadarHTTPTopAseListByBotClassParamsO = "CHROMEOS"
	RadarHTTPTopAseListByBotClassParamsOLinux    RadarHTTPTopAseListByBotClassParamsO = "LINUX"
	RadarHTTPTopAseListByBotClassParamsOSmartTv  RadarHTTPTopAseListByBotClassParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByBotClassParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsOWindows, RadarHTTPTopAseListByBotClassParamsOMacosx, RadarHTTPTopAseListByBotClassParamsOIos, RadarHTTPTopAseListByBotClassParamsOAndroid, RadarHTTPTopAseListByBotClassParamsOChromeos, RadarHTTPTopAseListByBotClassParamsOLinux, RadarHTTPTopAseListByBotClassParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBotClassParamsTlsVersion string

const (
	RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByBotClassParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByBotClassParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByBotClassParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByBotClassParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByBotClassParamsTlsVersionTlSvQuic RadarHTTPTopAseListByBotClassParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByBotClassParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByBotClassParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByBotClassParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsBotClass] `query:"botClass"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByBrowserFamilyParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByBrowserFamilyParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopAseListByBrowserFamilyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Browser family.
type RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyChrome  RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyEdge    RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyFirefox RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilySafari  RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyChrome, RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyEdge, RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilyFirefox, RadarHTTPTopAseListByBrowserFamilyParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsBotClass string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsBotClassLikelyAutomated RadarHTTPTopAseListByBrowserFamilyParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByBrowserFamilyParamsBotClassLikelyHuman     RadarHTTPTopAseListByBrowserFamilyParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsBotClassLikelyAutomated, RadarHTTPTopAseListByBrowserFamilyParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsDeviceType string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeDesktop RadarHTTPTopAseListByBrowserFamilyParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeMobile  RadarHTTPTopAseListByBrowserFamilyParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeOther   RadarHTTPTopAseListByBrowserFamilyParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeDesktop, RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeMobile, RadarHTTPTopAseListByBrowserFamilyParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByBrowserFamilyParamsFormat string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsFormatJson RadarHTTPTopAseListByBrowserFamilyParamsFormat = "JSON"
	RadarHTTPTopAseListByBrowserFamilyParamsFormatCsv  RadarHTTPTopAseListByBrowserFamilyParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsFormatJson, RadarHTTPTopAseListByBrowserFamilyParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocolHTTP  RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocolHTTPS RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocolHTTP, RadarHTTPTopAseListByBrowserFamilyParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv1 RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv2 RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv3 RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv1, RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv2, RadarHTTPTopAseListByBrowserFamilyParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsIPVersion string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsIPVersionIPv4 RadarHTTPTopAseListByBrowserFamilyParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByBrowserFamilyParamsIPVersionIPv6 RadarHTTPTopAseListByBrowserFamilyParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsIPVersionIPv4, RadarHTTPTopAseListByBrowserFamilyParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsO string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsOWindows  RadarHTTPTopAseListByBrowserFamilyParamsO = "WINDOWS"
	RadarHTTPTopAseListByBrowserFamilyParamsOMacosx   RadarHTTPTopAseListByBrowserFamilyParamsO = "MACOSX"
	RadarHTTPTopAseListByBrowserFamilyParamsOIos      RadarHTTPTopAseListByBrowserFamilyParamsO = "IOS"
	RadarHTTPTopAseListByBrowserFamilyParamsOAndroid  RadarHTTPTopAseListByBrowserFamilyParamsO = "ANDROID"
	RadarHTTPTopAseListByBrowserFamilyParamsOChromeos RadarHTTPTopAseListByBrowserFamilyParamsO = "CHROMEOS"
	RadarHTTPTopAseListByBrowserFamilyParamsOLinux    RadarHTTPTopAseListByBrowserFamilyParamsO = "LINUX"
	RadarHTTPTopAseListByBrowserFamilyParamsOSmartTv  RadarHTTPTopAseListByBrowserFamilyParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsOWindows, RadarHTTPTopAseListByBrowserFamilyParamsOMacosx, RadarHTTPTopAseListByBrowserFamilyParamsOIos, RadarHTTPTopAseListByBrowserFamilyParamsOAndroid, RadarHTTPTopAseListByBrowserFamilyParamsOChromeos, RadarHTTPTopAseListByBrowserFamilyParamsOLinux, RadarHTTPTopAseListByBrowserFamilyParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion string

const (
	RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSvQuic RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByBrowserFamilyParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByBrowserFamilyParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily] `query:"browserFamily"`
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
	Format param.Field[RadarHTTPTopAseListByDeviceTypeParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByDeviceTypeParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByDeviceTypeParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopAseListByDeviceTypeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Device type.
type RadarHTTPTopAseListByDeviceTypeParamsDeviceType string

const (
	RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeDesktop RadarHTTPTopAseListByDeviceTypeParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeMobile  RadarHTTPTopAseListByDeviceTypeParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeOther   RadarHTTPTopAseListByDeviceTypeParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeDesktop, RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeMobile, RadarHTTPTopAseListByDeviceTypeParamsDeviceTypeOther:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsBotClass string

const (
	RadarHTTPTopAseListByDeviceTypeParamsBotClassLikelyAutomated RadarHTTPTopAseListByDeviceTypeParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByDeviceTypeParamsBotClassLikelyHuman     RadarHTTPTopAseListByDeviceTypeParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsBotClassLikelyAutomated, RadarHTTPTopAseListByDeviceTypeParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily string

const (
	RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyChrome  RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyEdge    RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyFirefox RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilySafari  RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyChrome, RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyEdge, RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilyFirefox, RadarHTTPTopAseListByDeviceTypeParamsBrowserFamilySafari:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByDeviceTypeParamsFormat string

const (
	RadarHTTPTopAseListByDeviceTypeParamsFormatJson RadarHTTPTopAseListByDeviceTypeParamsFormat = "JSON"
	RadarHTTPTopAseListByDeviceTypeParamsFormatCsv  RadarHTTPTopAseListByDeviceTypeParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsFormatJson, RadarHTTPTopAseListByDeviceTypeParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocolHTTP  RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocolHTTPS RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocolHTTP, RadarHTTPTopAseListByDeviceTypeParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion string

const (
	RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv1 RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv2 RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv3 RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv1, RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv2, RadarHTTPTopAseListByDeviceTypeParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsIPVersion string

const (
	RadarHTTPTopAseListByDeviceTypeParamsIPVersionIPv4 RadarHTTPTopAseListByDeviceTypeParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByDeviceTypeParamsIPVersionIPv6 RadarHTTPTopAseListByDeviceTypeParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsIPVersionIPv4, RadarHTTPTopAseListByDeviceTypeParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsO string

const (
	RadarHTTPTopAseListByDeviceTypeParamsOWindows  RadarHTTPTopAseListByDeviceTypeParamsO = "WINDOWS"
	RadarHTTPTopAseListByDeviceTypeParamsOMacosx   RadarHTTPTopAseListByDeviceTypeParamsO = "MACOSX"
	RadarHTTPTopAseListByDeviceTypeParamsOIos      RadarHTTPTopAseListByDeviceTypeParamsO = "IOS"
	RadarHTTPTopAseListByDeviceTypeParamsOAndroid  RadarHTTPTopAseListByDeviceTypeParamsO = "ANDROID"
	RadarHTTPTopAseListByDeviceTypeParamsOChromeos RadarHTTPTopAseListByDeviceTypeParamsO = "CHROMEOS"
	RadarHTTPTopAseListByDeviceTypeParamsOLinux    RadarHTTPTopAseListByDeviceTypeParamsO = "LINUX"
	RadarHTTPTopAseListByDeviceTypeParamsOSmartTv  RadarHTTPTopAseListByDeviceTypeParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsOWindows, RadarHTTPTopAseListByDeviceTypeParamsOMacosx, RadarHTTPTopAseListByDeviceTypeParamsOIos, RadarHTTPTopAseListByDeviceTypeParamsOAndroid, RadarHTTPTopAseListByDeviceTypeParamsOChromeos, RadarHTTPTopAseListByDeviceTypeParamsOLinux, RadarHTTPTopAseListByDeviceTypeParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByDeviceTypeParamsTlsVersion string

const (
	RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByDeviceTypeParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByDeviceTypeParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByDeviceTypeParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByDeviceTypeParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSvQuic RadarHTTPTopAseListByDeviceTypeParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByDeviceTypeParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByDeviceTypeParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByHTTPProtocolParamsFormat] `query:"format"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByHTTPProtocolParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopAseListByHTTPProtocolParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP protocol (HTTP vs. HTTPS).
type RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocolHTTP  RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocolHTTPS RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocolHTTP, RadarHTTPTopAseListByHTTPProtocolParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsBotClass string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsBotClassLikelyAutomated RadarHTTPTopAseListByHTTPProtocolParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByHTTPProtocolParamsBotClassLikelyHuman     RadarHTTPTopAseListByHTTPProtocolParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsBotClassLikelyAutomated, RadarHTTPTopAseListByHTTPProtocolParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyChrome  RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyEdge    RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyFirefox RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilySafari  RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyChrome, RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyEdge, RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilyFirefox, RadarHTTPTopAseListByHTTPProtocolParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsDeviceType string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeDesktop RadarHTTPTopAseListByHTTPProtocolParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeMobile  RadarHTTPTopAseListByHTTPProtocolParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeOther   RadarHTTPTopAseListByHTTPProtocolParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeDesktop, RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeMobile, RadarHTTPTopAseListByHTTPProtocolParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByHTTPProtocolParamsFormat string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsFormatJson RadarHTTPTopAseListByHTTPProtocolParamsFormat = "JSON"
	RadarHTTPTopAseListByHTTPProtocolParamsFormatCsv  RadarHTTPTopAseListByHTTPProtocolParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsFormatJson, RadarHTTPTopAseListByHTTPProtocolParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv1 RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv2 RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv3 RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv1, RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv2, RadarHTTPTopAseListByHTTPProtocolParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsIPVersion string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsIPVersionIPv4 RadarHTTPTopAseListByHTTPProtocolParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByHTTPProtocolParamsIPVersionIPv6 RadarHTTPTopAseListByHTTPProtocolParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsIPVersionIPv4, RadarHTTPTopAseListByHTTPProtocolParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsO string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsOWindows  RadarHTTPTopAseListByHTTPProtocolParamsO = "WINDOWS"
	RadarHTTPTopAseListByHTTPProtocolParamsOMacosx   RadarHTTPTopAseListByHTTPProtocolParamsO = "MACOSX"
	RadarHTTPTopAseListByHTTPProtocolParamsOIos      RadarHTTPTopAseListByHTTPProtocolParamsO = "IOS"
	RadarHTTPTopAseListByHTTPProtocolParamsOAndroid  RadarHTTPTopAseListByHTTPProtocolParamsO = "ANDROID"
	RadarHTTPTopAseListByHTTPProtocolParamsOChromeos RadarHTTPTopAseListByHTTPProtocolParamsO = "CHROMEOS"
	RadarHTTPTopAseListByHTTPProtocolParamsOLinux    RadarHTTPTopAseListByHTTPProtocolParamsO = "LINUX"
	RadarHTTPTopAseListByHTTPProtocolParamsOSmartTv  RadarHTTPTopAseListByHTTPProtocolParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsOWindows, RadarHTTPTopAseListByHTTPProtocolParamsOMacosx, RadarHTTPTopAseListByHTTPProtocolParamsOIos, RadarHTTPTopAseListByHTTPProtocolParamsOAndroid, RadarHTTPTopAseListByHTTPProtocolParamsOChromeos, RadarHTTPTopAseListByHTTPProtocolParamsOLinux, RadarHTTPTopAseListByHTTPProtocolParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion string

const (
	RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSvQuic RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByHTTPProtocolParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByHTTPProtocolParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByHTTPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByHTTPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByHTTPVersionParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopAseListByHTTPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion string

const (
	RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv1 RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv2 RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv3 RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv1, RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv2, RadarHTTPTopAseListByHTTPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsBotClass string

const (
	RadarHTTPTopAseListByHTTPVersionParamsBotClassLikelyAutomated RadarHTTPTopAseListByHTTPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByHTTPVersionParamsBotClassLikelyHuman     RadarHTTPTopAseListByHTTPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsBotClassLikelyAutomated, RadarHTTPTopAseListByHTTPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily string

const (
	RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyChrome  RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyEdge    RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyFirefox RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilySafari  RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyChrome, RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyEdge, RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilyFirefox, RadarHTTPTopAseListByHTTPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsDeviceType string

const (
	RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeDesktop RadarHTTPTopAseListByHTTPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeMobile  RadarHTTPTopAseListByHTTPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeOther   RadarHTTPTopAseListByHTTPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeDesktop, RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeMobile, RadarHTTPTopAseListByHTTPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByHTTPVersionParamsFormat string

const (
	RadarHTTPTopAseListByHTTPVersionParamsFormatJson RadarHTTPTopAseListByHTTPVersionParamsFormat = "JSON"
	RadarHTTPTopAseListByHTTPVersionParamsFormatCsv  RadarHTTPTopAseListByHTTPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsFormatJson, RadarHTTPTopAseListByHTTPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocolHTTP  RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocolHTTPS RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocolHTTP, RadarHTTPTopAseListByHTTPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsIPVersion string

const (
	RadarHTTPTopAseListByHTTPVersionParamsIPVersionIPv4 RadarHTTPTopAseListByHTTPVersionParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByHTTPVersionParamsIPVersionIPv6 RadarHTTPTopAseListByHTTPVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsIPVersionIPv4, RadarHTTPTopAseListByHTTPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsO string

const (
	RadarHTTPTopAseListByHTTPVersionParamsOWindows  RadarHTTPTopAseListByHTTPVersionParamsO = "WINDOWS"
	RadarHTTPTopAseListByHTTPVersionParamsOMacosx   RadarHTTPTopAseListByHTTPVersionParamsO = "MACOSX"
	RadarHTTPTopAseListByHTTPVersionParamsOIos      RadarHTTPTopAseListByHTTPVersionParamsO = "IOS"
	RadarHTTPTopAseListByHTTPVersionParamsOAndroid  RadarHTTPTopAseListByHTTPVersionParamsO = "ANDROID"
	RadarHTTPTopAseListByHTTPVersionParamsOChromeos RadarHTTPTopAseListByHTTPVersionParamsO = "CHROMEOS"
	RadarHTTPTopAseListByHTTPVersionParamsOLinux    RadarHTTPTopAseListByHTTPVersionParamsO = "LINUX"
	RadarHTTPTopAseListByHTTPVersionParamsOSmartTv  RadarHTTPTopAseListByHTTPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsOWindows, RadarHTTPTopAseListByHTTPVersionParamsOMacosx, RadarHTTPTopAseListByHTTPVersionParamsOIos, RadarHTTPTopAseListByHTTPVersionParamsOAndroid, RadarHTTPTopAseListByHTTPVersionParamsOChromeos, RadarHTTPTopAseListByHTTPVersionParamsOLinux, RadarHTTPTopAseListByHTTPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByHTTPVersionParamsTlsVersion string

const (
	RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByHTTPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByHTTPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByHTTPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByHTTPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSvQuic RadarHTTPTopAseListByHTTPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByHTTPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByHTTPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByIPVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByIPVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByIPVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByIPVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByIPVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByIPVersionParamsHTTPVersion] `query:"httpVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByIPVersionParamsO] `query:"os"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByIPVersionParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByIPVersionParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseListByIPVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPTopAseListByIPVersionParamsIPVersion string

const (
	RadarHTTPTopAseListByIPVersionParamsIPVersionIPv4 RadarHTTPTopAseListByIPVersionParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByIPVersionParamsIPVersionIPv6 RadarHTTPTopAseListByIPVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByIPVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsIPVersionIPv4, RadarHTTPTopAseListByIPVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsBotClass string

const (
	RadarHTTPTopAseListByIPVersionParamsBotClassLikelyAutomated RadarHTTPTopAseListByIPVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByIPVersionParamsBotClassLikelyHuman     RadarHTTPTopAseListByIPVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByIPVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsBotClassLikelyAutomated, RadarHTTPTopAseListByIPVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsBrowserFamily string

const (
	RadarHTTPTopAseListByIPVersionParamsBrowserFamilyChrome  RadarHTTPTopAseListByIPVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByIPVersionParamsBrowserFamilyEdge    RadarHTTPTopAseListByIPVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByIPVersionParamsBrowserFamilyFirefox RadarHTTPTopAseListByIPVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByIPVersionParamsBrowserFamilySafari  RadarHTTPTopAseListByIPVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByIPVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsBrowserFamilyChrome, RadarHTTPTopAseListByIPVersionParamsBrowserFamilyEdge, RadarHTTPTopAseListByIPVersionParamsBrowserFamilyFirefox, RadarHTTPTopAseListByIPVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsDeviceType string

const (
	RadarHTTPTopAseListByIPVersionParamsDeviceTypeDesktop RadarHTTPTopAseListByIPVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByIPVersionParamsDeviceTypeMobile  RadarHTTPTopAseListByIPVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByIPVersionParamsDeviceTypeOther   RadarHTTPTopAseListByIPVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByIPVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsDeviceTypeDesktop, RadarHTTPTopAseListByIPVersionParamsDeviceTypeMobile, RadarHTTPTopAseListByIPVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByIPVersionParamsFormat string

const (
	RadarHTTPTopAseListByIPVersionParamsFormatJson RadarHTTPTopAseListByIPVersionParamsFormat = "JSON"
	RadarHTTPTopAseListByIPVersionParamsFormatCsv  RadarHTTPTopAseListByIPVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByIPVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsFormatJson, RadarHTTPTopAseListByIPVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByIPVersionParamsHTTPProtocolHTTP  RadarHTTPTopAseListByIPVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByIPVersionParamsHTTPProtocolHTTPS RadarHTTPTopAseListByIPVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByIPVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsHTTPProtocolHTTP, RadarHTTPTopAseListByIPVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsHTTPVersion string

const (
	RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv1 RadarHTTPTopAseListByIPVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv2 RadarHTTPTopAseListByIPVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv3 RadarHTTPTopAseListByIPVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByIPVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv1, RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv2, RadarHTTPTopAseListByIPVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsO string

const (
	RadarHTTPTopAseListByIPVersionParamsOWindows  RadarHTTPTopAseListByIPVersionParamsO = "WINDOWS"
	RadarHTTPTopAseListByIPVersionParamsOMacosx   RadarHTTPTopAseListByIPVersionParamsO = "MACOSX"
	RadarHTTPTopAseListByIPVersionParamsOIos      RadarHTTPTopAseListByIPVersionParamsO = "IOS"
	RadarHTTPTopAseListByIPVersionParamsOAndroid  RadarHTTPTopAseListByIPVersionParamsO = "ANDROID"
	RadarHTTPTopAseListByIPVersionParamsOChromeos RadarHTTPTopAseListByIPVersionParamsO = "CHROMEOS"
	RadarHTTPTopAseListByIPVersionParamsOLinux    RadarHTTPTopAseListByIPVersionParamsO = "LINUX"
	RadarHTTPTopAseListByIPVersionParamsOSmartTv  RadarHTTPTopAseListByIPVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByIPVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsOWindows, RadarHTTPTopAseListByIPVersionParamsOMacosx, RadarHTTPTopAseListByIPVersionParamsOIos, RadarHTTPTopAseListByIPVersionParamsOAndroid, RadarHTTPTopAseListByIPVersionParamsOChromeos, RadarHTTPTopAseListByIPVersionParamsOLinux, RadarHTTPTopAseListByIPVersionParamsOSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByIPVersionParamsTlsVersion string

const (
	RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByIPVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByIPVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByIPVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByIPVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSvQuic RadarHTTPTopAseListByIPVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByIPVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByIPVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByOsParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByOsParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByOsParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByOsParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByOsParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByOsParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByOsParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by TLS version.
	TlsVersion param.Field[[]RadarHTTPTopAseListByOsParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListByOsParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseListByOsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Operating system.
type RadarHTTPTopAseListByOsParamsOs string

const (
	RadarHTTPTopAseListByOsParamsOsWindows  RadarHTTPTopAseListByOsParamsOs = "WINDOWS"
	RadarHTTPTopAseListByOsParamsOsMacosx   RadarHTTPTopAseListByOsParamsOs = "MACOSX"
	RadarHTTPTopAseListByOsParamsOsIos      RadarHTTPTopAseListByOsParamsOs = "IOS"
	RadarHTTPTopAseListByOsParamsOsAndroid  RadarHTTPTopAseListByOsParamsOs = "ANDROID"
	RadarHTTPTopAseListByOsParamsOsChromeos RadarHTTPTopAseListByOsParamsOs = "CHROMEOS"
	RadarHTTPTopAseListByOsParamsOsLinux    RadarHTTPTopAseListByOsParamsOs = "LINUX"
	RadarHTTPTopAseListByOsParamsOsSmartTv  RadarHTTPTopAseListByOsParamsOs = "SMART_TV"
)

func (r RadarHTTPTopAseListByOsParamsOs) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsOsWindows, RadarHTTPTopAseListByOsParamsOsMacosx, RadarHTTPTopAseListByOsParamsOsIos, RadarHTTPTopAseListByOsParamsOsAndroid, RadarHTTPTopAseListByOsParamsOsChromeos, RadarHTTPTopAseListByOsParamsOsLinux, RadarHTTPTopAseListByOsParamsOsSmartTv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsBotClass string

const (
	RadarHTTPTopAseListByOsParamsBotClassLikelyAutomated RadarHTTPTopAseListByOsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByOsParamsBotClassLikelyHuman     RadarHTTPTopAseListByOsParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByOsParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsBotClassLikelyAutomated, RadarHTTPTopAseListByOsParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsBrowserFamily string

const (
	RadarHTTPTopAseListByOsParamsBrowserFamilyChrome  RadarHTTPTopAseListByOsParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByOsParamsBrowserFamilyEdge    RadarHTTPTopAseListByOsParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByOsParamsBrowserFamilyFirefox RadarHTTPTopAseListByOsParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByOsParamsBrowserFamilySafari  RadarHTTPTopAseListByOsParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByOsParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsBrowserFamilyChrome, RadarHTTPTopAseListByOsParamsBrowserFamilyEdge, RadarHTTPTopAseListByOsParamsBrowserFamilyFirefox, RadarHTTPTopAseListByOsParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsDeviceType string

const (
	RadarHTTPTopAseListByOsParamsDeviceTypeDesktop RadarHTTPTopAseListByOsParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByOsParamsDeviceTypeMobile  RadarHTTPTopAseListByOsParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByOsParamsDeviceTypeOther   RadarHTTPTopAseListByOsParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByOsParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsDeviceTypeDesktop, RadarHTTPTopAseListByOsParamsDeviceTypeMobile, RadarHTTPTopAseListByOsParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByOsParamsFormat string

const (
	RadarHTTPTopAseListByOsParamsFormatJson RadarHTTPTopAseListByOsParamsFormat = "JSON"
	RadarHTTPTopAseListByOsParamsFormatCsv  RadarHTTPTopAseListByOsParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByOsParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsFormatJson, RadarHTTPTopAseListByOsParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByOsParamsHTTPProtocolHTTP  RadarHTTPTopAseListByOsParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByOsParamsHTTPProtocolHTTPS RadarHTTPTopAseListByOsParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByOsParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsHTTPProtocolHTTP, RadarHTTPTopAseListByOsParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsHTTPVersion string

const (
	RadarHTTPTopAseListByOsParamsHTTPVersionHttPv1 RadarHTTPTopAseListByOsParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByOsParamsHTTPVersionHttPv2 RadarHTTPTopAseListByOsParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByOsParamsHTTPVersionHttPv3 RadarHTTPTopAseListByOsParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByOsParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsHTTPVersionHttPv1, RadarHTTPTopAseListByOsParamsHTTPVersionHttPv2, RadarHTTPTopAseListByOsParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsIPVersion string

const (
	RadarHTTPTopAseListByOsParamsIPVersionIPv4 RadarHTTPTopAseListByOsParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByOsParamsIPVersionIPv6 RadarHTTPTopAseListByOsParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByOsParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsIPVersionIPv4, RadarHTTPTopAseListByOsParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByOsParamsTlsVersion string

const (
	RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByOsParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByOsParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByOsParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByOsParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByOsParamsTlsVersionTlSvQuic RadarHTTPTopAseListByOsParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByOsParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByOsParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByOsParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParams struct {
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListByTlsVersionParamsBotClass] `query:"botClass"`
	// Filters results by browser family.
	BrowserFamily param.Field[[]RadarHTTPTopAseListByTlsVersionParamsBrowserFamily] `query:"browserFamily"`
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
	// Filters results by device type.
	DeviceType param.Field[[]RadarHTTPTopAseListByTlsVersionParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[RadarHTTPTopAseListByTlsVersionParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListByTlsVersionParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]RadarHTTPTopAseListByTlsVersionParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	Os param.Field[[]RadarHTTPTopAseListByTlsVersionParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTopAseListByTlsVersionParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopAseListByTlsVersionParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type RadarHTTPTopAseListByTlsVersionParamsTlsVersion string

const (
	RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_0  RadarHTTPTopAseListByTlsVersionParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_1  RadarHTTPTopAseListByTlsVersionParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_2  RadarHTTPTopAseListByTlsVersionParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_3  RadarHTTPTopAseListByTlsVersionParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSvQuic RadarHTTPTopAseListByTlsVersionParamsTlsVersion = "TLSvQUIC"
)

func (r RadarHTTPTopAseListByTlsVersionParamsTlsVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_0, RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_1, RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_2, RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSv1_3, RadarHTTPTopAseListByTlsVersionParamsTlsVersionTlSvQuic:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsBotClass string

const (
	RadarHTTPTopAseListByTlsVersionParamsBotClassLikelyAutomated RadarHTTPTopAseListByTlsVersionParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListByTlsVersionParamsBotClassLikelyHuman     RadarHTTPTopAseListByTlsVersionParamsBotClass = "LIKELY_HUMAN"
)

func (r RadarHTTPTopAseListByTlsVersionParamsBotClass) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsBotClassLikelyAutomated, RadarHTTPTopAseListByTlsVersionParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsBrowserFamily string

const (
	RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyChrome  RadarHTTPTopAseListByTlsVersionParamsBrowserFamily = "CHROME"
	RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyEdge    RadarHTTPTopAseListByTlsVersionParamsBrowserFamily = "EDGE"
	RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyFirefox RadarHTTPTopAseListByTlsVersionParamsBrowserFamily = "FIREFOX"
	RadarHTTPTopAseListByTlsVersionParamsBrowserFamilySafari  RadarHTTPTopAseListByTlsVersionParamsBrowserFamily = "SAFARI"
)

func (r RadarHTTPTopAseListByTlsVersionParamsBrowserFamily) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyChrome, RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyEdge, RadarHTTPTopAseListByTlsVersionParamsBrowserFamilyFirefox, RadarHTTPTopAseListByTlsVersionParamsBrowserFamilySafari:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsDeviceType string

const (
	RadarHTTPTopAseListByTlsVersionParamsDeviceTypeDesktop RadarHTTPTopAseListByTlsVersionParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListByTlsVersionParamsDeviceTypeMobile  RadarHTTPTopAseListByTlsVersionParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListByTlsVersionParamsDeviceTypeOther   RadarHTTPTopAseListByTlsVersionParamsDeviceType = "OTHER"
)

func (r RadarHTTPTopAseListByTlsVersionParamsDeviceType) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsDeviceTypeDesktop, RadarHTTPTopAseListByTlsVersionParamsDeviceTypeMobile, RadarHTTPTopAseListByTlsVersionParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type RadarHTTPTopAseListByTlsVersionParamsFormat string

const (
	RadarHTTPTopAseListByTlsVersionParamsFormatJson RadarHTTPTopAseListByTlsVersionParamsFormat = "JSON"
	RadarHTTPTopAseListByTlsVersionParamsFormatCsv  RadarHTTPTopAseListByTlsVersionParamsFormat = "CSV"
)

func (r RadarHTTPTopAseListByTlsVersionParamsFormat) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsFormatJson, RadarHTTPTopAseListByTlsVersionParamsFormatCsv:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol string

const (
	RadarHTTPTopAseListByTlsVersionParamsHTTPProtocolHTTP  RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListByTlsVersionParamsHTTPProtocolHTTPS RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol = "HTTPS"
)

func (r RadarHTTPTopAseListByTlsVersionParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsHTTPProtocolHTTP, RadarHTTPTopAseListByTlsVersionParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsHTTPVersion string

const (
	RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv1 RadarHTTPTopAseListByTlsVersionParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv2 RadarHTTPTopAseListByTlsVersionParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv3 RadarHTTPTopAseListByTlsVersionParamsHTTPVersion = "HTTPv3"
)

func (r RadarHTTPTopAseListByTlsVersionParamsHTTPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv1, RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv2, RadarHTTPTopAseListByTlsVersionParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsIPVersion string

const (
	RadarHTTPTopAseListByTlsVersionParamsIPVersionIPv4 RadarHTTPTopAseListByTlsVersionParamsIPVersion = "IPv4"
	RadarHTTPTopAseListByTlsVersionParamsIPVersionIPv6 RadarHTTPTopAseListByTlsVersionParamsIPVersion = "IPv6"
)

func (r RadarHTTPTopAseListByTlsVersionParamsIPVersion) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsIPVersionIPv4, RadarHTTPTopAseListByTlsVersionParamsIPVersionIPv6:
		return true
	}
	return false
}

type RadarHTTPTopAseListByTlsVersionParamsO string

const (
	RadarHTTPTopAseListByTlsVersionParamsOWindows  RadarHTTPTopAseListByTlsVersionParamsO = "WINDOWS"
	RadarHTTPTopAseListByTlsVersionParamsOMacosx   RadarHTTPTopAseListByTlsVersionParamsO = "MACOSX"
	RadarHTTPTopAseListByTlsVersionParamsOIos      RadarHTTPTopAseListByTlsVersionParamsO = "IOS"
	RadarHTTPTopAseListByTlsVersionParamsOAndroid  RadarHTTPTopAseListByTlsVersionParamsO = "ANDROID"
	RadarHTTPTopAseListByTlsVersionParamsOChromeos RadarHTTPTopAseListByTlsVersionParamsO = "CHROMEOS"
	RadarHTTPTopAseListByTlsVersionParamsOLinux    RadarHTTPTopAseListByTlsVersionParamsO = "LINUX"
	RadarHTTPTopAseListByTlsVersionParamsOSmartTv  RadarHTTPTopAseListByTlsVersionParamsO = "SMART_TV"
)

func (r RadarHTTPTopAseListByTlsVersionParamsO) IsKnown() bool {
	switch r {
	case RadarHTTPTopAseListByTlsVersionParamsOWindows, RadarHTTPTopAseListByTlsVersionParamsOMacosx, RadarHTTPTopAseListByTlsVersionParamsOIos, RadarHTTPTopAseListByTlsVersionParamsOAndroid, RadarHTTPTopAseListByTlsVersionParamsOChromeos, RadarHTTPTopAseListByTlsVersionParamsOLinux, RadarHTTPTopAseListByTlsVersionParamsOSmartTv:
		return true
	}
	return false
}
