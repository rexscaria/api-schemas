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

// ZoneEmailRoutingRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneEmailRoutingRuleService] method instead.
type ZoneEmailRoutingRuleService struct {
	Options  []option.RequestOption
	CatchAll *ZoneEmailRoutingRuleCatchAllService
}

// NewZoneEmailRoutingRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingRuleService(opts ...option.RequestOption) (r *ZoneEmailRoutingRuleService) {
	r = &ZoneEmailRoutingRuleService{}
	r.Options = opts
	r.CatchAll = NewZoneEmailRoutingRuleCatchAllService(opts...)
	return
}

// Rules consist of a set of criteria for matching emails (such as an email being
// sent to a specific custom email address) plus a set of actions to take on the
// email (like forwarding it to a specific destination address).
func (r *ZoneEmailRoutingRuleService) New(ctx context.Context, zoneID string, body ZoneEmailRoutingRuleNewParams, opts ...option.RequestOption) (res *EmailRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information for a specific routing rule already created.
func (r *ZoneEmailRoutingRuleService) Get(ctx context.Context, zoneID string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update actions and matches, or enable/disable specific routing rules.
func (r *ZoneEmailRoutingRuleService) Update(ctx context.Context, zoneID string, ruleIdentifier string, body ZoneEmailRoutingRuleUpdateParams, opts ...option.RequestOption) (res *EmailRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists existing routing rules.
func (r *ZoneEmailRoutingRuleService) List(ctx context.Context, zoneID string, query ZoneEmailRoutingRuleListParams, opts ...option.RequestOption) (res *ZoneEmailRoutingRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a specific routing rule.
func (r *ZoneEmailRoutingRuleService) Delete(ctx context.Context, zoneID string, ruleIdentifier string, opts ...option.RequestOption) (res *EmailRuleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleIdentifier == "" {
		err = errors.New("missing required rule_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/rules/%s", zoneID, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Actions pattern.
type EmailRuleActionPattern struct {
	// Type of supported action.
	Type  EmailRuleActionPatternType `json:"type,required"`
	Value []string                   `json:"value"`
	JSON  emailRuleActionPatternJSON `json:"-"`
}

// emailRuleActionPatternJSON contains the JSON metadata for the struct
// [EmailRuleActionPattern]
type emailRuleActionPatternJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleActionPattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleActionPatternJSON) RawJSON() string {
	return r.raw
}

// Type of supported action.
type EmailRuleActionPatternType string

const (
	EmailRuleActionPatternTypeDrop    EmailRuleActionPatternType = "drop"
	EmailRuleActionPatternTypeForward EmailRuleActionPatternType = "forward"
	EmailRuleActionPatternTypeWorker  EmailRuleActionPatternType = "worker"
)

func (r EmailRuleActionPatternType) IsKnown() bool {
	switch r {
	case EmailRuleActionPatternTypeDrop, EmailRuleActionPatternTypeForward, EmailRuleActionPatternTypeWorker:
		return true
	}
	return false
}

// Actions pattern.
type EmailRuleActionPatternParam struct {
	// Type of supported action.
	Type  param.Field[EmailRuleActionPatternType] `json:"type,required"`
	Value param.Field[[]string]                   `json:"value"`
}

func (r EmailRuleActionPatternParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Routing rule status.
type EmailRuleEnabled bool

const (
	EmailRuleEnabledTrue  EmailRuleEnabled = true
	EmailRuleEnabledFalse EmailRuleEnabled = false
)

func (r EmailRuleEnabled) IsKnown() bool {
	switch r {
	case EmailRuleEnabledTrue, EmailRuleEnabledFalse:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type EmailRuleMatcher struct {
	// Type of matcher.
	Type EmailRuleMatcherType `json:"type,required"`
	// Field for type matcher.
	Field EmailRuleMatcherField `json:"field"`
	// Value for matcher.
	Value string               `json:"value"`
	JSON  emailRuleMatcherJSON `json:"-"`
}

// emailRuleMatcherJSON contains the JSON metadata for the struct
// [EmailRuleMatcher]
type emailRuleMatcherJSON struct {
	Type        apijson.Field
	Field       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleMatcher) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleMatcherJSON) RawJSON() string {
	return r.raw
}

// Type of matcher.
type EmailRuleMatcherType string

const (
	EmailRuleMatcherTypeAll     EmailRuleMatcherType = "all"
	EmailRuleMatcherTypeLiteral EmailRuleMatcherType = "literal"
)

func (r EmailRuleMatcherType) IsKnown() bool {
	switch r {
	case EmailRuleMatcherTypeAll, EmailRuleMatcherTypeLiteral:
		return true
	}
	return false
}

// Field for type matcher.
type EmailRuleMatcherField string

const (
	EmailRuleMatcherFieldTo EmailRuleMatcherField = "to"
)

func (r EmailRuleMatcherField) IsKnown() bool {
	switch r {
	case EmailRuleMatcherFieldTo:
		return true
	}
	return false
}

// Matching pattern to forward your actions.
type EmailRuleMatcherParam struct {
	// Type of matcher.
	Type param.Field[EmailRuleMatcherType] `json:"type,required"`
	// Field for type matcher.
	Field param.Field[EmailRuleMatcherField] `json:"field"`
	// Value for matcher.
	Value param.Field[string] `json:"value"`
}

func (r EmailRuleMatcherParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRuleResponseSingle struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success EmailRuleResponseSingleSuccess `json:"success,required"`
	Result  EmailRules                     `json:"result"`
	JSON    emailRuleResponseSingleJSON    `json:"-"`
}

// emailRuleResponseSingleJSON contains the JSON metadata for the struct
// [EmailRuleResponseSingle]
type emailRuleResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRuleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRuleResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EmailRuleResponseSingleSuccess bool

const (
	EmailRuleResponseSingleSuccessTrue EmailRuleResponseSingleSuccess = true
)

func (r EmailRuleResponseSingleSuccess) IsKnown() bool {
	switch r {
	case EmailRuleResponseSingleSuccessTrue:
		return true
	}
	return false
}

type EmailRules struct {
	// Routing rule identifier.
	ID string `json:"id"`
	// List actions patterns.
	Actions []EmailRuleActionPattern `json:"actions"`
	// Routing rule status.
	Enabled EmailRuleEnabled `json:"enabled"`
	// Matching patterns to forward to your actions.
	Matchers []EmailRuleMatcher `json:"matchers"`
	// Routing rule name.
	Name string `json:"name"`
	// Priority of the routing rule.
	Priority float64 `json:"priority"`
	// Routing rule tag. (Deprecated, replaced by routing rule identifier)
	//
	// Deprecated: deprecated
	Tag  string         `json:"tag"`
	JSON emailRulesJSON `json:"-"`
}

// emailRulesJSON contains the JSON metadata for the struct [EmailRules]
type emailRulesJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	Enabled     apijson.Field
	Matchers    apijson.Field
	Name        apijson.Field
	Priority    apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRulesJSON) RawJSON() string {
	return r.raw
}

type ZoneEmailRoutingRuleListResponse struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneEmailRoutingRuleListResponseSuccess    `json:"success,required"`
	Result     []EmailRules                               `json:"result"`
	ResultInfo ZoneEmailRoutingRuleListResponseResultInfo `json:"result_info"`
	JSON       zoneEmailRoutingRuleListResponseJSON       `json:"-"`
}

// zoneEmailRoutingRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneEmailRoutingRuleListResponse]
type zoneEmailRoutingRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneEmailRoutingRuleListResponseSuccess bool

const (
	ZoneEmailRoutingRuleListResponseSuccessTrue ZoneEmailRoutingRuleListResponseSuccess = true
)

func (r ZoneEmailRoutingRuleListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneEmailRoutingRuleListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneEmailRoutingRuleListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       zoneEmailRoutingRuleListResponseResultInfoJSON `json:"-"`
}

// zoneEmailRoutingRuleListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneEmailRoutingRuleListResponseResultInfo]
type zoneEmailRoutingRuleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEmailRoutingRuleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEmailRoutingRuleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneEmailRoutingRuleNewParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRuleActionPatternParam] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRuleMatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRuleEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneEmailRoutingRuleUpdateParams struct {
	// List actions patterns.
	Actions param.Field[[]EmailRuleActionPatternParam] `json:"actions,required"`
	// Matching patterns to forward to your actions.
	Matchers param.Field[[]EmailRuleMatcherParam] `json:"matchers,required"`
	// Routing rule status.
	Enabled param.Field[EmailRuleEnabled] `json:"enabled"`
	// Routing rule name.
	Name param.Field[string] `json:"name"`
	// Priority of the routing rule.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneEmailRoutingRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneEmailRoutingRuleListParams struct {
	// Filter by enabled routing rules.
	Enabled param.Field[ZoneEmailRoutingRuleListParamsEnabled] `query:"enabled"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneEmailRoutingRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneEmailRoutingRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by enabled routing rules.
type ZoneEmailRoutingRuleListParamsEnabled bool

const (
	ZoneEmailRoutingRuleListParamsEnabledTrue  ZoneEmailRoutingRuleListParamsEnabled = true
	ZoneEmailRoutingRuleListParamsEnabledFalse ZoneEmailRoutingRuleListParamsEnabled = false
)

func (r ZoneEmailRoutingRuleListParamsEnabled) IsKnown() bool {
	switch r {
	case ZoneEmailRoutingRuleListParamsEnabledTrue, ZoneEmailRoutingRuleListParamsEnabledFalse:
		return true
	}
	return false
}
