// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountRulesetRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetRuleService] method instead.
type AccountRulesetRuleService struct {
	Options []option.RequestOption
}

// NewAccountRulesetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetRuleService(opts ...option.RequestOption) (r *AccountRulesetRuleService) {
	r = &AccountRulesetRuleService{}
	r.Options = opts
	return
}

// Adds a new rule to an account ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset by default.
func (r *AccountRulesetRuleService) New(ctx context.Context, accountID string, rulesetID string, body AccountRulesetRuleNewParams, opts ...option.RequestOption) (res *AccountRulesetRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing rule in an account ruleset.
func (r *AccountRulesetRuleService) Update(ctx context.Context, accountID string, rulesetID string, ruleID string, body AccountRulesetRuleUpdateParams, opts ...option.RequestOption) (res *AccountRulesetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from an account ruleset.
func (r *AccountRulesetRuleService) Delete(ctx context.Context, accountID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *AccountRulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountRulesetRuleNewResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetRuleNewResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetRuleNewResponseSuccess `json:"success,required"`
	JSON    accountRulesetRuleNewResponseJSON    `json:"-"`
}

// accountRulesetRuleNewResponseJSON contains the JSON metadata for the struct
// [AccountRulesetRuleNewResponse]
type accountRulesetRuleNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetRuleNewResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleNewResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleNewResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleNewResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetRuleNewResponseMessage]
type accountRulesetRuleNewResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetRuleNewResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    accountRulesetRuleNewResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleNewResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountRulesetRuleNewResponseMessagesSource]
type accountRulesetRuleNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetRuleNewResponseSuccess bool

const (
	AccountRulesetRuleNewResponseSuccessTrue AccountRulesetRuleNewResponseSuccess = true
)

func (r AccountRulesetRuleNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetRuleNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetRuleUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetRuleUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetRuleUpdateResponseSuccess `json:"success,required"`
	JSON    accountRulesetRuleUpdateResponseJSON    `json:"-"`
}

// accountRulesetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRulesetRuleUpdateResponse]
type accountRulesetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetRuleUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleUpdateResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleUpdateResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetRuleUpdateResponseMessage]
type accountRulesetRuleUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetRuleUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetRuleUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleUpdateResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetRuleUpdateResponseMessagesSource]
type accountRulesetRuleUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetRuleUpdateResponseSuccess bool

const (
	AccountRulesetRuleUpdateResponseSuccessTrue AccountRulesetRuleUpdateResponseSuccess = true
)

func (r AccountRulesetRuleUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetRuleUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetRuleDeleteResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetRuleDeleteResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetRuleDeleteResponseSuccess `json:"success,required"`
	JSON    accountRulesetRuleDeleteResponseJSON    `json:"-"`
}

// accountRulesetRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRulesetRuleDeleteResponse]
type accountRulesetRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetRuleDeleteResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetRuleDeleteResponseMessagesSource `json:"source"`
	JSON   accountRulesetRuleDeleteResponseMessageJSON    `json:"-"`
}

// accountRulesetRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetRuleDeleteResponseMessage]
type accountRulesetRuleDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetRuleDeleteResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer,required"`
	JSON    accountRulesetRuleDeleteResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetRuleDeleteResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRulesetRuleDeleteResponseMessagesSource]
type accountRulesetRuleDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetRuleDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetRuleDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetRuleDeleteResponseSuccess bool

const (
	AccountRulesetRuleDeleteResponseSuccessTrue AccountRulesetRuleDeleteResponseSuccess = true
)

