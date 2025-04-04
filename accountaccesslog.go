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

// AccountAccessLogService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountAccessLogService] method instead.
type AccountAccessLogService struct {
	Options []option.RequestOption
	Scim    *AccountAccessLogScimService
}

// NewAccountAccessLogService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAccessLogService(opts ...option.RequestOption) (r *AccountAccessLogService) {
	r = &AccountAccessLogService{}
	r.Options = opts
	r.Scim = NewAccountAccessLogScimService(opts...)
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccountAccessLogService) AccessRequests(ctx context.Context, accountID string, query AccountAccessLogAccessRequestsParams, opts ...option.RequestOption) (res *AccountAccessLogAccessRequestsResponse, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountAccessLogAccessRequestsResponse struct {
	Result []AccountAccessLogAccessRequestsResponseResult `json:"result"`
	JSON   accountAccessLogAccessRequestsResponseJSON     `json:"-"`
	APIResponseAccess
}

// accountAccessLogAccessRequestsResponseJSON contains the JSON metadata for the
// struct [AccountAccessLogAccessRequestsResponse]
type accountAccessLogAccessRequestsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessLogAccessRequestsResponseJSON) RawJSON() string {
	return r.raw
}

type AccountAccessLogAccessRequestsResponseResult struct {
	// The event that occurred, such as a login attempt.
	Action string `json:"action"`
	// The result of the authentication event.
	Allowed bool `json:"allowed"`
	// The URL of the Access application.
	AppDomain string `json:"app_domain"`
	// The unique identifier for the Access application.
	AppUid string `json:"app_uid"`
	// The IdP used to authenticate.
	Connection string    `json:"connection"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// The IP address of the authenticating user.
	IPAddress string `json:"ip_address"`
	// The unique identifier for the request to Cloudflare.
	RayID string `json:"ray_id"`
	// The email address of the authenticating user.
	UserEmail string                                           `json:"user_email" format:"email"`
	JSON      accountAccessLogAccessRequestsResponseResultJSON `json:"-"`
}

// accountAccessLogAccessRequestsResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessLogAccessRequestsResponseResult]
type accountAccessLogAccessRequestsResponseResultJSON struct {
	Action      apijson.Field
	Allowed     apijson.Field
	AppDomain   apijson.Field
	AppUid      apijson.Field
	Connection  apijson.Field
	CreatedAt   apijson.Field
	IPAddress   apijson.Field
	RayID       apijson.Field
	UserEmail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessLogAccessRequestsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountAccessLogAccessRequestsResponseResultJSON) RawJSON() string {
	return r.raw
}

type AccountAccessLogAccessRequestsParams struct {
	// The chronological sorting order for the logs.
	Direction param.Field[AccountAccessLogAccessRequestsParamsDirection] `query:"direction"`
	// The maximum number of log entries to retrieve.
	Limit param.Field[int64] `query:"limit"`
	// The earliest event timestamp to query.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The latest event timestamp to query.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountAccessLogAccessRequestsParams]'s query parameters as
// `url.Values`.
func (r AccountAccessLogAccessRequestsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The chronological sorting order for the logs.
type AccountAccessLogAccessRequestsParamsDirection string

const (
	AccountAccessLogAccessRequestsParamsDirectionDesc AccountAccessLogAccessRequestsParamsDirection = "desc"
	AccountAccessLogAccessRequestsParamsDirectionAsc  AccountAccessLogAccessRequestsParamsDirection = "asc"
)

func (r AccountAccessLogAccessRequestsParamsDirection) IsKnown() bool {
	switch r {
	case AccountAccessLogAccessRequestsParamsDirectionDesc, AccountAccessLogAccessRequestsParamsDirectionAsc:
		return true
	}
	return false
}
