// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/apiquery"
	"github.com/stainless-sdks/cf-rex-go/internal/param"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// UserBillingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserBillingService] method instead.
type UserBillingService struct {
	Options []option.RequestOption
}

// NewUserBillingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserBillingService(opts ...option.RequestOption) (r *UserBillingService) {
	r = &UserBillingService{}
	r.Options = opts
	return
}

// Accesses your billing history object.
func (r *UserBillingService) ListHistory(ctx context.Context, query UserBillingListHistoryParams, opts ...option.RequestOption) (res *UserBillingListHistoryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/billing/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Accesses your billing profile object.
func (r *UserBillingService) GetProfile(ctx context.Context, opts ...option.RequestOption) (res *BillingResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserBillingListHistoryResponse struct {
	Result []UserBillingListHistoryResponseResult `json:"result"`
	JSON   userBillingListHistoryResponseJSON     `json:"-"`
	BillSubsAPIResponseCollection
}

// userBillingListHistoryResponseJSON contains the JSON metadata for the struct
// [UserBillingListHistoryResponse]
type userBillingListHistoryResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingListHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingListHistoryResponseJSON) RawJSON() string {
	return r.raw
}

type UserBillingListHistoryResponseResult struct {
	// Billing item identifier tag.
	ID string `json:"id,required"`
	// The billing item action.
	Action string `json:"action,required"`
	// The amount associated with this billing item.
	Amount float64 `json:"amount,required"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency,required"`
	// The billing item description.
	Description string `json:"description,required"`
	// When the billing item was created.
	OccurredAt time.Time `json:"occurred_at,required" format:"date-time"`
	// The billing item type.
	Type string                                   `json:"type,required"`
	Zone UserBillingListHistoryResponseResultZone `json:"zone,required"`
	JSON userBillingListHistoryResponseResultJSON `json:"-"`
}

// userBillingListHistoryResponseResultJSON contains the JSON metadata for the
// struct [UserBillingListHistoryResponseResult]
type userBillingListHistoryResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Amount      apijson.Field
	Currency    apijson.Field
	Description apijson.Field
	OccurredAt  apijson.Field
	Type        apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingListHistoryResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingListHistoryResponseResultJSON) RawJSON() string {
	return r.raw
}

type UserBillingListHistoryResponseResultZone struct {
	Name string                                       `json:"name"`
	JSON userBillingListHistoryResponseResultZoneJSON `json:"-"`
}

// userBillingListHistoryResponseResultZoneJSON contains the JSON metadata for the
// struct [UserBillingListHistoryResponseResultZone]
type userBillingListHistoryResponseResultZoneJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserBillingListHistoryResponseResultZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userBillingListHistoryResponseResultZoneJSON) RawJSON() string {
	return r.raw
}

type UserBillingListHistoryParams struct {
	// The billing item action.
	Action param.Field[string] `query:"action"`
	// When the billing item was created.
	OccurredAt param.Field[time.Time] `query:"occurred_at" format:"date-time"`
	// Field to order billing history by.
	Order param.Field[UserBillingListHistoryParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The billing item type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserBillingListHistoryParams]'s query parameters as
// `url.Values`.
func (r UserBillingListHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Field to order billing history by.
type UserBillingListHistoryParamsOrder string

const (
	UserBillingListHistoryParamsOrderType       UserBillingListHistoryParamsOrder = "type"
	UserBillingListHistoryParamsOrderOccurredAt UserBillingListHistoryParamsOrder = "occurred_at"
	UserBillingListHistoryParamsOrderAction     UserBillingListHistoryParamsOrder = "action"
)

func (r UserBillingListHistoryParamsOrder) IsKnown() bool {
	switch r {
	case UserBillingListHistoryParamsOrderType, UserBillingListHistoryParamsOrderOccurredAt, UserBillingListHistoryParamsOrderAction:
		return true
	}
	return false
}
