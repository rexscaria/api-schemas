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

// AccountLoadBalancerPoolService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLoadBalancerPoolService] method instead.
type AccountLoadBalancerPoolService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPoolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPoolService(opts ...option.RequestOption) (r *AccountLoadBalancerPoolService) {
	r = &AccountLoadBalancerPoolService{}
	r.Options = opts
	return
}

// Create a new pool.
func (r *AccountLoadBalancerPoolService) New(ctx context.Context, accountID string, body AccountLoadBalancerPoolNewParams, opts ...option.RequestOption) (res *SchemasLoadBalancingSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured pool.
func (r *AccountLoadBalancerPoolService) Get(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *SchemasLoadBalancingSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify a configured pool.
func (r *AccountLoadBalancerPoolService) Update(ctx context.Context, accountID string, poolID string, body AccountLoadBalancerPoolUpdateParams, opts ...option.RequestOption) (res *SchemasLoadBalancingSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured pools.
func (r *AccountLoadBalancerPoolService) List(ctx context.Context, accountID string, query AccountLoadBalancerPoolListParams, opts ...option.RequestOption) (res *SchemasResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a configured pool.
func (r *AccountLoadBalancerPoolService) Delete(ctx context.Context, accountID string, poolID string, body AccountLoadBalancerPoolDeleteParams, opts ...option.RequestOption) (res *SchemasIDResponseLoadBalancing, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Fetch the latest pool health status for a single pool.
func (r *AccountLoadBalancerPoolService) Health(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *HealthDetails, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/health", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the list of resources that reference the provided pool.
func (r *AccountLoadBalancerPoolService) ListReferences(ctx context.Context, accountID string, poolID string, opts ...option.RequestOption) (res *ReferencesPoolResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Apply changes to an existing pool, overwriting the supplied properties.
func (r *AccountLoadBalancerPoolService) Patch(ctx context.Context, accountID string, poolID string, body AccountLoadBalancerPoolPatchParams, opts ...option.RequestOption) (res *SchemasLoadBalancingSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Preview pool health using provided monitor details. The returned preview_id can
// be used in the preview endpoint to retrieve the results.
func (r *AccountLoadBalancerPoolService) Preview(ctx context.Context, accountID string, poolID string, body AccountLoadBalancerPoolPreviewParams, opts ...option.RequestOption) (res *PreviewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/preview", accountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type CheckRegions string

const (
	CheckRegionsWnam       CheckRegions = "WNAM"
	CheckRegionsEnam       CheckRegions = "ENAM"
	CheckRegionsWeu        CheckRegions = "WEU"
	CheckRegionsEeu        CheckRegions = "EEU"
	CheckRegionsNsam       CheckRegions = "NSAM"
	CheckRegionsSsam       CheckRegions = "SSAM"
	CheckRegionsOc         CheckRegions = "OC"
	CheckRegionsMe         CheckRegions = "ME"
	CheckRegionsNaf        CheckRegions = "NAF"
	CheckRegionsSaf        CheckRegions = "SAF"
	CheckRegionsSas        CheckRegions = "SAS"
	CheckRegionsSeas       CheckRegions = "SEAS"
	CheckRegionsNeas       CheckRegions = "NEAS"
	CheckRegionsAllRegions CheckRegions = "ALL_REGIONS"
)

func (r CheckRegions) IsKnown() bool {
	switch r {
	case CheckRegionsWnam, CheckRegionsEnam, CheckRegionsWeu, CheckRegionsEeu, CheckRegionsNsam, CheckRegionsSsam, CheckRegionsOc, CheckRegionsMe, CheckRegionsNaf, CheckRegionsSaf, CheckRegionsSas, CheckRegionsSeas, CheckRegionsNeas, CheckRegionsAllRegions:
		return true
	}
	return false
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptions struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool              `json:"healthy,nullable"`
	JSON    filterOptionsJSON `json:"-"`
}

// filterOptionsJSON contains the JSON metadata for the struct [FilterOptions]
type filterOptionsJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterOptionsJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptionsParam struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r FilterOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HealthDetails struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result HealthDetailsResult `json:"result"`
	JSON   healthDetailsJSON   `json:"-"`
	SingleResponseMonitor
}

// healthDetailsJSON contains the JSON metadata for the struct [HealthDetails]
type healthDetailsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDetailsJSON) RawJSON() string {
	return r.raw
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
type HealthDetailsResult struct {
	// Pool ID
	PoolID string `json:"pool_id"`
	// List of regions and associated health status.
	PopHealth HealthDetailsResultPopHealth `json:"pop_health"`
	JSON      healthDetailsResultJSON      `json:"-"`
}

// healthDetailsResultJSON contains the JSON metadata for the struct
// [HealthDetailsResult]
type healthDetailsResultJSON struct {
	PoolID      apijson.Field
	PopHealth   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetailsResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDetailsResultJSON) RawJSON() string {
	return r.raw
}

// List of regions and associated health status.
type HealthDetailsResultPopHealth struct {
	// Whether health check in region is healthy.
	Healthy bool                                 `json:"healthy"`
	Origins []HealthDetailsResultPopHealthOrigin `json:"origins"`
	JSON    healthDetailsResultPopHealthJSON     `json:"-"`
}

// healthDetailsResultPopHealthJSON contains the JSON metadata for the struct
// [HealthDetailsResultPopHealth]
type healthDetailsResultPopHealthJSON struct {
	Healthy     apijson.Field
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetailsResultPopHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDetailsResultPopHealthJSON) RawJSON() string {
	return r.raw
}

type HealthDetailsResultPopHealthOrigin struct {
	IP   HealthDetailsResultPopHealthOriginsIP  `json:"ip"`
	JSON healthDetailsResultPopHealthOriginJSON `json:"-"`
}

// healthDetailsResultPopHealthOriginJSON contains the JSON metadata for the struct
// [HealthDetailsResultPopHealthOrigin]
type healthDetailsResultPopHealthOriginJSON struct {
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetailsResultPopHealthOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDetailsResultPopHealthOriginJSON) RawJSON() string {
	return r.raw
}

type HealthDetailsResultPopHealthOriginsIP struct {
	// Failure reason.
	FailureReason string `json:"failure_reason"`
	// Origin health status.
	Healthy bool `json:"healthy"`
	// Response code from origin health check.
	ResponseCode float64 `json:"response_code"`
	// Origin RTT (Round Trip Time) response.
	Rtt  string                                    `json:"rtt"`
	JSON healthDetailsResultPopHealthOriginsIPJSON `json:"-"`
}

// healthDetailsResultPopHealthOriginsIPJSON contains the JSON metadata for the
// struct [HealthDetailsResultPopHealthOriginsIP]
type healthDetailsResultPopHealthOriginsIPJSON struct {
	FailureReason apijson.Field
	Healthy       apijson.Field
	ResponseCode  apijson.Field
	Rtt           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *HealthDetailsResultPopHealthOriginsIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthDetailsResultPopHealthOriginsIPJSON) RawJSON() string {
	return r.raw
}

// Configures load shedding policies and percentages for the pool.
type LoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadSheddingJSON          `json:"-"`
}

// loadSheddingJSON contains the JSON metadata for the struct [LoadShedding]
type loadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadSheddingDefaultPolicy string

const (
	LoadSheddingDefaultPolicyRandom LoadSheddingDefaultPolicy = "random"
	LoadSheddingDefaultPolicyHash   LoadSheddingDefaultPolicy = "hash"
)

func (r LoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingDefaultPolicyRandom, LoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadSheddingSessionPolicy string

const (
	LoadSheddingSessionPolicyHash LoadSheddingSessionPolicy = "hash"
)

func (r LoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

// Configures load shedding policies and percentages for the pool.
type LoadSheddingParam struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadSheddingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin FilterOptions `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool FilterOptions          `json:"pool,nullable"`
	JSON notificationFilterJSON `json:"-"`
}

// notificationFilterJSON contains the JSON metadata for the struct
// [NotificationFilter]
type notificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilterParam struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[FilterOptionsParam] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[FilterOptionsParam] `json:"pool"`
}

func (r NotificationFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Origin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header OriginHeader `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64    `json:"weight"`
	JSON   originJSON `json:"-"`
}

// originJSON contains the JSON metadata for the struct [Origin]
type originJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Origin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type OriginHeader struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []string         `json:"Host"`
	JSON originHeaderJSON `json:"-"`
}

// originHeaderJSON contains the JSON metadata for the struct [OriginHeader]
type originHeaderJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originHeaderJSON) RawJSON() string {
	return r.raw
}

type OriginParam struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address param.Field[string] `json:"address"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled param.Field[bool] `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header param.Field[OriginHeaderParam] `json:"header"`
	// A human-identifiable name for the origin.
	Name param.Field[string] `json:"name"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight param.Field[float64] `json:"weight"`
}

func (r OriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type OriginHeaderParam struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]string] `json:"Host"`
}

func (r OriginHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy OriginSteeringPolicy `json:"policy"`
	JSON   originSteeringJSON   `json:"-"`
}

// originSteeringJSON contains the JSON metadata for the struct [OriginSteering]
type originSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originSteeringJSON) RawJSON() string {
	return r.raw
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type OriginSteeringPolicy string

const (
	OriginSteeringPolicyRandom                   OriginSteeringPolicy = "random"
	OriginSteeringPolicyHash                     OriginSteeringPolicy = "hash"
	OriginSteeringPolicyLeastOutstandingRequests OriginSteeringPolicy = "least_outstanding_requests"
	OriginSteeringPolicyLeastConnections         OriginSteeringPolicy = "least_connections"
)

func (r OriginSteeringPolicy) IsKnown() bool {
	switch r {
	case OriginSteeringPolicyRandom, OriginSteeringPolicyHash, OriginSteeringPolicyLeastOutstandingRequests, OriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteeringParam struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy param.Field[OriginSteeringPolicy] `json:"policy"`
}

func (r OriginSteeringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Pool struct {
	ID string `json:"id"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions []CheckRegions `json:"check_regions,nullable"`
	CreatedOn    time.Time      `json:"created_on" format:"date-time"`
	// A human-readable description of the pool.
	Description string `json:"description"`
	// This field shows up only if the pool is disabled. This field is set with the
	// time the pool was disabled at.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled bool `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude float64 `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding LoadShedding `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude float64 `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins int64     `json:"minimum_origins"`
	ModifiedOn     time.Time `json:"modified_on" format:"date-time"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor string `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name string `json:"name"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks []string `json:"networks"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail string `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter NotificationFilter `json:"notification_filter,nullable"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering OriginSteering `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins"`
	JSON    poolJSON `json:"-"`
}

// poolJSON contains the JSON metadata for the struct [Pool]
type poolJSON struct {
	ID                 apijson.Field
	CheckRegions       apijson.Field
	CreatedOn          apijson.Field
	Description        apijson.Field
	DisabledAt         apijson.Field
	Enabled            apijson.Field
	Latitude           apijson.Field
	LoadShedding       apijson.Field
	Longitude          apijson.Field
	MinimumOrigins     apijson.Field
	ModifiedOn         apijson.Field
	Monitor            apijson.Field
	Name               apijson.Field
	Networks           apijson.Field
	NotificationEmail  apijson.Field
	NotificationFilter apijson.Field
	OriginSteering     apijson.Field
	Origins            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Pool) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolJSON) RawJSON() string {
	return r.raw
}

type ReferencesPoolResponse struct {
	// List of resources that reference a given pool.
	Result []ReferencesPoolResponseResult `json:"result"`
	JSON   referencesPoolResponseJSON     `json:"-"`
	CommonResponseLoadBalancers
}

// referencesPoolResponseJSON contains the JSON metadata for the struct
// [ReferencesPoolResponse]
type referencesPoolResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReferencesPoolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesPoolResponseJSON) RawJSON() string {
	return r.raw
}

type ReferencesPoolResponseResult struct {
	ReferenceType ReferencesPoolResponseResultReferenceType `json:"reference_type"`
	ResourceID    string                                    `json:"resource_id"`
	ResourceName  string                                    `json:"resource_name"`
	ResourceType  string                                    `json:"resource_type"`
	JSON          referencesPoolResponseResultJSON          `json:"-"`
}

// referencesPoolResponseResultJSON contains the JSON metadata for the struct
// [ReferencesPoolResponseResult]
type referencesPoolResponseResultJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ReferencesPoolResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r referencesPoolResponseResultJSON) RawJSON() string {
	return r.raw
}

type ReferencesPoolResponseResultReferenceType string

const (
	ReferencesPoolResponseResultReferenceTypeStar     ReferencesPoolResponseResultReferenceType = "*"
	ReferencesPoolResponseResultReferenceTypeReferral ReferencesPoolResponseResultReferenceType = "referral"
	ReferencesPoolResponseResultReferenceTypeReferrer ReferencesPoolResponseResultReferenceType = "referrer"
)

func (r ReferencesPoolResponseResultReferenceType) IsKnown() bool {
	switch r {
	case ReferencesPoolResponseResultReferenceTypeStar, ReferencesPoolResponseResultReferenceTypeReferral, ReferencesPoolResponseResultReferenceTypeReferrer:
		return true
	}
	return false
}

type SchemasIDResponseLoadBalancing struct {
	Result SchemasIDResponseLoadBalancingResult `json:"result"`
	JSON   schemasIDResponseLoadBalancingJSON   `json:"-"`
	SingleResponseMonitor
}

// schemasIDResponseLoadBalancingJSON contains the JSON metadata for the struct
// [SchemasIDResponseLoadBalancing]
type schemasIDResponseLoadBalancingJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLoadBalancing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseLoadBalancingJSON) RawJSON() string {
	return r.raw
}

type SchemasIDResponseLoadBalancingResult struct {
	ID   string                                   `json:"id"`
	JSON schemasIDResponseLoadBalancingResultJSON `json:"-"`
}

// schemasIDResponseLoadBalancingResultJSON contains the JSON metadata for the
// struct [SchemasIDResponseLoadBalancingResult]
type schemasIDResponseLoadBalancingResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseLoadBalancingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseLoadBalancingResultJSON) RawJSON() string {
	return r.raw
}

