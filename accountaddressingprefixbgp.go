// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAddressingPrefixBgpService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAddressingPrefixBgpService] method instead.
type AccountAddressingPrefixBgpService struct {
	Options  []option.RequestOption
	Prefixes *AccountAddressingPrefixBgpPrefixService
	Status   *AccountAddressingPrefixBgpStatusService
}

// NewAccountAddressingPrefixBgpService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingPrefixBgpService(opts ...option.RequestOption) (r *AccountAddressingPrefixBgpService) {
	r = &AccountAddressingPrefixBgpService{}
	r.Options = opts
	r.Prefixes = NewAccountAddressingPrefixBgpPrefixService(opts...)
	r.Status = NewAccountAddressingPrefixBgpStatusService(opts...)
	return
}
