// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
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
	// Limits the returned results to logs older than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Before param.Field[time.Time] `query:"before,required" format:"date"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Since          param.Field[time.Time]                                  `query:"since,required" format:"date"`
	AccountName    param.Field[AccountLogGetAuditLogsParamsAccountName]    `query:"account_name"`
	ActionResult   param.Field[AccountLogGetAuditLogsParamsActionResult]   `query:"action_result"`
	ActionType     param.Field[AccountLogGetAuditLogsParamsActionType]     `query:"action_type"`
	ActorContext   param.Field[AccountLogGetAuditLogsParamsActorContext]   `query:"actor_context"`
	ActorEmail     param.Field[AccountLogGetAuditLogsParamsActorEmail]     `query:"actor_email"`
	ActorID        param.Field[AccountLogGetAuditLogsParamsActorID]        `query:"actor_id"`
	ActorIPAddress param.Field[AccountLogGetAuditLogsParamsActorIPAddress] `query:"actor_ip_address"`
	ActorTokenID   param.Field[AccountLogGetAuditLogsParamsActorTokenID]   `query:"actor_token_id"`
	ActorTokenName param.Field[AccountLogGetAuditLogsParamsActorTokenName] `query:"actor_token_name"`
	ActorType      param.Field[AccountLogGetAuditLogsParamsActorType]      `query:"actor_type"`
	AuditLogID     param.Field[AccountLogGetAuditLogsParamsAuditLogID]     `query:"audit_log_id"`
	// The cursor is an opaque token used to paginate through large sets of records. It
	// indicates the position from which to continue when requesting the next set of
	// records. A valid cursor value can be obtained from the cursor object in the
	// result_info structure of a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Sets sorting order.
	Direction param.Field[AccountLogGetAuditLogsParamsDirection] `query:"direction"`
	// The number limits the objects to return. The cursor attribute may be used to
	// iterate over the next batch of objects if there are more than the limit.
	Limit           param.Field[float64]                                     `query:"limit"`
	RawCfRayID      param.Field[AccountLogGetAuditLogsParamsRawCfRayID]      `query:"raw_cf_ray_id"`
	RawMethod       param.Field[AccountLogGetAuditLogsParamsRawMethod]       `query:"raw_method"`
	RawStatusCode   param.Field[AccountLogGetAuditLogsParamsRawStatusCode]   `query:"raw_status_code"`
	RawUri          param.Field[AccountLogGetAuditLogsParamsRawUri]          `query:"raw_uri"`
	ResourceID      param.Field[AccountLogGetAuditLogsParamsResourceID]      `query:"resource_id"`
	ResourceProduct param.Field[AccountLogGetAuditLogsParamsResourceProduct] `query:"resource_product"`
	ResourceScope   param.Field[AccountLogGetAuditLogsParamsResourceScope]   `query:"resource_scope"`
	ResourceType    param.Field[AccountLogGetAuditLogsParamsResourceType]    `query:"resource_type"`
	ZoneID          param.Field[AccountLogGetAuditLogsParamsZoneID]          `query:"zone_id"`
	ZoneName        param.Field[AccountLogGetAuditLogsParamsZoneName]        `query:"zone_name"`
}

// URLQuery serializes [AccountLogGetAuditLogsParams]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsAccountName struct {
	// Filters out audit logs by the account name.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsAccountName]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsAccountName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActionResult struct {
	// Filters out audit logs by whether the action was successful or not.
	Not param.Field[[]AccountLogGetAuditLogsParamsActionResultNot] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActionResult]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsActionResult) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActionResultNot string

const (
	AccountLogGetAuditLogsParamsActionResultNotSuccess AccountLogGetAuditLogsParamsActionResultNot = "success"
	AccountLogGetAuditLogsParamsActionResultNotFailure AccountLogGetAuditLogsParamsActionResultNot = "failure"
)

func (r AccountLogGetAuditLogsParamsActionResultNot) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActionResultNotSuccess, AccountLogGetAuditLogsParamsActionResultNotFailure:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParamsActionType struct {
	// Filters out audit logs by the action type.
	Not param.Field[[]AccountLogGetAuditLogsParamsActionTypeNot] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActionType]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsActionType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActionTypeNot string

const (
	AccountLogGetAuditLogsParamsActionTypeNotCreate AccountLogGetAuditLogsParamsActionTypeNot = "create"
	AccountLogGetAuditLogsParamsActionTypeNotDelete AccountLogGetAuditLogsParamsActionTypeNot = "delete"
	AccountLogGetAuditLogsParamsActionTypeNotView   AccountLogGetAuditLogsParamsActionTypeNot = "view"
	AccountLogGetAuditLogsParamsActionTypeNotUpdate AccountLogGetAuditLogsParamsActionTypeNot = "update"
)

func (r AccountLogGetAuditLogsParamsActionTypeNot) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActionTypeNotCreate, AccountLogGetAuditLogsParamsActionTypeNotDelete, AccountLogGetAuditLogsParamsActionTypeNotView, AccountLogGetAuditLogsParamsActionTypeNotUpdate:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParamsActorContext struct {
	// Filters out audit logs by the actor context.
	Not param.Field[[]AccountLogGetAuditLogsParamsActorContextNot] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorContext]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorContext) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorContextNot string

