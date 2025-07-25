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

// ZoneFirewallRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallRuleService] method instead.
type ZoneFirewallRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallRuleService(opts ...option.RequestOption) (r *ZoneFirewallRuleService) {
	r = &ZoneFirewallRuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) New(ctx context.Context, zoneID string, body ZoneFirewallRuleNewParams, opts ...option.RequestOption) (res *FirewallFilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a firewall rule.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) Get(ctx context.Context, zoneID string, ruleID string, opts ...option.RequestOption) (res *FirewallFilterRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing firewall rule.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) Update(ctx context.Context, zoneID string, ruleID string, body ZoneFirewallRuleUpdateParams, opts ...option.RequestOption) (res *FirewallFilterRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) List(ctx context.Context, zoneID string, query ZoneFirewallRuleListParams, opts ...option.RequestOption) (res *FirewallFilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing firewall rule.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) Delete(ctx context.Context, zoneID string, ruleID string, opts ...option.RequestOption) (res *ZoneFirewallRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Updates the priority of an existing firewall rule.
//
// Deprecated: deprecated
func (r *ZoneFirewallRuleService) UpdatePriority(ctx context.Context, zoneID string, ruleID string, body ZoneFirewallRuleUpdatePriorityParams, opts ...option.RequestOption) (res *FirewallFilterRulesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type FirewallAction struct {
	// The action to perform.
	Mode FirewallActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response FirewallActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64            `json:"timeout"`
	JSON    firewallActionJSON `json:"-"`
}

// firewallActionJSON contains the JSON metadata for the struct [FirewallAction]
type firewallActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallActionJSON) RawJSON() string {
	return r.raw
}

// The action to perform.
type FirewallActionMode string

const (
	FirewallActionModeSimulate         FirewallActionMode = "simulate"
	FirewallActionModeBan              FirewallActionMode = "ban"
	FirewallActionModeChallenge        FirewallActionMode = "challenge"
	FirewallActionModeJsChallenge      FirewallActionMode = "js_challenge"
	FirewallActionModeManagedChallenge FirewallActionMode = "managed_challenge"
)

func (r FirewallActionMode) IsKnown() bool {
	switch r {
	case FirewallActionModeSimulate, FirewallActionModeBan, FirewallActionModeChallenge, FirewallActionModeJsChallenge, FirewallActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type FirewallActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                     `json:"content_type"`
	JSON        firewallActionResponseJSON `json:"-"`
}

// firewallActionResponseJSON contains the JSON metadata for the struct
// [FirewallActionResponse]
type firewallActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallActionResponseJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type FirewallActionParam struct {
	// The action to perform.
	Mode param.Field[FirewallActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[FirewallActionResponseParam] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r FirewallActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type FirewallActionResponseParam struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r FirewallActionResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string             `json:"ref"`
	JSON firewallFilterJSON `json:"-"`
}

// firewallFilterJSON contains the JSON metadata for the struct [FirewallFilter]
type firewallFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallFilter) implementsFirewallFilterRuleResponseFilter() {}

type FirewallFilterParam struct {
	// An informative summary of the filter.
	Description param.Field[string] `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression param.Field[string] `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref param.Field[string] `json:"ref"`
}

func (r FirewallFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallFilterRuleResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action FirewallFilterRuleResponseAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                           `json:"description"`
	Filter      FirewallFilterRuleResponseFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                             `json:"priority"`
	Products []FirewallFilterRuleResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                         `json:"ref"`
	JSON firewallFilterRuleResponseJSON `json:"-"`
}

// firewallFilterRuleResponseJSON contains the JSON metadata for the struct
// [FirewallFilterRuleResponse]
type firewallFilterRuleResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRuleResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type FirewallFilterRuleResponseAction string

const (
	FirewallFilterRuleResponseActionBlock            FirewallFilterRuleResponseAction = "block"
	FirewallFilterRuleResponseActionChallenge        FirewallFilterRuleResponseAction = "challenge"
	FirewallFilterRuleResponseActionJsChallenge      FirewallFilterRuleResponseAction = "js_challenge"
	FirewallFilterRuleResponseActionManagedChallenge FirewallFilterRuleResponseAction = "managed_challenge"
	FirewallFilterRuleResponseActionAllow            FirewallFilterRuleResponseAction = "allow"
	FirewallFilterRuleResponseActionLog              FirewallFilterRuleResponseAction = "log"
	FirewallFilterRuleResponseActionBypass           FirewallFilterRuleResponseAction = "bypass"
)

func (r FirewallFilterRuleResponseAction) IsKnown() bool {
	switch r {
	case FirewallFilterRuleResponseActionBlock, FirewallFilterRuleResponseActionChallenge, FirewallFilterRuleResponseActionJsChallenge, FirewallFilterRuleResponseActionManagedChallenge, FirewallFilterRuleResponseActionAllow, FirewallFilterRuleResponseActionLog, FirewallFilterRuleResponseActionBypass:
		return true
	}
	return false
}

type FirewallFilterRuleResponseFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool `json:"deleted"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref   string                               `json:"ref"`
	JSON  firewallFilterRuleResponseFilterJSON `json:"-"`
	union FirewallFilterRuleResponseFilterUnion
}

// firewallFilterRuleResponseFilterJSON contains the JSON metadata for the struct
// [FirewallFilterRuleResponseFilter]
type firewallFilterRuleResponseFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallFilterRuleResponseFilterJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallFilterRuleResponseFilter) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallFilterRuleResponseFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallFilterRuleResponseFilterUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [FirewallFilter],
// [FirewallFilterRuleResponseFilterFirewallDeletedFilter].
func (r FirewallFilterRuleResponseFilter) AsUnion() FirewallFilterRuleResponseFilterUnion {
	return r.union
}

// Union satisfied by [FirewallFilter] or
// [FirewallFilterRuleResponseFilterFirewallDeletedFilter].
type FirewallFilterRuleResponseFilterUnion interface {
	implementsFirewallFilterRuleResponseFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallFilterRuleResponseFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallFilterRuleResponseFilterFirewallDeletedFilter{}),
		},
	)
}

type FirewallFilterRuleResponseFilterFirewallDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                      `json:"deleted,required"`
	JSON    firewallFilterRuleResponseFilterFirewallDeletedFilterJSON `json:"-"`
}

// firewallFilterRuleResponseFilterFirewallDeletedFilterJSON contains the JSON
// metadata for the struct [FirewallFilterRuleResponseFilterFirewallDeletedFilter]
type firewallFilterRuleResponseFilterFirewallDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRuleResponseFilterFirewallDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRuleResponseFilterFirewallDeletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r FirewallFilterRuleResponseFilterFirewallDeletedFilter) implementsFirewallFilterRuleResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type FirewallFilterRuleResponseProduct string

const (
	FirewallFilterRuleResponseProductZoneLockdown  FirewallFilterRuleResponseProduct = "zoneLockdown"
	FirewallFilterRuleResponseProductUaBlock       FirewallFilterRuleResponseProduct = "uaBlock"
	FirewallFilterRuleResponseProductBic           FirewallFilterRuleResponseProduct = "bic"
	FirewallFilterRuleResponseProductHot           FirewallFilterRuleResponseProduct = "hot"
	FirewallFilterRuleResponseProductSecurityLevel FirewallFilterRuleResponseProduct = "securityLevel"
	FirewallFilterRuleResponseProductRateLimit     FirewallFilterRuleResponseProduct = "rateLimit"
	FirewallFilterRuleResponseProductWaf           FirewallFilterRuleResponseProduct = "waf"
)

func (r FirewallFilterRuleResponseProduct) IsKnown() bool {
	switch r {
	case FirewallFilterRuleResponseProductZoneLockdown, FirewallFilterRuleResponseProductUaBlock, FirewallFilterRuleResponseProductBic, FirewallFilterRuleResponseProductHot, FirewallFilterRuleResponseProductSecurityLevel, FirewallFilterRuleResponseProductRateLimit, FirewallFilterRuleResponseProductWaf:
		return true
	}
	return false
}

type FirewallFilterRulesResponseCollection struct {
	Errors   []FirewallFilterRulesResponseCollectionError   `json:"errors,required"`
	Messages []FirewallFilterRulesResponseCollectionMessage `json:"messages,required"`
	Result   []FirewallFilterRuleResponse                   `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    FirewallFilterRulesResponseCollectionSuccess    `json:"success,required"`
	ResultInfo FirewallFilterRulesResponseCollectionResultInfo `json:"result_info"`
	JSON       firewallFilterRulesResponseCollectionJSON       `json:"-"`
}

// firewallFilterRulesResponseCollectionJSON contains the JSON metadata for the
// struct [FirewallFilterRulesResponseCollection]
type firewallFilterRulesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesResponseCollectionError struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           FirewallFilterRulesResponseCollectionErrorsSource `json:"source"`
	JSON             firewallFilterRulesResponseCollectionErrorJSON    `json:"-"`
}

