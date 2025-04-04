// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountDeviceResilienceService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDeviceResilienceService] method instead.
type AccountDeviceResilienceService struct {
	Options    []option.RequestOption
	Disconnect *AccountDeviceResilienceDisconnectService
}

// NewAccountDeviceResilienceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceResilienceService(opts ...option.RequestOption) (r *AccountDeviceResilienceService) {
	r = &AccountDeviceResilienceService{}
	r.Options = opts
	r.Disconnect = NewAccountDeviceResilienceDisconnectService(opts...)
	return
}
