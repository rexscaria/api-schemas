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

// AccountCustomNService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountCustomNService] method instead.
type AccountCustomNService struct {
	Options []option.RequestOption
}

// NewAccountCustomNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountCustomNService(opts ...option.RequestOption) (r *AccountCustomNService) {
	r = &AccountCustomNService{}
	r.Options = opts
	return
}

// Add Account Custom Nameserver
func (r *AccountCustomNService) New(ctx context.Context, accountID string, body AccountCustomNNewParams, opts ...option.RequestOption) (res *AccountCustomNNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List an account's custom nameservers.
func (r *AccountCustomNService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountCustomNListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Account Custom Nameserver
func (r *AccountCustomNService) Delete(ctx context.Context, accountID string, customNsID string, opts ...option.RequestOption) (res *AccountCustomNDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customNsID == "" {
		err = errors.New("missing required custom_ns_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", accountID, customNsID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// A single account custom nameserver.
type CustomNs struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNsDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NsName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	//
	// Deprecated: deprecated
	Status CustomNsStatus `json:"status,required"`
	// Identifier
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NsSet float64      `json:"ns_set"`
	JSON  customNsJSON `json:"-"`
}

// customNsJSON contains the JSON metadata for the struct [CustomNs]
type customNsJSON struct {
	DNSRecords  apijson.Field
	NsName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NsSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsJSON) RawJSON() string {
	return r.raw
}

type CustomNsDNSRecord struct {
	// DNS record type.
	Type CustomNsDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                `json:"value"`
	JSON  customNsDNSRecordJSON `json:"-"`
}

// customNsDNSRecordJSON contains the JSON metadata for the struct
// [CustomNsDNSRecord]
type customNsDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type CustomNsDNSRecordsType string

const (
	CustomNsDNSRecordsTypeA    CustomNsDNSRecordsType = "A"
	CustomNsDNSRecordsTypeAaaa CustomNsDNSRecordsType = "AAAA"
)

func (r CustomNsDNSRecordsType) IsKnown() bool {
	switch r {
	case CustomNsDNSRecordsTypeA, CustomNsDNSRecordsTypeAaaa:
		return true
	}
	return false
}

// Verification status of the nameserver.
type CustomNsStatus string

const (
	CustomNsStatusMoved    CustomNsStatus = "moved"
	CustomNsStatusPending  CustomNsStatus = "pending"
	CustomNsStatusVerified CustomNsStatus = "verified"
)

func (r CustomNsStatus) IsKnown() bool {
	switch r {
	case CustomNsStatusMoved, CustomNsStatusPending, CustomNsStatusVerified:
		return true
	}
	return false
}

type CustomNsMessages struct {
	Code             int64                  `json:"code,required"`
	Message          string                 `json:"message,required"`
	DocumentationURL string                 `json:"documentation_url"`
	Source           CustomNsMessagesSource `json:"source"`
	JSON             customNsMessagesJSON   `json:"-"`
}

// customNsMessagesJSON contains the JSON metadata for the struct
// [CustomNsMessages]
type customNsMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomNsMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomNsMessagesSource struct {
	Pointer string                     `json:"pointer"`
	JSON    customNsMessagesSourceJSON `json:"-"`
}

// customNsMessagesSourceJSON contains the JSON metadata for the struct
// [CustomNsMessagesSource]
type customNsMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNsMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNsMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponse struct {
	Errors   []AccountCustomNNewResponseError   `json:"errors,required"`
	Messages []AccountCustomNNewResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success AccountCustomNNewResponseSuccess `json:"success,required"`
	// A single account custom nameserver.
	Result CustomNs                      `json:"result"`
	JSON   accountCustomNNewResponseJSON `json:"-"`
}

// accountCustomNNewResponseJSON contains the JSON metadata for the struct
// [AccountCustomNNewResponse]
type accountCustomNNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponseError struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AccountCustomNNewResponseErrorsSource `json:"source"`
	JSON             accountCustomNNewResponseErrorJSON    `json:"-"`
}

// accountCustomNNewResponseErrorJSON contains the JSON metadata for the struct
// [AccountCustomNNewResponseError]
type accountCustomNNewResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponseErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    accountCustomNNewResponseErrorsSourceJSON `json:"-"`
}

// accountCustomNNewResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountCustomNNewResponseErrorsSource]
type accountCustomNNewResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNNewResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponseMessage struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AccountCustomNNewResponseMessagesSource `json:"source"`
	JSON             accountCustomNNewResponseMessageJSON    `json:"-"`
}

// accountCustomNNewResponseMessageJSON contains the JSON metadata for the struct
// [AccountCustomNNewResponseMessage]
type accountCustomNNewResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewResponseMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    accountCustomNNewResponseMessagesSourceJSON `json:"-"`
}

// accountCustomNNewResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountCustomNNewResponseMessagesSource]
type accountCustomNNewResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNNewResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNNewResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountCustomNNewResponseSuccess bool

const (
	AccountCustomNNewResponseSuccessTrue AccountCustomNNewResponseSuccess = true
)

func (r AccountCustomNNewResponseSuccess) IsKnown() bool {
	switch r {
	case AccountCustomNNewResponseSuccessTrue:
		return true
	}
	return false
}

type AccountCustomNListResponse struct {
	Errors   []AccountCustomNListResponseError   `json:"errors,required"`
	Messages []AccountCustomNListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountCustomNListResponseSuccess    `json:"success,required"`
	Result     []CustomNs                           `json:"result"`
	ResultInfo AccountCustomNListResponseResultInfo `json:"result_info"`
	JSON       accountCustomNListResponseJSON       `json:"-"`
}

// accountCustomNListResponseJSON contains the JSON metadata for the struct
// [AccountCustomNListResponse]
type accountCustomNListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNListResponseError struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AccountCustomNListResponseErrorsSource `json:"source"`
	JSON             accountCustomNListResponseErrorJSON    `json:"-"`
}

// accountCustomNListResponseErrorJSON contains the JSON metadata for the struct
// [AccountCustomNListResponseError]
type accountCustomNListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNListResponseErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    accountCustomNListResponseErrorsSourceJSON `json:"-"`
}

// accountCustomNListResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountCustomNListResponseErrorsSource]
type accountCustomNListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNListResponseMessage struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountCustomNListResponseMessagesSource `json:"source"`
	JSON             accountCustomNListResponseMessageJSON    `json:"-"`
}

// accountCustomNListResponseMessageJSON contains the JSON metadata for the struct
// [AccountCustomNListResponseMessage]
type accountCustomNListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNListResponseMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountCustomNListResponseMessagesSourceJSON `json:"-"`
}

// accountCustomNListResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AccountCustomNListResponseMessagesSource]
type accountCustomNListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountCustomNListResponseSuccess bool

const (
	AccountCustomNListResponseSuccessTrue AccountCustomNListResponseSuccess = true
)

func (r AccountCustomNListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountCustomNListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountCustomNListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       accountCustomNListResponseResultInfoJSON `json:"-"`
}

// accountCustomNListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountCustomNListResponseResultInfo]
type accountCustomNListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponse struct {
	Errors   []AccountCustomNDeleteResponseError   `json:"errors,required"`
	Messages []AccountCustomNDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountCustomNDeleteResponseSuccess    `json:"success,required"`
	Result     []string                               `json:"result"`
	ResultInfo AccountCustomNDeleteResponseResultInfo `json:"result_info"`
	JSON       accountCustomNDeleteResponseJSON       `json:"-"`
}

// accountCustomNDeleteResponseJSON contains the JSON metadata for the struct
// [AccountCustomNDeleteResponse]
type accountCustomNDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponseError struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           AccountCustomNDeleteResponseErrorsSource `json:"source"`
	JSON             accountCustomNDeleteResponseErrorJSON    `json:"-"`
}

// accountCustomNDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountCustomNDeleteResponseError]
type accountCustomNDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponseErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    accountCustomNDeleteResponseErrorsSourceJSON `json:"-"`
}

// accountCustomNDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AccountCustomNDeleteResponseErrorsSource]
type accountCustomNDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponseMessage struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AccountCustomNDeleteResponseMessagesSource `json:"source"`
	JSON             accountCustomNDeleteResponseMessageJSON    `json:"-"`
}

// accountCustomNDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountCustomNDeleteResponseMessage]
type accountCustomNDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNDeleteResponseMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    accountCustomNDeleteResponseMessagesSourceJSON `json:"-"`
}

// accountCustomNDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountCustomNDeleteResponseMessagesSource]
type accountCustomNDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountCustomNDeleteResponseSuccess bool

const (
	AccountCustomNDeleteResponseSuccessTrue AccountCustomNDeleteResponseSuccess = true
)

func (r AccountCustomNDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountCustomNDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountCustomNDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       accountCustomNDeleteResponseResultInfoJSON `json:"-"`
}

// accountCustomNDeleteResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountCustomNDeleteResponseResultInfo]
type accountCustomNDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCustomNDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountCustomNDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountCustomNNewParams struct {
	// The FQDN of the name server.
	NsName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NsSet param.Field[float64] `json:"ns_set"`
}

func (r AccountCustomNNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
