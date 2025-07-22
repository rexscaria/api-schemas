// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountRulesetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetService] method instead.
type AccountRulesetService struct {
	Options  []option.RequestOption
	Rules    *AccountRulesetRuleService
	Versions *AccountRulesetVersionService
	Phases   *AccountRulesetPhaseService
}

// NewAccountRulesetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRulesetService(opts ...option.RequestOption) (r *AccountRulesetService) {
	r = &AccountRulesetService{}
	r.Options = opts
	r.Rules = NewAccountRulesetRuleService(opts...)
	r.Versions = NewAccountRulesetVersionService(opts...)
	r.Phases = NewAccountRulesetPhaseService(opts...)
	return
}

// Creates a ruleset at the account level.
func (r *AccountRulesetService) New(ctx context.Context, accountID string, body AccountRulesetNewParams, opts ...option.RequestOption) (res *AccountRulesetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the latest version of an account ruleset.
func (r *AccountRulesetService) Get(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *AccountRulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account ruleset, creating a new version.
func (r *AccountRulesetService) Update(ctx context.Context, accountID string, rulesetID string, body AccountRulesetUpdateParams, opts ...option.RequestOption) (res *AccountRulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the account level.
func (r *AccountRulesetService) List(ctx context.Context, accountID string, query AccountRulesetListParams, opts ...option.RequestOption) (res *AccountRulesetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes all versions of an existing account ruleset.
func (r *AccountRulesetService) Delete(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AccountRulesetNewResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetNewResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetNewResponseSuccess `json:"success,required"`
	JSON    accountRulesetNewResponseJSON    `json:"-"`
}

// accountRulesetNewResponseJSON contains the JSON metadata for the struct
// [AccountRulesetNewResponse]
type accountRulesetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetNewResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetNewResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetNewResponseMessagesSource `json:"source"`
	JSON   accountRulesetNewResponseMessageJSON    `json:"-"`
}

// accountRulesetNewResponseMessageJSON contains the JSON metadata for the struct
// [AccountRulesetNewResponseMessage]
type accountRulesetNewResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetNewResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    accountRulesetNewResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetNewResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRulesetNewResponseMessagesSource]
type accountRulesetNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetNewResponseSuccess bool

const (
	AccountRulesetNewResponseSuccessTrue AccountRulesetNewResponseSuccess = true
)

func (r AccountRulesetNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetGetResponseSuccess `json:"success,required"`
	JSON    accountRulesetGetResponseJSON    `json:"-"`
}

// accountRulesetGetResponseJSON contains the JSON metadata for the struct
// [AccountRulesetGetResponse]
type accountRulesetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetGetResponseMessagesSource `json:"source"`
	JSON   accountRulesetGetResponseMessageJSON    `json:"-"`
}

// accountRulesetGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountRulesetGetResponseMessage]
type accountRulesetGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    accountRulesetGetResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRulesetGetResponseMessagesSource]
type accountRulesetGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetGetResponseSuccess bool

const (
	AccountRulesetGetResponseSuccessTrue AccountRulesetGetResponseSuccess = true
)

func (r AccountRulesetGetResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetGetResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetUpdateResponseSuccess `json:"success,required"`
	JSON    accountRulesetUpdateResponseJSON    `json:"-"`
}

// accountRulesetUpdateResponseJSON contains the JSON metadata for the struct
// [AccountRulesetUpdateResponse]
type accountRulesetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetUpdateResponseMessagesSource `json:"source"`
	JSON   accountRulesetUpdateResponseMessageJSON    `json:"-"`
}

// accountRulesetUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountRulesetUpdateResponseMessage]
type accountRulesetUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                         `json:"pointer,required"`
	JSON    accountRulesetUpdateResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetUpdateResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountRulesetUpdateResponseMessagesSource]
type accountRulesetUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetUpdateResponseSuccess bool

