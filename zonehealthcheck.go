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

// ZoneHealthcheckService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneHealthcheckService] method instead.
type ZoneHealthcheckService struct {
	Options []option.RequestOption
	Preview *ZoneHealthcheckPreviewService
}

// NewZoneHealthcheckService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneHealthcheckService(opts ...option.RequestOption) (r *ZoneHealthcheckService) {
	r = &ZoneHealthcheckService{}
	r.Options = opts
	r.Preview = NewZoneHealthcheckPreviewService(opts...)
	return
}

// Create a new health check.
func (r *ZoneHealthcheckService) New(ctx context.Context, zoneID string, body ZoneHealthcheckNewParams, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured health check.
func (r *ZoneHealthcheckService) Get(ctx context.Context, zoneID string, healthcheckID string, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured health check.
func (r *ZoneHealthcheckService) Update(ctx context.Context, zoneID string, healthcheckID string, body ZoneHealthcheckUpdateParams, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured health checks.
func (r *ZoneHealthcheckService) List(ctx context.Context, zoneID string, query ZoneHealthcheckListParams, opts ...option.RequestOption) (res *ZoneHealthcheckListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a health check.
func (r *ZoneHealthcheckService) Delete(ctx context.Context, zoneID string, healthcheckID string, opts ...option.RequestOption) (res *HealthcheckIDResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patch a configured health check.
func (r *ZoneHealthcheckService) Patch(ctx context.Context, zoneID string, healthcheckID string, body ZoneHealthcheckPatchParams, opts ...option.RequestOption) (res *HealthcheckSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if healthcheckID == "" {
		err = errors.New("missing required healthcheck_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/healthchecks/%s", zoneID, healthcheckID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type Healthcheck struct {
	// Identifier
	ID string `json:"id"`
	// The hostname or IP address of the origin server to run health checks on.
	Address string `json:"address"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions []HealthcheckCheckRegion `json:"check_regions,nullable"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails int64 `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses int64     `json:"consecutive_successes"`
	CreatedOn            time.Time `json:"created_on" format:"date-time"`
	// A human-readable description of the health check.
	Description string `json:"description"`
	// The current failure reason if status is unhealthy.
	FailureReason string `json:"failure_reason"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig HealthcheckHTTPConfig `json:"http_config,nullable"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval   int64     `json:"interval"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name string `json:"name"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The current status of the origin server according to the health check.
	Status HealthcheckStatus `json:"status"`
	// If suspended, no health checks are sent to the origin.
	Suspended bool `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig HealthcheckTcpConfig `json:"tcp_config,nullable"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type string          `json:"type"`
	JSON healthcheckJSON `json:"-"`
}

// healthcheckJSON contains the JSON metadata for the struct [Healthcheck]
type healthcheckJSON struct {
	ID                   apijson.Field
	Address              apijson.Field
	CheckRegions         apijson.Field
	ConsecutiveFails     apijson.Field
	ConsecutiveSuccesses apijson.Field
	CreatedOn            apijson.Field
	Description          apijson.Field
	FailureReason        apijson.Field
	HTTPConfig           apijson.Field
	Interval             apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	Retries              apijson.Field
	Status               apijson.Field
	Suspended            apijson.Field
	TcpConfig            apijson.Field
	Timeout              apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *Healthcheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckJSON) RawJSON() string {
	return r.raw
}

// The current status of the origin server according to the health check.
type HealthcheckStatus string

const (
	HealthcheckStatusUnknown   HealthcheckStatus = "unknown"
	HealthcheckStatusHealthy   HealthcheckStatus = "healthy"
	HealthcheckStatusUnhealthy HealthcheckStatus = "unhealthy"
	HealthcheckStatusSuspended HealthcheckStatus = "suspended"
)

func (r HealthcheckStatus) IsKnown() bool {
	switch r {
	case HealthcheckStatusUnknown, HealthcheckStatusHealthy, HealthcheckStatusUnhealthy, HealthcheckStatusSuspended:
		return true
	}
	return false
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, IN: India,
// SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all regions (BUSINESS
// and ENTERPRISE customers only).
type HealthcheckCheckRegion string

const (
	HealthcheckCheckRegionWnam       HealthcheckCheckRegion = "WNAM"
	HealthcheckCheckRegionEnam       HealthcheckCheckRegion = "ENAM"
	HealthcheckCheckRegionWeu        HealthcheckCheckRegion = "WEU"
	HealthcheckCheckRegionEeu        HealthcheckCheckRegion = "EEU"
	HealthcheckCheckRegionNsam       HealthcheckCheckRegion = "NSAM"
	HealthcheckCheckRegionSsam       HealthcheckCheckRegion = "SSAM"
	HealthcheckCheckRegionOc         HealthcheckCheckRegion = "OC"
	HealthcheckCheckRegionMe         HealthcheckCheckRegion = "ME"
	HealthcheckCheckRegionNaf        HealthcheckCheckRegion = "NAF"
	HealthcheckCheckRegionSaf        HealthcheckCheckRegion = "SAF"
	HealthcheckCheckRegionIn         HealthcheckCheckRegion = "IN"
	HealthcheckCheckRegionSeas       HealthcheckCheckRegion = "SEAS"
	HealthcheckCheckRegionNeas       HealthcheckCheckRegion = "NEAS"
	HealthcheckCheckRegionAllRegions HealthcheckCheckRegion = "ALL_REGIONS"
)

func (r HealthcheckCheckRegion) IsKnown() bool {
	switch r {
	case HealthcheckCheckRegionWnam, HealthcheckCheckRegionEnam, HealthcheckCheckRegionWeu, HealthcheckCheckRegionEeu, HealthcheckCheckRegionNsam, HealthcheckCheckRegionSsam, HealthcheckCheckRegionOc, HealthcheckCheckRegionMe, HealthcheckCheckRegionNaf, HealthcheckCheckRegionSaf, HealthcheckCheckRegionIn, HealthcheckCheckRegionSeas, HealthcheckCheckRegionNeas, HealthcheckCheckRegionAllRegions:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckHTTPConfig struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure bool `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes []string `json:"expected_codes,nullable"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header map[string][]string `json:"header,nullable"`
	// The HTTP method to use for the health check.
	Method HealthcheckHTTPConfigMethod `json:"method"`
	// The endpoint path to health check against.
	Path string `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port int64                     `json:"port"`
	JSON healthcheckHTTPConfigJSON `json:"-"`
}

// healthcheckHTTPConfigJSON contains the JSON metadata for the struct
// [HealthcheckHTTPConfig]
type healthcheckHTTPConfigJSON struct {
	AllowInsecure   apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HealthcheckHTTPConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckHTTPConfigJSON) RawJSON() string {
	return r.raw
}

// The HTTP method to use for the health check.
type HealthcheckHTTPConfigMethod string

const (
	HealthcheckHTTPConfigMethodGet  HealthcheckHTTPConfigMethod = "GET"
	HealthcheckHTTPConfigMethodHead HealthcheckHTTPConfigMethod = "HEAD"
)

func (r HealthcheckHTTPConfigMethod) IsKnown() bool {
	switch r {
	case HealthcheckHTTPConfigMethodGet, HealthcheckHTTPConfigMethodHead:
		return true
	}
	return false
}

// Parameters specific to an HTTP or HTTPS health check.
type HealthcheckHTTPConfigParam struct {
	// Do not validate the certificate when the health check uses HTTPS.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response codes (e.g. "200") or code ranges (e.g. "2xx" for all
	// codes starting with 2) of the health check.
	ExpectedCodes param.Field[[]string] `json:"expected_codes"`
	// Follow redirects if the origin returns a 3xx status code.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden.
	Header param.Field[map[string][]string] `json:"header"`
	// The HTTP method to use for the health check.
	Method param.Field[HealthcheckHTTPConfigMethod] `json:"method"`
	// The endpoint path to health check against.
	Path param.Field[string] `json:"path"`
	// Port number to connect to for the health check. Defaults to 80 if type is HTTP
	// or 443 if type is HTTPS.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckHTTPConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HealthcheckIDResponse struct {
	Errors   []HealthcheckIDResponseError   `json:"errors,required"`
	Messages []HealthcheckIDResponseMessage `json:"messages,required"`
	Result   HealthcheckIDResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckIDResponseSuccess `json:"success,required"`
	JSON    healthcheckIDResponseJSON    `json:"-"`
}

// healthcheckIDResponseJSON contains the JSON metadata for the struct
// [HealthcheckIDResponse]
type healthcheckIDResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckIDResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseJSON) RawJSON() string {
	return r.raw
}

type HealthcheckIDResponseError struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           HealthcheckIDResponseErrorsSource `json:"source"`
	JSON             healthcheckIDResponseErrorJSON    `json:"-"`
}

// healthcheckIDResponseErrorJSON contains the JSON metadata for the struct
// [HealthcheckIDResponseError]
type healthcheckIDResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HealthcheckIDResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseErrorJSON) RawJSON() string {
	return r.raw
}

type HealthcheckIDResponseErrorsSource struct {
	Pointer string                                `json:"pointer"`
	JSON    healthcheckIDResponseErrorsSourceJSON `json:"-"`
}

// healthcheckIDResponseErrorsSourceJSON contains the JSON metadata for the struct
// [HealthcheckIDResponseErrorsSource]
type healthcheckIDResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckIDResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HealthcheckIDResponseMessage struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           HealthcheckIDResponseMessagesSource `json:"source"`
	JSON             healthcheckIDResponseMessageJSON    `json:"-"`
}

// healthcheckIDResponseMessageJSON contains the JSON metadata for the struct
// [HealthcheckIDResponseMessage]
type healthcheckIDResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HealthcheckIDResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseMessageJSON) RawJSON() string {
	return r.raw
}

type HealthcheckIDResponseMessagesSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    healthcheckIDResponseMessagesSourceJSON `json:"-"`
}

// healthcheckIDResponseMessagesSourceJSON contains the JSON metadata for the
// struct [HealthcheckIDResponseMessagesSource]
type healthcheckIDResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckIDResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type HealthcheckIDResponseResult struct {
	// Identifier
	ID   string                          `json:"id"`
	JSON healthcheckIDResponseResultJSON `json:"-"`
}

// healthcheckIDResponseResultJSON contains the JSON metadata for the struct
// [HealthcheckIDResponseResult]
type healthcheckIDResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckIDResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckIDResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckIDResponseSuccess bool

const (
	HealthcheckIDResponseSuccessTrue HealthcheckIDResponseSuccess = true
)

func (r HealthcheckIDResponseSuccess) IsKnown() bool {
	switch r {
	case HealthcheckIDResponseSuccessTrue:
		return true
	}
	return false
}

type HealthcheckMessage struct {
	Code             int64                    `json:"code,required"`
	Message          string                   `json:"message,required"`
	DocumentationURL string                   `json:"documentation_url"`
	Source           HealthcheckMessageSource `json:"source"`
	JSON             healthcheckMessageJSON   `json:"-"`
}

// healthcheckMessageJSON contains the JSON metadata for the struct
// [HealthcheckMessage]
type healthcheckMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HealthcheckMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckMessageJSON) RawJSON() string {
	return r.raw
}

type HealthcheckMessageSource struct {
	Pointer string                       `json:"pointer"`
	JSON    healthcheckMessageSourceJSON `json:"-"`
}

// healthcheckMessageSourceJSON contains the JSON metadata for the struct
// [HealthcheckMessageSource]
type healthcheckMessageSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckMessageSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckMessageSourceJSON) RawJSON() string {
	return r.raw
}

