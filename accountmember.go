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

// AccountMemberService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMemberService] method instead.
type AccountMemberService struct {
	Options []option.RequestOption
}

// NewAccountMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMemberService(opts ...option.RequestOption) (r *AccountMemberService) {
	r = &AccountMemberService{}
	r.Options = opts
	return
}

// Get information about a specific member of an account.
func (r *AccountMemberService) Get(ctx context.Context, accountID string, memberID string, opts ...option.RequestOption) (res *IamSingleMemberResponseWithPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify an account member.
func (r *AccountMemberService) Update(ctx context.Context, accountID string, memberID string, body AccountMemberUpdateParams, opts ...option.RequestOption) (res *IamSingleMemberResponseWithPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all members of an account.
func (r *AccountMemberService) List(ctx context.Context, accountID string, query AccountMemberListParams, opts ...option.RequestOption) (res *AccountMemberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Add a user to the list of members for this account.
func (r *AccountMemberService) Add(ctx context.Context, accountID string, body AccountMemberAddParams, opts ...option.RequestOption) (res *IamSingleMemberResponseWithPolicies, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove a member from an account.
func (r *AccountMemberService) Remove(ctx context.Context, accountID string, memberID string, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", accountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Allow or deny operations against the resources.
type IamAccess string

const (
	IamAccessAllow IamAccess = "allow"
	IamAccessDeny  IamAccess = "deny"
)

func (r IamAccess) IsKnown() bool {
	switch r {
	case IamAccessAllow, IamAccessDeny:
		return true
	}
	return false
}

type IamCreateMemberPolicyParam struct {
	// Allow or deny operations against the resources.
	Access param.Field[IamAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]IamCreateMemberPolicyPermissionGroupParam] `json:"permission_groups,required"`
	// A list of resource groups that the policy applies to.
	ResourceGroups param.Field[[]IamCreateMemberPolicyResourceGroupParam] `json:"resource_groups,required"`
}

func (r IamCreateMemberPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of permissions.
type IamCreateMemberPolicyPermissionGroupParam struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r IamCreateMemberPolicyPermissionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type IamCreateMemberPolicyResourceGroupParam struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r IamCreateMemberPolicyResourceGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IamListMemberPolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access IamAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []IamPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []IamResourceGroup      `json:"resource_groups"`
	JSON           iamListMemberPolicyJSON `json:"-"`
}

// iamListMemberPolicyJSON contains the JSON metadata for the struct
// [IamListMemberPolicy]
type iamListMemberPolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamListMemberPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamListMemberPolicyJSON) RawJSON() string {
	return r.raw
}

type IamMemberWithPolicies struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []IamListMemberPolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []IamRole `json:"roles"`
	// A member's status in the account.
	Status IamMemberWithPoliciesStatus `json:"status"`
	// Details of the user associated to the membership.
	User IamMemberWithPoliciesUser `json:"user"`
	JSON iamMemberWithPoliciesJSON `json:"-"`
}

// iamMemberWithPoliciesJSON contains the JSON metadata for the struct
// [IamMemberWithPolicies]
type iamMemberWithPoliciesJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamMemberWithPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamMemberWithPoliciesJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type IamMemberWithPoliciesStatus string

const (
	IamMemberWithPoliciesStatusAccepted IamMemberWithPoliciesStatus = "accepted"
	IamMemberWithPoliciesStatusPending  IamMemberWithPoliciesStatus = "pending"
)

func (r IamMemberWithPoliciesStatus) IsKnown() bool {
	switch r {
	case IamMemberWithPoliciesStatusAccepted, IamMemberWithPoliciesStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type IamMemberWithPoliciesUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool                          `json:"two_factor_authentication_enabled"`
	JSON                           iamMemberWithPoliciesUserJSON `json:"-"`
}

// iamMemberWithPoliciesUserJSON contains the JSON metadata for the struct
// [IamMemberWithPoliciesUser]
type iamMemberWithPoliciesUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *IamMemberWithPoliciesUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamMemberWithPoliciesUserJSON) RawJSON() string {
	return r.raw
}

type IamSingleMemberResponseWithPolicies struct {
	Errors   []IamSingleMemberResponseWithPoliciesError   `json:"errors,required"`
	Messages []IamSingleMemberResponseWithPoliciesMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IamSingleMemberResponseWithPoliciesSuccess `json:"success,required"`
	Result  IamMemberWithPolicies                      `json:"result"`
	JSON    iamSingleMemberResponseWithPoliciesJSON    `json:"-"`
}

// iamSingleMemberResponseWithPoliciesJSON contains the JSON metadata for the
// struct [IamSingleMemberResponseWithPolicies]
type iamSingleMemberResponseWithPoliciesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleMemberResponseWithPolicies) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleMemberResponseWithPoliciesJSON) RawJSON() string {
	return r.raw
}

