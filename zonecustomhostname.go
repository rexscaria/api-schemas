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

// ZoneCustomHostnameService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCustomHostnameService] method instead.
type ZoneCustomHostnameService struct {
	Options         []option.RequestOption
	CertificatePack *ZoneCustomHostnameCertificatePackService
	FallbackOrigin  *ZoneCustomHostnameFallbackOriginService
}

// NewZoneCustomHostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCustomHostnameService(opts ...option.RequestOption) (r *ZoneCustomHostnameService) {
	r = &ZoneCustomHostnameService{}
	r.Options = opts
	r.CertificatePack = NewZoneCustomHostnameCertificatePackService(opts...)
	r.FallbackOrigin = NewZoneCustomHostnameFallbackOriginService(opts...)
	return
}

// Add a new custom hostname and request that an SSL certificate be issued for it.
// One of three validation methods—http, txt, email—should be used, with 'http'
// recommended if the CNAME is already in place (or will be soon). Specifying
// 'email' will send an email to the WHOIS contacts on file for the base domain
// plus hostmaster, postmaster, webmaster, admin, administrator. If http is used
// and the domain is not already pointing to the Managed CNAME host, the PATCH
// method must be used once it is (to complete validation). Enable bundling of
// certificates using the custom_cert_bundle field. The bundling process requires
// the following condition One certificate in the bundle must use an RSA, and the
// other must use an ECDSA.
func (r *ZoneCustomHostnameService) New(ctx context.Context, zoneID string, body ZoneCustomHostnameNewParams, opts ...option.RequestOption) (res *CustomHostnameResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Custom Hostname Details
func (r *ZoneCustomHostnameService) Get(ctx context.Context, zoneID string, customHostnameID string, opts ...option.RequestOption) (res *CustomHostnameResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'. Bundle an existing certificate with another certificate by
// using the "custom_cert_bundle" field. The bundling process supports combining
// certificates as long as the following condition is met. One certificate must use
// the RSA algorithm, and the other must use the ECDSA algorithm.
func (r *ZoneCustomHostnameService) Update(ctx context.Context, zoneID string, customHostnameID string, body ZoneCustomHostnameUpdateParams, opts ...option.RequestOption) (res *CustomHostnameResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *ZoneCustomHostnameService) List(ctx context.Context, zoneID string, query ZoneCustomHostnameListParams, opts ...option.RequestOption) (res *ZoneCustomHostnameListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete Custom Hostname (and any issued SSL certificates)
func (r *ZoneCustomHostnameService) Delete(ctx context.Context, zoneID string, customHostnameID string, body ZoneCustomHostnameDeleteParams, opts ...option.RequestOption) (res *ZoneCustomHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if customHostnameID == "" {
		err = errors.New("missing required custom_hostname_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// The Certificate Authority that will issue the certificate
type CertificateAuthorityCustomHostname string

const (
	CertificateAuthorityCustomHostnameDigicert    CertificateAuthorityCustomHostname = "digicert"
	CertificateAuthorityCustomHostnameGoogle      CertificateAuthorityCustomHostname = "google"
	CertificateAuthorityCustomHostnameLetsEncrypt CertificateAuthorityCustomHostname = "lets_encrypt"
	CertificateAuthorityCustomHostnameSslCom      CertificateAuthorityCustomHostname = "ssl_com"
)

func (r CertificateAuthorityCustomHostname) IsKnown() bool {
	switch r {
	case CertificateAuthorityCustomHostnameDigicert, CertificateAuthorityCustomHostnameGoogle, CertificateAuthorityCustomHostnameLetsEncrypt, CertificateAuthorityCustomHostnameSslCom:
		return true
	}
	return false
}

type CustomHostname struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string            `json:"hostname,required"`
	Ssl      CustomHostnameSsl `json:"ssl,required"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata map[string]string `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSni string `json:"custom_origin_sni"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// Status of the hostname's activation.
	Status CustomHostnameStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []string           `json:"verification_errors"`
	JSON               customHostnameJSON `json:"-"`
}

// customHostnameJSON contains the JSON metadata for the struct [CustomHostname]
type customHostnameJSON struct {
	ID                        apijson.Field
	Hostname                  apijson.Field
	Ssl                       apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSni           apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CertificateAuthorityCustomHostname `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []string `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method CustomHostnameSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string                    `json:"serial_number"`
	Settings     CustomHostnameSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameSslValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                  `json:"wildcard"`
	JSON     customHostnameSslJSON `json:"-"`
}

// customHostnameSslJSON contains the JSON metadata for the struct
// [CustomHostnameSsl]
type customHostnameSslJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CustomHostnameSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameSslJSON) RawJSON() string {
	return r.raw
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameSslBundleMethod string

const (
	CustomHostnameSslBundleMethodUbiquitous CustomHostnameSslBundleMethod = "ubiquitous"
	CustomHostnameSslBundleMethodOptimal    CustomHostnameSslBundleMethod = "optimal"
	CustomHostnameSslBundleMethodForce      CustomHostnameSslBundleMethod = "force"
)

func (r CustomHostnameSslBundleMethod) IsKnown() bool {
	switch r {
	case CustomHostnameSslBundleMethodUbiquitous, CustomHostnameSslBundleMethodOptimal, CustomHostnameSslBundleMethodForce:
		return true
	}
	return false
}

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameSslMethod string

const (
	CustomHostnameSslMethodHTTP  CustomHostnameSslMethod = "http"
	CustomHostnameSslMethodTxt   CustomHostnameSslMethod = "txt"
	CustomHostnameSslMethodEmail CustomHostnameSslMethod = "email"
)

func (r CustomHostnameSslMethod) IsKnown() bool {
	switch r {
	case CustomHostnameSslMethodHTTP, CustomHostnameSslMethodTxt, CustomHostnameSslMethodEmail:
		return true
	}
	return false
}

type CustomHostnameSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 CustomHostnameSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion CustomHostnameSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 CustomHostnameSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   customHostnameSslSettingsJSON   `json:"-"`
}

// customHostnameSslSettingsJSON contains the JSON metadata for the struct
// [CustomHostnameSslSettings]
type customHostnameSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameSslSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type CustomHostnameSslSettingsEarlyHints string

const (
	CustomHostnameSslSettingsEarlyHintsOn  CustomHostnameSslSettingsEarlyHints = "on"
	CustomHostnameSslSettingsEarlyHintsOff CustomHostnameSslSettingsEarlyHints = "off"
)

func (r CustomHostnameSslSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameSslSettingsEarlyHintsOn, CustomHostnameSslSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameSslSettingsHttp2 string

const (
	CustomHostnameSslSettingsHttp2On  CustomHostnameSslSettingsHttp2 = "on"
	CustomHostnameSslSettingsHttp2Off CustomHostnameSslSettingsHttp2 = "off"
)

func (r CustomHostnameSslSettingsHttp2) IsKnown() bool {
	switch r {
	case CustomHostnameSslSettingsHttp2On, CustomHostnameSslSettingsHttp2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameSslSettingsMinTlsVersion string

const (
	CustomHostnameSslSettingsMinTlsVersion1_0 CustomHostnameSslSettingsMinTlsVersion = "1.0"
	CustomHostnameSslSettingsMinTlsVersion1_1 CustomHostnameSslSettingsMinTlsVersion = "1.1"
	CustomHostnameSslSettingsMinTlsVersion1_2 CustomHostnameSslSettingsMinTlsVersion = "1.2"
	CustomHostnameSslSettingsMinTlsVersion1_3 CustomHostnameSslSettingsMinTlsVersion = "1.3"
)

func (r CustomHostnameSslSettingsMinTlsVersion) IsKnown() bool {
	switch r {
	case CustomHostnameSslSettingsMinTlsVersion1_0, CustomHostnameSslSettingsMinTlsVersion1_1, CustomHostnameSslSettingsMinTlsVersion1_2, CustomHostnameSslSettingsMinTlsVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameSslSettingsTls1_3 string

const (
	CustomHostnameSslSettingsTls1_3On  CustomHostnameSslSettingsTls1_3 = "on"
	CustomHostnameSslSettingsTls1_3Off CustomHostnameSslSettingsTls1_3 = "off"
)

func (r CustomHostnameSslSettingsTls1_3) IsKnown() bool {
	switch r {
	case CustomHostnameSslSettingsTls1_3On, CustomHostnameSslSettingsTls1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type CustomHostnameSslStatus string

const (
	CustomHostnameSslStatusInitializing         CustomHostnameSslStatus = "initializing"
	CustomHostnameSslStatusPendingValidation    CustomHostnameSslStatus = "pending_validation"
	CustomHostnameSslStatusDeleted              CustomHostnameSslStatus = "deleted"
	CustomHostnameSslStatusPendingIssuance      CustomHostnameSslStatus = "pending_issuance"
	CustomHostnameSslStatusPendingDeployment    CustomHostnameSslStatus = "pending_deployment"
	CustomHostnameSslStatusPendingDeletion      CustomHostnameSslStatus = "pending_deletion"
	CustomHostnameSslStatusPendingExpiration    CustomHostnameSslStatus = "pending_expiration"
	CustomHostnameSslStatusExpired              CustomHostnameSslStatus = "expired"
	CustomHostnameSslStatusActive               CustomHostnameSslStatus = "active"
	CustomHostnameSslStatusInitializingTimedOut CustomHostnameSslStatus = "initializing_timed_out"
	CustomHostnameSslStatusValidationTimedOut   CustomHostnameSslStatus = "validation_timed_out"
	CustomHostnameSslStatusIssuanceTimedOut     CustomHostnameSslStatus = "issuance_timed_out"
	CustomHostnameSslStatusDeploymentTimedOut   CustomHostnameSslStatus = "deployment_timed_out"
	CustomHostnameSslStatusDeletionTimedOut     CustomHostnameSslStatus = "deletion_timed_out"
	CustomHostnameSslStatusPendingCleanup       CustomHostnameSslStatus = "pending_cleanup"
	CustomHostnameSslStatusStagingDeployment    CustomHostnameSslStatus = "staging_deployment"
	CustomHostnameSslStatusStagingActive        CustomHostnameSslStatus = "staging_active"
	CustomHostnameSslStatusDeactivating         CustomHostnameSslStatus = "deactivating"
	CustomHostnameSslStatusInactive             CustomHostnameSslStatus = "inactive"
	CustomHostnameSslStatusBackupIssued         CustomHostnameSslStatus = "backup_issued"
	CustomHostnameSslStatusHoldingDeployment    CustomHostnameSslStatus = "holding_deployment"
)

func (r CustomHostnameSslStatus) IsKnown() bool {
	switch r {
	case CustomHostnameSslStatusInitializing, CustomHostnameSslStatusPendingValidation, CustomHostnameSslStatusDeleted, CustomHostnameSslStatusPendingIssuance, CustomHostnameSslStatusPendingDeployment, CustomHostnameSslStatusPendingDeletion, CustomHostnameSslStatusPendingExpiration, CustomHostnameSslStatusExpired, CustomHostnameSslStatusActive, CustomHostnameSslStatusInitializingTimedOut, CustomHostnameSslStatusValidationTimedOut, CustomHostnameSslStatusIssuanceTimedOut, CustomHostnameSslStatusDeploymentTimedOut, CustomHostnameSslStatusDeletionTimedOut, CustomHostnameSslStatusPendingCleanup, CustomHostnameSslStatusStagingDeployment, CustomHostnameSslStatusStagingActive, CustomHostnameSslStatusDeactivating, CustomHostnameSslStatusInactive, CustomHostnameSslStatusBackupIssued, CustomHostnameSslStatusHoldingDeployment:
		return true
	}
	return false
}

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameSslType string

const (
	CustomHostnameSslTypeDv CustomHostnameSslType = "dv"
)

func (r CustomHostnameSslType) IsKnown() bool {
	switch r {
	case CustomHostnameSslTypeDv:
		return true
	}
	return false
}

type CustomHostnameSslValidationError struct {
	// A domain validation error.
	Message string                               `json:"message"`
	JSON    customHostnameSslValidationErrorJSON `json:"-"`
}

// customHostnameSslValidationErrorJSON contains the JSON metadata for the struct
// [CustomHostnameSslValidationError]
type customHostnameSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameSslValidationErrorJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameSslValidationRecord struct {
	// The set of email addresses that the certificate authority (CA) will use to
	// complete domain validation.
	Emails []string `json:"emails"`
	// The content that the certificate authority (CA) will expect to find at the
	// http_url during the domain validation.
	HTTPBody string `json:"http_body"`
	// The url that will be checked during domain validation.
	HTTPURL string `json:"http_url"`
	// The hostname that the certificate authority (CA) will check for a TXT record
	// during domain validation .
	TxtName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TxtValue string                                `json:"txt_value"`
	JSON     customHostnameSslValidationRecordJSON `json:"-"`
}

// customHostnameSslValidationRecordJSON contains the JSON metadata for the struct
// [CustomHostnameSslValidationRecord]
type customHostnameSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameSslValidationRecordJSON) RawJSON() string {
	return r.raw
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                  `json:"value"`
	JSON  customHostnameOwnershipVerificationJSON `json:"-"`
}

// customHostnameOwnershipVerificationJSON contains the JSON metadata for the
// struct [CustomHostnameOwnershipVerification]
type customHostnameOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameOwnershipVerificationJSON) RawJSON() string {
	return r.raw
}

// DNS Record type.
type CustomHostnameOwnershipVerificationType string

const (
	CustomHostnameOwnershipVerificationTypeTxt CustomHostnameOwnershipVerificationType = "txt"
)

func (r CustomHostnameOwnershipVerificationType) IsKnown() bool {
	switch r {
	case CustomHostnameOwnershipVerificationTypeTxt:
		return true
	}
	return false
}

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                      `json:"http_url"`
	JSON    customHostnameOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameOwnershipVerificationHTTPJSON contains the JSON metadata for the
// struct [CustomHostnameOwnershipVerificationHTTP]
type customHostnameOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameOwnershipVerificationHTTPJSON) RawJSON() string {
	return r.raw
}

// Status of the hostname's activation.
type CustomHostnameStatus string

const (
	CustomHostnameStatusActive             CustomHostnameStatus = "active"
	CustomHostnameStatusPending            CustomHostnameStatus = "pending"
	CustomHostnameStatusActiveRedeploying  CustomHostnameStatus = "active_redeploying"
	CustomHostnameStatusMoved              CustomHostnameStatus = "moved"
	CustomHostnameStatusPendingDeletion    CustomHostnameStatus = "pending_deletion"
	CustomHostnameStatusDeleted            CustomHostnameStatus = "deleted"
	CustomHostnameStatusPendingBlocked     CustomHostnameStatus = "pending_blocked"
	CustomHostnameStatusPendingMigration   CustomHostnameStatus = "pending_migration"
	CustomHostnameStatusPendingProvisioned CustomHostnameStatus = "pending_provisioned"
	CustomHostnameStatusTestPending        CustomHostnameStatus = "test_pending"
	CustomHostnameStatusTestActive         CustomHostnameStatus = "test_active"
	CustomHostnameStatusTestActiveApex     CustomHostnameStatus = "test_active_apex"
	CustomHostnameStatusTestBlocked        CustomHostnameStatus = "test_blocked"
	CustomHostnameStatusTestFailed         CustomHostnameStatus = "test_failed"
	CustomHostnameStatusProvisioned        CustomHostnameStatus = "provisioned"
	CustomHostnameStatusBlocked            CustomHostnameStatus = "blocked"
)

func (r CustomHostnameStatus) IsKnown() bool {
	switch r {
	case CustomHostnameStatusActive, CustomHostnameStatusPending, CustomHostnameStatusActiveRedeploying, CustomHostnameStatusMoved, CustomHostnameStatusPendingDeletion, CustomHostnameStatusDeleted, CustomHostnameStatusPendingBlocked, CustomHostnameStatusPendingMigration, CustomHostnameStatusPendingProvisioned, CustomHostnameStatusTestPending, CustomHostnameStatusTestActive, CustomHostnameStatusTestActiveApex, CustomHostnameStatusTestBlocked, CustomHostnameStatusTestFailed, CustomHostnameStatusProvisioned, CustomHostnameStatusBlocked:
		return true
	}
	return false
}

type CustomHostnameResponseSingle struct {
	Result CustomHostname                   `json:"result"`
	JSON   customHostnameResponseSingleJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// customHostnameResponseSingleJSON contains the JSON metadata for the struct
// [CustomHostnameResponseSingle]
type customHostnameResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameResponseSingleJSON) RawJSON() string {
	return r.raw
}

// SSL properties used when creating the custom hostname.
type SslPostPropertiesParam struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[SslPostPropertiesBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CertificateAuthorityCustomHostname] `json:"certificate_authority"`
	// Whether or not to add Cloudflare Branding for the order. This will add a
	// subdomain of sni.cloudflaressl.com as the Common Name if set to true
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
	// Array of custom certificate and key pairs (1 or 2 pairs allowed)
	CustomCertBundle param.Field[[]CustomCertAndKeyParam] `json:"custom_cert_bundle"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[SslPostPropertiesMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[SslSettingsParam] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[SslPostPropertiesType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r SslPostPropertiesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type SslPostPropertiesBundleMethod string

const (
	SslPostPropertiesBundleMethodUbiquitous SslPostPropertiesBundleMethod = "ubiquitous"
	SslPostPropertiesBundleMethodOptimal    SslPostPropertiesBundleMethod = "optimal"
	SslPostPropertiesBundleMethodForce      SslPostPropertiesBundleMethod = "force"
)

func (r SslPostPropertiesBundleMethod) IsKnown() bool {
	switch r {
	case SslPostPropertiesBundleMethodUbiquitous, SslPostPropertiesBundleMethodOptimal, SslPostPropertiesBundleMethodForce:
		return true
	}
	return false
}

// Domain control validation (DCV) method used for this hostname.
type SslPostPropertiesMethod string

const (
	SslPostPropertiesMethodHTTP  SslPostPropertiesMethod = "http"
	SslPostPropertiesMethodTxt   SslPostPropertiesMethod = "txt"
	SslPostPropertiesMethodEmail SslPostPropertiesMethod = "email"
)

func (r SslPostPropertiesMethod) IsKnown() bool {
	switch r {
	case SslPostPropertiesMethodHTTP, SslPostPropertiesMethodTxt, SslPostPropertiesMethodEmail:
		return true
	}
	return false
}

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type SslPostPropertiesType string

const (
	SslPostPropertiesTypeDv SslPostPropertiesType = "dv"
)

func (r SslPostPropertiesType) IsKnown() bool {
	switch r {
	case SslPostPropertiesTypeDv:
		return true
	}
	return false
}

// SSL specific settings.
type SslSettingsParam struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[SslSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 param.Field[SslSettingsHttp2] `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion param.Field[SslSettingsMinTlsVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 param.Field[SslSettingsTls1_3] `json:"tls_1_3"`
}

func (r SslSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type SslSettingsEarlyHints string

const (
	SslSettingsEarlyHintsOn  SslSettingsEarlyHints = "on"
	SslSettingsEarlyHintsOff SslSettingsEarlyHints = "off"
)

func (r SslSettingsEarlyHints) IsKnown() bool {
	switch r {
	case SslSettingsEarlyHintsOn, SslSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type SslSettingsHttp2 string

const (
	SslSettingsHttp2On  SslSettingsHttp2 = "on"
	SslSettingsHttp2Off SslSettingsHttp2 = "off"
)

func (r SslSettingsHttp2) IsKnown() bool {
	switch r {
	case SslSettingsHttp2On, SslSettingsHttp2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type SslSettingsMinTlsVersion string

const (
	SslSettingsMinTlsVersion1_0 SslSettingsMinTlsVersion = "1.0"
	SslSettingsMinTlsVersion1_1 SslSettingsMinTlsVersion = "1.1"
	SslSettingsMinTlsVersion1_2 SslSettingsMinTlsVersion = "1.2"
	SslSettingsMinTlsVersion1_3 SslSettingsMinTlsVersion = "1.3"
)

func (r SslSettingsMinTlsVersion) IsKnown() bool {
	switch r {
	case SslSettingsMinTlsVersion1_0, SslSettingsMinTlsVersion1_1, SslSettingsMinTlsVersion1_2, SslSettingsMinTlsVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type SslSettingsTls1_3 string

const (
	SslSettingsTls1_3On  SslSettingsTls1_3 = "on"
	SslSettingsTls1_3Off SslSettingsTls1_3 = "off"
)

func (r SslSettingsTls1_3) IsKnown() bool {
	switch r {
	case SslSettingsTls1_3On, SslSettingsTls1_3Off:
		return true
	}
	return false
}

type ZoneCustomHostnameListResponse struct {
	Result []CustomHostname                   `json:"result"`
	JSON   zoneCustomHostnameListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneCustomHostnameListResponseJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameListResponse]
type zoneCustomHostnameListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomHostnameDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON zoneCustomHostnameDeleteResponseJSON `json:"-"`
}

// zoneCustomHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameDeleteResponse]
type zoneCustomHostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCustomHostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCustomHostnameNewParams struct {
	// The custom hostname that will point to your hostname via CNAME.
	Hostname param.Field[string] `json:"hostname,required"`
	// SSL properties used when creating the custom hostname.
	Ssl param.Field[SslPostPropertiesParam] `json:"ssl,required"`
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata param.Field[map[string]string] `json:"custom_metadata"`
}

func (r ZoneCustomHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomHostnameUpdateParams struct {
	// Unique key/value metadata for this hostname. These are per-hostname (customer)
	// settings.
	CustomMetadata param.Field[map[string]string] `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer param.Field[string] `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSni param.Field[string] `json:"custom_origin_sni"`
	// SSL properties used when creating the custom hostname.
	Ssl param.Field[SslPostPropertiesParam] `json:"ssl"`
}

func (r ZoneCustomHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomHostnameListParams struct {
	// Hostname ID to match against. This ID was generated and returned during the
	// initial custom_hostname creation. This parameter cannot be used with the
	// 'hostname' parameter.
	ID param.Field[string] `query:"id"`
	// Direction to order hostnames.
	Direction param.Field[ZoneCustomHostnameListParamsDirection] `query:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with
	// the 'id' parameter.
	Hostname param.Field[string] `query:"hostname"`
	// Field to order hostnames by.
	Order param.Field[ZoneCustomHostnameListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of hostnames per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether to filter hostnames based on if they have SSL enabled.
	Ssl param.Field[ZoneCustomHostnameListParamsSsl] `query:"ssl"`
}

// URLQuery serializes [ZoneCustomHostnameListParams]'s query parameters as
// `url.Values`.
func (r ZoneCustomHostnameListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order hostnames.
type ZoneCustomHostnameListParamsDirection string

const (
	ZoneCustomHostnameListParamsDirectionAsc  ZoneCustomHostnameListParamsDirection = "asc"
	ZoneCustomHostnameListParamsDirectionDesc ZoneCustomHostnameListParamsDirection = "desc"
)

func (r ZoneCustomHostnameListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneCustomHostnameListParamsDirectionAsc, ZoneCustomHostnameListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order hostnames by.
type ZoneCustomHostnameListParamsOrder string

const (
	ZoneCustomHostnameListParamsOrderSsl       ZoneCustomHostnameListParamsOrder = "ssl"
	ZoneCustomHostnameListParamsOrderSslStatus ZoneCustomHostnameListParamsOrder = "ssl_status"
)

func (r ZoneCustomHostnameListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneCustomHostnameListParamsOrderSsl, ZoneCustomHostnameListParamsOrderSslStatus:
		return true
	}
	return false
}

// Whether to filter hostnames based on if they have SSL enabled.
type ZoneCustomHostnameListParamsSsl float64

const (
	ZoneCustomHostnameListParamsSsl0 ZoneCustomHostnameListParamsSsl = 0
	ZoneCustomHostnameListParamsSsl1 ZoneCustomHostnameListParamsSsl = 1
)

func (r ZoneCustomHostnameListParamsSsl) IsKnown() bool {
	switch r {
	case ZoneCustomHostnameListParamsSsl0, ZoneCustomHostnameListParamsSsl1:
		return true
	}
	return false
}

type ZoneCustomHostnameDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneCustomHostnameDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
