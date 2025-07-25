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

// AccountRegistrarDomainService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountRegistrarDomainService] method instead.
type AccountRegistrarDomainService struct {
	Options []option.RequestOption
}

// NewAccountRegistrarDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRegistrarDomainService(opts ...option.RequestOption) (r *AccountRegistrarDomainService) {
	r = &AccountRegistrarDomainService{}
	r.Options = opts
	return
}

// Show individual domain.
func (r *AccountRegistrarDomainService) Get(ctx context.Context, accountID string, domainName string, opts ...option.RequestOption) (res *RegistrarAPIDomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update individual domain.
func (r *AccountRegistrarDomainService) Update(ctx context.Context, accountID string, domainName string, body AccountRegistrarDomainUpdateParams, opts ...option.RequestOption) (res *RegistrarAPIDomainResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", accountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List domains handled by Registrar.
func (r *AccountRegistrarDomainService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountRegistrarDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RegistrarAPIDomainResponseSingle struct {
	Errors   []RegistrarAPIDomainResponseSingleError   `json:"errors,required"`
	Messages []RegistrarAPIDomainResponseSingleMessage `json:"messages,required"`
	Result   interface{}                               `json:"result,required,nullable"`
	// Whether the API call was successful
	Success RegistrarAPIDomainResponseSingleSuccess `json:"success,required"`
	JSON    registrarAPIDomainResponseSingleJSON    `json:"-"`
}

// registrarAPIDomainResponseSingleJSON contains the JSON metadata for the struct
// [RegistrarAPIDomainResponseSingle]
type registrarAPIDomainResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleJSON) RawJSON() string {
	return r.raw
}

type RegistrarAPIDomainResponseSingleError struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RegistrarAPIDomainResponseSingleErrorsSource `json:"source"`
	JSON             registrarAPIDomainResponseSingleErrorJSON    `json:"-"`
}

// registrarAPIDomainResponseSingleErrorJSON contains the JSON metadata for the
// struct [RegistrarAPIDomainResponseSingleError]
type registrarAPIDomainResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type RegistrarAPIDomainResponseSingleErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    registrarAPIDomainResponseSingleErrorsSourceJSON `json:"-"`
}

// registrarAPIDomainResponseSingleErrorsSourceJSON contains the JSON metadata for
// the struct [RegistrarAPIDomainResponseSingleErrorsSource]
type registrarAPIDomainResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegistrarAPIDomainResponseSingleMessage struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           RegistrarAPIDomainResponseSingleMessagesSource `json:"source"`
	JSON             registrarAPIDomainResponseSingleMessageJSON    `json:"-"`
}

// registrarAPIDomainResponseSingleMessageJSON contains the JSON metadata for the
// struct [RegistrarAPIDomainResponseSingleMessage]
type registrarAPIDomainResponseSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

type RegistrarAPIDomainResponseSingleMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    registrarAPIDomainResponseSingleMessagesSourceJSON `json:"-"`
}

// registrarAPIDomainResponseSingleMessagesSourceJSON contains the JSON metadata
// for the struct [RegistrarAPIDomainResponseSingleMessagesSource]
type registrarAPIDomainResponseSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegistrarAPIDomainResponseSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r registrarAPIDomainResponseSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegistrarAPIDomainResponseSingleSuccess bool

const (
	RegistrarAPIDomainResponseSingleSuccessTrue RegistrarAPIDomainResponseSingleSuccess = true
)

func (r RegistrarAPIDomainResponseSingleSuccess) IsKnown() bool {
	switch r {
	case RegistrarAPIDomainResponseSingleSuccessTrue:
		return true
	}
	return false
}

type AccountRegistrarDomainListResponse struct {
	Errors   []AccountRegistrarDomainListResponseError   `json:"errors,required"`
	Messages []AccountRegistrarDomainListResponseMessage `json:"messages,required"`
	Result   []AccountRegistrarDomainListResponseResult  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountRegistrarDomainListResponseSuccess    `json:"success,required"`
	ResultInfo AccountRegistrarDomainListResponseResultInfo `json:"result_info"`
	JSON       accountRegistrarDomainListResponseJSON       `json:"-"`
}

