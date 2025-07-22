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

// ZoneFirewallUaRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallUaRuleService] method instead.
type ZoneFirewallUaRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallUaRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallUaRuleService(opts ...option.RequestOption) (r *ZoneFirewallUaRuleService) {
	r = &ZoneFirewallUaRuleService{}
	r.Options = opts
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *ZoneFirewallUaRuleService) New(ctx context.Context, zoneID string, body ZoneFirewallUaRuleNewParams, opts ...option.RequestOption) (res *FirewallUablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Get(ctx context.Context, zoneID string, uaRuleID string, opts ...option.RequestOption) (res *FirewallUablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Update(ctx context.Context, zoneID string, uaRuleID string, body ZoneFirewallUaRuleUpdateParams, opts ...option.RequestOption) (res *FirewallUablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *ZoneFirewallUaRuleService) List(ctx context.Context, zoneID string, query ZoneFirewallUaRuleListParams, opts ...option.RequestOption) (res *ZoneFirewallUaRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Delete(ctx context.Context, zoneID string, uaRuleID string, body ZoneFirewallUaRuleDeleteParams, opts ...option.RequestOption) (res *ZoneFirewallUaRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if uaRuleID == "" {
		err = errors.New("missing required ua_rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneID, uaRuleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type FirewallUablockResponseSingle struct {
	Result interface{}                       `json:"result"`
	JSON   firewallUablockResponseSingleJSON `json:"-"`
	FirewallAPIResponseSingle
}

// firewallUablockResponseSingleJSON contains the JSON metadata for the struct
// [FirewallUablockResponseSingle]
type firewallUablockResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUablockResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponse struct {
	Result []ZoneFirewallUaRuleListResponseResult `json:"result"`
	JSON   zoneFirewallUaRuleListResponseJSON     `json:"-"`
	FirewallAPIResponseCollection
}

// zoneFirewallUaRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleListResponse]
type zoneFirewallUaRuleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponseResult struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration ZoneFirewallUaRuleListResponseResultConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode ZoneFirewallUaRuleListResponseResultMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                     `json:"paused"`
	JSON   zoneFirewallUaRuleListResponseResultJSON `json:"-"`
}

// zoneFirewallUaRuleListResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleListResponseResult]
type zoneFirewallUaRuleListResponseResultJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseResultJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type ZoneFirewallUaRuleListResponseResultConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                                `json:"value"`
	JSON  zoneFirewallUaRuleListResponseResultConfigurationJSON `json:"-"`
}

// zoneFirewallUaRuleListResponseResultConfigurationJSON contains the JSON metadata
// for the struct [ZoneFirewallUaRuleListResponseResultConfiguration]
type zoneFirewallUaRuleListResponseResultConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseResultConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseResultConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type ZoneFirewallUaRuleListResponseResultMode string

const (
	ZoneFirewallUaRuleListResponseResultModeBlock            ZoneFirewallUaRuleListResponseResultMode = "block"
	ZoneFirewallUaRuleListResponseResultModeChallenge        ZoneFirewallUaRuleListResponseResultMode = "challenge"
	ZoneFirewallUaRuleListResponseResultModeJsChallenge      ZoneFirewallUaRuleListResponseResultMode = "js_challenge"
	ZoneFirewallUaRuleListResponseResultModeManagedChallenge ZoneFirewallUaRuleListResponseResultMode = "managed_challenge"
)

func (r ZoneFirewallUaRuleListResponseResultMode) IsKnown() bool {
	switch r {
	case ZoneFirewallUaRuleListResponseResultModeBlock, ZoneFirewallUaRuleListResponseResultModeChallenge, ZoneFirewallUaRuleListResponseResultModeJsChallenge, ZoneFirewallUaRuleListResponseResultModeManagedChallenge:
		return true
	}
	return false
}

type ZoneFirewallUaRuleDeleteResponse struct {
	Result ZoneFirewallUaRuleDeleteResponseResult `json:"result"`
	JSON   zoneFirewallUaRuleDeleteResponseJSON   `json:"-"`
	FirewallUablockResponseSingle
}

// zoneFirewallUaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleDeleteResponse]
type zoneFirewallUaRuleDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleDeleteResponseResult struct {
	// The unique identifier of the User Agent Blocking rule.
	ID   string                                     `json:"id"`
	JSON zoneFirewallUaRuleDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleDeleteResponseResult]
type zoneFirewallUaRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
}

func (r ZoneFirewallUaRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallUaRuleUpdateParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
}

func (r ZoneFirewallUaRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallUaRuleListParams struct {
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// A string to search for in the user agent values of existing rules.
	UaSearch param.Field[string] `query:"ua_search"`
}

// URLQuery serializes [ZoneFirewallUaRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallUaRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallUaRuleDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneFirewallUaRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
