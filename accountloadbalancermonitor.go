// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountLoadBalancerMonitorService) Delete(ctx context.Context, accountID string, monitorID string, body AccountLoadBalancerMonitorDeleteParams, opts ...option.RequestOption) (res *IDResponseLoadBalancing, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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

type CommonResponseLoadBalancers struct {
	Errors   []LoadBalancingMessages `json:"errors,required"`
	Messages []LoadBalancingMessages `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseLoadBalancersSuccess `json:"success,required"`
	JSON    commonResponseLoadBalancersJSON    `json:"-"`
}

// commonResponseLoadBalancersJSON contains the JSON metadata for the struct
// [CommonResponseLoadBalancers]
type commonResponseLoadBalancersJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseLoadBalancers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseLoadBalancersJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseLoadBalancersSuccess bool

const (
	CommonResponseLoadBalancersSuccessTrue CommonResponseLoadBalancersSuccess = true
)

func (r CommonResponseLoadBalancersSuccess) IsKnown() bool {
	switch r {
	case CommonResponseLoadBalancersSuccessTrue:
		return true
	}
	return false
}

type EditableMonitor struct {
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure bool `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown int64 `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp int64 `json:"consecutive_up"`
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
	Method string `json:"method"`
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
	Type EditableMonitorType `json:"type"`
	JSON editableMonitorJSON `json:"-"`
}

// editableMonitorJSON contains the JSON metadata for the struct [EditableMonitor]
type editableMonitorJSON struct {
	AllowInsecure   apijson.Field
	ConsecutiveDown apijson.Field
	ConsecutiveUp   apijson.Field
	Description     apijson.Field
	ExpectedBody    apijson.Field
	ExpectedCodes   apijson.Field
	FollowRedirects apijson.Field
	Header          apijson.Field
	Interval        apijson.Field
	Method          apijson.Field
	Path            apijson.Field
	Port            apijson.Field
	ProbeZone       apijson.Field
	Retries         apijson.Field
	Timeout         apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EditableMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r editableMonitorJSON) RawJSON() string {
	return r.raw
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

type IDResponseLoadBalancing struct {
	Result IDResponseLoadBalancingResult `json:"result"`
	JSON   idResponseLoadBalancingJSON   `json:"-"`
	SingleResponseMonitor
}

// idResponseLoadBalancingJSON contains the JSON metadata for the struct
// [IDResponseLoadBalancing]
type idResponseLoadBalancingJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseLoadBalancing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseLoadBalancingJSON) RawJSON() string {
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

type LoadBalancingMessages struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    loadBalancingMessagesJSON `json:"-"`
}

// loadBalancingMessagesJSON contains the JSON metadata for the struct
// [LoadBalancingMessages]
type loadBalancingMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingMessagesJSON) RawJSON() string {
	return r.raw
}

type Monitor struct {
	ID         string      `json:"id"`
	CreatedOn  time.Time   `json:"created_on" format:"date-time"`
	ModifiedOn time.Time   `json:"modified_on" format:"date-time"`
	JSON       monitorJSON `json:"-"`
	EditableMonitor
}

// monitorJSON contains the JSON metadata for the struct [Monitor]
type monitorJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Monitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorJSON) RawJSON() string {
	return r.raw
}

type PreviewResponse struct {
	Result PreviewResponseResult `json:"result"`
	JSON   previewResponseJSON   `json:"-"`
	SingleResponseMonitor
}

// previewResponseJSON contains the JSON metadata for the struct [PreviewResponse]
type previewResponseJSON struct {
	Result      apijson.Field
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

type ReferencesMonitorResponse struct {
	// List of resources that reference a given monitor.
	Result []ReferencesMonitorResponseResult `json:"result"`
	JSON   referencesMonitorResponseJSON     `json:"-"`
	CommonResponseLoadBalancers
}

// referencesMonitorResponseJSON contains the JSON metadata for the struct
// [ReferencesMonitorResponse]
type referencesMonitorResponseJSON struct {
	Result      apijson.Field
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

type ResponseCollectionMonitor struct {
	Result []Monitor                     `json:"result"`
	JSON   responseCollectionMonitorJSON `json:"-"`
	PaginatedResponseCollection
}

// responseCollectionMonitorJSON contains the JSON metadata for the struct
// [ResponseCollectionMonitor]
type responseCollectionMonitorJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionMonitorJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleMonitor struct {
	Result Monitor                   `json:"result"`
	JSON   responseSingleMonitorJSON `json:"-"`
	SingleResponseMonitor
}

// responseSingleMonitorJSON contains the JSON metadata for the struct
// [ResponseSingleMonitor]
type responseSingleMonitorJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleMonitorJSON) RawJSON() string {
	return r.raw
}

type SingleResponseMonitor struct {
	Result interface{}               `json:"result"`
	JSON   singleResponseMonitorJSON `json:"-"`
	CommonResponseLoadBalancers
}

// singleResponseMonitorJSON contains the JSON metadata for the struct
// [SingleResponseMonitor]
type singleResponseMonitorJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseMonitorJSON) RawJSON() string {
	return r.raw
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

type AccountLoadBalancerMonitorDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountLoadBalancerMonitorDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
