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

// ZoneOriginTlsClientAuthService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneOriginTlsClientAuthService] method instead.
type ZoneOriginTlsClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *ZoneOriginTlsClientAuthHostnameService
	Settings  *ZoneOriginTlsClientAuthSettingService
}

// NewZoneOriginTlsClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneOriginTlsClientAuthService(opts ...option.RequestOption) (r *ZoneOriginTlsClientAuthService) {
	r = &ZoneOriginTlsClientAuthService{}
	r.Options = opts
	r.Hostnames = NewZoneOriginTlsClientAuthHostnameService(opts...)
	r.Settings = NewZoneOriginTlsClientAuthSettingService(opts...)
	return
}

// Get Certificate Details
func (r *ZoneOriginTlsClientAuthService) Get(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleOrigin, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Certificates
func (r *ZoneOriginTlsClientAuthService) List(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneOriginTlsClientAuthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Certificate
func (r *ZoneOriginTlsClientAuthService) Delete(ctx context.Context, zoneID string, certificateID string, opts ...option.RequestOption) (res *CertificateResponseSingleOrigin, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", zoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *ZoneOriginTlsClientAuthService) Upload(ctx context.Context, zoneID string, body ZoneOriginTlsClientAuthUploadParams, opts ...option.RequestOption) (res *CertificateResponseSingleOrigin, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CertificateResponseSingleOrigin struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success CertificateResponseSingleOriginSuccess `json:"success,required"`
	Result  ZoneAuthenticatedOriginPull            `json:"result"`
	JSON    certificateResponseSingleOriginJSON    `json:"-"`
}

// certificateResponseSingleOriginJSON contains the JSON metadata for the struct
// [CertificateResponseSingleOrigin]
type certificateResponseSingleOriginJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateResponseSingleOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateResponseSingleOriginJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CertificateResponseSingleOriginSuccess bool

const (
	CertificateResponseSingleOriginSuccessTrue CertificateResponseSingleOriginSuccess = true
)

func (r CertificateResponseSingleOriginSuccess) IsKnown() bool {
	switch r {
	case CertificateResponseSingleOriginSuccessTrue:
		return true
	}
	return false
}

type ZoneAuthenticatedOriginPull struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The zone's private key.
	PrivateKey string `json:"private_key"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate activation.
	Status ZoneAuthenticatedOriginPullStatus `json:"status"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                       `json:"uploaded_on" format:"date-time"`
	JSON       zoneAuthenticatedOriginPullJSON `json:"-"`
}

// zoneAuthenticatedOriginPullJSON contains the JSON metadata for the struct
// [ZoneAuthenticatedOriginPull]
type zoneAuthenticatedOriginPullJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	PrivateKey  apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAuthenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate activation.
type ZoneAuthenticatedOriginPullStatus string

const (
	ZoneAuthenticatedOriginPullStatusInitializing       ZoneAuthenticatedOriginPullStatus = "initializing"
	ZoneAuthenticatedOriginPullStatusPendingDeployment  ZoneAuthenticatedOriginPullStatus = "pending_deployment"
	ZoneAuthenticatedOriginPullStatusPendingDeletion    ZoneAuthenticatedOriginPullStatus = "pending_deletion"
	ZoneAuthenticatedOriginPullStatusActive             ZoneAuthenticatedOriginPullStatus = "active"
	ZoneAuthenticatedOriginPullStatusDeleted            ZoneAuthenticatedOriginPullStatus = "deleted"
	ZoneAuthenticatedOriginPullStatusDeploymentTimedOut ZoneAuthenticatedOriginPullStatus = "deployment_timed_out"
	ZoneAuthenticatedOriginPullStatusDeletionTimedOut   ZoneAuthenticatedOriginPullStatus = "deletion_timed_out"
)

func (r ZoneAuthenticatedOriginPullStatus) IsKnown() bool {
	switch r {
	case ZoneAuthenticatedOriginPullStatusInitializing, ZoneAuthenticatedOriginPullStatusPendingDeployment, ZoneAuthenticatedOriginPullStatusPendingDeletion, ZoneAuthenticatedOriginPullStatusActive, ZoneAuthenticatedOriginPullStatusDeleted, ZoneAuthenticatedOriginPullStatusDeploymentTimedOut, ZoneAuthenticatedOriginPullStatusDeletionTimedOut:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthListResponse struct {
	Errors   []MessagesTlsCertificatesItem `json:"errors,required"`
	Messages []MessagesTlsCertificatesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    ZoneOriginTlsClientAuthListResponseSuccess    `json:"success,required"`
	Result     []ZoneAuthenticatedOriginPull                 `json:"result"`
	ResultInfo ZoneOriginTlsClientAuthListResponseResultInfo `json:"result_info"`
	JSON       zoneOriginTlsClientAuthListResponseJSON       `json:"-"`
}

// zoneOriginTlsClientAuthListResponseJSON contains the JSON metadata for the
// struct [ZoneOriginTlsClientAuthListResponse]
type zoneOriginTlsClientAuthListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneOriginTlsClientAuthListResponseSuccess bool

const (
	ZoneOriginTlsClientAuthListResponseSuccessTrue ZoneOriginTlsClientAuthListResponseSuccess = true
)

func (r ZoneOriginTlsClientAuthListResponseSuccess) IsKnown() bool {
	switch r {
	case ZoneOriginTlsClientAuthListResponseSuccessTrue:
		return true
	}
	return false
}

type ZoneOriginTlsClientAuthListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                           `json:"total_count"`
	JSON       zoneOriginTlsClientAuthListResponseResultInfoJSON `json:"-"`
}

// zoneOriginTlsClientAuthListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneOriginTlsClientAuthListResponseResultInfo]
type zoneOriginTlsClientAuthListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOriginTlsClientAuthListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOriginTlsClientAuthListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZoneOriginTlsClientAuthUploadParams struct {
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r ZoneOriginTlsClientAuthUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
