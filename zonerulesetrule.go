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

// ZoneRulesetRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetRuleService] method instead.
type ZoneRulesetRuleService struct {
	Options []option.RequestOption
}

// NewZoneRulesetRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetRuleService(opts ...option.RequestOption) (r *ZoneRulesetRuleService) {
	r = &ZoneRulesetRuleService{}
	r.Options = opts
	return
}

// Adds a new rule to a zone ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset by default.
func (r *ZoneRulesetRuleService) New(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetRuleNewParams, opts ...option.RequestOption) (res *ZoneRulesetRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an existing rule in a zone ruleset.
func (r *ZoneRulesetRuleService) Update(ctx context.Context, zoneID string, rulesetID string, ruleID string, body ZoneRulesetRuleUpdateParams, opts ...option.RequestOption) (res *ZoneRulesetRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
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
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from a zone ruleset.
func (r *ZoneRulesetRuleService) Delete(ctx context.Context, zoneID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *ZoneRulesetRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
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
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ZoneRulesetRuleNewResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleNewResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleNewResponseSuccess `json:"success,required"`
	JSON    zoneRulesetRuleNewResponseJSON    `json:"-"`
}

// zoneRulesetRuleNewResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleNewResponse]
type zoneRulesetRuleNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetRuleNewResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleNewResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleNewResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleNewResponseMessage]
type zoneRulesetRuleNewResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetRuleNewResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    zoneRulesetRuleNewResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleNewResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleNewResponseMessagesSource]
type zoneRulesetRuleNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetRuleNewResponseSuccess bool

const (
	ZoneRulesetRuleNewResponseSuccessTrue ZoneRulesetRuleNewResponseSuccess = true
)

func (r ZoneRulesetRuleNewResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleNewResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetRuleUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleUpdateResponseSuccess `json:"success,required"`
	JSON    zoneRulesetRuleUpdateResponseJSON    `json:"-"`
}

// zoneRulesetRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleUpdateResponse]
type zoneRulesetRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetRuleUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleUpdateResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleUpdateResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleUpdateResponseMessage]
type zoneRulesetRuleUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetRuleUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetRuleUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleUpdateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetRuleUpdateResponseMessagesSource]
type zoneRulesetRuleUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetRuleUpdateResponseSuccess bool

const (
	ZoneRulesetRuleUpdateResponseSuccessTrue ZoneRulesetRuleUpdateResponseSuccess = true
)

func (r ZoneRulesetRuleUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetRuleDeleteResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetRuleDeleteResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetRuleDeleteResponseSuccess `json:"success,required"`
	JSON    zoneRulesetRuleDeleteResponseJSON    `json:"-"`
}

// zoneRulesetRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetRuleDeleteResponse]
type zoneRulesetRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetRuleDeleteResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetRuleDeleteResponseMessagesSource `json:"source"`
	JSON   zoneRulesetRuleDeleteResponseMessageJSON    `json:"-"`
}

// zoneRulesetRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneRulesetRuleDeleteResponseMessage]
type zoneRulesetRuleDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetRuleDeleteResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    zoneRulesetRuleDeleteResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetRuleDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneRulesetRuleDeleteResponseMessagesSource]
type zoneRulesetRuleDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetRuleDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetRuleDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetRuleDeleteResponseSuccess bool

const (
	ZoneRulesetRuleDeleteResponseSuccessTrue ZoneRulesetRuleDeleteResponseSuccess = true
)

func (r ZoneRulesetRuleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetRuleNewParams struct {
	Body ZoneRulesetRuleNewParamsBodyUnion `json:"body,required"`
}

func (r ZoneRulesetRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneRulesetRuleNewParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetRuleNewParamsBodyAction] `json:"action"`
	ActionParameters param.Field[interface{}]                        `json:"action_parameters"`
	Categories       param.Field[interface{}]                        `json:"categories"`
	Description      param.Field[interface{}]                        `json:"description"`
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

func (r ZoneRulesetRuleNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBody) implementsZoneRulesetRuleNewParamsBodyUnion() {}

// Satisfied by [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBodyObject],
// [ZoneRulesetRuleNewParamsBodyObject], [ZoneRulesetRuleNewParamsBody].
type ZoneRulesetRuleNewParamsBodyUnion interface {
	implementsZoneRulesetRuleNewParamsBodyUnion()
}

type ZoneRulesetRuleNewParamsBodyObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetRuleNewParamsBodyObjectAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetRuleNewParamsBodyObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ZoneRulesetRuleNewParamsBodyObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging  param.Field[ZoneRulesetRuleNewParamsBodyObjectLogging]       `json:"logging"`
	Position param.Field[ZoneRulesetRuleNewParamsBodyObjectPositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ZoneRulesetRuleNewParamsBodyObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetRuleNewParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBodyObject) implementsZoneRulesetRuleNewParamsBodyUnion() {}

// The action to perform when the rule matches.
type ZoneRulesetRuleNewParamsBodyObjectAction string

const (
	ZoneRulesetRuleNewParamsBodyObjectActionBlock ZoneRulesetRuleNewParamsBodyObjectAction = "block"
)

func (r ZoneRulesetRuleNewParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleNewParamsBodyObjectActionBlock:
		return true
	}
	return false
}

type ZoneRulesetRuleNewParamsBodyObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetRuleNewParamsBodyObjectActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetRuleNewParamsBodyObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ZoneRulesetRuleNewParamsBodyObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetRuleNewParamsBodyObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetRuleNewParamsBodyObjectPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBodyObjectPosition) implementsZoneRulesetRuleNewParamsBodyObjectPositionUnion() {
}

// Satisfied by [ZoneRulesetRuleNewParamsBodyObjectPositionBefore],
// [ZoneRulesetRuleNewParamsBodyObjectPositionAfter],
// [ZoneRulesetRuleNewParamsBodyObjectPositionIndex],
// [ZoneRulesetRuleNewParamsBodyObjectPosition].
type ZoneRulesetRuleNewParamsBodyObjectPositionUnion interface {
	implementsZoneRulesetRuleNewParamsBodyObjectPositionUnion()
}

type ZoneRulesetRuleNewParamsBodyObjectPositionBefore struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionBefore) implementsZoneRulesetRuleNewParamsBodyObjectPositionUnion() {
}

type ZoneRulesetRuleNewParamsBodyObjectPositionAfter struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionAfter) implementsZoneRulesetRuleNewParamsBodyObjectPositionUnion() {
}

type ZoneRulesetRuleNewParamsBodyObjectPositionIndex struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleNewParamsBodyObjectPositionIndex) implementsZoneRulesetRuleNewParamsBodyObjectPositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type ZoneRulesetRuleNewParamsBodyObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod] `json:"period,required"`
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

func (r ZoneRulesetRuleNewParamsBodyObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod int64

const (
	ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod10   ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod = 10
	ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod60   ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod = 60
	ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod600  ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod = 600
	ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod3600 ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod = 3600
)

func (r ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod10, ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod60, ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod600, ZoneRulesetRuleNewParamsBodyObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type ZoneRulesetRuleNewParamsBodyAction string

const (
	ZoneRulesetRuleNewParamsBodyActionBlock                ZoneRulesetRuleNewParamsBodyAction = "block"
	ZoneRulesetRuleNewParamsBodyActionChallenge            ZoneRulesetRuleNewParamsBodyAction = "challenge"
	ZoneRulesetRuleNewParamsBodyActionCompressResponse     ZoneRulesetRuleNewParamsBodyAction = "compress_response"
	ZoneRulesetRuleNewParamsBodyActionExecute              ZoneRulesetRuleNewParamsBodyAction = "execute"
	ZoneRulesetRuleNewParamsBodyActionJsChallenge          ZoneRulesetRuleNewParamsBodyAction = "js_challenge"
	ZoneRulesetRuleNewParamsBodyActionLog                  ZoneRulesetRuleNewParamsBodyAction = "log"
	ZoneRulesetRuleNewParamsBodyActionManagedChallenge     ZoneRulesetRuleNewParamsBodyAction = "managed_challenge"
	ZoneRulesetRuleNewParamsBodyActionRedirect             ZoneRulesetRuleNewParamsBodyAction = "redirect"
	ZoneRulesetRuleNewParamsBodyActionRewrite              ZoneRulesetRuleNewParamsBodyAction = "rewrite"
	ZoneRulesetRuleNewParamsBodyActionRoute                ZoneRulesetRuleNewParamsBodyAction = "route"
	ZoneRulesetRuleNewParamsBodyActionScore                ZoneRulesetRuleNewParamsBodyAction = "score"
	ZoneRulesetRuleNewParamsBodyActionServeError           ZoneRulesetRuleNewParamsBodyAction = "serve_error"
	ZoneRulesetRuleNewParamsBodyActionSetConfig            ZoneRulesetRuleNewParamsBodyAction = "set_config"
	ZoneRulesetRuleNewParamsBodyActionSkip                 ZoneRulesetRuleNewParamsBodyAction = "skip"
	ZoneRulesetRuleNewParamsBodyActionSetCacheSettings     ZoneRulesetRuleNewParamsBodyAction = "set_cache_settings"
	ZoneRulesetRuleNewParamsBodyActionLogCustomField       ZoneRulesetRuleNewParamsBodyAction = "log_custom_field"
	ZoneRulesetRuleNewParamsBodyActionDdosDynamic          ZoneRulesetRuleNewParamsBodyAction = "ddos_dynamic"
	ZoneRulesetRuleNewParamsBodyActionForceConnectionClose ZoneRulesetRuleNewParamsBodyAction = "force_connection_close"
)

func (r ZoneRulesetRuleNewParamsBodyAction) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleNewParamsBodyActionBlock, ZoneRulesetRuleNewParamsBodyActionChallenge, ZoneRulesetRuleNewParamsBodyActionCompressResponse, ZoneRulesetRuleNewParamsBodyActionExecute, ZoneRulesetRuleNewParamsBodyActionJsChallenge, ZoneRulesetRuleNewParamsBodyActionLog, ZoneRulesetRuleNewParamsBodyActionManagedChallenge, ZoneRulesetRuleNewParamsBodyActionRedirect, ZoneRulesetRuleNewParamsBodyActionRewrite, ZoneRulesetRuleNewParamsBodyActionRoute, ZoneRulesetRuleNewParamsBodyActionScore, ZoneRulesetRuleNewParamsBodyActionServeError, ZoneRulesetRuleNewParamsBodyActionSetConfig, ZoneRulesetRuleNewParamsBodyActionSkip, ZoneRulesetRuleNewParamsBodyActionSetCacheSettings, ZoneRulesetRuleNewParamsBodyActionLogCustomField, ZoneRulesetRuleNewParamsBodyActionDdosDynamic, ZoneRulesetRuleNewParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}

type ZoneRulesetRuleUpdateParams struct {
	Body ZoneRulesetRuleUpdateParamsBodyUnion `json:"body,required"`
}

func (r ZoneRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneRulesetRuleUpdateParamsBody struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetRuleUpdateParamsBodyAction] `json:"action"`
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

func (r ZoneRulesetRuleUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBody) implementsZoneRulesetRuleUpdateParamsBodyUnion() {}

// Satisfied by [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject],
// [ZoneRulesetRuleUpdateParamsBodyObject], [ZoneRulesetRuleUpdateParamsBody].
type ZoneRulesetRuleUpdateParamsBodyUnion interface {
	implementsZoneRulesetRuleUpdateParamsBodyUnion()
}

type ZoneRulesetRuleUpdateParamsBodyObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetRuleUpdateParamsBodyObjectAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetRuleUpdateParamsBodyObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ZoneRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging  param.Field[ZoneRulesetRuleUpdateParamsBodyObjectLogging]       `json:"logging"`
	Position param.Field[ZoneRulesetRuleUpdateParamsBodyObjectPositionUnion] `json:"position"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ZoneRulesetRuleUpdateParamsBodyObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBodyObject) implementsZoneRulesetRuleUpdateParamsBodyUnion() {}

// The action to perform when the rule matches.
type ZoneRulesetRuleUpdateParamsBodyObjectAction string

const (
	ZoneRulesetRuleUpdateParamsBodyObjectActionBlock ZoneRulesetRuleUpdateParamsBodyObjectAction = "block"
)

func (r ZoneRulesetRuleUpdateParamsBodyObjectAction) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleUpdateParamsBodyObjectActionBlock:
		return true
	}
	return false
}

type ZoneRulesetRuleUpdateParamsBodyObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetRuleUpdateParamsBodyObjectActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetRuleUpdateParamsBodyObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ZoneRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetRuleUpdateParamsBodyObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetRuleUpdateParamsBodyObjectPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPosition) implementsZoneRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

// Satisfied by [ZoneRulesetRuleUpdateParamsBodyObjectPositionBefore],
// [ZoneRulesetRuleUpdateParamsBodyObjectPositionAfter],
// [ZoneRulesetRuleUpdateParamsBodyObjectPositionIndex],
// [ZoneRulesetRuleUpdateParamsBodyObjectPosition].
type ZoneRulesetRuleUpdateParamsBodyObjectPositionUnion interface {
	implementsZoneRulesetRuleUpdateParamsBodyObjectPositionUnion()
}

type ZoneRulesetRuleUpdateParamsBodyObjectPositionBefore struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionBefore) implementsZoneRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

type ZoneRulesetRuleUpdateParamsBodyObjectPositionAfter struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionAfter) implementsZoneRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

type ZoneRulesetRuleUpdateParamsBodyObjectPositionIndex struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetRuleUpdateParamsBodyObjectPositionIndex) implementsZoneRulesetRuleUpdateParamsBodyObjectPositionUnion() {
}

// An object configuring the rule's ratelimit behavior.
type ZoneRulesetRuleUpdateParamsBodyObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod] `json:"period,required"`
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

func (r ZoneRulesetRuleUpdateParamsBodyObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod int64

const (
	ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod10   ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 10
	ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod60   ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 60
	ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod600  ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 600
	ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod3600 ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod = 3600
)

func (r ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod10, ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod60, ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod600, ZoneRulesetRuleUpdateParamsBodyObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type ZoneRulesetRuleUpdateParamsBodyAction string

const (
	ZoneRulesetRuleUpdateParamsBodyActionBlock                ZoneRulesetRuleUpdateParamsBodyAction = "block"
	ZoneRulesetRuleUpdateParamsBodyActionChallenge            ZoneRulesetRuleUpdateParamsBodyAction = "challenge"
	ZoneRulesetRuleUpdateParamsBodyActionCompressResponse     ZoneRulesetRuleUpdateParamsBodyAction = "compress_response"
	ZoneRulesetRuleUpdateParamsBodyActionExecute              ZoneRulesetRuleUpdateParamsBodyAction = "execute"
	ZoneRulesetRuleUpdateParamsBodyActionJsChallenge          ZoneRulesetRuleUpdateParamsBodyAction = "js_challenge"
	ZoneRulesetRuleUpdateParamsBodyActionLog                  ZoneRulesetRuleUpdateParamsBodyAction = "log"
	ZoneRulesetRuleUpdateParamsBodyActionManagedChallenge     ZoneRulesetRuleUpdateParamsBodyAction = "managed_challenge"
	ZoneRulesetRuleUpdateParamsBodyActionRedirect             ZoneRulesetRuleUpdateParamsBodyAction = "redirect"
	ZoneRulesetRuleUpdateParamsBodyActionRewrite              ZoneRulesetRuleUpdateParamsBodyAction = "rewrite"
	ZoneRulesetRuleUpdateParamsBodyActionRoute                ZoneRulesetRuleUpdateParamsBodyAction = "route"
	ZoneRulesetRuleUpdateParamsBodyActionScore                ZoneRulesetRuleUpdateParamsBodyAction = "score"
	ZoneRulesetRuleUpdateParamsBodyActionServeError           ZoneRulesetRuleUpdateParamsBodyAction = "serve_error"
	ZoneRulesetRuleUpdateParamsBodyActionSetConfig            ZoneRulesetRuleUpdateParamsBodyAction = "set_config"
	ZoneRulesetRuleUpdateParamsBodyActionSkip                 ZoneRulesetRuleUpdateParamsBodyAction = "skip"
	ZoneRulesetRuleUpdateParamsBodyActionSetCacheSettings     ZoneRulesetRuleUpdateParamsBodyAction = "set_cache_settings"
	ZoneRulesetRuleUpdateParamsBodyActionLogCustomField       ZoneRulesetRuleUpdateParamsBodyAction = "log_custom_field"
	ZoneRulesetRuleUpdateParamsBodyActionDdosDynamic          ZoneRulesetRuleUpdateParamsBodyAction = "ddos_dynamic"
	ZoneRulesetRuleUpdateParamsBodyActionForceConnectionClose ZoneRulesetRuleUpdateParamsBodyAction = "force_connection_close"
)

func (r ZoneRulesetRuleUpdateParamsBodyAction) IsKnown() bool {
	switch r {
	case ZoneRulesetRuleUpdateParamsBodyActionBlock, ZoneRulesetRuleUpdateParamsBodyActionChallenge, ZoneRulesetRuleUpdateParamsBodyActionCompressResponse, ZoneRulesetRuleUpdateParamsBodyActionExecute, ZoneRulesetRuleUpdateParamsBodyActionJsChallenge, ZoneRulesetRuleUpdateParamsBodyActionLog, ZoneRulesetRuleUpdateParamsBodyActionManagedChallenge, ZoneRulesetRuleUpdateParamsBodyActionRedirect, ZoneRulesetRuleUpdateParamsBodyActionRewrite, ZoneRulesetRuleUpdateParamsBodyActionRoute, ZoneRulesetRuleUpdateParamsBodyActionScore, ZoneRulesetRuleUpdateParamsBodyActionServeError, ZoneRulesetRuleUpdateParamsBodyActionSetConfig, ZoneRulesetRuleUpdateParamsBodyActionSkip, ZoneRulesetRuleUpdateParamsBodyActionSetCacheSettings, ZoneRulesetRuleUpdateParamsBodyActionLogCustomField, ZoneRulesetRuleUpdateParamsBodyActionDdosDynamic, ZoneRulesetRuleUpdateParamsBodyActionForceConnectionClose:
		return true
	}
	return false
}
