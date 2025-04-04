// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountDNSFirewallDNSAnalyticService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallDNSAnalyticService] method instead.
type AccountDNSFirewallDNSAnalyticService struct {
	Options []option.RequestOption
	Report  *AccountDNSFirewallDNSAnalyticReportService
}

// NewAccountDNSFirewallDNSAnalyticService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDNSFirewallDNSAnalyticService(opts ...option.RequestOption) (r *AccountDNSFirewallDNSAnalyticService) {
	r = &AccountDNSFirewallDNSAnalyticService{}
	r.Options = opts
	r.Report = NewAccountDNSFirewallDNSAnalyticReportService(opts...)
	return
}