type IamSingleMemberResponseWithPoliciesError struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           IamSingleMemberResponseWithPoliciesErrorsSource `json:"source"`
	JSON             iamSingleMemberResponseWithPoliciesErrorJSON    `json:"-"`
}

// iamSingleMemberResponseWithPoliciesErrorJSON contains the JSON metadata for the
// struct [IamSingleMemberResponseWithPoliciesError]
type iamSingleMemberResponseWithPoliciesErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleMemberResponseWithPoliciesError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleMemberResponseWithPoliciesErrorJSON) RawJSON() string {
	return r.raw
}

type IamSingleMemberResponseWithPoliciesErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    iamSingleMemberResponseWithPoliciesErrorsSourceJSON `json:"-"`
}

// iamSingleMemberResponseWithPoliciesErrorsSourceJSON contains the JSON metadata
// for the struct [IamSingleMemberResponseWithPoliciesErrorsSource]
type iamSingleMemberResponseWithPoliciesErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleMemberResponseWithPoliciesErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleMemberResponseWithPoliciesErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamSingleMemberResponseWithPoliciesMessage struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           IamSingleMemberResponseWithPoliciesMessagesSource `json:"source"`
	JSON             iamSingleMemberResponseWithPoliciesMessageJSON    `json:"-"`
}

// iamSingleMemberResponseWithPoliciesMessageJSON contains the JSON metadata for
// the struct [IamSingleMemberResponseWithPoliciesMessage]
type iamSingleMemberResponseWithPoliciesMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleMemberResponseWithPoliciesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleMemberResponseWithPoliciesMessageJSON) RawJSON() string {
	return r.raw
}

type IamSingleMemberResponseWithPoliciesMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    iamSingleMemberResponseWithPoliciesMessagesSourceJSON `json:"-"`
}

// iamSingleMemberResponseWithPoliciesMessagesSourceJSON contains the JSON metadata
// for the struct [IamSingleMemberResponseWithPoliciesMessagesSource]
type iamSingleMemberResponseWithPoliciesMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleMemberResponseWithPoliciesMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleMemberResponseWithPoliciesMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamSingleMemberResponseWithPoliciesSuccess bool

const (
	IamSingleMemberResponseWithPoliciesSuccessTrue IamSingleMemberResponseWithPoliciesSuccess = true
)

func (r IamSingleMemberResponseWithPoliciesSuccess) IsKnown() bool {
	switch r {
	case IamSingleMemberResponseWithPoliciesSuccessTrue:
		return true
	}
	return false
}

type AccountMemberListResponse struct {
	Errors   []AccountMemberListResponseError   `json:"errors,required"`
	Messages []AccountMemberListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountMemberListResponseSuccess    `json:"success,required"`
	Result     []IamMemberWithPolicies             `json:"result"`
	ResultInfo AccountMemberListResponseResultInfo `json:"result_info"`
	JSON       accountMemberListResponseJSON       `json:"-"`
}

// accountMemberListResponseJSON contains the JSON metadata for the struct
// [AccountMemberListResponse]
type accountMemberListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMemberListResponseError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AccountMemberListResponseErrorsSource `json:"source"`
	JSON             accountMemberListResponseErrorJSON    `json:"-"`
}

// accountMemberListResponseErrorJSON contains the JSON metadata for the struct
// [AccountMemberListResponseError]
type accountMemberListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMemberListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountMemberListResponseErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    accountMemberListResponseErrorsSourceJSON `json:"-"`
}

// accountMemberListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountMemberListResponseErrorsSource]
type accountMemberListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountMemberListResponseMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AccountMemberListResponseMessagesSource `json:"source"`
	JSON             accountMemberListResponseMessageJSON    `json:"-"`
}

