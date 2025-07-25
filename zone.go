// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneService contains methods and other services that help with interacting with
// the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options                []option.RequestOption
	Subscription           *ZoneSubscriptionService
	Access                 *ZoneAccessService
	Acm                    *ZoneAcmService
	Addressing             *ZoneAddressingService
	Analytics              *ZoneAnalyticsService
	APIGateway             *ZoneAPIGatewayService
	Argo                   *ZoneArgoService
	AvailablePlans         *ZoneAvailablePlanService
	BotManagement          *ZoneBotManagementService
	Cache                  *ZoneCacheService
	CertificateAuthorities *ZoneCertificateAuthorityService
	ClientCertificates     *ZoneClientCertificateService
	CloudConnector         *ZoneCloudConnectorService
	ContentUploadScan      *ZoneContentUploadScanService
	CustomCertificates     *ZoneCustomCertificateService
	CustomHostnames        *ZoneCustomHostnameService
	CustomNs               *ZoneCustomNService
	DcvDelegation          *ZoneDcvDelegationService
	Devices                *ZoneDeviceService
	DNSAnalytics           *ZoneDNSAnalyticService
	DNSRecords             *ZoneDNSRecordService
	DNSSettings            *ZoneDNSSettingService
	Dnssec                 *ZoneDnssecService
	Email                  *ZoneEmailService
	Filters                *ZoneFilterService
	Firewall               *ZoneFirewallService
	Healthchecks           *ZoneHealthcheckService
	Hold                   *ZoneHoldService
	Hostnames              *ZoneHostnameService
	KeylessCertificates    *ZoneKeylessCertificateService
	LeakedCredentialChecks *ZoneLeakedCredentialCheckService
	LoadBalancers          *ZoneLoadBalancerService
	Logpush                *ZoneLogpushService
	Logs                   *ZoneLogService
	ManagedHeaders         *ZoneManagedHeaderService
	OriginTlsClientAuth    *ZoneOriginTlsClientAuthService
	PageShield             *ZonePageShieldService
	Pagerules              *ZonePageruleService
	RateLimits             *ZoneRateLimitService
	Rulesets               *ZoneRulesetService
	SecondaryDNS           *ZoneSecondaryDNSService
	SecurityCenter         *ZoneSecurityCenterService
	Settings               *ZoneSettingService
	Snippets               *ZoneSnippetService
	Spectrum               *ZoneSpectrumService
	SpeedAPI               *ZoneSpeedAPIService
	Ssl                    *ZoneSslService
	URLNormalization       *ZoneURLNormalizationService
	WaitingRooms           *ZoneWaitingRoomService
	Web3                   *ZoneWeb3Service
	Workers                *ZoneWorkerService
	CustomPages            *ZoneCustomPageService
	CachePurge             *ZoneCachePurgeService
	ZoneActivation         *ZoneZoneActivationService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	r.Subscription = NewZoneSubscriptionService(opts...)
	r.Access = NewZoneAccessService(opts...)
	r.Acm = NewZoneAcmService(opts...)
	r.Addressing = NewZoneAddressingService(opts...)
	r.Analytics = NewZoneAnalyticsService(opts...)
	r.APIGateway = NewZoneAPIGatewayService(opts...)
	r.Argo = NewZoneArgoService(opts...)
	r.AvailablePlans = NewZoneAvailablePlanService(opts...)
	r.BotManagement = NewZoneBotManagementService(opts...)
	r.Cache = NewZoneCacheService(opts...)
	r.CertificateAuthorities = NewZoneCertificateAuthorityService(opts...)
	r.ClientCertificates = NewZoneClientCertificateService(opts...)
	r.CloudConnector = NewZoneCloudConnectorService(opts...)
	r.ContentUploadScan = NewZoneContentUploadScanService(opts...)
	r.CustomCertificates = NewZoneCustomCertificateService(opts...)
	r.CustomHostnames = NewZoneCustomHostnameService(opts...)
	r.CustomNs = NewZoneCustomNService(opts...)
	r.DcvDelegation = NewZoneDcvDelegationService(opts...)
	r.Devices = NewZoneDeviceService(opts...)
	r.DNSAnalytics = NewZoneDNSAnalyticService(opts...)
	r.DNSRecords = NewZoneDNSRecordService(opts...)
	r.DNSSettings = NewZoneDNSSettingService(opts...)
	r.Dnssec = NewZoneDnssecService(opts...)
	r.Email = NewZoneEmailService(opts...)
	r.Filters = NewZoneFilterService(opts...)
	r.Firewall = NewZoneFirewallService(opts...)
	r.Healthchecks = NewZoneHealthcheckService(opts...)
	r.Hold = NewZoneHoldService(opts...)
	r.Hostnames = NewZoneHostnameService(opts...)
	r.KeylessCertificates = NewZoneKeylessCertificateService(opts...)
	r.LeakedCredentialChecks = NewZoneLeakedCredentialCheckService(opts...)
	r.LoadBalancers = NewZoneLoadBalancerService(opts...)
	r.Logpush = NewZoneLogpushService(opts...)
	r.Logs = NewZoneLogService(opts...)
	r.ManagedHeaders = NewZoneManagedHeaderService(opts...)
	r.OriginTlsClientAuth = NewZoneOriginTlsClientAuthService(opts...)
	r.PageShield = NewZonePageShieldService(opts...)
	r.Pagerules = NewZonePageruleService(opts...)
	r.RateLimits = NewZoneRateLimitService(opts...)
	r.Rulesets = NewZoneRulesetService(opts...)
	r.SecondaryDNS = NewZoneSecondaryDNSService(opts...)
	r.SecurityCenter = NewZoneSecurityCenterService(opts...)
	r.Settings = NewZoneSettingService(opts...)
	r.Snippets = NewZoneSnippetService(opts...)
	r.Spectrum = NewZoneSpectrumService(opts...)
	r.SpeedAPI = NewZoneSpeedAPIService(opts...)
	r.Ssl = NewZoneSslService(opts...)
	r.URLNormalization = NewZoneURLNormalizationService(opts...)
	r.WaitingRooms = NewZoneWaitingRoomService(opts...)
	r.Web3 = NewZoneWeb3Service(opts...)
	r.Workers = NewZoneWorkerService(opts...)
	r.CustomPages = NewZoneCustomPageService(opts...)
	r.CachePurge = NewZoneCachePurgeService(opts...)
	r.ZoneActivation = NewZoneZoneActivationService(opts...)
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *ZoneNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Edits a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Update(ctx context.Context, zoneID string, body ZoneUpdateParams, opts ...option.RequestOption) (res *ZoneUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, sorts, and filters your zones. Listing zones across more than
// 500 accounts is currently not allowed.
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *ZoneListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an existing zone.
func (r *ZoneService) Delete(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *ZoneService) ListAvailableRatePlans(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneListAvailableRatePlansResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_rate_plans", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// ### Purge All Cached Content
//
// Removes ALL files from Cloudflare's cache. All tiers can purge everything.
//
// ```
// {"purge_everything": true}
// ```
//
// ### Purge Cached Content by URL
//
// Granularly removes one or more files from Cloudflare's cache by specifying URLs.
// All tiers can purge by URL.
//
// To purge files with custom cache keys, include the headers used to compute the
// cache key as in the example. If you have a device type or geo in your cache key,
// you will need to include the CF-Device-Type or CF-IPCountry headers. If you have
// lang in your cache key, you will need to include the Accept-Language header.
//
// **NB:** When including the Origin header, be sure to include the **scheme** and
// **hostname**. The port number can be omitted if it is the default port (80 for
// http, 443 for https), but must be included otherwise.
//
// Single file purge example with files:
//
// ```
// {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}
// ```
//
// Single file purge example with url and header pairs:
//
// ```
// {"files": [{url: "http://www.example.com/cat_picture.jpg", headers: { "CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN" }}, {url: "http://www.example.com/dog_picture.jpg", headers: { "CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US" }}]}
// ```
//
// ### Purge Cached Content by Tag, Host or Prefix
//
// Granularly removes one or more files from Cloudflare's cache either by
// specifying the host, the associated Cache-Tag, or a Prefix.
//
// Flex purge with tags:
//
// ```
// {"tags": ["a-cache-tag", "another-cache-tag"]}
// ```
//
// Flex purge with hosts:
//
// ```
// {"hosts": ["www.example.com", "images.example.com"]}
// ```
//
// Flex purge with prefixes:
//
// ```
// {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}
// ```
//
// ### Availability and limits
//
// please refer to
// [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *ZoneService) PurgeCache(ctx context.Context, zoneID string, body ZonePurgeCacheParams, opts ...option.RequestOption) (res *ZonePurgeCacheResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/purge_cache", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Triggeres a new activation check for a PENDING Zone. This can be triggered every
// 5 min for paygo/ent customers, every hour for FREE Zones.
func (r *ZoneService) RerunActivationCheck(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneRerunActivationCheckResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/activation_check", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type MessagesZonesItem struct {
	Code             int64                   `json:"code,required"`
	Message          string                  `json:"message,required"`
	DocumentationURL string                  `json:"documentation_url"`
	Source           MessagesZonesItemSource `json:"source"`
	JSON             messagesZonesItemJSON   `json:"-"`
}

// messagesZonesItemJSON contains the JSON metadata for the struct
// [MessagesZonesItem]
type messagesZonesItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessagesZonesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesZonesItemJSON) RawJSON() string {
	return r.raw
}

type MessagesZonesItemSource struct {
	Pointer string                      `json:"pointer"`
	JSON    messagesZonesItemSourceJSON `json:"-"`
}

// messagesZonesItemSourceJSON contains the JSON metadata for the struct
// [MessagesZonesItemSource]
type messagesZonesItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesZonesItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesZonesItemSourceJSON) RawJSON() string {
	return r.raw
}

