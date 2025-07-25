// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessAppCaService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessAppCaService] method instead.
type AccountAccessAppCaService struct {
	Options []option.RequestOption
}

// NewAccountAccessAppCaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessAppCaService(opts ...option.RequestOption) (r *AccountAccessAppCaService) {
	r = &AccountAccessAppCaService{}
	r.Options = opts
	return
}
