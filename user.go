// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// UserService contains methods and other services that help with interacting with
// the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options                []option.RequestOption
	Billing                *UserBillingService
	Firewall               *UserFirewallService
	Invites                *UserInviteService
	LoadBalancers          *UserLoadBalancerService
	LoadBalancingAnalytics *UserLoadBalancingAnalyticService
	Organizations          *UserOrganizationService
	Subscriptions          *UserSubscriptionService
	Tokens                 *UserTokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.Billing = NewUserBillingService(opts...)
	r.Firewall = NewUserFirewallService(opts...)
	r.Invites = NewUserInviteService(opts...)
	r.LoadBalancers = NewUserLoadBalancerService(opts...)
	r.LoadBalancingAnalytics = NewUserLoadBalancingAnalyticService(opts...)
	r.Organizations = NewUserOrganizationService(opts...)
	r.Subscriptions = NewUserSubscriptionService(opts...)
	r.Tokens = NewUserTokenService(opts...)
	return
}

// User Details
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *IamSingleUser, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Update(ctx context.Context, body UserUpdateParams, opts ...option.RequestOption) (res *IamSingleUser, err error) {
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *UserService) ListAuditLogs(ctx context.Context, query UserListAuditLogsParams, opts ...option.RequestOption) (res *AaaAuditLogs, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/audit_logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IamSingleUser struct {
	Result interface{}       `json:"result"`
	JSON   iamSingleUserJSON `json:"-"`
	APIResponseSingleIam
}

// iamSingleUserJSON contains the JSON metadata for the struct [IamSingleUser]
type iamSingleUserJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IamSingleUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r iamSingleUserJSON) RawJSON() string {
	return r.raw
}

type UserUpdateParams struct {
	// The country in which the user lives.
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
	// User's telephone number
	Telephone param.Field[string] `json:"telephone"`
	// The zipcode or postal code where the user lives.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserListAuditLogsParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                        `query:"id"`
	Action param.Field[UserListAuditLogsParamsAction] `query:"action"`
	Actor  param.Field[UserListAuditLogsParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. A `full-date`
	// that conforms to RFC3339.
	Before param.Field[UserListAuditLogsParamsBeforeUnion] `query:"before" format:"date"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[UserListAuditLogsParamsDirection] `query:"direction"`
	// Indicates that this request is an export of logs in CSV format.
	Export param.Field[bool] `query:"export"`
	// Indicates whether or not to hide user level audit logs.
	HideUserLogs param.Field[bool] `query:"hide_user_logs"`
	// Defines which page of results to return.
	Page param.Field[float64] `query:"page"`
	// Sets the number of results to return per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limits the returned results to logs newer than the specified date. A `full-date`
	// that conforms to RFC3339.
	Since param.Field[UserListAuditLogsParamsSinceUnion] `query:"since" format:"date"`
	Zone  param.Field[UserListAuditLogsParamsZone]       `query:"zone"`
}

// URLQuery serializes [UserListAuditLogsParams]'s query parameters as
// `url.Values`.
func (r UserListAuditLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListAuditLogsParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserListAuditLogsParamsAction]'s query parameters as
// `url.Values`.
func (r UserListAuditLogsParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserListAuditLogsParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [UserListAuditLogsParamsActor]'s query parameters as
// `url.Values`.
func (r UserListAuditLogsParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Limits the returned results to logs older than the specified date. A `full-date`
// that conforms to RFC3339.
//
// Satisfied by [shared.UnionTime], [shared.UnionTime].
type UserListAuditLogsParamsBeforeUnion interface {
	ImplementsUserListAuditLogsParamsBeforeUnion()
}

// Changes the direction of the chronological sorting.
type UserListAuditLogsParamsDirection string

const (
	UserListAuditLogsParamsDirectionDesc UserListAuditLogsParamsDirection = "desc"
	UserListAuditLogsParamsDirectionAsc  UserListAuditLogsParamsDirection = "asc"
)

func (r UserListAuditLogsParamsDirection) IsKnown() bool {
	switch r {
	case UserListAuditLogsParamsDirectionDesc, UserListAuditLogsParamsDirectionAsc:
		return true
	}
	return false
}

// Limits the returned results to logs newer than the specified date. A `full-date`
// that conforms to RFC3339.
//
// Satisfied by [shared.UnionTime], [shared.UnionTime].
type UserListAuditLogsParamsSinceUnion interface {
	ImplementsUserListAuditLogsParamsSinceUnion()
}

type UserListAuditLogsParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [UserListAuditLogsParamsZone]'s query parameters as
// `url.Values`.
func (r UserListAuditLogsParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
