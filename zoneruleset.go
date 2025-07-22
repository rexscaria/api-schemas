// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneRulesetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetService] method instead.
type ZoneRulesetService struct {
	Options  []option.RequestOption
	Rules    *ZoneRulesetRuleService
	Versions *ZoneRulesetVersionService
	Phases   *ZoneRulesetPhaseService
}

// NewZoneRulesetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetService(opts ...option.RequestOption) (r *ZoneRulesetService) {
	r = &ZoneRulesetService{}
	r.Options = opts
	r.Rules = NewZoneRulesetRuleService(opts...)
	r.Versions = NewZoneRulesetVersionService(opts...)
	r.Phases = NewZoneRulesetPhaseService(opts...)
	return
}

// Creates a ruleset at the zone level.
func (r *ZoneRulesetService) New(ctx context.Context, zoneID string, body ZoneRulesetNewParams, opts ...option.RequestOption) (res *ZoneRulesetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the latest version of a zone ruleset.
func (r *ZoneRulesetService) Get(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (res *ZoneRulesetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a zone ruleset, creating a new version.
func (r *ZoneRulesetService) Update(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetUpdateParams, opts ...option.RequestOption) (res *ZoneRulesetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the zone level.
func (r *ZoneRulesetService) List(ctx context.Context, zoneID string, query ZoneRulesetListParams, opts ...option.RequestOption) (res *ZoneRulesetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes all versions of an existing zone ruleset.
func (r *ZoneRulesetService) Delete(ctx context.Context, zoneID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rulesets/%s", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type ZoneRulesetNewResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetNewResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetNewResponseSuccess `json:"success,required"`
	JSON    zoneRulesetNewResponseJSON    `json:"-"`
}

// zoneRulesetNewResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetNewResponse]
type zoneRulesetNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetNewResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetNewResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetNewResponseMessagesSource `json:"source"`
	JSON   zoneRulesetNewResponseMessageJSON    `json:"-"`
}

// zoneRulesetNewResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetNewResponseMessage]
type zoneRulesetNewResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetNewResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    zoneRulesetNewResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetNewResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetNewResponseMessagesSource]
type zoneRulesetNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetNewResponseSuccess bool

const (
	ZoneRulesetNewResponseSuccessTrue ZoneRulesetNewResponseSuccess = true
)

func (r ZoneRulesetNewResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetNewResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetGetResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetGetResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetGetResponseSuccess `json:"success,required"`
	JSON    zoneRulesetGetResponseJSON    `json:"-"`
}

// zoneRulesetGetResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetGetResponse]
type zoneRulesetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetGetResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetGetResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetGetResponseMessagesSource `json:"source"`
	JSON   zoneRulesetGetResponseMessageJSON    `json:"-"`
}

// zoneRulesetGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetGetResponseMessage]
type zoneRulesetGetResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetGetResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetGetResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    zoneRulesetGetResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetGetResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetGetResponseMessagesSource]
type zoneRulesetGetResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetGetResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetGetResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetGetResponseSuccess bool

const (
	ZoneRulesetGetResponseSuccessTrue ZoneRulesetGetResponseSuccess = true
)

func (r ZoneRulesetGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetUpdateResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetUpdateResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetUpdateResponseSuccess `json:"success,required"`
	JSON    zoneRulesetUpdateResponseJSON    `json:"-"`
}

// zoneRulesetUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetUpdateResponse]
type zoneRulesetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetUpdateResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetUpdateResponseMessagesSource `json:"source"`
	JSON   zoneRulesetUpdateResponseMessageJSON    `json:"-"`
}

// zoneRulesetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetUpdateResponseMessage]
type zoneRulesetUpdateResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetUpdateResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetUpdateResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer,required"`
	JSON    zoneRulesetUpdateResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetUpdateResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetUpdateResponseMessagesSource]
type zoneRulesetUpdateResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetUpdateResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetUpdateResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetUpdateResponseSuccess bool

const (
	ZoneRulesetUpdateResponseSuccessTrue ZoneRulesetUpdateResponseSuccess = true
)

