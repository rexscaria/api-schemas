// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
	"github.com/stainless-sdks/cf-rex-go/shared"
	"github.com/tidwall/gjson"
)

// ZoneFirewallWafPackageGroupService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallWafPackageGroupService] method instead.
type ZoneFirewallWafPackageGroupService struct {
	Options []option.RequestOption
}

// NewZoneFirewallWafPackageGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneFirewallWafPackageGroupService(opts ...option.RequestOption) (r *ZoneFirewallWafPackageGroupService) {
	r = &ZoneFirewallWafPackageGroupService{}
	r.Options = opts
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) Get(ctx context.Context, zoneID string, packageID string, groupID string, opts ...option.RequestOption) (res *WafManagedRulesRuleGroupResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) Update(ctx context.Context, zoneID string, packageID string, groupID string, body ZoneFirewallWafPackageGroupUpdateParams, opts ...option.RequestOption) (res *WafManagedRulesRuleGroupResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", zoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
func (r *ZoneFirewallWafPackageGroupService) List(ctx context.Context, zoneID string, packageID string, query ZoneFirewallWafPackageGroupListParams, opts ...option.RequestOption) (res *ZoneFirewallWafPackageGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", zoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WafManagedRulesAPIResponseCollection struct {
	Result     []interface{}                                  `json:"result,nullable"`
	ResultInfo WafManagedRulesAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       wafManagedRulesAPIResponseCollectionJSON       `json:"-"`
	WafManagedRulesAPIResponseCommon
}

// wafManagedRulesAPIResponseCollectionJSON contains the JSON metadata for the
// struct [WafManagedRulesAPIResponseCollection]
type wafManagedRulesAPIResponseCollectionJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type WafManagedRulesAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       wafManagedRulesAPIResponseCollectionResultInfoJSON `json:"-"`
}

// wafManagedRulesAPIResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [WafManagedRulesAPIResponseCollectionResultInfo]
type wafManagedRulesAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type WafManagedRulesAPIResponseCommon struct {
	Errors   []WafManagedRulesMessage                    `json:"errors,required"`
	Messages []WafManagedRulesMessage                    `json:"messages,required"`
	Result   WafManagedRulesAPIResponseCommonResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success WafManagedRulesAPIResponseCommonSuccess `json:"success,required"`
	JSON    wafManagedRulesAPIResponseCommonJSON    `json:"-"`
}

// wafManagedRulesAPIResponseCommonJSON contains the JSON metadata for the struct
// [WafManagedRulesAPIResponseCommon]
type wafManagedRulesAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [WafManagedRulesAPIResponseCommonResultArray] or
// [shared.UnionString].
type WafManagedRulesAPIResponseCommonResultUnion interface {
	ImplementsWafManagedRulesAPIResponseCommonResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WafManagedRulesAPIResponseCommonResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WafManagedRulesAPIResponseCommonResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WafManagedRulesAPIResponseCommonResultArray []interface{}

func (r WafManagedRulesAPIResponseCommonResultArray) ImplementsWafManagedRulesAPIResponseCommonResultUnion() {
}

// Whether the API call was successful
type WafManagedRulesAPIResponseCommonSuccess bool

const (
	WafManagedRulesAPIResponseCommonSuccessTrue WafManagedRulesAPIResponseCommonSuccess = true
)

func (r WafManagedRulesAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case WafManagedRulesAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type WafManagedRulesAPIResponseSingle struct {
	Result interface{}                          `json:"result"`
	JSON   wafManagedRulesAPIResponseSingleJSON `json:"-"`
	WafManagedRulesAPIResponseCommon
}

// wafManagedRulesAPIResponseSingleJSON contains the JSON metadata for the struct
// [WafManagedRulesAPIResponseSingle]
type wafManagedRulesAPIResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

type WafManagedRulesMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    wafManagedRulesMessageJSON `json:"-"`
}

// wafManagedRulesMessageJSON contains the JSON metadata for the struct
// [WafManagedRulesMessage]
type wafManagedRulesMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesMessageJSON) RawJSON() string {
	return r.raw
}

// The state of the rules contained in the rule group. When `on`, the rules in the
// group are configurable/usable.
type WafManagedRulesMode string

const (
	WafManagedRulesModeOn  WafManagedRulesMode = "on"
	WafManagedRulesModeOff WafManagedRulesMode = "off"
)

func (r WafManagedRulesMode) IsKnown() bool {
	switch r {
	case WafManagedRulesModeOn, WafManagedRulesModeOff:
		return true
	}
	return false
}

type WafManagedRulesRuleGroupResponseSingle struct {
	Result interface{}                                `json:"result"`
	JSON   wafManagedRulesRuleGroupResponseSingleJSON `json:"-"`
	WafManagedRulesAPIResponseSingle
}