type HealthcheckQueryParam struct {
	// The hostname or IP address of the origin server to run health checks on.
	Address param.Field[string] `json:"address,required"`
	// A short name to identify the health check. Only alphanumeric characters, hyphens
	// and underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// A list of regions from which to run health checks. Null means Cloudflare will
	// pick a default region.
	CheckRegions param.Field[[]HealthcheckCheckRegion] `json:"check_regions"`
	// The number of consecutive fails required from a health check before changing the
	// health to unhealthy.
	ConsecutiveFails param.Field[int64] `json:"consecutive_fails"`
	// The number of consecutive successes required from a health check before changing
	// the health to healthy.
	ConsecutiveSuccesses param.Field[int64] `json:"consecutive_successes"`
	// A human-readable description of the health check.
	Description param.Field[string] `json:"description"`
	// Parameters specific to an HTTP or HTTPS health check.
	HTTPConfig param.Field[HealthcheckHTTPConfigParam] `json:"http_config"`
	// The interval between each health check. Shorter intervals may give quicker
	// notifications if the origin status changes, but will increase load on the origin
	// as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// If suspended, no health checks are sent to the origin.
	Suspended param.Field[bool] `json:"suspended"`
	// Parameters specific to TCP health check.
	TcpConfig param.Field[HealthcheckTcpConfigParam] `json:"tcp_config"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP', 'HTTPS' and 'TCP'.
	Type param.Field[string] `json:"type"`
}

func (r HealthcheckQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HealthcheckSingleResponse struct {
	Errors   []HealthcheckMessage `json:"errors,required"`
	Messages []HealthcheckMessage `json:"messages,required"`
	Result   Healthcheck          `json:"result,required"`
	// Whether the API call was successful
	Success HealthcheckSingleResponseSuccess `json:"success,required"`
	JSON    healthcheckSingleResponseJSON    `json:"-"`
}

// healthcheckSingleResponseJSON contains the JSON metadata for the struct
// [HealthcheckSingleResponse]
type healthcheckSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HealthcheckSingleResponseSuccess bool

const (
	HealthcheckSingleResponseSuccessTrue HealthcheckSingleResponseSuccess = true
)

func (r HealthcheckSingleResponseSuccess) IsKnown() bool {
	switch r {
	case HealthcheckSingleResponseSuccessTrue:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type HealthcheckTcpConfig struct {
	// The TCP connection method to use for the health check.
	Method HealthcheckTcpConfigMethod `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port int64                    `json:"port"`
	JSON healthcheckTcpConfigJSON `json:"-"`
}

