// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountRulesetPhaseService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRulesetPhaseService] method instead.
type AccountRulesetPhaseService struct {
	Options    []option.RequestOption
	Entrypoint *AccountRulesetPhaseEntrypointService
}

// NewAccountRulesetPhaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetPhaseService(opts ...option.RequestOption) (r *AccountRulesetPhaseService) {
	r = &AccountRulesetPhaseService{}
	r.Options = opts
	r.Entrypoint = NewAccountRulesetPhaseEntrypointService(opts...)
	return
}
