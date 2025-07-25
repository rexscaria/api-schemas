// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneWaitingRoomRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneWaitingRoomRuleService] method instead.
type ZoneWaitingRoomRuleService struct {
	Options []option.RequestOption
}

// NewZoneWaitingRoomRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneWaitingRoomRuleService(opts ...option.RequestOption) (r *ZoneWaitingRoomRuleService) {
	r = &ZoneWaitingRoomRuleService{}
	r.Options = opts
	return
}

// Only available for the Waiting Room Advanced subscription. Creates a rule for a
// waiting room.
func (r *ZoneWaitingRoomRuleService) New(ctx context.Context, zoneID string, waitingRoomID string, body ZoneWaitingRoomRuleNewParams, opts ...option.RequestOption) (res *ResponseCollectionWaitingRoomRules, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists rules for a waiting room.
func (r *ZoneWaitingRoomRuleService) List(ctx context.Context, zoneID string, waitingRoomID string, opts ...option.RequestOption) (res *ResponseCollectionWaitingRoomRules, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a rule for a waiting room.
func (r *ZoneWaitingRoomRuleService) Delete(ctx context.Context, zoneID string, waitingRoomID string, ruleID string, opts ...option.RequestOption) (res *ResponseCollectionWaitingRoomRules, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", zoneID, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Patches a rule for a waiting room.
func (r *ZoneWaitingRoomRuleService) Patch(ctx context.Context, zoneID string, waitingRoomID string, ruleID string, body ZoneWaitingRoomRulePatchParams, opts ...option.RequestOption) (res *ResponseCollectionWaitingRoomRules, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules/%s", zoneID, waitingRoomID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Only available for the Waiting Room Advanced subscription. Replaces all rules
// for a waiting room.
func (r *ZoneWaitingRoomRuleService) Replace(ctx context.Context, zoneID string, waitingRoomID string, body ZoneWaitingRoomRuleReplaceParams, opts ...option.RequestOption) (res *ResponseCollectionWaitingRoomRules, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if waitingRoomID == "" {
		err = errors.New("missing required waiting_room_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/waiting_rooms/%s/rules", zoneID, waitingRoomID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CreateWaitingRoomRuleParam struct {
	// The action to take when the expression matches.
	Action param.Field[RuleAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r CreateWaitingRoomRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseCollectionWaitingRoomRules struct {
	Errors   []WaitingroomMessage `json:"errors,required"`
	Messages []WaitingroomMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    ResponseCollectionWaitingRoomRulesSuccess    `json:"success,required"`
	Result     []ResponseCollectionWaitingRoomRulesResult   `json:"result"`
	ResultInfo ResponseCollectionWaitingRoomRulesResultInfo `json:"result_info"`
	JSON       responseCollectionWaitingRoomRulesJSON       `json:"-"`
}

// responseCollectionWaitingRoomRulesJSON contains the JSON metadata for the struct
// [ResponseCollectionWaitingRoomRules]
type responseCollectionWaitingRoomRulesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionWaitingRoomRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionWaitingRoomRulesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ResponseCollectionWaitingRoomRulesSuccess bool

const (
	ResponseCollectionWaitingRoomRulesSuccessTrue ResponseCollectionWaitingRoomRulesSuccess = true
)

func (r ResponseCollectionWaitingRoomRulesSuccess) IsKnown() bool {
	switch r {
	case ResponseCollectionWaitingRoomRulesSuccessTrue:
		return true
	}
	return false
}

type ResponseCollectionWaitingRoomRulesResult struct {
	// The ID of the rule.
	ID string `json:"id"`
	// The action to take when the expression matches.
	Action RuleAction `json:"action"`
	// The description of the rule.
	Description string `json:"description"`
	// When set to true, the rule is enabled.
	Enabled bool `json:"enabled"`
	// Criteria defining when there is a match for the current rule.
	Expression  string    `json:"expression"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The version of the rule.
	Version string                                       `json:"version"`
	JSON    responseCollectionWaitingRoomRulesResultJSON `json:"-"`
}

// responseCollectionWaitingRoomRulesResultJSON contains the JSON metadata for the
// struct [ResponseCollectionWaitingRoomRulesResult]
type responseCollectionWaitingRoomRulesResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionWaitingRoomRulesResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionWaitingRoomRulesResultJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionWaitingRoomRulesResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                          `json:"total_count"`
	JSON       responseCollectionWaitingRoomRulesResultInfoJSON `json:"-"`
}

// responseCollectionWaitingRoomRulesResultInfoJSON contains the JSON metadata for
// the struct [ResponseCollectionWaitingRoomRulesResultInfo]
type responseCollectionWaitingRoomRulesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionWaitingRoomRulesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionWaitingRoomRulesResultInfoJSON) RawJSON() string {
	return r.raw
}

// The action to take when the expression matches.
type RuleAction string

const (
	RuleActionBypassWaitingRoom RuleAction = "bypass_waiting_room"
)

func (r RuleAction) IsKnown() bool {
	switch r {
	case RuleActionBypassWaitingRoom:
		return true
	}
	return false
}

type ZoneWaitingRoomRuleNewParams struct {
	CreateWaitingRoomRule CreateWaitingRoomRuleParam `json:"create_waiting_room_rule,required"`
}

func (r ZoneWaitingRoomRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateWaitingRoomRule)
}

type ZoneWaitingRoomRulePatchParams struct {
	// The action to take when the expression matches.
	Action param.Field[RuleAction] `json:"action,required"`
	// Criteria defining when there is a match for the current rule.
	Expression param.Field[string] `json:"expression,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// When set to true, the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Reorder the position of a rule
	Position param.Field[ZoneWaitingRoomRulePatchParamsPositionUnion] `json:"position"`
}

func (r ZoneWaitingRoomRulePatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reorder the position of a rule
type ZoneWaitingRoomRulePatchParamsPosition struct {
	// Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the last rule in the ruleset.
	After param.Field[string] `json:"after"`
	// Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the first rule in the ruleset.
	Before param.Field[string] `json:"before"`
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r ZoneWaitingRoomRulePatchParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneWaitingRoomRulePatchParamsPosition) implementsZoneWaitingRoomRulePatchParamsPositionUnion() {
}

// Reorder the position of a rule
//
// Satisfied by [ZoneWaitingRoomRulePatchParamsPositionIndex],
// [ZoneWaitingRoomRulePatchParamsPositionBefore],
// [ZoneWaitingRoomRulePatchParamsPositionAfter],
// [ZoneWaitingRoomRulePatchParamsPosition].
type ZoneWaitingRoomRulePatchParamsPositionUnion interface {
	implementsZoneWaitingRoomRulePatchParamsPositionUnion()
}

type ZoneWaitingRoomRulePatchParamsPositionIndex struct {
	// Places the rule in the exact position specified by the integer number
	// <POSITION_NUMBER>. Position numbers start with 1. Existing rules in the ruleset
	// from the specified position number onward are shifted one position (no rule is
	// overwritten).
	Index param.Field[int64] `json:"index"`
}

func (r ZoneWaitingRoomRulePatchParamsPositionIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneWaitingRoomRulePatchParamsPositionIndex) implementsZoneWaitingRoomRulePatchParamsPositionUnion() {
}

type ZoneWaitingRoomRulePatchParamsPositionBefore struct {
	// Places the rule before rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the first rule in the ruleset.
	Before param.Field[string] `json:"before"`
}

func (r ZoneWaitingRoomRulePatchParamsPositionBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneWaitingRoomRulePatchParamsPositionBefore) implementsZoneWaitingRoomRulePatchParamsPositionUnion() {
}

type ZoneWaitingRoomRulePatchParamsPositionAfter struct {
	// Places the rule after rule <RULE_ID>. Use this argument with an empty rule ID
	// value ("") to set the rule as the last rule in the ruleset.
	After param.Field[string] `json:"after"`
}

func (r ZoneWaitingRoomRulePatchParamsPositionAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneWaitingRoomRulePatchParamsPositionAfter) implementsZoneWaitingRoomRulePatchParamsPositionUnion() {
}

type ZoneWaitingRoomRuleReplaceParams struct {
	Body []CreateWaitingRoomRuleParam `json:"body,required"`
}

func (r ZoneWaitingRoomRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
