// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountDlpEmailService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDlpEmailService] method instead.
type AccountDlpEmailService struct {
	Options        []option.RequestOption
	AccountMapping *AccountDlpEmailAccountMappingService
	Rules          *AccountDlpEmailRuleService
}

// NewAccountDlpEmailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountDlpEmailService(opts ...option.RequestOption) (r *AccountDlpEmailService) {
	r = &AccountDlpEmailService{}
	r.Options = opts
	r.AccountMapping = NewAccountDlpEmailAccountMappingService(opts...)
	r.Rules = NewAccountDlpEmailRuleService(opts...)
	return
}
