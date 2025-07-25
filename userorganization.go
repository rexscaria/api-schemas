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
func (r *UserOrganizationService) Leave(ctx context.Context, organizationID string, opts ...option.RequestOption) (res *UserOrganizationLeaveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if organizationID == "" {
		err = errors.New("missing required organization_id parameter")
		return
	}
	path := fmt.Sprintf("user/organizations/%s", organizationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type UserOrganizationGetResponse struct {
	Errors   []UserOrganizationGetResponseError   `json:"errors,required"`
	Messages []UserOrganizationGetResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success UserOrganizationGetResponseSuccess `json:"success,required"`
	Result  interface{}                        `json:"result"`
	JSON    userOrganizationGetResponseJSON    `json:"-"`
}

// userOrganizationGetResponseJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponse]
type userOrganizationGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type UserOrganizationGetResponseError struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           UserOrganizationGetResponseErrorsSource `json:"source"`
	JSON             userOrganizationGetResponseErrorJSON    `json:"-"`
}

// userOrganizationGetResponseErrorJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponseError]
type userOrganizationGetResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserOrganizationGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationGetResponseErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    userOrganizationGetResponseErrorsSourceJSON `json:"-"`
}

// userOrganizationGetResponseErrorsSourceJSON contains the JSON metadata for the
// struct [UserOrganizationGetResponseErrorsSource]
type userOrganizationGetResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationGetResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationGetResponseMessage struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           UserOrganizationGetResponseMessagesSource `json:"source"`
	JSON             userOrganizationGetResponseMessageJSON    `json:"-"`
}

// userOrganizationGetResponseMessageJSON contains the JSON metadata for the struct
// [UserOrganizationGetResponseMessage]
type userOrganizationGetResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserOrganizationGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationGetResponseMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    userOrganizationGetResponseMessagesSourceJSON `json:"-"`
}

// userOrganizationGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [UserOrganizationGetResponseMessagesSource]
type userOrganizationGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserOrganizationGetResponseSuccess bool

const (
	UserOrganizationGetResponseSuccessTrue UserOrganizationGetResponseSuccess = true
)

func (r UserOrganizationGetResponseSuccess) IsKnown() bool {
	switch r {
	case UserOrganizationGetResponseSuccessTrue:
		return true
	}
	return false
}

type UserOrganizationListResponse struct {
	Errors   []UserOrganizationListResponseError   `json:"errors,required"`
	Messages []UserOrganizationListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    UserOrganizationListResponseSuccess    `json:"success,required"`
	Result     []UserOrganizationListResponseResult   `json:"result"`
	ResultInfo UserOrganizationListResponseResultInfo `json:"result_info"`
	JSON       userOrganizationListResponseJSON       `json:"-"`
}

// userOrganizationListResponseJSON contains the JSON metadata for the struct
// [UserOrganizationListResponse]
type userOrganizationListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           UserOrganizationListResponseErrorsSource `json:"source"`
	JSON             userOrganizationListResponseErrorJSON    `json:"-"`
}

// userOrganizationListResponseErrorJSON contains the JSON metadata for the struct
// [UserOrganizationListResponseError]
type userOrganizationListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserOrganizationListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    userOrganizationListResponseErrorsSourceJSON `json:"-"`
}

// userOrganizationListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [UserOrganizationListResponseErrorsSource]
type userOrganizationListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           UserOrganizationListResponseMessagesSource `json:"source"`
	JSON             userOrganizationListResponseMessageJSON    `json:"-"`
}

// userOrganizationListResponseMessageJSON contains the JSON metadata for the
// struct [UserOrganizationListResponseMessage]
type userOrganizationListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserOrganizationListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UserOrganizationListResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    userOrganizationListResponseMessagesSourceJSON `json:"-"`
}

// userOrganizationListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [UserOrganizationListResponseMessagesSource]
type userOrganizationListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserOrganizationListResponseSuccess bool

const (
	UserOrganizationListResponseSuccessTrue UserOrganizationListResponseSuccess = true
)

func (r UserOrganizationListResponseSuccess) IsKnown() bool {
	switch r {
	case UserOrganizationListResponseSuccessTrue:
		return true
	}
	return false
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

type UserOrganizationListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       userOrganizationListResponseResultInfoJSON `json:"-"`
}

// userOrganizationListResponseResultInfoJSON contains the JSON metadata for the
// struct [UserOrganizationListResponseResultInfo]
type userOrganizationListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserOrganizationListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userOrganizationListResponseResultInfoJSON) RawJSON() string {
	return r.raw
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
