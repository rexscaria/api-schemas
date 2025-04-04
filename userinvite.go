// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
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
	Result IamUserInvite       `json:"result"`
	JSON   iamSingleInviteJSON `json:"-"`
	APIResponseSingleIam
}

// iamSingleInviteJSON contains the JSON metadata for the struct [IamSingleInvite]
type iamSingleInviteJSON struct {
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
	Result []IamUserInvite            `json:"result"`
	JSON   userInviteListResponseJSON `json:"-"`
	IamAPIResponseCollection
}

// userInviteListResponseJSON contains the JSON metadata for the struct
// [UserInviteListResponse]
type userInviteListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserInviteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userInviteListResponseJSON) RawJSON() string {
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
