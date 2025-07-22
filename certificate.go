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

// CertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificateService] method instead.
type CertificateService struct {
	Options []option.RequestOption
}

// NewCertificateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificateService(opts ...option.RequestOption) (r *CertificateService) {
	r = &CertificateService{}
	r.Options = opts
	return
}

// Create an Origin CA certificate. You can use an Origin CA Key as your User
// Service Key or an API token when calling this endpoint ([see above](#requests)).
func (r *CertificateService) New(ctx context.Context, body CertificateNewParams, opts ...option.RequestOption) (res *SingleCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an existing Origin CA certificate by its serial number. You can use an
// Origin CA Key as your User Service Key or an API token when calling this
// endpoint ([see above](#requests)).
func (r *CertificateService) Get(ctx context.Context, certificateID string, opts ...option.RequestOption) (res *SingleCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all existing Origin CA certificates for a given zone. You can use an Origin
// CA Key as your User Service Key or an API token when calling this endpoint
// ([see above](#requests)).
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *CertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Revoke an existing Origin CA certificate by its serial number. You can use an
// Origin CA Key as your User Service Key or an API token when calling this
// endpoint ([see above](#requests)).
func (r *CertificateService) Revoke(ctx context.Context, certificateID string, body CertificateRevokeParams, opts ...option.RequestOption) (res *CertificateRevokeResponse, err error) {
	opts = append(r.Options[:], opts...)
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("certificates/%s", certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type RequestType string

const (
	RequestTypeOriginRsa          RequestType = "origin-rsa"
	RequestTypeOriginEcc          RequestType = "origin-ecc"
	RequestTypeKeylessCertificate RequestType = "keyless-certificate"
)

func (r RequestType) IsKnown() bool {
	switch r {
	case RequestTypeOriginRsa, RequestTypeOriginEcc, RequestTypeKeylessCertificate:
		return true
	}
	return false
}

// The number of days for which the certificate should be valid.
type RequestedValidity float64

const (
	RequestedValidity7    RequestedValidity = 7
	RequestedValidity30   RequestedValidity = 30
	RequestedValidity90   RequestedValidity = 90
	RequestedValidity365  RequestedValidity = 365
	RequestedValidity730  RequestedValidity = 730
	RequestedValidity1095 RequestedValidity = 1095
	RequestedValidity5475 RequestedValidity = 5475
)

func (r RequestedValidity) IsKnown() bool {
	switch r {
	case RequestedValidity7, RequestedValidity30, RequestedValidity90, RequestedValidity365, RequestedValidity730, RequestedValidity1095, RequestedValidity5475:
		return true
	}
	return false
}

type SingleCertificateResponse struct {
	Result TlsCertificates               `json:"result"`
	JSON   singleCertificateResponseJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// singleCertificateResponseJSON contains the JSON metadata for the struct
// [SingleCertificateResponse]
type singleCertificateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleCertificateResponseJSON) RawJSON() string {
	return r.raw
}

type TlsCertificates struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr,required"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames []string `json:"hostnames,required"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType RequestType `json:"request_type,required"`
	// The number of days for which the certificate should be valid.
	RequestedValidity RequestedValidity `json:"requested_validity,required"`
	// Identifier
	ID string `json:"id"`
	// The Origin CA certificate. Will be newline-encoded.
	Certificate string `json:"certificate"`
	// When the certificate will expire.
	ExpiresOn string              `json:"expires_on"`
	JSON      tlsCertificatesJSON `json:"-"`
}

// tlsCertificatesJSON contains the JSON metadata for the struct [TlsCertificates]
type tlsCertificatesJSON struct {
	Csr               apijson.Field
	Hostnames         apijson.Field
	RequestType       apijson.Field
	RequestedValidity apijson.Field
	ID                apijson.Field
	Certificate       apijson.Field
	ExpiresOn         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TlsCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesJSON) RawJSON() string {
	return r.raw
}

type CertificateListResponse struct {
	Result []TlsCertificates           `json:"result"`
	JSON   certificateListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// certificateListResponseJSON contains the JSON metadata for the struct
// [CertificateListResponse]
type certificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateListResponseJSON) RawJSON() string {
	return r.raw
}

type CertificateRevokeResponse struct {
	Result CertificateRevokeResponseResult `json:"result"`
	JSON   certificateRevokeResponseJSON   `json:"-"`
}

// certificateRevokeResponseJSON contains the JSON metadata for the struct
// [CertificateRevokeResponse]
type certificateRevokeResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateRevokeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateRevokeResponseJSON) RawJSON() string {
	return r.raw
}

type CertificateRevokeResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// When the certificate was revoked.
	RevokedAt time.Time                           `json:"revoked_at" format:"date-time"`
	JSON      certificateRevokeResponseResultJSON `json:"-"`
}

// certificateRevokeResponseResultJSON contains the JSON metadata for the struct
// [CertificateRevokeResponseResult]
type certificateRevokeResponseResultJSON struct {
	ID          apijson.Field
	RevokedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateRevokeResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateRevokeResponseResultJSON) RawJSON() string {
	return r.raw
}

type CertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr"`
	// Array of hostnames or wildcard names (e.g., \*.example.com) bound to the
	// certificate.
	Hostnames param.Field[[]string] `json:"hostnames"`
	// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
	// or "keyless-certificate" (for Keyless SSL servers).
	RequestType param.Field[RequestType] `json:"request_type"`
	// The number of days for which the certificate should be valid.
	RequestedValidity param.Field[RequestedValidity] `json:"requested_validity"`
}

func (r CertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `query:"zone_id,required"`
}

// URLQuery serializes [CertificateListParams]'s query parameters as `url.Values`.
func (r CertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CertificateRevokeParams struct {
	Body interface{} `json:"body,required"`
}

func (r CertificateRevokeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
