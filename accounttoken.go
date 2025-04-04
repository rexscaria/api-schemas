// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountTokenService) Delete(ctx context.Context, accountID string, tokenID string, body AccountTokenDeleteParams, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Find all available permission groups for Account Owned API Tokens
func (r *AccountTokenService) ListPermissionGroups(ctx context.Context, accountID string, opts ...option.RequestOption) (res *IamPermissionsGroupResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/permission_groups", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *AccountTokenService) Verify(ctx context.Context, accountID string, opts ...option.RequestOption) (res *IamResponseSingleSegment, err error) {
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
	Result []IamTokenBase                  `json:"result"`
	JSON   iamCollectionTokensResponseJSON `json:"-"`
	IamAPIResponseCollection
}

// iamCollectionTokensResponseJSON contains the JSON metadata for the struct
// [IamCollectionTokensResponse]
type iamCollectionTokensResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCollectionTokensResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCollectionTokensResponseJSON) RawJSON() string {
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
	Result []IamPermissionsGroupResponseCollectionResult `json:"result"`
	JSON   iamPermissionsGroupResponseCollectionJSON     `json:"-"`
	IamAPIResponseCollection
}

// iamPermissionsGroupResponseCollectionJSON contains the JSON metadata for the
// struct [IamPermissionsGroupResponseCollection]
type iamPermissionsGroupResponseCollectionJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionsGroupResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionsGroupResponseCollectionJSON) RawJSON() string {
	return r.raw
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

type IamPolicyWithPermissionGroupsAndResources struct {
	// Policy identifier.
	ID string `json:"id,required"`
	// Allow or deny operations against the resources.
	Effect IamPolicyWithPermissionGroupsAndResourcesEffect `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []IamPermissionGroup `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources map[string]string                             `json:"resources,required"`
	JSON      iamPolicyWithPermissionGroupsAndResourcesJSON `json:"-"`
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

type IamPolicyWithPermissionGroupsAndResourcesParam struct {
	// Allow or deny operations against the resources.
	Effect param.Field[IamPolicyWithPermissionGroupsAndResourcesEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]IamPermissionGroupParam] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[map[string]string] `json:"resources,required"`
}

func (r IamPolicyWithPermissionGroupsAndResourcesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamResponseSingleSegment struct {
	Result IamResponseSingleSegmentResult `json:"result"`
	JSON   iamResponseSingleSegmentJSON   `json:"-"`
	APIResponseSingleIam
}

// iamResponseSingleSegmentJSON contains the JSON metadata for the struct
// [IamResponseSingleSegment]
type iamResponseSingleSegmentJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResponseSingleSegment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleSegmentJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleSegmentResult struct {
	// Token identifier tag.
	ID string `json:"id,required"`
	// Status of the token.
	Status IamStatus `json:"status,required"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time                          `json:"not_before" format:"date-time"`
	JSON      iamResponseSingleSegmentResultJSON `json:"-"`
}

// iamResponseSingleSegmentResultJSON contains the JSON metadata for the struct
// [IamResponseSingleSegmentResult]
type iamResponseSingleSegmentResultJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	ExpiresOn   apijson.Field
	NotBefore   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResponseSingleSegmentResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResponseSingleSegmentResultJSON) RawJSON() string {
	return r.raw
}

type IamResponseSingleValue struct {
	// The token value.
	Result string                     `json:"result"`
	JSON   iamResponseSingleValueJSON `json:"-"`
	APIResponseSingleIam
}

// iamResponseSingleValueJSON contains the JSON metadata for the struct
// [IamResponseSingleValue]
type iamResponseSingleValueJSON struct {
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

type IamSingleTokenCreateResponse struct {
	Result IamSingleTokenCreateResponseResult `json:"result"`
	JSON   iamSingleTokenCreateResponseJSON   `json:"-"`
	APIResponseSingleIam
}

// iamSingleTokenCreateResponseJSON contains the JSON metadata for the struct
// [IamSingleTokenCreateResponse]
type iamSingleTokenCreateResponseJSON struct {
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

type IamSingleTokenCreateResponseResult struct {
	// The token value.
	Value string                                 `json:"value"`
	JSON  iamSingleTokenCreateResponseResultJSON `json:"-"`
	IamTokenBase
}

// iamSingleTokenCreateResponseResultJSON contains the JSON metadata for the struct
// [IamSingleTokenCreateResponseResult]
type iamSingleTokenCreateResponseResultJSON struct {
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

type IamSingleTokenResponse struct {
	Result IamTokenBase               `json:"result"`
	JSON   iamSingleTokenResponseJSON `json:"-"`
	APIResponseSingleIam
}

// iamSingleTokenResponseJSON contains the JSON metadata for the struct
// [IamSingleTokenResponse]
type iamSingleTokenResponseJSON struct {
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

// Status of the token.
type IamStatus string

const (
	IamStatusActive   IamStatus = "active"
	IamStatusDisabled IamStatus = "disabled"
	IamStatusExpired  IamStatus = "expired"
)

func (r IamStatus) IsKnown() bool {
	switch r {
	case IamStatusActive, IamStatusDisabled, IamStatusExpired:
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
	Status IamStatus        `json:"status"`
	JSON   iamTokenBaseJSON `json:"-"`
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

type IamTokenBaseParam struct {
	Condition param.Field[IamConditionParam] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// Token name.
	Name param.Field[string] `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies param.Field[[]IamPolicyWithPermissionGroupsAndResourcesParam] `json:"policies"`
	// Status of the token.
	Status param.Field[IamStatus] `json:"status"`
}

func (r IamTokenBaseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamTokenBodyParam struct {
	IamTokenBaseParam
}

func (r IamTokenBodyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
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

type AccountTokenDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountTokenDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountTokenRollParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountTokenRollParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