// wafManagedRulesRuleGroupResponseSingleJSON contains the JSON metadata for the
// struct [WafManagedRulesRuleGroupResponseSingle]
type wafManagedRulesRuleGroupResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WafManagedRulesRuleGroupResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafManagedRulesRuleGroupResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafPackageGroupListResponse struct {
	Result []ZoneFirewallWafPackageGroupListResponseResult `json:"result"`
	JSON   zoneFirewallWafPackageGroupListResponseJSON     `json:"-"`
	WafManagedRulesAPIResponseCollection
}

// zoneFirewallWafPackageGroupListResponseJSON contains the JSON metadata for the
// struct [ZoneFirewallWafPackageGroupListResponse]
type zoneFirewallWafPackageGroupListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafPackageGroupListResponseResult struct {
	// The unique identifier of the rule group.
	ID string `json:"id"`
	// The available states for the rule group.
	AllowedModes []WafManagedRulesMode `json:"allowed_modes"`
	// An informative summary of what the rule group does.
	Description string `json:"description,nullable"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode WafManagedRulesMode `json:"mode"`
	// The number of rules within the group that have been modified from their default
	// configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// The name of the rule group.
	Name string `json:"name"`
	// The unique identifier of a WAF package.
	PackageID string `json:"package_id"`
	// The number of rules in the current rule group.
	RulesCount float64                                           `json:"rules_count"`
	JSON       zoneFirewallWafPackageGroupListResponseResultJSON `json:"-"`
}

// zoneFirewallWafPackageGroupListResponseResultJSON contains the JSON metadata for
// the struct [ZoneFirewallWafPackageGroupListResponseResult]
type zoneFirewallWafPackageGroupListResponseResultJSON struct {
	ID                 apijson.Field
	AllowedModes       apijson.Field
	Description        apijson.Field
	Mode               apijson.Field
	ModifiedRulesCount apijson.Field
	Name               apijson.Field
	PackageID          apijson.Field
	RulesCount         apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneFirewallWafPackageGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallWafPackageGroupListResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallWafPackageGroupUpdateParams struct {
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[WafManagedRulesMode] `json:"mode"`
}

func (r ZoneFirewallWafPackageGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallWafPackageGroupListParams struct {
	// The direction used to sort returned rule groups.
	Direction param.Field[ZoneFirewallWafPackageGroupListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZoneFirewallWafPackageGroupListParamsMatch] `query:"match"`
	// The state of the rules contained in the rule group. When `on`, the rules in the
	// group are configurable/usable.
	Mode param.Field[WafManagedRulesMode] `query:"mode"`
	// The name of the rule group.
	Name param.Field[string] `query:"name"`
	// The field used to sort returned rule groups.
	Order param.Field[ZoneFirewallWafPackageGroupListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of rule groups per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The number of rules in the current rule group.
	RulesCount param.Field[float64] `query:"rules_count"`
}

// URLQuery serializes [ZoneFirewallWafPackageGroupListParams]'s query parameters
// as `url.Values`.
func (r ZoneFirewallWafPackageGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned rule groups.
type ZoneFirewallWafPackageGroupListParamsDirection string

const (
	ZoneFirewallWafPackageGroupListParamsDirectionAsc  ZoneFirewallWafPackageGroupListParamsDirection = "asc"
	ZoneFirewallWafPackageGroupListParamsDirectionDesc ZoneFirewallWafPackageGroupListParamsDirection = "desc"
)

func (r ZoneFirewallWafPackageGroupListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageGroupListParamsDirectionAsc, ZoneFirewallWafPackageGroupListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZoneFirewallWafPackageGroupListParamsMatch string

const (
	ZoneFirewallWafPackageGroupListParamsMatchAny ZoneFirewallWafPackageGroupListParamsMatch = "any"
	ZoneFirewallWafPackageGroupListParamsMatchAll ZoneFirewallWafPackageGroupListParamsMatch = "all"
)

func (r ZoneFirewallWafPackageGroupListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageGroupListParamsMatchAny, ZoneFirewallWafPackageGroupListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned rule groups.
type ZoneFirewallWafPackageGroupListParamsOrder string

const (
	ZoneFirewallWafPackageGroupListParamsOrderMode       ZoneFirewallWafPackageGroupListParamsOrder = "mode"
	ZoneFirewallWafPackageGroupListParamsOrderRulesCount ZoneFirewallWafPackageGroupListParamsOrder = "rules_count"
)

func (r ZoneFirewallWafPackageGroupListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneFirewallWafPackageGroupListParamsOrderMode, ZoneFirewallWafPackageGroupListParamsOrderRulesCount:
		return true
	}
	return false
}
