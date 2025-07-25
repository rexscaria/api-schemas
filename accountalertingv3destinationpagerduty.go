// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountAlertingV3DestinationPagerdutyService contains methods and other services
// that help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3DestinationPagerdutyService] method instead.
type AccountAlertingV3DestinationPagerdutyService struct {
	Options []option.RequestOption
	Connect *AccountAlertingV3DestinationPagerdutyConnectService
}

// NewAccountAlertingV3DestinationPagerdutyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountAlertingV3DestinationPagerdutyService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationPagerdutyService) {
	r = &AccountAlertingV3DestinationPagerdutyService{}
	r.Options = opts
	r.Connect = NewAccountAlertingV3DestinationPagerdutyConnectService(opts...)
	return
}

// Get a list of all configured PagerDuty services.
func (r *AccountAlertingV3DestinationPagerdutyService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationPagerdutyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *AccountAlertingV3DestinationPagerdutyService) DeleteAll(ctx context.Context, accountID string, opts ...option.RequestOption) (res *APIResponseCollectionAlerting, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIResponseAlerting struct {
	Errors   []APIResponseAlertingError   `json:"errors,required"`
	Messages []APIResponseAlertingMessage `json:"messages,required"`
	// Whether the API call was successful
	Success APIResponseAlertingSuccess `json:"success,required"`
	JSON    apiResponseAlertingJSON    `json:"-"`
}

// apiResponseAlertingJSON contains the JSON metadata for the struct
// [APIResponseAlerting]
type apiResponseAlertingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAlerting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAlertingJSON) RawJSON() string {
	return r.raw
}

func (r APIResponseAlerting) implementsAaaAuditLogs() {}

type APIResponseAlertingError struct {
	Code             int64                           `json:"code,required"`
	Message          string                          `json:"message,required"`
	DocumentationURL string                          `json:"documentation_url"`
	Source           APIResponseAlertingErrorsSource `json:"source"`
	JSON             apiResponseAlertingErrorJSON    `json:"-"`
}

// apiResponseAlertingErrorJSON contains the JSON metadata for the struct
// [APIResponseAlertingError]
type apiResponseAlertingErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseAlertingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAlertingErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseAlertingErrorsSource struct {
	Pointer string                              `json:"pointer"`
	JSON    apiResponseAlertingErrorsSourceJSON `json:"-"`
}

// apiResponseAlertingErrorsSourceJSON contains the JSON metadata for the struct
// [APIResponseAlertingErrorsSource]
type apiResponseAlertingErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAlertingErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAlertingErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseAlertingMessage struct {
	Code             int64                             `json:"code,required"`
	Message          string                            `json:"message,required"`
	DocumentationURL string                            `json:"documentation_url"`
	Source           APIResponseAlertingMessagesSource `json:"source"`
	JSON             apiResponseAlertingMessageJSON    `json:"-"`
}

// apiResponseAlertingMessageJSON contains the JSON metadata for the struct
// [APIResponseAlertingMessage]
type apiResponseAlertingMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseAlertingMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAlertingMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseAlertingMessagesSource struct {
	Pointer string                                `json:"pointer"`
	JSON    apiResponseAlertingMessagesSourceJSON `json:"-"`
}

// apiResponseAlertingMessagesSourceJSON contains the JSON metadata for the struct
// [APIResponseAlertingMessagesSource]
type apiResponseAlertingMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseAlertingMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseAlertingMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseAlertingSuccess bool

const (
	APIResponseAlertingSuccessTrue APIResponseAlertingSuccess = true
)

func (r APIResponseAlertingSuccess) IsKnown() bool {
	switch r {
	case APIResponseAlertingSuccessTrue:
		return true
	}
	return false
}

type APIResponseCollectionAlerting struct {
	Errors   []APIResponseCollectionAlertingError   `json:"errors,required"`
	Messages []APIResponseCollectionAlertingMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    APIResponseCollectionAlertingSuccess    `json:"success,required"`
	ResultInfo APIResponseCollectionAlertingResultInfo `json:"result_info"`
	JSON       apiResponseCollectionAlertingJSON       `json:"-"`
}

// apiResponseCollectionAlertingJSON contains the JSON metadata for the struct
// [APIResponseCollectionAlerting]
type apiResponseCollectionAlertingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAlerting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAlertingError struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           APIResponseCollectionAlertingErrorsSource `json:"source"`
	JSON             apiResponseCollectionAlertingErrorJSON    `json:"-"`
}

