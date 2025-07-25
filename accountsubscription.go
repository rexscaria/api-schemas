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

// AccountSubscriptionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountSubscriptionService] method instead.
type AccountSubscriptionService struct {
	Options []option.RequestOption
}

// NewAccountSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountSubscriptionService(opts ...option.RequestOption) (r *AccountSubscriptionService) {
	r = &AccountSubscriptionService{}
	r.Options = opts
	return
}

// Creates an account subscription.
func (r *AccountSubscriptionService) New(ctx context.Context, accountID string, body AccountSubscriptionNewParams, opts ...option.RequestOption) (res *AccountSubscriptionResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates an account subscription.
func (r *AccountSubscriptionService) Update(ctx context.Context, accountID string, subscriptionIdentifier string, body AccountSubscriptionUpdateParams, opts ...option.RequestOption) (res *AccountSubscriptionResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all of an account's subscriptions.
func (r *AccountSubscriptionService) List(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountSubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an account's subscription.
func (r *AccountSubscriptionService) Delete(ctx context.Context, accountID string, subscriptionIdentifier string, opts ...option.RequestOption) (res *AccountSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if subscriptionIdentifier == "" {
		err = errors.New("missing required subscription_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/subscriptions/%s", accountID, subscriptionIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccountSubscriptionResponseSingle struct {
	Errors   []AccountSubscriptionResponseSingleError   `json:"errors,required"`
	Messages []AccountSubscriptionResponseSingleMessage `json:"messages,required"`
	Result   Subscription                               `json:"result,required"`
	// Whether the API call was successful
	Success AccountSubscriptionResponseSingleSuccess `json:"success,required"`
	JSON    accountSubscriptionResponseSingleJSON    `json:"-"`
}

// accountSubscriptionResponseSingleJSON contains the JSON metadata for the struct
// [AccountSubscriptionResponseSingle]
type accountSubscriptionResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionResponseSingleError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountSubscriptionResponseSingleErrorsSource `json:"source"`
	JSON             accountSubscriptionResponseSingleErrorJSON    `json:"-"`
}

// accountSubscriptionResponseSingleErrorJSON contains the JSON metadata for the
// struct [AccountSubscriptionResponseSingleError]
type accountSubscriptionResponseSingleErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionResponseSingleErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountSubscriptionResponseSingleErrorsSourceJSON `json:"-"`
}

// accountSubscriptionResponseSingleErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSubscriptionResponseSingleErrorsSource]
type accountSubscriptionResponseSingleErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingleErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionResponseSingleMessage struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccountSubscriptionResponseSingleMessagesSource `json:"source"`
	JSON             accountSubscriptionResponseSingleMessageJSON    `json:"-"`
}

// accountSubscriptionResponseSingleMessageJSON contains the JSON metadata for the
// struct [AccountSubscriptionResponseSingleMessage]
type accountSubscriptionResponseSingleMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleMessageJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionResponseSingleMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accountSubscriptionResponseSingleMessagesSourceJSON `json:"-"`
}

// accountSubscriptionResponseSingleMessagesSourceJSON contains the JSON metadata
// for the struct [AccountSubscriptionResponseSingleMessagesSource]
type accountSubscriptionResponseSingleMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingleMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSubscriptionResponseSingleSuccess bool

const (
	AccountSubscriptionResponseSingleSuccessTrue AccountSubscriptionResponseSingleSuccess = true
)

func (r AccountSubscriptionResponseSingleSuccess) IsKnown() bool {
	switch r {
	case AccountSubscriptionResponseSingleSuccessTrue:
		return true
	}
	return false
}

type BillSubsAPIMessages struct {
	Code             int64                     `json:"code,required"`
	Message          string                    `json:"message,required"`
	DocumentationURL string                    `json:"documentation_url"`
	Source           BillSubsAPIMessagesSource `json:"source"`
	JSON             billSubsAPIMessagesJSON   `json:"-"`
}

// billSubsAPIMessagesJSON contains the JSON metadata for the struct
// [BillSubsAPIMessages]
type billSubsAPIMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BillSubsAPIMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIMessagesJSON) RawJSON() string {
	return r.raw
}

type BillSubsAPIMessagesSource struct {
	Pointer string                        `json:"pointer"`
	JSON    billSubsAPIMessagesSourceJSON `json:"-"`
}

// billSubsAPIMessagesSourceJSON contains the JSON metadata for the struct
// [BillSubsAPIMessagesSource]
type billSubsAPIMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillSubsAPIMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type Subscription struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan SubscriptionRatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionState `json:"state"`
	JSON  subscriptionJSON  `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyWeekly    SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly   SubscriptionFrequency = "monthly"
	SubscriptionFrequencyQuarterly SubscriptionFrequency = "quarterly"
	SubscriptionFrequencyYearly    SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyQuarterly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// The rate plan applied to the subscription.
type SubscriptionRatePlan struct {
	// The ID of the rate plan.
	ID SubscriptionRatePlanID `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string                 `json:"sets"`
	JSON subscriptionRatePlanJSON `json:"-"`
}

