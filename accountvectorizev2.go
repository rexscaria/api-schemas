// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountVectorizeV2Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountVectorizeV2Service] method instead.
type AccountVectorizeV2Service struct {
	Options []option.RequestOption
	Indexes *AccountVectorizeV2IndexService
}

// NewAccountVectorizeV2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountVectorizeV2Service(opts ...option.RequestOption) (r *AccountVectorizeV2Service) {
	r = &AccountVectorizeV2Service{}
	r.Options = opts
	r.Indexes = NewAccountVectorizeV2IndexService(opts...)
	return
}