// accountMemberListResponseMessageJSON contains the JSON metadata for the struct
// [AccountMemberListResponseMessage]
type accountMemberListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountMemberListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountMemberListResponseMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    accountMemberListResponseMessagesSourceJSON `json:"-"`
}

// accountMemberListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountMemberListResponseMessagesSource]
type accountMemberListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountMemberListResponseSuccess bool

const (
	AccountMemberListResponseSuccessTrue AccountMemberListResponseSuccess = true
)

func (r AccountMemberListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountMemberListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountMemberListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       accountMemberListResponseResultInfoJSON `json:"-"`
}

// accountMemberListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountMemberListResponseResultInfo]
type accountMemberListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountMemberUpdateParams struct {
	Body AccountMemberUpdateParamsBodyUnion `json:"body,required"`
}

func (r AccountMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMemberUpdateParamsBody struct {
	Policies param.Field[interface{}] `json:"policies"`
	Roles    param.Field[interface{}] `json:"roles"`
	User     param.Field[interface{}] `json:"user"`
}

func (r AccountMemberUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberUpdateParamsBody) implementsAccountMemberUpdateParamsBodyUnion() {}

// Satisfied by [AccountMemberUpdateParamsBodyIamUpdateMemberWithRoles],
// [AccountMemberUpdateParamsBodyIamUpdateMemberWithPolicies],
// [AccountMemberUpdateParamsBody].
type AccountMemberUpdateParamsBodyUnion interface {
	implementsAccountMemberUpdateParamsBodyUnion()
}

type AccountMemberUpdateParamsBodyIamUpdateMemberWithRoles struct {
	// Roles assigned to this member.
	Roles param.Field[[]IamRoleParam] `json:"roles"`
}

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithRoles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithRoles) implementsAccountMemberUpdateParamsBodyUnion() {
}

// A member's status in the account.
type AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatus string

