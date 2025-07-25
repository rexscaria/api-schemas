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

// ZoneOriginTlsClientAuthHostnameService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneOriginTlsClientAuthHostnameService] method instead.
type ZoneOriginTlsClientAuthHostnameService struct {
	Options      []option.RequestOption
	Certificates *ZoneOriginTlsClientAuthHostnameCertificateService
}

// NewZoneOriginTlsClientAuthHostnameService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthHostnameService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthHostnameService) {
	r = &ZoneOriginTlsClientAuthHostnameService{}
	r.Options = opts
	r.Certificates = NewZoneOriginTlsClientAuthHostnameCertificateService(opts...)
	return
}

// Get the Hostname Status for Client Authentication
func (r *ZoneOriginTlsClientAuthHostnameService) Get(ctx context.Context, zoneID string, hostname string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/%s", zoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Associate a hostname to a certificate and enable, disable or invalidate the
// association. If disabled, client certificate will not be sent to the hostname
// even if activated at the zone level. 100 maximum associations on a single
// certificate are allowed. Note: Use a null value for parameter _enabled_ to
// invalidate the association.
func (r *ZoneOriginTlsClientAuthHostnameService) Update(ctx context.Context, zoneID string, body ZoneOriginTlsClientAuthHostnameUpdateParams, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CertidObject struct {
	// Identifier.
	CertID string `json:"cert_id"`
	// Status of the certificate or the association.
	CertStatus StatusCertificate `json:"cert_status"`
	// The time when the certificate was updated.
	CertUpdatedAt time.Time `json:"cert_updated_at" format:"date-time"`
	// The time when the certificate was uploaded.
	CertUploadedOn time.Time `json:"cert_uploaded_on" format:"date-time"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The time when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status StatusCertificate `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      certidObjectJSON `json:"-"`
}

// certidObjectJSON contains the JSON metadata for the struct [CertidObject]
type certidObjectJSON struct {
	CertID         apijson.Field
	CertStatus     apijson.Field
	CertUpdatedAt  apijson.Field
	CertUploadedOn apijson.Field
	Certificate    apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresOn      apijson.Field
	Hostname       apijson.Field
	Issuer         apijson.Field
	SerialNumber   apijson.Field
	Signature      apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CertidObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certidObjectJSON) RawJSON() string {
	return r.raw
}

type HostnameAuthenticatedOriginPull struct {
	// Identifier.
	ID string `json:"id"`
	// Identifier.
	CertID string `json:"cert_id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The hostname certificate's private key.
	PrivateKey string                              `json:"private_key"`
	JSON       hostnameAuthenticatedOriginPullJSON `json:"-"`
	CertidObject
}

// hostnameAuthenticatedOriginPullJSON contains the JSON metadata for the struct
// [HostnameAuthenticatedOriginPull]
type hostnameAuthenticatedOriginPullJSON struct {
	ID          apijson.Field
	CertID      apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	Hostname    apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameAuthenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type StatusCertificate string

const (
	StatusCertificateInitializing       StatusCertificate = "initializing"
	StatusCertificatePendingDeployment  StatusCertificate = "pending_deployment"
	StatusCertificatePendingDeletion    StatusCertificate = "pending_deletion"
	StatusCertificateActive             StatusCertificate = "active"
	StatusCertificateDeleted            StatusCertificate = "deleted"
	StatusCertificateDeploymentTimedOut StatusCertificate = "deployment_timed_out"
	StatusCertificateDeletionTimedOut   StatusCertificate = "deletion_timed_out"
)

func (r StatusCertificate) IsKnown() bool {
	switch r {
	case StatusCertificateInitializing, StatusCertificatePendingDeployment, StatusCertificatePendingDeletion, StatusCertificateActive, StatusCertificateDeleted, StatusCertificateDeploymentTimedOut, StatusCertificateDeletionTimedOut:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthHostnameGetResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneOriginTlsClientAuthHostnameGetResponseSuccess `json:"success,required"`
	Result  CertidObject                                      `json:"result"`
	JSON    zoneOriginTlsClientAuthHostnameGetResponseJSON    `json:"-"`
}

// zoneOriginTlsClientAuthHostnameGetResponseJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthHostnameGetResponse]
type zoneOriginTlsClientAuthHostnameGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneOriginTlsClientAuthHostnameGetResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameGetResponseSuccessTrue ZoneOriginTlsClientAuthHostnameGetResponseSuccess = true
)

func (r ZoneOriginTlsClientAuthHostnameGetResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneOriginTlsClientAuthHostnameGetResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthHostnameUpdateResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneOriginTlsClientAuthHostnameUpdateResponseSuccess    `json:"success,required"`
	Result     []HostnameAuthenticatedOriginPull                       `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthHostnameUpdateResponseResultInfo `json:"result_info"`
	JSON       zoneOriginTlsClientAuthHostnameUpdateResponseJSON       `json:"-"`
}

// zoneOriginTlsClientAuthHostnameUpdateResponseJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthHostnameUpdateResponse]
type zoneOriginTlsClientAuthHostnameUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneOriginTlsClientAuthHostnameUpdateResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameUpdateResponseSuccessTrue ZoneOriginTlsClientAuthHostnameUpdateResponseSuccess = true
)

func (r ZoneOriginTlsClientAuthHostnameUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneOriginTlsClientAuthHostnameUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthHostnameUpdateResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                     `json:"total_count"`
	JSON       zoneOriginTlsClientAuthHostnameUpdateResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct
// [ZoneOriginTlsClientAuthHostnameUpdateResponseResultInfo]
type zoneOriginTlsClientAuthHostnameUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameUpdateResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthHostnameUpdateParams struct {
	Config param.Field[[]ZoneOriginTlsClientAuthHostnameUpdateParamsConfig] `json:"config,required"`
}

func (r ZoneOriginTlsClientAuthHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneOriginTlsClientAuthHostnameUpdateParamsConfig struct {
	// Certificate identifier tag.
	CertID param.Field[string] `json:"cert_id"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname param.Field[string] `json:"hostname"`
}

func (r ZoneOriginTlsClientAuthHostnameUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
