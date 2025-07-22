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

// AccountInfrastructureTargetService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountInfrastructureTargetService] method instead.
type AccountInfrastructureTargetService struct {
	Options []option.RequestOption
	Batch   *AccountInfrastructureTargetBatchService
}

// NewAccountInfrastructureTargetService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountInfrastructureTargetService(opts ...option.RequestOption) (r *AccountInfrastructureTargetService) {
	r = &AccountInfrastructureTargetService{}
	r.Options = opts
	r.Batch = NewAccountInfrastructureTargetBatchService(opts...)
	return
}

// Create new target
func (r *AccountInfrastructureTargetService) New(ctx context.Context, accountID string, body AccountInfrastructureTargetNewParams, opts ...option.RequestOption) (res *AccountInfrastructureTargetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get target
func (r *AccountInfrastructureTargetService) Get(ctx context.Context, accountID string, targetID string, opts ...option.RequestOption) (res *AccountInfrastructureTargetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", accountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update target
func (r *AccountInfrastructureTargetService) Update(ctx context.Context, accountID string, targetID string, body AccountInfrastructureTargetUpdateParams, opts ...option.RequestOption) (res *AccountInfrastructureTargetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", accountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists and sorts an account’s targets. Filters are optional and are ANDed
// together.
func (r *AccountInfrastructureTargetService) List(ctx context.Context, accountID string, query AccountInfrastructureTargetListParams, opts ...option.RequestOption) (res *AccountInfrastructureTargetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete target
func (r *AccountInfrastructureTargetService) Delete(ctx context.Context, accountID string, targetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/infrastructure/targets/%s", accountID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type APIResponseInfrastructure struct {
	Errors   []MessagesInfraItem `json:"errors,required"`
	Messages []MessagesInfraItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseInfrastructureSuccess `json:"success,required"`
	JSON    apiResponseInfrastructureJSON    `json:"-"`
}

// apiResponseInfrastructureJSON contains the JSON metadata for the struct
// [APIResponseInfrastructure]
type apiResponseInfrastructureJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseInfrastructure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseInfrastructureJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseInfrastructureSuccess bool

const (
	APIResponseInfrastructureSuccessTrue APIResponseInfrastructureSuccess = true
)

func (r APIResponseInfrastructureSuccess) IsKnown() bool {
	switch r {
	case APIResponseInfrastructureSuccessTrue:
		return true
	}
	return false
}

type APIResponseSingleInfrastructure struct {
	Errors   []MessagesInfraItem `json:"errors,required"`
	Messages []MessagesInfraItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleInfrastructureSuccess `json:"success,required"`
	JSON    apiResponseSingleInfrastructureJSON    `json:"-"`
}

// apiResponseSingleInfrastructureJSON contains the JSON metadata for the struct
// [APIResponseSingleInfrastructure]
type apiResponseSingleInfrastructureJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleInfrastructure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleInfrastructureJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleInfrastructureSuccess bool

const (
	APIResponseSingleInfrastructureSuccessTrue APIResponseSingleInfrastructureSuccess = true
)

func (r APIResponseSingleInfrastructureSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleInfrastructureSuccessTrue:
		return true
	}
	return false
}

// The IPv4/IPv6 address that identifies where to reach a target
type IPInfoTarget struct {
	// The target's IPv4 address
	Ipv4 IPInfoTargetIpv4 `json:"ipv4"`
	// The target's IPv6 address
	Ipv6 IPInfoTargetIpv6 `json:"ipv6"`
	JSON ipInfoTargetJSON `json:"-"`
}

// ipInfoTargetJSON contains the JSON metadata for the struct [IPInfoTarget]
type ipInfoTargetJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPInfoTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipInfoTargetJSON) RawJSON() string {
	return r.raw
}

// The target's IPv4 address
type IPInfoTargetIpv4 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string               `json:"virtual_network_id" format:"uuid"`
	JSON             ipInfoTargetIpv4JSON `json:"-"`
}

// ipInfoTargetIpv4JSON contains the JSON metadata for the struct
// [IPInfoTargetIpv4]
type ipInfoTargetIpv4JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPInfoTargetIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipInfoTargetIpv4JSON) RawJSON() string {
	return r.raw
}

