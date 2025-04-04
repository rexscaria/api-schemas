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

// AccountRumV2RuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRumV2RuleService] method instead.
type AccountRumV2RuleService struct {
	Options []option.RequestOption
}

// NewAccountRumV2RuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRumV2RuleService(opts ...option.RequestOption) (r *AccountRumV2RuleService) {
	r = &AccountRumV2RuleService{}
	r.Options = opts
	return
}

// Creates a new rule in a Web Analytics ruleset.
func (r *AccountRumV2RuleService) New(ctx context.Context, accountID string, rulesetID string, body AccountRumV2RuleNewParams, opts ...option.RequestOption) (res *ResponseSingleRule, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all the rules in a Web Analytics ruleset.
func (r *AccountRumV2RuleService) List(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *ResponseCollectionRules, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an existing rule from a Web Analytics ruleset.
func (r *AccountRumV2RuleService) Delete(ctx context.Context, accountID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *AccountRumV2RuleDeleteResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates a rule in a Web Analytics ruleset.
func (r *AccountRumV2RuleService) Update0(ctx context.Context, accountID string, rulesetID string, ruleID string, body AccountRumV2RuleUpdate0Params, opts ...option.RequestOption) (res *ResponseSingleRule, err error) {
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
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rule/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Modifies one or more rules in a Web Analytics ruleset with a single request.
func (r *AccountRumV2RuleService) Update1(ctx context.Context, accountID string, rulesetID string, body AccountRumV2RuleUpdate1Params, opts ...option.RequestOption) (res *ResponseCollectionRules, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rum/v2/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateRequestParam struct {
	Host param.Field[string] `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive param.Field[bool] `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused param.Field[bool]     `json:"is_paused"`
	Paths    param.Field[[]string] `json:"paths"`
}

func (r CreateRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseCollectionRules struct {
	Result ResponseCollectionRulesResult `json:"result"`
	JSON   responseCollectionRulesJSON   `json:"-"`
	Common
}

// responseCollectionRulesJSON contains the JSON metadata for the struct
// [ResponseCollectionRules]
type responseCollectionRulesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionRulesJSON) RawJSON() string {
	return r.raw
}

type ResponseCollectionRulesResult struct {
	// A list of rules.
	Rules   []RumRule                         `json:"rules"`
	Ruleset Ruleset                           `json:"ruleset"`
	JSON    responseCollectionRulesResultJSON `json:"-"`
}

// responseCollectionRulesResultJSON contains the JSON metadata for the struct
// [ResponseCollectionRulesResult]
type responseCollectionRulesResultJSON struct {
	Rules       apijson.Field
	Ruleset     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionRulesResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseCollectionRulesResultJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleRule struct {
	Result RumRule                `json:"result"`
	JSON   responseSingleRuleJSON `json:"-"`
	Common
}

// responseSingleRuleJSON contains the JSON metadata for the struct
// [ResponseSingleRule]
type responseSingleRuleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleRuleJSON) RawJSON() string {
	return r.raw
}

type Ruleset struct {
	// The Web Analytics ruleset identifier.
	ID string `json:"id"`
	// Whether the ruleset is enabled.
	Enabled  bool   `json:"enabled"`
	ZoneName string `json:"zone_name"`
	// The zone identifier.
	ZoneTag string      `json:"zone_tag"`
	JSON    rulesetJSON `json:"-"`
}

// rulesetJSON contains the JSON metadata for the struct [Ruleset]
type rulesetJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	ZoneName    apijson.Field
	ZoneTag     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ruleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetJSON) RawJSON() string {
	return r.raw
}

type AccountRumV2RuleDeleteResponse struct {
	Result AccountRumV2RuleDeleteResponseResult `json:"result"`
	JSON   accountRumV2RuleDeleteResponseJSON   `json:"-"`
	Common
}

// accountRumV2RuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountRumV2RuleDeleteResponse]
type accountRumV2RuleDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumV2RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumV2RuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRumV2RuleDeleteResponseResult struct {
	// The Web Analytics rule identifier.
	ID   string                                   `json:"id"`
	JSON accountRumV2RuleDeleteResponseResultJSON `json:"-"`
}

// accountRumV2RuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountRumV2RuleDeleteResponseResult]
type accountRumV2RuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRumV2RuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRumV2RuleDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountRumV2RuleNewParams struct {
	CreateRequest CreateRequestParam `json:"create_request,required"`
}

func (r AccountRumV2RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRequest)
}

type AccountRumV2RuleUpdate0Params struct {
	CreateRequest CreateRequestParam `json:"create_request,required"`
}

func (r AccountRumV2RuleUpdate0Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CreateRequest)
}

type AccountRumV2RuleUpdate1Params struct {
	// A list of rule identifiers to delete.
	DeleteRules param.Field[[]string] `json:"delete_rules"`
	// A list of rules to create or update.
	Rules param.Field[[]AccountRumV2RuleUpdate1ParamsRule] `json:"rules"`
}

func (r AccountRumV2RuleUpdate1Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRumV2RuleUpdate1ParamsRule struct {
	// The Web Analytics rule identifier.
	ID        param.Field[string]   `json:"id"`
	Host      param.Field[string]   `json:"host"`
	Inclusive param.Field[bool]     `json:"inclusive"`
	IsPaused  param.Field[bool]     `json:"is_paused"`
	Paths     param.Field[[]string] `json:"paths"`
}

func (r AccountRumV2RuleUpdate1ParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