type Zone struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to.
	Account ZoneAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active.
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone.
	Meta ZoneMeta `json:"meta,required"`
	// When the zone was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name.
	Name string `json:"name,required"`
	// The name servers Cloudflare assigns to a zone.
	NameServers []string `json:"name_servers,required" format:"hostname"`
	// DNS host at the time of switching to Cloudflare.
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare.
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare.
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone.
	Owner ZoneOwner `json:"owner,required"`
	// A Zones subscription information.
	//
	// Deprecated: deprecated
	Plan ZonePlan `json:"plan,required"`
	// Allows the customer to use a custom apex. _Tenants Only Configuration_.
	CnameSuffix string `json:"cname_suffix"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused bool `json:"paused"`
	// Legacy permissions based on legacy user membership information.
	//
	// Deprecated: deprecated
	Permissions []string `json:"permissions"`
	// The zone status on Cloudflare.
	Status ZoneStatus `json:"status"`
	// The root organizational unit that this zone belongs to (such as a tenant or
	// organization).
	Tenant ZoneTenant `json:"tenant"`
	// The immediate parent organizational unit that this zone belongs to (such as
	// under a tenant or sub-organization).
	TenantUnit ZoneTenantUnit `json:"tenant_unit"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type ZoneType `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string `json:"vanity_name_servers" format:"hostname"`
	// Verification key for partial zone setup.
	VerificationKey string   `json:"verification_key"`
	JSON            zoneJSON `json:"-"`
}

// zoneJSON contains the JSON metadata for the struct [Zone]
type zoneJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	NameServers         apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	Plan                apijson.Field
	CnameSuffix         apijson.Field
	Paused              apijson.Field
	Permissions         apijson.Field
	Status              apijson.Field
	Tenant              apijson.Field
	TenantUnit          apijson.Field
	Type                apijson.Field
	VanityNameServers   apijson.Field
	VerificationKey     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Zone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneJSON) RawJSON() string {
	return r.raw
}

// The account the zone belongs to.
type ZoneAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account.
	Name string          `json:"name"`
	JSON zoneAccountJSON `json:"-"`
}

// zoneAccountJSON contains the JSON metadata for the struct [ZoneAccount]
type zoneAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccountJSON) RawJSON() string {
	return r.raw
}

// Metadata about the zone.
type ZoneMeta struct {
	// The zone is only configured for CDN.
	CdnOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have.
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS.
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have.
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing.
	PhishingDetected bool         `json:"phishing_detected"`
	Step             int64        `json:"step"`
	JSON             zoneMetaJSON `json:"-"`
}

// zoneMetaJSON contains the JSON metadata for the struct [ZoneMeta]
type zoneMetaJSON struct {
	CdnOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneMetaJSON) RawJSON() string {
	return r.raw
}

// The owner of the zone.
type ZoneOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner.
	Name string `json:"name"`
	// The type of owner.
	Type string        `json:"type"`
	JSON zoneOwnerJSON `json:"-"`
}

// zoneOwnerJSON contains the JSON metadata for the struct [ZoneOwner]
type zoneOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOwnerJSON) RawJSON() string {
	return r.raw
}

// A Zones subscription information.
//
// Deprecated: deprecated
type ZonePlan struct {
	// Identifier
	ID string `json:"id"`
	// States if the subscription can be activated.
	CanSubscribe bool `json:"can_subscribe"`
	// The denomination of the customer.
	Currency string `json:"currency"`
	// If this Zone is managed by another company.
	ExternallyManaged bool `json:"externally_managed"`
	// How often the customer is billed.
	Frequency string `json:"frequency"`
	// States if the subscription active.
	IsSubscribed bool `json:"is_subscribed"`
	// If the legacy discount applies to this Zone.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy name of the plan.
	LegacyID string `json:"legacy_id"`
	// Name of the owner.
	Name string `json:"name"`
	// How much the customer is paying.
	Price float64      `json:"price"`
	JSON  zonePlanJSON `json:"-"`
}

// zonePlanJSON contains the JSON metadata for the struct [ZonePlan]
type zonePlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePlanJSON) RawJSON() string {
	return r.raw
}

// The zone status on Cloudflare.
type ZoneStatus string

const (
	ZoneStatusInitializing ZoneStatus = "initializing"
	ZoneStatusPending      ZoneStatus = "pending"
	ZoneStatusActive       ZoneStatus = "active"
	ZoneStatusMoved        ZoneStatus = "moved"
)

func (r ZoneStatus) IsKnown() bool {
	switch r {
	case ZoneStatusInitializing, ZoneStatusPending, ZoneStatusActive, ZoneStatusMoved:
		return true
	}
	return false
}

// The root organizational unit that this zone belongs to (such as a tenant or
// organization).
type ZoneTenant struct {
	// Identifier
	ID string `json:"id"`
	// The name of the Tenant account.
	Name string         `json:"name"`
	JSON zoneTenantJSON `json:"-"`
}

// zoneTenantJSON contains the JSON metadata for the struct [ZoneTenant]
type zoneTenantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTenantJSON) RawJSON() string {
	return r.raw
}

// The immediate parent organizational unit that this zone belongs to (such as
// under a tenant or sub-organization).
type ZoneTenantUnit struct {
	// Identifier
	ID   string             `json:"id"`
	JSON zoneTenantUnitJSON `json:"-"`
}

// zoneTenantUnitJSON contains the JSON metadata for the struct [ZoneTenantUnit]
type zoneTenantUnitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTenantUnitJSON) RawJSON() string {
	return r.raw
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type ZoneType string

const (
	ZoneTypeFull      ZoneType = "full"
	ZoneTypePartial   ZoneType = "partial"
	ZoneTypeSecondary ZoneType = "secondary"
	ZoneTypeInternal  ZoneType = "internal"
)

func (r ZoneType) IsKnown() bool {
	switch r {
	case ZoneTypeFull, ZoneTypePartial, ZoneTypeSecondary, ZoneTypeInternal:
		return true
	}
	return false
}

type ZoneNewResponse struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                `json:"success,required"`
	Result  Zone                `json:"result"`
	JSON    zoneNewResponseJSON `json:"-"`
}

// zoneNewResponseJSON contains the JSON metadata for the struct [ZoneNewResponse]
type zoneNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneGetResponse struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                `json:"success,required"`
	Result  Zone                `json:"result"`
	JSON    zoneGetResponseJSON `json:"-"`
}

// zoneGetResponseJSON contains the JSON metadata for the struct [ZoneGetResponse]
type zoneGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneUpdateResponse struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                   `json:"success,required"`
	Result  Zone                   `json:"result"`
	JSON    zoneUpdateResponseJSON `json:"-"`
}

// zoneUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneUpdateResponse]
type zoneUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneListResponse struct {
	Errors   []MessagesZonesItem `json:"errors,required"`
	Messages []MessagesZonesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                       `json:"success,required"`
	Result     []Zone                     `json:"result"`
	ResultInfo ZoneListResponseResultInfo `json:"result_info"`
	JSON       zoneListResponseJSON       `json:"-"`
}

// zoneListResponseJSON contains the JSON metadata for the struct
// [ZoneListResponse]
type zoneListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64 `json:"total_count"`
	// Total number of pages
	TotalPages float64                        `json:"total_pages"`
	JSON       zoneListResponseResultInfoJSON `json:"-"`
}

// zoneListResponseResultInfoJSON contains the JSON metadata for the struct
// [ZoneListResponseResultInfo]
type zoneListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponse struct {
	Errors   []ZoneDeleteResponseError   `json:"errors,required"`
	Messages []ZoneDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                     `json:"success,required"`
	Result  ZoneDeleteResponseResult `json:"result,nullable"`
	JSON    zoneDeleteResponseJSON   `json:"-"`
}

// zoneDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDeleteResponse]
type zoneDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponseError struct {
	Code             int64                          `json:"code,required"`
	Message          string                         `json:"message,required"`
	DocumentationURL string                         `json:"documentation_url"`
	Source           ZoneDeleteResponseErrorsSource `json:"source"`
	JSON             zoneDeleteResponseErrorJSON    `json:"-"`
}

// zoneDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseError]
type zoneDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponseErrorsSource struct {
	Pointer string                             `json:"pointer"`
	JSON    zoneDeleteResponseErrorsSourceJSON `json:"-"`
}

// zoneDeleteResponseErrorsSourceJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseErrorsSource]
type zoneDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponseMessage struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           ZoneDeleteResponseMessagesSource `json:"source"`
	JSON             zoneDeleteResponseMessageJSON    `json:"-"`
}

// zoneDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseMessage]
type zoneDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponseMessagesSource struct {
	Pointer string                               `json:"pointer"`
	JSON    zoneDeleteResponseMessagesSourceJSON `json:"-"`
}

// zoneDeleteResponseMessagesSourceJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseMessagesSource]
type zoneDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponseResult struct {
	// Identifier
	ID   string                       `json:"id,required"`
	JSON zoneDeleteResponseResultJSON `json:"-"`
}

// zoneDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseResult]
type zoneDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneListAvailableRatePlansResponse struct {
	Errors   []BillSubsAPIMessages                      `json:"errors,required"`
	Messages []BillSubsAPIMessages                      `json:"messages,required"`
	Result   []ZoneListAvailableRatePlansResponseResult `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZoneListAvailableRatePlansResponseSuccess    `json:"success,required"`
	ResultInfo ZoneListAvailableRatePlansResponseResultInfo `json:"result_info"`
	JSON       zoneListAvailableRatePlansResponseJSON       `json:"-"`
}

