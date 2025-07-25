// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/tidwall/gjson"
)

// ZoneFirewallWafPackageRuleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallWafPackageRuleService] method instead.
type ZoneFirewallWafPackageRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafPackageRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageRuleService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageRuleService) {
	r = &ZoneFirewallWafPackageRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageRuleService) Get(ctx context.Context, zoneID string, packageID string, ruleID string, opts ...option.RequestOption) (res *WafManagedRulesRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageRuleService) Update(ctx context.Context, zoneID string, packageID string, ruleID string, body ZoneFirewallWafPackageRuleUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", zoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *ZoneFirewallWafPackageRuleService) List(ctx context.Context, zoneID string, packageID string, query ZoneFirewallWafPackageRuleListParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WafManagedRulesAnomalyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []WafManagedRulesModeAnomaly `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WafManagedRulesAnomalyRuleGroup `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode WafManagedRulesModeAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                         `json:"priority,required"`
	JSON     wafManagedRulesAnomalyRuleJSON `json:"-"`
}

// wafManagedRulesAnomalyRuleJSON contains the JSON metadata for the struct
// [WafManagedRulesAnomalyRule]
type wafManagedRulesAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WafManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WafManagedRulesAnomalyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {}

func (r WafManagedRulesAnomalyRule) implementsZoneFirewallWafPackageRuleListResponseResult() {}

// Defines the rule group to which the current WAF rule belongs.
type WafManagedRulesAnomalyRuleGroup struct {
	// Defines the unique identifier of the rule group.
	ID string `json:"id"`
	// Defines the name of the rule group.
	Name string                              `json:"name"`
	JSON wafManagedRulesAnomalyRuleGroupJSON `json:"-"`
}

// wafManagedRulesAnomalyRuleGroupJSON contains the JSON metadata for the struct
// [WafManagedRulesAnomalyRuleGroup]
type wafManagedRulesAnomalyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesAnomalyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAnomalyRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WafManagedRulesModeAllowTraditional string

const (
	WafManagedRulesModeAllowTraditionalOn  WafManagedRulesModeAllowTraditional = "on"
	WafManagedRulesModeAllowTraditionalOff WafManagedRulesModeAllowTraditional = "off"
)

func (r WafManagedRulesModeAllowTraditional) IsKnown() bool {
	switch r {
	case WafManagedRulesModeAllowTraditionalOn, WafManagedRulesModeAllowTraditionalOff:
		return true
	}
	return false
}

// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
// when evaluating the request. Applies to anomaly detection WAF rules.
type WafManagedRulesModeAnomaly string

const (
	WafManagedRulesModeAnomalyOn  WafManagedRulesModeAnomaly = "on"
	WafManagedRulesModeAnomalyOff WafManagedRulesModeAnomaly = "off"
)

func (r WafManagedRulesModeAnomaly) IsKnown() bool {
	switch r {
	case WafManagedRulesModeAnomalyOn, WafManagedRulesModeAnomalyOff:
		return true
	}
	return false
}

// Defines the action that the current WAF rule will perform when triggered.
// Applies to traditional (deny) WAF rules.
type WafManagedRulesModeDenyTraditional string

const (
	WafManagedRulesModeDenyTraditionalDefault   WafManagedRulesModeDenyTraditional = "default"
	WafManagedRulesModeDenyTraditionalDisable   WafManagedRulesModeDenyTraditional = "disable"
	WafManagedRulesModeDenyTraditionalSimulate  WafManagedRulesModeDenyTraditional = "simulate"
	WafManagedRulesModeDenyTraditionalBlock     WafManagedRulesModeDenyTraditional = "block"
	WafManagedRulesModeDenyTraditionalChallenge WafManagedRulesModeDenyTraditional = "challenge"
)

func (r WafManagedRulesModeDenyTraditional) IsKnown() bool {
	switch r {
	case WafManagedRulesModeDenyTraditionalDefault, WafManagedRulesModeDenyTraditionalDisable, WafManagedRulesModeDenyTraditionalSimulate, WafManagedRulesModeDenyTraditionalBlock, WafManagedRulesModeDenyTraditionalChallenge:
		return true
	}
	return false
}

type WafManagedRulesRuleResponseSingle struct {
	Errors   []WafManagedRulesMessage `json:"errors,required"`
	Messages []WafManagedRulesMessage `json:"messages,required"`
	Result   interface{}              `json:"result,required"`
	// Defines whether the API call was successful.
	Success WafManagedRulesRuleResponseSingleSuccess `json:"success,required"`
	JSON    wafManagedRulesRuleResponseSingleJSON    `json:"-"`
}

