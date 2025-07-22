// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountMnmRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMnmRuleService] method instead.
type AccountMnmRuleService struct {
	Options []option.RequestOption
}

// NewAccountMnmRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMnmRuleService(opts ...option.RequestOption) (r *AccountMnmRuleService) {
	r = &AccountMnmRuleService{}
	r.Options = opts
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *AccountMnmRuleService) New(ctx context.Context, accountID string, body AccountMnmRuleNewParams, opts ...option.RequestOption) (res *RulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List a single network monitoring rule for account.
func (r *AccountMnmRuleService) Get(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *RulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update network monitoring rules for account.
func (r *AccountMnmRuleService) Update(ctx context.Context, accountID string, body AccountMnmRuleUpdateParams, opts ...option.RequestOption) (res *RulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists network monitoring rules for account.
func (r *AccountMnmRuleService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMnmRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a network monitoring rule for account.
func (r *AccountMnmRuleService) Delete(ctx context.Context, accountID string, ruleID string, body AccountMnmRuleDeleteParams, opts ...option.RequestOption) (res *RulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update advertisement for rule.
func (r *AccountMnmRuleService) UpdateAdvertisement(ctx context.Context, accountID string, ruleID string, body AccountMnmRuleUpdateAdvertisementParams, opts ...option.RequestOption) (res *AccountMnmRuleUpdateAdvertisementResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s/advertisement", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Update a network monitoring rule for account.
func (r *AccountMnmRuleService) UpdateRule(ctx context.Context, accountID string, ruleID string, body AccountMnmRuleUpdateRuleParams, opts ...option.RequestOption) (res *RulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type MnmRule struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string   `json:"name,required"`
	Prefixes []string `json:"prefixes,required"`
	// MNM rule type.
	Type MnmRuleType `json:"type,required"`
	// The id of the rule. Must be unique.
	ID string `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration RuleDuration `json:"duration"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64 `json:"packet_threshold"`
	// Prefix match type to be applied for a prefix auto advertisement when using an
	// advanced_ddos rule.
	PrefixMatch MnmRulePrefixMatch `json:"prefix_match,nullable"`
	// Level of sensitivity set for zscore rules.
	ZscoreSensitivity MnmRuleZscoreSensitivity `json:"zscore_sensitivity,nullable"`
	// Target of the zscore rule analysis.
	ZscoreTarget MnmRuleZscoreTarget `json:"zscore_target,nullable"`
	JSON         mnmRuleJSON         `json:"-"`
}

// mnmRuleJSON contains the JSON metadata for the struct [MnmRule]
type mnmRuleJSON struct {
	AutomaticAdvertisement apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	Type                   apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	Duration               apijson.Field
	PacketThreshold        apijson.Field
	PrefixMatch            apijson.Field
	ZscoreSensitivity      apijson.Field
	ZscoreTarget           apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *MnmRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mnmRuleJSON) RawJSON() string {
	return r.raw
}

// MNM rule type.
type MnmRuleType string

const (
	MnmRuleTypeThreshold    MnmRuleType = "threshold"
	MnmRuleTypeZscore       MnmRuleType = "zscore"
	MnmRuleTypeAdvancedDdos MnmRuleType = "advanced_ddos"
)

func (r MnmRuleType) IsKnown() bool {
	switch r {
	case MnmRuleTypeThreshold, MnmRuleTypeZscore, MnmRuleTypeAdvancedDdos:
		return true
	}
	return false
}

// Prefix match type to be applied for a prefix auto advertisement when using an
// advanced_ddos rule.
type MnmRulePrefixMatch string

const (
	MnmRulePrefixMatchExact    MnmRulePrefixMatch = "exact"
	MnmRulePrefixMatchSubnet   MnmRulePrefixMatch = "subnet"
	MnmRulePrefixMatchSupernet MnmRulePrefixMatch = "supernet"
)

func (r MnmRulePrefixMatch) IsKnown() bool {
	switch r {
	case MnmRulePrefixMatchExact, MnmRulePrefixMatchSubnet, MnmRulePrefixMatchSupernet:
		return true
	}
	return false
}

// Level of sensitivity set for zscore rules.
type MnmRuleZscoreSensitivity string

const (
	MnmRuleZscoreSensitivityLow    MnmRuleZscoreSensitivity = "low"
	MnmRuleZscoreSensitivityMedium MnmRuleZscoreSensitivity = "medium"
	MnmRuleZscoreSensitivityHigh   MnmRuleZscoreSensitivity = "high"
)

func (r MnmRuleZscoreSensitivity) IsKnown() bool {
	switch r {
	case MnmRuleZscoreSensitivityLow, MnmRuleZscoreSensitivityMedium, MnmRuleZscoreSensitivityHigh:
		return true
	}
	return false
}

// Target of the zscore rule analysis.
type MnmRuleZscoreTarget string

const (
	MnmRuleZscoreTargetBits    MnmRuleZscoreTarget = "bits"
	MnmRuleZscoreTargetPackets MnmRuleZscoreTarget = "packets"
)

func (r MnmRuleZscoreTarget) IsKnown() bool {
	switch r {
	case MnmRuleZscoreTargetBits, MnmRuleZscoreTargetPackets:
		return true
	}
	return false
}

// The amount of time that the rule threshold must be exceeded to send an alert
// notification. The final value must be equivalent to one of the following 8
// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
type RuleDuration string

const (
	RuleDuration1m  RuleDuration = "1m"
	RuleDuration5m  RuleDuration = "5m"
	RuleDuration10m RuleDuration = "10m"
	RuleDuration15m RuleDuration = "15m"
	RuleDuration20m RuleDuration = "20m"
	RuleDuration30m RuleDuration = "30m"
	RuleDuration45m RuleDuration = "45m"
	RuleDuration60m RuleDuration = "60m"
)

func (r RuleDuration) IsKnown() bool {
	switch r {
	case RuleDuration1m, RuleDuration5m, RuleDuration10m, RuleDuration15m, RuleDuration20m, RuleDuration30m, RuleDuration45m, RuleDuration60m:
		return true
	}
	return false
}

type RulesSingleResponse struct {
	Result MnmRule                 `json:"result,nullable"`
	JSON   rulesSingleResponseJSON `json:"-"`
	APIResponseSingleMagicVisibility
}

// rulesSingleResponseJSON contains the JSON metadata for the struct
// [RulesSingleResponse]
type rulesSingleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesSingleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMnmRuleListResponse struct {
	Result     []MnmRule                            `json:"result,nullable"`
	ResultInfo AccountMnmRuleListResponseResultInfo `json:"result_info"`
	JSON       accountMnmRuleListResponseJSON       `json:"-"`
	APIResponseMagicVisibilityMnm
}

// accountMnmRuleListResponseJSON contains the JSON metadata for the struct
// [AccountMnmRuleListResponse]
type accountMnmRuleListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMnmRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMnmRuleListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       accountMnmRuleListResponseResultInfoJSON `json:"-"`
}

// accountMnmRuleListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountMnmRuleListResponseResultInfo]
type accountMnmRuleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMnmRuleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountMnmRuleUpdateAdvertisementResponse struct {
	Result AccountMnmRuleUpdateAdvertisementResponseResult `json:"result,nullable"`
	JSON   accountMnmRuleUpdateAdvertisementResponseJSON   `json:"-"`
	APIResponseSingleMagicVisibility
}

// accountMnmRuleUpdateAdvertisementResponseJSON contains the JSON metadata for the
// struct [AccountMnmRuleUpdateAdvertisementResponse]
type accountMnmRuleUpdateAdvertisementResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateAdvertisementResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMnmRuleUpdateAdvertisementResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMnmRuleUpdateAdvertisementResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool                                                `json:"automatic_advertisement,required,nullable"`
	JSON                   accountMnmRuleUpdateAdvertisementResponseResultJSON `json:"-"`
}

// accountMnmRuleUpdateAdvertisementResponseResultJSON contains the JSON metadata
// for the struct [AccountMnmRuleUpdateAdvertisementResponseResult]
type accountMnmRuleUpdateAdvertisementResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateAdvertisementResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMnmRuleUpdateAdvertisementResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMnmRuleNewParams struct {
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleDuration] `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name,required"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r AccountMnmRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMnmRuleUpdateParams struct {
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleDuration] `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name,required"`
	// The id of the rule. Must be unique.
	ID param.Field[string] `json:"id"`
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r AccountMnmRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountMnmRuleDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMnmRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMnmRuleUpdateAdvertisementParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMnmRuleUpdateAdvertisementParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMnmRuleUpdateRuleParams struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement param.Field[bool] `json:"automatic_advertisement"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	Bandwidth param.Field[float64] `json:"bandwidth"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"].
	Duration param.Field[RuleDuration] `json:"duration"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name param.Field[string] `json:"name"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold param.Field[float64]  `json:"packet_threshold"`
	Prefixes        param.Field[[]string] `json:"prefixes"`
}

func (r AccountMnmRuleUpdateRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
