// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneDeviceService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDeviceService] method instead.
type ZoneDeviceService struct {
	Options []option.RequestOption
	Policy  *ZoneDevicePolicyService
}

// NewZoneDeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDeviceService(opts ...option.RequestOption) (r *ZoneDeviceService) {
	r = &ZoneDeviceService{}
	r.Options = opts
	r.Policy = NewZoneDevicePolicyService(opts...)
	return
}
