// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountDNSSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSSettingService] method instead.
type AccountDNSSettingService struct {
	Options []option.RequestOption
	Views   *AccountDNSSettingViewService
}

// NewAccountDNSSettingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDNSSettingService(opts ...option.RequestOption) (r *AccountDNSSettingService) {
	r = &AccountDNSSettingService{}
	r.Options = opts
	r.Views = NewAccountDNSSettingViewService(opts...)
	return
}

// Show DNS settings for an account
func (r *AccountDNSSettingService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *DNSResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update DNS settings for an account
func (r *AccountDNSSettingService) Update(ctx context.Context, accountID string, body AccountDNSSettingUpdateParams, opts ...option.RequestOption) (res *DNSResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountSettings struct {
	ZoneDefaults AccountSettingsZoneDefaults `json:"zone_defaults"`
	JSON         accountSettingsJSON         `json:"-"`
}

// accountSettingsJSON contains the JSON metadata for the struct [AccountSettings]
type accountSettingsJSON struct {
	ZoneDefaults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsJSON) RawJSON() string {
	return r.raw
}

type AccountSettingsZoneDefaults struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers AccountSettingsZoneDefaultsNameservers `json:"nameservers"`
	JSON        accountSettingsZoneDefaultsJSON        `json:"-"`
	DNSSettings
}

// accountSettingsZoneDefaultsJSON contains the JSON metadata for the struct
// [AccountSettingsZoneDefaults]
type accountSettingsZoneDefaultsJSON struct {
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type AccountSettingsZoneDefaultsNameservers struct {
	// Nameserver type
	Type AccountSettingsZoneDefaultsNameserversType `json:"type,required"`
	JSON accountSettingsZoneDefaultsNameserversJSON `json:"-"`
}

// accountSettingsZoneDefaultsNameserversJSON contains the JSON metadata for the
// struct [AccountSettingsZoneDefaultsNameservers]
type accountSettingsZoneDefaultsNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsZoneDefaultsNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsZoneDefaultsNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type AccountSettingsZoneDefaultsNameserversType string

const (
	AccountSettingsZoneDefaultsNameserversTypeCloudflareStandard       AccountSettingsZoneDefaultsNameserversType = "cloudflare.standard"
	AccountSettingsZoneDefaultsNameserversTypeCloudflareStandardRandom AccountSettingsZoneDefaultsNameserversType = "cloudflare.standard.random"
	AccountSettingsZoneDefaultsNameserversTypeCustomAccount            AccountSettingsZoneDefaultsNameserversType = "custom.account"
	AccountSettingsZoneDefaultsNameserversTypeCustomTenant             AccountSettingsZoneDefaultsNameserversType = "custom.tenant"
)

func (r AccountSettingsZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case AccountSettingsZoneDefaultsNameserversTypeCloudflareStandard, AccountSettingsZoneDefaultsNameserversTypeCloudflareStandardRandom, AccountSettingsZoneDefaultsNameserversTypeCustomAccount, AccountSettingsZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

type AccountSettingsParam struct {
	ZoneDefaults param.Field[AccountSettingsZoneDefaultsParam] `json:"zone_defaults"`
}

func (r AccountSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSettingsZoneDefaultsParam struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[AccountSettingsZoneDefaultsNameserversParam] `json:"nameservers"`
	DNSSettingsParam
}

func (r AccountSettingsZoneDefaultsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type AccountSettingsZoneDefaultsNameserversParam struct {
	// Nameserver type
	Type param.Field[AccountSettingsZoneDefaultsNameserversType] `json:"type,required"`
}

func (r AccountSettingsZoneDefaultsNameserversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CommonResponseDNSSettings struct {
	Errors   []DNSSettingsMessages `json:"errors,required"`
	Messages []DNSSettingsMessages `json:"messages,required"`
	// Whether the API call was successful
	Success CommonResponseDNSSettingsSuccess `json:"success,required"`
	JSON    commonResponseDNSSettingsJSON    `json:"-"`
}

// commonResponseDNSSettingsJSON contains the JSON metadata for the struct
// [CommonResponseDNSSettings]
type commonResponseDNSSettingsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CommonResponseDNSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r commonResponseDNSSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CommonResponseDNSSettingsSuccess bool

const (
	CommonResponseDNSSettingsSuccessTrue CommonResponseDNSSettingsSuccess = true
)

func (r CommonResponseDNSSettingsSuccess) IsKnown() bool {
	switch r {
	case CommonResponseDNSSettingsSuccessTrue:
		return true
	}
	return false
}

type DNSResponseSingle struct {
	Result AccountSettings       `json:"result"`
	JSON   dnsResponseSingleJSON `json:"-"`
	SingleResponseDNSSettings
}

// dnsResponseSingleJSON contains the JSON metadata for the struct
// [DNSResponseSingle]
type dnsResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResponseSingleJSON) RawJSON() string {
	return r.raw
}

type DNSSettingsMessages struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    dnsSettingsMessagesJSON `json:"-"`
}

// dnsSettingsMessagesJSON contains the JSON metadata for the struct
// [DNSSettingsMessages]
type dnsSettingsMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsMessagesJSON) RawJSON() string {
	return r.raw
}

type SingleResponseDNSSettings struct {
	Errors   []DNSSettingsMessages `json:"errors,required"`
	Messages []DNSSettingsMessages `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseDNSSettingsSuccess `json:"success,required"`
	JSON    singleResponseDNSSettingsJSON    `json:"-"`
}

// singleResponseDNSSettingsJSON contains the JSON metadata for the struct
// [SingleResponseDNSSettings]
type singleResponseDNSSettingsJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseDNSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseDNSSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseDNSSettingsSuccess bool

const (
	SingleResponseDNSSettingsSuccessTrue SingleResponseDNSSettingsSuccess = true
)

func (r SingleResponseDNSSettingsSuccess) IsKnown() bool {
	switch r {
	case SingleResponseDNSSettingsSuccessTrue:
		return true
	}
	return false
}

type AccountDNSSettingUpdateParams struct {
	AccountSettings AccountSettingsParam `json:"account_settings,required"`
}

func (r AccountDNSSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountSettings)
}
