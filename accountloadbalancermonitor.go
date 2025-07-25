// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountLoadBalancerMonitorService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLoadBalancerMonitorService] method instead.
type AccountLoadBalancerMonitorService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerMonitorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerMonitorService(opts ...option.RequestOption) (r *AccountLoadBalancerMonitorService) {
	r = &AccountLoadBalancerMonitorService{}
	r.Options = opts
	return
}

// Create a configured monitor.
func (r *AccountLoadBalancerMonitorService) New(ctx context.Context, accountID string, body AccountLoadBalancerMonitorNewParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List a single configured monitor for an account.
func (r *AccountLoadBalancerMonitorService) Get(ctx context.Context, accountID string, monitorID string, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured monitor.
func (r *AccountLoadBalancerMonitorService) Update(ctx context.Context, accountID string, monitorID string, body AccountLoadBalancerMonitorUpdateParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured monitors for an account.
func (r *AccountLoadBalancerMonitorService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ResponseCollectionMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured monitor.
func (r *AccountLoadBalancerMonitorService) Delete(ctx context.Context, accountID string, monitorID string, opts ...option.RequestOption) (res *IDResponseLoadBalancing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the list of resources that reference the provided monitor.
func (r *AccountLoadBalancerMonitorService) ListReferences(ctx context.Context, accountID string, monitorID string, opts ...option.RequestOption) (res *ReferencesMonitorResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/references", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply changes to an existing monitor, overwriting the supplied properties.
func (r *AccountLoadBalancerMonitorService) Patch(ctx context.Context, accountID string, monitorID string, body AccountLoadBalancerMonitorPatchParams, opts ...option.RequestOption) (res *ResponseSingleMonitor, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *AccountLoadBalancerMonitorService) Preview(ctx context.Context, accountID string, monitorID string, body AccountLoadBalancerMonitorPreviewParams, opts ...option.RequestOption) (res *PreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/preview", accountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EditableMonitorParam struct {
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[map[string][]string] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[EditableMonitorType] `json:"type"`
}

func (r EditableMonitorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type EditableMonitorType string

const (
	EditableMonitorTypeHTTP     EditableMonitorType = "http"
	EditableMonitorTypeHTTPS    EditableMonitorType = "https"
	EditableMonitorTypeTcp      EditableMonitorType = "tcp"
	EditableMonitorTypeUdpIcmp  EditableMonitorType = "udp_icmp"
	EditableMonitorTypeIcmpPing EditableMonitorType = "icmp_ping"
	EditableMonitorTypeSmtp     EditableMonitorType = "smtp"
)

func (r EditableMonitorType) IsKnown() bool {
	switch r {
	case EditableMonitorTypeHTTP, EditableMonitorTypeHTTPS, EditableMonitorTypeTcp, EditableMonitorTypeUdpIcmp, EditableMonitorTypeIcmpPing, EditableMonitorTypeSmtp:
		return true
	}
	return false
}

type IDResponseLoadBalancing struct {
	Errors   []IDResponseLoadBalancingError   `json:"errors,required"`
	Messages []IDResponseLoadBalancingMessage `json:"messages,required"`
	Result   IDResponseLoadBalancingResult    `json:"result,required"`
	// Whether the API call was successful
	Success IDResponseLoadBalancingSuccess `json:"success,required"`
	JSON    idResponseLoadBalancingJSON    `json:"-"`
}

// idResponseLoadBalancingJSON contains the JSON metadata for the struct
// [IDResponseLoadBalancing]
type idResponseLoadBalancingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseLoadBalancing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingJSON) RawJSON() string {
	return r.raw
}

type IDResponseLoadBalancingError struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           IDResponseLoadBalancingErrorsSource `json:"source"`
	JSON             idResponseLoadBalancingErrorJSON    `json:"-"`
}

// idResponseLoadBalancingErrorJSON contains the JSON metadata for the struct
// [IDResponseLoadBalancingError]
type idResponseLoadBalancingErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IDResponseLoadBalancingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingErrorJSON) RawJSON() string {
	return r.raw
}

type IDResponseLoadBalancingErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    idResponseLoadBalancingErrorsSourceJSON `json:"-"`
}

// idResponseLoadBalancingErrorsSourceJSON contains the JSON metadata for the
// struct [IDResponseLoadBalancingErrorsSource]
type idResponseLoadBalancingErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseLoadBalancingErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IDResponseLoadBalancingMessage struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           IDResponseLoadBalancingMessagesSource `json:"source"`
	JSON             idResponseLoadBalancingMessageJSON    `json:"-"`
}

// idResponseLoadBalancingMessageJSON contains the JSON metadata for the struct
// [IDResponseLoadBalancingMessage]
type idResponseLoadBalancingMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IDResponseLoadBalancingMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingMessageJSON) RawJSON() string {
	return r.raw
}

type IDResponseLoadBalancingMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    idResponseLoadBalancingMessagesSourceJSON `json:"-"`
}