func (r AccountRulesetRuleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetRuleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetRuleNewParams struct {
	Body AccountRulesetRuleNewParamsBodyUnion `json:"body,required"`
}

func (r AccountRulesetRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRulesetRuleNewParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetRuleNewParamsBodyAction] `json:"action"`
	ActionParameters param.Field[interface{}]                           `json:"action_parameters"`
	Categories       param.Field[interface{}]                           `json:"categories"`
	Description      param.Field[interface{}]                           `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string]      `json:"expression"`
	Logging    param.Field[interface{}] `json:"logging"`
	Position   param.Field[interface{}] `json:"position"`
	Ratelimit  param.Field[interface{}] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBody) implementsAccountRulesetRuleNewParamsBodyUnion() {}

// Satisfied by [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject],
// [AccountRulesetRuleNewParamsBodyObject], [AccountRulesetRuleNewParamsBody].
type AccountRulesetRuleNewParamsBodyUnion interface {
	implementsAccountRulesetRuleNewParamsBodyUnion()
}

type AccountRulesetRuleNewParamsBodyObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetRuleNewParamsBodyObjectAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetRuleNewParamsBodyObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[AccountRulesetRuleNewParamsBodyObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging  param.Field[AccountRulesetRuleNewParamsBodyObjectLogging]       `json:"logging"`
	Position param.Field[AccountRulesetRuleNewParamsBodyObjectPositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[AccountRulesetRuleNewParamsBodyObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBodyObject) implementsAccountRulesetRuleNewParamsBodyUnion() {}

// The action to perform when the rule matches.
type AccountRulesetRuleNewParamsBodyObjectAction string

const (
	AccountRulesetRuleNewParamsBodyObjectActionBlock AccountRulesetRuleNewParamsBodyObjectAction = "block"
)

func (r AccountRulesetRuleNewParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case AccountRulesetRuleNewParamsBodyObjectActionBlock:
		return true
	}
	return false
}

type AccountRulesetRuleNewParamsBodyObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetRuleNewParamsBodyObjectActionParametersResponse] `json:"response"`
}

func (r AccountRulesetRuleNewParamsBodyObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetRuleNewParamsBodyObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetRuleNewParamsBodyObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type AccountRulesetRuleNewParamsBodyObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r AccountRulesetRuleNewParamsBodyObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetRuleNewParamsBodyObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetRuleNewParamsBodyObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetRuleNewParamsBodyObjectPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r AccountRulesetRuleNewParamsBodyObjectPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBodyObjectPosition) implementsAccountRulesetRuleNewParamsBodyObjectPositionUnion() {
}

// Satisfied by [AccountRulesetRuleNewParamsBodyObjectPositionBefore],
// [AccountRulesetRuleNewParamsBodyObjectPositionAfter],
// [AccountRulesetRuleNewParamsBodyObjectPositionIndex],
// [AccountRulesetRuleNewParamsBodyObjectPosition].
type AccountRulesetRuleNewParamsBodyObjectPositionUnion interface {
	implementsAccountRulesetRuleNewParamsBodyObjectPositionUnion()
}

type AccountRulesetRuleNewParamsBodyObjectPositionBefore struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionBefore) implementsAccountRulesetRuleNewParamsBodyObjectPositionUnion() {
}

type AccountRulesetRuleNewParamsBodyObjectPositionAfter struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionAfter) implementsAccountRulesetRuleNewParamsBodyObjectPositionUnion() {
}

type AccountRulesetRuleNewParamsBodyObjectPositionIndex struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleNewParamsBodyObjectPositionIndex) implementsAccountRulesetRuleNewParamsBodyObjectPositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type AccountRulesetRuleNewParamsBodyObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r AccountRulesetRuleNewParamsBodyObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod int64

const (
	AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod10   AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod = 10
	AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod60   AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod = 60
	AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod600  AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod = 600
	AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod3600 AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod = 3600
)

func (r AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod10, AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod60, AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod600, AccountRulesetRuleNewParamsBodyObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type AccountRulesetRuleNewParamsBodyAction string

const (
	AccountRulesetRuleNewParamsBodyActionBlock                AccountRulesetRuleNewParamsBodyAction = "block"
	AccountRulesetRuleNewParamsBodyActionChallenge            AccountRulesetRuleNewParamsBodyAction = "challenge"
	AccountRulesetRuleNewParamsBodyActionCompressResponse     AccountRulesetRuleNewParamsBodyAction = "compress_response"
	AccountRulesetRuleNewParamsBodyActionExecute              AccountRulesetRuleNewParamsBodyAction = "execute"
	AccountRulesetRuleNewParamsBodyActionJsChallenge          AccountRulesetRuleNewParamsBodyAction = "js_challenge"
	AccountRulesetRuleNewParamsBodyActionLog                  AccountRulesetRuleNewParamsBodyAction = "log"
	AccountRulesetRuleNewParamsBodyActionManagedChallenge     AccountRulesetRuleNewParamsBodyAction = "managed_challenge"
	AccountRulesetRuleNewParamsBodyActionRedirect             AccountRulesetRuleNewParamsBodyAction = "redirect"
	AccountRulesetRuleNewParamsBodyActionRewrite              AccountRulesetRuleNewParamsBodyAction = "rewrite"
	AccountRulesetRuleNewParamsBodyActionRoute                AccountRulesetRuleNewParamsBodyAction = "route"
	AccountRulesetRuleNewParamsBodyActionScore                AccountRulesetRuleNewParamsBodyAction = "score"
	AccountRulesetRuleNewParamsBodyActionServeError           AccountRulesetRuleNewParamsBodyAction = "serve_error"
	AccountRulesetRuleNewParamsBodyActionSetConfig            AccountRulesetRuleNewParamsBodyAction = "set_config"
	AccountRulesetRuleNewParamsBodyActionSkip                 AccountRulesetRuleNewParamsBodyAction = "skip"
	AccountRulesetRuleNewParamsBodyActionSetCacheSettings     AccountRulesetRuleNewParamsBodyAction = "set_cache_settings"
	AccountRulesetRuleNewParamsBodyActionLogCustomField       AccountRulesetRuleNewParamsBodyAction = "log_custom_field"
	AccountRulesetRuleNewParamsBodyActionDdosDynamic          AccountRulesetRuleNewParamsBodyAction = "ddos_dynamic"
	AccountRulesetRuleNewParamsBodyActionForceConnectionClose AccountRulesetRuleNewParamsBodyAction = "force_connection_close"
)

func (r AccountRulesetRuleNewParamsBodyAction) IsKnown() bool {
	switch r {
	case AccountRulesetRuleNewParamsBodyActionBlock, AccountRulesetRuleNewParamsBodyActionChallenge, AccountRulesetRuleNewParamsBodyActionCompressResponse, AccountRulesetRuleNewParamsBodyActionExecute, AccountRulesetRuleNewParamsBodyActionJsChallenge, AccountRulesetRuleNewParamsBodyActionLog, AccountRulesetRuleNewParamsBodyActionManagedChallenge, AccountRulesetRuleNewParamsBodyActionRedirect, AccountRulesetRuleNewParamsBodyActionRewrite, AccountRulesetRuleNewParamsBodyActionRoute, AccountRulesetRuleNewParamsBodyActionScore, AccountRulesetRuleNewParamsBodyActionServeError, AccountRulesetRuleNewParamsBodyActionSetConfig, AccountRulesetRuleNewParamsBodyActionSkip, AccountRulesetRuleNewParamsBodyActionSetCacheSettings, AccountRulesetRuleNewParamsBodyActionLogCustomField, AccountRulesetRuleNewParamsBodyActionDdosDynamic, AccountRulesetRuleNewParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}

type AccountRulesetRuleUpdateParams struct {
	Body AccountRulesetRuleUpdateParamsBodyUnion `json:"body,required"`
}

func (r AccountRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountRulesetRuleUpdateParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetRuleUpdateParamsBodyAction] `json:"action"`
	ActionParameters param.Field[interface{}]                              `json:"action_parameters"`
	Categories       param.Field[interface{}]                              `json:"categories"`
	Description      param.Field[interface{}]                              `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string]      `json:"expression"`
	Logging    param.Field[interface{}] `json:"logging"`
	Position   param.Field[interface{}] `json:"position"`
	Ratelimit  param.Field[interface{}] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBody) implementsAccountRulesetRuleUpdateParamsBodyUnion() {}

