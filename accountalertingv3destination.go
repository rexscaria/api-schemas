// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAlertingV3DestinationService contains methods and other services that
// help with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3DestinationService] method instead.
type AccountAlertingV3DestinationService struct {
	Options   []option.RequestOption
	Pagerduty *AccountAlertingV3DestinationPagerdutyService
	Webhooks  *AccountAlertingV3DestinationWebhookService
}

// NewAccountAlertingV3DestinationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAlertingV3DestinationService(opts ...option.RequestOption) (r *AccountAlertingV3DestinationService) {
	r = &AccountAlertingV3DestinationService{}
	r.Options = opts
	r.Pagerduty = NewAccountAlertingV3DestinationPagerdutyService(opts...)
	r.Webhooks = NewAccountAlertingV3DestinationWebhookService(opts...)
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *AccountAlertingV3DestinationService) ListEligibility(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3DestinationListEligibilityResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAlertingV3DestinationListEligibilityResponse struct {
	Result map[string][]AccountAlertingV3DestinationListEligibilityResponseResult `json:"result"`
	JSON   accountAlertingV3DestinationListEligibilityResponseJSON                `json:"-"`
	APIResponseCollectionAlerting
}

// accountAlertingV3DestinationListEligibilityResponseJSON contains the JSON
// metadata for the struct [AccountAlertingV3DestinationListEligibilityResponse]
type accountAlertingV3DestinationListEligibilityResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationListEligibilityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationListEligibilityResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3DestinationListEligibilityResponseResult struct {
	// Determines whether or not the account is eligible for the delivery mechanism.
	Eligible bool `json:"eligible"`
	// Beta flag. Users can create a policy with a mechanism that is not ready, but we
	// cannot guarantee successful delivery of notifications.
	Ready bool `json:"ready"`
	// Determines type of delivery mechanism.
	Type AccountAlertingV3DestinationListEligibilityResponseResultType `json:"type"`
	JSON accountAlertingV3DestinationListEligibilityResponseResultJSON `json:"-"`
}

// accountAlertingV3DestinationListEligibilityResponseResultJSON contains the JSON
// metadata for the struct
// [AccountAlertingV3DestinationListEligibilityResponseResult]
type accountAlertingV3DestinationListEligibilityResponseResultJSON struct {
	Eligible    apijson.Field
	Ready       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationListEligibilityResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationListEligibilityResponseResultJSON) RawJSON() string {
	return r.raw
}

// Determines type of delivery mechanism.
type AccountAlertingV3DestinationListEligibilityResponseResultType string

const (
	AccountAlertingV3DestinationListEligibilityResponseResultTypeEmail     AccountAlertingV3DestinationListEligibilityResponseResultType = "email"
	AccountAlertingV3DestinationListEligibilityResponseResultTypePagerduty AccountAlertingV3DestinationListEligibilityResponseResultType = "pagerduty"
	AccountAlertingV3DestinationListEligibilityResponseResultTypeWebhook   AccountAlertingV3DestinationListEligibilityResponseResultType = "webhook"
)

func (r AccountAlertingV3DestinationListEligibilityResponseResultType) IsKnown() bool {
	switch r {
	case AccountAlertingV3DestinationListEligibilityResponseResultTypeEmail, AccountAlertingV3DestinationListEligibilityResponseResultTypePagerduty, AccountAlertingV3DestinationListEligibilityResponseResultTypeWebhook:
		return true
	}
	return false
}