// idResponseLoadBalancingMessagesSourceJSON contains the JSON metadata for the
// struct [IDResponseLoadBalancingMessagesSource]
type idResponseLoadBalancingMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseLoadBalancingMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type IDResponseLoadBalancingResult struct {
	ID   string                            `json:"id"`
	JSON idResponseLoadBalancingResultJSON `json:"-"`
}

// idResponseLoadBalancingResultJSON contains the JSON metadata for the struct
// [IDResponseLoadBalancingResult]
type idResponseLoadBalancingResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseLoadBalancingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IDResponseLoadBalancingSuccess bool

const (
	IDResponseLoadBalancingSuccessTrue IDResponseLoadBalancingSuccess = true
)

func (r IDResponseLoadBalancingSuccess) IsKnown() bool {
	switch r {
	case IDResponseLoadBalancingSuccessTrue:
		return true
	}
	return false
}

type LoadBalancingMessages struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           LoadBalancingMessagesSource `json:"source"`
	JSON             loadBalancingMessagesJSON   `json:"-"`
}

// loadBalancingMessagesJSON contains the JSON metadata for the struct
// [LoadBalancingMessages]
type loadBalancingMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancingMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMessagesJSON) RawJSON() string {
	return r.raw
}

type LoadBalancingMessagesSource struct {
	Pointer string                          `json:"pointer"`
	JSON    loadBalancingMessagesSourceJSON `json:"-"`
}

