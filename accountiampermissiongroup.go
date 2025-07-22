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
	// Identifier of the group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta IamPermissionGroupMeta `json:"meta"`
	// Name of the group.
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
	// Identifier of the group.
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
	// A set of permission groups that are specified to the policy.
	Result []IamPermissionGroup                      `json:"result"`
	JSON   accountIamPermissionGroupListResponseJSON `json:"-"`
	IamAPIResponseCollection
}

// accountIamPermissionGroupListResponseJSON contains the JSON metadata for the
// struct [AccountIamPermissionGroupListResponse]
type accountIamPermissionGroupListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIamPermissionGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountIamPermissionGroupListResponseJSON) RawJSON() string {
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