type SchemasLoadBalancingSingleResponse struct {
	Result Pool                                   `json:"result"`
	JSON   schemasLoadBalancingSingleResponseJSON `json:"-"`
	SingleResponseMonitor
}

// schemasLoadBalancingSingleResponseJSON contains the JSON metadata for the struct
// [SchemasLoadBalancingSingleResponse]
type schemasLoadBalancingSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasLoadBalancingSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasLoadBalancingSingleResponseJSON) RawJSON() string {
	return r.raw
}

type SchemasResponseCollection struct {
	Result []Pool                        `json:"result"`
	JSON   schemasResponseCollectionJSON `json:"-"`
	PaginatedResponseCollection
}

// schemasResponseCollectionJSON contains the JSON metadata for the struct
// [SchemasResponseCollection]
type schemasResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type AccountLoadBalancerPoolNewParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginParam] `json:"origins,required"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[string] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
}

func (r AccountLoadBalancerPoolNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLoadBalancerPoolUpdateParams struct {
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name,required"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginParam] `json:"origins,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]CheckRegions] `json:"check_regions"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[string] `json:"monitor"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
}

func (r AccountLoadBalancerPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLoadBalancerPoolListParams struct {
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[string] `query:"monitor"`
}

// URLQuery serializes [AccountLoadBalancerPoolListParams]'s query parameters as
// `url.Values`.
func (r AccountLoadBalancerPoolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLoadBalancerPoolDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountLoadBalancerPoolDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountLoadBalancerPoolPatchParams struct {
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	CheckRegions param.Field[[]CheckRegions] `json:"check_regions"`
	// A human-readable description of the pool.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) or disable this pool. Disabled pools will not
	// receive traffic and are excluded from health checks. Disabling a pool will cause
	// any load balancers using it to failover to the next pool (if any).
	Enabled param.Field[bool] `json:"enabled"`
	// The latitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, longitude must also be set.
	Latitude param.Field[float64] `json:"latitude"`
	// Configures load shedding policies and percentages for the pool.
	LoadShedding param.Field[LoadSheddingParam] `json:"load_shedding"`
	// The longitude of the data center containing the origins used in this pool in
	// decimal degrees. If this is set, latitude must also be set.
	Longitude param.Field[float64] `json:"longitude"`
	// The minimum number of origins that must be healthy for this pool to serve
	// traffic. If the number of healthy origins falls below this number, the pool will
	// be marked unhealthy and will failover to the next available pool.
	MinimumOrigins param.Field[int64] `json:"minimum_origins"`
	// The ID of the Monitor to use for checking the health of origins within this
	// pool.
	Monitor param.Field[string] `json:"monitor"`
	// A short name (tag) for the pool. Only alphanumeric characters, hyphens, and
	// underscores are allowed.
	Name param.Field[string] `json:"name"`
	// This field is now deprecated. It has been moved to Cloudflare's Centralized
	// Notification service
	// https://developers.cloudflare.com/fundamentals/notifications/. The email address
	// to send health status notifications to. This can be an individual mailbox or a
	// mailing list. Multiple emails can be supplied as a comma delimited list.
	NotificationEmail param.Field[string] `json:"notification_email"`
	// Filter pool and origin health notifications by resource type or health status.
	// Use null to reset.
	NotificationFilter param.Field[NotificationFilterParam] `json:"notification_filter"`
	// Configures origin steering for the pool. Controls how origins are selected for
	// new sessions and traffic without session affinity.
	OriginSteering param.Field[OriginSteeringParam] `json:"origin_steering"`
	// The list of origins within this pool. Traffic directed at this pool is balanced
	// across all currently healthy origins, provided the pool itself is healthy.
	Origins param.Field[[]OriginParam] `json:"origins"`
}

func (r AccountLoadBalancerPoolPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountLoadBalancerPoolPreviewParams struct {
	EditableMonitor EditableMonitorParam `json:"editable_monitor,required"`
}

func (r AccountLoadBalancerPoolPreviewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EditableMonitor)
}
