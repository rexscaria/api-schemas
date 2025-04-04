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

// UserLoadBalancingAnalyticService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserLoadBalancingAnalyticService] method instead.
type UserLoadBalancingAnalyticService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancingAnalyticService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancingAnalyticService(opts ...option.RequestOption) (r *UserLoadBalancingAnalyticService) {
	r = &UserLoadBalancingAnalyticService{}
	r.Options = opts
	return
}

// List origin health changes.
func (r *UserLoadBalancingAnalyticService) ListEvents(ctx context.Context, query UserLoadBalancingAnalyticListEventsParams, opts ...option.RequestOption) (res *UserLoadBalancingAnalyticListEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserLoadBalancingAnalyticListEventsResponse struct {
	Result []UserLoadBalancingAnalyticListEventsResponseResult `json:"result"`
	JSON   userLoadBalancingAnalyticListEventsResponseJSON     `json:"-"`
	PaginatedResponseCollection
}

// userLoadBalancingAnalyticListEventsResponseJSON contains the JSON metadata for
// the struct [UserLoadBalancingAnalyticListEventsResponse]
type userLoadBalancingAnalyticListEventsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticListEventsResponseResult struct {
	ID        int64                                                     `json:"id"`
	Origins   []UserLoadBalancingAnalyticListEventsResponseResultOrigin `json:"origins"`
	Pool      interface{}                                               `json:"pool"`
	Timestamp time.Time                                                 `json:"timestamp" format:"date-time"`
	JSON      userLoadBalancingAnalyticListEventsResponseResultJSON     `json:"-"`
}

// userLoadBalancingAnalyticListEventsResponseResultJSON contains the JSON metadata
// for the struct [UserLoadBalancingAnalyticListEventsResponseResult]
type userLoadBalancingAnalyticListEventsResponseResultJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticListEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticListEventsResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticListEventsResponseResultOrigin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// Whether the origin has changed health status.
	Changed bool `json:"changed"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// Failure reason for un-healthy origin health check.
	FailureReason string `json:"failure_reason"`
	// Whether the origin is reported as healthy.
	Healthy bool `json:"healthy"`
	// The IP address (IPv4 or IPv6) of the origin.
	IP string `json:"ip"`
	// A human-identifiable name for the origin.
	Name string                                                      `json:"name"`
	JSON userLoadBalancingAnalyticListEventsResponseResultOriginJSON `json:"-"`
}

// userLoadBalancingAnalyticListEventsResponseResultOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancingAnalyticListEventsResponseResultOrigin]
type userLoadBalancingAnalyticListEventsResponseResultOriginJSON struct {
	Address       apijson.Field
	Changed       apijson.Field
	Enabled       apijson.Field
	FailureReason apijson.Field
	Healthy       apijson.Field
	IP            apijson.Field
	Name          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticListEventsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticListEventsResponseResultOriginJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticListEventsParams struct {
	// If true, filter events where the origin status is healthy. If false, filter
	// events where the origin status is unhealthy.
	OriginHealthy param.Field[bool] `query:"origin_healthy"`
	// The name for the origin to filter.
	OriginName param.Field[string] `query:"origin_name"`
	// If true, filter events where the pool status is healthy. If false, filter events
	// where the pool status is unhealthy.
	PoolHealthy param.Field[bool]   `query:"pool_healthy"`
	PoolID      param.Field[string] `query:"pool_id"`
	// The name for the pool to filter.
	PoolName param.Field[string] `query:"pool_name"`
	// Start date and time of requesting data period in the ISO8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// End date and time of requesting data period in the ISO8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [UserLoadBalancingAnalyticListEventsParams]'s query
// parameters as `url.Values`.
func (r UserLoadBalancingAnalyticListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