// wafManagedRulesRuleResponseSingleJSON contains the JSON metadata for the struct
// [WafManagedRulesRuleResponseSingle]
type wafManagedRulesRuleResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesRuleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WafManagedRulesRuleResponseSingleSuccess bool

const (
	WafManagedRulesRuleResponseSingleSuccessTrue WafManagedRulesRuleResponseSingleSuccess = true
)

func (r WafManagedRulesRuleResponseSingleSuccess) IsKnown() bool {
	switch r {
	case WafManagedRulesRuleResponseSingleSuccessTrue:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WafManagedRulesTraditionalAllowRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WafManagedRulesModeAllowTraditional `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WafManagedRulesTraditionalAllowRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode WafManagedRulesModeAllowTraditional `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                  `json:"priority,required"`
	JSON     wafManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// wafManagedRulesTraditionalAllowRuleJSON contains the JSON metadata for the
// struct [WafManagedRulesTraditionalAllowRule]
type wafManagedRulesTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WafManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r WafManagedRulesTraditionalAllowRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

func (r WafManagedRulesTraditionalAllowRule) implementsZoneFirewallWafPackageRuleListResponseResult() {
}

// Defines the rule group to which the current WAF rule belongs.
type WafManagedRulesTraditionalAllowRuleGroup struct {
	// Defines the unique identifier of the rule group.
	ID string `json:"id"`
	// Defines the name of the rule group.
	Name string                                       `json:"name"`
	JSON wafManagedRulesTraditionalAllowRuleGroupJSON `json:"-"`
}

// wafManagedRulesTraditionalAllowRuleGroupJSON contains the JSON metadata for the
// struct [WafManagedRulesTraditionalAllowRuleGroup]
type wafManagedRulesTraditionalAllowRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesTraditionalAllowRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesTraditionalAllowRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type WafManagedRulesTraditionalDenyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the list of possible actions of the WAF rule when it is triggered.
	AllowedModes []WafManagedRulesModeDenyTraditional `json:"allowed_modes,required"`
	// Defines the default action/mode of a rule.
	DefaultMode WafManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WafManagedRulesTraditionalDenyRuleGroup `json:"group,required"`
	// Defines the action that the current WAF rule will perform when triggered.
	// Applies to traditional (deny) WAF rules.
	Mode WafManagedRulesModeDenyTraditional `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                 `json:"priority,required"`
	JSON     wafManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// wafManagedRulesTraditionalDenyRuleJSON contains the JSON metadata for the struct
// [WafManagedRulesTraditionalDenyRule]
type wafManagedRulesTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WafManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WafManagedRulesTraditionalDenyRule) implementsZoneFirewallWafPackageRuleUpdateResponseResult() {
}

func (r WafManagedRulesTraditionalDenyRule) implementsZoneFirewallWafPackageRuleListResponseResult() {
}

// Defines the default action/mode of a rule.
type WafManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WafManagedRulesTraditionalDenyRuleDefaultModeDisable   WafManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WafManagedRulesTraditionalDenyRuleDefaultModeSimulate  WafManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WafManagedRulesTraditionalDenyRuleDefaultModeBlock     WafManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WafManagedRulesTraditionalDenyRuleDefaultModeChallenge WafManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r WafManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case WafManagedRulesTraditionalDenyRuleDefaultModeDisable, WafManagedRulesTraditionalDenyRuleDefaultModeSimulate, WafManagedRulesTraditionalDenyRuleDefaultModeBlock, WafManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
}

// Defines the rule group to which the current WAF rule belongs.
type WafManagedRulesTraditionalDenyRuleGroup struct {
	// Defines the unique identifier of the rule group.
	ID string `json:"id"`
	// Defines the name of the rule group.
	Name string                                      `json:"name"`
	JSON wafManagedRulesTraditionalDenyRuleGroupJSON `json:"-"`
}

// wafManagedRulesTraditionalDenyRuleGroupJSON contains the JSON metadata for the
// struct [WafManagedRulesTraditionalDenyRuleGroup]
type wafManagedRulesTraditionalDenyRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesTraditionalDenyRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesTraditionalDenyRuleGroupJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafPackageRuleUpdateResponse struct {
	Errors   []WafManagedRulesMessage `json:"errors,required"`
	Messages []WafManagedRulesMessage `json:"messages,required"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result ZoneFirewallWafPackageRuleUpdateResponseResult `json:"result,required"`
	// Defines whether the API call was successful.
	Success ZoneFirewallWafPackageRuleUpdateResponseSuccess `json:"success,required"`
	JSON    zoneFirewallWafPackageRuleUpdateResponseJSON    `json:"-"`
}

// zoneFirewallWafPackageRuleUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageRuleUpdateResponse]
type zoneFirewallWafPackageRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type ZoneFirewallWafPackageRuleUpdateResponseResult struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// This field can have the runtime type of [[]WafManagedRulesModeAnomaly],
	// [[]WafManagedRulesModeDenyTraditional], [[]WafManagedRulesModeAllowTraditional].
	AllowedModes interface{} `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// This field can have the runtime type of [WafManagedRulesAnomalyRuleGroup],
	// [WafManagedRulesTraditionalDenyRuleGroup],
	// [WafManagedRulesTraditionalAllowRuleGroup].
	Group interface{} `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode WafManagedRulesModeAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string `json:"priority,required"`
	// Defines the default action/mode of a rule.
	DefaultMode ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode `json:"default_mode"`
	JSON        zoneFirewallWafPackageRuleUpdateResponseResultJSON        `json:"-"`
	union       ZoneFirewallWafPackageRuleUpdateResponseResultUnion
}

