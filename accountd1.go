// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountD1Service contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountD1Service] method instead.
type AccountD1Service struct {
	Options  []option.RequestOption
	Database *AccountD1DatabaseService
}

// NewAccountD1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountD1Service(opts ...option.RequestOption) (r *AccountD1Service) {
	r = &AccountD1Service{}
	r.Options = opts
	r.Database = NewAccountD1DatabaseService(opts...)
	return
}
