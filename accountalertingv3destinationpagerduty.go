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
	Errors   []AaaMessage `json:"errors,required"`
	Messages []AaaMessage `json:"messages,required"`
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
	ResultInfo APIResponseCollectionAlertingResultInfo `json:"result_info"`
	JSON       apiResponseCollectionAlertingJSON       `json:"-"`
	APIResponseAlerting
}

// apiResponseCollectionAlertingJSON contains the JSON metadata for the struct
// [APIResponseCollectionAlerting]
type apiResponseCollectionAlertingJSON struct {
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
	Result []AccountAlertingV3DestinationPagerdutyListResponseResult `json:"result"`
	JSON   accountAlertingV3DestinationPagerdutyListResponseJSON     `json:"-"`
	APIResponseCollectionAlerting
}

// accountAlertingV3DestinationPagerdutyListResponseJSON contains the JSON metadata
// for the struct [AccountAlertingV3DestinationPagerdutyListResponse]
type accountAlertingV3DestinationPagerdutyListResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAlertingV3DestinationPagerdutyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAlertingV3DestinationPagerdutyListResponseJSON) RawJSON() string {
	return r.raw
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
