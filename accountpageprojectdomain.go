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
func (r *AccountPageProjectDomainService) Delete(ctx context.Context, accountID string, projectName string, domainName string, opts ...option.RequestOption) (res *AccountPageProjectDomainDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	Errors   []ResponseSingleDomainError   `json:"errors,required"`
	Messages []ResponseSingleDomainMessage `json:"messages,required"`
	Result   DomainProject                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success ResponseSingleDomainSuccess `json:"success,required"`
	JSON    responseSingleDomainJSON    `json:"-"`
}

// responseSingleDomainJSON contains the JSON metadata for the struct
// [ResponseSingleDomain]
type responseSingleDomainJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainError struct {
	Code             int64                            `json:"code,required"`
	Message          string                           `json:"message,required"`
	DocumentationURL string                           `json:"documentation_url"`
	Source           ResponseSingleDomainErrorsSource `json:"source"`
	JSON             responseSingleDomainErrorJSON    `json:"-"`
}

// responseSingleDomainErrorJSON contains the JSON metadata for the struct
// [ResponseSingleDomainError]
type responseSingleDomainErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseSingleDomainError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainErrorJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainErrorsSource struct {
	Pointer string                               `json:"pointer"`
	JSON    responseSingleDomainErrorsSourceJSON `json:"-"`
}

// responseSingleDomainErrorsSourceJSON contains the JSON metadata for the struct
// [ResponseSingleDomainErrorsSource]
type responseSingleDomainErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleDomainErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainMessage struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           ResponseSingleDomainMessagesSource `json:"source"`
	JSON             responseSingleDomainMessageJSON    `json:"-"`
}

// responseSingleDomainMessageJSON contains the JSON metadata for the struct
// [ResponseSingleDomainMessage]
type responseSingleDomainMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseSingleDomainMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainMessageJSON) RawJSON() string {
	return r.raw
}

type ResponseSingleDomainMessagesSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    responseSingleDomainMessagesSourceJSON `json:"-"`
}

// responseSingleDomainMessagesSourceJSON contains the JSON metadata for the struct
// [ResponseSingleDomainMessagesSource]
type responseSingleDomainMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleDomainMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseSingleDomainMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ResponseSingleDomainSuccess bool

const (
	ResponseSingleDomainSuccessFalse ResponseSingleDomainSuccess = false
	ResponseSingleDomainSuccessTrue  ResponseSingleDomainSuccess = true
)

func (r ResponseSingleDomainSuccess) IsKnown() bool {
	switch r {
	case ResponseSingleDomainSuccessFalse, ResponseSingleDomainSuccessTrue:
		return true
	}
	return false
}

type AccountPageProjectDomainListResponse struct {
	Errors   []AccountPageProjectDomainListResponseError   `json:"errors,required"`
	Messages []AccountPageProjectDomainListResponseMessage `json:"messages,required"`
	Result   []DomainProject                               `json:"result,required"`
	// Whether the API call was successful
	Success    AccountPageProjectDomainListResponseSuccess    `json:"success,required"`
	ResultInfo AccountPageProjectDomainListResponseResultInfo `json:"result_info"`
	JSON       accountPageProjectDomainListResponseJSON       `json:"-"`
}

// accountPageProjectDomainListResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainListResponse]
type accountPageProjectDomainListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainListResponseError struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountPageProjectDomainListResponseErrorsSource `json:"source"`
	JSON             accountPageProjectDomainListResponseErrorJSON    `json:"-"`
}

// accountPageProjectDomainListResponseErrorJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainListResponseError]
type accountPageProjectDomainListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainListResponseErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountPageProjectDomainListResponseErrorsSourceJSON `json:"-"`
}

// accountPageProjectDomainListResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountPageProjectDomainListResponseErrorsSource]
type accountPageProjectDomainListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainListResponseMessage struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccountPageProjectDomainListResponseMessagesSource `json:"source"`
	JSON             accountPageProjectDomainListResponseMessageJSON    `json:"-"`
}

// accountPageProjectDomainListResponseMessageJSON contains the JSON metadata for
// the struct [AccountPageProjectDomainListResponseMessage]
type accountPageProjectDomainListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainListResponseMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accountPageProjectDomainListResponseMessagesSourceJSON `json:"-"`
}

// accountPageProjectDomainListResponseMessagesSourceJSON contains the JSON
// metadata for the struct [AccountPageProjectDomainListResponseMessagesSource]
type accountPageProjectDomainListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPageProjectDomainListResponseSuccess bool

const (
	AccountPageProjectDomainListResponseSuccessFalse AccountPageProjectDomainListResponseSuccess = false
	AccountPageProjectDomainListResponseSuccessTrue  AccountPageProjectDomainListResponseSuccess = true
)

func (r AccountPageProjectDomainListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPageProjectDomainListResponseSuccessFalse, AccountPageProjectDomainListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountPageProjectDomainListResponseResultInfo struct {
	// The number of items on the current page.
	Count int64 `json:"count,required"`
	// The page currently being requested.
	Page int64 `json:"page,required"`
	// The number of items per page being returned.
	PerPage int64 `json:"per_page,required"`
	// The total count of items.
	TotalCount int64 `json:"total_count,required"`
	// The total count of pages.
	TotalPages int64                                              `json:"total_pages"`
	JSON       accountPageProjectDomainListResponseResultInfoJSON `json:"-"`
}

// accountPageProjectDomainListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountPageProjectDomainListResponseResultInfo]
type accountPageProjectDomainListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountPageProjectDomainDeleteResponse struct {
	Errors   []MessagesPageItem `json:"errors,required"`
	Messages []MessagesPageItem `json:"messages,required"`
	Result   interface{}        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AccountPageProjectDomainDeleteResponseSuccess `json:"success,required"`
	JSON    accountPageProjectDomainDeleteResponseJSON    `json:"-"`
}

// accountPageProjectDomainDeleteResponseJSON contains the JSON metadata for the
// struct [AccountPageProjectDomainDeleteResponse]
type accountPageProjectDomainDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountPageProjectDomainDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountPageProjectDomainDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountPageProjectDomainDeleteResponseSuccess bool

const (
	AccountPageProjectDomainDeleteResponseSuccessFalse AccountPageProjectDomainDeleteResponseSuccess = false
	AccountPageProjectDomainDeleteResponseSuccessTrue  AccountPageProjectDomainDeleteResponseSuccess = true
)

func (r AccountPageProjectDomainDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountPageProjectDomainDeleteResponseSuccessFalse, AccountPageProjectDomainDeleteResponseSuccessTrue:
		return true
	}
	return false
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