// zoneFirewallWafPackageRuleUpdateResponseResultJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageRuleUpdateResponseResult]
type zoneFirewallWafPackageRuleUpdateResponseResultJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r zoneFirewallWafPackageRuleUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneFirewallWafPackageRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneFirewallWafPackageRuleUpdateResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneFirewallWafPackageRuleUpdateResponseResultUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [WafManagedRulesAnomalyRule],
// [WafManagedRulesTraditionalDenyRule], [WafManagedRulesTraditionalAllowRule].
func (r ZoneFirewallWafPackageRuleUpdateResponseResult) AsUnion() ZoneFirewallWafPackageRuleUpdateResponseResultUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [WafManagedRulesAnomalyRule],
// [WafManagedRulesTraditionalDenyRule] or [WafManagedRulesTraditionalAllowRule].
type ZoneFirewallWafPackageRuleUpdateResponseResultUnion interface {
	implementsZoneFirewallWafPackageRuleUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageRuleUpdateResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesTraditionalAllowRule{}),
		},
	)
}

// Defines the default action/mode of a rule.
type ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode string

const (
	ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeDisable   ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode = "disable"
	ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeSimulate  ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeBlock     ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode = "block"
	ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeChallenge ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode = "challenge"
)

func (r ZoneFirewallWafPackageRuleUpdateResponseResultDefaultMode) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeDisable, ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeSimulate, ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeBlock, ZoneFirewallWafPackageRuleUpdateResponseResultDefaultModeChallenge:
		return true
	}
	return false
}

// Defines whether the API call was successful.
type ZoneFirewallWafPackageRuleUpdateResponseSuccess bool

const (
	ZoneFirewallWafPackageRuleUpdateResponseSuccessTrue ZoneFirewallWafPackageRuleUpdateResponseSuccess = true
)

func (r ZoneFirewallWafPackageRuleUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallWafPackageRuleListResponse struct {
	Errors   []WafManagedRulesMessage                       `json:"errors,required"`
	Messages []WafManagedRulesMessage                       `json:"messages,required"`
	Result   []ZoneFirewallWafPackageRuleListResponseResult `json:"result,required"`
	// Defines whether the API call was successful.
	Success    ZoneFirewallWafPackageRuleListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneFirewallWafPackageRuleListResponseResultInfo `json:"result_info"`
	JSON       zoneFirewallWafPackageRuleListResponseJSON       `json:"-"`
}

// zoneFirewallWafPackageRuleListResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageRuleListResponse]
type zoneFirewallWafPackageRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type ZoneFirewallWafPackageRuleListResponseResult struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// This field can have the runtime type of [[]WafManagedRulesModeAnomaly],
	// [[]WafManagedRulesModeDenyTraditional], [[]WafManagedRulesModeAllowTraditional].
	AllowedModes interface{} `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// This field can have the runtime type of [WafManagedRulesAnomalyRuleGroup],
	// [WafManagedRulesTraditionalDenyRuleGroup],
	// [WafManagedRulesTraditionalAllowRuleGroup].
	Group interface{} `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode WafManagedRulesModeAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string `json:"priority,required"`
	// Defines the default action/mode of a rule.
	DefaultMode ZoneFirewallWafPackageRuleListResponseResultDefaultMode `json:"default_mode"`
	JSON        zoneFirewallWafPackageRuleListResponseResultJSON        `json:"-"`
	union       ZoneFirewallWafPackageRuleListResponseResultUnion
}

// zoneFirewallWafPackageRuleListResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageRuleListResponseResult]
type zoneFirewallWafPackageRuleListResponseResultJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r zoneFirewallWafPackageRuleListResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r *ZoneFirewallWafPackageRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	*r = ZoneFirewallWafPackageRuleListResponseResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ZoneFirewallWafPackageRuleListResponseResultUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [WafManagedRulesAnomalyRule],
// [WafManagedRulesTraditionalDenyRule], [WafManagedRulesTraditionalAllowRule].
func (r ZoneFirewallWafPackageRuleListResponseResult) AsUnion() ZoneFirewallWafPackageRuleListResponseResultUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [WafManagedRulesAnomalyRule],
// [WafManagedRulesTraditionalDenyRule] or [WafManagedRulesTraditionalAllowRule].
type ZoneFirewallWafPackageRuleListResponseResultUnion interface {
	implementsZoneFirewallWafPackageRuleListResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneFirewallWafPackageRuleListResponseResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesTraditionalAllowRule{}),
		},
	)
}

// Defines the default action/mode of a rule.
type ZoneFirewallWafPackageRuleListResponseResultDefaultMode string

const (
	ZoneFirewallWafPackageRuleListResponseResultDefaultModeDisable   ZoneFirewallWafPackageRuleListResponseResultDefaultMode = "disable"
	ZoneFirewallWafPackageRuleListResponseResultDefaultModeSimulate  ZoneFirewallWafPackageRuleListResponseResultDefaultMode = "simulate"
	ZoneFirewallWafPackageRuleListResponseResultDefaultModeBlock     ZoneFirewallWafPackageRuleListResponseResultDefaultMode = "block"
	ZoneFirewallWafPackageRuleListResponseResultDefaultModeChallenge ZoneFirewallWafPackageRuleListResponseResultDefaultMode = "challenge"
)

func (r ZoneFirewallWafPackageRuleListResponseResultDefaultMode) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListResponseResultDefaultModeDisable, ZoneFirewallWafPackageRuleListResponseResultDefaultModeSimulate, ZoneFirewallWafPackageRuleListResponseResultDefaultModeBlock, ZoneFirewallWafPackageRuleListResponseResultDefaultModeChallenge:
		return true
	}
	return false
}

// Defines whether the API call was successful.
type ZoneFirewallWafPackageRuleListResponseSuccess bool

const (
	ZoneFirewallWafPackageRuleListResponseSuccessTrue ZoneFirewallWafPackageRuleListResponseSuccess = true
)

func (r ZoneFirewallWafPackageRuleListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallWafPackageRuleListResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                              `json:"total_count"`
	JSON       zoneFirewallWafPackageRuleListResponseResultInfoJSON `json:"-"`
}

// zoneFirewallWafPackageRuleListResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneFirewallWafPackageRuleListResponseResultInfo]
type zoneFirewallWafPackageRuleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageRuleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageRuleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafPackageRuleUpdateParams struct {
	// Defines the mode/action of the rule when triggered. You must use a value from
	// the `allowed_modes` array of the current rule.
	Mode param.Field[ZoneFirewallWafPackageRuleUpdateParamsMode] `json:"mode"`
}

func (r ZoneFirewallWafPackageRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the mode/action of the rule when triggered. You must use a value from
// the `allowed_modes` array of the current rule.
type ZoneFirewallWafPackageRuleUpdateParamsMode string

const (
	ZoneFirewallWafPackageRuleUpdateParamsModeDefault   ZoneFirewallWafPackageRuleUpdateParamsMode = "default"
	ZoneFirewallWafPackageRuleUpdateParamsModeDisable   ZoneFirewallWafPackageRuleUpdateParamsMode = "disable"
	ZoneFirewallWafPackageRuleUpdateParamsModeSimulate  ZoneFirewallWafPackageRuleUpdateParamsMode = "simulate"
	ZoneFirewallWafPackageRuleUpdateParamsModeBlock     ZoneFirewallWafPackageRuleUpdateParamsMode = "block"
	ZoneFirewallWafPackageRuleUpdateParamsModeChallenge ZoneFirewallWafPackageRuleUpdateParamsMode = "challenge"
	ZoneFirewallWafPackageRuleUpdateParamsModeOn        ZoneFirewallWafPackageRuleUpdateParamsMode = "on"
	ZoneFirewallWafPackageRuleUpdateParamsModeOff       ZoneFirewallWafPackageRuleUpdateParamsMode = "off"
)

func (r ZoneFirewallWafPackageRuleUpdateParamsMode) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleUpdateParamsModeDefault, ZoneFirewallWafPackageRuleUpdateParamsModeDisable, ZoneFirewallWafPackageRuleUpdateParamsModeSimulate, ZoneFirewallWafPackageRuleUpdateParamsModeBlock, ZoneFirewallWafPackageRuleUpdateParamsModeChallenge, ZoneFirewallWafPackageRuleUpdateParamsModeOn, ZoneFirewallWafPackageRuleUpdateParamsModeOff:
		return true
	}
	return false
}

