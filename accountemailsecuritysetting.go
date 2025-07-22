// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountEmailSecuritySettingService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingService] method instead.
type AccountEmailSecuritySettingService struct {
	Options               []option.RequestOption
	AllowPolicies         *AccountEmailSecuritySettingAllowPolicyService
	BlockSenders          *AccountEmailSecuritySettingBlockSenderService
	Domains               *AccountEmailSecuritySettingDomainService
	ImpersonationRegistry *AccountEmailSecuritySettingImpersonationRegistryService
	TrustedDomains        *AccountEmailSecuritySettingTrustedDomainService
}

// NewAccountEmailSecuritySettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEmailSecuritySettingService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingService) {
	r = &AccountEmailSecuritySettingService{}
	r.Options = opts
	r.AllowPolicies = NewAccountEmailSecuritySettingAllowPolicyService(opts...)
	r.BlockSenders = NewAccountEmailSecuritySettingBlockSenderService(opts...)
	r.Domains = NewAccountEmailSecuritySettingDomainService(opts...)
	r.ImpersonationRegistry = NewAccountEmailSecuritySettingImpersonationRegistryService(opts...)
	r.TrustedDomains = NewAccountEmailSecuritySettingTrustedDomainService(opts...)
	return
}
