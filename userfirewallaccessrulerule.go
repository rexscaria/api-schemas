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

// UserFirewallAccessRuleRuleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserFirewallAccessRuleRuleService] method instead.
type UserFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewUserFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *UserFirewallAccessRuleRuleService) {
	r = &UserFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for all zones owned by the current user.
//
// Note: To create an IP Access rule that applies to a specific zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *UserFirewallAccessRuleRuleService) New(ctx context.Context, body UserFirewallAccessRuleRuleNewParams, opts ...option.RequestOption) (res *FirewallRuleSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an IP Access rule defined at the user level. You can only update the
// rule action (`mode` parameter) and notes.
func (r *UserFirewallAccessRuleRuleService) Update(ctx context.Context, ruleID string, body UserFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *FirewallRuleSingle, err error) {
	opts = append(r.Options[:], opts...)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of the user. You can filter the results using several
// optional parameters.
func (r *UserFirewallAccessRuleRuleService) List(ctx context.Context, query UserFirewallAccessRuleRuleListParams, opts ...option.RequestOption) (res *FirewallRuleCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/firewall/access_rules/rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an IP Access rule at the user level.
//
// Note: Deleting a user-level rule will affect all zones owned by the user.
func (r *UserFirewallAccessRuleRuleService) Delete(ctx context.Context, ruleID string, opts ...option.RequestOption) (res *FirewallRuleSingleID, err error) {
	opts = append(r.Options[:], opts...)
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("user/firewall/access_rules/rules/%s", ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FirewallRuleCollection struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   []FirewallRule         `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    FirewallRuleCollectionSuccess    `json:"success,required"`
	ResultInfo FirewallRuleCollectionResultInfo `json:"result_info"`
	JSON       firewallRuleCollectionJSON       `json:"-"`
}

// firewallRuleCollectionJSON contains the JSON metadata for the struct
// [FirewallRuleCollection]
type firewallRuleCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleCollectionJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallRuleCollectionSuccess bool

const (
	FirewallRuleCollectionSuccessTrue FirewallRuleCollectionSuccess = true
)

func (r FirewallRuleCollectionSuccess) IsKnown() bool {
	switch r {
	case FirewallRuleCollectionSuccessTrue:
		return true
	}
	return false
}

type FirewallRuleCollectionResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                              `json:"total_count"`
	JSON       firewallRuleCollectionResultInfoJSON `json:"-"`
}

// firewallRuleCollectionResultInfoJSON contains the JSON metadata for the struct
// [FirewallRuleCollectionResultInfo]
type firewallRuleCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleSingle struct {
	Errors   []FirewallMessagesItem `json:"errors,required"`
	Messages []FirewallMessagesItem `json:"messages,required"`
	Result   FirewallRule           `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallRuleSingleSuccess `json:"success,required"`
	JSON    firewallRuleSingleJSON    `json:"-"`
}

// firewallRuleSingleJSON contains the JSON metadata for the struct
// [FirewallRuleSingle]
type firewallRuleSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleSingleJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallRuleSingleSuccess bool

const (
	FirewallRuleSingleSuccessTrue FirewallRuleSingleSuccess = true
)

func (r FirewallRuleSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallRuleSingleSuccessTrue:
		return true
	}
	return false
}

type FirewallRuleSingleID struct {
	Errors   []FirewallMessagesItem     `json:"errors,required"`
	Messages []FirewallMessagesItem     `json:"messages,required"`
	Result   FirewallRuleSingleIDResult `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallRuleSingleIDSuccess `json:"success,required"`
	JSON    firewallRuleSingleIDJSON    `json:"-"`
}

// firewallRuleSingleIDJSON contains the JSON metadata for the struct
// [FirewallRuleSingleID]
type firewallRuleSingleIDJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleSingleIDJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleSingleIDResult struct {
	// The unique identifier of the IP Access rule.
	ID   string                         `json:"id"`
	JSON firewallRuleSingleIDResultJSON `json:"-"`
}

// firewallRuleSingleIDResultJSON contains the JSON metadata for the struct
// [FirewallRuleSingleIDResult]
type firewallRuleSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleSingleIDResultJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallRuleSingleIDSuccess bool

const (
	FirewallRuleSingleIDSuccessTrue FirewallRuleSingleIDSuccess = true
)

func (r FirewallRuleSingleIDSuccess) IsKnown() bool {
	switch r {
	case FirewallRuleSingleIDSuccessTrue:
		return true
	}
	return false
}

type UserFirewallAccessRuleRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserFirewallAccessRuleRuleUpdateParams struct {
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r UserFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserFirewallAccessRuleRuleListParams struct {
	Configuration param.Field[UserFirewallAccessRuleRuleListParamsConfiguration] `query:"configuration"`
	// Defines the direction used to sort returned rules.
	Direction param.Field[UserFirewallAccessRuleRuleListParamsDirection] `query:"direction"`
	// Defines the search requirements. When set to `all`, all the search requirements
	// must match. When set to `any`, only one of the search requirements has to match.
	Match param.Field[UserFirewallAccessRuleRuleListParamsMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `query:"mode"`
	// Defines the string to search for in the notes of existing IP Access rules.
	// Notes: For example, the string 'attack' would match IP Access rules with notes
	// 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
	// Defines the field used to sort returned rules.
	Order param.Field[UserFirewallAccessRuleRuleListParamsOrder] `query:"order"`
	// Defines the requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Defines the maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParams]'s query parameters as
// `url.Values`.
func (r UserFirewallAccessRuleRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserFirewallAccessRuleRuleListParamsConfiguration struct {
	// Defines the target to search in existing rules.
	Target param.Field[UserFirewallAccessRuleRuleListParamsConfigurationTarget] `query:"target"`
	// Defines the target value to search for in existing rules: an IP address, an IP
	// address range, or a country code, depending on the provided
	// `configuration.target`. Notes: You can search for a single IPv4 address, an IP
	// address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1
	// alpha-2 country code.
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [UserFirewallAccessRuleRuleListParamsConfiguration]'s query
// parameters as `url.Values`.
func (r UserFirewallAccessRuleRuleListParamsConfiguration) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Defines the target to search in existing rules.
type UserFirewallAccessRuleRuleListParamsConfigurationTarget string

const (
	UserFirewallAccessRuleRuleListParamsConfigurationTargetIP      UserFirewallAccessRuleRuleListParamsConfigurationTarget = "ip"
	UserFirewallAccessRuleRuleListParamsConfigurationTargetIPRange UserFirewallAccessRuleRuleListParamsConfigurationTarget = "ip_range"
	UserFirewallAccessRuleRuleListParamsConfigurationTargetAsn     UserFirewallAccessRuleRuleListParamsConfigurationTarget = "asn"
	UserFirewallAccessRuleRuleListParamsConfigurationTargetCountry UserFirewallAccessRuleRuleListParamsConfigurationTarget = "country"
)

func (r UserFirewallAccessRuleRuleListParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case UserFirewallAccessRuleRuleListParamsConfigurationTargetIP, UserFirewallAccessRuleRuleListParamsConfigurationTargetIPRange, UserFirewallAccessRuleRuleListParamsConfigurationTargetAsn, UserFirewallAccessRuleRuleListParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// Defines the direction used to sort returned rules.
type UserFirewallAccessRuleRuleListParamsDirection string

const (
	UserFirewallAccessRuleRuleListParamsDirectionAsc  UserFirewallAccessRuleRuleListParamsDirection = "asc"
	UserFirewallAccessRuleRuleListParamsDirectionDesc UserFirewallAccessRuleRuleListParamsDirection = "desc"
)

func (r UserFirewallAccessRuleRuleListParamsDirection) IsKnown() bool {
	switch r {
	case UserFirewallAccessRuleRuleListParamsDirectionAsc, UserFirewallAccessRuleRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// Defines the search requirements. When set to `all`, all the search requirements
// must match. When set to `any`, only one of the search requirements has to match.
type UserFirewallAccessRuleRuleListParamsMatch string

const (
	UserFirewallAccessRuleRuleListParamsMatchAny UserFirewallAccessRuleRuleListParamsMatch = "any"
	UserFirewallAccessRuleRuleListParamsMatchAll UserFirewallAccessRuleRuleListParamsMatch = "all"
)

func (r UserFirewallAccessRuleRuleListParamsMatch) IsKnown() bool {
	switch r {
	case UserFirewallAccessRuleRuleListParamsMatchAny, UserFirewallAccessRuleRuleListParamsMatchAll:
		return true
	}
	return false
}

// Defines the field used to sort returned rules.
type UserFirewallAccessRuleRuleListParamsOrder string

const (
	UserFirewallAccessRuleRuleListParamsOrderConfigurationTarget UserFirewallAccessRuleRuleListParamsOrder = "configuration.target"
	UserFirewallAccessRuleRuleListParamsOrderConfigurationValue  UserFirewallAccessRuleRuleListParamsOrder = "configuration.value"
	UserFirewallAccessRuleRuleListParamsOrderMode                UserFirewallAccessRuleRuleListParamsOrder = "mode"
)

func (r UserFirewallAccessRuleRuleListParamsOrder) IsKnown() bool {
	switch r {
	case UserFirewallAccessRuleRuleListParamsOrderConfigurationTarget, UserFirewallAccessRuleRuleListParamsOrderConfigurationValue, UserFirewallAccessRuleRuleListParamsOrderMode:
		return true
	}
	return false
}