// healthcheckTcpConfigJSON contains the JSON metadata for the struct
// [HealthcheckTcpConfig]
type healthcheckTcpConfigJSON struct {
	Method      apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthcheckTcpConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthcheckTcpConfigJSON) RawJSON() string {
	return r.raw
}

// The TCP connection method to use for the health check.
type HealthcheckTcpConfigMethod string

const (
	HealthcheckTcpConfigMethodConnectionEstablished HealthcheckTcpConfigMethod = "connection_established"
)

func (r HealthcheckTcpConfigMethod) IsKnown() bool {
	switch r {
	case HealthcheckTcpConfigMethodConnectionEstablished:
		return true
	}
	return false
}

// Parameters specific to TCP health check.
type HealthcheckTcpConfigParam struct {
	// The TCP connection method to use for the health check.
	Method param.Field[HealthcheckTcpConfigMethod] `json:"method"`
	// Port number to connect to for the health check. Defaults to 80.
	Port param.Field[int64] `json:"port"`
}

func (r HealthcheckTcpConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneHealthcheckListResponse struct {
	Errors   []HealthcheckMessage `json:"errors,required"`
	Messages []HealthcheckMessage `json:"messages,required"`
	Result   []Healthcheck        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZoneHealthcheckListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneHealthcheckListResponseResultInfo `json:"result_info"`
	JSON       zoneHealthcheckListResponseJSON       `json:"-"`
}

// zoneHealthcheckListResponseJSON contains the JSON metadata for the struct
// [ZoneHealthcheckListResponse]
type zoneHealthcheckListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHealthcheckListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneHealthcheckListResponseSuccess bool

const (
	ZoneHealthcheckListResponseSuccessTrue ZoneHealthcheckListResponseSuccess = true
)

func (r ZoneHealthcheckListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneHealthcheckListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneHealthcheckListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       zoneHealthcheckListResponseResultInfoJSON `json:"-"`
}

// zoneHealthcheckListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneHealthcheckListResponseResultInfo]
type zoneHealthcheckListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneHealthcheckListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHealthcheckListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneHealthcheckNewParams struct {
	HealthcheckQuery HealthcheckQueryParam `json:"healthcheck_query,required"`
}

func (r ZoneHealthcheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HealthcheckQuery)
}

type ZoneHealthcheckUpdateParams struct {
	HealthcheckQuery HealthcheckQueryParam `json:"healthcheck_query,required"`
}

func (r ZoneHealthcheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HealthcheckQuery)
}

type ZoneHealthcheckListParams struct {
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page. Must be a multiple of 5.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneHealthcheckListParams]'s query parameters as
// `url.Values`.
func (r ZoneHealthcheckListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneHealthcheckPatchParams struct {
	HealthcheckQuery HealthcheckQueryParam `json:"healthcheck_query,required"`
}

func (r ZoneHealthcheckPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.HealthcheckQuery)
}
