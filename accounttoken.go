// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// AccountTokenService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountTokenService] method instead.
type AccountTokenService struct {
	Options []option.RequestOption
}

// NewAccountTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountTokenService(opts ...option.RequestOption) (r *AccountTokenService) {
	r = &AccountTokenService{}
	r.Options = opts
	return
}

// Create a new Account Owned API token.
func (r *AccountTokenService) New(ctx context.Context, accountID string, body AccountTokenNewParams, opts ...option.RequestOption) (res *IamSingleTokenCreateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a specific Account Owned API token.
func (r *AccountTokenService) Get(ctx context.Context, accountID string, tokenID string, opts ...option.RequestOption) (res *IamSingleTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing token.
func (r *AccountTokenService) Update(ctx context.Context, accountID string, tokenID string, body AccountTokenUpdateParams, opts ...option.RequestOption) (res *IamSingleTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all Account Owned API tokens created for this account.
func (r *AccountTokenService) List(ctx context.Context, accountID string, query AccountTokenListParams, opts ...option.RequestOption) (res *IamCollectionTokensResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Destroy an Account Owned API token.
func (r *AccountTokenService) Delete(ctx context.Context, accountID string, tokenID string, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Find all available permission groups for Account Owned API Tokens
func (r *AccountTokenService) ListPermissionGroups(ctx context.Context, accountID string, query AccountTokenListPermissionGroupsParams, opts ...option.RequestOption) (res *IamPermissionsGroupResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/permission_groups", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Roll the Account Owned API token secret.
func (r *AccountTokenService) Roll(ctx context.Context, accountID string, tokenID string, body AccountTokenRollParams, opts ...option.RequestOption) (res *IamResponseSingleValue, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/%s/value", accountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Test whether a token works.
func (r *AccountTokenService) Verify(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountTokenVerifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/verify", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type IamCollectionTokensResponse struct {
	Errors   []IamCollectionTokensResponseError   `json:"errors,required"`
	Messages []IamCollectionTokensResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    IamCollectionTokensResponseSuccess    `json:"success,required"`
	Result     []IamTokenBase                        `json:"result"`
	ResultInfo IamCollectionTokensResponseResultInfo `json:"result_info"`
	JSON       iamCollectionTokensResponseJSON       `json:"-"`
}

// iamCollectionTokensResponseJSON contains the JSON metadata for the struct
// [IamCollectionTokensResponse]
type iamCollectionTokensResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCollectionTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseJSON) RawJSON() string {
	return r.raw
}

type IamCollectionTokensResponseError struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           IamCollectionTokensResponseErrorsSource `json:"source"`
	JSON             iamCollectionTokensResponseErrorJSON    `json:"-"`
}

// iamCollectionTokensResponseErrorJSON contains the JSON metadata for the struct
// [IamCollectionTokensResponseError]
type iamCollectionTokensResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamCollectionTokensResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseErrorJSON) RawJSON() string {
	return r.raw
}

type IamCollectionTokensResponseErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    iamCollectionTokensResponseErrorsSourceJSON `json:"-"`
}

// iamCollectionTokensResponseErrorsSourceJSON contains the JSON metadata for the
// struct [IamCollectionTokensResponseErrorsSource]
type iamCollectionTokensResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCollectionTokensResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamCollectionTokensResponseMessage struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           IamCollectionTokensResponseMessagesSource `json:"source"`
	JSON             iamCollectionTokensResponseMessageJSON    `json:"-"`
}

// iamCollectionTokensResponseMessageJSON contains the JSON metadata for the struct
// [IamCollectionTokensResponseMessage]
type iamCollectionTokensResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamCollectionTokensResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseMessageJSON) RawJSON() string {
	return r.raw
}

type IamCollectionTokensResponseMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    iamCollectionTokensResponseMessagesSourceJSON `json:"-"`
}

// iamCollectionTokensResponseMessagesSourceJSON contains the JSON metadata for the
// struct [IamCollectionTokensResponseMessagesSource]
type iamCollectionTokensResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCollectionTokensResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamCollectionTokensResponseSuccess bool

