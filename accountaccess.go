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

// AccountAccessService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessService] method instead.
type AccountAccessService struct {
	Options           []option.RequestOption
	Apps              *AccountAccessAppService
	Bookmarks         *AccountAccessBookmarkService
	Certificates      *AccountAccessCertificateService
	CustomPages       *AccountAccessCustomPageService
	GatewayCa         *AccountAccessGatewayCaService
	Groups            *AccountAccessGroupService
	IdentityProviders *AccountAccessIdentityProviderService
	Keys              *AccountAccessKeyService
	Logs              *AccountAccessLogService
	Organizations     *AccountAccessOrganizationService
	Policies          *AccountAccessPolicyService
	PolicyTests       *AccountAccessPolicyTestService
	ServiceTokens     *AccountAccessServiceTokenService
	Tags              *AccountAccessTagService
	Users             *AccountAccessUserService
}

// NewAccountAccessService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAccessService(opts ...option.RequestOption) (r *AccountAccessService) {
	r = &AccountAccessService{}
	r.Options = opts
	r.Apps = NewAccountAccessAppService(opts...)
	r.Bookmarks = NewAccountAccessBookmarkService(opts...)
	r.Certificates = NewAccountAccessCertificateService(opts...)
	r.CustomPages = NewAccountAccessCustomPageService(opts...)
	r.GatewayCa = NewAccountAccessGatewayCaService(opts...)
	r.Groups = NewAccountAccessGroupService(opts...)
	r.IdentityProviders = NewAccountAccessIdentityProviderService(opts...)
	r.Keys = NewAccountAccessKeyService(opts...)
	r.Logs = NewAccountAccessLogService(opts...)
	r.Organizations = NewAccountAccessOrganizationService(opts...)
	r.Policies = NewAccountAccessPolicyService(opts...)
	r.PolicyTests = NewAccountAccessPolicyTestService(opts...)
	r.ServiceTokens = NewAccountAccessServiceTokenService(opts...)
	r.Tags = NewAccountAccessTagService(opts...)
	r.Users = NewAccountAccessUserService(opts...)
	return
}

// Removes a user from a Zero Trust seat when both `access_seat` and `gateway_seat`
// are set to false.
func (r *AccountAccessService) UpdateSeats(ctx context.Context, accountID string, body AccountAccessUpdateSeatsParams, opts ...option.RequestOption) (res *AccountAccessUpdateSeatsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/seats", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type AccountAccessUpdateSeatsResponse struct {
	Errors   []MessagesAccessItem `json:"errors,required"`
	Messages []MessagesAccessItem `json:"messages,required"`
	// Whether the API call was successful.
	Success    AccountAccessUpdateSeatsResponseSuccess    `json:"success,required"`
	Result     []AccountAccessUpdateSeatsResponseResult   `json:"result"`
	ResultInfo AccountAccessUpdateSeatsResponseResultInfo `json:"result_info"`
	JSON       accountAccessUpdateSeatsResponseJSON       `json:"-"`
}

// accountAccessUpdateSeatsResponseJSON contains the JSON metadata for the struct
// [AccountAccessUpdateSeatsResponse]
type accountAccessUpdateSeatsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUpdateSeatsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUpdateSeatsResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccountAccessUpdateSeatsResponseSuccess bool

const (
	AccountAccessUpdateSeatsResponseSuccessTrue AccountAccessUpdateSeatsResponseSuccess = true
)

func (r AccountAccessUpdateSeatsResponseSuccess) IsKnown() bool {
	switch r {
	case AccountAccessUpdateSeatsResponseSuccessTrue:
		return true
	}
	return false
}

type AccountAccessUpdateSeatsResponseResult struct {
	// True if the seat is part of Access.
	AccessSeat bool      `json:"access_seat"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// True if the seat is part of Gateway.
	GatewaySeat bool `json:"gateway_seat"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid   string                                     `json:"seat_uid"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountAccessUpdateSeatsResponseResultJSON `json:"-"`
}

// accountAccessUpdateSeatsResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessUpdateSeatsResponseResult]
type accountAccessUpdateSeatsResponseResultJSON struct {
	AccessSeat  apijson.Field
	CreatedAt   apijson.Field
	GatewaySeat apijson.Field
	SeatUid     apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUpdateSeatsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUpdateSeatsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUpdateSeatsResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                        `json:"total_count"`
	JSON       accountAccessUpdateSeatsResponseResultInfoJSON `json:"-"`
}

// accountAccessUpdateSeatsResponseResultInfoJSON contains the JSON metadata for
// the struct [AccountAccessUpdateSeatsResponseResultInfo]
type accountAccessUpdateSeatsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUpdateSeatsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUpdateSeatsResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUpdateSeatsParams struct {
	Body []AccountAccessUpdateSeatsParamsBody `json:"body,required"`
}

func (r AccountAccessUpdateSeatsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountAccessUpdateSeatsParamsBody struct {
	// True if the seat is part of Access.
	AccessSeat param.Field[bool] `json:"access_seat,required"`
	// True if the seat is part of Gateway.
	GatewaySeat param.Field[bool] `json:"gateway_seat,required"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid param.Field[string] `json:"seat_uid,required"`
}

func (r AccountAccessUpdateSeatsParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