func (r ZoneRulesetUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRulesetListResponse struct {
	// A list of error messages.
	Errors interface{} `json:"errors,required"`
	// A list of warning messages.
	Messages []ZoneRulesetListResponseMessage `json:"messages,required"`
	// A result.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success ZoneRulesetListResponseSuccess `json:"success,required"`
	// Cursors information to navigate the results.
	ResultInfo ZoneRulesetListResponseResultInfo `json:"result_info"`
	JSON       zoneRulesetListResponseJSON       `json:"-"`
}

// zoneRulesetListResponseJSON contains the JSON metadata for the struct
// [ZoneRulesetListResponse]
type zoneRulesetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetListResponseJSON) RawJSON() string {
	return r.raw
}

// A message.
type ZoneRulesetListResponseMessage struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ZoneRulesetListResponseMessagesSource `json:"source"`
	JSON   zoneRulesetListResponseMessageJSON    `json:"-"`
}

// zoneRulesetListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneRulesetListResponseMessage]
type zoneRulesetListResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetListResponseMessageJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ZoneRulesetListResponseMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    zoneRulesetListResponseMessagesSourceJSON `json:"-"`
}

// zoneRulesetListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZoneRulesetListResponseMessagesSource]
type zoneRulesetListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRulesetListResponseSuccess bool

const (
	ZoneRulesetListResponseSuccessTrue ZoneRulesetListResponseSuccess = true
)

func (r ZoneRulesetListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRulesetListResponseSuccessTrue:
		return true
	}
	return false
}

// Cursors information to navigate the results.
type ZoneRulesetListResponseResultInfo struct {
	// Set of cursors.
	Cursors ZoneRulesetListResponseResultInfoCursors `json:"cursors"`
	JSON    zoneRulesetListResponseResultInfoJSON    `json:"-"`
}

// zoneRulesetListResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneRulesetListResponseResultInfo]
type zoneRulesetListResponseResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

// Set of cursors.
type ZoneRulesetListResponseResultInfoCursors struct {
	// Cursor to use for the next page.
	After string                                       `json:"after"`
	JSON  zoneRulesetListResponseResultInfoCursorsJSON `json:"-"`
}