// firewallFilterRulesResponseCollectionErrorJSON contains the JSON metadata for
// the struct [FirewallFilterRulesResponseCollectionError]
type firewallFilterRulesResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesResponseCollectionErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    firewallFilterRulesResponseCollectionErrorsSourceJSON `json:"-"`
}

// firewallFilterRulesResponseCollectionErrorsSourceJSON contains the JSON metadata
// for the struct [FirewallFilterRulesResponseCollectionErrorsSource]
type firewallFilterRulesResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesResponseCollectionMessage struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           FirewallFilterRulesResponseCollectionMessagesSource `json:"source"`
	JSON             firewallFilterRulesResponseCollectionMessageJSON    `json:"-"`
}

// firewallFilterRulesResponseCollectionMessageJSON contains the JSON metadata for
// the struct [FirewallFilterRulesResponseCollectionMessage]
type firewallFilterRulesResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesResponseCollectionMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    firewallFilterRulesResponseCollectionMessagesSourceJSON `json:"-"`
}

// firewallFilterRulesResponseCollectionMessagesSourceJSON contains the JSON
// metadata for the struct [FirewallFilterRulesResponseCollectionMessagesSource]
type firewallFilterRulesResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallFilterRulesResponseCollectionSuccess bool

const (
	FirewallFilterRulesResponseCollectionSuccessTrue FirewallFilterRulesResponseCollectionSuccess = true
)

