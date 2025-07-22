// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
)

// AccountFirewallAccessRuleRuleService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountFirewallAccessRuleRuleService] method instead.
type AccountFirewallAccessRuleRuleService struct {
	Options []option.RequestOption
}

// NewAccountFirewallAccessRuleRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountFirewallAccessRuleRuleService(opts ...option.RequestOption) (r *AccountFirewallAccessRuleRuleService) {
	r = &AccountFirewallAccessRuleRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account. The rule will apply to all zones in
// the account.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *AccountFirewallAccessRuleRuleService) New(ctx context.Context, accountID string, body AccountFirewallAccessRuleRuleNewParams, opts ...option.RequestOption) (res *FirewallResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/firewall/access_rules/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches the details of an IP Access rule defined at the account level.
func (r *AccountFirewallAccessRuleRuleService) Get(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *FirewallResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/firewall/access_rules/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Update(ctx context.Context, accountID string, ruleID string, body AccountFirewallAccessRuleRuleUpdateParams, opts ...option.RequestOption) (res *FirewallResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/firewall/access_rules/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches IP Access rules of an account. These rules apply to all the zones in the
// account. You can filter the results using several optional parameters.
func (r *AccountFirewallAccessRuleRuleService) List(ctx context.Context, accountID string, query AccountFirewallAccessRuleRuleListParams, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/firewall/access_rules/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing IP Access rule defined at the account level.
//
// Note: This operation will affect all zones in the account.
func (r *AccountFirewallAccessRuleRuleService) Delete(ctx context.Context, accountID string, ruleID string, body AccountFirewallAccessRuleRuleDeleteParams, opts ...option.RequestOption) (res *AccountFirewallAccessRuleRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/firewall/access_rules/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type FirewallAPIResponseCollection struct {
	Result     []interface{}                           `json:"result,nullable"`
	ResultInfo FirewallAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       firewallAPIResponseCollectionJSON       `json:"-"`
	FirewallAPIResponseCommon
}

// firewallAPIResponseCollectionJSON contains the JSON metadata for the struct
// [FirewallAPIResponseCollection]
type firewallAPIResponseCollectionJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAPIResponseCollection) implementsZoneFirewallWafPackageListResponse() {}

type FirewallAPIResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       firewallAPIResponseCollectionResultInfoJSON `json:"-"`
}

// firewallAPIResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [FirewallAPIResponseCollectionResultInfo]
type firewallAPIResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionResultInfoJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseCommon struct {
	Errors   []FirewallMessagesItem               `json:"errors,required"`
	Messages []FirewallMessagesItem               `json:"messages,required"`
	Result   FirewallAPIResponseCommonResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAPIResponseCommonSuccess `json:"success,required"`
	JSON    firewallAPIResponseCommonJSON    `json:"-"`
}

// firewallAPIResponseCommonJSON contains the JSON metadata for the struct
// [FirewallAPIResponseCommon]
type firewallAPIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCommonJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [FirewallAPIResponseCommonResultArray] or
// [shared.UnionString].
type FirewallAPIResponseCommonResultUnion interface {
	ImplementsFirewallAPIResponseCommonResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallAPIResponseCommonResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallAPIResponseCommonResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallAPIResponseCommonResultArray []interface{}

func (r FirewallAPIResponseCommonResultArray) ImplementsFirewallAPIResponseCommonResultUnion() {}

// Whether the API call was successful
type FirewallAPIResponseCommonSuccess bool

const (
	FirewallAPIResponseCommonSuccessTrue FirewallAPIResponseCommonSuccess = true
)

func (r FirewallAPIResponseCommonSuccess) IsKnown() bool {
	switch r {
	case FirewallAPIResponseCommonSuccessTrue:
		return true
	}
	return false
}

type FirewallAPIResponseSingle struct {
	Result interface{}                   `json:"result"`
	JSON   firewallAPIResponseSingleJSON `json:"-"`
	FirewallAPIResponseCommon
}

// firewallAPIResponseSingleJSON contains the JSON metadata for the struct
// [FirewallAPIResponseSingle]
type firewallAPIResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r FirewallAPIResponseSingle) implementsFirewallPackageResponseSingle() {}

type FirewallMessagesItem struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    firewallMessagesItemJSON `json:"-"`
}

// firewallMessagesItemJSON contains the JSON metadata for the struct
// [FirewallMessagesItem]
type firewallMessagesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMessagesItemJSON) RawJSON() string {
	return r.raw
}

type FirewallResponseSingle struct {
	Result FirewallSchemasRule        `json:"result"`
	JSON   firewallResponseSingleJSON `json:"-"`
	FirewallAPIResponseSingle
}

