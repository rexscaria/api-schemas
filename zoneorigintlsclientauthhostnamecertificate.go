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

// ZoneOriginTlsClientAuthHostnameCertificateService contains methods and other
// services that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneOriginTlsClientAuthHostnameCertificateService] method instead.
type ZoneOriginTlsClientAuthHostnameCertificateService struct {
	Options []option.RequestOption
}

// NewZoneOriginTlsClientAuthHostnameCertificateService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneOriginTlsClientAuthHostnameCertificateService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthHostnameCertificateService) {
	r = &ZoneOriginTlsClientAuthHostnameCertificateService{}
	r.Options = opts
	return
}

// Get the certificate by ID to be used for client authentication on a hostname.
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Get(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleHostname, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Certificates
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthHostnameCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Hostname Client Certificate
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Delete(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleHostname, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload a certificate to be used for client authentication on a hostname. 10
// hostname certificates per zone are allowed.
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Upload(ctx context.Context, zoneID string, body ZoneOriginTlsClientAuthHostnameCertificateUploadParams, opts ...option.RequestOption) (res *CertificateResponseSingleHostname, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CertificateResponseSingleHostname struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificateResponseSingleHostnameSuccess `json:"success,required"`
	Result  CertificateResponseSingleHostnameResult  `json:"result"`
	JSON    certificateResponseSingleHostnameJSON    `json:"-"`
}

// certificateResponseSingleHostnameJSON contains the JSON metadata for the struct
// [CertificateResponseSingleHostname]
type certificateResponseSingleHostnameJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleHostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseSingleHostnameJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificateResponseSingleHostnameSuccess bool

const (
	CertificateResponseSingleHostnameSuccessTrue CertificateResponseSingleHostnameSuccess = true
)

func (r CertificateResponseSingleHostnameSuccess) IsKnown() bool {
	switch r {
	case CertificateResponseSingleHostnameSuccessTrue:
		return true
	}
	return false
}

type CertificateResponseSingleHostnameResult struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status StatusCertificate `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                                   `json:"uploaded_on" format:"date-time"`
	JSON       certificateResponseSingleHostnameResultJSON `json:"-"`
}

// certificateResponseSingleHostnameResultJSON contains the JSON metadata for the
// struct [CertificateResponseSingleHostnameResult]
type certificateResponseSingleHostnameResultJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CertificateResponseSingleHostnameResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseSingleHostnameResultJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthHostnameCertificateListResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccess    `json:"success,required"`
	Result     []HostnameAuthenticatedOriginPull                                `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthHostnameCertificateListResponseResultInfo `json:"result_info"`
	JSON       zoneOriginTlsClientAuthHostnameCertificateListResponseJSON       `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateListResponseJSON contains the JSON
// metadata for the struct [ZoneOriginTlsClientAuthHostnameCertificateListResponse]
type zoneOriginTlsClientAuthHostnameCertificateListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccess bool

const (
	ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccessTrue ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccess = true
)

func (r ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneOriginTlsClientAuthHostnameCertificateListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthHostnameCertificateListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                              `json:"total_count"`
	JSON       zoneOriginTlsClientAuthHostnameCertificateListResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthHostnameCertificateListResponseResultInfoJSON contains
// the JSON metadata for the struct
// [ZoneOriginTlsClientAuthHostnameCertificateListResponseResultInfo]
type zoneOriginTlsClientAuthHostnameCertificateListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameCertificateListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthHostnameCertificateUploadParams struct {
	// The hostname certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The hostname certificate's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r ZoneOriginTlsClientAuthHostnameCertificateUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
