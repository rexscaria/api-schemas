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

// AccountSecondaryDNSACLService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSecondaryDNSACLService] method instead.
type AccountSecondaryDNSACLService struct {
	Options []option.RequestOption
}

// NewAccountSecondaryDNSACLService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSecondaryDNSACLService(opts ...option.RequestOption) (r *AccountSecondaryDNSACLService) {
	r = &AccountSecondaryDNSACLService{}
	r.Options = opts
	return
}

// Create ACL.
func (r *AccountSecondaryDNSACLService) New(ctx context.Context, accountID string, body AccountSecondaryDNSACLNewParams, opts ...option.RequestOption) (res *SchemasSecondaryDNSComponentsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get ACL.
func (r *AccountSecondaryDNSACLService) Get(ctx context.Context, accountID string, aclID string, opts ...option.RequestOption) (res *SchemasSecondaryDNSComponentsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify ACL.
func (r *AccountSecondaryDNSACLService) Update(ctx context.Context, accountID string, aclID string, body AccountSecondaryDNSACLUpdateParams, opts ...option.RequestOption) (res *SchemasSecondaryDNSComponentsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List ACLs.
func (r *AccountSecondaryDNSACLService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSecondaryDnsaclListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete ACL.
func (r *AccountSecondaryDNSACLService) Delete(ctx context.Context, accountID string, aclID string, body AccountSecondaryDNSACLDeleteParams, opts ...option.RequestOption) (res *SchemasIDResponseSecondaryDNS, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", accountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ACL struct {
	ID string `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string  `json:"name,required"`
	JSON aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

type ACLParam struct {
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r ACLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseCollectionDNSACLs struct {
	ResultInfo ResponseCollectionDNSACLsResultInfo `json:"result_info"`
	JSON       responseCollectionDnsacLsJSON       `json:"-"`
	ResponseCommonACLs
}

// responseCollectionDnsacLsJSON contains the JSON metadata for the struct
// [ResponseCollectionDNSACLs]
type responseCollectionDnsacLsJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDNSACLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDnsacLsJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionDNSACLsResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       responseCollectionDnsacLsResultInfoJSON `json:"-"`
}

// responseCollectionDnsacLsResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionDNSACLsResultInfo]
type responseCollectionDnsacLsResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionDNSACLsResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionDnsacLsResultInfoJSON) RawJSON() string {
	return r.raw
}

type ResponseCommonACLs struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful
	Success ResponseCommonACLsSuccess `json:"success,required"`
	JSON    responseCommonACLsJSON    `json:"-"`
}

// responseCommonACLsJSON contains the JSON metadata for the struct
// [ResponseCommonACLs]
type responseCommonACLsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCommonACLs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCommonACLsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseCommonACLsSuccess bool

const (
	ResponseCommonACLsSuccessTrue ResponseCommonACLsSuccess = true
)

func (r ResponseCommonACLsSuccess) IsKnown() bool {
	switch r {
	case ResponseCommonACLsSuccessTrue:
		return true
	}
	return false
}

type ResponseSingleACL struct {
	Errors   []SecondaryDNSMessages `json:"errors,required"`
	Messages []SecondaryDNSMessages `json:"messages,required"`
	// Whether the API call was successful
	Success ResponseSingleACLSuccess `json:"success,required"`
	JSON    responseSingleACLJSON    `json:"-"`
}

// responseSingleACLJSON contains the JSON metadata for the struct
// [ResponseSingleACL]
type responseSingleACLJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleACLJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseSingleACLSuccess bool

const (
	ResponseSingleACLSuccessTrue ResponseSingleACLSuccess = true
)

func (r ResponseSingleACLSuccess) IsKnown() bool {
	switch r {
	case ResponseSingleACLSuccessTrue:
		return true
	}
	return false
}

type SchemasIDResponseSecondaryDNS struct {
	Result SchemasIDResponseSecondaryDNSResult `json:"result"`
	JSON   schemasIDResponseSecondaryDNSJSON   `json:"-"`
	ResponseSingleACL
}

// schemasIDResponseSecondaryDNSJSON contains the JSON metadata for the struct
// [SchemasIDResponseSecondaryDNS]
type schemasIDResponseSecondaryDNSJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseSecondaryDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseSecondaryDNSJSON) RawJSON() string {
	return r.raw
}

type SchemasIDResponseSecondaryDNSResult struct {
	ID   string                                  `json:"id"`
	JSON schemasIDResponseSecondaryDNSResultJSON `json:"-"`
}

// schemasIDResponseSecondaryDNSResultJSON contains the JSON metadata for the
// struct [SchemasIDResponseSecondaryDNSResult]
type schemasIDResponseSecondaryDNSResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasIDResponseSecondaryDNSResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasIDResponseSecondaryDNSResultJSON) RawJSON() string {
	return r.raw
}

type SchemasSecondaryDNSComponentsSingleResponse struct {
	Result ACL                                             `json:"result"`
	JSON   schemasSecondaryDNSComponentsSingleResponseJSON `json:"-"`
	ResponseSingleACL
}

// schemasSecondaryDNSComponentsSingleResponseJSON contains the JSON metadata for
// the struct [SchemasSecondaryDNSComponentsSingleResponse]
type schemasSecondaryDNSComponentsSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasSecondaryDNSComponentsSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemasSecondaryDNSComponentsSingleResponseJSON) RawJSON() string {
	return r.raw
}

type SecondaryDNSMessages struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    secondaryDNSMessagesJSON `json:"-"`
}

// secondaryDNSMessagesJSON contains the JSON metadata for the struct
// [SecondaryDNSMessages]
type secondaryDNSMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r secondaryDNSMessagesJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDnsaclListResponse struct {
	Result []ACL                                  `json:"result"`
	JSON   accountSecondaryDnsaclListResponseJSON `json:"-"`
	ResponseCollectionDNSACLs
}

// accountSecondaryDnsaclListResponseJSON contains the JSON metadata for the struct
// [AccountSecondaryDnsaclListResponse]
type accountSecondaryDnsaclListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSecondaryDnsaclListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSecondaryDnsaclListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSecondaryDNSACLNewParams struct {
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r AccountSecondaryDNSACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSecondaryDNSACLUpdateParams struct {
	ACL ACLParam `json:"acl,required"`
}

func (r AccountSecondaryDNSACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ACL)
}

type AccountSecondaryDNSACLDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountSecondaryDNSACLDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
