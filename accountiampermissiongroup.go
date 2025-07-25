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

// AccountIamPermissionGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountIamPermissionGroupService] method instead.
type AccountIamPermissionGroupService struct {
	Options []option.RequestOption
}

// NewAccountIamPermissionGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountIamPermissionGroupService(opts ...option.RequestOption) (r *AccountIamPermissionGroupService) {
	r = &AccountIamPermissionGroupService{}
	r.Options = opts
	return
}

// Get information about a specific permission group in an account.
func (r *AccountIamPermissionGroupService) Get(ctx context.Context, accountID string, permissionGroupID string, opts ...option.RequestOption) (res *IamPermissionGroup, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if permissionGroupID == "" {
		err = errors.New("missing required permission_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/permission_groups/%s", accountID, permissionGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all the permissions groups for an account.
func (r *AccountIamPermissionGroupService) List(ctx context.Context, accountID string, query AccountIamPermissionGroupListParams, opts ...option.RequestOption) (res *AccountIamPermissionGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/permission_groups", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// A named group of permissions that map to a group of operations against
// resources.
type IamPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta IamPermissionGroupMeta `json:"meta"`
	// Name of the permission group.
	Name string                 `json:"name"`
	JSON iamPermissionGroupJSON `json:"-"`
}

// iamPermissionGroupJSON contains the JSON metadata for the struct
// [IamPermissionGroup]
type iamPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type IamPermissionGroupMeta struct {
	Key   string                     `json:"key"`
	Value string                     `json:"value"`
	JSON  iamPermissionGroupMetaJSON `json:"-"`
}

// iamPermissionGroupMetaJSON contains the JSON metadata for the struct
// [IamPermissionGroupMeta]
type iamPermissionGroupMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamPermissionGroupMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamPermissionGroupMetaJSON) RawJSON() string {
	return r.raw
}

// A named group of permissions that map to a group of operations against
// resources.
type IamPermissionGroupParam struct {
	// Identifier of the permission group.
	ID param.Field[string] `json:"id,required"`
	// Attributes associated to the permission group.
	Meta param.Field[IamPermissionGroupMetaParam] `json:"meta"`
}

func (r IamPermissionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes associated to the permission group.
type IamPermissionGroupMetaParam struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r IamPermissionGroupMetaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountIamPermissionGroupListResponse struct {
	Errors   []AccountIamPermissionGroupListResponseError   `json:"errors,required"`
	Messages []AccountIamPermissionGroupListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountIamPermissionGroupListResponseSuccess `json:"success,required"`
	// A set of permission groups that are specified to the policy.
	Result     []IamPermissionGroup                            `json:"result"`
	ResultInfo AccountIamPermissionGroupListResponseResultInfo `json:"result_info"`
	JSON       accountIamPermissionGroupListResponseJSON       `json:"-"`
}

// accountIamPermissionGroupListResponseJSON contains the JSON metadata for the
// struct [AccountIamPermissionGroupListResponse]
type accountIamPermissionGroupListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountIamPermissionGroupListResponseError struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccountIamPermissionGroupListResponseErrorsSource `json:"source"`
	JSON             accountIamPermissionGroupListResponseErrorJSON    `json:"-"`
}

// accountIamPermissionGroupListResponseErrorJSON contains the JSON metadata for
// the struct [AccountIamPermissionGroupListResponseError]
type accountIamPermissionGroupListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountIamPermissionGroupListResponseErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accountIamPermissionGroupListResponseErrorsSourceJSON `json:"-"`
}

// accountIamPermissionGroupListResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountIamPermissionGroupListResponseErrorsSource]
type accountIamPermissionGroupListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountIamPermissionGroupListResponseMessage struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccountIamPermissionGroupListResponseMessagesSource `json:"source"`
	JSON             accountIamPermissionGroupListResponseMessageJSON    `json:"-"`
}

// accountIamPermissionGroupListResponseMessageJSON contains the JSON metadata for
// the struct [AccountIamPermissionGroupListResponseMessage]
type accountIamPermissionGroupListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountIamPermissionGroupListResponseMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accountIamPermissionGroupListResponseMessagesSourceJSON `json:"-"`
}

// accountIamPermissionGroupListResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountIamPermissionGroupListResponseMessagesSource]
type accountIamPermissionGroupListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountIamPermissionGroupListResponseSuccess bool

const (
	AccountIamPermissionGroupListResponseSuccessTrue AccountIamPermissionGroupListResponseSuccess = true
)

func (r AccountIamPermissionGroupListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountIamPermissionGroupListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountIamPermissionGroupListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accountIamPermissionGroupListResponseResultInfoJSON `json:"-"`
}

// accountIamPermissionGroupListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountIamPermissionGroupListResponseResultInfo]
type accountIamPermissionGroupListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountIamPermissionGroupListParams struct {
	// ID of the permission group to be fetched.
	ID param.Field[string] `query:"id"`
	// Label of the permission group to be fetched.
	Label param.Field[string] `query:"label"`
	// Name of the permission group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountIamPermissionGroupListParams]'s query parameters as
// `url.Values`.
func (r AccountIamPermissionGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
