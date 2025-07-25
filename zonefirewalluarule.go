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
func (r *ZoneFirewallUaRuleService) Delete(ctx context.Context, zoneID string, uaRuleID string, opts ...option.RequestOption) (res *ZoneFirewallUaRuleDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FirewallUablockResponseSingle struct {
	Errors   []FirewallUablockResponseSingleError   `json:"errors,required"`
	Messages []FirewallUablockResponseSingleMessage `json:"messages,required"`
	Result   FirewallUablockResponseSingleResult    `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallUablockResponseSingleSuccess `json:"success,required"`
	JSON    firewallUablockResponseSingleJSON    `json:"-"`
}

// firewallUablockResponseSingleJSON contains the JSON metadata for the struct
// [FirewallUablockResponseSingle]
type firewallUablockResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUablockResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleJSON) RawJSON() string {
	return r.raw
}

type FirewallUablockResponseSingleError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           FirewallUablockResponseSingleErrorsSource `json:"source"`
	JSON             firewallUablockResponseSingleErrorJSON    `json:"-"`
}

// firewallUablockResponseSingleErrorJSON contains the JSON metadata for the struct
// [FirewallUablockResponseSingleError]
type firewallUablockResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallUablockResponseSingleErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    firewallUablockResponseSingleErrorsSourceJSON `json:"-"`
}

// firewallUablockResponseSingleErrorsSourceJSON contains the JSON metadata for the
// struct [FirewallUablockResponseSingleErrorsSource]
type firewallUablockResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallUablockResponseSingleMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           FirewallUablockResponseSingleMessagesSource `json:"source"`
	JSON             firewallUablockResponseSingleMessageJSON    `json:"-"`
}

// firewallUablockResponseSingleMessageJSON contains the JSON metadata for the
// struct [FirewallUablockResponseSingleMessage]
type firewallUablockResponseSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallUablockResponseSingleMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    firewallUablockResponseSingleMessagesSourceJSON `json:"-"`
}

// firewallUablockResponseSingleMessagesSourceJSON contains the JSON metadata for
// the struct [FirewallUablockResponseSingleMessagesSource]
type firewallUablockResponseSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallUablockResponseSingleResult struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration FirewallUablockResponseSingleResultConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode FirewallUablockResponseSingleResultMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                    `json:"paused"`
	JSON   firewallUablockResponseSingleResultJSON `json:"-"`
}

// firewallUablockResponseSingleResultJSON contains the JSON metadata for the
// struct [FirewallUablockResponseSingleResult]
type firewallUablockResponseSingleResultJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type FirewallUablockResponseSingleResultConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                               `json:"value"`
	JSON  firewallUablockResponseSingleResultConfigurationJSON `json:"-"`
}

// firewallUablockResponseSingleResultConfigurationJSON contains the JSON metadata
// for the struct [FirewallUablockResponseSingleResultConfiguration]
type firewallUablockResponseSingleResultConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUablockResponseSingleResultConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUablockResponseSingleResultConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallUablockResponseSingleResultMode string

const (
	FirewallUablockResponseSingleResultModeBlock            FirewallUablockResponseSingleResultMode = "block"
	FirewallUablockResponseSingleResultModeChallenge        FirewallUablockResponseSingleResultMode = "challenge"
	FirewallUablockResponseSingleResultModeJsChallenge      FirewallUablockResponseSingleResultMode = "js_challenge"
	FirewallUablockResponseSingleResultModeManagedChallenge FirewallUablockResponseSingleResultMode = "managed_challenge"
)

func (r FirewallUablockResponseSingleResultMode) IsKnown() bool {
	switch r {
	case FirewallUablockResponseSingleResultModeBlock, FirewallUablockResponseSingleResultModeChallenge, FirewallUablockResponseSingleResultModeJsChallenge, FirewallUablockResponseSingleResultModeManagedChallenge:
		return true
	}
	return false
}

// Defines whether the API call was successful.
type FirewallUablockResponseSingleSuccess bool

const (
	FirewallUablockResponseSingleSuccessTrue FirewallUablockResponseSingleSuccess = true
)

func (r FirewallUablockResponseSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallUablockResponseSingleSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallUaRuleListResponse struct {
	Errors   []ZoneFirewallUaRuleListResponseError   `json:"errors,required"`
	Messages []ZoneFirewallUaRuleListResponseMessage `json:"messages,required"`
	Result   []ZoneFirewallUaRuleListResponseResult  `json:"result,required"`
	// Defines whether the API call was successful.
	Success    ZoneFirewallUaRuleListResponseSuccess    `json:"success,required"`
	ResultInfo ZoneFirewallUaRuleListResponseResultInfo `json:"result_info"`
	JSON       zoneFirewallUaRuleListResponseJSON       `json:"-"`
}

// zoneFirewallUaRuleListResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleListResponse]
type zoneFirewallUaRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ZoneFirewallUaRuleListResponseErrorsSource `json:"source"`
	JSON             zoneFirewallUaRuleListResponseErrorJSON    `json:"-"`
}

// zoneFirewallUaRuleListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleListResponseError]
type zoneFirewallUaRuleListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    zoneFirewallUaRuleListResponseErrorsSourceJSON `json:"-"`
}

// zoneFirewallUaRuleListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [ZoneFirewallUaRuleListResponseErrorsSource]
type zoneFirewallUaRuleListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponseMessage struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ZoneFirewallUaRuleListResponseMessagesSource `json:"source"`
	JSON             zoneFirewallUaRuleListResponseMessageJSON    `json:"-"`
}

// zoneFirewallUaRuleListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleListResponseMessage]
type zoneFirewallUaRuleListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleListResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    zoneFirewallUaRuleListResponseMessagesSourceJSON `json:"-"`
}

// zoneFirewallUaRuleListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [ZoneFirewallUaRuleListResponseMessagesSource]
type zoneFirewallUaRuleListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseMessagesSourceJSON) RawJSON() string {
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

// Defines whether the API call was successful.
type ZoneFirewallUaRuleListResponseSuccess bool

const (
	ZoneFirewallUaRuleListResponseSuccessTrue ZoneFirewallUaRuleListResponseSuccess = true
)

func (r ZoneFirewallUaRuleListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallUaRuleListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallUaRuleListResponseResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
	TotalCount float64                                      `json:"total_count"`
	JSON       zoneFirewallUaRuleListResponseResultInfoJSON `json:"-"`
}

// zoneFirewallUaRuleListResponseResultInfoJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleListResponseResultInfo]
type zoneFirewallUaRuleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneFirewallUaRuleDeleteResponse struct {
	Errors   []FirewallMessagesItem                 `json:"errors,required"`
	Messages []FirewallMessagesItem                 `json:"messages,required"`
	Result   ZoneFirewallUaRuleDeleteResponseResult `json:"result,required"`
	// Defines whether the API call was successful.
	Success ZoneFirewallUaRuleDeleteResponseSuccess `json:"success,required"`
	JSON    zoneFirewallUaRuleDeleteResponseJSON    `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleDeleteResponse]
type zoneFirewallUaRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration ZoneFirewallUaRuleDeleteResponseResultConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode ZoneFirewallUaRuleDeleteResponseResultMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                       `json:"paused"`
	JSON   zoneFirewallUaRuleDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleDeleteResponseResult]
