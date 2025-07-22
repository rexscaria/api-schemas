// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAccessCertificateService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessCertificateService] method instead.
type AccountAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *AccountAccessCertificateSettingService
}

// NewAccountAccessCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateService(opts ...option.RequestOption) (r *AccountAccessCertificateService) {
	r = &AccountAccessCertificateService{}
	r.Options = opts
	r.Settings = NewAccountAccessCertificateSettingService(opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccountAccessCertificateService) New(ctx context.Context, accountID string, body AccountAccessCertificateNewParams, opts ...option.RequestOption) (res *SingleResponseCertificate, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *AccountAccessCertificateService) Get(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *SingleResponseCertificate, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *AccountAccessCertificateService) Update(ctx context.Context, accountID string, certificateID string, body AccountAccessCertificateUpdateParams, opts ...option.RequestOption) (res *SingleResponseCertificate, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all mTLS root certificates.
func (r *AccountAccessCertificateService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAccessCertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *AccountAccessCertificateService) Delete(ctx context.Context, accountID string, certificateID string, opts ...option.RequestOption) (res *IDResponseCertificates, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", accountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseAccess struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseAccessSuccess `json:"success,required"`
	JSON    apiResponseAccessJSON    `json:"-"`
}

// apiResponseAccessJSON contains the JSON metadata for the struct
// [APIResponseAccess]
type apiResponseAccessJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAccessJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseAccessSuccess bool

const (
	APIResponseAccessSuccessTrue APIResponseAccessSuccess = true
)

func (r APIResponseAccessSuccess) IsKnown() bool {
	switch r {
	case APIResponseAccessSuccessTrue:
		return true
	}
	return false
}

type Certificates struct {
	// The ID of the application that will use this certificate.
	ID string `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string           `json:"name"`
	UpdatedAt time.Time        `json:"updated_at" format:"date-time"`
	JSON      certificatesJSON `json:"-"`
}

// certificatesJSON contains the JSON metadata for the struct [Certificates]
type certificatesJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Certificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificatesJSON) RawJSON() string {
	return r.raw
}

type IDResponseCertificates struct {
	Result IDResponseCertificatesResult `json:"result"`
	JSON   idResponseCertificatesJSON   `json:"-"`
	APIResponseAccess
}

// idResponseCertificatesJSON contains the JSON metadata for the struct
// [IDResponseCertificates]
type idResponseCertificatesJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseCertificatesJSON) RawJSON() string {
	return r.raw
}

type IDResponseCertificatesResult struct {
	// UUID
	ID   string                           `json:"id"`
	JSON idResponseCertificatesResultJSON `json:"-"`
}

// idResponseCertificatesResultJSON contains the JSON metadata for the struct
// [IDResponseCertificatesResult]
type idResponseCertificatesResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IDResponseCertificatesResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r idResponseCertificatesResultJSON) RawJSON() string {
	return r.raw
}

type MessagesAccessItem struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    messagesAccessItemJSON `json:"-"`
}

// messagesAccessItemJSON contains the JSON metadata for the struct
// [MessagesAccessItem]
type messagesAccessItemJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessagesAccessItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messagesAccessItemJSON) RawJSON() string {
	return r.raw
}

type SingleResponseCertificate struct {
	Result Certificates                  `json:"result"`
	JSON   singleResponseCertificateJSON `json:"-"`
	APIResponseSingleAccess
}

// singleResponseCertificateJSON contains the JSON metadata for the struct
// [SingleResponseCertificate]
type singleResponseCertificateJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r singleResponseCertificateJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCertificateListResponse struct {
	Result []Certificates                           `json:"result"`
	JSON   accountAccessCertificateListResponseJSON `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessCertificateListResponseJSON contains the JSON metadata for the
// struct [AccountAccessCertificateListResponse]
type accountAccessCertificateListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessCertificateNewParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccountAccessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