// zoneRulesetListResponseResultInfoCursorsJSON contains the JSON metadata for the
// struct [ZoneRulesetListResponseResultInfoCursors]
type zoneRulesetListResponseResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRulesetListResponseResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRulesetListResponseResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type ZoneRulesetNewParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetNewParamsKind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[RulesetPhase] `json:"phase,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetNewParamsRuleUnion] `json:"rules"`
}

func (r ZoneRulesetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type ZoneRulesetNewParamsKind string

const (
	ZoneRulesetNewParamsKindManaged ZoneRulesetNewParamsKind = "managed"
	ZoneRulesetNewParamsKindCustom  ZoneRulesetNewParamsKind = "custom"
	ZoneRulesetNewParamsKindRoot    ZoneRulesetNewParamsKind = "root"
	ZoneRulesetNewParamsKindZone    ZoneRulesetNewParamsKind = "zone"
)

func (r ZoneRulesetNewParamsKind) IsKnown() bool {
	switch r {
	case ZoneRulesetNewParamsKindManaged, ZoneRulesetNewParamsKindCustom, ZoneRulesetNewParamsKindRoot, ZoneRulesetNewParamsKindZone:
		return true
	}
	return false
}

type ZoneRulesetNewParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                     `json:"action_parameters"`
	Categories       param.Field[interface{}]                     `json:"categories"`
	Description      param.Field[interface{}]                     `json:"description"`
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

func (r ZoneRulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetNewParamsRule) implementsZoneRulesetNewParamsRuleUnion() {}

// Satisfied by [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRulesObject],
// [ZoneRulesetNewParamsRulesObject], [ZoneRulesetNewParamsRule].
type ZoneRulesetNewParamsRuleUnion interface {
	implementsZoneRulesetNewParamsRuleUnion()
}

type ZoneRulesetNewParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetNewParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetNewParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ZoneRulesetNewParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetNewParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ZoneRulesetNewParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetNewParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetNewParamsRulesObject) implementsZoneRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type ZoneRulesetNewParamsRulesObjectAction string

const (
	ZoneRulesetNewParamsRulesObjectActionBlock ZoneRulesetNewParamsRulesObjectAction = "block"
)

func (r ZoneRulesetNewParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case ZoneRulesetNewParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type ZoneRulesetNewParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetNewParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetNewParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetNewParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetNewParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ZoneRulesetNewParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ZoneRulesetNewParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetNewParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetNewParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ZoneRulesetNewParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[ZoneRulesetNewParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r ZoneRulesetNewParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type ZoneRulesetNewParamsRulesObjectRatelimitPeriod int64

const (
	ZoneRulesetNewParamsRulesObjectRatelimitPeriod10   ZoneRulesetNewParamsRulesObjectRatelimitPeriod = 10
	ZoneRulesetNewParamsRulesObjectRatelimitPeriod60   ZoneRulesetNewParamsRulesObjectRatelimitPeriod = 60
	ZoneRulesetNewParamsRulesObjectRatelimitPeriod600  ZoneRulesetNewParamsRulesObjectRatelimitPeriod = 600
	ZoneRulesetNewParamsRulesObjectRatelimitPeriod3600 ZoneRulesetNewParamsRulesObjectRatelimitPeriod = 3600
)

func (r ZoneRulesetNewParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case ZoneRulesetNewParamsRulesObjectRatelimitPeriod10, ZoneRulesetNewParamsRulesObjectRatelimitPeriod60, ZoneRulesetNewParamsRulesObjectRatelimitPeriod600, ZoneRulesetNewParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type ZoneRulesetNewParamsRulesAction string

const (
	ZoneRulesetNewParamsRulesActionBlock                ZoneRulesetNewParamsRulesAction = "block"
	ZoneRulesetNewParamsRulesActionChallenge            ZoneRulesetNewParamsRulesAction = "challenge"
	ZoneRulesetNewParamsRulesActionCompressResponse     ZoneRulesetNewParamsRulesAction = "compress_response"
	ZoneRulesetNewParamsRulesActionExecute              ZoneRulesetNewParamsRulesAction = "execute"
	ZoneRulesetNewParamsRulesActionJsChallenge          ZoneRulesetNewParamsRulesAction = "js_challenge"
	ZoneRulesetNewParamsRulesActionLog                  ZoneRulesetNewParamsRulesAction = "log"
	ZoneRulesetNewParamsRulesActionManagedChallenge     ZoneRulesetNewParamsRulesAction = "managed_challenge"
	ZoneRulesetNewParamsRulesActionRedirect             ZoneRulesetNewParamsRulesAction = "redirect"
	ZoneRulesetNewParamsRulesActionRewrite              ZoneRulesetNewParamsRulesAction = "rewrite"
	ZoneRulesetNewParamsRulesActionRoute                ZoneRulesetNewParamsRulesAction = "route"
	ZoneRulesetNewParamsRulesActionScore                ZoneRulesetNewParamsRulesAction = "score"
	ZoneRulesetNewParamsRulesActionServeError           ZoneRulesetNewParamsRulesAction = "serve_error"
	ZoneRulesetNewParamsRulesActionSetConfig            ZoneRulesetNewParamsRulesAction = "set_config"
	ZoneRulesetNewParamsRulesActionSkip                 ZoneRulesetNewParamsRulesAction = "skip"
	ZoneRulesetNewParamsRulesActionSetCacheSettings     ZoneRulesetNewParamsRulesAction = "set_cache_settings"
	ZoneRulesetNewParamsRulesActionLogCustomField       ZoneRulesetNewParamsRulesAction = "log_custom_field"
	ZoneRulesetNewParamsRulesActionDdosDynamic          ZoneRulesetNewParamsRulesAction = "ddos_dynamic"
	ZoneRulesetNewParamsRulesActionForceConnectionClose ZoneRulesetNewParamsRulesAction = "force_connection_close"
)

func (r ZoneRulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case ZoneRulesetNewParamsRulesActionBlock, ZoneRulesetNewParamsRulesActionChallenge, ZoneRulesetNewParamsRulesActionCompressResponse, ZoneRulesetNewParamsRulesActionExecute, ZoneRulesetNewParamsRulesActionJsChallenge, ZoneRulesetNewParamsRulesActionLog, ZoneRulesetNewParamsRulesActionManagedChallenge, ZoneRulesetNewParamsRulesActionRedirect, ZoneRulesetNewParamsRulesActionRewrite, ZoneRulesetNewParamsRulesActionRoute, ZoneRulesetNewParamsRulesActionScore, ZoneRulesetNewParamsRulesActionServeError, ZoneRulesetNewParamsRulesActionSetConfig, ZoneRulesetNewParamsRulesActionSkip, ZoneRulesetNewParamsRulesActionSetCacheSettings, ZoneRulesetNewParamsRulesActionLogCustomField, ZoneRulesetNewParamsRulesActionDdosDynamic, ZoneRulesetNewParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

type ZoneRulesetUpdateParams struct {
	// The unique ID of the ruleset.
	ID param.Field[string] `json:"id,required"`
	// The version of the ruleset.
	Version param.Field[string] `json:"version,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[ZoneRulesetUpdateParamsKind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[RulesetPhase] `json:"phase"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetUpdateParamsRuleUnion] `json:"rules"`
}

