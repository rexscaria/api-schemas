// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountR2BucketDomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountR2BucketDomainService] method instead.
type AccountR2BucketDomainService struct {
	Options []option.RequestOption
	Custom  *AccountR2BucketDomainCustomService
	Managed *AccountR2BucketDomainManagedService
}

// NewAccountR2BucketDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountR2BucketDomainService(opts ...option.RequestOption) (r *AccountR2BucketDomainService) {
	r = &AccountR2BucketDomainService{}
	r.Options = opts
	r.Custom = NewAccountR2BucketDomainCustomService(opts...)
	r.Managed = NewAccountR2BucketDomainManagedService(opts...)
	return
}
