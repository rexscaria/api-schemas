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

// AccountMtlsCertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountMtlsCertificateService] method instead.
type AccountMtlsCertificateService struct {
	Options []option.RequestOption
}

// NewAccountMtlsCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountMtlsCertificateService(opts ...option.RequestOption) (r *AccountMtlsCertificateService) {
	r = &AccountMtlsCertificateService{}
	r.Options = opts
	return
}

// Fetches a single mTLS certificate.
func (r *AccountMtlsCertificateService) Get(ctx context.Context, accountID string, mtlsCertificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleMtls, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *AccountMtlsCertificateService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountMtlsCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes the mTLS certificate unless the certificate is in use by one or more
// Cloudflare services.
func (r *AccountMtlsCertificateService) Delete(ctx context.Context, accountID string, mtlsCertificateID string, body AccountMtlsCertificateDeleteParams, opts ...option.RequestOption) (res *CertificateResponseSingleMtls, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *AccountMtlsCertificateService) ListAssociations(ctx context.Context, accountID string, mtlsCertificateID string, opts ...option.RequestOption) (res *AccountMtlsCertificateListAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", accountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
func (r *AccountMtlsCertificateService) Upload(ctx context.Context, accountID string, body AccountMtlsCertificateUploadParams, opts ...option.RequestOption) (res *AccountMtlsCertificateUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type APIResponseCollectionTlsCertificates struct {
	ResultInfo APIResponseCollectionTlsCertificatesResultInfo `json:"result_info"`
	JSON       apiResponseCollectionTlsCertificatesJSON       `json:"-"`
	APIResponseTlsCertificates
}

// apiResponseCollectionTlsCertificatesJSON contains the JSON metadata for the
// struct [APIResponseCollectionTlsCertificates]
type apiResponseCollectionTlsCertificatesJSON struct {
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTlsCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTlsCertificatesJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionTlsCertificatesResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       apiResponseCollectionTlsCertificatesResultInfoJSON `json:"-"`
}

// apiResponseCollectionTlsCertificatesResultInfoJSON contains the JSON metadata
// for the struct [APIResponseCollectionTlsCertificatesResultInfo]
type apiResponseCollectionTlsCertificatesResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionTlsCertificatesResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionTlsCertificatesResultInfoJSON) RawJSON() string {
	return r.raw
}

type APIResponseSingleTlsCertificates struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseSingleTlsCertificatesSuccess `json:"success,required"`
	JSON    apiResponseSingleTlsCertificatesJSON    `json:"-"`
}

// apiResponseSingleTlsCertificatesJSON contains the JSON metadata for the struct
// [APIResponseSingleTlsCertificates]
type apiResponseSingleTlsCertificatesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleTlsCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleTlsCertificatesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseSingleTlsCertificatesSuccess bool

const (
	APIResponseSingleTlsCertificatesSuccessTrue APIResponseSingleTlsCertificatesSuccess = true
)

func (r APIResponseSingleTlsCertificatesSuccess) IsKnown() bool {
	switch r {
	case APIResponseSingleTlsCertificatesSuccessTrue:
		return true
	}
	return false
}

type APIResponseTlsCertificates struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseTlsCertificatesSuccess `json:"success,required"`
	JSON    apiResponseTlsCertificatesJSON    `json:"-"`
}

// apiResponseTlsCertificatesJSON contains the JSON metadata for the struct
// [APIResponseTlsCertificates]
type apiResponseTlsCertificatesJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseTlsCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseTlsCertificatesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseTlsCertificatesSuccess bool

const (
	APIResponseTlsCertificatesSuccessTrue APIResponseTlsCertificatesSuccess = true
)

func (r APIResponseTlsCertificatesSuccess) IsKnown() bool {
	switch r {
	case APIResponseTlsCertificatesSuccessTrue:
		return true
	}
	return false
}

type CertificateObject struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time             `json:"uploaded_on" format:"date-time"`
	JSON       certificateObjectJSON `json:"-"`
}

// certificateObjectJSON contains the JSON metadata for the struct
// [CertificateObject]
type certificateObjectJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CertificateObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateObjectJSON) RawJSON() string {
	return r.raw
}

