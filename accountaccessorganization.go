// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountAccessOrganizationService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessOrganizationService] method instead.
type AccountAccessOrganizationService struct {
	Options []option.RequestOption
	Doh     *AccountAccessOrganizationDohService
}

// NewAccountAccessOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessOrganizationService(opts ...option.RequestOption) (r *AccountAccessOrganizationService) {
	r = &AccountAccessOrganizationService{}
	r.Options = opts
	r.Doh = NewAccountAccessOrganizationDohService(opts...)
	return
}