// apiResponseCollectionAlertingErrorJSON contains the JSON metadata for the struct
// [APIResponseCollectionAlertingError]
type apiResponseCollectionAlertingErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseCollectionAlertingError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingErrorJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAlertingErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    apiResponseCollectionAlertingErrorsSourceJSON `json:"-"`
}

// apiResponseCollectionAlertingErrorsSourceJSON contains the JSON metadata for the
// struct [APIResponseCollectionAlertingErrorsSource]
type apiResponseCollectionAlertingErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAlertingErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAlertingMessage struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           APIResponseCollectionAlertingMessagesSource `json:"source"`
	JSON             apiResponseCollectionAlertingMessageJSON    `json:"-"`
}

// apiResponseCollectionAlertingMessageJSON contains the JSON metadata for the
// struct [APIResponseCollectionAlertingMessage]
type apiResponseCollectionAlertingMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *APIResponseCollectionAlertingMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingMessageJSON) RawJSON() string {
	return r.raw
}

type APIResponseCollectionAlertingMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    apiResponseCollectionAlertingMessagesSourceJSON `json:"-"`
}

// apiResponseCollectionAlertingMessagesSourceJSON contains the JSON metadata for
// the struct [APIResponseCollectionAlertingMessagesSource]
type apiResponseCollectionAlertingMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAlertingMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type APIResponseCollectionAlertingSuccess bool

const (
	APIResponseCollectionAlertingSuccessTrue APIResponseCollectionAlertingSuccess = true
)

func (r APIResponseCollectionAlertingSuccess) IsKnown() bool {
	switch r {
	case APIResponseCollectionAlertingSuccessTrue:
		return true
	}
	return false
}

type APIResponseCollectionAlertingResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       apiResponseCollectionAlertingResultInfoJSON `json:"-"`
}

// apiResponseCollectionAlertingResultInfoJSON contains the JSON metadata for the
// struct [APIResponseCollectionAlertingResultInfo]
type apiResponseCollectionAlertingResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseCollectionAlertingResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseCollectionAlertingResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponse struct {
	Errors   []AccountAlertingV3DestinationPagerdutyListResponseError   `json:"errors,required"`
	Messages []AccountAlertingV3DestinationPagerdutyListResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success    AccountAlertingV3DestinationPagerdutyListResponseSuccess    `json:"success,required"`
	Result     []AccountAlertingV3DestinationPagerdutyListResponseResult   `json:"result"`
	ResultInfo AccountAlertingV3DestinationPagerdutyListResponseResultInfo `json:"result_info"`
	JSON       accountAlertingV3DestinationPagerdutyListResponseJSON       `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationPagerdutyListResponse]
type accountAlertingV3DestinationPagerdutyListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponseError struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           AccountAlertingV3DestinationPagerdutyListResponseErrorsSource `json:"source"`
	JSON             accountAlertingV3DestinationPagerdutyListResponseErrorJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseErrorJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationPagerdutyListResponseError]
type accountAlertingV3DestinationPagerdutyListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponseErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    accountAlertingV3DestinationPagerdutyListResponseErrorsSourceJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseErrorsSourceJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseErrorsSource]
type accountAlertingV3DestinationPagerdutyListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponseMessage struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           AccountAlertingV3DestinationPagerdutyListResponseMessagesSource `json:"source"`
	JSON             accountAlertingV3DestinationPagerdutyListResponseMessageJSON    `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseMessage]
type accountAlertingV3DestinationPagerdutyListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponseMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    accountAlertingV3DestinationPagerdutyListResponseMessagesSourceJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseMessagesSource]
type accountAlertingV3DestinationPagerdutyListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountAlertingV3DestinationPagerdutyListResponseSuccess bool

const (
	AccountAlertingV3DestinationPagerdutyListResponseSuccessTrue AccountAlertingV3DestinationPagerdutyListResponseSuccess = true
)

func (r AccountAlertingV3DestinationPagerdutyListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAlertingV3DestinationPagerdutyListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAlertingV3DestinationPagerdutyListResponseResult struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string                                                      `json:"name"`
	JSON accountAlertingV3DestinationPagerdutyListResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseResult]
type accountAlertingV3DestinationPagerdutyListResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationPagerdutyListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON `json:"-"`
}

// accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON contains the
// JSON metadata for the struct
// [AccountAlertingV3DestinationPagerdutyListResponseResultInfo]
type accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}
