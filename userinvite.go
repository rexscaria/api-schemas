// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// UserInviteService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserInviteService] method instead.
type UserInviteService struct {
	Options []option.RequestOption
}

// NewUserInviteService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserInviteService(opts ...option.RequestOption) (r *UserInviteService) {
	r = &UserInviteService{}
	r.Options = opts
	return
}

// Gets the details of an invitation.
func (r *UserInviteService) Get(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *IamSingleInvite, err error) {
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all invitations associated with my user.
func (r *UserInviteService) List(ctx context.Context, opts ...option.RequestOption) (res *UserInviteListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Responds to an invitation.
func (r *UserInviteService) Respond(ctx context.Context, inviteID string, body UserInviteRespondParams, opts ...option.RequestOption) (res *IamSingleInvite, err error) {
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type IamSingleInvite struct {
	Errors   []IamSingleInviteError   `json:"errors,required"`
	Messages []IamSingleInviteMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success IamSingleInviteSuccess `json:"success,required"`
	Result  IamUserInvite          `json:"result"`
	JSON    iamSingleInviteJSON    `json:"-"`
}

// iamSingleInviteJSON contains the JSON metadata for the struct [IamSingleInvite]
type iamSingleInviteJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleInvite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleInviteJSON) RawJSON() string {
	return r.raw
}

type IamSingleInviteError struct {
	Code             int64                       `json:"code,required"`
	Message          string                      `json:"message,required"`
	DocumentationURL string                      `json:"documentation_url"`
	Source           IamSingleInviteErrorsSource `json:"source"`
	JSON             iamSingleInviteErrorJSON    `json:"-"`
}

// iamSingleInviteErrorJSON contains the JSON metadata for the struct
// [IamSingleInviteError]
type iamSingleInviteErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleInviteError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleInviteErrorJSON) RawJSON() string {
	return r.raw
}

type IamSingleInviteErrorsSource struct {
	Pointer string                          `json:"pointer"`
	JSON    iamSingleInviteErrorsSourceJSON `json:"-"`
}

// iamSingleInviteErrorsSourceJSON contains the JSON metadata for the struct
// [IamSingleInviteErrorsSource]
type iamSingleInviteErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleInviteErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleInviteErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IamSingleInviteMessage struct {
	Code             int64                         `json:"code,required"`
	Message          string                        `json:"message,required"`
	DocumentationURL string                        `json:"documentation_url"`
	Source           IamSingleInviteMessagesSource `json:"source"`
	JSON             iamSingleInviteMessageJSON    `json:"-"`
}

// iamSingleInviteMessageJSON contains the JSON metadata for the struct
// [IamSingleInviteMessage]
type iamSingleInviteMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IamSingleInviteMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleInviteMessageJSON) RawJSON() string {
	return r.raw
}

type IamSingleInviteMessagesSource struct {
	Pointer string                            `json:"pointer"`
	JSON    iamSingleInviteMessagesSourceJSON `json:"-"`
}

// iamSingleInviteMessagesSourceJSON contains the JSON metadata for the struct
// [IamSingleInviteMessagesSource]
type iamSingleInviteMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleInviteMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleInviteMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IamSingleInviteSuccess bool

const (
	IamSingleInviteSuccessTrue IamSingleInviteSuccess = true
)

func (r IamSingleInviteSuccess) IsKnown() bool {
	switch r {
	case IamSingleInviteSuccessTrue:
		return true
	}
	return false
}

type IamUserInvite struct {
	// ID of the user to add to the organization.
	InvitedMemberID string `json:"invited_member_id,required,nullable"`
	// ID of the organization the user will be added to.
	OrganizationID string `json:"organization_id,required"`
	// Invite identifier tag.
	ID string `json:"id"`
	// When the invite is no longer active.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The email address of the user who created the invite.
	InvitedBy string `json:"invited_by"`
	// Email address of the user to add to the organization.
	InvitedMemberEmail string `json:"invited_member_email"`
	// When the invite was sent.
	InvitedOn                        time.Time `json:"invited_on" format:"date-time"`
	OrganizationIsEnforcingTwofactor bool      `json:"organization_is_enforcing_twofactor"`
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Current status of the invitation.
	Status IamUserInviteStatus `json:"status"`
	JSON   iamUserInviteJSON   `json:"-"`
}

