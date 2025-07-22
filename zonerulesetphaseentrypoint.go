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

// ZoneRulesetPhaseEntrypointService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseEntrypointService] method instead.
type ZoneRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *ZoneRulesetPhaseEntrypointVersionService
}

// NewZoneRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointService) {
	r = &ZoneRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewZoneRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches the latest version of the zone entry point ruleset for a given phase.
func (r *ZoneRulesetPhaseEntrypointService) Get(ctx context.Context, zoneID string, rulesetPhase RulesetPhase, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a zone entry point ruleset, creating a new version.
func (r *ZoneRulesetPhaseEntrypointService) Update(ctx context.Context, zoneID string, rulesetPhase RulesetPhase, body ZoneRulesetPhaseEntrypointUpdateParams, opts ...option.RequestOption) (res *ZoneRulesetPhaseEntrypointUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/phases/%v/entrypoint", zoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneRulesetPhaseEntrypointGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointGetResponseSuccess `json:"success,required"`
	JSON    zoneRulesetPhaseEntrypointGetResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointGetResponseJSON contains the JSON metadata for the
// struct [ZoneRulesetPhaseEntrypointGetResponse]
type zoneRulesetPhaseEntrypointGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetPhaseEntrypointGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointGetResponseMessageJSON contains the JSON metadata for
// the struct [ZoneRulesetPhaseEntrypointGetResponseMessage]
type zoneRulesetPhaseEntrypointGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetPhaseEntrypointGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                  `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointGetResponseMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointGetResponseMessagesSource]
type zoneRulesetPhaseEntrypointGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointGetResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointGetResponseSuccessTrue ZoneRulesetPhaseEntrypointGetResponseSuccess = true
)

func (r ZoneRulesetPhaseEntrypointGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetPhaseEntrypointUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetPhaseEntrypointUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetPhaseEntrypointUpdateResponseSuccess `json:"success,required"`
	JSON    zoneRulesetPhaseEntrypointUpdateResponseJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneRulesetPhaseEntrypointUpdateResponse]
type zoneRulesetPhaseEntrypointUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetPhaseEntrypointUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetPhaseEntrypointUpdateResponseMessagesSource `json:"source"`
	JSON   zoneRulesetPhaseEntrypointUpdateResponseMessageJSON    `json:"-"`
}

// zoneRulesetPhaseEntrypointUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneRulesetPhaseEntrypointUpdateResponseMessage]
type zoneRulesetPhaseEntrypointUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetPhaseEntrypointUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                     `json:"pointer,required"`
	JSON    zoneRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneRulesetPhaseEntrypointUpdateResponseMessagesSource]
type zoneRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetPhaseEntrypointUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetPhaseEntrypointUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetPhaseEntrypointUpdateResponseSuccess bool

const (
	ZoneRulesetPhaseEntrypointUpdateResponseSuccessTrue ZoneRulesetPhaseEntrypointUpdateResponseSuccess = true
)

func (r ZoneRulesetPhaseEntrypointUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetPhaseEntrypointUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseEntrypointUpdateParamsRuleUnion] `json:"rules"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetPhaseEntrypointUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                                       `json:"action_parameters"`
	Categories       param.Field[interface{}]                                       `json:"categories"`
	Description      param.Field[interface{}]                                       `json:"description"`
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

func (r ZoneRulesetPhaseEntrypointUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRule) implementsZoneRulesetPhaseEntrypointUpdateParamsRuleUnion() {
}

// Satisfied by [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRulesObject],
// [ZoneRulesetPhaseEntrypointUpdateParamsRule].
type ZoneRulesetPhaseEntrypointUpdateParamsRuleUnion interface {
	implementsZoneRulesetPhaseEntrypointUpdateParamsRuleUnion()
}

type ZoneRulesetPhaseEntrypointUpdateParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObject) implementsZoneRulesetPhaseEntrypointUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectAction string

const (
	ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionBlock ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectAction = "block"
)

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod int64

const (
	ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod10   ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 10
	ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod60   ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 60
	ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod600  ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 600
	ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod3600 ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod = 3600
)

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod10, ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod60, ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod600, ZoneRulesetPhaseEntrypointUpdateParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type ZoneRulesetPhaseEntrypointUpdateParamsRulesAction string

const (
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionBlock                ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "block"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionChallenge            ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "challenge"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionCompressResponse     ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "compress_response"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionExecute              ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "execute"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionJsChallenge          ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "js_challenge"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionLog                  ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "log"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionManagedChallenge     ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "managed_challenge"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRedirect             ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "redirect"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRewrite              ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "rewrite"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRoute                ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "route"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionScore                ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "score"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionServeError           ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "serve_error"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSetConfig            ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "set_config"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSkip                 ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "skip"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSetCacheSettings     ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "set_cache_settings"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionLogCustomField       ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "log_custom_field"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionDdosDynamic          ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "ddos_dynamic"
	ZoneRulesetPhaseEntrypointUpdateParamsRulesActionForceConnectionClose ZoneRulesetPhaseEntrypointUpdateParamsRulesAction = "force_connection_close"
)

func (r ZoneRulesetPhaseEntrypointUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case ZoneRulesetPhaseEntrypointUpdateParamsRulesActionBlock, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionChallenge, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionCompressResponse, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionExecute, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionJsChallenge, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionLog, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionManagedChallenge, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRedirect, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRewrite, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionRoute, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionScore, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionServeError, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSetConfig, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSkip, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionSetCacheSettings, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionLogCustomField, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionDdosDynamic, ZoneRulesetPhaseEntrypointUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}