func (r FirewallFilterRulesResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case FirewallFilterRulesResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type FirewallFilterRulesResponseCollectionResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                             `json:"total_count"`
	JSON       firewallFilterRulesResponseCollectionResultInfoJSON `json:"-"`
}

// firewallFilterRulesResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [FirewallFilterRulesResponseCollectionResultInfo]
type firewallFilterRulesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesSingleResponse struct {
	Errors   []FirewallFilterRulesSingleResponseError   `json:"errors,required"`
	Messages []FirewallFilterRulesSingleResponseMessage `json:"messages,required"`
	Result   FirewallFilterRuleResponse                 `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallFilterRulesSingleResponseSuccess `json:"success,required"`
	JSON    firewallFilterRulesSingleResponseJSON    `json:"-"`
}

// firewallFilterRulesSingleResponseJSON contains the JSON metadata for the struct
// [FirewallFilterRulesSingleResponse]
type firewallFilterRulesSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesSingleResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesSingleResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           FirewallFilterRulesSingleResponseErrorsSource `json:"source"`
	JSON             firewallFilterRulesSingleResponseErrorJSON    `json:"-"`
}

// firewallFilterRulesSingleResponseErrorJSON contains the JSON metadata for the
// struct [FirewallFilterRulesSingleResponseError]
type firewallFilterRulesSingleResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterRulesSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesSingleResponseErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesSingleResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    firewallFilterRulesSingleResponseErrorsSourceJSON `json:"-"`
}

// firewallFilterRulesSingleResponseErrorsSourceJSON contains the JSON metadata for
// the struct [FirewallFilterRulesSingleResponseErrorsSource]
type firewallFilterRulesSingleResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesSingleResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesSingleResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesSingleResponseMessage struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           FirewallFilterRulesSingleResponseMessagesSource `json:"source"`
	JSON             firewallFilterRulesSingleResponseMessageJSON    `json:"-"`
}

// firewallFilterRulesSingleResponseMessageJSON contains the JSON metadata for the
// struct [FirewallFilterRulesSingleResponseMessage]
type firewallFilterRulesSingleResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallFilterRulesSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesSingleResponseMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallFilterRulesSingleResponseMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    firewallFilterRulesSingleResponseMessagesSourceJSON `json:"-"`
}

// firewallFilterRulesSingleResponseMessagesSourceJSON contains the JSON metadata
// for the struct [FirewallFilterRulesSingleResponseMessagesSource]
type firewallFilterRulesSingleResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallFilterRulesSingleResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallFilterRulesSingleResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallFilterRulesSingleResponseSuccess bool

const (
	FirewallFilterRulesSingleResponseSuccessTrue FirewallFilterRulesSingleResponseSuccess = true
)

func (r FirewallFilterRulesSingleResponseSuccess) IsKnown() bool {
	switch r {
	case FirewallFilterRulesSingleResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallRuleDeleteResponse struct {
	Errors   []ZoneFirewallRuleDeleteResponseError   `json:"errors,required"`
	Messages []ZoneFirewallRuleDeleteResponseMessage `json:"messages,required"`
	Result   FirewallFilterRuleResponse              `json:"result,required"`
	// Defines whether the API call was successful.
	Success ZoneFirewallRuleDeleteResponseSuccess `json:"success,required"`
	JSON    zoneFirewallRuleDeleteResponseJSON    `json:"-"`
}

// zoneFirewallRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleDeleteResponse]
type zoneFirewallRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallRuleDeleteResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ZoneFirewallRuleDeleteResponseErrorsSource `json:"source"`
	JSON             zoneFirewallRuleDeleteResponseErrorJSON    `json:"-"`
}

// zoneFirewallRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleDeleteResponseError]
type zoneFirewallRuleDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallRuleDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallRuleDeleteResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    zoneFirewallRuleDeleteResponseErrorsSourceJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZoneFirewallRuleDeleteResponseErrorsSource]
type zoneFirewallRuleDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallRuleDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallRuleDeleteResponseMessage struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZoneFirewallRuleDeleteResponseMessagesSource `json:"source"`
	JSON             zoneFirewallRuleDeleteResponseMessageJSON    `json:"-"`
}

// zoneFirewallRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleDeleteResponseMessage]
type zoneFirewallRuleDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallRuleDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallRuleDeleteResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zoneFirewallRuleDeleteResponseMessagesSourceJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneFirewallRuleDeleteResponseMessagesSource]
type zoneFirewallRuleDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallRuleDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type ZoneFirewallRuleDeleteResponseSuccess bool

const (
	ZoneFirewallRuleDeleteResponseSuccessTrue ZoneFirewallRuleDeleteResponseSuccess = true
)

func (r ZoneFirewallRuleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallRuleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallRuleNewParams struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[FirewallActionParam] `json:"action,required"`
	Filter param.Field[FirewallFilterParam] `json:"filter,required"`
}

func (r ZoneFirewallRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallRuleUpdateParams struct {
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[FirewallActionParam] `json:"action,required"`
	Filter param.Field[FirewallFilterParam] `json:"filter,required"`
}

func (r ZoneFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallRuleListParams struct {
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallRuleUpdatePriorityParams struct {
}

func (r ZoneFirewallRuleUpdatePriorityParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
