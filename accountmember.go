// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
func (r *AccountMemberService) Remove(ctx context.Context, accountID string, memberID string, body AccountMemberRemoveParams, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
	Result IamMemberWithPolicies                   `json:"result"`
	JSON   iamSingleMemberResponseWithPoliciesJSON `json:"-"`
	APIResponseSingleIam
}

// iamSingleMemberResponseWithPoliciesJSON contains the JSON metadata for the
// struct [IamSingleMemberResponseWithPolicies]
type iamSingleMemberResponseWithPoliciesJSON struct {
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

type AccountMemberListResponse struct {
	Result []IamMemberWithPolicies       `json:"result"`
	JSON   accountMemberListResponseJSON `json:"-"`
	IamAPIResponseCollection
}

// accountMemberListResponseJSON contains the JSON metadata for the struct
// [AccountMemberListResponse]
type accountMemberListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMemberListResponseJSON) RawJSON() string {
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

type AccountMemberRemoveParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMemberRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
