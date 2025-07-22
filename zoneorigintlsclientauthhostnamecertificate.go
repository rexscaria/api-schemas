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
func (r *ZoneOriginTlsClientAuthHostnameCertificateService) Delete(ctx context.Context, zoneID string, certificateID string, body ZoneOriginTlsClientAuthHostnameCertificateDeleteParams, opts ...option.RequestOption) (res *CertificateResponseSingleHostname, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
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
	Result CertificateResponseSingleHostnameResult `json:"result"`
	JSON   certificateResponseSingleHostnameJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// certificateResponseSingleHostnameJSON contains the JSON metadata for the struct
// [CertificateResponseSingleHostname]
type certificateResponseSingleHostnameJSON struct {
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

type CertificateResponseSingleHostnameResult struct {
	// Identifier
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
	Result []HostnameAuthenticatedOriginPull                          `json:"result"`
	JSON   zoneOriginTlsClientAuthHostnameCertificateListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneOriginTlsClientAuthHostnameCertificateListResponseJSON contains the JSON
// metadata for the struct [ZoneOriginTlsClientAuthHostnameCertificateListResponse]
type zoneOriginTlsClientAuthHostnameCertificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthHostnameCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthHostnameCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthHostnameCertificateDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r ZoneOriginTlsClientAuthHostnameCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
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
