// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountRumV2Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRumV2Service] method instead.
type AccountRumV2Service struct {
	Options []option.RequestOption
	Rules   *AccountRumV2RuleService
}

// NewAccountRumV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRumV2Service(opts ...option.RequestOption) (r *AccountRumV2Service) {
	r = &AccountRumV2Service{}
	r.Options = opts
	r.Rules = NewAccountRumV2RuleService(opts...)
	return
}

type RumRule struct {
	// The Web Analytics rule identifier.
	ID      string    `json:"id"`
	Created time.Time `json:"created" format:"date-time"`
	// The hostname the rule will be applied to.
	Host string `json:"host"`
	// Whether the rule includes or excludes traffic from being measured.
	Inclusive bool `json:"inclusive"`
	// Whether the rule is paused or not.
	IsPaused bool `json:"is_paused"`
	// The paths the rule will be applied to.
	Paths    []string    `json:"paths"`
	Priority float64     `json:"priority"`
	JSON     rumRuleJSON `json:"-"`
}

// rumRuleJSON contains the JSON metadata for the struct [RumRule]
type rumRuleJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Host        apijson.Field
	Inclusive   apijson.Field
	IsPaused    apijson.Field
	Paths       apijson.Field
	Priority    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RumRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rumRuleJSON) RawJSON() string {
	return r.raw
}
