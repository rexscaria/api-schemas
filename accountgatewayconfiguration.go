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
//
// Deprecated: deprecated
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
	Enabled bool `json:"enabled,required,nullable"`
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
	Errors   []GatewayAccountConfigError   `json:"errors,required"`
	Messages []GatewayAccountConfigMessage `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayAccountConfigSuccess `json:"success,required"`
	// Account settings
	Result GatewayAccountConfigResult `json:"result"`
	JSON   gatewayAccountConfigJSON   `json:"-"`
}

// gatewayAccountConfigJSON contains the JSON metadata for the struct
// [GatewayAccountConfig]
type gatewayAccountConfigJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
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

type GatewayAccountConfigError struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           GatewayAccountConfigErrorsSource `json:"source"`
	JSON             gatewayAccountConfigErrorJSON    `json:"-"`
}

// gatewayAccountConfigErrorJSON contains the JSON metadata for the struct
// [GatewayAccountConfigError]
type gatewayAccountConfigErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GatewayAccountConfigError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigErrorJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountConfigErrorsSource struct {
	Pointer string                               `json:"pointer"`
	JSON    gatewayAccountConfigErrorsSourceJSON `json:"-"`
}

// gatewayAccountConfigErrorsSourceJSON contains the JSON metadata for the struct
// [GatewayAccountConfigErrorsSource]
type gatewayAccountConfigErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountConfigMessage struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           GatewayAccountConfigMessagesSource `json:"source"`
	JSON             gatewayAccountConfigMessageJSON    `json:"-"`
}

// gatewayAccountConfigMessageJSON contains the JSON metadata for the struct
// [GatewayAccountConfigMessage]
type gatewayAccountConfigMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GatewayAccountConfigMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigMessageJSON) RawJSON() string {
	return r.raw
}

type GatewayAccountConfigMessagesSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    gatewayAccountConfigMessagesSourceJSON `json:"-"`
}

// gatewayAccountConfigMessagesSourceJSON contains the JSON metadata for the struct
// [GatewayAccountConfigMessagesSource]
type gatewayAccountConfigMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayAccountConfigSuccess bool

const (
	GatewayAccountConfigSuccessTrue GatewayAccountConfigSuccess = true
)

func (r GatewayAccountConfigSuccess) IsKnown() bool {
	switch r {
	case GatewayAccountConfigSuccessTrue:
		return true
	}
	return false
}

// Account settings
type GatewayAccountConfigResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Account settings
	Settings  GatewayAccountConfigResultSettings `json:"settings"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      gatewayAccountConfigResultJSON     `json:"-"`
}

