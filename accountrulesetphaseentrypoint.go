// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountRulesetPhaseEntrypointService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetPhaseEntrypointService] method instead.
type AccountRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *AccountRulesetPhaseEntrypointVersionService
}

// NewAccountRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointService) {
	r = &AccountRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewAccountRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches the latest version of the account entry point ruleset for a given phase.
func (r *AccountRulesetPhaseEntrypointService) Get(ctx context.Context, accountID string, rulesetPhase RulesetPhase, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account entry point ruleset, creating a new version.
func (r *AccountRulesetPhaseEntrypointService) Update(ctx context.Context, accountID string, rulesetPhase RulesetPhase, body AccountRulesetPhaseEntrypointUpdateParams, opts ...option.RequestOption) (res *AccountRulesetPhaseEntrypointUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%v/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// The phase of the ruleset.
type RulesetPhase string

const (
	RulesetPhaseDdosL4                       RulesetPhase = "ddos_l4"
	RulesetPhaseDdosL7                       RulesetPhase = "ddos_l7"
	RulesetPhaseHTTPConfigSettings           RulesetPhase = "http_config_settings"
	RulesetPhaseHTTPCustomErrors             RulesetPhase = "http_custom_errors"
	RulesetPhaseHTTPLogCustomFields          RulesetPhase = "http_log_custom_fields"
	RulesetPhaseHTTPRatelimit                RulesetPhase = "http_ratelimit"
	RulesetPhaseHTTPRequestCacheSettings     RulesetPhase = "http_request_cache_settings"
	RulesetPhaseHTTPRequestDynamicRedirect   RulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseHTTPRequestFirewallCustom    RulesetPhase = "http_request_firewall_custom"
	RulesetPhaseHTTPRequestFirewallManaged   RulesetPhase = "http_request_firewall_managed"
	RulesetPhaseHTTPRequestLateTransform     RulesetPhase = "http_request_late_transform"
	RulesetPhaseHTTPRequestOrigin            RulesetPhase = "http_request_origin"
	RulesetPhaseHTTPRequestRedirect          RulesetPhase = "http_request_redirect"
	RulesetPhaseHTTPRequestSanitize          RulesetPhase = "http_request_sanitize"
	RulesetPhaseHTTPRequestSbfm              RulesetPhase = "http_request_sbfm"
	RulesetPhaseHTTPRequestTransform         RulesetPhase = "http_request_transform"
	RulesetPhaseHTTPResponseCompression      RulesetPhase = "http_response_compression"
	RulesetPhaseHTTPResponseFirewallManaged  RulesetPhase = "http_response_firewall_managed"
	RulesetPhaseHTTPResponseHeadersTransform RulesetPhase = "http_response_headers_transform"
	RulesetPhaseMagicTransit                 RulesetPhase = "magic_transit"
	RulesetPhaseMagicTransitIDsManaged       RulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseMagicTransitManaged          RulesetPhase = "magic_transit_managed"
	RulesetPhaseMagicTransitRatelimit        RulesetPhase = "magic_transit_ratelimit"
)

func (r RulesetPhase) IsKnown() bool {
	switch r {
	case RulesetPhaseDdosL4, RulesetPhaseDdosL7, RulesetPhaseHTTPConfigSettings, RulesetPhaseHTTPCustomErrors, RulesetPhaseHTTPLogCustomFields, RulesetPhaseHTTPRatelimit, RulesetPhaseHTTPRequestCacheSettings, RulesetPhaseHTTPRequestDynamicRedirect, RulesetPhaseHTTPRequestFirewallCustom, RulesetPhaseHTTPRequestFirewallManaged, RulesetPhaseHTTPRequestLateTransform, RulesetPhaseHTTPRequestOrigin, RulesetPhaseHTTPRequestRedirect, RulesetPhaseHTTPRequestSanitize, RulesetPhaseHTTPRequestSbfm, RulesetPhaseHTTPRequestTransform, RulesetPhaseHTTPResponseCompression, RulesetPhaseHTTPResponseFirewallManaged, RulesetPhaseHTTPResponseHeadersTransform, RulesetPhaseMagicTransit, RulesetPhaseMagicTransitIDsManaged, RulesetPhaseMagicTransitManaged, RulesetPhaseMagicTransitRatelimit:
		return true
	}
	return false
}

type AccountRulesetPhaseEntrypointGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointGetResponseSuccess `json:"success,required"`
	JSON    accountRulesetPhaseEntrypointGetResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointGetResponseJSON contains the JSON metadata for the
// struct [AccountRulesetPhaseEntrypointGetResponse]
type accountRulesetPhaseEntrypointGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetPhaseEntrypointGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointGetResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointGetResponseMessageJSON contains the JSON metadata
// for the struct [AccountRulesetPhaseEntrypointGetResponseMessage]
type accountRulesetPhaseEntrypointGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetPhaseEntrypointGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                     `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointGetResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountRulesetPhaseEntrypointGetResponseMessagesSource]
type accountRulesetPhaseEntrypointGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointGetResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointGetResponseSuccessTrue AccountRulesetPhaseEntrypointGetResponseSuccess = true
)

func (r AccountRulesetPhaseEntrypointGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetPhaseEntrypointUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetPhaseEntrypointUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetPhaseEntrypointUpdateResponseSuccess `json:"success,required"`
	JSON    accountRulesetPhaseEntrypointUpdateResponseJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointUpdateResponseJSON contains the JSON metadata for
// the struct [AccountRulesetPhaseEntrypointUpdateResponse]
type accountRulesetPhaseEntrypointUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetPhaseEntrypointUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetPhaseEntrypointUpdateResponseMessagesSource `json:"source"`
	JSON   accountRulesetPhaseEntrypointUpdateResponseMessageJSON    `json:"-"`
}

// accountRulesetPhaseEntrypointUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountRulesetPhaseEntrypointUpdateResponseMessage]
type accountRulesetPhaseEntrypointUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetPhaseEntrypointUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                        `json:"pointer,required"`
	JSON    accountRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccountRulesetPhaseEntrypointUpdateResponseMessagesSource]
type accountRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetPhaseEntrypointUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetPhaseEntrypointUpdateResponseSuccess bool

const (
	AccountRulesetPhaseEntrypointUpdateResponseSuccessTrue AccountRulesetPhaseEntrypointUpdateResponseSuccess = true
)

func (r AccountRulesetPhaseEntrypointUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetPhaseEntrypointUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetPhaseEntrypointUpdateParamsRuleUnion] `json:"rules"`
}

func (r AccountRulesetPhaseEntrypointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetPhaseEntrypointUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                          `json:"action_parameters"`
	Categories       param.Field[interface{}]                                          `json:"categories"`
	Description      param.Field[interface{}]                                          `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string]      `json:"expression"`
	Logging    param.Field[interface{}] `json:"logging"`
	Ratelimit  param.Field[interface{}] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRule) implementsAccountRulesetPhaseEntrypointUpdateParamsRuleUnion() {
}

