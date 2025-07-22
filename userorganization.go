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

// UserOrganizationService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserOrganizationService] method instead.
type UserOrganizationService struct {
	Options []option.RequestOption
}

// NewUserOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserOrganizationService(opts ...option.RequestOption) (r *UserOrganizationService) {
	r = &UserOrganizationService{}
	r.Options = opts
	return
}

// Gets a specific organization the user is associated with.
//
// Deprecated: deprecated
func (r *UserOrganizationService) Get(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *UserOrganizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists organizations the user is associated with.
//
// Deprecated: deprecated
func (r *UserOrganizationService) List(ctx context.Context, query UserOrganizationListParams, opts ...option.RequestOption) (res *UserOrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/organizations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Removes association to an organization.
//
// Deprecated: deprecated
func (r *UserOrganizationService) Leave(ctx context.Context, organizationID string, body UserOrganizationLeaveParams, opts ...option.RequestOption) (res *UserOrganizationLeaveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type UserOrganizationGetResponse struct {
	Result interface{}                     `json:"result"`
	JSON   userOrganizationGetResponseJSON `json:"-"`
	APIResponseSingleIam
}

// userOrganizationGetResponseJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponse]
type userOrganizationGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponse struct {
	Result []UserOrganizationListResponseResult `json:"result"`
	JSON   userOrganizationListResponseJSON     `json:"-"`
	IamAPIResponseCollection
}

// userOrganizationListResponseJSON contains the JSON metadata for the struct
// [UserOrganizationListResponse]
type userOrganizationListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Organization name.
	Name string `json:"name"`
	// Access permissions for this User.
	Permissions []string `json:"permissions"`
	// List of roles that a user has within an organization.
	Roles []string `json:"roles"`
	// Whether the user is a member of the organization or has an invitation pending.
	Status UserOrganizationListResponseResultStatus `json:"status"`
	JSON   userOrganizationListResponseResultJSON   `json:"-"`
}

// userOrganizationListResponseResultJSON contains the JSON metadata for the struct
// [UserOrganizationListResponseResult]
type userOrganizationListResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the user is a member of the organization or has an invitation pending.
type UserOrganizationListResponseResultStatus string

const (
	UserOrganizationListResponseResultStatusMember  UserOrganizationListResponseResultStatus = "member"
	UserOrganizationListResponseResultStatusInvited UserOrganizationListResponseResultStatus = "invited"
)

func (r UserOrganizationListResponseResultStatus) IsKnown() bool {
	switch r {
	case UserOrganizationListResponseResultStatusMember, UserOrganizationListResponseResultStatusInvited:
		return true
	}
	return false
}

type UserOrganizationLeaveResponse struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON userOrganizationLeaveResponseJSON `json:"-"`
}

// userOrganizationLeaveResponseJSON contains the JSON metadata for the struct
// [UserOrganizationLeaveResponse]
type userOrganizationLeaveResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationLeaveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationLeaveResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListParams struct {
	// Direction to order organizations.
	Direction param.Field[UserOrganizationListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[UserOrganizationListParamsMatch] `query:"match"`
	// Organization name.
	Name param.Field[string] `query:"name"`
	// Field to order organizations by.
	Order param.Field[UserOrganizationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of organizations per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the user is a member of the organization or has an inivitation pending.
	Status param.Field[UserOrganizationListParamsStatus] `query:"status"`
}

// URLQuery serializes [UserOrganizationListParams]'s query parameters as
// `url.Values`.
func (r UserOrganizationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order organizations.
type UserOrganizationListParamsDirection string

const (
	UserOrganizationListParamsDirectionAsc  UserOrganizationListParamsDirection = "asc"
	UserOrganizationListParamsDirectionDesc UserOrganizationListParamsDirection = "desc"
)

func (r UserOrganizationListParamsDirection) IsKnown() bool {
	switch r {
	case UserOrganizationListParamsDirectionAsc, UserOrganizationListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any).
type UserOrganizationListParamsMatch string

const (
	UserOrganizationListParamsMatchAny UserOrganizationListParamsMatch = "any"
	UserOrganizationListParamsMatchAll UserOrganizationListParamsMatch = "all"
)

func (r UserOrganizationListParamsMatch) IsKnown() bool {
	switch r {
	case UserOrganizationListParamsMatchAny, UserOrganizationListParamsMatchAll:
		return true
	}
	return false
}

// Field to order organizations by.
type UserOrganizationListParamsOrder string

const (
	UserOrganizationListParamsOrderID     UserOrganizationListParamsOrder = "id"
	UserOrganizationListParamsOrderName   UserOrganizationListParamsOrder = "name"
	UserOrganizationListParamsOrderStatus UserOrganizationListParamsOrder = "status"
)

func (r UserOrganizationListParamsOrder) IsKnown() bool {
	switch r {
	case UserOrganizationListParamsOrderID, UserOrganizationListParamsOrderName, UserOrganizationListParamsOrderStatus:
		return true
	}
	return false
}

// Whether the user is a member of the organization or has an inivitation pending.
type UserOrganizationListParamsStatus string

const (
	UserOrganizationListParamsStatusMember  UserOrganizationListParamsStatus = "member"
	UserOrganizationListParamsStatusInvited UserOrganizationListParamsStatus = "invited"
)

func (r UserOrganizationListParamsStatus) IsKnown() bool {
	switch r {
	case UserOrganizationListParamsStatusMember, UserOrganizationListParamsStatusInvited:
		return true
	}
	return false
}

type UserOrganizationLeaveParams struct {
	Body interface{} `json:"body,required"`
}

func (r UserOrganizationLeaveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
