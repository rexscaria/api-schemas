// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountEventNotificationR2Service contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEventNotificationR2Service] method instead.
type AccountEventNotificationR2Service struct {
	Options       []option.RequestOption
	Configuration *AccountEventNotificationR2ConfigurationService
}

// NewAccountEventNotificationR2Service generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEventNotificationR2Service(opts ...option.RequestOption) (r *AccountEventNotificationR2Service) {
	r = &AccountEventNotificationR2Service{}
	r.Options = opts
	r.Configuration = NewAccountEventNotificationR2ConfigurationService(opts...)
	return
}