const (
	AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatusAccepted AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatus = "accepted"
	AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatusPending  AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatus = "pending"
)

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatus) IsKnown() bool {
	switch r {
	case AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatusAccepted, AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesUser struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Identifier
	ID param.Field[string] `json:"id"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
}

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithRolesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMemberUpdateParamsBodyIamUpdateMemberWithPolicies struct {
	// Array of policies associated with this member.
	Policies param.Field[[]IamCreateMemberPolicyParam] `json:"policies,required"`
}

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberUpdateParamsBodyIamUpdateMemberWithPolicies) implementsAccountMemberUpdateParamsBodyUnion() {
}

// A member's status in the account.
type AccountMemberUpdateParamsBodyStatus string

const (
	AccountMemberUpdateParamsBodyStatusAccepted AccountMemberUpdateParamsBodyStatus = "accepted"
	AccountMemberUpdateParamsBodyStatusPending  AccountMemberUpdateParamsBodyStatus = "pending"
)

func (r AccountMemberUpdateParamsBodyStatus) IsKnown() bool {
	switch r {
	case AccountMemberUpdateParamsBodyStatusAccepted, AccountMemberUpdateParamsBodyStatusPending:
		return true
	}
	return false
}

type AccountMemberListParams struct {
	// Direction to order results.
	Direction param.Field[AccountMemberListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[AccountMemberListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[AccountMemberListParamsStatus] `query:"status"`
}

// URLQuery serializes [AccountMemberListParams]'s query parameters as
// `url.Values`.
func (r AccountMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountMemberListParamsDirection string

const (
	AccountMemberListParamsDirectionAsc  AccountMemberListParamsDirection = "asc"
	AccountMemberListParamsDirectionDesc AccountMemberListParamsDirection = "desc"
)

func (r AccountMemberListParamsDirection) IsKnown() bool {
	switch r {
	case AccountMemberListParamsDirectionAsc, AccountMemberListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type AccountMemberListParamsOrder string

const (
	AccountMemberListParamsOrderUserFirstName AccountMemberListParamsOrder = "user.first_name"
	AccountMemberListParamsOrderUserLastName  AccountMemberListParamsOrder = "user.last_name"
	AccountMemberListParamsOrderUserEmail     AccountMemberListParamsOrder = "user.email"
	AccountMemberListParamsOrderStatus        AccountMemberListParamsOrder = "status"
)

func (r AccountMemberListParamsOrder) IsKnown() bool {
	switch r {
	case AccountMemberListParamsOrderUserFirstName, AccountMemberListParamsOrderUserLastName, AccountMemberListParamsOrderUserEmail, AccountMemberListParamsOrderStatus:
		return true
	}
	return false
}

// A member's status in the account.
type AccountMemberListParamsStatus string

const (
	AccountMemberListParamsStatusAccepted AccountMemberListParamsStatus = "accepted"
	AccountMemberListParamsStatusPending  AccountMemberListParamsStatus = "pending"
	AccountMemberListParamsStatusRejected AccountMemberListParamsStatus = "rejected"
)

func (r AccountMemberListParamsStatus) IsKnown() bool {
	switch r {
	case AccountMemberListParamsStatusAccepted, AccountMemberListParamsStatusPending, AccountMemberListParamsStatusRejected:
		return true
	}
	return false
}

type AccountMemberAddParams struct {
	Body AccountMemberAddParamsBodyUnion `json:"body,required"`
}

func (r AccountMemberAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMemberAddParamsBody struct {
	// The contact email address of the user.
	Email    param.Field[string]                           `json:"email,required"`
	Policies param.Field[interface{}]                      `json:"policies"`
	Roles    param.Field[interface{}]                      `json:"roles"`
	Status   param.Field[AccountMemberAddParamsBodyStatus] `json:"status"`
}

func (r AccountMemberAddParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberAddParamsBody) implementsAccountMemberAddParamsBodyUnion() {}

// Satisfied by [AccountMemberAddParamsBodyIamCreateMemberWithRoles],
// [AccountMemberAddParamsBodyIamCreateMemberWithPolicies],
// [AccountMemberAddParamsBody].
type AccountMemberAddParamsBodyUnion interface {
	implementsAccountMemberAddParamsBodyUnion()
}

type AccountMemberAddParamsBodyIamCreateMemberWithRoles struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                                                 `json:"roles,required"`
	Status param.Field[AccountMemberAddParamsBodyIamCreateMemberWithRolesStatus] `json:"status"`
}

func (r AccountMemberAddParamsBodyIamCreateMemberWithRoles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberAddParamsBodyIamCreateMemberWithRoles) implementsAccountMemberAddParamsBodyUnion() {
}

type AccountMemberAddParamsBodyIamCreateMemberWithRolesStatus string

const (
	AccountMemberAddParamsBodyIamCreateMemberWithRolesStatusAccepted AccountMemberAddParamsBodyIamCreateMemberWithRolesStatus = "accepted"
	AccountMemberAddParamsBodyIamCreateMemberWithRolesStatusPending  AccountMemberAddParamsBodyIamCreateMemberWithRolesStatus = "pending"
)

func (r AccountMemberAddParamsBodyIamCreateMemberWithRolesStatus) IsKnown() bool {
	switch r {
	case AccountMemberAddParamsBodyIamCreateMemberWithRolesStatusAccepted, AccountMemberAddParamsBodyIamCreateMemberWithRolesStatusPending:
		return true
	}
	return false
}

type AccountMemberAddParamsBodyIamCreateMemberWithPolicies struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of policies associated with this member.
	Policies param.Field[[]IamCreateMemberPolicyParam]                                `json:"policies,required"`
	Status   param.Field[AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatus] `json:"status"`
}

func (r AccountMemberAddParamsBodyIamCreateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountMemberAddParamsBodyIamCreateMemberWithPolicies) implementsAccountMemberAddParamsBodyUnion() {
}

type AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatus string

const (
	AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatusAccepted AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatus = "accepted"
	AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatusPending  AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatus = "pending"
)

func (r AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatus) IsKnown() bool {
	switch r {
	case AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatusAccepted, AccountMemberAddParamsBodyIamCreateMemberWithPoliciesStatusPending:
		return true
	}
	return false
}

type AccountMemberAddParamsBodyStatus string

const (
	AccountMemberAddParamsBodyStatusAccepted AccountMemberAddParamsBodyStatus = "accepted"
	AccountMemberAddParamsBodyStatusPending  AccountMemberAddParamsBodyStatus = "pending"
)

func (r AccountMemberAddParamsBodyStatus) IsKnown() bool {
	switch r {
	case AccountMemberAddParamsBodyStatusAccepted, AccountMemberAddParamsBodyStatusPending:
		return true
	}
	return false
}
