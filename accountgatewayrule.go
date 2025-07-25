// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountGatewayRuleService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayRuleService] method instead.
type AccountGatewayRuleService struct {
	Options []option.RequestOption
}

// NewAccountGatewayRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayRuleService(opts ...option.RequestOption) (r *AccountGatewayRuleService) {
	r = &AccountGatewayRuleService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) New(ctx context.Context, accountID string, body AccountGatewayRuleNewParams, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Get(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Update(ctx context.Context, accountID string, ruleID string, body AccountGatewayRuleUpdateParams, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *AccountGatewayRuleService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a Zero Trust Gateway rule.
func (r *AccountGatewayRuleService) Delete(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *ZeroTrustGatewayEmptyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Resets the expiration of a Zero Trust Gateway Rule if its duration has elapsed
// and it has a default duration.
//
// The Zero Trust Gateway Rule must have values for both `expiration.expires_at`
// and `expiration.duration`.
func (r *AccountGatewayRuleService) ResetExpiration(ctx context.Context, accountID string, ruleID string, opts ...option.RequestOption) (res *SingleResponseLocation, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s/reset_expiration", accountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// The action to perform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type ActionPerform string

const (
	ActionPerformOn           ActionPerform = "on"
	ActionPerformOff          ActionPerform = "off"
	ActionPerformAllow        ActionPerform = "allow"
	ActionPerformBlock        ActionPerform = "block"
	ActionPerformScan         ActionPerform = "scan"
	ActionPerformNoscan       ActionPerform = "noscan"
	ActionPerformSafesearch   ActionPerform = "safesearch"
	ActionPerformYtrestricted ActionPerform = "ytrestricted"
	ActionPerformIsolate      ActionPerform = "isolate"
	ActionPerformNoisolate    ActionPerform = "noisolate"
	ActionPerformOverride     ActionPerform = "override"
	ActionPerformL4Override   ActionPerform = "l4_override"
	ActionPerformEgress       ActionPerform = "egress"
	ActionPerformResolve      ActionPerform = "resolve"
	ActionPerformQuarantine   ActionPerform = "quarantine"
	ActionPerformRedirect     ActionPerform = "redirect"
)

func (r ActionPerform) IsKnown() bool {
	switch r {
	case ActionPerformOn, ActionPerformOff, ActionPerformAllow, ActionPerformBlock, ActionPerformScan, ActionPerformNoscan, ActionPerformSafesearch, ActionPerformYtrestricted, ActionPerformIsolate, ActionPerformNoisolate, ActionPerformOverride, ActionPerformL4Override, ActionPerformEgress, ActionPerformResolve, ActionPerformQuarantine, ActionPerformRedirect:
		return true
	}
	return false
}

// The expiration time stamp and default duration of a DNS policy. Takes precedence
// over the policy's `schedule` configuration, if any.
//
// This does not apply to HTTP or network policies.
type Expiration struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero offsets are accepted
	// but will be converted to the equivalent value with offset zero (UTC+00:00) and
	// will be returned as time stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of clients they are
	// applied to, and expire "globally" at the point given by their `expires_at`
	// value.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The default duration a policy will be active in minutes. Must be set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration int64 `json:"duration"`
	// Whether the policy has expired.
	Expired bool           `json:"expired"`
	JSON    expirationJSON `json:"-"`
}

// expirationJSON contains the JSON metadata for the struct [Expiration]
type expirationJSON struct {
	ExpiresAt   apijson.Field
	Duration    apijson.Field
	Expired     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Expiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r expirationJSON) RawJSON() string {
	return r.raw
}

// The expiration time stamp and default duration of a DNS policy. Takes precedence
// over the policy's `schedule` configuration, if any.
//
// This does not apply to HTTP or network policies.
type ExpirationParam struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero offsets are accepted
	// but will be converted to the equivalent value with offset zero (UTC+00:00) and
	// will be returned as time stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of clients they are
	// applied to, and expire "globally" at the point given by their `expires_at`
	// value.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// The default duration a policy will be active in minutes. Must be set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration param.Field[int64] `json:"duration"`
}

func (r ExpirationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol or layer to use.
type GatewayFilters string

const (
	GatewayFiltersHTTP        GatewayFilters = "http"
	GatewayFiltersDNS         GatewayFilters = "dns"
	GatewayFiltersL4          GatewayFilters = "l4"
	GatewayFiltersEgress      GatewayFilters = "egress"
	GatewayFiltersDNSResolver GatewayFilters = "dns_resolver"
)

func (r GatewayFilters) IsKnown() bool {
	switch r {
	case GatewayFiltersHTTP, GatewayFiltersDNS, GatewayFiltersL4, GatewayFiltersEgress, GatewayFiltersDNSResolver:
		return true
	}
	return false
}

type GatewayRule struct {
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action ActionPerform `json:"action,required"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled,required"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []GatewayFilters `json:"filters,required"`
	// The name of the rule.
	Name string `json:"name,required"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence int64 `json:"precedence,required"`
	// The wirefilter expression used for traffic matching.
	Traffic string `json:"traffic,required"`
	// The API resource UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the rule.
	Description string `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture string `json:"device_posture"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration Expiration `json:"expiration,nullable"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The rule cannot be shared via the Orgs API
	NotSharable bool `json:"not_sharable"`
	// The rule was shared via the Orgs API and cannot be edited by the current account
	ReadOnly bool `json:"read_only"`
	// Additional settings that modify the rule's action.
	RuleSettings RuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule Schedule `json:"schedule,nullable"`
	// account tag of account that created the rule
	SourceAccount string    `json:"source_account"`
	UpdatedAt     time.Time `json:"updated_at" format:"date-time"`
	// version number of the rule
	Version int64 `json:"version"`
	// Warning for a misconfigured rule, if any.
	WarningStatus string          `json:"warning_status,nullable"`
	JSON          gatewayRuleJSON `json:"-"`
}

// gatewayRuleJSON contains the JSON metadata for the struct [GatewayRule]
type gatewayRuleJSON struct {
	Action        apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Name          apijson.Field
	Precedence    apijson.Field
	Traffic       apijson.Field
	ID            apijson.Field
	CreatedAt     apijson.Field
	DeletedAt     apijson.Field
	Description   apijson.Field
	DevicePosture apijson.Field
	Expiration    apijson.Field
	Identity      apijson.Field
	NotSharable   apijson.Field
	ReadOnly      apijson.Field
	RuleSettings  apijson.Field
	Schedule      apijson.Field
	SourceAccount apijson.Field
	UpdatedAt     apijson.Field
	Version       apijson.Field
	WarningStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleJSON) RawJSON() string {
	return r.raw
}

// Additional settings that modify the rule's action.
type RuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders map[string][]string `json:"add_headers,nullable"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass,nullable"`
	// Settings for the Audit SSH action.
	AuditSSH RuleSettingsAuditSSH `json:"audit_ssh,nullable"`
	// Configure how browser isolation behaves.
	BisoAdminControls RuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Custom block page settings. If missing/null, blocking will use the the account
	// settings.
	BlockPage RuleSettingsBlockPage `json:"block_page,nullable"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule,nullable"`
	// Configure how session check behaves.
	CheckSession RuleSettingsCheckSession `json:"check_session,nullable"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// are set. DNS queries will route to the address closest to their origin. Only
	// valid when a rule's action is set to 'resolve'.
	DNSResolvers RuleSettingsDNSResolvers `json:"dns_resolvers,nullable"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress RuleSettingsEgress `json:"egress,nullable"`
	// Set to true, to ignore the category matches at CNAME domains in a response. If
	// unchecked, the categories in this rule will be checked against all the CNAME
	// domain categories in a response.
	IgnoreCnameCategoryMatches bool `json:"ignore_cname_category_matches"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override RuleSettingsL4override `json:"l4override,nullable"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings RuleSettingsNotificationSettings `json:"notification_settings,nullable"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips,nullable"`
	// Configure DLP payload logging.
	PayloadLog RuleSettingsPayloadLog `json:"payload_log,nullable"`
	// Settings that apply to quarantine rules
	Quarantine RuleSettingsQuarantine `json:"quarantine,nullable"`
	// Settings that apply to redirect rules
	Redirect RuleSettingsRedirect `json:"redirect,nullable"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
	// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
	// set to 'resolve'.
	ResolveDNSInternally RuleSettingsResolveDNSInternally `json:"resolve_dns_internally,nullable"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when 'dns_resolvers' are specified or
	// 'resolve_dns_internally' is set. Only valid when a rule's action is set to
	// 'resolve'.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare,nullable"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert RuleSettingsUntrustedCert `json:"untrusted_cert,nullable"`
	JSON          ruleSettingsJSON          `json:"-"`
}

// ruleSettingsJSON contains the JSON metadata for the struct [RuleSettings]
type ruleSettingsJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BisoAdminControls               apijson.Field
	BlockPage                       apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	DNSResolvers                    apijson.Field
	Egress                          apijson.Field
	IgnoreCnameCategoryMatches      apijson.Field
	InsecureDisableDnssecValidation apijson.Field
	IPCategories                    apijson.Field
	IPIndicatorFeeds                apijson.Field
	L4override                      apijson.Field
	NotificationSettings            apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	Quarantine                      apijson.Field
	Redirect                        apijson.Field
	ResolveDNSInternally            apijson.Field
	ResolveDNSThroughCloudflare     apijson.Field
	UntrustedCert                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *RuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type RuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                     `json:"command_logging"`
	JSON           ruleSettingsAuditSSHJSON `json:"-"`
}

// ruleSettingsAuditSSHJSON contains the JSON metadata for the struct
// [RuleSettingsAuditSSH]
type ruleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type RuleSettingsBisoAdminControls struct {
	// Configure whether copy is enabled or not. When set with "remote_only", copying
	// isolated content from the remote browser to the user's local clipboard is
	// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
	Copy RuleSettingsBisoAdminControlsCopy `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	Dcp bool `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	Dd bool `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	Dk bool `json:"dk"`
	// Configure whether downloading enabled or not. When set with "remote_only",
	// downloads are only available for viewing. Only applies when `version == "v2"`.
	Download RuleSettingsBisoAdminControlsDownload `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	Dp bool `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	Du bool `json:"du"`
	// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
	// is enabled. Only applies when `version == "v2"`.
	Keyboard RuleSettingsBisoAdminControlsKeyboard `json:"keyboard"`
	// Configure whether pasting is enabled or not. When set with "remote_only",
	// pasting content from the user's local clipboard into isolated pages is disabled.
	// When absent, paste is enabled. Only applies when `version == "v2"`.
	Paste RuleSettingsBisoAdminControlsPaste `json:"paste"`
	// Configure whether printing is enabled or not. When absent, printing is enabled.
	// Only applies when `version == "v2"`.
	Printing RuleSettingsBisoAdminControlsPrinting `json:"printing"`
	// Configure whether uploading is enabled or not. When absent, uploading is
	// enabled. Only applies when `version == "v2"`.
	Upload RuleSettingsBisoAdminControlsUpload `json:"upload"`
	// Indicates which version of the browser isolation controls should apply.
	Version RuleSettingsBisoAdminControlsVersion `json:"version"`
	JSON    ruleSettingsBisoAdminControlsJSON    `json:"-"`
}

// ruleSettingsBisoAdminControlsJSON contains the JSON metadata for the struct
// [RuleSettingsBisoAdminControls]
type ruleSettingsBisoAdminControlsJSON struct {
	Copy        apijson.Field
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Download    apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	Keyboard    apijson.Field
	Paste       apijson.Field
	Printing    apijson.Field
	Upload      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure whether copy is enabled or not. When set with "remote_only", copying
// isolated content from the remote browser to the user's local clipboard is
// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsCopy string

const (
	RuleSettingsBisoAdminControlsCopyEnabled    RuleSettingsBisoAdminControlsCopy = "enabled"
	RuleSettingsBisoAdminControlsCopyDisabled   RuleSettingsBisoAdminControlsCopy = "disabled"
	RuleSettingsBisoAdminControlsCopyRemoteOnly RuleSettingsBisoAdminControlsCopy = "remote_only"
)

func (r RuleSettingsBisoAdminControlsCopy) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsCopyEnabled, RuleSettingsBisoAdminControlsCopyDisabled, RuleSettingsBisoAdminControlsCopyRemoteOnly:
		return true
	}
	return false
}

// Configure whether downloading enabled or not. When set with "remote_only",
// downloads are only available for viewing. Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsDownload string

const (
	RuleSettingsBisoAdminControlsDownloadEnabled    RuleSettingsBisoAdminControlsDownload = "enabled"
	RuleSettingsBisoAdminControlsDownloadDisabled   RuleSettingsBisoAdminControlsDownload = "disabled"
	RuleSettingsBisoAdminControlsDownloadRemoteOnly RuleSettingsBisoAdminControlsDownload = "remote_only"
)

func (r RuleSettingsBisoAdminControlsDownload) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsDownloadEnabled, RuleSettingsBisoAdminControlsDownloadDisabled, RuleSettingsBisoAdminControlsDownloadRemoteOnly:
		return true
	}
	return false
}

// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
// is enabled. Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsKeyboard string

const (
	RuleSettingsBisoAdminControlsKeyboardEnabled  RuleSettingsBisoAdminControlsKeyboard = "enabled"
	RuleSettingsBisoAdminControlsKeyboardDisabled RuleSettingsBisoAdminControlsKeyboard = "disabled"
)

func (r RuleSettingsBisoAdminControlsKeyboard) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsKeyboardEnabled, RuleSettingsBisoAdminControlsKeyboardDisabled:
		return true
	}
	return false
}

// Configure whether pasting is enabled or not. When set with "remote_only",
// pasting content from the user's local clipboard into isolated pages is disabled.
// When absent, paste is enabled. Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsPaste string

const (
	RuleSettingsBisoAdminControlsPasteEnabled    RuleSettingsBisoAdminControlsPaste = "enabled"
	RuleSettingsBisoAdminControlsPasteDisabled   RuleSettingsBisoAdminControlsPaste = "disabled"
	RuleSettingsBisoAdminControlsPasteRemoteOnly RuleSettingsBisoAdminControlsPaste = "remote_only"
)

func (r RuleSettingsBisoAdminControlsPaste) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsPasteEnabled, RuleSettingsBisoAdminControlsPasteDisabled, RuleSettingsBisoAdminControlsPasteRemoteOnly:
		return true
	}
	return false
}

// Configure whether printing is enabled or not. When absent, printing is enabled.
// Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsPrinting string

const (
	RuleSettingsBisoAdminControlsPrintingEnabled  RuleSettingsBisoAdminControlsPrinting = "enabled"
	RuleSettingsBisoAdminControlsPrintingDisabled RuleSettingsBisoAdminControlsPrinting = "disabled"
)

func (r RuleSettingsBisoAdminControlsPrinting) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsPrintingEnabled, RuleSettingsBisoAdminControlsPrintingDisabled:
		return true
	}
	return false
}

// Configure whether uploading is enabled or not. When absent, uploading is
// enabled. Only applies when `version == "v2"`.
type RuleSettingsBisoAdminControlsUpload string

const (
	RuleSettingsBisoAdminControlsUploadEnabled  RuleSettingsBisoAdminControlsUpload = "enabled"
	RuleSettingsBisoAdminControlsUploadDisabled RuleSettingsBisoAdminControlsUpload = "disabled"
)

func (r RuleSettingsBisoAdminControlsUpload) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsUploadEnabled, RuleSettingsBisoAdminControlsUploadDisabled:
		return true
	}
	return false
}

// Indicates which version of the browser isolation controls should apply.
type RuleSettingsBisoAdminControlsVersion string

const (
	RuleSettingsBisoAdminControlsVersionV1 RuleSettingsBisoAdminControlsVersion = "v1"
	RuleSettingsBisoAdminControlsVersionV2 RuleSettingsBisoAdminControlsVersion = "v2"
)

func (r RuleSettingsBisoAdminControlsVersion) IsKnown() bool {
	switch r {
	case RuleSettingsBisoAdminControlsVersionV1, RuleSettingsBisoAdminControlsVersionV2:
		return true
	}
	return false
}

// Custom block page settings. If missing/null, blocking will use the the account
// settings.
type RuleSettingsBlockPage struct {
	// URI to which the user will be redirected
	TargetUri string `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext bool                      `json:"include_context"`
	JSON           ruleSettingsBlockPageJSON `json:"-"`
}

// ruleSettingsBlockPageJSON contains the JSON metadata for the struct
// [RuleSettingsBlockPage]
type ruleSettingsBlockPageJSON struct {
	TargetUri      apijson.Field
	IncludeContext apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type RuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                         `json:"enforce"`
	JSON    ruleSettingsCheckSessionJSON `json:"-"`
}

// ruleSettingsCheckSessionJSON contains the JSON metadata for the struct
// [RuleSettingsCheckSession]
type ruleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// are set. DNS queries will route to the address closest to their origin. Only
// valid when a rule's action is set to 'resolve'.
type RuleSettingsDNSResolvers struct {
	Ipv4 []RuleSettingsDNSResolversIpv4 `json:"ipv4"`
	Ipv6 []RuleSettingsDNSResolversIpv6 `json:"ipv6"`
	JSON ruleSettingsDNSResolversJSON   `json:"-"`
}

// ruleSettingsDNSResolversJSON contains the JSON metadata for the struct
// [RuleSettingsDNSResolvers]
type ruleSettingsDNSResolversJSON struct {
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
}

type RuleSettingsDNSResolversIpv4 struct {
	// IPv4 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                           `json:"vnet_id"`
	JSON   ruleSettingsDNSResolversIpv4JSON `json:"-"`
}

// ruleSettingsDNSResolversIpv4JSON contains the JSON metadata for the struct
// [RuleSettingsDNSResolversIpv4]
type ruleSettingsDNSResolversIpv4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *RuleSettingsDNSResolversIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsDNSResolversIpv4JSON) RawJSON() string {
	return r.raw
}

type RuleSettingsDNSResolversIpv6 struct {
	// IPv6 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                           `json:"vnet_id"`
	JSON   ruleSettingsDNSResolversIpv6JSON `json:"-"`
}

// ruleSettingsDNSResolversIpv6JSON contains the JSON metadata for the struct
// [RuleSettingsDNSResolversIpv6]
type ruleSettingsDNSResolversIpv6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *RuleSettingsDNSResolversIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsDNSResolversIpv6JSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	Ipv4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 string                 `json:"ipv6"`
	JSON ruleSettingsEgressJSON `json:"-"`
}

// ruleSettingsEgressJSON contains the JSON metadata for the struct
// [RuleSettingsEgress]
type ruleSettingsEgressJSON struct {
	Ipv4         apijson.Field
	Ipv4Fallback apijson.Field
	Ipv6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                      `json:"port"`
	JSON ruleSettingsL4overrideJSON `json:"-"`
}

// ruleSettingsL4overrideJSON contains the JSON metadata for the struct
// [RuleSettingsL4override]
type ruleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                               `json:"support_url"`
	JSON       ruleSettingsNotificationSettingsJSON `json:"-"`
}

