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

// AccountEmailSecuritySettingDomainService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailSecuritySettingDomainService] method instead.
type AccountEmailSecuritySettingDomainService struct {
	Options []option.RequestOption
}

// NewAccountEmailSecuritySettingDomainService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEmailSecuritySettingDomainService(opts ...option.RequestOption) (r *AccountEmailSecuritySettingDomainService) {
	r = &AccountEmailSecuritySettingDomainService{}
	r.Options = opts
	return
}

// Get an email domain
func (r *AccountEmailSecuritySettingDomainService) Get(ctx context.Context, accountID string, domainID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an email domain
func (r *AccountEmailSecuritySettingDomainService) Update(ctx context.Context, accountID string, domainID int64, body AccountEmailSecuritySettingDomainUpdateParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Lists, searches, and sorts an account’s email domains.
func (r *AccountEmailSecuritySettingDomainService) List(ctx context.Context, accountID string, query AccountEmailSecuritySettingDomainListParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Unprotect an email domain
func (r *AccountEmailSecuritySettingDomainService) Unprotect(ctx context.Context, accountID string, domainID int64, opts ...option.RequestOption) (res *AccountEmailSecuritySettingDomainUnprotectResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Unprotect multiple email domains
func (r *AccountEmailSecuritySettingDomainService) UnprotectMultiple(ctx context.Context, accountID string, body AccountEmailSecuritySettingDomainUnprotectMultipleParams, opts ...option.RequestOption) (res *AccountEmailSecuritySettingDomainUnprotectMultipleResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/settings/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type DeliveryMode string

const (
	DeliveryModeDirect    DeliveryMode = "DIRECT"
	DeliveryModeBcc       DeliveryMode = "BCC"
	DeliveryModeJournal   DeliveryMode = "JOURNAL"
	DeliveryModeAPI       DeliveryMode = "API"
	DeliveryModeRetroScan DeliveryMode = "RETRO_SCAN"
)

func (r DeliveryMode) IsKnown() bool {
	switch r {
	case DeliveryModeDirect, DeliveryModeBcc, DeliveryModeJournal, DeliveryModeAPI, DeliveryModeRetroScan:
		return true
	}
	return false
}

type DispositionLabel string

const (
	DispositionLabelMalicious    DispositionLabel = "MALICIOUS"
	DispositionLabelMaliciousBec DispositionLabel = "MALICIOUS-BEC"
	DispositionLabelSuspicious   DispositionLabel = "SUSPICIOUS"
	DispositionLabelSpoof        DispositionLabel = "SPOOF"
	DispositionLabelSpam         DispositionLabel = "SPAM"
	DispositionLabelBulk         DispositionLabel = "BULK"
	DispositionLabelEncrypted    DispositionLabel = "ENCRYPTED"
	DispositionLabelExternal     DispositionLabel = "EXTERNAL"
	DispositionLabelUnknown      DispositionLabel = "UNKNOWN"
	DispositionLabelNone         DispositionLabel = "NONE"
)

func (r DispositionLabel) IsKnown() bool {
	switch r {
	case DispositionLabelMalicious, DispositionLabelMaliciousBec, DispositionLabelSuspicious, DispositionLabelSpoof, DispositionLabelSpam, DispositionLabelBulk, DispositionLabelEncrypted, DispositionLabelExternal, DispositionLabelUnknown, DispositionLabelNone:
		return true
	}
	return false
}

type ScannableFolder string

const (
	ScannableFolderAllItems ScannableFolder = "AllItems"
	ScannableFolderInbox    ScannableFolder = "Inbox"
)

func (r ScannableFolder) IsKnown() bool {
	switch r {
	case ScannableFolderAllItems, ScannableFolderInbox:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainGetResponse struct {
	Result AccountEmailSecuritySettingDomainGetResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingDomainGetResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingDomainGetResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecuritySettingDomainGetResponse]
type accountEmailSecuritySettingDomainGetResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainGetResponseResult struct {
	// The unique identifier for the domain.
	ID                   int64                                                             `json:"id,required"`
	AllowedDeliveryModes []DeliveryMode                                                    `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                                         `json:"created_at,required" format:"date-time"`
	Domain               string                                                            `json:"domain,required"`
	DropDispositions     []DispositionLabel                                                `json:"drop_dispositions,required"`
	IPRestrictions       []string                                                          `json:"ip_restrictions,required"`
	LastModified         time.Time                                                         `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                                             `json:"lookback_hops,required"`
	Transport            string                                                            `json:"transport,required"`
	Authorization        AccountEmailSecuritySettingDomainGetResponseResultAuthorization   `json:"authorization"`
	EmailsProcessed      AccountEmailSecuritySettingDomainGetResponseResultEmailsProcessed `json:"emails_processed"`
	Folder               AccountEmailSecuritySettingDomainGetResponseResultFolder          `json:"folder,nullable"`
	InboxProvider        AccountEmailSecuritySettingDomainGetResponseResultInboxProvider   `json:"inbox_provider,nullable"`
	IntegrationID        string                                                            `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                                            `json:"o365_tenant_id,nullable"`
	RequireTlsInbound    bool                                                              `json:"require_tls_inbound,nullable"`
	RequireTlsOutbound   bool                                                              `json:"require_tls_outbound,nullable"`
	JSON                 accountEmailSecuritySettingDomainGetResponseResultJSON            `json:"-"`
}

// accountEmailSecuritySettingDomainGetResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingDomainGetResponseResult]
type accountEmailSecuritySettingDomainGetResponseResultJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTlsInbound    apijson.Field
	RequireTlsOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainGetResponseResultAuthorization struct {
	Authorized    bool                                                                `json:"authorized,required"`
	Timestamp     time.Time                                                           `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                                              `json:"status_message,nullable"`
	JSON          accountEmailSecuritySettingDomainGetResponseResultAuthorizationJSON `json:"-"`
}

// accountEmailSecuritySettingDomainGetResponseResultAuthorizationJSON contains the
// JSON metadata for the struct
// [AccountEmailSecuritySettingDomainGetResponseResultAuthorization]
type accountEmailSecuritySettingDomainGetResponseResultAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainGetResponseResultAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainGetResponseResultAuthorizationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainGetResponseResultEmailsProcessed struct {
	Timestamp                    time.Time                                                             `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                                                 `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                                                 `json:"total_emails_processed_previous,required"`
	JSON                         accountEmailSecuritySettingDomainGetResponseResultEmailsProcessedJSON `json:"-"`
}

// accountEmailSecuritySettingDomainGetResponseResultEmailsProcessedJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainGetResponseResultEmailsProcessed]
type accountEmailSecuritySettingDomainGetResponseResultEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainGetResponseResultEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainGetResponseResultEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainGetResponseResultFolder string

const (
	AccountEmailSecuritySettingDomainGetResponseResultFolderAllItems AccountEmailSecuritySettingDomainGetResponseResultFolder = "AllItems"
	AccountEmailSecuritySettingDomainGetResponseResultFolderInbox    AccountEmailSecuritySettingDomainGetResponseResultFolder = "Inbox"
)

func (r AccountEmailSecuritySettingDomainGetResponseResultFolder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainGetResponseResultFolderAllItems, AccountEmailSecuritySettingDomainGetResponseResultFolderInbox:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainGetResponseResultInboxProvider string

const (
	AccountEmailSecuritySettingDomainGetResponseResultInboxProviderMicrosoft AccountEmailSecuritySettingDomainGetResponseResultInboxProvider = "Microsoft"
	AccountEmailSecuritySettingDomainGetResponseResultInboxProviderGoogle    AccountEmailSecuritySettingDomainGetResponseResultInboxProvider = "Google"
)

func (r AccountEmailSecuritySettingDomainGetResponseResultInboxProvider) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainGetResponseResultInboxProviderMicrosoft, AccountEmailSecuritySettingDomainGetResponseResultInboxProviderGoogle:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainUpdateResponse struct {
	Result AccountEmailSecuritySettingDomainUpdateResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingDomainUpdateResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingDomainUpdateResponseJSON contains the JSON metadata
// for the struct [AccountEmailSecuritySettingDomainUpdateResponse]
type accountEmailSecuritySettingDomainUpdateResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUpdateResponseResult struct {
	// The unique identifier for the domain.
	ID                   int64                                                                `json:"id,required"`
	AllowedDeliveryModes []DeliveryMode                                                       `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                                            `json:"created_at,required" format:"date-time"`
	Domain               string                                                               `json:"domain,required"`
	DropDispositions     []DispositionLabel                                                   `json:"drop_dispositions,required"`
	IPRestrictions       []string                                                             `json:"ip_restrictions,required"`
	LastModified         time.Time                                                            `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                                                `json:"lookback_hops,required"`
	Transport            string                                                               `json:"transport,required"`
	Authorization        AccountEmailSecuritySettingDomainUpdateResponseResultAuthorization   `json:"authorization"`
	EmailsProcessed      AccountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessed `json:"emails_processed"`
	Folder               AccountEmailSecuritySettingDomainUpdateResponseResultFolder          `json:"folder,nullable"`
	InboxProvider        AccountEmailSecuritySettingDomainUpdateResponseResultInboxProvider   `json:"inbox_provider,nullable"`
	IntegrationID        string                                                               `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                                               `json:"o365_tenant_id,nullable"`
	RequireTlsInbound    bool                                                                 `json:"require_tls_inbound,nullable"`
	RequireTlsOutbound   bool                                                                 `json:"require_tls_outbound,nullable"`
	JSON                 accountEmailSecuritySettingDomainUpdateResponseResultJSON            `json:"-"`
}

// accountEmailSecuritySettingDomainUpdateResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingDomainUpdateResponseResult]
type accountEmailSecuritySettingDomainUpdateResponseResultJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTlsInbound    apijson.Field
	RequireTlsOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUpdateResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUpdateResponseResultAuthorization struct {
	Authorized    bool                                                                   `json:"authorized,required"`
	Timestamp     time.Time                                                              `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                                                 `json:"status_message,nullable"`
	JSON          accountEmailSecuritySettingDomainUpdateResponseResultAuthorizationJSON `json:"-"`
}

// accountEmailSecuritySettingDomainUpdateResponseResultAuthorizationJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainUpdateResponseResultAuthorization]
type accountEmailSecuritySettingDomainUpdateResponseResultAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUpdateResponseResultAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUpdateResponseResultAuthorizationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessed struct {
	Timestamp                    time.Time                                                                `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                                                    `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                                                    `json:"total_emails_processed_previous,required"`
	JSON                         accountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessedJSON `json:"-"`
}

// accountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessedJSON
// contains the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessed]
type accountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUpdateResponseResultEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUpdateResponseResultFolder string

const (
	AccountEmailSecuritySettingDomainUpdateResponseResultFolderAllItems AccountEmailSecuritySettingDomainUpdateResponseResultFolder = "AllItems"
	AccountEmailSecuritySettingDomainUpdateResponseResultFolderInbox    AccountEmailSecuritySettingDomainUpdateResponseResultFolder = "Inbox"
)

func (r AccountEmailSecuritySettingDomainUpdateResponseResultFolder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainUpdateResponseResultFolderAllItems, AccountEmailSecuritySettingDomainUpdateResponseResultFolderInbox:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainUpdateResponseResultInboxProvider string

const (
	AccountEmailSecuritySettingDomainUpdateResponseResultInboxProviderMicrosoft AccountEmailSecuritySettingDomainUpdateResponseResultInboxProvider = "Microsoft"
	AccountEmailSecuritySettingDomainUpdateResponseResultInboxProviderGoogle    AccountEmailSecuritySettingDomainUpdateResponseResultInboxProvider = "Google"
)

func (r AccountEmailSecuritySettingDomainUpdateResponseResultInboxProvider) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainUpdateResponseResultInboxProviderMicrosoft, AccountEmailSecuritySettingDomainUpdateResponseResultInboxProviderGoogle:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainListResponse struct {
	Result     []AccountEmailSecuritySettingDomainListResponseResult `json:"result,required"`
	ResultInfo ResultInfoEmailSecurity                               `json:"result_info,required"`
	JSON       accountEmailSecuritySettingDomainListResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingDomainListResponseJSON contains the JSON metadata for
// the struct [AccountEmailSecuritySettingDomainListResponse]
type accountEmailSecuritySettingDomainListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainListResponseResult struct {
	// The unique identifier for the domain.
	ID                   int64                                                              `json:"id,required"`
	AllowedDeliveryModes []DeliveryMode                                                     `json:"allowed_delivery_modes,required"`
	CreatedAt            time.Time                                                          `json:"created_at,required" format:"date-time"`
	Domain               string                                                             `json:"domain,required"`
	DropDispositions     []DispositionLabel                                                 `json:"drop_dispositions,required"`
	IPRestrictions       []string                                                           `json:"ip_restrictions,required"`
	LastModified         time.Time                                                          `json:"last_modified,required" format:"date-time"`
	LookbackHops         int64                                                              `json:"lookback_hops,required"`
	Transport            string                                                             `json:"transport,required"`
	Authorization        AccountEmailSecuritySettingDomainListResponseResultAuthorization   `json:"authorization"`
	EmailsProcessed      AccountEmailSecuritySettingDomainListResponseResultEmailsProcessed `json:"emails_processed"`
	Folder               AccountEmailSecuritySettingDomainListResponseResultFolder          `json:"folder,nullable"`
	InboxProvider        AccountEmailSecuritySettingDomainListResponseResultInboxProvider   `json:"inbox_provider,nullable"`
	IntegrationID        string                                                             `json:"integration_id,nullable" format:"uuid"`
	O365TenantID         string                                                             `json:"o365_tenant_id,nullable"`
	RequireTlsInbound    bool                                                               `json:"require_tls_inbound,nullable"`
	RequireTlsOutbound   bool                                                               `json:"require_tls_outbound,nullable"`
	JSON                 accountEmailSecuritySettingDomainListResponseResultJSON            `json:"-"`
}

// accountEmailSecuritySettingDomainListResponseResultJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingDomainListResponseResult]
type accountEmailSecuritySettingDomainListResponseResultJSON struct {
	ID                   apijson.Field
	AllowedDeliveryModes apijson.Field
	CreatedAt            apijson.Field
	Domain               apijson.Field
	DropDispositions     apijson.Field
	IPRestrictions       apijson.Field
	LastModified         apijson.Field
	LookbackHops         apijson.Field
	Transport            apijson.Field
	Authorization        apijson.Field
	EmailsProcessed      apijson.Field
	Folder               apijson.Field
	InboxProvider        apijson.Field
	IntegrationID        apijson.Field
	O365TenantID         apijson.Field
	RequireTlsInbound    apijson.Field
	RequireTlsOutbound   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainListResponseResultAuthorization struct {
	Authorized    bool                                                                 `json:"authorized,required"`
	Timestamp     time.Time                                                            `json:"timestamp,required" format:"date-time"`
	StatusMessage string                                                               `json:"status_message,nullable"`
	JSON          accountEmailSecuritySettingDomainListResponseResultAuthorizationJSON `json:"-"`
}

// accountEmailSecuritySettingDomainListResponseResultAuthorizationJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainListResponseResultAuthorization]
type accountEmailSecuritySettingDomainListResponseResultAuthorizationJSON struct {
	Authorized    apijson.Field
	Timestamp     apijson.Field
	StatusMessage apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainListResponseResultAuthorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainListResponseResultAuthorizationJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainListResponseResultEmailsProcessed struct {
	Timestamp                    time.Time                                                              `json:"timestamp,required" format:"date-time"`
	TotalEmailsProcessed         int64                                                                  `json:"total_emails_processed,required"`
	TotalEmailsProcessedPrevious int64                                                                  `json:"total_emails_processed_previous,required"`
	JSON                         accountEmailSecuritySettingDomainListResponseResultEmailsProcessedJSON `json:"-"`
}

// accountEmailSecuritySettingDomainListResponseResultEmailsProcessedJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainListResponseResultEmailsProcessed]
type accountEmailSecuritySettingDomainListResponseResultEmailsProcessedJSON struct {
	Timestamp                    apijson.Field
	TotalEmailsProcessed         apijson.Field
	TotalEmailsProcessedPrevious apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainListResponseResultEmailsProcessed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainListResponseResultEmailsProcessedJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainListResponseResultFolder string

const (
	AccountEmailSecuritySettingDomainListResponseResultFolderAllItems AccountEmailSecuritySettingDomainListResponseResultFolder = "AllItems"
	AccountEmailSecuritySettingDomainListResponseResultFolderInbox    AccountEmailSecuritySettingDomainListResponseResultFolder = "Inbox"
)

func (r AccountEmailSecuritySettingDomainListResponseResultFolder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainListResponseResultFolderAllItems, AccountEmailSecuritySettingDomainListResponseResultFolderInbox:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainListResponseResultInboxProvider string

const (
	AccountEmailSecuritySettingDomainListResponseResultInboxProviderMicrosoft AccountEmailSecuritySettingDomainListResponseResultInboxProvider = "Microsoft"
	AccountEmailSecuritySettingDomainListResponseResultInboxProviderGoogle    AccountEmailSecuritySettingDomainListResponseResultInboxProvider = "Google"
)

func (r AccountEmailSecuritySettingDomainListResponseResultInboxProvider) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainListResponseResultInboxProviderMicrosoft, AccountEmailSecuritySettingDomainListResponseResultInboxProviderGoogle:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainUnprotectResponse struct {
	Result AccountEmailSecuritySettingDomainUnprotectResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingDomainUnprotectResponseJSON   `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingDomainUnprotectResponseJSON contains the JSON
// metadata for the struct [AccountEmailSecuritySettingDomainUnprotectResponse]
type accountEmailSecuritySettingDomainUnprotectResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUnprotectResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUnprotectResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUnprotectResponseResult struct {
	// The unique identifier for the domain.
	ID   int64                                                        `json:"id,required"`
	JSON accountEmailSecuritySettingDomainUnprotectResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingDomainUnprotectResponseResultJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingDomainUnprotectResponseResult]
type accountEmailSecuritySettingDomainUnprotectResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUnprotectResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUnprotectResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUnprotectMultipleResponse struct {
	Result []AccountEmailSecuritySettingDomainUnprotectMultipleResponseResult `json:"result,required"`
	JSON   accountEmailSecuritySettingDomainUnprotectMultipleResponseJSON     `json:"-"`
	APIResponseEmailSecurity
}

// accountEmailSecuritySettingDomainUnprotectMultipleResponseJSON contains the JSON
// metadata for the struct
// [AccountEmailSecuritySettingDomainUnprotectMultipleResponse]
type accountEmailSecuritySettingDomainUnprotectMultipleResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUnprotectMultipleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUnprotectMultipleResponseJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUnprotectMultipleResponseResult struct {
	// The unique identifier for the domain.
	ID   int64                                                                `json:"id,required"`
	JSON accountEmailSecuritySettingDomainUnprotectMultipleResponseResultJSON `json:"-"`
}

// accountEmailSecuritySettingDomainUnprotectMultipleResponseResultJSON contains
// the JSON metadata for the struct
// [AccountEmailSecuritySettingDomainUnprotectMultipleResponseResult]
type accountEmailSecuritySettingDomainUnprotectMultipleResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailSecuritySettingDomainUnprotectMultipleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailSecuritySettingDomainUnprotectMultipleResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountEmailSecuritySettingDomainUpdateParams struct {
	IPRestrictions     param.Field[[]string]           `json:"ip_restrictions,required"`
	Domain             param.Field[string]             `json:"domain"`
	DropDispositions   param.Field[[]DispositionLabel] `json:"drop_dispositions"`
	Folder             param.Field[ScannableFolder]    `json:"folder"`
	IntegrationID      param.Field[string]             `json:"integration_id" format:"uuid"`
	LookbackHops       param.Field[int64]              `json:"lookback_hops"`
	RequireTlsInbound  param.Field[bool]               `json:"require_tls_inbound"`
	RequireTlsOutbound param.Field[bool]               `json:"require_tls_outbound"`
	Transport          param.Field[string]             `json:"transport"`
}

func (r AccountEmailSecuritySettingDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailSecuritySettingDomainListParams struct {
	// Filters response to domains with the currently active delivery mode.
	ActiveDeliveryMode param.Field[DeliveryMode] `query:"active_delivery_mode"`
	// Filters response to domains with the provided delivery mode.
	AllowedDeliveryMode param.Field[DeliveryMode] `query:"allowed_delivery_mode"`
	// The sorting direction.
	Direction param.Field[SortingDirection] `query:"direction"`
	// Filters results by the provided domains, allowing for multiple occurrences.
	Domain param.Field[[]string] `query:"domain"`
	// The field to sort by.
	Order param.Field[AccountEmailSecuritySettingDomainListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Allows searching in multiple properties of a record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountEmailSecuritySettingDomainListParams]'s query
// parameters as `url.Values`.
func (r AccountEmailSecuritySettingDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort by.
type AccountEmailSecuritySettingDomainListParamsOrder string

const (
	AccountEmailSecuritySettingDomainListParamsOrderDomain    AccountEmailSecuritySettingDomainListParamsOrder = "domain"
	AccountEmailSecuritySettingDomainListParamsOrderCreatedAt AccountEmailSecuritySettingDomainListParamsOrder = "created_at"
)

func (r AccountEmailSecuritySettingDomainListParamsOrder) IsKnown() bool {
	switch r {
	case AccountEmailSecuritySettingDomainListParamsOrderDomain, AccountEmailSecuritySettingDomainListParamsOrderCreatedAt:
		return true
	}
	return false
}

type AccountEmailSecuritySettingDomainUnprotectMultipleParams struct {
	Body []AccountEmailSecuritySettingDomainUnprotectMultipleParamsBody `json:"body,required"`
}

func (r AccountEmailSecuritySettingDomainUnprotectMultipleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountEmailSecuritySettingDomainUnprotectMultipleParamsBody struct {
	// The unique identifier for the domain.
	ID param.Field[int64] `json:"id,required"`
}

func (r AccountEmailSecuritySettingDomainUnprotectMultipleParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
