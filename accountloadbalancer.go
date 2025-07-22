// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountLoadBalancerService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLoadBalancerService] method instead.
type AccountLoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *AccountLoadBalancerMonitorService
	Pools    *AccountLoadBalancerPoolService
	Regions  *AccountLoadBalancerRegionService
}

// NewAccountLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerService(opts ...option.RequestOption) (r *AccountLoadBalancerService) {
	r = &AccountLoadBalancerService{}
	r.Options = opts
	r.Monitors = NewAccountLoadBalancerMonitorService(opts...)
	r.Pools = NewAccountLoadBalancerPoolService(opts...)
	r.Regions = NewAccountLoadBalancerRegionService(opts...)
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *AccountLoadBalancerService) PreviewResult(ctx context.Context, accountID string, previewID string, opts ...option.RequestOption) (res *PreviewResultResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if previewID == "" {
		err = errors.New("missing required preview_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%s", accountID, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Search for Load Balancing resources.
func (r *AccountLoadBalancerService) Search(ctx context.Context, accountID string, query AccountLoadBalancerSearchParams, opts ...option.RequestOption) (res *AccountLoadBalancerSearchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/search", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PaginatedResponseCollection struct {
	ResultInfo PaginatedResponseCollectionResultInfo `json:"result_info"`
	JSON       paginatedResponseCollectionJSON       `json:"-"`
	CommonResponseLoadBalancers
}

// paginatedResponseCollectionJSON contains the JSON metadata for the struct
// [PaginatedResponseCollection]
type paginatedResponseCollectionJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type PaginatedResponseCollectionResultInfo struct {
	// Total number of results on the current page
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64 `json:"total_count"`
	// Total number of pages available
	TotalPages float64                                   `json:"total_pages"`
	JSON       paginatedResponseCollectionResultInfoJSON `json:"-"`
}

// paginatedResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [PaginatedResponseCollectionResultInfo]
type paginatedResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PaginatedResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r paginatedResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type PreviewResultResponse struct {
	// Resulting health data from a preview operation.
	Result map[string]PreviewResultResponseResult `json:"result"`
	JSON   previewResultResponseJSON              `json:"-"`
	SingleResponseMonitor
}

// previewResultResponseJSON contains the JSON metadata for the struct
// [PreviewResultResponse]
type previewResultResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResultResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewResultResponseJSON) RawJSON() string {
	return r.raw
}

type PreviewResultResponseResult struct {
	Healthy bool                                           `json:"healthy"`
	Origins []map[string]PreviewResultResponseResultOrigin `json:"origins"`
	JSON    previewResultResponseResultJSON                `json:"-"`
}

// previewResultResponseResultJSON contains the JSON metadata for the struct
// [PreviewResultResponseResult]
type previewResultResponseResultJSON struct {
	Healthy     apijson.Field
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PreviewResultResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewResultResponseResultJSON) RawJSON() string {
	return r.raw
}

// The origin ipv4/ipv6 address or domain name mapped to it's health data.
type PreviewResultResponseResultOrigin struct {
	FailureReason string                                `json:"failure_reason"`
	Healthy       bool                                  `json:"healthy"`
	ResponseCode  float64                               `json:"response_code"`
	Rtt           string                                `json:"rtt"`
	JSON          previewResultResponseResultOriginJSON `json:"-"`
}

// previewResultResponseResultOriginJSON contains the JSON metadata for the struct
// [PreviewResultResponseResultOrigin]
type previewResultResponseResultOriginJSON struct {
	FailureReason apijson.Field
	Healthy       apijson.Field
	ResponseCode  apijson.Field
	Rtt           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PreviewResultResponseResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r previewResultResponseResultOriginJSON) RawJSON() string {
	return r.raw
}

type AccountLoadBalancerSearchResponse struct {
	Result AccountLoadBalancerSearchResponseResult `json:"result"`
	JSON   accountLoadBalancerSearchResponseJSON   `json:"-"`
	PaginatedResponseCollection
}

// accountLoadBalancerSearchResponseJSON contains the JSON metadata for the struct
// [AccountLoadBalancerSearchResponse]
type accountLoadBalancerSearchResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLoadBalancerSearchResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLoadBalancerSearchResponseResult struct {
	// A list of resources matching the search query.
	Resources []AccountLoadBalancerSearchResponseResultResource `json:"resources"`
	JSON      accountLoadBalancerSearchResponseResultJSON       `json:"-"`
}

// accountLoadBalancerSearchResponseResultJSON contains the JSON metadata for the
// struct [AccountLoadBalancerSearchResponseResult]
type accountLoadBalancerSearchResponseResultJSON struct {
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLoadBalancerSearchResponseResultJSON) RawJSON() string {
	return r.raw
}

// A reference to a load balancer resource.
type AccountLoadBalancerSearchResponseResultResource struct {
	// When listed as a reference, the type (direction) of the reference.
	ReferenceType AccountLoadBalancerSearchResponseResultResourcesReferenceType `json:"reference_type"`
	// A list of references to (referrer) or from (referral) this resource.
	References []interface{} `json:"references"`
	ResourceID interface{}   `json:"resource_id"`
	// The human-identifiable name of the resource.
	ResourceName string `json:"resource_name"`
	// The type of the resource.
	ResourceType AccountLoadBalancerSearchResponseResultResourcesResourceType `json:"resource_type"`
	JSON         accountLoadBalancerSearchResponseResultResourceJSON          `json:"-"`
}

// accountLoadBalancerSearchResponseResultResourceJSON contains the JSON metadata
// for the struct [AccountLoadBalancerSearchResponseResultResource]
type accountLoadBalancerSearchResponseResultResourceJSON struct {
	ReferenceType apijson.Field
	References    apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountLoadBalancerSearchResponseResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLoadBalancerSearchResponseResultResourceJSON) RawJSON() string {
	return r.raw
}

// When listed as a reference, the type (direction) of the reference.
type AccountLoadBalancerSearchResponseResultResourcesReferenceType string

const (
	AccountLoadBalancerSearchResponseResultResourcesReferenceTypeReferral AccountLoadBalancerSearchResponseResultResourcesReferenceType = "referral"
	AccountLoadBalancerSearchResponseResultResourcesReferenceTypeReferrer AccountLoadBalancerSearchResponseResultResourcesReferenceType = "referrer"
)

func (r AccountLoadBalancerSearchResponseResultResourcesReferenceType) IsKnown() bool {
	switch r {
	case AccountLoadBalancerSearchResponseResultResourcesReferenceTypeReferral, AccountLoadBalancerSearchResponseResultResourcesReferenceTypeReferrer:
		return true
	}
	return false
}

// The type of the resource.
type AccountLoadBalancerSearchResponseResultResourcesResourceType string

const (
	AccountLoadBalancerSearchResponseResultResourcesResourceTypeLoadBalancer AccountLoadBalancerSearchResponseResultResourcesResourceType = "load_balancer"
	AccountLoadBalancerSearchResponseResultResourcesResourceTypeMonitor      AccountLoadBalancerSearchResponseResultResourcesResourceType = "monitor"
	AccountLoadBalancerSearchResponseResultResourcesResourceTypePool         AccountLoadBalancerSearchResponseResultResourcesResourceType = "pool"
)

func (r AccountLoadBalancerSearchResponseResultResourcesResourceType) IsKnown() bool {
	switch r {
	case AccountLoadBalancerSearchResponseResultResourcesResourceTypeLoadBalancer, AccountLoadBalancerSearchResponseResultResourcesResourceTypeMonitor, AccountLoadBalancerSearchResponseResultResourcesResourceTypePool:
		return true
	}
	return false
}

type AccountLoadBalancerSearchParams struct {
	Page         param.Field[float64]                                     `query:"page"`
	PerPage      param.Field[float64]                                     `query:"per_page"`
	SearchParams param.Field[AccountLoadBalancerSearchParamsSearchParams] `query:"search_params"`
}

// URLQuery serializes [AccountLoadBalancerSearchParams]'s query parameters as
// `url.Values`.
func (r AccountLoadBalancerSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLoadBalancerSearchParamsSearchParams struct {
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include ("\*" for all).
	References param.Field[AccountLoadBalancerSearchParamsSearchParamsReferences] `query:"references"`
}

// URLQuery serializes [AccountLoadBalancerSearchParamsSearchParams]'s query
// parameters as `url.Values`.
func (r AccountLoadBalancerSearchParamsSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The type of references to include ("\*" for all).
type AccountLoadBalancerSearchParamsSearchParamsReferences string

const (
	AccountLoadBalancerSearchParamsSearchParamsReferencesEmpty    AccountLoadBalancerSearchParamsSearchParamsReferences = ""
	AccountLoadBalancerSearchParamsSearchParamsReferencesStar     AccountLoadBalancerSearchParamsSearchParamsReferences = "*"
	AccountLoadBalancerSearchParamsSearchParamsReferencesReferral AccountLoadBalancerSearchParamsSearchParamsReferences = "referral"
	AccountLoadBalancerSearchParamsSearchParamsReferencesReferrer AccountLoadBalancerSearchParamsSearchParamsReferences = "referrer"
)

func (r AccountLoadBalancerSearchParamsSearchParamsReferences) IsKnown() bool {
	switch r {
	case AccountLoadBalancerSearchParamsSearchParamsReferencesEmpty, AccountLoadBalancerSearchParamsSearchParamsReferencesStar, AccountLoadBalancerSearchParamsSearchParamsReferencesReferral, AccountLoadBalancerSearchParamsSearchParamsReferencesReferrer:
		return true
	}
	return false
}