func (r ZoneRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type ZoneRulesetUpdateParamsKind string

const (
	ZoneRulesetUpdateParamsKindManaged ZoneRulesetUpdateParamsKind = "managed"
	ZoneRulesetUpdateParamsKindCustom  ZoneRulesetUpdateParamsKind = "custom"
	ZoneRulesetUpdateParamsKindRoot    ZoneRulesetUpdateParamsKind = "root"
	ZoneRulesetUpdateParamsKindZone    ZoneRulesetUpdateParamsKind = "zone"
)

func (r ZoneRulesetUpdateParamsKind) IsKnown() bool {
	switch r {
	case ZoneRulesetUpdateParamsKindManaged, ZoneRulesetUpdateParamsKindCustom, ZoneRulesetUpdateParamsKindRoot, ZoneRulesetUpdateParamsKindZone:
		return true
	}
	return false
}

type ZoneRulesetUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetUpdateParamsRulesAction] `json:"action"`
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

func (r ZoneRulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRule) implementsZoneRulesetUpdateParamsRuleUnion() {}

// Satisfied by [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRulesObject],
// [ZoneRulesetUpdateParamsRulesObject], [ZoneRulesetUpdateParamsRule].
type ZoneRulesetUpdateParamsRuleUnion interface {
	implementsZoneRulesetUpdateParamsRuleUnion()
}

