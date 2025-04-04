// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountCallService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCallService] method instead.
type AccountCallService struct {
	Options  []option.RequestOption
	Apps     *AccountCallAppService
	TurnKeys *AccountCallTurnKeyService
}

// NewAccountCallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCallService(opts ...option.RequestOption) (r *AccountCallService) {
	r = &AccountCallService{}
	r.Options = opts
	r.Apps = NewAccountCallAppService(opts...)
	r.TurnKeys = NewAccountCallTurnKeyService(opts...)
	return
}
