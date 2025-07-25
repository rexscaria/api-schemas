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

// AccountIamResourceGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIamResourceGroupService] method instead.
type AccountIamResourceGroupService struct {
	Options []option.RequestOption
}

// NewAccountIamResourceGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIamResourceGroupService(opts ...option.RequestOption) (r *AccountIamResourceGroupService) {
	r = &AccountIamResourceGroupService{}
	r.Options = opts
	return
}

// Create a new Resource Group under the specified account.
func (r *AccountIamResourceGroupService) New(ctx context.Context, accountID string, body AccountIamResourceGroupNewParams, opts ...option.RequestOption) (res *AccountIamResourceGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a specific resource group in an account.
func (r *AccountIamResourceGroupService) Get(ctx context.Context, accountID string, resourceGroupID string, opts ...option.RequestOption) (res *IamResourceGroup, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", accountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify an existing resource group.
func (r *AccountIamResourceGroupService) Update(ctx context.Context, accountID string, resourceGroupID string, body AccountIamResourceGroupUpdateParams, opts ...option.RequestOption) (res *IamResourceGroup, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", accountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all the resource groups for an account.
func (r *AccountIamResourceGroupService) List(ctx context.Context, accountID string, query AccountIamResourceGroupListParams, opts ...option.RequestOption) (res *AccountIamResourceGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove a resource group from an account.
func (r *AccountIamResourceGroupService) Delete(ctx context.Context, accountID string, resourceGroupID string, opts ...option.RequestOption) (res *IamAPIResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if resourceGroupID == "" {
		err = errors.New("missing required resource_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/resource_groups/%s", accountID, resourceGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A scope is a combination of scope objects which provides additional context.
type IamCreateScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects []IamCreateScopeObject `json:"objects,required"`
	JSON    iamCreateScopeJSON     `json:"-"`
}

// iamCreateScopeJSON contains the JSON metadata for the struct [IamCreateScope]
type iamCreateScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCreateScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCreateScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamCreateScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                   `json:"key,required"`
	JSON iamCreateScopeObjectJSON `json:"-"`
}

// iamCreateScopeObjectJSON contains the JSON metadata for the struct
// [IamCreateScopeObject]
type iamCreateScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamCreateScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamCreateScopeObjectJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IamCreateScopeParam struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key param.Field[string] `json:"key,required"`
	// A list of scope objects for additional context. The number of Scope objects
	// should not be zero.
	Objects param.Field[[]IamCreateScopeObjectParam] `json:"objects,required"`
}

func (r IamCreateScopeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamCreateScopeObjectParam struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key param.Field[string] `json:"key,required"`
}

func (r IamCreateScopeObjectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type IamResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []IamResourceGroupScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta IamResourceGroupMeta `json:"meta"`
	// Name of the resource group.
	Name string               `json:"name"`
	JSON iamResourceGroupJSON `json:"-"`
}

// iamResourceGroupJSON contains the JSON metadata for the struct
// [IamResourceGroup]
type iamResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type IamResourceGroupScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []IamResourceGroupScopeObject `json:"objects,required"`
	JSON    iamResourceGroupScopeJSON     `json:"-"`
}

// iamResourceGroupScopeJSON contains the JSON metadata for the struct
// [IamResourceGroupScope]
type iamResourceGroupScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type IamResourceGroupScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                          `json:"key,required"`
	JSON iamResourceGroupScopeObjectJSON `json:"-"`
}

// iamResourceGroupScopeObjectJSON contains the JSON metadata for the struct
// [IamResourceGroupScopeObject]
type iamResourceGroupScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type IamResourceGroupMeta struct {
	Key   string                   `json:"key"`
	Value string                   `json:"value"`
	JSON  iamResourceGroupMetaJSON `json:"-"`
}

// iamResourceGroupMetaJSON contains the JSON metadata for the struct
// [IamResourceGroupMeta]
type iamResourceGroupMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamResourceGroupMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamResourceGroupMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type AccountIamResourceGroupNewResponse struct {
	// Identifier of the group.
	ID string `json:"id"`
	// Attributes associated to the resource group.
	Meta interface{} `json:"meta"`
	// A scope is a combination of scope objects which provides additional context.
	Scope IamCreateScope                         `json:"scope"`
	JSON  accountIamResourceGroupNewResponseJSON `json:"-"`
}

// accountIamResourceGroupNewResponseJSON contains the JSON metadata for the struct
// [AccountIamResourceGroupNewResponse]
type accountIamResourceGroupNewResponseJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamResourceGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupListResponse struct {
	Errors   []AccountIamResourceGroupListResponseError   `json:"errors,required"`
	Messages []AccountIamResourceGroupListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIamResourceGroupListResponseSuccess `json:"success,required"`
	// A list of resource groups that the policy applies to.
	Result     []IamResourceGroup                            `json:"result"`
	ResultInfo AccountIamResourceGroupListResponseResultInfo `json:"result_info"`
	JSON       accountIamResourceGroupListResponseJSON       `json:"-"`
}

// accountIamResourceGroupListResponseJSON contains the JSON metadata for the
// struct [AccountIamResourceGroupListResponse]
type accountIamResourceGroupListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupListResponseError struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccountIamResourceGroupListResponseErrorsSource `json:"source"`
	JSON             accountIamResourceGroupListResponseErrorJSON    `json:"-"`
}

// accountIamResourceGroupListResponseErrorJSON contains the JSON metadata for the
// struct [AccountIamResourceGroupListResponseError]
type accountIamResourceGroupListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupListResponseErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accountIamResourceGroupListResponseErrorsSourceJSON `json:"-"`
}

// accountIamResourceGroupListResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountIamResourceGroupListResponseErrorsSource]
type accountIamResourceGroupListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupListResponseMessage struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccountIamResourceGroupListResponseMessagesSource `json:"source"`
	JSON             accountIamResourceGroupListResponseMessageJSON    `json:"-"`
}

// accountIamResourceGroupListResponseMessageJSON contains the JSON metadata for
// the struct [AccountIamResourceGroupListResponseMessage]
type accountIamResourceGroupListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupListResponseMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accountIamResourceGroupListResponseMessagesSourceJSON `json:"-"`
}

// accountIamResourceGroupListResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountIamResourceGroupListResponseMessagesSource]
type accountIamResourceGroupListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIamResourceGroupListResponseSuccess bool

const (
	AccountIamResourceGroupListResponseSuccessTrue AccountIamResourceGroupListResponseSuccess = true
)

func (r AccountIamResourceGroupListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIamResourceGroupListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIamResourceGroupListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       accountIamResourceGroupListResponseResultInfoJSON `json:"-"`
}

// accountIamResourceGroupListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountIamResourceGroupListResponseResultInfo]
type accountIamResourceGroupListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamResourceGroupListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamResourceGroupListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountIamResourceGroupNewParams struct {
	// Name of the resource group
	Name param.Field[string] `json:"name,required"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IamCreateScopeParam] `json:"scope,required"`
}

func (r AccountIamResourceGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIamResourceGroupUpdateParams struct {
	// Name of the resource group
	Name param.Field[string] `json:"name"`
	// A scope is a combination of scope objects which provides additional context.
	Scope param.Field[IamCreateScopeParam] `json:"scope"`
}

func (r AccountIamResourceGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIamResourceGroupListParams struct {
	// ID of the resource group to be fetched.
	ID param.Field[string] `query:"id"`
	// Name of the resource group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountIamResourceGroupListParams]'s query parameters as
// `url.Values`.
func (r AccountIamResourceGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
