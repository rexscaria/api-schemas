// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountZtRiskScoringService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountZtRiskScoringService] method instead.
type AccountZtRiskScoringService struct {
	Options      []option.RequestOption
	Behaviors    *AccountZtRiskScoringBehaviorService
	Integrations *AccountZtRiskScoringIntegrationService
}

// NewAccountZtRiskScoringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountZtRiskScoringService(opts ...option.RequestOption) (r *AccountZtRiskScoringService) {
	r = &AccountZtRiskScoringService{}
	r.Options = opts
	r.Behaviors = NewAccountZtRiskScoringBehaviorService(opts...)
	r.Integrations = NewAccountZtRiskScoringIntegrationService(opts...)
	return
}

// Clear the risk score for a particular user
func (r *AccountZtRiskScoringService) ResetRiskScore(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *AccountZtRiskScoringResetRiskScoreResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s/reset", accountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get risk event/score information for a specific user
func (r *AccountZtRiskScoringService) GetRiskScore(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *AccountZtRiskScoringGetRiskScoreResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/%s", accountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get risk score info for all users in the account
func (r *AccountZtRiskScoringService) GetSummary(ctx context.Context, accountID string, opts ...option.RequestOption) (res *AccountZtRiskScoringGetSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zt_risk_scoring/summary", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RiskLevel string

const (
	RiskLevelLow    RiskLevel = "low"
	RiskLevelMedium RiskLevel = "medium"
	RiskLevelHigh   RiskLevel = "high"
)

func (r RiskLevel) IsKnown() bool {
	switch r {
	case RiskLevelLow, RiskLevelMedium, RiskLevelHigh:
		return true
	}
	return false
}

type AccountZtRiskScoringResetRiskScoreResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success AccountZtRiskScoringResetRiskScoreResponseSuccess `json:"success,required"`
	Result  interface{}                                       `json:"result,nullable"`
	JSON    accountZtRiskScoringResetRiskScoreResponseJSON    `json:"-"`
}

// accountZtRiskScoringResetRiskScoreResponseJSON contains the JSON metadata for
// the struct [AccountZtRiskScoringResetRiskScoreResponse]
type accountZtRiskScoringResetRiskScoreResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringResetRiskScoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringResetRiskScoreResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountZtRiskScoringResetRiskScoreResponseSuccess bool

const (
	AccountZtRiskScoringResetRiskScoreResponseSuccessTrue AccountZtRiskScoringResetRiskScoreResponseSuccess = true
)

func (r AccountZtRiskScoringResetRiskScoreResponseSuccess) IsKnown() bool {
	switch r {
	case AccountZtRiskScoringResetRiskScoreResponseSuccessTrue:
		return true
	}
	return false
}

type AccountZtRiskScoringGetRiskScoreResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountZtRiskScoringGetRiskScoreResponseSuccess    `json:"success,required"`
	Result     AccountZtRiskScoringGetRiskScoreResponseResult     `json:"result"`
	ResultInfo AccountZtRiskScoringGetRiskScoreResponseResultInfo `json:"result_info"`
	JSON       accountZtRiskScoringGetRiskScoreResponseJSON       `json:"-"`
}

// accountZtRiskScoringGetRiskScoreResponseJSON contains the JSON metadata for the
// struct [AccountZtRiskScoringGetRiskScoreResponse]
type accountZtRiskScoringGetRiskScoreResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetRiskScoreResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetRiskScoreResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountZtRiskScoringGetRiskScoreResponseSuccess bool

const (
	AccountZtRiskScoringGetRiskScoreResponseSuccessTrue AccountZtRiskScoringGetRiskScoreResponseSuccess = true
)

func (r AccountZtRiskScoringGetRiskScoreResponseSuccess) IsKnown() bool {
	switch r {
	case AccountZtRiskScoringGetRiskScoreResponseSuccessTrue:
		return true
	}
	return false
}

type AccountZtRiskScoringGetRiskScoreResponseResult struct {
	Email         string                                                `json:"email,required"`
	Events        []AccountZtRiskScoringGetRiskScoreResponseResultEvent `json:"events,required"`
	Name          string                                                `json:"name,required"`
	LastResetTime time.Time                                             `json:"last_reset_time,nullable" format:"date-time"`
	RiskLevel     RiskLevel                                             `json:"risk_level"`
	JSON          accountZtRiskScoringGetRiskScoreResponseResultJSON    `json:"-"`
}

// accountZtRiskScoringGetRiskScoreResponseResultJSON contains the JSON metadata
// for the struct [AccountZtRiskScoringGetRiskScoreResponseResult]
type accountZtRiskScoringGetRiskScoreResponseResultJSON struct {
	Email         apijson.Field
	Events        apijson.Field
	Name          apijson.Field
	LastResetTime apijson.Field
	RiskLevel     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetRiskScoreResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetRiskScoreResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringGetRiskScoreResponseResultEvent struct {
	ID           string                                                  `json:"id,required"`
	Name         string                                                  `json:"name,required"`
	RiskLevel    RiskLevel                                               `json:"risk_level,required"`
	Timestamp    time.Time                                               `json:"timestamp,required" format:"date-time"`
	EventDetails interface{}                                             `json:"event_details"`
	JSON         accountZtRiskScoringGetRiskScoreResponseResultEventJSON `json:"-"`
}

// accountZtRiskScoringGetRiskScoreResponseResultEventJSON contains the JSON
// metadata for the struct [AccountZtRiskScoringGetRiskScoreResponseResultEvent]
type accountZtRiskScoringGetRiskScoreResponseResultEventJSON struct {
	ID           apijson.Field
	Name         apijson.Field
	RiskLevel    apijson.Field
	Timestamp    apijson.Field
	EventDetails apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetRiskScoreResponseResultEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetRiskScoreResponseResultEventJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringGetRiskScoreResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                                `json:"total_count"`
	JSON       accountZtRiskScoringGetRiskScoreResponseResultInfoJSON `json:"-"`
}

// accountZtRiskScoringGetRiskScoreResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountZtRiskScoringGetRiskScoreResponseResultInfo]
type accountZtRiskScoringGetRiskScoreResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetRiskScoreResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetRiskScoreResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringGetSummaryResponse struct {
	Errors   []MessagesDlpItems `json:"errors,required"`
	Messages []MessagesDlpItems `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountZtRiskScoringGetSummaryResponseSuccess    `json:"success,required"`
	Result     AccountZtRiskScoringGetSummaryResponseResult     `json:"result"`
	ResultInfo AccountZtRiskScoringGetSummaryResponseResultInfo `json:"result_info"`
	JSON       accountZtRiskScoringGetSummaryResponseJSON       `json:"-"`
}

// accountZtRiskScoringGetSummaryResponseJSON contains the JSON metadata for the
// struct [AccountZtRiskScoringGetSummaryResponse]
type accountZtRiskScoringGetSummaryResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountZtRiskScoringGetSummaryResponseSuccess bool

const (
	AccountZtRiskScoringGetSummaryResponseSuccessTrue AccountZtRiskScoringGetSummaryResponseSuccess = true
)

func (r AccountZtRiskScoringGetSummaryResponseSuccess) IsKnown() bool {
	switch r {
	case AccountZtRiskScoringGetSummaryResponseSuccessTrue:
		return true
	}
	return false
}

type AccountZtRiskScoringGetSummaryResponseResult struct {
	Users []AccountZtRiskScoringGetSummaryResponseResultUser `json:"users,required"`
	JSON  accountZtRiskScoringGetSummaryResponseResultJSON   `json:"-"`
}

// accountZtRiskScoringGetSummaryResponseResultJSON contains the JSON metadata for
// the struct [AccountZtRiskScoringGetSummaryResponseResult]
type accountZtRiskScoringGetSummaryResponseResultJSON struct {
	Users       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetSummaryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetSummaryResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringGetSummaryResponseResultUser struct {
	Email        string                                               `json:"email,required"`
	EventCount   int64                                                `json:"event_count,required"`
	LastEvent    time.Time                                            `json:"last_event,required" format:"date-time"`
	MaxRiskLevel RiskLevel                                            `json:"max_risk_level,required"`
	Name         string                                               `json:"name,required"`
	UserID       string                                               `json:"user_id,required" format:"uuid"`
	JSON         accountZtRiskScoringGetSummaryResponseResultUserJSON `json:"-"`
}

// accountZtRiskScoringGetSummaryResponseResultUserJSON contains the JSON metadata
// for the struct [AccountZtRiskScoringGetSummaryResponseResultUser]
type accountZtRiskScoringGetSummaryResponseResultUserJSON struct {
	Email        apijson.Field
	EventCount   apijson.Field
	LastEvent    apijson.Field
	MaxRiskLevel apijson.Field
	Name         apijson.Field
	UserID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetSummaryResponseResultUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetSummaryResponseResultUserJSON) RawJSON() string {
	return r.raw
}

type AccountZtRiskScoringGetSummaryResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                              `json:"total_count"`
	JSON       accountZtRiskScoringGetSummaryResponseResultInfoJSON `json:"-"`
}

// accountZtRiskScoringGetSummaryResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountZtRiskScoringGetSummaryResponseResultInfo]
type accountZtRiskScoringGetSummaryResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountZtRiskScoringGetSummaryResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountZtRiskScoringGetSummaryResponseResultInfoJSON) RawJSON() string {
	return r.raw
}