const (
	AccountRulesetUpdateResponseSuccessTrue AccountRulesetUpdateResponseSuccess = true
)

func (r AccountRulesetUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRulesetListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []AccountRulesetListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success AccountRulesetListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo AccountRulesetListResponseResultInfo `json:"result_info"`
	JSON       accountRulesetListResponseJSON       `json:"-"`
}

// accountRulesetListResponseJSON contains the JSON metadata for the struct
// [AccountRulesetListResponse]
type accountRulesetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type AccountRulesetListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source AccountRulesetListResponseMessagesSource `json:"source"`
	JSON   accountRulesetListResponseMessageJSON    `json:"-"`
}

// accountRulesetListResponseMessageJSON contains the JSON metadata for the struct
// [AccountRulesetListResponseMessage]
type accountRulesetListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type AccountRulesetListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    accountRulesetListResponseMessagesSourceJSON `json:"-"`
}

// accountRulesetListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountRulesetListResponseMessagesSource]
type accountRulesetListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountRulesetListResponseSuccess bool

const (
	AccountRulesetListResponseSuccessTrue AccountRulesetListResponseSuccess = true
)

func (r AccountRulesetListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRulesetListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type AccountRulesetListResponseResultInfo struct {
	// Set of cursors.
	Cursors AccountRulesetListResponseResultInfoCursors `json:"cursors"`
	JSON    accountRulesetListResponseResultInfoJSON    `json:"-"`
}

// accountRulesetListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountRulesetListResponseResultInfo]
type accountRulesetListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type AccountRulesetListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                          `json:"after"`
	JSON  accountRulesetListResponseResultInfoCursorsJSON `json:"-"`
}

