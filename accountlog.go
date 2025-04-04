// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountLogService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogService] method instead.
type AccountLogService struct {
	Options []option.RequestOption
	Control *AccountLogControlService
}

// NewAccountLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountLogService(opts ...option.RequestOption) (r *AccountLogService) {
	r = &AccountLogService{}
	r.Options = opts
	r.Control = NewAccountLogControlService(opts...)
	return
}

// Gets a list of audit logs for an account. <br /> <br /> This is the beta release
// of Audit Logs Version 2. Since this is a beta version, there may be gaps or
// missing entries in the available audit logs. Be aware of the following
// limitations. <br /> <ul> <li>Audit logs are available only for the past 30 days.
// <br /></li> <li>Error handling is not yet implemented. <br /> </li> </ul>
func (r *AccountLogService) GetAuditLogs(ctx context.Context, accountID string, query AccountLogGetAuditLogsParams, opts ...option.RequestOption) (res *AccountLogGetAuditLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/audit", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountLogGetAuditLogsResponse struct {
	Errors []AccountLogGetAuditLogsResponseError  `json:"errors"`
	Result []AccountLogGetAuditLogsResponseResult `json:"result"`
	// Provides information about the result of the request, including count and
	// cursor.
	ResultInfo AccountLogGetAuditLogsResponseResultInfo `json:"result_info"`
	// Indicates whether the API call was successful
	Success AccountLogGetAuditLogsResponseSuccess `json:"success"`
	JSON    accountLogGetAuditLogsResponseJSON    `json:"-"`
}

// accountLogGetAuditLogsResponseJSON contains the JSON metadata for the struct
// [AccountLogGetAuditLogsResponse]
type accountLogGetAuditLogsResponseJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountLogGetAuditLogsResponseError struct {
	Message string                                  `json:"message,required"`
	JSON    accountLogGetAuditLogsResponseErrorJSON `json:"-"`
}

// accountLogGetAuditLogsResponseErrorJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseError]
type accountLogGetAuditLogsResponseErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountLogGetAuditLogsResponseResult struct {
	// A unique identifier for the audit log entry.
	ID string `json:"id"`
	// Contains account related information.
	Account AccountLogGetAuditLogsResponseResultAccount `json:"account"`
	// Provides information about the action performed.
	Action AccountLogGetAuditLogsResponseResultAction `json:"action"`
	// Provides details about the actor who performed the action.
	Actor AccountLogGetAuditLogsResponseResultActor `json:"actor"`
	// Provides raw information about the request and response.
	Raw AccountLogGetAuditLogsResponseResultRaw `json:"raw"`
	// Provides details about the affected resource.
	Resource AccountLogGetAuditLogsResponseResultResource `json:"resource"`
	// Provides details about the zone affected by the action.
	Zone AccountLogGetAuditLogsResponseResultZone `json:"zone"`
	JSON accountLogGetAuditLogsResponseResultJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseResult]
type accountLogGetAuditLogsResponseResultJSON struct {
	ID          apijson.Field
	Account     apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Raw         apijson.Field
	Resource    apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultJSON) RawJSON() string {
	return r.raw
}

// Contains account related information.
type AccountLogGetAuditLogsResponseResultAccount struct {
	// A unique identifier for the account.
	ID string `json:"id"`
	// A string that identifies the account name.
	Name string                                          `json:"name"`
	JSON accountLogGetAuditLogsResponseResultAccountJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultAccountJSON contains the JSON metadata for
// the struct [AccountLogGetAuditLogsResponseResultAccount]
type accountLogGetAuditLogsResponseResultAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultAccountJSON) RawJSON() string {
	return r.raw
}

