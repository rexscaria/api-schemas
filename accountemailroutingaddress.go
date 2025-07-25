// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rexscaria/api-schemas/internal/apijson"
	"github.com/rexscaria/api-schemas/internal/apiquery"
	"github.com/rexscaria/api-schemas/internal/param"
	"github.com/rexscaria/api-schemas/internal/requestconfig"
	"github.com/rexscaria/api-schemas/option"
)

// AccountEmailRoutingAddressService contains methods and other services that help
// with interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountEmailRoutingAddressService] method instead.
type AccountEmailRoutingAddressService struct {
	Options []option.RequestOption
}

// NewAccountEmailRoutingAddressService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountEmailRoutingAddressService(opts ...option.RequestOption) (r *AccountEmailRoutingAddressService) {
	r = &AccountEmailRoutingAddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *AccountEmailRoutingAddressService) New(ctx context.Context, accountID string, body AccountEmailRoutingAddressNewParams, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets information for a specific destination email already created.
func (r *AccountEmailRoutingAddressService) Get(ctx context.Context, accountID string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists existing destination addresses.
func (r *AccountEmailRoutingAddressService) List(ctx context.Context, accountID string, query AccountEmailRoutingAddressListParams, opts ...option.RequestOption) (res *AccountEmailRoutingAddressListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specific destination address.
func (r *AccountEmailRoutingAddressService) Delete(ctx context.Context, accountID string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *DestinationAddressResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DestinationAddressResponseSingle struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success DestinationAddressResponseSingleSuccess `json:"success,required"`
	Result  EmailAddress                            `json:"result"`
	JSON    destinationAddressResponseSingleJSON    `json:"-"`
}

// destinationAddressResponseSingleJSON contains the JSON metadata for the struct
// [DestinationAddressResponseSingle]
type destinationAddressResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationAddressResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationAddressResponseSingleJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DestinationAddressResponseSingleSuccess bool

const (
	DestinationAddressResponseSingleSuccessTrue DestinationAddressResponseSingleSuccess = true
)

func (r DestinationAddressResponseSingleSuccess) IsKnown() bool {
	switch r {
	case DestinationAddressResponseSingleSuccessTrue:
		return true
	}
	return false
}

type EmailAddress struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	//
	// Deprecated: deprecated
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time        `json:"verified" format:"date-time"`
	JSON     emailAddressJSON `json:"-"`
}

// emailAddressJSON contains the JSON metadata for the struct [EmailAddress]
type emailAddressJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailAddress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailAddressJSON) RawJSON() string {
	return r.raw
}

type AccountEmailRoutingAddressListResponse struct {
	Errors   []EmailMessagesItem `json:"errors,required"`
	Messages []EmailMessagesItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountEmailRoutingAddressListResponseSuccess    `json:"success,required"`
	Result     []EmailAddress                                   `json:"result"`
	ResultInfo AccountEmailRoutingAddressListResponseResultInfo `json:"result_info"`
	JSON       accountEmailRoutingAddressListResponseJSON       `json:"-"`
}

// accountEmailRoutingAddressListResponseJSON contains the JSON metadata for the
// struct [AccountEmailRoutingAddressListResponse]
type accountEmailRoutingAddressListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailRoutingAddressListResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountEmailRoutingAddressListResponseSuccess bool

const (
	AccountEmailRoutingAddressListResponseSuccessTrue AccountEmailRoutingAddressListResponseSuccess = true
)

func (r AccountEmailRoutingAddressListResponseSuccess) IsKnown() bool {
	switch r {
	case AccountEmailRoutingAddressListResponseSuccessTrue:
		return true
	}
	return false
}

type AccountEmailRoutingAddressListResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                              `json:"total_count"`
	JSON       accountEmailRoutingAddressListResponseResultInfoJSON `json:"-"`
}

// accountEmailRoutingAddressListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccountEmailRoutingAddressListResponseResultInfo]
type accountEmailRoutingAddressListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountEmailRoutingAddressListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountEmailRoutingAddressListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountEmailRoutingAddressNewParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r AccountEmailRoutingAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountEmailRoutingAddressListParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[AccountEmailRoutingAddressListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[AccountEmailRoutingAddressListParamsVerified] `query:"verified"`
}

// URLQuery serializes [AccountEmailRoutingAddressListParams]'s query parameters as
// `url.Values`.
func (r AccountEmailRoutingAddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type AccountEmailRoutingAddressListParamsDirection string

const (
	AccountEmailRoutingAddressListParamsDirectionAsc  AccountEmailRoutingAddressListParamsDirection = "asc"
	AccountEmailRoutingAddressListParamsDirectionDesc AccountEmailRoutingAddressListParamsDirection = "desc"
)

func (r AccountEmailRoutingAddressListParamsDirection) IsKnown() bool {
	switch r {
	case AccountEmailRoutingAddressListParamsDirectionAsc, AccountEmailRoutingAddressListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter by verified destination addresses.
type AccountEmailRoutingAddressListParamsVerified bool

const (
	AccountEmailRoutingAddressListParamsVerifiedTrue  AccountEmailRoutingAddressListParamsVerified = true
	AccountEmailRoutingAddressListParamsVerifiedFalse AccountEmailRoutingAddressListParamsVerified = false
)

func (r AccountEmailRoutingAddressListParamsVerified) IsKnown() bool {
	switch r {
	case AccountEmailRoutingAddressListParamsVerifiedTrue, AccountEmailRoutingAddressListParamsVerifiedFalse:
		return true
	}
	return false
}