// subscriptionRatePlanJSON contains the JSON metadata for the struct
// [SubscriptionRatePlan]
type subscriptionRatePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SubscriptionRatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionRatePlanJSON) RawJSON() string {
	return r.raw
}

// The ID of the rate plan.
type SubscriptionRatePlanID string

const (
	SubscriptionRatePlanIDFree               SubscriptionRatePlanID = "free"
	SubscriptionRatePlanIDLite               SubscriptionRatePlanID = "lite"
	SubscriptionRatePlanIDPro                SubscriptionRatePlanID = "pro"
	SubscriptionRatePlanIDProPlus            SubscriptionRatePlanID = "pro_plus"
	SubscriptionRatePlanIDBusiness           SubscriptionRatePlanID = "business"
	SubscriptionRatePlanIDEnterprise         SubscriptionRatePlanID = "enterprise"
	SubscriptionRatePlanIDPartnersFree       SubscriptionRatePlanID = "partners_free"
	SubscriptionRatePlanIDPartnersPro        SubscriptionRatePlanID = "partners_pro"
	SubscriptionRatePlanIDPartnersBusiness   SubscriptionRatePlanID = "partners_business"
	SubscriptionRatePlanIDPartnersEnterprise SubscriptionRatePlanID = "partners_enterprise"
)