type ZoneFirewallWafPackageRuleListParams struct {
	// Defines the public description of the WAF rule.
	Description param.Field[string] `query:"description"`
	// Defines the direction used to sort returned rules.
	Direction param.Field[ZoneFirewallWafPackageRuleListParamsDirection] `query:"direction"`
	// Defines the unique identifier of the rule group.
	GroupID param.Field[string] `query:"group_id"`
	// Defines the search requirements. When set to `all`, all the search requirements
	// must match. When set to `any`, only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageRuleListParamsMatch] `query:"match"`
	// Defines the action/mode a rule has been overridden to perform.
	Mode param.Field[ZoneFirewallWafPackageRuleListParamsMode] `query:"mode"`
	// Defines the field used to sort returned rules.
	Order param.Field[ZoneFirewallWafPackageRuleListParamsOrder] `query:"order"`
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority param.Field[string] `query:"priority"`
}

// URLQuery serializes [ZoneFirewallWafPackageRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallWafPackageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Defines the direction used to sort returned rules.
type ZoneFirewallWafPackageRuleListParamsDirection string

const (
	ZoneFirewallWafPackageRuleListParamsDirectionAsc  ZoneFirewallWafPackageRuleListParamsDirection = "asc"
	ZoneFirewallWafPackageRuleListParamsDirectionDesc ZoneFirewallWafPackageRuleListParamsDirection = "desc"
)

func (r ZoneFirewallWafPackageRuleListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListParamsDirectionAsc, ZoneFirewallWafPackageRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// Defines the search requirements. When set to `all`, all the search requirements
// must match. When set to `any`, only one of the search requirements has to match.
type ZoneFirewallWafPackageRuleListParamsMatch string

const (
	ZoneFirewallWafPackageRuleListParamsMatchAny ZoneFirewallWafPackageRuleListParamsMatch = "any"
	ZoneFirewallWafPackageRuleListParamsMatchAll ZoneFirewallWafPackageRuleListParamsMatch = "all"
)

func (r ZoneFirewallWafPackageRuleListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListParamsMatchAny, ZoneFirewallWafPackageRuleListParamsMatchAll:
		return true
	}
	return false
}

// Defines the action/mode a rule has been overridden to perform.
type ZoneFirewallWafPackageRuleListParamsMode string

const (
	ZoneFirewallWafPackageRuleListParamsModeDis ZoneFirewallWafPackageRuleListParamsMode = "DIS"
	ZoneFirewallWafPackageRuleListParamsModeChl ZoneFirewallWafPackageRuleListParamsMode = "CHL"
	ZoneFirewallWafPackageRuleListParamsModeBlk ZoneFirewallWafPackageRuleListParamsMode = "BLK"
	ZoneFirewallWafPackageRuleListParamsModeSim ZoneFirewallWafPackageRuleListParamsMode = "SIM"
)

func (r ZoneFirewallWafPackageRuleListParamsMode) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListParamsModeDis, ZoneFirewallWafPackageRuleListParamsModeChl, ZoneFirewallWafPackageRuleListParamsModeBlk, ZoneFirewallWafPackageRuleListParamsModeSim:
		return true
	}
	return false
}

// Defines the field used to sort returned rules.
type ZoneFirewallWafPackageRuleListParamsOrder string

const (
	ZoneFirewallWafPackageRuleListParamsOrderPriority    ZoneFirewallWafPackageRuleListParamsOrder = "priority"
	ZoneFirewallWafPackageRuleListParamsOrderGroupID     ZoneFirewallWafPackageRuleListParamsOrder = "group_id"
	ZoneFirewallWafPackageRuleListParamsOrderDescription ZoneFirewallWafPackageRuleListParamsOrder = "description"
)

func (r ZoneFirewallWafPackageRuleListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageRuleListParamsOrderPriority, ZoneFirewallWafPackageRuleListParamsOrderGroupID, ZoneFirewallWafPackageRuleListParamsOrderDescription:
		return true
	}
	return false
}
