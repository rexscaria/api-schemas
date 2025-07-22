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

// ZoneFirewallWafOverrideService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallWafOverrideService] method instead.
type ZoneFirewallWafOverrideService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafOverrideService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallWafOverrideService(opts ...option.RequestOption) (r *ZoneFirewallWafOverrideService) {
	r = &ZoneFirewallWafOverrideService{}
	r.Options = opts
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) New(ctx context.Context, zoneID string, body ZoneFirewallWafOverrideNewParams, opts ...option.RequestOption) (res *FirewallOverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Get(ctx context.Context, zoneID string, overridesID string, opts ...option.RequestOption) (res *FirewallOverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Update(ctx context.Context, zoneID string, overridesID string, body ZoneFirewallWafOverrideUpdateParams, opts ...option.RequestOption) (res *FirewallOverrideResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) List(ctx context.Context, zoneID string, query ZoneFirewallWafOverrideListParams, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafOverrideService) Delete(ctx context.Context, zoneID string, overridesID string, body ZoneFirewallWafOverrideDeleteParams, opts ...option.RequestOption) (res *ZoneFirewallWafOverrideDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", zoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type FirewallOverride struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the WAF package is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction FirewallRewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules map[string]FirewallOverrideRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []string             `json:"urls"`
	JSON firewallOverrideJSON `json:"-"`
}

// firewallOverrideJSON contains the JSON metadata for the struct
// [FirewallOverride]
type firewallOverrideJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallOverride) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallOverrideJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type FirewallOverrideRule string

const (
	FirewallOverrideRuleChallenge FirewallOverrideRule = "challenge"
	FirewallOverrideRuleBlock     FirewallOverrideRule = "block"
	FirewallOverrideRuleSimulate  FirewallOverrideRule = "simulate"
	FirewallOverrideRuleDisable   FirewallOverrideRule = "disable"
	FirewallOverrideRuleDefault   FirewallOverrideRule = "default"
)

func (r FirewallOverrideRule) IsKnown() bool {
	switch r {
	case FirewallOverrideRuleChallenge, FirewallOverrideRuleBlock, FirewallOverrideRuleSimulate, FirewallOverrideRuleDisable, FirewallOverrideRuleDefault:
		return true
	}
	return false
}

type FirewallOverrideResponseSingle struct {
	Result FirewallOverride                   `json:"result,required"`
	JSON   firewallOverrideResponseSingleJSON `json:"-"`
	FirewallAPIResponseSingle
}

// firewallOverrideResponseSingleJSON contains the JSON metadata for the struct
// [FirewallOverrideResponseSingle]
type firewallOverrideResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallOverrideResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallOverrideResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallRewriteAction struct {
	// The WAF rule action to apply.
	Block FirewallWafRewriteAction `json:"block"`
	// The WAF rule action to apply.
	Challenge FirewallWafRewriteAction `json:"challenge"`
	// The WAF rule action to apply.
	Default FirewallWafRewriteAction `json:"default"`
	// The WAF rule action to apply.
	Disable FirewallWafRewriteAction `json:"disable"`
	// The WAF rule action to apply.
	Simulate FirewallWafRewriteAction  `json:"simulate"`
	JSON     firewallRewriteActionJSON `json:"-"`
}

// firewallRewriteActionJSON contains the JSON metadata for the struct
// [FirewallRewriteAction]
type firewallRewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRewriteActionJSON) RawJSON() string {
	return r.raw
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type FirewallRewriteActionParam struct {
	// The WAF rule action to apply.
	Block param.Field[FirewallWafRewriteAction] `json:"block"`
	// The WAF rule action to apply.
	Challenge param.Field[FirewallWafRewriteAction] `json:"challenge"`
	// The WAF rule action to apply.
	Default param.Field[FirewallWafRewriteAction] `json:"default"`
	// The WAF rule action to apply.
	Disable param.Field[FirewallWafRewriteAction] `json:"disable"`
	// The WAF rule action to apply.
	Simulate param.Field[FirewallWafRewriteAction] `json:"simulate"`
}

func (r FirewallRewriteActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The WAF rule action to apply.
type FirewallWafRewriteAction string

const (
	FirewallWafRewriteActionChallenge FirewallWafRewriteAction = "challenge"
	FirewallWafRewriteActionBlock     FirewallWafRewriteAction = "block"
	FirewallWafRewriteActionSimulate  FirewallWafRewriteAction = "simulate"
	FirewallWafRewriteActionDisable   FirewallWafRewriteAction = "disable"
	FirewallWafRewriteActionDefault   FirewallWafRewriteAction = "default"
)

func (r FirewallWafRewriteAction) IsKnown() bool {
	switch r {
	case FirewallWafRewriteActionChallenge, FirewallWafRewriteActionBlock, FirewallWafRewriteActionSimulate, FirewallWafRewriteActionDisable, FirewallWafRewriteActionDefault:
		return true
	}
	return false
}

type ZoneFirewallWafOverrideListResponse struct {
	Result []ZoneFirewallWafOverrideListResponseResult `json:"result,required"`
	JSON   zoneFirewallWafOverrideListResponseJSON     `json:"-"`
	FirewallAPIResponseCollection
}

// zoneFirewallWafOverrideListResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideListResponse]
type zoneFirewallWafOverrideListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafOverrideListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafOverrideListResponseResult struct {
	JSON zoneFirewallWafOverrideListResponseResultJSON `json:"-"`
	FirewallOverride
}

// zoneFirewallWafOverrideListResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideListResponseResult]
type zoneFirewallWafOverrideListResponseResultJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafOverrideListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafOverrideDeleteResponse struct {
	Result ZoneFirewallWafOverrideDeleteResponseResult `json:"result"`
	JSON   zoneFirewallWafOverrideDeleteResponseJSON   `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafOverrideDeleteResponse]
type zoneFirewallWafOverrideDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafOverrideDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafOverrideDeleteResponseResult struct {
	// The unique identifier of the WAF override.
	ID   string                                          `json:"id"`
	JSON zoneFirewallWafOverrideDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallWafOverrideDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafOverrideDeleteResponseResult]
type zoneFirewallWafOverrideDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafOverrideDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafOverrideDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafOverrideNewParams struct {
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]string] `json:"urls,required"`
}

func (r ZoneFirewallWafOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallWafOverrideUpdateParams struct {
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction param.Field[FirewallRewriteActionParam] `json:"rewrite_action,required"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules param.Field[map[string]ZoneFirewallWafOverrideUpdateParamsRules] `json:"rules,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]string] `json:"urls,required"`
}

func (r ZoneFirewallWafOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The WAF rule action to apply.
type ZoneFirewallWafOverrideUpdateParamsRules string

const (
	ZoneFirewallWafOverrideUpdateParamsRulesChallenge ZoneFirewallWafOverrideUpdateParamsRules = "challenge"
	ZoneFirewallWafOverrideUpdateParamsRulesBlock     ZoneFirewallWafOverrideUpdateParamsRules = "block"
	ZoneFirewallWafOverrideUpdateParamsRulesSimulate  ZoneFirewallWafOverrideUpdateParamsRules = "simulate"
	ZoneFirewallWafOverrideUpdateParamsRulesDisable   ZoneFirewallWafOverrideUpdateParamsRules = "disable"
	ZoneFirewallWafOverrideUpdateParamsRulesDefault   ZoneFirewallWafOverrideUpdateParamsRules = "default"
)

func (r ZoneFirewallWafOverrideUpdateParamsRules) IsKnown() bool {
	switch r {
	case ZoneFirewallWafOverrideUpdateParamsRulesChallenge, ZoneFirewallWafOverrideUpdateParamsRulesBlock, ZoneFirewallWafOverrideUpdateParamsRulesSimulate, ZoneFirewallWafOverrideUpdateParamsRulesDisable, ZoneFirewallWafOverrideUpdateParamsRulesDefault:
		return true
	}
	return false
}

type ZoneFirewallWafOverrideListParams struct {
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallWafOverrideListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallWafOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallWafOverrideDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneFirewallWafOverrideDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
