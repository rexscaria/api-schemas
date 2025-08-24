// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountDNSFirewallDNSAnalyticsService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallDNSAnalyticsService] method instead.
type AccountDNSFirewallDNSAnalyticsService struct {
	Options []option.RequestOption
	Report  *AccountDNSFirewallDNSAnalyticsReportService
}

// NewAccountDNSFirewallDNSAnalyticsService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDNSFirewallDNSAnalyticsService(opts ...option.RequestOption) (r *AccountDNSFirewallDNSAnalyticsService) {
	r = &AccountDNSFirewallDNSAnalyticsService{}
	r.Options = opts
	r.Report = NewAccountDNSFirewallDNSAnalyticsReportService(opts...)
	return
}
