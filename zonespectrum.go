// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneSpectrumService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneSpectrumService] method instead.
type ZoneSpectrumService struct {
	Options   []option.RequestOption
	Analytics *ZoneSpectrumAnalyticsService
	Apps      *ZoneSpectrumAppService
}

// NewZoneSpectrumService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSpectrumService(opts ...option.RequestOption) (r *ZoneSpectrumService) {
	r = &ZoneSpectrumService{}
	r.Options = opts
	r.Analytics = NewZoneSpectrumAnalyticsService(opts...)
	r.Apps = NewZoneSpectrumAppService(opts...)
	return
}
