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

// AccountAccessUserService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessUserService] method instead.
type AccountAccessUserService struct {
	Options        []option.RequestOption
	ActiveSessions *AccountAccessUserActiveSessionService
}

// NewAccountAccessUserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessUserService(opts ...option.RequestOption) (r *AccountAccessUserService) {
	r = &AccountAccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccountAccessUserActiveSessionService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccountAccessUserService) List(ctx context.Context, accountID string, query AccountAccessUserListParams, opts ...option.RequestOption) (res *AccountAccessUserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get all failed login attempts for a single user.
func (r *AccountAccessUserService) FailedLogins(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *AccountAccessUserFailedLoginsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/failed_logins", accountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get last seen identity for a single user.
func (r *AccountAccessUserService) LastSeenIdentity(ctx context.Context, accountID string, userID string, opts ...option.RequestOption) (res *AccountAccessUserLastSeenIdentityResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", accountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessUserListResponse struct {
	Result     []AccountAccessUserListResponseResult   `json:"result"`
	ResultInfo AccountAccessUserListResponseResultInfo `json:"result_info"`
	JSON       accountAccessUserListResponseJSON       `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessUserListResponseJSON contains the JSON metadata for the struct
// [AccountAccessUserListResponse]
type accountAccessUserListResponseJSON struct {
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserListResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserListResponseResult struct {
	// UUID
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUid string `json:"seat_uid"`
	// The unique API identifier for the user.
	Uid       string                                  `json:"uid"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accountAccessUserListResponseResultJSON `json:"-"`
}

// accountAccessUserListResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessUserListResponseResult]
type accountAccessUserListResponseResultJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUid             apijson.Field
	Uid                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccountAccessUserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserListResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserListResponseResultInfo struct {
	Count      interface{}                                 `json:"count"`
	Page       interface{}                                 `json:"page"`
	PerPage    interface{}                                 `json:"per_page"`
	TotalCount interface{}                                 `json:"total_count"`
	JSON       accountAccessUserListResponseResultInfoJSON `json:"-"`
}

// accountAccessUserListResponseResultInfoJSON contains the JSON metadata for the
// struct [AccountAccessUserListResponseResultInfo]
type accountAccessUserListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserListResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserFailedLoginsResponse struct {
	Result []AccountAccessUserFailedLoginsResponseResult `json:"result"`
	JSON   accountAccessUserFailedLoginsResponseJSON     `json:"-"`
	APIResponseCollectionAccess
}

// accountAccessUserFailedLoginsResponseJSON contains the JSON metadata for the
// struct [AccountAccessUserFailedLoginsResponse]
type accountAccessUserFailedLoginsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserFailedLoginsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserFailedLoginsResponseResult struct {
	Expiration int64                                           `json:"expiration"`
	Metadata   interface{}                                     `json:"metadata"`
	JSON       accountAccessUserFailedLoginsResponseResultJSON `json:"-"`
}

// accountAccessUserFailedLoginsResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessUserFailedLoginsResponseResult]
type accountAccessUserFailedLoginsResponseResultJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserFailedLoginsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserFailedLoginsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserLastSeenIdentityResponse struct {
	Result Identity                                      `json:"result"`
	JSON   accountAccessUserLastSeenIdentityResponseJSON `json:"-"`
	APIResponseSingleAccess
}

// accountAccessUserLastSeenIdentityResponseJSON contains the JSON metadata for the
// struct [AccountAccessUserLastSeenIdentityResponse]
type accountAccessUserLastSeenIdentityResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessUserLastSeenIdentityResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessUserLastSeenIdentityResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessUserListParams struct {
	// The email of the user.
	Email param.Field[string] `query:"email"`
	// The name of the user.
	Name param.Field[string] `query:"name"`
	// Search for users by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccountAccessUserListParams]'s query parameters as
// `url.Values`.
func (r AccountAccessUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