// firewallResponseSingleJSON contains the JSON metadata for the struct
// [FirewallResponseSingle]
type firewallResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallResponseSingleJSON) RawJSON() string {
	return r.raw
}

type FirewallRule struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []FirewallSchemasMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration FirewallRuleConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode FirewallSchemasMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string           `json:"notes"`
	JSON  firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleParam struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r FirewallRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type FirewallRuleConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallRuleConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                        `json:"value"`
	JSON  firewallRuleConfigurationJSON `json:"-"`
	union FirewallRuleConfigurationUnion
}

// firewallRuleConfigurationJSON contains the JSON metadata for the struct
// [FirewallRuleConfiguration]
type firewallRuleConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallRuleConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallRuleConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallRuleConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallRuleConfigurationUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [FirewallRuleConfigurationFirewallIPConfiguration],
// [FirewallRuleConfigurationFirewallIpv6Configuration],
// [FirewallRuleConfigurationFirewallCidrConfiguration],
// [FirewallRuleConfigurationFirewallAsnConfiguration],
// [FirewallRuleConfigurationFirewallCountryConfiguration].
func (r FirewallRuleConfiguration) AsUnion() FirewallRuleConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by [FirewallRuleConfigurationFirewallIPConfiguration],
// [FirewallRuleConfigurationFirewallIpv6Configuration],
// [FirewallRuleConfigurationFirewallCidrConfiguration],
// [FirewallRuleConfigurationFirewallAsnConfiguration] or
// [FirewallRuleConfigurationFirewallCountryConfiguration].
type FirewallRuleConfigurationUnion interface {
	implementsFirewallRuleConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationFirewallIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationFirewallIpv6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationFirewallCidrConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationFirewallAsnConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallRuleConfigurationFirewallCountryConfiguration{}),
		},
	)
}

type FirewallRuleConfigurationFirewallIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target FirewallRuleConfigurationFirewallIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                               `json:"value"`
	JSON  firewallRuleConfigurationFirewallIPConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationFirewallIPConfigurationJSON contains the JSON metadata
// for the struct [FirewallRuleConfigurationFirewallIPConfiguration]
type firewallRuleConfigurationFirewallIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationFirewallIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationFirewallIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationFirewallIPConfiguration) implementsFirewallRuleConfiguration() {}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallRuleConfigurationFirewallIPConfigurationTarget string

const (
	FirewallRuleConfigurationFirewallIPConfigurationTargetIP FirewallRuleConfigurationFirewallIPConfigurationTarget = "ip"
)

func (r FirewallRuleConfigurationFirewallIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationFirewallIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallRuleConfigurationFirewallIpv6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target FirewallRuleConfigurationFirewallIpv6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                                                 `json:"value"`
	JSON  firewallRuleConfigurationFirewallIpv6ConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationFirewallIpv6ConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationFirewallIpv6Configuration]
type firewallRuleConfigurationFirewallIpv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationFirewallIpv6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationFirewallIpv6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationFirewallIpv6Configuration) implementsFirewallRuleConfiguration() {}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type FirewallRuleConfigurationFirewallIpv6ConfigurationTarget string

const (
	FirewallRuleConfigurationFirewallIpv6ConfigurationTargetIp6 FirewallRuleConfigurationFirewallIpv6ConfigurationTarget = "ip6"
)

func (r FirewallRuleConfigurationFirewallIpv6ConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationFirewallIpv6ConfigurationTargetIp6:
		return true
	}
	return false
}

type FirewallRuleConfigurationFirewallCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target FirewallRuleConfigurationFirewallCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                                                 `json:"value"`
	JSON  firewallRuleConfigurationFirewallCidrConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationFirewallCidrConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationFirewallCidrConfiguration]
type firewallRuleConfigurationFirewallCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationFirewallCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationFirewallCidrConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationFirewallCidrConfiguration) implementsFirewallRuleConfiguration() {}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type FirewallRuleConfigurationFirewallCidrConfigurationTarget string

const (
	FirewallRuleConfigurationFirewallCidrConfigurationTargetIPRange FirewallRuleConfigurationFirewallCidrConfigurationTarget = "ip_range"
)

func (r FirewallRuleConfigurationFirewallCidrConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationFirewallCidrConfigurationTargetIPRange:
		return true
	}
	return false
}

type FirewallRuleConfigurationFirewallAsnConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target FirewallRuleConfigurationFirewallAsnConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string                                                `json:"value"`
	JSON  firewallRuleConfigurationFirewallAsnConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationFirewallAsnConfigurationJSON contains the JSON metadata
// for the struct [FirewallRuleConfigurationFirewallAsnConfiguration]
type firewallRuleConfigurationFirewallAsnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationFirewallAsnConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationFirewallAsnConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationFirewallAsnConfiguration) implementsFirewallRuleConfiguration() {}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type FirewallRuleConfigurationFirewallAsnConfigurationTarget string

const (
	FirewallRuleConfigurationFirewallAsnConfigurationTargetAsn FirewallRuleConfigurationFirewallAsnConfigurationTarget = "asn"
)

func (r FirewallRuleConfigurationFirewallAsnConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationFirewallAsnConfigurationTargetAsn:
		return true
	}
	return false
}

type FirewallRuleConfigurationFirewallCountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target FirewallRuleConfigurationFirewallCountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                                                    `json:"value"`
	JSON  firewallRuleConfigurationFirewallCountryConfigurationJSON `json:"-"`
}

// firewallRuleConfigurationFirewallCountryConfigurationJSON contains the JSON
// metadata for the struct [FirewallRuleConfigurationFirewallCountryConfiguration]
type firewallRuleConfigurationFirewallCountryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRuleConfigurationFirewallCountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleConfigurationFirewallCountryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallRuleConfigurationFirewallCountryConfiguration) implementsFirewallRuleConfiguration() {
}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type FirewallRuleConfigurationFirewallCountryConfigurationTarget string

const (
	FirewallRuleConfigurationFirewallCountryConfigurationTargetCountry FirewallRuleConfigurationFirewallCountryConfigurationTarget = "country"
)

func (r FirewallRuleConfigurationFirewallCountryConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationFirewallCountryConfigurationTargetCountry:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type FirewallRuleConfigurationTarget string

const (
	FirewallRuleConfigurationTargetIP      FirewallRuleConfigurationTarget = "ip"
	FirewallRuleConfigurationTargetIp6     FirewallRuleConfigurationTarget = "ip6"
	FirewallRuleConfigurationTargetIPRange FirewallRuleConfigurationTarget = "ip_range"
	FirewallRuleConfigurationTargetAsn     FirewallRuleConfigurationTarget = "asn"
	FirewallRuleConfigurationTargetCountry FirewallRuleConfigurationTarget = "country"
)

func (r FirewallRuleConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallRuleConfigurationTargetIP, FirewallRuleConfigurationTargetIp6, FirewallRuleConfigurationTargetIPRange, FirewallRuleConfigurationTargetAsn, FirewallRuleConfigurationTargetCountry:
		return true
	}
	return false
}

// The rule configuration.
type FirewallRuleConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallRuleConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {}

// The rule configuration.
//
// Satisfied by [FirewallRuleConfigurationFirewallIPConfigurationParam],
// [FirewallRuleConfigurationFirewallIpv6ConfigurationParam],
// [FirewallRuleConfigurationFirewallCidrConfigurationParam],
// [FirewallRuleConfigurationFirewallAsnConfigurationParam],
// [FirewallRuleConfigurationFirewallCountryConfigurationParam],
// [FirewallRuleConfigurationParam].
type FirewallRuleConfigurationUnionParam interface {
	implementsFirewallRuleConfigurationUnionParam()
}

type FirewallRuleConfigurationFirewallIPConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[FirewallRuleConfigurationFirewallIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationFirewallIPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationFirewallIPConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {
}

type FirewallRuleConfigurationFirewallIpv6ConfigurationParam struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[FirewallRuleConfigurationFirewallIpv6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationFirewallIpv6ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationFirewallIpv6ConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {
}

type FirewallRuleConfigurationFirewallCidrConfigurationParam struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[FirewallRuleConfigurationFirewallCidrConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationFirewallCidrConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationFirewallCidrConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {
}

type FirewallRuleConfigurationFirewallAsnConfigurationParam struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[FirewallRuleConfigurationFirewallAsnConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationFirewallAsnConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationFirewallAsnConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {
}

type FirewallRuleConfigurationFirewallCountryConfigurationParam struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[FirewallRuleConfigurationFirewallCountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r FirewallRuleConfigurationFirewallCountryConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r FirewallRuleConfigurationFirewallCountryConfigurationParam) implementsFirewallRuleConfigurationUnionParam() {
}

// The action to apply to a matched request.
type FirewallSchemasMode string

const (
	FirewallSchemasModeBlock            FirewallSchemasMode = "block"
	FirewallSchemasModeChallenge        FirewallSchemasMode = "challenge"
	FirewallSchemasModeWhitelist        FirewallSchemasMode = "whitelist"
	FirewallSchemasModeJsChallenge      FirewallSchemasMode = "js_challenge"
	FirewallSchemasModeManagedChallenge FirewallSchemasMode = "managed_challenge"
)

func (r FirewallSchemasMode) IsKnown() bool {
	switch r {
	case FirewallSchemasModeBlock, FirewallSchemasModeChallenge, FirewallSchemasModeWhitelist, FirewallSchemasModeJsChallenge, FirewallSchemasModeManagedChallenge:
		return true
	}
	return false
}

type FirewallSchemasRule struct {
	// All zones owned by the user will have the rule applied.
	Scope FirewallSchemasRuleScope `json:"scope"`
	JSON  firewallSchemasRuleJSON  `json:"-"`
	FirewallRule
}

// firewallSchemasRuleJSON contains the JSON metadata for the struct
// [FirewallSchemasRule]
type firewallSchemasRuleJSON struct {
	Scope       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallSchemasRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallSchemasRuleJSON) RawJSON() string {
	return r.raw
}

// All zones owned by the user will have the rule applied.
type FirewallSchemasRuleScope struct {
	// Identifier
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The scope of the rule.
	Type FirewallSchemasRuleScopeType `json:"type"`
	JSON firewallSchemasRuleScopeJSON `json:"-"`
}

// firewallSchemasRuleScopeJSON contains the JSON metadata for the struct
// [FirewallSchemasRuleScope]
type firewallSchemasRuleScopeJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallSchemasRuleScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallSchemasRuleScopeJSON) RawJSON() string {
	return r.raw
}

// The scope of the rule.
type FirewallSchemasRuleScopeType string

const (
	FirewallSchemasRuleScopeTypeUser         FirewallSchemasRuleScopeType = "user"
	FirewallSchemasRuleScopeTypeOrganization FirewallSchemasRuleScopeType = "organization"
)

func (r FirewallSchemasRuleScopeType) IsKnown() bool {
	switch r {
	case FirewallSchemasRuleScopeTypeUser, FirewallSchemasRuleScopeTypeOrganization:
		return true
	}
	return false
}

type FirewallSchemasRuleParam struct {
	FirewallRuleParam
}

func (r FirewallSchemasRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// All zones owned by the user will have the rule applied.
type FirewallSchemasRuleScopeParam struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email"`
}

func (r FirewallSchemasRuleScopeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountFirewallAccessRuleRuleListResponse struct {
	Result []FirewallSchemasRule                         `json:"result"`
	JSON   accountFirewallAccessRuleRuleListResponseJSON `json:"-"`
	FirewallAPIResponseCollection
}

// accountFirewallAccessRuleRuleListResponseJSON contains the JSON metadata for the
// struct [AccountFirewallAccessRuleRuleListResponse]
type accountFirewallAccessRuleRuleListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFirewallAccessRuleRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountFirewallAccessRuleRuleDeleteResponse struct {
	Result AccountFirewallAccessRuleRuleDeleteResponseResult `json:"result,nullable"`
	JSON   accountFirewallAccessRuleRuleDeleteResponseJSON   `json:"-"`
	FirewallAPIResponseCommon
}

// accountFirewallAccessRuleRuleDeleteResponseJSON contains the JSON metadata for
// the struct [AccountFirewallAccessRuleRuleDeleteResponse]
type accountFirewallAccessRuleRuleDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFirewallAccessRuleRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountFirewallAccessRuleRuleDeleteResponseResult struct {
	// Identifier
	ID   string                                                `json:"id,required"`
	JSON accountFirewallAccessRuleRuleDeleteResponseResultJSON `json:"-"`
}

// accountFirewallAccessRuleRuleDeleteResponseResultJSON contains the JSON metadata
// for the struct [AccountFirewallAccessRuleRuleDeleteResponseResult]
type accountFirewallAccessRuleRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountFirewallAccessRuleRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountFirewallAccessRuleRuleDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountFirewallAccessRuleRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[FirewallRuleConfigurationUnionParam] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `json:"mode,required"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccountFirewallAccessRuleRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountFirewallAccessRuleRuleUpdateParams struct {
	FirewallSchemasRule FirewallSchemasRuleParam `json:"firewall_schemas_rule,required"`
}

func (r AccountFirewallAccessRuleRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.FirewallSchemasRule)
}

type AccountFirewallAccessRuleRuleListParams struct {
	Configuration param.Field[AccountFirewallAccessRuleRuleListParamsConfiguration] `query:"configuration"`
	// The direction used to sort returned rules.
	Direction param.Field[AccountFirewallAccessRuleRuleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[AccountFirewallAccessRuleRuleListParamsMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[FirewallSchemasMode] `query:"mode"`
	// The string to search for in the notes of existing IP Access rules. Notes: For
	// example, the string 'attack' would match IP Access rules with notes 'Attack
	// 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
	// The field used to sort returned rules.
	Order param.Field[AccountFirewallAccessRuleRuleListParamsOrder] `query:"order"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountFirewallAccessRuleRuleListParams]'s query parameters
// as `url.Values`.
func (r AccountFirewallAccessRuleRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountFirewallAccessRuleRuleListParamsConfiguration struct {
	// The target to search in existing rules.
	Target param.Field[AccountFirewallAccessRuleRuleListParamsConfigurationTarget] `query:"target"`
	// The target value to search for in existing rules: an IP address, an IP address
	// range, or a country code, depending on the provided `configuration.target`.
	// Notes: You can search for a single IPv4 address, an IP address range with a
	// subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [AccountFirewallAccessRuleRuleListParamsConfiguration]'s
// query parameters as `url.Values`.
func (r AccountFirewallAccessRuleRuleListParamsConfiguration) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The target to search in existing rules.
type AccountFirewallAccessRuleRuleListParamsConfigurationTarget string

const (
	AccountFirewallAccessRuleRuleListParamsConfigurationTargetIP      AccountFirewallAccessRuleRuleListParamsConfigurationTarget = "ip"
	AccountFirewallAccessRuleRuleListParamsConfigurationTargetIPRange AccountFirewallAccessRuleRuleListParamsConfigurationTarget = "ip_range"
	AccountFirewallAccessRuleRuleListParamsConfigurationTargetAsn     AccountFirewallAccessRuleRuleListParamsConfigurationTarget = "asn"
	AccountFirewallAccessRuleRuleListParamsConfigurationTargetCountry AccountFirewallAccessRuleRuleListParamsConfigurationTarget = "country"
)

func (r AccountFirewallAccessRuleRuleListParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccountFirewallAccessRuleRuleListParamsConfigurationTargetIP, AccountFirewallAccessRuleRuleListParamsConfigurationTargetIPRange, AccountFirewallAccessRuleRuleListParamsConfigurationTargetAsn, AccountFirewallAccessRuleRuleListParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The direction used to sort returned rules.
type AccountFirewallAccessRuleRuleListParamsDirection string

const (
	AccountFirewallAccessRuleRuleListParamsDirectionAsc  AccountFirewallAccessRuleRuleListParamsDirection = "asc"
	AccountFirewallAccessRuleRuleListParamsDirectionDesc AccountFirewallAccessRuleRuleListParamsDirection = "desc"
)

func (r AccountFirewallAccessRuleRuleListParamsDirection) IsKnown() bool {
	switch r {
	case AccountFirewallAccessRuleRuleListParamsDirectionAsc, AccountFirewallAccessRuleRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type AccountFirewallAccessRuleRuleListParamsMatch string

const (
	AccountFirewallAccessRuleRuleListParamsMatchAny AccountFirewallAccessRuleRuleListParamsMatch = "any"
	AccountFirewallAccessRuleRuleListParamsMatchAll AccountFirewallAccessRuleRuleListParamsMatch = "all"
)

func (r AccountFirewallAccessRuleRuleListParamsMatch) IsKnown() bool {
	switch r {
	case AccountFirewallAccessRuleRuleListParamsMatchAny, AccountFirewallAccessRuleRuleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned rules.
type AccountFirewallAccessRuleRuleListParamsOrder string

const (
	AccountFirewallAccessRuleRuleListParamsOrderConfigurationTarget AccountFirewallAccessRuleRuleListParamsOrder = "configuration.target"
	AccountFirewallAccessRuleRuleListParamsOrderConfigurationValue  AccountFirewallAccessRuleRuleListParamsOrder = "configuration.value"
	AccountFirewallAccessRuleRuleListParamsOrderMode                AccountFirewallAccessRuleRuleListParamsOrder = "mode"
)

func (r AccountFirewallAccessRuleRuleListParamsOrder) IsKnown() bool {
	switch r {
	case AccountFirewallAccessRuleRuleListParamsOrderConfigurationTarget, AccountFirewallAccessRuleRuleListParamsOrderConfigurationValue, AccountFirewallAccessRuleRuleListParamsOrderMode:
		return true
	}
	return false
}

type AccountFirewallAccessRuleRuleDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountFirewallAccessRuleRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