// ruleSettingsNotificationSettingsJSON contains the JSON metadata for the struct
// [RuleSettingsNotificationSettings]
type ruleSettingsNotificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type RuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                       `json:"enabled"`
	JSON    ruleSettingsPayloadLogJSON `json:"-"`
}

// ruleSettingsPayloadLogJSON contains the JSON metadata for the struct
// [RuleSettingsPayloadLog]
type ruleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Settings that apply to quarantine rules
type RuleSettingsQuarantine struct {
	// Types of files to sandbox.
	FileTypes []RuleSettingsQuarantineFileType `json:"file_types"`
	JSON      ruleSettingsQuarantineJSON       `json:"-"`
}

// ruleSettingsQuarantineJSON contains the JSON metadata for the struct
// [RuleSettingsQuarantine]
type ruleSettingsQuarantineJSON struct {
	FileTypes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsQuarantine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsQuarantineJSON) RawJSON() string {
	return r.raw
}

type RuleSettingsQuarantineFileType string

const (
	RuleSettingsQuarantineFileTypeExe  RuleSettingsQuarantineFileType = "exe"
	RuleSettingsQuarantineFileTypePdf  RuleSettingsQuarantineFileType = "pdf"
	RuleSettingsQuarantineFileTypeDoc  RuleSettingsQuarantineFileType = "doc"
	RuleSettingsQuarantineFileTypeDocm RuleSettingsQuarantineFileType = "docm"
	RuleSettingsQuarantineFileTypeDocx RuleSettingsQuarantineFileType = "docx"
	RuleSettingsQuarantineFileTypeRtf  RuleSettingsQuarantineFileType = "rtf"
	RuleSettingsQuarantineFileTypePpt  RuleSettingsQuarantineFileType = "ppt"
	RuleSettingsQuarantineFileTypePptx RuleSettingsQuarantineFileType = "pptx"
	RuleSettingsQuarantineFileTypeXls  RuleSettingsQuarantineFileType = "xls"
	RuleSettingsQuarantineFileTypeXlsm RuleSettingsQuarantineFileType = "xlsm"
	RuleSettingsQuarantineFileTypeXlsx RuleSettingsQuarantineFileType = "xlsx"
	RuleSettingsQuarantineFileTypeZip  RuleSettingsQuarantineFileType = "zip"
	RuleSettingsQuarantineFileTypeRar  RuleSettingsQuarantineFileType = "rar"
)

