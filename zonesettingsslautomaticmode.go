// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// ZoneSettingSslAutomaticModeService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSettingSslAutomaticModeService] method instead.
type ZoneSettingSslAutomaticModeService struct {
	Options []option.RequestOption
}

// NewZoneSettingSslAutomaticModeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSettingSslAutomaticModeService(opts ...option.RequestOption) (r *ZoneSettingSslAutomaticModeService) {
	r = &ZoneSettingSslAutomaticModeService{}
	r.Options = opts
	return
}
