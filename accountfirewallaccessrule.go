// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountFirewallAccessRuleService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountFirewallAccessRuleService] method instead.
type AccountFirewallAccessRuleService struct {
	Options []option.RequestOption
	Rules   *AccountFirewallAccessRuleRuleService
}

// NewAccountFirewallAccessRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountFirewallAccessRuleService(opts ...option.RequestOption) (r *AccountFirewallAccessRuleService) {
	r = &AccountFirewallAccessRuleService{}
	r.Options = opts
	r.Rules = NewAccountFirewallAccessRuleRuleService(opts...)
	return
}
