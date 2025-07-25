// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/option"
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

type FirewallAPIResponseCollection struct {
	Errors   []FirewallAPIResponseCollectionError   `json:"errors,required"`
	Messages []FirewallAPIResponseCollectionMessage `json:"messages,required"`
	Result   []interface{}                          `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success    FirewallAPIResponseCollectionSuccess    `json:"success,required"`
	ResultInfo FirewallAPIResponseCollectionResultInfo `json:"result_info"`
	JSON       firewallAPIResponseCollectionJSON       `json:"-"`
}

// firewallAPIResponseCollectionJSON contains the JSON metadata for the struct
// [FirewallAPIResponseCollection]
type firewallAPIResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type FirewallAPIResponseCollectionError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           FirewallAPIResponseCollectionErrorsSource `json:"source"`
	JSON             firewallAPIResponseCollectionErrorJSON    `json:"-"`
}

// firewallAPIResponseCollectionErrorJSON contains the JSON metadata for the struct
// [FirewallAPIResponseCollectionError]
type firewallAPIResponseCollectionErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallAPIResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseCollectionErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    firewallAPIResponseCollectionErrorsSourceJSON `json:"-"`
}

// firewallAPIResponseCollectionErrorsSourceJSON contains the JSON metadata for the
// struct [FirewallAPIResponseCollectionErrorsSource]
type firewallAPIResponseCollectionErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseCollectionErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseCollectionMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           FirewallAPIResponseCollectionMessagesSource `json:"source"`
	JSON             firewallAPIResponseCollectionMessageJSON    `json:"-"`
}

// firewallAPIResponseCollectionMessageJSON contains the JSON metadata for the
// struct [FirewallAPIResponseCollectionMessage]
type firewallAPIResponseCollectionMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallAPIResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseCollectionMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    firewallAPIResponseCollectionMessagesSourceJSON `json:"-"`
}

// firewallAPIResponseCollectionMessagesSourceJSON contains the JSON metadata for
// the struct [FirewallAPIResponseCollectionMessagesSource]
type firewallAPIResponseCollectionMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseCollectionMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseCollectionMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallAPIResponseCollectionSuccess bool

const (
	FirewallAPIResponseCollectionSuccessTrue FirewallAPIResponseCollectionSuccess = true
)

func (r FirewallAPIResponseCollectionSuccess) IsKnown() bool {
	switch r {
	case FirewallAPIResponseCollectionSuccessTrue:
		return true
	}
	return false
}

type FirewallAPIResponseCollectionResultInfo struct {
	// Defines the total number of results for the requested service.
	Count float64 `json:"count"`
	// Defines the current page within paginated list of results.
	Page float64 `json:"page"`
	// Defines the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Defines the total results available without any search parameters.
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

type FirewallAPIResponseSingle struct {
	Errors   []FirewallAPIResponseSingleError   `json:"errors,required"`
	Messages []FirewallAPIResponseSingleMessage `json:"messages,required"`
	Result   interface{}                        `json:"result,required"`
	// Defines whether the API call was successful.
	Success FirewallAPIResponseSingleSuccess `json:"success,required"`
	JSON    firewallAPIResponseSingleJSON    `json:"-"`
}

// firewallAPIResponseSingleJSON contains the JSON metadata for the struct
// [FirewallAPIResponseSingle]
type firewallAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
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

type FirewallAPIResponseSingleError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           FirewallAPIResponseSingleErrorsSource `json:"source"`
	JSON             firewallAPIResponseSingleErrorJSON    `json:"-"`
}

// firewallAPIResponseSingleErrorJSON contains the JSON metadata for the struct
// [FirewallAPIResponseSingleError]
type firewallAPIResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallAPIResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseSingleErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    firewallAPIResponseSingleErrorsSourceJSON `json:"-"`
}

// firewallAPIResponseSingleErrorsSourceJSON contains the JSON metadata for the
// struct [FirewallAPIResponseSingleErrorsSource]
type firewallAPIResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseSingleMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           FirewallAPIResponseSingleMessagesSource `json:"source"`
	JSON             firewallAPIResponseSingleMessageJSON    `json:"-"`
}

// firewallAPIResponseSingleMessageJSON contains the JSON metadata for the struct
// [FirewallAPIResponseSingleMessage]
type firewallAPIResponseSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallAPIResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

type FirewallAPIResponseSingleMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    firewallAPIResponseSingleMessagesSourceJSON `json:"-"`
}

// firewallAPIResponseSingleMessagesSourceJSON contains the JSON metadata for the
// struct [FirewallAPIResponseSingleMessagesSource]
type firewallAPIResponseSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAPIResponseSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAPIResponseSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type FirewallAPIResponseSingleSuccess bool

const (
	FirewallAPIResponseSingleSuccessTrue FirewallAPIResponseSingleSuccess = true
)

func (r FirewallAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case FirewallAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type FirewallMessagesItem struct {
	Code             int64                      `json:"code,required"`
	Message          string                     `json:"message,required"`
	DocumentationURL string                     `json:"documentation_url"`
	Source           FirewallMessagesItemSource `json:"source"`
	JSON             firewallMessagesItemJSON   `json:"-"`
}

// firewallMessagesItemJSON contains the JSON metadata for the struct
// [FirewallMessagesItem]
type firewallMessagesItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FirewallMessagesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMessagesItemJSON) RawJSON() string {
	return r.raw
}

type FirewallMessagesItemSource struct {
	Pointer string                         `json:"pointer"`
	JSON    firewallMessagesItemSourceJSON `json:"-"`
}

// firewallMessagesItemSourceJSON contains the JSON metadata for the struct
// [FirewallMessagesItemSource]
type firewallMessagesItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallMessagesItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallMessagesItemSourceJSON) RawJSON() string {
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