// Provides information about the action performed.
type AccountLogGetAuditLogsResponseResultAction struct {
	// A short description of the action performed.
	Description string `json:"description"`
	// The result of the action, indicating success or failure.
	Result string `json:"result"`
	// A timestamp indicating when the action was logged.
	Time time.Time `json:"time" format:"date-time"`
	// A short string that describes the action that was performed.
	Type string                                         `json:"type"`
	JSON accountLogGetAuditLogsResponseResultActionJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultActionJSON contains the JSON metadata for
// the struct [AccountLogGetAuditLogsResponseResultAction]
type accountLogGetAuditLogsResponseResultActionJSON struct {
	Description apijson.Field
	Result      apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultActionJSON) RawJSON() string {
	return r.raw
}

// Provides details about the actor who performed the action.
type AccountLogGetAuditLogsResponseResultActor struct {
	// The ID of the actor who performed the action. If a user performed the action,
	// this will be their User ID.
	ID      string                                           `json:"id"`
	Context AccountLogGetAuditLogsResponseResultActorContext `json:"context"`
	// The email of the actor who performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IPAddress string `json:"ip_address" format:"ipv4 | ipv6"`
	// Filters by the API token ID when the actor context is an api_token.
	TokenID string `json:"token_id"`
	// Filters by the API token name when the actor context is an api_token.
	TokenName string `json:"token_name"`
	// The type of actor.
	Type AccountLogGetAuditLogsResponseResultActorType `json:"type"`
	JSON accountLogGetAuditLogsResponseResultActorJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultActorJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseResultActor]
type accountLogGetAuditLogsResponseResultActorJSON struct {
	ID          apijson.Field
	Context     apijson.Field
	Email       apijson.Field
	IPAddress   apijson.Field
	TokenID     apijson.Field
	TokenName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultActorJSON) RawJSON() string {
	return r.raw
}

type AccountLogGetAuditLogsResponseResultActorContext string

const (
	AccountLogGetAuditLogsResponseResultActorContextAPIKey      AccountLogGetAuditLogsResponseResultActorContext = "api_key"
	AccountLogGetAuditLogsResponseResultActorContextAPIToken    AccountLogGetAuditLogsResponseResultActorContext = "api_token"
	AccountLogGetAuditLogsResponseResultActorContextDash        AccountLogGetAuditLogsResponseResultActorContext = "dash"
	AccountLogGetAuditLogsResponseResultActorContextOAuth       AccountLogGetAuditLogsResponseResultActorContext = "oauth"
	AccountLogGetAuditLogsResponseResultActorContextOriginCaKey AccountLogGetAuditLogsResponseResultActorContext = "origin_ca_key"
)

func (r AccountLogGetAuditLogsResponseResultActorContext) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsResponseResultActorContextAPIKey, AccountLogGetAuditLogsResponseResultActorContextAPIToken, AccountLogGetAuditLogsResponseResultActorContextDash, AccountLogGetAuditLogsResponseResultActorContextOAuth, AccountLogGetAuditLogsResponseResultActorContextOriginCaKey:
		return true
	}
	return false
}

// The type of actor.
type AccountLogGetAuditLogsResponseResultActorType string

const (
	AccountLogGetAuditLogsResponseResultActorTypeAccount         AccountLogGetAuditLogsResponseResultActorType = "account"
	AccountLogGetAuditLogsResponseResultActorTypeCloudflareAdmin AccountLogGetAuditLogsResponseResultActorType = "cloudflare_admin"
	AccountLogGetAuditLogsResponseResultActorTypeSystem          AccountLogGetAuditLogsResponseResultActorType = "system"
	AccountLogGetAuditLogsResponseResultActorTypeUser            AccountLogGetAuditLogsResponseResultActorType = "user"
)

func (r AccountLogGetAuditLogsResponseResultActorType) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsResponseResultActorTypeAccount, AccountLogGetAuditLogsResponseResultActorTypeCloudflareAdmin, AccountLogGetAuditLogsResponseResultActorTypeSystem, AccountLogGetAuditLogsResponseResultActorTypeUser:
		return true
	}
	return false
}

// Provides raw information about the request and response.
type AccountLogGetAuditLogsResponseResultRaw struct {
	// The Cloudflare Ray ID for the request.
	CfRayID string `json:"cf_ray_id"`
	// The HTTP method of the request.
	Method string `json:"method"`
	// The HTTP response status code returned by the API.
	StatusCode int64 `json:"status_code"`
	// The URI of the request.
	Uri string `json:"uri"`
	// The client's user agent string sent with the request.
	UserAgent string                                      `json:"user_agent"`
	JSON      accountLogGetAuditLogsResponseResultRawJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultRawJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseResultRaw]
type accountLogGetAuditLogsResponseResultRawJSON struct {
	CfRayID     apijson.Field
	Method      apijson.Field
	StatusCode  apijson.Field
	Uri         apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultRaw) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultRawJSON) RawJSON() string {
	return r.raw
}

