// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountPageProjectDomainService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountPageProjectDomainService] method instead.
type AccountPageProjectDomainService struct {
	Options []option.RequestOption
}

// NewAccountPageProjectDomainService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountPageProjectDomainService(opts ...option.RequestOption) (r *AccountPageProjectDomainService) {
	r = &AccountPageProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *AccountPageProjectDomainService) New(ctx context.Context, accountID string, projectName string, body AccountPageProjectDomainNewParams, opts ...option.RequestOption) (res *ResponseSingleDomain, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single domain.
func (r *AccountPageProjectDomainService) Get(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *ResponseSingleDomain, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retry the validation status of a single domain.
func (r *AccountPageProjectDomainService) Update(ctx context.Context, accountID string, projectName string, domainName string, body AccountPageProjectDomainUpdateParams, opts ...option.RequestOption) (res *ResponseSingleDomain, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *AccountPageProjectDomainService) List(ctx context.Context, accountID string, projectName string, opts ...option.RequestOption) (res *AccountPageProjectDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", accountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a Pages project's domain.
func (r *AccountPageProjectDomainService) Delete(ctx context.Context, accountID string, projectName string, domainName string, body AccountPageProjectDomainDeleteParams, opts ...option.RequestOption) (res *AccountPageProjectDomainDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", accountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type DomainProject struct {
	ID                   string                            `json:"id"`
	CertificateAuthority DomainProjectCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                            `json:"created_on"`
	DomainID             string                            `json:"domain_id"`
	Name                 string                            `json:"name"`
	Status               DomainProjectStatus               `json:"status"`
	ValidationData       DomainProjectValidationData       `json:"validation_data"`
	VerificationData     DomainProjectVerificationData     `json:"verification_data"`
	ZoneTag              string                            `json:"zone_tag"`
	JSON                 domainProjectJSON                 `json:"-"`
}

// domainProjectJSON contains the JSON metadata for the struct [DomainProject]
type domainProjectJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DomainProject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainProjectJSON) RawJSON() string {
	return r.raw
}

type DomainProjectCertificateAuthority string

const (
	DomainProjectCertificateAuthorityGoogle      DomainProjectCertificateAuthority = "google"
	DomainProjectCertificateAuthorityLetsEncrypt DomainProjectCertificateAuthority = "lets_encrypt"
)

func (r DomainProjectCertificateAuthority) IsKnown() bool {
	switch r {
	case DomainProjectCertificateAuthorityGoogle, DomainProjectCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type DomainProjectStatus string

const (
	DomainProjectStatusInitializing DomainProjectStatus = "initializing"
	DomainProjectStatusPending      DomainProjectStatus = "pending"
	DomainProjectStatusActive       DomainProjectStatus = "active"
	DomainProjectStatusDeactivated  DomainProjectStatus = "deactivated"
	DomainProjectStatusBlocked      DomainProjectStatus = "blocked"
	DomainProjectStatusError        DomainProjectStatus = "error"
)

func (r DomainProjectStatus) IsKnown() bool {
	switch r {
	case DomainProjectStatusInitializing, DomainProjectStatusPending, DomainProjectStatusActive, DomainProjectStatusDeactivated, DomainProjectStatusBlocked, DomainProjectStatusError:
		return true
	}
	return false
}

type DomainProjectValidationData struct {
	ErrorMessage string                            `json:"error_message"`
	Method       DomainProjectValidationDataMethod `json:"method"`
	Status       DomainProjectValidationDataStatus `json:"status"`
	TxtName      string                            `json:"txt_name"`
	TxtValue     string                            `json:"txt_value"`
	JSON         domainProjectValidationDataJSON   `json:"-"`
}

// domainProjectValidationDataJSON contains the JSON metadata for the struct
// [DomainProjectValidationData]
type domainProjectValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TxtName      apijson.Field
	TxtValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DomainProjectValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainProjectValidationDataJSON) RawJSON() string {
	return r.raw
}

type DomainProjectValidationDataMethod string

const (
	DomainProjectValidationDataMethodHTTP DomainProjectValidationDataMethod = "http"
	DomainProjectValidationDataMethodTxt  DomainProjectValidationDataMethod = "txt"
)

func (r DomainProjectValidationDataMethod) IsKnown() bool {
	switch r {
	case DomainProjectValidationDataMethodHTTP, DomainProjectValidationDataMethodTxt:
		return true
	}
	return false
}

type DomainProjectValidationDataStatus string

const (
	DomainProjectValidationDataStatusInitializing DomainProjectValidationDataStatus = "initializing"
	DomainProjectValidationDataStatusPending      DomainProjectValidationDataStatus = "pending"
	DomainProjectValidationDataStatusActive       DomainProjectValidationDataStatus = "active"
	DomainProjectValidationDataStatusDeactivated  DomainProjectValidationDataStatus = "deactivated"
	DomainProjectValidationDataStatusError        DomainProjectValidationDataStatus = "error"
)

func (r DomainProjectValidationDataStatus) IsKnown() bool {
	switch r {
	case DomainProjectValidationDataStatusInitializing, DomainProjectValidationDataStatusPending, DomainProjectValidationDataStatusActive, DomainProjectValidationDataStatusDeactivated, DomainProjectValidationDataStatusError:
		return true
	}
	return false
}

type DomainProjectVerificationData struct {
	ErrorMessage string                              `json:"error_message"`
	Status       DomainProjectVerificationDataStatus `json:"status"`
	JSON         domainProjectVerificationDataJSON   `json:"-"`
}

// domainProjectVerificationDataJSON contains the JSON metadata for the struct
// [DomainProjectVerificationData]
type domainProjectVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DomainProjectVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainProjectVerificationDataJSON) RawJSON() string {
	return r.raw
}

type DomainProjectVerificationDataStatus string

const (
	DomainProjectVerificationDataStatusPending     DomainProjectVerificationDataStatus = "pending"
	DomainProjectVerificationDataStatusActive      DomainProjectVerificationDataStatus = "active"
	DomainProjectVerificationDataStatusDeactivated DomainProjectVerificationDataStatus = "deactivated"
	DomainProjectVerificationDataStatusBlocked     DomainProjectVerificationDataStatus = "blocked"
	DomainProjectVerificationDataStatusError       DomainProjectVerificationDataStatus = "error"
)

func (r DomainProjectVerificationDataStatus) IsKnown() bool {
	switch r {
	case DomainProjectVerificationDataStatusPending, DomainProjectVerificationDataStatusActive, DomainProjectVerificationDataStatusDeactivated, DomainProjectVerificationDataStatusBlocked, DomainProjectVerificationDataStatusError:
		return true
	}
	return false
}

type ResponseSingleDomain struct {
	Result ResponseSingleDomainResult `json:"result,nullable"`
	JSON   responseSingleDomainJSON   `json:"-"`
	APIResponsePages
}

// responseSingleDomainJSON contains the JSON metadata for the struct
// [ResponseSingleDomain]
type responseSingleDomainJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainResult struct {
	ID                   string                                         `json:"id"`
	CertificateAuthority ResponseSingleDomainResultCertificateAuthority `json:"certificate_authority"`
	CreatedOn            string                                         `json:"created_on"`
	DomainID             string                                         `json:"domain_id"`
	Name                 string                                         `json:"name"`
	Status               ResponseSingleDomainResultStatus               `json:"status"`
	ValidationData       ResponseSingleDomainResultValidationData       `json:"validation_data"`
	VerificationData     ResponseSingleDomainResultVerificationData     `json:"verification_data"`
	ZoneTag              string                                         `json:"zone_tag"`
	JSON                 responseSingleDomainResultJSON                 `json:"-"`
}

// responseSingleDomainResultJSON contains the JSON metadata for the struct
// [ResponseSingleDomainResult]
type responseSingleDomainResultJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ResponseSingleDomainResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainResultJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainResultCertificateAuthority string

const (
	ResponseSingleDomainResultCertificateAuthorityGoogle      ResponseSingleDomainResultCertificateAuthority = "google"
	ResponseSingleDomainResultCertificateAuthorityLetsEncrypt ResponseSingleDomainResultCertificateAuthority = "lets_encrypt"
)

func (r ResponseSingleDomainResultCertificateAuthority) IsKnown() bool {
	switch r {
	case ResponseSingleDomainResultCertificateAuthorityGoogle, ResponseSingleDomainResultCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ResponseSingleDomainResultStatus string

const (
	ResponseSingleDomainResultStatusInitializing ResponseSingleDomainResultStatus = "initializing"
	ResponseSingleDomainResultStatusPending      ResponseSingleDomainResultStatus = "pending"
	ResponseSingleDomainResultStatusActive       ResponseSingleDomainResultStatus = "active"
	ResponseSingleDomainResultStatusDeactivated  ResponseSingleDomainResultStatus = "deactivated"
	ResponseSingleDomainResultStatusBlocked      ResponseSingleDomainResultStatus = "blocked"
	ResponseSingleDomainResultStatusError        ResponseSingleDomainResultStatus = "error"
)

func (r ResponseSingleDomainResultStatus) IsKnown() bool {
	switch r {
	case ResponseSingleDomainResultStatusInitializing, ResponseSingleDomainResultStatusPending, ResponseSingleDomainResultStatusActive, ResponseSingleDomainResultStatusDeactivated, ResponseSingleDomainResultStatusBlocked, ResponseSingleDomainResultStatusError:
		return true
	}
	return false
}

type ResponseSingleDomainResultValidationData struct {
	ErrorMessage string                                         `json:"error_message"`
	Method       ResponseSingleDomainResultValidationDataMethod `json:"method"`
	Status       ResponseSingleDomainResultValidationDataStatus `json:"status"`
	TxtName      string                                         `json:"txt_name"`
	TxtValue     string                                         `json:"txt_value"`
	JSON         responseSingleDomainResultValidationDataJSON   `json:"-"`
}

// responseSingleDomainResultValidationDataJSON contains the JSON metadata for the
// struct [ResponseSingleDomainResultValidationData]
type responseSingleDomainResultValidationDataJSON struct {
	ErrorMessage apijson.Field
	Method       apijson.Field
	Status       apijson.Field
	TxtName      apijson.Field
	TxtValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseSingleDomainResultValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainResultValidationDataJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainResultValidationDataMethod string

const (
	ResponseSingleDomainResultValidationDataMethodHTTP ResponseSingleDomainResultValidationDataMethod = "http"
	ResponseSingleDomainResultValidationDataMethodTxt  ResponseSingleDomainResultValidationDataMethod = "txt"
)

func (r ResponseSingleDomainResultValidationDataMethod) IsKnown() bool {
	switch r {
	case ResponseSingleDomainResultValidationDataMethodHTTP, ResponseSingleDomainResultValidationDataMethodTxt:
		return true
	}
	return false
}

type ResponseSingleDomainResultValidationDataStatus string

const (
	ResponseSingleDomainResultValidationDataStatusInitializing ResponseSingleDomainResultValidationDataStatus = "initializing"
	ResponseSingleDomainResultValidationDataStatusPending      ResponseSingleDomainResultValidationDataStatus = "pending"
	ResponseSingleDomainResultValidationDataStatusActive       ResponseSingleDomainResultValidationDataStatus = "active"
	ResponseSingleDomainResultValidationDataStatusDeactivated  ResponseSingleDomainResultValidationDataStatus = "deactivated"
	ResponseSingleDomainResultValidationDataStatusError        ResponseSingleDomainResultValidationDataStatus = "error"
)

func (r ResponseSingleDomainResultValidationDataStatus) IsKnown() bool {
	switch r {
	case ResponseSingleDomainResultValidationDataStatusInitializing, ResponseSingleDomainResultValidationDataStatusPending, ResponseSingleDomainResultValidationDataStatusActive, ResponseSingleDomainResultValidationDataStatusDeactivated, ResponseSingleDomainResultValidationDataStatusError:
		return true
	}
	return false
}

type ResponseSingleDomainResultVerificationData struct {
	ErrorMessage string                                           `json:"error_message"`
	Status       ResponseSingleDomainResultVerificationDataStatus `json:"status"`
	JSON         responseSingleDomainResultVerificationDataJSON   `json:"-"`
}

// responseSingleDomainResultVerificationDataJSON contains the JSON metadata for
// the struct [ResponseSingleDomainResultVerificationData]
type responseSingleDomainResultVerificationDataJSON struct {
	ErrorMessage apijson.Field
	Status       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseSingleDomainResultVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainResultVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainResultVerificationDataStatus string

const (
	ResponseSingleDomainResultVerificationDataStatusPending     ResponseSingleDomainResultVerificationDataStatus = "pending"
	ResponseSingleDomainResultVerificationDataStatusActive      ResponseSingleDomainResultVerificationDataStatus = "active"
	ResponseSingleDomainResultVerificationDataStatusDeactivated ResponseSingleDomainResultVerificationDataStatus = "deactivated"
	ResponseSingleDomainResultVerificationDataStatusBlocked     ResponseSingleDomainResultVerificationDataStatus = "blocked"
	ResponseSingleDomainResultVerificationDataStatusError       ResponseSingleDomainResultVerificationDataStatus = "error"
)

func (r ResponseSingleDomainResultVerificationDataStatus) IsKnown() bool {
	switch r {
	case ResponseSingleDomainResultVerificationDataStatusPending, ResponseSingleDomainResultVerificationDataStatusActive, ResponseSingleDomainResultVerificationDataStatusDeactivated, ResponseSingleDomainResultVerificationDataStatusBlocked, ResponseSingleDomainResultVerificationDataStatusError:
		return true
	}
	return false
}

type AccountPageProjectDomainListResponse struct {
	Result []DomainProject                          `json:"result"`
	JSON   accountPageProjectDomainListResponseJSON `json:"-"`
	APIResponsePages
	APIResponsePagination
}

// accountPageProjectDomainListResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainListResponse]
type accountPageProjectDomainListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainDeleteResponse struct {
	Result interface{}                                `json:"result,nullable"`
	JSON   accountPageProjectDomainDeleteResponseJSON `json:"-"`
	APIResponsePages
}

// accountPageProjectDomainDeleteResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainDeleteResponse]
type accountPageProjectDomainDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainNewParams struct {
	Name param.Field[string] `json:"name"`
}

func (r AccountPageProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountPageProjectDomainUpdateParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountPageProjectDomainDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountPageProjectDomainDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
