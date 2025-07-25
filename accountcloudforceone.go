// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountCloudforceOneService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCloudforceOneService] method instead.
type AccountCloudforceOneService struct {
	Options  []option.RequestOption
	Events   *AccountCloudforceOneEventService
	Scans    *AccountCloudforceOneScanService
	Requests *AccountCloudforceOneRequestService
}

// NewAccountCloudforceOneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCloudforceOneService(opts ...option.RequestOption) (r *AccountCloudforceOneService) {
	r = &AccountCloudforceOneService{}
	r.Options = opts
	r.Events = NewAccountCloudforceOneEventService(opts...)
	r.Scans = NewAccountCloudforceOneScanService(opts...)
	r.Requests = NewAccountCloudforceOneRequestService(opts...)
	return
}
