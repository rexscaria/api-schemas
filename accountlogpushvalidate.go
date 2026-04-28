// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountLogpushValidateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushValidateService] method instead.
type AccountLogpushValidateService struct {
	Options     []option.RequestOption
	Destination *AccountLogpushValidateDestinationService
}

// NewAccountLogpushValidateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushValidateService(opts ...option.RequestOption) (r *AccountLogpushValidateService) {
	r = &AccountLogpushValidateService{}
	r.Options = opts
	r.Destination = NewAccountLogpushValidateDestinationService(opts...)
	return
}