const (
	IamCollectionTokensResponseSuccessTrue IamCollectionTokensResponseSuccess = true
)

func (r IamCollectionTokensResponseSuccess) IsKnown() bool {
	switch r {
	case IamCollectionTokensResponseSuccessTrue:
		return true
	}
	return false
}

type IamCollectionTokensResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       iamCollectionTokensResponseResultInfoJSON `json:"-"`
}

// iamCollectionTokensResponseResultInfoJSON contains the JSON metadata for the
// struct [IamCollectionTokensResponseResultInfo]
type iamCollectionTokensResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCollectionTokensResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type IamCondition struct {
	// Client IP restrictions.
	RequestIP IamConditionRequestIP `json:"request_ip"`
	JSON      iamConditionJSON      `json:"-"`
}

// iamConditionJSON contains the JSON metadata for the struct [IamCondition]
type iamConditionJSON struct {
	RequestIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamConditionJSON) RawJSON() string {
	return r.raw
}

// Client IP restrictions.
type IamConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In []string `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn []string                  `json:"not_in"`
	JSON  iamConditionRequestIPJSON `json:"-"`
}

// iamConditionRequestIPJSON contains the JSON metadata for the struct
// [IamConditionRequestIP]
type iamConditionRequestIPJSON struct {
	In          apijson.Field
	NotIn       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamConditionRequestIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamConditionRequestIPJSON) RawJSON() string {
	return r.raw
}

type IamConditionParam struct {
	// Client IP restrictions.
	RequestIP param.Field[IamConditionRequestIPParam] `json:"request_ip"`
}

func (r IamConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type IamConditionRequestIPParam struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]string] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]string] `json:"not_in"`
}

func (r IamConditionRequestIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamCreatePayloadParam struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]IamPolicyWithPermissionGroupsAndResourcesParam] `json:"policies,required"`
	Condition param.Field[IamConditionParam]                                `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
}

func (r IamCreatePayloadParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamPermissionsGroupResponseCollection struct {
	Errors   []IamPermissionsGroupResponseCollectionError   `json:"errors,required"`
	Messages []IamPermissionsGroupResponseCollectionMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    IamPermissionsGroupResponseCollectionSuccess    `json:"success,required"`
	Result     []IamPermissionsGroupResponseCollectionResult   `json:"result"`
	ResultInfo IamPermissionsGroupResponseCollectionResultInfo `json:"result_info"`
	JSON       iamPermissionsGroupResponseCollectionJSON       `json:"-"`
}

// iamPermissionsGroupResponseCollectionJSON contains the JSON metadata for the
// struct [IamPermissionsGroupResponseCollection]
type iamPermissionsGroupResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsGroupResponseCollectionError struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           IamPermissionsGroupResponseCollectionErrorsSource `json:"source"`
	JSON             iamPermissionsGroupResponseCollectionErrorJSON    `json:"-"`
}

// iamPermissionsGroupResponseCollectionErrorJSON contains the JSON metadata for
// the struct [IamPermissionsGroupResponseCollectionError]
type iamPermissionsGroupResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsGroupResponseCollectionErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    iamPermissionsGroupResponseCollectionErrorsSourceJSON `json:"-"`
}

// iamPermissionsGroupResponseCollectionErrorsSourceJSON contains the JSON metadata
// for the struct [IamPermissionsGroupResponseCollectionErrorsSource]
type iamPermissionsGroupResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsGroupResponseCollectionMessage struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           IamPermissionsGroupResponseCollectionMessagesSource `json:"source"`
	JSON             iamPermissionsGroupResponseCollectionMessageJSON    `json:"-"`
}

// iamPermissionsGroupResponseCollectionMessageJSON contains the JSON metadata for
// the struct [IamPermissionsGroupResponseCollectionMessage]
type iamPermissionsGroupResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsGroupResponseCollectionMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    iamPermissionsGroupResponseCollectionMessagesSourceJSON `json:"-"`
}

// iamPermissionsGroupResponseCollectionMessagesSourceJSON contains the JSON
// metadata for the struct [IamPermissionsGroupResponseCollectionMessagesSource]
type iamPermissionsGroupResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamPermissionsGroupResponseCollectionSuccess bool

