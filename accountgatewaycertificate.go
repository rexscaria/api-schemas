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

// AccountGatewayCertificateService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountGatewayCertificateService] method instead.
type AccountGatewayCertificateService struct {
	Options []option.RequestOption
}

// NewAccountGatewayCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayCertificateService(opts ...option.RequestOption) (r *AccountGatewayCertificateService) {
	r = &AccountGatewayCertificateService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust certificate.
func (r *AccountGatewayCertificateService) New(ctx context.Context, accountID string, body AccountGatewayCertificateNewParams, opts ...option.RequestOption) (res *SingleResponseCertificateGateway, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust certificate.
func (r *AccountGatewayCertificateService) Get(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *SingleResponseCertificateGateway, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches all Zero Trust certificates for an account.
func (r *AccountGatewayCertificateService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountGatewayCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a gateway-managed Zero Trust certificate. A certificate must be
// deactivated from the edge (inactive) before it is deleted.
func (r *AccountGatewayCertificateService) Delete(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *SingleResponseCertificateGateway, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Binds a single Zero Trust certificate to the edge.
func (r *AccountGatewayCertificateService) Activate(ctx context.Context, accountID string, certificateID string, body AccountGatewayCertificateActivateParams, opts ...option.RequestOption) (res *SingleResponseCertificateGateway, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s/activate", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unbinds a single Zero Trust certificate from the edge
func (r *AccountGatewayCertificateService) Deactivate(ctx context.Context, accountID string, certificateID string, body AccountGatewayCertificateDeactivateParams, opts ...option.RequestOption) (res *SingleResponseCertificateGateway, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/certificates/%s/deactivate", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Certificate struct {
	// Certificate UUID tag.
	ID string `json:"id"`
	// The deployment status of the certificate on Cloudflare's edge. Certificates in
	// the 'available' (previously called 'active') state may be used for Gateway TLS
	// interception.
	BindingStatus CertificateBindingStatus `json:"binding_status"`
	// The CA certificate
	Certificate string    `json:"certificate"`
	CreatedAt   time.Time `json:"created_at" format:"date-time"`
	ExpiresOn   time.Time `json:"expires_on" format:"date-time"`
	// The SHA256 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// Use this certificate for Gateway TLS interception
	InUse bool `json:"in_use"`
	// The organization that issued the certificate.
	IssuerOrg string `json:"issuer_org"`
	// The entire issuer field of the certificate.
	IssuerRaw string `json:"issuer_raw"`
	// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
	Type       CertificateType `json:"type"`
	UpdatedAt  time.Time       `json:"updated_at" format:"date-time"`
	UploadedOn time.Time       `json:"uploaded_on" format:"date-time"`
	JSON       certificateJSON `json:"-"`
}

// certificateJSON contains the JSON metadata for the struct [Certificate]
type certificateJSON struct {
	ID            apijson.Field
	BindingStatus apijson.Field
	Certificate   apijson.Field
	CreatedAt     apijson.Field
	ExpiresOn     apijson.Field
	Fingerprint   apijson.Field
	InUse         apijson.Field
	IssuerOrg     apijson.Field
	IssuerRaw     apijson.Field
	Type          apijson.Field
	UpdatedAt     apijson.Field
	UploadedOn    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Certificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateJSON) RawJSON() string {
	return r.raw
}

// The deployment status of the certificate on Cloudflare's edge. Certificates in
// the 'available' (previously called 'active') state may be used for Gateway TLS
// interception.
type CertificateBindingStatus string

const (
	CertificateBindingStatusPendingDeployment CertificateBindingStatus = "pending_deployment"
	CertificateBindingStatusAvailable         CertificateBindingStatus = "available"
	CertificateBindingStatusPendingDeletion   CertificateBindingStatus = "pending_deletion"
	CertificateBindingStatusInactive          CertificateBindingStatus = "inactive"
)

func (r CertificateBindingStatus) IsKnown() bool {
	switch r {
	case CertificateBindingStatusPendingDeployment, CertificateBindingStatusAvailable, CertificateBindingStatusPendingDeletion, CertificateBindingStatusInactive:
		return true
	}
	return false
}

// The type of certificate, either BYO-PKI (custom) or Gateway-managed.
type CertificateType string

const (
	CertificateTypeCustom         CertificateType = "custom"
	CertificateTypeGatewayManaged CertificateType = "gateway_managed"
)

func (r CertificateType) IsKnown() bool {
	switch r {
	case CertificateTypeCustom, CertificateTypeGatewayManaged:
		return true
	}
	return false
}

type SingleResponseCertificateGateway struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success SingleResponseCertificateGatewaySuccess `json:"success,required"`
	Result  Certificate                             `json:"result"`
	JSON    singleResponseCertificateGatewayJSON    `json:"-"`
}

// singleResponseCertificateGatewayJSON contains the JSON metadata for the struct
// [SingleResponseCertificateGateway]
type singleResponseCertificateGatewayJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCertificateGateway) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseCertificateGatewayJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SingleResponseCertificateGatewaySuccess bool

const (
	SingleResponseCertificateGatewaySuccessTrue SingleResponseCertificateGatewaySuccess = true
)

func (r SingleResponseCertificateGatewaySuccess) IsKnown() bool {
	switch r {
	case SingleResponseCertificateGatewaySuccessTrue:
		return true
	}
	return false
}

type AccountGatewayCertificateListResponse struct {
	Errors   []ZeroTrustGatewayMessages `json:"errors,required"`
	Messages []ZeroTrustGatewayMessages `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountGatewayCertificateListResponseSuccess    `json:"success,required"`
	Result     []Certificate                                   `json:"result"`
	ResultInfo AccountGatewayCertificateListResponseResultInfo `json:"result_info"`
	JSON       accountGatewayCertificateListResponseJSON       `json:"-"`
}

// accountGatewayCertificateListResponseJSON contains the JSON metadata for the
// struct [AccountGatewayCertificateListResponse]
type accountGatewayCertificateListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountGatewayCertificateListResponseSuccess bool

const (
	AccountGatewayCertificateListResponseSuccessTrue AccountGatewayCertificateListResponseSuccess = true
)

func (r AccountGatewayCertificateListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountGatewayCertificateListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountGatewayCertificateListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       accountGatewayCertificateListResponseResultInfoJSON `json:"-"`
}

// accountGatewayCertificateListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountGatewayCertificateListResponseResultInfo]
type accountGatewayCertificateListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayCertificateListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountGatewayCertificateListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountGatewayCertificateNewParams struct {
	// Number of days the generated certificate will be valid, minimum 1 day and
	// maximum 30 years. Defaults to 5 years. In terraform, validity_period_days can
	// only be used while creating a certificate, and this CAN NOT be used to extend
	// the validity of an already generated certificate.
	ValidityPeriodDays param.Field[int64] `json:"validity_period_days"`
}

func (r AccountGatewayCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayCertificateActivateParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGatewayCertificateActivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountGatewayCertificateDeactivateParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountGatewayCertificateDeactivateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