// accountRulesetListResponseResultInfoCursorsJSON contains the JSON metadata for
// the struct [AccountRulesetListResponseResultInfoCursors]
type accountRulesetListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRulesetListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRulesetListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type AccountRulesetNewParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetNewParamsKind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[RulesetPhase] `json:"phase,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetNewParamsRuleUnion] `json:"rules"`
}

func (r AccountRulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type AccountRulesetNewParamsKind string

const (
	AccountRulesetNewParamsKindManaged AccountRulesetNewParamsKind = "managed"
	AccountRulesetNewParamsKindCustom  AccountRulesetNewParamsKind = "custom"
	AccountRulesetNewParamsKindRoot    AccountRulesetNewParamsKind = "root"
	AccountRulesetNewParamsKindZone    AccountRulesetNewParamsKind = "zone"
)

func (r AccountRulesetNewParamsKind) IsKnown() bool {
	switch r {
	case AccountRulesetNewParamsKindManaged, AccountRulesetNewParamsKindCustom, AccountRulesetNewParamsKindRoot, AccountRulesetNewParamsKindZone:
		return true
	}
	return false
}

type AccountRulesetNewParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                        `json:"action_parameters"`
	Categories       param.Field[interface{}]                        `json:"categories"`
	Description      param.Field[interface{}]                        `json:"description"`
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

func (r AccountRulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetNewParamsRule) implementsAccountRulesetNewParamsRuleUnion() {}

// Satisfied by [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRulesObject],
// [AccountRulesetNewParamsRulesObject], [AccountRulesetNewParamsRule].
type AccountRulesetNewParamsRuleUnion interface {
	implementsAccountRulesetNewParamsRuleUnion()
}

type AccountRulesetNewParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetNewParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetNewParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[AccountRulesetNewParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetNewParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[AccountRulesetNewParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetNewParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetNewParamsRulesObject) implementsAccountRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type AccountRulesetNewParamsRulesObjectAction string

const (
	AccountRulesetNewParamsRulesObjectActionBlock AccountRulesetNewParamsRulesObjectAction = "block"
)

func (r AccountRulesetNewParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case AccountRulesetNewParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type AccountRulesetNewParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetNewParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r AccountRulesetNewParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetNewParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetNewParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type AccountRulesetNewParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r AccountRulesetNewParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetNewParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetNewParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type AccountRulesetNewParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[AccountRulesetNewParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r AccountRulesetNewParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type AccountRulesetNewParamsRulesObjectRatelimitPeriod int64

const (
	AccountRulesetNewParamsRulesObjectRatelimitPeriod10   AccountRulesetNewParamsRulesObjectRatelimitPeriod = 10
	AccountRulesetNewParamsRulesObjectRatelimitPeriod60   AccountRulesetNewParamsRulesObjectRatelimitPeriod = 60
	AccountRulesetNewParamsRulesObjectRatelimitPeriod600  AccountRulesetNewParamsRulesObjectRatelimitPeriod = 600
	AccountRulesetNewParamsRulesObjectRatelimitPeriod3600 AccountRulesetNewParamsRulesObjectRatelimitPeriod = 3600
)

func (r AccountRulesetNewParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case AccountRulesetNewParamsRulesObjectRatelimitPeriod10, AccountRulesetNewParamsRulesObjectRatelimitPeriod60, AccountRulesetNewParamsRulesObjectRatelimitPeriod600, AccountRulesetNewParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type AccountRulesetNewParamsRulesAction string

const (
	AccountRulesetNewParamsRulesActionBlock                AccountRulesetNewParamsRulesAction = "block"
	AccountRulesetNewParamsRulesActionChallenge            AccountRulesetNewParamsRulesAction = "challenge"
	AccountRulesetNewParamsRulesActionCompressResponse     AccountRulesetNewParamsRulesAction = "compress_response"
	AccountRulesetNewParamsRulesActionExecute              AccountRulesetNewParamsRulesAction = "execute"
	AccountRulesetNewParamsRulesActionJsChallenge          AccountRulesetNewParamsRulesAction = "js_challenge"
	AccountRulesetNewParamsRulesActionLog                  AccountRulesetNewParamsRulesAction = "log"
	AccountRulesetNewParamsRulesActionManagedChallenge     AccountRulesetNewParamsRulesAction = "managed_challenge"
	AccountRulesetNewParamsRulesActionRedirect             AccountRulesetNewParamsRulesAction = "redirect"
	AccountRulesetNewParamsRulesActionRewrite              AccountRulesetNewParamsRulesAction = "rewrite"
	AccountRulesetNewParamsRulesActionRoute                AccountRulesetNewParamsRulesAction = "route"
	AccountRulesetNewParamsRulesActionScore                AccountRulesetNewParamsRulesAction = "score"
	AccountRulesetNewParamsRulesActionServeError           AccountRulesetNewParamsRulesAction = "serve_error"
	AccountRulesetNewParamsRulesActionSetConfig            AccountRulesetNewParamsRulesAction = "set_config"
	AccountRulesetNewParamsRulesActionSkip                 AccountRulesetNewParamsRulesAction = "skip"
	AccountRulesetNewParamsRulesActionSetCacheSettings     AccountRulesetNewParamsRulesAction = "set_cache_settings"
	AccountRulesetNewParamsRulesActionLogCustomField       AccountRulesetNewParamsRulesAction = "log_custom_field"
	AccountRulesetNewParamsRulesActionDdosDynamic          AccountRulesetNewParamsRulesAction = "ddos_dynamic"
	AccountRulesetNewParamsRulesActionForceConnectionClose AccountRulesetNewParamsRulesAction = "force_connection_close"
)

func (r AccountRulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case AccountRulesetNewParamsRulesActionBlock, AccountRulesetNewParamsRulesActionChallenge, AccountRulesetNewParamsRulesActionCompressResponse, AccountRulesetNewParamsRulesActionExecute, AccountRulesetNewParamsRulesActionJsChallenge, AccountRulesetNewParamsRulesActionLog, AccountRulesetNewParamsRulesActionManagedChallenge, AccountRulesetNewParamsRulesActionRedirect, AccountRulesetNewParamsRulesActionRewrite, AccountRulesetNewParamsRulesActionRoute, AccountRulesetNewParamsRulesActionScore, AccountRulesetNewParamsRulesActionServeError, AccountRulesetNewParamsRulesActionSetConfig, AccountRulesetNewParamsRulesActionSkip, AccountRulesetNewParamsRulesActionSetCacheSettings, AccountRulesetNewParamsRulesActionLogCustomField, AccountRulesetNewParamsRulesActionDdosDynamic, AccountRulesetNewParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

type AccountRulesetUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[RulesetPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetUpdateParamsRuleUnion] `json:"rules"`
}

func (r AccountRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type AccountRulesetUpdateParamsKind string

const (
	AccountRulesetUpdateParamsKindManaged AccountRulesetUpdateParamsKind = "managed"
	AccountRulesetUpdateParamsKindCustom  AccountRulesetUpdateParamsKind = "custom"
	AccountRulesetUpdateParamsKindRoot    AccountRulesetUpdateParamsKind = "root"
	AccountRulesetUpdateParamsKindZone    AccountRulesetUpdateParamsKind = "zone"
)

func (r AccountRulesetUpdateParamsKind) IsKnown() bool {
	switch r {
	case AccountRulesetUpdateParamsKindManaged, AccountRulesetUpdateParamsKindCustom, AccountRulesetUpdateParamsKindRoot, AccountRulesetUpdateParamsKindZone:
		return true
	}
	return false
}

type AccountRulesetUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                           `json:"action_parameters"`
	Categories       param.Field[interface{}]                           `json:"categories"`
	Description      param.Field[interface{}]                           `json:"description"`
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

func (r AccountRulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRule) implementsAccountRulesetUpdateParamsRuleUnion() {}

// Satisfied by [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject],
// [AccountRulesetUpdateParamsRulesObject], [AccountRulesetUpdateParamsRule].
type AccountRulesetUpdateParamsRuleUnion interface {
	implementsAccountRulesetUpdateParamsRuleUnion()
}

type AccountRulesetUpdateParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[AccountRulesetUpdateParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[AccountRulesetUpdateParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[AccountRulesetUpdateParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[AccountRulesetUpdateParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[AccountRulesetUpdateParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesObject) implementsAccountRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type AccountRulesetUpdateParamsRulesObjectAction string

const (
	AccountRulesetUpdateParamsRulesObjectActionBlock AccountRulesetUpdateParamsRulesObjectAction = "block"
)

func (r AccountRulesetUpdateParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case AccountRulesetUpdateParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type AccountRulesetUpdateParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[AccountRulesetUpdateParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r AccountRulesetUpdateParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type AccountRulesetUpdateParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r AccountRulesetUpdateParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type AccountRulesetUpdateParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r AccountRulesetUpdateParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r AccountRulesetUpdateParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type AccountRulesetUpdateParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[AccountRulesetUpdateParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r AccountRulesetUpdateParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type AccountRulesetUpdateParamsRulesObjectRatelimitPeriod int64

const (
	AccountRulesetUpdateParamsRulesObjectRatelimitPeriod10   AccountRulesetUpdateParamsRulesObjectRatelimitPeriod = 10
	AccountRulesetUpdateParamsRulesObjectRatelimitPeriod60   AccountRulesetUpdateParamsRulesObjectRatelimitPeriod = 60
	AccountRulesetUpdateParamsRulesObjectRatelimitPeriod600  AccountRulesetUpdateParamsRulesObjectRatelimitPeriod = 600
	AccountRulesetUpdateParamsRulesObjectRatelimitPeriod3600 AccountRulesetUpdateParamsRulesObjectRatelimitPeriod = 3600
)

func (r AccountRulesetUpdateParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case AccountRulesetUpdateParamsRulesObjectRatelimitPeriod10, AccountRulesetUpdateParamsRulesObjectRatelimitPeriod60, AccountRulesetUpdateParamsRulesObjectRatelimitPeriod600, AccountRulesetUpdateParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type AccountRulesetUpdateParamsRulesAction string

const (
	AccountRulesetUpdateParamsRulesActionBlock                AccountRulesetUpdateParamsRulesAction = "block"
	AccountRulesetUpdateParamsRulesActionChallenge            AccountRulesetUpdateParamsRulesAction = "challenge"
	AccountRulesetUpdateParamsRulesActionCompressResponse     AccountRulesetUpdateParamsRulesAction = "compress_response"
	AccountRulesetUpdateParamsRulesActionExecute              AccountRulesetUpdateParamsRulesAction = "execute"
	AccountRulesetUpdateParamsRulesActionJsChallenge          AccountRulesetUpdateParamsRulesAction = "js_challenge"
	AccountRulesetUpdateParamsRulesActionLog                  AccountRulesetUpdateParamsRulesAction = "log"
	AccountRulesetUpdateParamsRulesActionManagedChallenge     AccountRulesetUpdateParamsRulesAction = "managed_challenge"
	AccountRulesetUpdateParamsRulesActionRedirect             AccountRulesetUpdateParamsRulesAction = "redirect"
	AccountRulesetUpdateParamsRulesActionRewrite              AccountRulesetUpdateParamsRulesAction = "rewrite"
	AccountRulesetUpdateParamsRulesActionRoute                AccountRulesetUpdateParamsRulesAction = "route"
	AccountRulesetUpdateParamsRulesActionScore                AccountRulesetUpdateParamsRulesAction = "score"
	AccountRulesetUpdateParamsRulesActionServeError           AccountRulesetUpdateParamsRulesAction = "serve_error"
	AccountRulesetUpdateParamsRulesActionSetConfig            AccountRulesetUpdateParamsRulesAction = "set_config"
	AccountRulesetUpdateParamsRulesActionSkip                 AccountRulesetUpdateParamsRulesAction = "skip"
	AccountRulesetUpdateParamsRulesActionSetCacheSettings     AccountRulesetUpdateParamsRulesAction = "set_cache_settings"
	AccountRulesetUpdateParamsRulesActionLogCustomField       AccountRulesetUpdateParamsRulesAction = "log_custom_field"
	AccountRulesetUpdateParamsRulesActionDdosDynamic          AccountRulesetUpdateParamsRulesAction = "ddos_dynamic"
	AccountRulesetUpdateParamsRulesActionForceConnectionClose AccountRulesetUpdateParamsRulesAction = "force_connection_close"
)

func (r AccountRulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case AccountRulesetUpdateParamsRulesActionBlock, AccountRulesetUpdateParamsRulesActionChallenge, AccountRulesetUpdateParamsRulesActionCompressResponse, AccountRulesetUpdateParamsRulesActionExecute, AccountRulesetUpdateParamsRulesActionJsChallenge, AccountRulesetUpdateParamsRulesActionLog, AccountRulesetUpdateParamsRulesActionManagedChallenge, AccountRulesetUpdateParamsRulesActionRedirect, AccountRulesetUpdateParamsRulesActionRewrite, AccountRulesetUpdateParamsRulesActionRoute, AccountRulesetUpdateParamsRulesActionScore, AccountRulesetUpdateParamsRulesActionServeError, AccountRulesetUpdateParamsRulesActionSetConfig, AccountRulesetUpdateParamsRulesActionSkip, AccountRulesetUpdateParamsRulesActionSetCacheSettings, AccountRulesetUpdateParamsRulesActionLogCustomField, AccountRulesetUpdateParamsRulesActionDdosDynamic, AccountRulesetUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

type AccountRulesetListParams struct {
	// Cursor to use for the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Number of rulesets to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccountRulesetListParams]'s query parameters as
// `url.Values`.
func (r AccountRulesetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