const (
	IamPermissionsGroupResponseCollectionSuccessTrue IamPermissionsGroupResponseCollectionSuccess = true
)

func (r IamPermissionsGroupResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case IamPermissionsGroupResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type IamPermissionsGroupResponseCollectionResult struct {
	// Public ID.
	ID string `json:"id"`
	// Permission Group Name
	Name string `json:"name"`
	// Resources to which the Permission Group is scoped
	Scopes []IamPermissionsGroupResponseCollectionResultScope `json:"scopes"`
	JSON   iamPermissionsGroupResponseCollectionResultJSON    `json:"-"`
}

// iamPermissionsGroupResponseCollectionResultJSON contains the JSON metadata for
// the struct [IamPermissionsGroupResponseCollectionResult]
type iamPermissionsGroupResponseCollectionResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionResultJSON) RawJSON() string {
	return r.raw
}

type IamPermissionsGroupResponseCollectionResultScope string

const (
	IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIAccount     IamPermissionsGroupResponseCollectionResultScope = "com.cloudflare.api.account"
	IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIAccountZone IamPermissionsGroupResponseCollectionResultScope = "com.cloudflare.api.account.zone"
	IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIUser        IamPermissionsGroupResponseCollectionResultScope = "com.cloudflare.api.user"
	IamPermissionsGroupResponseCollectionResultScopeComCloudflareEdgeR2Bucket   IamPermissionsGroupResponseCollectionResultScope = "com.cloudflare.edge.r2.bucket"
)

func (r IamPermissionsGroupResponseCollectionResultScope) IsKnown() bool {
	switch r {
	case IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIAccount, IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIAccountZone, IamPermissionsGroupResponseCollectionResultScopeComCloudflareAPIUser, IamPermissionsGroupResponseCollectionResultScopeComCloudflareEdgeR2Bucket:
		return true
	}
	return false
}

type IamPermissionsGroupResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       iamPermissionsGroupResponseCollectionResultInfoJSON `json:"-"`
}

// iamPermissionsGroupResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [IamPermissionsGroupResponseCollectionResultInfo]
type iamPermissionsGroupResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type IamPolicyWithPermissionGroupsAndResources struct {
	// Policy identifier.
	ID string `json:"id,required"`
	// Allow or deny operations against the resources.
	Effect IamPolicyWithPermissionGroupsAndResourcesEffect `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []IamPermissionGroup `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources IamPolicyWithPermissionGroupsAndResourcesResourcesUnion `json:"resources,required"`
	JSON      iamPolicyWithPermissionGroupsAndResourcesJSON           `json:"-"`
}

// iamPolicyWithPermissionGroupsAndResourcesJSON contains the JSON metadata for the
// struct [IamPolicyWithPermissionGroupsAndResources]
type iamPolicyWithPermissionGroupsAndResourcesJSON struct {
	ID               apijson.Field
	Effect           apijson.Field
	PermissionGroups apijson.Field
	Resources        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamPolicyWithPermissionGroupsAndResources) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPolicyWithPermissionGroupsAndResourcesJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type IamPolicyWithPermissionGroupsAndResourcesEffect string

const (
	IamPolicyWithPermissionGroupsAndResourcesEffectAllow IamPolicyWithPermissionGroupsAndResourcesEffect = "allow"
	IamPolicyWithPermissionGroupsAndResourcesEffectDeny  IamPolicyWithPermissionGroupsAndResourcesEffect = "deny"
)

func (r IamPolicyWithPermissionGroupsAndResourcesEffect) IsKnown() bool {
	switch r {
	case IamPolicyWithPermissionGroupsAndResourcesEffectAllow, IamPolicyWithPermissionGroupsAndResourcesEffectDeny:
		return true
	}
	return false
}

// A list of resource names that the policy applies to.
//
// Union satisfied by
// [IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectString]
// or
// [IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNested].
type IamPolicyWithPermissionGroupsAndResourcesResourcesUnion interface {
	implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IamPolicyWithPermissionGroupsAndResourcesResourcesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectString{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNested{}),
		},
	)
}

type IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectString map[string]string

func (r IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectString) implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnion() {
}

type IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNested map[string]map[string]string

func (r IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNested) implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnion() {
}

type IamPolicyWithPermissionGroupsAndResourcesParam struct {
	// Allow or deny operations against the resources.
	Effect param.Field[IamPolicyWithPermissionGroupsAndResourcesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]IamPermissionGroupParam] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[IamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam] `json:"resources,required"`
}

func (r IamPolicyWithPermissionGroupsAndResourcesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of resource names that the policy applies to.
//
// Satisfied by
// [IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectStringParam],
// [IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNestedParam].
type IamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam interface {
	implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam()
}

type IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectStringParam map[string]string

func (r IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectStringParam) implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam() {
}

type IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNestedParam map[string]map[string]string

func (r IamPolicyWithPermissionGroupsAndResourcesResourcesIamResourcesTypeObjectNestedParam) implementsIamPolicyWithPermissionGroupsAndResourcesResourcesUnionParam() {
}

type IamResponseSingleValue struct {
	Errors   []IamResponseSingleValueError   `json:"errors,required"`
	Messages []IamResponseSingleValueMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IamResponseSingleValueSuccess `json:"success,required"`
	// The token value.
	Result string                     `json:"result"`
	JSON   iamResponseSingleValueJSON `json:"-"`
}

// iamResponseSingleValueJSON contains the JSON metadata for the struct
// [IamResponseSingleValue]
type iamResponseSingleValueJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResponseSingleValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleValueJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleValueError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           IamResponseSingleValueErrorsSource `json:"source"`
	JSON             iamResponseSingleValueErrorJSON    `json:"-"`
}

// iamResponseSingleValueErrorJSON contains the JSON metadata for the struct
// [IamResponseSingleValueError]
type iamResponseSingleValueErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamResponseSingleValueError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleValueErrorJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleValueErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    iamResponseSingleValueErrorsSourceJSON `json:"-"`
}

// iamResponseSingleValueErrorsSourceJSON contains the JSON metadata for the struct
// [IamResponseSingleValueErrorsSource]
type iamResponseSingleValueErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResponseSingleValueErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleValueErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleValueMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           IamResponseSingleValueMessagesSource `json:"source"`
	JSON             iamResponseSingleValueMessageJSON    `json:"-"`
}

// iamResponseSingleValueMessageJSON contains the JSON metadata for the struct
// [IamResponseSingleValueMessage]
type iamResponseSingleValueMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamResponseSingleValueMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleValueMessageJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleValueMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    iamResponseSingleValueMessagesSourceJSON `json:"-"`
}

// iamResponseSingleValueMessagesSourceJSON contains the JSON metadata for the
// struct [IamResponseSingleValueMessagesSource]
type iamResponseSingleValueMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResponseSingleValueMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleValueMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamResponseSingleValueSuccess bool

const (
	IamResponseSingleValueSuccessTrue IamResponseSingleValueSuccess = true
)

func (r IamResponseSingleValueSuccess) IsKnown() bool {
	switch r {
	case IamResponseSingleValueSuccessTrue:
		return true
	}
	return false
}

type IamSingleTokenCreateResponse struct {
	Errors   []IamSingleTokenCreateResponseError   `json:"errors,required"`
	Messages []IamSingleTokenCreateResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IamSingleTokenCreateResponseSuccess `json:"success,required"`
	Result  IamSingleTokenCreateResponseResult  `json:"result"`
	JSON    iamSingleTokenCreateResponseJSON    `json:"-"`
}

// iamSingleTokenCreateResponseJSON contains the JSON metadata for the struct
// [IamSingleTokenCreateResponse]
type iamSingleTokenCreateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenCreateResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           IamSingleTokenCreateResponseErrorsSource `json:"source"`
	JSON             iamSingleTokenCreateResponseErrorJSON    `json:"-"`
}

// iamSingleTokenCreateResponseErrorJSON contains the JSON metadata for the struct
// [IamSingleTokenCreateResponseError]
type iamSingleTokenCreateResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseErrorJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenCreateResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    iamSingleTokenCreateResponseErrorsSourceJSON `json:"-"`
}

