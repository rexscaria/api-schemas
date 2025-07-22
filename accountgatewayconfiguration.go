// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountGatewayConfigurationService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayConfigurationService] method instead.
type AccountGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountGatewayConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayConfigurationService(opts ...option.RequestOption) (r *AccountGatewayConfigurationService) {
	r = &AccountGatewayConfigurationService{}
	r.Options = opts
	return
}

// Fetches the current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) Update(ctx context.Context, accountID string, body AccountGatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `certificate`, without updating the entire configuration object. Returns an
// error if any collection of settings is not properly configured.
func (r *AccountGatewayConfigurationService) Patch(ctx context.Context, accountID string, body AccountGatewayConfigurationPatchParams, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetches the current Zero Trust certificate configuration.
func (r *AccountGatewayConfigurationService) GetCustomCertificate(ctx context.Context, accountID string, opts ...option.RequestOption) (res *CustomCertificateSettings, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration/custom_certificate", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Custom certificate settings for BYO-PKI. (deprecated and replaced by
// `certificate`)
//
// Deprecated: deprecated
type CustomCertificateSettings struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                        `json:"binding_status"`
	UpdatedAt     time.Time                     `json:"updated_at" format:"date-time"`
	JSON          customCertificateSettingsJSON `json:"-"`
}

// customCertificateSettingsJSON contains the JSON metadata for the struct
// [CustomCertificateSettings]
type customCertificateSettingsJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomCertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateSettingsJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI. (deprecated and replaced by
// `certificate`)
//
// Deprecated: deprecated
type CustomCertificateSettingsParam struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r CustomCertificateSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayAccountConfig struct {
	// Account settings
	Result GatewayAccountConfigResult `json:"result"`
	JSON   gatewayAccountConfigJSON   `json:"-"`
	APIResponseSingleZeroTrustGateway
}

// gatewayAccountConfigJSON contains the JSON metadata for the struct
// [GatewayAccountConfig]
type gatewayAccountConfigJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayAccountConfigResult struct {
	CreatedAt time.Time                      `json:"created_at" format:"date-time"`
	UpdatedAt time.Time                      `json:"updated_at" format:"date-time"`
	JSON      gatewayAccountConfigResultJSON `json:"-"`
	GatewayAccountSettings
}

// gatewayAccountConfigResultJSON contains the JSON metadata for the struct
// [GatewayAccountConfigResult]
type gatewayAccountConfigResultJSON struct {
	CreatedAt   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayAccountSettings struct {
	// Account settings
	Settings GatewayAccountSettingsSettings `json:"settings"`
	JSON     gatewayAccountSettingsJSON     `json:"-"`
}

// gatewayAccountSettingsJSON contains the JSON metadata for the struct
// [GatewayAccountSettings]
type gatewayAccountSettingsJSON struct {
	Settings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayAccountSettingsSettings struct {
	// Activity log settings.
	ActivityLog GatewayAccountSettingsSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayAccountSettingsSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayAccountSettingsSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayAccountSettingsSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayAccountSettingsSettingsBrowserIsolation `json:"browser_isolation"`
	// Certificate settings for Gateway TLS interception. If not specified, the
	// Cloudflare Root CA will be used.
	Certificate GatewayAccountSettingsSettingsCertificate `json:"certificate"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by
	// `certificate`)
	//
	// Deprecated: deprecated
	CustomCertificate CustomCertificateSettings `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayAccountSettingsSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayAccountSettingsSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayAccountSettingsSettingsProtocolDetection `json:"protocol_detection"`
	// Sandbox settings.
	Sandbox GatewayAccountSettingsSettingsSandbox `json:"sandbox"`
	// TLS interception settings.
	TlsDecrypt GatewayAccountSettingsSettingsTlsDecrypt `json:"tls_decrypt"`
	JSON       gatewayAccountSettingsSettingsJSON       `json:"-"`
}

// gatewayAccountSettingsSettingsJSON contains the JSON metadata for the struct
// [GatewayAccountSettingsSettings]
type gatewayAccountSettingsSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	Certificate           apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	ProtocolDetection     apijson.Field
	Sandbox               apijson.Field
	TlsDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type GatewayAccountSettingsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                          `json:"enabled"`
	JSON    gatewayAccountSettingsSettingsActivityLogJSON `json:"-"`
}

// gatewayAccountSettingsSettingsActivityLogJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsActivityLog]
type gatewayAccountSettingsSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type GatewayAccountSettingsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayAccountSettingsSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayAccountSettingsSettingsAntivirusJSON                 `json:"-"`
}

// gatewayAccountSettingsSettingsAntivirusJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsAntivirus]
type gatewayAccountSettingsSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayAccountSettingsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                          `json:"support_url"`
	JSON       gatewayAccountSettingsSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayAccountSettingsSettingsAntivirusNotificationSettingsJSON contains the
// JSON metadata for the struct
// [GatewayAccountSettingsSettingsAntivirusNotificationSettings]
type gatewayAccountSettingsSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type GatewayAccountSettingsSettingsBlockPage struct {
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled"`
	// If mode is customized_block_page: block page footer text.
	FooterText string `json:"footer_text"`
	// If mode is customized_block_page: block page header text.
	HeaderText string `json:"header_text"`
	// If mode is redirect_uri: when enabled, context will be appended to target_uri as
	// query parameters.
	IncludeContext bool `json:"include_context"`
	// If mode is customized_block_page: full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// If mode is customized_block_page: admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// If mode is customized_block_page: subject line for emails created from block
	// page.
	MailtoSubject string `json:"mailto_subject"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode GatewayAccountSettingsSettingsBlockPageMode `json:"mode"`
	// If mode is customized_block_page: block page title.
	Name string `json:"name"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the
	// block page.
	SuppressFooter bool `json:"suppress_footer"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	TargetUri string                                      `json:"target_uri" format:"uri"`
	JSON      gatewayAccountSettingsSettingsBlockPageJSON `json:"-"`
}

// gatewayAccountSettingsSettingsBlockPageJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsBlockPage]
type gatewayAccountSettingsSettingsBlockPageJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	IncludeContext  apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Mode            apijson.Field
	Name            apijson.Field
	SuppressFooter  apijson.Field
	TargetUri       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// Controls whether the user is redirected to a Cloudflare-hosted block page or to
// a customer-provided URI.
type GatewayAccountSettingsSettingsBlockPageMode string

const (
	GatewayAccountSettingsSettingsBlockPageModeCustomizedBlockPage GatewayAccountSettingsSettingsBlockPageMode = "customized_block_page"
	GatewayAccountSettingsSettingsBlockPageModeRedirectUri         GatewayAccountSettingsSettingsBlockPageMode = "redirect_uri"
)

func (r GatewayAccountSettingsSettingsBlockPageMode) IsKnown() bool {
	switch r {
	case GatewayAccountSettingsSettingsBlockPageModeCustomizedBlockPage, GatewayAccountSettingsSettingsBlockPageModeRedirectUri:
		return true
	}
	return false
}

// DLP body scanning settings.
type GatewayAccountSettingsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                         `json:"inspection_mode"`
	JSON           gatewayAccountSettingsSettingsBodyScanningJSON `json:"-"`
}

// gatewayAccountSettingsSettingsBodyScanningJSON contains the JSON metadata for
// the struct [GatewayAccountSettingsSettingsBodyScanning]
type gatewayAccountSettingsSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type GatewayAccountSettingsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                               `json:"url_browser_isolation_enabled"`
	JSON                       gatewayAccountSettingsSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayAccountSettingsSettingsBrowserIsolationJSON contains the JSON metadata
// for the struct [GatewayAccountSettingsSettingsBrowserIsolation]
type gatewayAccountSettingsSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Certificate settings for Gateway TLS interception. If not specified, the
// Cloudflare Root CA will be used.
type GatewayAccountSettingsSettingsCertificate struct {
	// UUID of certificate to be used for interception. Certificate must be available
	// (previously called 'active') on the edge. A nil UUID will indicate the
	// Cloudflare Root CA should be used.
	ID   string                                        `json:"id,required"`
	JSON gatewayAccountSettingsSettingsCertificateJSON `json:"-"`
}

// gatewayAccountSettingsSettingsCertificateJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsCertificate]
type gatewayAccountSettingsSettingsCertificateJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type GatewayAccountSettingsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                    `json:"enabled"`
	JSON    gatewayAccountSettingsSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayAccountSettingsSettingsExtendedEmailMatchingJSON contains the JSON
// metadata for the struct [GatewayAccountSettingsSettingsExtendedEmailMatching]
type gatewayAccountSettingsSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type GatewayAccountSettingsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                   `json:"tls"`
	JSON gatewayAccountSettingsSettingsFipsJSON `json:"-"`
}

// gatewayAccountSettingsSettingsFipsJSON contains the JSON metadata for the struct
// [GatewayAccountSettingsSettingsFips]
type gatewayAccountSettingsSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type GatewayAccountSettingsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                `json:"enabled"`
	JSON    gatewayAccountSettingsSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayAccountSettingsSettingsProtocolDetectionJSON contains the JSON metadata
// for the struct [GatewayAccountSettingsSettingsProtocolDetection]
type gatewayAccountSettingsSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// Sandbox settings.
type GatewayAccountSettingsSettingsSandbox struct {
	// Enable sandbox.
	Enabled bool `json:"enabled"`
	// Action to take when the file cannot be scanned.
	FallbackAction GatewayAccountSettingsSettingsSandboxFallbackAction `json:"fallback_action"`
	JSON           gatewayAccountSettingsSettingsSandboxJSON           `json:"-"`
}

// gatewayAccountSettingsSettingsSandboxJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsSandbox]
type gatewayAccountSettingsSettingsSandboxJSON struct {
	Enabled        apijson.Field
	FallbackAction apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsSandbox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsSandboxJSON) RawJSON() string {
	return r.raw
}

// Action to take when the file cannot be scanned.
type GatewayAccountSettingsSettingsSandboxFallbackAction string

const (
	GatewayAccountSettingsSettingsSandboxFallbackActionAllow GatewayAccountSettingsSettingsSandboxFallbackAction = "allow"
	GatewayAccountSettingsSettingsSandboxFallbackActionBlock GatewayAccountSettingsSettingsSandboxFallbackAction = "block"
)

func (r GatewayAccountSettingsSettingsSandboxFallbackAction) IsKnown() bool {
	switch r {
	case GatewayAccountSettingsSettingsSandboxFallbackActionAllow, GatewayAccountSettingsSettingsSandboxFallbackActionBlock:
		return true
	}
	return false
}

// TLS interception settings.
type GatewayAccountSettingsSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                         `json:"enabled"`
	JSON    gatewayAccountSettingsSettingsTlsDecryptJSON `json:"-"`
}

// gatewayAccountSettingsSettingsTlsDecryptJSON contains the JSON metadata for the
// struct [GatewayAccountSettingsSettingsTlsDecrypt]
type gatewayAccountSettingsSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountSettingsSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountSettingsSettingsTlsDecryptJSON) RawJSON() string {
	return r.raw
}

// Account settings
type GatewayAccountSettingsParam struct {
	// Account settings
	Settings param.Field[GatewayAccountSettingsSettingsParam] `json:"settings"`
}

func (r GatewayAccountSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type GatewayAccountSettingsSettingsParam struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayAccountSettingsSettingsActivityLogParam] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayAccountSettingsSettingsAntivirusParam] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayAccountSettingsSettingsBlockPageParam] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayAccountSettingsSettingsBodyScanningParam] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayAccountSettingsSettingsBrowserIsolationParam] `json:"browser_isolation"`
	// Certificate settings for Gateway TLS interception. If not specified, the
	// Cloudflare Root CA will be used.
	Certificate param.Field[GatewayAccountSettingsSettingsCertificateParam] `json:"certificate"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by
	// `certificate`)
	//
	// Deprecated: deprecated
	CustomCertificate param.Field[CustomCertificateSettingsParam] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayAccountSettingsSettingsExtendedEmailMatchingParam] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayAccountSettingsSettingsFipsParam] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayAccountSettingsSettingsProtocolDetectionParam] `json:"protocol_detection"`
	// Sandbox settings.
	Sandbox param.Field[GatewayAccountSettingsSettingsSandboxParam] `json:"sandbox"`
	// TLS interception settings.
	TlsDecrypt param.Field[GatewayAccountSettingsSettingsTlsDecryptParam] `json:"tls_decrypt"`
}

