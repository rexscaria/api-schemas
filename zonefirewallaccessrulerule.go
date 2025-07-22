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

// ZoneFirewallAccessRuleRuleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallAccessRuleRuleService] method instead.
type ZoneFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *ZoneFirewallAccessRuleRuleService) {
	r = &ZoneFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for a zone.
//
// Note: To create an IP Access rule that applies to multiple zones, refer to
// [IP Access rules for a user](#ip-access-rules-for-a-user) or
// [IP Access rules for an account](#ip-access-rules-for-an-account) as
// appropriate.
func (r *ZoneFirewallAccessRuleRuleService) New(ctx context.Context, zoneID string, body ZoneFirewallAccessRuleRuleNewParams, opts ...option.RequestOption) (res *FirewallRuleSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an IP Access rule defined at the zone level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *ZoneFirewallAccessRuleRuleService) Update(ctx context.Context, zoneID string, ruleID string, body ZoneFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *FirewallRuleSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of a zone. You can filter the results using several
// optional parameters.
func (r *ZoneFirewallAccessRuleRuleService) List(ctx context.Context, zoneID string, query ZoneFirewallAccessRuleRuleListParams, opts ...option.RequestOption) (res *FirewallRuleCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an IP Access rule defined at the zone level.
//
// Optionally, you can use the `cascade` property to specify that you wish to
// delete similar rules in other zones managed by the same zone owner.
func (r *ZoneFirewallAccessRuleRuleService) Delete(ctx context.Context, zoneID string, ruleID string, body ZoneFirewallAccessRuleRuleDeleteParams, opts ...option.RequestOption) (res *FirewallRuleSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/access_rules/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type ZoneFirewallAccessRuleRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes,required"`
}

func (r ZoneFirewallAccessRuleRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallAccessRuleRuleUpdateParams struct {
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r ZoneFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallAccessRuleRuleListParams struct {
	Configuration param.Field[ZoneFirewallAccessRuleRuleListParamsConfiguration] `query:"configuration"`
	// The direction used to sort returned rules.
	Direction param.Field[ZoneFirewallAccessRuleRuleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallAccessRuleRuleListParamsMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
	// The field used to sort returned rules.
	Order param.Field[ZoneFirewallAccessRuleRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallAccessRuleRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallAccessRuleRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallAccessRuleRuleListParamsConfiguration struct {
	// The target to search in existing rules.
	Target param.Field[ZoneFirewallAccessRuleRuleListParamsConfigurationTarget] `query:"target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [ZoneFirewallAccessRuleRuleListParamsConfiguration]'s query
// parameters as `url.Values`.
func (r ZoneFirewallAccessRuleRuleListParamsConfiguration) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type ZoneFirewallAccessRuleRuleListParamsConfigurationTarget string

const (
	ZoneFirewallAccessRuleRuleListParamsConfigurationTargetIP      ZoneFirewallAccessRuleRuleListParamsConfigurationTarget = "ip"
	ZoneFirewallAccessRuleRuleListParamsConfigurationTargetIPRange ZoneFirewallAccessRuleRuleListParamsConfigurationTarget = "ip_range"
	ZoneFirewallAccessRuleRuleListParamsConfigurationTargetAsn     ZoneFirewallAccessRuleRuleListParamsConfigurationTarget = "asn"
	ZoneFirewallAccessRuleRuleListParamsConfigurationTargetCountry ZoneFirewallAccessRuleRuleListParamsConfigurationTarget = "country"
)

func (r ZoneFirewallAccessRuleRuleListParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case ZoneFirewallAccessRuleRuleListParamsConfigurationTargetIP, ZoneFirewallAccessRuleRuleListParamsConfigurationTargetIPRange, ZoneFirewallAccessRuleRuleListParamsConfigurationTargetAsn, ZoneFirewallAccessRuleRuleListParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The direction used to sort returned rules.
type ZoneFirewallAccessRuleRuleListParamsDirection string

const (
	ZoneFirewallAccessRuleRuleListParamsDirectionAsc  ZoneFirewallAccessRuleRuleListParamsDirection = "asc"
	ZoneFirewallAccessRuleRuleListParamsDirectionDesc ZoneFirewallAccessRuleRuleListParamsDirection = "desc"
)

func (r ZoneFirewallAccessRuleRuleListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneFirewallAccessRuleRuleListParamsDirectionAsc, ZoneFirewallAccessRuleRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallAccessRuleRuleListParamsMatch string

const (
	ZoneFirewallAccessRuleRuleListParamsMatchAny ZoneFirewallAccessRuleRuleListParamsMatch = "any"
	ZoneFirewallAccessRuleRuleListParamsMatchAll ZoneFirewallAccessRuleRuleListParamsMatch = "all"
)

func (r ZoneFirewallAccessRuleRuleListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneFirewallAccessRuleRuleListParamsMatchAny, ZoneFirewallAccessRuleRuleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned rules.
type ZoneFirewallAccessRuleRuleListParamsOrder string

const (
	ZoneFirewallAccessRuleRuleListParamsOrderConfigurationTarget ZoneFirewallAccessRuleRuleListParamsOrder = "configuration.target"
	ZoneFirewallAccessRuleRuleListParamsOrderConfigurationValue  ZoneFirewallAccessRuleRuleListParamsOrder = "configuration.value"
	ZoneFirewallAccessRuleRuleListParamsOrderMode                ZoneFirewallAccessRuleRuleListParamsOrder = "mode"
)

func (r ZoneFirewallAccessRuleRuleListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneFirewallAccessRuleRuleListParamsOrderConfigurationTarget, ZoneFirewallAccessRuleRuleListParamsOrderConfigurationValue, ZoneFirewallAccessRuleRuleListParamsOrderMode:
		return true
	}
	return false
}

type ZoneFirewallAccessRuleRuleDeleteParams struct {
	// The level to attempt to delete similar rules defined for other zones with the
	// same owner. The default value is `none`, which will only delete the current
	// rule. Using `basic` will delete rules that match the same action (mode) and
	// configuration, while using `aggressive` will delete rules that match the same
	// configuration.
	Cascade param.Field[ZoneFirewallAccessRuleRuleDeleteParamsCascade] `json:"cascade"`
}

func (r ZoneFirewallAccessRuleRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The level to attempt to delete similar rules defined for other zones with the
// same owner. The default value is `none`, which will only delete the current
// rule. Using `basic` will delete rules that match the same action (mode) and
// configuration, while using `aggressive` will delete rules that match the same
// configuration.
type ZoneFirewallAccessRuleRuleDeleteParamsCascade string

const (
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeNone       ZoneFirewallAccessRuleRuleDeleteParamsCascade = "none"
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeBasic      ZoneFirewallAccessRuleRuleDeleteParamsCascade = "basic"
	ZoneFirewallAccessRuleRuleDeleteParamsCascadeAggressive ZoneFirewallAccessRuleRuleDeleteParamsCascade = "aggressive"
)

func (r ZoneFirewallAccessRuleRuleDeleteParamsCascade) IsKnown() bool {
	switch r {
	case ZoneFirewallAccessRuleRuleDeleteParamsCascadeNone, ZoneFirewallAccessRuleRuleDeleteParamsCascadeBasic, ZoneFirewallAccessRuleRuleDeleteParamsCascadeAggressive:
		return true
	}
	return false
}