// iamUserInviteJSON contains the JSON metadata for the struct [IamUserInvite]
type iamUserInviteJSON struct {
	InvitedMemberID                  apijson.Field
	OrganizationID                   apijson.Field
	ID                               apijson.Field
	ExpiresOn                        apijson.Field
	InvitedBy                        apijson.Field
	InvitedMemberEmail               apijson.Field
	InvitedOn                        apijson.Field
	OrganizationIsEnforcingTwofactor apijson.Field
	OrganizationName                 apijson.Field
	Roles                            apijson.Field
	Status                           apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *IamUserInvite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamUserInviteJSON) RawJSON() string {
	return r.raw
}

// Current status of the invitation.
type IamUserInviteStatus string

const (
	IamUserInviteStatusPending  IamUserInviteStatus = "pending"
	IamUserInviteStatusAccepted IamUserInviteStatus = "accepted"
	IamUserInviteStatusRejected IamUserInviteStatus = "rejected"
	IamUserInviteStatusExpired  IamUserInviteStatus = "expired"
)

func (r IamUserInviteStatus) IsKnown() bool {
	switch r {
	case IamUserInviteStatusPending, IamUserInviteStatusAccepted, IamUserInviteStatusRejected, IamUserInviteStatusExpired:
		return true
	}
	return false
}

type UserInviteListResponse struct {
	Errors   []UserInviteListResponseError   `json:"errors,required"`
	Messages []UserInviteListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    UserInviteListResponseSuccess    `json:"success,required"`
	Result     []IamUserInvite                  `json:"result"`
	ResultInfo UserInviteListResponseResultInfo `json:"result_info"`
	JSON       userInviteListResponseJSON       `json:"-"`
}

// userInviteListResponseJSON contains the JSON metadata for the struct
// [UserInviteListResponse]
type userInviteListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseJSON) RawJSON() string {
	return r.raw
}

type UserInviteListResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           UserInviteListResponseErrorsSource `json:"source"`
	JSON             userInviteListResponseErrorJSON    `json:"-"`
}

// userInviteListResponseErrorJSON contains the JSON metadata for the struct
// [UserInviteListResponseError]
type userInviteListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserInviteListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type UserInviteListResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    userInviteListResponseErrorsSourceJSON `json:"-"`
}

// userInviteListResponseErrorsSourceJSON contains the JSON metadata for the struct
// [UserInviteListResponseErrorsSource]
type userInviteListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserInviteListResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           UserInviteListResponseMessagesSource `json:"source"`
	JSON             userInviteListResponseMessageJSON    `json:"-"`
}

// userInviteListResponseMessageJSON contains the JSON metadata for the struct
// [UserInviteListResponseMessage]
type userInviteListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserInviteListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type UserInviteListResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    userInviteListResponseMessagesSourceJSON `json:"-"`
}

// userInviteListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [UserInviteListResponseMessagesSource]
type userInviteListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserInviteListResponseSuccess bool

const (
	UserInviteListResponseSuccessTrue UserInviteListResponseSuccess = true
)

func (r UserInviteListResponseSuccess) IsKnown() bool {
	switch r {
	case UserInviteListResponseSuccessTrue:
		return true
	}
	return false
}

type UserInviteListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       userInviteListResponseResultInfoJSON `json:"-"`
}

// userInviteListResponseResultInfoJSON contains the JSON metadata for the struct
// [UserInviteListResponseResultInfo]
type userInviteListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type UserInviteRespondParams struct {
	// Status of your response to the invitation (rejected or accepted).
	Status param.Field[UserInviteRespondParamsStatus] `json:"status,required"`
}

func (r UserInviteRespondParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of your response to the invitation (rejected or accepted).
type UserInviteRespondParamsStatus string

const (
	UserInviteRespondParamsStatusAccepted UserInviteRespondParamsStatus = "accepted"
	UserInviteRespondParamsStatusRejected UserInviteRespondParamsStatus = "rejected"
)

func (r UserInviteRespondParamsStatus) IsKnown() bool {
	switch r {
	case UserInviteRespondParamsStatusAccepted, UserInviteRespondParamsStatusRejected:
		return true
	}
	return false
}