func (r GatewayAccountSettingsSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayAccountSettingsSettingsActivityLogParam struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayAccountSettingsSettingsActivityLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayAccountSettingsSettingsAntivirusParam struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayAccountSettingsSettingsAntivirusNotificationSettingsParam] `json:"notification_settings"`
}

func (r GatewayAccountSettingsSettingsAntivirusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayAccountSettingsSettingsAntivirusNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayAccountSettingsSettingsAntivirusNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayAccountSettingsSettingsBlockPageParam struct {
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled"`
	// If mode is customized_block_page: block page footer text.
	FooterText param.Field[string] `json:"footer_text"`
	// If mode is customized_block_page: block page header text.
	HeaderText param.Field[string] `json:"header_text"`
	// If mode is redirect_uri: when enabled, context will be appended to target_uri as
	// query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
	// If mode is customized_block_page: full URL to the logo file.
	LogoPath param.Field[string] `json:"logo_path"`
	// If mode is customized_block_page: admin email for users to contact.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// If mode is customized_block_page: subject line for emails created from block
	// page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode param.Field[GatewayAccountSettingsSettingsBlockPageMode] `json:"mode"`
	// If mode is customized_block_page: block page title.
	Name param.Field[string] `json:"name"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the
	// block page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	TargetUri param.Field[string] `json:"target_uri" format:"uri"`
}

func (r GatewayAccountSettingsSettingsBlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayAccountSettingsSettingsBodyScanningParam struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayAccountSettingsSettingsBodyScanningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayAccountSettingsSettingsBrowserIsolationParam struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayAccountSettingsSettingsBrowserIsolationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate settings for Gateway TLS interception. If not specified, the
// Cloudflare Root CA will be used.
type GatewayAccountSettingsSettingsCertificateParam struct {
	// UUID of certificate to be used for interception. Certificate must be available
	// (previously called 'active') on the edge. A nil UUID will indicate the
	// Cloudflare Root CA should be used.
	ID param.Field[string] `json:"id,required"`
}

func (r GatewayAccountSettingsSettingsCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayAccountSettingsSettingsExtendedEmailMatchingParam struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayAccountSettingsSettingsExtendedEmailMatchingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayAccountSettingsSettingsFipsParam struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls param.Field[bool] `json:"tls"`
}

func (r GatewayAccountSettingsSettingsFipsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayAccountSettingsSettingsProtocolDetectionParam struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayAccountSettingsSettingsProtocolDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sandbox settings.
type GatewayAccountSettingsSettingsSandboxParam struct {
	// Enable sandbox.
	Enabled param.Field[bool] `json:"enabled"`
	// Action to take when the file cannot be scanned.
	FallbackAction param.Field[GatewayAccountSettingsSettingsSandboxFallbackAction] `json:"fallback_action"`
}

func (r GatewayAccountSettingsSettingsSandboxParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayAccountSettingsSettingsTlsDecryptParam struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayAccountSettingsSettingsTlsDecryptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayConfigurationUpdateParams struct {
	// Account settings
	GatewayAccountSettings GatewayAccountSettingsParam `json:"gateway_account_settings,required"`
}

func (r AccountGatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GatewayAccountSettings)
}

type AccountGatewayConfigurationPatchParams struct {
	// Account settings
	GatewayAccountSettings GatewayAccountSettingsParam `json:"gateway_account_settings,required"`
}

func (r AccountGatewayConfigurationPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.GatewayAccountSettings)
}