func (r SubscriptionRatePlanID) IsKnown() bool {
	switch r {
	case SubscriptionRatePlanIDFree, SubscriptionRatePlanIDLite, SubscriptionRatePlanIDPro, SubscriptionRatePlanIDProPlus, SubscriptionRatePlanIDBusiness, SubscriptionRatePlanIDEnterprise, SubscriptionRatePlanIDPartnersFree, SubscriptionRatePlanIDPartnersPro, SubscriptionRatePlanIDPartnersBusiness, SubscriptionRatePlanIDPartnersEnterprise:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionState string

const (
	SubscriptionStateTrial           SubscriptionState = "Trial"
	SubscriptionStateProvisioned     SubscriptionState = "Provisioned"
	SubscriptionStatePaid            SubscriptionState = "Paid"
	SubscriptionStateAwaitingPayment SubscriptionState = "AwaitingPayment"
	SubscriptionStateCancelled       SubscriptionState = "Cancelled"
	SubscriptionStateFailed          SubscriptionState = "Failed"
	SubscriptionStateExpired         SubscriptionState = "Expired"
)

func (r SubscriptionState) IsKnown() bool {
	switch r {
	case SubscriptionStateTrial, SubscriptionStateProvisioned, SubscriptionStatePaid, SubscriptionStateAwaitingPayment, SubscriptionStateCancelled, SubscriptionStateFailed, SubscriptionStateExpired:
		return true
	}
	return false
}

type SubscriptionV2Param struct {
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionV2Frequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionV2RatePlanParam] `json:"rate_plan"`
}

func (r SubscriptionV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// How often the subscription is renewed automatically.
type SubscriptionV2Frequency string

const (
	SubscriptionV2FrequencyWeekly    SubscriptionV2Frequency = "weekly"
	SubscriptionV2FrequencyMonthly   SubscriptionV2Frequency = "monthly"
	SubscriptionV2FrequencyQuarterly SubscriptionV2Frequency = "quarterly"
	SubscriptionV2FrequencyYearly    SubscriptionV2Frequency = "yearly"
)

func (r SubscriptionV2Frequency) IsKnown() bool {
	switch r {
	case SubscriptionV2FrequencyWeekly, SubscriptionV2FrequencyMonthly, SubscriptionV2FrequencyQuarterly, SubscriptionV2FrequencyYearly:
		return true
	}
	return false
}

// The rate plan applied to the subscription.
type SubscriptionV2RatePlanParam struct {
	// The ID of the rate plan.
	ID param.Field[SubscriptionV2RatePlanID] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r SubscriptionV2RatePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The ID of the rate plan.
type SubscriptionV2RatePlanID string

const (
	SubscriptionV2RatePlanIDFree               SubscriptionV2RatePlanID = "free"
	SubscriptionV2RatePlanIDLite               SubscriptionV2RatePlanID = "lite"
	SubscriptionV2RatePlanIDPro                SubscriptionV2RatePlanID = "pro"
	SubscriptionV2RatePlanIDProPlus            SubscriptionV2RatePlanID = "pro_plus"
	SubscriptionV2RatePlanIDBusiness           SubscriptionV2RatePlanID = "business"
	SubscriptionV2RatePlanIDEnterprise         SubscriptionV2RatePlanID = "enterprise"
	SubscriptionV2RatePlanIDPartnersFree       SubscriptionV2RatePlanID = "partners_free"
	SubscriptionV2RatePlanIDPartnersPro        SubscriptionV2RatePlanID = "partners_pro"
	SubscriptionV2RatePlanIDPartnersBusiness   SubscriptionV2RatePlanID = "partners_business"
	SubscriptionV2RatePlanIDPartnersEnterprise SubscriptionV2RatePlanID = "partners_enterprise"
)

func (r SubscriptionV2RatePlanID) IsKnown() bool {
	switch r {
	case SubscriptionV2RatePlanIDFree, SubscriptionV2RatePlanIDLite, SubscriptionV2RatePlanIDPro, SubscriptionV2RatePlanIDProPlus, SubscriptionV2RatePlanIDBusiness, SubscriptionV2RatePlanIDEnterprise, SubscriptionV2RatePlanIDPartnersFree, SubscriptionV2RatePlanIDPartnersPro, SubscriptionV2RatePlanIDPartnersBusiness, SubscriptionV2RatePlanIDPartnersEnterprise:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionV2State string

const (
	SubscriptionV2StateTrial           SubscriptionV2State = "Trial"
	SubscriptionV2StateProvisioned     SubscriptionV2State = "Provisioned"
	SubscriptionV2StatePaid            SubscriptionV2State = "Paid"
	SubscriptionV2StateAwaitingPayment SubscriptionV2State = "AwaitingPayment"
	SubscriptionV2StateCancelled       SubscriptionV2State = "Cancelled"
	SubscriptionV2StateFailed          SubscriptionV2State = "Failed"
	SubscriptionV2StateExpired         SubscriptionV2State = "Expired"
)

func (r SubscriptionV2State) IsKnown() bool {
	switch r {
	case SubscriptionV2StateTrial, SubscriptionV2StateProvisioned, SubscriptionV2StatePaid, SubscriptionV2StateAwaitingPayment, SubscriptionV2StateCancelled, SubscriptionV2StateFailed, SubscriptionV2StateExpired:
		return true
	}
	return false
}

type AccountSubscriptionListResponse struct {
	Errors   []AccountSubscriptionListResponseError   `json:"errors,required"`
	Messages []AccountSubscriptionListResponseMessage `json:"messages,required"`
	Result   []Subscription                           `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccountSubscriptionListResponseSuccess    `json:"success,required"`
	ResultInfo AccountSubscriptionListResponseResultInfo `json:"result_info"`
	JSON       accountSubscriptionListResponseJSON       `json:"-"`
}

// accountSubscriptionListResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionListResponse]
type accountSubscriptionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionListResponseError struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AccountSubscriptionListResponseErrorsSource `json:"source"`
	JSON             accountSubscriptionListResponseErrorJSON    `json:"-"`
}

// accountSubscriptionListResponseErrorJSON contains the JSON metadata for the
// struct [AccountSubscriptionListResponseError]
type accountSubscriptionListResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSubscriptionListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionListResponseErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    accountSubscriptionListResponseErrorsSourceJSON `json:"-"`
}

// accountSubscriptionListResponseErrorsSourceJSON contains the JSON metadata for
// the struct [AccountSubscriptionListResponseErrorsSource]
type accountSubscriptionListResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionListResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionListResponseMessage struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccountSubscriptionListResponseMessagesSource `json:"source"`
	JSON             accountSubscriptionListResponseMessageJSON    `json:"-"`
}

// accountSubscriptionListResponseMessageJSON contains the JSON metadata for the
// struct [AccountSubscriptionListResponseMessage]
type accountSubscriptionListResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccountSubscriptionListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionListResponseMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accountSubscriptionListResponseMessagesSourceJSON `json:"-"`
}

// accountSubscriptionListResponseMessagesSourceJSON contains the JSON metadata for
// the struct [AccountSubscriptionListResponseMessagesSource]
type accountSubscriptionListResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionListResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSubscriptionListResponseSuccess bool

const (
	AccountSubscriptionListResponseSuccessTrue AccountSubscriptionListResponseSuccess = true
)

func (r AccountSubscriptionListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSubscriptionListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSubscriptionListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       accountSubscriptionListResponseResultInfoJSON `json:"-"`
}

// accountSubscriptionListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountSubscriptionListResponseResultInfo]
type accountSubscriptionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionDeleteResponse struct {
	Errors   []BillSubsAPIMessages                   `json:"errors,required"`
	Messages []BillSubsAPIMessages                   `json:"messages,required"`
	Result   AccountSubscriptionDeleteResponseResult `json:"result,required"`
	// Whether the API call was successful
	Success AccountSubscriptionDeleteResponseSuccess `json:"success,required"`
	JSON    accountSubscriptionDeleteResponseJSON    `json:"-"`
}

// accountSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionDeleteResponse]
type accountSubscriptionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionDeleteResponseResult struct {
	// Subscription identifier tag.
	SubscriptionID string                                      `json:"subscription_id"`
	JSON           accountSubscriptionDeleteResponseResultJSON `json:"-"`
}

// accountSubscriptionDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccountSubscriptionDeleteResponseResult]
type accountSubscriptionDeleteResponseResultJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountSubscriptionDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSubscriptionDeleteResponseSuccess bool

const (
	AccountSubscriptionDeleteResponseSuccessTrue AccountSubscriptionDeleteResponseSuccess = true
)

func (r AccountSubscriptionDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AccountSubscriptionDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AccountSubscriptionNewParams struct {
	SubscriptionV2 SubscriptionV2Param `json:"subscription_v2,required"`
}

func (r AccountSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubscriptionV2)
}

type AccountSubscriptionUpdateParams struct {
	SubscriptionV2 SubscriptionV2Param `json:"subscription_v2,required"`
}

func (r AccountSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubscriptionV2)
}