// iamSingleTokenCreateResponseErrorsSourceJSON contains the JSON metadata for the
// struct [IamSingleTokenCreateResponseErrorsSource]
type iamSingleTokenCreateResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenCreateResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           IamSingleTokenCreateResponseMessagesSource `json:"source"`
	JSON             iamSingleTokenCreateResponseMessageJSON    `json:"-"`
}

// iamSingleTokenCreateResponseMessageJSON contains the JSON metadata for the
// struct [IamSingleTokenCreateResponseMessage]
type iamSingleTokenCreateResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseMessageJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenCreateResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    iamSingleTokenCreateResponseMessagesSourceJSON `json:"-"`
}

// iamSingleTokenCreateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [IamSingleTokenCreateResponseMessagesSource]
type iamSingleTokenCreateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamSingleTokenCreateResponseSuccess bool

const (
	IamSingleTokenCreateResponseSuccessTrue IamSingleTokenCreateResponseSuccess = true
)

func (r IamSingleTokenCreateResponseSuccess) IsKnown() bool {
	switch r {
	case IamSingleTokenCreateResponseSuccessTrue:
		return true
	}
	return false
}

type IamSingleTokenCreateResponseResult struct {
	// Token identifier tag.
	ID        string       `json:"id"`
	Condition IamCondition `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time on which the token was created.
	IssuedOn time.Time `json:"issued_on" format:"date-time"`
	// Last time the token was used.
	LastUsedOn time.Time `json:"last_used_on" format:"date-time"`
	// Last time the token was modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Token name.
	Name string `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies []IamPolicyWithPermissionGroupsAndResources `json:"policies"`
	// Status of the token.
	Status IamSingleTokenCreateResponseResultStatus `json:"status"`
	// The token value.
	Value string                                 `json:"value"`
	JSON  iamSingleTokenCreateResponseResultJSON `json:"-"`
}

// iamSingleTokenCreateResponseResultJSON contains the JSON metadata for the struct
// [IamSingleTokenCreateResponseResult]
type iamSingleTokenCreateResponseResultJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	ExpiresOn   apijson.Field
	IssuedOn    apijson.Field
	LastUsedOn  apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	NotBefore   apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenCreateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenCreateResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type IamSingleTokenCreateResponseResultStatus string

const (
	IamSingleTokenCreateResponseResultStatusActive   IamSingleTokenCreateResponseResultStatus = "active"
	IamSingleTokenCreateResponseResultStatusDisabled IamSingleTokenCreateResponseResultStatus = "disabled"
	IamSingleTokenCreateResponseResultStatusExpired  IamSingleTokenCreateResponseResultStatus = "expired"
)

func (r IamSingleTokenCreateResponseResultStatus) IsKnown() bool {
	switch r {
	case IamSingleTokenCreateResponseResultStatusActive, IamSingleTokenCreateResponseResultStatusDisabled, IamSingleTokenCreateResponseResultStatusExpired:
		return true
	}
	return false
}

type IamSingleTokenResponse struct {
	Errors   []IamSingleTokenResponseError   `json:"errors,required"`
	Messages []IamSingleTokenResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IamSingleTokenResponseSuccess `json:"success,required"`
	Result  IamTokenBase                  `json:"result"`
	JSON    iamSingleTokenResponseJSON    `json:"-"`
}

// iamSingleTokenResponseJSON contains the JSON metadata for the struct
// [IamSingleTokenResponse]
type iamSingleTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenResponseJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           IamSingleTokenResponseErrorsSource `json:"source"`
	JSON             iamSingleTokenResponseErrorJSON    `json:"-"`
}

// iamSingleTokenResponseErrorJSON contains the JSON metadata for the struct
// [IamSingleTokenResponseError]
type iamSingleTokenResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenResponseErrorJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    iamSingleTokenResponseErrorsSourceJSON `json:"-"`
}

