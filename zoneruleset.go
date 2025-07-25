// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// ZoneRulesetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetService] method instead.
type ZoneRulesetService struct {
	Options  []option.RequestOption
	Rules    *ZoneRulesetRuleService
	Versions *ZoneRulesetVersionService
	Phases   *ZoneRulesetPhaseService
}

// NewZoneRulesetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetService(opts ...option.RequestOption) (r *ZoneRulesetService) {
	r = &ZoneRulesetService{}
	r.Options = opts
	r.Rules = NewZoneRulesetRuleService(opts...)
	r.Versions = NewZoneRulesetVersionService(opts...)
	r.Phases = NewZoneRulesetPhaseService(opts...)
	return
}