// The target's IPv6 address
type IPInfoTargetIpv6 struct {
	// IP address of the target
	IPAddr string `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID string               `json:"virtual_network_id" format:"uuid"`
	JSON             ipInfoTargetIpv6JSON `json:"-"`
}

// ipInfoTargetIpv6JSON contains the JSON metadata for the struct
// [IPInfoTargetIpv6]
type ipInfoTargetIpv6JSON struct {
	IPAddr           apijson.Field
	VirtualNetworkID apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPInfoTargetIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipInfoTargetIpv6JSON) RawJSON() string {
	return r.raw
}

// The IPv4/IPv6 address that identifies where to reach a target
type IPInfoTargetParam struct {
	// The target's IPv4 address
	Ipv4 param.Field[IPInfoTargetIpv4Param] `json:"ipv4"`
	// The target's IPv6 address
	Ipv6 param.Field[IPInfoTargetIpv6Param] `json:"ipv6"`
}

func (r IPInfoTargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv4 address
type IPInfoTargetIpv4Param struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r IPInfoTargetIpv4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The target's IPv6 address
type IPInfoTargetIpv6Param struct {
	// IP address of the target
	IPAddr param.Field[string] `json:"ip_addr"`
	// (optional) Private virtual network identifier for the target. If omitted, the
	// default virtual network ID will be used.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r IPInfoTargetIpv6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MessagesInfraItem struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    messagesInfraItemJSON `json:"-"`
}

// messagesInfraItemJSON contains the JSON metadata for the struct
// [MessagesInfraItem]
type messagesInfraItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesInfraItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesInfraItemJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetNewResponse struct {
	Result TargetBatch                                `json:"result"`
	JSON   accountInfrastructureTargetNewResponseJSON `json:"-"`
	APIResponseSingleInfrastructure
}

// accountInfrastructureTargetNewResponseJSON contains the JSON metadata for the
// struct [AccountInfrastructureTargetNewResponse]
type accountInfrastructureTargetNewResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetGetResponse struct {
	Result TargetBatch                                `json:"result"`
	JSON   accountInfrastructureTargetGetResponseJSON `json:"-"`
	APIResponseSingleInfrastructure
}

// accountInfrastructureTargetGetResponseJSON contains the JSON metadata for the
// struct [AccountInfrastructureTargetGetResponse]
type accountInfrastructureTargetGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetUpdateResponse struct {
	Result TargetBatch                                   `json:"result"`
	JSON   accountInfrastructureTargetUpdateResponseJSON `json:"-"`
	APIResponseSingleInfrastructure
}

// accountInfrastructureTargetUpdateResponseJSON contains the JSON metadata for the
// struct [AccountInfrastructureTargetUpdateResponse]
type accountInfrastructureTargetUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetListResponse struct {
	Result     []TargetBatch                                     `json:"result"`
	ResultInfo AccountInfrastructureTargetListResponseResultInfo `json:"result_info"`
	JSON       accountInfrastructureTargetListResponseJSON       `json:"-"`
	APIResponseInfrastructure
}

// accountInfrastructureTargetListResponseJSON contains the JSON metadata for the
// struct [AccountInfrastructureTargetListResponse]
type accountInfrastructureTargetListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       accountInfrastructureTargetListResponseResultInfoJSON `json:"-"`
}

// accountInfrastructureTargetListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountInfrastructureTargetListResponseResultInfo]
type accountInfrastructureTargetListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountInfrastructureTargetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountInfrastructureTargetListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountInfrastructureTargetNewParams struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[IPInfoTargetParam] `json:"ip,required"`
}

func (r AccountInfrastructureTargetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountInfrastructureTargetUpdateParams struct {
	// A non-unique field that refers to a target. Case insensitive, maximum length of
	// 255 characters, supports the use of special characters dash and period, does not
	// support spaces, and must start and end with an alphanumeric character.
	Hostname param.Field[string] `json:"hostname,required"`
	// The IPv4/IPv6 address that identifies where to reach a target
	IP param.Field[IPInfoTargetParam] `json:"ip,required"`
}

func (r AccountInfrastructureTargetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountInfrastructureTargetListParams struct {
	// Date and time at which the target was created after (inclusive)
	CreatedAfter param.Field[time.Time] `query:"created_after" format:"date-time"`
	// Date and time at which the target was created before (inclusive)
	CreatedBefore param.Field[time.Time] `query:"created_before" format:"date-time"`
	// The sorting direction.
	Direction param.Field[AccountInfrastructureTargetListParamsDirection] `query:"direction"`
	// Hostname of a target
	Hostname param.Field[string] `query:"hostname"`
	// Partial match to the hostname of a target
	HostnameContains param.Field[string] `query:"hostname_contains"`
	// Filters for targets whose IP addresses look like the specified string. Supports
	// `*` as a wildcard character
	IPLike param.Field[string] `query:"ip_like"`
	// IPv4 address of the target
	IPV4 param.Field[string] `query:"ip_v4"`
	// IPv6 address of the target
	IPV6 param.Field[string] `query:"ip_v6"`
	// Filters for targets that have any of the following IP addresses. Specify `ips`
	// multiple times in query parameter to build list of candidates.
	IPs param.Field[[]string] `query:"ips"`
	// Defines an IPv4 filter range's ending value (inclusive). Requires `ipv4_start`
	// to be specified as well.
	Ipv4End param.Field[string] `query:"ipv4_end"`
	// Defines an IPv4 filter range's starting value (inclusive). Requires `ipv4_end`
	// to be specified as well.
	Ipv4Start param.Field[string] `query:"ipv4_start"`
	// Defines an IPv6 filter range's ending value (inclusive). Requires `ipv6_start`
	// to be specified as well.
	Ipv6End param.Field[string] `query:"ipv6_end"`
	// Defines an IPv6 filter range's starting value (inclusive). Requires `ipv6_end`
	// to be specified as well.
	Ipv6Start param.Field[string] `query:"ipv6_start"`
	// Date and time at which the target was modified after (inclusive)
	ModifiedAfter param.Field[time.Time] `query:"modified_after" format:"date-time"`
	// Date and time at which the target was modified before (inclusive)
	ModifiedBefore param.Field[time.Time] `query:"modified_before" format:"date-time"`
	// The field to sort by.
	Order param.Field[AccountInfrastructureTargetListParamsOrder] `query:"order"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64] `query:"per_page"`
	// Filters for targets that have any of the following UUIDs. Specify `target_ids`
	// multiple times in query parameter to build list of candidates.
	TargetIDs param.Field[[]string] `query:"target_ids" format:"uuid"`
	// Private virtual network identifier of the target
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [AccountInfrastructureTargetListParams]'s query parameters
// as `url.Values`.
func (r AccountInfrastructureTargetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The sorting direction.
type AccountInfrastructureTargetListParamsDirection string

const (
	AccountInfrastructureTargetListParamsDirectionAsc  AccountInfrastructureTargetListParamsDirection = "asc"
	AccountInfrastructureTargetListParamsDirectionDesc AccountInfrastructureTargetListParamsDirection = "desc"
)

func (r AccountInfrastructureTargetListParamsDirection) IsKnown() bool {
	switch r {
	case AccountInfrastructureTargetListParamsDirectionAsc, AccountInfrastructureTargetListParamsDirectionDesc:
		return true
	}
	return false
}

// The field to sort by.
type AccountInfrastructureTargetListParamsOrder string

const (
	AccountInfrastructureTargetListParamsOrderHostname  AccountInfrastructureTargetListParamsOrder = "hostname"
	AccountInfrastructureTargetListParamsOrderCreatedAt AccountInfrastructureTargetListParamsOrder = "created_at"
)

func (r AccountInfrastructureTargetListParamsOrder) IsKnown() bool {
	switch r {
	case AccountInfrastructureTargetListParamsOrderHostname, AccountInfrastructureTargetListParamsOrderCreatedAt:
		return true
	}
	return false
}
