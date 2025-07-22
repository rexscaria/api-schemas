// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// ZoneClientCertificateService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneClientCertificateService] method instead.
type ZoneClientCertificateService struct {
	Options []option.RequestOption
}

// NewZoneClientCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneClientCertificateService(opts ...option.RequestOption) (r *ZoneClientCertificateService) {
	r = &ZoneClientCertificateService{}
	r.Options = opts
	return
}

// Create a new API Shield mTLS Client Certificate
func (r *ZoneClientCertificateService) New(ctx context.Context, zoneID string, body ZoneClientCertificateNewParams, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/client_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Details for a single mTLS API Shield Client Certificate
func (r *ZoneClientCertificateService) Get(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if clientCertificateID == "" {
		err = errors.New("missing required client_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all of your Zone's API Shield mTLS Client Certificates by Status and/or
// using Pagination
func (r *ZoneClientCertificateService) List(ctx context.Context, zoneID string, query ZoneClientCertificateListParams, opts ...option.RequestOption) (res *ZoneClientCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/client_certificates", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// If a API Shield mTLS Client Certificate is in a pending_revocation state, you
// may reactivate it with this endpoint.
func (r *ZoneClientCertificateService) Reactivate(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if clientCertificateID == "" {
		err = errors.New("missing required client_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Set a API Shield mTLS Client Certificate to pending_revocation status for
// processing to revoked status.
func (r *ZoneClientCertificateService) Revoke(ctx context.Context, zoneID string, clientCertificateID string, opts ...option.RequestOption) (res *ClientCertificateResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if zoneID == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if clientCertificateID == "" {
		err = errors.New("missing required client_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/client_certificates/%s", zoneID, clientCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ClientCertificate struct {
	// Identifier
	ID string `json:"id"`
	// The Client Certificate PEM
	Certificate string `json:"certificate"`
	// Certificate Authority used to issue the Client Certificate
	CertificateAuthority ClientCertificateCertificateAuthority `json:"certificate_authority"`
	// Common Name of the Client Certificate
	CommonName string `json:"common_name"`
	// Country, provided by the CSR
	Country string `json:"country"`
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr string `json:"csr"`
	// Date that the Client Certificate expires
	ExpiresOn string `json:"expires_on"`
	// Unique identifier of the Client Certificate
	FingerprintSha256 string `json:"fingerprint_sha256"`
	// Date that the Client Certificate was issued by the Certificate Authority
	IssuedOn string `json:"issued_on"`
	// Location, provided by the CSR
	Location string `json:"location"`
	// Organization, provided by the CSR
	Organization string `json:"organization"`
	// Organizational Unit, provided by the CSR
	OrganizationalUnit string `json:"organizational_unit"`
	// The serial number on the created Client Certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the Client Certificate..
	Signature string `json:"signature"`
	// Subject Key Identifier
	Ski string `json:"ski"`
	// State, provided by the CSR
	State string `json:"state"`
	// Client Certificates may be active or revoked, and the pending_reactivation or
	// pending_revocation represent in-progress asynchronous transitions
	Status ClientCertificateStatus `json:"status"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays int64                 `json:"validity_days"`
	JSON         clientCertificateJSON `json:"-"`
}

// clientCertificateJSON contains the JSON metadata for the struct
// [ClientCertificate]
type clientCertificateJSON struct {
	ID                   apijson.Field
	Certificate          apijson.Field
	CertificateAuthority apijson.Field
	CommonName           apijson.Field
	Country              apijson.Field
	Csr                  apijson.Field
	ExpiresOn            apijson.Field
	FingerprintSha256    apijson.Field
	IssuedOn             apijson.Field
	Location             apijson.Field
	Organization         apijson.Field
	OrganizationalUnit   apijson.Field
	SerialNumber         apijson.Field
	Signature            apijson.Field
	Ski                  apijson.Field
	State                apijson.Field
	Status               apijson.Field
	ValidityDays         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ClientCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateJSON) RawJSON() string {
	return r.raw
}

// Certificate Authority used to issue the Client Certificate
type ClientCertificateCertificateAuthority struct {
	ID   string                                    `json:"id"`
	Name string                                    `json:"name"`
	JSON clientCertificateCertificateAuthorityJSON `json:"-"`
}

// clientCertificateCertificateAuthorityJSON contains the JSON metadata for the
// struct [ClientCertificateCertificateAuthority]
type clientCertificateCertificateAuthorityJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Client Certificates may be active or revoked, and the pending_reactivation or
// pending_revocation represent in-progress asynchronous transitions
type ClientCertificateStatus string

const (
	ClientCertificateStatusActive              ClientCertificateStatus = "active"
	ClientCertificateStatusPendingReactivation ClientCertificateStatus = "pending_reactivation"
	ClientCertificateStatusPendingRevocation   ClientCertificateStatus = "pending_revocation"
	ClientCertificateStatusRevoked             ClientCertificateStatus = "revoked"
)

func (r ClientCertificateStatus) IsKnown() bool {
	switch r {
	case ClientCertificateStatusActive, ClientCertificateStatusPendingReactivation, ClientCertificateStatusPendingRevocation, ClientCertificateStatusRevoked:
		return true
	}
	return false
}

type ClientCertificateResponseSingle struct {
	Result ClientCertificate                   `json:"result"`
	JSON   clientCertificateResponseSingleJSON `json:"-"`
	APIResponseSingleTlsCertificates
}

// clientCertificateResponseSingleJSON contains the JSON metadata for the struct
// [ClientCertificateResponseSingle]
type clientCertificateResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ClientCertificateResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientCertificateResponseSingleJSON) RawJSON() string {
	return r.raw
}

type ZoneClientCertificateListResponse struct {
	Result []ClientCertificate                   `json:"result"`
	JSON   zoneClientCertificateListResponseJSON `json:"-"`
	APIResponseCollectionTlsCertificates
}

// zoneClientCertificateListResponseJSON contains the JSON metadata for the struct
// [ZoneClientCertificateListResponse]
type zoneClientCertificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneClientCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneClientCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneClientCertificateNewParams struct {
	// The Certificate Signing Request (CSR). Must be newline-encoded.
	Csr param.Field[string] `json:"csr,required"`
	// The number of days the Client Certificate will be valid after the issued_on date
	ValidityDays param.Field[int64] `json:"validity_days,required"`
}

func (r ZoneClientCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneClientCertificateListParams struct {
	// Limit to the number of records returned.
	Limit param.Field[int64] `query:"limit"`
	// Offset the results
	Offset param.Field[int64] `query:"offset"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Client Certitifcate Status to filter results by.
	Status param.Field[ZoneClientCertificateListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneClientCertificateListParams]'s query parameters as
// `url.Values`.
func (r ZoneClientCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Client Certitifcate Status to filter results by.
type ZoneClientCertificateListParamsStatus string

const (
	ZoneClientCertificateListParamsStatusAll                 ZoneClientCertificateListParamsStatus = "all"
	ZoneClientCertificateListParamsStatusActive              ZoneClientCertificateListParamsStatus = "active"
	ZoneClientCertificateListParamsStatusPendingReactivation ZoneClientCertificateListParamsStatus = "pending_reactivation"
	ZoneClientCertificateListParamsStatusPendingRevocation   ZoneClientCertificateListParamsStatus = "pending_revocation"
	ZoneClientCertificateListParamsStatusRevoked             ZoneClientCertificateListParamsStatus = "revoked"
)

func (r ZoneClientCertificateListParamsStatus) IsKnown() bool {
	switch r {
	case ZoneClientCertificateListParamsStatusAll, ZoneClientCertificateListParamsStatusActive, ZoneClientCertificateListParamsStatusPendingReactivation, ZoneClientCertificateListParamsStatusPendingRevocation, ZoneClientCertificateListParamsStatusRevoked:
		return true
	}
	return false
}