func (r RuleSettingsQuarantineFileType) IsKnown() bool {
	switch r {
	case RuleSettingsQuarantineFileTypeExe, RuleSettingsQuarantineFileTypePdf, RuleSettingsQuarantineFileTypeDoc, RuleSettingsQuarantineFileTypeDocm, RuleSettingsQuarantineFileTypeDocx, RuleSettingsQuarantineFileTypeRtf, RuleSettingsQuarantineFileTypePpt, RuleSettingsQuarantineFileTypePptx, RuleSettingsQuarantineFileTypeXls, RuleSettingsQuarantineFileTypeXlsm, RuleSettingsQuarantineFileTypeXlsx, RuleSettingsQuarantineFileTypeZip, RuleSettingsQuarantineFileTypeRar:
		return true
	}
	return false
}

// Settings that apply to redirect rules
type RuleSettingsRedirect struct {
	// URI to which the user will be redirected
	TargetUri string `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// If true, the path and query parameters from the original request will be
	// appended to target_uri
	PreservePathAndQuery bool                     `json:"preserve_path_and_query"`
	JSON                 ruleSettingsRedirectJSON `json:"-"`
}

// ruleSettingsRedirectJSON contains the JSON metadata for the struct
// [RuleSettingsRedirect]
type ruleSettingsRedirectJSON struct {
	TargetUri            apijson.Field
	IncludeContext       apijson.Field
	PreservePathAndQuery apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RuleSettingsRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsRedirectJSON) RawJSON() string {
	return r.raw
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingsResolveDNSInternally struct {
	// The fallback behavior to apply when the internal DNS response code is different
	// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
	// 'AAAA' queries.
	Fallback RuleSettingsResolveDNSInternallyFallback `json:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	ViewID string                               `json:"view_id"`
	JSON   ruleSettingsResolveDNSInternallyJSON `json:"-"`
}