// accountRegistrarDomainListResponseJSON contains the JSON metadata for the struct
// [AccountRegistrarDomainListResponse]
type accountRegistrarDomainListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseError struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccountRegistrarDomainListResponseErrorsSource `json:"source"`
	JSON             accountRegistrarDomainListResponseErrorJSON    `json:"-"`
}

// accountRegistrarDomainListResponseErrorJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainListResponseError]
type accountRegistrarDomainListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accountRegistrarDomainListResponseErrorsSourceJSON `json:"-"`
}

// accountRegistrarDomainListResponseErrorsSourceJSON contains the JSON metadata
// for the struct [AccountRegistrarDomainListResponseErrorsSource]
type accountRegistrarDomainListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseMessage struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccountRegistrarDomainListResponseMessagesSource `json:"source"`
	JSON             accountRegistrarDomainListResponseMessageJSON    `json:"-"`
}

// accountRegistrarDomainListResponseMessageJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainListResponseMessage]
type accountRegistrarDomainListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accountRegistrarDomainListResponseMessagesSourceJSON `json:"-"`
}

// accountRegistrarDomainListResponseMessagesSourceJSON contains the JSON metadata
// for the struct [AccountRegistrarDomainListResponseMessagesSource]
type accountRegistrarDomainListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainListResponseResult struct {
	// Domain identifier.
	ID string `json:"id"`
	// Shows if a domain is available for transferring into Cloudflare Registrar.
	Available bool `json:"available"`
	// Indicates if the domain can be registered as a new domain.
	CanRegister bool `json:"can_register"`
	// Shows time of creation.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Shows name of current registrar.
	CurrentRegistrar string `json:"current_registrar"`
	// Shows when domain name registration expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Shows whether a registrar lock is in place for a domain.
	Locked bool `json:"locked"`
	// Shows contact information for domain registrant.
	RegistrantContact AccountRegistrarDomainListResponseResultRegistrantContact `json:"registrant_contact"`
	// A comma-separated list of registry status codes. A full list of status codes can
	// be found at
	// [EPP Status Codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
	RegistryStatuses string `json:"registry_statuses"`
	// Whether a particular TLD is currently supported by Cloudflare Registrar. Refer
	// to [TLD Policies](https://www.cloudflare.com/tld-policies/) for a list of
	// supported TLDs.
	SupportedTld bool `json:"supported_tld"`
	// Statuses for domain transfers into Cloudflare Registrar.
	TransferIn AccountRegistrarDomainListResponseResultTransferIn `json:"transfer_in"`
	// Last updated.
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      accountRegistrarDomainListResponseResultJSON `json:"-"`
}

// accountRegistrarDomainListResponseResultJSON contains the JSON metadata for the
// struct [AccountRegistrarDomainListResponseResult]
type accountRegistrarDomainListResponseResultJSON struct {
	ID                apijson.Field
	Available         apijson.Field
	CanRegister       apijson.Field
	CreatedAt         apijson.Field
	CurrentRegistrar  apijson.Field
	ExpiresAt         apijson.Field
	Locked            apijson.Field
	RegistrantContact apijson.Field
	RegistryStatuses  apijson.Field
	SupportedTld      apijson.Field
	TransferIn        apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseResultJSON) RawJSON() string {
	return r.raw
}