type zoneFirewallUaRuleDeleteResponseResultJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type ZoneFirewallUaRuleDeleteResponseResultConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                                  `json:"value"`
	JSON  zoneFirewallUaRuleDeleteResponseResultConfigurationJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseResultConfigurationJSON contains the JSON
// metadata for the struct [ZoneFirewallUaRuleDeleteResponseResultConfiguration]
type zoneFirewallUaRuleDeleteResponseResultConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseResultConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneFirewallUaRuleDeleteResponseResultConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type ZoneFirewallUaRuleDeleteResponseResultMode string

const (
	ZoneFirewallUaRuleDeleteResponseResultModeBlock            ZoneFirewallUaRuleDeleteResponseResultMode = "block"
	ZoneFirewallUaRuleDeleteResponseResultModeChallenge        ZoneFirewallUaRuleDeleteResponseResultMode = "challenge"
	ZoneFirewallUaRuleDeleteResponseResultModeJsChallenge      ZoneFirewallUaRuleDeleteResponseResultMode = "js_challenge"
	ZoneFirewallUaRuleDeleteResponseResultModeManagedChallenge ZoneFirewallUaRuleDeleteResponseResultMode = "managed_challenge"
)

func (r ZoneFirewallUaRuleDeleteResponseResultMode) IsKnown() bool {
	switch r {
	case ZoneFirewallUaRuleDeleteResponseResultModeBlock, ZoneFirewallUaRuleDeleteResponseResultModeChallenge, ZoneFirewallUaRuleDeleteResponseResultModeJsChallenge, ZoneFirewallUaRuleDeleteResponseResultModeManagedChallenge:
		return true
	}
	return false
}

// Defines whether the API call was successful.
type ZoneFirewallUaRuleDeleteResponseSuccess bool

const (
	ZoneFirewallUaRuleDeleteResponseSuccessTrue ZoneFirewallUaRuleDeleteResponseSuccess = true
)

func (r ZoneFirewallUaRuleDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneFirewallUaRuleDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneFirewallUaRuleNewParams struct {
	Configuration param.Field[ZoneFirewallUaRuleNewParamsConfiguration] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
}

func (r ZoneFirewallUaRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallUaRuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ua` when specifying a user
	// agent in the rule.
	Target param.Field[ZoneFirewallUaRuleNewParamsConfigurationTarget] `json:"target"`
	// the user agent to exactly match
	Value param.Field[string] `json:"value"`
}

func (r ZoneFirewallUaRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration target. You must set the target to `ua` when specifying a user
// agent in the rule.
type ZoneFirewallUaRuleNewParamsConfigurationTarget string

const (
	ZoneFirewallUaRuleNewParamsConfigurationTargetUa ZoneFirewallUaRuleNewParamsConfigurationTarget = "ua"
)

func (r ZoneFirewallUaRuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case ZoneFirewallUaRuleNewParamsConfigurationTargetUa:
		return true
	}
	return false
}

type ZoneFirewallUaRuleUpdateParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description param.Field[string] `json:"description"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `json:"paused"`
}

func (r ZoneFirewallUaRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallUaRuleListParams struct {
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// A string to search for in the user agent values of existing rules.
	UserAgent param.Field[string] `query:"user_agent"`
}

// URLQuery serializes [ZoneFirewallUaRuleListParams]'s query parameters as
// `url.Values`.
func (r ZoneFirewallUaRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