const (
	AccountLogGetAuditLogsParamsActorContextNotAPIKey      AccountLogGetAuditLogsParamsActorContextNot = "api_key"
	AccountLogGetAuditLogsParamsActorContextNotAPIToken    AccountLogGetAuditLogsParamsActorContextNot = "api_token"
	AccountLogGetAuditLogsParamsActorContextNotDash        AccountLogGetAuditLogsParamsActorContextNot = "dash"
	AccountLogGetAuditLogsParamsActorContextNotOAuth       AccountLogGetAuditLogsParamsActorContextNot = "oauth"
	AccountLogGetAuditLogsParamsActorContextNotOriginCaKey AccountLogGetAuditLogsParamsActorContextNot = "origin_ca_key"
)

func (r AccountLogGetAuditLogsParamsActorContextNot) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActorContextNotAPIKey, AccountLogGetAuditLogsParamsActorContextNotAPIToken, AccountLogGetAuditLogsParamsActorContextNotDash, AccountLogGetAuditLogsParamsActorContextNotOAuth, AccountLogGetAuditLogsParamsActorContextNotOriginCaKey:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParamsActorEmail struct {
	// Filters out audit logs by the actor's email address.
	Not param.Field[[]string] `query:"not" format:"email"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorEmail]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorEmail) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorID struct {
	// Filters out audit logs by the actor ID. This can be either the Account ID or
	// User ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorID]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParamsActorID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorIPAddress struct {
	// Filters out audit logs IP address where the action was initiated.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorIPAddress]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorIPAddress) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorTokenID struct {
	// Filters out audit logs by the API token ID when the actor context is an
	// api_token or oauth.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorTokenID]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorTokenID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorTokenName struct {
	// Filters out audit logs by the API token name when the actor context is an
	// api_token or oauth.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorTokenName]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorTokenName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorType struct {
	// Filters out audit logs by the actor type.
	Not param.Field[[]AccountLogGetAuditLogsParamsActorTypeNot] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsActorType]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsActorType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsActorTypeNot string

const (
	AccountLogGetAuditLogsParamsActorTypeNotAccount         AccountLogGetAuditLogsParamsActorTypeNot = "account"
	AccountLogGetAuditLogsParamsActorTypeNotCloudflareAdmin AccountLogGetAuditLogsParamsActorTypeNot = "cloudflare_admin"
	AccountLogGetAuditLogsParamsActorTypeNotSystem          AccountLogGetAuditLogsParamsActorTypeNot = "system"
	AccountLogGetAuditLogsParamsActorTypeNotUser            AccountLogGetAuditLogsParamsActorTypeNot = "user"
)

func (r AccountLogGetAuditLogsParamsActorTypeNot) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsActorTypeNotAccount, AccountLogGetAuditLogsParamsActorTypeNotCloudflareAdmin, AccountLogGetAuditLogsParamsActorTypeNotSystem, AccountLogGetAuditLogsParamsActorTypeNotUser:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParamsAuditLogID struct {
	// Filters out audit logs by their IDs.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsAuditLogID]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsAuditLogID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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

type AccountLogGetAuditLogsParamsRawCfRayID struct {
	// Filters out audit logs by the response CF Ray ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsRawCfRayID]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsRawCfRayID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsRawMethod struct {
	// Filters out audit logs by the HTTP method for the API call.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsRawMethod]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsRawMethod) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsRawStatusCode struct {
	// Filters out audit logs by the response status code that was returned.
	Not param.Field[[]int64] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsRawStatusCode]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsRawStatusCode) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsRawUri struct {
	// Filters out audit logs by the request URI.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsRawUri]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParamsRawUri) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsResourceID struct {
	// Filters out audit logs by the resource ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsResourceID]'s query parameters
// as `url.Values`.
func (r AccountLogGetAuditLogsParamsResourceID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsResourceProduct struct {
	// Filters out audit logs by the Cloudflare product associated with the changed
	// resource.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsResourceProduct]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsResourceProduct) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsResourceScope struct {
	// Filters out audit logs by the resource scope, specifying whether the resource is
	// associated with an user, an account, or a zone.
	Not param.Field[[]AccountLogGetAuditLogsParamsResourceScopeNot] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsResourceScope]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsResourceScope) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsResourceScopeNot string

const (
	AccountLogGetAuditLogsParamsResourceScopeNotAccounts AccountLogGetAuditLogsParamsResourceScopeNot = "accounts"
	AccountLogGetAuditLogsParamsResourceScopeNotUser     AccountLogGetAuditLogsParamsResourceScopeNot = "user"
	AccountLogGetAuditLogsParamsResourceScopeNotZones    AccountLogGetAuditLogsParamsResourceScopeNot = "zones"
)

func (r AccountLogGetAuditLogsParamsResourceScopeNot) IsKnown() bool {
	switch r {
	case AccountLogGetAuditLogsParamsResourceScopeNotAccounts, AccountLogGetAuditLogsParamsResourceScopeNotUser, AccountLogGetAuditLogsParamsResourceScopeNotZones:
		return true
	}
	return false
}

type AccountLogGetAuditLogsParamsResourceType struct {
	// Filters out audit logs based on the unique type of resource changed by the
	// action.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsResourceType]'s query
// parameters as `url.Values`.
func (r AccountLogGetAuditLogsParamsResourceType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsZoneID struct {
	// Filters out audit logs by the zone ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsZoneID]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParamsZoneID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountLogGetAuditLogsParamsZoneName struct {
	// Filters out audit logs by the zone name associated with the change.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [AccountLogGetAuditLogsParamsZoneName]'s query parameters as
// `url.Values`.
func (r AccountLogGetAuditLogsParamsZoneName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