// ruleSettingsResolveDNSInternallyJSON contains the JSON metadata for the struct
// [RuleSettingsResolveDNSInternally]
type ruleSettingsResolveDNSInternallyJSON struct {
	Fallback    apijson.Field
	ViewID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsResolveDNSInternally) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsResolveDNSInternallyJSON) RawJSON() string {
	return r.raw
}

// The fallback behavior to apply when the internal DNS response code is different
// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
// 'AAAA' queries.
type RuleSettingsResolveDNSInternallyFallback string

const (
	RuleSettingsResolveDNSInternallyFallbackNone      RuleSettingsResolveDNSInternallyFallback = "none"
	RuleSettingsResolveDNSInternallyFallbackPublicDNS RuleSettingsResolveDNSInternallyFallback = "public_dns"
)

func (r RuleSettingsResolveDNSInternallyFallback) IsKnown() bool {
	switch r {
	case RuleSettingsResolveDNSInternallyFallbackNone, RuleSettingsResolveDNSInternallyFallbackPublicDNS:
		return true
	}
	return false
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingsUntrustedCert struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action RuleSettingsUntrustedCertAction `json:"action"`
	JSON   ruleSettingsUntrustedCertJSON   `json:"-"`
}

// ruleSettingsUntrustedCertJSON contains the JSON metadata for the struct
// [RuleSettingsUntrustedCert]
type ruleSettingsUntrustedCertJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingsUntrustedCert) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingsUntrustedCertJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type RuleSettingsUntrustedCertAction string

