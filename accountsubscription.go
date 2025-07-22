// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
	"github.com/rexscaria/api-schemas/shared"
	"github.com/tidwall/gjson"
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
func (r *AccountSubscriptionService) Delete(ctx context.Context, accountID string, subscriptionIdentifier string, body AccountSubscriptionDeleteParams, opts ...option.RequestOption) (res *AccountSubscriptionDeleteResponse, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountSubscriptionResponseSingle struct {
	Result interface{}                           `json:"result"`
	JSON   accountSubscriptionResponseSingleJSON `json:"-"`
	APIResponseSingleBilling
}

// accountSubscriptionResponseSingleJSON contains the JSON metadata for the struct
// [AccountSubscriptionResponseSingle]
type accountSubscriptionResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionResponseSingleJSON) RawJSON() string {
	return r.raw
}

type APIResponseBilling struct {
	Errors   []BillSubsAPIMessages         `json:"errors,required"`
	Messages []BillSubsAPIMessages         `json:"messages,required"`
	Result   APIResponseBillingResultUnion `json:"result,required"`
	// Whether the API call was successful
	Success APIResponseBillingSuccess `json:"success,required"`
	JSON    apiResponseBillingJSON    `json:"-"`
}

// apiResponseBillingJSON contains the JSON metadata for the struct
// [APIResponseBilling]
type apiResponseBillingJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseBillingJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [APIResponseBillingResultArray] or [shared.UnionString].
type APIResponseBillingResultUnion interface {
	ImplementsAPIResponseBillingResultUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*APIResponseBillingResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(APIResponseBillingResultArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type APIResponseBillingResultArray []interface{}

func (r APIResponseBillingResultArray) ImplementsAPIResponseBillingResultUnion() {}

// Whether the API call was successful
type APIResponseBillingSuccess bool

const (
	APIResponseBillingSuccessTrue APIResponseBillingSuccess = true
)

func (r APIResponseBillingSuccess) IsKnown() bool {
	switch r {
	case APIResponseBillingSuccessTrue:
		return true
	}
	return false
}

type APIResponseSingleBilling struct {
	Result interface{}                  `json:"result"`
	JSON   apiResponseSingleBillingJSON `json:"-"`
	APIResponseBilling
}

// apiResponseSingleBillingJSON contains the JSON metadata for the struct
// [APIResponseSingleBilling]
type apiResponseSingleBillingJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *APIResponseSingleBilling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiResponseSingleBillingJSON) RawJSON() string {
	return r.raw
}

type BillSubsAPIMessages struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    billSubsAPIMessagesJSON `json:"-"`
}

// billSubsAPIMessagesJSON contains the JSON metadata for the struct
// [BillSubsAPIMessages]
type billSubsAPIMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillSubsAPIMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billSubsAPIMessagesJSON) RawJSON() string {
	return r.raw
}

type Subscription struct {
	// Subscription identifier tag.
	ID  string          `json:"id"`
	App SubscriptionApp `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues []SubscriptionComponentValue `json:"component_values"`
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
	// A simple zone object. May have null properties if not a zone subscription.
	Zone SubscriptionZone `json:"zone"`
	JSON subscriptionJSON `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                 apijson.Field
	App                apijson.Field
	ComponentValues    apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	Zone               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

type SubscriptionApp struct {
	// app install id.
	InstallID string              `json:"install_id"`
	JSON      subscriptionAppJSON `json:"-"`
}

// subscriptionAppJSON contains the JSON metadata for the struct [SubscriptionApp]
type subscriptionAppJSON struct {
	InstallID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionAppJSON) RawJSON() string {
	return r.raw
}

// A component value for a subscription.
type SubscriptionComponentValue struct {
	// The default amount assigned.
	Default float64 `json:"default"`
	// The name of the component value.
	Name string `json:"name"`
	// The unit price for the component value.
	Price float64 `json:"price"`
	// The amount of the component value assigned.
	Value float64                        `json:"value"`
	JSON  subscriptionComponentValueJSON `json:"-"`
}

// subscriptionComponentValueJSON contains the JSON metadata for the struct
// [SubscriptionComponentValue]
type subscriptionComponentValueJSON struct {
	Default     apijson.Field
	Name        apijson.Field
	Price       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionComponentValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionComponentValueJSON) RawJSON() string {
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
	ID string `json:"id"`
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

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionZone struct {
	// Identifier
	ID string `json:"id"`
	// The domain name
	Name string               `json:"name"`
	JSON subscriptionZoneJSON `json:"-"`
}

// subscriptionZoneJSON contains the JSON metadata for the struct
// [SubscriptionZone]
type subscriptionZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionZoneJSON) RawJSON() string {
	return r.raw
}

type SubscriptionV2Param struct {
	App param.Field[SubscriptionV2AppParam] `json:"app"`
	// The list of add-ons subscribed to.
	ComponentValues param.Field[[]SubscriptionV2ComponentValueParam] `json:"component_values"`
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionV2Frequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[SubscriptionV2RatePlanParam] `json:"rate_plan"`
	// A simple zone object. May have null properties if not a zone subscription.
	Zone param.Field[SubscriptionV2ZoneParam] `json:"zone"`
}

func (r SubscriptionV2Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SubscriptionV2AppParam struct {
	// app install id.
	InstallID param.Field[string] `json:"install_id"`
}

func (r SubscriptionV2AppParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A component value for a subscription.
type SubscriptionV2ComponentValueParam struct {
	// The default amount assigned.
	Default param.Field[float64] `json:"default"`
	// The name of the component value.
	Name param.Field[string] `json:"name"`
	// The unit price for the component value.
	Price param.Field[float64] `json:"price"`
	// The amount of the component value assigned.
	Value param.Field[float64] `json:"value"`
}

func (r SubscriptionV2ComponentValueParam) MarshalJSON() (data []byte, err error) {
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
	ID param.Field[string] `json:"id"`
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

// A simple zone object. May have null properties if not a zone subscription.
type SubscriptionV2ZoneParam struct {
}

func (r SubscriptionV2ZoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountSubscriptionListResponse struct {
	Result []Subscription                      `json:"result"`
	JSON   accountSubscriptionListResponseJSON `json:"-"`
	BillSubsAPIResponseCollection
}

// accountSubscriptionListResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionListResponse]
type accountSubscriptionListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSubscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSubscriptionDeleteResponse struct {
	Result AccountSubscriptionDeleteResponseResult `json:"result"`
	JSON   accountSubscriptionDeleteResponseJSON   `json:"-"`
	APIResponseSingleBilling
}

// accountSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [AccountSubscriptionDeleteResponse]
type accountSubscriptionDeleteResponseJSON struct {
	Result      apijson.Field
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

type AccountSubscriptionDeleteParams struct {
	Body interface{} `json:"body,required"`
}

func (r AccountSubscriptionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