// Provides details about the affected resource.
type AccountLogGetAuditLogsResponseResultResource struct {
	// The unique identifier for the affected resource.
	ID string `json:"id"`
	// The Cloudflare product associated with the resource.
	Product  string      `json:"product"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
	// The scope of the resource.
	Scope interface{} `json:"scope"`
	// The type of the resource.
	Type string                                           `json:"type"`
	JSON accountLogGetAuditLogsResponseResultResourceJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultResourceJSON contains the JSON metadata for
// the struct [AccountLogGetAuditLogsResponseResultResource]
type accountLogGetAuditLogsResponseResultResourceJSON struct {
	ID          apijson.Field
	Product     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultResourceJSON) RawJSON() string {
	return r.raw
}

// Provides details about the zone affected by the action.
type AccountLogGetAuditLogsResponseResultZone struct {
	// A string that identifies the zone id.
	ID string `json:"id"`
	// A string that identifies the zone name.
	Name string                                       `json:"name"`
	JSON accountLogGetAuditLogsResponseResultZoneJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultZoneJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseResultZone]
type accountLogGetAuditLogsResponseResultZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultZoneJSON) RawJSON() string {
	return r.raw
}

// Provides information about the result of the request, including count and
// cursor.
type AccountLogGetAuditLogsResponseResultInfo struct {
	// The number of records returned in the response.
	Count string `json:"count"`
	// The cursor token used for pagination.
	Cursor string                                       `json:"cursor"`
	JSON   accountLogGetAuditLogsResponseResultInfoJSON `json:"-"`
}

// accountLogGetAuditLogsResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountLogGetAuditLogsResponseResultInfo]
type accountLogGetAuditLogsResponseResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogGetAuditLogsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountLogGetAuditLogsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Indicates whether the API call was successful
type AccountLogGetAuditLogsResponseSuccess bool

const (
	AccountLogGetAuditLogsResponseSuccessTrue AccountLogGetAuditLogsResponseSuccess = true
)

func (r AccountLogGetAuditLogsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParams struct {
	// Filters actions based on a given timestamp in the format yyyy-mm-dd, returning
	// only logs that occurred on and before the specified date.
	Before param.Field[time.Time] `query:"before,required" format:"date"`
	// Filters actions based on a given timestamp in the format yyyy-mm-dd, returning
	// only logs that occurred on and after the specified date.
	Since param.Field[time.Time] `query:"since,required" format:"date"`
	// Filters by the account name.
	AccountName param.Field[string] `query:"account_name"`
	// Whether the action was successful or not.
	ActionResult param.Field[AccountLogGetAuditLogsParamsActionResult] `query:"action_result"`
	// Filters by the action type.
	ActionType param.Field[AccountLogGetAuditLogsParamsActionType] `query:"action_type"`
	// Filters by the actor context.
	ActorContext param.Field[AccountLogGetAuditLogsParamsActorContext] `query:"actor_context"`
	// Filters by the actor's email address.
	ActorEmail param.Field[string] `query:"actor_email" format:"email"`
	// Filters by the actor ID. This can be either the Account ID or User ID.
	ActorID param.Field[string] `query:"actor_id"`
	// The IP address where the action was initiated.
	ActorIPAddress param.Field[string] `query:"actor_ip_address"`
	// Filters by the API token ID when the actor context is an api_token or oauth.
	ActorTokenID param.Field[string] `query:"actor_token_id"`
	// Filters by the API token name when the actor context is an api_token or oauth.
	ActorTokenName param.Field[string] `query:"actor_token_name"`
	// Filters by the actor type.
	ActorType param.Field[AccountLogGetAuditLogsParamsActorType] `query:"actor_type"`
	// Finds a specific log by its ID.
	AuditLogID param.Field[string] `query:"audit_log_id"`
	// The cursor is an opaque token used to paginate through large sets of records. It
	// indicates the position from which to continue when requesting the next set of
	// records. A valid cursor value can be obtained from the cursor object in the
	// result_info structure of a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Sets sorting order.
	Direction param.Field[AccountLogGetAuditLogsParamsDirection] `query:"direction"`
	// The number limits the objects to return. The cursor attribute may be used to
	// iterate over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// Filters by the response CF Ray ID.
	RawCfRayID param.Field[string] `query:"raw_cf_ray_id"`
	// The HTTP method for the API call.
	RawMethod param.Field[string] `query:"raw_method"`
	// The response status code that was returned.
	RawStatusCode param.Field[int64] `query:"raw_status_code"`
	// Filters by the request URI.
	RawUri param.Field[string] `query:"raw_uri"`
	// Filters by the resource ID.
	ResourceID param.Field[string] `query:"resource_id"`
	// Filters audit logs by the Cloudflare product associated with the changed
	// resource.
	ResourceProduct param.Field[string] `query:"resource_product"`
	// Filters by the resource scope, specifying whether the resource is associated
	// with an user, an account, or a zone.
	ResourceScope param.Field[AccountLogGetAuditLogsParamsResourceScope] `query:"resource_scope"`
	// Filters audit logs based on the unique type of resource changed by the action.
	ResourceType param.Field[string] `query:"resource_type"`
	// Filters by the zone ID.
	ZoneID param.Field[string] `query:"zone_id"`
	// Filters by the zone name associated with the change.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [AccountLogGetAuditLogsParams]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether the action was successful or not.
type AccountLogGetAuditLogsParamsActionResult string

const (
	AccountLogGetAuditLogsParamsActionResultSuccess AccountLogGetAuditLogsParamsActionResult = "success"
	AccountLogGetAuditLogsParamsActionResultFailure AccountLogGetAuditLogsParamsActionResult = "failure"
)

func (r AccountLogGetAuditLogsParamsActionResult) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActionResultSuccess, AccountLogGetAuditLogsParamsActionResultFailure:
		return true
	}
	return false
}

// Filters by the action type.
type AccountLogGetAuditLogsParamsActionType string

const (
	AccountLogGetAuditLogsParamsActionTypeCreate AccountLogGetAuditLogsParamsActionType = "create"
	AccountLogGetAuditLogsParamsActionTypeDelete AccountLogGetAuditLogsParamsActionType = "delete"
	AccountLogGetAuditLogsParamsActionTypeView   AccountLogGetAuditLogsParamsActionType = "view"
	AccountLogGetAuditLogsParamsActionTypeUpdate AccountLogGetAuditLogsParamsActionType = "update"
)

func (r AccountLogGetAuditLogsParamsActionType) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActionTypeCreate, AccountLogGetAuditLogsParamsActionTypeDelete, AccountLogGetAuditLogsParamsActionTypeView, AccountLogGetAuditLogsParamsActionTypeUpdate:
		return true
	}
	return false
}

// Filters by the actor context.
type AccountLogGetAuditLogsParamsActorContext string

const (
	AccountLogGetAuditLogsParamsActorContextAPIKey      AccountLogGetAuditLogsParamsActorContext = "api_key"
	AccountLogGetAuditLogsParamsActorContextAPIToken    AccountLogGetAuditLogsParamsActorContext = "api_token"
	AccountLogGetAuditLogsParamsActorContextDash        AccountLogGetAuditLogsParamsActorContext = "dash"
	AccountLogGetAuditLogsParamsActorContextOAuth       AccountLogGetAuditLogsParamsActorContext = "oauth"
	AccountLogGetAuditLogsParamsActorContextOriginCaKey AccountLogGetAuditLogsParamsActorContext = "origin_ca_key"
)

func (r AccountLogGetAuditLogsParamsActorContext) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActorContextAPIKey, AccountLogGetAuditLogsParamsActorContextAPIToken, AccountLogGetAuditLogsParamsActorContextDash, AccountLogGetAuditLogsParamsActorContextOAuth, AccountLogGetAuditLogsParamsActorContextOriginCaKey:
		return true
	}
	return false
}

// Filters by the actor type.
type AccountLogGetAuditLogsParamsActorType string

const (
	AccountLogGetAuditLogsParamsActorTypeCloudflareAdmin AccountLogGetAuditLogsParamsActorType = "cloudflare_admin"
	AccountLogGetAuditLogsParamsActorTypeAccount         AccountLogGetAuditLogsParamsActorType = "account"
	AccountLogGetAuditLogsParamsActorTypeUser            AccountLogGetAuditLogsParamsActorType = "user"
)

func (r AccountLogGetAuditLogsParamsActorType) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActorTypeCloudflareAdmin, AccountLogGetAuditLogsParamsActorTypeAccount, AccountLogGetAuditLogsParamsActorTypeUser:
		return true
	}
	return false
}

// Sets sorting order.
type AccountLogGetAuditLogsParamsDirection string

const (
	AccountLogGetAuditLogsParamsDirectionDesc AccountLogGetAuditLogsParamsDirection = "desc"
	AccountLogGetAuditLogsParamsDirectionAsc  AccountLogGetAuditLogsParamsDirection = "asc"
)

func (r AccountLogGetAuditLogsParamsDirection) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsDirectionDesc, AccountLogGetAuditLogsParamsDirectionAsc:
		return true
	}
	return false
}

// Filters by the resource scope, specifying whether the resource is associated
// with an user, an account, or a zone.
type AccountLogGetAuditLogsParamsResourceScope string

const (
	AccountLogGetAuditLogsParamsResourceScopeAccounts AccountLogGetAuditLogsParamsResourceScope = "accounts"
	AccountLogGetAuditLogsParamsResourceScopeUser     AccountLogGetAuditLogsParamsResourceScope = "user"
	AccountLogGetAuditLogsParamsResourceScopeZones    AccountLogGetAuditLogsParamsResourceScope = "zones"
)

func (r AccountLogGetAuditLogsParamsResourceScope) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsResourceScopeAccounts, AccountLogGetAuditLogsParamsResourceScopeUser, AccountLogGetAuditLogsParamsResourceScopeZones:
		return true
	}
	return false
}
