// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountAlertingV3Service contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAlertingV3Service] method instead.
type AccountAlertingV3Service struct {
	Options      []option.RequestOption
	Destinations *AccountAlertingV3DestinationService
	Policies     *AccountAlertingV3PolicyService
}

// NewAccountAlertingV3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAlertingV3Service(opts ...option.RequestOption) (r *AccountAlertingV3Service) {
	r = &AccountAlertingV3Service{}
	r.Options = opts
	r.Destinations = NewAccountAlertingV3DestinationService(opts...)
	r.Policies = NewAccountAlertingV3PolicyService(opts...)
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AccountAlertingV3Service) ListAvailableAlerts(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountAlertingV3ListAvailableAlertsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets a list of history records for notifications sent to an account. The records
// are displayed for last `x` number of days based on the zone plan (free = 30, pro
// = 30, biz = 30, ent = 90).
func (r *AccountAlertingV3Service) ListHistory(ctx context.Context, accountID string, query AccountAlertingV3ListHistoryParams, opts ...option.RequestOption) (res *AccountAlertingV3ListHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAlertingV3ListAvailableAlertsResponse struct {
	Result map[string][]AccountAlertingV3ListAvailableAlertsResponseResult `json:"result"`
	JSON   accountAlertingV3ListAvailableAlertsResponseJSON                `json:"-"`
	APIResponseCollectionAlerting
}

// accountAlertingV3ListAvailableAlertsResponseJSON contains the JSON metadata for
// the struct [AccountAlertingV3ListAvailableAlertsResponse]
type accountAlertingV3ListAvailableAlertsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3ListAvailableAlertsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3ListAvailableAlertsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3ListAvailableAlertsResponseResult struct {
	// Describes the alert type.
	Description string `json:"description"`
	// Alert type name.
	DisplayName string `json:"display_name"`
	// Format of additional configuration options (filters) for the alert type. Data
	// type of filters during policy creation: Array of strings.
	FilterOptions []interface{} `json:"filter_options"`
	// Use this value when creating and updating a notification policy.
	Type string                                                 `json:"type"`
	JSON accountAlertingV3ListAvailableAlertsResponseResultJSON `json:"-"`
}

// accountAlertingV3ListAvailableAlertsResponseResultJSON contains the JSON
// metadata for the struct [AccountAlertingV3ListAvailableAlertsResponseResult]
type accountAlertingV3ListAvailableAlertsResponseResultJSON struct {
	Description   apijson.Field
	DisplayName   apijson.Field
	FilterOptions apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAlertingV3ListAvailableAlertsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3ListAvailableAlertsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3ListHistoryResponse struct {
	Result     []AccountAlertingV3ListHistoryResponseResult `json:"result"`
	ResultInfo interface{}                                  `json:"result_info"`
	JSON       accountAlertingV3ListHistoryResponseJSON     `json:"-"`
	APIResponseCollectionAlerting
}

// accountAlertingV3ListHistoryResponseJSON contains the JSON metadata for the
// struct [AccountAlertingV3ListHistoryResponse]
type accountAlertingV3ListHistoryResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3ListHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3ListHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAlertingV3ListHistoryResponseResult struct {
	// UUID
	ID string `json:"id"`
	// Message body included in the notification sent.
	AlertBody string `json:"alert_body"`
	// Type of notification that has been dispatched.
	AlertType string `json:"alert_type"`
	// Description of the notification policy (if present).
	Description string `json:"description"`
	// The mechanism to which the notification has been dispatched.
	Mechanism string `json:"mechanism"`
	// The type of mechanism to which the notification has been dispatched. This can be
	// email/pagerduty/webhook based on the mechanism configured.
	MechanismType AccountAlertingV3ListHistoryResponseResultMechanismType `json:"mechanism_type"`
	// Name of the policy.
	Name string `json:"name"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// Timestamp of when the notification was dispatched in ISO 8601 format.
	Sent time.Time                                      `json:"sent" format:"date-time"`
	JSON accountAlertingV3ListHistoryResponseResultJSON `json:"-"`
}

// accountAlertingV3ListHistoryResponseResultJSON contains the JSON metadata for
// the struct [AccountAlertingV3ListHistoryResponseResult]
type accountAlertingV3ListHistoryResponseResultJSON struct {
	ID            apijson.Field
	AlertBody     apijson.Field
	AlertType     apijson.Field
	Description   apijson.Field
	Mechanism     apijson.Field
	MechanismType apijson.Field
	Name          apijson.Field
	PolicyID      apijson.Field
	Sent          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountAlertingV3ListHistoryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3ListHistoryResponseResultJSON) RawJSON() string {
	return r.raw
}

// The type of mechanism to which the notification has been dispatched. This can be
// email/pagerduty/webhook based on the mechanism configured.
type AccountAlertingV3ListHistoryResponseResultMechanismType string

const (
	AccountAlertingV3ListHistoryResponseResultMechanismTypeEmail     AccountAlertingV3ListHistoryResponseResultMechanismType = "email"
	AccountAlertingV3ListHistoryResponseResultMechanismTypePagerduty AccountAlertingV3ListHistoryResponseResultMechanismType = "pagerduty"
	AccountAlertingV3ListHistoryResponseResultMechanismTypeWebhook   AccountAlertingV3ListHistoryResponseResultMechanismType = "webhook"
)

func (r AccountAlertingV3ListHistoryResponseResultMechanismType) IsKnown() bool {
	switch r {
	case AccountAlertingV3ListHistoryResponseResultMechanismTypeEmail, AccountAlertingV3ListHistoryResponseResultMechanismTypePagerduty, AccountAlertingV3ListHistoryResponseResultMechanismTypeWebhook:
		return true
	}
	return false
}

type AccountAlertingV3ListHistoryParams struct {
	// Limit the returned results to history records older than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limit the returned results to history records newer than the specified date.
	// This must be a timestamp that conforms to RFC3339.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
}

// URLQuery serializes [AccountAlertingV3ListHistoryParams]'s query parameters as
// `url.Values`.
func (r AccountAlertingV3ListHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
