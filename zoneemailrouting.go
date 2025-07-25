// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneEmailRoutingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneEmailRoutingService] method instead.
type ZoneEmailRoutingService struct {
	Options []option.RequestOption
	DNS     *ZoneEmailRoutingDNSService
	Rules   *ZoneEmailRoutingRuleService
}

// NewZoneEmailRoutingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneEmailRoutingService(opts ...option.RequestOption) (r *ZoneEmailRoutingService) {
	r = &ZoneEmailRoutingService{}
	r.Options = opts
	r.DNS = NewZoneEmailRoutingDNSService(opts...)
	r.Rules = NewZoneEmailRoutingRuleService(opts...)
	return
}

// Get information about the settings for your Email Routing zone.
func (r *ZoneEmailRoutingService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *EmailEmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Disable your Email Routing zone. Also removes additional MX records previously
// required for Email Routing to work.
//
// Deprecated: deprecated
func (r *ZoneEmailRoutingService) Disable(ctx context.Context, zoneID string, body ZoneEmailRoutingDisableParams, opts ...option.RequestOption) (res *EmailEmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/disable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
//
// Deprecated: deprecated
func (r *ZoneEmailRoutingService) Enable(ctx context.Context, zoneID string, body ZoneEmailRoutingEnableParams, opts ...option.RequestOption) (res *EmailEmailSettingsResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/email/routing/enable", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmailEmailSettingsResponseSingle struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success EmailEmailSettingsResponseSingleSuccess `json:"success,required"`
	Result  EmailEmailSettingsResponseSingleResult  `json:"result"`
	JSON    emailEmailSettingsResponseSingleJSON    `json:"-"`
}

// emailEmailSettingsResponseSingleJSON contains the JSON metadata for the struct
// [EmailEmailSettingsResponseSingle]
type emailEmailSettingsResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailEmailSettingsResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailEmailSettingsResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EmailEmailSettingsResponseSingleSuccess bool

const (
	EmailEmailSettingsResponseSingleSuccessTrue EmailEmailSettingsResponseSingleSuccess = true
)

func (r EmailEmailSettingsResponseSingleSuccess) IsKnown() bool {
	switch r {
	case EmailEmailSettingsResponseSingleSuccessTrue:
		return true
	}
	return false
}

type EmailEmailSettingsResponseSingleResult struct {
	// Email Routing settings identifier.
	ID string `json:"id,required"`
	// State of the zone settings for Email Routing.
	Enabled EmailEmailSettingsResponseSingleResultEnabled `json:"enabled,required"`
	// Domain of your zone.
	Name string `json:"name,required"`
	// The date and time the settings have been created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the settings have been modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard EmailEmailSettingsResponseSingleResultSkipWizard `json:"skip_wizard"`
	// Show the state of your account, and the type or configuration error.
	Status EmailEmailSettingsResponseSingleResultStatus `json:"status"`
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings
	// identifier)
	//
	// Deprecated: deprecated
	Tag  string                                     `json:"tag"`
	JSON emailEmailSettingsResponseSingleResultJSON `json:"-"`
}

// emailEmailSettingsResponseSingleResultJSON contains the JSON metadata for the
// struct [EmailEmailSettingsResponseSingleResult]
type emailEmailSettingsResponseSingleResultJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	Created     apijson.Field
	Modified    apijson.Field
	SkipWizard  apijson.Field
	Status      apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailEmailSettingsResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailEmailSettingsResponseSingleResultJSON) RawJSON() string {
	return r.raw
}

// State of the zone settings for Email Routing.
type EmailEmailSettingsResponseSingleResultEnabled bool

const (
	EmailEmailSettingsResponseSingleResultEnabledTrue  EmailEmailSettingsResponseSingleResultEnabled = true
	EmailEmailSettingsResponseSingleResultEnabledFalse EmailEmailSettingsResponseSingleResultEnabled = false
)

func (r EmailEmailSettingsResponseSingleResultEnabled) IsKnown() bool {
	switch r {
	case EmailEmailSettingsResponseSingleResultEnabledTrue, EmailEmailSettingsResponseSingleResultEnabledFalse:
		return true
	}
	return false
}

// Flag to check if the user skipped the configuration wizard.
type EmailEmailSettingsResponseSingleResultSkipWizard bool

const (
	EmailEmailSettingsResponseSingleResultSkipWizardTrue  EmailEmailSettingsResponseSingleResultSkipWizard = true
	EmailEmailSettingsResponseSingleResultSkipWizardFalse EmailEmailSettingsResponseSingleResultSkipWizard = false
)

func (r EmailEmailSettingsResponseSingleResultSkipWizard) IsKnown() bool {
	switch r {
	case EmailEmailSettingsResponseSingleResultSkipWizardTrue, EmailEmailSettingsResponseSingleResultSkipWizardFalse:
		return true
	}
	return false
}

// Show the state of your account, and the type or configuration error.
type EmailEmailSettingsResponseSingleResultStatus string

const (
	EmailEmailSettingsResponseSingleResultStatusReady               EmailEmailSettingsResponseSingleResultStatus = "ready"
	EmailEmailSettingsResponseSingleResultStatusUnconfigured        EmailEmailSettingsResponseSingleResultStatus = "unconfigured"
	EmailEmailSettingsResponseSingleResultStatusMisconfigured       EmailEmailSettingsResponseSingleResultStatus = "misconfigured"
	EmailEmailSettingsResponseSingleResultStatusMisconfiguredLocked EmailEmailSettingsResponseSingleResultStatus = "misconfigured/locked"
	EmailEmailSettingsResponseSingleResultStatusUnlocked            EmailEmailSettingsResponseSingleResultStatus = "unlocked"
)

func (r EmailEmailSettingsResponseSingleResultStatus) IsKnown() bool {
	switch r {
	case EmailEmailSettingsResponseSingleResultStatusReady, EmailEmailSettingsResponseSingleResultStatusUnconfigured, EmailEmailSettingsResponseSingleResultStatusMisconfigured, EmailEmailSettingsResponseSingleResultStatusMisconfiguredLocked, EmailEmailSettingsResponseSingleResultStatusUnlocked:
		return true
	}
	return false
}

type ZoneEmailRoutingDisableParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneEmailRoutingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneEmailRoutingEnableParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneEmailRoutingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