type CertificateResponseSingleMtls struct {
	Result CertificateObject                 `json:"result"`
	JSON   certificateResponseSingleMtlsJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// certificateResponseSingleMtlsJSON contains the JSON metadata for the struct
// [CertificateResponseSingleMtls]
type certificateResponseSingleMtlsJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleMtls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseSingleMtlsJSON) RawJSON() string {
	return r.raw
}

type MessagesTlsCertificatesItem struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    messagesTlsCertificatesItemJSON `json:"-"`
}

// messagesTlsCertificatesItemJSON contains the JSON metadata for the struct
// [MessagesTlsCertificatesItem]
type messagesTlsCertificatesItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesTlsCertificatesItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesTlsCertificatesItemJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateListResponse struct {
	Result     []CertificateObject                          `json:"result"`
	ResultInfo AccountMtlsCertificateListResponseResultInfo `json:"result_info"`
	JSON       accountMtlsCertificateListResponseJSON       `json:"-"`
	APIResponseCollectionTlsCertificates
}

// accountMtlsCertificateListResponseJSON contains the JSON metadata for the struct
// [AccountMtlsCertificateListResponse]
type accountMtlsCertificateListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateListResponseResultInfo struct {
	Count      interface{} `json:"count"`
	Page       interface{} `json:"page"`
	PerPage    interface{} `json:"per_page"`
	TotalCount interface{} `json:"total_count"`
	// Total pages available of results
	TotalPages float64                                          `json:"total_pages"`
	JSON       accountMtlsCertificateListResponseResultInfoJSON `json:"-"`
}

// accountMtlsCertificateListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountMtlsCertificateListResponseResultInfo]
type accountMtlsCertificateListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateListAssociationsResponse struct {
	Result []AccountMtlsCertificateListAssociationsResponseResult `json:"result"`
	JSON   accountMtlsCertificateListAssociationsResponseJSON     `json:"-"`
	APIResponseCollectionTlsCertificates
}

// accountMtlsCertificateListAssociationsResponseJSON contains the JSON metadata
// for the struct [AccountMtlsCertificateListAssociationsResponse]
type accountMtlsCertificateListAssociationsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateListAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateListAssociationsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateListAssociationsResponseResult struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                                                   `json:"status"`
	JSON   accountMtlsCertificateListAssociationsResponseResultJSON `json:"-"`
}

// accountMtlsCertificateListAssociationsResponseResultJSON contains the JSON
// metadata for the struct [AccountMtlsCertificateListAssociationsResponseResult]
type accountMtlsCertificateListAssociationsResponseResultJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateListAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateListAssociationsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateUploadResponse struct {
	Result AccountMtlsCertificateUploadResponseResult `json:"result"`
	JSON   accountMtlsCertificateUploadResponseJSON   `json:"-"`
	APIResponseSingleTlsCertificates
}

// accountMtlsCertificateUploadResponseJSON contains the JSON metadata for the
// struct [AccountMtlsCertificateUploadResponse]
type accountMtlsCertificateUploadResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMtlsCertificateUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateUploadResponseJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateUploadResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                                      `json:"uploaded_on" format:"date-time"`
	JSON       accountMtlsCertificateUploadResponseResultJSON `json:"-"`
}

// accountMtlsCertificateUploadResponseResultJSON contains the JSON metadata for
// the struct [AccountMtlsCertificateUploadResponseResult]
type accountMtlsCertificateUploadResponseResultJSON struct {
	ID           apijson.Field
	Ca           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UpdatedAt    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountMtlsCertificateUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountMtlsCertificateUploadResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountMtlsCertificateDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountMtlsCertificateDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountMtlsCertificateUploadParams struct {
	// Indicates whether the certificate is a CA or leaf certificate.
	Ca param.Field[bool] `json:"ca,required"`
	// The uploaded root CA certificate.
	Certificates param.Field[string] `json:"certificates,required"`
	// Optional unique name for the certificate. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// The private key for the certificate
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r AccountMtlsCertificateUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