// iamSingleTokenResponseErrorsSourceJSON contains the JSON metadata for the struct
// [IamSingleTokenResponseErrorsSource]
type iamSingleTokenResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           IamSingleTokenResponseMessagesSource `json:"source"`
	JSON             iamSingleTokenResponseMessageJSON    `json:"-"`
}

// iamSingleTokenResponseMessageJSON contains the JSON metadata for the struct
// [IamSingleTokenResponseMessage]
type iamSingleTokenResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenResponseMessageJSON) RawJSON() string {
	return r.raw
}

type IamSingleTokenResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    iamSingleTokenResponseMessagesSourceJSON `json:"-"`
}

// iamSingleTokenResponseMessagesSourceJSON contains the JSON metadata for the
// struct [IamSingleTokenResponseMessagesSource]
type iamSingleTokenResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleTokenResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleTokenResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamSingleTokenResponseSuccess bool

const (
	IamSingleTokenResponseSuccessTrue IamSingleTokenResponseSuccess = true
)

func (r IamSingleTokenResponseSuccess) IsKnown() bool {
	switch r {
	case IamSingleTokenResponseSuccessTrue:
		return true
	}
	return false
}

type IamTokenBase struct {
	// Token identifier tag.
	ID        string       `json:"id"`
	Condition IamCondition `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time on which the token was created.
	IssuedOn time.Time `json:"issued_on" format:"date-time"`
	// Last time the token was used.
	LastUsedOn time.Time `json:"last_used_on" format:"date-time"`
	// Last time the token was modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Token name.
	Name string `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies []IamPolicyWithPermissionGroupsAndResources `json:"policies"`
	// Status of the token.
	Status IamTokenBaseStatus `json:"status"`
	JSON   iamTokenBaseJSON   `json:"-"`
}

// iamTokenBaseJSON contains the JSON metadata for the struct [IamTokenBase]
type iamTokenBaseJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	ExpiresOn   apijson.Field
	IssuedOn    apijson.Field
	LastUsedOn  apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	NotBefore   apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamTokenBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamTokenBaseJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type IamTokenBaseStatus string

const (
	IamTokenBaseStatusActive   IamTokenBaseStatus = "active"
	IamTokenBaseStatusDisabled IamTokenBaseStatus = "disabled"
	IamTokenBaseStatusExpired  IamTokenBaseStatus = "expired"
)

func (r IamTokenBaseStatus) IsKnown() bool {
	switch r {
	case IamTokenBaseStatusActive, IamTokenBaseStatusDisabled, IamTokenBaseStatusExpired:
		return true
	}
	return false
}

