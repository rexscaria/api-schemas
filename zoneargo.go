// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneArgoService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneArgoService] method instead.
type ZoneArgoService struct {
	Options       []option.RequestOption
	SmartRouting  *ZoneArgoSmartRoutingService
	TieredCaching *ZoneArgoTieredCachingService
}

// NewZoneArgoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneArgoService(opts ...option.RequestOption) (r *ZoneArgoService) {
	r = &ZoneArgoService{}
	r.Options = opts
	r.SmartRouting = NewZoneArgoSmartRoutingService(opts...)
	r.TieredCaching = NewZoneArgoTieredCachingService(opts...)
	return
}
