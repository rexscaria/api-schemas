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
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCnames bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS AccountSettingsZoneDefaultsInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers AccountSettingsZoneDefaultsNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NsTtl float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa AccountSettingsZoneDefaultsSoa `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode AccountSettingsZoneDefaultsZoneMode `json:"zone_mode"`
	JSON     accountSettingsZoneDefaultsJSON     `json:"-"`
}

// accountSettingsZoneDefaultsJSON contains the JSON metadata for the struct
// [AccountSettingsZoneDefaults]
type accountSettingsZoneDefaultsJSON struct {
	FlattenAllCnames   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NsTtl              apijson.Field
	SecondaryOverrides apijson.Field
	Soa                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountSettingsZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type AccountSettingsZoneDefaultsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                     `json:"reference_zone_id"`
	JSON            accountSettingsZoneDefaultsInternalDNSJSON `json:"-"`
}

// accountSettingsZoneDefaultsInternalDNSJSON contains the JSON metadata for the
// struct [AccountSettingsZoneDefaultsInternalDNS]
type accountSettingsZoneDefaultsInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountSettingsZoneDefaultsInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsZoneDefaultsInternalDNSJSON) RawJSON() string {
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

// Components of the zone's SOA record.
type AccountSettingsZoneDefaultsSoa struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTtl float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	Mname string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	Rname string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	Ttl  float64                            `json:"ttl,required"`
	JSON accountSettingsZoneDefaultsSoaJSON `json:"-"`
}

// accountSettingsZoneDefaultsSoaJSON contains the JSON metadata for the struct
// [AccountSettingsZoneDefaultsSoa]
type accountSettingsZoneDefaultsSoaJSON struct {
	Expire      apijson.Field
	MinTtl      apijson.Field
	Mname       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	Rname       apijson.Field
	Ttl         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsZoneDefaultsSoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingsZoneDefaultsSoaJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type AccountSettingsZoneDefaultsZoneMode string

const (
	AccountSettingsZoneDefaultsZoneModeStandard AccountSettingsZoneDefaultsZoneMode = "standard"
	AccountSettingsZoneDefaultsZoneModeCdnOnly  AccountSettingsZoneDefaultsZoneMode = "cdn_only"
	AccountSettingsZoneDefaultsZoneModeDNSOnly  AccountSettingsZoneDefaultsZoneMode = "dns_only"
)

func (r AccountSettingsZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case AccountSettingsZoneDefaultsZoneModeStandard, AccountSettingsZoneDefaultsZoneModeCdnOnly, AccountSettingsZoneDefaultsZoneModeDNSOnly:
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
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCnames param.Field[bool] `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS param.Field[AccountSettingsZoneDefaultsInternalDNSParam] `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[AccountSettingsZoneDefaultsNameserversParam] `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NsTtl param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa param.Field[AccountSettingsZoneDefaultsSoaParam] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[AccountSettingsZoneDefaultsZoneMode] `json:"zone_mode"`
}

func (r AccountSettingsZoneDefaultsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for this internal zone.
type AccountSettingsZoneDefaultsInternalDNSParam struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID param.Field[string] `json:"reference_zone_id"`
}

func (r AccountSettingsZoneDefaultsInternalDNSParam) MarshalJSON() (data []byte, err error) {
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

// Components of the zone's SOA record.
type AccountSettingsZoneDefaultsSoaParam struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire param.Field[float64] `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTtl param.Field[float64] `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	Mname param.Field[string] `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh param.Field[float64] `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry param.Field[float64] `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	Rname param.Field[string] `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	Ttl param.Field[float64] `json:"ttl,required"`
}

func (r AccountSettingsZoneDefaultsSoaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSResponseSingle struct {
	Errors   []DNSSettingsMessages `json:"errors,required"`
	Messages []DNSSettingsMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSResponseSingleSuccess `json:"success,required"`
	Result  AccountSettings          `json:"result"`
	JSON    dnsResponseSingleJSON    `json:"-"`
}

// dnsResponseSingleJSON contains the JSON metadata for the struct
// [DNSResponseSingle]
type dnsResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

// Whether the API call was successful.
type DNSResponseSingleSuccess bool

const (
	DNSResponseSingleSuccessTrue DNSResponseSingleSuccess = true
)

func (r DNSResponseSingleSuccess) IsKnown() bool {
	switch r {
	case DNSResponseSingleSuccessTrue:
		return true
	}
	return false
}

type DNSSettingsMessages struct {
	Code             int64                     `json:"code,required"`
	Message          string                    `json:"message,required"`
	DocumentationURL string                    `json:"documentation_url"`
	Source           DNSSettingsMessagesSource `json:"source"`
	JSON             dnsSettingsMessagesJSON   `json:"-"`
}

// dnsSettingsMessagesJSON contains the JSON metadata for the struct
// [DNSSettingsMessages]
type dnsSettingsMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSSettingsMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSSettingsMessagesSource struct {
	Pointer string                        `json:"pointer"`
	JSON    dnsSettingsMessagesSourceJSON `json:"-"`
}

// dnsSettingsMessagesSourceJSON contains the JSON metadata for the struct
// [DNSSettingsMessagesSource]
type dnsSettingsMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSSettingsMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountDNSSettingUpdateParams struct {
	AccountSettings AccountSettingsParam `json:"account_settings,required"`
}

func (r AccountDNSSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.AccountSettings)
}