const (
	RuleSettingsUntrustedCertActionPassThrough RuleSettingsUntrustedCertAction = "pass_through"
	RuleSettingsUntrustedCertActionBlock       RuleSettingsUntrustedCertAction = "block"
	RuleSettingsUntrustedCertActionError       RuleSettingsUntrustedCertAction = "error"
)

func (r RuleSettingsUntrustedCertAction) IsKnown() bool {
	switch r {
	case RuleSettingsUntrustedCertActionPassThrough, RuleSettingsUntrustedCertActionBlock, RuleSettingsUntrustedCertActionError:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type RuleSettingsParam struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[map[string][]string] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[RuleSettingsAuditSSHParam] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[RuleSettingsBisoAdminControlsParam] `json:"biso_admin_controls"`
	// Custom block page settings. If missing/null, blocking will use the the account
	// settings.
	BlockPage param.Field[RuleSettingsBlockPageParam] `json:"block_page"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[RuleSettingsCheckSessionParam] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// are set. DNS queries will route to the address closest to their origin. Only
	// valid when a rule's action is set to 'resolve'.
	DNSResolvers param.Field[RuleSettingsDNSResolversParam] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[RuleSettingsEgressParam] `json:"egress"`
	// Set to true, to ignore the category matches at CNAME domains in a response. If
	// unchecked, the categories in this rule will be checked against all the CNAME
	// domain categories in a response.
	IgnoreCnameCategoryMatches param.Field[bool] `json:"ignore_cname_category_matches"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDnssecValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[RuleSettingsL4overrideParam] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[RuleSettingsNotificationSettingsParam] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[RuleSettingsPayloadLogParam] `json:"payload_log"`
	// Settings that apply to quarantine rules
	Quarantine param.Field[RuleSettingsQuarantineParam] `json:"quarantine"`
	// Settings that apply to redirect rules
	Redirect param.Field[RuleSettingsRedirectParam] `json:"redirect"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
	// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
	// set to 'resolve'.
	ResolveDNSInternally param.Field[RuleSettingsResolveDNSInternallyParam] `json:"resolve_dns_internally"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when 'dns_resolvers' are specified or
	// 'resolve_dns_internally' is set. Only valid when a rule's action is set to
	// 'resolve'.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCert param.Field[RuleSettingsUntrustedCertParam] `json:"untrusted_cert"`
}

func (r RuleSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type RuleSettingsAuditSSHParam struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r RuleSettingsAuditSSHParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type RuleSettingsBisoAdminControlsParam struct {
	// Configure whether copy is enabled or not. When set with "remote_only", copying
	// isolated content from the remote browser to the user's local clipboard is
	// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
	Copy param.Field[RuleSettingsBisoAdminControlsCopy] `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	Dcp param.Field[bool] `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	Dd param.Field[bool] `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	Dk param.Field[bool] `json:"dk"`
	// Configure whether downloading enabled or not. When set with "remote_only",
	// downloads are only available for viewing. Only applies when `version == "v2"`.
	Download param.Field[RuleSettingsBisoAdminControlsDownload] `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	Dp param.Field[bool] `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	Du param.Field[bool] `json:"du"`
	// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
	// is enabled. Only applies when `version == "v2"`.
	Keyboard param.Field[RuleSettingsBisoAdminControlsKeyboard] `json:"keyboard"`
	// Configure whether pasting is enabled or not. When set with "remote_only",
	// pasting content from the user's local clipboard into isolated pages is disabled.
	// When absent, paste is enabled. Only applies when `version == "v2"`.
	Paste param.Field[RuleSettingsBisoAdminControlsPaste] `json:"paste"`
	// Configure whether printing is enabled or not. When absent, printing is enabled.
	// Only applies when `version == "v2"`.
	Printing param.Field[RuleSettingsBisoAdminControlsPrinting] `json:"printing"`
	// Configure whether uploading is enabled or not. When absent, uploading is
	// enabled. Only applies when `version == "v2"`.
	Upload param.Field[RuleSettingsBisoAdminControlsUpload] `json:"upload"`
	// Indicates which version of the browser isolation controls should apply.
	Version param.Field[RuleSettingsBisoAdminControlsVersion] `json:"version"`
}

func (r RuleSettingsBisoAdminControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom block page settings. If missing/null, blocking will use the the account
// settings.
type RuleSettingsBlockPageParam struct {
	// URI to which the user will be redirected
	TargetUri param.Field[string] `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
}

func (r RuleSettingsBlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type RuleSettingsCheckSessionParam struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r RuleSettingsCheckSessionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// are set. DNS queries will route to the address closest to their origin. Only
// valid when a rule's action is set to 'resolve'.
type RuleSettingsDNSResolversParam struct {
	Ipv4 param.Field[[]RuleSettingsDNSResolversIpv4Param] `json:"ipv4"`
	Ipv6 param.Field[[]RuleSettingsDNSResolversIpv6Param] `json:"ipv6"`
}

func (r RuleSettingsDNSResolversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleSettingsDNSResolversIpv4Param struct {
	// IPv4 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r RuleSettingsDNSResolversIpv4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleSettingsDNSResolversIpv6Param struct {
	// IPv6 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r RuleSettingsDNSResolversIpv6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingsEgressParam struct {
	// The IPv4 address to be used for egress.
	Ipv4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	Ipv4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	Ipv6 param.Field[string] `json:"ipv6"`
}

func (r RuleSettingsEgressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingsL4overrideParam struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r RuleSettingsL4overrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingsNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r RuleSettingsNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type RuleSettingsPayloadLogParam struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleSettingsPayloadLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings that apply to quarantine rules
type RuleSettingsQuarantineParam struct {
	// Types of files to sandbox.
	FileTypes param.Field[[]RuleSettingsQuarantineFileType] `json:"file_types"`
}

func (r RuleSettingsQuarantineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings that apply to redirect rules
type RuleSettingsRedirectParam struct {
	// URI to which the user will be redirected
	TargetUri param.Field[string] `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
	// If true, the path and query parameters from the original request will be
	// appended to target_uri
	PreservePathAndQuery param.Field[bool] `json:"preserve_path_and_query"`
}

func (r RuleSettingsRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingsResolveDNSInternallyParam struct {
	// The fallback behavior to apply when the internal DNS response code is different
	// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
	// 'AAAA' queries.
	Fallback param.Field[RuleSettingsResolveDNSInternallyFallback] `json:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	ViewID param.Field[string] `json:"view_id"`
}

func (r RuleSettingsResolveDNSInternallyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingsUntrustedCertParam struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[RuleSettingsUntrustedCertAction] `json:"action"`
}

func (r RuleSettingsUntrustedCertParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type Schedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri string `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon string `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat string `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun string `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu string `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string       `json:"wed"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Fri         apijson.Field
	Mon         apijson.Field
	Sat         apijson.Field
	Sun         apijson.Field
	Thu         apijson.Field
	TimeZone    apijson.Field
	Tue         apijson.Field
	Wed         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ScheduleParam struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri param.Field[string] `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon param.Field[string] `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat param.Field[string] `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun param.Field[string] `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu param.Field[string] `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone param.Field[string] `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue param.Field[string] `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed param.Field[string] `json:"wed"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleListResponse struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayRuleListResponseSuccess    `json:"success,required"`
	Result     []GatewayRule                            `json:"result"`
	ResultInfo AccountGatewayRuleListResponseResultInfo `json:"result_info"`
	JSON       accountGatewayRuleListResponseJSON       `json:"-"`
}

// accountGatewayRuleListResponseJSON contains the JSON metadata for the struct
// [AccountGatewayRuleListResponse]
type accountGatewayRuleListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayRuleListResponseSuccess bool

const (
	AccountGatewayRuleListResponseSuccessTrue AccountGatewayRuleListResponseSuccess = true
)

func (r AccountGatewayRuleListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayRuleListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayRuleListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       accountGatewayRuleListResponseResultInfoJSON `json:"-"`
}

// accountGatewayRuleListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountGatewayRuleListResponseResultInfo]
type accountGatewayRuleListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayRuleListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayRuleListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayRuleNewParams struct {
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[ActionPerform] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration param.Field[ExpirationParam] `json:"expiration"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayFilters] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingsParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayRuleUpdateParams struct {
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[ActionPerform] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration param.Field[ExpirationParam] `json:"expiration"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayFilters] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingsParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r AccountGatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
