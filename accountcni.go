// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"github.com/rexscaria/api-schemas/option"
)

// AccountCniService contains methods and other services that help with interacting
// with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCniService] method instead.
type AccountCniService struct {
	Options       []option.RequestOption
	Cnis          *AccountCniCniService
	Interconnects *AccountCniInterconnectService
	Settings      *AccountCniSettingService
	Slots         *AccountCniSlotService
}

// NewAccountCniService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCniService(opts ...option.RequestOption) (r *AccountCniService) {
	r = &AccountCniService{}
	r.Options = opts
	r.Cnis = NewAccountCniCniService(opts...)
	r.Interconnects = NewAccountCniInterconnectService(opts...)
	r.Settings = NewAccountCniSettingService(opts...)
	r.Slots = NewAccountCniSlotService(opts...)
	return
}
