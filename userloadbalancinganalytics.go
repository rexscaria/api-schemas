// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// UserLoadBalancingAnalyticsService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserLoadBalancingAnalyticsService] method instead.
type UserLoadBalancingAnalyticsService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancingAnalyticsService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancingAnalyticsService(opts ...option.RequestOption) (r *UserLoadBalancingAnalyticsService) {
	r = &UserLoadBalancingAnalyticsService{}
	r.Options = opts
	return
}

// List origin health changes.
func (r *UserLoadBalancingAnalyticsService) ListEvents(ctx context.Context, query UserLoadBalancingAnalyticsListEventsParams, opts ...option.RequestOption) (res *UserLoadBalancingAnalyticsListEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserLoadBalancingAnalyticsListEventsResponse struct {
	Errors   []UserLoadBalancingAnalyticsListEventsResponseError   `json:"errors,required"`
	Messages []UserLoadBalancingAnalyticsListEventsResponseMessage `json:"messages,required"`
	Result   []UserLoadBalancingAnalyticsListEventsResponseResult  `json:"result,required"`
	// Whether the API call was successful
	Success    UserLoadBalancingAnalyticsListEventsResponseSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancingAnalyticsListEventsResponseResultInfo `json:"result_info"`
	JSON       userLoadBalancingAnalyticsListEventsResponseJSON       `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseJSON contains the JSON metadata for
// the struct [UserLoadBalancingAnalyticsListEventsResponse]
type userLoadBalancingAnalyticsListEventsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseError struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           UserLoadBalancingAnalyticsListEventsResponseErrorsSource `json:"source"`
	JSON             userLoadBalancingAnalyticsListEventsResponseErrorJSON    `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseErrorJSON contains the JSON metadata
// for the struct [UserLoadBalancingAnalyticsListEventsResponseError]
type userLoadBalancingAnalyticsListEventsResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    userLoadBalancingAnalyticsListEventsResponseErrorsSourceJSON `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseErrorsSourceJSON contains the JSON
// metadata for the struct
// [UserLoadBalancingAnalyticsListEventsResponseErrorsSource]
type userLoadBalancingAnalyticsListEventsResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseMessage struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           UserLoadBalancingAnalyticsListEventsResponseMessagesSource `json:"source"`
	JSON             userLoadBalancingAnalyticsListEventsResponseMessageJSON    `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseMessageJSON contains the JSON
// metadata for the struct [UserLoadBalancingAnalyticsListEventsResponseMessage]
type userLoadBalancingAnalyticsListEventsResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    userLoadBalancingAnalyticsListEventsResponseMessagesSourceJSON `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [UserLoadBalancingAnalyticsListEventsResponseMessagesSource]
type userLoadBalancingAnalyticsListEventsResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseResult struct {
	ID        int64                                                      `json:"id"`
	Origins   []UserLoadBalancingAnalyticsListEventsResponseResultOrigin `json:"origins"`
	Pool      interface{}                                                `json:"pool"`
	Timestamp time.Time                                                  `json:"timestamp" format:"date-time"`
	JSON      userLoadBalancingAnalyticsListEventsResponseResultJSON     `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseResultJSON contains the JSON
// metadata for the struct [UserLoadBalancingAnalyticsListEventsResponseResult]
type userLoadBalancingAnalyticsListEventsResponseResultJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsResponseResultOrigin struct {
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
	Name string                                                       `json:"name"`
	JSON userLoadBalancingAnalyticsListEventsResponseResultOriginJSON `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseResultOriginJSON contains the JSON
// metadata for the struct
// [UserLoadBalancingAnalyticsListEventsResponseResultOrigin]
type userLoadBalancingAnalyticsListEventsResponseResultOriginJSON struct {
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

func (r *UserLoadBalancingAnalyticsListEventsResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseResultOriginJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserLoadBalancingAnalyticsListEventsResponseSuccess bool

const (
	UserLoadBalancingAnalyticsListEventsResponseSuccessTrue UserLoadBalancingAnalyticsListEventsResponseSuccess = true
)

func (r UserLoadBalancingAnalyticsListEventsResponseSuccess) IsKnown() bool {
	switch r {
	case UserLoadBalancingAnalyticsListEventsResponseSuccessTrue:
		return true
	}
	return false
}

type UserLoadBalancingAnalyticsListEventsResponseResultInfo struct {
	// Total number of results on the current page
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total number of pages available
	TotalPages float64                                                    `json:"total_pages"`
	JSON       userLoadBalancingAnalyticsListEventsResponseResultInfoJSON `json:"-"`
}

// userLoadBalancingAnalyticsListEventsResponseResultInfoJSON contains the JSON
// metadata for the struct [UserLoadBalancingAnalyticsListEventsResponseResultInfo]
type userLoadBalancingAnalyticsListEventsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticsListEventsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoadBalancingAnalyticsListEventsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type UserLoadBalancingAnalyticsListEventsParams struct {
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

// URLQuery serializes [UserLoadBalancingAnalyticsListEventsParams]'s query
// parameters as `url.Values`.
func (r UserLoadBalancingAnalyticsListEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