// gatewayAccountConfigResultJSON contains the JSON metadata for the struct
// [GatewayAccountConfigResult]
type gatewayAccountConfigResultJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
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
type GatewayAccountConfigResultSettings struct {
	// Activity log settings.
	ActivityLog GatewayAccountConfigResultSettingsActivityLog `json:"activity_log,nullable"`
	// Anti-virus settings.
	Antivirus GatewayAccountConfigResultSettingsAntivirus `json:"antivirus,nullable"`
	// Block page layout settings.
	BlockPage GatewayAccountConfigResultSettingsBlockPage `json:"block_page,nullable"`
	// DLP body scanning settings.
	BodyScanning GatewayAccountConfigResultSettingsBodyScanning `json:"body_scanning,nullable"`
	// Browser isolation settings.
	BrowserIsolation GatewayAccountConfigResultSettingsBrowserIsolation `json:"browser_isolation,nullable"`
	// Certificate settings for Gateway TLS interception. If not specified, the
	// Cloudflare Root CA will be used.
	Certificate GatewayAccountConfigResultSettingsCertificate `json:"certificate,nullable"`
	// Custom certificate settings for BYO-PKI. (deprecated and replaced by
	// `certificate`)
	//
	// Deprecated: deprecated
	CustomCertificate CustomCertificateSettings `json:"custom_certificate,nullable"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayAccountConfigResultSettingsExtendedEmailMatching `json:"extended_email_matching,nullable"`
	// FIPS settings.
	Fips GatewayAccountConfigResultSettingsFips `json:"fips,nullable"`
	// Setting to enable host selector in egress policies.
	HostSelector GatewayAccountConfigResultSettingsHostSelector `json:"host_selector,nullable"`
	// Setting to define inspection settings
	Inspection GatewayAccountConfigResultSettingsInspection `json:"inspection,nullable"`
	// Protocol Detection settings.
	ProtocolDetection GatewayAccountConfigResultSettingsProtocolDetection `json:"protocol_detection,nullable"`
	// Sandbox settings.
	Sandbox GatewayAccountConfigResultSettingsSandbox `json:"sandbox,nullable"`
	// TLS interception settings.
	TlsDecrypt GatewayAccountConfigResultSettingsTlsDecrypt `json:"tls_decrypt,nullable"`
	JSON       gatewayAccountConfigResultSettingsJSON       `json:"-"`
}

// gatewayAccountConfigResultSettingsJSON contains the JSON metadata for the struct
// [GatewayAccountConfigResultSettings]
type gatewayAccountConfigResultSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	Certificate           apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	HostSelector          apijson.Field
	Inspection            apijson.Field
	ProtocolDetection     apijson.Field
	Sandbox               apijson.Field
	TlsDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type GatewayAccountConfigResultSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                              `json:"enabled,nullable"`
	JSON    gatewayAccountConfigResultSettingsActivityLogJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsActivityLogJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsActivityLog]
type gatewayAccountConfigResultSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type GatewayAccountConfigResultSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase,nullable"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase,nullable"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed,nullable"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayAccountConfigResultSettingsAntivirusNotificationSettings `json:"notification_settings,nullable"`
	JSON                 gatewayAccountConfigResultSettingsAntivirusJSON                 `json:"-"`
}

// gatewayAccountConfigResultSettingsAntivirusJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsAntivirus]
type gatewayAccountConfigResultSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayAccountConfigResultSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                              `json:"support_url"`
	JSON       gatewayAccountConfigResultSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsAntivirusNotificationSettingsJSON contains the
// JSON metadata for the struct
// [GatewayAccountConfigResultSettingsAntivirusNotificationSettings]
type gatewayAccountConfigResultSettingsAntivirusNotificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type GatewayAccountConfigResultSettingsBlockPage struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled,required,nullable"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode GatewayAccountConfigResultSettingsBlockPageMode `json:"mode,required"`
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
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
	// If mode is customized_block_page: block page title.
	Name string `json:"name"`
	// This setting was shared via the Orgs API and cannot be edited by the current
	// account
	ReadOnly bool `json:"read_only,nullable"`
	// Account tag of account that shared this setting
	SourceAccount string `json:"source_account,nullable"`
	// If mode is customized_block_page: suppress detailed info at the bottom of the
	// block page.
	SuppressFooter bool `json:"suppress_footer"`
	// If mode is redirect_uri: URI to which the user should be redirected.
	TargetUri string `json:"target_uri" format:"uri"`
	// Version number of the setting
	Version int64                                           `json:"version,nullable"`
	JSON    gatewayAccountConfigResultSettingsBlockPageJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsBlockPageJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsBlockPage]
type gatewayAccountConfigResultSettingsBlockPageJSON struct {
	Enabled         apijson.Field
	Mode            apijson.Field
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	IncludeContext  apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Name            apijson.Field
	ReadOnly        apijson.Field
	SourceAccount   apijson.Field
	SuppressFooter  apijson.Field
	TargetUri       apijson.Field
	Version         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// Controls whether the user is redirected to a Cloudflare-hosted block page or to
// a customer-provided URI.
type GatewayAccountConfigResultSettingsBlockPageMode string

const (
	GatewayAccountConfigResultSettingsBlockPageModeCustomizedBlockPage GatewayAccountConfigResultSettingsBlockPageMode = "customized_block_page"
	GatewayAccountConfigResultSettingsBlockPageModeRedirectUri         GatewayAccountConfigResultSettingsBlockPageMode = "redirect_uri"
)

func (r GatewayAccountConfigResultSettingsBlockPageMode) IsKnown() bool {
	switch r {
	case GatewayAccountConfigResultSettingsBlockPageModeCustomizedBlockPage, GatewayAccountConfigResultSettingsBlockPageModeRedirectUri:
		return true
	}
	return false
}

// DLP body scanning settings.
type GatewayAccountConfigResultSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode GatewayAccountConfigResultSettingsBodyScanningInspectionMode `json:"inspection_mode"`
	JSON           gatewayAccountConfigResultSettingsBodyScanningJSON           `json:"-"`
}

// gatewayAccountConfigResultSettingsBodyScanningJSON contains the JSON metadata
// for the struct [GatewayAccountConfigResultSettingsBodyScanning]
type gatewayAccountConfigResultSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Set the inspection mode to either `deep` or `shallow`.
type GatewayAccountConfigResultSettingsBodyScanningInspectionMode string

const (
	GatewayAccountConfigResultSettingsBodyScanningInspectionModeDeep    GatewayAccountConfigResultSettingsBodyScanningInspectionMode = "deep"
	GatewayAccountConfigResultSettingsBodyScanningInspectionModeShallow GatewayAccountConfigResultSettingsBodyScanningInspectionMode = "shallow"
)

func (r GatewayAccountConfigResultSettingsBodyScanningInspectionMode) IsKnown() bool {
	switch r {
	case GatewayAccountConfigResultSettingsBodyScanningInspectionModeDeep, GatewayAccountConfigResultSettingsBodyScanningInspectionModeShallow:
		return true
	}
	return false
}

// Browser isolation settings.
type GatewayAccountConfigResultSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                   `json:"url_browser_isolation_enabled"`
	JSON                       gatewayAccountConfigResultSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsBrowserIsolationJSON contains the JSON
// metadata for the struct [GatewayAccountConfigResultSettingsBrowserIsolation]
type gatewayAccountConfigResultSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Certificate settings for Gateway TLS interception. If not specified, the
// Cloudflare Root CA will be used.
type GatewayAccountConfigResultSettingsCertificate struct {
	// UUID of certificate to be used for interception. Certificate must be available
	// (previously called 'active') on the edge. A nil UUID will indicate the
	// Cloudflare Root CA should be used.
	ID   string                                            `json:"id,required"`
	JSON gatewayAccountConfigResultSettingsCertificateJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsCertificateJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsCertificate]
type gatewayAccountConfigResultSettingsCertificateJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type GatewayAccountConfigResultSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool `json:"enabled,nullable"`
	// This setting was shared via the Orgs API and cannot be edited by the current
	// account
	ReadOnly bool `json:"read_only,nullable"`
	// Account tag of account that shared this setting
	SourceAccount string `json:"source_account,nullable"`
	// Version number of the setting
	Version int64                                                       `json:"version,nullable"`
	JSON    gatewayAccountConfigResultSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsExtendedEmailMatchingJSON contains the JSON
// metadata for the struct
// [GatewayAccountConfigResultSettingsExtendedEmailMatching]
type gatewayAccountConfigResultSettingsExtendedEmailMatchingJSON struct {
	Enabled       apijson.Field
	ReadOnly      apijson.Field
	SourceAccount apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type GatewayAccountConfigResultSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                       `json:"tls"`
	JSON gatewayAccountConfigResultSettingsFipsJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsFipsJSON contains the JSON metadata for the
// struct [GatewayAccountConfigResultSettingsFips]
type gatewayAccountConfigResultSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Setting to enable host selector in egress policies.
type GatewayAccountConfigResultSettingsHostSelector struct {
	// Enable filtering via hosts for egress policies.
	Enabled bool                                               `json:"enabled,nullable"`
	JSON    gatewayAccountConfigResultSettingsHostSelectorJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsHostSelectorJSON contains the JSON metadata
// for the struct [GatewayAccountConfigResultSettingsHostSelector]
type gatewayAccountConfigResultSettingsHostSelectorJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsHostSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsHostSelectorJSON) RawJSON() string {
	return r.raw
}

// Setting to define inspection settings
type GatewayAccountConfigResultSettingsInspection struct {
	// Defines the mode of inspection the proxy will use.
	//
	//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
	//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
	//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
	//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
	Mode GatewayAccountConfigResultSettingsInspectionMode `json:"mode"`
	JSON gatewayAccountConfigResultSettingsInspectionJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsInspectionJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsInspection]
type gatewayAccountConfigResultSettingsInspectionJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsInspection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsInspectionJSON) RawJSON() string {
	return r.raw
}

// Defines the mode of inspection the proxy will use.
//
//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
type GatewayAccountConfigResultSettingsInspectionMode string

const (
	GatewayAccountConfigResultSettingsInspectionModeStatic  GatewayAccountConfigResultSettingsInspectionMode = "static"
	GatewayAccountConfigResultSettingsInspectionModeDynamic GatewayAccountConfigResultSettingsInspectionMode = "dynamic"
)

func (r GatewayAccountConfigResultSettingsInspectionMode) IsKnown() bool {
	switch r {
	case GatewayAccountConfigResultSettingsInspectionModeStatic, GatewayAccountConfigResultSettingsInspectionModeDynamic:
		return true
	}
	return false
}

// Protocol Detection settings.
type GatewayAccountConfigResultSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                    `json:"enabled,nullable"`
	JSON    gatewayAccountConfigResultSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsProtocolDetectionJSON contains the JSON
// metadata for the struct [GatewayAccountConfigResultSettingsProtocolDetection]
type gatewayAccountConfigResultSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// Sandbox settings.
type GatewayAccountConfigResultSettingsSandbox struct {
	// Enable sandbox.
	Enabled bool `json:"enabled,nullable"`
	// Action to take when the file cannot be scanned.
	FallbackAction GatewayAccountConfigResultSettingsSandboxFallbackAction `json:"fallback_action"`
	JSON           gatewayAccountConfigResultSettingsSandboxJSON           `json:"-"`
}

// gatewayAccountConfigResultSettingsSandboxJSON contains the JSON metadata for the
// struct [GatewayAccountConfigResultSettingsSandbox]
type gatewayAccountConfigResultSettingsSandboxJSON struct {
	Enabled        apijson.Field
	FallbackAction apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsSandbox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsSandboxJSON) RawJSON() string {
	return r.raw
}

// Action to take when the file cannot be scanned.
type GatewayAccountConfigResultSettingsSandboxFallbackAction string

const (
	GatewayAccountConfigResultSettingsSandboxFallbackActionAllow GatewayAccountConfigResultSettingsSandboxFallbackAction = "allow"
	GatewayAccountConfigResultSettingsSandboxFallbackActionBlock GatewayAccountConfigResultSettingsSandboxFallbackAction = "block"
)

func (r GatewayAccountConfigResultSettingsSandboxFallbackAction) IsKnown() bool {
	switch r {
	case GatewayAccountConfigResultSettingsSandboxFallbackActionAllow, GatewayAccountConfigResultSettingsSandboxFallbackActionBlock:
		return true
	}
	return false
}

// TLS interception settings.
type GatewayAccountConfigResultSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                             `json:"enabled"`
	JSON    gatewayAccountConfigResultSettingsTlsDecryptJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsTlsDecryptJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsTlsDecrypt]
type gatewayAccountConfigResultSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayAccountConfigResultSettingsTlsDecryptJSON) RawJSON() string {
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
	// Setting to enable host selector in egress policies.
	HostSelector param.Field[GatewayAccountSettingsSettingsHostSelectorParam] `json:"host_selector"`
	// Setting to define inspection settings
	Inspection param.Field[GatewayAccountSettingsSettingsInspectionParam] `json:"inspection"`
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
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
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
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Controls whether the user is redirected to a Cloudflare-hosted block page or to
	// a customer-provided URI.
	Mode param.Field[GatewayAccountSettingsSettingsBlockPageMode] `json:"mode,required"`
	// If mode is customized_block_page: block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
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
type GatewayAccountSettingsSettingsBodyScanningParam struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[GatewayAccountSettingsSettingsBodyScanningInspectionMode] `json:"inspection_mode"`
}

func (r GatewayAccountSettingsSettingsBodyScanningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set the inspection mode to either `deep` or `shallow`.
type GatewayAccountSettingsSettingsBodyScanningInspectionMode string

const (
	GatewayAccountSettingsSettingsBodyScanningInspectionModeDeep    GatewayAccountSettingsSettingsBodyScanningInspectionMode = "deep"
	GatewayAccountSettingsSettingsBodyScanningInspectionModeShallow GatewayAccountSettingsSettingsBodyScanningInspectionMode = "shallow"
)

func (r GatewayAccountSettingsSettingsBodyScanningInspectionMode) IsKnown() bool {
	switch r {
	case GatewayAccountSettingsSettingsBodyScanningInspectionModeDeep, GatewayAccountSettingsSettingsBodyScanningInspectionModeShallow:
		return true
	}
	return false
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

// Setting to enable host selector in egress policies.
type GatewayAccountSettingsSettingsHostSelectorParam struct {
	// Enable filtering via hosts for egress policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayAccountSettingsSettingsHostSelectorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Setting to define inspection settings
type GatewayAccountSettingsSettingsInspectionParam struct {
	// Defines the mode of inspection the proxy will use.
	//
	//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
	//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
	//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
	//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
	Mode param.Field[GatewayAccountSettingsSettingsInspectionMode] `json:"mode"`
}

func (r GatewayAccountSettingsSettingsInspectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the mode of inspection the proxy will use.
//
//   - static: Gateway will use static inspection to inspect HTTP on TCP(80). If TLS
//     decryption is on, Gateway will inspect HTTPS traffic on TCP(443) & UDP(443).
//   - dynamic: Gateway will use protocol detection to dynamically inspect HTTP and
//     HTTPS traffic on any port. TLS decryption must be on to inspect HTTPS traffic.
type GatewayAccountSettingsSettingsInspectionMode string

const (
	GatewayAccountSettingsSettingsInspectionModeStatic  GatewayAccountSettingsSettingsInspectionMode = "static"
	GatewayAccountSettingsSettingsInspectionModeDynamic GatewayAccountSettingsSettingsInspectionMode = "dynamic"
)

func (r GatewayAccountSettingsSettingsInspectionMode) IsKnown() bool {
	switch r {
	case GatewayAccountSettingsSettingsInspectionModeStatic, GatewayAccountSettingsSettingsInspectionModeDynamic:
		return true
	}
	return false
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