type IamTokenBodyParam struct {
	// Token name.
	Name param.Field[string] `json:"name,required"`
	// List of access policies assigned to the token.
	Policies  param.Field[[]IamPolicyWithPermissionGroupsAndResourcesParam] `json:"policies,required"`
	Condition param.Field[IamConditionParam]                                `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
	// Status of the token.
	Status param.Field[IamTokenBodyStatus] `json:"status"`
}

func (r IamTokenBodyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of the token.
type IamTokenBodyStatus string

const (
	IamTokenBodyStatusActive   IamTokenBodyStatus = "active"
	IamTokenBodyStatusDisabled IamTokenBodyStatus = "disabled"
	IamTokenBodyStatusExpired  IamTokenBodyStatus = "expired"
)

func (r IamTokenBodyStatus) IsKnown() bool {
	switch r {
	case IamTokenBodyStatusActive, IamTokenBodyStatusDisabled, IamTokenBodyStatusExpired:
		return true
	}
	return false
}

type AccountTokenVerifyResponse struct {
	Errors   []AccountTokenVerifyResponseError   `json:"errors,required"`
	Messages []AccountTokenVerifyResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountTokenVerifyResponseSuccess `json:"success,required"`
	Result  AccountTokenVerifyResponseResult  `json:"result"`
	JSON    accountTokenVerifyResponseJSON    `json:"-"`
}

// accountTokenVerifyResponseJSON contains the JSON metadata for the struct
// [AccountTokenVerifyResponse]
type accountTokenVerifyResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTokenVerifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseJSON) RawJSON() string {
	return r.raw
}

type AccountTokenVerifyResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AccountTokenVerifyResponseErrorsSource `json:"source"`
	JSON             accountTokenVerifyResponseErrorJSON    `json:"-"`
}

// accountTokenVerifyResponseErrorJSON contains the JSON metadata for the struct
// [AccountTokenVerifyResponseError]
type accountTokenVerifyResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTokenVerifyResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountTokenVerifyResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    accountTokenVerifyResponseErrorsSourceJSON `json:"-"`
}

// accountTokenVerifyResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountTokenVerifyResponseErrorsSource]
type accountTokenVerifyResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTokenVerifyResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountTokenVerifyResponseMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountTokenVerifyResponseMessagesSource `json:"source"`
	JSON             accountTokenVerifyResponseMessageJSON    `json:"-"`
}

// accountTokenVerifyResponseMessageJSON contains the JSON metadata for the struct
// [AccountTokenVerifyResponseMessage]
type accountTokenVerifyResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountTokenVerifyResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountTokenVerifyResponseMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountTokenVerifyResponseMessagesSourceJSON `json:"-"`
}

// accountTokenVerifyResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountTokenVerifyResponseMessagesSource]
type accountTokenVerifyResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTokenVerifyResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountTokenVerifyResponseSuccess bool

const (
	AccountTokenVerifyResponseSuccessTrue AccountTokenVerifyResponseSuccess = true
)

func (r AccountTokenVerifyResponseSuccess) IsKnown() bool {
	switch r {
	case AccountTokenVerifyResponseSuccessTrue:
		return true
	}
	return false
}

type AccountTokenVerifyResponseResult struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status AccountTokenVerifyResponseResultStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                            `json:"not_before" format:"date-time"`
	JSON      accountTokenVerifyResponseResultJSON `json:"-"`
}

// accountTokenVerifyResponseResultJSON contains the JSON metadata for the struct
// [AccountTokenVerifyResponseResult]
type accountTokenVerifyResponseResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTokenVerifyResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountTokenVerifyResponseResultJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type AccountTokenVerifyResponseResultStatus string

const (
	AccountTokenVerifyResponseResultStatusActive   AccountTokenVerifyResponseResultStatus = "active"
	AccountTokenVerifyResponseResultStatusDisabled AccountTokenVerifyResponseResultStatus = "disabled"
	AccountTokenVerifyResponseResultStatusExpired  AccountTokenVerifyResponseResultStatus = "expired"
)

func (r AccountTokenVerifyResponseResultStatus) IsKnown() bool {
	switch r {
	case AccountTokenVerifyResponseResultStatusActive, AccountTokenVerifyResponseResultStatusDisabled, AccountTokenVerifyResponseResultStatusExpired:
		return true
	}
	return false
}

type AccountTokenNewParams struct {
	IamCreatePayload IamCreatePayloadParam `json:"iam_create_payload,required"`
}

func (r AccountTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IamCreatePayload)
}

type AccountTokenUpdateParams struct {
	IamTokenBody IamTokenBodyParam `json:"iam_token_body,required"`
}

func (r AccountTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IamTokenBody)
}

type AccountTokenListParams struct {
	// Direction to order results.
	Direction param.Field[AccountTokenListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountTokenListParams]'s query parameters as `url.Values`.
func (r AccountTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountTokenListParamsDirection string

const (
	AccountTokenListParamsDirectionAsc  AccountTokenListParamsDirection = "asc"
	AccountTokenListParamsDirectionDesc AccountTokenListParamsDirection = "desc"
)

func (r AccountTokenListParamsDirection) IsKnown() bool {
	switch r {
	case AccountTokenListParamsDirectionAsc, AccountTokenListParamsDirectionDesc:
		return true
	}
	return false
}

type AccountTokenListPermissionGroupsParams struct {
	// Filter by the name of the permission group. The value must be URL-encoded.
	Name param.Field[string] `query:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	Scope param.Field[string] `query:"scope"`
}

// URLQuery serializes [AccountTokenListPermissionGroupsParams]'s query parameters
// as `url.Values`.
func (r AccountTokenListPermissionGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountTokenRollParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountTokenRollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