// loadBalancingMessagesSourceJSON contains the JSON metadata for the struct
// [LoadBalancingMessagesSource]
type loadBalancingMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type Monitor struct {
	ID string `json:"id"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64  `json:"consecutive_up"`
	CreatedOn     string `json:"created_on"`
	// Object description.
	Description string `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody string `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes string `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects bool `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header map[string][]string `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval int64 `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method     string `json:"method"`
	ModifiedOn string `json:"modified_on"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path string `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port int64 `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone string `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries int64 `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout int64 `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type MonitorType `json:"type"`
	JSON monitorJSON `json:"-"`
}

// monitorJSON contains the JSON metadata for the struct [Monitor]
type monitorJSON struct {
	ID              apijson.Field
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	CreatedOn       apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	ModifiedOn      apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Monitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorJSON) RawJSON() string {
	return r.raw
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorType string

const (
	MonitorTypeHTTP     MonitorType = "http"
	MonitorTypeHTTPS    MonitorType = "https"
	MonitorTypeTcp      MonitorType = "tcp"
	MonitorTypeUdpIcmp  MonitorType = "udp_icmp"
	MonitorTypeIcmpPing MonitorType = "icmp_ping"
	MonitorTypeSmtp     MonitorType = "smtp"
)

func (r MonitorType) IsKnown() bool {
	switch r {
	case MonitorTypeHTTP, MonitorTypeHTTPS, MonitorTypeTcp, MonitorTypeUdpIcmp, MonitorTypeIcmpPing, MonitorTypeSmtp:
		return true
	}
	return false
}

type PreviewResponse struct {
	Errors   []LoadBalancingMessages `json:"errors,required"`
	Messages []LoadBalancingMessages `json:"messages,required"`
	Result   PreviewResponseResult   `json:"result,required"`
	// Whether the API call was successful
	Success PreviewResponseSuccess `json:"success,required"`
	JSON    previewResponseJSON    `json:"-"`
}

// previewResponseJSON contains the JSON metadata for the struct [PreviewResponse]
type previewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewResponseJSON) RawJSON() string {
	return r.raw
}

type PreviewResponseResult struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string         `json:"pools"`
	PreviewID string                    `json:"preview_id"`
	JSON      previewResponseResultJSON `json:"-"`
}

// previewResponseResultJSON contains the JSON metadata for the struct
// [PreviewResponseResult]
type previewResponseResultJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PreviewResponseSuccess bool

const (
	PreviewResponseSuccessTrue PreviewResponseSuccess = true
)

func (r PreviewResponseSuccess) IsKnown() bool {
	switch r {
	case PreviewResponseSuccessTrue:
		return true
	}
	return false
}

type ReferencesMonitorResponse struct {
	Errors   []LoadBalancingMessages `json:"errors,required"`
	Messages []LoadBalancingMessages `json:"messages,required"`
	// List of resources that reference a given monitor.
	Result []ReferencesMonitorResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success ReferencesMonitorResponseSuccess `json:"success,required"`
	JSON    referencesMonitorResponseJSON    `json:"-"`
}

// referencesMonitorResponseJSON contains the JSON metadata for the struct
// [ReferencesMonitorResponse]
type referencesMonitorResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesMonitorResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesMonitorResponseJSON) RawJSON() string {
	return r.raw
}

type ReferencesMonitorResponseResult struct {
	ReferenceType ReferencesMonitorResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                       `json:"resource_id"`
	ResourceName  string                                       `json:"resource_name"`
	ResourceType  string                                       `json:"resource_type"`
	JSON          referencesMonitorResponseResultJSON          `json:"-"`
}

// referencesMonitorResponseResultJSON contains the JSON metadata for the struct
// [ReferencesMonitorResponseResult]
type referencesMonitorResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ReferencesMonitorResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesMonitorResponseResultJSON) RawJSON() string {
	return r.raw
}

type ReferencesMonitorResponseResultReferenceType string

const (
	ReferencesMonitorResponseResultReferenceTypeStar     ReferencesMonitorResponseResultReferenceType = "*"
	ReferencesMonitorResponseResultReferenceTypeReferral ReferencesMonitorResponseResultReferenceType = "referral"
	ReferencesMonitorResponseResultReferenceTypeReferrer ReferencesMonitorResponseResultReferenceType = "referrer"
)

func (r ReferencesMonitorResponseResultReferenceType) IsKnown() bool {
	switch r {
	case ReferencesMonitorResponseResultReferenceTypeStar, ReferencesMonitorResponseResultReferenceTypeReferral, ReferencesMonitorResponseResultReferenceTypeReferrer:
		return true
	}
	return false
}

// Whether the API call was successful
type ReferencesMonitorResponseSuccess bool

const (
	ReferencesMonitorResponseSuccessTrue ReferencesMonitorResponseSuccess = true
)

func (r ReferencesMonitorResponseSuccess) IsKnown() bool {
	switch r {
	case ReferencesMonitorResponseSuccessTrue:
		return true
	}
	return false
}

type ResponseCollectionMonitor struct {
	Errors   []LoadBalancingMessages `json:"errors,required"`
	Messages []LoadBalancingMessages `json:"messages,required"`
	Result   []Monitor               `json:"result,required"`
	// Whether the API call was successful
	Success    ResponseCollectionMonitorSuccess    `json:"success,required"`
	ResultInfo ResponseCollectionMonitorResultInfo `json:"result_info"`
	JSON       responseCollectionMonitorJSON       `json:"-"`
}

// responseCollectionMonitorJSON contains the JSON metadata for the struct
// [ResponseCollectionMonitor]
type responseCollectionMonitorJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionMonitorJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseCollectionMonitorSuccess bool

const (
	ResponseCollectionMonitorSuccessTrue ResponseCollectionMonitorSuccess = true
)

func (r ResponseCollectionMonitorSuccess) IsKnown() bool {
	switch r {
	case ResponseCollectionMonitorSuccessTrue:
		return true
	}
	return false
}

type ResponseCollectionMonitorResultInfo struct {
	// Total number of results on the current page
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total number of pages available
	TotalPages float64                                 `json:"total_pages"`
	JSON       responseCollectionMonitorResultInfoJSON `json:"-"`
}

// responseCollectionMonitorResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionMonitorResultInfo]
type responseCollectionMonitorResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionMonitorResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionMonitorResultInfoJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleMonitor struct {
	Errors   []LoadBalancingMessages `json:"errors,required"`
	Messages []LoadBalancingMessages `json:"messages,required"`
	Result   Monitor                 `json:"result,required"`
	// Whether the API call was successful
	Success ResponseSingleMonitorSuccess `json:"success,required"`
	JSON    responseSingleMonitorJSON    `json:"-"`
}

// responseSingleMonitorJSON contains the JSON metadata for the struct
// [ResponseSingleMonitor]
type responseSingleMonitorJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleMonitorJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseSingleMonitorSuccess bool

const (
	ResponseSingleMonitorSuccessTrue ResponseSingleMonitorSuccess = true
)

func (r ResponseSingleMonitorSuccess) IsKnown() bool {
	switch r {
	case ResponseSingleMonitorSuccessTrue:
		return true
	}
	return false
}

type AccountLoadBalancerMonitorNewParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r AccountLoadBalancerMonitorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type AccountLoadBalancerMonitorUpdateParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r AccountLoadBalancerMonitorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type AccountLoadBalancerMonitorPatchParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r AccountLoadBalancerMonitorPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}

type AccountLoadBalancerMonitorPreviewParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r AccountLoadBalancerMonitorPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}
