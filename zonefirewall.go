// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneFirewallService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneFirewallService] method instead.
type ZoneFirewallService struct {
	Options     []option.RequestOption
	AccessRules *ZoneFirewallAccessRuleService
	Lockdowns   *ZoneFirewallLockdownService
	Rules       *ZoneFirewallRuleService
	UaRules     *ZoneFirewallUaRuleService
	Waf         *ZoneFirewallWafService
}

// NewZoneFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneFirewallService(opts ...option.RequestOption) (r *ZoneFirewallService) {
	r = &ZoneFirewallService{}
	r.Options = opts
	r.AccessRules = NewZoneFirewallAccessRuleService(opts...)
	r.Lockdowns = NewZoneFirewallLockdownService(opts...)
	r.Rules = NewZoneFirewallRuleService(opts...)
	r.UaRules = NewZoneFirewallUaRuleService(opts...)
	r.Waf = NewZoneFirewallWafService(opts...)
	return
}
