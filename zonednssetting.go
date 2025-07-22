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

// ZoneDNSSettingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneDNSSettingService] method instead.
type ZoneDNSSettingService struct {
	Options []option.RequestOption
}

// NewZoneDNSSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneDNSSettingService(opts ...option.RequestOption) (r *ZoneDNSSettingService) {
	r = &ZoneDNSSettingService{}
	r.Options = opts
	return
}

// Update DNS settings for a zone
func (r *ZoneDNSSettingService) Update(ctx context.Context, zoneID string, body ZoneDNSSettingUpdateParams, opts ...option.RequestOption) (res *SingleResponseDNSSettingsZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Show DNS settings for a zone
func (r *ZoneDNSSettingService) Show(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SingleResponseDNSSettingsZone, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_settings", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DNSSettings struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCnames bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS DNSSettingsInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NsTtl float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa DNSSettingsSoa `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode DNSSettingsZoneMode `json:"zone_mode"`
	JSON     dnsSettingsJSON     `json:"-"`
}

// dnsSettingsJSON contains the JSON metadata for the struct [DNSSettings]
type dnsSettingsJSON struct {
	FlattenAllCnames   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	NsTtl              apijson.Field
	SecondaryOverrides apijson.Field
	Soa                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type DNSSettingsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                     `json:"reference_zone_id"`
	JSON            dnsSettingsInternalDNSJSON `json:"-"`
}

// dnsSettingsInternalDNSJSON contains the JSON metadata for the struct
// [DNSSettingsInternalDNS]
type dnsSettingsInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DNSSettingsInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Components of the zone's SOA record.
type DNSSettingsSoa struct {
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
	Ttl  float64            `json:"ttl,required"`
	JSON dnsSettingsSoaJSON `json:"-"`
}

// dnsSettingsSoaJSON contains the JSON metadata for the struct [DNSSettingsSoa]
type dnsSettingsSoaJSON struct {
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

func (r *DNSSettingsSoa) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsSettingsSoaJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type DNSSettingsZoneMode string

const (
	DNSSettingsZoneModeStandard DNSSettingsZoneMode = "standard"
	DNSSettingsZoneModeCdnOnly  DNSSettingsZoneMode = "cdn_only"
	DNSSettingsZoneModeDNSOnly  DNSSettingsZoneMode = "dns_only"
)

func (r DNSSettingsZoneMode) IsKnown() bool {
	switch r {
	case DNSSettingsZoneModeStandard, DNSSettingsZoneModeCdnOnly, DNSSettingsZoneModeDNSOnly:
		return true
	}
	return false
}

type DNSSettingsParam struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCnames param.Field[bool] `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS param.Field[DNSSettingsInternalDNSParam] `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NsTtl param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	Soa param.Field[DNSSettingsSoaParam] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[DNSSettingsZoneMode] `json:"zone_mode"`
}

func (r DNSSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for this internal zone.
type DNSSettingsInternalDNSParam struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID param.Field[string] `json:"reference_zone_id"`
}

func (r DNSSettingsInternalDNSParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Components of the zone's SOA record.
type DNSSettingsSoaParam struct {
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

func (r DNSSettingsSoaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingsZone struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingsZoneNameservers `json:"nameservers"`
	JSON        settingsZoneJSON        `json:"-"`
	DNSSettings
}

// settingsZoneJSON contains the JSON metadata for the struct [SettingsZone]
type settingsZoneJSON struct {
	Nameservers apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsZoneJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingsZoneNameservers struct {
	// Nameserver type
	Type SettingsZoneNameserversType `json:"type,required"`
	// Configured nameserver set to be used for this zone
	NsSet int64                       `json:"ns_set"`
	JSON  settingsZoneNameserversJSON `json:"-"`
}

// settingsZoneNameserversJSON contains the JSON metadata for the struct
// [SettingsZoneNameservers]
type settingsZoneNameserversJSON struct {
	Type        apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingsZoneNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsZoneNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingsZoneNameserversType string

const (
	SettingsZoneNameserversTypeCloudflareStandard SettingsZoneNameserversType = "cloudflare.standard"
	SettingsZoneNameserversTypeCustomAccount      SettingsZoneNameserversType = "custom.account"
	SettingsZoneNameserversTypeCustomTenant       SettingsZoneNameserversType = "custom.tenant"
	SettingsZoneNameserversTypeCustomZone         SettingsZoneNameserversType = "custom.zone"
)

func (r SettingsZoneNameserversType) IsKnown() bool {
	switch r {
	case SettingsZoneNameserversTypeCloudflareStandard, SettingsZoneNameserversTypeCustomAccount, SettingsZoneNameserversTypeCustomTenant, SettingsZoneNameserversTypeCustomZone:
		return true
	}
	return false
}

type SettingsZoneParam struct {
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[SettingsZoneNameserversParam] `json:"nameservers"`
	DNSSettingsParam
}

func (r SettingsZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type SettingsZoneNameserversParam struct {
	// Nameserver type
	Type param.Field[SettingsZoneNameserversType] `json:"type,required"`
	// Configured nameserver set to be used for this zone
	NsSet param.Field[int64] `json:"ns_set"`
}

func (r SettingsZoneNameserversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SingleResponseDNSSettingsZone struct {
	Result SettingsZone                      `json:"result"`
	JSON   singleResponseDNSSettingsZoneJSON `json:"-"`
	SingleResponseDNSSettings
}

// singleResponseDNSSettingsZoneJSON contains the JSON metadata for the struct
// [SingleResponseDNSSettingsZone]
type singleResponseDNSSettingsZoneJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseDNSSettingsZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseDNSSettingsZoneJSON) RawJSON() string {
	return r.raw
}

type ZoneDNSSettingUpdateParams struct {
	SettingsZone SettingsZoneParam `json:"settings_zone,required"`
}

func (r ZoneDNSSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SettingsZone)
}