// Satisfied by [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRulesObject],
// [AccountRulesetPhaseEntrypointUpdateParamsRule].
type AccountRulesetPhaseEntrypointUpdateParamsRuleUnion interface {
	implementsAccountRulesetPhaseEntrypointUpdateParamsRuleUnion()
}

type AccountRulesetPhaseEntrypointUpdateParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObject) implementsAccountRulesetPhaseEntrypointUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectAction string

const (
	AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionBlock AccountRulesetPhaseEntrypointUpdateParamsRulesObjectAction = "block"
)

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod int64

const (
	AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod10   AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 10
	AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod60   AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 60
	AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod600  AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 600
	AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod3600 AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 3600
)

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod10, AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod60, AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod600, AccountRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type AccountRulesetPhaseEntrypointUpdateParamsRulesAction string

const (
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionBlock                AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "block"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionChallenge            AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "challenge"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionCompressResponse     AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "compress_response"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionExecute              AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "execute"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionJsChallenge          AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "js_challenge"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionLog                  AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "log"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionManagedChallenge     AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "managed_challenge"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionRedirect             AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "redirect"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionRewrite              AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "rewrite"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionRoute                AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "route"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionScore                AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "score"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionServeError           AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "serve_error"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionSetConfig            AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "set_config"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionSkip                 AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "skip"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionSetCacheSettings     AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "set_cache_settings"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionLogCustomField       AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "log_custom_field"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionDdosDynamic          AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "ddos_dynamic"
	AccountRulesetPhaseEntrypointUpdateParamsRulesActionForceConnectionClose AccountRulesetPhaseEntrypointUpdateParamsRulesAction = "force_connection_close"
)

func (r AccountRulesetPhaseEntrypointUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case AccountRulesetPhaseEntrypointUpdateParamsRulesActionBlock, AccountRulesetPhaseEntrypointUpdateParamsRulesActionChallenge, AccountRulesetPhaseEntrypointUpdateParamsRulesActionCompressResponse, AccountRulesetPhaseEntrypointUpdateParamsRulesActionExecute, AccountRulesetPhaseEntrypointUpdateParamsRulesActionJsChallenge, AccountRulesetPhaseEntrypointUpdateParamsRulesActionLog, AccountRulesetPhaseEntrypointUpdateParamsRulesActionManagedChallenge, AccountRulesetPhaseEntrypointUpdateParamsRulesActionRedirect, AccountRulesetPhaseEntrypointUpdateParamsRulesActionRewrite, AccountRulesetPhaseEntrypointUpdateParamsRulesActionRoute, AccountRulesetPhaseEntrypointUpdateParamsRulesActionScore, AccountRulesetPhaseEntrypointUpdateParamsRulesActionServeError, AccountRulesetPhaseEntrypointUpdateParamsRulesActionSetConfig, AccountRulesetPhaseEntrypointUpdateParamsRulesActionSkip, AccountRulesetPhaseEntrypointUpdateParamsRulesActionSetCacheSettings, AccountRulesetPhaseEntrypointUpdateParamsRulesActionLogCustomField, AccountRulesetPhaseEntrypointUpdateParamsRulesActionDdosDynamic, AccountRulesetPhaseEntrypointUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}
