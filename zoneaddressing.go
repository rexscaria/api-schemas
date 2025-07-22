// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// ZoneAddressingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneAddressingService] method instead.
type ZoneAddressingService struct {
	Options           []option.RequestOption
	RegionalHostnames *ZoneAddressingRegionalHostnameService
}

// NewZoneAddressingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneAddressingService(opts ...option.RequestOption) (r *ZoneAddressingService) {
	r = &ZoneAddressingService{}
	r.Options = opts
	r.RegionalHostnames = NewZoneAddressingRegionalHostnameService(opts...)
	return
}
