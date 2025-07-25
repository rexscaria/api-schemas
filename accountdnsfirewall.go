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

// AccountDNSFirewallService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallService] method instead.
type AccountDNSFirewallService struct {
	Options      []option.RequestOption
	DNSAnalytics *AccountDNSFirewallDNSAnalyticService
	ReverseDNS   *AccountDNSFirewallReverseDNSService
}

// NewAccountDNSFirewallService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDNSFirewallService(opts ...option.RequestOption) (r *AccountDNSFirewallService) {
	r = &AccountDNSFirewallService{}
	r.Options = opts
	r.DNSAnalytics = NewAccountDNSFirewallDNSAnalyticService(opts...)
	r.ReverseDNS = NewAccountDNSFirewallReverseDNSService(opts...)
	return
}

// Create a DNS Firewall cluster
func (r *AccountDNSFirewallService) New(ctx context.Context, accountID string, body AccountDNSFirewallNewParams, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Show a single DNS Firewall cluster for an account
func (r *AccountDNSFirewallService) Get(ctx context.Context, accountID string, dnsFirewallID string, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify the configuration of a DNS Firewall cluster
func (r *AccountDNSFirewallService) Update(ctx context.Context, accountID string, dnsFirewallID string, body AccountDNSFirewallUpdateParams, opts ...option.RequestOption) (res *DNSFirewallSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List DNS Firewall clusters for an account
func (r *AccountDNSFirewallService) List(ctx context.Context, accountID string, query AccountDNSFirewallListParams, opts ...option.RequestOption) (res *AccountDNSFirewallListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a DNS Firewall cluster
func (r *AccountDNSFirewallService) Delete(ctx context.Context, accountID string, dnsFirewallID string, opts ...option.RequestOption) (res *AccountDNSFirewallDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", accountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DNSFirewallClusterParam struct {
	// Attack mitigation settings
	AttackMitigation param.Field[DNSFirewallClusterAttackMitigationParam] `json:"attack_mitigation"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	EcsFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTtl param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTtl param.Field[float64] `json:"minimum_cache_ttl"`
	// DNS Firewall cluster name
	Name param.Field[string] `json:"name"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTtl param.Field[float64] `json:"negative_cache_ttl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     param.Field[float64]  `json:"retries"`
	UpstreamIPs param.Field[[]string] `json:"upstream_ips" format:"ipv4"`
}

func (r DNSFirewallClusterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attack mitigation settings
type DNSFirewallClusterAttackMitigationParam struct {
	// When enabled, automatically mitigate random-prefix attacks to protect upstream
	// DNS servers
	Enabled param.Field[bool] `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r DNSFirewallClusterAttackMitigationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallClusterResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests bool     `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []string `json:"dns_firewall_ips,required" format:"ipv4"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	EcsFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTtl float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTtl float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall cluster name
	Name string `json:"name,required"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTtl float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     float64  `json:"retries,required"`
	UpstreamIPs []string `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation DNSFirewallClusterResponseAttackMitigation `json:"attack_mitigation,nullable"`
	JSON             dnsFirewallClusterResponseJSON             `json:"-"`
}

// dnsFirewallClusterResponseJSON contains the JSON metadata for the struct
// [DNSFirewallClusterResponse]
type dnsFirewallClusterResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	EcsFallback          apijson.Field
	MaximumCacheTtl      apijson.Field
	MinimumCacheTtl      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTtl     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallClusterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallClusterResponseJSON) RawJSON() string {
	return r.raw
}

// Attack mitigation settings
type DNSFirewallClusterResponseAttackMitigation struct {
	// When enabled, automatically mitigate random-prefix attacks to protect upstream
	// DNS servers
	Enabled bool `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy
	OnlyWhenUpstreamUnhealthy bool                                           `json:"only_when_upstream_unhealthy"`
	JSON                      dnsFirewallClusterResponseAttackMitigationJSON `json:"-"`
}

// dnsFirewallClusterResponseAttackMitigationJSON contains the JSON metadata for
// the struct [DNSFirewallClusterResponseAttackMitigation]
type dnsFirewallClusterResponseAttackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *DNSFirewallClusterResponseAttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallClusterResponseAttackMitigationJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallSingleResponse struct {
	Errors   []MessagesDNSFirewallItem `json:"errors,required"`
	Messages []MessagesDNSFirewallItem `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSFirewallSingleResponseSuccess `json:"success,required"`
	Result  DNSFirewallClusterResponse       `json:"result"`
	JSON    dnsFirewallSingleResponseJSON    `json:"-"`
}

// dnsFirewallSingleResponseJSON contains the JSON metadata for the struct
// [DNSFirewallSingleResponse]
type dnsFirewallSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallSingleResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSFirewallSingleResponseSuccess bool

const (
	DNSFirewallSingleResponseSuccessTrue DNSFirewallSingleResponseSuccess = true
)

func (r DNSFirewallSingleResponseSuccess) IsKnown() bool {
	switch r {
	case DNSFirewallSingleResponseSuccessTrue:
		return true
	}
	return false
}

type MessagesDNSFirewallItem struct {
	Code             int64                         `json:"code,required"`
	Message          string                        `json:"message,required"`
	DocumentationURL string                        `json:"documentation_url"`
	Source           MessagesDNSFirewallItemSource `json:"source"`
	JSON             messagesDNSFirewallItemJSON   `json:"-"`
}

// messagesDNSFirewallItemJSON contains the JSON metadata for the struct
// [MessagesDNSFirewallItem]
type messagesDNSFirewallItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesDNSFirewallItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDNSFirewallItemJSON) RawJSON() string {
	return r.raw
}

type MessagesDNSFirewallItemSource struct {
	Pointer string                            `json:"pointer"`
	JSON    messagesDNSFirewallItemSourceJSON `json:"-"`
}

// messagesDNSFirewallItemSourceJSON contains the JSON metadata for the struct
// [MessagesDNSFirewallItemSource]
type messagesDNSFirewallItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesDNSFirewallItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesDNSFirewallItemSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallListResponse struct {
	Errors   []MessagesDNSFirewallItem `json:"errors,required"`
	Messages []MessagesDNSFirewallItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountDNSFirewallListResponseSuccess    `json:"success,required"`
	Result     []DNSFirewallClusterResponse             `json:"result"`
	ResultInfo AccountDNSFirewallListResponseResultInfo `json:"result_info"`
	JSON       accountDNSFirewallListResponseJSON       `json:"-"`
}

// accountDNSFirewallListResponseJSON contains the JSON metadata for the struct
// [AccountDNSFirewallListResponse]
type accountDNSFirewallListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDNSFirewallListResponseSuccess bool

const (
	AccountDNSFirewallListResponseSuccessTrue AccountDNSFirewallListResponseSuccess = true
)

func (r AccountDNSFirewallListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDNSFirewallListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDNSFirewallListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                      `json:"total_count"`
	JSON       accountDNSFirewallListResponseResultInfoJSON `json:"-"`
}

// accountDNSFirewallListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountDNSFirewallListResponseResultInfo]
type accountDNSFirewallListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallDeleteResponse struct {
	Errors   []MessagesDNSFirewallItem `json:"errors,required"`
	Messages []MessagesDNSFirewallItem `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountDNSFirewallDeleteResponseSuccess `json:"success,required"`
	Result  AccountDNSFirewallDeleteResponseResult  `json:"result"`
	JSON    accountDNSFirewallDeleteResponseJSON    `json:"-"`
}

// accountDNSFirewallDeleteResponseJSON contains the JSON metadata for the struct
// [AccountDNSFirewallDeleteResponse]
type accountDNSFirewallDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountDNSFirewallDeleteResponseSuccess bool

const (
	AccountDNSFirewallDeleteResponseSuccessTrue AccountDNSFirewallDeleteResponseSuccess = true
)

func (r AccountDNSFirewallDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountDNSFirewallDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountDNSFirewallDeleteResponseResult struct {
	// Identifier.
	ID   string                                     `json:"id"`
	JSON accountDNSFirewallDeleteResponseResultJSON `json:"-"`
}

// accountDNSFirewallDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountDNSFirewallDeleteResponseResult]
type accountDNSFirewallDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountDNSFirewallDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountDNSFirewallNewParams struct {
	DNSFirewallCluster DNSFirewallClusterParam `json:"dns_firewall_cluster,required"`
}

func (r AccountDNSFirewallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSFirewallCluster)
}

type AccountDNSFirewallUpdateParams struct {
	DNSFirewallCluster DNSFirewallClusterParam `json:"dns_firewall_cluster,required"`
}

func (r AccountDNSFirewallUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DNSFirewallCluster)
}

type AccountDNSFirewallListParams struct {
	// Page number of paginated results
	Page param.Field[float64] `query:"page"`
	// Number of clusters per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountDNSFirewallListParams]'s query parameters as
// `url.Values`.
func (r AccountDNSFirewallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