type ZoneRulesetUpdateParamsRulesObject struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[ZoneRulesetUpdateParamsRulesObjectAction]           `json:"action"`
	ActionParameters param.Field[ZoneRulesetUpdateParamsRulesObjectActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[interface{}] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configure checks for exposed credentials.
	ExposedCredentialCheck param.Field[ZoneRulesetUpdateParamsRulesObjectExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetUpdateParamsRulesObjectLogging] `json:"logging"`
	// An object configuring the rule's ratelimit behavior.
	Ratelimit param.Field[ZoneRulesetUpdateParamsRulesObjectRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetUpdateParamsRulesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetUpdateParamsRulesObject) implementsZoneRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type ZoneRulesetUpdateParamsRulesObjectAction string

const (
	ZoneRulesetUpdateParamsRulesObjectActionBlock ZoneRulesetUpdateParamsRulesObjectAction = "block"
)

func (r ZoneRulesetUpdateParamsRulesObjectAction) IsKnown() bool {
	switch r {
	case ZoneRulesetUpdateParamsRulesObjectActionBlock:
		return true
	}
	return false
}

type ZoneRulesetUpdateParamsRulesObjectActionParameters struct {
	// The response to show when the block is applied.
	Response param.Field[ZoneRulesetUpdateParamsRulesObjectActionParametersResponse] `json:"response"`
}

func (r ZoneRulesetUpdateParamsRulesObjectActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type ZoneRulesetUpdateParamsRulesObjectActionParametersResponse struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r ZoneRulesetUpdateParamsRulesObjectActionParametersResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure checks for exposed credentials.
type ZoneRulesetUpdateParamsRulesObjectExposedCredentialCheck struct {
	// Expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression,required"`
	// Expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression,required"`
}

func (r ZoneRulesetUpdateParamsRulesObjectExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetUpdateParamsRulesObjectLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ZoneRulesetUpdateParamsRulesObjectLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's ratelimit behavior.
type ZoneRulesetUpdateParamsRulesObjectRatelimit struct {
	// Characteristics of the request on which the ratelimiter counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics,required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod] `json:"period,required"`
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

func (r ZoneRulesetUpdateParamsRulesObjectRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Period in seconds over which the counter is being incremented.
type ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod int64

const (
	ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod10   ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod = 10
	ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod60   ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod = 60
	ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod600  ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod = 600
	ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod3600 ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod = 3600
)

func (r ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod) IsKnown() bool {
	switch r {
	case ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod10, ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod60, ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod600, ZoneRulesetUpdateParamsRulesObjectRatelimitPeriod3600:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type ZoneRulesetUpdateParamsRulesAction string

const (
	ZoneRulesetUpdateParamsRulesActionBlock                ZoneRulesetUpdateParamsRulesAction = "block"
	ZoneRulesetUpdateParamsRulesActionChallenge            ZoneRulesetUpdateParamsRulesAction = "challenge"
	ZoneRulesetUpdateParamsRulesActionCompressResponse     ZoneRulesetUpdateParamsRulesAction = "compress_response"
	ZoneRulesetUpdateParamsRulesActionExecute              ZoneRulesetUpdateParamsRulesAction = "execute"
	ZoneRulesetUpdateParamsRulesActionJsChallenge          ZoneRulesetUpdateParamsRulesAction = "js_challenge"
	ZoneRulesetUpdateParamsRulesActionLog                  ZoneRulesetUpdateParamsRulesAction = "log"
	ZoneRulesetUpdateParamsRulesActionManagedChallenge     ZoneRulesetUpdateParamsRulesAction = "managed_challenge"
	ZoneRulesetUpdateParamsRulesActionRedirect             ZoneRulesetUpdateParamsRulesAction = "redirect"
	ZoneRulesetUpdateParamsRulesActionRewrite              ZoneRulesetUpdateParamsRulesAction = "rewrite"
	ZoneRulesetUpdateParamsRulesActionRoute                ZoneRulesetUpdateParamsRulesAction = "route"
	ZoneRulesetUpdateParamsRulesActionScore                ZoneRulesetUpdateParamsRulesAction = "score"
	ZoneRulesetUpdateParamsRulesActionServeError           ZoneRulesetUpdateParamsRulesAction = "serve_error"
	ZoneRulesetUpdateParamsRulesActionSetConfig            ZoneRulesetUpdateParamsRulesAction = "set_config"
	ZoneRulesetUpdateParamsRulesActionSkip                 ZoneRulesetUpdateParamsRulesAction = "skip"
	ZoneRulesetUpdateParamsRulesActionSetCacheSettings     ZoneRulesetUpdateParamsRulesAction = "set_cache_settings"
	ZoneRulesetUpdateParamsRulesActionLogCustomField       ZoneRulesetUpdateParamsRulesAction = "log_custom_field"
	ZoneRulesetUpdateParamsRulesActionDdosDynamic          ZoneRulesetUpdateParamsRulesAction = "ddos_dynamic"
	ZoneRulesetUpdateParamsRulesActionForceConnectionClose ZoneRulesetUpdateParamsRulesAction = "force_connection_close"
)

func (r ZoneRulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case ZoneRulesetUpdateParamsRulesActionBlock, ZoneRulesetUpdateParamsRulesActionChallenge, ZoneRulesetUpdateParamsRulesActionCompressResponse, ZoneRulesetUpdateParamsRulesActionExecute, ZoneRulesetUpdateParamsRulesActionJsChallenge, ZoneRulesetUpdateParamsRulesActionLog, ZoneRulesetUpdateParamsRulesActionManagedChallenge, ZoneRulesetUpdateParamsRulesActionRedirect, ZoneRulesetUpdateParamsRulesActionRewrite, ZoneRulesetUpdateParamsRulesActionRoute, ZoneRulesetUpdateParamsRulesActionScore, ZoneRulesetUpdateParamsRulesActionServeError, ZoneRulesetUpdateParamsRulesActionSetConfig, ZoneRulesetUpdateParamsRulesActionSkip, ZoneRulesetUpdateParamsRulesActionSetCacheSettings, ZoneRulesetUpdateParamsRulesActionLogCustomField, ZoneRulesetUpdateParamsRulesActionDdosDynamic, ZoneRulesetUpdateParamsRulesActionForceConnectionClose:
		return true
	}
	return false
}

type ZoneRulesetListParams struct {
	// Cursor to use for the next page.
	Cursor param.Field[string] `query:"cursor"`
	// Number of rulesets to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ZoneRulesetListParams]'s query parameters as `url.Values`.
func (r ZoneRulesetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
