// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountLogpushService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushService] method instead.
type AccountLogpushService struct {
	Options   []option.RequestOption
	Datasets  *AccountLogpushDatasetService
	Jobs      *AccountLogpushJobService
	Ownership *AccountLogpushOwnershipService
	Validate  *AccountLogpushValidateService
}

// NewAccountLogpushService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountLogpushService(opts ...option.RequestOption) (r *AccountLogpushService) {
	r = &AccountLogpushService{}
	r.Options = opts
	r.Datasets = NewAccountLogpushDatasetService(opts...)
	r.Jobs = NewAccountLogpushJobService(opts...)
	r.Ownership = NewAccountLogpushOwnershipService(opts...)
	r.Validate = NewAccountLogpushValidateService(opts...)
	return
}