// Satisfied by [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBodyObject],
// [AccountRulesetRuleUpdateParamsBody].
type AccountRulesetRuleUpdateParamsBodyUnion interface {
	implementsAccountRulesetRuleUpdateParamsBodyUnion()
}

type AccountRulesetRuleUpdateParamsBodyObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetRuleUpdateParamsBodyObjectAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetRuleUpdateParamsBodyObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[AccountRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging  param.Field[AccountRulesetRuleUpdateParamsBodyObjectLogging]       `json:"logging"`
	Position param.Field[AccountRulesetRuleUpdateParamsBodyObjectPositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[AccountRulesetRuleUpdateParamsBodyObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBodyObject) implementsAccountRulesetRuleUpdateParamsBodyUnion() {
}

// The action to perform when the rule matches.
type AccountRulesetRuleUpdateParamsBodyObjectAction string

const (
	AccountRulesetRuleUpdateParamsBodyObjectActionBlock AccountRulesetRuleUpdateParamsBodyObjectAction = "block"
)

func (r AccountRulesetRuleUpdateParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case AccountRulesetRuleUpdateParamsBodyObjectActionBlock:
		return true
	}
	return false
}

type AccountRulesetRuleUpdateParamsBodyObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetRuleUpdateParamsBodyObjectActionParametersResponse] `json:"response"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetRuleUpdateParamsBodyObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type AccountRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetRuleUpdateParamsBodyObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetRuleUpdateParamsBodyObjectPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPosition) implementsAccountRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

// Satisfied by [AccountRulesetRuleUpdateParamsBodyObjectPositionBefore],
// [AccountRulesetRuleUpdateParamsBodyObjectPositionAfter],
// [AccountRulesetRuleUpdateParamsBodyObjectPositionIndex],
// [AccountRulesetRuleUpdateParamsBodyObjectPosition].
type AccountRulesetRuleUpdateParamsBodyObjectPositionUnion interface {
	implementsAccountRulesetRuleUpdateParamsBodyObjectPositionUnion()
}

type AccountRulesetRuleUpdateParamsBodyObjectPositionBefore struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionBefore) implementsAccountRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

type AccountRulesetRuleUpdateParamsBodyObjectPositionAfter struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionAfter) implementsAccountRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

type AccountRulesetRuleUpdateParamsBodyObjectPositionIndex struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetRuleUpdateParamsBodyObjectPositionIndex) implementsAccountRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type AccountRulesetRuleUpdateParamsBodyObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod] `json:"period,required"`
	// Defines when the ratelimit counter should be incremented. It is optional and
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Defines if ratelimit counting is only done when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// The response header name provided by the origin which should contain the score
	// to increment ratelimit counter on.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r AccountRulesetRuleUpdateParamsBodyObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod int64

const (
	AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod10   AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 10
	AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod60   AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 60
	AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod600  AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 600
	AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod3600 AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 3600
)

func (r AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod10, AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod60, AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod600, AccountRulesetRuleUpdateParamsBodyObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type AccountRulesetRuleUpdateParamsBodyAction string

const (
	AccountRulesetRuleUpdateParamsBodyActionBlock                AccountRulesetRuleUpdateParamsBodyAction = "block"
	AccountRulesetRuleUpdateParamsBodyActionChallenge            AccountRulesetRuleUpdateParamsBodyAction = "challenge"
	AccountRulesetRuleUpdateParamsBodyActionCompressResponse     AccountRulesetRuleUpdateParamsBodyAction = "compress_response"
	AccountRulesetRuleUpdateParamsBodyActionExecute              AccountRulesetRuleUpdateParamsBodyAction = "execute"
	AccountRulesetRuleUpdateParamsBodyActionJsChallenge          AccountRulesetRuleUpdateParamsBodyAction = "js_challenge"
	AccountRulesetRuleUpdateParamsBodyActionLog                  AccountRulesetRuleUpdateParamsBodyAction = "log"
	AccountRulesetRuleUpdateParamsBodyActionManagedChallenge     AccountRulesetRuleUpdateParamsBodyAction = "managed_challenge"
	AccountRulesetRuleUpdateParamsBodyActionRedirect             AccountRulesetRuleUpdateParamsBodyAction = "redirect"
	AccountRulesetRuleUpdateParamsBodyActionRewrite              AccountRulesetRuleUpdateParamsBodyAction = "rewrite"
	AccountRulesetRuleUpdateParamsBodyActionRoute                AccountRulesetRuleUpdateParamsBodyAction = "route"
	AccountRulesetRuleUpdateParamsBodyActionScore                AccountRulesetRuleUpdateParamsBodyAction = "score"
	AccountRulesetRuleUpdateParamsBodyActionServeError           AccountRulesetRuleUpdateParamsBodyAction = "serve_error"
	AccountRulesetRuleUpdateParamsBodyActionSetConfig            AccountRulesetRuleUpdateParamsBodyAction = "set_config"
	AccountRulesetRuleUpdateParamsBodyActionSkip                 AccountRulesetRuleUpdateParamsBodyAction = "skip"
	AccountRulesetRuleUpdateParamsBodyActionSetCacheSettings     AccountRulesetRuleUpdateParamsBodyAction = "set_cache_settings"
	AccountRulesetRuleUpdateParamsBodyActionLogCustomField       AccountRulesetRuleUpdateParamsBodyAction = "log_custom_field"
	AccountRulesetRuleUpdateParamsBodyActionDdosDynamic          AccountRulesetRuleUpdateParamsBodyAction = "ddos_dynamic"
	AccountRulesetRuleUpdateParamsBodyActionForceConnectionClose AccountRulesetRuleUpdateParamsBodyAction = "force_connection_close"
)

func (r AccountRulesetRuleUpdateParamsBodyAction) IsKnown() bool {
	switch r {
	case AccountRulesetRuleUpdateParamsBodyActionBlock, AccountRulesetRuleUpdateParamsBodyActionChallenge, AccountRulesetRuleUpdateParamsBodyActionCompressResponse, AccountRulesetRuleUpdateParamsBodyActionExecute, AccountRulesetRuleUpdateParamsBodyActionJsChallenge, AccountRulesetRuleUpdateParamsBodyActionLog, AccountRulesetRuleUpdateParamsBodyActionManagedChallenge, AccountRulesetRuleUpdateParamsBodyActionRedirect, AccountRulesetRuleUpdateParamsBodyActionRewrite, AccountRulesetRuleUpdateParamsBodyActionRoute, AccountRulesetRuleUpdateParamsBodyActionScore, AccountRulesetRuleUpdateParamsBodyActionServeError, AccountRulesetRuleUpdateParamsBodyActionSetConfig, AccountRulesetRuleUpdateParamsBodyActionSkip, AccountRulesetRuleUpdateParamsBodyActionSetCacheSettings, AccountRulesetRuleUpdateParamsBodyActionLogCustomField, AccountRulesetRuleUpdateParamsBodyActionDdosDynamic, AccountRulesetRuleUpdateParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}