// zoneListAvailableRatePlansResponseJSON contains the JSON metadata for the struct
// [ZoneListAvailableRatePlansResponse]
type zoneListAvailableRatePlansResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListAvailableRatePlansResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListAvailableRatePlansResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneListAvailableRatePlansResponseResult struct {
	// Plan identifier tag.
	ID string `json:"id"`
	// Array of available components values for the plan.
	Components []ZoneListAvailableRatePlansResponseResultComponent `json:"components"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The duration of the plan subscription.
	Duration float64 `json:"duration"`
	// The frequency at which you will be billed for this plan.
	Frequency BillSubsAPIFrequency `json:"frequency"`
	// The plan name.
	Name string                                       `json:"name"`
	JSON zoneListAvailableRatePlansResponseResultJSON `json:"-"`
}

// zoneListAvailableRatePlansResponseResultJSON contains the JSON metadata for the
// struct [ZoneListAvailableRatePlansResponseResult]
type zoneListAvailableRatePlansResponseResultJSON struct {
	ID          apijson.Field
	Components  apijson.Field
	Currency    apijson.Field
	Duration    apijson.Field
	Frequency   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListAvailableRatePlansResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListAvailableRatePlansResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneListAvailableRatePlansResponseResultComponent struct {
	// The default amount allocated.
	Default float64 `json:"default"`
	// The unique component.
	Name ZoneListAvailableRatePlansResponseResultComponentsName `json:"name"`
	// The unit price of the addon.
	UnitPrice float64                                               `json:"unit_price"`
	JSON      zoneListAvailableRatePlansResponseResultComponentJSON `json:"-"`
}

// zoneListAvailableRatePlansResponseResultComponentJSON contains the JSON metadata
// for the struct [ZoneListAvailableRatePlansResponseResultComponent]
type zoneListAvailableRatePlansResponseResultComponentJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	UnitPrice   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListAvailableRatePlansResponseResultComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListAvailableRatePlansResponseResultComponentJSON) RawJSON() string {
	return r.raw
}

// The unique component.
type ZoneListAvailableRatePlansResponseResultComponentsName string

const (
	ZoneListAvailableRatePlansResponseResultComponentsNameZones                       ZoneListAvailableRatePlansResponseResultComponentsName = "zones"
	ZoneListAvailableRatePlansResponseResultComponentsNamePageRules                   ZoneListAvailableRatePlansResponseResultComponentsName = "page_rules"
	ZoneListAvailableRatePlansResponseResultComponentsNameDedicatedCertificates       ZoneListAvailableRatePlansResponseResultComponentsName = "dedicated_certificates"
	ZoneListAvailableRatePlansResponseResultComponentsNameDedicatedCertificatesCustom ZoneListAvailableRatePlansResponseResultComponentsName = "dedicated_certificates_custom"
)

func (r ZoneListAvailableRatePlansResponseResultComponentsName) IsKnown() bool {
	switch r {
	case ZoneListAvailableRatePlansResponseResultComponentsNameZones, ZoneListAvailableRatePlansResponseResultComponentsNamePageRules, ZoneListAvailableRatePlansResponseResultComponentsNameDedicatedCertificates, ZoneListAvailableRatePlansResponseResultComponentsNameDedicatedCertificatesCustom:
		return true
	}
	return false
}

// Whether the API call was successful
type ZoneListAvailableRatePlansResponseSuccess bool

const (
	ZoneListAvailableRatePlansResponseSuccessTrue ZoneListAvailableRatePlansResponseSuccess = true
)

func (r ZoneListAvailableRatePlansResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneListAvailableRatePlansResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneListAvailableRatePlansResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       zoneListAvailableRatePlansResponseResultInfoJSON `json:"-"`
}

// zoneListAvailableRatePlansResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneListAvailableRatePlansResponseResultInfo]
type zoneListAvailableRatePlansResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneListAvailableRatePlansResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneListAvailableRatePlansResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponse struct {
	Errors   []ZonePurgeCacheResponseError   `json:"errors,required"`
	Messages []ZonePurgeCacheResponseMessage `json:"messages,required"`
	// Indicates the API call's success or failure.
	Success bool                         `json:"success,required"`
	Result  ZonePurgeCacheResponseResult `json:"result,nullable"`
	JSON    zonePurgeCacheResponseJSON   `json:"-"`
}

// zonePurgeCacheResponseJSON contains the JSON metadata for the struct
// [ZonePurgeCacheResponse]
type zonePurgeCacheResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePurgeCacheResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponseError struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           ZonePurgeCacheResponseErrorsSource `json:"source"`
	JSON             zonePurgeCacheResponseErrorJSON    `json:"-"`
}

// zonePurgeCacheResponseErrorJSON contains the JSON metadata for the struct
// [ZonePurgeCacheResponseError]
type zonePurgeCacheResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePurgeCacheResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponseErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    zonePurgeCacheResponseErrorsSourceJSON `json:"-"`
}

// zonePurgeCacheResponseErrorsSourceJSON contains the JSON metadata for the struct
// [ZonePurgeCacheResponseErrorsSource]
type zonePurgeCacheResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePurgeCacheResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponseMessage struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           ZonePurgeCacheResponseMessagesSource `json:"source"`
	JSON             zonePurgeCacheResponseMessageJSON    `json:"-"`
}

// zonePurgeCacheResponseMessageJSON contains the JSON metadata for the struct
// [ZonePurgeCacheResponseMessage]
type zonePurgeCacheResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePurgeCacheResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseMessageJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponseMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    zonePurgeCacheResponseMessagesSourceJSON `json:"-"`
}

// zonePurgeCacheResponseMessagesSourceJSON contains the JSON metadata for the
// struct [ZonePurgeCacheResponseMessagesSource]
type zonePurgeCacheResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePurgeCacheResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type ZonePurgeCacheResponseResult struct {
	ID   string                           `json:"id,required"`
	JSON zonePurgeCacheResponseResultJSON `json:"-"`
}

// zonePurgeCacheResponseResultJSON contains the JSON metadata for the struct
// [ZonePurgeCacheResponseResult]
type zonePurgeCacheResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePurgeCacheResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePurgeCacheResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneRerunActivationCheckResponse struct {
	Errors   []MessageItem `json:"errors,required"`
	Messages []MessageItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneRerunActivationCheckResponseSuccess `json:"success,required"`
	Result  ZoneRerunActivationCheckResponseResult  `json:"result"`
	JSON    zoneRerunActivationCheckResponseJSON    `json:"-"`
}

// zoneRerunActivationCheckResponseJSON contains the JSON metadata for the struct
// [ZoneRerunActivationCheckResponse]
type zoneRerunActivationCheckResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRerunActivationCheckResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRerunActivationCheckResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneRerunActivationCheckResponseSuccess bool

const (
	ZoneRerunActivationCheckResponseSuccessTrue ZoneRerunActivationCheckResponseSuccess = true
)

func (r ZoneRerunActivationCheckResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneRerunActivationCheckResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneRerunActivationCheckResponseResult struct {
	// Identifier.
	ID   string                                     `json:"id"`
	JSON zoneRerunActivationCheckResponseResultJSON `json:"-"`
}

// zoneRerunActivationCheckResponseResultJSON contains the JSON metadata for the
// struct [ZoneRerunActivationCheckResponseResult]
type zoneRerunActivationCheckResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneRerunActivationCheckResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneRerunActivationCheckResponseResultJSON) RawJSON() string {
	return r.raw
}

type ZoneNewParams struct {
	Account param.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name.
	Name param.Field[string] `json:"name,required"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type param.Field[ZoneType] `json:"type"`
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneNewParamsAccount struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneNewParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneUpdateParams struct {
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused param.Field[bool] `json:"paused"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup. This parameter is only
	// available to Enterprise customers or if it has been explicitly enabled on a
	// zone.
	Type param.Field[ZoneUpdateParamsType] `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers param.Field[[]string] `json:"vanity_name_servers" format:"hostname"`
}

func (r ZoneUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup. This parameter is only
// available to Enterprise customers or if it has been explicitly enabled on a
// zone.
type ZoneUpdateParamsType string

const (
	ZoneUpdateParamsTypeFull      ZoneUpdateParamsType = "full"
	ZoneUpdateParamsTypePartial   ZoneUpdateParamsType = "partial"
	ZoneUpdateParamsTypeSecondary ZoneUpdateParamsType = "secondary"
	ZoneUpdateParamsTypeInternal  ZoneUpdateParamsType = "internal"
)

func (r ZoneUpdateParamsType) IsKnown() bool {
	switch r {
	case ZoneUpdateParamsTypeFull, ZoneUpdateParamsTypePartial, ZoneUpdateParamsTypeSecondary, ZoneUpdateParamsTypeInternal:
		return true
	}
	return false
}

type ZoneListParams struct {
	Account param.Field[ZoneListParamsAccount] `query:"account"`
	// Direction to order zones.
	Direction param.Field[ZoneListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneListParamsMatch] `query:"match"`
	// A domain name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
	// Field to order zones by.
	Order param.Field[ZoneListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Specify a zone status to filter by.
	Status param.Field[ZoneListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneListParamsAccount struct {
	// Filter by an account ID.
	ID param.Field[string] `query:"id"`
	// An account Name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [ZoneListParamsAccount]'s query parameters as `url.Values`.
func (r ZoneListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order zones.
type ZoneListParamsDirection string

const (
	ZoneListParamsDirectionAsc  ZoneListParamsDirection = "asc"
	ZoneListParamsDirectionDesc ZoneListParamsDirection = "desc"
)

func (r ZoneListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneListParamsDirectionAsc, ZoneListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any).
type ZoneListParamsMatch string

const (
	ZoneListParamsMatchAny ZoneListParamsMatch = "any"
	ZoneListParamsMatchAll ZoneListParamsMatch = "all"
)

func (r ZoneListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneListParamsMatchAny, ZoneListParamsMatchAll:
		return true
	}
	return false
}

// Field to order zones by.
type ZoneListParamsOrder string

const (
	ZoneListParamsOrderName        ZoneListParamsOrder = "name"
	ZoneListParamsOrderStatus      ZoneListParamsOrder = "status"
	ZoneListParamsOrderAccountID   ZoneListParamsOrder = "account.id"
	ZoneListParamsOrderAccountName ZoneListParamsOrder = "account.name"
	ZoneListParamsOrderPlanID      ZoneListParamsOrder = "plan.id"
)

func (r ZoneListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneListParamsOrderName, ZoneListParamsOrderStatus, ZoneListParamsOrderAccountID, ZoneListParamsOrderAccountName, ZoneListParamsOrderPlanID:
		return true
	}
	return false
}

// Specify a zone status to filter by.
type ZoneListParamsStatus string

const (
	ZoneListParamsStatusInitializing ZoneListParamsStatus = "initializing"
	ZoneListParamsStatusPending      ZoneListParamsStatus = "pending"
	ZoneListParamsStatusActive       ZoneListParamsStatus = "active"
	ZoneListParamsStatusMoved        ZoneListParamsStatus = "moved"
)

func (r ZoneListParamsStatus) IsKnown() bool {
	switch r {
	case ZoneListParamsStatusInitializing, ZoneListParamsStatusPending, ZoneListParamsStatusActive, ZoneListParamsStatusMoved:
		return true
	}
	return false
}

type ZonePurgeCacheParams struct {
	Body ZonePurgeCacheParamsBodyUnion `json:"body,required"`
}

func (r ZonePurgeCacheParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZonePurgeCacheParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r ZonePurgeCacheParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBody) implementsZonePurgeCacheParamsBodyUnion() {}

// Satisfied by [ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByTags],
// [ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByHostnames],
// [ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByPrefixes],
// [ZonePurgeCacheParamsBodyCachePurgeEverything],
// [ZonePurgeCacheParamsBodyCachePurgeSingleFile],
// [ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [ZonePurgeCacheParamsBody].
type ZonePurgeCacheParamsBodyUnion interface {
	implementsZonePurgeCacheParamsBodyUnion()
}

type ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByTags struct {
	// For more information on cache tags and purging by tags, please refer to
	// [purge by cache-tags documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByTags) implementsZonePurgeCacheParamsBodyUnion() {
}

type ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// For more information purging by hostnames, please refer to
	// [purge by hostname documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByHostnames) implementsZonePurgeCacheParamsBodyUnion() {
}

type ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// For more information on purging by prefixes, please refer to
	// [purge by prefix documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeFlexPurgeByPrefixes) implementsZonePurgeCacheParamsBodyUnion() {
}

type ZonePurgeCacheParamsBodyCachePurgeEverything struct {
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeEverything) implementsZonePurgeCacheParamsBodyUnion() {}

type ZonePurgeCacheParamsBodyCachePurgeSingleFile struct {
	// For more information on purging files, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeSingleFile) implementsZonePurgeCacheParamsBodyUnion() {}

type ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// For more information on purging files with URL and headers, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsZonePurgeCacheParamsBodyUnion() {
}

type ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	Headers param.Field[map[string]string] `json:"headers"`
	URL     param.Field[string]            `json:"url"`
}

func (r ZonePurgeCacheParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
