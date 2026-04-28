// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountSlurperService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSlurperService] method instead.
type AccountSlurperService struct {
	Options []option.RequestOption
	Jobs    *AccountSlurperJobService
	Source  *AccountSlurperSourceService
	Target  *AccountSlurperTargetService
}

// NewAccountSlurperService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountSlurperService(opts ...option.RequestOption) (r *AccountSlurperService) {
	r = &AccountSlurperService{}
	r.Options = opts
	r.Jobs = NewAccountSlurperJobService(opts...)
	r.Source = NewAccountSlurperSourceService(opts...)
	r.Target = NewAccountSlurperTargetService(opts...)
	return
}
