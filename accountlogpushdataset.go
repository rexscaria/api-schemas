// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountLogpushDatasetService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountLogpushDatasetService] method instead.
type AccountLogpushDatasetService struct {
	Options []option.RequestOption
}

// NewAccountLogpushDatasetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushDatasetService(opts ...option.RequestOption) (r *AccountLogpushDatasetService) {
	r = &AccountLogpushDatasetService{}
	r.Options = opts
	return
}
