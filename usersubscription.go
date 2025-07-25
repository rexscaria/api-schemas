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

// UserSubscriptionService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSubscriptionService] method instead.
type UserSubscriptionService struct {
	Options []option.RequestOption
}

// NewUserSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSubscriptionService(opts ...option.RequestOption) (r *UserSubscriptionService) {
	r = &UserSubscriptionService{}
	r.Options = opts
	return
}

// Updates a user's subscriptions.
func (r *UserSubscriptionService) Update(ctx context.Context, identifier string, body UserSubscriptionUpdateParams, opts ...option.RequestOption) (res *UserSubscriptionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all of a user's subscriptions.
func (r *UserSubscriptionService) List(ctx context.Context, opts ...option.RequestOption) (res *UserSubscriptionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a user's subscription.
func (r *UserSubscriptionService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserSubscriptionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("user/subscriptions/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type UserSubscriptionUpdateResponse struct {
	Errors   []BillSubsAPIMessages `json:"errors,required"`
	Messages []BillSubsAPIMessages `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Whether the API call was successful
	Success UserSubscriptionUpdateResponseSuccess `json:"success,required"`
	JSON    userSubscriptionUpdateResponseJSON    `json:"-"`
}

// userSubscriptionUpdateResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionUpdateResponse]
type userSubscriptionUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserSubscriptionUpdateResponseSuccess bool

const (
	UserSubscriptionUpdateResponseSuccessTrue UserSubscriptionUpdateResponseSuccess = true
)

func (r UserSubscriptionUpdateResponseSuccess) IsKnown() bool {
	switch r {
	case UserSubscriptionUpdateResponseSuccessTrue:
		return true
	}
	return false
}

type UserSubscriptionListResponse struct {
	Errors   []BillSubsAPIMessages `json:"errors,required"`
	Messages []BillSubsAPIMessages `json:"messages,required"`
	Result   []Subscription        `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserSubscriptionListResponseSuccess    `json:"success,required"`
	ResultInfo UserSubscriptionListResponseResultInfo `json:"result_info"`
	JSON       userSubscriptionListResponseJSON       `json:"-"`
}

// userSubscriptionListResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionListResponse]
type userSubscriptionListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UserSubscriptionListResponseSuccess bool

const (
	UserSubscriptionListResponseSuccessTrue UserSubscriptionListResponseSuccess = true
)

func (r UserSubscriptionListResponseSuccess) IsKnown() bool {
	switch r {
	case UserSubscriptionListResponseSuccessTrue:
		return true
	}
	return false
}

type UserSubscriptionListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                    `json:"total_count"`
	JSON       userSubscriptionListResponseResultInfoJSON `json:"-"`
}

// userSubscriptionListResponseResultInfoJSON contains the JSON metadata for the
// struct [UserSubscriptionListResponseResultInfo]
type userSubscriptionListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSubscriptionListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type UserSubscriptionDeleteResponse struct {
	// Subscription identifier tag.
	SubscriptionID string                             `json:"subscription_id"`
	JSON           userSubscriptionDeleteResponseJSON `json:"-"`
}

// userSubscriptionDeleteResponseJSON contains the JSON metadata for the struct
// [UserSubscriptionDeleteResponse]
type userSubscriptionDeleteResponseJSON struct {
	SubscriptionID apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserSubscriptionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSubscriptionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UserSubscriptionUpdateParams struct {
	SubscriptionV2 SubscriptionV2Param `json:"subscription_v2,required"`
}

func (r UserSubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SubscriptionV2)
}
