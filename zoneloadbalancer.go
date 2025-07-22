// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneLoadBalancerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneLoadBalancerService] method instead.
type ZoneLoadBalancerService struct {
	Options []option.RequestOption
}

// NewZoneLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLoadBalancerService(opts ...option.RequestOption) (r *ZoneLoadBalancerService) {
	r = &ZoneLoadBalancerService{}
	r.Options = opts
	return
}

// Create a new load balancer.
func (r *ZoneLoadBalancerService) New(ctx context.Context, zoneID string, body ZoneLoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancerSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single configured load balancer.
func (r *ZoneLoadBalancerService) Get(ctx context.Context, zoneID string, loadBalancerID string, opts ...option.RequestOption) (res *LoadBalancerSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured load balancer.
func (r *ZoneLoadBalancerService) Update(ctx context.Context, zoneID string, loadBalancerID string, body ZoneLoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancerSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List configured load balancers.
func (r *ZoneLoadBalancerService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneLoadBalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a configured load balancer.
func (r *ZoneLoadBalancerService) Delete(ctx context.Context, zoneID string, loadBalancerID string, body ZoneLoadBalancerDeleteParams, opts ...option.RequestOption) (res *ZoneLoadBalancerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Apply changes to an existing load balancer, overwriting the supplied properties.
func (r *ZoneLoadBalancerService) Patch(ctx context.Context, zoneID string, loadBalancerID string, body ZoneLoadBalancerPatchParams, opts ...option.RequestOption) (res *LoadBalancerSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", zoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type AdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                `json:"failover_across_pools"`
	JSON                adaptiveRoutingJSON `json:"-"`
}

// adaptiveRoutingJSON contains the JSON metadata for the struct [AdaptiveRouting]
type adaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adaptiveRoutingJSON) RawJSON() string {
	return r.raw
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type AdaptiveRoutingParam struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r AdaptiveRoutingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancer struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting AdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools map[string][]string `json:"country_pools"`
	CreatedOn    time.Time           `json:"created_on" format:"date-time"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []string `json:"default_pools"`
	// Object description.
	Description string `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled bool `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool string `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LocationStrategy `json:"location_strategy"`
	ModifiedOn       time.Time        `json:"modified_on" format:"date-time"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name string `json:"name"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks []string `json:"networks"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools map[string][]string `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied bool `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering RandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools map[string][]string `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []LoadBalancingRule `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity SessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes SessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy SteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl      float64          `json:"ttl"`
	ZoneName string           `json:"zone_name"`
	JSON     loadBalancerJSON `json:"-"`
}

// loadBalancerJSON contains the JSON metadata for the struct [LoadBalancer]
type loadBalancerJSON struct {
	ID                        apijson.Field
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	CreatedOn                 apijson.Field
	DefaultPools              apijson.Field
	Description               apijson.Field
	Enabled                   apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	ModifiedOn                apijson.Field
	Name                      apijson.Field
	Networks                  apijson.Field
	PopPools                  apijson.Field
	Proxied                   apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	Rules                     apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	ZoneName                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerSingleResponse struct {
	Result LoadBalancer                   `json:"result"`
	JSON   loadBalancerSingleResponseJSON `json:"-"`
	SingleResponseMonitor
}

// loadBalancerSingleResponseJSON contains the JSON metadata for the struct
// [LoadBalancerSingleResponse]
type loadBalancerSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerSingleResponseJSON) RawJSON() string {
	return r.raw
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancingRule struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition string `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled bool `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse LoadBalancingRuleFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides LoadBalancingRuleOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool                  `json:"terminates"`
	JSON       loadBalancingRuleJSON `json:"-"`
}

// loadBalancingRuleJSON contains the JSON metadata for the struct
// [LoadBalancingRule]
type loadBalancingRuleJSON struct {
	Condition     apijson.Field
	Disabled      apijson.Field
	FixedResponse apijson.Field
	Name          apijson.Field
	Overrides     apijson.Field
	Priority      apijson.Field
	Terminates    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancingRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingRuleJSON) RawJSON() string {
	return r.raw
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancingRuleFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                              `json:"status_code"`
	JSON       loadBalancingRuleFixedResponseJSON `json:"-"`
}

// loadBalancingRuleFixedResponseJSON contains the JSON metadata for the struct
// [LoadBalancingRuleFixedResponse]
type loadBalancingRuleFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingRuleFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingRuleFixedResponseJSON) RawJSON() string {
	return r.raw
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancingRuleOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting AdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools map[string][]string `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []string `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool string `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LocationStrategy `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools map[string][]string `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering RandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools map[string][]string `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity SessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes SessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy SteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl  float64                        `json:"ttl"`
	JSON loadBalancingRuleOverridesJSON `json:"-"`
}

// loadBalancingRuleOverridesJSON contains the JSON metadata for the struct
// [LoadBalancingRuleOverrides]
type loadBalancingRuleOverridesJSON struct {
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	DefaultPools              apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	PopPools                  apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTtl        apijson.Field
	SteeringPolicy            apijson.Field
	Ttl                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancingRuleOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingRuleOverridesJSON) RawJSON() string {
	return r.raw
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type LoadBalancingRuleParam struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition param.Field[string] `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled param.Field[bool] `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse param.Field[LoadBalancingRuleFixedResponseParam] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[LoadBalancingRuleOverridesParam] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r LoadBalancingRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type LoadBalancingRuleFixedResponseParam struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r LoadBalancingRuleFixedResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type LoadBalancingRuleOverridesParam struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[map[string][]string] `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
}

func (r LoadBalancingRuleOverridesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs LocationStrategyPreferEcs `json:"prefer_ecs"`
	JSON      locationStrategyJSON      `json:"-"`
}

// locationStrategyJSON contains the JSON metadata for the struct
// [LocationStrategy]
type locationStrategyJSON struct {
	Mode        apijson.Field
	PreferEcs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationStrategyJSON) RawJSON() string {
	return r.raw
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LocationStrategyMode string

const (
	LocationStrategyModePop        LocationStrategyMode = "pop"
	LocationStrategyModeResolverIP LocationStrategyMode = "resolver_ip"
)

func (r LocationStrategyMode) IsKnown() bool {
	switch r {
	case LocationStrategyModePop, LocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LocationStrategyPreferEcs string

const (
	LocationStrategyPreferEcsAlways    LocationStrategyPreferEcs = "always"
	LocationStrategyPreferEcsNever     LocationStrategyPreferEcs = "never"
	LocationStrategyPreferEcsProximity LocationStrategyPreferEcs = "proximity"
	LocationStrategyPreferEcsGeo       LocationStrategyPreferEcs = "geo"
)

func (r LocationStrategyPreferEcs) IsKnown() bool {
	switch r {
	case LocationStrategyPreferEcsAlways, LocationStrategyPreferEcsNever, LocationStrategyPreferEcsProximity, LocationStrategyPreferEcsGeo:
		return true
	}
	return false
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LocationStrategyParam struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferEcs param.Field[LocationStrategyPreferEcs] `json:"prefer_ecs"`
}

func (r LocationStrategyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type RandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights map[string]float64 `json:"pool_weights"`
	JSON        randomSteeringJSON `json:"-"`
}

// randomSteeringJSON contains the JSON metadata for the struct [RandomSteering]
type randomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r randomSteeringJSON) RawJSON() string {
	return r.raw
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type RandomSteeringParam struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[map[string]float64] `json:"pool_weights"`
}

func (r RandomSteeringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"`. The supported types are:
//
//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
//     generated, encoding information of which origin the request will be forwarded
//     to. Subsequent requests, by the same client to the same load balancer, will be
//     sent to the origin server the cookie encodes, for the duration of the cookie
//     and as long as the origin server remains healthy. If the cookie has expired or
//     the origin server is unhealthy, then a new origin server is calculated and
//     used.
//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
//     selection is stable and based on the client's ip address.
//   - `"header"`: On the first request to a proxied load balancer, a session key
//     based on the configured HTTP headers (see
//     `session_affinity_attributes.headers`) is generated, encoding the request
//     headers used for storing in the load balancer session state which origin the
//     request will be forwarded to. Subsequent requests to the load balancer with
//     the same headers will be sent to the same origin server, for the duration of
//     the session and as long as the origin server remains healthy. If the session
//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
//     server is unhealthy, then a new origin server is calculated and used. See
//     `headers` in `session_affinity_attributes` for additional required
//     configuration.
type SessionAffinity string

const (
	SessionAffinityNone     SessionAffinity = "none"
	SessionAffinityCookie   SessionAffinity = "cookie"
	SessionAffinityIPCookie SessionAffinity = "ip_cookie"
	SessionAffinityHeader   SessionAffinity = "header"
)

func (r SessionAffinity) IsKnown() bool {
	switch r {
	case SessionAffinityNone, SessionAffinityCookie, SessionAffinityIPCookie, SessionAffinityHeader:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type SessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers []string `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders bool `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite SessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure SessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover SessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 sessionAffinityAttributesJSON                 `json:"-"`
}

// sessionAffinityAttributesJSON contains the JSON metadata for the struct
// [SessionAffinityAttributes]
type sessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionAffinityAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type SessionAffinityAttributesSamesite string

const (
	SessionAffinityAttributesSamesiteAuto   SessionAffinityAttributesSamesite = "Auto"
	SessionAffinityAttributesSamesiteLax    SessionAffinityAttributesSamesite = "Lax"
	SessionAffinityAttributesSamesiteNone   SessionAffinityAttributesSamesite = "None"
	SessionAffinityAttributesSamesiteStrict SessionAffinityAttributesSamesite = "Strict"
)

func (r SessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesSamesiteAuto, SessionAffinityAttributesSamesiteLax, SessionAffinityAttributesSamesiteNone, SessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type SessionAffinityAttributesSecure string

const (
	SessionAffinityAttributesSecureAuto   SessionAffinityAttributesSecure = "Auto"
	SessionAffinityAttributesSecureAlways SessionAffinityAttributesSecure = "Always"
	SessionAffinityAttributesSecureNever  SessionAffinityAttributesSecure = "Never"
)

func (r SessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesSecureAuto, SessionAffinityAttributesSecureAlways, SessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are:
//
//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
//     originally pinned origin is available; note that this can potentially result
//     in heavy origin flapping.
//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
//     currently not supported for session affinity by header.
type SessionAffinityAttributesZeroDowntimeFailover string

const (
	SessionAffinityAttributesZeroDowntimeFailoverNone      SessionAffinityAttributesZeroDowntimeFailover = "none"
	SessionAffinityAttributesZeroDowntimeFailoverTemporary SessionAffinityAttributesZeroDowntimeFailover = "temporary"
	SessionAffinityAttributesZeroDowntimeFailoverSticky    SessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r SessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesZeroDowntimeFailoverNone, SessionAffinityAttributesZeroDowntimeFailoverTemporary, SessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type SessionAffinityAttributesParam struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are:
	//
	//   - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	//     specified by the `headers` session affinity attribute, otherwise sessions
	//     aren't created.
	//   - `"false"`: Load balancing requests must contain _at least one_ of the HTTP
	//     headers specified by the `headers` session affinity attribute, otherwise
	//     sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[SessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[SessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are:
	//
	//   - `"none"`: No failover takes place for sessions pinned to the origin (default).
	//   - `"temporary"`: Traffic will be sent to another other healthy origin until the
	//     originally pinned origin is available; note that this can potentially result
	//     in heavy origin flapping.
	//   - `"sticky"`: The session affinity cookie is updated and subsequent requests are
	//     sent to the new origin. Note: Zero-downtime failover with sticky sessions is
	//     currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[SessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r SessionAffinityAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type SteeringPolicy string

const (
	SteeringPolicyOff                      SteeringPolicy = "off"
	SteeringPolicyGeo                      SteeringPolicy = "geo"
	SteeringPolicyRandom                   SteeringPolicy = "random"
	SteeringPolicyDynamicLatency           SteeringPolicy = "dynamic_latency"
	SteeringPolicyProximity                SteeringPolicy = "proximity"
	SteeringPolicyLeastOutstandingRequests SteeringPolicy = "least_outstanding_requests"
	SteeringPolicyLeastConnections         SteeringPolicy = "least_connections"
	SteeringPolicyEmpty                    SteeringPolicy = ""
)

func (r SteeringPolicy) IsKnown() bool {
	switch r {
	case SteeringPolicyOff, SteeringPolicyGeo, SteeringPolicyRandom, SteeringPolicyDynamicLatency, SteeringPolicyProximity, SteeringPolicyLeastOutstandingRequests, SteeringPolicyLeastConnections, SteeringPolicyEmpty:
		return true
	}
	return false
}

type ZoneLoadBalancerListResponse struct {
	Result []LoadBalancer                   `json:"result"`
	JSON   zoneLoadBalancerListResponseJSON `json:"-"`
	PaginatedResponseCollection
}

// zoneLoadBalancerListResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerListResponse]
type zoneLoadBalancerListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLoadBalancerListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLoadBalancerDeleteResponse struct {
	Result ZoneLoadBalancerDeleteResponseResult `json:"result"`
	JSON   zoneLoadBalancerDeleteResponseJSON   `json:"-"`
	SingleResponseMonitor
}

// zoneLoadBalancerDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneLoadBalancerDeleteResponse]
type zoneLoadBalancerDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLoadBalancerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneLoadBalancerDeleteResponseResult struct {
	ID   string                                   `json:"id"`
	JSON zoneLoadBalancerDeleteResponseResultJSON `json:"-"`
}

// zoneLoadBalancerDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneLoadBalancerDeleteResponseResult]
type zoneLoadBalancerDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneLoadBalancerDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneLoadBalancerDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneLoadBalancerNewParams struct {
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks param.Field[[]string] `json:"networks"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancingRuleParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneLoadBalancerUpdateParams struct {
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks param.Field[[]string] `json:"networks"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancingRuleParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneLoadBalancerDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneLoadBalancerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneLoadBalancerPatchParams struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]string] `json:"default_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name"`
	// (Enterprise only): A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	PopPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]LoadBalancingRuleParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are:
	//
	//   - `"cookie"`: On the first request to a proxied load balancer, a cookie is
	//     generated, encoding information of which origin the request will be forwarded
	//     to. Subsequent requests, by the same client to the same load balancer, will be
	//     sent to the origin server the cookie encodes, for the duration of the cookie
	//     and as long as the origin server remains healthy. If the cookie has expired or
	//     the origin server is unhealthy, then a new origin server is calculated and
	//     used.
	//   - `"ip_cookie"`: Behaves the same as `"cookie"` except the initial origin
	//     selection is stable and based on the client's ip address.
	//   - `"header"`: On the first request to a proxied load balancer, a session key
	//     based on the configured HTTP headers (see
	//     `session_affinity_attributes.headers`) is generated, encoding the request
	//     headers used for storing in the load balancer session state which origin the
	//     request will be forwarded to. Subsequent requests to the load balancer with
	//     the same headers will be sent to the same origin server, for the duration of
	//     the session and as long as the origin server remains healthy. If the session
	//     has been idle for the duration of `session_affinity_ttl` seconds or the origin
	//     server is unhealthy, then a new origin server is calculated and used. See
	//     `headers` in `session_affinity_attributes` for additional required
	//     configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are:
	//
	//   - `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used
	//     unless explicitly set. The accepted range of values is between [1800, 604800].
	//   - `"header"`: The current default of 1800 seconds will be used unless explicitly
	//     set. The accepted range of values is between [30, 3600]. Note: With session
	//     affinity by header, sessions only expire after they haven't been used for the
	//     number of seconds specified.
	SessionAffinityTtl param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	Ttl param.Field[float64] `json:"ttl"`
}

func (r ZoneLoadBalancerPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