// Shows contact information for domain registrant.
type AccountRegistrarDomainListResponseResultRegistrantContact struct {
	// Address.
	Address string `json:"address,required"`
	// City.
	City string `json:"city,required"`
	// The country in which the user lives.
	Country string `json:"country,required,nullable"`
	// User's first name
	FirstName string `json:"first_name,required,nullable"`
	// User's last name
	LastName string `json:"last_name,required,nullable"`
	// Name of organization.
	Organization string `json:"organization,required"`
	// User's telephone number
	Phone string `json:"phone,required,nullable"`
	// State.
	State string `json:"state,required"`
	// The zipcode or postal code where the user lives.
	Zip string `json:"zip,required,nullable"`
	// Contact Identifier.
	ID string `json:"id"`
	// Optional address line for unit, floor, suite, etc.
	Address2 string `json:"address2"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Contact fax number.
	Fax  string                                                        `json:"fax"`
	JSON accountRegistrarDomainListResponseResultRegistrantContactJSON `json:"-"`
}

// accountRegistrarDomainListResponseResultRegistrantContactJSON contains the JSON
// metadata for the struct
// [AccountRegistrarDomainListResponseResultRegistrantContact]
type accountRegistrarDomainListResponseResultRegistrantContactJSON struct {
	Address      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	Organization apijson.Field
	Phone        apijson.Field
	State        apijson.Field
	Zip          apijson.Field
	ID           apijson.Field
	Address2     apijson.Field
	Email        apijson.Field
	Fax          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseResultRegistrantContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseResultRegistrantContactJSON) RawJSON() string {
	return r.raw
}

// Statuses for domain transfers into Cloudflare Registrar.
type AccountRegistrarDomainListResponseResultTransferIn struct {
	// Form of authorization has been accepted by the registrant.
	AcceptFoa AccountRegistrarDomainListResponseResultTransferInAcceptFoa `json:"accept_foa"`
	// Shows transfer status with the registry.
	ApproveTransfer AccountRegistrarDomainListResponseResultTransferInApproveTransfer `json:"approve_transfer"`
	// Indicates if cancellation is still possible.
	CanCancelTransfer bool `json:"can_cancel_transfer"`
	// Privacy guards are disabled at the foreign registrar.
	DisablePrivacy AccountRegistrarDomainListResponseResultTransferInDisablePrivacy `json:"disable_privacy"`
	// Auth code has been entered and verified.
	EnterAuthCode AccountRegistrarDomainListResponseResultTransferInEnterAuthCode `json:"enter_auth_code"`
	// Domain is unlocked at the foreign registrar.
	UnlockDomain AccountRegistrarDomainListResponseResultTransferInUnlockDomain `json:"unlock_domain"`
	JSON         accountRegistrarDomainListResponseResultTransferInJSON         `json:"-"`
}

// accountRegistrarDomainListResponseResultTransferInJSON contains the JSON
// metadata for the struct [AccountRegistrarDomainListResponseResultTransferIn]
type accountRegistrarDomainListResponseResultTransferInJSON struct {
	AcceptFoa         apijson.Field
	ApproveTransfer   apijson.Field
	CanCancelTransfer apijson.Field
	DisablePrivacy    apijson.Field
	EnterAuthCode     apijson.Field
	UnlockDomain      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseResultTransferIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseResultTransferInJSON) RawJSON() string {
	return r.raw
}

// Form of authorization has been accepted by the registrant.
type AccountRegistrarDomainListResponseResultTransferInAcceptFoa string

const (
	AccountRegistrarDomainListResponseResultTransferInAcceptFoaNeeded AccountRegistrarDomainListResponseResultTransferInAcceptFoa = "needed"
	AccountRegistrarDomainListResponseResultTransferInAcceptFoaOk     AccountRegistrarDomainListResponseResultTransferInAcceptFoa = "ok"
)

func (r AccountRegistrarDomainListResponseResultTransferInAcceptFoa) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseResultTransferInAcceptFoaNeeded, AccountRegistrarDomainListResponseResultTransferInAcceptFoaOk:
		return true
	}
	return false
}

// Shows transfer status with the registry.
type AccountRegistrarDomainListResponseResultTransferInApproveTransfer string

const (
	AccountRegistrarDomainListResponseResultTransferInApproveTransferNeeded   AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "needed"
	AccountRegistrarDomainListResponseResultTransferInApproveTransferOk       AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "ok"
	AccountRegistrarDomainListResponseResultTransferInApproveTransferPending  AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "pending"
	AccountRegistrarDomainListResponseResultTransferInApproveTransferTrying   AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "trying"
	AccountRegistrarDomainListResponseResultTransferInApproveTransferRejected AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "rejected"
	AccountRegistrarDomainListResponseResultTransferInApproveTransferUnknown  AccountRegistrarDomainListResponseResultTransferInApproveTransfer = "unknown"
)

func (r AccountRegistrarDomainListResponseResultTransferInApproveTransfer) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseResultTransferInApproveTransferNeeded, AccountRegistrarDomainListResponseResultTransferInApproveTransferOk, AccountRegistrarDomainListResponseResultTransferInApproveTransferPending, AccountRegistrarDomainListResponseResultTransferInApproveTransferTrying, AccountRegistrarDomainListResponseResultTransferInApproveTransferRejected, AccountRegistrarDomainListResponseResultTransferInApproveTransferUnknown:
		return true
	}
	return false
}

// Privacy guards are disabled at the foreign registrar.
type AccountRegistrarDomainListResponseResultTransferInDisablePrivacy string

const (
	AccountRegistrarDomainListResponseResultTransferInDisablePrivacyNeeded  AccountRegistrarDomainListResponseResultTransferInDisablePrivacy = "needed"
	AccountRegistrarDomainListResponseResultTransferInDisablePrivacyOk      AccountRegistrarDomainListResponseResultTransferInDisablePrivacy = "ok"
	AccountRegistrarDomainListResponseResultTransferInDisablePrivacyUnknown AccountRegistrarDomainListResponseResultTransferInDisablePrivacy = "unknown"
)

func (r AccountRegistrarDomainListResponseResultTransferInDisablePrivacy) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseResultTransferInDisablePrivacyNeeded, AccountRegistrarDomainListResponseResultTransferInDisablePrivacyOk, AccountRegistrarDomainListResponseResultTransferInDisablePrivacyUnknown:
		return true
	}
	return false
}

// Auth code has been entered and verified.
type AccountRegistrarDomainListResponseResultTransferInEnterAuthCode string

const (
	AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeNeeded   AccountRegistrarDomainListResponseResultTransferInEnterAuthCode = "needed"
	AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeOk       AccountRegistrarDomainListResponseResultTransferInEnterAuthCode = "ok"
	AccountRegistrarDomainListResponseResultTransferInEnterAuthCodePending  AccountRegistrarDomainListResponseResultTransferInEnterAuthCode = "pending"
	AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeTrying   AccountRegistrarDomainListResponseResultTransferInEnterAuthCode = "trying"
	AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeRejected AccountRegistrarDomainListResponseResultTransferInEnterAuthCode = "rejected"
)

func (r AccountRegistrarDomainListResponseResultTransferInEnterAuthCode) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeNeeded, AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeOk, AccountRegistrarDomainListResponseResultTransferInEnterAuthCodePending, AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeTrying, AccountRegistrarDomainListResponseResultTransferInEnterAuthCodeRejected:
		return true
	}
	return false
}

// Domain is unlocked at the foreign registrar.
type AccountRegistrarDomainListResponseResultTransferInUnlockDomain string

const (
	AccountRegistrarDomainListResponseResultTransferInUnlockDomainNeeded  AccountRegistrarDomainListResponseResultTransferInUnlockDomain = "needed"
	AccountRegistrarDomainListResponseResultTransferInUnlockDomainOk      AccountRegistrarDomainListResponseResultTransferInUnlockDomain = "ok"
	AccountRegistrarDomainListResponseResultTransferInUnlockDomainPending AccountRegistrarDomainListResponseResultTransferInUnlockDomain = "pending"
	AccountRegistrarDomainListResponseResultTransferInUnlockDomainTrying  AccountRegistrarDomainListResponseResultTransferInUnlockDomain = "trying"
	AccountRegistrarDomainListResponseResultTransferInUnlockDomainUnknown AccountRegistrarDomainListResponseResultTransferInUnlockDomain = "unknown"
)

func (r AccountRegistrarDomainListResponseResultTransferInUnlockDomain) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseResultTransferInUnlockDomainNeeded, AccountRegistrarDomainListResponseResultTransferInUnlockDomainOk, AccountRegistrarDomainListResponseResultTransferInUnlockDomainPending, AccountRegistrarDomainListResponseResultTransferInUnlockDomainTrying, AccountRegistrarDomainListResponseResultTransferInUnlockDomainUnknown:
		return true
	}
	return false
}

// Whether the API call was successful
type AccountRegistrarDomainListResponseSuccess bool

const (
	AccountRegistrarDomainListResponseSuccessTrue AccountRegistrarDomainListResponseSuccess = true
)

func (r AccountRegistrarDomainListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountRegistrarDomainListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountRegistrarDomainListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       accountRegistrarDomainListResponseResultInfoJSON `json:"-"`
}

// accountRegistrarDomainListResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountRegistrarDomainListResponseResultInfo]
type accountRegistrarDomainListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountRegistrarDomainListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountRegistrarDomainListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountRegistrarDomainUpdateParams struct {
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r AccountRegistrarDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
